---
subcategory: "Development and Testing"
---
# databricks_workspace_setting Resource


## Example Usage


## Arguments
The following arguments are supported:
* `boolean_val` (BooleanMessage, optional)
* `integer_val` (IntegerMessage, optional)
* `name` (string, optional) - Name of the setting
* `string_val` (StringMessage, optional)

### BooleanMessage
* `value` (boolean, optional)

### IntegerMessage
* `value` (integer, optional)

### StringMessage
* `value` (string, optional) - Represents a generic string value



## Attributes
In addition to the above arguments, the following attributes are exported:
* `effective_boolean_val` (BooleanMessage)
* `effective_integer_val` (IntegerMessage)
* `effective_string_val` (StringMessage)

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = name
  to = databricks_workspace_setting.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_setting name
```