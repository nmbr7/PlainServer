package types

type Glue_CatalogDatabaseFederatedDatabase struct {
	// Name of the connection to the external metastore.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// Unique identifier for the federated database.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
}
