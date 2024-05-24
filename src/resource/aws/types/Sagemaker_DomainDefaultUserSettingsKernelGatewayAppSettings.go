package types

type Sagemaker_DomainDefaultUserSettingsKernelGatewayAppSettings struct {
	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see `custom_image` Block below.
	CustomImages []Sagemaker_DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImage `json:"customImages,omitempty" yaml:"customImages,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see `default_resource_spec` Block below.
	DefaultResourceSpec Sagemaker_DomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`

	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" yaml:"lifecycleConfigArns,omitempty"`
}
