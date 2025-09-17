---
subcategory: "Settings"
---
# databricks_account_settings_v2_setting Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to get a single account setting. 

## Example Usage
Referring to a setting by id

```hcl
data "databricks_account_setting" "this" {
    name="llm_proxy_partner_powered"
    boolean_val={
        value=false
    }
}
```

## Arguments
The following arguments are supported:
* `name` (string, required) - Name of the setting

## Attributes
The following attributes are exported:
* `aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy)
* `aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains)
* `automatic_cluster_update_workspace` (ClusterAutoRestartMessage) - todo: Mark these Public after onboarded to DSL
* `boolean_val` (BooleanMessage)
* `default_data_security_mode` (DefaultDataSecurityModeMessage)
* `effective_aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy)
* `effective_aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains)
* `effective_automatic_cluster_update_workspace` (ClusterAutoRestartMessage)
* `effective_boolean_val` (BooleanMessage)
* `effective_default_data_security_mode` (DefaultDataSecurityModeMessage)
* `effective_integer_val` (IntegerMessage)
* `effective_personal_compute` (PersonalComputeMessage)
* `effective_restrict_workspace_admins` (RestrictWorkspaceAdminsMessage)
* `effective_string_val` (StringMessage)
* `integer_val` (IntegerMessage)
* `name` (string) - Name of the setting
* `personal_compute` (PersonalComputeMessage)
* `restrict_workspace_admins` (RestrictWorkspaceAdminsMessage)
* `string_val` (StringMessage)

### AibiDashboardEmbeddingAccessPolicy
* `access_policy_type` (string) - . Possible values are: `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`

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
* `day_of_week` (string) - . Possible values are: `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`
* `frequency` (string) - . Possible values are: `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`
* `window_start_time` (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime)

### ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
* `hours` (integer)
* `minutes` (integer)

### DefaultDataSecurityModeMessage
* `status` (string) - . Possible values are: `NOT_SET`, `SINGLE_USER`, `USER_ISOLATION`

### IntegerMessage
* `value` (integer)

### PersonalComputeMessage
* `value` (string) - . Possible values are: `DELEGATE`, `ON`

### RestrictWorkspaceAdminsMessage
* `status` (string) - . Possible values are: `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`

### StringMessage
* `value` (string) - Represents a generic string value