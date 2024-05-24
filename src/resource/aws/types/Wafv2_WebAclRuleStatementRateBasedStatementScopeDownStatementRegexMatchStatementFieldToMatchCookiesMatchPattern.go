package types

type Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPattern struct {
	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`

	//
	ExcludedCookies []string `json:"excludedCookies,omitempty" yaml:"excludedCookies,omitempty"`

	//
	IncludedCookies []string `json:"includedCookies,omitempty" yaml:"includedCookies,omitempty"`
}
