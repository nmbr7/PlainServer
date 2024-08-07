package types

type Storage_TransferJobTransferSpecGcsDataSink struct {
	// Google Cloud Storage bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
