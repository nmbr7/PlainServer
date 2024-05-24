package types

type Vmwareengine_getPrivateCloudManagementCluster struct {
	/*
	   The user-provided identifier of the new Cluster. The identifier must meet the following requirements:
	     - Only contains 1-63 alphanumeric characters and hyphens
	     - Begins with an alphabetical character
	     - Ends with a non-hyphen character
	     - Not formatted as a UUID
	     - Complies with RFC 1034 (https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	*/
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	/*
	   The map of cluster node types in this cluster,
	   where the key is canonical identifier of the node type (corresponds to the NodeType).
	*/
	NodeTypeConfigs []Vmwareengine_getPrivateCloudManagementClusterNodeTypeConfig `json:"nodeTypeConfigs,omitempty" yaml:"nodeTypeConfigs,omitempty"`
}