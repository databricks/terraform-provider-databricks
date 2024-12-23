---
subcategory: "Unity Catalog"
---
# databricks_function Resource

-> This resource source can only be used with a workspace-level provider.

Creates a [User-Defined Function (UDF)](https://docs.databricks.com/en/udf/unity-catalog.html) in Unity Catalog. UDFs can be defined using SQL, or external languages (e.g., Python) and are stored within [Unity Catalog schemas](../resources/schema.md). 

## Example Usage

### SQL-based function:

```hcl
resource "databricks_catalog" "sandbox" {
  name    = "sandbox_example"
  comment = "Catalog managed by Terraform"
}

resource "databricks_schema" "functions" {
  catalog_name = databricks_catalog.sandbox.name
  name         = "functions_example"
  comment      = "Schema managed by Terraform"
}

resource "databricks_function" "calculate_bmi" {
  name              = "calculate_bmi"
  catalog_name      = databricks_catalog.sandbox.name
  schema_name       = databricks_schema.functions.name
  input_params      = {
    parameters  = [
        {
          name = "weight"
          type_name = "DOUBLE"
        },
        {
          name = "height"
          type_name = "DOUBLE"
        }
      ]
  }
  data_type         = "DOUBLE"
  routine_body      = "SQL"
  routine_definition = "weight / (height * height)"
  is_deterministic  = true
  sql_data_access   = "CONTAINS_SQL"
  security_type     = "DEFINER"
}
```

### Python-based function: 

```hcl
resource "databricks_function" "calculate_bmi_py" {
  name              = "calculate_bmi_py"
  catalog_name      = databricks_catalog.sandbox.name
  schema_name       = databricks_schema.functions.name
  input_params      = {
   parameters = [
      {
        name = "weight_kg"
        type = "DOUBLE"
      },
      {
        name = "height_m"
        type = "DOUBLE"
      }
    ]
  }
  data_type         = "DOUBLE"
  routine_body      = "EXTERNAL"
  routine_definition = "return weight_kg / (height_m ** 2)"
  language          = "Python"
  is_deterministic  = false
  sql_data_access   = "NO_SQL"
  security_type     = "DEFINER"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the function. 
* `catalog_name` - (Required) The name of the parent [databricks_catalog](../resources/catalog.md). 
* `schema_name` - (Required) The name of [databricks_schema](../resources/schema.md) where the function will reside.
* `input_params` - (Required) A list of objects specifying the input parameters for the function. 
    * `name` - (Required) The name of the parameter.
    * `type` - (Required) The data type of the parameter (e.g., `DOUBLE`, `INT`, etc.).
* `data_type` - (Required) The return data type of the function (e.g., `DOUBLE`).
* `full_data_type` - (Required) Pretty printed function data type (e.g. `string`).
* `return_params` - (Optional) A list of objects specifying the function's return parameters.
    * `parameters` - (Required) An array of objects describing the function's return parameters. Each object includes:
        * `name` - (Required) The name of the return parameter.
        * `type_text` - (Required) The full data type specification as SQL/catalog string text.
        * `type_json` - The full data type specification as JSON-serialized text.
        * `type_name` - (Required) The name of the data type (e.g., `BOOLEAN`, `INT`, `STRING`, etc.).
        * `type_precision` - (Required for `DecimalTypes`) Digits of precision for the type.
        * `type_scale` - (Required for `DecimalTypes`) Digits to the right of the decimal for the type.
        * `type_interval_type` - The format of `IntervalType`.
        * `position` - (Required) The ordinal position of the parameter (starting at 0).
        * `parameter_mode` - The mode of the parameter. Possible value: `IN`.
        * `parameter_type` - The type of the parameter. Possible values:
            * `PARAM` - Represents a generic parameter.
            * `COLUMN` - Represents a column parameter.
        * `parameter_default` - The default value for the parameter, if any.
        * `comment` - User-provided free-form text description of the parameter.
* `routine_definition` - (Required) The actual definition of the function, expressed in SQL or the specified external language.
* `routine_dependencies` - (Optional) A list of objects specifying the function's dependencies. Each object includes:
    * `dependencies` - (Optional) An array of objects describing the dependencies. Each object includes:
        * `table` - (Optional) An object representing a table that is dependent on the SQL object.
        * `function` - (Optional) An object representing a function that is dependent on the SQL object.
* `is_deterministic`- (Required, `bool`) Whether the function is deterministic. Default is `true`.
* `is_null_call` - (Required, `bool`) Indicates whether the function should handle `NULL` input arguments explicitly. 
* `specific_name` - (Required) Specific name of the function. Reserverd for future use.
* `external_name` - (Optional) External function name.
* `sql_path` - (Optional) The fully qualified SQL path where the function resides, including catalog and schema information.
* `comment` - (Optional) User-provided free-form text description.
* `properties` - (Optional) A key-value pair object representing additional metadata or attributes associated with the function.
* `routine_body` - (Required) Specifies the body type of the function, either `SQL` for SQL-based functions or `EXTERNAL` for functions in external languages.
* `security_type` - (Required) The security type of the function, generally `DEFINER`.
* `sql_data_access`- (Required) The SQL data access level for the function. Possible values are: 
    * `CONTAINS_SQL` - The function contains SQL statements.
    * `READS_SQL_DATA` - The function reads SQL data but does not modify it. 
    * `NO_SQL` - The function does not contain SQL.
* `parameter_style` - (Required) Function parameter style (e.g, `S` for SQL).

## Attribute Reference

In addition to all arguments above, the following attributes are exported: 
* `full_name` - Full name of the function in the form of `catalog_name.schema_name.function_name`.
* `created_at` - The time when this function was created, in epoch milliseconds.
* `created_by` - The username of the function's creator. 
* `updated_at` - The time when this function was last updated, in epoch milliseconds.
* `updated_by` - The username of the last user to modify the function.

## Related Resources

The following resources are used in the same context:

* [databricks_schema](./schema.md) to get information about a single schema
*  Data source [databricks_functions](../data-sources/functions.md) to get a list of functions under a specified location. 
