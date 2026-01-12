---
subcategory: "Settings"
---
# databricks_account_setting_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

Setting is a configurable value or control that determines how a feature or behavior works within the databricks platform.

[//]: # (todo: add public link to metadata api after production doc link available)
See settings-metadata api for list of settings that can be modified using this resource. 

## Example Usage
Getting an account level setting:

```hcl
resource "databricks_account_setting_v2" "this" {
  name = "llm_proxy_partner_powered"
  boolean_val = {
    value = false
  }
}
```

## Arguments
The following arguments are supported:
* `aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy, optional) - Setting value for aibi_dashboard_embedding_access_policy setting. This is the setting value set by consumers, check effective_aibi_dashboard_embedding_access_policy for final setting value
* `aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains, optional) - Setting value for aibi_dashboard_embedding_approved_domains setting. This is the setting value set by consumers, check effective_aibi_dashboard_embedding_approved_domains for final setting value
* `automatic_cluster_update_workspace` (ClusterAutoRestartMessage, optional) - Setting value for automatic_cluster_update_workspace setting. This is the setting value set by consumers, check effective_automatic_cluster_update_workspace for final setting value
* `boolean_val` (BooleanMessage, optional) - Setting value for boolean type setting. This is the setting value set by consumers, check effective_boolean_val for final setting value
* `effective_aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy, optional) - Effective setting value for aibi_dashboard_embedding_access_policy setting. This is the final effective value of setting. To set a value use aibi_dashboard_embedding_access_policy
* `effective_aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains, optional) - Effective setting value for aibi_dashboard_embedding_approved_domains setting. This is the final effective value of setting. To set a value use aibi_dashboard_embedding_approved_domains
* `effective_automatic_cluster_update_workspace` (ClusterAutoRestartMessage, optional) - Effective setting value for automatic_cluster_update_workspace setting. This is the final effective value of setting. To set a value use automatic_cluster_update_workspace
* `effective_personal_compute` (PersonalComputeMessage, optional) - Effective setting value for personal_compute setting. This is the final effective value of setting. To set a value use personal_compute
* `effective_restrict_workspace_admins` (RestrictWorkspaceAdminsMessage, optional) - Effective setting value for restrict_workspace_admins setting. This is the final effective value of setting. To set a value use restrict_workspace_admins
* `integer_val` (IntegerMessage, optional) - Setting value for integer type setting. This is the setting value set by consumers, check effective_integer_val for final setting value
* `name` (string, optional) - Name of the setting
* `personal_compute` (PersonalComputeMessage, optional) - Setting value for personal_compute setting. This is the setting value set by consumers, check effective_personal_compute for final setting value
* `restrict_workspace_admins` (RestrictWorkspaceAdminsMessage, optional) - Setting value for restrict_workspace_admins setting. This is the setting value set by consumers, check effective_restrict_workspace_admins for final setting value
* `string_val` (StringMessage, optional) - Setting value for string type setting. This is the setting value set by consumers, check effective_string_val for final setting value

### AibiDashboardEmbeddingAccessPolicy
* `access_policy_type` (string, required) - Possible values are: `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`

### AibiDashboardEmbeddingApprovedDomains
* `approved_domains` (list of string, optional)

### BooleanMessage
* `value` (boolean, optional)

### ClusterAutoRestartMessage
* `can_toggle` (boolean, optional)
* `enabled` (boolean, optional)
* `enablement_details` (ClusterAutoRestartMessageEnablementDetails, optional)
* `maintenance_window` (ClusterAutoRestartMessageMaintenanceWindow, optional)
* `restart_even_if_no_updates_available` (boolean, optional)

### ClusterAutoRestartMessageEnablementDetails
* `forced_for_compliance_mode` (boolean, optional) - The feature is force enabled if compliance mode is active
* `unavailable_for_disabled_entitlement` (boolean, optional) - The feature is unavailable if the corresponding entitlement disabled (see getShieldEntitlementEnable)
* `unavailable_for_non_enterprise_tier` (boolean, optional) - The feature is unavailable if the customer doesn't have enterprise tier

### ClusterAutoRestartMessageMaintenanceWindow
* `week_day_based_schedule` (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, optional)

### ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
* `day_of_week` (string, optional) - Possible values are: `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`
* `frequency` (string, optional) - Possible values are: `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`
* `window_start_time` (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, optional)

### ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
* `hours` (integer, optional)
* `minutes` (integer, optional)

### IntegerMessage
* `value` (integer, optional)

### PersonalComputeMessage
* `value` (string, optional) - Possible values are: `DELEGATE`, `ON`

### RestrictWorkspaceAdminsMessage
* `status` (string, required) - Possible values are: `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`

### StringMessage
* `value` (string, optional) - Represents a generic string value

## Attributes
In addition to the above arguments, the following attributes are exported:
* `effective_boolean_val` (BooleanMessage) - Effective setting value for boolean type setting. This is the final effective value of setting. To set a value use boolean_val
* `effective_integer_val` (IntegerMessage) - Effective setting value for integer type setting. This is the final effective value of setting. To set a value use integer_val
* `effective_string_val` (StringMessage) - Effective setting value for string type setting. This is the final effective value of setting. To set a value use string_val

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_account_setting_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_setting_v2.this "name"
```
