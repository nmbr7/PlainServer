package types

type Autoscaling_GroupInitialLifecycleHook struct {
	//
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	//
	DefaultResult string `json:"defaultResult,omitempty" yaml:"defaultResult,omitempty"`

	//
	HeartbeatTimeout int `json:"heartbeatTimeout,omitempty" yaml:"heartbeatTimeout,omitempty"`

	//
	LifecycleTransition string `json:"lifecycleTransition,omitempty" yaml:"lifecycleTransition,omitempty"`

	// Name of the Auto Scaling Group. By default generated by Pulumi. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	NotificationMetadata string `json:"notificationMetadata,omitempty" yaml:"notificationMetadata,omitempty"`

	//
	NotificationTargetArn string `json:"notificationTargetArn,omitempty" yaml:"notificationTargetArn,omitempty"`
}
