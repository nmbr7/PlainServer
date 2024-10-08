package ec2

type VpcDhcpOptions struct {
	// List of NETBIOS name servers.
	NetbiosNameServers []string `json:"netbiosNameServers,omitempty" yaml:"netbiosNameServers,omitempty"`

	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType string `json:"netbiosNodeType,omitempty" yaml:"netbiosNodeType,omitempty"`

	// List of NTP servers to configure.
	NtpServers []string `json:"ntpServers,omitempty" yaml:"ntpServers,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers []string `json:"domainNameServers,omitempty" yaml:"domainNameServers,omitempty"`

	// How frequently, in seconds, a running instance with an IPv6 assigned to it goes through DHCPv6 lease renewal. Acceptable values are between 140 and 2147483647 (approximately 68 years). If no value is entered, the default lease time is 140 seconds. If you use long-term addressing for EC2 instances, you can increase the lease time and avoid frequent lease renewal requests. Lease renewal typically occurs when half of the lease time has elapsed.
	Ipv6AddressPreferredLeaseTime string `json:"ipv6AddressPreferredLeaseTime,omitempty" yaml:"ipv6AddressPreferredLeaseTime,omitempty"`
}
