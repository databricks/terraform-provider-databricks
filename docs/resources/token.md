# databricks_token Resource

This resource creates an api token that can be used to create Databricks resources. This will create an API token for the user that has authenticated on the provider. So if you have used an admin user to setup the provider then you will be making API tokens for that admin user. 

## Example Usage

```hcl
// initialize provider in normal mode
provider "databricks" {
  alias = "created_workspace" 
  
  host  = databricks_mws_workspaces.this.workspace_url
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

## Argument Reference

The following arguments are available:

* `lifetime_seconds` - (Optional) (Integer) The lifetime of the token, in seconds. If no lifetime is specified, the token remains valid indefinitely.
* `comment` - (Optional) (String) Comment that will appear on the user’s settings page for this token.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the token.
* `token_value` - **Sensitive** value of the newly-created token.