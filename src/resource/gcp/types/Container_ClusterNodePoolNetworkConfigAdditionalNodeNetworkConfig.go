package types

type Container_ClusterNodePoolNetworkConfigAdditionalNodeNetworkConfig struct {
	/*
	   The name or self_link of the Google Compute Engine
	   subnetwork in which the cluster's instances are launched.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   The name or self_link of the Google Compute Engine
	   network to which the cluster is connected. For Shared VPC, set this to the self link of the
	   shared network.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
