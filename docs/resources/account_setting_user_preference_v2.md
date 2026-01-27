---
subcategory: "Settings"
---
# databricks_account_setting_user_preference_v2 Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

User preference is a configurable value that determines how a feature or behavior works for a specific user within the Databricks platform.

See user settings-metadata API for list of user preferences that can be modified using this resource.


## Example Usage
Setting an account user preference:

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
    }
  }
}

provider "databricks" {
  profile = "ACCOUNT-<account-id>"
}

resource "databricks_account_user_setting_v2" "theme_setting" {
  user_id = "<user-id>"
  name    = "enableDarkMode"
  string_val = {
    value = "dark"
  }
}
```

Setting a boolean user preference:

```hcl
resource "databricks_account_user_setting_v2" "enable_line_numbers" {
  user_id = "<user-id>"
  name    = "enableLineNumbers"
  boolean_val = {
    value = true
  }
}
```


## Arguments
The following arguments are supported:
* `boolean_val` (BooleanMessage, optional)
* `name` (string, optional) - Name of the setting
* `string_val` (StringMessage, optional)
* `user_id` (string, optional) - User ID of the user

### BooleanMessage
* `value` (boolean, optional)

### StringMessage
* `value` (string, optional) - Represents a generic string value

## Attributes
In addition to the above arguments, the following attributes are exported:
* `effective_boolean_val` (BooleanMessage)
* `effective_string_val` (StringMessage)

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name,user_id"
  to = databricks_account_setting_user_preference_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_setting_user_preference_v2.this "name,user_id"
```