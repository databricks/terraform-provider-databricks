package common

import (
	"context"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func configureAndAuthenticate(dc *DatabricksClient) (*DatabricksClient, error) {
	req, err := http.NewRequest("GET", dc.Config.Host, nil)
	if err != nil {
		return dc, err
	}
	return dc, dc.Config.Authenticate(req)
}

func failsToAuthenticateWith(t *testing.T, dc *DatabricksClient, message string) {
	_, err := configureAndAuthenticate(dc)
	if dc.Config.AuthType != "" {
		log.Printf("[INFO] Auth is: %s", dc.Config.AuthType)
	}
	if assert.NotNil(t, err, "expected to have error: %s", message) {
		assert.True(t, strings.HasPrefix(err.Error(), message), "Expected to have '%s' error, but got '%s'", message, err.Error())
	}
}

func TestDatabricksClientConfigure_Nothing(t *testing.T) {
	t.Setenv("PATH", "testdata:/bin")
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_BasicAuth_NoHost(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Username: "foo",
				Password: "bar",
			},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_BasicAuth(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:     "https://localhost:443",
				Username: "foo",
				Password: "bar",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "basic", dc.Config.AuthType)
}

func TestDatabricksClientConfigure_HostWithoutScheme(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "localhost:443",
				Token: "...",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.Config.AuthType)
	assert.Equal(t, "...", dc.Config.Token)
	assert.Equal(t, "https://localhost:443", dc.Config.Host)
}

func TestDatabricksClientConfigure_Token_NoHost(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token: "dapi345678",
			},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_HostTokensTakePrecedence(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:       "foo",
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.Config.AuthType)
}

func TestDatabricksClientConfigure_BasicAuthDoesNotTakePrecedence(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:       "foo",
				Token:      "configured",
				Username:   "foo",
				Password:   "bar",
				ConfigFile: "testdata/.databrickscfg",
			},
		},
	}, "validate: more than one authorization method configured: basic and pat.")
}

func TestDatabricksClientConfigure_ConfigRead(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				ConfigFile: "testdata/.databrickscfg",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.Config.AuthType)
	assert.Equal(t, "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ", dc.Config.Token)
}

func TestDatabricksClientConfigure_NoHostGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "nohost",
			},
		},
	}, NoAuth+
		". Config: token=***, profile=nohost, config_file=testdata/.databrickscfg")
}

func TestDatabricksClientConfigure_InvalidProfileGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "invalidhost",
			},
		},
	}, "resolve: testdata/.databrickscfg has no invalidhost profile configured. "+
		"Config: token=***, profile=invalidhost, config_file=testdata/.databrickscfg")
}

func TestDatabricksClientConfigure_MissingFile(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.invalid file",
				Profile:    "invalidhost",
			},
		},
	}, NoAuth)
}

func TestDatabricksClientConfigure_InvalidConfigFilePath(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/az",
				Profile:    "invalidhost",
			},
		},
	}, `resolve: cannot parse config file`)
}

func TestDatabricksClient_FormatURL(t *testing.T) {
	client := DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://some.host",
			},
		},
	}
	assert.Equal(t, "https://some.host/#job/123", client.FormatURL("#job/123"))
}

func TestDatabricksIsGcp(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://demo.gcp.databricks.com/",
				Token: "dapi123",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, true, dc.IsGcp())
}

func TestIsAzure_Error(t *testing.T) {
	dc := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "notoken",
			},
		},
	}
	assert.Equal(t, false, dc.IsAzure())
}

func TestClientForHost(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:     "https://accounts.cloud.databricks.com/",
				Username: "abc",
				Password: "bcd",
			},
		},
	})
	assert.NoError(t, err)
	assert.True(t, dc.IsAws())
	cc, err := dc.ClientForHost(context.Background(), "https://e2-workspace.cloud.databricks.com/")
	assert.NoError(t, err)
	assert.Equal(t, dc.Config.Username, cc.Config.Username)
	assert.Equal(t, dc.Config.Password, cc.Config.Password)
	assert.NotEqual(t, dc.Config.Host, cc.Config.Host)
}

func TestClientForHostAuthError(t *testing.T) {
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "notoken",
			},
		},
	}
	_, err := c.ClientForHost(context.Background(), "https://e2-workspace.cloud.databricks.com/")
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_NonsenseAuth(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				AuthType: "nonsense",
			},
		},
	}, NoAuth)
}

func TestGetJWTProperty_AzureCLI_SP(t *testing.T) {
	p, _ := filepath.Abs("./testdata")
	t.Setenv("PATH", p+":/bin")

	aa := DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				AzureClientID:     "a",
				AzureClientSecret: "b",
				AzureTenantID:     "c",
				Host:              "https://adb-1232.azuredatabricks.net",
			},
		},
	}
	tid, err := aa.GetAzureJwtProperty("tid")
	assert.NoError(t, err)
	assert.Equal(t, "c", tid)
}

func TestGetJWTProperty_NonAzure(t *testing.T) {
	p, _ := filepath.Abs("./testdata")
	t.Setenv("PATH", p+":/bin")

	aa := DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://abc.cloud.databricks.com",
				Token: "abc",
			},
		},
	}
	_, err := aa.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "can't get Azure JWT token in non-Azure environment")
}

func TestGetJWTProperty_Authenticate_Fail(t *testing.T) {
	p, _ := filepath.Abs("./testdata")
	t.Setenv("PATH", p+":/bin")
	t.Setenv("FAIL", "yes")
	t.Setenv("ARM_TENANT_ID", "tenant-id")

	client := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host: "https://adb-1232.azuredatabricks.net",
			},
		},
	}
	_, err := client.GetAzureJwtProperty("tid")
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(),
		"default auth: azure-cli: cannot get access token: This is just a failing script"))
}
