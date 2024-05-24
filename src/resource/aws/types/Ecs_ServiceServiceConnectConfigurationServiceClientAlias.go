package types

type Ecs_ServiceServiceConnectConfigurationServiceClientAlias struct {
	// The name that you use in the applications of client tasks to connect to this service.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	// The listening port number for the Service Connect proxy. This port is available inside of all of the tasks within the same namespace.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
