// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/google/uuid"
)

// unexported type that holds implementations of just BillableUsage API methods
type billableUsageImpl struct {
	client *client.DatabricksClient
}

func (a *billableUsageImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	var downloadResponse DownloadResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/usage/download", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &downloadResponse)
	return &downloadResponse, err
}

// unexported type that holds implementations of just BudgetPolicy API methods
type budgetPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *budgetPolicyImpl) Create(ctx context.Context, request CreateBudgetPolicyRequest) (*BudgetPolicy, error) {
	var budgetPolicy BudgetPolicy
	if request.RequestId == "" {
		request.RequestId = uuid.New().String()
	}
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &budgetPolicy)
	return &budgetPolicy, err
}

func (a *budgetPolicyImpl) Delete(ctx context.Context, request DeleteBudgetPolicyRequest) error {
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *budgetPolicyImpl) Get(ctx context.Context, request GetBudgetPolicyRequest) (*BudgetPolicy, error) {
	var budgetPolicy BudgetPolicy
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &budgetPolicy)
	return &budgetPolicy, err
}

// Lists the usage policies that the caller has permission to view. Policies are
// sorted by name in ascending order unless `sort_spec` requests a different
// order.
func (a *budgetPolicyImpl) List(ctx context.Context, request ListBudgetPoliciesRequest) listing.Iterator[BudgetPolicy] {

	getNextPage := func(ctx context.Context, req ListBudgetPoliciesRequest) (*ListBudgetPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListBudgetPoliciesResponse) []BudgetPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListBudgetPoliciesResponse) *ListBudgetPoliciesRequest {
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

// Lists the usage policies that the caller has permission to view. Policies are
// sorted by name in ascending order unless `sort_spec` requests a different
// order.
func (a *budgetPolicyImpl) ListAll(ctx context.Context, request ListBudgetPoliciesRequest) ([]BudgetPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BudgetPolicy](ctx, iterator)
}

func (a *budgetPolicyImpl) internalList(ctx context.Context, request ListBudgetPoliciesRequest) (*ListBudgetPoliciesResponse, error) {
	var listBudgetPoliciesResponse ListBudgetPoliciesResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBudgetPoliciesResponse)
	return &listBudgetPoliciesResponse, err
}

func (a *budgetPolicyImpl) Update(ctx context.Context, request UpdateBudgetPolicyRequest) (*BudgetPolicy, error) {
	var budgetPolicy BudgetPolicy
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)

	if request.LimitConfig != nil {
		queryParams["limit_config"] = request.LimitConfig
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Policy, &budgetPolicy)
	return &budgetPolicy, err
}

// unexported type that holds implementations of just Budgets API methods
type budgetsImpl struct {
	client *client.DatabricksClient
}

func (a *budgetsImpl) Create(ctx context.Context, request CreateBudgetConfigurationRequest) (*CreateBudgetConfigurationResponse, error) {
	var createBudgetConfigurationResponse CreateBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createBudgetConfigurationResponse)
	return &createBudgetConfigurationResponse, err
}

func (a *budgetsImpl) Delete(ctx context.Context, request DeleteBudgetConfigurationRequest) error {
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *budgetsImpl) Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error) {
	var getBudgetConfigurationResponse GetBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getBudgetConfigurationResponse)
	return &getBudgetConfigurationResponse, err
}

func (a *budgetsImpl) GetBudgetSpendStatus(ctx context.Context, request GetBudgetSpendStatusRequest) (*BudgetSpendStatus, error) {
	var budgetSpendStatus BudgetSpendStatus
	path := fmt.Sprintf("/api/2.1/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &budgetSpendStatus)
	return &budgetSpendStatus, err
}

// Gets all budgets associated with this account.
func (a *budgetsImpl) List(ctx context.Context, request ListBudgetConfigurationsRequest) listing.Iterator[BudgetConfiguration] {

	getNextPage := func(ctx context.Context, req ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListBudgetConfigurationsResponse) []BudgetConfiguration {
		return resp.Budgets
	}
	getNextReq := func(resp *ListBudgetConfigurationsResponse) *ListBudgetConfigurationsRequest {
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

// Gets all budgets associated with this account.
func (a *budgetsImpl) ListAll(ctx context.Context, request ListBudgetConfigurationsRequest) ([]BudgetConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BudgetConfiguration](ctx, iterator)
}

func (a *budgetsImpl) internalList(ctx context.Context, request ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error) {
	var listBudgetConfigurationsResponse ListBudgetConfigurationsResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBudgetConfigurationsResponse)
	return &listBudgetConfigurationsResponse, err
}

func (a *budgetsImpl) ListBudgetSpendStatuses(ctx context.Context, request ListBudgetSpendStatusesRequest) (*ListBudgetSpendStatusesResponse, error) {
	var listBudgetSpendStatusesResponse ListBudgetSpendStatusesResponse
	path := fmt.Sprintf("/api/2.1/%v/budget-spend-statuses", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBudgetSpendStatusesResponse)
	return &listBudgetSpendStatusesResponse, err
}

func (a *budgetsImpl) Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error) {
	var updateBudgetConfigurationResponse UpdateBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateBudgetConfigurationResponse)
	return &updateBudgetConfigurationResponse, err
}

// unexported type that holds implementations of just LogDelivery API methods
type logDeliveryImpl struct {
	client *client.DatabricksClient
}

func (a *logDeliveryImpl) Create(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

func (a *logDeliveryImpl) Get(ctx context.Context, request GetLogDeliveryRequest) (*GetLogDeliveryConfigurationResponse, error) {
	var getLogDeliveryConfigurationResponse GetLogDeliveryConfigurationResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), request.LogDeliveryConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getLogDeliveryConfigurationResponse)
	return &getLogDeliveryConfigurationResponse, err
}

// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *logDeliveryImpl) List(ctx context.Context, request ListLogDeliveryRequest) listing.Iterator[LogDeliveryConfiguration] {

	getNextPage := func(ctx context.Context, req ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *WrappedLogDeliveryConfigurations) []LogDeliveryConfiguration {
		return resp.LogDeliveryConfigurations
	}
	getNextReq := func(resp *WrappedLogDeliveryConfigurations) *ListLogDeliveryRequest {
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

// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *logDeliveryImpl) ListAll(ctx context.Context, request ListLogDeliveryRequest) ([]LogDeliveryConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[LogDeliveryConfiguration](ctx, iterator)
}

func (a *logDeliveryImpl) internalList(ctx context.Context, request ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {
	var wrappedLogDeliveryConfigurations WrappedLogDeliveryConfigurations
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &wrappedLogDeliveryConfigurations)
	return &wrappedLogDeliveryConfigurations, err
}

func (a *logDeliveryImpl) PatchStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), request.LogDeliveryConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just UsageDashboards API methods
type usageDashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *usageDashboardsImpl) Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error) {
	var createBillingUsageDashboardResponse CreateBillingUsageDashboardResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createBillingUsageDashboardResponse)
	return &createBillingUsageDashboardResponse, err
}

func (a *usageDashboardsImpl) Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error) {
	var getBillingUsageDashboardResponse GetBillingUsageDashboardResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getBillingUsageDashboardResponse)
	return &getBillingUsageDashboardResponse, err
}
