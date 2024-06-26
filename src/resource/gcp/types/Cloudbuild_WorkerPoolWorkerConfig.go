package types

type Cloudbuild_WorkerPoolWorkerConfig struct {
	// Size of the disk attached to the worker, in GB. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If `0` is specified, Cloud Build will use a standard disk size.
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	// Machine type of a worker, such as `n1-standard-1`. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use `n1-standard-1`.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// If true, workers are created without any public address, which prevents network egress to public IPs.
	NoExternalIp bool `json:"noExternalIp,omitempty" yaml:"noExternalIp,omitempty"`
}
