package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeaderMatchPattern struct {
	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeaderMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`

	// An array of strings that will be used for inspecting headers that do not have a key that matches one of the provided values.
	ExcludedHeaders []string `json:"excludedHeaders,omitempty" yaml:"excludedHeaders,omitempty"`

	// An array of strings that will be used for inspecting headers that have a key that matches one of the provided values.
	IncludedHeaders []string `json:"includedHeaders,omitempty" yaml:"includedHeaders,omitempty"`
}
