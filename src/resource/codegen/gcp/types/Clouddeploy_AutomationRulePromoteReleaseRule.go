package types

type Clouddeploy_AutomationRulePromoteReleaseRule struct {
	// Optional. How long the release need to be paused until being promoted to the next target.
	Wait string `json:"wait,omitempty" yaml:"wait,omitempty"`

	// Optional. The starting phase of the rollout created by this operation. Default to the first phase.
	DestinationPhase string `json:"destinationPhase,omitempty" yaml:"destinationPhase,omitempty"`

	// Optional. The ID of the stage in the pipeline to which this `Release` is deploying. If unspecified, default it to the next stage in the promotion flow. The value of this field could be one of the following: - The last segment of a target name. It only needs the ID to determine if the target is one of the stages in the promotion sequence defined in the pipeline. - "@next", the next target in the promotion sequence.
	DestinationTargetId string `json:"destinationTargetId,omitempty" yaml:"destinationTargetId,omitempty"`

	// Required. ID of the rule. This id must be unique in the `Automation` resource to which this rule belongs. The format is `a-z{0,62}`.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
