package types

type Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements struct {
	// List of instance generation names. Default is any generation.
	InstanceGenerations []string `json:"instanceGenerations,omitempty" yaml:"instanceGenerations,omitempty"`

	// Block describing the minimum and maximum amount of memory (GiB) per vCPU. Default is no minimum or maximum.
	MemoryGibPerVcpu Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpu `json:"memoryGibPerVcpu,omitempty" yaml:"memoryGibPerVcpu,omitempty"`

	/*
	   Price protection threshold for On-Demand Instances. This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 20.

	   If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is applied based on the per vCPU or per memory price instead of the per instance price.
	*/
	OnDemandMaxPricePercentageOverLowestPrice int `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" yaml:"onDemandMaxPricePercentageOverLowestPrice,omitempty"`

	// Block describing the minimum and maximum number of vCPUs. Default is no maximum.
	VcpuCount Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCount `json:"vcpuCount,omitempty" yaml:"vcpuCount,omitempty"`

	// Block describing the minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips). Default is no minimum or maximum.
	AcceleratorCount Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCount `json:"acceleratorCount,omitempty" yaml:"acceleratorCount,omitempty"`

	// Indicate whether bare metal instace types should be `included`, `excluded`, or `required`. Default is `excluded`.
	BareMetal string `json:"bareMetal,omitempty" yaml:"bareMetal,omitempty"`

	// Indicate whether burstable performance instance types should be `included`, `excluded`, or `required`. Default is `excluded`.
	BurstablePerformance string `json:"burstablePerformance,omitempty" yaml:"burstablePerformance,omitempty"`

	/*
	   List of CPU manufacturer names. Default is any manufacturer.

	   > --NOTE:-- Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	*/
	CpuManufacturers []string `json:"cpuManufacturers,omitempty" yaml:"cpuManufacturers,omitempty"`

	// Indicate whether instance types with local storage volumes are `included`, `excluded`, or `required`. Default is `included`.
	LocalStorage string `json:"localStorage,omitempty" yaml:"localStorage,omitempty"`

	// Indicate whether instance types must support On-Demand Instance Hibernation, either `true` or `false`. Default is `false`.
	RequireHibernateSupport bool `json:"requireHibernateSupport,omitempty" yaml:"requireHibernateSupport,omitempty"`

	// Block describing the minimum and maximum total local storage (GB). Default is no minimum or maximum.
	TotalLocalStorageGb Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGb `json:"totalLocalStorageGb,omitempty" yaml:"totalLocalStorageGb,omitempty"`

	// List of accelerator names. Default is any acclerator.
	AcceleratorNames []string `json:"acceleratorNames,omitempty" yaml:"acceleratorNames,omitempty"`

	// Block describing the minimum and maximum total memory of the accelerators. Default is no minimum or maximum.
	AcceleratorTotalMemoryMib Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMib `json:"acceleratorTotalMemoryMib,omitempty" yaml:"acceleratorTotalMemoryMib,omitempty"`

	/*
	   List of instance types to apply your specified attributes against. All other instance types are ignored, even if they match your specified attributes. You can use strings with one or more wild cards, represented by an asterisk (\-), to allow an instance type, size, or generation. The following are examples: `m5.8xlarge`, `c5-.-`, `m5a.-`, `r-`, `-3-`. For example, if you specify `c5-`, you are allowing the entire C5 instance family, which includes all C5a and C5n instance  If you specify `m5a.-`, you are allowing all the M5a instance types, but not the M5n instance  Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is all instance

	   > --NOTE:-- If you specify `allowed_instance_types`, you can't specify `excluded_instance_types`.
	*/
	AllowedInstanceTypes []string `json:"allowedInstanceTypes,omitempty" yaml:"allowedInstanceTypes,omitempty"`

	// Block describing the minimum and maximum amount of network bandwidth, in gigabits per second (Gbps). Default is no minimum or maximum.
	NetworkBandwidthGbps Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkBandwidthGbps `json:"networkBandwidthGbps,omitempty" yaml:"networkBandwidthGbps,omitempty"`

	// Block describing the minimum and maximum number of network interfaces. Default is no minimum or maximum.
	NetworkInterfaceCount Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCount `json:"networkInterfaceCount,omitempty" yaml:"networkInterfaceCount,omitempty"`

	// List of accelerator  Default is any accelerator type.
	AcceleratorTypes []string `json:"acceleratorTypes,omitempty" yaml:"acceleratorTypes,omitempty"`

	// Block describing the minimum and maximum baseline EBS bandwidth, in Mbps. Default is no minimum or maximum.
	BaselineEbsBandwidthMbps Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps `json:"baselineEbsBandwidthMbps,omitempty" yaml:"baselineEbsBandwidthMbps,omitempty"`

	/*
	   List of instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk (\-), to exclude an instance type, size, or generation. The following are examples: `m5.8xlarge`, `c5-.-`, `m5a.-`, `r-`, `-3-`. For example, if you specify `c5-`, you are excluding the entire C5 instance family, which includes all C5a and C5n instance  If you specify `m5a.-`, you are excluding all the M5a instance types, but not the M5n instance  Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is no excluded instance

	   > --NOTE:-- If you specify `excluded_instance_types`, you can't specify `allowed_instance_types`.
	*/
	ExcludedInstanceTypes []string `json:"excludedInstanceTypes,omitempty" yaml:"excludedInstanceTypes,omitempty"`

	/*
	   Price protection threshold for Spot Instances. This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 100.

	   If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is applied based on the per vCPU or per memory price instead of the per instance price.
	*/
	SpotMaxPricePercentageOverLowestPrice int `json:"spotMaxPricePercentageOverLowestPrice,omitempty" yaml:"spotMaxPricePercentageOverLowestPrice,omitempty"`

	// List of accelerator manufacturer names. Default is any manufacturer.
	AcceleratorManufacturers []string `json:"acceleratorManufacturers,omitempty" yaml:"acceleratorManufacturers,omitempty"`

	// List of local storage type names. Default any storage type.
	LocalStorageTypes []string `json:"localStorageTypes,omitempty" yaml:"localStorageTypes,omitempty"`

	// Block describing the minimum and maximum amount of memory (MiB). Default is no maximum.
	MemoryMib Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib `json:"memoryMib,omitempty" yaml:"memoryMib,omitempty"`
}
