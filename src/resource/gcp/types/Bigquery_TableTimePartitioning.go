package types

type Bigquery_TableTimePartitioning struct {
	/*
	   Number of milliseconds for which to keep the
	   storage for a partition.
	*/
	ExpirationMs int `json:"expirationMs,omitempty" yaml:"expirationMs,omitempty"`

	/*
	   The field used to determine how to create a time-based
	   partition. If time-based partitioning is enabled without this value, the
	   table is partitioned based on the load time.
	*/
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   If set to true, queries over this table
	   require a partition filter that can be used for partition elimination to be
	   specified. `require_partition_filter` is deprecated and will be removed in
	   a future major release. Use the top level field with the same name instead.
	*/
	RequirePartitionFilter bool `json:"requirePartitionFilter,omitempty" yaml:"requirePartitionFilter,omitempty"`

	/*
	   The supported types are DAY, HOUR, MONTH, and YEAR,
	   which will generate one partition per day, hour, month, and year, respectively.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
