package types

type Mskconnect_CustomPluginLocation struct {
	// Information of the plugin file stored in Amazon S3. See below.
	S3 Mskconnect_CustomPluginLocationS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
