package networkmanager

type CoreNetworkPolicyAttachment struct {
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument string `json:"policyDocument,omitempty" yaml:"policyDocument,omitempty"`

	// The ID of the core network that a policy will be attached to and made `LIVE`.
	CoreNetworkId string `json:"coreNetworkId,omitempty" yaml:"coreNetworkId,omitempty"`
}
