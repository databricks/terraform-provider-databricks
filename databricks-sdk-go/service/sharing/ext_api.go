package sharing

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type sharesAPIUtilities interface {
	// Gets an array of data object shares from the metastore. The caller must be a
	// metastore admin or the owner of the share. There is no guarantee of a
	// specific ordering of the elements in the array.
	//
	// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
	List(ctx context.Context, request ListSharesRequest) listing.Iterator[ShareInfo]

	// Gets an array of data object shares from the metastore. The caller must be a
	// metastore admin or the owner of the share. There is no guarantee of a
	// specific ordering of the elements in the array.
	//
	// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
	ListAll(ctx context.Context, request ListSharesRequest) ([]ShareInfo, error)
}

// Gets an array of data object shares from the metastore. The caller must be a
// metastore admin or the owner of the share. There is no guarantee of a
// specific ordering of the elements in the array.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
func (a *SharesAPI) List(ctx context.Context, request ListSharesRequest) listing.Iterator[ShareInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSharesRequest) (*ListSharesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListSharesResponse) []ShareInfo {
		return resp.Shares
	}
	getNextReq := func(resp *ListSharesResponse) *ListSharesRequest {
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

// Gets an array of data object shares from the metastore. The caller must be a
// metastore admin or the owner of the share. There is no guarantee of a
// specific ordering of the elements in the array.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
func (a *SharesAPI) ListAll(ctx context.Context, request ListSharesRequest) ([]ShareInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ShareInfo](ctx, iterator)
}

func (a *SharesAPI) internalList(ctx context.Context, request ListSharesRequest) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSharesResponse)
	return &listSharesResponse, err
}
