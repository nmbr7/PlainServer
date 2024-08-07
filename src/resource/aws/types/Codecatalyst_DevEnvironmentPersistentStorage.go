package types

type Codecatalyst_DevEnvironmentPersistentStorage struct {
	// The size of the persistent storage in gigabytes (specifically GiB). Valid values for storage are based on memory sizes in 16GB increments. Valid values are 16, 32, and 64.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`
}
