package acceptance

import (
	"context"
	"os"
	"testing"

	. "github.com/databrickslabs/terraform-provider-databricks/access"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
	resource "databricks_secret_scope" "this" {
		name = "tf-scope-{var.RANDOM}"
	}
	resource "databricks_secret" "this" {
		scope = databricks_secret_scope.this.name
		string_value = "{var.RANDOM}"
		key = "password"
	}`)
	scope := qa.FirstKeyValue(t, config, "name")
	key := qa.FirstKeyValue(t, config, "key")
	secret := qa.FirstKeyValue(t, config, "string_value")

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewSecretsAPI(context.Background(), client).Delete(scope, key)
					assert.NoError(t, err, err)
				},
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_secret.this", "scope", scope),
					resource.TestCheckResourceAttr("databricks_secret.this", "key", key),
					resource.TestCheckResourceAttr("databricks_secret.this", "string_value", secret),
				),
			},
		},
	})
}
