package healthcare

import types "libds/gcp/types"

type ConsentStoreIamBinding struct {
	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Healthcare_ConsentStoreIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// Used to find the parent resource to bind the IAM policy to
	ConsentStoreId string `json:"consentStoreId,omitempty" yaml:"consentStoreId,omitempty"`

	/*
	   Identifies the dataset addressed by this request. Must be in the format
	   'projects/{project}/locations/{location}/datasets/{dataset}'
	   Used to find the parent resource to bind the IAM policy to

	   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
	   Each entry can have one of the following values:
	   - --allUsers--: A special identifier that represents anyone who is on the internet; with or without a Google account.
	   - --allAuthenticatedUsers--: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	   - --user:{emailid}--: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	   - --serviceAccount:{emailid}--: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	   - --group:{emailid}--: An email address that represents a Google group. For example, admins@example.com.
	   - --domain:{domain}--: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	   - --projectOwner:projectid--: Owners of the given project. For example, "projectOwner:my-example-project"
	   - --projectEditor:projectid--: Editors of the given project. For example, "projectEditor:my-example-project"
	   - --projectViewer:projectid--: Viewers of the given project. For example, "projectViewer:my-example-project"
	*/
	Dataset string `json:"dataset,omitempty" yaml:"dataset,omitempty"`
}
