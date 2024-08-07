package types

type Wafv2_WebAclRuleStatementXssMatchStatementFieldToMatchBody struct {
	// What WAF should do if the body is larger than WAF can inspect. WAF does not support inspecting the entire contents of the body of a web request when the body exceeds 8 KB (8192 bytes). Only the first 8 KB of the request body are forwarded to WAF by the underlying host service. Valid values: `CONTINUE`, `MATCH`, `NO_MATCH`.
	OversizeHandling string `json:"oversizeHandling,omitempty" yaml:"oversizeHandling,omitempty"`
}
