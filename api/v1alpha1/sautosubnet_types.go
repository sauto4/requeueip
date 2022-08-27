/*
Copyright 2022 The RequeueIP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SautoSubnetSpec defines the desired state of SautoSubnet
type SautoSubnetSpec struct {
	// Foo is an example field of SautoSubnet. Edit sautosubnet_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SautoSubnetStatus defines the observed state of SautoSubnet
type SautoSubnetStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SautoSubnet is the Schema for the sautosubnets API
type SautoSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SautoSubnetSpec   `json:"spec,omitempty"`
	Status SautoSubnetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SautoSubnetList contains a list of SautoSubnet
type SautoSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SautoSubnet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SautoSubnet{}, &SautoSubnetList{})
}
