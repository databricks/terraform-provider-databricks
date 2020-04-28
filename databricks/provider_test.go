package databricks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider("").(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"databricks": testAccProvider,
	}
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Failed to load environment")
	}
	code := m.Run()
	os.Exit(code)
}
