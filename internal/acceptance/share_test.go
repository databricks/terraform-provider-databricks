package acceptance

import (
	"fmt"
	"testing"
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
		name = "abc"
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
		name = "def"
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
		name = "xyz"
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

	resource "databricks_table" "mytable_4" {
		catalog_name = databricks_catalog.sandbox.id
		schema_name = databricks_schema.things.name
		name = "pqr"
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
	unityWorkspaceLevel(t, step{
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
			}								
		}`, owner, comment)
}

func TestUcAccShare_MultipleObjects(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: preTestTemplate + preTestTemplateUpdate + `		
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable_2.id
				comment = "def"
				data_object_type = "TABLE"
			}					
		}		
		`,
	}, step{
		Template: preTestTemplate + preTestTemplateUpdate + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "abc"
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable_2.id
				comment = "def - updated"
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable_3.id
				comment = "xyz"
				data_object_type = "TABLE"
			}
		}`,
	}, step{
		Template: preTestTemplate + preTestTemplateUpdate + `
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable_2.id
				comment = "def - updated 2"
				data_object_type = "TABLE"
			}
			object {
				name = databricks_table.mytable_4.id
				comment = "xyz - updated"
				data_object_type = "TABLE"
			}
		}`,
	})
}

func TestUcAccUpdateShare(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("c", "account users"),
	}, step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("e", "account users"),
	}, step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("e", "{env.TEST_DATA_ENG_GROUP}"),
	}, step{
		Template: preTestTemplate + preTestTemplateUpdate + shareTemplateWithOwner("f", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}
