/*
Copyright 2020 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VesselGpsTypeParameters are the configurable fields of a MyType.
type VesselGpsTypeParameters struct {
	IMO string `json:"imo"`
}

// VesselGpsTypeObservation are the observable fields of a MyType.
type VesselGpsTypeObservation struct {
	Lat string `json:"lat,omitempty"`
	Lng string `json:"lng,omitempty"`
}

// A VesselGpsTypeSpec defines the desired state of a MyType.
type VesselGpsTypeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VesselGpsTypeParameters `json:"forProvider"`
}

// A VesselGpsTypeStatus represents the observed state of a MyType.
type VesselGpsTypeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VesselGpsTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A VesselGpsType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="IMO",type="string",JSONPath=".spec.forProvider.imo"
// +kubebuilder:printcolumn:name="LAT",type="string",JSONPath=".status.atProvider.lat"
// +kubebuilder:printcolumn:name="LNG",type="string",JSONPath=".status.atProvider.lng"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type VesselGpsType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VesselGpsTypeSpec   `json:"spec"`
	Status VesselGpsTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VesselGpsTypeList contains a list of MyType
type VesselGpsTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VesselGpsType `json:"items"`
}

// MyType type metadata.
var (
	VesselGpsTypeKind             = reflect.TypeOf(VesselGpsType{}).Name()
	VesselGpsTypeGroupKind        = schema.GroupKind{Group: Group, Kind: VesselGpsTypeKind}.String()
	VesselGpsTypeKindAPIVersion   = VesselGpsTypeKind + "." + SchemeGroupVersion.String()
	VesselGpsTypeGroupVersionKind = SchemeGroupVersion.WithKind(VesselGpsTypeKind)
)

func init() {
	SchemeBuilder.Register(&VesselGpsType{}, &VesselGpsTypeList{})
}
