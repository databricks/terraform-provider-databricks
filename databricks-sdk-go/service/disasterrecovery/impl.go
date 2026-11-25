// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disasterrecovery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just DisasterRecovery API methods
type disasterRecoveryImpl struct {
	client *client.DatabricksClient
}

func (a *disasterRecoveryImpl) CreateFailoverGroup(ctx context.Context, request CreateFailoverGroupRequest) (*FailoverGroup, error) {
	var failoverGroup FailoverGroup
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v/failover-groups", request.Parent)
	queryParams := make(map[string]any)

	if request.FailoverGroupId != "" || slices.Contains(request.ForceSendFields, "FailoverGroupId") {
		queryParams["failover_group_id"] = request.FailoverGroupId
	}

	if request.ValidateOnly != false || slices.Contains(request.ForceSendFields, "ValidateOnly") {
		queryParams["validate_only"] = request.ValidateOnly
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.FailoverGroup, &failoverGroup)
	return &failoverGroup, err
}

func (a *disasterRecoveryImpl) CreateStableUrl(ctx context.Context, request CreateStableUrlRequest) (*StableUrl, error) {
	var stableUrl StableUrl
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v/stable-urls", request.Parent)
	queryParams := make(map[string]any)

	if request.StableUrlId != "" || slices.Contains(request.ForceSendFields, "StableUrlId") {
		queryParams["stable_url_id"] = request.StableUrlId
	}

	if request.ValidateOnly != false || slices.Contains(request.ForceSendFields, "ValidateOnly") {
		queryParams["validate_only"] = request.ValidateOnly
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.StableUrl, &stableUrl)
	return &stableUrl, err
}

func (a *disasterRecoveryImpl) DeleteFailoverGroup(ctx context.Context, request DeleteFailoverGroupRequest) error {
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *disasterRecoveryImpl) DeleteStableUrl(ctx context.Context, request DeleteStableUrlRequest) error {
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *disasterRecoveryImpl) FailoverFailoverGroup(ctx context.Context, request FailoverFailoverGroupRequest) (*FailoverGroup, error) {
	var failoverGroup FailoverGroup
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v/failover", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &failoverGroup)
	return &failoverGroup, err
}

func (a *disasterRecoveryImpl) GetFailoverGroup(ctx context.Context, request GetFailoverGroupRequest) (*FailoverGroup, error) {
	var failoverGroup FailoverGroup
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &failoverGroup)
	return &failoverGroup, err
}

func (a *disasterRecoveryImpl) GetStableUrl(ctx context.Context, request GetStableUrlRequest) (*StableUrl, error) {
	var stableUrl StableUrl
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &stableUrl)
	return &stableUrl, err
}

// List the failover groups in an account.
//
// List entries are abbreviated: `state`, `replication_point` and
// `unity_catalog_assets.location_mappings` are not populated. Call
// GetFailoverGroup to retrieve the full resource.
func (a *disasterRecoveryImpl) ListFailoverGroups(ctx context.Context, request ListFailoverGroupsRequest) listing.Iterator[FailoverGroup] {

	getNextPage := func(ctx context.Context, req ListFailoverGroupsRequest) (*ListFailoverGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListFailoverGroups(ctx, req)
	}
	getItems := func(resp *ListFailoverGroupsResponse) []FailoverGroup {
		return resp.FailoverGroups
	}
	getNextReq := func(resp *ListFailoverGroupsResponse) *ListFailoverGroupsRequest {
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

// List the failover groups in an account.
//
// List entries are abbreviated: `state`, `replication_point` and
// `unity_catalog_assets.location_mappings` are not populated. Call
// GetFailoverGroup to retrieve the full resource.
func (a *disasterRecoveryImpl) ListFailoverGroupsAll(ctx context.Context, request ListFailoverGroupsRequest) ([]FailoverGroup, error) {
	iterator := a.ListFailoverGroups(ctx, request)
	return listing.ToSlice[FailoverGroup](ctx, iterator)
}

func (a *disasterRecoveryImpl) internalListFailoverGroups(ctx context.Context, request ListFailoverGroupsRequest) (*ListFailoverGroupsResponse, error) {
	var listFailoverGroupsResponse ListFailoverGroupsResponse
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v/failover-groups", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFailoverGroupsResponse)
	return &listFailoverGroupsResponse, err
}

// List stable URLs for an account.
func (a *disasterRecoveryImpl) ListStableUrls(ctx context.Context, request ListStableUrlsRequest) listing.Iterator[StableUrl] {

	getNextPage := func(ctx context.Context, req ListStableUrlsRequest) (*ListStableUrlsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListStableUrls(ctx, req)
	}
	getItems := func(resp *ListStableUrlsResponse) []StableUrl {
		return resp.StableUrls
	}
	getNextReq := func(resp *ListStableUrlsResponse) *ListStableUrlsRequest {
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

// List stable URLs for an account.
func (a *disasterRecoveryImpl) ListStableUrlsAll(ctx context.Context, request ListStableUrlsRequest) ([]StableUrl, error) {
	iterator := a.ListStableUrls(ctx, request)
	return listing.ToSlice[StableUrl](ctx, iterator)
}

func (a *disasterRecoveryImpl) internalListStableUrls(ctx context.Context, request ListStableUrlsRequest) (*ListStableUrlsResponse, error) {
	var listStableUrlsResponse ListStableUrlsResponse
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v/stable-urls", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listStableUrlsResponse)
	return &listStableUrlsResponse, err
}

func (a *disasterRecoveryImpl) UpdateFailoverGroup(ctx context.Context, request UpdateFailoverGroupRequest) (*FailoverGroup, error) {
	var failoverGroup FailoverGroup
	path := fmt.Sprintf("/api/disaster-recovery/v1/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.FailoverGroup, &failoverGroup)
	return &failoverGroup, err
}
