package types

type Workbench_InstanceGceSetupServiceAccount struct {
	/*
	   (Output)
	   Output only. The list of scopes to be made available for this
	   service account. Set by the CLH to https://www.googleapis.com/auth/cloud-platform
	*/
	Scopes []string `json:"scopes,omitempty" yaml:"scopes,omitempty"`

	// Optional. Email address of the service account.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}
