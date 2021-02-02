# databricks_secret Resource

With this resource you can insert a secret under the provided scope with the given name. If a secret already exists with the same name, this command overwrites the existing secret’s value. The server encrypts the secret using the secret scope’s encryption settings before storing it. You must have WRITE or MANAGE permission on the secret scope. The secret key must consist of alphanumeric characters, dashes, underscores, and periods, and cannot exceed 128 characters. The maximum allowed secret value size is 128 KB. The maximum number of secrets in a given scope is 1000. You can read a secret value only from within a command on a [cluster](cluster.md) (for example, through a notebook); there is no API to read a secret value outside of a cluster. The permission applied is based on who is invoking the command and you must have at least READ permission. Please consult [Secrets User Guide](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) for more details.

## Example Usage

```hcl
resource "databricks_secret_scope" "app" {
    name = "application-secret-scope"
}
resource "databricks_secret" "publishing_api" {
    key = "publishing_api"
    string_value = data.azurerm_key_vault_secret.example.value
    scope = databricks_secret_scope.app.id
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


## Import

The resource secret can be imported using `scopeName|||secretKey` combination. **This may change in future versions.**

```bash
$ terraform import databricks_secret.app `scopeName|||secretKey`
```