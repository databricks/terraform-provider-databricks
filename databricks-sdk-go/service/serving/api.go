// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Serving Endpoints, Serving Endpoints Data Plane, etc.
package serving

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/dataplane"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type ServingEndpointsInterface interface {

	// WaitGetServingEndpointNotUpdating repeatedly calls [ServingEndpointsAPI.Get] and waits to reach NOT_UPDATING state
	WaitGetServingEndpointNotUpdating(ctx context.Context, name string,
		timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error)

	// Retrieves the build logs associated with the provided served model.
	BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error)

	// Retrieves the build logs associated with the provided served model.
	BuildLogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*BuildLogsResponse, error)

	// Create a new serving endpoint.
	Create(ctx context.Context, createServingEndpoint CreateServingEndpoint) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)

	// Calls [ServingEndpointsAPIInterface.Create] and waits to reach NOT_UPDATING state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
	//
	// Deprecated: use [ServingEndpointsAPIInterface.Create].Get() or [ServingEndpointsAPIInterface.WaitGetServingEndpointNotUpdating]
	CreateAndWait(ctx context.Context, createServingEndpoint CreateServingEndpoint, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error)

	// Create a new PT serving endpoint.
	CreateProvisionedThroughputEndpoint(ctx context.Context, createPtEndpointRequest CreatePtEndpointRequest) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)

	// Calls [ServingEndpointsAPIInterface.CreateProvisionedThroughputEndpoint] and waits to reach NOT_UPDATING state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
	//
	// Deprecated: use [ServingEndpointsAPIInterface.CreateProvisionedThroughputEndpoint].Get() or [ServingEndpointsAPIInterface.WaitGetServingEndpointNotUpdating]
	CreateProvisionedThroughputEndpointAndWait(ctx context.Context, createPtEndpointRequest CreatePtEndpointRequest, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error)

	// Delete a serving endpoint.
	Delete(ctx context.Context, request DeleteServingEndpointRequest) error

	// Delete a serving endpoint.
	DeleteByName(ctx context.Context, name string) error

	// Retrieves the metrics associated with the provided serving endpoint in either
	// Prometheus or OpenMetrics exposition format.
	ExportMetrics(ctx context.Context, request ExportMetricsRequest) (*ExportMetricsResponse, error)

	// Retrieves the metrics associated with the provided serving endpoint in either
	// Prometheus or OpenMetrics exposition format.
	ExportMetricsByName(ctx context.Context, name string) (*ExportMetricsResponse, error)

	// Retrieves the details for a single serving endpoint.
	Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error)

	// Retrieves the details for a single serving endpoint.
	GetByName(ctx context.Context, name string) (*ServingEndpointDetailed, error)

	// Get the query schema of the serving endpoint in OpenAPI format. The schema
	// contains information for the supported paths, input and output format and
	// datatypes.
	GetOpenApi(ctx context.Context, request GetOpenApiRequest) (*GetOpenApiResponse, error)

	// Get the query schema of the serving endpoint in OpenAPI format. The schema
	// contains information for the supported paths, input and output format and
	// datatypes.
	GetOpenApiByName(ctx context.Context, name string) (*GetOpenApiResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetServingEndpointPermissionLevelsRequest) (*GetServingEndpointPermissionLevelsResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByServingEndpointId(ctx context.Context, servingEndpointId string) (*GetServingEndpointPermissionLevelsResponse, error)

	// Gets the permissions of a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)

	// Gets the permissions of a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	GetPermissionsByServingEndpointId(ctx context.Context, servingEndpointId string) (*ServingEndpointPermissions, error)

	// Make external services call using the credentials stored in UC Connection.
	HttpRequest(ctx context.Context, request ExternalFunctionRequest) (*HttpRequestResponse, error)

	// Get all serving endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context) listing.Iterator[ServingEndpoint]

	// Get all serving endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context) ([]ServingEndpoint, error)

	// Retrieves the service logs associated with the provided served model.
	Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error)

	// Retrieves the service logs associated with the provided served model.
	LogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*ServerLogsResponse, error)

	// Used to batch add and delete tags from a serving endpoint with a single API
	// call.
	Patch(ctx context.Context, request PatchServingEndpointTags) (*EndpointTags, error)

	// Updates the telemetry configuration of a serving endpoint.
	PatchTelemetryConfig(ctx context.Context, request PatchTelemetryConfigRequest) (*ServingEndpointDetailed, error)

	// Deprecated: Please use AI Gateway to manage rate limits instead.
	Put(ctx context.Context, request PutRequest) (*PutResponse, error)

	// Used to update the AI Gateway of a serving endpoint. NOTE: External model,
	// provisioned throughput, and pay-per-token endpoints are fully supported;
	// agent endpoints currently only support inference tables.
	PutAiGateway(ctx context.Context, request PutAiGatewayRequest) (*PutAiGatewayResponse, error)

	// Query a serving endpoint
	Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error)

	// Sets permissions on an object, replacing existing permissions if they exist.
	// Deletes all direct permissions if none are specified. Objects can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)

	// Updates any combination of the serving endpoint's served entities, the
	// compute configuration of those served entities, and the endpoint's traffic
	// config. An endpoint that already has an update in progress can not be updated
	// until the current update completes or fails.
	UpdateConfig(ctx context.Context, endpointCoreConfigInput EndpointCoreConfigInput) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)

	// Calls [ServingEndpointsAPIInterface.UpdateConfig] and waits to reach NOT_UPDATING state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
	//
	// Deprecated: use [ServingEndpointsAPIInterface.UpdateConfig].Get() or [ServingEndpointsAPIInterface.WaitGetServingEndpointNotUpdating]
	UpdateConfigAndWait(ctx context.Context, endpointCoreConfigInput EndpointCoreConfigInput, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error)

	// Updates the email and webhook notification settings for an endpoint.
	UpdateNotifications(ctx context.Context, request UpdateInferenceEndpointNotifications) (*UpdateInferenceEndpointNotificationsResponse, error)

	// Updates the permissions on a serving endpoint. Serving endpoints can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error)

	// Updates any combination of the pt endpoint's served entities, the compute
	// configuration of those served entities, and the endpoint's traffic config.
	// Updates are instantaneous and endpoint should be updated instantly
	UpdateProvisionedThroughputEndpointConfig(ctx context.Context, updateProvisionedThroughputEndpointConfigRequest UpdateProvisionedThroughputEndpointConfigRequest) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)

	// Calls [ServingEndpointsAPIInterface.UpdateProvisionedThroughputEndpointConfig] and waits to reach NOT_UPDATING state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
	//
	// Deprecated: use [ServingEndpointsAPIInterface.UpdateProvisionedThroughputEndpointConfig].Get() or [ServingEndpointsAPIInterface.WaitGetServingEndpointNotUpdating]
	UpdateProvisionedThroughputEndpointConfigAndWait(ctx context.Context, updateProvisionedThroughputEndpointConfigRequest UpdateProvisionedThroughputEndpointConfigRequest, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error)
}

func NewServingEndpoints(client *client.DatabricksClient) *ServingEndpointsAPI {
	return &ServingEndpointsAPI{
		servingEndpointsImpl: servingEndpointsImpl{
			client: client,
		},
	}
}

// The Serving Endpoints API allows you to create, update, and delete model
// serving endpoints.
//
// You can use a serving endpoint to serve models from the Databricks Model
// Registry or from Unity Catalog. Endpoints expose the underlying models as
// scalable REST API endpoints using serverless compute. This means the
// endpoints and associated compute resources are fully managed by Databricks
// and will not appear in your cloud account. A serving endpoint can consist of
// one or more MLflow models from the Databricks Model Registry, called served
// entities. A serving endpoint can have at most ten served entities. You can
// configure traffic settings to define how requests should be routed to your
// served entities behind an endpoint. Additionally, you can configure the scale
// of resources that should be applied to each served entity.
type ServingEndpointsAPI struct {
	servingEndpointsImpl
}

// WaitGetServingEndpointNotUpdating repeatedly calls [ServingEndpointsAPI.Get] and waits to reach NOT_UPDATING state
func (a *ServingEndpointsAPI) WaitGetServingEndpointNotUpdating(ctx context.Context, name string,
	timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[ServingEndpointDetailed](ctx, timeout, func() (*ServingEndpointDetailed, *retries.Err) {
		servingEndpointDetailed, err := a.Get(ctx, GetServingEndpointRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(servingEndpointDetailed)
		}
		status := servingEndpointDetailed.State.ConfigUpdate
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case EndpointStateConfigUpdateNotUpdating: // target state
			return servingEndpointDetailed, nil
		case EndpointStateConfigUpdateUpdateFailed, EndpointStateConfigUpdateUpdateCanceled:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				EndpointStateConfigUpdateNotUpdating, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetServingEndpointNotUpdating is a wrapper that calls [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating] and waits to reach NOT_UPDATING state.
type WaitGetServingEndpointNotUpdating[R any] struct {
	Response *R
	Name     string `json:"name"`
	Poll     func(time.Duration, func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error)
	callback func(*ServingEndpointDetailed)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetServingEndpointNotUpdating[R]) OnProgress(callback func(*ServingEndpointDetailed)) *WaitGetServingEndpointNotUpdating[R] {
	w.callback = callback
	return w
}

// Get the ServingEndpointDetailed with the default timeout of 20 minutes.
func (w *WaitGetServingEndpointNotUpdating[R]) Get() (*ServingEndpointDetailed, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the ServingEndpointDetailed with custom timeout.
func (w *WaitGetServingEndpointNotUpdating[R]) GetWithTimeout(timeout time.Duration) (*ServingEndpointDetailed, error) {
	return w.Poll(timeout, w.callback)
}

// Retrieves the build logs associated with the provided served model.
func (a *ServingEndpointsAPI) BuildLogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*BuildLogsResponse, error) {
	return a.servingEndpointsImpl.BuildLogs(ctx, BuildLogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Create a new serving endpoint.
func (a *ServingEndpointsAPI) Create(ctx context.Context, createServingEndpoint CreateServingEndpoint) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error) {
	servingEndpointDetailed, err := a.servingEndpointsImpl.Create(ctx, createServingEndpoint)
	if err != nil {
		return nil, err
	}
	return &WaitGetServingEndpointNotUpdating[ServingEndpointDetailed]{
		Response: servingEndpointDetailed,
		Name:     servingEndpointDetailed.Name,
		Poll: func(timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
			return a.WaitGetServingEndpointNotUpdating(ctx, servingEndpointDetailed.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ServingEndpointsAPI.Create] and waits to reach NOT_UPDATING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
//
// Deprecated: use [ServingEndpointsAPI.Create].Get() or [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating]
func (a *ServingEndpointsAPI) CreateAndWait(ctx context.Context, createServingEndpoint CreateServingEndpoint, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error) {
	wait, err := a.Create(ctx, createServingEndpoint)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ServingEndpointDetailed]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ServingEndpointDetailed) {
		for _, o := range options {
			o(&retries.Info[ServingEndpointDetailed]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Create a new PT serving endpoint.
func (a *ServingEndpointsAPI) CreateProvisionedThroughputEndpoint(ctx context.Context, createPtEndpointRequest CreatePtEndpointRequest) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error) {
	servingEndpointDetailed, err := a.servingEndpointsImpl.CreateProvisionedThroughputEndpoint(ctx, createPtEndpointRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetServingEndpointNotUpdating[ServingEndpointDetailed]{
		Response: servingEndpointDetailed,
		Name:     servingEndpointDetailed.Name,
		Poll: func(timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
			return a.WaitGetServingEndpointNotUpdating(ctx, servingEndpointDetailed.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ServingEndpointsAPI.CreateProvisionedThroughputEndpoint] and waits to reach NOT_UPDATING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
//
// Deprecated: use [ServingEndpointsAPI.CreateProvisionedThroughputEndpoint].Get() or [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating]
func (a *ServingEndpointsAPI) CreateProvisionedThroughputEndpointAndWait(ctx context.Context, createPtEndpointRequest CreatePtEndpointRequest, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error) {
	wait, err := a.CreateProvisionedThroughputEndpoint(ctx, createPtEndpointRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ServingEndpointDetailed]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ServingEndpointDetailed) {
		for _, o := range options {
			o(&retries.Info[ServingEndpointDetailed]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete a serving endpoint.
func (a *ServingEndpointsAPI) DeleteByName(ctx context.Context, name string) error {
	return a.servingEndpointsImpl.Delete(ctx, DeleteServingEndpointRequest{
		Name: name,
	})
}

// Retrieves the metrics associated with the provided serving endpoint in either
// Prometheus or OpenMetrics exposition format.
func (a *ServingEndpointsAPI) ExportMetricsByName(ctx context.Context, name string) (*ExportMetricsResponse, error) {
	return a.servingEndpointsImpl.ExportMetrics(ctx, ExportMetricsRequest{
		Name: name,
	})
}

// Retrieves the details for a single serving endpoint.
func (a *ServingEndpointsAPI) GetByName(ctx context.Context, name string) (*ServingEndpointDetailed, error) {
	return a.servingEndpointsImpl.Get(ctx, GetServingEndpointRequest{
		Name: name,
	})
}

// Get the query schema of the serving endpoint in OpenAPI format. The schema
// contains information for the supported paths, input and output format and
// datatypes.
func (a *ServingEndpointsAPI) GetOpenApiByName(ctx context.Context, name string) (*GetOpenApiResponse, error) {
	return a.servingEndpointsImpl.GetOpenApi(ctx, GetOpenApiRequest{
		Name: name,
	})
}

// Gets the permission levels that a user can have on an object.
func (a *ServingEndpointsAPI) GetPermissionLevelsByServingEndpointId(ctx context.Context, servingEndpointId string) (*GetServingEndpointPermissionLevelsResponse, error) {
	return a.servingEndpointsImpl.GetPermissionLevels(ctx, GetServingEndpointPermissionLevelsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Gets the permissions of a serving endpoint. Serving endpoints can inherit
// permissions from their root object.
func (a *ServingEndpointsAPI) GetPermissionsByServingEndpointId(ctx context.Context, servingEndpointId string) (*ServingEndpointPermissions, error) {
	return a.servingEndpointsImpl.GetPermissions(ctx, GetServingEndpointPermissionsRequest{
		ServingEndpointId: servingEndpointId,
	})
}

// Retrieves the service logs associated with the provided served model.
func (a *ServingEndpointsAPI) LogsByNameAndServedModelName(ctx context.Context, name string, servedModelName string) (*ServerLogsResponse, error) {
	return a.servingEndpointsImpl.Logs(ctx, LogsRequest{
		Name:            name,
		ServedModelName: servedModelName,
	})
}

// Updates any combination of the serving endpoint's served entities, the
// compute configuration of those served entities, and the endpoint's traffic
// config. An endpoint that already has an update in progress can not be updated
// until the current update completes or fails.
func (a *ServingEndpointsAPI) UpdateConfig(ctx context.Context, endpointCoreConfigInput EndpointCoreConfigInput) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error) {
	servingEndpointDetailed, err := a.servingEndpointsImpl.UpdateConfig(ctx, endpointCoreConfigInput)
	if err != nil {
		return nil, err
	}
	return &WaitGetServingEndpointNotUpdating[ServingEndpointDetailed]{
		Response: servingEndpointDetailed,
		Name:     servingEndpointDetailed.Name,
		Poll: func(timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
			return a.WaitGetServingEndpointNotUpdating(ctx, servingEndpointDetailed.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ServingEndpointsAPI.UpdateConfig] and waits to reach NOT_UPDATING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
//
// Deprecated: use [ServingEndpointsAPI.UpdateConfig].Get() or [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating]
func (a *ServingEndpointsAPI) UpdateConfigAndWait(ctx context.Context, endpointCoreConfigInput EndpointCoreConfigInput, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error) {
	wait, err := a.UpdateConfig(ctx, endpointCoreConfigInput)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ServingEndpointDetailed]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ServingEndpointDetailed) {
		for _, o := range options {
			o(&retries.Info[ServingEndpointDetailed]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Updates any combination of the pt endpoint's served entities, the compute
// configuration of those served entities, and the endpoint's traffic config.
// Updates are instantaneous and endpoint should be updated instantly
func (a *ServingEndpointsAPI) UpdateProvisionedThroughputEndpointConfig(ctx context.Context, updateProvisionedThroughputEndpointConfigRequest UpdateProvisionedThroughputEndpointConfigRequest) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error) {
	servingEndpointDetailed, err := a.servingEndpointsImpl.UpdateProvisionedThroughputEndpointConfig(ctx, updateProvisionedThroughputEndpointConfigRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetServingEndpointNotUpdating[ServingEndpointDetailed]{
		Response: servingEndpointDetailed,
		Name:     servingEndpointDetailed.Name,
		Poll: func(timeout time.Duration, callback func(*ServingEndpointDetailed)) (*ServingEndpointDetailed, error) {
			return a.WaitGetServingEndpointNotUpdating(ctx, servingEndpointDetailed.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ServingEndpointsAPI.UpdateProvisionedThroughputEndpointConfig] and waits to reach NOT_UPDATING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ServingEndpointDetailed](60*time.Minute) functional option.
//
// Deprecated: use [ServingEndpointsAPI.UpdateProvisionedThroughputEndpointConfig].Get() or [ServingEndpointsAPI.WaitGetServingEndpointNotUpdating]
func (a *ServingEndpointsAPI) UpdateProvisionedThroughputEndpointConfigAndWait(ctx context.Context, updateProvisionedThroughputEndpointConfigRequest UpdateProvisionedThroughputEndpointConfigRequest, options ...retries.Option[ServingEndpointDetailed]) (*ServingEndpointDetailed, error) {
	wait, err := a.UpdateProvisionedThroughputEndpointConfig(ctx, updateProvisionedThroughputEndpointConfigRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ServingEndpointDetailed]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ServingEndpointDetailed) {
		for _, o := range options {
			o(&retries.Info[ServingEndpointDetailed]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

type ServingEndpointsDataPlaneInterface interface {

	// Query a serving endpoint
	Query(ctx context.Context, request QueryEndpointInput) (*QueryEndpointResponse, error)
}

func NewServingEndpointsDataPlane(client *client.DatabricksClient,
	controlPlane *ServingEndpointsAPI,
) *ServingEndpointsDataPlaneAPI {
	return &ServingEndpointsDataPlaneAPI{
		servingEndpointsDataPlaneImpl: servingEndpointsDataPlaneImpl{
			client:       client,
			controlPlane: controlPlane,
			dpts: dataplane.NewEndpointTokenSource(
				client,
				client.Config.GetTokenSource(),
			),
		},
	}
}

// Serving endpoints DataPlane provides a set of operations to interact with
// data plane endpoints for Serving endpoints service.
type ServingEndpointsDataPlaneAPI struct {
	servingEndpointsDataPlaneImpl
}
