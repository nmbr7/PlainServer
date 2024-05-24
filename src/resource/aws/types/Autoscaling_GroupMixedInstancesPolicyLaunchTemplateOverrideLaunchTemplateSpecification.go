package types

type Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification struct {
	// Name of the launch template. Conflicts with `launch_template_id`.
	LaunchTemplateName string `json:"launchTemplateName,omitempty" yaml:"launchTemplateName,omitempty"`

	// Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// ID of the launch template. Conflicts with `launch_template_name`.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`
}
