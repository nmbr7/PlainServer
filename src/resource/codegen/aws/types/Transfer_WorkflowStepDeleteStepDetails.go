package types

type Transfer_WorkflowStepDeleteStepDetails struct {
	// The name of the step, used as an identifier.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	SourceFileLocation string `json:"sourceFileLocation,omitempty" yaml:"sourceFileLocation,omitempty"`
}
