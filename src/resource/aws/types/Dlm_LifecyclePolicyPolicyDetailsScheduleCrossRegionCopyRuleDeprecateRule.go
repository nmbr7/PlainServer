package types

type Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleDeprecateRule struct {
	// The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value. Conflicts with `cron_expression`. Must be set if `interval` is set.
	IntervalUnit string `json:"intervalUnit,omitempty" yaml:"intervalUnit,omitempty"`

	// How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values. Conflicts with `cron_expression`. If set, `interval_unit` and `times` must also be set.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`
}
