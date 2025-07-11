---
subcategory: "Machine Learning"
---
# databricks_online_store Resource


## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - The name of the online store. This is the unique identifier for the online store
* `capacity` (string, optional) - The capacity of the online store. Valid values are "CU_1", "CU_2", "CU_4", "CU_8"

## Attributes
In addition to the above arguments, the following attributes are exported:
* `creation_time` (string) - The timestamp when the online store was created
* `creator` (string) - The email of the creator of the online store
* `state` (string) - The current state of the online store. Possible values are: `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = name
  to = databricks_online_store.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_online_store name
```