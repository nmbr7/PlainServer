package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfiguration struct {
	// Describes an application's checkpointing configuration.
	CheckpointConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfiguration `json:"checkpointConfiguration,omitempty" yaml:"checkpointConfiguration,omitempty"`

	// Describes configuration parameters for CloudWatch logging for an application.
	MonitoringConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration `json:"monitoringConfiguration,omitempty" yaml:"monitoringConfiguration,omitempty"`

	// Describes parameters for how an application executes multiple tasks simultaneously.
	ParallelismConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfiguration `json:"parallelismConfiguration,omitempty" yaml:"parallelismConfiguration,omitempty"`
}
