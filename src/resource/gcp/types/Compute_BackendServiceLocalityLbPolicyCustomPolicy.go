package types

type Compute_BackendServiceLocalityLbPolicyCustomPolicy struct {
	/*
	   An optional, arbitrary JSON object with configuration data, understood
	   by a locally installed custom policy implementation.
	*/
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	/*
	   Identifies the custom policy.
	   The value should match the type the custom implementation is registered
	   with on the gRPC clients. It should follow protocol buffer
	   message naming conventions and include the full path (e.g.
	   myorg.CustomLbPolicy). The maximum length is 256 characters.
	   Note that specifying the same custom policy more than once for a
	   backend is not a valid configuration and will be rejected.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
