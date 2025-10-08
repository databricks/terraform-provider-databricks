package sharing_test

import (
	"context"
	"fmt"
	"maps"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
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

			// add object to share outside terraform
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
	})
}
