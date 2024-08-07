package types

type Bigquery_JobQueryUserDefinedFunctionResource struct {
	/*
	   An inline resource that contains code for a user-defined function (UDF).
	   Providing a inline code resource is equivalent to providing a URI for a file containing the same code.
	*/
	InlineCode string `json:"inlineCode,omitempty" yaml:"inlineCode,omitempty"`

	// A code resource to load from a Google Cloud Storage URI (gs://bucket/path).
	ResourceUri string `json:"resourceUri,omitempty" yaml:"resourceUri,omitempty"`
}
