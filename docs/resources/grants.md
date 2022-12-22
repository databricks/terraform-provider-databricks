---
subcategory: "Unity Catalog"
---
# databricks_grants Resource

-> **Note**
  This article refers to the privileges and inheritance model in Privilege Model version 1.0. If you created your metastore during the public preview (before August 25, 2022), you can upgrade to Privilege Model version 1.0 following [Upgrade to privilege inheritance](https://docs.databricks.com/data-governance/unity-catalog/hive-metastore.html)

In Unity Catalog all users initially have no access to data. Only Metastore Admins can create objects and can grant/revoke access on individual objects to users and groups. Every securable object in Unity Catalog has an owner. The owner can be any account-level user or group, called principals in general. The principal that creates an object becomes its owner. Owners receive `ALL_PRIVILEGES` on the securable object (e.g., `SELECT` and `MODIFY` on a table), as well as the permission to grant privileges to other principals.

Securable objects are hierarchical and privileges are inherited downward. The highest level object that privileges are inherited from is the catalog. This means that granting a privilege on a catalog or schema automatically grants the privilege to all current and future objects within the catalog or schema. Privileges that are granted on a metastore are not inherited.

Every `databricks_grants` resource must have exactly one securable identifier and one or more `grant` blocks with the following arguments:

- `principal` - User name, group name or service principal application ID.
- `privileges` - One or more privileges that are specific to a securable type.

The securable objects are:

- `METASTORE`: The top-level container for metadata. Each metastore exposes a three-level namespace (`catalog`.`schema`.`table`) that organizes your data.
- `CATALOG`: The first layer of the object hierarchy, used to organize your data assets.
- `SCHEMA`: Also known as databases, schemas are the second layer of the object hierarchy and contain tables and views.
- `TABLE`: The lowest level in the object hierarchy, tables can be  _external_ (stored in external locations in your cloud storage of choice) or _managed_ tables (stored in a storage container in your cloud storage that you create expressly for UC).
- `VIEW`: A read-only object created from one or more tables that is contained within a schema.
- `EXTERNAL LOCATION`: An object that contains a reference to a storage credential and a cloud storage path that is contained within a metatore.
- `STORAGE CREDENTIAL`: An object that encapsulates a long-term cloud credential that provides access to cloud storage that is contained within a metatore.
- `SHARE`: A logical grouping for the tables you intend to share using Delta Sharing. A share is contained within a Unity Catalog metastore.

Terraform will handle any configuration drift on every `terraform apply` run, even when grants are changed outside of Terraform state.

It is required to define all permissions for a securable in a single resource, otherwise Terraform cannot guarantee config drift prevention.

Below summarizes which privilege types apply to each securable object in the catalog:

## Metastore grants

You can grant `CREATE_CATALOG`, `CREATE_EXTERNAL_LOCATION`, `CREATE_SHARE`, `CREATE_RECIPIENT` and `CREATE_PROVIDER` privileges to [databricks_metastore](metastore.md) id specified in `metastore` attribute.

```hcl
resource "databricks_grants" "sandbox" {
  metastore = databricks_metastore.this.id
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

You can grant `ALL_PRIVILEGES`, `CREATE_SCHEMA`, `USE_CATALOG` privileges to [databricks_catalog](catalog.md) specified in the `catalog` attribute. You can also grant `CREATE_FUNCTION`, `CREATE_TABLE`, `CREATE_VIEW`, `EXECUTE`, `MODIFY`, `SELECT` and `USE_SCHEMA` at the catalog level to apply them to the pertinent current and future securable objects within the catalog:

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

You can grant `ALL_PRIVILEGES`, `CREATE_FUNCTION`, `CREATE_TABLE`, `CREATE_VIEW` and `USE_SCHEMA` privileges to [_`catalog.schema`_](schema.md) specified in the `schema` attribute. You can also grant `EXECUTE`, `MODIFY` and `SELECT` at the schema level to apply them to the pertinent current and future securable objects within the schema:

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

You can grant `ALL_PRIVILEGES`, `SELECT` and `MODIFY` privileges to [_`catalog.schema.table`_](tables.md) specified in the `table` attribute.

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

You can grant `ALL_PRIVILEGES` and `SELECT` privileges to [_`catalog.schema.view`_](views.md) specified in `table` attribute.

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

You can grant `ALL_PRIVILEGES`, `CREATE_EXTERNAL_TABLE`, `READ_FILES` and `WRITE_FILES` privileges to [databricks_storage_credential](storage_credential.md) id specified in `storage_credential` attribute:

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

You can grant `ALL_PRIVILEGES`, `CREATE_EXTERNAL_TABLE`, `CREATE MANAGED STORAGE`, `READ_FILES` and `WRITE_FILES` privileges to [databricks_external_location](external_location.md) id specified in `external_location` attribute:

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
