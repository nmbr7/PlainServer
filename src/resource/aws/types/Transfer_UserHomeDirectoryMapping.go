package types

type Transfer_UserHomeDirectoryMapping struct {
	// Represents an entry and a target.
	Entry string `json:"entry,omitempty" yaml:"entry,omitempty"`

	/*
	   Represents the map target.

	   The `Restricted` option is achieved using the following mapping:
	*/
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
