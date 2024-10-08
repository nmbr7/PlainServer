package types

type Lex_V2modelsIntentConfirmationSettingDeclinationResponse struct {
	// Configuration blocks for responses that Amazon Lex can send to the user. Amazon Lex chooses the actual response to send at runtime. See `message_group`.
	MessageGroups []Lex_V2modelsIntentConfirmationSettingDeclinationResponseMessageGroup `json:"messageGroups,omitempty" yaml:"messageGroups,omitempty"`

	// Whether the user can interrupt a speech response from Amazon Lex.
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`
}
