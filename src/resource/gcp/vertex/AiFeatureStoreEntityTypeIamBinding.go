package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiFeatureStoreEntityTypeIamBinding struct {
	//
	Condition types.Vertex_AiFeatureStoreEntityTypeIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// Used to find the parent resource to bind the IAM policy to
	Entitytype string `json:"entitytype,omitempty" yaml:"entitytype,omitempty"`

	/*
	   The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to

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
	Featurestore string `json:"featurestore,omitempty" yaml:"featurestore,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.vertex.AiFeatureStoreEntityTypeIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
