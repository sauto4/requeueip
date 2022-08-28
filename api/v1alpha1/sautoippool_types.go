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

// SautoIPPoolSpec defines the desired state of SautoIPPool.
type SautoIPPoolSpec struct {
	// Foo is an example field of SautoIPPool. Edit sautoippool_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SautoIPPoolStatus defines the observed state of SautoIPPool.
type SautoIPPoolStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SautoIPPool is the Schema for the sautoippools API.
type SautoIPPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SautoIPPoolSpec   `json:"spec,omitempty"`
	Status SautoIPPoolStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SautoIPPoolList contains a list of SautoIPPool.
type SautoIPPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SautoIPPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SautoIPPool{}, &SautoIPPoolList{})
}
