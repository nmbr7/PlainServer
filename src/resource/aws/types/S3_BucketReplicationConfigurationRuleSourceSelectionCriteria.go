package types

type S3_BucketReplicationConfigurationRuleSourceSelectionCriteria struct {
	/*
	   Match SSE-KMS encrypted objects (documented below). If specified, `replica_kms_key_id`
	   in `destination` must be specified as well.
	*/
	SseKmsEncryptedObjects S3_BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects `json:"sseKmsEncryptedObjects,omitempty" yaml:"sseKmsEncryptedObjects,omitempty"`
}
