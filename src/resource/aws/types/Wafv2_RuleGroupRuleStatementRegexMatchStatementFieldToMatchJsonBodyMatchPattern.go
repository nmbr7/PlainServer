package types

type Wafv2_RuleGroupRuleStatementRegexMatchStatementFieldToMatchJsonBodyMatchPattern struct {
	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_RuleGroupRuleStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`

	//
	IncludedPaths []string `json:"includedPaths,omitempty" yaml:"includedPaths,omitempty"`
}
