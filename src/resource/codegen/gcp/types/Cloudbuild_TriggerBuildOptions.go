package types

type Cloudbuild_TriggerBuildOptions struct {
	/*
	   Option to specify behavior when there is an error in the substitution checks.
	   NOTE this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden
	   in the build configuration file.
	   Possible values are: `MUST_MATCH`, `ALLOW_LOOSE`.
	*/
	SubstitutionOption string `json:"substitutionOption,omitempty" yaml:"substitutionOption,omitempty"`

	/*
	   Global list of volumes to mount for ALL build steps
	   Each volume is created as an empty volume prior to starting the build process.
	   Upon completion of the build, volumes and their contents are discarded. Global
	   volume names and paths cannot conflict with the volumes defined a build step.
	   Using a global volume in a build with only one step is not valid as it is indicative
	   of a build request with an incorrect configuration.
	   Structure is documented below.
	*/
	Volumes []Cloudbuild_TriggerBuildOptionsVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	/*
	   Requested disk size for the VM that runs the build. Note that this is NOT "disk free";
	   some of the space will be used by the operating system and build utilities.
	   Also note that this is the minimum disk size that will be allocated for the build --
	   the build may run with a larger disk than requested. At present, the maximum disk size
	   is 1000GB; builds that request more than the maximum are rejected with an error.
	*/
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	// Compute Engine machine type on which to run the build.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	/*
	   Requested hash for SourceProvenance.
	   Each value may be one of: `NONE`, `SHA256`, `MD5`.
	*/
	SourceProvenanceHashes []string `json:"sourceProvenanceHashes,omitempty" yaml:"sourceProvenanceHashes,omitempty"`

	/*
	   Option to specify the logging mode, which determines if and where build logs are stored.
	   Possible values are: `LOGGING_UNSPECIFIED`, `LEGACY`, `GCS_ONLY`, `STACKDRIVER_ONLY`, `CLOUD_LOGGING_ONLY`, `NONE`.
	*/
	Logging string `json:"logging,omitempty" yaml:"logging,omitempty"`

	/*
	   Requested verifiability options.
	   Possible values are: `NOT_VERIFIED`, `VERIFIED`.
	*/
	RequestedVerifyOption string `json:"requestedVerifyOption,omitempty" yaml:"requestedVerifyOption,omitempty"`

	/*
	   A list of global environment variables, which are encrypted using a Cloud Key Management
	   Service crypto key. These values must be specified in the build's Secret. These variables
	   will be available to all build steps in this build.
	*/
	SecretEnvs []string `json:"secretEnvs,omitempty" yaml:"secretEnvs,omitempty"`

	/*
	   Option to specify a WorkerPool for the build. Format projects/{project}/workerPools/{workerPool}
	   This field is experimental.
	*/
	WorkerPool string `json:"workerPool,omitempty" yaml:"workerPool,omitempty"`

	/*
	   Option to specify whether or not to apply bash style string operations to the substitutions.
	   NOTE this is always enabled for triggered builds and cannot be overridden in the build configuration file.
	*/
	DynamicSubstitutions bool `json:"dynamicSubstitutions,omitempty" yaml:"dynamicSubstitutions,omitempty"`

	/*
	   A list of global environment variable definitions that will exist for all build steps
	   in this build. If a variable is defined in both globally and in a build step,
	   the variable will use the build step value.
	   The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
	*/
	Envs []string `json:"envs,omitempty" yaml:"envs,omitempty"`

	/*
	   Option to define build log streaming behavior to Google Cloud Storage.
	   Possible values are: `STREAM_DEFAULT`, `STREAM_ON`, `STREAM_OFF`.
	*/
	LogStreamingOption string `json:"logStreamingOption,omitempty" yaml:"logStreamingOption,omitempty"`
}
