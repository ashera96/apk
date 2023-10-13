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

package commoncontroller

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	envoy_cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	xdsv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	enforcerCallbacks "github.com/wso2/apk/common-controller/internal/xds/enforcercallbacks"
	routercb "github.com/wso2/apk/common-controller/internal/xds/routercallbacks"
	apiservice "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/api"
	configservice "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/config"
	subscriptionservice "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/subscription"
	wso2_server "github.com/wso2/apk/adapter/pkg/discovery/protocol/server/v3"
	"github.com/wso2/apk/adapter/pkg/health"
	healthservice "github.com/wso2/apk/adapter/pkg/health/api/wso2/health/service"
	"github.com/wso2/apk/common-controller/internal/config"
	"github.com/wso2/apk/common-controller/internal/operator"
	utils "github.com/wso2/apk/common-controller/internal/utils"
	xds "github.com/wso2/apk/common-controller/internal/xds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"github.com/wso2/apk/adapter/pkg/logging"
	"github.com/wso2/apk/common-controller/internal/loggers"
)

var (
	rlsPort uint
	cache   envoy_cachev3.SnapshotCache

	debug       bool
	onlyLogging bool

	port    uint
	alsPort uint

	mode string
)

const (
	maxRandomInt             int    = 999999999
	grpcMaxConcurrentStreams        = 1000000
	apiKeyFieldSeparator     string = ":"
	ads                             = "ads"
	amqpProtocol                    = "amqp"
)

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
	flag.UintVar(&rlsPort, "rls-port", 18005, "Rate Limiter management server port")

	flag.BoolVar(&debug, "debug", true, "Use debug logging")
	flag.BoolVar(&onlyLogging, "onlyLogging", false, "Only demo AccessLogging Service")
	flag.UintVar(&port, "port", 18002, "Management server port")
	flag.UintVar(&alsPort, "als", 18090, "Accesslog server port")
	flag.StringVar(&mode, "ads", ads, "Management server type (ads, xds, rest)")
}

func runRatelimitServer(rlsServer xdsv3.Server) {
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	publicKeyLocation, privateKeyLocation, truststoreLocation := utils.GetKeyLocations()
	cert, err := utils.GetServerCertificate(publicKeyLocation, privateKeyLocation)

	caCertPool := utils.GetTrustedCertPool(truststoreLocation)
	if err == nil {
		loggers.LoggerAPKOperator.Info("initiate the ssl context: ", err)
		grpcOptions = append(grpcOptions, grpc.Creds(
			credentials.NewTLS(&tls.Config{
				Certificates: []tls.Certificate{cert},
				ClientAuth:   tls.RequireAndVerifyClientCert,
				ClientCAs:    caCertPool,
			}),
		))
	} else {
		loggers.LoggerAPKOperator.Warn("failed to initiate the ssl context: ", err)
		panic(err)
	}

	grpcOptions = append(grpcOptions, grpc.KeepaliveParams(
		keepalive.ServerParameters{
			Time:    time.Duration(5 * time.Minute),
			Timeout: time.Duration(20 * time.Second),
		}),
	)
	rlsGrpcServer := grpc.NewServer(grpcOptions...)
	// It is required a separate gRPC server for the rate limit xDS, since it is the same RPC method
	// ADS used in both envoy xDS and rate limiter xDS.
	// According to https://github.com/envoyproxy/ratelimit/pull/368#discussion_r995831078 a separate RPC service is not
	// defined specifically to the rate limit xDS, instead using the ADS.
	loggers.LoggerAPKOperator.Info("port: ", rlsPort, " ratelimiter management server listening")
	rlsLis, err := net.Listen("tcp", fmt.Sprintf(":%d", rlsPort))
	if err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error1100, logging.BLOCKER, "Failed to listen on port: %v, error: %v", rlsPort, err.Error()))
	}

	discoveryv3.RegisterAggregatedDiscoveryServiceServer(rlsGrpcServer, rlsServer)
	// register health service
	healthservice.RegisterHealthServer(rlsGrpcServer, &health.Server{})
	go func() {
		loggers.LoggerAPKOperator.Info("Starting Rate Limiter xDS gRPC server.")
		if err = rlsGrpcServer.Serve(rlsLis); err != nil {
			loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error1105, logging.BLOCKER,
				"Error serving Rate Limiter xDS gRPC server on port %v, error: %v", rlsPort, err.Error()))
		}
	}()
}

func runCommonEnforcerServer(server xdsv3.Server, enforcerServer wso2_server.Server, enforcerSdsServer wso2_server.Server,
	enforcerAppDsSrv wso2_server.Server, enforcerAppKeyMappingDsSrv wso2_server.Server, enforcerAppMappingDsSrv wso2_server.Server,
	port uint) {
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	// TODO(Ashera): Add TLS support for Common Controller - Enforcer connection
	publicKeyLocation, privateKeyLocation, truststoreLocation := utils.GetKeyLocations()
	cert, err := utils.GetServerCertificate(publicKeyLocation, privateKeyLocation)

	caCertPool := utils.GetTrustedCertPool(truststoreLocation)

	if err == nil {
		grpcOptions = append(grpcOptions, grpc.Creds(
			credentials.NewTLS(&tls.Config{
				Certificates: []tls.Certificate{cert},
				ClientAuth:   tls.RequireAndVerifyClientCert,
				ClientCAs:    caCertPool,
			}),
		))
	} else {
		loggers.LoggerAPKOperator.Warn("failed to initiate the ssl context: ", err)
		panic(err)
	}

	grpcOptions = append(grpcOptions, grpc.KeepaliveParams(
		keepalive.ServerParameters{
			Time:    time.Duration(5 * time.Minute),
			Timeout: time.Duration(20 * time.Second),
		}),
	)
	grpcServer := grpc.NewServer(grpcOptions...)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error1100, logging.BLOCKER, "Failed to listen on port: %v, error: %v", port, err.Error()))
	}

	// register services
	discoveryv3.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	configservice.RegisterConfigDiscoveryServiceServer(grpcServer, enforcerServer)
	apiservice.RegisterApiDiscoveryServiceServer(grpcServer, enforcerServer)
	subscriptionservice.RegisterSubscriptionDiscoveryServiceServer(grpcServer, enforcerSdsServer)
	subscriptionservice.RegisterApplicationDiscoveryServiceServer(grpcServer, enforcerAppDsSrv)
	subscriptionservice.RegisterApplicationKeyMappingDiscoveryServiceServer(grpcServer, enforcerAppKeyMappingDsSrv)
	subscriptionservice.RegisterApplicationMappingDiscoveryServiceServer(grpcServer, enforcerAppMappingDsSrv)
	// register health service
	healthservice.RegisterHealthServer(grpcServer, &health.Server{})

	loggers.LoggerAPKOperator.Info("port: ", port, " common enforcer server listening")
	go func() {
		loggers.LoggerAPKOperator.Info("Starting XDS GRPC server.")
		if err = grpcServer.Serve(lis); err != nil {
			loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error1101, logging.BLOCKER, "Failed to start XDS GRPS server, error: %v", err.Error()))
		}
	}()
}

// InitCommonControllerServer initializes the gRPC server for the common controller.
func InitCommonControllerServer(conf *config.Config) {
	sig := make(chan os.Signal, 2)
	flag.Parse()
	signal.Notify(sig, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	loggers.LoggerAPKOperator.Info("Starting common controller ....")

	rateLimiterCache := xds.GetRateLimiterCache()
	rlsSrv := xdsv3.NewServer(ctx, rateLimiterCache, &xds.Callbacks{})

	// Start Rate Limiter xDS gRPC server
	runRatelimitServer(rlsSrv)
	// Set empty snapshot to initiate ratelimit service
	xds.SetEmptySnapshotupdate(conf.CommonController.Server.Label)

	cache := xds.GetXdsCache()
	enforcerCache := xds.GetEnforcerCache()
	enforcerSubscriptionCache := xds.GetEnforcerSubscriptionCache()
	enforcerApplicationCache := xds.GetEnforcerApplicationCache()
	enforcerApplicationKeyMappingCache := xds.GetEnforcerApplicationKeyMappingCache()
	enforcerApplicationMappingCache := xds.GetEnforcerApplicationMappingCache()
	srv := xdsv3.NewServer(ctx, cache, &routercb.Callbacks{})
	enforcerXdsSrv := wso2_server.NewServer(ctx, enforcerCache, &enforcerCallbacks.Callbacks{})
	enforcerSdsSrv := wso2_server.NewServer(ctx, enforcerSubscriptionCache, &enforcerCallbacks.Callbacks{})
	enforcerAppDsSrv := wso2_server.NewServer(ctx, enforcerApplicationCache, &enforcerCallbacks.Callbacks{})
	enforcerAppKeyMappingDsSrv := wso2_server.NewServer(ctx, enforcerApplicationKeyMappingCache, &enforcerCallbacks.Callbacks{})
	enforcerAppMappingDsSrv := wso2_server.NewServer(ctx, enforcerApplicationMappingCache, &enforcerCallbacks.Callbacks{})

	// Start Enforcer xDS gRPC server
	runCommonEnforcerServer(srv, enforcerXdsSrv, enforcerSdsSrv, enforcerAppDsSrv, enforcerAppKeyMappingDsSrv,
		enforcerAppMappingDsSrv, port)

	go operator.InitOperator()

OUTER:
	for {
		select {
		case s := <-sig:
			switch s {
			case os.Interrupt:
				loggers.LoggerAPKOperator.Info("Shutting down...")
				break OUTER
			}
		}
	}
}
