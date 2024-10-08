package types

type Workspaces_DirectoryWorkspaceAccessProperties struct {
	// Indicates whether users can access their WorkSpaces through a web browser.
	DeviceTypeWeb string `json:"deviceTypeWeb,omitempty" yaml:"deviceTypeWeb,omitempty"`

	// Indicates whether users can use Windows clients to access their WorkSpaces.
	DeviceTypeWindows string `json:"deviceTypeWindows,omitempty" yaml:"deviceTypeWindows,omitempty"`

	// Indicates whether users can use zero client devices to access their WorkSpaces.
	DeviceTypeZeroclient string `json:"deviceTypeZeroclient,omitempty" yaml:"deviceTypeZeroclient,omitempty"`

	// Indicates whether users can use Android devices to access their WorkSpaces.
	DeviceTypeAndroid string `json:"deviceTypeAndroid,omitempty" yaml:"deviceTypeAndroid,omitempty"`

	// Indicates whether users can use Chromebooks to access their WorkSpaces.
	DeviceTypeChromeos string `json:"deviceTypeChromeos,omitempty" yaml:"deviceTypeChromeos,omitempty"`

	// Indicates whether users can use iOS devices to access their WorkSpaces.
	DeviceTypeIos string `json:"deviceTypeIos,omitempty" yaml:"deviceTypeIos,omitempty"`

	// Indicates whether users can use Linux clients to access their WorkSpaces.
	DeviceTypeLinux string `json:"deviceTypeLinux,omitempty" yaml:"deviceTypeLinux,omitempty"`

	// Indicates whether users can use macOS clients to access their WorkSpaces.
	DeviceTypeOsx string `json:"deviceTypeOsx,omitempty" yaml:"deviceTypeOsx,omitempty"`
}
