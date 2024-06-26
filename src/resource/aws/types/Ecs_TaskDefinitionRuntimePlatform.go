package types

type Ecs_TaskDefinitionRuntimePlatform struct {
	// Must be set to either `X86_64` or `ARM64`; see [cpu architecture](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform)
	CpuArchitecture string `json:"cpuArchitecture,omitempty" yaml:"cpuArchitecture,omitempty"`

	// If the `requires_compatibilities` is `FARGATE` this field is required; must be set to a valid option from the [operating system family in the runtime platform](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform) setting
	OperatingSystemFamily string `json:"operatingSystemFamily,omitempty" yaml:"operatingSystemFamily,omitempty"`
}
