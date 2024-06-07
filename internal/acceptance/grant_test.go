package acceptance

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var grantTemplate = `
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

resource "databricks_storage_credential" "external" {
	name = "cred-{var.STICKY_RANDOM}"
	aws_iam_role {
		role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
	}
	comment = "Managed by TF"
}

resource "databricks_external_location" "some" {
	name            = "external-{var.STICKY_RANDOM}"
	url             = "s3://{env.TEST_BUCKET}/some{var.STICKY_RANDOM}"
	credential_name = databricks_storage_credential.external.id
	comment         = "Managed by TF"
}

# resource "databricks_grant" "metastore" {
# 	metastore = "{env.TEST_METASTORE_ID}"
#
# 	principal  = "%s"
# 	privileges = ["CREATE_STORAGE_CREDENTIAL"]
# }

resource "databricks_grant" "catalog" {
	catalog = databricks_catalog.sandbox.id

	principal  = "%s"
	privileges = ["ALL_PRIVILEGES"]
}

resource "databricks_grant" "schema" {
	schema = databricks_schema.things.id

	principal  = "%s"
	privileges = ["ALL_PRIVILEGES"]
}

resource "databricks_grant" "table" {
	table = databricks_table.mytable.id

	principal  = "%s"
	privileges = ["ALL_PRIVILEGES"]
}

resource "databricks_grant" "cred" {
	storage_credential = databricks_storage_credential.external.id

	principal  = "%s"
	privileges = ["ALL_PRIVILEGES"]
}

resource "databricks_grant" "some" {
	external_location = databricks_external_location.some.id

	principal  = "%s"
	privileges = ["ALL_PRIVILEGES"]
}`

func TestUcAccGrant(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: strings.ReplaceAll(grantTemplate, "%s", "{env.TEST_DATA_ENG_GROUP}"),
	}, step{
		Template: strings.ReplaceAll(grantTemplate, "%s", "{env.TEST_DATA_SCI_GROUP}"),
	}, step{
		Template: strings.ReplaceAll(grantTemplate, `"%s"`, `upper("{env.TEST_DATA_SCI_GROUP}")`),
	}, step{
		Template: strings.ReplaceAll(grantTemplate, "ALL_PRIVILEGES", "ALL PRIVILEGES"),
	})
}

func grantTemplateForNamePermissionChange(suffix string, permission string) string {
	return fmt.Sprintf(`
	resource "databricks_storage_credential" "external" {
		name = "cred-{var.STICKY_RANDOM}%s"
		aws_iam_role {
			role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
		}
		comment = "Managed by TF"
	}
	
	resource "databricks_grant" "cred" {
		storage_credential = databricks_storage_credential.external.id
		principal  = "{env.TEST_DATA_ENG_GROUP}"
		privileges = ["%s"]
	}
	`, suffix, permission)
}

func TestUcAccGrantForIdChange(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: grantTemplateForNamePermissionChange("-old", "ALL_PRIVILEGES"),
	}, step{
		Template: grantTemplateForNamePermissionChange("-new", "ALL_PRIVILEGES"),
	}, step{
		Template:    grantTemplateForNamePermissionChange("-fail", "abc"),
		ExpectError: regexp.MustCompile(`cannot create grant: Privilege ABC is not applicable to this entity`),
	})
}
