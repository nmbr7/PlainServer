package types

type Notebooks_RuntimeVirtualMachineVirtualMachineConfigDataDisk struct {
	/*
	   (Output)
	   Indicates a list of features to enable on the guest operating
	   system. Applicable only for bootable images. To see a list of
	   available features, read `https://cloud.google.com/compute/docs/
	   images/create-delete-deprecate-private-images#guest-os-features`
	   options. ``
	*/
	GuestOsFeatures []string `json:"guestOsFeatures,omitempty" yaml:"guestOsFeatures,omitempty"`

	/*
	   Input only. Specifies the parameters for a new disk that will
	   be created alongside the new instance. Use initialization
	   parameters to create boot disks or local SSDs attached to the
	   new instance. This property is mutually exclusive with the
	   source property; you can only define one or the other, but not
	   both.
	   Structure is documented below.
	*/
	InitializeParams Notebooks_RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams `json:"initializeParams,omitempty" yaml:"initializeParams,omitempty"`

	/*
	   Specifies a valid partial or full URL to an existing
	   Persistent Disk resource.
	*/
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   (Output)
	   Optional. Indicates that this is a boot disk. The virtual
	   machine will use the first partition of the disk for its
	   root filesystem.
	*/
	Boot bool `json:"boot,omitempty" yaml:"boot,omitempty"`

	/*
	   (Output)
	   Optional. Specifies a unique device name of your choice
	   that is reflected into the /dev/disk/by-id/google-- tree
	   of a Linux operating system running within the instance.
	   This name can be used to reference the device for mounting,
	   resizing, and so on, from within the instance.
	   If not specified, the server chooses a default device name
	   to apply to this disk, in the form persistent-disk-x, where
	   x is a number assigned by Google Compute Engine. This field
	   is only applicable for persistent disks.
	*/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	/*
	   (Output)
	   Output only. A zero-based index to this disk, where 0 is
	   reserved for the boot disk. If you have many disks attached
	   to an instance, each disk would have a unique index number.
	*/
	Index int `json:"index,omitempty" yaml:"index,omitempty"`

	/*
	   "Specifies the disk interface to use for attaching this disk,
	   which is either SCSI or NVME. The default is SCSI. Persistent
	   disks must always use SCSI and the request will fail if you attempt
	   to attach a persistent disk in any other format than SCSI. Local SSDs
	   can use either NVME or SCSI. For performance characteristics of SCSI
	   over NVMe, see Local SSD performance. Valid values: - NVME - SCSI".
	*/
	Interface string `json:"interface,omitempty" yaml:"interface,omitempty"`

	/*
	   (Output)
	   Type of the resource. Always compute#attachedDisk for attached
	   disks.
	*/
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	/*
	   (Output)
	   Output only. Any valid publicly visible licenses.
	*/
	Licenses []string `json:"licenses,omitempty" yaml:"licenses,omitempty"`

	/*
	   The mode in which to attach this disk, either READ_WRITE
	   or READ_ONLY. If not specified, the default is to attach
	   the disk in READ_WRITE mode.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   Specifies the type of the disk, either SCRATCH or PERSISTENT.
	   If not specified, the default is PERSISTENT.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   (Output)
	   Optional. Specifies whether the disk will be auto-deleted
	   when the instance is deleted (but not when the disk is
	   detached from the instance).
	*/
	AutoDelete bool `json:"autoDelete,omitempty" yaml:"autoDelete,omitempty"`
}
