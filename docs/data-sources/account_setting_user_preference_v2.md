---
subcategory: "Settings"
---
# databricks_account_setting_user_preference_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/settingsv2)

This data source can be used to get a single account user preference setting.


## Example Usage
Referring to a user preference by name:

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

data "databricks_account_user_setting_v2" "theme_setting" {
  user_id = "<user-id>"
  name    = "enableDarkMode"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Identifies the user preference. The "List user preferences and their metadata" method returns
  the valid names
* `user_id` (string, required) - Identifies the user that the preference belongs to. The value is the numeric account user ID
  that the Account Users API (https://docs.databricks.com/api/account/accountusersv2/list) returns

## Attributes
The following attributes are exported:
* `boolean_val` (BooleanMessage) - Holds the value that the user set for a boolean preference. Each preference uses either
  `boolean_val` or `string_val`, depending on its type. Responses omit the user-set value when
  the user has not set one
* `effective_boolean_val` (BooleanMessage) - Holds the value that takes effect for a boolean preference. The server evaluates it from the
  user-set value and the preference's policy, so it can differ from `boolean_val`. When the user
  has not set a value, it reflects the preference's default
* `effective_string_val` (StringMessage) - Holds the value that takes effect for a string preference, as evaluated by the server
* `name` (string) - Identifies the user preference. The "List user preferences and their metadata" method returns
  the valid names
* `string_val` (StringMessage) - Holds the value that the user set for a string preference
* `user_id` (string) - Identifies the user that the preference belongs to. The value is the numeric account user ID
  that the Account Users API (https://docs.databricks.com/api/account/accountusersv2/list) returns

### BooleanMessage
* `value` (boolean)

### StringMessage
* `value` (string) - Represents a generic string value