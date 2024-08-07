package types

type Elasticsearch_DomainEncryptAtRest struct {
	// Whether to enable encryption at rest. If the `encrypt_at_rest` block is not provided then this defaults to `false`. Enabling encryption on new domains requires `elasticsearch_version` 5.1 or greater.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// KMS key ARN to encrypt the Elasticsearch domain with. If not specified then it defaults to using the `aws/es` service KMS key. Note that KMS will accept a KMS key ID but will return the key ARN. To prevent the provider detecting unwanted changes, use the key ARN instead.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
