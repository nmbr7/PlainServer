package types

type Apprunner_ServiceInstanceConfiguration struct {
	// Number of CPU units reserved for each instance of your App Runner service represented as a String. Defaults to `1024`. Valid values: `256|512|1024|2048|4096|(0.25|0.5|1|2|4) vCPU`.
	Cpu string `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// ARN of an IAM role that provides permissions to your App Runner service. These are permissions that your code needs when it calls any AWS APIs.
	InstanceRoleArn string `json:"instanceRoleArn,omitempty" yaml:"instanceRoleArn,omitempty"`

	// Amount of memory, in MB or GB, reserved for each instance of your App Runner service. Defaults to `2048`. Valid values: `512|1024|2048|3072|4096|6144|8192|10240|12288|(0.5|1|2|3|4|6|8|10|12) GB`.
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}
