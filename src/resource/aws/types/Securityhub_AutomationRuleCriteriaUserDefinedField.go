package types

type Securityhub_AutomationRuleCriteriaUserDefinedField struct {
	// The key of the map filter.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// A date range value for the date filter, provided as an Integer.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The condition to apply to a string value when querying for findings. Valid values include: `EQUALS` and `NOT_EQUALS`.
	Comparison string `json:"comparison,omitempty" yaml:"comparison,omitempty"`
}
