package types

type Ec2_SecurityGroupIngress struct {
	// Description of this ingress rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of Prefix List IDs.
	PrefixListIds []string `json:"prefixListIds,omitempty" yaml:"prefixListIds,omitempty"`

	/*
	   Protocol. If you select a protocol of `-1` (semantically equivalent to `all`, which is not a valid value here), you must specify a `from_port` and `to_port` equal to 0. The supported values are defined in the `IpProtocol` argument on the [IpPermission](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpPermission.html) API reference.

	   The following arguments are optional:

	   > --Note-- Although `cidr_blocks`, `ipv6_cidr_blocks`, `prefix_list_ids`, and `security_groups` are all marked as optional, you _must_ provide one of them in order to configure the source of the traffic.
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Whether the security group itself will be added as a source to this ingress rule.
	Self bool `json:"self,omitempty" yaml:"self,omitempty"`

	// End range port (or ICMP code if protocol is `icmp`).
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// List of CIDR blocks.
	CidrBlocks []string `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`

	// Start port (or ICMP type number if protocol is `icmp` or `icmpv6`).
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// List of IPv6 CIDR blocks.
	Ipv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty" yaml:"ipv6CidrBlocks,omitempty"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
}
