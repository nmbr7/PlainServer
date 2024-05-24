package types

type Appmesh_GatewayRouteSpecHttp2RouteActionRewritePrefix struct {
	// Value used to replace the incoming route prefix when rewritten.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Default prefix used to replace the incoming route prefix when rewritten. Valid values: `ENABLED`, `DISABLED`.
	DefaultPrefix string `json:"defaultPrefix,omitempty" yaml:"defaultPrefix,omitempty"`
}
