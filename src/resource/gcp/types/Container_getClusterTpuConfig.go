package types

type Container_getClusterTpuConfig struct {
	// Whether Cloud TPU integration is enabled or not
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// IPv4 CIDR block reserved for Cloud TPU in the VPC.
	Ipv4CidrBlock string `json:"ipv4CidrBlock,omitempty" yaml:"ipv4CidrBlock,omitempty"`

	// Whether to use service networking for Cloud TPU or not
	UseServiceNetworking bool `json:"useServiceNetworking,omitempty" yaml:"useServiceNetworking,omitempty"`
}
