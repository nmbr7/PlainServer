package types

type Lex_V2modelsIntentConfirmationSettingElicitationCodeHook struct {
	// Label that indicates the dialog step from which the dialog code hook is happening.
	InvocationLabel string `json:"invocationLabel,omitempty" yaml:"invocationLabel,omitempty"`

	// Whether a Lambda function should be invoked for the dialog.
	EnableCodeHookInvocation bool `json:"enableCodeHookInvocation,omitempty" yaml:"enableCodeHookInvocation,omitempty"`
}
