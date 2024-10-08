package types

type Compute_URLMapPathMatcherRouteRuleMatchRule struct {
	/*
	   For satisfying the matchRule condition, the path of the request
	   must match the wildcard pattern specified in pathTemplateMatch
	   after removing any query parameters and anchor that may be part
	   of the original URL.
	   pathTemplateMatch must be between 1 and 255 characters
	   (inclusive).  The pattern specified by pathTemplateMatch may
	   have at most 5 wildcard operators and at most 5 variable
	   captures in total.
	*/
	PathTemplateMatch string `json:"pathTemplateMatch,omitempty" yaml:"pathTemplateMatch,omitempty"`

	/*
	   For satisfying the matchRule condition, the request's path must begin with the
	   specified prefixMatch. prefixMatch must begin with a /. The value must be
	   between 1 and 1024 characters. Only one of prefixMatch, fullPathMatch or
	   regexMatch must be specified.
	*/
	PrefixMatch string `json:"prefixMatch,omitempty" yaml:"prefixMatch,omitempty"`

	/*
	   Specifies a list of query parameter match criteria, all of which must match
	   corresponding query parameters in the request.
	   Structure is documented below.
	*/
	QueryParameterMatches []Compute_URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatch `json:"queryParameterMatches,omitempty" yaml:"queryParameterMatches,omitempty"`

	/*
	   For satisfying the matchRule condition, the path of the request must satisfy the
	   regular expression specified in regexMatch after removing any query parameters
	   and anchor supplied with the original URL. For regular expression grammar please
	   see en.cppreference.com/w/cpp/regex/ecmascript  Only one of prefixMatch,
	   fullPathMatch or regexMatch must be specified.
	*/
	RegexMatch string `json:"regexMatch,omitempty" yaml:"regexMatch,omitempty"`

	/*
	   For satisfying the matchRule condition, the path of the request must exactly
	   match the value specified in fullPathMatch after removing any query parameters
	   and anchor that may be part of the original URL. FullPathMatch must be between 1
	   and 1024 characters. Only one of prefixMatch, fullPathMatch or regexMatch must
	   be specified.
	*/
	FullPathMatch string `json:"fullPathMatch,omitempty" yaml:"fullPathMatch,omitempty"`

	/*
	   Specifies a list of header match criteria, all of which must match corresponding
	   headers in the request.
	   Structure is documented below.
	*/
	HeaderMatches []Compute_URLMapPathMatcherRouteRuleMatchRuleHeaderMatch `json:"headerMatches,omitempty" yaml:"headerMatches,omitempty"`

	/*
	   Specifies that prefixMatch and fullPathMatch matches are case sensitive.
	   Defaults to false.
	*/
	IgnoreCase bool `json:"ignoreCase,omitempty" yaml:"ignoreCase,omitempty"`

	/*
	   Opaque filter criteria used by Loadbalancer to restrict routing configuration to
	   a limited set xDS compliant clients. In their xDS requests to Loadbalancer, xDS
	   clients present node metadata. If a match takes place, the relevant routing
	   configuration is made available to those proxies. For each metadataFilter in
	   this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the
	   filterLabels must match the corresponding label provided in the metadata. If its
	   filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match
	   with corresponding labels in the provided metadata. metadataFilters specified
	   here can be overrides those specified in ForwardingRule that refers to this
	   UrlMap. metadataFilters only applies to Loadbalancers that have their
	   loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	   Structure is documented below.
	*/
	MetadataFilters []Compute_URLMapPathMatcherRouteRuleMatchRuleMetadataFilter `json:"metadataFilters,omitempty" yaml:"metadataFilters,omitempty"`
}
