package types

type Wafregional_WebAclLoggingConfiguration struct {
	// Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
	LogDestination string `json:"logDestination,omitempty" yaml:"logDestination,omitempty"`

	// Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
	RedactedFields Wafregional_WebAclLoggingConfigurationRedactedFields `json:"redactedFields,omitempty" yaml:"redactedFields,omitempty"`
}
