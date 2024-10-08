package types

type Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchJa3Fingerprint struct {
	// The match status to assign to the web request if the request doesn't have a JA3 fingerprint. Valid values include: `MATCH` or `NO_MATCH`.
	FallbackBehavior string `json:"fallbackBehavior,omitempty" yaml:"fallbackBehavior,omitempty"`
}
