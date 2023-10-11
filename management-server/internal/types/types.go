/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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

package types

// ApplicationEvent is a data holder for an application event
type ApplicationEvent struct {
	Label         string
	UUID          string
	Name          string
	Owner         string
	// Policy        string
	Attributes    map[string]string
	// Keys          []*ApplicationKey
	Organization  string
	TimeStamp     string
	IsRemoveEvent bool
}

// ApplicationKey is a data holder for an application key
// type ApplicationKey struct {
// 	Key        string
// 	KeyManager string
// }

// SubscriptionEvent is a data holder for a subscription event
type SubscriptionEvent struct {
	Label          string
	UUID           string
	// ApplicationRef string
	// APIRef         string
	// PolicyID       string
	SubStatus      string
	// Subscriber     string
	Organization   string
	TimeStamp      string
	IsRemoveEvent  bool
	IsUpdateEvent  bool
}
