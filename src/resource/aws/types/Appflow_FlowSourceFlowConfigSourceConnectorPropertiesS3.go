package types

type Appflow_FlowSourceFlowConfigSourceConnectorPropertiesS3 struct {
	// Amazon S3 bucket prefix.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// When you use Amazon S3 as the source, the configuration format that you provide the flow input data. See S3 Input Format Config for details.
	S3InputFormatConfig Appflow_FlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfig `json:"s3InputFormatConfig,omitempty" yaml:"s3InputFormatConfig,omitempty"`

	// Name of the Amazon S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
