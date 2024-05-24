package lakeformation

import types "DesignSphere_Server/src/resource/aws/types"

type DataLakeSettings struct {
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions []types.Lakeformation_DataLakeSettingsCreateDatabaseDefaultPermission `json:"createDatabaseDefaultPermissions,omitempty" yaml:"createDatabaseDefaultPermissions,omitempty"`

	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions []types.Lakeformation_DataLakeSettingsCreateTableDefaultPermission `json:"createTableDefaultPermissions,omitempty" yaml:"createTableDefaultPermissions,omitempty"`

	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
	ExternalDataFilteringAllowLists []string `json:"externalDataFilteringAllowLists,omitempty" yaml:"externalDataFilteringAllowLists,omitempty"`

	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners []string `json:"trustedResourceOwners,omitempty" yaml:"trustedResourceOwners,omitempty"`

	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	AllowExternalDataFiltering bool `json:"allowExternalDataFiltering,omitempty" yaml:"allowExternalDataFiltering,omitempty"`

	/*
	   Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.

	   > --NOTE:-- Although optional, not including `admins`, `create_database_default_permissions`, `create_table_default_permissions`, and/or `trusted_resource_owners` results in the setting being cleared.
	*/
	AuthorizedSessionTagValueLists []string `json:"authorizedSessionTagValueLists,omitempty" yaml:"authorizedSessionTagValueLists,omitempty"`

	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Set of ARNs of AWS Lake Formation principals (IAM users or roles) with only view access to the resources.
	ReadOnlyAdmins []string `json:"readOnlyAdmins,omitempty" yaml:"readOnlyAdmins,omitempty"`

	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins []string `json:"admins,omitempty" yaml:"admins,omitempty"`
}
