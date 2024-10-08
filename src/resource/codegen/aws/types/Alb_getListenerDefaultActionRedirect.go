package types

type Alb_getListenerDefaultActionRedirect struct {
	//
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	//
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	//
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Port of the listener. Required if `arn` is not set.
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	Query string `json:"query,omitempty" yaml:"query,omitempty"`
}
