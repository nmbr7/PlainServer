package types

type Wafregional_SizeConstraintSetSizeConstraint struct {
	/*
	   The type of comparison you want to perform.
	   e.g., `EQ`, `NE`, `LT`, `GT`.
	   See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_SizeConstraint.html) for all supported values.
	*/
	ComparisonOperator string `json:"comparisonOperator,omitempty" yaml:"comparisonOperator,omitempty"`

	// Specifies where in a web request to look for the size constraint.
	FieldToMatch Wafregional_SizeConstraintSetSizeConstraintFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	/*
	   The size in bytes that you want to compare against the size of the specified `field_to_match`.
	   Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
	*/
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	/*
	   Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	   If you specify a transformation, AWS WAF performs the transformation on `field_to_match` before inspecting a request for a match.
	   e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	   See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-TextTransformation)
	   for all supported values.
	   --Note:-- if you choose `BODY` as `type`, you must choose `NONE` because CloudFront forwards only the first 8192 bytes for inspection.
	*/
	TextTransformation string `json:"textTransformation,omitempty" yaml:"textTransformation,omitempty"`
}
