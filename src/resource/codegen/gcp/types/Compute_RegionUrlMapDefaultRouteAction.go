package types

type Compute_RegionUrlMapDefaultRouteAction struct {
	/*
	   The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure.
	   As part of fault injection, when clients send requests to a backend service, delays can be introduced by a load balancer on a percentage of requests before sending those requests to the backend service.
	   Similarly requests from clients can be aborted by the load balancer for a percentage of requests.
	   timeout and retryPolicy is ignored by clients that are configured with a faultInjectionPolicy if: 1. The traffic is generated by fault injection AND 2. The fault injection is not a delay fault injection.
	   Fault injection is not supported with the global external HTTP(S) load balancer (classic). To see which load balancers support fault injection, see Load balancing: [Routing and traffic management features](https://cloud.google.com/load-balancing/docs/features#routing-traffic-management).
	   Structure is documented below.
	*/
	FaultInjectionPolicy Compute_RegionUrlMapDefaultRouteActionFaultInjectionPolicy `json:"faultInjectionPolicy,omitempty" yaml:"faultInjectionPolicy,omitempty"`

	/*
	   Specifies the policy on how requests intended for the route's backends are shadowed to a separate mirrored backend service.
	   The load balancer does not wait for responses from the shadow service. Before sending traffic to the shadow service, the host / authority header is suffixed with -shadow.
	   Not supported when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	   Structure is documented below.
	*/
	RequestMirrorPolicy Compute_RegionUrlMapDefaultRouteActionRequestMirrorPolicy `json:"requestMirrorPolicy,omitempty" yaml:"requestMirrorPolicy,omitempty"`

	/*
	   Specifies the retry policy associated with this route.
	   Structure is documented below.
	*/
	RetryPolicy Compute_RegionUrlMapDefaultRouteActionRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	/*
	   Specifies the timeout for the selected route. Timeout is computed from the time the request has been fully processed (known as end-of-stream) up until the response has been processed. Timeout includes all retries.
	   If not specified, this field uses the largest timeout among all backend services associated with the route.
	   Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	   Structure is documented below.
	*/
	Timeout Compute_RegionUrlMapDefaultRouteActionTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	/*
	   The spec to modify the URL of the request, before forwarding the request to the matched service.
	   urlRewrite is the only action supported in UrlMaps for external HTTP(S) load balancers.
	   Not supported when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	   Structure is documented below.
	*/
	UrlRewrite Compute_RegionUrlMapDefaultRouteActionUrlRewrite `json:"urlRewrite,omitempty" yaml:"urlRewrite,omitempty"`

	/*
	   A list of weighted backend services to send traffic to when a route match occurs. The weights determine the fraction of traffic that flows to their corresponding backend service. If all traffic needs to go to a single backend service, there must be one weightedBackendService with weight set to a non-zero number.
	   After a backend service is identified and before forwarding the request to the backend service, advanced routing actions such as URL rewrites and header transformations are applied depending on additional settings specified in this HttpRouteAction.
	   Structure is documented below.
	*/
	WeightedBackendServices []Compute_RegionUrlMapDefaultRouteActionWeightedBackendService `json:"weightedBackendServices,omitempty" yaml:"weightedBackendServices,omitempty"`

	/*
	   The specification for allowing client side cross-origin requests. Please see
	   [W3C Recommendation for Cross Origin Resource Sharing](https://www.w3.org/TR/cors/)
	   Structure is documented below.
	*/
	CorsPolicy Compute_RegionUrlMapDefaultRouteActionCorsPolicy `json:"corsPolicy,omitempty" yaml:"corsPolicy,omitempty"`
}