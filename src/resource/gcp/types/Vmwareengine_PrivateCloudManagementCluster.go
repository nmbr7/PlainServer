package types

type Vmwareengine_PrivateCloudManagementCluster struct {
	/*
	   The map of cluster node types in this cluster,
	   where the key is canonical identifier of the node type (corresponds to the NodeType).
	   Structure is documented below.
	*/
	NodeTypeConfigs []Vmwareengine_PrivateCloudManagementClusterNodeTypeConfig `json:"nodeTypeConfigs,omitempty" yaml:"nodeTypeConfigs,omitempty"`

	/*
	   The user-provided identifier of the new Cluster. The identifier must meet the following requirements:
	   - Only contains 1-63 alphanumeric characters and hyphens
	   - Begins with an alphabetical character
	   - Ends with a non-hyphen character
	   - Not formatted as a UUID
	   - Complies with RFC 1034 (https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	*/
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
}
