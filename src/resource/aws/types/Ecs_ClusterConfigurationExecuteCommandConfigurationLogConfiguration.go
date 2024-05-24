package types

type Ecs_ClusterConfigurationExecuteCommandConfigurationLogConfiguration struct {
	// Whether or not to enable encryption on the CloudWatch logs. If not specified, encryption will be disabled.
	CloudWatchEncryptionEnabled bool `json:"cloudWatchEncryptionEnabled,omitempty" yaml:"cloudWatchEncryptionEnabled,omitempty"`

	// The name of the CloudWatch log group to send logs to.
	CloudWatchLogGroupName string `json:"cloudWatchLogGroupName,omitempty" yaml:"cloudWatchLogGroupName,omitempty"`

	// Whether or not to enable encryption on the logs sent to S3. If not specified, encryption will be disabled.
	S3BucketEncryptionEnabled bool `json:"s3BucketEncryptionEnabled,omitempty" yaml:"s3BucketEncryptionEnabled,omitempty"`

	// The name of the S3 bucket to send logs to.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// An optional folder in the S3 bucket to place logs in.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`
}
