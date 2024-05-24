package types

type Macie2_ClassificationJobS3JobDefinitionScopingExcludesAnd struct {
	// A property-based condition that defines a property, operator, and one or more values for including or excluding an object from the job. (documented below)
	SimpleScopeTerm Macie2_ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm `json:"simpleScopeTerm,omitempty" yaml:"simpleScopeTerm,omitempty"`

	// A tag-based condition that defines the operator and tag keys or tag key and value pairs for including or excluding an object from the job. (documented below)
	TagScopeTerm Macie2_ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm `json:"tagScopeTerm,omitempty" yaml:"tagScopeTerm,omitempty"`
}
