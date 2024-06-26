package types

type Bigquery_DataTransferConfigSensitiveParams struct {
	/*
	   The Secret Access Key of the AWS account transferring data from.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	SecretAccessKey string `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`
}
