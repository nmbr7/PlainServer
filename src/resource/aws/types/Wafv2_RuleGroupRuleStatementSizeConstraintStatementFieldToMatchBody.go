package types

type Wafv2_RuleGroupRuleStatementSizeConstraintStatementFieldToMatchBody struct {
	// What AWS WAF should do if the cookies of the request are larger than AWS WAF can inspect. AWS WAF does not support inspecting the entire contents of request cookies when they exceed 8 KB (8192 bytes) or 200 total cookies. The underlying host service forwards a maximum of 200 cookies and at most 8 KB of cookie contents to AWS WAF. Valid values: `CONTINUE`, `MATCH`, `NO_MATCH`
	OversizeHandling string `json:"oversizeHandling,omitempty" yaml:"oversizeHandling,omitempty"`
}
