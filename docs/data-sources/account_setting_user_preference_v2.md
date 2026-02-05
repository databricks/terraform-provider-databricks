---
subcategory: "Settings"
---
# databricks_account_setting_user_preference_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `name` (string, required) - Name of the setting
* `user_id` (string, required) - User ID of the user

## Attributes
The following attributes are exported:
* `boolean_val` (BooleanMessage)
* `effective_boolean_val` (BooleanMessage)
* `effective_string_val` (StringMessage)
* `name` (string) - Name of the setting
* `string_val` (StringMessage)
* `user_id` (string) - User ID of the user

### BooleanMessage
* `value` (boolean)

### StringMessage
* `value` (string) - Represents a generic string value