package types

type S3_BucketLogging struct {
	// The name of the bucket that will receive the log objects.
	TargetBucket string `json:"targetBucket,omitempty" yaml:"targetBucket,omitempty"`

	// To specify a key prefix for log objects.
	TargetPrefix string `json:"targetPrefix,omitempty" yaml:"targetPrefix,omitempty"`
}
