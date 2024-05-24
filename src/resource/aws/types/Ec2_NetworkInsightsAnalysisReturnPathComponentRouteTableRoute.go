package types

type Ec2_NetworkInsightsAnalysisReturnPathComponentRouteTableRoute struct {
	//
	NatGatewayId string `json:"natGatewayId,omitempty" yaml:"natGatewayId,omitempty"`

	//
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	//
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	//
	Origin string `json:"origin,omitempty" yaml:"origin,omitempty"`

	//
	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId,omitempty" yaml:"vpcPeeringConnectionId,omitempty"`

	//
	DestinationCidr string `json:"destinationCidr,omitempty" yaml:"destinationCidr,omitempty"`

	//
	DestinationPrefixListId string `json:"destinationPrefixListId,omitempty" yaml:"destinationPrefixListId,omitempty"`

	//
	EgressOnlyInternetGatewayId string `json:"egressOnlyInternetGatewayId,omitempty" yaml:"egressOnlyInternetGatewayId,omitempty"`

	//
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`
}
