package types

type Eventarc_TriggerDestinationGke struct {
	// Required. The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// Required. The name of the Google Compute Engine in which the cluster resides, which can either be compute zone (for example, us-central1-a) for the zonal clusters or region (for example, us-central1) for regional clusters.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Required. The namespace the GKE service is running in.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Required. Name of the GKE service.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
