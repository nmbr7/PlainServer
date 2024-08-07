package types

type Cloudrunv2_ServiceTemplateContainerStartupProbeTcpSocket struct {
	/*
	   Port number to access on the container. Must be in the range 1 to 65535.
	   If not specified, defaults to the same value as container.ports[0].containerPort.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
