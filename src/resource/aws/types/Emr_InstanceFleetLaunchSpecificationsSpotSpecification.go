package types

type Emr_InstanceFleetLaunchSpecificationsSpotSpecification struct {
	// The action to take when TargetSpotCapacity has not been fulfilled when the TimeoutDurationMinutes has expired; that is, when all Spot instances could not be provisioned within the Spot provisioning timeout. Valid values are `TERMINATE_CLUSTER` and `SWITCH_TO_ON_DEMAND`. SWITCH_TO_ON_DEMAND specifies that if no Spot instances are available, On-Demand Instances should be provisioned to fulfill any remaining Spot capacity.
	TimeoutAction string `json:"timeoutAction,omitempty" yaml:"timeoutAction,omitempty"`

	// The spot provisioning timeout period in minutes. If Spot instances are not provisioned within this time period, the TimeOutAction is taken. Minimum value is 5 and maximum value is 1440. The timeout applies only during initial provisioning, when the cluster is first created.
	TimeoutDurationMinutes int `json:"timeoutDurationMinutes,omitempty" yaml:"timeoutDurationMinutes,omitempty"`

	// Specifies one of the following strategies to launch Spot Instance fleets: `price-capacity-optimized`, `capacity-optimized`, `lowest-price`, or `diversified`. For more information on the provisioning strategies, see [Allocation strategies for Spot Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-allocation-strategy.html).
	AllocationStrategy string `json:"allocationStrategy,omitempty" yaml:"allocationStrategy,omitempty"`

	// The defined duration for Spot instances (also known as Spot blocks) in minutes. When specified, the Spot instance does not terminate before the defined duration expires, and defined duration pricing for Spot instances applies. Valid values are 60, 120, 180, 240, 300, or 360. The duration period starts as soon as a Spot instance receives its instance ID. At the end of the duration, Amazon EC2 marks the Spot instance for termination and provides a Spot instance termination notice, which gives the instance a two-minute warning before it terminates.
	BlockDurationMinutes int `json:"blockDurationMinutes,omitempty" yaml:"blockDurationMinutes,omitempty"`
}
