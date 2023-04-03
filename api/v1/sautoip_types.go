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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SautoIPStatus defines the observed state of SautoIP.
type SautoIPStatus struct{}

// +kubebuilder:resource:categories={requeueip},path="sautoips",scope="Cluster",shortName={si},singular="sautoip"
// +kubebuilder:object:root=true

// SautoIP is the Schema for the sautoips API.
type SautoIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status SautoIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SautoIPList contains a list of SautoIP.
type SautoIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SautoIP `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SautoIP{}, &SautoIPList{})
}
