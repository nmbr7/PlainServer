package types

type Accesscontextmanager_ServicePerimeterIngressPolicyIngressFromSource struct {
	/*
	   An `AccessLevel` resource name that allow resources within the
	   `ServicePerimeters` to be accessed from the internet. `AccessLevels` listed
	   must be in the same policy as this `ServicePerimeter`. Referencing a nonexistent
	   `AccessLevel` will cause an error. If no `AccessLevel` names are listed,
	   resources within the perimeter can only be accessed via Google Cloud calls
	   with request origins within the perimeter.
	   Example `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.`
	   If - is specified, then all IngressSources will be allowed.
	*/
	AccessLevel string `json:"accessLevel,omitempty" yaml:"accessLevel,omitempty"`

	/*
	   A Google Cloud resource that is allowed to ingress the perimeter.
	   Requests from these resources will be allowed to access perimeter data.
	   Currently only projects are allowed. Format `projects/{project_number}`
	   The project may be in any Google Cloud organization, not just the
	   organization that the perimeter is defined in. `-` is not allowed, the case
	   of allowing all Google Cloud resources only is not supported.
	*/
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
