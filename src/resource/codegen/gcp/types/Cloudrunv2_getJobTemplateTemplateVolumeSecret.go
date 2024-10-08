package types

type Cloudrunv2_getJobTemplateTemplateVolumeSecret struct {
	// Integer representation of mode bits to use on created files by default. Must be a value between 0000 and 0777 (octal), defaulting to 0444. Directories within the path are not affected by this setting.
	DefaultMode int `json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`

	// If unspecified, the volume will expose a file whose name is the secret, relative to VolumeMount.mount_path. If specified, the key will be used as the version to fetch from Cloud Secret Manager and the path will be the name of the file exposed in the volume. When items are defined, they must specify a path and a version.
	Items []Cloudrunv2_getJobTemplateTemplateVolumeSecretItem `json:"items,omitempty" yaml:"items,omitempty"`

	// The name of the secret in Cloud Secret Manager. Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project.
	Secret string `json:"secret,omitempty" yaml:"secret,omitempty"`
}
