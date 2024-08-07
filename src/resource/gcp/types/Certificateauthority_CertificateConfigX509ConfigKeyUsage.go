package types

type Certificateauthority_CertificateConfigX509ConfigKeyUsage struct {
	/*
	   Describes high-level ways in which a key may be used.
	   Structure is documented below.
	*/
	BaseKeyUsage Certificateauthority_CertificateConfigX509ConfigKeyUsageBaseKeyUsage `json:"baseKeyUsage,omitempty" yaml:"baseKeyUsage,omitempty"`

	/*
	   Describes high-level ways in which a key may be used.
	   Structure is documented below.
	*/
	ExtendedKeyUsage Certificateauthority_CertificateConfigX509ConfigKeyUsageExtendedKeyUsage `json:"extendedKeyUsage,omitempty" yaml:"extendedKeyUsage,omitempty"`

	/*
	   An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	   Structure is documented below.
	*/
	UnknownExtendedKeyUsages []Certificateauthority_CertificateConfigX509ConfigKeyUsageUnknownExtendedKeyUsage `json:"unknownExtendedKeyUsages,omitempty" yaml:"unknownExtendedKeyUsages,omitempty"`
}
