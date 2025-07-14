---
subcategory: "Machine Learning"
---
# databricks_materialized_features_feature_tag Resource
Preview Stage: `PRIVATE_PREVIEW`



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
  to = databricks_materialized_features_feature_tag.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_materialized_features_feature_tag key
```