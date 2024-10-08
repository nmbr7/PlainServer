package types

type Cloudrun_getServiceTemplateSpecContainerStartupProbeHttpGet struct {
	// Custom headers to set in the request. HTTP allows repeated headers.
	HttpHeaders []Cloudrun_getServiceTemplateSpecContainerStartupProbeHttpGetHttpHeader `json:"httpHeaders,omitempty" yaml:"httpHeaders,omitempty"`

	// Path to access on the HTTP server. If set, it should not be empty string.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   Port number to access on the container. Number must be in the range 1 to 65535.
	   If not specified, defaults to the same value as container.ports[0].containerPort.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
