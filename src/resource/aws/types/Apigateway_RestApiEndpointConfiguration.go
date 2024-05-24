package types

type Apigateway_RestApiEndpointConfiguration struct {
	// List of endpoint  This resource currently only supports managing a single value. Valid values: `EDGE`, `REGIONAL` or `PRIVATE`. If unspecified, defaults to `EDGE`. If set to `PRIVATE` recommend to set `put_rest_api_mode` = `merge` to not cause the endpoints and associated Route53 records to be deleted. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
	Types string `json:"types,omitempty" yaml:"types,omitempty"`

	// Set of VPC Endpoint identifiers. It is only supported for `PRIVATE` endpoint type. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-endpoint-configuration` extension `vpcEndpointIds` property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-endpoint-configuration.html). If the argument value is provided and is different than the OpenAPI value, --the argument value will override the OpenAPI value--.
	VpcEndpointIds []string `json:"vpcEndpointIds,omitempty" yaml:"vpcEndpointIds,omitempty"`
}
