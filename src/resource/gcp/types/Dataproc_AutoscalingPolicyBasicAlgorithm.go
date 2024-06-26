package types

type Dataproc_AutoscalingPolicyBasicAlgorithm struct {
	/*
	   YARN autoscaling configuration.
	   Structure is documented below.
	*/
	YarnConfig Dataproc_AutoscalingPolicyBasicAlgorithmYarnConfig `json:"yarnConfig,omitempty" yaml:"yarnConfig,omitempty"`

	/*
	   Duration between scaling events. A scaling period starts after the
	   update operation from the previous event has completed.
	   Bounds: [2m, 1d]. Default: 2m.
	*/
	CooldownPeriod string `json:"cooldownPeriod,omitempty" yaml:"cooldownPeriod,omitempty"`
}
