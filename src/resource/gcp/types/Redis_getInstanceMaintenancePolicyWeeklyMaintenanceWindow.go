package types

type Redis_getInstanceMaintenancePolicyWeeklyMaintenanceWindow struct {
	/*
	   Required. The day of week that maintenance updates occur.

	   - DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.
	   - MONDAY: Monday
	   - TUESDAY: Tuesday
	   - WEDNESDAY: Wednesday
	   - THURSDAY: Thursday
	   - FRIDAY: Friday
	   - SATURDAY: Saturday
	   - SUNDAY: Sunday Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	*/
	Day string `json:"day,omitempty" yaml:"day,omitempty"`

	/*
	   Output only. Duration of the maintenance window.
	   The current window is fixed at 1 hour.
	   A duration in seconds with up to nine fractional digits,
	   terminated by 's'. Example: "3.5s".
	*/
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`

	// Required. Start time of the window in UTC time.
	StartTimes []Redis_getInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime `json:"startTimes,omitempty" yaml:"startTimes,omitempty"`
}
