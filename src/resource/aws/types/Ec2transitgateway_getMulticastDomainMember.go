package types

type Ec2transitgateway_getMulticastDomainMember struct {
	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress string `json:"groupIpAddress,omitempty" yaml:"groupIpAddress,omitempty"`

	// The group members' network interface ID.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`
}
