package types

type Datastream_StreamDestinationConfigBigqueryDestinationConfig struct {
	/*
	   The guaranteed data freshness (in seconds) when querying tables created by the stream.
	   Editing this field will only affect new tables created in the future, but existing tables
	   will not be impacted. Lower values mean that queries will return fresher data, but may result in higher cost.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
	*/
	DataFreshness string `json:"dataFreshness,omitempty" yaml:"dataFreshness,omitempty"`

	/*
	   A single target dataset to which all data will be streamed.
	   Structure is documented below.
	*/
	SingleTargetDataset Datastream_StreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset `json:"singleTargetDataset,omitempty" yaml:"singleTargetDataset,omitempty"`

	/*
	   Destination datasets are created so that hierarchy of the destination data objects matches the source hierarchy.
	   Structure is documented below.
	*/
	SourceHierarchyDatasets Datastream_StreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets `json:"sourceHierarchyDatasets,omitempty" yaml:"sourceHierarchyDatasets,omitempty"`
}
