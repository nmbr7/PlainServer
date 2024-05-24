package types

type Ec2_getSecurityGroupFilter struct {
	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html).
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Set of values that are accepted for the given field.
	   A Security Group will be selected if any one of the given values matches.
	*/
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
