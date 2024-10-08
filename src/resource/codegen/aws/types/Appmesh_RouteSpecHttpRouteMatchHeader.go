package types

type Appmesh_RouteSpecHttpRouteMatchHeader struct {
	// Name for the HTTP header in the client request that will be matched on.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If `true`, the match is on the opposite of the `match` method and value. Default is `false`.
	Invert bool `json:"invert,omitempty" yaml:"invert,omitempty"`

	// Method and value to match the header value sent with a request. Specify one match method.
	Match Appmesh_RouteSpecHttpRouteMatchHeaderMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
