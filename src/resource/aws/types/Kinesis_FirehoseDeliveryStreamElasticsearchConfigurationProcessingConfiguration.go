package types

type Kinesis_FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfiguration struct {
	// Enables or disables data processing.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Specifies the data processors as multiple blocks. See `processors` block below for details.
	Processors []Kinesis_FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessor `json:"processors,omitempty" yaml:"processors,omitempty"`
}
