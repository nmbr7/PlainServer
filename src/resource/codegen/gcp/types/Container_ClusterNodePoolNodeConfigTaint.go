package types

type Container_ClusterNodePoolNodeConfigTaint struct {
	// Key for taint.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value for taint.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Effect for taint. Accepted values are `NO_SCHEDULE`, `PREFER_NO_SCHEDULE`, and `NO_EXECUTE`.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`
}
