package types

type Cloudrunv2_getJobTemplateTemplateVolumeCloudSqlInstance struct {
	// The Cloud SQL instance connection names, as can be found in https://console.cloud.google.com/sql/instances. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run. Format: {project}:{location}:{instance}
	Instances []string `json:"instances,omitempty" yaml:"instances,omitempty"`
}
