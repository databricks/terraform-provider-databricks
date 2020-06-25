package databricks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	urlParse "net/url"
	"time"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

// List of management information
const (
	ADBResourceID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"
)

// TokenPayload contains all the auth information for azure sp authentication
type TokenPayload struct {
	ManagedResourceGroup string
	AzureRegion          string
	WorkspaceName        string
	ResourceGroup        string
	SubscriptionID       string
	ClientSecret         string
	ClientID             string
	TenantID             string
	PatTokenDuration     int32
}

type Workspace struct {
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

// WsProps contains information about the workspace properties
type WsProps struct {
	ManagedResourceGroupID string `json:"managedResourceGroupId"`
}

// WorkspaceRequest contains the request information for getting workspace information
type WorkspaceRequest struct {
	Properties *WsProps `json:"properties"`
	Name       string   `json:"name"`
	Location   string   `json:"location"`
}

func (t *TokenPayload) getManagementToken() (string, error) {
	log.Println("[DEBUG] Creating Azure Databricks management OAuth token.")
	mgmtTokenOAuthCfg, err := adal.NewOAuthConfigWithAPIVersion(azure.PublicCloud.ActiveDirectoryEndpoint,
		t.TenantID,
		nil)
	if err != nil {
		return "", err
	}
	mgmtToken, err := adal.NewServicePrincipalToken(
		*mgmtTokenOAuthCfg,
		t.ClientID,
		t.ClientSecret,
		azure.PublicCloud.ServiceManagementEndpoint)
	if err != nil {
		return "", err
	}
	err = mgmtToken.Refresh()
	if err != nil {
		return "", err
	}

	return mgmtToken.OAuthToken(), nil
}

func (t *TokenPayload) getWorkspace(config *service.DBApiClientConfig, managementToken string) (*Workspace, error) {
	log.Println("[DEBUG] Getting Workspace ID via management token.")
	// Escape all the ids
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s"+
		"/providers/Microsoft.Databricks/workspaces/%s",
		urlParse.PathEscape(t.SubscriptionID),
		urlParse.PathEscape(t.ResourceGroup),
		urlParse.PathEscape(t.WorkspaceName))
	headers := map[string]string{
		"Content-Type":  "application/json",
		"cache-control": "no-cache",
		"Authorization": "Bearer " + managementToken,
	}
	type apiVersion struct {
		APIVersion string `url:"api-version"`
	}
	uriPayload := apiVersion{
		APIVersion: "2018-04-01",
	}
	resp, err := service.PerformQuery(config, http.MethodGet, url, "2.0", headers, false, true, uriPayload, nil)
	if err != nil {
		return nil, err
	}

	var workspace Workspace
	err = json.Unmarshal(resp, &workspace)
	if err != nil {
		return nil, err
	}
	return &workspace, err
}

func (t *TokenPayload) getADBPlatformToken() (string, error) {
	log.Println("[DEBUG] Creating Azure Databricks management OAuth token.")
	platformTokenOAuthCfg, err := adal.NewOAuthConfigWithAPIVersion(azure.PublicCloud.ActiveDirectoryEndpoint,
		t.TenantID,
		nil)
	if err != nil {
		return "", err
	}
	platformToken, err := adal.NewServicePrincipalToken(
		*platformTokenOAuthCfg,
		t.ClientID,
		t.ClientSecret,
		ADBResourceID)
	if err != nil {
		return "", err
	}

	err = platformToken.Refresh()
	if err != nil {
		return "", err
	}

	return platformToken.OAuthToken(), nil
}

func getWorkspaceAccessToken(config *service.DBApiClientConfig, managementToken, adbWorkspaceUrl, adbWorkspaceResourceID, adbPlatformToken string) (*model.TokenResponse, error) {
	log.Println("[DEBUG] Creating workspace token")
	url := adbWorkspaceUrl + "/api/2.0/token/create"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Databricks-Azure-Workspace-Resource-Id": adbWorkspaceResourceID,
		"X-Databricks-Azure-SP-Management-Token":   managementToken,
		"cache-control":                            "no-cache",
		"Authorization":                            "Bearer " + adbPlatformToken,
	}

	var tokenResponse model.TokenResponse
	resp, err := service.PerformQuery(config, http.MethodPost, url, "2.0",
		headers, true, true, model.TokenRequest{
			LifetimeSeconds: int32(3600),
			Comment:         "Secret made via SP",
		}, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tokenResponse)
	if err != nil {
		return nil, err
	}
	return &tokenResponse, nil
}

// Main function call that gets made and it follows 4 steps at the moment:
// 1. Get Management OAuth Token using management endpoint
// 2. Get Workspace ID and URL
// 3. Get Azure Databricks Platform OAuth Token using Databricks resource id
// 4. Get Azure Databricks Workspace Personal Access Token for the SP (60 min duration)
func (t *TokenPayload) initWorkspaceAndGetClient(config *service.DBApiClientConfig) error {
	// Get management token
	managementToken, err := t.getManagementToken()
	if err != nil {
		return err
	}

	// Get workspace access token
	adbWorkspace, err := t.getWorkspace(config, managementToken)
	if err != nil {
		return err
	}
	adbWorkspaceUrl := "https://" + adbWorkspace.Properties.WorkspaceURL

	// Get platform token
	adbPlatformToken, err := t.getADBPlatformToken()
	if err != nil {
		return err
	}

	// Get workspace personal access token
	workspaceAccessTokenResp, err := getWorkspaceAccessToken(config, managementToken, adbWorkspaceUrl, adbWorkspace.ID, adbPlatformToken)
	if err != nil {
		return err
	}

	config.Host = adbWorkspaceUrl
	config.Token = workspaceAccessTokenResp.TokenValue
	if workspaceAccessTokenResp.TokenInfo != nil {
		config.TokenExpiryTime = workspaceAccessTokenResp.TokenInfo.ExpiryTime
	}

	return nil
}
