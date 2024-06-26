package types

type Migrationcenter_PreferenceSetVirtualMachinePreferencesVmwareEnginePreferences struct {
	/*
	   Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with.
	   Possible values:
	   COMMITMENT_PLAN_UNSPECIFIED
	   ON_DEMAND
	   COMMITMENT_1_YEAR_MONTHLY_PAYMENTS
	   COMMITMENT_3_YEAR_MONTHLY_PAYMENTS
	   COMMITMENT_1_YEAR_UPFRONT_PAYMENT
	   COMMITMENT_3_YEAR_UPFRONT_PAYMENT
	*/
	CommitmentPlan string `json:"commitmentPlan,omitempty" yaml:"commitmentPlan,omitempty"`

	// CPU overcommit ratio. Acceptable values are between 1.0 and 8.0, with 0.1 increment.
	CpuOvercommitRatio float64 `json:"cpuOvercommitRatio,omitempty" yaml:"cpuOvercommitRatio,omitempty"`

	// Memory overcommit ratio. Acceptable values are 1.0, 1.25, 1.5, 1.75 and 2.0.
	MemoryOvercommitRatio float64 `json:"memoryOvercommitRatio,omitempty" yaml:"memoryOvercommitRatio,omitempty"`

	// The Deduplication and Compression ratio is based on the logical (Used Before) space required to store data before applying deduplication and compression, in relation to the physical (Used After) space required after applying deduplication and compression. Specifically, the ratio is the Used Before space divided by the Used After space. For example, if the Used Before space is 3 GB, but the physical Used After space is 1 GB, the deduplication and compression ratio is 3x. Acceptable values are between 1.0 and 4.0.
	StorageDeduplicationCompressionRatio float64 `json:"storageDeduplicationCompressionRatio,omitempty" yaml:"storageDeduplicationCompressionRatio,omitempty"`
}
