package common

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertErrorStartsWith(t *testing.T, err error, message string) bool {
	return assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
}

func configureAndAuthenticate(dc *DatabricksClient) (*DatabricksClient, error) {
	err := dc.Configure()
	if err != nil {
		return dc, err
	}
	return dc, dc.Authenticate(context.Background())
}

func TestDatabricksClientConfigure_Nothing(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	_, err := configureAndAuthenticate(&DatabricksClient{})
	AssertErrorStartsWith(t, err, "authentication is not configured for provider")
}

func TestDatabricksClientConfigure_BasicAuth_NoHost(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Username: "foo",
		Password: "bar",
	})

	AssertErrorStartsWith(t, err, "cannot configure direct auth: host is empty, but is required by basic_auth")
	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
}

func TestDatabricksClientConfigure_BasicAuth(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:     "https://localhost:443",
		Username: "foo",
		Password: "bar",
	})

	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_HostWithoutScheme(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:  "localhost:443",
		Token: "...",
	})

	assert.Equal(t, "...", dc.Token)
	assert.Equal(t, "https://localhost:443", dc.Host)
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_Token_NoHost(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Token: "dapi345678",
	})

	AssertErrorStartsWith(t, err, "cannot configure direct auth: host is empty, but is required by token")
	assert.Equal(t, "dapi345678", dc.Token)
}

func TestDatabricksClientConfigure_HostTokensTakePrecedence(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Host:       "foo",
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_BasicAuthTakePrecedence(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:       "foo",
		Token:      "connfigured",
		Username:   "foo",
		Password:   "bar",
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
}

func TestDatabricksClientConfigure_ConfigRead(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
	assert.Equal(t, "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ", dc.Token)
}

func TestDatabricksClientConfigure_NoHostGivesError(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "nohost",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_NoTokenGivesError(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "notoken",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_InvalidProfileGivesError(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "invalidhost",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_MissingFile(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.invalid file",
		Profile:    "invalidhost",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_InvalidConfigFilePath(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/policy01.json",
		Profile:    "invalidhost",
	})
	assert.Error(t, err)
}

func TestDatabricksClient_FormatURL(t *testing.T) {
	client := DatabricksClient{Host: "https://some.host"}
	assert.Equal(t, "https://some.host/#job/123", client.FormatURL("#job/123"))
}

func TestClientAttributes(t *testing.T) {
	ca := ClientAttributes()
	assert.Len(t, ca, 23)
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
