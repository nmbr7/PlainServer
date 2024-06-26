package types

type Container_AzureNodePoolConfigSshConfig struct {
	// The SSH public key data for VMs managed by Anthos. This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	AuthorizedKey string `json:"authorizedKey,omitempty" yaml:"authorizedKey,omitempty"`
}
