---
subcategory: "Unity Catalog"
---
# databricks_entity_tag_assignments Data Source


## Example Usage


## Arguments
The following arguments are supported:
* `entity_name` (string, required) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string, required) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `max_results` (integer, optional) - Optional. Maximum number of tag assignments to return in a single page



## Attributes
This data source exports a single attribute, `tag_assignments`. It is a list of resources, each with the following attributes:
* `entity_name` (string) - The fully qualified name of the entity to which the tag is assigned
* `entity_type` (string) - The type of the entity to which the tag is assigned. Allowed values are: catalogs, schemas, tables, columns, volumes
* `tag_key` (string) - The key of the tag
* `tag_value` (string) - The value of the tag