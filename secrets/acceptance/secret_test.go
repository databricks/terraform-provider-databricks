package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/secrets"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccSecretResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	config := acceptance.EnvironmentTemplate(t, `
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

	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				PreConfig: func() {
					err := w.Secrets.DeleteSecret(ctx, secrets.DeleteSecret{
						Scope: scope,
						Key:   key,
					})
					assert.NoError(t, err)
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
