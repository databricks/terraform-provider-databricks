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
**Filter Results**

Starts With:

```hcl
data "databricks_instance_profiles" "all" {
  filter {
    name    = "name"
    pattern = "^aws"
  }
}
```

Ends With:

```hcl
data "databricks_instance_profiles" "all" {
  filter {
    name    = "arn"
    pattern = "prod$"
  }
}
```

Contains:

```hcl
data "databricks_instance_profiles" "all" {
  filter {
    name    = "role_arn"
    pattern = "prod"
  }
}
```

Equals:

```hcl
data "databricks_instance_profiles" "all" {
  filter {
    name    = "is_meta"
    pattern = "^false$"
  }
}
```

## Argument Reference

* `filter` - (Optional) Configuration block for filtering. Detailed below.

### `filter` Configuration Block
The filter configuration block supports the following arguments:

* `name` - (Required) Name of the filter field. Valid values can be found in the [attribute reference](#instance_profiles) below.
* `pattern` - (Required) Regex pattern to filter using. Follows the [Go Regexp syntax](https://pkg.go.dev/regexp/syntax)

## Attribute Reference

This data source exports the following attributes:
* `instance_profiles` - Set of objects for a [databricks_instance_profile](../resources/instance_profile.md). This contains the following attributes:
  * `name` - Name of the instance profile.
  * `arn` - ARN of the instance profile.
  * `role_arn` - ARN of the role attached to the instance profile.
  * `is_meta` - Whether the instance profile is a meta instance profile or not.
