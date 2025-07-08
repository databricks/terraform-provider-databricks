---
subcategory: "Machine Learning"
---
# databricks_feature_tag Resource


## Example Usage


## Arguments
The following arguments are supported:
* `key` (string, required) - 
* `value` (string, optional) - 

## Attributes
In addition to the above arguments, the following attributes are exported:

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = key
  to = databricks_feature_tag.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_feature_tag key
```