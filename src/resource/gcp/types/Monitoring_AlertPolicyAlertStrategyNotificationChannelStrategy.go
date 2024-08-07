package types

type Monitoring_AlertPolicyAlertStrategyNotificationChannelStrategy struct {
	/*
	   The notification channels that these settings apply to. Each of these
	   correspond to the name field in one of the NotificationChannel objects
	   referenced in the notification_channels field of this AlertPolicy. The format is
	   `projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]`
	*/
	NotificationChannelNames []string `json:"notificationChannelNames,omitempty" yaml:"notificationChannelNames,omitempty"`

	// The frequency at which to send reminder notifications for open incidents.
	RenotifyInterval string `json:"renotifyInterval,omitempty" yaml:"renotifyInterval,omitempty"`
}
