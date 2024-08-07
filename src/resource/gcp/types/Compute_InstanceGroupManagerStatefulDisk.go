package types

type Compute_InstanceGroupManagerStatefulDisk struct {
	// , A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`. `NEVER` - detach the disk when the VM is deleted, but do not delete the disk. `ON_PERMANENT_INSTANCE_DELETION` will delete the stateful disk when the VM is permanently deleted from the instance group. The default is `NEVER`.
	DeleteRule string `json:"deleteRule,omitempty" yaml:"deleteRule,omitempty"`

	// , The device name of the disk to be attached.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`
}
