package types

type Compute_InstanceGroupManagerStatefulExternalIp struct {
	// , A value that prescribes what should happen to the external ip when the VM instance is deleted. The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`. `NEVER` - detach the ip when the VM is deleted, but do not delete the ip. `ON_PERMANENT_INSTANCE_DELETION` will delete the external ip when the VM is permanently deleted from the instance group.
	DeleteRule string `json:"deleteRule,omitempty" yaml:"deleteRule,omitempty"`

	// , The network interface name of the external Ip. Possible value: `nic0`
	InterfaceName string `json:"interfaceName,omitempty" yaml:"interfaceName,omitempty"`
}
