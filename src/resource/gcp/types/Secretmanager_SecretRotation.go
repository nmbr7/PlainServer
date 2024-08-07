package types

type Secretmanager_SecretRotation struct {
	/*
	   Timestamp in UTC at which the Secret is scheduled to rotate.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	*/
	NextRotationTime string `json:"nextRotationTime,omitempty" yaml:"nextRotationTime,omitempty"`

	/*
	   The Duration between rotation notifications. Must be in seconds and at least 3600s (1h) and at most 3153600000s (100 years).
	   If rotationPeriod is set, `next_rotation_time` must be set. `next_rotation_time` will be advanced by this period when the service automatically sends rotation notifications.
	*/
	RotationPeriod string `json:"rotationPeriod,omitempty" yaml:"rotationPeriod,omitempty"`
}
