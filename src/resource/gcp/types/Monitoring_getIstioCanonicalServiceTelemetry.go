package types

type Monitoring_getIstioCanonicalServiceTelemetry struct {
	/*
	   The full name of the resource that defines this service.
	   Formatted as described in
	   https://cloud.google.com/apis/design/resource_names.
	*/
	ResourceName string `json:"resourceName,omitempty" yaml:"resourceName,omitempty"`
}
