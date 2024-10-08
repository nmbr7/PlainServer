package types

type Cloudfunctionsv2_getFunctionServiceConfig struct {
	/*
	   The limit on the maximum number of function instances that may coexist at a
	   given time.
	*/
	MaxInstanceCount int `json:"maxInstanceCount,omitempty" yaml:"maxInstanceCount,omitempty"`

	// The email of the service account for this function.
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	/*
	   The function execution timeout. Execution is considered failed and
	   can be terminated if the function is not completed at the end of the
	   timeout period. Defaults to 60 seconds.
	*/
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`

	// URI of the Service deployed.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	// Available egress settings. Possible values: ["VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED", "PRIVATE_RANGES_ONLY", "ALL_TRAFFIC"]
	VpcConnectorEgressSettings string `json:"vpcConnectorEgressSettings,omitempty" yaml:"vpcConnectorEgressSettings,omitempty"`

	// Whether 100%!!(MISSING)o(MISSING)f traffic is routed to the latest revision. Defaults to true.
	AllTrafficOnLatestRevision bool `json:"allTrafficOnLatestRevision,omitempty" yaml:"allTrafficOnLatestRevision,omitempty"`

	// The number of CPUs used in a single container instance. Default value is calculated from available memory.
	AvailableCpu string `json:"availableCpu,omitempty" yaml:"availableCpu,omitempty"`

	/*
	   The amount of memory available for a function.
	   Defaults to 256M. Supported units are k, M, G, Mi, Gi. If no unit is
	   supplied the value is interpreted as bytes.
	*/
	AvailableMemory string `json:"availableMemory,omitempty" yaml:"availableMemory,omitempty"`

	// Environment variables that shall be available during function execution.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	// Sets the maximum number of concurrent requests that each instance can receive. Defaults to 1.
	MaxInstanceRequestConcurrency int `json:"maxInstanceRequestConcurrency,omitempty" yaml:"maxInstanceRequestConcurrency,omitempty"`

	// Secret volumes configuration.
	SecretVolumes []Cloudfunctionsv2_getFunctionServiceConfigSecretVolume `json:"secretVolumes,omitempty" yaml:"secretVolumes,omitempty"`

	// Name of the service associated with a Function.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// Available ingress settings. Defaults to "ALLOW_ALL" if unspecified. Default value: "ALLOW_ALL" Possible values: ["ALLOW_ALL", "ALLOW_INTERNAL_ONLY", "ALLOW_INTERNAL_AND_GCLB"]
	IngressSettings string `json:"ingressSettings,omitempty" yaml:"ingressSettings,omitempty"`

	// The Serverless VPC Access connector that this cloud function can connect to.
	VpcConnector string `json:"vpcConnector,omitempty" yaml:"vpcConnector,omitempty"`

	// URIs of the Service deployed
	GcfUri string `json:"gcfUri,omitempty" yaml:"gcfUri,omitempty"`

	/*
	   The limit on the minimum number of function instances that may coexist at a
	   given time.
	*/
	MinInstanceCount int `json:"minInstanceCount,omitempty" yaml:"minInstanceCount,omitempty"`

	// Secret environment variables configuration.
	SecretEnvironmentVariables []Cloudfunctionsv2_getFunctionServiceConfigSecretEnvironmentVariable `json:"secretEnvironmentVariables,omitempty" yaml:"secretEnvironmentVariables,omitempty"`
}
