package types

type Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessage struct {
	//
	CustomPayloads []string `json:"customPayloads,omitempty" yaml:"customPayloads,omitempty"`

	//
	ImageResponseCard Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageImageResponseCard `json:"imageResponseCard,omitempty" yaml:"imageResponseCard,omitempty"`

	//
	PlainTextMessage Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePlainTextMessage `json:"plainTextMessage,omitempty" yaml:"plainTextMessage,omitempty"`

	//
	SsmlMessage Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessageSsmlMessage `json:"ssmlMessage,omitempty" yaml:"ssmlMessage,omitempty"`
}
