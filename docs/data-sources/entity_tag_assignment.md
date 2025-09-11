---
subcategory: "Unity Catalog"
---
# databricks_entity_tag_assignment Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `entity_name` (string, required) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string, required) - The key of the tag
* `workspace_id` (string, optional) - Workspace ID of the resource

## Attributes
The following attributes are exported:
* `entity_name` (string) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string) - The key of the tag
* `tag_value` (string) - The value of the tag