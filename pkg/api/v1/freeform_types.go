package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FreeFormSpec defines the desired state of FreeForm
type FreeFormSpec struct {
}

// FreeFormStatus defines the observed state of FreeForm
type FreeFormStatus struct {
	ID *string `json:"id,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FreeForm is the Schema for the FreeForms API
type FreeForm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FreeFormSpec   `json:"spec,omitempty"`
	Status FreeFormStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FreeFormList contains a list of FreeForm
type FreeFormList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FreeForm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FreeForm{}, &FreeFormList{})
}
