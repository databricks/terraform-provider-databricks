---
subcategory: "Deployment"
---

# databricks_instance_profiles Data Source

Lists all available [databricks_instance_profiles](../resources/instance_profile.md).

## Example Usage

Get all instance profiles:

```hcl
data "databricks_instance_profiles" "all" {
}

output "all_instance_profiles" {
  value = data.databricks_instance_profiles.all.instance_profiles
}
```

## Argument Reference

There are no arguments available for this data source.

## Attribute Reference

This data source exports the following attributes:
* `instance_profiles` - Set of objects for a [databricks_instance_profile](../resources/instance_profile.md). This contains the following attributes:
  * `name` - Name of the instance profile.
  * `arn` - ARN of the instance profile.
  * `role_arn` - ARN of the role attached to the instance profile.
  * `is_meta` - Whether the instance profile is a meta instance profile or not.
