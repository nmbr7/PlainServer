package types

type Compute_URLMapPathMatcherRouteRule struct {
	/*
	   For routeRules within a given pathMatcher, priority determines the order
	   in which load balancer will interpret routeRules. RouteRules are evaluated
	   in order of priority, from the lowest to highest number. The priority of
	   a rule decreases as its number increases (1, 2, 3, N+1). The first rule
	   that matches the request is applied.
	   You cannot configure two or more routeRules with the same priority.
	   Priority for each rule must be set to a number between 0 and
	   2147483647 inclusive.
	   Priority numbers can have gaps, which enable you to add or remove rules
	   in the future without affecting the rest of the rules. For example,
	   1, 2, 3, 4, 5, 9, 12, 16 is a valid series of priority numbers to which
	   you could add rules numbered from 6 to 8, 10 to 11, and 13 to 15 in the
	   future without any impact on existing rules.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   In response to a matching matchRule, the load balancer performs advanced routing
	   actions like URL rewrites, header transformations, etc. prior to forwarding the
	   request to the selected backend. If  routeAction specifies any
	   weightedBackendServices, service must not be set. Conversely if service is set,
	   routeAction cannot contain any  weightedBackendServices. Only one of routeAction
	   or urlRedirect must be set.
	   Structure is documented below.
	*/
	RouteAction Compute_URLMapPathMatcherRouteRuleRouteAction `json:"routeAction,omitempty" yaml:"routeAction,omitempty"`

	/*
	   The backend service resource to which traffic is
	   directed if this rule is matched. If routeAction is additionally specified,
	   advanced routing actions like URL Rewrites, etc. take effect prior to sending
	   the request to the backend. However, if service is specified, routeAction cannot
	   contain any weightedBackendService s. Conversely, if routeAction specifies any
	   weightedBackendServices, service must not be specified. Only one of urlRedirect,
	   service or routeAction.weightedBackendService must be set.
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   When this rule is matched, the request is redirected to a URL specified by
	   urlRedirect. If urlRedirect is specified, service or routeAction must not be
	   set.
	   Structure is documented below.
	*/
	UrlRedirect Compute_URLMapPathMatcherRouteRuleUrlRedirect `json:"urlRedirect,omitempty" yaml:"urlRedirect,omitempty"`

	/*
	   Specifies changes to request and response headers that need to take effect for
	   the selected backendService. The headerAction specified here are applied before
	   the matching pathMatchers[].headerAction and after pathMatchers[].routeRules[].r
	   outeAction.weightedBackendService.backendServiceWeightAction[].headerAction
	   Structure is documented below.
	*/
	HeaderAction Compute_URLMapPathMatcherRouteRuleHeaderAction `json:"headerAction,omitempty" yaml:"headerAction,omitempty"`

	/*
	   The rules for determining a match.
	   Structure is documented below.
	*/
	MatchRules []Compute_URLMapPathMatcherRouteRuleMatchRule `json:"matchRules,omitempty" yaml:"matchRules,omitempty"`
}
