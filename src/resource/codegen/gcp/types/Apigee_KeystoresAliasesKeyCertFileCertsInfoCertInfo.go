package types

type Apigee_KeystoresAliasesKeyCertFileCertsInfoCertInfo struct {
	/*
	   (Output)
	   X.509 version.
	*/
	Version int `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   (Output)
	   X.509 basic constraints extension.
	*/
	BasicConstraints string `json:"basicConstraints,omitempty" yaml:"basicConstraints,omitempty"`

	/*
	   (Output)
	   X.509 issuer.
	*/
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	/*
	   (Output)
	   Public key component of the X.509 subject public key info.
	*/
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`

	/*
	   (Output)
	   X.509 serial number.
	*/
	SerialNumber string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`

	/*
	   (Output)
	   X.509 signatureAlgorithm.
	*/
	SigAlgName string `json:"sigAlgName,omitempty" yaml:"sigAlgName,omitempty"`

	/*
	   (Output)
	   X.509 subject.
	*/
	Subject string `json:"subject,omitempty" yaml:"subject,omitempty"`

	/*
	   (Output)
	   X.509 subject alternative names (SANs) extension.
	*/
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	/*
	   (Output)
	   X.509 notBefore validity period in milliseconds since epoch.
	*/
	ValidFrom string `json:"validFrom,omitempty" yaml:"validFrom,omitempty"`

	/*
	   (Output)
	   X.509 notAfter validity period in milliseconds since epoch.
	*/
	ExpiryDate string `json:"expiryDate,omitempty" yaml:"expiryDate,omitempty"`

	/*
	   (Output)
	   Flag that specifies whether the certificate is valid.
	   Flag is set to Yes if the certificate is valid, No if expired, or Not yet if not yet valid.
	*/
	IsValid string `json:"isValid,omitempty" yaml:"isValid,omitempty"`
}
