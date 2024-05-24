package types

type Opensearch_getDomainEbsOption struct {
	// The throughput (in MiB/s) of the EBS volumes attached to data nodes.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Size of EBS volumes attached to data nodes (in GB).
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	// Type of EBS volumes attached to data nodes.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Whether EBS volumes are attached to data nodes in the domain.
	EbsEnabled bool `json:"ebsEnabled,omitempty" yaml:"ebsEnabled,omitempty"`

	// Baseline input/output (I/O) performance of EBS volumes attached to data nodes.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`
}
