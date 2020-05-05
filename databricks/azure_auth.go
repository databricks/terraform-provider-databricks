package databricks

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"log"
	"net/http"
	urlParse "net/url"
	"strings"
	"time"
)

// List of management information
const (
	ADBResourceID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"
)

// AzureAuth is a struct that contains information about the azure sp authentication
type AzureAuth struct {
	TokenPayload           *TokenPayload
	ManagementToken        string
	AdbWorkspaceResourceID string
	AdbAccessToken         string
	AdbPlatformToken       string
}

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

func (a *AzureAuth) getManagementToken(config *service.DBApiClientConfig) error {
	log.Println("[DEBUG] Creating Azure Databricks management OAuth token.")
	mgmtTokenOAuthCfg, err := adal.NewOAuthConfigWithAPIVersion(azure.PublicCloud.ActiveDirectoryEndpoint,
		a.TokenPayload.TenantID,
		nil)
	if err != nil {
		return err
	}
	mgmtToken, err := adal.NewServicePrincipalToken(
		*mgmtTokenOAuthCfg,
		a.TokenPayload.ClientID,
		a.TokenPayload.ClientSecret,
		azure.PublicCloud.ServiceManagementEndpoint)
	if err != nil {
		return err
	}
	err = mgmtToken.EnsureFresh()
	if err != nil {
		return err
	}
	a.ManagementToken = mgmtToken.OAuthToken()
	return nil
}

func (a *AzureAuth) getWorkspaceID(config *service.DBApiClientConfig) error {
	log.Println("[DEBUG] Getting Workspace ID via management token.")
	// Escape all the ids
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s"+
		"/providers/Microsoft.Databricks/workspaces/%s?api-version=2018-04-01",
		urlParse.PathEscape(a.TokenPayload.SubscriptionID),
		urlParse.PathEscape(a.TokenPayload.ResourceGroup),
		urlParse.PathEscape(a.TokenPayload.WorkspaceName))
	payload := &WorkspaceRequest{
		Properties: &WsProps{ManagedResourceGroupID: "/subscriptions/" + a.TokenPayload.SubscriptionID + "/resourceGroups/" + a.TokenPayload.ManagedResourceGroup},
		Name:       a.TokenPayload.WorkspaceName,
		Location:   a.TokenPayload.AzureRegion,
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"cache-control": "no-cache",
		"Authorization": "Bearer " + a.ManagementToken,
	}

	var responseMap map[string]interface{}
	resp, err := service.PerformQuery(config, http.MethodPut, url, "2.0", headers, true, true, payload, nil)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &responseMap)
	if err != nil {
		return err
	}
	a.AdbWorkspaceResourceID = responseMap["id"].(string)
	return err
}

func (a *AzureAuth) getADBPlatformToken(clientConfig *service.DBApiClientConfig) error {
	log.Println("[DEBUG] Creating Azure Databricks management OAuth token.")
	platformTokenOAuthCfg, err := adal.NewOAuthConfigWithAPIVersion(azure.PublicCloud.ActiveDirectoryEndpoint,
		a.TokenPayload.TenantID,
		nil)
	if err != nil {
		return err
	}
	platformToken, err := adal.NewServicePrincipalToken(
		*platformTokenOAuthCfg,
		a.TokenPayload.ClientID,
		a.TokenPayload.ClientSecret,
		ADBResourceID)
	if err != nil {
		return err
	}

	err = platformToken.EnsureFresh()
	if err != nil {
		return err
	}
	a.AdbPlatformToken = platformToken.OAuthToken()
	return nil
}

func (a *AzureAuth) initWorkspaceAndGetClient(config *service.DBApiClientConfig) (service.DBApiClient, error) {
	var dbClient service.DBApiClient

	// Get management token
	err := a.getManagementToken(config)
	if err != nil {
		return dbClient, err
	}

	// Get workspace access token
	err = a.getWorkspaceID(config)
	if err != nil {
		return dbClient, err
	}

	// Get platform token
	err = a.getADBPlatformToken(config)
	if err != nil {
		return dbClient, err
	}

	var newOption service.DBApiClientConfig
	// Host for azure
	newOption.Host = "https://" + a.TokenPayload.AzureRegion + ".azuredatabricks.net"
	// Host for platform token
	newOption.Token = a.AdbPlatformToken
	// Headers to use aad tokens
	newOption.DefaultHeaders = map[string]string{
		"X-Databricks-Azure-Workspace-Resource-Id": a.AdbWorkspaceResourceID,
		"X-Databricks-Azure-SP-Management-Token":   a.ManagementToken,
	}
	dbClient.SetConfig(&newOption)

	// So when the workspace is initially created sometimes it fails to perform api calls so this is a simple test
	// to verify that the workspace has been created successfully. The retry is intentional as sometimes after workspace
	// creation the API's will not work correctly. This may also
	err = validateWorkspaceApis(10, 30, func(attempt int) error {
		_, err = dbClient.Clusters().ListNodeTypes()
		return err
	})

	return dbClient, err
}

func validateWorkspaceApis(sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration, do func(attempt int) error) error {
	errChan := make(chan error, 1)
	var timeoutBool = false
	var attempts int
	var expectedError error
	apisAreNotYetReadErr := "com.databricks.backend.manager.util.UnknownWorkerEnvironmentException: Unknown worker environment WorkerEnvId"
	go func(attempts *int, expectedError *error, timeout *bool) {
		for {
			err := do(*attempts)
			// Timeout and terminate go routine so it does not leak
			if *timeout {
				errChan <- err
				return
			}
			if err == nil {
				errChan <- err
				return
			}
			if !strings.Contains(err.Error(), apisAreNotYetReadErr) {
				errChan <- err
				return
			}
			log.Println(fmt.Sprintf("Waiting for cluster apis to not throw error: %v", err))
			*attempts++
			*expectedError = err
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}(&attempts, &expectedError, &timeoutBool)
	select {
	case err := <-errChan:
		if err == nil {
			log.Printf("Returned nil error after %v attempts\n", attempts)
		}
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		timeoutBool = true
		return fmt.Errorf("timed out waiting for ready state after %v attempts with error %v", attempts, expectedError)
	}
}
