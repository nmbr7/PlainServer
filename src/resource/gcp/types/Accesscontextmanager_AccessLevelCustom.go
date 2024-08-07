package types

type Accesscontextmanager_AccessLevelCustom struct {
	/*
	   Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
	   This page details the objects and attributes that are used to the build the CEL expressions for
	   custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.
	   Structure is documented below.
	*/
	Expr Accesscontextmanager_AccessLevelCustomExpr `json:"expr,omitempty" yaml:"expr,omitempty"`
}
