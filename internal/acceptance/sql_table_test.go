package acceptance

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/catalog"
)

func TestUcAccResourceSqlTable_Managed(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			properties         = {
				this      = "that"
				something = "else"
			}

			column {
				name      = "id"
				type      = "int"
			}
			column {
				name      = "name"
				type      = "string"
			}
			comment = "this table is managed by terraform"
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			properties         = {
				that      = "this"
				something = "else2"
			}
			
			column {
				name      = "id"
				type      = "int"
			}
			column {
				name      = "name"
				type      = "string"
			}
			comment = "this table is managed by terraform..."
		}`,
	})
}

func TestUcAccResourceSqlTable_External(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}
		
		resource "databricks_external_location" "some" {
			name            = "external-{var.RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/some{var.RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
			force_destroy   = true
		}
				
		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "EXTERNAL"
			data_source_format = "DELTA"
			storage_location   = "s3://{env.TEST_BUCKET}/some{var.RANDOM}"
			comment 		   = "this table is managed by terraform"
		}`,
	})
}

func TestUcAccResourceSqlTable_View(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			data_source_format = "DELTA"
			comment 		   = "this table is managed by terraform..."

			column {
				name      = "id"
				type      = "string"
			}

			column {
				name      = "name"
				type      = "string"
			}
		}
		
		resource "databricks_sql_table" "view" {
			name               = "bar_view"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "VIEW"
			comment 		   = "this view is managed by terraform..."
			view_definition    = format("SELECT id, name FROM %s", databricks_sql_table.this.id)

			column {
				name      = "id"
			}

			column {
				name      = "name"
				comment   = "view column comment"
			}			
		}`,
	})
}

func TestUcAccResourceSqlTable_WarehousePartition(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_sql_endpoint" "this" {
			name = "tf-{var.RANDOM}"
			cluster_size = "2X-Small"
			max_num_clusters = 1
		}

		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			warehouse_id       = databricks_sql_endpoint.this.id
			properties         = {
				them      = "that"
				something = "else"
			}
			options         = {
				this      = "blue"
				that      = "green"
			}			
			column {
				name      = "id"
				type      = "int"
			}
			column {
				name      = "name"
				type      = "string"
			}
			partitions = ["id"]
			comment = "this table is managed by terraform"
		}`,
	})
}
func TestUcAccResourceSqlTable_Liquid(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			properties         = {
				them      = "that"
				something = "else"
			}
			options         = {
				this      = "blue"
				that      = "green"
			}			
			column {
				name      = "id"
				type      = "int"
			}
			column {
				name      = "name"
				type      = "varchar(64)"
			}
			cluster_keys = ["id"]
			comment = "this table is managed by terraform"
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "bar"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			properties         = {
				them      = "that"
				something = "else"
			}
			options         = {
				this      = "blue"
				that      = "green"
			}
			column {
				name      = "id"
				type      = "int"
			}
			column {
				name      = "name"
				type      = "string"
			}
			cluster_keys = ["id", "name"]			
			comment = "this table is managed by terraform..."
		}`,
	})
}

func constructManagedSqlTableTemplate(tableName string, columnInfos []catalog.SqlColumnInfo) string {
	columnsTemplate := ""

	for _, ci := range columnInfos {
		ciTemplate := fmt.Sprintf(
			`
			column {
				name      = "%s"
				type      = "%s"
				nullable  = %t
				comment   = "%s"
			}
			`, ci.Name, ci.Type, ci.Nullable, ci.Comment,
		)
		columnsTemplate += ciTemplate
	}

	return fmt.Sprintf(`
		resource "databricks_schema" "this" {
			name         = "{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_sql_table" "this" {
			name               = "%s"
			catalog_name       = "main"
			schema_name        = databricks_schema.this.name
			table_type         = "MANAGED"
			properties         = {
				this      = "that"
				something = "else"
			}

			%s
			comment = "this table is managed by terraform"
		}`, tableName, columnsTemplate)
}

func TestUcAccResourceSqlTable_RenameColumn(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	unityWorkspaceLevel(t, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "new_name", Type: "string", Nullable: true, Comment: "comment"}}),
	})
}

func TestUcAccResourceSqlTable_AddColumnComment(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	unityWorkspaceLevel(t, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "new comment"}}),
	})
}

func TestUcAccResourceSqlTable_DropColumnNullable(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	unityWorkspaceLevel(t, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: false, Comment: "comment"}}),
	})
}

func TestUcAccResourceSqlTable_MultipleColumnUpdates(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	unityWorkspaceLevel(t, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: false, Comment: "new comment"}}),
	})
}

func TestUcAccResourceSqlTable_ChangeColumnTypeThrows(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	pattern := "changing the 'type' of an existing column is not supported"
	r := regexp.MustCompile(pattern)

	unityWorkspaceLevel(t, step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, step{
		Template:    constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "int", Nullable: true, Comment: "comment"}}),
		ExpectError: r,
	})
}
