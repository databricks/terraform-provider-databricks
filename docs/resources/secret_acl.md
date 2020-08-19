# databricks_secret_acl Resource

Create or overwrite the ACL associated with the given principal (user or group) on the specified scope point. Please consult [Secrets User Guide](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) for more details.

## Example Usage

This way data scientists can read Publishing API key that is synchronized from example Azure Key Vault.

```hcl
resource "databricks_group" "ds" {
  display_name = "data-scientists"
}

resource "databricks_secret_scope" "app" {
    name = "Application Secret Scope"
}

resource "databricks_secret_acl" "my_secret_acl" {
    principal = databricks_group.ds.display_name
    permission = "READ"
    scope = databricks_secret_scope.app.name
}

resource "databricks_secret" "publishing_api" {
    key = "publishing_api"
    // replace it with secret management solution of your choice :-)
    string_value = data.azurerm_key_vault_secret.example.value
    scope = databricks_secret_scope.app.name
}
```

## Argument Reference

The following arguments are required:

* `scope` - (Required) name of the scope
* `principal` - (Required) name of the principals. It can be `users` for all users or name or `display_name` of [databricks_group](group.md)
* `permission` - (Required) `READ`, `WRITE` or `MANAGE`. 

## Import

The resource secret acl can be imported using `scopeName|||principalName` combination. **This may change in future versions.**

```bash
$ terraform import databricks_secret_acl.object `scopeName|||principalName`
```