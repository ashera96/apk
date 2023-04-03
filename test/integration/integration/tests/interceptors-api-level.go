/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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

package tests

import (
	"testing"

	"github.com/wso2/apk/test/integration/integration/utils/http"
	"github.com/wso2/apk/test/integration/integration/utils/suite"
)

func init() {
	// TODO (Amila): Uncomment once AKS test runner is enabled
	// IntegrationTests = append(IntegrationTests, APILevelInterceptors)
}

// APILevelInterceptors test
var APILevelInterceptors = suite.IntegrationTest{
	ShortName:   "APILevelInterceptors",
	Description: "Tests API with with request and response interceptors",
	Manifests:   []string{"tests/interceptors-api-level.yaml"},
	Test: func(t *testing.T, suite *suite.IntegrationTestSuite) {
		gwAddr := "interceptor-api.test.gw.wso2.com:9095"
		token := http.GetTestToken(t)

		testCases := []http.ExpectedResponse{
			{
				Request: http.Request{
					Host: "interceptor-api.test.gw.wso2.com",
					Path: "/interceptor/2.0.0/books",
					Headers: map[string]string{
						"content-type": "application/json",
					},
					Method: "POST",
					Body:   `{"name":"The Prisoner"}`,
				},
				Response: http.Response{StatusCode: 201},
			},
		}
		for i := range testCases {
			tc := testCases[i]
			tc.Request.Headers = http.AddBearerTokenToHeader(token, tc.Request.Headers)
			t.Run(tc.GetTestCaseName(i), func(t *testing.T) {
				t.Parallel()
				http.MakeRequestAndExpectEventuallyConsistentResponse(t, suite.RoundTripper, suite.TimeoutConfig, gwAddr, tc)
			})
		}
	},
}
