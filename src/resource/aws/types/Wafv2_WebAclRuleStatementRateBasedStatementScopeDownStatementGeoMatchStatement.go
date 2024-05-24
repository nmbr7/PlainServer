package types

type Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement struct {
	// Configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. See `forwarded_ip_config` below for details.
	ForwardedIpConfig Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatementForwardedIpConfig `json:"forwardedIpConfig,omitempty" yaml:"forwardedIpConfig,omitempty"`

	// Array of two-character country codes, for example, [ "US", "CN" ], from the alpha-2 country ISO codes of the `ISO 3166` international standard. See the [documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_GeoMatchStatement.html) for valid values.
	CountryCodes []string `json:"countryCodes,omitempty" yaml:"countryCodes,omitempty"`
}
