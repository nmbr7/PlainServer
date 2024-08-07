package lightsail

type KeyPair struct {
	/*
	   A map of tags to assign to the collection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   > --NOTE:-- a PGP key is not required, however it is strongly encouraged. Without a PGP key, the private key material will be stored in state unencrypted.`pgp_key` is ignored if `public_key` is supplied.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the Lightsail Key Pair. If omitted, a unique name will be generated by this provider
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// An optional PGP key to encrypt the resulting private key material. Only used when creating a new key pair
	PgpKey string `json:"pgpKey,omitempty" yaml:"pgpKey,omitempty"`

	// The public key material. This public key will be imported into Lightsail
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`
}
