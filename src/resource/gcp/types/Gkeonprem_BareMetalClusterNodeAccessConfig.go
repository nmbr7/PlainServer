package types

type Gkeonprem_BareMetalClusterNodeAccessConfig struct {
	/*
	   LoginUser is the user name used to access node machines.
	   It defaults to "root" if not set.
	*/
	LoginUser string `json:"loginUser,omitempty" yaml:"loginUser,omitempty"`
}
