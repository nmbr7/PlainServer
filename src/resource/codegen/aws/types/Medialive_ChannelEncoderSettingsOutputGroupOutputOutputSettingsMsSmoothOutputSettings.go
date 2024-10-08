package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsMsSmoothOutputSettings struct {
	//
	H265PackagingType string `json:"h265PackagingType,omitempty" yaml:"h265PackagingType,omitempty"`

	// String concatenated to the end of the destination filename. Required for multiple outputs of the same type.
	NameModifier string `json:"nameModifier,omitempty" yaml:"nameModifier,omitempty"`
}
