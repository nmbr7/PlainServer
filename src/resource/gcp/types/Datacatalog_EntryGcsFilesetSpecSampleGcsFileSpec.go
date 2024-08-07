package types

type Datacatalog_EntryGcsFilesetSpecSampleGcsFileSpec struct {
	/*
	   (Output)
	   The size of the file, in bytes.
	*/
	SizeBytes int `json:"sizeBytes,omitempty" yaml:"sizeBytes,omitempty"`

	/*
	   (Output)
	   The full file path
	*/
	FilePath string `json:"filePath,omitempty" yaml:"filePath,omitempty"`
}
