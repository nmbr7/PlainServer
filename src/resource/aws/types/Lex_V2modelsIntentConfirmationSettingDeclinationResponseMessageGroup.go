package types

type Lex_V2modelsIntentConfirmationSettingDeclinationResponseMessageGroup struct {
	// Configuration block for the primary message that Amazon Lex should send to the user. See `message`.
	Message Lex_V2modelsIntentConfirmationSettingDeclinationResponseMessageGroupMessage `json:"message,omitempty" yaml:"message,omitempty"`

	// Configuration blocks for message variations to send to the user. When variations are defined, Amazon Lex chooses the primary message or one of the variations to send to the user. See `variation`.
	Variations []Lex_V2modelsIntentConfirmationSettingDeclinationResponseMessageGroupVariation `json:"variations,omitempty" yaml:"variations,omitempty"`
}