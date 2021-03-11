package common

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureCliAuth(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	// fake expiration date for az mock cli
	os.Setenv("EXPIRE", "15M")

	cnt := []int{0}
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			cnt[0]++
			if req.RequestURI == "/api/2.0/clusters/list-zones" {
				assert.Equal(t, "Bearer ...", req.Header.Get("Authorization"))
				_, err := rw.Write([]byte(`{"zones": ["a", "b", "c"]}`))
				assert.NoError(t, err)
				return
			}
			assert.Fail(t, fmt.Sprintf("Received unexpected call: %s %s",
				req.Method, req.RequestURI))
		}))
	defer server.Close()

	client := DatabricksClient{
		Host: server.URL,
		AzureAuth: AzureAuth{
			ResourceID: "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c",
		},
		InsecureSkipVerify: true,
	}
	err := client.Configure()
	assert.NoError(t, err)

	type ZonesInfo struct {
		Zones       []string `json:"zones,omitempty"`
		DefaultZone string   `json:"default_zone,omitempty"`
	}
	var zi ZonesInfo
	err = client.Get(context.Background(), "/clusters/list-zones", nil, &zi)
	assert.NoError(t, err)
	assert.NotNil(t, zi)
	assert.Len(t, zi.Zones, 3)

	err = client.Get(context.Background(), "/clusters/list-zones", nil, &zi)
	assert.NoError(t, err)

	assert.Equal(t, 2, cnt[0], "There should be only one HTTP call")
}

func TestOAuthToken_CornerCases(t *testing.T) {
	rct := refreshableCliToken{}
	assert.Empty(t, rct.OAuthToken())
}

func TestEnsureFreshWithContext(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	rct := refreshableCliToken{
		token: &adal.Token{
			ExpiresIn: "10",
		},
		lock: &sync.RWMutex{},
	}
	err := rct.EnsureFreshWithContext(context.Background())
	assert.NoError(t, err)
}

func TestRefreshWithContext(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	rct := refreshableCliToken{
		token: &adal.Token{
			ExpiresIn: "10",
		},
		lock: &sync.RWMutex{},
	}
	err := rct.RefreshWithContext(context.Background())
	assert.NoError(t, err)
}

func TestRefreshExchangeWithContext(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	rct := refreshableCliToken{
		token: &adal.Token{
			ExpiresIn: "10",
		},
		lock: &sync.RWMutex{},
	}
	err := rct.RefreshExchangeWithContext(context.Background(), "a")
	assert.NoError(t, err)
}

func TestInternalRefresh_ExitError(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("FAIL", "yes")

	rct := refreshableCliToken{
		token: &adal.Token{
			ExpiresIn: "10",
		},
		lock: &sync.RWMutex{},
	}
	err := rct.refreshInternal("a")
	assert.EqualError(t, err, "Cannot get access token: This is just a failing script.\n")
}

func TestInternalRefresh_OtherError(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "whatever")

	rct := refreshableCliToken{
		token: &adal.Token{
			ExpiresIn: "10",
		},
		lock: &sync.RWMutex{},
	}
	err := rct.refreshInternal("a")
	assert.EqualError(t, err, "Cannot get access token: exec: \"az\": executable file not found in $PATH")
}

func TestInternalRefresh_Corrupt(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("FAIL", "corrupt")

	rct := refreshableCliToken{
		token: &adal.Token{
			ExpiresIn: "10",
		},
		lock: &sync.RWMutex{},
	}
	err := rct.refreshInternal("a")
	assert.EqualError(t, err, "invalid character 'a' looking for beginning of object key string")
}

func TestInternalRefresh_CorruptExpire(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("EXPIRE", "corrupt")

	rct := refreshableCliToken{
		token: &adal.Token{
			ExpiresIn: "10",
		},
		lock: &sync.RWMutex{},
	}
	err := rct.refreshInternal("a")
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Error parsing Token Expiration Date"),
		"Actual message: %s", err.Error())
}

func TestConfigureWithAzureCLI_SP(t *testing.T) {
	aa := AzureAuth{
		ClientID:     "a",
		ClientSecret: "b",
		TenantID:     "c",
		ResourceID:   "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c",
	}
	auth, err := aa.configureWithAzureCLI()
	assert.NoError(t, err)
	assert.Nil(t, auth)
}

func TestConfigureWithAzureCLI(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	// token without expiry in this case
	client, server := singleRequestServer(t, "POST", "/api/2.0/token/create", `{
		"token_value": "abc"
	}`)
	defer server.Close()

	client.AzureAuth.ResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	client.AzureAuth.UsePATForCLI = true

	auth, err := client.AzureAuth.configureWithAzureCLI()
	assert.NoError(t, err)

	err = auth(httptest.NewRequest("GET", "/clusters/list", http.NoBody))
	assert.NoError(t, err)
}

func TestConfigureWithAzureCLI_Error(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	// token without expiry in this case
	client, server := singleRequestServer(t, "POST", "/api/2.0/token/create", `{
		token_value": corrupt
	}`)
	defer server.Close()

	client.AzureAuth.ResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	client.AzureAuth.UsePATForCLI = true

	auth, err := client.AzureAuth.configureWithAzureCLI()
	assert.NoError(t, err)

	err = auth(httptest.NewRequest("GET", "/clusters/list", http.NoBody))
	assert.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Invalid JSON received"),
		"Actual message: %s", err.Error())
}

func TestConfigureWithAzureCLI_NotInstalled(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "whatever")

	// token without expiry in this case
	client, server := singleRequestServer(t, "POST", "/api/2.0/token/create", `{
		"token_value": "abc"
	}`)
	defer server.Close()

	client.AzureAuth.ResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	client.AzureAuth.UsePATForCLI = true

	_, err := client.AzureAuth.configureWithAzureCLI()
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Most likely Azure CLI is not installed"),
		"Actual message: %s", err.Error())
}

func TestCliAuthorizer_Error(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "whatever")
	aa := AzureAuth{}
	_, err := aa.cliAuthorizer("x")
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), `Cannot get access token: exec: "az"`),
		"Actual message: %s", err.Error())
}
