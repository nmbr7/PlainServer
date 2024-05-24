package types

type Codebuild_ProjectSourceBuildStatusConfig struct {
	// Specifies the target url of the build status CodeBuild sends to the source provider. The usage of this parameter depends on the source provider.
	TargetUrl string `json:"targetUrl,omitempty" yaml:"targetUrl,omitempty"`

	// Specifies the context of the build status CodeBuild sends to the source provider. The usage of this parameter depends on the source provider.
	Context string `json:"context,omitempty" yaml:"context,omitempty"`
}
