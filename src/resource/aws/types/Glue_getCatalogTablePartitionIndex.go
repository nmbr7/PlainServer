package types

type Glue_getCatalogTablePartitionIndex struct {
	//
	IndexStatus string `json:"indexStatus,omitempty" yaml:"indexStatus,omitempty"`

	// Keys for the partition index.
	Keys []string `json:"keys,omitempty" yaml:"keys,omitempty"`

	// Name of the partition index.
	IndexName string `json:"indexName,omitempty" yaml:"indexName,omitempty"`
}
