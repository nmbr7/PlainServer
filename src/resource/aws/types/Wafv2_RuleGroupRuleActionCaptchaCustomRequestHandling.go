package types

type Wafv2_RuleGroupRuleActionCaptchaCustomRequestHandling struct {
	// The `insert_header` blocks used to define HTTP headers added to the request. See Custom HTTP Header below for details.
	InsertHeaders []Wafv2_RuleGroupRuleActionCaptchaCustomRequestHandlingInsertHeader `json:"insertHeaders,omitempty" yaml:"insertHeaders,omitempty"`
}
