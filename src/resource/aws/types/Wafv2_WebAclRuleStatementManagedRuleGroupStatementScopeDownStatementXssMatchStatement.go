package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementXssMatchStatement struct {
	// Part of a web request that you want AWS WAF to inspect. See `field_to_match` below for details.
	FieldToMatch Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. At least one transformation is required. See `text_transformation` below for details.
	TextTransformations []Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
