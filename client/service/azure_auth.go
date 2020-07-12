package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func (c *DatabricksClient) configureAzureAuth() (bool, error) {
	if c.AzureAuth.WorkspaceName == "" {
		return false, nil
	}
	c.AzureAuth.databricksClient = c
	patTokenDuration, err := strconv.Atoi(c.AzureAuth.PATTokenDurationSeconds)
	if err != nil {
		return false, fmt.Errorf("failed to parse pat_token_duration_seconds[%v], %w", patTokenDuration, err)
	}
	c.AzureAuth.patTokenSeconds = int32(patTokenDuration)
	//c.customAuthorizer = c.AzureAuth.initWorkspaceAndGetClient

	return true, nil
}

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

	PATTokenDurationSeconds string
	PatTokenSeconds         int32
	patTokenSeconds         int32

	// private property to give resource access
	databricksClient *DatabricksClient
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
	return fmt.Sprintf("https://management.azure.com%s", aa.resourceID())
}

func (aa *AzureAuth) isClientSecretSet() bool {
	return aa.ClientID != "" && aa.ClientSecret != "" && aa.TenantID != ""
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
	aa.databricksClient.Token = tokenResponse.TokenValue
	if tokenResponse.TokenInfo != nil {
		aa.databricksClient.tokenCreateTime = tokenResponse.TokenInfo.CreationTime
		aa.databricksClient.tokenExpiryTime = tokenResponse.TokenInfo.ExpiryTime
	}
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

// Main function call that gets made and it follows 4 steps at the moment:
// 1. Get Management OAuth Token using management endpoint
// 2. Get Workspace ID and URL
// 3. Get Azure Databricks Platform OAuth Token using Databricks resource id
// 4. Get Azure Databricks Workspace Personal Access Token for the SP (60 min duration)
func (aa *AzureAuth) stuff(managementAuthorizer, platformAuthorizer autorest.Authorizer,
	interceptor func(r *http.Request) (*http.Request, error)) (*model.TokenResponse, error) {
	log.Println("[DEBUG] Getting Workspace ID via management token.")
	var workspace azureDatabricksWorkspace
	resp, err := aa.databricksClient.genericQuery2(http.MethodGet, aa.workspaceInfoURL(), map[string]string{
		"api-version": "2018-04-01",
	}, func(r *http.Request) (*http.Request, error) {
		r.Header.Set("Cache-Control", "no-cache")
		return autorest.Prepare(r, managementAuthorizer.WithAuthorization())
	})
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resp, &workspace)
	aa.databricksClient.Host = fmt.Sprintf("https://%s/", workspace.Properties.WorkspaceURL)
	log.Println("[DEBUG] Creating workspace token")
	url := fmt.Sprintf("%sapi/2.0/token/create", aa.databricksClient.Host)
	tokenLifetimeSeconds := (time.Duration(aa.patTokenSeconds) * time.Second).Seconds()
	var tokenResponse model.TokenResponse
	resp, err = aa.databricksClient.genericQuery2(http.MethodPost, url, model.TokenRequest{
		LifetimeSeconds: int32(tokenLifetimeSeconds),
		Comment:         "Secret made via SP",
	}, func(r *http.Request) (*http.Request, error) {
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
	es := auth.EnvironmentSettings{}
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
