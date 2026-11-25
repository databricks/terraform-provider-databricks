// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package httpcallv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just HttpCallV2 API methods
type httpCallV2Impl struct {
	client *client.DatabricksClient
}

func (a *httpCallV2Impl) CreateResource(ctx context.Context, request CreateResourceRequest) (*Resource, error) {
	var resource Resource
	path := fmt.Sprintf("/api/2.0/http-call/%v/%v/%v", request.PathParamString, request.PathParamInt, request.PathParamBool)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resource)
	return &resource, err
}

func (a *httpCallV2Impl) GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error) {
	var resource Resource
	path := fmt.Sprintf("/api/2.0/http-call/%v/%v/%v", request.PathParamString, request.PathParamInt, request.PathParamBool)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &resource)
	return &resource, err
}

func (a *httpCallV2Impl) SyncResource(ctx context.Context, request SyncResourceRequest) (*Resource, error) {
	var resource Resource
	path := fmt.Sprintf("/api/2.0/http-call/%v/%v/%v/state:sync", request.PathParamString, request.PathParamInt, request.PathParamBool)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resource)
	return &resource, err
}

func (a *httpCallV2Impl) UpdateResource(ctx context.Context, request UpdateResourceRequest) (*Resource, error) {
	var resource Resource
	path := fmt.Sprintf("/api/2.0/http-call/%v/%v/%v", request.NestedPathParamString, request.NestedPathParamInt, request.NestedPathParamBool)
	queryParams := make(map[string]any)

	if request.FieldMask != nil || slices.Contains(request.ForceSendFields, "FieldMask") {
		fieldMaskJson, fieldMaskMarshallError := json.Marshal(request.FieldMask)
		if fieldMaskMarshallError != nil {
			return nil, fieldMaskMarshallError
		}

		queryParams["field_mask"] = strings.Trim(string(fieldMaskJson), `"`)
	}

	if request.OptionalComplexQueryParam != nil || slices.Contains(request.ForceSendFields, "OptionalComplexQueryParam") {
		queryParams["optional_complex_query_param"] = request.OptionalComplexQueryParam
	}

	if request.QueryParamBool != false || slices.Contains(request.ForceSendFields, "QueryParamBool") {
		queryParams["query_param_bool"] = request.QueryParamBool
	}

	if request.QueryParamInt != 0 || slices.Contains(request.ForceSendFields, "QueryParamInt") {
		queryParams["query_param_int"] = request.QueryParamInt
	}

	if request.QueryParamString != "" || slices.Contains(request.ForceSendFields, "QueryParamString") {
		queryParams["query_param_string"] = request.QueryParamString
	}

	if request.RepeatedComplexQueryParam != nil || slices.Contains(request.ForceSendFields, "RepeatedComplexQueryParam") {
		queryParams["repeated_complex_query_param"] = request.RepeatedComplexQueryParam
	}

	if request.RepeatedQueryParam != nil || slices.Contains(request.ForceSendFields, "RepeatedQueryParam") {
		queryParams["repeated_query_param"] = request.RepeatedQueryParam
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Resource, &resource)
	return &resource, err
}
