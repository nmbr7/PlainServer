package types

type Quicksight_DataSetPhysicalTableMapCustomSqlColumn struct {
	// Name of this column in the underlying data source.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Data type of the column.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
