---
subcategory: "Settings"
---
# databricks_account_setting_user_preference_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/settingsv2)

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
* `boolean_val` (BooleanMessage, optional) - Holds the value that the user set for a boolean preference. Each preference uses either
  `boolean_val` or `string_val`, depending on its type. Responses omit the user-set value when
  the user has not set one
* `name` (string, optional) - Identifies the user preference. The "List user preferences and their metadata" method returns
  the valid names
* `string_val` (StringMessage, optional) - Holds the value that the user set for a string preference
* `user_id` (string, optional) - Identifies the user that the preference belongs to. The value is the numeric account user ID
  that the Account Users API (https://docs.databricks.com/api/account/accountusersv2/list) returns

### BooleanMessage
* `value` (boolean, optional)

### StringMessage
* `value` (string, optional) - Represents a generic string value

## Attributes
In addition to the above arguments, the following attributes are exported:
* `effective_boolean_val` (BooleanMessage) - Holds the value that takes effect for a boolean preference. The server evaluates it from the
  user-set value and the preference's policy, so it can differ from `boolean_val`. When the user
  has not set a value, it reflects the preference's default
* `effective_string_val` (StringMessage) - Holds the value that takes effect for a string preference, as evaluated by the server

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