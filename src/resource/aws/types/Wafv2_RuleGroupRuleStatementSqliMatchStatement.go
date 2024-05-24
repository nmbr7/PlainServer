package types

type Wafv2_RuleGroupRuleStatementSqliMatchStatement struct {
	/*
	   Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	   At least one required.
	   See Text Transformation below for details.
	*/
	TextTransformations []Wafv2_RuleGroupRuleStatementSqliMatchStatementTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`

	// The part of a web request that you want AWS WAF to inspect. See Field to Match below for details.
	FieldToMatch Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`
}
