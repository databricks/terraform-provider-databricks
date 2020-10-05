# databricks_secret_scope Resource

This resource creates a Databricks-backed secret scope in which secrets are stored in Databricks-managed storage and encrypted with a cloud-based specific encryption key. 

The scope name:

* Must be unique within a workspace.
* Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.

Please consult [Secrets User Guide](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) for more details.

## Example Usage

```hcl
resource "databricks_secret_scope" "my-scope" {
  name = "terraform-demo-scope"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Scope name requested by the user. Scope names are unique. This field is required.
* `initial_manage_principal` - (Optional) The principal with the only possible value `users` that is initially granted `MANAGE` permission to the created scope.  If it's omitted, then the [databricks_secret_acl](secret_acl.md) with `MANAGE` permission applied to the scope is assigned to the API request issuer's user identity (see [documentation](https://docs.databricks.com/dev-tools/api/latest/secrets.html#create-secret-scope)). 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the secret scope object.

## Import

The secret resource scope can be imported using the scope name:

```bash
$ terraform import databricks_secret_scope.object <scopeName>
```
