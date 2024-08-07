package types

type Sagemaker_UserProfileUserSettingsRSessionAppSettings struct {
	// A list of custom SageMaker images that are configured to run as a KernelGateway app. see Custom Image below.
	CustomImages []Sagemaker_UserProfileUserSettingsRSessionAppSettingsCustomImage `json:"customImages,omitempty" yaml:"customImages,omitempty"`

	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
	DefaultResourceSpec Sagemaker_UserProfileUserSettingsRSessionAppSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`
}
