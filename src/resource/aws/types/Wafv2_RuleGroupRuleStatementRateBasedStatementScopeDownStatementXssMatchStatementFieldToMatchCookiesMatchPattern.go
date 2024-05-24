package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchCookiesMatchPattern struct {
	//
	IncludedCookies []string `json:"includedCookies,omitempty" yaml:"includedCookies,omitempty"`

	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchCookiesMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`

	//
	ExcludedCookies []string `json:"excludedCookies,omitempty" yaml:"excludedCookies,omitempty"`
}
