package types

type Composer_getEnvironmentConfigNodeConfig struct {
	// The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and region/zone. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// The maximum pods per node in the GKE cluster allocated during environment creation. Lowering this value reduces IP address consumption by the Cloud Composer Kubernetes cluster. This value can only be set during environment creation, and only if the environment is VPC-Native. The range of possible values is 8-110, and the default is 32. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`

	// The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and region/zone. The network must belong to the environment's project. If unspecified, the "default" network ID in the environment's project is used. If a Custom Subnet Network is provided, subnetwork must also be provided.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// The list of instance tags applied to all node VMs. Tags are used to identify valid sources or targets for network firewalls. Each tag within the list must comply with RFC1035. Cannot be updated.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// IPv4 cidr range that will be used by Composer internal components.
	ComposerInternalIpv4CidrBlock string `json:"composerInternalIpv4CidrBlock,omitempty" yaml:"composerInternalIpv4CidrBlock,omitempty"`

	// Deploys 'ip-masq-agent' daemon set in the GKE cluster and defines nonMasqueradeCIDRs equals to pod IP range so IP masquerading is used for all destination addresses, except between pods traffic. See: https://cloud.google.com/kubernetes-engine/docs/how-to/ip-masquerade-agent
	EnableIpMasqAgent bool `json:"enableIpMasqAgent,omitempty" yaml:"enableIpMasqAgent,omitempty"`

	// Configuration for controlling how IPs are allocated in the GKE cluster. Cannot be updated.
	IpAllocationPolicies []Composer_getEnvironmentConfigNodeConfigIpAllocationPolicy `json:"ipAllocationPolicies,omitempty" yaml:"ipAllocationPolicies,omitempty"`

	// The Google Cloud Platform Service Account to be used by the node VMs. If a service account is not specified, the "default" Compute Engine service account is used. Cannot be updated. If given, note that the service account must have roles/composer.worker for any GCP resources created under the Cloud Composer Environment.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// The Compute Engine subnetwork to be used for machine communications, specified as a self-link, relative resource name (e.g. "projects/{project}/regions/{region}/subnetworks/{subnetwork}"), or by name. If subnetwork is provided, network must also be provided and the subnetwork must belong to the enclosing environment's project and region.
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	// The Compute Engine zone in which to deploy the VMs running the Apache Airflow software, specified as the zone name or relative resource name (e.g. "projects/{project}/zones/{zone}"). Must belong to the enclosing environment's project and region. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	// PSC (Private Service Connect) Network entry point. Customers can pre-create the Network Attachment and point Cloud Composer environment to use. It is possible to share network attachment among many environments, provided enough IP addresses are available.
	ComposerNetworkAttachment string `json:"composerNetworkAttachment,omitempty" yaml:"composerNetworkAttachment,omitempty"`

	// The disk size in GB used for node VMs. Minimum size is 20GB. If unspecified, defaults to 100GB. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	// The set of Google API scopes to be made available on all node VMs. Cannot be updated. If empty, defaults to ["https://www.googleapis.com/auth/cloud-platform"]. This field is supported for Cloud Composer environments in versions composer-1.-.--airflow--.-.-.
	OauthScopes []string `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`
}
