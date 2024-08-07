package types

type Compute_getRegionNetworkEndpointGroupCloudFunction struct {
	/*
	   A user-defined name of the Cloud Function.
	   The function name is case-sensitive and must be 1-63 characters long.
	   Example value: "func1".
	*/
	Function string `json:"function,omitempty" yaml:"function,omitempty"`

	/*
	   A template to parse function field from a request URL. URL mask allows
	   for routing to multiple Cloud Functions without having to create
	   multiple Network Endpoint Groups and backend services.

	   For example, request URLs "mydomain.com/function1" and "mydomain.com/function2"
	   can be backed by the same Serverless NEG with URL mask "/". The URL mask
	   will parse them to { function = "function1" } and { function = "function2" } respectively.
	*/
	UrlMask string `json:"urlMask,omitempty" yaml:"urlMask,omitempty"`
}
