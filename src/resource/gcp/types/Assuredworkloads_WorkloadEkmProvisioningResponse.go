package types

type Assuredworkloads_WorkloadEkmProvisioningResponse struct {
	// Indicates Ekm provisioning error if any. Possible values: EKM_PROVISIONING_ERROR_DOMAIN_UNSPECIFIED, UNSPECIFIED_ERROR, GOOGLE_SERVER_ERROR, EXTERNAL_USER_ERROR, EXTERNAL_PARTNER_ERROR, TIMEOUT_ERROR
	EkmProvisioningErrorDomain string `json:"ekmProvisioningErrorDomain,omitempty" yaml:"ekmProvisioningErrorDomain,omitempty"`

	// Detailed error message if Ekm provisioning fails Possible values: EKM_PROVISIONING_ERROR_MAPPING_UNSPECIFIED, INVALID_SERVICE_ACCOUNT, MISSING_METRICS_SCOPE_ADMIN_PERMISSION, MISSING_EKM_CONNECTION_ADMIN_PERMISSION
	EkmProvisioningErrorMapping string `json:"ekmProvisioningErrorMapping,omitempty" yaml:"ekmProvisioningErrorMapping,omitempty"`

	// Indicates Ekm enrollment Provisioning of a given workload. Possible values: EKM_PROVISIONING_STATE_UNSPECIFIED, EKM_PROVISIONING_STATE_PENDING, EKM_PROVISIONING_STATE_FAILED, EKM_PROVISIONING_STATE_COMPLETED
	EkmProvisioningState string `json:"ekmProvisioningState,omitempty" yaml:"ekmProvisioningState,omitempty"`
}
