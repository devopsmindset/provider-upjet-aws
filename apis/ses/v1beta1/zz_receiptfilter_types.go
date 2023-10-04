// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ReceiptFilterInitParameters struct {

	// The IP address or address range to filter, in CIDR notation
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Block or Allow
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type ReceiptFilterObservation struct {

	// The SES receipt filter ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The IP address or address range to filter, in CIDR notation
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The SES receipt filter name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Block or Allow
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type ReceiptFilterParameters struct {

	// The IP address or address range to filter, in CIDR notation
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Block or Allow
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ReceiptFilterSpec defines the desired state of ReceiptFilter
type ReceiptFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReceiptFilterParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ReceiptFilterInitParameters `json:"initProvider,omitempty"`
}

// ReceiptFilterStatus defines the observed state of ReceiptFilter.
type ReceiptFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReceiptFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReceiptFilter is the Schema for the ReceiptFilters API. Provides an SES receipt filter
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReceiptFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidr) || has(self.initProvider.cidr)",message="cidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || has(self.initProvider.policy)",message="policy is a required parameter"
	Spec   ReceiptFilterSpec   `json:"spec"`
	Status ReceiptFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReceiptFilterList contains a list of ReceiptFilters
type ReceiptFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReceiptFilter `json:"items"`
}

// Repository type metadata.
var (
	ReceiptFilter_Kind             = "ReceiptFilter"
	ReceiptFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReceiptFilter_Kind}.String()
	ReceiptFilter_KindAPIVersion   = ReceiptFilter_Kind + "." + CRDGroupVersion.String()
	ReceiptFilter_GroupVersionKind = CRDGroupVersion.WithKind(ReceiptFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&ReceiptFilter{}, &ReceiptFilterList{})
}
