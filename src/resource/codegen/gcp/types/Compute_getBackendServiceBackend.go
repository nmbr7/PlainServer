package types

type Compute_getBackendServiceBackend struct {
	/*
	   Specifies the balancing mode for this backend.

	   For global HTTP(S) or TCP/SSL load balancing, the default is
	   UTILIZATION. Valid values are UTILIZATION, RATE (for HTTP(S))
	   and CONNECTION (for TCP/SSL).

	   See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
	   for an explanation of load balancing modes. Default value: "UTILIZATION" Possible values: ["UTILIZATION", "RATE", "CONNECTION"]
	*/
	BalancingMode string `json:"balancingMode,omitempty" yaml:"balancingMode,omitempty"`

	/*
	   A multiplier applied to the group's maximum servicing capacity
	   (based on UTILIZATION, RATE or CONNECTION).

	   Default value is 1, which means the group will serve up to 100%!!(MISSING)
	   (MISSING)of its configured capacity (depending on balancingMode). A
	   setting of 0 means the group is completely drained, offering
	   0%!!(MISSING)o(MISSING)f its available Capacity. Valid range is [0.0,1.0].
	*/
	CapacityScaler float64 `json:"capacityScaler,omitempty" yaml:"capacityScaler,omitempty"`

	/*
	   The max number of simultaneous connections that a single
	   backend instance can handle. This is used to calculate the
	   capacity of the group. Can be used in either CONNECTION or
	   UTILIZATION balancing modes.

	   For CONNECTION mode, either maxConnections or
	   maxConnectionsPerInstance must be set.
	*/
	MaxConnectionsPerInstance int `json:"maxConnectionsPerInstance,omitempty" yaml:"maxConnectionsPerInstance,omitempty"`

	/*
	   The max requests per second (RPS) that a single backend
	   instance can handle. This is used to calculate the capacity of
	   the group. Can be used in either balancing mode. For RATE mode,
	   either maxRate or maxRatePerInstance must be set.
	*/
	MaxRatePerInstance float64 `json:"maxRatePerInstance,omitempty" yaml:"maxRatePerInstance,omitempty"`

	/*
	   Used when balancingMode is UTILIZATION. This ratio defines the
	   CPU utilization target for the group. Valid range is [0.0, 1.0].
	*/
	MaxUtilization float64 `json:"maxUtilization,omitempty" yaml:"maxUtilization,omitempty"`

	// Textual description for the Backend Service.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The fully-qualified URL of an Instance Group or Network Endpoint
	   Group resource. In case of instance group this defines the list
	   of instances that serve traffic. Member virtual machine
	   instances from each instance group must live in the same zone as
	   the instance group itself. No two backends in a backend service
	   are allowed to use same Instance Group resource.

	   For Network Endpoint Groups this defines list of endpoints. All
	   endpoints of Network Endpoint Group must be hosted on instances
	   located in the same zone as the Network Endpoint Group.

	   Backend services cannot mix Instance Group and
	   Network Endpoint Group backends.

	   Note that you must specify an Instance Group or Network Endpoint
	   Group resource using the fully-qualified URL, rather than a
	   partial URL.
	*/
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	/*
	   The max number of simultaneous connections for the group. Can
	   be used with either CONNECTION or UTILIZATION balancing modes.

	   For CONNECTION mode, either maxConnections or one
	   of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
	   as appropriate for group type, must be set.
	*/
	MaxConnections int `json:"maxConnections,omitempty" yaml:"maxConnections,omitempty"`

	/*
	   The max number of simultaneous connections that a single backend
	   network endpoint can handle. This is used to calculate the
	   capacity of the group. Can be used in either CONNECTION or
	   UTILIZATION balancing modes.

	   For CONNECTION mode, either
	   maxConnections or maxConnectionsPerEndpoint must be set.
	*/
	MaxConnectionsPerEndpoint int `json:"maxConnectionsPerEndpoint,omitempty" yaml:"maxConnectionsPerEndpoint,omitempty"`

	/*
	   The max requests per second (RPS) of the group.

	   Can be used with either RATE or UTILIZATION balancing modes,
	   but required if RATE mode. For RATE mode, either maxRate or one
	   of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
	   group type, must be set.
	*/
	MaxRate int `json:"maxRate,omitempty" yaml:"maxRate,omitempty"`

	/*
	   The max requests per second (RPS) that a single backend network
	   endpoint can handle. This is used to calculate the capacity of
	   the group. Can be used in either balancing mode. For RATE mode,
	   either maxRate or maxRatePerEndpoint must be set.
	*/
	MaxRatePerEndpoint float64 `json:"maxRatePerEndpoint,omitempty" yaml:"maxRatePerEndpoint,omitempty"`
}
