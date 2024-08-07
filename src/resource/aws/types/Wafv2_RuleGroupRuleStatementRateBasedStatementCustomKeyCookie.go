package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyCookie struct {
	// A friendly name of the rule group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	   At least one required.
	   See Text Transformation below for details.
	*/
	TextTransformations []Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyCookieTextTransformation `json:"textTransformations,omitempty" yaml:"textTransformations,omitempty"`
}
