package types

type Appconfig_ExtensionActionPoint struct {
	// An action defines the tasks the extension performs during the AppConfig workflow. Detailed below.
	Actions []Appconfig_ExtensionActionPointAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	// The point at which to perform the defined actions. Valid points are `PRE_CREATE_HOSTED_CONFIGURATION_VERSION`, `PRE_START_DEPLOYMENT`, `ON_DEPLOYMENT_START`, `ON_DEPLOYMENT_STEP`, `ON_DEPLOYMENT_BAKING`, `ON_DEPLOYMENT_COMPLETE`, `ON_DEPLOYMENT_ROLLED_BACK`.
	Point string `json:"point,omitempty" yaml:"point,omitempty"`
}