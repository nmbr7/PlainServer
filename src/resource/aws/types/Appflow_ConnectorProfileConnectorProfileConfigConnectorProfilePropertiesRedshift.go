package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift struct {
	// ARN of the IAM role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A name for the associated Amazon S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// The unique ID that's assigned to an Amazon Redshift cluster.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// ARN of the IAM role that permits AppFlow to access the database through Data API.
	DataApiRoleArn string `json:"dataApiRoleArn,omitempty" yaml:"dataApiRoleArn,omitempty"`

	// The name of an Amazon Redshift database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The JDBC URL of the Amazon Redshift cluster.
	DatabaseUrl string `json:"databaseUrl,omitempty" yaml:"databaseUrl,omitempty"`
}
