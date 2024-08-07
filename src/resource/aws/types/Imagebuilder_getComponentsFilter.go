package types

type Imagebuilder_getComponentsFilter struct {
	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Name of the filter field. Valid values can be found in the [Image Builder ListComponents API Reference](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_ListComponents.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
