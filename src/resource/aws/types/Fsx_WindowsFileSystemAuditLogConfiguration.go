package types

type Fsx_WindowsFileSystemAuditLogConfiguration struct {
	// The Amazon Resource Name (ARN) for the destination of the audit logs. The destination can be any Amazon CloudWatch Logs log group ARN or Amazon Kinesis Data Firehose delivery stream ARN. Can be specified when `file_access_audit_log_level` and `file_share_access_audit_log_level` are not set to `DISABLED`. The name of the Amazon CloudWatch Logs log group must begin with the `/aws/fsx` prefix. The name of the Amazon Kinesis Data Firehouse delivery stream must begin with the `aws-fsx` prefix. If you do not provide a destination in `audit_log_destionation`, Amazon FSx will create and use a log stream in the CloudWatch Logs /aws/fsx/windows log group.
	AuditLogDestination string `json:"auditLogDestination,omitempty" yaml:"auditLogDestination,omitempty"`

	// Sets which attempt type is logged by Amazon FSx for file and folder accesses. Valid values are `SUCCESS_ONLY`, `FAILURE_ONLY`, `SUCCESS_AND_FAILURE`, and `DISABLED`. Default value is `DISABLED`.
	FileAccessAuditLogLevel string `json:"fileAccessAuditLogLevel,omitempty" yaml:"fileAccessAuditLogLevel,omitempty"`

	// Sets which attempt type is logged by Amazon FSx for file share accesses. Valid values are `SUCCESS_ONLY`, `FAILURE_ONLY`, `SUCCESS_AND_FAILURE`, and `DISABLED`. Default value is `DISABLED`.
	FileShareAccessAuditLogLevel string `json:"fileShareAccessAuditLogLevel,omitempty" yaml:"fileShareAccessAuditLogLevel,omitempty"`
}
