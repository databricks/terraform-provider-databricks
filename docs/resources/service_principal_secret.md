---
subcategory: "Security"
---
# databricks_service_principal_service Resource

With this resource you can create a secret under given [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html)

With the secret you can configure Databricks Terraform Provider to authenticate with service principal UUID and secret. See [Authenticating with service principal ID and secret](../index.md#authenticating-with-hostname-client_id-and-client_secret).

You can also use the secret to request OAuth tokens for the service prinicipal, which can be used to authenticate to Databricks APIs. See [Authentication using OAuth tokens for service principals](https://docs.databricks.com/dev-tools/api/latest/authentication-oauth.html).

-> **Note** This resource in only available in AWS.

## Example Usage

Create service principal secret

```hcl
resource "databricks_service_principal_secret" "terraform_sp" {
  service_principal_id = "123456789012345"
}
```

## Argument Reference

The following arguments are available:

* `service_principal_id` - ID of the [databricks_service_principal](service_principal.md)


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - ID of the secret
- `secret` - Generated secret for the service principal


## Related Resources

The following resources are often used in the same context:

* [databricks_service_principal](service_principal.md) to manage [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html) in Databricks