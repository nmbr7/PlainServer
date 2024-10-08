package types

type Bigquery_JobLoadDestinationEncryptionConfiguration struct {
	/*
	   (Output)
	   Describes the Cloud KMS encryption key version used to protect destination BigQuery table.
	*/
	KmsKeyVersion string `json:"kmsKeyVersion,omitempty" yaml:"kmsKeyVersion,omitempty"`

	/*
	   Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.
	   The BigQuery Service Account associated with your project requires access to this encryption key.
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
