---
subcategory: "Unity Catalog"
---
# databricks_entity_tag_assignment Resource


## Example Usage


## Arguments
The following arguments are supported:
* `entity_name` (string, required) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string, required) - The key of the tag
* `tag_value` (string, optional) - The value of the tag

## Attributes
In addition to the above arguments, the following attributes are exported:

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = entity_type,entity_name,tag_key
  to = databricks_entity_tag_assignment.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_entity_tag_assignment entity_type,entity_name,tag_key
```