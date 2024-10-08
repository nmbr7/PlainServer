package dynamodb

type Tag struct {
	// Tag name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Amazon Resource Name (ARN) of the DynamoDB resource to tag.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// Tag value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
