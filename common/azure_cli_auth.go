package common

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/cli"
)

type refreshableCliToken struct {
	resource       string
	token          *adal.Token
	lock           *sync.RWMutex
	refreshMinutes int
}

// OAuthToken implements adal.OAuthTokenProvider
func (rct *refreshableCliToken) OAuthToken() string {
	if rct.token == nil {
		return ""
	}
	return rct.token.OAuthToken()
}

// EnsureFreshWithContext implements adal.RefresherWithContext
func (rct *refreshableCliToken) EnsureFreshWithContext(ctx context.Context) error {
	refreshInterval := time.Duration(rct.refreshMinutes) * time.Minute
	if rct.token != nil && !rct.token.WillExpireIn(refreshInterval) {
		return nil
	}
	rct.lock.Lock()
	defer rct.lock.Unlock()
	if rct.token != nil && !rct.token.WillExpireIn(refreshInterval) {
		return nil
	}
	return rct.refreshInternal(rct.resource)
}

// RefreshWithContext implements adal.RefresherWithContext
func (rct *refreshableCliToken) RefreshWithContext(ctx context.Context) error {
	rct.lock.Lock()
	defer rct.lock.Unlock()
	return rct.refreshInternal(rct.resource)
}

// RefreshExchangeWithContext implements adal.RefresherWithContext
func (rct *refreshableCliToken) RefreshExchangeWithContext(ctx context.Context, resource string) error {
	rct.lock.Lock()
	defer rct.lock.Unlock()
	return rct.refreshInternal(rct.resource)
}

func (rct *refreshableCliToken) refreshInternal(resource string) (err error) {
	out, err := exec.Command("az", "account", "get-access-token", "--resource", resource).Output()
	if ee, ok := err.(*exec.ExitError); ok {
		err = fmt.Errorf("Cannot get access token: %s", string(ee.Stderr))
		return
	}
	if err != nil {
		err = fmt.Errorf("Cannot get access token: %v", err)
		return
	}
	var cliToken cli.Token
	err = json.Unmarshal(out, &cliToken)
	if err != nil {
		return
	}
	token, err := cliToken.ToADALToken()
	if err != nil {
		return err
	}
	log.Printf("[INFO] Refreshed OAuth token for %s from Azure CLI, which expires on %s", resource, cliToken.ExpiresOn)
	rct.token = &token
	return
}

func (aa *AzureAuth) cliAuthorizer(resource string) (autorest.Authorizer, error) {
	rct := refreshableCliToken{
		lock:           &sync.RWMutex{},
		resource:       resource,
		refreshMinutes: 6,
	}
	err := rct.refreshInternal(resource)
	if err != nil {
		return nil, err
	}
	return autorest.NewBearerAuthorizer(&rct), nil
}

func (aa *AzureAuth) configureWithAzureCLI() (func(r *http.Request) error, error) {
	if aa.resourceID() == "" {
		return nil, nil
	}
	if aa.IsClientSecretSet() {
		return nil, nil
	}
	// verify that Azure CLI is authenticated
	_, err := cli.GetTokenFromCLI(AzureDatabricksResourceID)
	if err != nil {
		if err.Error() == "Invoking Azure CLI failed with the following error: " {
			return nil, fmt.Errorf("Most likely Azure CLI is not installed. " +
				"See https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest for details.")
		}
		return nil, err
	}
	if aa.UsePATForCLI {
		log.Printf("[INFO] Using Azure CLI authentication with session-generated PAT")
		return func(r *http.Request) error {
			pat, err := aa.acquirePAT(r.Context(), aa.cliAuthorizer)
			if err != nil {
				return err
			}
			r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pat.TokenValue))
			return nil
		}, nil
	}
	log.Printf("[INFO] Using Azure CLI authentication with AAD tokens")
	return aa.simpleAADRequestVisitor(context.TODO(), aa.cliAuthorizer)
}
