package types

type Ec2_NetworkInsightsAnalysisForwardPathComponent struct {
	//
	OutboundHeaders []Ec2_NetworkInsightsAnalysisForwardPathComponentOutboundHeader `json:"outboundHeaders,omitempty" yaml:"outboundHeaders,omitempty"`

	//
	RouteTableRoutes []Ec2_NetworkInsightsAnalysisForwardPathComponentRouteTableRoute `json:"routeTableRoutes,omitempty" yaml:"routeTableRoutes,omitempty"`

	//
	SequenceNumber int `json:"sequenceNumber,omitempty" yaml:"sequenceNumber,omitempty"`

	//
	Vpcs []Ec2_NetworkInsightsAnalysisForwardPathComponentVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`

	//
	AclRules []Ec2_NetworkInsightsAnalysisForwardPathComponentAclRule `json:"aclRules,omitempty" yaml:"aclRules,omitempty"`

	//
	AdditionalDetails []Ec2_NetworkInsightsAnalysisForwardPathComponentAdditionalDetail `json:"additionalDetails,omitempty" yaml:"additionalDetails,omitempty"`

	//
	Components []Ec2_NetworkInsightsAnalysisForwardPathComponentComponent `json:"components,omitempty" yaml:"components,omitempty"`

	//
	DestinationVpcs []Ec2_NetworkInsightsAnalysisForwardPathComponentDestinationVpc `json:"destinationVpcs,omitempty" yaml:"destinationVpcs,omitempty"`

	//
	AttachedTos []Ec2_NetworkInsightsAnalysisForwardPathComponentAttachedTo `json:"attachedTos,omitempty" yaml:"attachedTos,omitempty"`

	//
	SecurityGroupRules []Ec2_NetworkInsightsAnalysisForwardPathComponentSecurityGroupRule `json:"securityGroupRules,omitempty" yaml:"securityGroupRules,omitempty"`

	//
	Subnets []Ec2_NetworkInsightsAnalysisForwardPathComponentSubnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	//
	InboundHeaders []Ec2_NetworkInsightsAnalysisForwardPathComponentInboundHeader `json:"inboundHeaders,omitempty" yaml:"inboundHeaders,omitempty"`

	//
	SourceVpcs []Ec2_NetworkInsightsAnalysisForwardPathComponentSourceVpc `json:"sourceVpcs,omitempty" yaml:"sourceVpcs,omitempty"`

	//
	TransitGatewayRouteTableRoutes []Ec2_NetworkInsightsAnalysisForwardPathComponentTransitGatewayRouteTableRoute `json:"transitGatewayRouteTableRoutes,omitempty" yaml:"transitGatewayRouteTableRoutes,omitempty"`

	//
	TransitGateways []Ec2_NetworkInsightsAnalysisForwardPathComponentTransitGateway `json:"transitGateways,omitempty" yaml:"transitGateways,omitempty"`
}
