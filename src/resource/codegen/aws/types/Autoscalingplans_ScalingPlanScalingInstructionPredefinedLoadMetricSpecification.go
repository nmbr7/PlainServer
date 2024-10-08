package types

type Autoscalingplans_ScalingPlanScalingInstructionPredefinedLoadMetricSpecification struct {
	// Metric type. Valid values: `ALBTargetGroupRequestCount`, `ASGTotalCPUUtilization`, `ASGTotalNetworkIn`, `ASGTotalNetworkOut`.
	PredefinedLoadMetricType string `json:"predefinedLoadMetricType,omitempty" yaml:"predefinedLoadMetricType,omitempty"`

	// Identifies the resource associated with the metric type.
	ResourceLabel string `json:"resourceLabel,omitempty" yaml:"resourceLabel,omitempty"`
}
