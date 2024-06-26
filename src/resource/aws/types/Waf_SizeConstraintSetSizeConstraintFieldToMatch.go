package types

type Waf_SizeConstraintSetSizeConstraintFieldToMatch struct {
	// Part of the web request that you want AWS WAF to search for a specified string. For example, `HEADER`, `METHOD`, or `BODY`. See the [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html) for all supported values.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// When the `type` is `HEADER`, specify the name of the header that you want to search using the `data` field, for example, `User-Agent` or `Referer`. If the `type` is any other value, you can omit this field.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`
}
