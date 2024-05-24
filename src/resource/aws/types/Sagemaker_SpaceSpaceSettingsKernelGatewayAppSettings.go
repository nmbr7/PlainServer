package types

type Sagemaker_SpaceSpaceSettingsKernelGatewayAppSettings struct {
	// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
	LifecycleConfigArns []string `json:"lifecycleConfigArns,omitempty" yaml:"lifecycleConfigArns,omitempty"`

	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	CustomImages []Sagemaker_SpaceSpaceSettingsKernelGatewayAppSettingsCustomImage `json:"customImages,omitempty" yaml:"customImages,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	DefaultResourceSpec Sagemaker_SpaceSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`
}
