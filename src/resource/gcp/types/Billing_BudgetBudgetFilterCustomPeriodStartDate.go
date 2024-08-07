package types

type Billing_BudgetBudgetFilterCustomPeriodStartDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month.
	Day int `json:"day,omitempty" yaml:"day,omitempty"`

	// Month of a year. Must be from 1 to 12.
	Month int `json:"month,omitempty" yaml:"month,omitempty"`

	// Year of the date. Must be from 1 to 9999.
	Year int `json:"year,omitempty" yaml:"year,omitempty"`
}
