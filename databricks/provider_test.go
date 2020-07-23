package databricks

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

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
	testAccProvider = Provider("dev").(*schema.Provider)
	testMWSProvider = Provider("dev").(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"databricks": testAccProvider,
	}
}

func getMWSClient() *service.DatabricksClient {
	client := service.DatabricksClient{
		Host:     os.Getenv("DATABRICKS_MWS_HOST"),
		Username: os.Getenv("DATABRICKS_USERNAME"),
		Password: os.Getenv("DATABRICKS_PASSWORD"),
	}
	err := client.Configure("dev-mws")
	if err != nil {
		panic(err)
	}
	return &client
}

func TestMain(m *testing.M) {
	// This should not be asserted as it may not always be set for all tests
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

	client := testAccProvider.Meta().(*service.DatabricksClient)
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
	client := testAccProvider.Meta().(*service.DatabricksClient)
	assert.Equal(t, expectedToken, client.Token)
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

func TestProvider_DurationToSecondsString(t *testing.T) {
	assert.Equal(t, durationToSecondsString(time.Hour), "3600")
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
