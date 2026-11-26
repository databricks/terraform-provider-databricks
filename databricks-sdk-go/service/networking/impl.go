// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networking

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Endpoints API methods
type endpointsImpl struct {
	client *client.DatabricksClient
}

func (a *endpointsImpl) CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/networking/v1/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Endpoint, &endpoint)
	return &endpoint, err
}

func (a *endpointsImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error {
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *endpointsImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &endpoint)
	return &endpoint, err
}

// Lists all network connectivity endpoints for the account.
func (a *endpointsImpl) ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[Endpoint] {

	getNextPage := func(ctx context.Context, req ListEndpointsRequest) (*ListEndpointsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListEndpoints(ctx, req)
	}
	getItems := func(resp *ListEndpointsResponse) []Endpoint {
		return resp.Items
	}
	getNextReq := func(resp *ListEndpointsResponse) *ListEndpointsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists all network connectivity endpoints for the account.
func (a *endpointsImpl) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]Endpoint, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[Endpoint](ctx, iterator)
}

func (a *endpointsImpl) internalListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := fmt.Sprintf("/api/networking/v1/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listEndpointsResponse)
	return &listEndpointsResponse, err
}

// unexported type that holds implementations of just PrivateNetworkGateways API methods
type privateNetworkGatewaysImpl struct {
	client *client.DatabricksClient
}

func (a *privateNetworkGatewaysImpl) CreatePrivateNetworkGateway(ctx context.Context, request CreatePrivateNetworkGatewayRequest) (*Operation, error) {
	var operation Operation
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := fmt.Sprintf("/api/networking/v1/%v/private-network-gateways", request.Parent)
	queryParams := make(map[string]any)

	if request.RequestId != "" || slices.Contains(request.ForceSendFields, "RequestId") {
		queryParams["request_id"] = request.RequestId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.PrivateNetworkGateway, &operation)
	return &operation, err
}

func (a *privateNetworkGatewaysImpl) DeletePrivateNetworkGateway(ctx context.Context, request DeletePrivateNetworkGatewayRequest) error {
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *privateNetworkGatewaysImpl) GetPrivateNetworkGateway(ctx context.Context, request GetPrivateNetworkGatewayRequest) (*PrivateNetworkGateway, error) {
	var privateNetworkGateway PrivateNetworkGateway
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &privateNetworkGateway)
	return &privateNetworkGateway, err
}

func (a *privateNetworkGatewaysImpl) GetPrivateNetworkGatewayOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &operation)
	return &operation, err
}

// Lists private network gateways under a network connectivity configuration.
func (a *privateNetworkGatewaysImpl) ListPrivateNetworkGateways(ctx context.Context, request ListPrivateNetworkGatewaysRequest) listing.Iterator[PrivateNetworkGateway] {

	getNextPage := func(ctx context.Context, req ListPrivateNetworkGatewaysRequest) (*ListPrivateNetworkGatewaysResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListPrivateNetworkGateways(ctx, req)
	}
	getItems := func(resp *ListPrivateNetworkGatewaysResponse) []PrivateNetworkGateway {
		return resp.PrivateNetworkGateways
	}
	getNextReq := func(resp *ListPrivateNetworkGatewaysResponse) *ListPrivateNetworkGatewaysRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists private network gateways under a network connectivity configuration.
func (a *privateNetworkGatewaysImpl) ListPrivateNetworkGatewaysAll(ctx context.Context, request ListPrivateNetworkGatewaysRequest) ([]PrivateNetworkGateway, error) {
	iterator := a.ListPrivateNetworkGateways(ctx, request)
	return listing.ToSlice[PrivateNetworkGateway](ctx, iterator)
}

func (a *privateNetworkGatewaysImpl) internalListPrivateNetworkGateways(ctx context.Context, request ListPrivateNetworkGatewaysRequest) (*ListPrivateNetworkGatewaysResponse, error) {
	var listPrivateNetworkGatewaysResponse ListPrivateNetworkGatewaysResponse
	path := fmt.Sprintf("/api/networking/v1/%v/private-network-gateways", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listPrivateNetworkGatewaysResponse)
	return &listPrivateNetworkGatewaysResponse, err
}

func (a *privateNetworkGatewaysImpl) UpdatePrivateNetworkGateway(ctx context.Context, request UpdatePrivateNetworkGatewayRequest) (*PrivateNetworkGateway, error) {
	var privateNetworkGateway PrivateNetworkGateway
	path := fmt.Sprintf("/api/networking/v1/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.PrivateNetworkGateway, &privateNetworkGateway)
	return &privateNetworkGateway, err
}
