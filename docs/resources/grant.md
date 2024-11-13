---
subcategory: "Unity Catalog"
---
# databricks_grant Resource

-> This article refers to the privileges and inheritance model in Privilege Model version 1.0. If you created your metastore during the public preview (before August 25, 2022), you can upgrade to Privilege Model version 1.0 following [Upgrade to privilege inheritance](https://docs.databricks.com/data-governance/unity-catalog/hive-metastore.html)

-> Most of Unity Catalog APIs are only accessible via **workspace-level APIs**. This design may change in the future. Account-level principal grants can be assigned with any valid workspace as the Unity Catalog is decoupled from specific workspaces. More information in [the official documentation](https://docs.databricks.com/data-governance/unity-catalog/index.html).

In Unity Catalog all users initially have no access to data. Only Metastore Admins can create objects and can grant/revoke access on individual objects to users and groups. Every securable object in Unity Catalog has an owner. The owner can be any account-level user or group, called principals in general. The principal that creates an object becomes its owner. Owners receive `ALL_PRIVILEGES` on the securable object (e.g., `SELECT` and `MODIFY` on a table), as well as the permission to grant privileges to other principals.

Securable objects are hierarchical and privileges are inherited downward. The highest level object that privileges are inherited from is the catalog. This means that granting a privilege on a catalog or schema automatically grants the privilege to all current and future objects within the catalog or schema. Privileges that are granted on a metastore are not inherited.

Every `databricks_grant` resource must have exactly one securable identifier and the following arguments:

- `principal` - User name, group name or service principal application ID.
- `privileges` - One or more privileges that are specific to a securable type.

For the latest list of privilege types that apply to each securable object in Unity Catalog, please refer to the [official documentation](https://docs.databricks.com/en/data-governance/unity-catalog/manage-privileges/privileges.html#privilege-types-by-securable-object-in-unity-catalog)

Terraform will handle any configuration drift for the specified principal on every `terraform apply` run, even when grants are changed outside of Terraform state.

See [databricks_grants](grants.md) for the list of privilege types that apply to each securable object.

## Examples

## Metastore grants

See [databricks_grants Metastore grants](grants.md#metastore-grants) for the list of privileges that apply to Metastores.

```hcl
resource "databricks_grant" "sandbox_data_engineers" {
  metastore = "metastore_id"

  principal  = "Data Engineers"
  privileges = ["CREATE_CATALOG", "CREATE_EXTERNAL_LOCATION"]
}

resource "databricks_grant" "sandbox_data_sharer" {
  metastore = "metastore_id"

  principal  = "Data Sharer"
  privileges = ["CREATE_RECIPIENT", "CREATE_SHARE"]
}
```

## Catalog grants

See [databricks_grants Catalog grants](grants.md#catalog-grants) for the list of privileges that apply to Catalogs.

```hcl
resource "databricks_catalog" "sandbox" {
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_grant" "sandbox_data_scientists" {
  catalog = databricks_catalog.sandbox.name

  principal  = "Data Scientists"
  privileges = ["USE_CATALOG", "USE_SCHEMA", "CREATE_TABLE", "SELECT"]
}

resource "databricks_grant" "sandbox_data_engineers" {
  catalog = databricks_catalog.sandbox.name

  principal  = "Data Engineers"
  privileges = ["USE_CATALOG", "USE_SCHEMA", "CREATE_SCHEMA", "CREATE_TABLE", "MODIFY"]
}

resource "databricks_grant" "sandbox_data_analyst" {
  catalog = databricks_catalog.sandbox.name

  principal  = "Data Analyst"
  privileges = ["USE_CATALOG", "USE_SCHEMA", "SELECT"]
}
```

## Schema grants

See [databricks_grants Schema grants](grants.md#schema-grants) for the list of privileges that apply to Schemas.

```hcl
resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this schema is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_grant" "things" {
  schema = databricks_schema.things.id

  principal  = "Data Engineers"
  privileges = ["USE_SCHEMA", "MODIFY"]
}
```

## Table grants

See [databricks_grants Table grants](grants.md#table-grants) for the list of privileges that apply to Tables.

```hcl
resource "databricks_grant" "customers_data_engineers" {
  table = "main.reporting.customers"

  principal  = "Data Engineers"
  privileges = ["MODIFY", "SELECT"]
}

resource "databricks_grant" "customers_data_analysts" {
  table = "main.reporting.customers"

  principal  = "Data Analysts"
  privileges = ["SELECT"]
}
```

You can also apply grants dynamically with [databricks_tables](../data-sources/tables.md) data resource:

```hcl
data "databricks_tables" "things" {
  catalog_name = "sandbox"
  schema_name  = "things"
}

resource "databricks_grant" "things" {
  for_each = data.databricks_tables.things.ids

  table = each.value

  principal  = "sensitive"
  privileges = ["SELECT", "MODIFY"]
}
```

## View grants

See [databricks_grants View grants](grants.md#view-grants) for the list of privileges that apply to Views.

```hcl
resource "databricks_grant" "customer360" {
  table = "main.reporting.customer360"

  principal  = "Data Analysts"
  privileges = ["SELECT"]
}
```

You can also apply grants dynamically with [databricks_views](../data-sources/views.md) data resource:

```hcl
data "databricks_views" "customers" {
  catalog_name = "main"
  schema_name  = "customers"
}

resource "databricks_grant" "customers" {
  for_each = data.databricks_views.customers.ids

  table = each.value

  principal  = "sensitive"
  privileges = ["SELECT", "MODIFY"]
}
```

## Volume grants

See [databricks_grants Volume grants](grants.md#volume-grants) for the list of privileges that apply to Volumes.

```hcl
resource "databricks_volume" "this" {
  name             = "quickstart_volume"
  catalog_name     = databricks_catalog.sandbox.name
  schema_name      = databricks_schema.things.name
  volume_type      = "EXTERNAL"
  storage_location = databricks_external_location.some.url
  comment          = "this volume is managed by terraform"
}

resource "databricks_grant" "volume" {
  volume = databricks_volume.this.id

  principal  = "Data Engineers"
  privileges = ["WRITE_VOLUME"]
}
```

## Registered model grants

See [databricks_grants Registered model grants](grants.md#registered-model-grants) for the list of privileges that apply to Registered models.

```hcl
resource "databricks_grant" "customers_data_engineers" {
  model = "main.reporting.customer_model"

  principal  = "Data Engineers"
  privileges = ["APPLY_TAG", "EXECUTE"]
}

resource "databricks_grant" "customers_data_analysts" {
  model = "main.reporting.customer_model"

  principal  = "Data Analysts"
  privileges = ["EXECUTE"]
}
```

## Function grants

See [databricks_grants Function grants](grants.md#function-grants) for the list of privileges that apply to Registered models.

```hcl
resource "databricks_grant" "udf_data_engineers" {
  function = "main.reporting.udf"

  principal  = "Data Engineers"
  privileges = ["EXECUTE"]
}

resource "databricks_grant" "udf_data_analysts" {
  function = "main.reporting.udf"

  principal  = "Data Analysts"
  privileges = ["EXECUTE"]
}
```

## Service credential grants

See [databricks_grants Service credential grants](grants.md#service-credential-grants) for the list of privileges that apply to Service credentials.

```hcl
resource "databricks_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_grant" "external_creds" {
  credential = databricks_credential.external.id

  principal  = "Data Engineers"
  privileges = ["ACCESS"]
}
```

## Storage credential grants

See [databricks_grants Storage credential grants](grants.md#storage-credential-grants) for the list of privileges that apply to Storage credentials.

```hcl
resource "databricks_storage_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_grant" "external_creds" {
  storage_credential = databricks_storage_credential.external.id

  principal  = "Data Engineers"
  privileges = ["CREATE_EXTERNAL_TABLE"]
}
```

## External location grants

See [databricks_grants External location grants](grants.md#external-location-grants) for the list of privileges that apply to External locations.

```hcl
resource "databricks_external_location" "some" {
  name            = "external"
  url             = "s3://${aws_s3_bucket.external.id}/some"
  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
}

resource "databricks_grant" "some_data_engineers" {
  external_location = databricks_external_location.some.id

  principal  = "Data Engineers"
  privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
}

resource "databricks_grant" "some_service_principal" {
  external_location = databricks_external_location.some.id

  principal  = databricks_service_principal.my_sp.application_id
  privileges = ["USE_SCHEMA", "MODIFY"]
}

resource "databricks_grant" "some_group" {
  external_location = databricks_external_location.some.id

  principal  = databricks_group.my_group.display_name
  privileges = ["USE_SCHEMA", "MODIFY"]
}

resource "databricks_grant" "some_user" {
  external_location = databricks_external_location.some.id

  principal  = databricks_group.my_user.user_name
  privileges = ["USE_SCHEMA", "MODIFY"]
}
```

## Connection grants

See [databricks_grants Connection grants](grants.md#connection-grants) for the list of privileges that apply to Connections.

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

resource "databricks_grant" "some" {
  foreign_connection = databricks_connection.mysql.name

  principal  = "Data Engineers"
  privileges = ["CREATE_FOREIGN_CATALOG", "USE_CONNECTION"]
}
```

## Delta Sharing share grants

See [databricks_grants Delta Sharing share grants](grants.md#delta-sharing-share-grants) for the list of privileges that apply to Delta Sharing shares.

```hcl
resource "databricks_share" "some" {
  name = "my_share"
}

resource "databricks_recipient" "some" {
  name = "my_recipient"
}

resource "databricks_grant" "some" {
  share = databricks_share.some.name

  principal  = databricks_recipient.some.name
  privileges = ["SELECT"]
}
```

## Other access control

You can control Databricks General Permissions through [databricks_permissions](permissions.md) resource.

## Import

The resource can be imported using combination of securable type (`table`, `catalog`, `foreign_connection`, ...), it's name and `principal`:

```bash
terraform import databricks_grant.this catalog/abc/user_name
```
