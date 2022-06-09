---
subcategory: "Unity Catalog"
---
# databricks_grants Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

In Unity Catalog all users initially have no access to data. Only Metastore Admins can create objects and can grant/revoke access on individual objects to users and groups. Every securable object in Unity Catalog has an owner. The owner can be any account-level user or group, called principals in general. The principal that creates an object becomes its owner. Owners receive all privileges on the securable object (e.g., `SELECT` and `MODIFY` on a table), as well as the permission to grant privileges to other principals.

Unity Catalog supports the following privileges on securable objects:
* `SELECT` - Allows the grantee to read data from the securable (applicable to tables and views).
* `MODIFY` - Allows the grantee to add, update and delete data to or from the securable. (applicable to tables)
* `CREATE` - Allows the grantee to create child objects within this securable.
* `USAGE` - This privilege does not grant access to the securable itself, but allows the grantee to traverse the securable in order to access its child objects. For example, to select data from a table, users need to have the `SELECT` privilege on that table and `USAGE` privileges on its parent schema and parent catalog. Thus, you can use this privilege to restrict access to sections of your data namespace to specific groups.

Every `databricks_grants` resource must have exactly one securable identifier and one or more `grant` blocks with the following arguments:

* `principal` - User or group name.
* `privileges` - One or more privileges that are specific to a securable type.

Terraform will handle any configuration drift on every `terraform apply` run, even when grants are changed outside of Terraform state.

It is required to define all permissions for a securable in a single resource, otherwise Terraform cannot guarantee config drift prevention.

## Catalog grants

You can grant `CREATE` and `USAGE` privileges to [databricks_catalog](catalog.md) specified in the `catalog` attribute:

```hcl
resource "databricks_catalog" "sandbox" {
  metastore_id = databricks_metastore.this.id
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_grants" "sandbox" {
  catalog = databricks_catalog.sandbox.name
  grant {
    principal  = "Data Scientists"
    privileges = ["USAGE", "CREATE"]
  }
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}
```

## Schema grants

You can grant `CREATE` and `USAGE` privileges to [*`catalog`*.*`database`*](schema.md) specified in the `schema` attribute:

```hcl
resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_grants" "things" {
  schema = databricks_schema.things.id
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}
```

## Table grants

You can grant `MODIFY` and `SELECT` privileges to [*`catalog`*.*`database`*.*`table`*](table.md) specified in the `table` attribute. You can define a table through [databricks_table](table.md) resource.

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

You can grant `SELECT` privileges to [*`catalog`*.*`database`*.*`view`*](table.md) specified in `view` attribute. You can define a view through [databricks_table](table.md) resource.

```hcl
resource "databricks_grants" "customer360" {
  view = "main.reporting.customer360"
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

  view = each.value

  grant {
    principal  = "sensitive"
    privileges = ["SELECT", "MODIFY"]
  }
}
```

## Storage credential grants

You can grant `CREATE_TABLE`, `READ_FILES`, and `WRITE_FILES` privileges to [databricks_storage_credential](storage_credential.md) id specified in `storage_credential` attribute:

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
    privileges = ["CREATE_TABLE"]
  }
}
```

## Storage location grants

You can grant `CREATE_TABLE`, `READ_FILES`, and `WRITE_FILES` privileges to [databricks_external_location](external_location.md) id specified in `external_location` attribute:

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
    privileges = ["CREATE_TABLE", "READ_FILES"]
  }
}
```

## Other access control

You can control Databricks General Permissions through [databricks_permissions](permissions.md) resource.
