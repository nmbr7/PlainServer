package types

type Compute_RegionInstanceTemplateConfidentialInstanceConfig struct {
	// Defines the confidential computing technology the instance uses. SEV is an AMD feature. One of the following values: `SEV`, `SEV_SNP`. `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM. If `SEV_SNP`, currently `min_cpu_platform` has to be set to `"AMD Milan"` or this will fail to create the VM.
	ConfidentialInstanceType string `json:"confidentialInstanceType,omitempty" yaml:"confidentialInstanceType,omitempty"`

	// Defines whether the instance should have confidential compute enabled on AMD SEV. `on_host_maintenance` has to be set to TERMINATE or this will fail to create the VM.
	EnableConfidentialCompute bool `json:"enableConfidentialCompute,omitempty" yaml:"enableConfidentialCompute,omitempty"`
}
