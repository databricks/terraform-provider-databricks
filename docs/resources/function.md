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
  input_params      = [
    {
      name = "weight"
      type = "DOUBLE"
    },
    {
      name = "height"
      type = "DOUBLE"
    }
  ]
  data_type         = "DOUBLE"
  routine_body      = "SQL"
  routine_definition = "weight / (height * height)"
  language          = "SQL"
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
  input_params      = [
    {
      name = "weight_kg"
      type = "DOUBLE"
    },
    {
      name = "height_m"
      type = "DOUBLE"
    }
  ]
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
* `routine_body` - (Required) Specifies the body type of the function, either `SQL` for SQL-based functions or `EXTERNAL` for functions in external languages.
* `routine_definition` - (Required) The actual definition of the function, expressed in SQL or the specified external language.
* `language` - (Required) The language of the function, e.g., `SQL` or `Python`. 
* `is_deterministic`- (Optional, `bool`) Whether the function is deterministic. Default is `true`.
* `sql_data_Access`- (Optional) The SQL data access level for the function. Possible values are: 
    * `CONTAINS_SQL` - The function contains SQL statements.
    * `READS_SQL_DATA` - The function reads SQL data but does not modify it. 
    * `NO_SQL` - The function does not contain SQL.
* `security_type` - (Optional) The security type of the function, generally `DEFINER`.

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
