package types

type Securitylake_DataLakeConfigurationLifecycleConfigurationExpiration struct {
	// Number of days before data transition to a different S3 Storage Class in the Amazon Security Lake object.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`
}
