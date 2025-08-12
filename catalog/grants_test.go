package catalog_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

var grantsTemplate = `
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
	name = "managed-{var.STICKY_RANDOM}"
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

resource "databricks_table" "metric_view_grants" {
	catalog_name = databricks_catalog.sandbox.id
	schema_name = databricks_schema.things.name
	name = "metric-view-{var.STICKY_RANDOM}"
	table_type = "METRIC_VIEW"
	data_source_format = ""

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

resource "databricks_grants" "metastore" {
	metastore = "{env.TEST_METASTORE_ID}"
	grant {
		principal  = "%s"
		privileges = ["CREATE_STORAGE_CREDENTIAL"]
	}
}

resource "databricks_grants" "catalog" {
	catalog = databricks_catalog.sandbox.id
	grant {
		principal  = "%s"
		privileges = ["ALL_PRIVILEGES"]
	}
}

resource "databricks_grants" "schema" {
	schema = databricks_schema.things.id
	grant {
		principal  = "%s"
		privileges = ["ALL_PRIVILEGES"]
	}
}

resource "databricks_grants" "table" {
	table = databricks_table.mytable.id
	grant {
		principal  = "%s"
		privileges = ["ALL_PRIVILEGES"]
	}
}

resource "databricks_grants" "metric_view_grants" {
	table = databricks_table.metric_view_grants.id
	grant {
		principal  = "%s"
		privileges = ["ALL_PRIVILEGES"]
	}
}

resource "databricks_grants" "cred" {
	storage_credential = databricks_storage_credential.external.id
	grant {
		principal  = "%s"
		privileges = ["ALL_PRIVILEGES"]
	}
}

resource "databricks_grants" "some" {
	external_location = databricks_external_location.some.id
	grant {
		principal  = "%s"
		privileges = ["ALL_PRIVILEGES"]
	}
}`

func TestUcAccGrants(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: strings.ReplaceAll(grantsTemplate, "%s", "{env.TEST_DATA_ENG_GROUP}"),
	}, acceptance.Step{
		Template: strings.ReplaceAll(grantsTemplate, "%s", "{env.TEST_DATA_SCI_GROUP}"),
	}, acceptance.Step{
		Template: strings.ReplaceAll(strings.ReplaceAll(grantsTemplate, "ALL_PRIVILEGES", "ALL PRIVILEGES"), "%s", "{env.TEST_DATA_ENG_GROUP}"),
	})
}

func grantsTemplateForNamePermissionChange(suffix string, permission string) string {
	return fmt.Sprintf(`
	resource "databricks_storage_credential" "external" {
		name = "cred-{var.STICKY_RANDOM}%s"
		aws_iam_role {
			role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
		}
		comment = "Managed by TF"
	}

	resource "databricks_grants" "cred" {
		storage_credential = databricks_storage_credential.external.id
		grant {
			principal  = "{env.TEST_DATA_ENG_GROUP}"
			privileges = ["%s"]
		}
	}
	`, suffix, permission)
}

func TestUcAccGrantsForIdChange(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: grantsTemplateForNamePermissionChange("-old", "ALL_PRIVILEGES"),
	}, acceptance.Step{
		Template: grantsTemplateForNamePermissionChange("-new", "ALL_PRIVILEGES"),
	}, acceptance.Step{
		Template:    grantsTemplateForNamePermissionChange("-fail", "abc"),
		ExpectError: regexp.MustCompile(`Error: cannot create grants: Privilege ABC is not applicable to this entity`),
	})
}
