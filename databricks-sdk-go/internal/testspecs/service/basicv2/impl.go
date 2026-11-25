// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package basicv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just BasicV2 API methods
type basicV2Impl struct {
	client *client.DatabricksClient
}

func (a *basicV2Impl) CreateBasic(ctx context.Context, request CreateBasicRequest) (*Basic, error) {
	var basic Basic
	path := "/api/2.0/basic"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Basic, &basic)
	return &basic, err
}

func (a *basicV2Impl) DeleteBasic(ctx context.Context, request DeleteBasicRequest) error {
	path := fmt.Sprintf("/api/2.0/basic/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *basicV2Impl) GetBasic(ctx context.Context, request GetBasicRequest) (*Basic, error) {
	var basic Basic
	path := fmt.Sprintf("/api/2.0/basic/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &basic)
	return &basic, err
}

func (a *basicV2Impl) ListBasics(ctx context.Context, request ListBasicsRequest) listing.Iterator[Basic] {

	getNextPage := func(ctx context.Context, req ListBasicsRequest) (*ListBasicsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListBasics(ctx, req)
	}
	getItems := func(resp *ListBasicsResponse) []Basic {
		return resp.Basics
	}
	getNextReq := func(resp *ListBasicsResponse) *ListBasicsRequest {
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

func (a *basicV2Impl) ListBasicsAll(ctx context.Context, request ListBasicsRequest) ([]Basic, error) {
	iterator := a.ListBasics(ctx, request)
	return listing.ToSlice[Basic](ctx, iterator)
}

func (a *basicV2Impl) internalListBasics(ctx context.Context, request ListBasicsRequest) (*ListBasicsResponse, error) {
	var listBasicsResponse ListBasicsResponse
	path := "/api/2.0/basics"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBasicsResponse)
	return &listBasicsResponse, err
}

func (a *basicV2Impl) UpdateBasic(ctx context.Context, request UpdateBasicRequest) (*Basic, error) {
	var basic Basic
	path := fmt.Sprintf("/api/2.0/basic/%v", request.Name)
	queryParams := make(map[string]any)
	if request.UpdateMask != "" || slices.Contains(request.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Basic, &basic)
	return &basic, err
}
