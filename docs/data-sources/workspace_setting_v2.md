---
subcategory: "Settings"
---
# databricks_workspace_setting_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to get a single account setting. 

## Example Usage
Referring to a setting by id
```hcl
data "databricks_workspace_setting_v2" "this" {
  name = "llm_proxy_partner_powered"
  boolean_val = {
    value = false
  }
}
```

## Arguments
The following arguments are supported:
* `name` (string, required) - Name of the setting

## Attributes
The following attributes are exported:
* `aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy) - Setting value for aibi_dashboard_embedding_access_policy setting. This is the setting value set by consumers, check effective_aibi_dashboard_embedding_access_policy for final setting value
* `aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains) - Setting value for aibi_dashboard_embedding_approved_domains setting. This is the setting value set by consumers, check effective_aibi_dashboard_embedding_approved_domains for final setting value
* `automatic_cluster_update_workspace` (ClusterAutoRestartMessage) - Setting value for automatic_cluster_update_workspace setting. This is the setting value set by consumers, check effective_automatic_cluster_update_workspace for final setting value
* `boolean_val` (BooleanMessage) - Setting value for boolean type setting. This is the setting value set by consumers, check effective_boolean_val for final setting value
* `effective_aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy) - Effective setting value for aibi_dashboard_embedding_access_policy setting. This is the final effective value of setting. To set a value use aibi_dashboard_embedding_access_policy
* `effective_aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains) - Effective setting value for aibi_dashboard_embedding_approved_domains setting. This is the final effective value of setting. To set a value use aibi_dashboard_embedding_approved_domains
* `effective_automatic_cluster_update_workspace` (ClusterAutoRestartMessage) - Effective setting value for automatic_cluster_update_workspace setting. This is the final effective value of setting. To set a value use automatic_cluster_update_workspace
* `effective_boolean_val` (BooleanMessage) - Effective setting value for boolean type setting. This is the final effective value of setting. To set a value use boolean_val
* `effective_integer_val` (IntegerMessage) - Effective setting value for integer type setting. This is the final effective value of setting. To set a value use integer_val
* `effective_personal_compute` (PersonalComputeMessage) - Effective setting value for personal_compute setting. This is the final effective value of setting. To set a value use personal_compute
* `effective_restrict_workspace_admins` (RestrictWorkspaceAdminsMessage) - Effective setting value for restrict_workspace_admins setting. This is the final effective value of setting. To set a value use restrict_workspace_admins
* `effective_string_val` (StringMessage) - Effective setting value for string type setting. This is the final effective value of setting. To set a value use string_val
* `integer_val` (IntegerMessage) - Setting value for integer type setting. This is the setting value set by consumers, check effective_integer_val for final setting value
* `name` (string) - Name of the setting
* `personal_compute` (PersonalComputeMessage) - Setting value for personal_compute setting. This is the setting value set by consumers, check effective_personal_compute for final setting value
* `restrict_workspace_admins` (RestrictWorkspaceAdminsMessage) - Setting value for restrict_workspace_admins setting. This is the setting value set by consumers, check effective_restrict_workspace_admins for final setting value
* `string_val` (StringMessage) - Setting value for string type setting. This is the setting value set by consumers, check effective_string_val for final setting value

### AibiDashboardEmbeddingAccessPolicy
* `access_policy_type` (string) - Possible values are: `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`

### AibiDashboardEmbeddingApprovedDomains
* `approved_domains` (list of string)

### BooleanMessage
* `value` (boolean)

### ClusterAutoRestartMessage
* `can_toggle` (boolean)
* `enabled` (boolean)
* `enablement_details` (ClusterAutoRestartMessageEnablementDetails)
* `maintenance_window` (ClusterAutoRestartMessageMaintenanceWindow)
* `restart_even_if_no_updates_available` (boolean)

### ClusterAutoRestartMessageEnablementDetails
* `forced_for_compliance_mode` (boolean) - The feature is force enabled if compliance mode is active
* `unavailable_for_disabled_entitlement` (boolean) - The feature is unavailable if the corresponding entitlement disabled (see getShieldEntitlementEnable)
* `unavailable_for_non_enterprise_tier` (boolean) - The feature is unavailable if the customer doesn't have enterprise tier

### ClusterAutoRestartMessageMaintenanceWindow
* `week_day_based_schedule` (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule)

### ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
* `day_of_week` (string) - Possible values are: `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`
* `frequency` (string) - Possible values are: `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`
* `window_start_time` (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime)

### ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
* `hours` (integer)
* `minutes` (integer)

### IntegerMessage
* `value` (integer)

### PersonalComputeMessage
* `value` (string) - Possible values are: `DELEGATE`, `ON`

### RestrictWorkspaceAdminsMessage
* `status` (string) - Possible values are: `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`

### StringMessage
* `value` (string) - Represents a generic string value
