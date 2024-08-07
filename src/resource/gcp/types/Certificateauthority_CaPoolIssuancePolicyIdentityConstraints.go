package types

type Certificateauthority_CaPoolIssuancePolicyIdentityConstraints struct {
	/*
	   If this is set, the SubjectAltNames extension may be copied from a certificate request into the signed certificate.
	   Otherwise, the requested SubjectAltNames will be discarded.
	*/
	AllowSubjectAltNamesPassthrough bool `json:"allowSubjectAltNamesPassthrough,omitempty" yaml:"allowSubjectAltNamesPassthrough,omitempty"`

	/*
	   If this is set, the Subject field may be copied from a certificate request into the signed certificate.
	   Otherwise, the requested Subject will be discarded.
	*/
	AllowSubjectPassthrough bool `json:"allowSubjectPassthrough,omitempty" yaml:"allowSubjectPassthrough,omitempty"`

	/*
	   A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a
	   certificate is signed. To see the full allowed syntax and some examples,
	   see https://cloud.google.com/certificate-authority-service/docs/cel-guide
	   Structure is documented below.
	*/
	CelExpression Certificateauthority_CaPoolIssuancePolicyIdentityConstraintsCelExpression `json:"celExpression,omitempty" yaml:"celExpression,omitempty"`
}
