package types

type Securitycenter_FolderCustomModuleCustomConfig struct {
	/*
	   The resource types that the custom module operates on. Each custom module
	   can specify up to 5 resource
	   Structure is documented below.
	*/
	ResourceSelector Securitycenter_FolderCustomModuleCustomConfigResourceSelector `json:"resourceSelector,omitempty" yaml:"resourceSelector,omitempty"`

	/*
	   The severity to assign to findings generated by the module.
	   Possible values are: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`.
	*/
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`

	/*
	   Custom output properties.
	   Structure is documented below.
	*/
	CustomOutput Securitycenter_FolderCustomModuleCustomConfigCustomOutput `json:"customOutput,omitempty" yaml:"customOutput,omitempty"`

	/*
	   Text that describes the vulnerability or misconfiguration that the custom
	   module detects. This explanation is returned with each finding instance to
	   help investigators understand the detected issue. The text must be enclosed in quotation marks.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The CEL expression to evaluate to produce findings. When the expression evaluates
	   to true against a resource, a finding is generated.
	   Structure is documented below.
	*/
	Predicate Securitycenter_FolderCustomModuleCustomConfigPredicate `json:"predicate,omitempty" yaml:"predicate,omitempty"`

	/*
	   An explanation of the recommended steps that security teams can take to resolve
	   the detected issue. This explanation is returned with each finding generated by
	   this module in the nextSteps property of the finding JSON.
	*/
	Recommendation string `json:"recommendation,omitempty" yaml:"recommendation,omitempty"`
}