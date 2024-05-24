package types

type Pipes_PipeTargetParametersEcsTaskParametersPlacementStrategy struct {
	// The field to apply the placement strategy against. For the spread placement strategy, valid values are instanceId (or host, which has the same effect), or any platform or custom attribute that is applied to a container instance, such as attribute:ecs.availability-zone. For the binpack placement strategy, valid values are cpu and memory. For the random placement strategy, this field is not used. Maximum length of 255.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// The type of placement strategy. The random placement strategy randomly places tasks on available candidates. The spread placement strategy spreads placement across available candidates evenly based on the field parameter. The binpack strategy places tasks on available candidates that have the least available amount of the resource that is specified with the field parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory (but still enough to run the task). Valid Values: random, spread, binpack.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
