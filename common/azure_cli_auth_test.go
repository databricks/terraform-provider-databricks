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
		Host:                      server.URL,
		AzureDatabricksResourceID: "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c",
		InsecureSkipVerify:        true,
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
	assert.EqualError(t, err, "cannot get access token: This is just a failing script.\n")
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
	assert.EqualError(t, err, "cannot get access token: exec: \"az\": executable file not found in $PATH")
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
	assert.EqualError(t, err, "cannot unmarshal CLI result: invalid character 'a' looking for beginning of object key string")
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
	assert.True(t, strings.HasPrefix(err.Error(), "cannot convert to ADAL token: Error parsing Token Expiration Date"),
		"Actual message: %s", err.Error())
}

func TestConfigureWithAzureCLI_SP(t *testing.T) {
	aa := DatabricksClient{
		AzureClientID:             "a",
		AzureClientSecret:         "b",
		AzureTenantID:             "c",
		AzureDatabricksResourceID: "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c",
	}
	ctx := context.Background()
	auth, err := aa.configureWithAzureCLI(ctx)
	assert.NoError(t, err)
	assert.Nil(t, auth)
}

func TestCliAuthorizer_Error(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "whatever")
	aa := DatabricksClient{}
	_, err := aa.cliAuthorizer("x")
	require.Error(t, err)
	require.EqualError(t, err, "cannot refresh: cannot get access token: exec: \"az\": executable file not found in $PATH")
}
