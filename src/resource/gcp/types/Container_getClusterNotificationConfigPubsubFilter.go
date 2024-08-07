package types

type Container_getClusterNotificationConfigPubsubFilter struct {
	// Can be used to filter what notifications are sent. Valid values include include UPGRADE_AVAILABLE_EVENT, UPGRADE_EVENT and SECURITY_BULLETIN_EVENT
	EventTypes []string `json:"eventTypes,omitempty" yaml:"eventTypes,omitempty"`
}
