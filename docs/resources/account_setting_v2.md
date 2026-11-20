---
subcategory: "Settings"
---
# databricks_account_setting_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/settingsv2)

Setting is a configurable value or control that determines how a feature or behavior works within the databricks platform.

[//]: # (todo: add public link to metadata api after production doc link available)
See settings-metadata api for list of settings that can be modified using this resource. 

## Example Usage
Getting an account level setting:

```hcl
resource "databricks_account_setting_v2" "this" {
    name="llm_proxy_partner_powered"
    boolean_val={
        value=false
    }
}
```

## Arguments
The following arguments are supported:
* `aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy, optional) - Setting value for aibi_dashboard_embedding_access_policy setting. This is the setting value set by consumers, check effective_aibi_dashboard_embedding_access_policy for final setting value
* `aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains, optional) - Setting value for aibi_dashboard_embedding_approved_domains setting. This is the setting value set by consumers, check effective_aibi_dashboard_embedding_approved_domains for final setting value
* `allowed_apps_user_api_scopes` (AllowedAppsUserApiScopesMessage, optional) - Setting value for allowed_apps_user_api_scopes setting. This is the setting value set by consumers, check effective_allowed_apps_user_api_scopes for final setting value
* `automatic_cluster_update_workspace` (ClusterAutoRestartMessage, optional) - Setting value for automatic_cluster_update_workspace setting. This is the setting value set by consumers, check effective_automatic_cluster_update_workspace for final setting value
* `boolean_val` (BooleanMessage, optional) - Setting value for boolean type setting. This is the setting value set by consumers, check effective_boolean_val for final setting value
* `collaboration_platform_connectivity` (CollaborationPlatformConnectivityMessage, optional) - Setting value for collaboration_platform_connectivity setting. This is the setting value set by consumers, check effective_collaboration_platform_connectivity for final setting value
* `effective_aibi_dashboard_embedding_access_policy` (AibiDashboardEmbeddingAccessPolicy, optional) - Effective setting value for aibi_dashboard_embedding_access_policy setting. This is the final effective value of setting. To set a value use aibi_dashboard_embedding_access_policy
* `effective_aibi_dashboard_embedding_approved_domains` (AibiDashboardEmbeddingApprovedDomains, optional) - Effective setting value for aibi_dashboard_embedding_approved_domains setting. This is the final effective value of setting. To set a value use aibi_dashboard_embedding_approved_domains
* `effective_automatic_cluster_update_workspace` (ClusterAutoRestartMessage, optional) - Effective setting value for automatic_cluster_update_workspace setting. This is the final effective value of setting. To set a value use automatic_cluster_update_workspace
* `effective_personal_compute` (PersonalComputeMessage, optional) - Effective setting value for personal_compute setting. This is the final effective value of setting. To set a value use personal_compute
* `effective_restrict_workspace_admins` (RestrictWorkspaceAdminsMessage, optional) - Effective setting value for restrict_workspace_admins setting. This is the final effective value of setting. To set a value use restrict_workspace_admins
* `effective_workspace_label` (WorkspaceLabelMessage, optional) - Effective setting value for workspace_label setting. This is the final effective value of setting. To set a value use workspace_label
* `integer_val` (IntegerMessage, optional) - Setting value for integer type setting. This is the setting value set by consumers, check effective_integer_val for final setting value
* `name` (string, optional) - Name of the setting. The settings metadata endpoint lists the names available at each scope
* `operational_email_custom_recipient` (OperationalEmailCustomRecipientMessage, optional) - Setting value for operational_email_custom_recipient setting. This is the setting value set by consumers, check effective_operational_email_custom_recipient for final setting value
* `personal_compute` (PersonalComputeMessage, optional) - Setting value for personal_compute setting. This is the setting value set by consumers, check effective_personal_compute for final setting value
* `restrict_workspace_admins` (RestrictWorkspaceAdminsMessage, optional) - Setting value for restrict_workspace_admins setting. This is the setting value set by consumers, check effective_restrict_workspace_admins for final setting value
* `string_val` (StringMessage, optional) - Setting value for string type setting. This is the setting value set by consumers, check effective_string_val for final setting value
* `workspace_dashboard_theme` (DashboardTheme, optional) - Setting value for the workspace_dashboard_theme setting. This is the setting value set by consumers. Check effective_workspace_dashboard_theme for the final setting value
* `workspace_label` (WorkspaceLabelMessage, optional) - Setting value for workspace_label setting. This is the setting value set by consumers, check effective_workspace_label for final setting value

### AibiDashboardEmbeddingAccessPolicy
* `access_policy_type` (string, required) - Where AI/BI dashboards can be embedded: in any domain (ALLOW_ALL_DOMAINS), only in the approved
  domains configured for the workspace (ALLOW_APPROVED_DOMAINS), or nowhere (DENY_ALL_DOMAINS).
  Defaults to ALLOW_APPROVED_DOMAINS. Possible values are: `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`

### AibiDashboardEmbeddingApprovedDomains
* `approved_domains` (list of string, optional) - Domains allowed to host embedded AI/BI dashboards. A leading wildcard allows every subdomain of a
  domain. The list applies, and can only be changed, while the embedding access policy is
  ALLOW_APPROVED_DOMAINS

### AllowedAppsUserApiScopesMessage
* `allowed_scopes` (list of string, optional) - OAuth scopes that app developers can request when an app acts on behalf of a user. A list with
  the single entry * allows every supported scope and cannot be combined with other scopes. An
  empty list disables user authorization for apps

### BooleanMessage
* `value` (boolean, optional)

### ClusterAutoRestartMessage
* `can_toggle` (boolean, optional) - Whether the workspace can change the enabled state. Databricks computes this value: it is false
  when the compliance security profile forces the feature on or when the workspace lacks the
  Enhanced Security and Compliance entitlement. A value sent in an update request has no effect
* `enabled` (boolean, optional) - Whether automatic cluster update is enabled for the workspace. The effective value is always
  true while the compliance security profile is enabled on the workspace
* `enablement_details` (ClusterAutoRestartMessageEnablementDetails, optional) - Why the feature is forced on or unavailable. Databricks sets these details for information only;
  they do not determine whether the feature is enabled
* `maintenance_window` (ClusterAutoRestartMessageMaintenanceWindow, optional) - When the maintenance window runs. When unset, it defaults to the first Sunday of every month at
  01:00 UTC
* `restart_even_if_no_updates_available` (boolean, optional) - Whether compute resources restart during every maintenance window even when no update is
  available. When false, they restart only when an update is available

### ClusterAutoRestartMessageEnablementDetails
* `forced_for_compliance_mode` (boolean, optional) - The feature is force enabled if compliance mode is active
* `unavailable_for_disabled_entitlement` (boolean, optional) - The feature is unavailable because the workspace does not have the Enhanced Security and
  Compliance entitlement
* `unavailable_for_non_enterprise_tier` (boolean, optional) - The feature is unavailable if the customer doesn't have enterprise tier

### ClusterAutoRestartMessageMaintenanceWindow
* `week_day_based_schedule` (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, optional) - Schedule anchored to a day of the week, either every week or on given weeks of the month

### ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
* `day_of_week` (string, optional) - Day of the week on which the maintenance window starts. Possible values are: `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`
* `frequency` (string, optional) - How often the maintenance window occurs: every week, or only on the listed occurrences of
  the day of the week within each month. Possible values are: `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`
* `window_start_time` (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, optional) - Time of day, in UTC, at which the maintenance window starts

### ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
* `hours` (integer, optional) - Hour of the day on a 24-hour clock
* `minutes` (integer, optional) - Minute of the hour

### CollaborationPlatformConnectivityMessage
* `connectivity` (string, required) - Possible values are: `ALLOW_ALL`, `ALLOW_SLACK`, `ALLOW_TEAMS`, `DENY_ALL`

### DashboardTheme
* `axis_line_color` (DashboardThemeDashboardThemeColor, optional) - Color of the zero-value grid line, making it visually distinct from other grid lines
* `canvas_background_color` (DashboardThemeDashboardThemeColor, optional) - The background color of the dashboard canvas
* `continuous_gradient` (DashboardThemeContinuousGradient, optional) - The continuous gradient configuration for the dashboard
* `font_color` (DashboardThemeDashboardThemeColor, optional) - The default font color applied to all text in the dashboard.
  Superseded by font_settings.base.font_color when font_settings is set.
  More specific font settings in turn override base
* `font_family` (string, optional) - Font family to customize all text within the dashboard.
  Unsupported font names fall back to the default font.
  Superseded by font_settings.base.font_family when font_settings is set
* `font_settings` (DashboardThemeFontSettings, optional) - The per-category typography overrides for the dashboard
* `grid_line_color` (DashboardThemeDashboardThemeColor, optional) - Color of chart grid lines in visualizations (horizontal and vertical guide lines)
* `mark_corner_radius` (integer, optional) - The corner radius in pixels for rectangular marks in visualizations (bars, columns, histogram bins, etc.)
* `selection_color` (DashboardThemeDashboardThemeColor, optional) - The highlight color for selected elements in widgets
* `sentiment_colors` (DashboardThemeSentimentColors, optional) - The sentiment color palette for the dashboard
* `text_box_vertical_align` (string, optional) - The default vertical alignment for all text box widgets in the dashboard.
  Individual text box widgets can override this. Possible values are: `VERTICAL_ALIGN_BOTTOM`, `VERTICAL_ALIGN_CENTER`, `VERTICAL_ALIGN_TOP`
* `visualization_color_01` (string, optional) - Color palette for visualizations. Colors are applied in order (visualization_color_01, visualization_color_02, etc.).
  Invalid color values are preserved but ignored during rendering
* `visualization_color_02` (string, optional) - The color at position 2 in the visualization color palette
* `visualization_color_03` (string, optional) - The color at position 3 in the visualization color palette
* `visualization_color_04` (string, optional) - The color at position 4 in the visualization color palette
* `visualization_color_05` (string, optional) - The color at position 5 in the visualization color palette
* `visualization_color_06` (string, optional) - The color at position 6 in the visualization color palette
* `visualization_color_07` (string, optional) - The color at position 7 in the visualization color palette
* `visualization_color_08` (string, optional) - The color at position 8 in the visualization color palette
* `visualization_color_09` (string, optional) - The color at position 9 in the visualization color palette
* `visualization_color_10` (string, optional) - The color at position 10 in the visualization color palette
* `visualization_color_11` (string, optional) - The color at position 11 in the visualization color palette
* `visualization_color_12` (string, optional) - The color at position 12 in the visualization color palette
* `visualization_color_13` (string, optional) - The color at position 13 in the visualization color palette
* `visualization_color_14` (string, optional) - The color at position 14 in the visualization color palette
* `visualization_color_15` (string, optional) - The color at position 15 in the visualization color palette
* `visualization_color_16` (string, optional) - The color at position 16 in the visualization color palette
* `visualization_color_17` (string, optional) - The color at position 17 in the visualization color palette
* `visualization_color_18` (string, optional) - The color at position 18 in the visualization color palette
* `visualization_color_19` (string, optional) - The color at position 19 in the visualization color palette
* `visualization_color_20` (string, optional) - The color at position 20 in the visualization color palette
* `widget_background_color` (DashboardThemeDashboardThemeColor, optional) - The background color of widgets
* `widget_border_color` (DashboardThemeDashboardThemeColor, optional) - The border color of widgets
* `widget_corner_radius` (integer, optional) - The widget border radius in pixels
* `widget_header_alignment` (string, optional) - Alignment setting for all widget titles and descriptions at the dashboard level. Possible values are: `CENTER`, `LEFT`, `RIGHT`
* `widget_margin` (integer, optional) - The gap between widgets in pixels
* `widget_padding` (integer, optional) - The content padding inside each widget in pixels
* `widget_shadow` (integer, optional) - The shadow intensity level applied to widgets (0 = no shadow, 100 = maximum).
  The actual CSS box-shadow is computed at render time from this value, factoring in
  background colors and color mode

### DashboardThemeContinuousGradient
* `dark_mode` (DashboardThemeContinuousGradientGradientValue, optional) - Gradient value used when the dashboard is rendered in dark mode
* `light_mode` (DashboardThemeContinuousGradientGradientValue, optional) - Gradient value used when the dashboard is rendered in light mode

### DashboardThemeContinuousGradientCustomDivergingRamp
* `end` (DashboardThemeDashboardThemeColor, optional) - The color at the end of the gradient (highest value)
* `mid` (DashboardThemeDashboardThemeColor, optional) - The color at the midpoint of the gradient
* `start` (DashboardThemeDashboardThemeColor, optional) - The color at the start of the gradient (lowest value)

### DashboardThemeContinuousGradientCustomSequentialRamp
* `end` (DashboardThemeDashboardThemeColor, optional) - The color at the end of the gradient (highest value)
* `start` (DashboardThemeDashboardThemeColor, optional) - The color at the start of the gradient (lowest value)

### DashboardThemeContinuousGradientGradientValue
* `color` (string, optional) - A static color string (for example, '#FF0000')
* `color_scheme` (string, optional) - A named color scheme (for example, 'blues', 'viridis', 'spectral')
* `custom_diverging` (DashboardThemeContinuousGradientCustomDivergingRamp, optional) - A custom three-stop diverging gradient
* `custom_sequential` (DashboardThemeContinuousGradientCustomSequentialRamp, optional) - A custom two-stop sequential gradient
* `from_categorical_index` (integer, optional) - A 0-based index into the visualization color palette, where 0 is visualization_color_01. Valid values are 0 to 19

### DashboardThemeDashboardThemeColor
* `dark` (string, optional) - The color used when the dashboard is rendered in dark mode (for example, '#FF0000')
* `light` (string, optional) - The color used when the dashboard is rendered in light mode (for example, '#FF0000')

### DashboardThemeFontConfig
* `font_color` (DashboardThemeDashboardThemeColor, optional) - The font color applied to text in this category
* `font_family` (string, optional) - The font family name used for text in this category
* `font_size` (integer, optional) - The font size in pixels
* `font_style` (string, optional) - The font style (slant) applied to text in this category. Possible values are: `FONT_STYLE_ITALIC`, `FONT_STYLE_NORMAL`
* `font_weight` (integer, optional) - The font weight value (for example, 400 for normal, 700 for bold)
* `text_case` (string, optional) - The text case transformation applied to text in this category. Possible values are: `TEXT_CASE_CAPITALIZE`, `TEXT_CASE_LOWERCASE`, `TEXT_CASE_UPPERCASE`
* `text_decoration` (string, optional) - The text decoration applied to text in this category. Possible values are: `TEXT_DECORATION_NONE`, `TEXT_DECORATION_UNDERLINE`

### DashboardThemeFontSettings
* `base` (DashboardThemeFontConfig, optional) - The default font config applied to all text unless overridden by a more specific category
* `field_title` (DashboardThemeFontConfig, optional) - The font config applied to field titles (for example, axis titles, table heading titles)
* `field_value` (DashboardThemeFontConfig, optional) - The font config applied to field values (for example, axis tick labels, table cell values, tooltip values)
* `widget_description` (DashboardThemeFontConfig, optional) - The font config applied to widget descriptions
* `widget_title` (DashboardThemeFontConfig, optional) - The font config applied to widget titles

### DashboardThemeSentimentColors
* `negative` (DashboardThemeDashboardThemeColor, optional) - The color applied to negative sentiment values
* `neutral` (DashboardThemeDashboardThemeColor, optional) - The color applied to neutral sentiment values
* `positive` (DashboardThemeDashboardThemeColor, optional) - The color applied to positive sentiment values

### IntegerMessage
* `value` (integer, optional)

### OperationalEmailCustomRecipientMessage
* `email` (string, optional) - Email address that receives this workspace's operational emails in addition to the workspace
  admins. Use a distribution list address to reach several people. An empty string clears the
  additional recipient

### PersonalComputeMessage
* `value` (string, optional) - ON grants all users in all workspaces of the account access to the Personal Compute default
  policy. DELEGATE moves control to each workspace, where users or groups must be added to the
  policy's permissions before they can create compute with it. Possible values are: `DELEGATE`, `ON`

### RestrictWorkspaceAdminsMessage
* `status` (string, required) - Controls what workspace admins can do with service principal tokens and job identities. With
  ALLOW_ALL, workspace admins can create personal access tokens on behalf of any service principal
  in the workspace, change a job owner to any workspace user, and set a job's run_as identity to
  any workspace user or to a service principal they have the Service Principal User role on. With
  RESTRICT_TOKENS_AND_JOB_RUN_AS, they can only create tokens for service principals they have the
  Service Principal User role on, change a job owner to themselves, and set run_as to themselves
  or to a service principal they have that role on. Possible values are: `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`
* `disable_gov_tag_creation` (boolean, optional) - When true, workspace admins cannot create governance tags.
  ALLOW_ALL status does not override this; they are independent

### StringMessage
* `value` (string, optional) - Represents a generic string value

### WorkspaceLabelMessage
* `color` (string, optional) - Color of the label badge. The supported colors are coral, lemon, lime, turquoise, indigo, purple
  and pink. It is an empty string when no label is set
* `label` (string, optional) - Text of the label shown next to the workspace name in the top navigation bar. An empty string
  means no label is set

## Attributes
In addition to the above arguments, the following attributes are exported:
* `effective_allowed_apps_user_api_scopes` (AllowedAppsUserApiScopesMessage) - Effective setting value for allowed_apps_user_api_scopes setting. This is the final effective value of setting. To set a value use allowed_apps_user_api_scopes
* `effective_boolean_val` (BooleanMessage) - Effective setting value for boolean type setting. This is the final effective value of setting. To set a value use boolean_val
* `effective_collaboration_platform_connectivity` (CollaborationPlatformConnectivityMessage) - Effective setting value for collaboration_platform_connectivity setting. This is the final effective value of setting. To set a value use collaboration_platform_connectivity
* `effective_integer_val` (IntegerMessage) - Effective setting value for integer type setting. This is the final effective value of setting. To set a value use integer_val
* `effective_operational_email_custom_recipient` (OperationalEmailCustomRecipientMessage) - Effective setting value for operational_email_custom_recipient setting. This is the final effective value of setting. To set a value use operational_email_custom_recipient
* `effective_string_val` (StringMessage) - Effective setting value for string type setting. This is the final effective value of setting. To set a value use string_val
* `effective_workspace_dashboard_theme` (DashboardTheme) - Effective setting value for the workspace_dashboard_theme setting. This is the final effective value of the setting. To set a value, use workspace_dashboard_theme

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