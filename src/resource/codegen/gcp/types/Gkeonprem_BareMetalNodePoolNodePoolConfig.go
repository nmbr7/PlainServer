package types

type Gkeonprem_BareMetalNodePoolNodePoolConfig struct {
	/*
	   The map of Kubernetes labels (key/value pairs) to be applied to
	   each node. These will added in addition to any default label(s)
	   that Kubernetes may apply to the node. In case of conflict in
	   label keys, the applied set may differ depending on the Kubernetes
	   version -- it's best to assume the behavior is undefined and
	   conflicts should be avoided. For more information, including usage
	   and the valid values, see:
	   http://kubernetes.io/v1.1/docs/user-guide/labels.html
	   An object containing a list of "key": value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The list of machine addresses in the Bare Metal Node Pool.
	   Structure is documented below.
	*/
	NodeConfigs []Gkeonprem_BareMetalNodePoolNodePoolConfigNodeConfig `json:"nodeConfigs,omitempty" yaml:"nodeConfigs,omitempty"`

	// Specifies the nodes operating system (default: LINUX).
	OperatingSystem string `json:"operatingSystem,omitempty" yaml:"operatingSystem,omitempty"`

	/*
	   The initial taints assigned to nodes of this node pool.
	   Structure is documented below.
	*/
	Taints []Gkeonprem_BareMetalNodePoolNodePoolConfigTaint `json:"taints,omitempty" yaml:"taints,omitempty"`
}
