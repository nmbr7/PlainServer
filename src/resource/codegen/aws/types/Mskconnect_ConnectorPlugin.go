package types

type Mskconnect_ConnectorPlugin struct {
	// Details about a custom plugin. See `custom_plugin` Block for details.
	CustomPlugin Mskconnect_ConnectorPluginCustomPlugin `json:"customPlugin,omitempty" yaml:"customPlugin,omitempty"`
}
