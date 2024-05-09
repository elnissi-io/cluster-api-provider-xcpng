/*
Copyright 2024 elnissi-io.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// XCPngClusterSpec defines the desired state of XCPngCluster
type XCPngClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	NetworkSettings string `json:"networkSettings,omitempty"`
}

// XCPngClusterStatus defines the observed state of XCPngCluster
type XCPngClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - observed state of cluster
	Ready bool `json:"ready"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// XCPngCluster is the Schema for the xcpngclusters API
type XCPngCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   XCPngClusterSpec   `json:"spec,omitempty"`
	Status XCPngClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// XCPngClusterList contains a list of XCPngCluster
type XCPngClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XCPngCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&XCPngCluster{}, &XCPngClusterList{})
}
