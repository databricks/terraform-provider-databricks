package common

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/golang-jwt/jwt/v4"
)

// List of management information
const (
	AzureDatabricksResourceID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"
)

type tokenRequest struct {
	LifetimeSeconds int64  `json:"lifetime_seconds,omitempty"`
	Comment         string `json:"comment,omitempty"`
}

type tokenResponse struct {
	TokenValue string     `json:"token_value,omitempty"`
	TokenInfo  *tokenInfo `json:"token_info,omitempty"`
}

// tokenInfo is a struct that contains metadata about a given token
type tokenInfo struct {
	TokenID      string `json:"token_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

//
func (aa *DatabricksClient) GetAzureJwtProperty(key string) (interface{}, error) {
	if !aa.IsAzure() {
		return "", fmt.Errorf("can't get Azure JWT token in non-Azure environment")
	}
	if key == "tid" && aa.AzureTenantID != "" {
		return aa.AzureTenantID, nil
	}
	err := aa.Authenticate(context.TODO())
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", aa.Host, nil)
	if err != nil {
		return nil, err
	}
	if err = aa.authVisitor(request); err != nil {
		return nil, err
	}

	header := request.Header.Get("Authorization")
	var stoken string
	if len(header) > 0 && strings.HasPrefix(string(header), "Bearer ") {
		log.Printf("[DEBUG] Got Bearer token")
		stoken = strings.TrimSpace(strings.TrimPrefix(string(header), "Bearer "))
	}

	if stoken == "" {
		return nil, fmt.Errorf("can't obtain Azure JWT token")
	}
	if strings.HasPrefix(stoken, "dapi") {
		return nil, fmt.Errorf("can't use Databricks PAT")
	}
	parser := jwt.Parser{SkipClaimsValidation: true}
	token, _, err := parser.ParseUnverified(stoken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	v, ok := claims[key]
	if !ok {
		return nil, fmt.Errorf("can't find field '%s' in parsed JWT", key)
	}
	return v, nil
}

func (aa *DatabricksClient) getAzureEnvironment() (azure.Environment, error) {
	if aa.AzureEnvironment != nil {
		// used for testing purposes
		return *aa.AzureEnvironment, nil
	}
	if aa.AzurermEnvironment == "" {
		return azure.PublicCloud, nil
	}
	envName := fmt.Sprintf("AZURE%sCLOUD", strings.ToUpper(aa.AzurermEnvironment))
	return azure.EnvironmentFromName(envName)
}

func (aa *DatabricksClient) resourceID() string {
	if aa.AzureDatabricksResourceID != "" {
		if aa.AzureSubscriptionID == "" || aa.AzureResourceGroup == "" {
			res, err := azure.ParseResourceID(aa.AzureDatabricksResourceID)
			if err != nil {
				log.Printf("[ERROR] %s", err)
				return ""
			}
			aa.AzureSubscriptionID = res.SubscriptionID
			aa.AzureResourceGroup = res.ResourceGroup
			aa.AzureWorkspaceName = res.ResourceName
		}
		return aa.AzureDatabricksResourceID
	}
	if aa.AzureSubscriptionID == "" || aa.AzureResourceGroup == "" || aa.AzureWorkspaceName == "" {
		return ""
	}
	r := azure.Resource{
		SubscriptionID: aa.AzureSubscriptionID,
		ResourceGroup:  aa.AzureResourceGroup,
		Provider:       "Microsoft.Databricks",
		ResourceType:   "workspaces",
		ResourceName:   aa.AzureWorkspaceName,
	}
	aa.AzureDatabricksResourceID = r.String()
	return aa.AzureDatabricksResourceID
}

// IsAzureClientSecretSet returns true if client id/secret and tenand id are supplied
func (aa *DatabricksClient) IsAzureClientSecretSet() bool {
	return aa.AzureClientID != "" && aa.AzureClientSecret != "" && aa.AzureTenantID != ""
}

func (aa *DatabricksClient) configureWithAzureClientSecret(ctx context.Context) (func(*http.Request) error, error) {
	if !aa.IsAzure() {
		return nil, nil
	}
	if !aa.IsAzureClientSecretSet() {
		return nil, nil
	}
	log.Printf("[INFO] Generating AAD token for Azure Service Principal")
	return aa.simpleAADRequestVisitor(ctx, aa.getClientSecretAuthorizer, aa.addSpManagementTokenVisitor)
}

func (aa *DatabricksClient) configureWithAzureManagedIdentity(ctx context.Context) (func(*http.Request) error, error) {
	if !aa.IsAzure() {
		return nil, nil
	}
	if !aa.AzureUseMSI {
		return nil, nil
	}
	if !adal.MSIAvailable(ctx, aa.httpClient.HTTPClient) {
		return nil, fmt.Errorf("managed identity is not available")
	}
	log.Printf("[INFO] Using Azure Managed Identity authentication")
	return aa.simpleAADRequestVisitor(ctx, func(resource string) (autorest.Authorizer, error) {
		return auth.MSIConfig{
			Resource: resource,
		}.Authorizer()
	}, aa.addSpManagementTokenVisitor)
}

func (aa *DatabricksClient) addSpManagementTokenVisitor(r *http.Request, management autorest.Authorizer) error {
	log.Printf("[DEBUG] Setting 'X-Databricks-Azure-SP-Management-Token' header")
	ba, ok := management.(*autorest.BearerAuthorizer)
	if !ok {
		return fmt.Errorf("supposed to get BearerAuthorizer, but got %#v", management)
	}
	tokenProvider := ba.TokenProvider()
	if tokenProvider == nil {
		return fmt.Errorf("token provider is nil")
	}
	var err error
	switch rf := tokenProvider.(type) {
	case adal.RefresherWithContext:
		err = rf.EnsureFreshWithContext(r.Context())
	case adal.Refresher:
		err = rf.EnsureFresh()
	}
	if err != nil {
		return fmt.Errorf("cannot refresh AAD token: %w", err)
	}
	accessToken := tokenProvider.OAuthToken()
	r.Header.Set("X-Databricks-Azure-SP-Management-Token", accessToken)
	return nil
}

// go nolint
func (aa *DatabricksClient) simpleAADRequestVisitor(
	ctx context.Context,
	authorizerFactory func(resource string) (autorest.Authorizer, error),
	visitors ...func(r *http.Request, ma autorest.Authorizer) error) (func(r *http.Request) error, error) {
	managementAuthorizer, err := authorizerFactory(aa.AzureEnvironment.ServiceManagementEndpoint)
	if err != nil {
		return nil, fmt.Errorf("cannot authorize management: %w", err)
	}
	err = aa.ensureWorkspaceURL(ctx, managementAuthorizer)
	if err != nil {
		return nil, fmt.Errorf("cannot get workspace: %w", err)
	}
	platformAuthorizer, err := authorizerFactory(AzureDatabricksResourceID)
	if err != nil {
		return nil, fmt.Errorf("cannot authorize databricks: %w", err)
	}
	return func(r *http.Request) error {
		if len(visitors) > 0 {
			err = visitors[0](r, managementAuthorizer)
			if err != nil {
				return err
			}
		}
		resourceID := aa.resourceID()
		if resourceID != "" {
			r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", resourceID)
		}
		_, err = autorest.Prepare(r, platformAuthorizer.WithAuthorization())
		if err != nil {
			return fmt.Errorf("cannot prepare request: %w", err)
		}
		return nil
	}, nil
}

func (aa *DatabricksClient) acquirePAT(
	ctx context.Context,
	factory func(resource string) (autorest.Authorizer, error),
	visitors ...func(r *http.Request, ma autorest.Authorizer) error) (*tokenResponse, error) {
	if aa.temporaryPat != nil {
		// todo: add IsExpired
		return aa.temporaryPat, nil
	}
	if aa.temporaryPat != nil {
		return aa.temporaryPat, nil
	}
	management, err := factory(aa.AzureEnvironment.ServiceManagementEndpoint)
	if err != nil {
		return nil, err
	}
	err = aa.ensureWorkspaceURL(ctx, management)
	if err != nil {
		return nil, err
	}
	token, err := aa.createPAT(ctx, func(r *http.Request) error {
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
		resourceID := aa.resourceID()
		if resourceID != "" {
			r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", resourceID)
		}
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

func (aa *DatabricksClient) patRequest() tokenRequest {
	seconds, err := strconv.ParseInt(aa.AzurePATTokenDurationSeconds, 10, 64)
	if err != nil {
		seconds = 60 * 60
	}
	return tokenRequest{
		LifetimeSeconds: seconds,
		Comment:         "Secret made via Terraform",
	}
}

func maybeExtendAuthzError(err error) error {
	fmtString := "Azure authorization error. Does your SPN have Contributor access to Databricks workspace? %v"
	if e, ok := err.(APIError); ok && e.StatusCode == 403 {
		return fmt.Errorf(fmtString, err)
	} else if strings.Contains(err.Error(), "does not have authorization to perform action") {
		return fmt.Errorf(fmtString, err)
	}
	return err
}

func (aa *DatabricksClient) ensureWorkspaceURL(ctx context.Context,
	managementAuthorizer autorest.Authorizer) error {
	if aa.Host != "" {
		return nil
	}
	resourceID := aa.resourceID()
	if resourceID == "" {
		return fmt.Errorf("somehow resource id is not set")
	}
	log.Println("[DEBUG] Getting Workspace ID via management token.")
	// All azure endpoints typically end with a trailing slash removing it because resourceID starts with slash
	managementResourceURL := strings.TrimSuffix(aa.AzureEnvironment.ResourceManagerEndpoint, "/") + resourceID
	var workspace azureDatabricksWorkspace
	resp, err := aa.genericQuery(ctx, http.MethodGet,
		managementResourceURL,
		map[string]string{
			"api-version": "2018-04-01",
		}, func(r *http.Request) error {
			_, err := autorest.Prepare(r, managementAuthorizer.WithAuthorization())
			if err != nil {
				return maybeExtendAuthzError(err)
			}
			return nil
		})
	if err != nil {
		return maybeExtendAuthzError(err)
	}
	err = json.Unmarshal(resp, &workspace)
	if err != nil {
		return err
	}
	aa.Host = fmt.Sprintf("https://%s/", workspace.Properties.WorkspaceURL)
	return nil
}

func (aa *DatabricksClient) createPAT(ctx context.Context,
	interceptor func(r *http.Request) error) (tr tokenResponse, err error) {
	log.Println("[DEBUG] Creating workspace token")
	url := fmt.Sprintf("%sapi/2.0/token/create", aa.Host)
	body, err := aa.genericQuery(ctx,
		http.MethodPost, url, aa.patRequest(), interceptor)
	if err != nil {
		return
	}
	err = aa.unmarshall("/api/2.0/token/create", body, &tr)
	return
}

func (aa *DatabricksClient) getClientSecretAuthorizer(resource string) (autorest.Authorizer, error) {
	if aa.azureAuthorizer != nil {
		return aa.azureAuthorizer, nil
	}
	if resource != AzureDatabricksResourceID {
		es := auth.EnvironmentSettings{
			Values: map[string]string{
				auth.ClientID:     aa.AzureClientID,
				auth.ClientSecret: aa.AzureClientSecret,
				auth.TenantID:     aa.AzureTenantID,
				auth.Resource:     resource,
			},
			Environment: *aa.AzureEnvironment,
		}
		return es.GetAuthorizer()
	}
	platformTokenOAuthCfg, err := adal.NewOAuthConfigWithAPIVersion(
		aa.AzureEnvironment.ActiveDirectoryEndpoint,
		aa.AzureTenantID,
		nil)
	if err != nil {
		return nil, maybeExtendAuthzError(err)
	}
	spt, err := adal.NewServicePrincipalToken(
		*platformTokenOAuthCfg,
		aa.AzureClientID,
		aa.AzureClientSecret,
		AzureDatabricksResourceID)
	if err != nil {
		return nil, maybeExtendAuthzError(err)
	}
	return autorest.NewBearerAuthorizer(spt), nil
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
