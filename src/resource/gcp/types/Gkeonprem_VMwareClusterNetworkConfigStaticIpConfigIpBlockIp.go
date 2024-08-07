package types

type Gkeonprem_VMwareClusterNetworkConfigStaticIpConfigIpBlockIp struct {
	// IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
	Ip string `json:"ip,omitempty" yaml:"ip,omitempty"`

	// Hostname of the machine. VM's name will be used if this field is empty.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
}
