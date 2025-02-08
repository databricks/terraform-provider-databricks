---
subcategory: "Security"
---
# databricks_token Resource

-> This resource can only be used with a workspace-level provider!

This resource creates [Personal Access Tokens](https://docs.databricks.com/sql/user/security/personal-access-tokens.html) for the same user that is authenticated with the provider. Most likely you should use [databricks_obo_token](obo_token.md) to create [On-Behalf-Of tokens](https://docs.databricks.com/administration-guide/users-groups/service-principals.html#manage-personal-access-tokens-for-a-service-principal) for a [databricks_service_principal](service_principal.md) in Databricks workspaces on AWS. Databricks workspaces on other clouds use their own native OAuth token flows.

## Example Usage

```hcl
// initialize provider in normal mode
provider "databricks" {
  alias = "created_workspace"

  host = databricks_mws_workspaces.this.workspace_url
}

// create PAT token to provision entities within workspace
resource "databricks_token" "pat" {
  provider = databricks.created_workspace
  comment  = "Terraform Provisioning"
  // 100 day token
  lifetime_seconds = 8640000
}

// output token for other modules
output "databricks_token" {
  value     = databricks_token.pat.token_value
  sensitive = true
}
```

A token can be automatically rotated by taking a dependency on the `time_rotating` resource:

```hcl
resource "time_rotating" "this" {
  rotation_days = 30
}

resource "databricks_token" "pat" {
  comment = "Terraform (created: ${time_rotating.this.rfc3339})"

  # Token is valid for 60 days but is rotated after 30 days.
  # Run `terraform apply` within 60 days to refresh before it expires.
  lifetime_seconds = 60 * 24 * 60 * 60
}
```

## Argument Reference

The following arguments are available:

* `lifetime_seconds` - (Optional) (Integer) The lifetime of the token, in seconds. If no lifetime is specified, then expire time will be set to maximum allowed by the workspace configuration or platform.
* `comment` - (Optional) (String) Comment that will appear on the user’s settings page for this token.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the token.
* `token_value` - **Sensitive** value of the newly-created token.

## Import

!> Importing this resource is not currently supported.
