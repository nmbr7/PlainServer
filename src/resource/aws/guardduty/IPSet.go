package guardduty

type IPSet struct {
	// The detector ID of the GuardDuty.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// The URI of the file that contains the IPSet.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The friendly name to identify the IPSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate bool `json:"activate,omitempty" yaml:"activate,omitempty"`
}
