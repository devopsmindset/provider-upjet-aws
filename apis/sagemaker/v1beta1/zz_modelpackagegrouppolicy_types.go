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

type ModelPackageGroupPolicyInitParameters struct {
	ResourcePolicy *string `json:"resourcePolicy,omitempty" tf:"resource_policy,omitempty"`
}

type ModelPackageGroupPolicyObservation struct {

	// The name of the Model Package Package Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the model package group.
	ModelPackageGroupName *string `json:"modelPackageGroupName,omitempty" tf:"model_package_group_name,omitempty"`

	ResourcePolicy *string `json:"resourcePolicy,omitempty" tf:"resource_policy,omitempty"`
}

type ModelPackageGroupPolicyParameters struct {

	// The name of the model package group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sagemaker/v1beta1.ModelPackageGroup
	// +kubebuilder:validation:Optional
	ModelPackageGroupName *string `json:"modelPackageGroupName,omitempty" tf:"model_package_group_name,omitempty"`

	// Reference to a ModelPackageGroup in sagemaker to populate modelPackageGroupName.
	// +kubebuilder:validation:Optional
	ModelPackageGroupNameRef *v1.Reference `json:"modelPackageGroupNameRef,omitempty" tf:"-"`

	// Selector for a ModelPackageGroup in sagemaker to populate modelPackageGroupName.
	// +kubebuilder:validation:Optional
	ModelPackageGroupNameSelector *v1.Selector `json:"modelPackageGroupNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourcePolicy *string `json:"resourcePolicy,omitempty" tf:"resource_policy,omitempty"`
}

// ModelPackageGroupPolicySpec defines the desired state of ModelPackageGroupPolicy
type ModelPackageGroupPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModelPackageGroupPolicyParameters `json:"forProvider"`
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
	InitProvider ModelPackageGroupPolicyInitParameters `json:"initProvider,omitempty"`
}

// ModelPackageGroupPolicyStatus defines the observed state of ModelPackageGroupPolicy.
type ModelPackageGroupPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModelPackageGroupPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ModelPackageGroupPolicy is the Schema for the ModelPackageGroupPolicys API. Provides a SageMaker Model Package Group Policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ModelPackageGroupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourcePolicy) || has(self.initProvider.resourcePolicy)",message="resourcePolicy is a required parameter"
	Spec   ModelPackageGroupPolicySpec   `json:"spec"`
	Status ModelPackageGroupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelPackageGroupPolicyList contains a list of ModelPackageGroupPolicys
type ModelPackageGroupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelPackageGroupPolicy `json:"items"`
}

// Repository type metadata.
var (
	ModelPackageGroupPolicy_Kind             = "ModelPackageGroupPolicy"
	ModelPackageGroupPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ModelPackageGroupPolicy_Kind}.String()
	ModelPackageGroupPolicy_KindAPIVersion   = ModelPackageGroupPolicy_Kind + "." + CRDGroupVersion.String()
	ModelPackageGroupPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ModelPackageGroupPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ModelPackageGroupPolicy{}, &ModelPackageGroupPolicyList{})
}
