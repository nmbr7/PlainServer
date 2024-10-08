package types

type Servicequotas_ServiceQuotaUsageMetric struct {
	// The namespace of the metric.
	MetricNamespace string `json:"metricNamespace,omitempty" yaml:"metricNamespace,omitempty"`

	// The metric statistic that AWS recommend you use when determining quota usage.
	MetricStatisticRecommendation string `json:"metricStatisticRecommendation,omitempty" yaml:"metricStatisticRecommendation,omitempty"`

	// The metric dimensions.
	MetricDimensions []Servicequotas_ServiceQuotaUsageMetricMetricDimension `json:"metricDimensions,omitempty" yaml:"metricDimensions,omitempty"`

	// The name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`
}
