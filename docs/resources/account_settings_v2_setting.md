---
subcategory: "Settings"
---
# databricks_account_settings_v2_setting Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

Setting is a configurable value or control that determines how a feature or behavior works within the databricks platform.

[//]: # (todo: add public link to metadata api after production doc link available)
See settings-metadata api for list of settings that can be modified using this resource. 

## Example Usage
Getting an account level setting:

```hcl
resource "databricks_account_setting" "this" {
    name="llm_proxy_partner_powered"
    boolean_val={
        value=false
    }
}
```

## Arguments
The following arguments are supported:
* `aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy, optional)
* `aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains, optional)
* `automatic_cluster_update_workspace` (ClusterAutoRestartMessage, optional) - todo: Mark these Public after onboarded to DSL
* `boolean_val` (BooleanMessage, optional)
* `default_data_security_mode` (DefaultDataSecurityModeMessage, optional)
* `effective_aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy, optional)
* `effective_aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains, optional)
* `effective_automatic_cluster_update_workspace` (ClusterAutoRestartMessage, optional)
* `effective_default_data_security_mode` (DefaultDataSecurityModeMessage, optional)
* `effective_personal_compute` (PersonalComputeMessage, optional)
* `effective_restrict_workspace_admins` (RestrictWorkspaceAdminsMessage, optional)
* `integer_val` (IntegerMessage, optional)
* `name` (string, optional) - Name of the setting
* `personal_compute` (PersonalComputeMessage, optional)
* `restrict_workspace_admins` (RestrictWorkspaceAdminsMessage, optional)
* `string_val` (StringMessage, optional)

### AibiDashboardEmbeddingAccessPolicy
* `access_policy_type` (string, required) - . Possible values are: `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`

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
* `day_of_week` (string, optional) - . Possible values are: `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`
* `frequency` (string, optional) - . Possible values are: `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`
* `window_start_time` (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, optional)

### ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
* `hours` (integer, optional)
* `minutes` (integer, optional)

### DefaultDataSecurityModeMessage
* `status` (string, required) - . Possible values are: `NOT_SET`, `SINGLE_USER`, `USER_ISOLATION`

### IntegerMessage
* `value` (integer, optional)

### PersonalComputeMessage
* `value` (string, optional) - . Possible values are: `DELEGATE`, `ON`

### RestrictWorkspaceAdminsMessage
* `status` (string, required) - . Possible values are: `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`

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
  id = "name"
  to = databricks_account_settings_v2_setting.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_account_settings_v2_setting "name"
```