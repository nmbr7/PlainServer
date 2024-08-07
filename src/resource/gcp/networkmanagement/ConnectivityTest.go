package networkmanagement

import types "DesignSphere_Server/src/resource/gcp/types"

type ConnectivityTest struct {
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	/*
	   Other projects that may be relevant for reachability analysis.
	   This is applicable to scenarios where a test can cross project
	   boundaries.
	*/
	RelatedProjects []string `json:"relatedProjects,omitempty" yaml:"relatedProjects,omitempty"`

	/*
	   Required. Source specification of the Connectivity Test.
	   You can use a combination of source IP address, virtual machine
	   (VM) instance, or Compute Engine network to uniquely identify the
	   source location.
	   Examples: If the source IP address is an internal IP address within
	   a Google Cloud Virtual Private Cloud (VPC) network, then you must
	   also specify the VPC network. Otherwise, specify the VM instance,
	   which already contains its internal IP address and VPC network
	   information.
	   If the source of the test is within an on-premises network, then
	   you must provide the destination VPC network.
	   If the source endpoint is a Compute Engine VM instance with multiple
	   network interfaces, the instance itself is not sufficient to
	   identify the endpoint. So, you must also specify the source IP
	   address or VPC network.
	   A reachability analysis proceeds even if the source location is
	   ambiguous. However, the test result may include endpoints that
	   you don't intend to test.
	   Structure is documented below.
	*/
	Source types.Networkmanagement_ConnectivityTestSource `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   The user-supplied description of the Connectivity Test.
	   Maximum of 512 characters.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Required. Destination specification of the Connectivity Test.
	   You can use a combination of destination IP address, Compute
	   Engine VM instance, or VPC network to uniquely identify the
	   destination location.
	   Even if the destination IP address is not unique, the source IP
	   location is unique. Usually, the analysis can infer the destination
	   endpoint from route information.
	   If the destination you specify is a VM instance and the instance has
	   multiple network interfaces, then you must also specify either a
	   destination IP address or VPC network to identify the destination
	   interface.
	   A reachability analysis proceeds even if the destination location
	   is ambiguous. However, the result can include endpoints that you
	   don't intend to test.
	   Structure is documented below.
	*/
	Destination types.Networkmanagement_ConnectivityTestDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Unique name for the connectivity test.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
