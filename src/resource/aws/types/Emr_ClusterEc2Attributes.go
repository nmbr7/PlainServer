package types

type Emr_ClusterEc2Attributes struct {
	// String containing a comma separated list of additional Amazon EC2 security group IDs for the slave nodes as a comma separated string.
	AdditionalSlaveSecurityGroups string `json:"additionalSlaveSecurityGroups,omitempty" yaml:"additionalSlaveSecurityGroups,omitempty"`

	// Identifier of the Amazon EC2 EMR-Managed security group for the master node.
	EmrManagedMasterSecurityGroup string `json:"emrManagedMasterSecurityGroup,omitempty" yaml:"emrManagedMasterSecurityGroup,omitempty"`

	// Identifier of the Amazon EC2 EMR-Managed security group for the slave nodes.
	EmrManagedSlaveSecurityGroup string `json:"emrManagedSlaveSecurityGroup,omitempty" yaml:"emrManagedSlaveSecurityGroup,omitempty"`

	// Instance Profile for EC2 instances of the cluster assume this role.
	InstanceProfile string `json:"instanceProfile,omitempty" yaml:"instanceProfile,omitempty"`

	// Amazon EC2 key pair that can be used to ssh to the master node as the user called `hadoop`.
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// Identifier of the Amazon EC2 service-access security group - required when the cluster runs on a private subnet.
	ServiceAccessSecurityGroup string `json:"serviceAccessSecurityGroup,omitempty" yaml:"serviceAccessSecurityGroup,omitempty"`

	// String containing a comma separated list of additional Amazon EC2 security group IDs for the master node.
	AdditionalMasterSecurityGroups string `json:"additionalMasterSecurityGroups,omitempty" yaml:"additionalMasterSecurityGroups,omitempty"`

	// VPC subnet id where you want the job flow to launch. Cannot specify the `cc1.4xlarge` instance type for nodes of a job flow launched in an Amazon VPC.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	/*
	   List of VPC subnet id-s where you want the job flow to launch.  Amazon EMR identifies the best Availability Zone to launch instances according to your fleet specifications.

	   > --NOTE on EMR-Managed security groups:-- These security groups will have any missing inbound or outbound access rules added and maintained by AWS, to ensure proper communication between instances in a cluster. The EMR service will maintain these rules for groups provided in `emr_managed_master_security_group` and `emr_managed_slave_security_group`; attempts to remove the required rules may succeed, only for the EMR service to re-add them in a matter of minutes. This may cause this provider to fail to destroy an environment that contains an EMR cluster, because the EMR service does not revoke rules added on deletion, leaving a cyclic dependency between the security groups that prevents their deletion. To avoid this, use the `revoke_rules_on_delete` optional attribute for any Security Group used in `emr_managed_master_security_group` and `emr_managed_slave_security_group`. See [Amazon EMR-Managed Security Groups](http://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-man-sec-groups.html) for more information about the EMR-managed security group rules.
	*/
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
