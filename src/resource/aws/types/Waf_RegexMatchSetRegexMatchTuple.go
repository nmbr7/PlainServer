package types

type Waf_RegexMatchSetRegexMatchTuple struct {
	/*
	   Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	   e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	   See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
	   for all supported values.
	*/
	TextTransformation string `json:"textTransformation,omitempty" yaml:"textTransformation,omitempty"`

	// The part of a web request that you want to search, such as a specified header or a query string.
	FieldToMatch Waf_RegexMatchSetRegexMatchTupleFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// The ID of a Regex Pattern Set.
	RegexPatternSetId string `json:"regexPatternSetId,omitempty" yaml:"regexPatternSetId,omitempty"`
}
