package kms

import types "libds/gcp/types"

type KeyRingIAMMember struct {
	/*
	   An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	   Structure is documented below.
	*/
	Condition types.Kms_KeyRingIAMMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	/*
	   The key ring ID, in the form
	   `{project_id}/{location_name}/{key_ring_name}` or
	   `{location_name}/{key_ring_name}`. In the second form, the provider's
	   project setting will be used as a fallback.

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	*/
	KeyRingId string `json:"keyRingId,omitempty" yaml:"keyRingId,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
