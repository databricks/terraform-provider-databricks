---
subcategory: "Unity Catalog"
---
# databricks_grants Resource

-> **Note**
  This article refers to the privileges and inheritance model in Privilege Model version 1.0. If you created your metastore during the public preview (before August 25, 2022), you can upgrade to Privilege Model version 1.0 following [Upgrade to privilege inheritance](https://docs.databricks.com/data-governance/unity-catalog/hive-metastore.html)

-> **Note**
  Unity Catalog APIs are accessible via **workspace-level APIs**. This design may change in the future. Account-level principal grants can be assigned with any valid workspace as the Unity Catalog is decoupled from specific workspaces. More information in [the official documentation](https://docs.databricks.com/data-governance/unity-catalog/index.html).

Two different resources help you manage your Unity Catalog grants for a securable. Each of these resources serves a different use case:

- [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants): Authoritative. Sets the grants of a securable and replaces any existing grants defined inside or outside of Terraform.
- [databricks_grant](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grant): Authoritative for a given principal. Updates the grants of a securable to a single principal. Other principals within the grants for the securables are preserved.

In Unity Catalog all users initially have no access to data. Only Metastore Admins can create objects and can grant/revoke access on individual objects to users and groups. Every securable object in Unity Catalog has an owner. The owner can be any account-level user or group, called principals in general. The principal that creates an object becomes its owner. Owners receive `ALL_PRIVILEGES` on the securable object (e.g., `SELECT` and `MODIFY` on a table), as well as the permission to grant privileges to other principals.

Securable objects are hierarchical and privileges are inherited downward. The highest level object that privileges are inherited from is the catalog. This means that granting a privilege on a catalog or schema automatically grants the privilege to all current and future objects within the catalog or schema. Privileges that are granted on a metastore are not inherited.

Every `databricks_grants` resource must have exactly one securable identifier and one or more `grant` blocks with the following arguments:

- `principal` - User name, group name or service principal application ID.
- `privileges` - One or more privileges that are specific to a securable type.

For the latest list of privilege types that apply to each securable object in Unity Catalog, please refer to the [official documentation](https://docs.databricks.com/en/data-governance/unity-catalog/manage-privileges/privileges.html#privilege-types-by-securable-object-in-unity-catalog)

Terraform will handle any configuration drift on every `terraform apply` run, even when grants are changed outside of Terraform state.

Unlike the [SQL specification](https://docs.databricks.com/sql/language-manual/sql-ref-privileges.html#privilege-types), all privileges to be written with underscore instead of space, e.g. `CREATE_TABLE` and not `CREATE TABLE`. Below summarizes which privilege types apply to each securable object in the catalog:

## Metastore grants

You can grant `CREATE_CATALOG`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `MANAGE_ALLOWLIST`, `SET_SHARE_PERMISSION`, `USE_MARKETPLACE_ASSETS`, `USE_CONNECTION`, `USE_PROVIDER`, `USE_RECIPIENT` and `USE_SHARE` privileges to [databricks_metastore](metastore.md) assigned to the workspace.

```hcl
resource "databricks_grants" "sandbox" {
  metastore = "metastore_id"
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_CATALOG", "CREATE_EXTERNAL_LOCATION"]
  }
  grant {
    principal  = "Data Sharer"
    privileges = ["CREATE_RECIPIENT", "CREATE_SHARE"]
  }
}
```

## Catalog grants

You can grant `ALL_PRIVILEGES`, `APPLY_TAG`, `CREATE_CONNECTION`, `CREATE_SCHEMA`, `USE_CATALOG` privileges to [databricks_catalog](catalog.md) specified in the `catalog` attribute. You can also grant `CREATE_FUNCTION`, `CREATE_TABLE`, `CREATE_VOLUME`, `EXECUTE`, `MODIFY`, `REFRESH`, `SELECT`, `READ_VOLUME`, `WRITE_VOLUME` and `USE_SCHEMA` at the catalog level to apply them to the pertinent current and future securable objects within the catalog:

```hcl
resource "databricks_catalog" "sandbox" {
  name    = "sandbox"
  comment = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_grants" "sandbox" {
  catalog = databricks_catalog.sandbox.name
  grant {
    principal  = "Data Scientists"
    privileges = ["USE_CATALOG", "USE_SCHEMA", "CREATE_TABLE", "SELECT"]
  }
  grant {
    principal  = "Data Engineers"
    privileges = ["USE_CATALOG", "USE_SCHEMA", "CREATE_SCHEMA", "CREATE_TABLE", "MODIFY"]
  }
  grant {
    principal  = "Data Analyst"
    privileges = ["USE_CATALOG", "USE_SCHEMA", "SELECT"]
  }
}
```

## Schema grants

You can grant `ALL_PRIVILEGES`, `APPLY_TAG`, `CREATE_FUNCTION`, `CREATE_TABLE`, `CREATE_VOLUME` and `USE_SCHEMA` privileges to [_`catalog.schema`_](schema.md) specified in the `schema` attribute. You can also grant `EXECUTE`, `MODIFY`, `REFRESH`, `SELECT`, `READ_VOLUME`, `WRITE_VOLUME` at the schema level to apply them to the pertinent current and future securable objects within the schema:

```hcl
resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this schema is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_grants" "things" {
  schema = databricks_schema.things.id
  grant {
    principal  = "Data Engineers"
    privileges = ["USE_SCHEMA", "MODIFY"]
  }
}
```

## Table grants

You can grant `ALL_PRIVILEGES`, `APPLY_TAG`, `SELECT` and `MODIFY` privileges to [_`catalog.schema.table`_](tables.md) specified in the `table` attribute.

```hcl
resource "databricks_grants" "customers" {
  table = "main.reporting.customers"
  grant {
    principal  = "Data Engineers"
    privileges = ["MODIFY", "SELECT"]
  }
  grant {
    principal  = "Data Analysts"
    privileges = ["SELECT"]
  }
}
```

You can also apply grants dynamically with [databricks_tables](../data-sources/tables.md) data resource:

```hcl
data "databricks_tables" "things" {
  catalog_name = "sandbox"
  schema_name  = "things"
}

resource "databricks_grants" "things" {
  for_each = data.databricks_tables.things.ids

  table = each.value

  grant {
    principal  = "sensitive"
    privileges = ["SELECT", "MODIFY"]
  }
}
```

## View grants

You can grant `ALL_PRIVILEGES`, `APPLY_TAG` and `SELECT` privileges to [_`catalog.schema.view`_](views.md) specified in `table` attribute.

```hcl
resource "databricks_grants" "customer360" {
  table = "main.reporting.customer360"
  grant {
    principal  = "Data Analysts"
    privileges = ["SELECT"]
  }
}
```

You can also apply grants dynamically with [databricks_views](../data-sources/views.md) data resource:

```hcl
data "databricks_views" "customers" {
  catalog_name = "main"
  schema_name  = "customers"
}

resource "databricks_grants" "customers" {
  for_each = data.databricks_views.customers.ids

  table = each.value

  grant {
    principal  = "sensitive"
    privileges = ["SELECT", "MODIFY"]
  }
}
```

## Volume grants

You can grant `ALL_PRIVILEGES`, `READ_VOLUME` and `WRITE_VOLUME` privileges to [_`catalog.schema.volume`_](volumes.md) specified in the `volume` attribute.

```hcl
resource "databricks_volume" "this" {
  name             = "quickstart_volume"
  catalog_name     = databricks_catalog.sandbox.name
  schema_name      = databricks_schema.things.name
  volume_type      = "EXTERNAL"
  storage_location = databricks_external_location.some.url
  comment          = "this volume is managed by terraform"
}

resource "databricks_grants" "volume" {
  volume = databricks_volume.this.id
  grant {
    principal  = "Data Engineers"
    privileges = ["WRITE_VOLUME"]
  }
}
```

## Registered model grants

You can grant `ALL_PRIVILEGES`, `APPLY_TAG`, and `EXECUTE` privileges to [_`catalog.schema.model`_](registered_model.md) specified in the `model` attribute.

```hcl
resource "databricks_grants" "customers" {
  model = "main.reporting.customer_model"
  grant {
    principal  = "Data Engineers"
    privileges = ["APPLY_TAG", "EXECUTE"]
  }
  grant {
    principal  = "Data Analysts"
    privileges = ["EXECUTE"]
  }
}
```

## Function grants

You can grant `ALL_PRIVILEGES` and `EXECUTE` privileges to _`catalog.schema.function`_ specified in the `function` attribute.

```hcl
resource "databricks_grants" "udf" {
  function = "main.reporting.udf"

  grant {
    principal  = "Data Engineers"
    privileges = ["EXECUTE"]
  }
  grant {
    principal  = "Data Analysts"
    privileges = ["EXECUTE"]
  }
}
```

## Storage credential grants

You can grant `ALL_PRIVILEGES`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `READ_FILES` and `WRITE_FILES` privileges to [databricks_storage_credential](storage_credential.md) id specified in `storage_credential` attribute:

```hcl
resource "databricks_storage_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_grants" "external_creds" {
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_EXTERNAL_TABLE"]
  }
}
```

## External location grants

You can grant `ALL_PRIVILEGES`, `CREATE_EXTERNAL_TABLE`, `CREATE_MANAGED_STORAGE`, `CREATE EXTERNAL VOLUME`, `READ_FILES` and `WRITE_FILES` privileges to [databricks_external_location](external_location.md) id specified in `external_location` attribute:

```hcl
resource "databricks_external_location" "some" {
  name            = "external"
  url             = "s3://${aws_s3_bucket.external.id}/some"
  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
}

resource "databricks_grants" "some" {
  external_location = databricks_external_location.some.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
  }
  grant {
    principal  = databricks_service_principal.my_sp.application_id
    privileges = ["USE_SCHEMA", "MODIFY"]
  }
  grant {
    principal  = databricks_group.my_group.display_name
    privileges = ["USE_SCHEMA", "MODIFY"]
  }
  grant {
    principal  = databricks_group.my_user.user_name
    privileges = ["USE_SCHEMA", "MODIFY"]
  }
}
```

## Connection grants

You can grant `ALL_PRIVILEGES`, `USE_CONNECTION` and `CREATE_FOREIGN_CATALOG` to [databricks_connection](connection.md) specified in `foreign_connection` attribute:

```hcl
resource "databricks_connection" "mysql" {
  name            = "mysql_connection"
  connection_type = "MYSQL"
  comment         = "this is a connection to mysql db"
  options = {
    host     = "test.mysql.database.azure.com"
    port     = "3306"
    user     = "user"
    password = "password"
  }
  properties = {
    purpose = "testing"
  }
}

resource "databricks_grants" "some" {
  foreign_connection = databricks_connection.mysql.name
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_FOREIGN_CATALOG", "USE_CONNECTION"]
  }
}
```

## Delta Sharing share grants

You can grant `SELECT` to [databricks_recipient](recipient.md) on [databricks_share](share.md) name specified in `share` attribute:

```hcl
resource "databricks_share" "some" {
  name = "my_share"
}

resource "databricks_recipient" "some" {
  name = "my_recipient"
}

resource "databricks_grants" "some" {
  share = databricks_share.some.name
  grant {
    principal  = databricks_recipient.some.name
    privileges = ["SELECT"]
  }
}
```

## Other access control

You can control Databricks General Permissions through [databricks_permissions](permissions.md) resource.

## Import

The resource can be imported using combination of securable type (`table`, `catalog`, `foreign_connection`, ...) and it's name:

```bash
terraform import databricks_grants.this catalog/abc
```
