package types

type Lb_getTargetGroupStickiness struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	CookieDuration int `json:"cookieDuration,omitempty" yaml:"cookieDuration,omitempty"`

	//
	CookieName string `json:"cookieName,omitempty" yaml:"cookieName,omitempty"`
}
