package types

type Codepipeline_PipelineTriggerGitConfigurationPullRequestBranches struct {
	// A list of patterns of Git branches that, when a commit is pushed, are to be excluded from starting the pipeline.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	// A list of patterns of Git branches that, when a commit is pushed, are to be included as criteria that starts the pipeline.
	Includes []string `json:"includes,omitempty" yaml:"includes,omitempty"`
}
