package common

import (
	"context"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func configureAndAuthenticate(dc *DatabricksClient) (*DatabricksClient, error) {
	err := dc.Configure()
	if err != nil {
		return dc, err
	}
	return dc, dc.Authenticate(context.Background())
}

func failsToAuthenticateWith(t *testing.T, dc *DatabricksClient, message string) {
	_, err := configureAndAuthenticate(dc)
	if dc.AuthType != "" {
		log.Printf("[INFO] Auth is: %s", dc.AuthType)
	}
	if assert.NotNil(t, err, "expected to have error: %s", message) {
		assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
	}
}

func TestDatabricksClientConfigure_Nothing(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	failsToAuthenticateWith(t, &DatabricksClient{},
		"authentication is not configured for provider")
}

func TestDatabricksClientConfigure_BasicAuth_NoHost(t *testing.T) {
	defer CleanupEnvironment()()
	failsToAuthenticateWith(t, &DatabricksClient{
		Username: "foo",
		Password: "bar",
	}, "authentication is not configured for provider.")
}

func TestDatabricksClientConfigure_BasicAuth(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:     "https://localhost:443",
		Username: "foo",
		Password: "bar",
	})
	assert.NoError(t, err)
	assert.Equal(t, "basic", dc.AuthType)
}

func TestDatabricksClientConfigure_HostWithoutScheme(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:  "localhost:443",
		Token: "...",
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.AuthType)
	assert.Equal(t, "...", dc.Token)
	assert.Equal(t, "https://localhost:443", dc.Host)
}

func TestDatabricksClientConfigure_Token_NoHost(t *testing.T) {
	defer CleanupEnvironment()()
	failsToAuthenticateWith(t, &DatabricksClient{
		Token: "dapi345678",
	}, "authentication is not configured for provider.")
}

func TestDatabricksClientConfigure_HostTokensTakePrecedence(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:       "foo",
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.AuthType)
}

func TestDatabricksClientConfigure_BasicAuthTakePrecedence(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:       "foo",
		Token:      "configured",
		Username:   "foo",
		Password:   "bar",
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
	assert.Equal(t, "pat", dc.AuthType)
	assert.Equal(t, "configured", dc.Token)
}

func TestDatabricksClientConfigure_ConfigRead(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
	assert.Equal(t, "databricks-cli", dc.AuthType)
	assert.Equal(t, "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ", dc.Token)
}

func TestDatabricksClientConfigure_NoHostGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "nohost",
	}, "cannot configure databricks-cli auth: config file "+
		"testdata/.databrickscfg is corrupt: cannot find host in nohost profile.")
}

func TestDatabricksClientConfigure_NoTokenGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "notoken",
	}, "cannot configure databricks-cli auth: config file "+
		"testdata/.databrickscfg is corrupt: cannot find token in notoken profile.")
}

func TestDatabricksClientConfigure_InvalidProfileGivesError(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "invalidhost",
	}, "cannot configure databricks-cli auth: testdata/.databrickscfg "+
		"has no invalidhost profile configured")
}

func TestDatabricksClientConfigure_MissingFile(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.invalid file",
		Profile:    "invalidhost",
	}, "authentication is not configured for provider.")
}

func TestDatabricksClientConfigure_InvalidConfigFilePath(t *testing.T) {
	failsToAuthenticateWith(t, &DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/az",
		Profile:    "invalidhost",
	}, "cannot configure databricks-cli auth: cannot parse config file")
}

func TestDatabricksClient_FormatURL(t *testing.T) {
	client := DatabricksClient{Host: "https://some.host"}
	assert.Equal(t, "https://some.host/#job/123", client.FormatURL("#job/123"))
}

func TestClientAttributes(t *testing.T) {
	ca := ClientAttributes()
	assert.Len(t, ca, 20)
}

func TestDatabricksClient_Authenticate(t *testing.T) {
	defer CleanupEnvironment()()
	dc := DatabricksClient{}
	err := dc.Configure("account_id", "username", "password")
	os.Setenv("DATABRICKS_PASSWORD", ".")
	assert.NoError(t, err)
	err = dc.Authenticate(context.WithValue(context.Background(), IsData, "yes"))
	assert.EqualError(t, err, "workspace is most likely not created yet, because the `host` is empty. "+
		"Please add `depends_on = [databricks_mws_workspaces.this]` or `depends_on = [azurerm_databricks"+
		"_workspace.this]` to every data resource. See https://www.terraform.io/docs/language/resources/behavior.html more info. "+
		"Attributes used: account_id, username. Environment variables used: DATABRICKS_PASSWORD. "+
		"Please check https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs#authentication for details")
}

func TestDatabricksClient_AuthenticateAzure(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("ARM_CLIENT_SECRET", ".")
	os.Setenv("ARM_CLIENT_ID", ".")
	dc := DatabricksClient{}
	err := dc.Configure("azure_client_id", "azure_client_secret")
	assert.NoError(t, err)
	err = dc.Authenticate(context.WithValue(context.Background(), IsData, "yes"))
	assert.EqualError(t, err, "workspace is most likely not created yet, because the `host` is empty. "+
		"Please add `depends_on = [databricks_mws_workspaces.this]` or `depends_on = [azurerm_databricks"+
		"_workspace.this]` to every data resource. See https://www.terraform.io/docs/language/resources/"+
		"behavior.html more info. Environment variables used: ARM_CLIENT_SECRET, ARM_CLIENT_ID. "+
		"Please check https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs#authentication for details")
}

func TestDatabricksIsGcp(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:  "https://demo.gcp.databricks.com/",
		Token: "dapi123",
	})
	assert.NoError(t, err)
	assert.Equal(t, true, dc.IsGcp())
}

func TestIsAzure_Error(t *testing.T) {
	dc := &DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "notoken",
	}
	assert.Equal(t, false, dc.IsAzure())
}

func TestClientForHost(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:     "https://accounts.cloud.databricks.com/",
		Username: "abc",
		Password: "bcd",
	})
	assert.NoError(t, err)
	assert.True(t, dc.IsAws())
	cc, err := dc.ClientForHost(context.Background(), "https://e2-workspace.cloud.databricks.com/")
	assert.NoError(t, err)
	assert.Equal(t, dc.Username, cc.Username)
	assert.Equal(t, dc.Password, cc.Password)
	assert.NotEqual(t, dc.Host, cc.Host)
}

func TestClientForHostAuthError(t *testing.T) {
	c := &DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "notoken",
	}
	_, err := c.ClientForHost(context.Background(), "https://e2-workspace.cloud.databricks.com/")
	if assert.NotNil(t, err) {
		assert.True(t, strings.HasPrefix(err.Error(),
			"cannot authenticate parent client: cannot configure databricks-cli auth"), err.Error())
	}
}

func TestDatabricksCliCouldNotFindHomeDir(t *testing.T) {
	_, err := (&DatabricksClient{
		ConfigFile: "~.databrickscfg",
	}).configureWithDatabricksCfg(context.Background())
	assert.EqualError(t, err, "cannot find homedir: cannot expand user-specific home dir")
}

func TestDatabricksCliCouldNotParseIni(t *testing.T) {
	_, err := (&DatabricksClient{
		ConfigFile: "testdata/az",
	}).configureWithDatabricksCfg(context.Background())
	if assert.NotNil(t, err) {
		assert.True(t, strings.HasPrefix(err.Error(),
			"cannot parse config file: key-value delimiter not found"), err.Error())
	}
}

func TestDatabricksCliWrongProfile(t *testing.T) {
	_, err := (&DatabricksClient{
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "🤣",
	}).configureWithDatabricksCfg(context.Background())
	assert.EqualError(t, err, "testdata/.databrickscfg has no 🤣 profile configured")
}

func TestDatabricksNoHost(t *testing.T) {
	_, err := (&DatabricksClient{
		ConfigFile: "testdata/corrupt/.databrickscfg",
		Profile:    "nohost",
	}).configureWithDatabricksCfg(context.Background())
	assert.EqualError(t, err, "config file testdata/corrupt/.databrickscfg is corrupt: cannot find host in nohost profile")
}

func TestDatabricksNoToken(t *testing.T) {
	_, err := (&DatabricksClient{
		ConfigFile: "testdata/corrupt/.databrickscfg",
		Profile:    "notoken",
	}).configureWithDatabricksCfg(context.Background())
	assert.EqualError(t, err, "config file testdata/corrupt/.databrickscfg is corrupt: cannot find token in notoken profile")
}

func TestDatabricksBasicAuth(t *testing.T) {
	c := &DatabricksClient{
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "basic",
	}
	_, err := c.configureWithDatabricksCfg(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "abc", c.Username)
	assert.Equal(t, "bcd", c.Password)
}

func TestDatabricksClientConfigure_NonsenseAuth(t *testing.T) {
	defer CleanupEnvironment()()
	failsToAuthenticateWith(t, &DatabricksClient{
		AuthType: "nonsense",
	}, "cannot configure nonsense auth.")
}

func TestConfigAttributeSetNonsense(t *testing.T) {
	err := (&ConfigAttribute{
		Kind: reflect.Chan,
	}).Set(&DatabricksClient{}, 1)
	assert.EqualError(t, err, "cannot set  of unknown type Chan")
}
