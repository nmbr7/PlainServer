package types

type Compute_BackendServiceOutlierDetectionInterval struct {
	/*
	   Span of time that's a fraction of a second at nanosecond resolution. Durations
	   less than one second are represented with a 0 `seconds` field and a positive
	   `nanos` field. Must be from 0 to 999,999,999 inclusive.
	*/
	Nanos int `json:"nanos,omitempty" yaml:"nanos,omitempty"`

	/*
	   Span of time at a resolution of a second. Must be from 0 to 315,576,000,000
	   inclusive.
	*/
	Seconds int `json:"seconds,omitempty" yaml:"seconds,omitempty"`
}
