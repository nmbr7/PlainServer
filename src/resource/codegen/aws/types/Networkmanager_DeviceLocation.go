package types

type Networkmanager_DeviceLocation struct {
	// The latitude.
	Latitude string `json:"latitude,omitempty" yaml:"latitude,omitempty"`

	// The longitude.
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty"`

	// The physical address.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
}
