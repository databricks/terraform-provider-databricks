---
subcategory: "Unity Catalog"
---
# databricks_grants Resource

In Unity Catalog all users initially have no access to data. Only Metastore Admins can create objects and can grant/revoke access on individual objects to users and groups. Every securable object in Unity Catalog has an owner. The owner can be any account-level user or group, called principals in general. The principal that creates an object becomes its owner. Owners receive all privileges on the securable object (e.g., `SELECT` and `MODIFY` on a table), as well as the permission to grant privileges to other principals.

Unity Catalog supports the following privileges on securable objects:

- `USAGE`: Applicable object types: `CATALOG`, `SCHEMA`. This privilege does not grant access to the securable itself, but is needed for a user to interact with any object within the securable. For example, to select data from a table, users need to have the `SELECT` privilege on that table and `USAGE` privileges on its parent schema and parent catalog.
  
  This is useful for allowing schema and catalog owners to be able to limit how far individual table owners can share data they produce. A table owner granting `SELECT` to another user does not allow that user read access to the table unless they also have `USAGE` on the schema and catalog.
- `SELECT`: Applicable object types: `TABLE`, `VIEW`. Allows a user to select from a table or view, if the user also has `USAGE` on its parent catalog and schema.
- `MODIFY`: Applicable object types: `TABLE`. Allows a user to add, update, and delete data to or from the table if the user also has `USAGE` on its parent catalog and schema.
- `CREATE`: Applicable object types: `CATALOG`, `SCHEMA`. If applied to a catalog, allows a user to create a schema. The user also requires the `USAGE` permission on the catalog.

  If applied to a schema, allows a user to create a table or view in the schema. The user also requires the `USAGE` permission on its parent catalog and the schema.
- `EXECUTE`: Applicable object types: `FUNCTION`. Allows a user to invoke a user defined function, if the user also has `USAGE` on its parent catalog and schema.
- `CREATE_TABLE`: Applicable object types: `EXTERNAL_LOCATION`, `STORAGE_CREDENTIAL`. Allows a user to create external tables directly in your cloud tenant using an external location or storage credential. Databricks recommends granting this privilege on an external location rather than storage credential; because the privilege is scoped to a path, it allows more control over where users can create external tables in your cloud tenant.
- `READ_FILES`: Applicable object types: `EXTERNAL_LOCATION`, `STORAGE_CREDENTIAL`. Allows a user to read files directly from your cloud tenant (for example from S3 or ADLS). Databricks recommends granting this privilege on an external location rather than storage credential; because the privilege is scoped to a path it allows more control over from where users can read data.
- `WRITE_FILES`: Applicable object types: `EXTERNAL_LOCATION`, `STORAGE_CREDENTIAL`. Allows a user to write files directly into your cloud tenant (for example into S3 or ADLS). We recommend granting this privilege on an external location rather than storage credential (since it is scoped to a path it allows more control over where users can write data to).
- `ALL_PRIVILEGES`: Applicable object types: All object types. Allows a user to grant or revoke all privileges applicable to the securable without explicitly specifying them. This expands to all available privileges at the time of the grant.

In Unity Catalog, privileges are not inherited on child securable objects. For example, if you grant the `CREATE` privilege on a catalog to a user, the user does not automatically have the `CREATE` privilege on all schemas in the catalog.

Every `databricks_grants` resource must have exactly one securable identifier and one or more `grant` blocks with the following arguments:

- `principal` - User or group name.
- `privileges` - One or more privileges that are specific to a securable type.

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

You can grant `CREATE` and `USAGE` privileges to [*`catalog`*.*`schema`*](schema.md) specified in the `schema` attribute:

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
    privileges = ["USAGE"]
  }
}
```

## Table grants

You can grant `MODIFY` and `SELECT` privileges to [*`catalog`*.*`schema`*.*`table`*](table.md) specified in the `table` attribute. You can define a table through [databricks_table](table.md) resource.

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

You can grant `SELECT` privileges to [*`catalog`*.*`schema`*.*`view`*](table.md) specified in `table` attribute. You can define a view through [databricks_table](table.md) resource.

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

## Storage credential grants

You can grant `CREATE_TABLE`, `READ_FILES`, `WRITE_FILES` and `CREATE_EXTERNAL_LOCATION` privileges to [databricks_storage_credential](storage_credential.md) id specified in `storage_credential` attribute:

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

## Metastore grants

You can grant `CREATE_CATALOG`, `CREATE_EXTERNAL_LOCATION`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_RECIPIENT`, and `CREATE_PROVIDER` privileges to [databricks_metastore](metastore.md) id specified in `metastore` attribute.

## Other access control

You can control Databricks General Permissions through [databricks_permissions](permissions.md) resource.
