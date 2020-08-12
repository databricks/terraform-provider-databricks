package provider

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// This should not be asserted as it may not always be set for all tests
	// TODO: add common instance pool & cluster for libs & stuff
	cloudEnv := os.Getenv("CLOUD_ENV")
	envFileName := fmt.Sprintf("../.%s.env", cloudEnv)
	err := godotenv.Load(envFileName)
	if !os.IsNotExist(err) {
		log.Printf("[WARN] Failed to load environment: %s", err)
	}
	code := m.Run()
	// epoch.tearDown()
	os.Exit(code)
}

func TestProvider(t *testing.T) {
	if err := DatabricksProvider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_HostTokensTakePrecedence(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["host"] = "foo"
	raw["token"] = "configured"
	raw["config_file"] = "testdata/.databrickscfg"
	p := DatabricksProvider().(*schema.Provider)
	err := p.Configure(terraform.NewResourceConfigRaw(raw))
	assert.Nil(t, err)
	client := p.Meta().(*common.DatabricksClient)
	assert.Equal(t, "configured", client.Token)
}

func TestProvider_BasicAuthTakePrecedence(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["host"] = "foo"
	raw["basic_auth"] = []interface{}{map[string]interface{}{"username": "user", "password": "pass"}}
	raw["config_file"] = "testdata/.databrickscfg"
	p := DatabricksProvider().(*schema.Provider)
	err := p.Configure(terraform.NewResourceConfigRaw(raw))
	assert.Nil(t, err)
	client := p.Meta().(*common.DatabricksClient)
	expectedToken := base64.StdEncoding.EncodeToString([]byte("user:pass"))
	assert.Equal(t, expectedToken, client.Token)
}

func TestProvider_NoHostGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "nohost"
	p := DatabricksProvider().(*schema.Provider)
	err := p.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_NoTokenGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "notoken"
	p := DatabricksProvider().(*schema.Provider)
	err := p.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_InvalidProfileGivesError(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.databrickscfg"
	raw["profile"] = "invalidhost"
	p := DatabricksProvider().(*schema.Provider)
	err := p.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_InvalidConfigFilePath(t *testing.T) {
	var raw = make(map[string]interface{})
	raw["config_file"] = "testdata/.invalid file"
	raw["profile"] = "invalidhost"
	p := DatabricksProvider().(*schema.Provider)
	err := p.Configure(terraform.NewResourceConfigRaw(raw))
	assert.NotNil(t, err)
}

func TestProvider_DurationToSecondsString(t *testing.T) {
	assert.Equal(t, durationToSecondsString(time.Hour), "3600")
}

func TestAccDatabricksCliConfigWorks(t *testing.T) {
	t.Skip("Skipping this test till the better times")
	resource.Test(t,
		resource.TestCase{
			Providers: map[string]terraform.ResourceProvider{
				"databricks": DatabricksProvider(),
			},
			Steps: []resource.TestStep{
				{
					Config:             `provider "databricks" {}`,
					ExpectNonEmptyPlan: true,
				},
			},
		},
	)
}
