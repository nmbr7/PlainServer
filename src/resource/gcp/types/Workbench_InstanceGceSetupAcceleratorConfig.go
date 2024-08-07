package types

type Workbench_InstanceGceSetupAcceleratorConfig struct {
	// Optional. Count of cores of this accelerator.
	CoreCount string `json:"coreCount,omitempty" yaml:"coreCount,omitempty"`

	/*
	   Optional. Type of this accelerator.
	   Possible values are: `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_A100`, `NVIDIA_A100_80GB`, `NVIDIA_L4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
