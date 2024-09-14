package types

type Sagemaker_UserProfileUserSettingsJupyterServerAppSettings struct {
	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" yaml:"lifecycleConfigArns,omitempty"`

	// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see Code Repository below.
	CodeRepositories []Sagemaker_UserProfileUserSettingsJupyterServerAppSettingsCodeRepository `json:"codeRepositories,omitempty" yaml:"codeRepositories,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	DefaultResourceSpec Sagemaker_UserProfileUserSettingsJupyterServerAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`
}