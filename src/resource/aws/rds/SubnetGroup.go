package rds

type SubnetGroup struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the DB subnet group. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the DB subnet group. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// A list of VPC subnet IDs.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
