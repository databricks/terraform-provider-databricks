package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// List of management information
const (
	AzureDatabricksResourceID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"
)

// func (c *DatabricksClient) configureAzureAuth() (bool, error) {
// 	if c.AzureAuth.WorkspaceName == "" {
// 		return false, nil
// 	}
// 	c.AzureAuth.databricksClient = c
// 	patTokenDuration, err := strconv.Atoi(c.AzureAuth.PATTokenDurationSeconds)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to parse pat_token_duration_seconds[%v], %w", patTokenDuration, err)
// 	}
// 	c.AzureAuth.patTokenSeconds = int32(patTokenDuration)
// 	//c.customAuthorizer = c.AzureAuth.initWorkspaceAndGetClient

// 	return true, nil
// }

// AzureAuth contains all the auth information for azure sp authentication
type AzureAuth struct {
	ManagedResourceGroup string
	AzureRegion          string

	WorkspaceName  string
	ResourceGroup  string
	SubscriptionID string

	// azurerm_databricks_workspace.this.id ->
	// /subscriptions/{subscription}/resourceGroups/{rg}/providers/Microsoft.Databricks/workspaces/{name}
	ResourceID string // todo: make working

	ClientSecret string
	ClientID     string
	TenantID     string

	// temporary workaround for SP-based auth
	PATTokenDurationSeconds string

	// private property to give resource access
	databricksClient *DatabricksClient
	authorizerMutex  *sync.Mutex

	azureManagementEndpoint string
	authorizer              autorest.Authorizer
	temporaryPat            *model.TokenResponse
}

func (aa *AzureAuth) resourceID() string {
	if aa.ResourceID != "" {
		return aa.ResourceID
	}
	if aa.SubscriptionID == "" || aa.ResourceGroup == "" || aa.WorkspaceName == "" {
		return ""
	}
	return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Databricks/workspaces/%s",
		aa.SubscriptionID, aa.ResourceGroup, aa.WorkspaceName)
}

func (aa *AzureAuth) workspaceInfoURL() string {
	// TODO: i guess i've overengineered it here with unit tests...
	endpoint := "https://management.azure.com"
	if aa.azureManagementEndpoint != "" {
		endpoint = aa.azureManagementEndpoint
	}
	return endpoint + aa.resourceID()
}

func (aa *AzureAuth) isClientSecretSet() bool {
	return aa.ClientID != "" && aa.ClientSecret != "" && aa.TenantID != ""
}

func (aa *AzureAuth) configureWithClientSecret2() (func(r *http.Request) (*http.Request, error), error) {
	if aa.resourceID() == "" {
		return nil, nil
	}
	if !aa.isClientSecretSet() {
		return nil, nil
	}
	managementAuthorizer, err := aa.getClientSecretAuthorizer(
		azure.PublicCloud.ServiceManagementEndpoint)
	if err != nil {
		return nil, err
	}
	platformAuthorizer, err := aa.getClientSecretAuthorizer(
		AzureDatabricksResourceID)
	if err != nil {
		return nil, err
	}
	return func(r *http.Request) (*http.Request, error) {
		pat, err := aa.acquirePAT(managementAuthorizer, platformAuthorizer)
		if err != nil {
			return nil, err
		}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pat.TokenValue))
		return r, nil
	}, nil
}

// acquirePAT is supposed to refresh temporary PAT if it's expired. E.g. if provisioning takes longer than an hour
func (aa *AzureAuth) acquirePAT(managementAuthorizer, platformAuthorizer autorest.Authorizer) (*model.TokenResponse, error) {
	if aa.temporaryPat != nil {
		// todo: add IsExpired
		return aa.temporaryPat, nil
	}
	aa.authorizerMutex.Lock()
	defer aa.authorizerMutex.Unlock()
	tokenResponse, err := aa.stuff(
		managementAuthorizer, platformAuthorizer, func(r *http.Request) (*http.Request, error) {
			log.Printf("[DEBUG] Setting 'X-Databricks-Azure-SP-Management-Token' header")
			bearerAuth, ok := managementAuthorizer.(*autorest.BearerAuthorizer)
			if !ok {
				return nil, fmt.Errorf("Supposed to get BearerAuthorizer, but got %v", managementAuthorizer)
			}
			accessToken := bearerAuth.TokenProvider().OAuthToken()
			r.Header.Set("X-Databricks-Azure-SP-Management-Token", accessToken)
			return r, nil
		})
	if err != nil {
		return nil, err
	}
	aa.temporaryPat = tokenResponse
	return aa.temporaryPat, nil
}

func (aa *AzureAuth) configureWithClientSecret() (bool, error) {
	if aa.resourceID() == "" {
		return false, nil
	}
	if !aa.isClientSecretSet() {
		return false, nil
	}
	managementAuthorizer, err := aa.getClientSecretAuthorizer(
		azure.PublicCloud.ServiceManagementEndpoint)
	if err != nil {
		return false, err
	}
	platformAuthorizer, err := aa.getClientSecretAuthorizer(
		AzureDatabricksResourceID)
	if err != nil {
		return false, err
	}
	tokenResponse, err := aa.stuff(
		managementAuthorizer, platformAuthorizer, func(r *http.Request) (*http.Request, error) {
			// TODO: when most of things are done, make this type assertion working, so that SP auth works!!!
			// managementToken := managementAuthorizer.(autorest.BearerAuthorizer).TokenProvider().OAuthToken()
			// r.Header.Set("X-Databricks-Azure-SP-Management-Token", managementToken)
			return r, nil
		})
	if err != nil {
		return false, err
	}
	aa.databricksClient.Token = tokenResponse.TokenValue
	if tokenResponse.TokenInfo != nil {
		aa.databricksClient.tokenCreateTime = tokenResponse.TokenInfo.CreationTime
		aa.databricksClient.tokenExpiryTime = tokenResponse.TokenInfo.ExpiryTime
	}
	// TODO: make it a function
	return true, nil
}

func (aa *AzureAuth) configureWithAzureCLI() (bool, error) {
	if aa.resourceID() == "" {
		return false, nil
	}
	if aa.isClientSecretSet() {
		return false, nil
	}
	managementAuthorizer, err := auth.NewAuthorizerFromCLIWithResource(
		azure.PublicCloud.ServiceManagementEndpoint)
	if err != nil {
		return false, err
	}
	platformAuthorizer, err := auth.NewAuthorizerFromCLIWithResource(
		AzureDatabricksResourceID)
	if err != nil {
		return false, err
	}
	tokenResponse, err := aa.stuff(
		managementAuthorizer, platformAuthorizer, func(r *http.Request) (*http.Request, error) {
			return r, nil
		})
	if err != nil {
		return false, err
	}
	aa.databricksClient.Token = tokenResponse.TokenValue
	if tokenResponse.TokenInfo != nil {
		aa.databricksClient.tokenCreateTime = tokenResponse.TokenInfo.CreationTime
		aa.databricksClient.tokenExpiryTime = tokenResponse.TokenInfo.ExpiryTime
	}
	return true, nil
}

func (aa *AzureAuth) patRequest() model.TokenRequest {
	seconds, err := strconv.Atoi(aa.PATTokenDurationSeconds)
	if err != nil {
		seconds = 60 * 60
	}
	return model.TokenRequest{
		LifetimeSeconds: int32(seconds),
		Comment:         "Secret made via SP",
	}
}

func (aa *AzureAuth) ensureWorkspaceURL(managementAuthorizer autorest.Authorizer) error {
	if aa.databricksClient == nil {
		return fmt.Errorf("DatabricksClient is not configured")
	}
	if aa.databricksClient.Host != "" {
		// TODO: it may have already been set in host property of provider... additional validation needed
		return nil
	}
	log.Println("[DEBUG] Getting Workspace ID via management token.")
	var workspace azureDatabricksWorkspace
	resp, err := aa.databricksClient.genericQuery2(http.MethodGet, aa.workspaceInfoURL(), map[string]string{
		"api-version": "2018-04-01",
	}, func(r *http.Request) (*http.Request, error) {
		r.Header.Set("Cache-Control", "no-cache")
		return autorest.Prepare(r, managementAuthorizer.WithAuthorization())
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &workspace)
	if err != nil {
		return err
	}
	aa.databricksClient.Host = fmt.Sprintf("https://%s/", workspace.Properties.WorkspaceURL)
	return nil
}

// Main function call that gets made and it follows 4 steps at the moment:
// 1. Get Management OAuth Token using management endpoint
// 2. Get Workspace ID and URL
// 3. Get Azure Databricks Platform OAuth Token using Databricks resource id
// 4. Get Azure Databricks Workspace Personal Access Token for the SP (60 min duration)
func (aa *AzureAuth) stuff(managementAuthorizer, platformAuthorizer autorest.Authorizer,
	interceptor func(r *http.Request) (*http.Request, error)) (*model.TokenResponse, error) {
	err := aa.ensureWorkspaceURL(managementAuthorizer)
	if err != nil {
		return nil, err
	}
	log.Println("[DEBUG] Creating workspace token")
	url := fmt.Sprintf("%sapi/2.0/token/create", aa.databricksClient.Host)

	var tokenResponse model.TokenResponse
	resp, err := aa.databricksClient.genericQuery2(http.MethodPost, url,
		aa.patRequest(), func(r *http.Request) (*http.Request, error) {
			r.Header.Set("Cache-Control", "no-cache")
			r, err := interceptor(r)
			if err != nil {
				return nil, err
			}
			r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", aa.resourceID())
			return autorest.Prepare(r, platformAuthorizer.WithAuthorization())
		})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tokenResponse)
	if err != nil {
		return nil, APIError{
			ErrorCode:  "UNKNOWN",
			StatusCode: 200,
			Resource:   "/api/2.0/token/create",
			Message:    fmt.Sprintf("Invalid JSON received: %v", string(resp)),
		}
	}
	return &tokenResponse, nil
}

func (aa *AzureAuth) getClientSecretAuthorizer(resource string) (autorest.Authorizer, error) {
	if aa.authorizer != nil {
		// todo: probably should be two different ones...
		return aa.authorizer, nil
	}
	es := auth.EnvironmentSettings{
		Values:      map[string]string{},
		Environment: azure.PublicCloud,
	}
	es.Values[auth.ClientID] = aa.ClientID
	es.Values[auth.ClientSecret] = aa.ClientSecret
	es.Values[auth.TenantID] = aa.TenantID
	es.Values[auth.Resource] = resource
	return es.GetAuthorizer()
}

// IsZero tells if there are any values inside
func (aa *AzureAuth) IsZero() bool {
	return *aa == AzureAuth{}
}

type azureDatabricksWorkspace struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Sku  struct {
		Name string `json:"name"`
	} `json:"sku"`
	Location   string `json:"location"`
	Properties struct {
		ManagedResourceGroupID string      `json:"managedResourceGroupId"`
		Parameters             interface{} `json:"parameters"`
		ProvisioningState      string      `json:"provisioningState"`
		UIDefinitionURI        string      `json:"uiDefinitionUri"`
		Authorizations         []struct {
			PrincipalID      string `json:"principalId"`
			RoleDefinitionID string `json:"roleDefinitionId"`
		} `json:"authorizations"`
		CreatedBy struct {
			Oid           string `json:"oid"`
			Puid          string `json:"puid"`
			ApplicationID string `json:"applicationId"`
		} `json:"createdBy"`
		UpdatedBy struct {
			Oid           string `json:"oid"`
			Puid          string `json:"puid"`
			ApplicationID string `json:"applicationId"`
		} `json:"updatedBy"`
		CreatedDateTime time.Time `json:"createdDateTime"`
		WorkspaceID     string    `json:"workspaceId"`
		WorkspaceURL    string    `json:"workspaceUrl"`
	} `json:"properties"`
}
