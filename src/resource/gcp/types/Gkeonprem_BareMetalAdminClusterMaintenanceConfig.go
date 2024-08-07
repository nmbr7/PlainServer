package types

type Gkeonprem_BareMetalAdminClusterMaintenanceConfig struct {
	/*
	   All IPv4 address from these ranges will be placed into maintenance mode.
	   Nodes in maintenance mode will be cordoned and drained. When both of these
	   are true, the "baremetal.cluster.gke.io/maintenance" annotation will be set
	   on the node resource.
	*/
	MaintenanceAddressCidrBlocks []string `json:"maintenanceAddressCidrBlocks,omitempty" yaml:"maintenanceAddressCidrBlocks,omitempty"`
}
