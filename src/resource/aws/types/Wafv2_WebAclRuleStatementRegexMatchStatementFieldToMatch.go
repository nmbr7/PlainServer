package types

type Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatch struct {
	// Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform.
	Method Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchMethod `json:"method,omitempty" yaml:"method,omitempty"`

	// Inspect a single header. See `single_header` below for details.
	SingleHeader Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchSingleHeader `json:"singleHeader,omitempty" yaml:"singleHeader,omitempty"`

	// Inspect a string containing the list of the request's header names, ordered as they appear in the web request that AWS WAF receives for inspection. See `header_order` below for details.
	HeaderOrders []Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchHeaderOrder `json:"headerOrders,omitempty" yaml:"headerOrders,omitempty"`

	// Inspect the request headers. See `headers` below for details.
	Headers []Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Inspect the JA3 fingerprint. See `ja3_fingerprint` below for details.
	Ja3Fingerprint Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchJa3Fingerprint `json:"ja3Fingerprint,omitempty" yaml:"ja3Fingerprint,omitempty"`

	// Inspect the request body as JSON. See `json_body` for details.
	JsonBody Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchJsonBody `json:"jsonBody,omitempty" yaml:"jsonBody,omitempty"`

	// Inspect the query string. This is the part of a URL that appears after a `?` character, if any.
	QueryString Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchQueryString `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// Inspect a single query argument. See `single_query_argument` below for details.
	SingleQueryArgument Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchSingleQueryArgument `json:"singleQueryArgument,omitempty" yaml:"singleQueryArgument,omitempty"`

	// Inspect the request URI path. This is the part of a web request that identifies a resource, for example, `/images/daily-ad.jpg`.
	UriPath Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchUriPath `json:"uriPath,omitempty" yaml:"uriPath,omitempty"`

	// Inspect all query arguments.
	AllQueryArguments Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchAllQueryArguments `json:"allQueryArguments,omitempty" yaml:"allQueryArguments,omitempty"`

	// Inspect the request body, which immediately follows the request headers. See `body` below for details.
	Body Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchBody `json:"body,omitempty" yaml:"body,omitempty"`

	// Inspect the cookies in the web request. See `cookies` below for details.
	Cookies Wafv2_WebAclRuleStatementRegexMatchStatementFieldToMatchCookies `json:"cookies,omitempty" yaml:"cookies,omitempty"`
}
