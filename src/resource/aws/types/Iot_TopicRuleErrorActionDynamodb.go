package types

type Iot_TopicRuleErrorActionDynamodb struct {
	// The name of the DynamoDB table.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// The hash key value.
	HashKeyValue string `json:"hashKeyValue,omitempty" yaml:"hashKeyValue,omitempty"`

	// The action payload.
	PayloadField string `json:"payloadField,omitempty" yaml:"payloadField,omitempty"`

	// The range key value.
	RangeKeyValue string `json:"rangeKeyValue,omitempty" yaml:"rangeKeyValue,omitempty"`

	// The range key name.
	RangeKeyField string `json:"rangeKeyField,omitempty" yaml:"rangeKeyField,omitempty"`

	// The range key type. Valid values are "STRING" or "NUMBER".
	RangeKeyType string `json:"rangeKeyType,omitempty" yaml:"rangeKeyType,omitempty"`

	// The ARN of the IAM role that grants access to the DynamoDB table.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The hash key name.
	HashKeyField string `json:"hashKeyField,omitempty" yaml:"hashKeyField,omitempty"`

	// The hash key type. Valid values are "STRING" or "NUMBER".
	HashKeyType string `json:"hashKeyType,omitempty" yaml:"hashKeyType,omitempty"`

	// The operation. Valid values are "INSERT", "UPDATE", or "DELETE".
	Operation string `json:"operation,omitempty" yaml:"operation,omitempty"`
}
