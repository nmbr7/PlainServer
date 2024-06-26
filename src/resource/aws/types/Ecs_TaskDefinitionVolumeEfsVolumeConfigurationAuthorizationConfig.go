package types

type Ecs_TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig struct {
	// Access point ID to use. If an access point is specified, the root directory value will be relative to the directory set for the access point. If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
	AccessPointId string `json:"accessPointId,omitempty" yaml:"accessPointId,omitempty"`

	// Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system. If enabled, transit encryption must be enabled in the EFSVolumeConfiguration. Valid values: `ENABLED`, `DISABLED`. If this parameter is omitted, the default value of `DISABLED` is used.
	Iam string `json:"iam,omitempty" yaml:"iam,omitempty"`
}
