package databricks

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/joho/godotenv"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider("").(*schema.Provider)
	cloudEnv := os.Getenv("CLOUD_ENV")

	// If Azure inject sp based auth, this should probably have a different environment variable
	// But for now best practice on azure is to use SP based auth
	if cloudEnv == "azure" {
		var config service.DBApiClientConfig
		testAccProvider.ConfigureFunc = func(data *schema.ResourceData) (i interface{}, e error) {
			return providerConfigureAzureClient(data, "", &config)
		}
	}

	testAccProviders = map[string]terraform.ResourceProvider{
		"databricks": testAccProvider,
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
