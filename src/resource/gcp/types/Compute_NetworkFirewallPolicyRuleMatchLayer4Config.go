package types

type Compute_NetworkFirewallPolicyRuleMatchLayer4Config struct {
	// The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (`tcp`, `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol number.
	IpProtocol string `json:"ipProtocol,omitempty" yaml:"ipProtocol,omitempty"`

	/*
	   An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ``.

	   - - -
	*/
	Ports []string `json:"ports,omitempty" yaml:"ports,omitempty"`
}
