package types

type Compute_SecurityPolicyRuleMatchExpr struct {
	/*
	   Textual representation of an expression in Common Expression Language syntax.
	   The application context of the containing message determines which well-known feature set of CEL is supported.
	*/
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`
}
