package types

type Rds_ProxyDefaultTargetGroupConnectionPoolConfig struct {
	// One or more SQL statements for the proxy to run when opening each new database connection. Typically used with `SET` statements to make sure that each connection has identical settings such as time zone and character set. This setting is empty by default. For multiple statements, use semicolons as the separator. You can also include multiple variables in a single `SET` statement, such as `SET x=1, y=2`.
	InitQuery string `json:"initQuery,omitempty" yaml:"initQuery,omitempty"`

	// The maximum size of the connection pool for each target in a target group. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	MaxConnectionsPercent int `json:"maxConnectionsPercent,omitempty" yaml:"maxConnectionsPercent,omitempty"`

	// Controls how actively the proxy closes idle database connections in the connection pool. A high value enables the proxy to leave a high percentage of idle connections open. A low value causes the proxy to close idle client connections and return the underlying database connections to the connection pool. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	MaxIdleConnectionsPercent int `json:"maxIdleConnectionsPercent,omitempty" yaml:"maxIdleConnectionsPercent,omitempty"`

	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection. Including an item in the list exempts that class of SQL operations from the pinning behavior. Currently, the only allowed value is `EXCLUDE_VARIABLE_SETS`.
	SessionPinningFilters []string `json:"sessionPinningFilters,omitempty" yaml:"sessionPinningFilters,omitempty"`

	// The number of seconds for a proxy to wait for a connection to become available in the connection pool. Only applies when the proxy has opened its maximum number of connections and all connections are busy with client sessions.
	ConnectionBorrowTimeout int `json:"connectionBorrowTimeout,omitempty" yaml:"connectionBorrowTimeout,omitempty"`
}
