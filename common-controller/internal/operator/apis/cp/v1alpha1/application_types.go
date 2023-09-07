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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	Name                  string                 `json:"name"`
	Owner                 string                 `json:"owner"`
	Attributes            map[string]string      `json:"attributes,omitempty"`
	AuthenticationOptions []AuthenticationOption `json:"authenticationOptions"`
}

// AuthenticationOption defines an authentication option
type AuthenticationOption struct {
	// Type denotes the application authentication type
	// Possible values are OAuth2, mTLS, BasicAuth, APIKey, IPRange
	//
	// +kubebuilder:validation:Enum=oauth2;mTLS;basicAuth;apiKey;IPRange
	Type string `json:"type"`

	// OAuth2 denotes the OAuth2 authentication option
	OAuth2 *OAuth2 `json:"oauth2,omitempty"`

	// mTLS denotes the mTLS authentication option
	MTLS *mTLS `json:"mTLS,omitempty"`

	// BasicAuth denotes the Basic authentication option
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"`

	// APIKey denotes the APIKey authentication option
	APIKey *APIKey `json:"apiKey,omitempty"`

	// IPRange denotes the IPRange authentication option
	IPRange *IPRange `json:"IPRange,omitempty"`
}

// OAuth2 defines the OAuth2 configurations
type OAuth2 struct {
	ConsumerKey string `json:"consumerKey"`
	KeyManager  string `json:"keyManager"`
	KeyType     string `json:"keyType"`
}

// mTLS defines the mTLS configurations
type mTLS struct {
	Cert    string `json:"cert"`
	Issuer  string `json:"issuer"`
	KeyType string `json:"keyType"`
}

// BasicAuth defines the BasicAuth configurations
type BasicAuth struct {
	BasicAuthIdentifier string `json:"basicAuthIdentifier"`
	KeyType             string `json:"keyType"`
}

// APIKey defines the APIKey configurations
type APIKey struct {
	Key     string `json:"key"`
	KeyType string `json:"keyType"`
}

// IPRange defines the IPRange configurations
type IPRange struct {
	IPs               string `json:"IPs"`
	IPRangeIdentifier string `json:"IPRangeIdentifier"`
	KeyType           string `json:"keyType"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
