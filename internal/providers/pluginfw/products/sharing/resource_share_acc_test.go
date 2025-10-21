package sharing_test

import (
	"context"
	"fmt"
	"maps"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const preTestTemplate = `
	resource "databricks_catalog" "sandbox" {
		name         = "sandbox{var.STICKY_RANDOM}"
		comment      = "this catalog is managed by terraform"
		properties = {
			purpose = "testing"
		}
	}

	resource "databricks_schema" "things" {
		catalog_name = databricks_catalog.sandbox.id
		name         = "things{var.STICKY_RANDOM}"
		comment      = "this database is managed by terraform"
		properties = {
			kind = "various"
		}
	}

	resource "databricks_sql_table" "mytable" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar"
		table_type = "MANAGED"
		warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"

		column {
			name = "id"
			type = "int"
		}
	}

	resource "databricks_sql_table" "mytable_2" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar_2"
		table_type = "MANAGED"
		warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"

		column {
			name = "id"
			type = "int"
		}
	}

	resource "databricks_sql_table" "mytable_3" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar_3"
		table_type = "MANAGED"
		warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"

		column {
			name = "id"
			type = "int"
		}
	}
`

const preTestTemplateUpdate = `
	resource "databricks_grants" "some" {
		catalog = databricks_catalog.sandbox.id
		grant {
			principal  = "account users"
			privileges = ["ALL_PRIVILEGES"]
		}
		grant {
			principal  = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
			privileges = ["ALL_PRIVILEGES"]
		}
	}
`

func TestUcAccCreateShare(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_sql_table.mytable.id
				comment = "c"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
         	}
			object {
				name = databricks_sql_table.mytable_2.id
				comment = "c"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
		}

		resource "databricks_recipient" "db2open" {
			name = "{var.STICKY_RANDOM}-terraform-db2open-recipient"
			comment = "made by terraform"
			authentication_type = "TOKEN"
			sharing_code = "{var.STICKY_RANDOM}"
			ip_access_list {
			// using private ip for acc testing
			allowed_ip_addresses = ["10.0.0.0/16"]
			}
		}

		resource "databricks_grants" "some" {
			share = databricks_share.myshare.name
			grant {
				principal  = databricks_recipient.db2open.name
				privileges = ["SELECT"]
			}
		}
		`,
	})
}

func shareTemplateWithOwner(comment string, owner string) string {
	return fmt.Sprintf(`
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "%s"
			object {
				name = databricks_sql_table.mytable.id
				comment = "%s"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}

		}`, owner, comment)
}

func TestUcAccUpdateShare(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("c", "account users"),
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("e", "account users"),
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("e", "{env.TEST_DATA_ENG_GROUP}"),
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("f", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}

func TestUcAccUpdateShareAddObject(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_sql_table.mytable.id
				comment = "A"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
			object {
				name = databricks_sql_table.mytable_3.id
				comment = "C"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}

		}`,
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_sql_table.mytable.id
				comment = "AA"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
			object {
				name = databricks_sql_table.mytable_2.id
				comment = "BB"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
			object {
				name = databricks_sql_table.mytable_3.id
				comment = "CC"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
		}`,
	})
}

func TestUcAccUpdateShareReorderObject(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_sql_table.mytable.id
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
			object {
				name = databricks_sql_table.mytable_3.id
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
		}`,
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_sql_table.mytable_3.id
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
			object {
				name = databricks_sql_table.mytable.id
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
		}`,
	})
}

// TestUcAccUpdateShareNoChanges tests that updating a share with no actual changes doesn't cause issues
func TestUcAccUpdateShareNoChanges(t *testing.T) {
	shareConfig := preTestTemplateSchema +
		`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
		}`

	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: shareConfig,
	}, acceptance.Step{
		PlanOnly: true,
		Template: shareConfig, // Same config - should not trigger any updates
	})
}

// TestUcAccUpdateShareComplexObjectChanges tests complex scenarios with multiple object updates
func TestUcAccUpdateShareComplexObjectChanges(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			object {
				name = databricks_schema.schema1.id
				comment = "original comment"
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema2.id
				comment = "second schema"
				data_object_type = "SCHEMA"
			}
		}`,
	}, acceptance.Step{
		// Remove one object, add another, and update comment on existing
		Template: preTestTemplateSchema +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			object {
				name = databricks_schema.schema1.id
				comment = "updated comment"
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema3.id
				comment = "third schema"
				data_object_type = "SCHEMA"
			}
		}`,
	})
}

// TestUcAccUpdateShareRemoveAllObjects tests removing all objects from a share
func TestUcAccUpdateShareRemoveAllObjects(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_schema.schema1.id
				comment = "to be removed"
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema2.id
				comment = "also to be removed"
				data_object_type = "SCHEMA"
			}
		}`,
	}, acceptance.Step{
		Template: preTestTemplateSchema +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
		}`,
	})
}

// TestUcAccShareMigrationFromSDKv2 tests the transition from sdkv2 to plugin framework.
// This test verifies that existing state created by SDK v2 implementation can be
// successfully managed by the plugin framework implementation without any changes.
func TestUcAccShareMigrationFromSDKv2(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t,
		// Step 1: Create share using SDK v2 implementation
		acceptance.Step{
			ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
				"databricks": func() (tfprotov6.ProviderServer, error) {
					sdkv2Provider, pluginfwProvider := acceptance.ProvidersWithResourceFallbacks([]string{"databricks_share"})
					return providers.GetProviderServer(context.Background(), providers.WithSdkV2Provider(sdkv2Provider), providers.WithPluginFrameworkProvider(pluginfwProvider))
				},
			},
			Template: preTestTemplateSchema + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share"
					object {
						name = databricks_schema.schema1.id
						comment = "Shared schema object for migration test"
						data_object_type = "SCHEMA"
					}
					object {
						name = databricks_schema.schema2.id
						comment = "Second shared schema object"
						data_object_type = "SCHEMA"
					}
				}`,
		},
		// Step 2: Update the share using plugin framework implementation (default)
		// This verifies no changes are needed when switching implementations
		acceptance.Step{
			Template: preTestTemplateSchema + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share"
					object {
						name = databricks_schema.schema1.id
						comment = "Updated comment for schema object after migration"
						data_object_type = "SCHEMA"
					}
					object {
						name = databricks_schema.schema2.id
						comment = "Second shared schema object"
						data_object_type = "SCHEMA"
					}
				}`,
		},
	)
}

// TestUcAccShareMigrationFromPluginFramework tests the transition from plugin framework to sdkv2.
// This test verifies that existing state created by plugin framework implementation can be
// successfully managed by the SDK v2 implementation without any changes.
func TestUcAccShareMigrationFromPluginFramework(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t,
		// Step 1: Create share using plugin framework implementation
		acceptance.Step{
			Template: preTestTemplateSchema + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share-rollback"
					owner = "account users"
					object {
						name = databricks_schema.schema1.id
						comment = "Shared schema object for migration test"
						data_object_type = "SCHEMA"
					}
					object {
						name = databricks_schema.schema2.id
						comment = "Second shared schema object"
						data_object_type = "SCHEMA"
					}
				}`,
		},
		// Step 2: Update the share using SDK v2 (default)
		// This verifies no changes are needed when switching implementations
		acceptance.Step{
			ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
				"databricks": func() (tfprotov6.ProviderServer, error) {
					sdkv2Provider, pluginfwProvider := acceptance.ProvidersWithResourceFallbacks([]string{"databricks_share"})
					return providers.GetProviderServer(context.Background(), providers.WithSdkV2Provider(sdkv2Provider), providers.WithPluginFrameworkProvider(pluginfwProvider))
				},
			},
			Template: preTestTemplateSchema + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share-rollback"
					owner = "account users"
					object {
						name = databricks_schema.schema1.id
						comment = "Shared schema object for migration test"
						data_object_type = "SCHEMA"
					}
					object {
						name = databricks_schema.schema2.id
						comment = "Second shared schema object"
						data_object_type = "SCHEMA"
					}
				}`,
		},
	)
}
func shareUpdateWithName(name string) string {
	return fmt.Sprintf(`resource "databricks_share" "myshare" {
			name  = "%s"
			owner = "account users"
			object {
				name = databricks_sql_table.mytable.id
				comment = "A"
				data_object_type = "TABLE"
				history_data_sharing_status = "ENABLED"
			}
		}`, name)
}

func shareCheckStateforID() func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources["databricks_share.myshare"]
		if !ok {
			return fmt.Errorf("resource not found in state")
		}
		id := r.Primary.Attributes["id"]
		name := r.Primary.Attributes["name"]
		if id != name {
			return fmt.Errorf("resource ID is not equal to the name. Attributes: %v", r.Primary.Attributes)
		}
		return nil
	}
}

func TestUcAccUpdateShareName(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + shareUpdateWithName("{var.STICKY_RANDOM}-terraform-delta-share-before"),
		Check:    shareCheckStateforID(),
	}, acceptance.Step{
		Template: preTestTemplate + shareUpdateWithName("{var.STICKY_RANDOM}-terraform-delta-share-after"),
		Check:    shareCheckStateforID(),
	})
}

const preTestTemplateSchema = `
	resource "databricks_catalog" "sandbox" {
		name         = "sandbox{var.STICKY_RANDOM}"
		comment      = "this catalog is managed by terraform"
		properties = {
			purpose = "testing"
		}
	}
	resource "databricks_schema" "schema1" {
		catalog_name = databricks_catalog.sandbox.id
		name         = "schema1{var.STICKY_RANDOM}"
		comment      = "this database is managed by terraform"
		properties = {
			kind = "various"
		}
	}
	resource "databricks_schema" "schema2" {
		catalog_name = databricks_catalog.sandbox.id
		name         = "schema2{var.STICKY_RANDOM}"
		comment      = "this database is managed by terraform"
		properties = {
			kind = "various"
		}
	}
	resource "databricks_schema" "schema3" {
		catalog_name = databricks_catalog.sandbox.id
		name         = "schema3{var.STICKY_RANDOM}"
		comment      = "this database is managed by terraform"
		properties = {
			kind = "various"
		}
	}
`

func TestUcAccShareReorderObject(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share-reorder-terraform"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema3.id
				data_object_type = "SCHEMA"
			}
		}`,
	}, acceptance.Step{
		Template: preTestTemplateSchema + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share-reorder-terraform"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema3.id
				data_object_type = "SCHEMA"
			}
		}`,
		PlanOnly: true,
	}, acceptance.Step{
		// Changing order of objects in the config leads to changes show up in plan as updates
		Template: preTestTemplateSchema + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share-reorder-terraform"
			object {
				name = databricks_schema.schema3.id
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}

		}`,
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
	})
}

func TestUcAccUpdateShareOutsideTerraform(t *testing.T) {
	shareName := ""
	sharedObjectNameToAdd := ""
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share-outside-terraform"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema3.id
				data_object_type = "SCHEMA"
			}
		}`,
		Check: func(s *terraform.State) error {
			resources := s.RootModule().Resources
			share := resources["databricks_share.myshare"]
			if share == nil {
				return fmt.Errorf("expected to find databricks_share.myshare in resources keys: %v", maps.Keys(resources))
			}
			shareName = share.Primary.Attributes["name"]
			assert.NotEmpty(t, shareName)

			schema := resources["databricks_schema.schema2"]
			if schema == nil {
				return fmt.Errorf("expected to find databricks_schema.schema2 in resources keys: %v", maps.Keys(resources))
			}
			sharedObjectNameToAdd = schema.Primary.Attributes["id"]
			assert.NotEmpty(t, sharedObjectNameToAdd)
			return nil
		},
	}, acceptance.Step{
		PreConfig: func() {
			w, err := databricks.NewWorkspaceClient(&databricks.Config{})
			require.NoError(t, err)

			// Add object to share outside terraform
			_, err = w.Shares.Update(context.Background(), sharing.UpdateShare{
				Name: shareName,
				Updates: []sharing.SharedDataObjectUpdate{
					{
						Action: sharing.SharedDataObjectUpdateActionAdd,
						DataObject: &sharing.SharedDataObject{
							Name:           sharedObjectNameToAdd,
							DataObjectType: "SCHEMA",
						},
					},
				},
			})
			require.NoError(t, err)
		},
		Template: preTestTemplateSchema + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share-outside-terraform"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
			object {
				name = databricks_schema.schema3.id
				data_object_type = "SCHEMA"
			}
		}`,
	})
}

func shareTemplate(provider_config string) string {
	return fmt.Sprintf(`
	resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-share-config"
			%s
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
	}
`, provider_config)
}

func TestAccShare_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(
			`(?s)Attribute provider_config\[0\]\.workspace_id ` +
				`workspace_id must be a valid.*integer, got: invalid`,
		),
		PlanOnly: true,
	})
}

func TestAccShare_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(
			`(?s)failed to get workspace client.*workspace_id mismatch` +
				`.*please check the workspace_id provided in ` +
				`provider_config`,
		),
	})
}

func TestAccShare_ProviderConfig_Multiple(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config {
				workspace_id = "123"
			}
			provider_config {
				workspace_id = "456"
			}
		`),
		ExpectError: regexp.MustCompile(
			`Attribute provider_config list must contain at most 1 element`,
		),
		PlanOnly: true,
	})
}

func TestAccShare_ProviderConfig_Required(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`(?s).*workspace_id.*is required`),
	})
}

func TestAccShare_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`Attribute provider_config\[0\]\.workspace_id string length must be at least 1`),
		PlanOnly:    true,
	})
}

func TestAccShare_ProviderConfig_NotProvided(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	})
}

func TestAccShare_ProviderConfig_Match(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_share.myshare", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccShare_ProviderConfig_Recreate(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_share.myshare", plancheck.ResourceActionDestroyBeforeCreate),
			},
		},
		ExpectError: regexp.MustCompile(`failed to validate workspace_id: workspace_id mismatch`),
	})
}

func TestAccShare_ProviderConfig_Remove(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_share.myshare", plancheck.ResourceActionUpdate),
			},
		},
	})
}
