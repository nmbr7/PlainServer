package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedMetricPairSpecification struct {
	// Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group. You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). Refer to [PredefinedMetricSpecification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredefinedMetricSpecification.html) for more information.
	ResourceLabel string `json:"resourceLabel,omitempty" yaml:"resourceLabel,omitempty"`

	// Which metrics to use. There are two different types of metrics for each metric type: one is a load metric and one is a scaling metric. For example, if the metric type is `ASGCPUUtilization`, the Auto Scaling group's total CPU metric is used as the load metric, and the average CPU metric is used for the scaling metric. Valid values are `ASGCPUUtilization`, `ASGNetworkIn`, `ASGNetworkOut`, or `ALBRequestCount`.
	PredefinedMetricType string `json:"predefinedMetricType,omitempty" yaml:"predefinedMetricType,omitempty"`
}
