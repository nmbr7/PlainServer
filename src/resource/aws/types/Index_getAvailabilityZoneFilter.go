package types

type Index_getAvailabilityZoneFilter struct {
	// Name of the filter field. Valid values can be found in the [EC2 DescribeAvailabilityZones API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
