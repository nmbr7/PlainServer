package types

type Orgpolicy_PolicyDryRunSpec struct {
	// In policies for boolean constraints, the following requirements apply: - There must be one and only one policy rule where condition is unset. - Boolean policy rules with conditions must set `enforced` to the opposite of the policy rule without a condition. - During policy evaluation, policy rules with conditions that are true for a target resource take precedence.
	Rules []Orgpolicy_PolicyDryRunSpecRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Output only. The time stamp this was previously updated. This represents the last time a call to `CreatePolicy` or `UpdatePolicy` was made for that policy.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	// An opaque tag indicating the current version of the policy, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current policy to use when executing a read-modify-write loop. When the policy is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
	Etag string `json:"etag,omitempty" yaml:"etag,omitempty"`

	// Determines the inheritance behavior for this policy. If `inherit_from_parent` is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
	InheritFromParent bool `json:"inheritFromParent,omitempty" yaml:"inheritFromParent,omitempty"`

	// Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
	Reset bool `json:"reset,omitempty" yaml:"reset,omitempty"`
}
