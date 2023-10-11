/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org).
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package xds

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	envoy_cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"

	// envoy_resource "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/subscription"
	wso2_cache "github.com/wso2/apk/adapter/pkg/discovery/protocol/cache/v3"
	wso2_resource "github.com/wso2/apk/adapter/pkg/discovery/protocol/resource/v3"
	eventhubTypes "github.com/wso2/apk/adapter/pkg/eventhub/types"
	"github.com/wso2/apk/adapter/pkg/logging"
	// "github.com/wso2/apk/common-controller/internal/config"
	"github.com/wso2/apk/common-controller/internal/loggers"
	// "github.com/wso2/apk/common-controller/internal/oasparser/model"
	dpv1alpha1 "github.com/wso2/apk/common-controller/internal/operator/apis/dp/v1alpha1"
)

// EnvoyInternalAPI struct use to hold envoy resources and adapter internal resources
type EnvoyInternalAPI struct {
	// commonControllerInternalAPI model.commonControllerInternalAPI
	envoyLabels                 []string
	routes                      []*routev3.Route
	clusters                    []*clusterv3.Cluster
	endpointAddresses           []*corev3.Address
	enforcerAPI                 types.Resource
}

// EnvoyGatewayConfig struct use to hold envoy gateway resources
type EnvoyGatewayConfig struct {
	listener                *listenerv3.Listener
	routeConfig             *routev3.RouteConfiguration
	clusters                []*clusterv3.Cluster
	endpoints               []*corev3.Address
	// customRateLimitPolicies []*model.CustomRateLimitPolicy
}

// EnforcerInternalAPI struct use to hold enforcer resources
type EnforcerInternalAPI struct {
	configs                []types.Resource
	subscriptions          []types.Resource
	applications           []types.Resource
	applicationKeyMappings []types.Resource
	applicationMappings    []types.Resource
}

var (
	// TODO: (VirajSalaka) Remove Unused mutexes.
	mutexForXdsUpdate         sync.Mutex
	mutexForInternalMapUpdate sync.Mutex

	cache                              envoy_cachev3.SnapshotCache
	enforcerCache                      wso2_cache.SnapshotCache
	enforcerSubscriptionCache          wso2_cache.SnapshotCache
	enforcerApplicationCache           wso2_cache.SnapshotCache
	enforcerApplicationKeyMappingCache wso2_cache.SnapshotCache
	enforcerApplicationMappingCache    wso2_cache.SnapshotCache

	orgAPIMap map[string]map[string]*EnvoyInternalAPI // organizationID -> Vhost:API_UUID -> EnvoyInternalAPI struct map

	orgIDvHostBasepathMap map[string]map[string]string   // organizationID -> Vhost:basepath -> Vhost:API_UUID
	orgIDAPIvHostsMap     map[string]map[string][]string // organizationID -> UUID -> prod/sand -> Envoy Vhost Array map

	// Envoy Label as map key
	gatewayLabelConfigMap map[string]*EnvoyGatewayConfig // GW-Label -> EnvoyGatewayConfig struct map

	// Listener as map key
	listenerToRouteArrayMap map[string][]*routev3.Route // Listener -> Routes map

	// Common Enforcer Label as map key
	// TODO(amali) This doesn't have a usage yet. It will be used to handle multiple enforcer labels in future.
	enforcerLabelMap map[string]*EnforcerInternalAPI // Enforcer Label -> EnforcerInternalAPI struct map

	// KeyManagerList to store data
	KeyManagerList = make([]eventhubTypes.KeyManager, 0)
	isReady        = false
)

const (
	maxRandomInt             int    = 999999999
	grpcMaxConcurrentStreams        = 1000000
	apiKeyFieldSeparator     string = ":"
	commonEnforcerLabel      string = "commonEnforcerLabel"
)

func maxRandomBigInt() *big.Int {
	return big.NewInt(int64(maxRandomInt))
}

// IDHash uses ID field as the node hash.
type IDHash struct{}

// ID uses the node ID field
func (IDHash) ID(node *corev3.Node) string {
	if node == nil {
		return "unknown"
	}
	return node.Id
}

var _ envoy_cachev3.NodeHash = IDHash{}

func init() {
	cache = envoy_cachev3.NewSnapshotCache(false, IDHash{}, nil)
	enforcerCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	enforcerSubscriptionCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	enforcerApplicationCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	enforcerApplicationKeyMappingCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	gatewayLabelConfigMap = make(map[string]*EnvoyGatewayConfig)
	listenerToRouteArrayMap = make(map[string][]*routev3.Route)
	orgAPIMap = make(map[string]map[string]*EnvoyInternalAPI)
	orgIDAPIvHostsMap = make(map[string]map[string][]string) // organizationID -> UUID-prod/sand -> Envoy Vhost Array map
	orgIDvHostBasepathMap = make(map[string]map[string]string)

	enforcerLabelMap = make(map[string]*EnforcerInternalAPI)
	//TODO(amali) currently subscriptions, configs, applications, applicationPolicies, subscriptionPolicies,
	// applicationKeyMappings, keyManagerConfigList, revokedTokens are supported with the hard coded label for Enforcer
	enforcerLabelMap[commonEnforcerLabel] = &EnforcerInternalAPI{}
	rand.Seed(time.Now().UnixNano())
	// go watchEnforcerResponse()
}

// GetRateLimiterCache returns xds server cache for rate limiter service.
func GetRateLimiterCache() envoy_cachev3.SnapshotCache {
	return rlsPolicyCache.xdsCache
}

// UpdateRateLimitXDSCache updates the xDS cache of the RateLimiter.
func UpdateRateLimitXDSCache(vhosts []string, resolveRatelimit dpv1alpha1.ResolveRateLimitAPIPolicy) {
	// Add Rate Limit inline policies in API to the cache
	rlsPolicyCache.AddAPILevelRateLimitPolicies(vhosts, resolveRatelimit)
}

// UpdateRateLimitXDSCacheForCustomPolicies updates the xDS cache of the RateLimiter for custom policies.
func UpdateRateLimitXDSCacheForCustomPolicies(customRateLimitPolicies dpv1alpha1.CustomRateLimitPolicyDef) {
	if customRateLimitPolicies.Key != "" {
		rlsPolicyCache.AddCustomRateLimitPolicies(customRateLimitPolicies)
	}
}

// DeleteAPILevelRateLimitPolicies delete the ratelimit xds cache
func DeleteAPILevelRateLimitPolicies(resolveRatelimit dpv1alpha1.ResolveRateLimitAPIPolicy) {
	var org = resolveRatelimit.Organization
	var vhost = resolveRatelimit.Vhost
	var basePath = resolveRatelimit.BasePath
	rlsPolicyCache.DeleteAPILevelRateLimitPolicies(org, vhost, basePath)
}

// DeleteResourceLevelRateLimitPolicies delete the ratelimit xds cache
func DeleteResourceLevelRateLimitPolicies(resolveRatelimit dpv1alpha1.ResolveRateLimitAPIPolicy) {
	var org = resolveRatelimit.Organization
	var vhost = resolveRatelimit.Vhost
	var basePath = resolveRatelimit.BasePath
	var path = resolveRatelimit.Resources[0].Path
	var method = resolveRatelimit.Resources[0].Method
	rlsPolicyCache.DeleteResourceLevelRateLimitPolicies(org, vhost, basePath, path, method)
}

// DeleteCustomRateLimitPolicies delete the ratelimit xds cache
func DeleteCustomRateLimitPolicies(customRateLimitPolicy dpv1alpha1.CustomRateLimitPolicyDef) {
	rlsPolicyCache.DeleteCustomRateLimitPolicies(customRateLimitPolicy)
}

// GenerateIdentifierForAPIWithUUID generates an identifier unique to the API
func GenerateIdentifierForAPIWithUUID(vhost, uuid string) string {
	return fmt.Sprint(vhost, apiKeyFieldSeparator, uuid)
}

// UpdateRateLimiterPolicies update the rate limiter xDS cache with latest rate limit policies
func UpdateRateLimiterPolicies(label string) {
	_ = rlsPolicyCache.updateXdsCache(label)
}

// SetEmptySnapshotupdate update empty snapshot
func SetEmptySnapshotupdate(lable string) bool {
	return rlsPolicyCache.SetEmptySnapshot(lable)
}

// GetXdsCache returns xds server cache.
func GetXdsCache() envoy_cachev3.SnapshotCache {
	return cache
}

// GetEnforcerCache returns xds server cache.
func GetEnforcerCache() wso2_cache.SnapshotCache {
	return enforcerCache
}

// GetEnforcerSubscriptionCache returns xds server cache.
func GetEnforcerSubscriptionCache() wso2_cache.SnapshotCache {
	return enforcerSubscriptionCache
}

// GetEnforcerApplicationCache returns xds server cache.
func GetEnforcerApplicationCache() wso2_cache.SnapshotCache {
	return enforcerApplicationCache
}

// GetEnforcerApplicationKeyMappingCache returns xds server cache.
func GetEnforcerApplicationKeyMappingCache() wso2_cache.SnapshotCache {
	return enforcerApplicationKeyMappingCache
}

// // UpdateEnforcerConfig Sets new update to the enforcer's configuration
// func UpdateEnforcerConfig(configFile *config.Config) {
// 	// TODO: (Praminda) handle labels
// 	label := commonEnforcerLabel
// 	config := config.ReadConfigs()
// 	// configs := []types.Resource{MarshalConfig(configFile)}
// 	version, _ := crand.Int(crand.Reader, maxRandomBigInt())
// 	snap, errNewSnap := wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
// 		wso2_resource.ConfigType: configs,
// 	})
// 	if errNewSnap != nil {
// 		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error1413, logging.MAJOR, "Error creating new snapshot : %v", errNewSnap.Error()))
// 	}
// 	snap.Consistent()

// 	errSetSnap := enforcerCache.SetSnapshot(context.Background(), label, snap)
// 	if errSetSnap != nil {
// 		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error1414, logging.MAJOR, "Error while setting the snapshot : %v", errSetSnap.Error()))
// 	}

// 	enforcerLabelMap[label].configs = configs
// 	loggers.LoggerXds.Infof("New Config cache update for the label: " + label + " version: " + fmt.Sprint(version))
// }

// UpdateEnforcerApplications sets new update to the enforcer's Applications
func UpdateEnforcerApplications(applications *subscription.ApplicationList) {
	loggers.LoggerXds.Debug("Updating Enforcer Application Cache")
	label := commonEnforcerLabel
	applicationList := append(enforcerLabelMap[label].applications, applications)
	version, _ := crand.Int(crand.Reader, maxRandomBigInt())
	snap, _ := wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
		wso2_resource.ApplicationListType: applicationList,
	})
	snap.Consistent()
	errSetSnap := enforcerApplicationCache.SetSnapshot(context.Background(), label, snap)
	if errSetSnap != nil {
		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error1414, logging.MAJOR, "Error while setting the snapshot : %v", errSetSnap.Error()))
	}
	enforcerLabelMap[label].applications = applicationList
	loggers.LoggerXds.Infof("New Application cache update for the label: " + label + " version: " + fmt.Sprint(version))
}

// UpdateEnforcerApplicationKeyMappings sets new update to the enforcer's Application Key Mappings
func UpdateEnforcerApplicationKeyMappings(applicationKeyMappings *subscription.ApplicationKeyMappingList) {
	loggers.LoggerXds.Debug("Updating Application Key Mapping Cache")
	label := commonEnforcerLabel
	applicationKeyMappingList := append(enforcerLabelMap[label].applicationKeyMappings, applicationKeyMappings)
	version, _ := crand.Int(crand.Reader, maxRandomBigInt())
	snap, _ := wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
		wso2_resource.ApplicationKeyMappingListType: applicationKeyMappingList,
	})
	snap.Consistent()
	errSetSnap := enforcerApplicationKeyMappingCache.SetSnapshot(context.Background(), label, snap)
	if errSetSnap != nil {
		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error1414, logging.MAJOR, "Error while setting the snapshot : %v", errSetSnap.Error()))
	}
	enforcerLabelMap[label].applicationKeyMappings = applicationKeyMappingList
	loggers.LoggerXds.Infof("New Application Key Mapping cache update for the label: " + label + " version: " + fmt.Sprint(version))
}

// UpdateEnforcerSubscriptions sets new update to the enforcer's Subscriptions
func UpdateEnforcerSubscriptions(subscriptions *subscription.SubscriptionList) {
	//TODO: (Dinusha) check this hardcoded value
	loggers.LoggerXds.Debug("Updating Enforcer Subscription Cache")
	label := commonEnforcerLabel
	subscriptionList := append(enforcerLabelMap[label].subscriptions, subscriptions)

	// TODO: (VirajSalaka) Decide if a map is required to keep version (just to avoid having the same version)
	version, _ := crand.Int(crand.Reader, maxRandomBigInt())
	snap, _ := wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
		wso2_resource.SubscriptionListType: subscriptionList,
	})
	snap.Consistent()

	errSetSnap := enforcerSubscriptionCache.SetSnapshot(context.Background(), label, snap)
	if errSetSnap != nil {
		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error1414, logging.MAJOR, "Error while setting the snapshot : %v", errSetSnap.Error()))
	}
	enforcerLabelMap[label].subscriptions = subscriptionList
	loggers.LoggerXds.Infof("New Subscription cache update for the label: " + label + " version: " + fmt.Sprint(version))
}

// UpdateEnforcerApplicationMappings sets new update to the enforcer's Application Mappings
func UpdateEnforcerApplicationMappings(applicationMappings *subscription.ApplicationMappingList) {
	loggers.LoggerXds.Debug("Updating Application Mapping Cache")
	label := commonEnforcerLabel
	applicationMappingList := append(enforcerLabelMap[label].applicationMappings, applicationMappings)
	version, _ := crand.Int(crand.Reader, maxRandomBigInt())
	snap, _ := wso2_cache.NewSnapshot(fmt.Sprint(version), map[wso2_resource.Type][]types.Resource{
		wso2_resource.ApplicationMappingListType: applicationMappingList,
	})
	snap.Consistent()
	errSetSnap := enforcerApplicationMappingCache.SetSnapshot(context.Background(), label, snap)
	if errSetSnap != nil {
		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error1414, logging.MAJOR, "Error while setting the snapshot : %v", errSetSnap.Error()))
	}
	enforcerLabelMap[label].applicationMappings = applicationMappingList
	loggers.LoggerXds.Infof("New Application Mapping cache update for the label: " + label + " version: " + fmt.Sprint(version))
}
