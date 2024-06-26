package secretsmanager

type SecretVersion struct {
	// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secret_string is not set. Needs to be encoded to base64.
	SecretBinary string `json:"secretBinary,omitempty" yaml:"secretBinary,omitempty"`

	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId string `json:"secretId,omitempty" yaml:"secretId,omitempty"`

	// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secret_binary is not set.
	SecretString string `json:"secretString,omitempty" yaml:"secretString,omitempty"`

	/*
	   Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.

	   > --NOTE:-- If `version_stages` is configured, you must include the `AWSCURRENT` staging label if this secret version is the only version or if the label is currently present on this secret version, otherwise this provider will show a perpetual difference.
	*/
	VersionStages []string `json:"versionStages,omitempty" yaml:"versionStages,omitempty"`
}
