---
subcategory: "Security"
---
# databricks_service_principal_secret Resource

With this resource you can create a secret for a given [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html).

-> This resource can only be used with an account-level provider!

This secret can be used to configure the Databricks Terraform Provider to authenticate with the service principal. See [Authenticating with service principal](../index.md#authenticating-with-service-principal).

Additionally, the secret can be used to request OAuth tokens for the service principal, which can be used to authenticate to Databricks REST APIs. See [Authentication using OAuth tokens for service principals](https://docs.databricks.com/dev-tools/authentication-oauth.html).

## Example Usage

Create service principal secret

```hcl
resource "databricks_service_principal_secret" "terraform_sp" {
  service_principal_id = databricks_service_principal.this.id
}
```

A secret can be automatically rotated by taking a dependency on the `time_rotating` resource:

```hcl
resource "time_rotating" "this" {
  rotation_days = 30
}

resource "databricks_service_principal_secret" "terraform_sp" {
  service_principal_id = databricks_service_principal.this.id

  # Token is valid for 60 days but is rotated after 30 days.
  time_rotating = "Terraform (created: ${time_rotating.this.rfc3339})"
}
```

## Argument Reference

The following arguments are available:

* `service_principal_id` (Required, string) - SCIM ID of the [databricks_service_principal](service_principal.md) (not application ID).
* `lifetime` (Optional, string) - The lifetime of the secret in seconds formatted as `NNNNs`. If this parameter is not provided, the secret will have a default lifetime of 730 days (`63072000s`).  Expiration of secret will lead to generation of new secret.
* `time_rotating` - (Optional, string) - Changing this argument forces recreation of the secret.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the secret
* `secret` - **Sensitive** Generated secret for the service principal.
* `create_time` - UTC time when the secret was created.
* `expire_time` - UTC time when the secret will expire. If the field is not present, the secret does not expire.
* `secret_hash` - Secret Hash.
* `status`  - Status of the secret (i.e., `ACTIVE` - see [REST API docs for full list](https://docs.databricks.com/api/account/serviceprincipalsecrets/list#secrets-status)).
* `update_time` - UTC time when the secret was updated.

## Related Resources

The following resources are often used in the same context:

* [databricks_service_principal](service_principal.md) to manage [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html) in Databricks
