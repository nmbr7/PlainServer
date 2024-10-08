package types

type Sagemaker_FeatureGroupOfflineStoreConfig struct {
	// The meta data of the Glue table that is autogenerated when an OfflineStore is created. See Data Catalog Config Below.
	DataCatalogConfig Sagemaker_FeatureGroupOfflineStoreConfigDataCatalogConfig `json:"dataCatalogConfig,omitempty" yaml:"dataCatalogConfig,omitempty"`

	// Set to `true` to turn Online Store On.
	DisableGlueTableCreation bool `json:"disableGlueTableCreation,omitempty" yaml:"disableGlueTableCreation,omitempty"`

	// The Amazon Simple Storage (Amazon S3) location of OfflineStore. See S3 Storage Config Below.
	S3StorageConfig Sagemaker_FeatureGroupOfflineStoreConfigS3StorageConfig `json:"s3StorageConfig,omitempty" yaml:"s3StorageConfig,omitempty"`

	// Format for the offline store table. Supported formats are `Glue` (Default) and Apache `Iceberg` (https://iceberg.apache.org/).
	TableFormat string `json:"tableFormat,omitempty" yaml:"tableFormat,omitempty"`
}
