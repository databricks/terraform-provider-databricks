package common

import (
	"context"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
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
		assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
	}
}

func TestDatabricksClientConfigure_Nothing(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{},
		},
	}, "default auth: cannot configure default credentials")
}

func TestDatabricksClientConfigure_BasicAuth_NoHost(t *testing.T) {
	defer CleanupEnvironment()()
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Username: "foo",
				Password: "bar",
			},
		},
	}, "default auth: cannot configure default credentials")
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
	defer CleanupEnvironment()()
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token: "dapi345678",
			},
		},
	}, "default auth: cannot configure default credentials")
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
	}, "default auth: cannot configure default credentials. "+
		"Config: token=***, profile=nohost, config_file=testdata/.databrickscfg")
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
	}, "default auth: cannot configure default credentials.")
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

// func TestDatabricksClient_Authenticate(t *testing.T) {
// 	defer CleanupEnvironment()()
// 	dc := DatabricksClient{}
// 	err := dc.Configure("account_id", "username", "password")
// 	os.Setenv("DATABRICKS_PASSWORD", ".")
// 	assert.NoError(t, err)
// 	err = dc.Authenticate(context.WithValue(context.Background(), IsData, "yes"))
// 	assert.EqualError(t, err, "workspace is most likely not created yet, because the `host` is empty. "+
// 		"Please add `depends_on = [databricks_mws_workspaces.this]` or `depends_on = [azurerm_databricks"+
// 		"_workspace.this]` to every data resource. See https://www.terraform.io/docs/language/resources/behavior.html more info. "+
// 		"Attributes used: account_id, username. Environment variables used: DATABRICKS_PASSWORD. "+
// 		"Please check https://registry.terraform.io/providers/databricks/databricks/latest/docs#authentication for details")
// }

// func TestDatabricksClient_AuthenticateAzure(t *testing.T) {
// 	defer CleanupEnvironment()()
// 	os.Setenv("ARM_CLIENT_SECRET", ".")
// 	os.Setenv("ARM_CLIENT_ID", ".")
// 	dc := DatabricksClient{}
// 	err := dc.Configure("azure_client_id", "azure_client_secret")
// 	assert.NoError(t, err)
// 	err = dc.Authenticate(context.WithValue(context.Background(), IsData, "yes"))
// 	assert.EqualError(t, err, "workspace is most likely not created yet, because the `host` is empty. "+
// 		"Please add `depends_on = [databricks_mws_workspaces.this]` or `depends_on = [azurerm_databricks"+
// 		"_workspace.this]` to every data resource. See https://www.terraform.io/docs/language/resources/"+
// 		"behavior.html more info. Environment variables used: ARM_CLIENT_SECRET, ARM_CLIENT_ID. "+
// 		"Please check https://registry.terraform.io/providers/databricks/databricks/latest/docs#authentication for details")
// }

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
	defer CleanupEnvironment()()
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				AuthType: "nonsense",
			},
		},
	}, "default auth: cannot configure default credentials")
}

func TestConfigAttributeSetNonsense(t *testing.T) {
	err := (&ConfigAttribute{
		Kind: reflect.Chan,
	}).Set(&DatabricksClient{}, 1)
	assert.EqualError(t, err, "cannot set  of unknown type Chan")
}
