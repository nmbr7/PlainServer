package types

type Dataproc_AutoscalingPolicyBasicAlgorithmYarnConfig struct {
	/*
	   Timeout for YARN graceful decommissioning of Node Managers. Specifies the
	   duration to wait for jobs to complete before forcefully removing workers
	   (and potentially interrupting jobs). Only applicable to downscaling operations.
	   Bounds: [0s, 1d].
	*/
	GracefulDecommissionTimeout string `json:"gracefulDecommissionTimeout,omitempty" yaml:"gracefulDecommissionTimeout,omitempty"`

	/*
	   Fraction of average pending memory in the last cooldown period for which to
	   remove workers. A scale-down factor of 1 will result in scaling down so that there
	   is no available memory remaining after the update (more aggressive scaling).
	   A scale-down factor of 0 disables removing workers, which can be beneficial for
	   autoscaling a single job.
	   Bounds: [0.0, 1.0].
	*/
	ScaleDownFactor float64 `json:"scaleDownFactor,omitempty" yaml:"scaleDownFactor,omitempty"`

	/*
	   Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
	   For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
	   recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
	   means the autoscaler will scale down on any recommended change.
	   Bounds: [0.0, 1.0]. Default: 0.0.
	*/
	ScaleDownMinWorkerFraction float64 `json:"scaleDownMinWorkerFraction,omitempty" yaml:"scaleDownMinWorkerFraction,omitempty"`

	/*
	   Fraction of average pending memory in the last cooldown period for which to
	   add workers. A scale-up factor of 1.0 will result in scaling up so that there
	   is no pending memory remaining after the update (more aggressive scaling).
	   A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
	   (less aggressive scaling).
	   Bounds: [0.0, 1.0].
	*/
	ScaleUpFactor float64 `json:"scaleUpFactor,omitempty" yaml:"scaleUpFactor,omitempty"`

	/*
	   Minimum scale-up threshold as a fraction of total cluster size before scaling
	   occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
	   must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
	   0 means the autoscaler will scale up on any recommended change.
	   Bounds: [0.0, 1.0]. Default: 0.0.
	*/
	ScaleUpMinWorkerFraction float64 `json:"scaleUpMinWorkerFraction,omitempty" yaml:"scaleUpMinWorkerFraction,omitempty"`
}
