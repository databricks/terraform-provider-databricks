// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package marketplace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just ConsumerFulfillments API methods
type consumerFulfillmentsImpl struct {
	client *client.DatabricksClient
}

// Get a high level preview of the metadata of listing installable content.
func (a *consumerFulfillmentsImpl) Get(ctx context.Context, request GetListingContentMetadataRequest) listing.Iterator[SharedDataObject] {

	getNextPage := func(ctx context.Context, req GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGet(ctx, req)
	}
	getItems := func(resp *GetListingContentMetadataResponse) []SharedDataObject {
		return resp.SharedDataObjects
	}
	getNextReq := func(resp *GetListingContentMetadataResponse) *GetListingContentMetadataRequest {
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

// Get a high level preview of the metadata of listing installable content.
func (a *consumerFulfillmentsImpl) GetAll(ctx context.Context, request GetListingContentMetadataRequest) ([]SharedDataObject, error) {
	iterator := a.Get(ctx, request)
	return listing.ToSlice[SharedDataObject](ctx, iterator)
}

func (a *consumerFulfillmentsImpl) internalGet(ctx context.Context, request GetListingContentMetadataRequest) (*GetListingContentMetadataResponse, error) {
	var getListingContentMetadataResponse GetListingContentMetadataResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/content", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingContentMetadataResponse)
	return &getListingContentMetadataResponse, err
}

// Get all listings fulfillments associated with a listing. A _fulfillment_ is a
// potential installation. Standard installations contain metadata about the
// attached share or git repo. Only one of these fields will be present.
// Personalized installations contain metadata about the attached share or git
// repo, as well as the Delta Sharing recipient type.
func (a *consumerFulfillmentsImpl) List(ctx context.Context, request ListFulfillmentsRequest) listing.Iterator[ListingFulfillment] {

	getNextPage := func(ctx context.Context, req ListFulfillmentsRequest) (*ListFulfillmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFulfillmentsResponse) []ListingFulfillment {
		return resp.Fulfillments
	}
	getNextReq := func(resp *ListFulfillmentsResponse) *ListFulfillmentsRequest {
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

// Get all listings fulfillments associated with a listing. A _fulfillment_ is a
// potential installation. Standard installations contain metadata about the
// attached share or git repo. Only one of these fields will be present.
// Personalized installations contain metadata about the attached share or git
// repo, as well as the Delta Sharing recipient type.
func (a *consumerFulfillmentsImpl) ListAll(ctx context.Context, request ListFulfillmentsRequest) ([]ListingFulfillment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListingFulfillment](ctx, iterator)
}

func (a *consumerFulfillmentsImpl) internalList(ctx context.Context, request ListFulfillmentsRequest) (*ListFulfillmentsResponse, error) {
	var listFulfillmentsResponse ListFulfillmentsResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/fulfillments", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFulfillmentsResponse)
	return &listFulfillmentsResponse, err
}

// unexported type that holds implementations of just ConsumerInstallations API methods
type consumerInstallationsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerInstallationsImpl) Create(ctx context.Context, request CreateInstallationRequest) (*Installation, error) {
	var installation Installation
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &installation)
	return &installation, err
}

func (a *consumerInstallationsImpl) Delete(ctx context.Context, request DeleteInstallationRequest) error {
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
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

// List all installations across all listings.
func (a *consumerInstallationsImpl) List(ctx context.Context, request ListAllInstallationsRequest) listing.Iterator[InstallationDetail] {

	getNextPage := func(ctx context.Context, req ListAllInstallationsRequest) (*ListAllInstallationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAllInstallationsResponse) []InstallationDetail {
		return resp.Installations
	}
	getNextReq := func(resp *ListAllInstallationsResponse) *ListAllInstallationsRequest {
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

// List all installations across all listings.
func (a *consumerInstallationsImpl) ListAll(ctx context.Context, request ListAllInstallationsRequest) ([]InstallationDetail, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[InstallationDetail](ctx, iterator)
}

func (a *consumerInstallationsImpl) internalList(ctx context.Context, request ListAllInstallationsRequest) (*ListAllInstallationsResponse, error) {
	var listAllInstallationsResponse ListAllInstallationsResponse
	path := "/api/2.1/marketplace-consumer/installations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllInstallationsResponse)
	return &listAllInstallationsResponse, err
}

// List all installations for a particular listing.
func (a *consumerInstallationsImpl) ListListingInstallations(ctx context.Context, request ListInstallationsRequest) listing.Iterator[InstallationDetail] {

	getNextPage := func(ctx context.Context, req ListInstallationsRequest) (*ListInstallationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListListingInstallations(ctx, req)
	}
	getItems := func(resp *ListInstallationsResponse) []InstallationDetail {
		return resp.Installations
	}
	getNextReq := func(resp *ListInstallationsResponse) *ListInstallationsRequest {
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

// List all installations for a particular listing.
func (a *consumerInstallationsImpl) ListListingInstallationsAll(ctx context.Context, request ListInstallationsRequest) ([]InstallationDetail, error) {
	iterator := a.ListListingInstallations(ctx, request)
	return listing.ToSlice[InstallationDetail](ctx, iterator)
}

func (a *consumerInstallationsImpl) internalListListingInstallations(ctx context.Context, request ListInstallationsRequest) (*ListInstallationsResponse, error) {
	var listInstallationsResponse ListInstallationsResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listInstallationsResponse)
	return &listInstallationsResponse, err
}

func (a *consumerInstallationsImpl) Update(ctx context.Context, request UpdateInstallationRequest) (*UpdateInstallationResponse, error) {
	var updateInstallationResponse UpdateInstallationResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/installations/%v", request.ListingId, request.InstallationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateInstallationResponse)
	return &updateInstallationResponse, err
}

// unexported type that holds implementations of just ConsumerListings API methods
type consumerListingsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerListingsImpl) BatchGet(ctx context.Context, request BatchGetListingsRequest) (*BatchGetListingsResponse, error) {
	var batchGetListingsResponse BatchGetListingsResponse
	path := "/api/2.1/marketplace-consumer/listings:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &batchGetListingsResponse)
	return &batchGetListingsResponse, err
}

func (a *consumerListingsImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	var getListingResponse GetListingResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingResponse)
	return &getListingResponse, err
}

// List all published listings in the Databricks Marketplace that the consumer
// has access to.
func (a *consumerListingsImpl) List(ctx context.Context, request ListListingsRequest) listing.Iterator[Listing] {

	getNextPage := func(ctx context.Context, req ListListingsRequest) (*ListListingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListListingsResponse) []Listing {
		return resp.Listings
	}
	getNextReq := func(resp *ListListingsResponse) *ListListingsRequest {
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

// List all published listings in the Databricks Marketplace that the consumer
// has access to.
func (a *consumerListingsImpl) ListAll(ctx context.Context, request ListListingsRequest) ([]Listing, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}

func (a *consumerListingsImpl) internalList(ctx context.Context, request ListListingsRequest) (*ListListingsResponse, error) {
	var listListingsResponse ListListingsResponse
	path := "/api/2.1/marketplace-consumer/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listListingsResponse)
	return &listListingsResponse, err
}

// Search published listings in the Databricks Marketplace that the consumer has
// access to. This query supports a variety of different search parameters and
// performs fuzzy matching.
func (a *consumerListingsImpl) Search(ctx context.Context, request SearchListingsRequest) listing.Iterator[Listing] {

	getNextPage := func(ctx context.Context, req SearchListingsRequest) (*SearchListingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearch(ctx, req)
	}
	getItems := func(resp *SearchListingsResponse) []Listing {
		return resp.Listings
	}
	getNextReq := func(resp *SearchListingsResponse) *SearchListingsRequest {
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

// Search published listings in the Databricks Marketplace that the consumer has
// access to. This query supports a variety of different search parameters and
// performs fuzzy matching.
func (a *consumerListingsImpl) SearchAll(ctx context.Context, request SearchListingsRequest) ([]Listing, error) {
	iterator := a.Search(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}

func (a *consumerListingsImpl) internalSearch(ctx context.Context, request SearchListingsRequest) (*SearchListingsResponse, error) {
	var searchListingsResponse SearchListingsResponse
	path := "/api/2.1/marketplace-consumer/search-listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &searchListingsResponse)
	return &searchListingsResponse, err
}

// unexported type that holds implementations of just ConsumerPersonalizationRequests API methods
type consumerPersonalizationRequestsImpl struct {
	client *client.DatabricksClient
}

func (a *consumerPersonalizationRequestsImpl) Create(ctx context.Context, request CreatePersonalizationRequest) (*CreatePersonalizationRequestResponse, error) {
	var createPersonalizationRequestResponse CreatePersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createPersonalizationRequestResponse)
	return &createPersonalizationRequestResponse, err
}

func (a *consumerPersonalizationRequestsImpl) Get(ctx context.Context, request GetPersonalizationRequestRequest) (*GetPersonalizationRequestResponse, error) {
	var getPersonalizationRequestResponse GetPersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/listings/%v/personalization-requests", request.ListingId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPersonalizationRequestResponse)
	return &getPersonalizationRequestResponse, err
}

// List personalization requests for a consumer across all listings.
func (a *consumerPersonalizationRequestsImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) listing.Iterator[PersonalizationRequest] {

	getNextPage := func(ctx context.Context, req ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAllPersonalizationRequestsResponse) []PersonalizationRequest {
		return resp.PersonalizationRequests
	}
	getNextReq := func(resp *ListAllPersonalizationRequestsResponse) *ListAllPersonalizationRequestsRequest {
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

// List personalization requests for a consumer across all listings.
func (a *consumerPersonalizationRequestsImpl) ListAll(ctx context.Context, request ListAllPersonalizationRequestsRequest) ([]PersonalizationRequest, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PersonalizationRequest](ctx, iterator)
}

func (a *consumerPersonalizationRequestsImpl) internalList(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	var listAllPersonalizationRequestsResponse ListAllPersonalizationRequestsResponse
	path := "/api/2.1/marketplace-consumer/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllPersonalizationRequestsResponse)
	return &listAllPersonalizationRequestsResponse, err
}

// unexported type that holds implementations of just ConsumerProviders API methods
type consumerProvidersImpl struct {
	client *client.DatabricksClient
}

func (a *consumerProvidersImpl) BatchGet(ctx context.Context, request BatchGetProvidersRequest) (*BatchGetProvidersResponse, error) {
	var batchGetProvidersResponse BatchGetProvidersResponse
	path := "/api/2.1/marketplace-consumer/providers:batchGet"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &batchGetProvidersResponse)
	return &batchGetProvidersResponse, err
}

func (a *consumerProvidersImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1/marketplace-consumer/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getProviderResponse)
	return &getProviderResponse, err
}

// List all providers in the Databricks Marketplace with at least one visible
// listing.
func (a *consumerProvidersImpl) List(ctx context.Context, request ListConsumerProvidersRequest) listing.Iterator[ProviderInfo] {

	getNextPage := func(ctx context.Context, req ListConsumerProvidersRequest) (*ListProvidersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListProvidersResponse) []ProviderInfo {
		return resp.Providers
	}
	getNextReq := func(resp *ListProvidersResponse) *ListConsumerProvidersRequest {
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

// List all providers in the Databricks Marketplace with at least one visible
// listing.
func (a *consumerProvidersImpl) ListAll(ctx context.Context, request ListConsumerProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}

func (a *consumerProvidersImpl) internalList(ctx context.Context, request ListConsumerProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/marketplace-consumer/providers"
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

// unexported type that holds implementations of just ProviderExchangeFilters API methods
type providerExchangeFiltersImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangeFiltersImpl) Create(ctx context.Context, request CreateExchangeFilterRequest) (*CreateExchangeFilterResponse, error) {
	var createExchangeFilterResponse CreateExchangeFilterResponse
	path := "/api/2.0/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExchangeFilterResponse)
	return &createExchangeFilterResponse, err
}

func (a *providerExchangeFiltersImpl) Delete(ctx context.Context, request DeleteExchangeFilterRequest) error {
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/filters/%v", request.Id)
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

// List exchange filter
func (a *providerExchangeFiltersImpl) List(ctx context.Context, request ListExchangeFiltersRequest) listing.Iterator[ExchangeFilter] {

	getNextPage := func(ctx context.Context, req ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListExchangeFiltersResponse) []ExchangeFilter {
		return resp.Filters
	}
	getNextReq := func(resp *ListExchangeFiltersResponse) *ListExchangeFiltersRequest {
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

// List exchange filter
func (a *providerExchangeFiltersImpl) ListAll(ctx context.Context, request ListExchangeFiltersRequest) ([]ExchangeFilter, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ExchangeFilter](ctx, iterator)
}

func (a *providerExchangeFiltersImpl) internalList(ctx context.Context, request ListExchangeFiltersRequest) (*ListExchangeFiltersResponse, error) {
	var listExchangeFiltersResponse ListExchangeFiltersResponse
	path := "/api/2.0/marketplace-exchange/filters"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangeFiltersResponse)
	return &listExchangeFiltersResponse, err
}

func (a *providerExchangeFiltersImpl) Update(ctx context.Context, request UpdateExchangeFilterRequest) (*UpdateExchangeFilterResponse, error) {
	var updateExchangeFilterResponse UpdateExchangeFilterResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/filters/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateExchangeFilterResponse)
	return &updateExchangeFilterResponse, err
}

// unexported type that holds implementations of just ProviderExchanges API methods
type providerExchangesImpl struct {
	client *client.DatabricksClient
}

func (a *providerExchangesImpl) AddListingToExchange(ctx context.Context, request AddExchangeForListingRequest) (*AddExchangeForListingResponse, error) {
	var addExchangeForListingResponse AddExchangeForListingResponse
	path := "/api/2.0/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &addExchangeForListingResponse)
	return &addExchangeForListingResponse, err
}

func (a *providerExchangesImpl) Create(ctx context.Context, request CreateExchangeRequest) (*CreateExchangeResponse, error) {
	var createExchangeResponse CreateExchangeResponse
	path := "/api/2.0/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createExchangeResponse)
	return &createExchangeResponse, err
}

func (a *providerExchangesImpl) Delete(ctx context.Context, request DeleteExchangeRequest) error {
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
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

func (a *providerExchangesImpl) DeleteListingFromExchange(ctx context.Context, request RemoveExchangeForListingRequest) error {
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges-for-listing/%v", request.Id)
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

func (a *providerExchangesImpl) Get(ctx context.Context, request GetExchangeRequest) (*GetExchangeResponse, error) {
	var getExchangeResponse GetExchangeResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getExchangeResponse)
	return &getExchangeResponse, err
}

// List exchanges visible to provider
func (a *providerExchangesImpl) List(ctx context.Context, request ListExchangesRequest) listing.Iterator[Exchange] {

	getNextPage := func(ctx context.Context, req ListExchangesRequest) (*ListExchangesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListExchangesResponse) []Exchange {
		return resp.Exchanges
	}
	getNextReq := func(resp *ListExchangesResponse) *ListExchangesRequest {
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

// List exchanges visible to provider
func (a *providerExchangesImpl) ListAll(ctx context.Context, request ListExchangesRequest) ([]Exchange, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Exchange](ctx, iterator)
}

func (a *providerExchangesImpl) internalList(ctx context.Context, request ListExchangesRequest) (*ListExchangesResponse, error) {
	var listExchangesResponse ListExchangesResponse
	path := "/api/2.0/marketplace-exchange/exchanges"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangesResponse)
	return &listExchangesResponse, err
}

// List exchanges associated with a listing
func (a *providerExchangesImpl) ListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) listing.Iterator[ExchangeListing] {

	getNextPage := func(ctx context.Context, req ListExchangesForListingRequest) (*ListExchangesForListingResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExchangesForListing(ctx, req)
	}
	getItems := func(resp *ListExchangesForListingResponse) []ExchangeListing {
		return resp.ExchangeListing
	}
	getNextReq := func(resp *ListExchangesForListingResponse) *ListExchangesForListingRequest {
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

// List exchanges associated with a listing
func (a *providerExchangesImpl) ListExchangesForListingAll(ctx context.Context, request ListExchangesForListingRequest) ([]ExchangeListing, error) {
	iterator := a.ListExchangesForListing(ctx, request)
	return listing.ToSlice[ExchangeListing](ctx, iterator)
}

func (a *providerExchangesImpl) internalListExchangesForListing(ctx context.Context, request ListExchangesForListingRequest) (*ListExchangesForListingResponse, error) {
	var listExchangesForListingResponse ListExchangesForListingResponse
	path := "/api/2.0/marketplace-exchange/exchanges-for-listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExchangesForListingResponse)
	return &listExchangesForListingResponse, err
}

// List listings associated with an exchange
func (a *providerExchangesImpl) ListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) listing.Iterator[ExchangeListing] {

	getNextPage := func(ctx context.Context, req ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListListingsForExchange(ctx, req)
	}
	getItems := func(resp *ListListingsForExchangeResponse) []ExchangeListing {
		return resp.ExchangeListings
	}
	getNextReq := func(resp *ListListingsForExchangeResponse) *ListListingsForExchangeRequest {
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

// List listings associated with an exchange
func (a *providerExchangesImpl) ListListingsForExchangeAll(ctx context.Context, request ListListingsForExchangeRequest) ([]ExchangeListing, error) {
	iterator := a.ListListingsForExchange(ctx, request)
	return listing.ToSlice[ExchangeListing](ctx, iterator)
}

func (a *providerExchangesImpl) internalListListingsForExchange(ctx context.Context, request ListListingsForExchangeRequest) (*ListListingsForExchangeResponse, error) {
	var listListingsForExchangeResponse ListListingsForExchangeResponse
	path := "/api/2.0/marketplace-exchange/listings-for-exchange"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listListingsForExchangeResponse)
	return &listListingsForExchangeResponse, err
}

func (a *providerExchangesImpl) Update(ctx context.Context, request UpdateExchangeRequest) (*UpdateExchangeResponse, error) {
	var updateExchangeResponse UpdateExchangeResponse
	path := fmt.Sprintf("/api/2.0/marketplace-exchange/exchanges/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateExchangeResponse)
	return &updateExchangeResponse, err
}

// unexported type that holds implementations of just ProviderFiles API methods
type providerFilesImpl struct {
	client *client.DatabricksClient
}

func (a *providerFilesImpl) Create(ctx context.Context, request CreateFileRequest) (*CreateFileResponse, error) {
	var createFileResponse CreateFileResponse
	path := "/api/2.0/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createFileResponse)
	return &createFileResponse, err
}

func (a *providerFilesImpl) Delete(ctx context.Context, request DeleteFileRequest) error {
	path := fmt.Sprintf("/api/2.0/marketplace-provider/files/%v", request.FileId)
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

func (a *providerFilesImpl) Get(ctx context.Context, request GetFileRequest) (*GetFileResponse, error) {
	var getFileResponse GetFileResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/files/%v", request.FileId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getFileResponse)
	return &getFileResponse, err
}

// Lists the files attached to a provider profile or a listing.
func (a *providerFilesImpl) List(ctx context.Context, request ListFilesRequest) listing.Iterator[FileInfo] {

	getNextPage := func(ctx context.Context, req ListFilesRequest) (*ListFilesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFilesResponse) []FileInfo {
		return resp.FileInfos
	}
	getNextReq := func(resp *ListFilesResponse) *ListFilesRequest {
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

// Lists the files attached to a provider profile or a listing.
func (a *providerFilesImpl) ListAll(ctx context.Context, request ListFilesRequest) ([]FileInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FileInfo](ctx, iterator)
}

func (a *providerFilesImpl) internalList(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error) {
	var listFilesResponse ListFilesResponse
	path := "/api/2.0/marketplace-provider/files"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFilesResponse)
	return &listFilesResponse, err
}

// unexported type that holds implementations of just ProviderListings API methods
type providerListingsImpl struct {
	client *client.DatabricksClient
}

func (a *providerListingsImpl) Create(ctx context.Context, request CreateListingRequest) (*CreateListingResponse, error) {
	var createListingResponse CreateListingResponse
	path := "/api/2.0/marketplace-provider/listing"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createListingResponse)
	return &createListingResponse, err
}

func (a *providerListingsImpl) Delete(ctx context.Context, request DeleteListingRequest) error {
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
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

func (a *providerListingsImpl) Get(ctx context.Context, request GetListingRequest) (*GetListingResponse, error) {
	var getListingResponse GetListingResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingResponse)
	return &getListingResponse, err
}

// Lists the listings owned by the provider profiles in the caller's account.
func (a *providerListingsImpl) List(ctx context.Context, request GetListingsRequest) listing.Iterator[Listing] {

	getNextPage := func(ctx context.Context, req GetListingsRequest) (*GetListingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetListingsResponse) []Listing {
		return resp.Listings
	}
	getNextReq := func(resp *GetListingsResponse) *GetListingsRequest {
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

// Lists the listings owned by the provider profiles in the caller's account.
func (a *providerListingsImpl) ListAll(ctx context.Context, request GetListingsRequest) ([]Listing, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Listing](ctx, iterator)
}

func (a *providerListingsImpl) internalList(ctx context.Context, request GetListingsRequest) (*GetListingsResponse, error) {
	var getListingsResponse GetListingsResponse
	path := "/api/2.0/marketplace-provider/listings"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getListingsResponse)
	return &getListingsResponse, err
}

func (a *providerListingsImpl) Update(ctx context.Context, request UpdateListingRequest) (*UpdateListingResponse, error) {
	var updateListingResponse UpdateListingResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateListingResponse)
	return &updateListingResponse, err
}

// unexported type that holds implementations of just ProviderPersonalizationRequests API methods
type providerPersonalizationRequestsImpl struct {
	client *client.DatabricksClient
}

// List personalization requests to this provider. This will return all
// personalization requests, regardless of which listing they are for.
func (a *providerPersonalizationRequestsImpl) List(ctx context.Context, request ListAllPersonalizationRequestsRequest) listing.Iterator[PersonalizationRequest] {

	getNextPage := func(ctx context.Context, req ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAllPersonalizationRequestsResponse) []PersonalizationRequest {
		return resp.PersonalizationRequests
	}
	getNextReq := func(resp *ListAllPersonalizationRequestsResponse) *ListAllPersonalizationRequestsRequest {
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

// List personalization requests to this provider. This will return all
// personalization requests, regardless of which listing they are for.
func (a *providerPersonalizationRequestsImpl) ListAll(ctx context.Context, request ListAllPersonalizationRequestsRequest) ([]PersonalizationRequest, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PersonalizationRequest](ctx, iterator)
}

func (a *providerPersonalizationRequestsImpl) internalList(ctx context.Context, request ListAllPersonalizationRequestsRequest) (*ListAllPersonalizationRequestsResponse, error) {
	var listAllPersonalizationRequestsResponse ListAllPersonalizationRequestsResponse
	path := "/api/2.0/marketplace-provider/personalization-requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAllPersonalizationRequestsResponse)
	return &listAllPersonalizationRequestsResponse, err
}

func (a *providerPersonalizationRequestsImpl) Update(ctx context.Context, request UpdatePersonalizationRequestRequest) (*UpdatePersonalizationRequestResponse, error) {
	var updatePersonalizationRequestResponse UpdatePersonalizationRequestResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/listings/%v/personalization-requests/%v/request-status", request.ListingId, request.RequestId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updatePersonalizationRequestResponse)
	return &updatePersonalizationRequestResponse, err
}

// unexported type that holds implementations of just ProviderProviderAnalyticsDashboards API methods
type providerProviderAnalyticsDashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *providerProviderAnalyticsDashboardsImpl) Create(ctx context.Context) (*ProviderAnalyticsDashboard, error) {
	var providerAnalyticsDashboard ProviderAnalyticsDashboard
	path := "/api/2.0/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, nil, &providerAnalyticsDashboard)
	return &providerAnalyticsDashboard, err
}

func (a *providerProviderAnalyticsDashboardsImpl) Get(ctx context.Context) (*ListProviderAnalyticsDashboardResponse, error) {
	var listProviderAnalyticsDashboardResponse ListProviderAnalyticsDashboardResponse
	path := "/api/2.0/marketplace-provider/analytics_dashboard"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listProviderAnalyticsDashboardResponse)
	return &listProviderAnalyticsDashboardResponse, err
}

func (a *providerProviderAnalyticsDashboardsImpl) GetLatestVersion(ctx context.Context) (*GetLatestVersionProviderAnalyticsDashboardResponse, error) {
	var getLatestVersionProviderAnalyticsDashboardResponse GetLatestVersionProviderAnalyticsDashboardResponse
	path := "/api/2.0/marketplace-provider/analytics_dashboard/latest"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getLatestVersionProviderAnalyticsDashboardResponse)
	return &getLatestVersionProviderAnalyticsDashboardResponse, err
}

func (a *providerProviderAnalyticsDashboardsImpl) Update(ctx context.Context, request UpdateProviderAnalyticsDashboardRequest) (*UpdateProviderAnalyticsDashboardResponse, error) {
	var updateProviderAnalyticsDashboardResponse UpdateProviderAnalyticsDashboardResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/analytics_dashboard/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateProviderAnalyticsDashboardResponse)
	return &updateProviderAnalyticsDashboardResponse, err
}

// unexported type that holds implementations of just ProviderProviders API methods
type providerProvidersImpl struct {
	client *client.DatabricksClient
}

func (a *providerProvidersImpl) Create(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.0/marketplace-provider/provider"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createProviderResponse)
	return &createProviderResponse, err
}

func (a *providerProvidersImpl) Delete(ctx context.Context, request DeleteProviderRequest) error {
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
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

func (a *providerProvidersImpl) Get(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getProviderResponse)
	return &getProviderResponse, err
}

// Lists the provider profiles in the caller's account.
func (a *providerProvidersImpl) List(ctx context.Context, request ListProvidersRequest) listing.Iterator[ProviderInfo] {

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

// Lists the provider profiles in the caller's account.
func (a *providerProvidersImpl) ListAll(ctx context.Context, request ListProvidersRequest) ([]ProviderInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ProviderInfo](ctx, iterator)
}

func (a *providerProvidersImpl) internalList(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.0/marketplace-provider/providers"
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

func (a *providerProvidersImpl) Update(ctx context.Context, request UpdateProviderRequest) (*UpdateProviderResponse, error) {
	var updateProviderResponse UpdateProviderResponse
	path := fmt.Sprintf("/api/2.0/marketplace-provider/providers/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateProviderResponse)
	return &updateProviderResponse, err
}
