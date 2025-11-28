---
subcategory: "Unity Catalog"
---
# databricks_policy_info Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

Attribute-Based Access Control (ABAC) policies in Unity Catalog provide high leverage governance for enforcing compliance policies. With ABAC policies, access is controlled in a hierarchical and scalable manner, based on data attributes rather than specific resources, enabling more flexible and comprehensive access control.

ABAC policies in Unity Catalog support conditions on governance tags and the user identity. Callers must have the `MANAGE` privilege on a securable to view, create, update, or delete ABAC policies.

### Policy Components

ABAC policies consist of:
- **Conditions**: Define when the policy applies based on governance tags and the user identity
- **Actions**: What operations to be taken when policy condition meets. Supported actions include applying a row filter or a column mask to a table
- **Scope**: The securable hierarchy (a catalog, a schema or a table) in which the policy could take effect.

### Supported Securables

ABAC policies can be applied to the following securable types:
- Catalogs
- Schemas
- Tables



## Example Usage
### Row Filter Policy

```hcl
resource "databricks_policy_info" "pii_row_filter" {
  on_securable_type     = "catalog"
  on_securable_fullname = "main"
  name                  = "pii_data_policy"
  
  policy_type           = "POLICY_TYPE_ROW_FILTER"
  for_securable_type    = "table"
  to_principals         = ["account users"]
  
  # Condition for when the policy applies
  when_condition = "hasTag('pii')"
  
  # Match specific columns
  match_columns = [
    {
      condition = "hasTag('pii')"
      alias     = "pii_col"
    }
  ]
  
  # Row filter function to apply
  row_filter = {
    function_name = "main.filters.mask_pii_rows"
    using = [
      {
        alias = "pii_col"
      }
    ]
  }
}
```

### Column Mask Policy

```hcl
resource "databricks_policy_info" "sensitive_column_mask" {
  on_securable_type     = "schema"
  on_securable_fullname = "main.finance"
  name                  = "sensitive_data_mask"
  
  policy_type           = "POLICY_TYPE_COLUMN_MASK"
  for_securable_type    = "table"
  to_principals         = ["account users"]
  except_principals     = ["finance_admins"]
  
  # Condition for when the policy applies
  when_condition = "hasTag('pii')"
  
  # Match columns to mask
  match_columns = [
    {
      condition = "hasTag('pii')"
      alias     = "sensitive_col"
    }
  ]
  
  # Column mask function to apply
  column_mask = {
    function_name = "main.masks.redact_sensitive"
    on_column     = "sensitive_col"
    using = [
      {
        constant = "4"
      }
    ]
  }
}
```



## Arguments
The following arguments are supported:
* `for_securable_type` (string, required) - Type of securables that the policy should take effect on.
  Only `TABLE` is supported at this moment.
  Required on create and optional on update. Possible values are: `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `EXTERNAL_METADATA`, `FUNCTION`, `METASTORE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STAGING_TABLE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`
* `policy_type` (string, required) - Type of the policy. Required on create and ignored on update. Possible values are: `POLICY_TYPE_COLUMN_MASK`, `POLICY_TYPE_ROW_FILTER`
* `to_principals` (list of string, required) - List of user or group names that the policy applies to.
  Required on create and optional on update
* `column_mask` (ColumnMaskOptions, optional) - Options for column mask policies. Valid only if `policy_type` is `POLICY_TYPE_COLUMN_MASK`.
  Required on create and optional on update. When specified on update,
  the new options will replace the existing options as a whole
* `comment` (string, optional) - Optional description of the policy
* `except_principals` (list of string, optional) - Optional list of user or group names that should be excluded from the policy
* `match_columns` (list of MatchColumn, optional) - Optional list of condition expressions used to match table columns.
  Only valid when `for_securable_type` is `TABLE`.
  When specified, the policy only applies to tables whose columns satisfy all match conditions
* `name` (string, optional) - Name of the policy. Required on create and optional on update.
  To rename the policy, set `name` to a different value on update
* `on_securable_fullname` (string, optional) - Full name of the securable on which the policy is defined.
  Required on create and ignored on update
* `on_securable_type` (string, optional) - Type of the securable on which the policy is defined.
  Only `CATALOG`, `SCHEMA` and `TABLE` are supported at this moment.
  Required on create and ignored on update. Possible values are: `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `EXTERNAL_METADATA`, `FUNCTION`, `METASTORE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STAGING_TABLE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`
* `row_filter` (RowFilterOptions, optional) - Options for row filter policies. Valid only if `policy_type` is `POLICY_TYPE_ROW_FILTER`.
  Required on create and optional on update. When specified on update,
  the new options will replace the existing options as a whole
* `when_condition` (string, optional) - Optional condition when the policy should take effect

### ColumnMaskOptions
* `function_name` (string, required) - The fully qualified name of the column mask function.
  The function is called on each row of the target table.
  The function's first argument and its return type should match the type of the masked column.
  Required on create and update
* `on_column` (string, required) - The alias of the column to be masked. The alias must refer to one of matched columns.
  The values of the column is passed to the column mask function as the first argument.
  Required on create and update
* `using` (list of FunctionArgument, optional) - Optional list of column aliases or constant literals to be passed as additional arguments to the column mask function.
  The type of each column should match the positional argument of the column mask function

### FunctionArgument
* `alias` (string, optional) - The alias of a matched column
* `constant` (string, optional) - A constant literal

### MatchColumn
* `alias` (string, optional) - Optional alias of the matched column
* `condition` (string, optional) - The condition expression used to match a table column

### RowFilterOptions
* `function_name` (string, required) - The fully qualified name of the row filter function.
  The function is called on each row of the target table. It should return a boolean value
  indicating whether the row should be visible to the user.
  Required on create and update
* `using` (list of FunctionArgument, optional) - Optional list of column aliases or constant literals to be passed as arguments to the row filter function.
  The type of each column should match the positional argument of the row filter function

## Attributes
In addition to the above arguments, the following attributes are exported:
* `created_at` (integer) - Time at which the policy was created, in epoch milliseconds. Output only
* `created_by` (string) - Username of the user who created the policy. Output only
* `id` (string) - Unique identifier of the policy. This field is output only and is generated by the system
* `updated_at` (integer) - Time at which the policy was last modified, in epoch milliseconds. Output only
* `updated_by` (string) - Username of the user who last modified the policy. Output only

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "on_securable_type,on_securable_fullname,name"
  to = databricks_policy_info.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_policy_info.this "on_securable_type,on_securable_fullname,name"
```