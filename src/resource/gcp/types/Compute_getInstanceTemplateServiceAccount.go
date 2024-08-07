package types

type Compute_getInstanceTemplateServiceAccount struct {
	/*
	   The service account e-mail address. If not given, the
	   default Google Compute Engine service account is used.
	*/
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	/*
	   A list of service scopes. Both OAuth2 URLs and gcloud
	   short names are supported. To allow full access to all Cloud APIs, use the
	   `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
	*/
	Scopes []string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}
