package types

type Compute_getRegionInstanceTemplateDiskSourceSnapshotEncryptionKey struct {
	/*
	   The service account being used for the encryption
	   request for the given KMS key. If absent, the Compute
	   Engine default service account is used.
	*/
	KmsKeyServiceAccount string `json:"kmsKeyServiceAccount,omitempty" yaml:"kmsKeyServiceAccount,omitempty"`

	// The self link of the encryption key that is stored in Google Cloud KMS
	KmsKeySelfLink string `json:"kmsKeySelfLink,omitempty" yaml:"kmsKeySelfLink,omitempty"`
}
