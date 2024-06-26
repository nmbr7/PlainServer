package types

type Cognito_IdentityPoolRoleAttachmentRoleMapping struct {
	// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
	MappingRules []Cognito_IdentityPoolRoleAttachmentRoleMappingMappingRule `json:"mappingRules,omitempty" yaml:"mappingRules,omitempty"`

	// The role mapping type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
	AmbiguousRoleResolution string `json:"ambiguousRoleResolution,omitempty" yaml:"ambiguousRoleResolution,omitempty"`

	// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id". Depends on `cognito_identity_providers` set on `aws.cognito.IdentityPool` resource or a `aws.cognito.IdentityProvider` resource.
	IdentityProvider string `json:"identityProvider,omitempty" yaml:"identityProvider,omitempty"`
}
