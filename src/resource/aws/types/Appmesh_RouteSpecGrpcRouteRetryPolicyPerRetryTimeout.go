package types

type Appmesh_RouteSpecGrpcRouteRetryPolicyPerRetryTimeout struct {
	// Retry value.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`

	// Retry unit. Valid values: `ms`, `s`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
