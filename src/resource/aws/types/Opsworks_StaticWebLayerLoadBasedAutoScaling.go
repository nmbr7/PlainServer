package types

type Opsworks_StaticWebLayerLoadBasedAutoScaling struct {
	//
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	//
	Upscaling Opsworks_StaticWebLayerLoadBasedAutoScalingUpscaling `json:"upscaling,omitempty" yaml:"upscaling,omitempty"`

	//
	Downscaling Opsworks_StaticWebLayerLoadBasedAutoScalingDownscaling `json:"downscaling,omitempty" yaml:"downscaling,omitempty"`
}
