package types

type Ec2_LaunchTemplateMaintenanceOptions struct {
	// Disables the automatic recovery behavior of your instance or sets it to default. Can be `"default"` or `"disabled"`. See [Recover your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-recover.html) for more details.
	AutoRecovery string `json:"autoRecovery,omitempty" yaml:"autoRecovery,omitempty"`
}
