package types

type Eks_FargateProfileSelector struct {
	// Key-value map of Kubernetes labels for selection.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Kubernetes namespace for selection.

	   The following arguments are optional:
	*/
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
