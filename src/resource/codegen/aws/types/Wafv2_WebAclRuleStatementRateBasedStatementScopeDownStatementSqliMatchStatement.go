package types

type Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement struct {
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection. At least one transformation is required. See `text_transformation` below for details.
	TextTransformations []Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`

	// Part of a web request that you want AWS WAF to inspect. See `field_to_match` below for details.
	FieldToMatch Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// Sensitivity that you want AWS WAF to use to inspect for SQL injection attacks. Valid values include: `LOW`, `HIGH`.
	SensitivityLevel string `json:"sensitivityLevel,omitempty" yaml:"sensitivityLevel,omitempty"`
}