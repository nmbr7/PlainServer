package types

type Gamelift_BuildStorageLocation struct {
	// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Name of your S3 bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Name of the zip file containing your build files.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// A specific version of the file. If not set, the latest version of the file is retrieved.
	ObjectVersion string `json:"objectVersion,omitempty" yaml:"objectVersion,omitempty"`
}
