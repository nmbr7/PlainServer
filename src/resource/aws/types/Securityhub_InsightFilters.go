package types

type Securityhub_InsightFilters struct {
	// The canonical identifier for the given resource type. See String Filter below for more details.
	ResourceIds []Securityhub_InsightFiltersResourceId `json:"resourceIds,omitempty" yaml:"resourceIds,omitempty"`

	// A URL that links to a page about the current finding in the security-findings provider's solution. See String Filter below for more details.
	SourceUrls []Securityhub_InsightFiltersSourceUrl `json:"sourceUrls,omitempty" yaml:"sourceUrls,omitempty"`

	// The principal that created a note. See String Filter below for more details.
	NoteUpdatedBies []Securityhub_InsightFiltersNoteUpdatedBy `json:"noteUpdatedBies,omitempty" yaml:"noteUpdatedBies,omitempty"`

	// The identifier of the subnet that the instance was launched in. See String Filter below for more details.
	ResourceAwsEc2InstanceSubnetIds []Securityhub_InsightFiltersResourceAwsEc2InstanceSubnetId `json:"resourceAwsEc2InstanceSubnetIds,omitempty" yaml:"resourceAwsEc2InstanceSubnetIds,omitempty"`

	// A list of AWS tags associated with a resource at the time the finding was processed. See Map Filter below for more details.
	ResourceTags []Securityhub_InsightFiltersResourceTag `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`

	// An ISO8601-formatted timestamp that indicates when the security-findings provider last updated the finding record. See Date Filter below for more details.
	UpdatedAts []Securityhub_InsightFiltersUpdatedAt `json:"updatedAts,omitempty" yaml:"updatedAts,omitempty"`

	// The level of importance assigned to the resources associated with the finding. A score of 0 means that the underlying resources have no criticality, and a score of 100 is reserved for the most critical resources. See Number Filter below for more details.
	Criticalities []Securityhub_InsightFiltersCriticality `json:"criticalities,omitempty" yaml:"criticalities,omitempty"`

	// The source domain of network-related information about a finding. See String Filter below for more details.
	NetworkSourceDomains []Securityhub_InsightFiltersNetworkSourceDomain `json:"networkSourceDomains,omitempty" yaml:"networkSourceDomains,omitempty"`

	// The path to the process executable. See String Filter below for more details.
	ProcessPaths []Securityhub_InsightFiltersProcessPath `json:"processPaths,omitempty" yaml:"processPaths,omitempty"`

	// The ARN generated by Security Hub that uniquely identifies a third-party company (security findings provider) after this provider's product (solution that generates findings) is registered with Security Hub. See String Filter below for more details.
	ProductArns []Securityhub_InsightFiltersProductArn `json:"productArns,omitempty" yaml:"productArns,omitempty"`

	// The IPv6 addresses associated with the instance. See Ip Filter below for more details.
	ResourceAwsEc2InstanceIpv6Addresses []Securityhub_InsightFiltersResourceAwsEc2InstanceIpv6Address `json:"resourceAwsEc2InstanceIpv6Addresses,omitempty" yaml:"resourceAwsEc2InstanceIpv6Addresses,omitempty"`

	// The date and time the instance was launched. See Date Filter below for more details.
	ResourceAwsEc2InstanceLaunchedAts []Securityhub_InsightFiltersResourceAwsEc2InstanceLaunchedAt `json:"resourceAwsEc2InstanceLaunchedAts,omitempty" yaml:"resourceAwsEc2InstanceLaunchedAts,omitempty"`

	// The canonical user ID of the owner of the S3 bucket. See String Filter below for more details.
	ResourceAwsS3BucketOwnerIds []Securityhub_InsightFiltersResourceAwsS3BucketOwnerId `json:"resourceAwsS3BucketOwnerIds,omitempty" yaml:"resourceAwsS3BucketOwnerIds,omitempty"`

	// The name of the process. See String Filter below for more details.
	ProcessNames []Securityhub_InsightFiltersProcessName `json:"processNames,omitempty" yaml:"processNames,omitempty"`

	// A keyword for a finding. See Keyword Filter below for more details.
	Keywords []Securityhub_InsightFiltersKeyword `json:"keywords,omitempty" yaml:"keywords,omitempty"`

	// An ISO8601-formatted timestamp that indicates when the security-findings provider most recently observed the potential security issue that a finding captured. See Date Filter below for more details.
	LastObservedAts []Securityhub_InsightFiltersLastObservedAt `json:"lastObservedAts,omitempty" yaml:"lastObservedAts,omitempty"`

	// The state of the malware that was observed. See String Filter below for more details.
	MalwareStates []Securityhub_InsightFiltersMalwareState `json:"malwareStates,omitempty" yaml:"malwareStates,omitempty"`

	// The type of the malware that was observed. See String Filter below for more details.
	MalwareTypes []Securityhub_InsightFiltersMalwareType `json:"malwareTypes,omitempty" yaml:"malwareTypes,omitempty"`

	// The destination domain of network-related information about a finding. See String Filter below for more details.
	NetworkDestinationDomains []Securityhub_InsightFiltersNetworkDestinationDomain `json:"networkDestinationDomains,omitempty" yaml:"networkDestinationDomains,omitempty"`

	// The source IPv6 address of network-related information about a finding. See Ip Filter below for more details.
	NetworkSourceIpv6s []Securityhub_InsightFiltersNetworkSourceIpv6 `json:"networkSourceIpv6s,omitempty" yaml:"networkSourceIpv6s,omitempty"`

	// The date/time that the process was launched. See Date Filter below for more details.
	ProcessLaunchedAts []Securityhub_InsightFiltersProcessLaunchedAt `json:"processLaunchedAts,omitempty" yaml:"processLaunchedAts,omitempty"`

	// The name of the container related to a finding. See String Filter below for more details.
	ResourceContainerNames []Securityhub_InsightFiltersResourceContainerName `json:"resourceContainerNames,omitempty" yaml:"resourceContainerNames,omitempty"`

	// The label of a finding's severity. See String Filter below for more details.
	SeverityLabels []Securityhub_InsightFiltersSeverityLabel `json:"severityLabels,omitempty" yaml:"severityLabels,omitempty"`

	// The URL for more details from the source of the threat intelligence. See String Filter below for more details.
	ThreatIntelIndicatorSourceUrls []Securityhub_InsightFiltersThreatIntelIndicatorSourceUrl `json:"threatIntelIndicatorSourceUrls,omitempty" yaml:"threatIntelIndicatorSourceUrls,omitempty"`

	// A finding type in the format of `namespace/category/classifier` that classifies a finding. See String Filter below for more details.
	Types []Securityhub_InsightFiltersType `json:"types,omitempty" yaml:"types,omitempty"`

	// AWS account ID that a finding is generated in. See String_Filter below for more details.
	AwsAccountIds []Securityhub_InsightFiltersAwsAccountId `json:"awsAccountIds,omitempty" yaml:"awsAccountIds,omitempty"`

	// The finding provider value for the finding confidence. Confidence is defined as the likelihood that a finding accurately identifies the behavior or issue that it was intended to identify. Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent confidence and 100 means 100 percent confidence. See Number Filter below for more details.
	FindingProviderFieldsConfidences []Securityhub_InsightFiltersFindingProviderFieldsConfidence `json:"findingProviderFieldsConfidences,omitempty" yaml:"findingProviderFieldsConfidences,omitempty"`

	// The IPv4 addresses associated with the instance. See Ip Filter below for more details.
	ResourceAwsEc2InstanceIpv4Addresses []Securityhub_InsightFiltersResourceAwsEc2InstanceIpv4Address `json:"resourceAwsEc2InstanceIpv4Addresses,omitempty" yaml:"resourceAwsEc2InstanceIpv4Addresses,omitempty"`

	// The details of a resource that doesn't have a specific subfield for the resource type defined. See Map Filter below for more details.
	ResourceDetailsOthers []Securityhub_InsightFiltersResourceDetailsOther `json:"resourceDetailsOthers,omitempty" yaml:"resourceDetailsOthers,omitempty"`

	// The canonical AWS external Region name where this resource is located. See String Filter below for more details.
	ResourceRegions []Securityhub_InsightFiltersResourceRegion `json:"resourceRegions,omitempty" yaml:"resourceRegions,omitempty"`

	// The type of a threat intelligence indicator. See String Filter below for more details.
	ThreatIntelIndicatorTypes []Securityhub_InsightFiltersThreatIntelIndicatorType `json:"threatIntelIndicatorTypes,omitempty" yaml:"threatIntelIndicatorTypes,omitempty"`

	// An ISO8601-formatted timestamp that indicates when the security-findings provider first observed the potential security issue that a finding captured. See Date Filter below for more details.
	FirstObservedAts []Securityhub_InsightFiltersFirstObservedAt `json:"firstObservedAts,omitempty" yaml:"firstObservedAts,omitempty"`

	// The identifier for the solution-specific component (a discrete unit of logic) that generated a finding. See String Filter below for more details.
	GeneratorIds []Securityhub_InsightFiltersGeneratorId `json:"generatorIds,omitempty" yaml:"generatorIds,omitempty"`

	// The security findings provider-specific identifier for a finding. See String Filter below for more details.
	Ids []Securityhub_InsightFiltersId `json:"ids,omitempty" yaml:"ids,omitempty"`

	// Indicates the direction of network traffic associated with a finding. See String Filter below for more details.
	NetworkDirections []Securityhub_InsightFiltersNetworkDirection `json:"networkDirections,omitempty" yaml:"networkDirections,omitempty"`

	// The source IPv4 address of network-related information about a finding. See Ip Filter below for more details.
	NetworkSourceIpv4s []Securityhub_InsightFiltersNetworkSourceIpv4 `json:"networkSourceIpv4s,omitempty" yaml:"networkSourceIpv4s,omitempty"`

	// The process ID. See Number Filter below for more details.
	ProcessPids []Securityhub_InsightFiltersProcessPid `json:"processPids,omitempty" yaml:"processPids,omitempty"`

	// The veracity of a finding. See String Filter below for more details.
	VerificationStates []Securityhub_InsightFiltersVerificationState `json:"verificationStates,omitempty" yaml:"verificationStates,omitempty"`

	// A finding's confidence. Confidence is defined as the likelihood that a finding accurately identifies the behavior or issue that it was intended to identify. Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent confidence and 100 means 100 percent confidence. See Number Filter below for more details.
	Confidences []Securityhub_InsightFiltersConfidence `json:"confidences,omitempty" yaml:"confidences,omitempty"`

	// The destination IPv4 address of network-related information about a finding. See Ip Filter below for more details.
	NetworkDestinationIpv4s []Securityhub_InsightFiltersNetworkDestinationIpv4 `json:"networkDestinationIpv4s,omitempty" yaml:"networkDestinationIpv4s,omitempty"`

	// A list of name/value string pairs associated with the finding. These are custom, user-defined fields added to a finding. See Map Filter below for more details.
	UserDefinedValues []Securityhub_InsightFiltersUserDefinedValue `json:"userDefinedValues,omitempty" yaml:"userDefinedValues,omitempty"`

	// The finding identifier of a related finding that is identified by the finding provider. See String Filter below for more details.
	FindingProviderFieldsRelatedFindingsIds []Securityhub_InsightFiltersFindingProviderFieldsRelatedFindingsId `json:"findingProviderFieldsRelatedFindingsIds,omitempty" yaml:"findingProviderFieldsRelatedFindingsIds,omitempty"`

	// The timestamp of when the note was updated. See Date Filter below for more details.
	NoteUpdatedAts []Securityhub_InsightFiltersNoteUpdatedAt `json:"noteUpdatedAts,omitempty" yaml:"noteUpdatedAts,omitempty"`

	// The parent process ID. See Number Filter below for more details.
	ProcessParentPids []Securityhub_InsightFiltersProcessParentPid `json:"processParentPids,omitempty" yaml:"processParentPids,omitempty"`

	// The instance type of the instance. See String Filter below for more details.
	ResourceAwsEc2InstanceTypes []Securityhub_InsightFiltersResourceAwsEc2InstanceType `json:"resourceAwsEc2InstanceTypes,omitempty" yaml:"resourceAwsEc2InstanceTypes,omitempty"`

	// The identifier of the image related to a finding. See String Filter below for more details.
	ResourceContainerImageIds []Securityhub_InsightFiltersResourceContainerImageId `json:"resourceContainerImageIds,omitempty" yaml:"resourceContainerImageIds,omitempty"`

	// The status of the investigation into a finding. See Workflow Status Filter below for more details.
	WorkflowStatuses []Securityhub_InsightFiltersWorkflowStatus `json:"workflowStatuses,omitempty" yaml:"workflowStatuses,omitempty"`

	// The category of a threat intelligence indicator. See String Filter below for more details.
	ThreatIntelIndicatorCategories []Securityhub_InsightFiltersThreatIntelIndicatorCategory `json:"threatIntelIndicatorCategories,omitempty" yaml:"threatIntelIndicatorCategories,omitempty"`

	// The finding provider value for the severity label. See String Filter below for more details.
	FindingProviderFieldsSeverityLabels []Securityhub_InsightFiltersFindingProviderFieldsSeverityLabel `json:"findingProviderFieldsSeverityLabels,omitempty" yaml:"findingProviderFieldsSeverityLabels,omitempty"`

	// The name of the solution (product) that generates findings. See String Filter below for more details.
	ProductNames []Securityhub_InsightFiltersProductName `json:"productNames,omitempty" yaml:"productNames,omitempty"`

	// The recommendation of what to do about the issue described in a finding. See String Filter below for more details.
	RecommendationTexts []Securityhub_InsightFiltersRecommendationText `json:"recommendationTexts,omitempty" yaml:"recommendationTexts,omitempty"`

	// The identifier of the VPC that the instance was launched in. See String Filter below for more details.
	ResourceAwsEc2InstanceVpcIds []Securityhub_InsightFiltersResourceAwsEc2InstanceVpcId `json:"resourceAwsEc2InstanceVpcIds,omitempty" yaml:"resourceAwsEc2InstanceVpcIds,omitempty"`

	// The user associated with the IAM access key related to a finding. See String Filter below for more details.
	ResourceAwsIamAccessKeyUserNames []Securityhub_InsightFiltersResourceAwsIamAccessKeyUserName `json:"resourceAwsIamAccessKeyUserNames,omitempty" yaml:"resourceAwsIamAccessKeyUserNames,omitempty"`

	// The name of the image related to a finding. See String Filter below for more details.
	ResourceContainerImageNames []Securityhub_InsightFiltersResourceContainerImageName `json:"resourceContainerImageNames,omitempty" yaml:"resourceContainerImageNames,omitempty"`

	// Specifies the type of the resource that details are provided for. See String Filter below for more details.
	ResourceTypes []Securityhub_InsightFiltersResourceType `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`

	// Exclusive to findings that are generated as the result of a check run against a specific rule in a supported standard, such as CIS AWS Foundations. Contains security standard-related finding details. See String Filter below for more details.
	ComplianceStatuses []Securityhub_InsightFiltersComplianceStatus `json:"complianceStatuses,omitempty" yaml:"complianceStatuses,omitempty"`

	// A finding's description. See String Filter below for more details.
	Descriptions []Securityhub_InsightFiltersDescription `json:"descriptions,omitempty" yaml:"descriptions,omitempty"`

	// The updated record state for the finding. See String Filter below for more details.
	RecordStates []Securityhub_InsightFiltersRecordState `json:"recordStates,omitempty" yaml:"recordStates,omitempty"`

	// The Amazon Machine Image (AMI) ID of the instance. See String Filter below for more details.
	ResourceAwsEc2InstanceImageIds []Securityhub_InsightFiltersResourceAwsEc2InstanceImageId `json:"resourceAwsEc2InstanceImageIds,omitempty" yaml:"resourceAwsEc2InstanceImageIds,omitempty"`

	// The canonical AWS partition name that the Region is assigned to. See String Filter below for more details.
	ResourcePartitions []Securityhub_InsightFiltersResourcePartition `json:"resourcePartitions,omitempty" yaml:"resourcePartitions,omitempty"`

	// The ARN of the solution that generated a related finding that is identified by the finding provider. See String Filter below for more details.
	FindingProviderFieldsRelatedFindingsProductArns []Securityhub_InsightFiltersFindingProviderFieldsRelatedFindingsProductArn `json:"findingProviderFieldsRelatedFindingsProductArns,omitempty" yaml:"findingProviderFieldsRelatedFindingsProductArns,omitempty"`

	// The filesystem path of the malware that was observed. See String Filter below for more details.
	MalwarePaths []Securityhub_InsightFiltersMalwarePath `json:"malwarePaths,omitempty" yaml:"malwarePaths,omitempty"`

	// The creation date/time of the IAM access key related to a finding. See Date Filter below for more details.
	ResourceAwsIamAccessKeyCreatedAts []Securityhub_InsightFiltersResourceAwsIamAccessKeyCreatedAt `json:"resourceAwsIamAccessKeyCreatedAts,omitempty" yaml:"resourceAwsIamAccessKeyCreatedAts,omitempty"`

	// The display name of the owner of the S3 bucket. See String Filter below for more details.
	ResourceAwsS3BucketOwnerNames []Securityhub_InsightFiltersResourceAwsS3BucketOwnerName `json:"resourceAwsS3BucketOwnerNames,omitempty" yaml:"resourceAwsS3BucketOwnerNames,omitempty"`

	// The date/time that the container was started. See Date Filter below for more details.
	ResourceContainerLaunchedAts []Securityhub_InsightFiltersResourceContainerLaunchedAt `json:"resourceContainerLaunchedAts,omitempty" yaml:"resourceContainerLaunchedAts,omitempty"`

	// The date/time of the last observation of a threat intelligence indicator. See Date Filter below for more details.
	ThreatIntelIndicatorLastObservedAts []Securityhub_InsightFiltersThreatIntelIndicatorLastObservedAt `json:"threatIntelIndicatorLastObservedAts,omitempty" yaml:"threatIntelIndicatorLastObservedAts,omitempty"`

	// The source of the threat intelligence. See String Filter below for more details.
	ThreatIntelIndicatorSources []Securityhub_InsightFiltersThreatIntelIndicatorSource `json:"threatIntelIndicatorSources,omitempty" yaml:"threatIntelIndicatorSources,omitempty"`

	// The destination IPv6 address of network-related information about a finding. See Ip Filter below for more details.
	NetworkDestinationIpv6s []Securityhub_InsightFiltersNetworkDestinationIpv6 `json:"networkDestinationIpv6s,omitempty" yaml:"networkDestinationIpv6s,omitempty"`

	// The source media access control (MAC) address of network-related information about a finding. See String Filter below for more details.
	NetworkSourceMacs []Securityhub_InsightFiltersNetworkSourceMac `json:"networkSourceMacs,omitempty" yaml:"networkSourceMacs,omitempty"`

	// The text of a note. See String Filter below for more details.
	NoteTexts []Securityhub_InsightFiltersNoteText `json:"noteTexts,omitempty" yaml:"noteTexts,omitempty"`

	// The solution-generated identifier for a related finding. See String Filter below for more details.
	RelatedFindingsIds []Securityhub_InsightFiltersRelatedFindingsId `json:"relatedFindingsIds,omitempty" yaml:"relatedFindingsIds,omitempty"`

	// The IAM profile ARN of the instance. See String Filter below for more details.
	ResourceAwsEc2InstanceIamInstanceProfileArns []Securityhub_InsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn `json:"resourceAwsEc2InstanceIamInstanceProfileArns,omitempty" yaml:"resourceAwsEc2InstanceIamInstanceProfileArns,omitempty"`

	// The key name associated with the instance. See String Filter below for more details.
	ResourceAwsEc2InstanceKeyNames []Securityhub_InsightFiltersResourceAwsEc2InstanceKeyName `json:"resourceAwsEc2InstanceKeyNames,omitempty" yaml:"resourceAwsEc2InstanceKeyNames,omitempty"`

	// The status of the IAM access key related to a finding. See String Filter below for more details.
	ResourceAwsIamAccessKeyStatuses []Securityhub_InsightFiltersResourceAwsIamAccessKeyStatus `json:"resourceAwsIamAccessKeyStatuses,omitempty" yaml:"resourceAwsIamAccessKeyStatuses,omitempty"`

	// The finding provider's original value for the severity. See String Filter below for more details.
	FindingProviderFieldsSeverityOriginals []Securityhub_InsightFiltersFindingProviderFieldsSeverityOriginal `json:"findingProviderFieldsSeverityOriginals,omitempty" yaml:"findingProviderFieldsSeverityOriginals,omitempty"`

	// One or more finding types that the finding provider assigned to the finding. Uses the format of `namespace/category/classifier` that classify a finding. Valid namespace values include: `Software and Configuration Checks`, `TTPs`, `Effects`, `Unusual Behaviors`, and `Sensitive Data Identifications`. See String Filter below for more details.
	FindingProviderFieldsTypes []Securityhub_InsightFiltersFindingProviderFieldsType `json:"findingProviderFieldsTypes,omitempty" yaml:"findingProviderFieldsTypes,omitempty"`

	// The protocol of network-related information about a finding. See String Filter below for more details.
	NetworkProtocols []Securityhub_InsightFiltersNetworkProtocol `json:"networkProtocols,omitempty" yaml:"networkProtocols,omitempty"`

	// The finding provider value for the level of importance assigned to the resources associated with the findings. A score of 0 means that the underlying resources have no criticality, and a score of 100 is reserved for the most critical resources. See Number Filter below for more details.
	FindingProviderFieldsCriticalities []Securityhub_InsightFiltersFindingProviderFieldsCriticality `json:"findingProviderFieldsCriticalities,omitempty" yaml:"findingProviderFieldsCriticalities,omitempty"`

	// The destination port of network-related information about a finding. See Number Filter below for more details.
	NetworkDestinationPorts []Securityhub_InsightFiltersNetworkDestinationPort `json:"networkDestinationPorts,omitempty" yaml:"networkDestinationPorts,omitempty"`

	// The source port of network-related information about a finding. See Number Filter below for more details.
	NetworkSourcePorts []Securityhub_InsightFiltersNetworkSourcePort `json:"networkSourcePorts,omitempty" yaml:"networkSourcePorts,omitempty"`

	// A data type where security-findings providers can include additional solution-specific details that aren't part of the defined `AwsSecurityFinding` format. See Map Filter below for more details.
	ProductFields []Securityhub_InsightFiltersProductField `json:"productFields,omitempty" yaml:"productFields,omitempty"`

	// The ARN of the solution that generated a related finding. See String Filter below for more details.
	RelatedFindingsProductArns []Securityhub_InsightFiltersRelatedFindingsProductArn `json:"relatedFindingsProductArns,omitempty" yaml:"relatedFindingsProductArns,omitempty"`

	// A finding's title. See String Filter below for more details.
	Titles []Securityhub_InsightFiltersTitle `json:"titles,omitempty" yaml:"titles,omitempty"`

	// The name of the malware that was observed. See String Filter below for more details.
	MalwareNames []Securityhub_InsightFiltersMalwareName `json:"malwareNames,omitempty" yaml:"malwareNames,omitempty"`

	// The date/time that the process was terminated. See Date Filter below for more details.
	ProcessTerminatedAts []Securityhub_InsightFiltersProcessTerminatedAt `json:"processTerminatedAts,omitempty" yaml:"processTerminatedAts,omitempty"`

	// The name of the findings provider (company) that owns the solution (product) that generates findings. See String_Filter below for more details.
	CompanyNames []Securityhub_InsightFiltersCompanyName `json:"companyNames,omitempty" yaml:"companyNames,omitempty"`

	// An ISO8601-formatted timestamp that indicates when the security-findings provider captured the potential security issue that a finding captured. See Date Filter below for more details.
	CreatedAts []Securityhub_InsightFiltersCreatedAt `json:"createdAts,omitempty" yaml:"createdAts,omitempty"`

	// The value of a threat intelligence indicator. See String Filter below for more details.
	ThreatIntelIndicatorValues []Securityhub_InsightFiltersThreatIntelIndicatorValue `json:"threatIntelIndicatorValues,omitempty" yaml:"threatIntelIndicatorValues,omitempty"`
}
