---
subcategory: "Security"
---
# databricks_obo_token Resource

This resource creates On-Behalf-Of-Token for a Service Principal in AWS workspaces.

## Example Usage

```hcl
resource "databricks_service_principal" "this" {
  display_name = "Automation-only SP"
}

resource "databricks_permissions" "token_usage" {
  authorization = "tokens"
  access_control {
    service_principal_name = databricks_service_principal.this.application_id
    permission_level = "CAN_USE"
  }
}

resource "databricks_obo_token" "this" {
  depends_on = [databricks_permissions.token_usage]
  application_id = databricks_service_principal.this.application_id
  comment = "PAT on behalf of ${databricks_service_principal.this.display_name}"
  lifetime_seconds = 3600
}

output "obo" {
  value = databricks_obo_token.this.token_value
  sensitive = true
}
```

## Argument Reference

The following arguments are required:

* `application_id` - Application ID of [databricks_service_principal](service_principal.md#application_id) to create PAT token for.
* `lifetime_seconds` - (Integer) The number of seconds before the token expires. Token resource is re-created when it expires.
* `comment` - (String) Comment that describes the purpose of the token.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the token.
* `token_value` - **Sensitive** value of the newly-created token.