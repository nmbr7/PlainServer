package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary struct {
	/*
	   Configures the manner in which the template library is installed on the cluster.
	   Possible values are: `INSTALATION_UNSPECIFIED`, `NOT_INSTALLED`, `ALL`.
	*/
	Installation string `json:"installation,omitempty" yaml:"installation,omitempty"`
}
