package types

type Cloudtasks_QueueStackdriverLoggingConfig struct {
	/*
	   Specifies the fraction of operations to write to Stackdriver Logging.
	   This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
	   default and means that no operations are logged.
	*/
	SamplingRatio float64 `json:"samplingRatio,omitempty" yaml:"samplingRatio,omitempty"`
}
