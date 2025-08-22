package sharing_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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

	resource "databricks_table" "mytable" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar"
		table_type = "MANAGED"
		data_source_format = "DELTA"

		column {
			name      = "id"
			position  = 0
			type_name = "INT"
			type_text = "int"
			type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
		}
	}

	resource "databricks_table" "mytable_2" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar_2"
		table_type = "MANAGED"
		data_source_format = "DELTA"

		column {
			name      = "id"
			position  = 0
			type_name = "INT"
			type_text = "int"
			type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
		}
	}

	resource "databricks_table" "mytable_3" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "bar_3"
		table_type = "MANAGED"
		data_source_format = "DELTA"

		column {
			name      = "id"
			position  = 0
			type_name = "INT"
			type_text = "int"
			type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
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
				name = databricks_table.mytable.id
				comment = "c"
				data_object_type = "TABLE"
         	}
			object {
				name = databricks_table.mytable_2.id
				cdf_enabled = false
				comment = "c"
				data_object_type = "TABLE"
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
				name = databricks_table.mytable.id
				comment = "%s"
				data_object_type = "TABLE"
				history_data_sharing_status = "DISABLED"
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
				name = databricks_table.mytable.id
				comment = "A"
				data_object_type = "TABLE"
				history_data_sharing_status = "DISABLED"
			}
			object {
				name = databricks_table.mytable_3.id
				comment = "C"
				data_object_type = "TABLE"
				history_data_sharing_status = "DISABLED"
			}

		}`,
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "AA"
				data_object_type = "TABLE"
				history_data_sharing_status = "DISABLED"
			}
			object {
				name = databricks_table.mytable_2.id
				comment = "BB"
				data_object_type = "TABLE"
				history_data_sharing_status = "DISABLED"
			}
			object {
				name = databricks_table.mytable_3.id
				comment = "CC"
				data_object_type = "TABLE"
				history_data_sharing_status = "DISABLED"
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
				name = databricks_table.mytable.id
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable_3.id
				data_object_type = "TABLE"
			}
		}`,
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable_3.id
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable.id
				data_object_type = "TABLE"
			}
		}`,
	})
}

// TestUcAccUpdateShareNoChanges tests that updating a share with no actual changes doesn't cause issues
func TestUcAccUpdateShareNoChanges(t *testing.T) {
	shareConfig := preTestTemplate + preTestTemplateUpdate +
		`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "stable comment"
				data_object_type = "TABLE"
			}
		}`

	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: shareConfig,
	}, acceptance.Step{
		Template: shareConfig, // Same config - should not trigger any updates
	})
}

// TestUcAccUpdateShareComplexObjectChanges tests complex scenarios with multiple object updates
func TestUcAccUpdateShareComplexObjectChanges(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "original comment"
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable_2.id
				comment = "second table"
				data_object_type = "TABLE"
			}
		}`,
	}, acceptance.Step{
		// Remove one object, add another, and update comment on existing
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "updated comment"
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable_3.id
				comment = "third table"
				data_object_type = "TABLE"
			}
		}`,
	})
}

// TestUcAccUpdateShareRemoveAllObjects tests removing all objects from a share
func TestUcAccUpdateShareRemoveAllObjects(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "to be removed"
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable_2.id
				comment = "also to be removed"
				data_object_type = "TABLE"
			}
		}`,
	}, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
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
			Template: preTestTemplate + preTestTemplateUpdate + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share"
					owner = "account users"
					object {
						name = databricks_table.mytable.id
						comment = "Shared table for migration test"
						data_object_type = "TABLE"
					}
					object {
						name = databricks_table.mytable_2.id
						comment = "Second shared table"
						data_object_type = "TABLE"
						cdf_enabled = false
					}
				}`,
		},
		// Step 2: Update the share using plugin framework implementation (default)
		// This verifies no changes are needed when switching implementations
		acceptance.Step{
			ExpectNonEmptyPlan: false,
			Template: preTestTemplate + preTestTemplateUpdate + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share"
					owner = "account users"
					object {
						name = databricks_table.mytable.id
						comment = "Updated comment after migration"
						data_object_type = "TABLE"
					}
					object {
						name = databricks_table.mytable_2.id
						comment = "Second shared table"
						data_object_type = "TABLE"
						cdf_enabled = false
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
			Template: preTestTemplate + preTestTemplateUpdate + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share-rollback"
					owner = "account users"
					object {
						name = databricks_table.mytable.id
						comment = "Shared table for migration test"
						data_object_type = "TABLE"
					}
					object {
						name = databricks_table.mytable_2.id
						comment = "Second shared table"
						data_object_type = "TABLE"
						cdf_enabled = false
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
			ExpectNonEmptyPlan: false,
			Template: preTestTemplate + preTestTemplateUpdate + `
				resource "databricks_share" "myshare" {
					name  = "{var.STICKY_RANDOM}-terraform-migration-share-rollback"
					owner = "account users"
					object {
						name = databricks_table.mytable.id
						comment = "Shared table for migration test"
						data_object_type = "TABLE"
					}
					object {
						name = databricks_table.mytable_2.id
						comment = "Second shared table"
						data_object_type = "TABLE"
						cdf_enabled = false
					}
				}`,
		},
	)
}
func shareUpdateWithName(name string) string {
	return fmt.Sprintf(`resource "databricks_share_pluginframework" "myshare" {
			name  = "%s"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "A"
				data_object_type = "TABLE"
				history_data_sharing_status = "DISABLED"
			}
		}`, name)
}

func shareCheckStateforID() func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources["databricks_share_pluginframework.myshare"]
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
