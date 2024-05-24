package apigateway

type ApiKey struct {
	// An Amazon Web Services Marketplace customer identifier, when integrating with the Amazon Web Services SaaS Marketplace.
	CustomerId string `json:"customerId,omitempty" yaml:"customerId,omitempty"`

	// API key description. Defaults to "Managed by Pulumi".
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether the API key can be used by callers. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Name of the API key.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Value of the API key. If specified, the value must be an alphanumeric string between 20 and 128 characters. If not specified, it will be automatically generated by AWS on creation.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
