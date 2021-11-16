package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	log.Printf("[INFO] Using Azure Service Principal client secret authentication with AAD token")
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
		_, err = autorest.Prepare(r, platformAuthorizer.WithAuthorization())
		if err != nil {
			return fmt.Errorf("cannot prepare request: %w", err)
		}
		return nil
	}, nil
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
