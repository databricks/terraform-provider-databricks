package databricks

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/client/service"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testMWSProvider *schema.Provider

func init() {
	testAccProvider = Provider("").(*schema.Provider)
	testMWSProvider = Provider("").(*schema.Provider)
	cloudEnv := os.Getenv("CLOUD_ENV")

	// If Azure inject sp based auth, this should probably have a different environment variable
	// But for now best practice on azure is to use SP based auth
	if cloudEnv == "azure" {
		var config service.DBApiClientConfig
		testAccProvider.ConfigureFunc = func(data *schema.ResourceData) (i interface{}, e error) {
			return providerConfigureAzureClient(data, &config)
		}
	}

	testAccProviders = map[string]terraform.ResourceProvider{
		"databricks": testAccProvider,
	}
}

func getMWSClient() *service.DBApiClient {
	// Configure MWS Provider
	mwsHost := os.Getenv("DATABRICKS_MWS_HOST")
	mwsUser := os.Getenv("DATABRICKS_USERNAME")
	mwsPass := os.Getenv("DATABRICKS_PASSWORD")

	tokenUnB64 := fmt.Sprintf("%s:%s", mwsUser, mwsPass)
	token := base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
	config := service.DBApiClientConfig{
		Host:     mwsHost,
		Token:    token,
		AuthType: service.BasicAuth,
	}
	return &service.DBApiClient{
		Config: &config,
	}
}

func TestMain(m *testing.M) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	envFileName := fmt.Sprintf("../.%s.env", cloudEnv)
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Println("Failed to load environment")
	}
	code := m.Run()
	os.Exit(code)
}

func TestAccProviderConfigureAzureSPAuth(t *testing.T) {
	resource.Test(t,
		resource.TestCase{
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					PlanOnly:           true,
					Config:             testInitialEmptyWorkspaceClusterDeployment(),
					ExpectNonEmptyPlan: true,
				},
			},
		},
	)
}

func testInitialEmptyWorkspaceClusterDeployment() string {
	return `
provider "databricks" {
  azure_auth = {
    managed_resource_group = "azurerm_databricks_workspace.demo.managed_resource_group_name"
    azure_region           = "westus"
    workspace_name         = "azurerm_databricks_workspace.demo.name"
    resource_group         = "azurerm_databricks_workspace.demo.resource_group_name"
  }
}

resource "databricks_scim_group" "my-group-azure3" {
  display_name = "Test terraform Group3"
}
`
}

func TestProvider(t *testing.T) {
	if err := testAccProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_NoOptionsResultsInError(t *testing.T) {
	var provider = Provider("")
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg_non_existent"
	err := provider.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_HostTokensTakePrecedence(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["host"] = "foo"
	raw["token"] = "configured"
	raw["config_file"] = "testdata/.databrickscfg"
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	assert.Nil(t, err)

	client := testAccProvider.Meta().(*service.DBApiClient).Config
	assert.Equal(t, "configured", client.Token)
}

func TestProvider_BasicAuthTakePrecedence(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["host"] = "foo"
	raw["basic_auth"] = []interface{}{map[string]interface{}{"username": "user", "password": "pass"}}
	raw["config_file"] = "testdata/.databrickscfg"
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	assert.Nil(t, err)

	// Basic auth convention
	expectedToken := base64.StdEncoding.EncodeToString([]byte("user:pass"))
	client := testAccProvider.Meta().(*service.DBApiClient).Config
	assert.Equal(t, expectedToken, client.Token)
}

func TestProvider_MissingEnvMakesConfigRead(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["token"] = "configured"
	raw["config_file"] = "testdata/.databrickscfg"
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	assert.Nil(t, err)

	client := testAccProvider.Meta().(*service.DBApiClient).Config
	assert.Equal(t, "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ", client.Token)
}

func TestProvider_NoHostGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "nohost"
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_NoTokenGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "notoken"
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_InvalidProfileGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "invalidhost"
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_InvalidConfigFilePath(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.invalid file"
	raw["profile"] = "invalidhost"
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(raw))
	log.Println(err)
	assert.NotNil(t, err)
}

func TestAccDatabricksCliConfigWorks(t *testing.T) {
	resource.Test(t,
		resource.TestCase{
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:             `provider "databricks" {}`,
					ExpectNonEmptyPlan: true,
				},
			},
		},
	)
}
