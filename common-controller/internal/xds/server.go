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
	"fmt"
	"context"
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"time"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	envoy_cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/subscription"
	"github.com/wso2/apk/adapter/pkg/logging"
	"github.com/wso2/apk/common-controller/internal/loggers"
	dpv1alpha1 "github.com/wso2/apk/common-controller/internal/operator/apis/dp/v1alpha1"
	wso2_cache "github.com/wso2/apk/adapter/pkg/discovery/protocol/cache/v3"
	wso2_resource "github.com/wso2/apk/adapter/pkg/discovery/protocol/resource/v3"
)

// EnforcerInternalAPI struct use to hold enforcer resources
type EnforcerInternalAPI struct {
	configs                []types.Resource
	keyManagers            []types.Resource
	subscriptions          []types.Resource
	applications           []types.Resource
	apiList                []types.Resource
	applicationPolicies    []types.Resource
	subscriptionPolicies   []types.Resource
	applicationKeyMappings []types.Resource
	revokedTokens          []types.Resource
	jwtIssuers             []types.Resource
}

var (
	enforcerApplicationCache           wso2_cache.SnapshotCache
	enforcerApplicationKeyMappingCache wso2_cache.SnapshotCache
	// Common Enforcer Label as map key
	// TODO(amali) This doesn't have a usage yet. It will be used to handle multiple enforcer labels in future.
	enforcerLabelMap map[string]*EnforcerInternalAPI // Enforcer Label -> EnforcerInternalAPI struct map
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
	enforcerApplicationCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	enforcerLabelMap = make(map[string]*EnforcerInternalAPI)
	//TODO(amali) currently subscriptions, configs, applications, applicationPolicies, subscriptionPolicies,
	// applicationKeyMappings, keyManagerConfigList, revokedTokens are supported with the hard coded label for Enforcer
	enforcerLabelMap[commonEnforcerLabel] = &EnforcerInternalAPI{}
	rand.Seed(time.Now().UnixNano())
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
