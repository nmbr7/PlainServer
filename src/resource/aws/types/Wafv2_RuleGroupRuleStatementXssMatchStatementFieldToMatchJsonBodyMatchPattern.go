package types

type Wafv2_RuleGroupRuleStatementXssMatchStatementFieldToMatchJsonBodyMatchPattern struct {
	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_RuleGroupRuleStatementXssMatchStatementFieldToMatchJsonBodyMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`

	//
	IncludedPaths []string `json:"includedPaths,omitempty" yaml:"includedPaths,omitempty"`
}
