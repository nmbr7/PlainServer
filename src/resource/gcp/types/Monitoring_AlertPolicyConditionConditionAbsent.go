package types

type Monitoring_AlertPolicyConditionConditionAbsent struct {
	/*
	   The amount of time that a time series must
	   fail to report new data to be considered
	   failing. Currently, only values that are a
	   multiple of a minute--e.g. 60s, 120s, or 300s
	   --are supported.
	*/
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	/*
	   A filter that identifies which time series
	   should be compared with the threshold.The
	   filter is similar to the one that is
	   specified in the
	   MetricService.ListTimeSeries request (that
	   call is useful to verify the time series
	   that will be retrieved / processed) and must
	   specify the metric type and optionally may
	   contain restrictions on resource type,
	   resource labels, and metric labels. This
	   field may not exceed 2048 Unicode characters
	   in length.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	/*
	   The number/percent of time series for which
	   the comparison must hold in order for the
	   condition to trigger. If unspecified, then
	   the condition will trigger if the comparison
	   is true for any of the time series that have
	   been identified by filter and aggregations.
	   Structure is documented below.
	*/
	Trigger Monitoring_AlertPolicyConditionConditionAbsentTrigger `json:"trigger,omitempty" yaml:"trigger,omitempty"`

	/*
	   Specifies the alignment of data points in
	   individual time series as well as how to
	   combine the retrieved time series together
	   (such as when aggregating multiple streams
	   on each resource to a single stream for each
	   resource or when aggregating streams across
	   all members of a group of resources).
	   Multiple aggregations are applied in the
	   order specified.
	   Structure is documented below.
	*/
	Aggregations []Monitoring_AlertPolicyConditionConditionAbsentAggregation `json:"aggregations,omitempty" yaml:"aggregations,omitempty"`
}
