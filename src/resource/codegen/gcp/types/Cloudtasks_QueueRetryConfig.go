package types

type Cloudtasks_QueueRetryConfig struct {
	/*
	   The time between retries will double maxDoublings times.
	   A task's retry interval starts at minBackoff, then doubles maxDoublings times,
	   then increases linearly, and finally retries retries at intervals of maxBackoff
	   up to maxAttempts times.
	*/
	MaxDoublings int `json:"maxDoublings,omitempty" yaml:"maxDoublings,omitempty"`

	/*
	   If positive, maxRetryDuration specifies the time limit for
	   retrying a failed task, measured from when the task was first
	   attempted. Once maxRetryDuration time has passed and the task has
	   been attempted maxAttempts times, no further attempts will be
	   made and the task will be deleted.
	   If zero, then the task age is unlimited.
	*/
	MaxRetryDuration string `json:"maxRetryDuration,omitempty" yaml:"maxRetryDuration,omitempty"`

	/*
	   A task will be scheduled for retry between minBackoff and
	   maxBackoff duration after it fails, if the queue's RetryConfig
	   specifies that the task should be retried.
	*/
	MinBackoff string `json:"minBackoff,omitempty" yaml:"minBackoff,omitempty"`

	/*
	   Number of attempts per task.
	   Cloud Tasks will attempt the task maxAttempts times (that is, if
	   the first attempt fails, then there will be maxAttempts - 1
	   retries). Must be >= -1.
	   If unspecified when the queue is created, Cloud Tasks will pick
	   the default.
	   -1 indicates unlimited attempts.
	*/
	MaxAttempts int `json:"maxAttempts,omitempty" yaml:"maxAttempts,omitempty"`

	/*
	   A task will be scheduled for retry between minBackoff and
	   maxBackoff duration after it fails, if the queue's RetryConfig
	   specifies that the task should be retried.
	*/
	MaxBackoff string `json:"maxBackoff,omitempty" yaml:"maxBackoff,omitempty"`
}
