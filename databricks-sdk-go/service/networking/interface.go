// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networking

import (
	"context"
)

// These APIs manage endpoint configurations for this account.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EndpointsService interface {

	// Creates a new network connectivity endpoint that enables private
	// connectivity between your network resources and Databricks services.
	//
	// After creation, the endpoint is initially in the PENDING state. The
	// Databricks endpoint service automatically reviews and approves the
	// endpoint within a few minutes. Use the GET method to retrieve the latest
	// endpoint state.
	//
	// An endpoint can be used only after it reaches the APPROVED state.
	CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Endpoint, error)

	// Deletes a network endpoint. This will remove the endpoint configuration
	// from Databricks. Depending on the endpoint type and use case, you may
	// also need to delete corresponding network resources in your cloud
	// provider account.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error

	// Gets details of a specific network endpoint.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error)

	// Lists all network connectivity endpoints for the account.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error)
}

// These APIs manage private network gateways under network connectivity
// configurations.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PrivateNetworkGatewaysService interface {

	// Creates a private network gateway.
	CreatePrivateNetworkGateway(ctx context.Context, request CreatePrivateNetworkGatewayRequest) (*Operation, error)

	// Permanently deletes a private network gateway.
	DeletePrivateNetworkGateway(ctx context.Context, request DeletePrivateNetworkGatewayRequest) error

	// Gets a private network gateway.
	GetPrivateNetworkGateway(ctx context.Context, request GetPrivateNetworkGatewayRequest) (*PrivateNetworkGateway, error)

	// Gets the status of a private network gateway create operation.
	GetPrivateNetworkGatewayOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Lists private network gateways under a network connectivity
	// configuration.
	ListPrivateNetworkGateways(ctx context.Context, request ListPrivateNetworkGatewaysRequest) (*ListPrivateNetworkGatewaysResponse, error)

	// Updates a private network gateway.
	UpdatePrivateNetworkGateway(ctx context.Context, request UpdatePrivateNetworkGatewayRequest) (*PrivateNetworkGateway, error)
}
