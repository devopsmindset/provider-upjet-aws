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

type LicenseAssociationInitParameters struct {

	// The type of license for the workspace license association. Valid values are ENTERPRISE and ENTERPRISE_FREE_TRIAL.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`
}

type LicenseAssociationObservation struct {

	// If license_type is set to ENTERPRISE_FREE_TRIAL, this is the expiration date of the free trial.
	FreeTrialExpiration *string `json:"freeTrialExpiration,omitempty" tf:"free_trial_expiration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If license_type is set to ENTERPRISE, this is the expiration date of the enterprise license.
	LicenseExpiration *string `json:"licenseExpiration,omitempty" tf:"license_expiration,omitempty"`

	// The type of license for the workspace license association. Valid values are ENTERPRISE and ENTERPRISE_FREE_TRIAL.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// The workspace id.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type LicenseAssociationParameters struct {

	// The type of license for the workspace license association. Valid values are ENTERPRISE and ENTERPRISE_FREE_TRIAL.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The workspace id.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/grafana/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in grafana to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in grafana to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// LicenseAssociationSpec defines the desired state of LicenseAssociation
type LicenseAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LicenseAssociationParameters `json:"forProvider"`
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
	InitProvider LicenseAssociationInitParameters `json:"initProvider,omitempty"`
}

// LicenseAssociationStatus defines the observed state of LicenseAssociation.
type LicenseAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LicenseAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseAssociation is the Schema for the LicenseAssociations API. Provides an Amazon Managed Grafana workspace license association resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LicenseAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.licenseType) || has(self.initProvider.licenseType)",message="licenseType is a required parameter"
	Spec   LicenseAssociationSpec   `json:"spec"`
	Status LicenseAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseAssociationList contains a list of LicenseAssociations
type LicenseAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseAssociation `json:"items"`
}

// Repository type metadata.
var (
	LicenseAssociation_Kind             = "LicenseAssociation"
	LicenseAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LicenseAssociation_Kind}.String()
	LicenseAssociation_KindAPIVersion   = LicenseAssociation_Kind + "." + CRDGroupVersion.String()
	LicenseAssociation_GroupVersionKind = CRDGroupVersion.WithKind(LicenseAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&LicenseAssociation{}, &LicenseAssociationList{})
}
