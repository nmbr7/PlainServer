package types

type Ssmcontacts_getPlanStageTargetContactTargetInfo struct {
	//
	IsEssential bool `json:"isEssential,omitempty" yaml:"isEssential,omitempty"`

	// The Amazon Resource Name (ARN) of the contact or escalation plan.
	ContactId string `json:"contactId,omitempty" yaml:"contactId,omitempty"`
}
