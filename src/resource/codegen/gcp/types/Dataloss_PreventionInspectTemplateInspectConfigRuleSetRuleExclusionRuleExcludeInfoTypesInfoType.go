package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypesInfoType struct {
	/*
	   Optional custom sensitivity for this InfoType. This only applies to data profiling.
	   Structure is documented below.
	*/
	SensitivityScore Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeInfoTypesInfoTypeSensitivityScore `json:"sensitivityScore,omitempty" yaml:"sensitivityScore,omitempty"`

	// Version name for this InfoType.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	   at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
