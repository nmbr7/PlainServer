package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigHybridOptions struct {
	// A short description of where the data is coming from. Will be stored once in the job. 256 max length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   To organize findings, these labels will be added to each finding.
	   Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z?`.
	   Label values must be between 0 and 63 characters long and must conform to the regular expression `(a-z?)?`.
	   No more than 10 labels can be associated with a given finding.
	   Examples:
	   - `"environment" : "production"`
	   - `"pipeline" : "etl"`
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   These are labels that each inspection request must include within their 'finding_labels' map. Request
	   may contain others, but any missing one of these will be rejected.
	   Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z?`.
	   No more than 10 keys can be required.
	*/
	RequiredFindingLabelKeys []string `json:"requiredFindingLabelKeys,omitempty" yaml:"requiredFindingLabelKeys,omitempty"`

	/*
	   If the container is a table, additional information to make findings meaningful such as the columns that are primary keys.
	   Structure is documented below.
	*/
	TableOptions Dataloss_PreventionJobTriggerInspectJobStorageConfigHybridOptionsTableOptions `json:"tableOptions,omitempty" yaml:"tableOptions,omitempty"`
}
