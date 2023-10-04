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

type DataCatalogInitParameters struct {

	// Description of the data catalog.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key value pairs that specifies the Lambda function or functions to use for the data catalog. The mapping used depends on the catalog type.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of data catalog: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataCatalogObservation struct {

	// ARN of the data catalog.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the data catalog.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the data catalog.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key value pairs that specifies the Lambda function or functions to use for the data catalog. The mapping used depends on the catalog type.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Type of data catalog: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataCatalogParameters struct {

	// Description of the data catalog.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key value pairs that specifies the Lambda function or functions to use for the data catalog. The mapping used depends on the catalog type.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of data catalog: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataCatalogSpec defines the desired state of DataCatalog
type DataCatalogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataCatalogParameters `json:"forProvider"`
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
	InitProvider DataCatalogInitParameters `json:"initProvider,omitempty"`
}

// DataCatalogStatus defines the observed state of DataCatalog.
type DataCatalogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataCatalogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataCatalog is the Schema for the DataCatalogs API. Provides an Athena data catalog.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DataCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || has(self.initProvider.description)",message="description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parameters) || has(self.initProvider.parameters)",message="parameters is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="type is a required parameter"
	Spec   DataCatalogSpec   `json:"spec"`
	Status DataCatalogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataCatalogList contains a list of DataCatalogs
type DataCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalog `json:"items"`
}

// Repository type metadata.
var (
	DataCatalog_Kind             = "DataCatalog"
	DataCatalog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataCatalog_Kind}.String()
	DataCatalog_KindAPIVersion   = DataCatalog_Kind + "." + CRDGroupVersion.String()
	DataCatalog_GroupVersionKind = CRDGroupVersion.WithKind(DataCatalog_Kind)
)

func init() {
	SchemeBuilder.Register(&DataCatalog{}, &DataCatalogList{})
}
