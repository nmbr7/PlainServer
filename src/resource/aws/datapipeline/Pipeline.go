package datapipeline

type Pipeline struct {
	// The name of Pipeline.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of Pipeline.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
