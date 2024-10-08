package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig struct {
	// Determines the level of granularity that's included in the prefix. Valid values are `YEAR`, `MONTH`, `DAY`, `HOUR`, and `MINUTE`.
	PrefixFormat string `json:"prefixFormat,omitempty" yaml:"prefixFormat,omitempty"`

	// Determines whether the destination file path includes either or both of the selected elements. Valid values are `EXECUTION_ID` and `SCHEMA_VERSION`
	PrefixHierarchies []string `json:"prefixHierarchies,omitempty" yaml:"prefixHierarchies,omitempty"`

	// Determines the format of the prefix, and whether it applies to the file name, file path, or both. Valid values are `FILENAME`, `PATH`, and `PATH_AND_FILENAME`.
	PrefixType string `json:"prefixType,omitempty" yaml:"prefixType,omitempty"`
}
