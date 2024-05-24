package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse struct {
	// Instructs AWS WAF to allow the web request. See `allow` below for details.
	Allow Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow `json:"allow,omitempty" yaml:"allow,omitempty"`

	// Instructs AWS WAF to block the web request. See `block` below for details.
	Block Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock `json:"block,omitempty" yaml:"block,omitempty"`

	// Instructs AWS WAF to run a Captcha check against the web request. See `captcha` below for details.
	Captcha Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha `json:"captcha,omitempty" yaml:"captcha,omitempty"`

	// Instructs AWS WAF to run a check against the request to verify that the request is coming from a legitimate client session. See `challenge` below for details.
	Challenge Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseChallenge `json:"challenge,omitempty" yaml:"challenge,omitempty"`

	// Instructs AWS WAF to count the web request and allow it. See `count` below for details.
	Count Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount `json:"count,omitempty" yaml:"count,omitempty"`
}
