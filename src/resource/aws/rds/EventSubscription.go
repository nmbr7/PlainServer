package rds

type EventSubscription struct {
	// The SNS topic to send events to.
	SnsTopic string `json:"snsTopic,omitempty" yaml:"snsTopic,omitempty"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	SourceIds []string `json:"sourceIds,omitempty" yaml:"sourceIds,omitempty"`

	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster`, `db-cluster-snapshot`, or `db-proxy`. If not set, all sources will be subscribed to.
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html or run `aws rds describe-event-categories`.
	EventCategories []string `json:"eventCategories,omitempty" yaml:"eventCategories,omitempty"`

	// The name of the DB event subscription. By default generated by this provider.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the DB event subscription. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`
}
