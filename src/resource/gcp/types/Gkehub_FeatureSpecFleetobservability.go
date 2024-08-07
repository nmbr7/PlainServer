package types

type Gkehub_FeatureSpecFleetobservability struct {
	/*
	   Specified if fleet logging feature is enabled for the entire fleet. If UNSPECIFIED, fleet logging feature is disabled for the entire fleet.
	   Structure is documented below.
	*/
	LoggingConfig Gkehub_FeatureSpecFleetobservabilityLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`
}
