package types

type Waf_ByteMatchSetByteMatchTuple struct {
	// The part of a web request that you want to search, such as a specified header or a query string.
	FieldToMatch Waf_ByteMatchSetByteMatchTupleFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	/*
	   Within the portion of a web request that you want to search
	   (for example, in the query string, if any), specify where you want to search.
	   e.g., `CONTAINS`, `CONTAINS_WORD` or `EXACTLY`.
	   See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-PositionalConstraint)
	   for all supported values.
	*/
	PositionalConstraint string `json:"positionalConstraint,omitempty" yaml:"positionalConstraint,omitempty"`

	/*
	   The value that you want to search for within the field specified by `field_to_match`, e.g., `badrefer1`.
	   See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_ByteMatchTuple.html)
	   for all supported values.
	*/
	TargetString string `json:"targetString,omitempty" yaml:"targetString,omitempty"`

	/*
	   Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	   If you specify a transformation, AWS WAF performs the transformation on `target_string` before inspecting a request for a match.
	   e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	   See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
	   for all supported values.
	*/
	TextTransformation string `json:"textTransformation,omitempty" yaml:"textTransformation,omitempty"`
}
