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
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// List of management information
const (
	AzureDatabricksResourceID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"
)

// AzureAuth contains all the auth information for azure sp authentication
type AzureAuth struct {
	ManagedResourceGroup string
	AzureRegion          string

	WorkspaceName  string
	ResourceGroup  string
	SubscriptionID string

	// azurerm_databricks_workspace.this.id ->
	// /subscriptions/{subscription}/resourceGroups/{rg}/providers/Microsoft.Databricks/workspaces/{name}
	ResourceID string

	ClientSecret string
	ClientID     string
	TenantID     string

	// temporary workaround for SP-based auth
	PATTokenDurationSeconds string

	// private property to give resource access
	databricksClient *DatabricksClient

	azureManagementEndpoint string
	authorizer              autorest.Authorizer

	temporaryPat *model.TokenResponse
}

var authorizerMutex sync.Mutex

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

func (aa *AzureAuth) isClientSecretSet() bool {
	return aa.ClientID != "" && aa.ClientSecret != "" && aa.TenantID != ""
}

func (aa *AzureAuth) configureWithClientSecret() (func(r *http.Request) error, error) {
	if aa.resourceID() == "" {
		return nil, nil
	}
	if !aa.isClientSecretSet() {
		return nil, nil
	}
	log.Printf("[INFO] Using Azure Service Principal client secret authentication")
	return func(r *http.Request) error {
		pat, err := aa.acquirePAT(aa.getClientSecretAuthorizer,
			func(r *http.Request, management autorest.Authorizer) error {
				log.Printf("[DEBUG] Setting 'X-Databricks-Azure-SP-Management-Token' header")
				bearerAuth, ok := management.(*autorest.BearerAuthorizer)
				if !ok {
					return fmt.Errorf("Supposed to get BearerAuthorizer, but got %v", management)
				}
				accessToken := bearerAuth.TokenProvider().OAuthToken()
				r.Header.Set("X-Databricks-Azure-SP-Management-Token", accessToken)
				return nil
			})
		if err != nil {
			return err
		}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pat.TokenValue))
		return nil
	}, nil
}

func (aa *AzureAuth) configureWithAzureCLI() (func(r *http.Request) error, error) {
	if aa.resourceID() == "" {
		return nil, nil
	}
	if aa.isClientSecretSet() {
		return nil, nil
	}
	log.Printf("[INFO] Using Azure CLI authentication")
	return func(r *http.Request) error {
		pat, err := aa.acquirePAT(auth.NewAuthorizerFromCLIWithResource)
		if err != nil {
			return err
		}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pat.TokenValue))
		return nil
	}, nil
}

// func (aa *AzureAuth) simpleAADRequestVisitor(authorizerFactory func(resource string) (autorest.Authorizer, error),
// 	interceptor ...func(ma autorest.Authorizer) error) (func(r *http.Request) error, error) {
// 	managementAuthorizer, err := authorizerFactory(azure.PublicCloud.ServiceManagementEndpoint)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = aa.ensureWorkspaceURL(managementAuthorizer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	platformAuthorizer, err := authorizerFactory(AzureDatabricksResourceID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return func(r *http.Request) error {
// 		if len(interceptor) > 0 {
// 			err = interceptor[0](managementAuthorizer)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", aa.resourceID())
// 		_, err = autorest.Prepare(r, platformAuthorizer.WithAuthorization())
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}, nil
// }

func (aa *AzureAuth) acquirePAT(
	factory func(resource string) (autorest.Authorizer, error),
	visitors ...func(r *http.Request, ma autorest.Authorizer) error) (*model.TokenResponse, error) {
	if aa.temporaryPat != nil {
		// todo: add IsExpired
		return aa.temporaryPat, nil
	}
	authorizerMutex.Lock()
	defer authorizerMutex.Unlock()
	if aa.temporaryPat != nil {
		return aa.temporaryPat, nil
	}
	management, err := factory(azure.PublicCloud.ServiceManagementEndpoint)
	if err != nil {
		return nil, err
	}
	err = aa.ensureWorkspaceURL(management)
	if err != nil {
		return nil, err
	}
	// pt, err := aa.getADBPlatformToken()
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("%v", len(pt))
	token, err := aa.createPAT(func(r *http.Request) error {
		if len(visitors) > 0 {
			err = visitors[0](r, management)
			if err != nil {
				return err
			}
		}
		platform, err := factory(AzureDatabricksResourceID)
		if err != nil {
			return err
		}
		r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", aa.resourceID())
		_, err = autorest.Prepare(r, platform.WithAuthorization())
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	aa.temporaryPat = &token
	return aa.temporaryPat, nil
}

func (aa *AzureAuth) patRequest() model.TokenRequest {
	seconds, err := strconv.Atoi(aa.PATTokenDurationSeconds)
	if err != nil {
		seconds = 60 * 60
	}
	return model.TokenRequest{
		LifetimeSeconds: int32(seconds),
		Comment:         "Secret made via Terraform",
	}
}

func (aa *AzureAuth) ensureWorkspaceURL(managementAuthorizer autorest.Authorizer) error {
	if aa.databricksClient == nil {
		return fmt.Errorf("DatabricksClient is not configured")
	}
	if aa.databricksClient.Host != "" {
		return nil
	}
	resourceID := aa.resourceID()
	if resourceID == "" {
		return fmt.Errorf("Somehow resource id is not set")
	}
	log.Println("[DEBUG] Getting Workspace ID via management token.")
	endpoint := "https://management.azure.com"
	if aa.azureManagementEndpoint != "" {
		// sets endpoint specified in unit test
		endpoint = aa.azureManagementEndpoint
	}
	var workspace azureDatabricksWorkspace
	resp, err := aa.databricksClient.genericQuery2(http.MethodGet,
		endpoint+resourceID,
		map[string]string{
			"api-version": "2018-04-01",
		}, func(r *http.Request) error {
			_, err := autorest.Prepare(r, managementAuthorizer.WithAuthorization())
			if err != nil {
				return err
			}
			return nil
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

func (aa *AzureAuth) createPAT(interceptor func(r *http.Request) error) (tr model.TokenResponse, err error) {
	log.Println("[DEBUG] Creating workspace token")
	url := fmt.Sprintf("%sapi/2.0/token/create", aa.databricksClient.Host)
	body, err := aa.databricksClient.genericQuery2(
		http.MethodPost, url, aa.patRequest(), interceptor)
	if err != nil {
		return
	}
	err = aa.databricksClient.unmarshall("/api/2.0/token/create", body, &tr)
	return
}

func (aa *AzureAuth) getClientSecretAuthorizer(resource string) (autorest.Authorizer, error) {
	if aa.authorizer != nil {
		// todo: probably should be two different ones...
		return aa.authorizer, nil
	}
	if resource != AzureDatabricksResourceID {
		es := auth.EnvironmentSettings{
			Values: map[string]string{
				auth.ClientID:     aa.ClientID,
				auth.ClientSecret: aa.ClientSecret,
				auth.TenantID:     aa.TenantID,
				auth.Resource:     resource,
			},
			Environment: azure.PublicCloud,
		}
		return es.GetAuthorizer()
	}
	platformTokenOAuthCfg, err := adal.NewOAuthConfigWithAPIVersion(
		azure.PublicCloud.ActiveDirectoryEndpoint,
		aa.TenantID,
		nil)
	if err != nil {
		return nil, err
	}
	spt, err := adal.NewServicePrincipalToken(
		*platformTokenOAuthCfg,
		aa.ClientID,
		aa.ClientSecret,
		AzureDatabricksResourceID)
	if err != nil {
		return nil, err
	}
	return autorest.NewBearerAuthorizer(spt), nil
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
