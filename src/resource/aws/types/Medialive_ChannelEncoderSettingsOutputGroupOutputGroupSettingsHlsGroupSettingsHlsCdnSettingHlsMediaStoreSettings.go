package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsMediaStoreSettings struct {
	// Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
	ConnectionRetryInterval int `json:"connectionRetryInterval,omitempty" yaml:"connectionRetryInterval,omitempty"`

	//
	FilecacheDuration int `json:"filecacheDuration,omitempty" yaml:"filecacheDuration,omitempty"`

	//
	MediaStoreStorageClass string `json:"mediaStoreStorageClass,omitempty" yaml:"mediaStoreStorageClass,omitempty"`

	// Number of retry attempts.
	NumRetries int `json:"numRetries,omitempty" yaml:"numRetries,omitempty"`

	// Number of seconds to wait until a restart is initiated.
	RestartDelay int `json:"restartDelay,omitempty" yaml:"restartDelay,omitempty"`
}
