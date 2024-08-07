package iot

type Policy struct {
	// The name of the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The policy document. This is a JSON formatted string. Use the [IoT Developer Guide](http://docs.aws.amazon.com/iot/latest/developerguide/iot-policies.html) for more information on IoT Policies.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
