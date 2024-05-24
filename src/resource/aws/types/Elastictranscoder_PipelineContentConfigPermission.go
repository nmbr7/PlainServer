package types

type Elastictranscoder_PipelineContentConfigPermission struct {
	// The permission that you want to give to the AWS user that you specified in `content_config_permissions.grantee`. Valid values are `Read`, `ReadAcp`, `WriteAcp` or `FullControl`.
	Accesses []string `json:"accesses,omitempty" yaml:"accesses,omitempty"`

	// The AWS user or group that you want to have access to transcoded files and playlists.
	Grantee string `json:"grantee,omitempty" yaml:"grantee,omitempty"`

	// Specify the type of value that appears in the `content_config_permissions.grantee` object. Valid values are `Canonical`, `Email` or `Group`.
	GranteeType string `json:"granteeType,omitempty" yaml:"granteeType,omitempty"`
}
