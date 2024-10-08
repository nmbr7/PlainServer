package types

type Compute_RegionBackendServiceBackend struct {
	/*
	   A multiplier applied to the group's maximum servicing capacity
	   (based on UTILIZATION, RATE or CONNECTION).
	   ~>--NOTE--: This field cannot be set for
	   INTERNAL region backend services (default loadBalancingScheme),
	   but is required for non-INTERNAL backend service. The total
	   capacity_scaler for all backends must be non-zero.
	   A setting of 0 means the group is completely drained, offering
	   0%!!(MISSING)o(MISSING)f its available Capacity. Valid range is [0.0,1.0].
	*/
	CapacityScaler float64 `json:"capacityScaler,omitempty" yaml:"capacityScaler,omitempty"`

	/*
	   An optional description of this resource.
	   Provide this property when you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The max number of simultaneous connections that a single
	   backend instance can handle. Cannot be set for INTERNAL backend
	   services.
	   This is used to calculate the capacity of the group.
	   Can be used in either CONNECTION or UTILIZATION balancing modes.
	   For CONNECTION mode, either maxConnections or
	   maxConnectionsPerInstance must be set.
	*/
	MaxConnectionsPerInstance int `json:"maxConnectionsPerInstance,omitempty" yaml:"maxConnectionsPerInstance,omitempty"`

	/*
	   The max requests per second (RPS) of the group. Cannot be set
	   for INTERNAL backend services.
	   Can be used with either RATE or UTILIZATION balancing modes,
	   but required if RATE mode. Either maxRate or one
	   of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
	   group type, must be set.
	*/
	MaxRate int `json:"maxRate,omitempty" yaml:"maxRate,omitempty"`

	/*
	   Specifies the balancing mode for this backend.
	   See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
	   for an explanation of load balancing modes.
	   Default value is `CONNECTION`.
	   Possible values are: `UTILIZATION`, `RATE`, `CONNECTION`.
	*/
	BalancingMode string `json:"balancingMode,omitempty" yaml:"balancingMode,omitempty"`

	/*
	   This field designates whether this is a failover backend. More
	   than one failover backend can be configured for a given RegionBackendService.
	*/
	Failover bool `json:"failover,omitempty" yaml:"failover,omitempty"`

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
	   When the `load_balancing_scheme` is INTERNAL, only instance groups
	   are supported.
	   Note that you must specify an Instance Group or Network Endpoint
	   Group resource using the fully-qualified URL, rather than a
	   partial URL.
	*/
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	/*
	   The max number of simultaneous connections for the group. Can
	   be used with either CONNECTION or UTILIZATION balancing modes.
	   Cannot be set for INTERNAL backend services.
	   For CONNECTION mode, either maxConnections or one
	   of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
	   as appropriate for group type, must be set.
	*/
	MaxConnections int `json:"maxConnections,omitempty" yaml:"maxConnections,omitempty"`

	/*
	   The max number of simultaneous connections that a single backend
	   network endpoint can handle. Cannot be set
	   for INTERNAL backend services.
	   This is used to calculate the capacity of the group. Can be
	   used in either CONNECTION or UTILIZATION balancing modes. For
	   CONNECTION mode, either maxConnections or
	   maxConnectionsPerEndpoint must be set.
	*/
	MaxConnectionsPerEndpoint int `json:"maxConnectionsPerEndpoint,omitempty" yaml:"maxConnectionsPerEndpoint,omitempty"`

	/*
	   The max requests per second (RPS) that a single backend network
	   endpoint can handle. This is used to calculate the capacity of
	   the group. Can be used in either balancing mode. For RATE mode,
	   either maxRate or maxRatePerEndpoint must be set. Cannot be set
	   for INTERNAL backend services.
	*/
	MaxRatePerEndpoint float64 `json:"maxRatePerEndpoint,omitempty" yaml:"maxRatePerEndpoint,omitempty"`

	/*
	   The max requests per second (RPS) that a single backend
	   instance can handle. This is used to calculate the capacity of
	   the group. Can be used in either balancing mode. For RATE mode,
	   either maxRate or maxRatePerInstance must be set. Cannot be set
	   for INTERNAL backend services.
	*/
	MaxRatePerInstance float64 `json:"maxRatePerInstance,omitempty" yaml:"maxRatePerInstance,omitempty"`

	/*
	   Used when balancingMode is UTILIZATION. This ratio defines the
	   CPU utilization target for the group. Valid range is [0.0, 1.0].
	   Cannot be set for INTERNAL backend services.
	*/
	MaxUtilization float64 `json:"maxUtilization,omitempty" yaml:"maxUtilization,omitempty"`
}
