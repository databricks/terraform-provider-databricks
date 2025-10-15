package sharing_test

import (
	"context"
	"fmt"
<<<<<<< HEAD:internal/providers/pluginfw/products/sharing/resource_share_acc_test.go
	"maps"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sharing"
=======
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
>>>>>>> unified-tf-plugin:internal/providers/pluginfw/products/sharing/resource_acc_test.go
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
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
		resource "databricks_share_pluginframework" "myshare" {
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
			share = databricks_share_pluginframework.myshare.name
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
		resource "databricks_share_pluginframework" "myshare" {
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
			`resource "databricks_share_pluginframework" "myshare" {
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
			`resource "databricks_share_pluginframework" "myshare" {
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
			provider_config = {
				workspace_id = "{env.THIS_WORKSPACE_ID}"
			}
		}`,
	})
}

func TestUcAccUpdateShareReorderObject(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplate + preTestTemplateUpdate +
			`resource "databricks_share_pluginframework" "myshare" {
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
			`resource "databricks_share_pluginframework" "myshare" {
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

func shareUpdateWithName(name string) string {
	return fmt.Sprintf(`resource "databricks_share_pluginframework" "myshare" {
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
		resource "databricks_share_pluginframework" "myshare" {
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
		resource "databricks_share_pluginframework" "myshare" {
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
		resource "databricks_share_pluginframework" "myshare" {
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

<<<<<<< HEAD:internal/providers/pluginfw/products/sharing/resource_share_acc_test.go
func TestUcAccUpdateShareOutsideTerraform(t *testing.T) {
	shareName := ""
	sharedObjectNameToAdd := ""
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + `
		resource "databricks_share_pluginframework" "myshare" {
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
			share := resources["databricks_share_pluginframework.myshare"]
			if share == nil {
				return fmt.Errorf("expected to find databricks_share_pluginframework.myshare in resources keys: %v", maps.Keys(resources))
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
		resource "databricks_share_pluginframework" "myshare" {
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
=======
func shareTemplate(provider_config string) string {
	return fmt.Sprintf(`
	resource "databricks_share_pluginframework" "myshare" {
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
			provider_config = {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`(?s)failed to get workspace client.*failed to parse workspace_id.*valid integer`),
	})
}

func TestAccJobCluster_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config = {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`(?s)failed to get workspace client.*workspace_id mismatch.*please check the workspace_id provided in provider_config`),
	})
}

func TestAccJobCluster_ProviderConfig_Required(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config = {
			}
		`),
		ExpectError: regexp.MustCompile(`(?s).*workspace_id.*is required`),
	})
}

func TestAccJobCluster_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config = {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`Attribute provider_config\.workspace_id string length must be at least 1`),
	})
}

func TestAccJobCluster_ProviderConfig_NotProvided(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	})
}

func TestAccJobCluster_ProviderConfig_Match(t *testing.T) {
	// acceptance.LoadWorkspaceEnv(t)
	// get workspace id here from workspace
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config = {
				workspace_id = "4220866301720038"
			}
		`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				common.CheckResourceUpdate{Address: "databricks_share_pluginframework.myshare"},
				common.CheckResourceNoDelete{Address: "databricks_share_pluginframework.myshare"},
				common.CheckResourceNoCreate{Address: "databricks_share_pluginframework.myshare"},
			},
		},
	})
}

func TestAccJobCluster_ProviderConfig_Recreate(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config = {
				workspace_id = "4220866301720038"
			}
		`),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config = {
				workspace_id = "123"
			}
		`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				common.CheckResourceCreate{Address: "databricks_share_pluginframework.myshare"},
				common.CheckResourceDelete{Address: "databricks_share_pluginframework.myshare"},
			},
		},
		ExpectError: regexp.MustCompile(`failed to validate workspace_id: workspace_id mismatch`),
	})
}

func TestAccJobCluster_ProviderConfig_Remove(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(`
			provider_config = {
				workspace_id = "4220866301720038"
			}
		`),
	}, acceptance.Step{
		Template: preTestTemplateSchema + shareTemplate(""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				common.CheckResourceUpdate{Address: "databricks_share_pluginframework.myshare"},
				common.CheckResourceNoDelete{Address: "databricks_share_pluginframework.myshare"},
				common.CheckResourceNoCreate{Address: "databricks_share_pluginframework.myshare"},
			},
		},
>>>>>>> unified-tf-plugin:internal/providers/pluginfw/products/sharing/resource_acc_test.go
	})
}
