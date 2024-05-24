package types

type Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings struct {
	// A value used in calculating the threshold below which MediaLive considers a pixel to be 'black'. For the input to be considered black, every pixel in a frame must be below this threshold. The threshold is calculated as a percentage (expressed as a decimal) of white. Therefore .1 means 10%!w(MISSING)hite (or 90%!b(MISSING)lack). Note how the formula works for any color depth. For example, if you set this field to 0.1 in 10-bit color depth: (10230.1=102.3), which means a pixel value of 102 or less is 'black'. If you set this field to .1 in an 8-bit color depth: (2550.1=25.5), which means a pixel value of 25 or less is 'black'. The range is 0.0 to 1.0, with any number of decimal places.
	BlackDetectThreshold float64 `json:"blackDetectThreshold,omitempty" yaml:"blackDetectThreshold,omitempty"`

	// The amount of time (in milliseconds) that the active input must be black before automatic input failover occurs.
	VideoBlackThresholdMsec int `json:"videoBlackThresholdMsec,omitempty" yaml:"videoBlackThresholdMsec,omitempty"`
}
