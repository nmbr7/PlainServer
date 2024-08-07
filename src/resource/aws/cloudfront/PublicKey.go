package cloudfront

type PublicKey struct {
	// An optional comment about the public key.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey string `json:"encodedKey,omitempty" yaml:"encodedKey,omitempty"`

	// The name for the public key. By default generated by this provider.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The name for the public key. Conflicts with `name`.

	   --NOTE:-- When setting `encoded_key` value, there needs a newline at the end of string. Otherwise, multiple runs of pulumi will want to recreate the `aws.cloudfront.PublicKey` resource.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`
}
