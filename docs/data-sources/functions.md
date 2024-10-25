---
subcategory: "Unity Catalog"
---
# databricks_functionss Data Source

-> **Note** This data source can only be used with a workspace-level provider!

Retrieves a list of [User-Defined Functions (UDFs) registered in the Unity Catalog](https://docs.databricks.com/en/udf/unity-catalog.html).

## Example Usage

List all functions defined in a specific schema (`main.default` in this example):

```hcl
data "databricks_functions" "all" {
  catalog_name = "main"
  schema_name  = "default"
}

output "all_external_locations" {
  value = data.databricks_functions.all.functions
}
```

## Argument Reference

The following arguments are supported:

* `catalog_name` - (Required) Name of [databricks_catalog](../resources/catalog.md).
* `schema_name` - (Required) Name of [databricks_schema](../resources/schema.md).
* `include_browse` - (Optional, Boolean) flag to specify if include UDFs in the response for which the principal can only access selective metadata for.

## Attribute Reference

This data source exports the following attributes:

* `functions` - list of objects describing individual UDF. Each object consists of the following attributes (refer to [REST API documentation](https://docs.databricks.com/api/workspace/functions/list#functions) for up-to-date list of attributes. Default type is String):
  * `name` - Name of function, relative to parent schema.
  * `catalog_name` - Name of parent catalog.
  * `schema_name` - Name of parent schema relative to its parent catalog.
  * `input_params` - object describing input parameters. Consists of the single attribute:
    * `parameters` - The array of definitions of the function's parameters:
      * `name` - Name of parameter.
      * `type_text` - Full data type spec, SQL/catalogString text.
      * `type_json` - Full data type spec, JSON-serialized.
      * `type_name` - Name of type (INT, STRUCT, MAP, etc.).
      * `type_precision` - Digits of precision; required on Create for DecimalTypes.
      * `type_scale` - Digits to right of decimal; Required on Create for DecimalTypes.
      * `type_interval_type` - Format of IntervalType.
      * `position` - Ordinal position of column (starting at position 0).
      * `parameter_mode` - The mode of the function parameter.
      * `parameter_type` - The type of function parameter (`PARAM` or `COLUMN`).
      * `parameter_default` - Default value of the parameter.
      * `comment` - User-provided free-form text description.
  * `return_params` - Table function return parameters.  See `input_params` for description.
  * `data_type` - Scalar function return data type.
  * `full_data_type` - Pretty printed function data type.
  * `routine_body` - Function language (`SQL` or `EXTERNAL`). When `EXTERNAL` is used, the language of the routine function should be specified in the `external_language` field, and the `return_params` of the function cannot be used (as `TABLE` return type is not supported), and the `sql_data_access` field must be `NO_SQL`.
  * `routine_definition` - Function body.
  * `routine_dependencies` - Function dependencies.
  * `parameter_style` - Function parameter style. `S` is the value for SQL.
  * `is_deterministic` - Boolean flag specifying whether the function is deterministic.
  * `sql_data_access` - Function SQL data access (`CONTAINS_SQL`, `READS_SQL_DATA`, `NO_SQL`).
  * `is_null_call` - Boolean flag whether function null call.
  * `security_type` - Function security type. (Enum: `DEFINER`).
  * `specific_name` - Specific name of the function; Reserved for future use.
  * `external_name` - External function name.
  * `external_language` - External function language.
  * `sql_path` - List of schemes whose objects can be referenced without qualification.
  * `owner` - Username of current owner of function.
  * `comment` - User-provided free-form text description.
  * `properties` - JSON-serialized key-value pair map, encoded (escaped) as a string.
  * `metastore_id` - Unique identifier of parent metastore.
  * `full_name` - Full name of function, in form of catalog_name.schema_name.function__name
  * `created_at` - Time at which this function was created, in epoch milliseconds.
  * `created_by` - Username of function creator.
  * `updated_at` - Time at which this function was created, in epoch milliseconds.
  * `updated_by` - Username of user who last modified function.
  * `function_id` - Id of Function, relative to parent schema.
  * `browse_only` - Indicates whether the principal is limited to retrieving metadata for the associated object through the `BROWSE` privilege when `include_browse` is enabled in the request.

## Related Resources

The following resources are used in the same context:

* [databricks_schema](./schema.md) to get information about a single schema
