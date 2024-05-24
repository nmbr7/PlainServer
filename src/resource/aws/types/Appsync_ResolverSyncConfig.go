package types

type Appsync_ResolverSyncConfig struct {
	// Conflict Detection strategy to use. Valid values are `NONE` and `VERSION`.
	ConflictDetection string `json:"conflictDetection,omitempty" yaml:"conflictDetection,omitempty"`

	// Conflict Resolution strategy to perform in the event of a conflict. Valid values are `NONE`, `OPTIMISTIC_CONCURRENCY`, `AUTOMERGE`, and `LAMBDA`.
	ConflictHandler string `json:"conflictHandler,omitempty" yaml:"conflictHandler,omitempty"`

	// Lambda Conflict Handler Config when configuring `LAMBDA` as the Conflict Handler. See Lambda Conflict Handler Config.
	LambdaConflictHandlerConfig Appsync_ResolverSyncConfigLambdaConflictHandlerConfig `json:"lambdaConflictHandlerConfig,omitempty" yaml:"lambdaConflictHandlerConfig,omitempty"`
}
