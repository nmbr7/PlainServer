package types

type Emr_ClusterCoreInstanceFleetInstanceTypeConfigEbsConfig struct {
	// Volume type. Valid options are `gp3`, `gp2`, `io1`, `io2`, `standard`, `st1` and `sc1`. See [EBS Volume Types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Number of EBS volumes with this configuration to attach to each EC2 instance in the instance group (default is 1).
	VolumesPerInstance int `json:"volumesPerInstance,omitempty" yaml:"volumesPerInstance,omitempty"`

	// Number of I/O operations per second (IOPS) that the volume supports.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Volume size, in gibibytes (GiB).
	Size int `json:"size,omitempty" yaml:"size,omitempty"`
}
