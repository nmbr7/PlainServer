package types

type Ec2_LaunchTemplateElasticGpuSpecification struct {
	// The [Elastic GPU Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-graphics.html#elastic-graphics-basics)
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}