// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just AccountFederationPolicy API methods
type accountFederationPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *accountFederationPolicyImpl) Create(ctx context.Context, request CreateAccountFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)

	if request.PolicyId != "" || slices.Contains(request.ForceSendFields, "PolicyId") {
		queryParams["policy_id"] = request.PolicyId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

func (a *accountFederationPolicyImpl) Delete(ctx context.Context, request DeleteAccountFederationPolicyRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountFederationPolicyImpl) Get(ctx context.Context, request GetAccountFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &federationPolicy)
	return &federationPolicy, err
}

// List account federation policies.
func (a *accountFederationPolicyImpl) List(ctx context.Context, request ListAccountFederationPoliciesRequest) listing.Iterator[FederationPolicy] {

	getNextPage := func(ctx context.Context, req ListAccountFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFederationPoliciesResponse) []FederationPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListFederationPoliciesResponse) *ListAccountFederationPoliciesRequest {
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

// List account federation policies.
func (a *accountFederationPolicyImpl) ListAll(ctx context.Context, request ListAccountFederationPoliciesRequest) ([]FederationPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FederationPolicy](ctx, iterator)
}

func (a *accountFederationPolicyImpl) internalList(ctx context.Context, request ListAccountFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
	var listFederationPoliciesResponse ListFederationPoliciesResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFederationPoliciesResponse)
	return &listFederationPoliciesResponse, err
}

func (a *accountFederationPolicyImpl) Update(ctx context.Context, request UpdateAccountFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" || slices.Contains(request.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

// unexported type that holds implementations of just CustomAppIntegration API methods
type customAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *customAppIntegrationImpl) Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error) {
	var createCustomAppIntegrationOutput CreateCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createCustomAppIntegrationOutput)
	return &createCustomAppIntegrationOutput, err
}

func (a *customAppIntegrationImpl) Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *customAppIntegrationImpl) Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error) {
	var getCustomAppIntegrationOutput GetCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getCustomAppIntegrationOutput)
	return &getCustomAppIntegrationOutput, err
}

// Get the list of custom OAuth app integrations for the specified Databricks
// account.
func (a *customAppIntegrationImpl) List(ctx context.Context, request ListCustomAppIntegrationsRequest) listing.Iterator[GetCustomAppIntegrationOutput] {

	getNextPage := func(ctx context.Context, req ListCustomAppIntegrationsRequest) (*GetCustomAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetCustomAppIntegrationsOutput) []GetCustomAppIntegrationOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetCustomAppIntegrationsOutput) *ListCustomAppIntegrationsRequest {
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

// Get the list of custom OAuth app integrations for the specified Databricks
// account.
func (a *customAppIntegrationImpl) ListAll(ctx context.Context, request ListCustomAppIntegrationsRequest) ([]GetCustomAppIntegrationOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[GetCustomAppIntegrationOutput](ctx, iterator)
}

func (a *customAppIntegrationImpl) internalList(ctx context.Context, request ListCustomAppIntegrationsRequest) (*GetCustomAppIntegrationsOutput, error) {
	var getCustomAppIntegrationsOutput GetCustomAppIntegrationsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getCustomAppIntegrationsOutput)
	return &getCustomAppIntegrationsOutput, err
}

func (a *customAppIntegrationImpl) Update(ctx context.Context, request UpdateCustomAppIntegration) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just OAuthPublishedApps API methods
type oAuthPublishedAppsImpl struct {
	client *client.DatabricksClient
}

// Get all the available published OAuth apps in Databricks.
func (a *oAuthPublishedAppsImpl) List(ctx context.Context, request ListOAuthPublishedAppsRequest) listing.Iterator[PublishedAppOutput] {

	getNextPage := func(ctx context.Context, req ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetPublishedAppsOutput) []PublishedAppOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetPublishedAppsOutput) *ListOAuthPublishedAppsRequest {
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

// Get all the available published OAuth apps in Databricks.
func (a *oAuthPublishedAppsImpl) ListAll(ctx context.Context, request ListOAuthPublishedAppsRequest) ([]PublishedAppOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PublishedAppOutput](ctx, iterator)
}

func (a *oAuthPublishedAppsImpl) internalList(ctx context.Context, request ListOAuthPublishedAppsRequest) (*GetPublishedAppsOutput, error) {
	var getPublishedAppsOutput GetPublishedAppsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-apps", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPublishedAppsOutput)
	return &getPublishedAppsOutput, err
}

// unexported type that holds implementations of just PublishedAppIntegration API methods
type publishedAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *publishedAppIntegrationImpl) Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error) {
	var createPublishedAppIntegrationOutput CreatePublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createPublishedAppIntegrationOutput)
	return &createPublishedAppIntegrationOutput, err
}

func (a *publishedAppIntegrationImpl) Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *publishedAppIntegrationImpl) Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error) {
	var getPublishedAppIntegrationOutput GetPublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPublishedAppIntegrationOutput)
	return &getPublishedAppIntegrationOutput, err
}

// Get the list of published OAuth app integrations for the specified Databricks
// account.
func (a *publishedAppIntegrationImpl) List(ctx context.Context, request ListPublishedAppIntegrationsRequest) listing.Iterator[GetPublishedAppIntegrationOutput] {

	getNextPage := func(ctx context.Context, req ListPublishedAppIntegrationsRequest) (*GetPublishedAppIntegrationsOutput, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *GetPublishedAppIntegrationsOutput) []GetPublishedAppIntegrationOutput {
		return resp.Apps
	}
	getNextReq := func(resp *GetPublishedAppIntegrationsOutput) *ListPublishedAppIntegrationsRequest {
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

// Get the list of published OAuth app integrations for the specified Databricks
// account.
func (a *publishedAppIntegrationImpl) ListAll(ctx context.Context, request ListPublishedAppIntegrationsRequest) ([]GetPublishedAppIntegrationOutput, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[GetPublishedAppIntegrationOutput](ctx, iterator)
}

func (a *publishedAppIntegrationImpl) internalList(ctx context.Context, request ListPublishedAppIntegrationsRequest) (*GetPublishedAppIntegrationsOutput, error) {
	var getPublishedAppIntegrationsOutput GetPublishedAppIntegrationsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPublishedAppIntegrationsOutput)
	return &getPublishedAppIntegrationsOutput, err
}

func (a *publishedAppIntegrationImpl) Update(ctx context.Context, request UpdatePublishedAppIntegration) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just ServicePrincipalFederationPolicy API methods
type servicePrincipalFederationPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalFederationPolicyImpl) Create(ctx context.Context, request CreateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	queryParams := make(map[string]any)

	if request.PolicyId != "" || slices.Contains(request.ForceSendFields, "PolicyId") {
		queryParams["policy_id"] = request.PolicyId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

func (a *servicePrincipalFederationPolicyImpl) Delete(ctx context.Context, request DeleteServicePrincipalFederationPolicyRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *servicePrincipalFederationPolicyImpl) Get(ctx context.Context, request GetServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &federationPolicy)
	return &federationPolicy, err
}

// List account federation policies.
func (a *servicePrincipalFederationPolicyImpl) List(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) listing.Iterator[FederationPolicy] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFederationPoliciesResponse) []FederationPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListFederationPoliciesResponse) *ListServicePrincipalFederationPoliciesRequest {
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

// List account federation policies.
func (a *servicePrincipalFederationPolicyImpl) ListAll(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) ([]FederationPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FederationPolicy](ctx, iterator)
}

func (a *servicePrincipalFederationPolicyImpl) internalList(ctx context.Context, request ListServicePrincipalFederationPoliciesRequest) (*ListFederationPoliciesResponse, error) {
	var listFederationPoliciesResponse ListFederationPoliciesResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFederationPoliciesResponse)
	return &listFederationPoliciesResponse, err
}

func (a *servicePrincipalFederationPolicyImpl) Update(ctx context.Context, request UpdateServicePrincipalFederationPolicyRequest) (*FederationPolicy, error) {
	var federationPolicy FederationPolicy
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/federationPolicies/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.PolicyId)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" || slices.Contains(request.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Policy, &federationPolicy)
	return &federationPolicy, err
}

// unexported type that holds implementations of just ServicePrincipalSecrets API methods
type servicePrincipalSecretsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalSecretsImpl) Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error) {
	var createServicePrincipalSecretResponse CreateServicePrincipalSecretResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createServicePrincipalSecretResponse)
	return &createServicePrincipalSecretResponse, err
}

func (a *servicePrincipalSecretsImpl) Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets/%v", a.client.ConfiguredAccountID(), request.ServicePrincipalId, request.SecretId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *servicePrincipalSecretsImpl) List(ctx context.Context, request ListServicePrincipalSecretsRequest) listing.Iterator[SecretInfo] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalSecretsResponse) []SecretInfo {
		return resp.Secrets
	}
	getNextReq := func(resp *ListServicePrincipalSecretsResponse) *ListServicePrincipalSecretsRequest {
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

// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *servicePrincipalSecretsImpl) ListAll(ctx context.Context, request ListServicePrincipalSecretsRequest) ([]SecretInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SecretInfo](ctx, iterator)
}

func (a *servicePrincipalSecretsImpl) internalList(ctx context.Context, request ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
	var listServicePrincipalSecretsResponse ListServicePrincipalSecretsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/servicePrincipals/%v/credentials/secrets", a.client.ConfiguredAccountID(), request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalSecretsResponse)
	return &listServicePrincipalSecretsResponse, err
}

// unexported type that holds implementations of just ServicePrincipalSecretsProxy API methods
type servicePrincipalSecretsProxyImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalSecretsProxyImpl) Create(ctx context.Context, request CreateServicePrincipalSecretRequest) (*CreateServicePrincipalSecretResponse, error) {
	var createServicePrincipalSecretResponse CreateServicePrincipalSecretResponse
	path := fmt.Sprintf("/api/2.0/accounts/servicePrincipals/%v/credentials/secrets", request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createServicePrincipalSecretResponse)
	return &createServicePrincipalSecretResponse, err
}

func (a *servicePrincipalSecretsProxyImpl) Delete(ctx context.Context, request DeleteServicePrincipalSecretRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/servicePrincipals/%v/credentials/secrets/%v", request.ServicePrincipalId, request.SecretId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *servicePrincipalSecretsProxyImpl) List(ctx context.Context, request ListServicePrincipalSecretsRequest) listing.Iterator[SecretInfo] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalSecretsResponse) []SecretInfo {
		return resp.Secrets
	}
	getNextReq := func(resp *ListServicePrincipalSecretsResponse) *ListServicePrincipalSecretsRequest {
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

// List all secrets associated with the given service principal. This operation
// only returns information about the secrets themselves and does not include
// the secret values.
func (a *servicePrincipalSecretsProxyImpl) ListAll(ctx context.Context, request ListServicePrincipalSecretsRequest) ([]SecretInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SecretInfo](ctx, iterator)
}

func (a *servicePrincipalSecretsProxyImpl) internalList(ctx context.Context, request ListServicePrincipalSecretsRequest) (*ListServicePrincipalSecretsResponse, error) {
	var listServicePrincipalSecretsResponse ListServicePrincipalSecretsResponse
	path := fmt.Sprintf("/api/2.0/accounts/servicePrincipals/%v/credentials/secrets", request.ServicePrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalSecretsResponse)
	return &listServicePrincipalSecretsResponse, err
}
