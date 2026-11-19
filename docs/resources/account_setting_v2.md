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
* `name` (string, optional) - Name of the setting
* `operational_email_custom_recipient` (OperationalEmailCustomRecipientMessage, optional) - Setting value for operational_email_custom_recipient setting. This is the setting value set by consumers, check effective_operational_email_custom_recipient for final setting value
* `personal_compute` (PersonalComputeMessage, optional) - Setting value for personal_compute setting. This is the setting value set by consumers, check effective_personal_compute for final setting value
* `restrict_workspace_admins` (RestrictWorkspaceAdminsMessage, optional) - Setting value for restrict_workspace_admins setting. This is the setting value set by consumers, check effective_restrict_workspace_admins for final setting value
* `string_val` (StringMessage, optional) - Setting value for string type setting. This is the setting value set by consumers, check effective_string_val for final setting value
* `workspace_dashboard_theme` (DashboardTheme, optional) - Setting value for the workspace_dashboard_theme setting. This is the setting value set by consumers. Check effective_workspace_dashboard_theme for the final setting value
* `workspace_label` (WorkspaceLabelMessage, optional) - Setting value for workspace_label setting. This is the setting value set by consumers, check effective_workspace_label for final setting value

### AibiDashboardEmbeddingAccessPolicy
* `access_policy_type` (string, required) - Possible values are: `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`

### AibiDashboardEmbeddingApprovedDomains
* `approved_domains` (list of string, optional)

### AllowedAppsUserApiScopesMessage
* `allowed_scopes` (list of string, optional)

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
* `email` (string, optional)

### PersonalComputeMessage
* `value` (string, optional) - Possible values are: `DELEGATE`, `ON`

### RestrictWorkspaceAdminsMessage
* `status` (string, required) - Possible values are: `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`
* `disable_gov_tag_creation` (boolean, optional) - When true, workspace admins cannot create governance tags.
  ALLOW_ALL status does not override this; they are independent

### StringMessage
* `value` (string, optional) - Represents a generic string value

### WorkspaceLabelMessage
* `color` (string, optional)
* `label` (string, optional)

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