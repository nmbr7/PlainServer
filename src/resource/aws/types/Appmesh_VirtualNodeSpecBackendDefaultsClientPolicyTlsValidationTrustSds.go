package types

type Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds struct {
	// Name of the secret for a virtual node's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	SecretName string `json:"secretName,omitempty" yaml:"secretName,omitempty"`
}
