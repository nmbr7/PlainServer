package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchSingleHeader struct {
	// The name of the query header to inspect. This setting must be provided as lower case characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
