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
	},
		"authentication is not configured for provider")
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
	}, "authentication is not configured for provider.")
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
	}, "authentication is not configured for provider.")
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

func TestDatabricksClientConfigure_BasicAuthTakePrecedence(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:       "foo",
				Token:      "configured",
				Username:   "foo",
				Password:   "bar",
				ConfigFile: "testdata/.databrickscfg",
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.Config.AuthType)
	assert.Equal(t, "configured", dc.Config.Token)
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
	assert.Equal(t, "databricks-cli", dc.Config.AuthType)
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
	}, "cannot configure databricks-cli auth: config file "+
		"testdata/.databrickscfg is corrupt: cannot find host in nohost profile.")
}

func TestDatabricksClientConfigure_NoTokenGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Token:      "connfigured",
				ConfigFile: "testdata/.databrickscfg",
				Profile:    "notoken",
			},
		},
	}, "cannot configure databricks-cli auth: config file "+
		"testdata/.databrickscfg is corrupt: cannot find token in notoken profile.")
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
	}, "cannot configure databricks-cli auth: testdata/.databrickscfg "+
		"has no invalidhost profile configured")
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
	}, "authentication is not configured for provider.")
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
	}, "cannot configure databricks-cli auth: cannot parse config file")
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
	if assert.NotNil(t, err) {
		assert.True(t, strings.HasPrefix(err.Error(),
			"cannot authenticate parent client: cannot configure databricks-cli auth"), err.Error())
	}
}

func TestDatabricksClientConfigure_NonsenseAuth(t *testing.T) {
	defer CleanupEnvironment()()
	failsToAuthenticateWith(t, &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				AuthType: "nonsense",
			},
		},
	}, "cannot configure nonsense auth.")
}

func TestConfigAttributeSetNonsense(t *testing.T) {
	err := (&ConfigAttribute{
		Kind: reflect.Chan,
	}).Set(&DatabricksClient{}, 1)
	assert.EqualError(t, err, "cannot set  of unknown type Chan")
}

func TestDatabricksClientFixHost(t *testing.T) {
	hostForInput := func(in string) (string, error) {
		client := &DatabricksClient{
			DatabricksClient: &client.DatabricksClient{
				Config: &config.Config{
					Host: in,
				},
			},
		}
		return client.Config.Host, nil
	}

	{
		// Strip trailing slash.
		out, err := hostForInput("https://accounts.gcp.databricks.com/")
		assert.Nil(t, err)
		assert.Equal(t, out, "https://accounts.gcp.databricks.com")
	}

	{
		// Keep port.
		out, err := hostForInput("https://accounts.gcp.databricks.com:443")
		assert.Nil(t, err)
		assert.Equal(t, out, "https://accounts.gcp.databricks.com:443")
	}

	{
		// Default scheme.
		out, err := hostForInput("accounts.gcp.databricks.com")
		assert.Nil(t, err)
		assert.Equal(t, out, "https://accounts.gcp.databricks.com")
	}

	{
		// Default scheme with port.
		out, err := hostForInput("accounts.gcp.databricks.com:443")
		assert.Nil(t, err)
		assert.Equal(t, out, "https://accounts.gcp.databricks.com:443")
	}

	{
		// Return error.
		_, err := hostForInput("://@@@accounts.gcp.databricks.com/")
		assert.NotNil(t, err)
	}
}
