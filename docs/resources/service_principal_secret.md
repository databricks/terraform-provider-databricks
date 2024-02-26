---
subcategory: "Security"
---
# databricks_service_principal_secret Resource

-> **Note** This resource can only be used with an account-level provider.

With this resource you can create a secret for a given [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html).

This secret can be used to configure the Databricks Terraform Provider to authenticate with the service principal. See [Authenticating with service principal](../index.md#authenticating-with-service-principal).

Additionally, the secret can be used to request OAuth tokens for the service principal, which can be used to authenticate to Databricks REST APIs. See [Authentication using OAuth tokens for service principals](https://docs.databricks.com/dev-tools/authentication-oauth.html).

## Example Usage

Create service principal secret

```hcl
resource "databricks_service_principal_secret" "terraform_sp" {
  service_principal_id = databricks_service_principal.this.id
}
```

## Argument Reference

The following arguments are available:

* `service_principal_id` - ID of the [databricks_service_principal](service_principal.md) (not application ID).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the secret
* `secret` - Generated secret for the service principal

## Related Resources

The following resources are often used in the same context:

* [databricks_service_principal](service_principal.md) to manage [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html) in Databricks
