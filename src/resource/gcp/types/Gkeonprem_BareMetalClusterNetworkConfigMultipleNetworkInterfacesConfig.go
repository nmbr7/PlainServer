package types

type Gkeonprem_BareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig struct {
	/*
	   Whether to enable multiple network interfaces for your pods.
	   When set network_config.advanced_networking is automatically
	   set to true.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
