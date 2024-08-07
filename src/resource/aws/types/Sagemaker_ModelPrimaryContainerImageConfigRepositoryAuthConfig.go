package types

type Sagemaker_ModelPrimaryContainerImageConfigRepositoryAuthConfig struct {
	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see [Create a Lambda function with the console](https://docs.aws.amazon.com/lambda/latest/dg/getting-started-create-function.html) in the _AWS Lambda Developer Guide_.
	RepositoryCredentialsProviderArn string `json:"repositoryCredentialsProviderArn,omitempty" yaml:"repositoryCredentialsProviderArn,omitempty"`
}
