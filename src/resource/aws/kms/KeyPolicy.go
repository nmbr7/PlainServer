package kms

type KeyPolicy struct {
	/*
	   A flag to indicate whether to bypass the key policy lockout safety check.
	   Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately. If this value is set, and the resource is destroyed, a warning will be shown, and the resource will be removed from state.
	   For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	*/
	BypassPolicyLockoutSafetyCheck bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" yaml:"bypassPolicyLockoutSafetyCheck,omitempty"`

	// The ID of the KMS Key to attach the policy.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	/*
	   A valid policy JSON document. Although this is a key policy, not an IAM policy, an `aws.iam.getPolicyDocument`, in the form that designates a principal, can be used. For more information about building policy documents, see the AWS IAM Policy Document Guide.

	   > --NOTE:-- Note: All KMS keys must have a key policy. If a key policy is not specified, or this resource is destroyed, AWS gives the KMS key a [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) that gives all principals in the owning account unlimited access to all KMS operations for the key. This default key policy effectively delegates all access control to IAM policies and KMS grants.
	*/
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
