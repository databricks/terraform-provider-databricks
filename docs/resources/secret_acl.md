---
subcategory: "Security"
---
# databricks_secret_acl Resource

Create or overwrite the ACL associated with the given principal (user or group) on the specified [databricks_secret_scope](secret_scope.md). Please consult [Secrets User Guide](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) for more details.

-> This resource can only be used with a workspace-level provider!

## Example Usage

This way, data scientists can read the Publishing API key that is synchronized from, for example, Azure Key Vault.

```hcl
resource "databricks_group" "ds" {
  display_name = "data-scientists"
}

resource "databricks_secret_scope" "app" {
  name = "app-secret-scope"
}

resource "databricks_secret_acl" "my_secret_acl" {
  principal  = databricks_group.ds.display_name
  permission = "READ"
  scope      = databricks_secret_scope.app.name
}

resource "databricks_secret" "publishing_api" {
  key = "publishing_api"
  // replace it with a secret management solution of your choice :-)
  string_value = data.azurerm_key_vault_secret.example.value
  scope        = databricks_secret_scope.app.name
}
```

## Argument Reference

The following arguments are required:

* `scope` - (Required) name of the scope
* `principal` - (Required) principal's identifier. It can be:
  * `user_name` attribute of [databricks_user](user.md).
  * `display_name` attribute of [databricks_group](group.md).  Use `users` to allow access for all workspace users.
  * `application_id` attribute of [databricks_service_principal](service_principal.md).
* `permission` - (Required) `READ`, `WRITE` or `MANAGE`.

## Import

The resource secret acl can be imported using `scopeName|||principalName` combination.

```bash
terraform import databricks_secret_acl.object `scopeName|||principalName`
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_permissions](permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
* [databricks_secret](secret.md) to manage [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_secret_scope](secret_scope.md) to create [secret scopes](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
