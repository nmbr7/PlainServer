package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotword struct {
	/*
	   Regular expression pattern defining what qualifies as a hotword.
	   Structure is documented below.
	*/
	HotwordRegex Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordHotwordRegex `json:"hotwordRegex,omitempty" yaml:"hotwordRegex,omitempty"`

	/*
	   Proximity of the finding within which the entire hotword must reside. The total length of the window cannot
	   exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be
	   used to match substrings of the finding itself. For example, the certainty of a phone number regex
	   `(\d{3}) \d{3}-\d{4}` could be adjusted upwards if the area code is known to be the local area code of a company
	   office using the hotword regex `(xxx)`, where `xxx` is the area code in question.
	   Structure is documented below.
	*/
	Proximity Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleExcludeByHotwordProximity `json:"proximity,omitempty" yaml:"proximity,omitempty"`
}
