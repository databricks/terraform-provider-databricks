// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Providers API methods
type providersImpl struct {
	client *client.DatabricksClient
}

func (a *providersImpl) Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error) {
	var providerInfo ProviderInfo
	path := "/api/2.1/unity-catalog/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &providerInfo)
	return &providerInfo, err
}

func (a *providersImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *providersImpl) Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error) {
	var providerInfo ProviderInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &providerInfo)
	return &providerInfo, err
}

// Gets an array of available authentication providers. The caller must either
// be a metastore admin, have the **USE_PROVIDER** privilege on the providers,
// or be the owner of the providers. Providers not owned by the caller and for
// which the caller does not have the **USE_PROVIDER** privilege are not
// included in the response. There is no guarantee of a specific ordering of the
// elements in the array.
func (a *providersImpl) List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListProvidersRequest) (*ListProvidersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListProvidersResponse) []ProviderInfo {
		return resp.Providers
	}
	getNextReq := func(resp *ListProvidersResponse) *ListProvidersRequest {
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

// Gets an array of available authentication providers. The caller must either
// be a metastore admin, have the **USE_PROVIDER** privilege on the providers,
// or be the owner of the providers. Providers not owned by the caller and for
// which the caller does not have the **USE_PROVIDER** privilege are not
// included in the response. There is no guarantee of a specific ordering of the
// elements in the array.
func (a *providersImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}

func (a *providersImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *providersImpl) ListProviderShareAssets(ctx context.Context, request ListProviderShareAssetsRequest) (*ListProviderShareAssetsResponse, error) {
	var listProviderShareAssetsResponse ListProviderShareAssetsResponse
	path := fmt.Sprintf("/api/2.1/data-sharing/providers/%v/shares/%v", request.ProviderName, request.ShareName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProviderShareAssetsResponse)
	return &listProviderShareAssetsResponse, err
}

// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *providersImpl) ListShares(ctx context.Context, request ListSharesRequest) listing.Iterator[ProviderShare] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSharesRequest) (*ListProviderSharesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListShares(ctx, req)
	}
	getItems := func(resp *ListProviderSharesResponse) []ProviderShare {
		return resp.Shares
	}
	getNextReq := func(resp *ListProviderSharesResponse) *ListSharesRequest {
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

// Gets an array of a specified provider's shares within the metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner.
func (a *providersImpl) ListSharesAll(ctx context.Context, request ListSharesRequest) ([]ProviderShare, error) {
	iterator := a.ListShares(ctx, request)
	return listing.ToSlice[ProviderShare](ctx, iterator)
}

func (a *providersImpl) internalListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
	var listProviderSharesResponse ListProviderSharesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProviderSharesResponse)
	return &listProviderSharesResponse, err
}

func (a *providersImpl) Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error) {
	var providerInfo ProviderInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &providerInfo)
	return &providerInfo, err
}

// unexported type that holds implementations of just RecipientActivation API methods
type recipientActivationImpl struct {
	client *client.DatabricksClient
}

func (a *recipientActivationImpl) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, nil)
	return err
}

func (a *recipientActivationImpl) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	var retrieveTokenResponse RetrieveTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation/%v", request.ActivationUrl)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &retrieveTokenResponse)
	return &retrieveTokenResponse, err
}

// unexported type that holds implementations of just RecipientFederationPolicies API methods
type recipientFederationPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *recipientFederationPoliciesImpl) Create(ctx context.Context, request CreateFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies", request.RecipientName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

func (a *recipientFederationPoliciesImpl) Delete(ctx context.Context, request DeleteFederationPolicyRequest) error {
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies/%v", request.RecipientName, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *recipientFederationPoliciesImpl) GetFederationPolicy(ctx context.Context, request GetFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies/%v", request.RecipientName, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &federationPolicy)
	return &federationPolicy, err
}

// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
// from Databricks to non-Databricks recipients. The caller must have read
// access to the recipient.
func (a *recipientFederationPoliciesImpl) List(ctx context.Context, request ListFederationPoliciesRequest) listing.Iterator[FederationPolicy] {

	getNextPage := func(ctx context.Context, req ListFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFederationPoliciesResponse) []FederationPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListFederationPoliciesResponse) *ListFederationPoliciesRequest {
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

// Lists federation policies for an OIDC_FEDERATION recipient for sharing data
// from Databricks to non-Databricks recipients. The caller must have read
// access to the recipient.
func (a *recipientFederationPoliciesImpl) ListAll(ctx context.Context, request ListFederationPoliciesRequest) ([]FederationPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FederationPolicy](ctx, iterator)
}

func (a *recipientFederationPoliciesImpl) internalList(ctx context.Context, request ListFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
	var listFederationPoliciesResponse ListFederationPoliciesResponse
	path := fmt.Sprintf("/api/2.0/data-sharing/recipients/%v/federation-policies", request.RecipientName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFederationPoliciesResponse)
	return &listFederationPoliciesResponse, err
}

// unexported type that holds implementations of just Recipients API methods
type recipientsImpl struct {
	client *client.DatabricksClient
}

func (a *recipientsImpl) Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := "/api/2.1/unity-catalog/recipients"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

func (a *recipientsImpl) Delete(ctx context.Context, request DeleteRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *recipientsImpl) Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

// Gets an array of all share recipients within the current metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *recipientsImpl) List(ctx context.Context, request ListRecipientsRequest) listing.Iterator[RecipientInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListRecipientsRequest) (*ListRecipientsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListRecipientsResponse) []RecipientInfo {
		return resp.Recipients
	}
	getNextReq := func(resp *ListRecipientsResponse) *ListRecipientsRequest {
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

// Gets an array of all share recipients within the current metastore where:
//
// * the caller is a metastore admin, or * the caller is the owner. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *recipientsImpl) ListAll(ctx context.Context, request ListRecipientsRequest) ([]RecipientInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RecipientInfo](ctx, iterator)
}

func (a *recipientsImpl) internalList(ctx context.Context, request ListRecipientsRequest) (*ListRecipientsResponse, error) {
	var listRecipientsResponse ListRecipientsResponse
	path := "/api/2.1/unity-catalog/recipients"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRecipientsResponse)
	return &listRecipientsResponse, err
}

func (a *recipientsImpl) RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

func (a *recipientsImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	var getRecipientSharePermissionsResponse GetRecipientSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getRecipientSharePermissionsResponse)
	return &getRecipientSharePermissionsResponse, err
}

func (a *recipientsImpl) Update(ctx context.Context, request UpdateRecipient) (*RecipientInfo, error) {
	var recipientInfo RecipientInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &recipientInfo)
	return &recipientInfo, err
}

// unexported type that holds implementations of just Shares API methods
type sharesImpl struct {
	client *client.DatabricksClient
}

func (a *sharesImpl) Create(ctx context.Context, request CreateShare) (*ShareInfo, error) {
	var shareInfo ShareInfo
	path := "/api/2.1/unity-catalog/shares"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &shareInfo)
	return &shareInfo, err
}

func (a *sharesImpl) Delete(ctx context.Context, request DeleteShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *sharesImpl) Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error) {
	var shareInfo ShareInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &shareInfo)
	return &shareInfo, err
}

// Gets an array of data object shares from the metastore. If the caller has the
// USE_SHARE privilege on the metastore, all shares are returned. Otherwise,
// only shares owned by the caller are returned. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *sharesImpl) ListShares(ctx context.Context, request SharesListRequest) listing.Iterator[ShareInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req SharesListRequest) (*ListSharesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListShares(ctx, req)
	}
	getItems := func(resp *ListSharesResponse) []ShareInfo {
		return resp.Shares
	}
	getNextReq := func(resp *ListSharesResponse) *SharesListRequest {
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

// Gets an array of data object shares from the metastore. If the caller has the
// USE_SHARE privilege on the metastore, all shares are returned. Otherwise,
// only shares owned by the caller are returned. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *sharesImpl) ListSharesAll(ctx context.Context, request SharesListRequest) ([]ShareInfo, error) {
	iterator := a.ListShares(ctx, request)
	return listing.ToSlice[ShareInfo](ctx, iterator)
}

func (a *sharesImpl) internalListShares(ctx context.Context, request SharesListRequest) (*ListSharesResponse, error) {
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

func (a *sharesImpl) SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetSharePermissionsResponse, error) {
	var getSharePermissionsResponse GetSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getSharePermissionsResponse)
	return &getSharePermissionsResponse, err
}

func (a *sharesImpl) Update(ctx context.Context, request UpdateShare) (*ShareInfo, error) {
	var shareInfo ShareInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &shareInfo)
	return &shareInfo, err
}

func (a *sharesImpl) UpdatePermissions(ctx context.Context, request UpdateSharePermissions) (*UpdateSharePermissionsResponse, error) {
	var updateSharePermissionsResponse UpdateSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateSharePermissionsResponse)
	return &updateSharePermissionsResponse, err
}
