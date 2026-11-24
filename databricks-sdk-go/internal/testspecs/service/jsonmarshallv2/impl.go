// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jsonmarshallv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just JsonMarshallV2 API methods
type jsonMarshallV2Impl struct {
	client *client.DatabricksClient
}

func (a *jsonMarshallV2Impl) GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error) {
	var resource Resource
	path := fmt.Sprintf("/api/2.0/json-marshall/%v", request.Name)
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
