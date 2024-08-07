package types

type Wafv2_RegexPatternSetRegularExpression struct {
	// The string representing the regular expression, see the AWS WAF [documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-regex-pattern-set-creating.html) for more information.
	RegexString string `json:"regexString,omitempty" yaml:"regexString,omitempty"`
}
