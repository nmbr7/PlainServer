package types

type Eks_ClusterOutpostConfig struct {
	/*
	   The Amazon EC2 instance type that you want to use for your local Amazon EKS cluster on Outposts. The instance type that you specify is used for all Kubernetes control plane instances. The instance type can't be changed after cluster creation. Choose an instance type based on the number of nodes that your cluster will have. If your cluster will have:

	   - 1–20 nodes, then we recommend specifying a large instance type.

	   - 21–100 nodes, then we recommend specifying an xlarge instance type.

	   - 101–250 nodes, then we recommend specifying a 2xlarge instance type.

	   For a list of the available Amazon EC2 instance types, see Compute and storage in AWS Outposts rack features  The control plane is not automatically scaled by Amazon EKS.
	*/
	ControlPlaneInstanceType string `json:"controlPlaneInstanceType,omitempty" yaml:"controlPlaneInstanceType,omitempty"`

	/*
	   An object representing the placement configuration for all the control plane instances of your local Amazon EKS cluster on AWS Outpost.
	   The `control_plane_placement` configuration block supports the following arguments:
	*/
	ControlPlanePlacement Eks_ClusterOutpostConfigControlPlanePlacement `json:"controlPlanePlacement,omitempty" yaml:"controlPlanePlacement,omitempty"`

	// The ARN of the Outpost that you want to use for your local Amazon EKS cluster on Outposts. This argument is a list of arns, but only a single Outpost ARN is supported currently.
	OutpostArns []string `json:"outpostArns,omitempty" yaml:"outpostArns,omitempty"`
}
