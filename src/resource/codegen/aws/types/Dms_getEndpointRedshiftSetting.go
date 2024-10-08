package types

type Dms_getEndpointRedshiftSetting struct {
	//
	BucketFolder string `json:"bucketFolder,omitempty" yaml:"bucketFolder,omitempty"`

	//
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	//
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`

	//
	ServerSideEncryptionKmsKeyId string `json:"serverSideEncryptionKmsKeyId,omitempty" yaml:"serverSideEncryptionKmsKeyId,omitempty"`

	//
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`
}
