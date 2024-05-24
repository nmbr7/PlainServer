package shield

import types "DesignSphere_Server/src/resource/aws/types"

type DrtAccessRoleArnAssociation struct {
	// The Amazon Resource Name (ARN) of the role the SRT will use to access your AWS account. Prior to making the AssociateDRTRole request, you must attach the `AWSShieldDRTAccessPolicy` managed policy to this role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	//
	Timeouts types.Shield_DrtAccessRoleArnAssociationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
