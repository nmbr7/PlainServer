package types

type Ec2_NetworkInsightsAnalysisReturnPathComponentInboundHeader struct {
	//
	DestinationAddresses []string `json:"destinationAddresses,omitempty" yaml:"destinationAddresses,omitempty"`

	//
	DestinationPortRanges []Ec2_NetworkInsightsAnalysisReturnPathComponentInboundHeaderDestinationPortRange `json:"destinationPortRanges,omitempty" yaml:"destinationPortRanges,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	SourceAddresses []string `json:"sourceAddresses,omitempty" yaml:"sourceAddresses,omitempty"`

	//
	SourcePortRanges []Ec2_NetworkInsightsAnalysisReturnPathComponentInboundHeaderSourcePortRange `json:"sourcePortRanges,omitempty" yaml:"sourcePortRanges,omitempty"`
}
