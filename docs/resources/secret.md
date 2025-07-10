---
subcategory: "Security"
---
# databricks_secret Resource

With this resource you can insert a secret under the provided scope with the given name. If a secret already exists with the same name, this command overwrites the existing secret's value. The server encrypts the secret using the secret scope's encryption settings before storing it. You must have WRITE or MANAGE permission on the secret scope. The secret key must consist of alphanumeric characters, dashes, underscores, and periods, and cannot exceed 128 characters. The maximum allowed secret value size is 128 KB. The maximum number of secrets in a given scope is 1000. You can read a secret value only from within a command on a [cluster](cluster.md) (for example, through a notebook); there is no API to read a secret value outside of a cluster. The permission applied is based on who is invoking the command and you must have at least READ permission. Please consult [Secrets User Guide](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) for more details.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_secret_scope" "app" {
  name = "application-secret-scope"
}

resource "databricks_secret" "publishing_api" {
  key          = "publishing_api"
  string_value = data.azurerm_key_vault_secret.example.value
  scope        = databricks_secret_scope.app.id
}

resource "databricks_cluster" "this" {
  # ...
  spark_conf = {
    # ...
    "fs.azure.account.oauth2.client.secret" = databricks_secret.publishing_api.config_reference
  }
}
```

## Argument Reference

The following arguments are required:

* `string_value` - (Required) (String) super secret sensitive value.
* `scope` - (Required) (String) name of databricks secret scope. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
* `key` - (Required) (String) key within secret scope. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the secret.
* `last_updated_timestamp` - (Integer) time secret was updated
* `config_reference` - (String) value to use as a secret reference in [Spark configuration and environment variables](https://docs.databricks.com/security/secrets/secrets.html#use-a-secret-in-a-spark-configuration-property-or-environment-variable): `{{secrets/scope/key}}`.

## Import

The resource secret can be imported using `scopeName|||secretKey` combination. **This may change in future versions.**

```hcl
import {
  to = databricks_secret.app
  id = "<scopeName>|||<secretKey>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_secret.app "<scopeName>|||<secretKey>"
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_pipeline](pipeline.md) to deploy [Lakeflow Declarative Pipelines](https://docs.databricks.com/aws/en/dlt).
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
* [databricks_secret_acl](secret_acl.md) to manage access to [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_secret_scope](secret_scope.md) to create [secret scopes](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
