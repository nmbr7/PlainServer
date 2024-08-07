package types

type Binaryauthorization_PolicyAdmissionWhitelistPattern struct {
	/*
	   An image name pattern to whitelist, in the form
	   `registry/path/to/image`. This supports a trailing - as a
	   wildcard, but this is allowed only in text after the registry/
	   part.
	*/
	NamePattern string `json:"namePattern,omitempty" yaml:"namePattern,omitempty"`
}
