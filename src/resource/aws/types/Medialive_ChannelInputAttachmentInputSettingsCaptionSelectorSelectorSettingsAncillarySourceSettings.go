package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsAncillarySourceSettings struct {
	// Specifies the number (1 to 4) of the captions channel you want to extract from the ancillary captions. If you plan to convert the ancillary captions to another format, complete this field. If you plan to choose Embedded as the captions destination in the output (to pass through all the channels in the ancillary captions), leave this field blank because MediaLive ignores the field.
	SourceAncillaryChannelNumber int `json:"sourceAncillaryChannelNumber,omitempty" yaml:"sourceAncillaryChannelNumber,omitempty"`
}
