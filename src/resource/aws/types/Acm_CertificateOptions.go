package types

type Acm_CertificateOptions struct {
	// Whether certificate details should be added to a certificate transparency log. Valid values are `ENABLED` or `DISABLED`. See https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency for more details.
	CertificateTransparencyLoggingPreference string `json:"certificateTransparencyLoggingPreference,omitempty" yaml:"certificateTransparencyLoggingPreference,omitempty"`
}
