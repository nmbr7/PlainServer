package types

type Cloudfront_ContinuousDeploymentPolicyTrafficConfigSingleWeightConfigSessionStickinessConfig struct {
	// The amount of time in seconds after which sessions will cease if no requests are received. Valid values are `300` – `3600` (5–60 minutes). The value must be less than or equal to `maximum_ttl`.
	IdleTtl int `json:"idleTtl,omitempty" yaml:"idleTtl,omitempty"`

	// The maximum amount of time in seconds to consider requests from the viewer as being part of the same session. Valid values are `300` – `3600` (5–60 minutes). The value must be greater than or equal to `idle_ttl`.
	MaximumTtl int `json:"maximumTtl,omitempty" yaml:"maximumTtl,omitempty"`
}
