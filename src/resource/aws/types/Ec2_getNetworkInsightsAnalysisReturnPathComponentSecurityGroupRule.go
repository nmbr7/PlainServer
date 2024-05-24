package types

type Ec2_getNetworkInsightsAnalysisReturnPathComponentSecurityGroupRule struct {
	//
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	//
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	//
	PortRanges []Ec2_getNetworkInsightsAnalysisReturnPathComponentSecurityGroupRulePortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`

	//
	PrefixListId string `json:"prefixListId,omitempty" yaml:"prefixListId,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	SecurityGroupId string `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`
}
