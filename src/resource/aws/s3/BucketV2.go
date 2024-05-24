package s3

import types "DesignSphere_Server/src/resource/aws/types"

type BucketV2 struct {
	// Name of the bucket. If omitted, the provider will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html). The name must not be in the format `[bucket_name]--[azid]--x-s3`. Use the `aws.s3.DirectoryBucket` resource to manage S3 Express buckets.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`. Must be lowercase and less than or equal to 37 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	/*
	   Configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). See Object Lock Configuration below for details.
	   The provider wil only perform drift detection if a configuration value is provided.
	   Use the `object_lock_enabled` parameter and the resource `aws.s3.BucketObjectLockConfigurationV2` instead.
	*/
	ObjectLockConfiguration types.S3_BucketV2ObjectLockConfiguration `json:"objectLockConfiguration,omitempty" yaml:"objectLockConfiguration,omitempty"`

	/*
	   Valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing. In this case, please make sure you use the verbose/specific version of the policy. For more information about building AWS IAM policy documents with this provider, see the AWS IAM Policy Document Guide.
	   The provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketPolicy` instead.
	*/
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	/*
	   Configuration of the [S3 bucket website](https://docs.aws.amazon.com/AmazonS3/latest/userguide/WebsiteHosting.html). See Website below for details. The provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketWebsiteConfigurationV2` instead.
	*/
	Websites []types.S3_BucketV2Website `json:"websites,omitempty" yaml:"websites,omitempty"`

	/*
	   Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`. Cannot be used in `cn-north-1` or `us-gov-west-1`. This provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketAccelerateConfigurationV2` instead.
	*/
	AccelerationStatus string `json:"accelerationStatus,omitempty" yaml:"accelerationStatus,omitempty"`

	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Valid values are `private`, `public-read`, `public-read-write`, `aws-exec-read`, `authenticated-read`, and `log-delivery-write`. Defaults to `private`.  Conflicts with `grant`. The provider will only perform drift detection if a configuration value is provided. Use the resource `aws.s3.BucketAclV2` instead.
	Acl string `json:"acl,omitempty" yaml:"acl,omitempty"`

	// Rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html). See CORS rule below for details. This provider will only perform drift detection if a configuration value is provided. Use the resource `aws.s3.BucketCorsConfigurationV2` instead.
	CorsRules []types.S3_BucketV2CorsRule `json:"corsRules,omitempty" yaml:"corsRules,omitempty"`

	/*
	   Configuration of [S3 bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) parameters. See Logging below for details. The provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketLoggingV2` instead.
	*/
	Loggings []types.S3_BucketV2Logging `json:"loggings,omitempty" yaml:"loggings,omitempty"`

	/*
	   Specifies who should bear the cost of Amazon S3 data transfer.
	   Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur the costs of any data transfer.
	   See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html) developer guide for more information.
	   The provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketRequestPaymentConfigurationV2` instead.
	*/
	RequestPayer string `json:"requestPayer,omitempty" yaml:"requestPayer,omitempty"`

	// Boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket -when the bucket is destroyed- so that the bucket can be destroyed without error. These objects are -not- recoverable. This only deletes objects when the bucket is destroyed, -not- when setting this parameter to `true`. Once this parameter is set to `true`, there must be a successful `pulumi up` run before a destroy is required to update this value in the resource state. Without a successful `pulumi up` after this parameter is set, this flag will have no effect. If setting this field in the same operation that would require replacing the bucket or destroying the bucket, this flag will not work. Additionally when importing a bucket, a successful `pulumi up` is required to set this value in state before it will take effect on a destroy operation.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	/*
	   Configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html). See Lifecycle Rule below for details. The provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketLifecycleConfigurationV2` instead.
	*/
	LifecycleRules []types.S3_BucketV2LifecycleRule `json:"lifecycleRules,omitempty" yaml:"lifecycleRules,omitempty"`

	/*
	   Configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html). See Replication Configuration below for details. The provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketReplicationConfig` instead.
	*/
	ReplicationConfigurations []types.S3_BucketV2ReplicationConfiguration `json:"replicationConfigurations,omitempty" yaml:"replicationConfigurations,omitempty"`

	/*
	   Map of tags to assign to the bucket. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   The following arguments are deprecated, and will be removed in a future major version:
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl). See Grant below for details. Conflicts with `acl`. The provider will only perform drift detection if a configuration value is provided. Use the resource `aws.s3.BucketAclV2` instead.
	Grants []types.S3_BucketV2Grant `json:"grants,omitempty" yaml:"grants,omitempty"`

	// Indicates whether this bucket has an Object Lock configuration enabled. Valid values are `true` or `false`. This argument is not supported in all regions or partitions.
	ObjectLockEnabled bool `json:"objectLockEnabled,omitempty" yaml:"objectLockEnabled,omitempty"`

	/*
	   Configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html). See Server Side Encryption Configuration below for details.
	   The provider will only perform drift detection if a configuration value is provided.
	   Use the resource `aws.s3.BucketServerSideEncryptionConfigurationV2` instead.
	*/
	ServerSideEncryptionConfigurations []types.S3_BucketV2ServerSideEncryptionConfiguration `json:"serverSideEncryptionConfigurations,omitempty" yaml:"serverSideEncryptionConfigurations,omitempty"`

	// Configuration of the [S3 bucket versioning state](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html). See Versioning below for details. The provider will only perform drift detection if a configuration value is provided. Use the resource `aws.s3.BucketVersioningV2` instead.
	Versionings []types.S3_BucketV2Versioning `json:"versionings,omitempty" yaml:"versionings,omitempty"`
}
