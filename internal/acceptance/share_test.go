package acceptance

import (
	"testing"
)

func TestUcAccCreateShare(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
		}
		
		resource "databricks_schema" "things" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things{var.RANDOM}"
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
		
		resource "databricks_share" "myshare" {
			name  = "{var.RANDOM}-terraform-delta-share"
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
			name = "{var.RANDOM}-terraform-db2open-recipient"
			comment = "made by terraform"
			authentication_type = "TOKEN"
			sharing_code = "{var.RANDOM}"
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

func TestUcAccCreateUpdate(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
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
		
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "c"
				data_object_type = "TABLE"
			}							
		}`,
	}, step{
		Template: `
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
		
		resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-terraform-delta-share"
			owner = "account users"
			object {
				name = databricks_table.mytable.id
				comment = "d"
				data_object_type = "TABLE"
			}								
		}`,
	})
}
