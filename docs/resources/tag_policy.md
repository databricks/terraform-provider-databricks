---
subcategory: "Tags"
---
# databricks_tag_policy Resource


## Example Usage


## Arguments
The following arguments are supported:
* `tag_key` (string, required)
* `description` (string, optional)
* `values` (list of Value, optional)

### Value
* `name` (string, required)

## Attributes
In addition to the above arguments, the following attributes are exported:
* `id` (string)

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = tag_key
  to = databricks_tag_policy.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_tag_policy tag_key
```