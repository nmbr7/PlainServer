package types

type Appmesh_VirtualGatewaySpecListenerTlsValidation struct {
	// SANs for a virtual gateway's listener's Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames Appmesh_VirtualGatewaySpecListenerTlsValidationSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	// TLS validation context trust.
	Trust Appmesh_VirtualGatewaySpecListenerTlsValidationTrust `json:"trust,omitempty" yaml:"trust,omitempty"`
}
