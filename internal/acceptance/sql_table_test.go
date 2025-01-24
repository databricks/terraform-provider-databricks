package acceptance

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/catalog"
)

func TestUcAccResourceSqlTable_Managed(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	UnityWorkspaceLevel(t, Step{
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
			owner = "account users"
		}`,
	}, Step{
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

func TestUcAccResourceSqlTableWithIdentityColumn_Managed(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	UnityWorkspaceLevel(t, Step{
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
				type      = "bigint"
				identity  = "default"
			}
			column {
				name      = "name"
				type      = "string"
			}
			comment = "this table is managed by terraform"
			owner = "account users"
		}`,
	}, Step{
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
				type      = "bigint"
				identity  = "default"
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
	UnityWorkspaceLevel(t, Step{
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
			owner              = "account users"
		}`,
	})
}

func TestUcAccResourceSqlTable_View(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	UnityWorkspaceLevel(t, Step{
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
			owner              = "account users"

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
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	UnityWorkspaceLevel(t, Step{
		Template: `
		resource "databricks_sql_endpoint" "this" {
			name = "tf-{var.RANDOM}"
			cluster_size = "2X-Small"
			max_num_clusters = 1

			tags {
				custom_tags {
					key   = "Owner"
					value = "eng-dev-ecosystem-team@databricks.com"
				}
			}
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
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	UnityWorkspaceLevel(t, Step{
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
	}, Step{
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
			cluster_keys = ["id", "name"]			
			comment = "this table is managed by terraform..."
		}`,
	})
}

func constructManagedSqlTableTemplate(tableName string, columnInfos []catalog.SqlColumnInfo) string {
	columnsTemplate := catalog.GetSqlColumnInfoHCL(columnInfos)

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
				"this"                        = "that"
				"something"                   = "else"
				"delta.feature.allowColumnDefaults = "supported"
				"delta.feature.columnMapping" = "supported"
				"delta.feature.invariants"    = "supported"
				"delta.minReaderVersion"      = 3
				"delta.minWriterVersion"      = 7
				"delta.columnMapping.mode"    = "name"
			}

			%s
			comment = "this table is managed by terraform"
		}`, tableName, columnsTemplate)
}

var typeUpdateErrorPattern = "changing the 'type' of an existing column is not supported"
var typeUpdateErrorRegex = regexp.MustCompile(typeUpdateErrorPattern)

var inlineAndMembershipChangeErrorPattern = "detected changes in both number of columns and existing column field values, please do not change number of columns and update column values at the same time"
var inlineAndMembershipChangeErrorRegex = regexp.MustCompile(inlineAndMembershipChangeErrorPattern)

func TestUcAccResourceSqlTable_RenameColumn(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "new_name", Type: "string", Nullable: true, Comment: "comment"}}),
	})
}

func constructManagedSqlTableTemplateWithColumnTypeUpdates(tableName string, columnName string, step string, columnTypes []string) string {
	colInfos := []catalog.SqlColumnInfo{}
	for index, colType := range columnTypes {
		colInfos = append(colInfos, catalog.SqlColumnInfo{
			Name:     columnName + strconv.Itoa(index),
			Type:     colType,
			Nullable: true,
			Comment:  "comment" + strconv.Itoa(index) + step,
		})
	}
	return constructManagedSqlTableTemplate(tableName, colInfos)
}

func TestUcAccResourceSqlTable_ColumnTypeSuppressDiff(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	columnName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplateWithColumnTypeUpdates(tableName, columnName, "0", []string{
			"integer",
			"long",
			"real",
			"short",
			"byte",
			"decimal",
			"dec",
			"numeric",
		}),
	}, Step{
		Template: constructManagedSqlTableTemplateWithColumnTypeUpdates(tableName, columnName, "1", []string{
			"INTEGER",
			"LONG",
			"REAL",
			"SHORT",
			"BYTE",
			"DECIMAL",
			"DEC",
			"NUMERIC",
		}),
	}, Step{
		Template: constructManagedSqlTableTemplateWithColumnTypeUpdates(tableName, columnName, "2", []string{
			"int",
			"bigint",
			"float",
			"smallint",
			"tinyint",
			"decimal(10,0)",
			"decimal(10,0)",
			"decimal(10,0)",
		}),
	})
}

func TestUcAccResourceSqlTable_AddColumnComment(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "new comment"}}),
	})
}

func TestUcAccResourceSqlTable_DropColumnNullable(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: false, Comment: "comment"}}),
	})
}

func TestUcAccResourceSqlTable_MultipleColumnUpdates(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: false, Comment: "new comment"}}),
	})
}

func TestUcAccResourceSqlTable_ChangeColumnTypeThrows(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()

	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template:    constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "int", Nullable: true, Comment: "comment"}}),
		ExpectError: typeUpdateErrorRegex,
	})
}

func TestUcAccResourceSqlTable_ChangeColumnTypeWithMultipleWords(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()

	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "TIMESTAMP DEFAULT current_timestamp()", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "timestamp", Nullable: true, Comment: "comment"}}),
	})
}

func TestUcAccResourceSqlTable_DropColumn(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{
			{Name: "name", Type: "string", Nullable: true, Comment: "comment"},
			{Name: "nametwo", Type: "string", Nullable: true, Comment: "comment"},
		}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	})
}

func TestUcAccResourceSqlTable_DropMultipleColumns(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{
			{Name: "name", Type: "string", Nullable: true, Comment: "comment"},
			{Name: "nametwo", Type: "string", Nullable: true, Comment: "comment"},
			{Name: "namethree", Type: "string", Nullable: true, Comment: "comment"},
		}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	})
}

func TestUcAccResourceSqlTable_AddColumn(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{
			{Name: "name", Type: "string", Nullable: true, Comment: "comment"},
			{Name: "nametwo", Type: "string", Nullable: true, Comment: "comment"},
		}),
	})
}

func TestUcAccResourceSqlTable_AddMultipleColumns(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}
	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{
			{Name: "name", Type: "string", Nullable: true, Comment: "comment"},
			{Name: "nametwo", Type: "string", Nullable: true, Comment: "comment"},
			{Name: "namethree", Type: "string", Nullable: true, Comment: "comment"},
		}),
	})
}

func TestUcAccResourceSqlTable_AddColumnAndUpdateThrows(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}

	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: true, Comment: "comment"}}),
	}, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{
			{Name: "name", Type: "string", Nullable: false, Comment: "new comment"},
			{Name: "nametwo", Type: "string", Nullable: true, Comment: "comment"},
		}),
		ExpectError: inlineAndMembershipChangeErrorRegex,
	})
}

func TestUcAccResourceSqlTable_DropColumnAndUpdateThrows(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		Skipf(t)("databricks_sql_table resource not available on GCP")
	}

	tableName := RandomName()
	UnityWorkspaceLevel(t, Step{
		Template: constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{
			{Name: "name", Type: "string", Nullable: true, Comment: "comment"},
			{Name: "nametwo", Type: "string", Nullable: true, Comment: "comment"},
		}),
	}, Step{
		Template:    constructManagedSqlTableTemplate(tableName, []catalog.SqlColumnInfo{{Name: "name", Type: "string", Nullable: false, Comment: "new comment"}}),
		ExpectError: inlineAndMembershipChangeErrorRegex,
	})
}
