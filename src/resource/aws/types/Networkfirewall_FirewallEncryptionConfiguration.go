package types

type Networkfirewall_FirewallEncryptionConfiguration struct {
	// The ID of the customer managed key. You can use any of the [key identifiers](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are `CUSTOMER_KMS` and `AWS_OWNED_KMS_KEY`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
