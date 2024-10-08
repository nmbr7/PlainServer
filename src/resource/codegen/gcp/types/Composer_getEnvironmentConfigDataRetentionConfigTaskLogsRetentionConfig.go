package types

type Composer_getEnvironmentConfigDataRetentionConfigTaskLogsRetentionConfig struct {
	// Whether logs in cloud logging only is enabled or not. This field is supported for Cloud Composer environments in versions composer-2.0.32-airflow-2.1.4 and newer.
	StorageMode string `json:"storageMode,omitempty" yaml:"storageMode,omitempty"`
}
