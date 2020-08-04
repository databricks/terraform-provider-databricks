package databricks

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/client/service"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider("dev").(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"databricks": testAccProvider,
	}
}

var (
	once  sync.Once
	epoch testEpoch
)

type testEpoch struct {
}

func (e *testEpoch) RandomShortName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
}

func (e *testEpoch) RandomLongName() string {
	return "Terraform Integration Test " + e.RandomShortName()
}

func (e *testEpoch) ResourceCheck(name string,
	cb func(client *service.DatabricksClient, id string) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		client := testAccProvider.Meta().(*service.DatabricksClient)
		return cb(client, rs.Primary.ID)
	}
}

func TestMain(m *testing.M) {
	// This should not be asserted as it may not always be set for all tests
	// TODO: add common instance pool & cluster for libs & stuff
	cloudEnv := os.Getenv("CLOUD_ENV")
	envFileName := fmt.Sprintf("../.%s.env", cloudEnv)
	err := godotenv.Load(envFileName)
	if !os.IsNotExist(err) {
		log.Printf("[WARN] Failed to load environment: %s", err)
	}
	once.Do(func() { // atomic
		log.Printf("[INFO] Initializing test epoch")
		epoch = testEpoch{} // thread safe
	})
	code := m.Run()

	// TODO: make a teardown
	// epoch.tearDown()
	os.Exit(code)
}

func TestProvider(t *testing.T) {
	if err := testAccProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
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
	t.Skip("Skipping this test till the better times")
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
