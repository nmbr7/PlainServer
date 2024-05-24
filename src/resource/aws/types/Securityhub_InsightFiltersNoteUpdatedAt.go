package types

type Securityhub_InsightFiltersNoteUpdatedAt struct {
	// A start date for the date filter. Required with `end` if `date_range` is not specified.
	Start string `json:"start,omitempty" yaml:"start,omitempty"`

	// A configuration block of the date range for the date filter. See date_range below for more details.
	DateRange Securityhub_InsightFiltersNoteUpdatedAtDateRange `json:"dateRange,omitempty" yaml:"dateRange,omitempty"`

	// An end date for the date filter. Required with `start` if `date_range` is not specified.
	End string `json:"end,omitempty" yaml:"end,omitempty"`
}
