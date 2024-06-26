package types

type Costexplorer_AnomalySubscriptionThresholdExpressionNot struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags Costexplorer_AnomalySubscriptionThresholdExpressionNotTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for the filter that's based on  values. See Cost Category below.
	CostCategory Costexplorer_AnomalySubscriptionThresholdExpressionNotCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific Dimension to use for.
	Dimension Costexplorer_AnomalySubscriptionThresholdExpressionNotDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`
}
