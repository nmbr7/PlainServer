package types

type Pipes_PipeTargetParametersBatchJobParametersDependsOn struct {
	// The type of placement strategy. The random placement strategy randomly places tasks on available candidates. The spread placement strategy spreads placement across available candidates evenly based on the field parameter. The binpack strategy places tasks on available candidates that have the least available amount of the resource that is specified with the field parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory (but still enough to run the task). Valid Values: random, spread, binpack.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The job ID of the AWS Batch job that's associated with this dependency.
	JobId string `json:"jobId,omitempty" yaml:"jobId,omitempty"`
}
