// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AibiDashboardEmbeddingAccessPolicy struct {
	// Where AI/BI dashboards can be embedded: in any domain
	// (ALLOW_ALL_DOMAINS), only in the approved domains configured for the
	// workspace (ALLOW_APPROVED_DOMAINS), or nowhere (DENY_ALL_DOMAINS).
	// Defaults to ALLOW_APPROVED_DOMAINS.
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType `json:"access_policy_type"`
}

func (s *AibiDashboardEmbeddingAccessPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AibiDashboardEmbeddingAccessPolicyAccessPolicyType string

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `ALLOW_ALL_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `ALLOW_APPROVED_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeDenyAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `DENY_ALL_DOMAINS`

// String representation for [fmt.Print]
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Set(v string) error {
	switch v {
	case `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`:
		*f = AibiDashboardEmbeddingAccessPolicyAccessPolicyType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL_DOMAINS", "ALLOW_APPROVED_DOMAINS", "DENY_ALL_DOMAINS"`, v)
	}
}

// Values returns all possible values for AibiDashboardEmbeddingAccessPolicyAccessPolicyType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Values() []AibiDashboardEmbeddingAccessPolicyAccessPolicyType {
	return []AibiDashboardEmbeddingAccessPolicyAccessPolicyType{
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowAllDomains,
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains,
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeDenyAllDomains,
	}
}

// Type always returns AibiDashboardEmbeddingAccessPolicyAccessPolicyType to satisfy [pflag.Value] interface
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Type() string {
	return "AibiDashboardEmbeddingAccessPolicyAccessPolicyType"
}

type AibiDashboardEmbeddingApprovedDomains struct {
	// Domains allowed to host embedded AI/BI dashboards. A leading wildcard
	// allows every subdomain of a domain. The list applies, and can only be
	// changed, while the embedding access policy is ALLOW_APPROVED_DOMAINS.
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

func (s *AibiDashboardEmbeddingApprovedDomains) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AllowedAppsUserApiScopesMessage struct {
	// OAuth scopes that app developers can request when an app acts on behalf
	// of a user. A list with the single entry * allows every supported scope
	// and cannot be combined with other scopes. An empty list disables user
	// authorization for apps.
	AllowedScopes []string `json:"allowed_scopes,omitempty"`
}

func (s *AllowedAppsUserApiScopesMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BooleanMessage struct {
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BooleanMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BooleanMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAutoRestartMessage struct {
	// Whether the workspace can change the enabled state. Databricks computes
	// this value: it is false when the compliance security profile forces the
	// feature on or when the workspace lacks the Enhanced Security and
	// Compliance entitlement. A value sent in an update request has no effect.
	CanToggle bool `json:"can_toggle,omitempty"`
	// Whether automatic cluster update is enabled for the workspace. The
	// effective value is always true while the compliance security profile is
	// enabled on the workspace.
	Enabled bool `json:"enabled,omitempty"`
	// Why the feature is forced on or unavailable. Databricks sets these
	// details for information only; they do not determine whether the feature
	// is enabled.
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails `json:"enablement_details,omitempty"`
	// When the maintenance window runs. When unset, it defaults to the first
	// Sunday of every month at 01:00 UTC.
	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow `json:"maintenance_window,omitempty"`
	// Whether compute resources restart during every maintenance window even
	// when no update is available. When false, they restart only when an update
	// is available.
	RestartEvenIfNoUpdatesAvailable bool `json:"restart_even_if_no_updates_available,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode bool `json:"forced_for_compliance_mode,omitempty"`
	// The feature is unavailable because the workspace does not have the
	// Enhanced Security and Compliance entitlement.
	UnavailableForDisabledEntitlement bool `json:"unavailable_for_disabled_entitlement,omitempty"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier bool `json:"unavailable_for_non_enterprise_tier,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessageEnablementDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageEnablementDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	// Schedule anchored to a day of the week, either every week or on given
	// weeks of the month.
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `json:"week_day_based_schedule,omitempty"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ClusterAutoRestartMessageMaintenanceWindowDayOfWeek string

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFriday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `FRIDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekMonday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `MONDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSaturday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `SATURDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `SUNDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekThursday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `THURSDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekTuesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `TUESDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekWednesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `WEDNESDAY`

// String representation for [fmt.Print]
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Set(v string) error {
	switch v {
	case `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = ClusterAutoRestartMessageMaintenanceWindowDayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Values returns all possible values for ClusterAutoRestartMessageMaintenanceWindowDayOfWeek.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Values() []ClusterAutoRestartMessageMaintenanceWindowDayOfWeek {
	return []ClusterAutoRestartMessageMaintenanceWindowDayOfWeek{
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFriday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekMonday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSaturday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekThursday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekTuesday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekWednesday,
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowDayOfWeek to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowDayOfWeek"
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	// Day of the week on which the maintenance window starts.
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek `json:"day_of_week,omitempty"`
	// How often the maintenance window occurs: every week, or only on the
	// listed occurrences of the day of the week within each month.
	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `json:"frequency,omitempty"`
	// Time of day, in UTC, at which the maintenance window starts.
	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `json:"window_start_time,omitempty"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency string

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `EVERY_WEEK`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_AND_THIRD_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_AND_FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `THIRD_OF_MONTH`

// String representation for [fmt.Print]
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Set(v string) error {
	switch v {
	case `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`:
		*f = ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EVERY_WEEK", "FIRST_AND_THIRD_OF_MONTH", "FIRST_OF_MONTH", "FOURTH_OF_MONTH", "SECOND_AND_FOURTH_OF_MONTH", "SECOND_OF_MONTH", "THIRD_OF_MONTH"`, v)
	}
}

// Values returns all possible values for ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Values() []ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency {
	return []ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency{
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth,
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency"
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	// Hour of the day on a 24-hour clock.
	Hours int `json:"hours,omitempty"`
	// Minute of the hour.
	Minutes int `json:"minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Controls which external collaboration platforms (Slack, Microsoft Teams) can
// connect to a workspace. Defaults to ALLOW_ALL.
type CollaborationPlatformConnectivityMessage struct {
	Connectivity CollaborationPlatformConnectivityMessageConnectivity `json:"connectivity"`
}

func (s *CollaborationPlatformConnectivityMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CollaborationPlatformConnectivityMessageConnectivity string

const CollaborationPlatformConnectivityMessageConnectivityAllowAll CollaborationPlatformConnectivityMessageConnectivity = `ALLOW_ALL`

const CollaborationPlatformConnectivityMessageConnectivityAllowSlack CollaborationPlatformConnectivityMessageConnectivity = `ALLOW_SLACK`

const CollaborationPlatformConnectivityMessageConnectivityAllowTeams CollaborationPlatformConnectivityMessageConnectivity = `ALLOW_TEAMS`

const CollaborationPlatformConnectivityMessageConnectivityDenyAll CollaborationPlatformConnectivityMessageConnectivity = `DENY_ALL`

// String representation for [fmt.Print]
func (f *CollaborationPlatformConnectivityMessageConnectivity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CollaborationPlatformConnectivityMessageConnectivity) Set(v string) error {
	switch v {
	case `ALLOW_ALL`, `ALLOW_SLACK`, `ALLOW_TEAMS`, `DENY_ALL`:
		*f = CollaborationPlatformConnectivityMessageConnectivity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL", "ALLOW_SLACK", "ALLOW_TEAMS", "DENY_ALL"`, v)
	}
}

// Values returns all possible values for CollaborationPlatformConnectivityMessageConnectivity.
//
// There is no guarantee on the order of the values in the slice.
func (f *CollaborationPlatformConnectivityMessageConnectivity) Values() []CollaborationPlatformConnectivityMessageConnectivity {
	return []CollaborationPlatformConnectivityMessageConnectivity{
		CollaborationPlatformConnectivityMessageConnectivityAllowAll,
		CollaborationPlatformConnectivityMessageConnectivityAllowSlack,
		CollaborationPlatformConnectivityMessageConnectivityAllowTeams,
		CollaborationPlatformConnectivityMessageConnectivityDenyAll,
	}
}

// Type always returns CollaborationPlatformConnectivityMessageConnectivity to satisfy [pflag.Value] interface
func (f *CollaborationPlatformConnectivityMessageConnectivity) Type() string {
	return "CollaborationPlatformConnectivityMessageConnectivity"
}

// Theme settings for customizing AI/BI dashboard UI, including visualization
// colors, font colors, background colors, and more.
type DashboardTheme struct {
	// Color of the zero-value grid line, making it visually distinct from other
	// grid lines.
	AxisLineColor *DashboardThemeDashboardThemeColor `json:"axis_line_color,omitempty"`
	// The background color of the dashboard canvas.
	CanvasBackgroundColor *DashboardThemeDashboardThemeColor `json:"canvas_background_color,omitempty"`
	// The continuous gradient configuration for the dashboard.
	ContinuousGradient *DashboardThemeContinuousGradient `json:"continuous_gradient,omitempty"`
	// The default font color applied to all text in the dashboard. Superseded
	// by font_settings.base.font_color when font_settings is set. More specific
	// font settings in turn override base.
	FontColor *DashboardThemeDashboardThemeColor `json:"font_color,omitempty"`
	// Font family to customize all text within the dashboard. Unsupported font
	// names fall back to the default font. Superseded by
	// font_settings.base.font_family when font_settings is set.
	FontFamily string `json:"font_family,omitempty"`
	// The per-category typography overrides for the dashboard.
	FontSettings *DashboardThemeFontSettings `json:"font_settings,omitempty"`
	// Color of chart grid lines in visualizations (horizontal and vertical
	// guide lines).
	GridLineColor *DashboardThemeDashboardThemeColor `json:"grid_line_color,omitempty"`
	// The corner radius in pixels for rectangular marks in visualizations
	// (bars, columns, histogram bins, etc.).
	MarkCornerRadius int `json:"mark_corner_radius,omitempty"`
	// The highlight color for selected elements in widgets.
	SelectionColor *DashboardThemeDashboardThemeColor `json:"selection_color,omitempty"`
	// The sentiment color palette for the dashboard.
	SentimentColors *DashboardThemeSentimentColors `json:"sentiment_colors,omitempty"`
	// The default vertical alignment for all text box widgets in the dashboard.
	// Individual text box widgets can override this.
	TextBoxVerticalAlign DashboardThemeVerticalAlign `json:"text_box_vertical_align,omitempty"`
	// Color palette for visualizations. Colors are applied in order
	// (visualization_color_01, visualization_color_02, etc.). Invalid color
	// values are preserved but ignored during rendering.
	VisualizationColor01 string `json:"visualization_color_01,omitempty"`
	// The color at position 2 in the visualization color palette.
	VisualizationColor02 string `json:"visualization_color_02,omitempty"`
	// The color at position 3 in the visualization color palette.
	VisualizationColor03 string `json:"visualization_color_03,omitempty"`
	// The color at position 4 in the visualization color palette.
	VisualizationColor04 string `json:"visualization_color_04,omitempty"`
	// The color at position 5 in the visualization color palette.
	VisualizationColor05 string `json:"visualization_color_05,omitempty"`
	// The color at position 6 in the visualization color palette.
	VisualizationColor06 string `json:"visualization_color_06,omitempty"`
	// The color at position 7 in the visualization color palette.
	VisualizationColor07 string `json:"visualization_color_07,omitempty"`
	// The color at position 8 in the visualization color palette.
	VisualizationColor08 string `json:"visualization_color_08,omitempty"`
	// The color at position 9 in the visualization color palette.
	VisualizationColor09 string `json:"visualization_color_09,omitempty"`
	// The color at position 10 in the visualization color palette.
	VisualizationColor10 string `json:"visualization_color_10,omitempty"`
	// The color at position 11 in the visualization color palette.
	VisualizationColor11 string `json:"visualization_color_11,omitempty"`
	// The color at position 12 in the visualization color palette.
	VisualizationColor12 string `json:"visualization_color_12,omitempty"`
	// The color at position 13 in the visualization color palette.
	VisualizationColor13 string `json:"visualization_color_13,omitempty"`
	// The color at position 14 in the visualization color palette.
	VisualizationColor14 string `json:"visualization_color_14,omitempty"`
	// The color at position 15 in the visualization color palette.
	VisualizationColor15 string `json:"visualization_color_15,omitempty"`
	// The color at position 16 in the visualization color palette.
	VisualizationColor16 string `json:"visualization_color_16,omitempty"`
	// The color at position 17 in the visualization color palette.
	VisualizationColor17 string `json:"visualization_color_17,omitempty"`
	// The color at position 18 in the visualization color palette.
	VisualizationColor18 string `json:"visualization_color_18,omitempty"`
	// The color at position 19 in the visualization color palette.
	VisualizationColor19 string `json:"visualization_color_19,omitempty"`
	// The color at position 20 in the visualization color palette.
	VisualizationColor20 string `json:"visualization_color_20,omitempty"`
	// The background color of widgets.
	WidgetBackgroundColor *DashboardThemeDashboardThemeColor `json:"widget_background_color,omitempty"`
	// The border color of widgets.
	WidgetBorderColor *DashboardThemeDashboardThemeColor `json:"widget_border_color,omitempty"`
	// The widget border radius in pixels.
	WidgetCornerRadius int `json:"widget_corner_radius,omitempty"`
	// Alignment setting for all widget titles and descriptions at the dashboard
	// level.
	WidgetHeaderAlignment DashboardThemeAlignment `json:"widget_header_alignment,omitempty"`
	// The gap between widgets in pixels.
	WidgetMargin int `json:"widget_margin,omitempty"`
	// The content padding inside each widget in pixels.
	WidgetPadding int `json:"widget_padding,omitempty"`
	// The shadow intensity level applied to widgets (0 = no shadow, 100 =
	// maximum). The actual CSS box-shadow is computed at render time from this
	// value, factoring in background colors and color mode.
	WidgetShadow int `json:"widget_shadow,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardTheme) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardTheme) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Horizontal alignment options for widget titles and descriptions.
type DashboardThemeAlignment string

const DashboardThemeAlignmentCenter DashboardThemeAlignment = `CENTER`

const DashboardThemeAlignmentLeft DashboardThemeAlignment = `LEFT`

const DashboardThemeAlignmentRight DashboardThemeAlignment = `RIGHT`

// String representation for [fmt.Print]
func (f *DashboardThemeAlignment) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardThemeAlignment) Set(v string) error {
	switch v {
	case `CENTER`, `LEFT`, `RIGHT`:
		*f = DashboardThemeAlignment(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CENTER", "LEFT", "RIGHT"`, v)
	}
}

// Values returns all possible values for DashboardThemeAlignment.
//
// There is no guarantee on the order of the values in the slice.
func (f *DashboardThemeAlignment) Values() []DashboardThemeAlignment {
	return []DashboardThemeAlignment{
		DashboardThemeAlignmentCenter,
		DashboardThemeAlignmentLeft,
		DashboardThemeAlignmentRight,
	}
}

// Type always returns DashboardThemeAlignment to satisfy [pflag.Value] interface
func (f *DashboardThemeAlignment) Type() string {
	return "DashboardThemeAlignment"
}

// Configuration for continuous color gradients, with separate values for light
// and dark mode. Each mode's gradient can reference the visualization color
// palette via from_categorical_index, use a named color scheme, or define a
// custom sequential or diverging ramp.
type DashboardThemeContinuousGradient struct {
	// Gradient value used when the dashboard is rendered in dark mode.
	DarkMode *DashboardThemeContinuousGradientGradientValue `json:"dark_mode,omitempty"`
	// Gradient value used when the dashboard is rendered in light mode.
	LightMode *DashboardThemeContinuousGradientGradientValue `json:"light_mode,omitempty"`
}

func (s *DashboardThemeContinuousGradient) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A custom three-stop diverging gradient ramp.
type DashboardThemeContinuousGradientCustomDivergingRamp struct {
	// The color at the end of the gradient (highest value).
	End *DashboardThemeDashboardThemeColor `json:"end,omitempty"`
	// The color at the midpoint of the gradient.
	Mid *DashboardThemeDashboardThemeColor `json:"mid,omitempty"`
	// The color at the start of the gradient (lowest value).
	Start *DashboardThemeDashboardThemeColor `json:"start,omitempty"`
}

func (s *DashboardThemeContinuousGradientCustomDivergingRamp) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A custom two-stop sequential gradient ramp.
type DashboardThemeContinuousGradientCustomSequentialRamp struct {
	// The color at the end of the gradient (highest value).
	End *DashboardThemeDashboardThemeColor `json:"end,omitempty"`
	// The color at the start of the gradient (lowest value).
	Start *DashboardThemeDashboardThemeColor `json:"start,omitempty"`
}

func (s *DashboardThemeContinuousGradientCustomSequentialRamp) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The gradient for one color mode. Set at most one field.
type DashboardThemeContinuousGradientGradientValue struct {
	// A static color string (for example, '#FF0000').
	Color string `json:"color,omitempty"`
	// A named color scheme (for example, 'blues', 'viridis', 'spectral').
	ColorScheme string `json:"color_scheme,omitempty"`
	// A custom three-stop diverging gradient.
	CustomDiverging *DashboardThemeContinuousGradientCustomDivergingRamp `json:"custom_diverging,omitempty"`
	// A custom two-stop sequential gradient.
	CustomSequential *DashboardThemeContinuousGradientCustomSequentialRamp `json:"custom_sequential,omitempty"`
	// A 0-based index into the visualization color palette, where 0 is
	// visualization_color_01. Valid values are 0 to 19.
	FromCategoricalIndex int `json:"from_categorical_index,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardThemeContinuousGradientGradientValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardThemeContinuousGradientGradientValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents a color value with optional overrides for light and dark modes.
// Authors can customize colors separately to cater to each mode, based on the
// viewer's user-level theme preference.
type DashboardThemeDashboardThemeColor struct {
	// The color used when the dashboard is rendered in dark mode (for example,
	// '#FF0000').
	Dark string `json:"dark,omitempty"`
	// The color used when the dashboard is rendered in light mode (for example,
	// '#FF0000').
	Light string `json:"light,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardThemeDashboardThemeColor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardThemeDashboardThemeColor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Defines the typography settings for a single text category.
type DashboardThemeFontConfig struct {
	// The font color applied to text in this category.
	FontColor *DashboardThemeDashboardThemeColor `json:"font_color,omitempty"`
	// The font family name used for text in this category.
	FontFamily string `json:"font_family,omitempty"`
	// The font size in pixels.
	FontSize int `json:"font_size,omitempty"`
	// The font style (slant) applied to text in this category.
	FontStyle DashboardThemeFontStyle `json:"font_style,omitempty"`
	// The font weight value (for example, 400 for normal, 700 for bold).
	FontWeight int `json:"font_weight,omitempty"`
	// The text case transformation applied to text in this category.
	TextCase DashboardThemeTextCase `json:"text_case,omitempty"`
	// The text decoration applied to text in this category.
	TextDecoration DashboardThemeTextDecoration `json:"text_decoration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardThemeFontConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardThemeFontConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Per-category typography overrides. When font_settings.base is set, it takes
// precedence over the top-level font_family and font_color. More specific
// categories (widget_title, widget_description, field_title, field_value) in
// turn override base.
type DashboardThemeFontSettings struct {
	// The default font config applied to all text unless overridden by a more
	// specific category.
	Base *DashboardThemeFontConfig `json:"base,omitempty"`
	// The font config applied to field titles (for example, axis titles, table
	// heading titles).
	FieldTitle *DashboardThemeFontConfig `json:"field_title,omitempty"`
	// The font config applied to field values (for example, axis tick labels,
	// table cell values, tooltip values).
	FieldValue *DashboardThemeFontConfig `json:"field_value,omitempty"`
	// The font config applied to widget descriptions.
	WidgetDescription *DashboardThemeFontConfig `json:"widget_description,omitempty"`
	// The font config applied to widget titles.
	WidgetTitle *DashboardThemeFontConfig `json:"widget_title,omitempty"`
}

func (s *DashboardThemeFontSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Defines the font style (slant) applied to rendered text.
type DashboardThemeFontStyle string

const DashboardThemeFontStyleFontStyleItalic DashboardThemeFontStyle = `FONT_STYLE_ITALIC`

const DashboardThemeFontStyleFontStyleNormal DashboardThemeFontStyle = `FONT_STYLE_NORMAL`

// String representation for [fmt.Print]
func (f *DashboardThemeFontStyle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardThemeFontStyle) Set(v string) error {
	switch v {
	case `FONT_STYLE_ITALIC`, `FONT_STYLE_NORMAL`:
		*f = DashboardThemeFontStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FONT_STYLE_ITALIC", "FONT_STYLE_NORMAL"`, v)
	}
}

// Values returns all possible values for DashboardThemeFontStyle.
//
// There is no guarantee on the order of the values in the slice.
func (f *DashboardThemeFontStyle) Values() []DashboardThemeFontStyle {
	return []DashboardThemeFontStyle{
		DashboardThemeFontStyleFontStyleItalic,
		DashboardThemeFontStyleFontStyleNormal,
	}
}

// Type always returns DashboardThemeFontStyle to satisfy [pflag.Value] interface
func (f *DashboardThemeFontStyle) Type() string {
	return "DashboardThemeFontStyle"
}

// Shared sentiment colors reusable across visualizations, primarily for
// conditional formatting, ensuring consistent positive, negative, and neutral
// meaning throughout the dashboard.
type DashboardThemeSentimentColors struct {
	// The color applied to negative sentiment values.
	Negative *DashboardThemeDashboardThemeColor `json:"negative,omitempty"`
	// The color applied to neutral sentiment values.
	Neutral *DashboardThemeDashboardThemeColor `json:"neutral,omitempty"`
	// The color applied to positive sentiment values.
	Positive *DashboardThemeDashboardThemeColor `json:"positive,omitempty"`
}

func (s *DashboardThemeSentimentColors) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Defines the text case transformation applied to rendered text.
type DashboardThemeTextCase string

const DashboardThemeTextCaseTextCaseCapitalize DashboardThemeTextCase = `TEXT_CASE_CAPITALIZE`

const DashboardThemeTextCaseTextCaseLowercase DashboardThemeTextCase = `TEXT_CASE_LOWERCASE`

const DashboardThemeTextCaseTextCaseUppercase DashboardThemeTextCase = `TEXT_CASE_UPPERCASE`

// String representation for [fmt.Print]
func (f *DashboardThemeTextCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardThemeTextCase) Set(v string) error {
	switch v {
	case `TEXT_CASE_CAPITALIZE`, `TEXT_CASE_LOWERCASE`, `TEXT_CASE_UPPERCASE`:
		*f = DashboardThemeTextCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TEXT_CASE_CAPITALIZE", "TEXT_CASE_LOWERCASE", "TEXT_CASE_UPPERCASE"`, v)
	}
}

// Values returns all possible values for DashboardThemeTextCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *DashboardThemeTextCase) Values() []DashboardThemeTextCase {
	return []DashboardThemeTextCase{
		DashboardThemeTextCaseTextCaseCapitalize,
		DashboardThemeTextCaseTextCaseLowercase,
		DashboardThemeTextCaseTextCaseUppercase,
	}
}

// Type always returns DashboardThemeTextCase to satisfy [pflag.Value] interface
func (f *DashboardThemeTextCase) Type() string {
	return "DashboardThemeTextCase"
}

// Defines the decoration line applied to rendered text.
type DashboardThemeTextDecoration string

const DashboardThemeTextDecorationTextDecorationNone DashboardThemeTextDecoration = `TEXT_DECORATION_NONE`

const DashboardThemeTextDecorationTextDecorationUnderline DashboardThemeTextDecoration = `TEXT_DECORATION_UNDERLINE`

// String representation for [fmt.Print]
func (f *DashboardThemeTextDecoration) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardThemeTextDecoration) Set(v string) error {
	switch v {
	case `TEXT_DECORATION_NONE`, `TEXT_DECORATION_UNDERLINE`:
		*f = DashboardThemeTextDecoration(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TEXT_DECORATION_NONE", "TEXT_DECORATION_UNDERLINE"`, v)
	}
}

// Values returns all possible values for DashboardThemeTextDecoration.
//
// There is no guarantee on the order of the values in the slice.
func (f *DashboardThemeTextDecoration) Values() []DashboardThemeTextDecoration {
	return []DashboardThemeTextDecoration{
		DashboardThemeTextDecorationTextDecorationNone,
		DashboardThemeTextDecorationTextDecorationUnderline,
	}
}

// Type always returns DashboardThemeTextDecoration to satisfy [pflag.Value] interface
func (f *DashboardThemeTextDecoration) Type() string {
	return "DashboardThemeTextDecoration"
}

// Vertical alignment options for text box content.
type DashboardThemeVerticalAlign string

const DashboardThemeVerticalAlignVerticalAlignBottom DashboardThemeVerticalAlign = `VERTICAL_ALIGN_BOTTOM`

const DashboardThemeVerticalAlignVerticalAlignCenter DashboardThemeVerticalAlign = `VERTICAL_ALIGN_CENTER`

const DashboardThemeVerticalAlignVerticalAlignTop DashboardThemeVerticalAlign = `VERTICAL_ALIGN_TOP`

// String representation for [fmt.Print]
func (f *DashboardThemeVerticalAlign) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardThemeVerticalAlign) Set(v string) error {
	switch v {
	case `VERTICAL_ALIGN_BOTTOM`, `VERTICAL_ALIGN_CENTER`, `VERTICAL_ALIGN_TOP`:
		*f = DashboardThemeVerticalAlign(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "VERTICAL_ALIGN_BOTTOM", "VERTICAL_ALIGN_CENTER", "VERTICAL_ALIGN_TOP"`, v)
	}
}

// Values returns all possible values for DashboardThemeVerticalAlign.
//
// There is no guarantee on the order of the values in the slice.
func (f *DashboardThemeVerticalAlign) Values() []DashboardThemeVerticalAlign {
	return []DashboardThemeVerticalAlign{
		DashboardThemeVerticalAlignVerticalAlignBottom,
		DashboardThemeVerticalAlignVerticalAlignCenter,
		DashboardThemeVerticalAlignVerticalAlignTop,
	}
}

// Type always returns DashboardThemeVerticalAlign to satisfy [pflag.Value] interface
func (f *DashboardThemeVerticalAlign) Type() string {
	return "DashboardThemeVerticalAlign"
}

type GetPublicAccountSettingRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetPublicAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPublicAccountUserPreferenceRequest struct {
	// User Setting name.
	Name string `json:"-" url:"-"`
	// User ID of the user whose setting is being retrieved.
	UserId string `json:"-" url:"-"`
}

func (s *GetPublicAccountUserPreferenceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPublicWorkspaceSettingRequest struct {
	// Name of the setting to retrieve.
	Name string `json:"-" url:"-"`
}

func (s *GetPublicWorkspaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type IntegerMessage struct {
	Value int `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IntegerMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IntegerMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountSettingsMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListAccountSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListAccountSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountSettingsMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountSettingsMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountSettingsMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountSettingsMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountUserPreferencesMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListAccountUserPreferencesMetadataRequest` call. Provide this to
	// retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListAccountUserPreferencesMetadataRequest` must match the call that
	// provided the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// User ID of the user whose settings metadata is being retrieved.
	UserId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountUserPreferencesMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountUserPreferencesMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountUserPreferencesMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountUserPreferencesMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountUserPreferencesMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceSettingsMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListWorkspaceSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListWorkspaceSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceSettingsMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceSettingsMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceSettingsMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceSettingsMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type OperationalEmailCustomRecipientMessage struct {
	// Email address that receives this workspace's operational emails in
	// addition to the workspace admins. Use a distribution list address to
	// reach several people. An empty string clears the additional recipient.
	Email string `json:"email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OperationalEmailCustomRecipientMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OperationalEmailCustomRecipientMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchPublicAccountSettingRequest struct {
	Name string `json:"-" url:"-"`

	Setting Setting `json:"setting"`
}

func (s *PatchPublicAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PatchPublicAccountUserPreferenceRequest struct {
	Name string `json:"-" url:"-"`

	Setting UserPreference `json:"setting"`
	// User ID of the user whose setting is being updated.
	UserId string `json:"-" url:"-"`
}

func (s *PatchPublicAccountUserPreferenceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PatchPublicWorkspaceSettingRequest struct {
	// Name of the setting to update.
	Name string `json:"-" url:"-"`

	Setting Setting `json:"setting"`
}

func (s *PatchPublicWorkspaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PersonalComputeMessage struct {
	// ON grants all users in all workspaces of the account access to the
	// Personal Compute default policy. DELEGATE moves control to each
	// workspace, where users or groups must be added to the policy's
	// permissions before they can create compute with it.
	Value PersonalComputeMessagePersonalComputeMessageEnum `json:"value,omitempty"`
}

func (s *PersonalComputeMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// ON: Grants all users in all workspaces access to the Personal Compute default
// policy, allowing all users to create single-machine compute resources.
// DELEGATE: Moves access control for the Personal Compute default policy to
// individual workspaces and requires a workspace’s users or groups to be
// added to the ACLs of that workspace’s Personal Compute default policy
// before they will be able to create compute resources through that policy.
type PersonalComputeMessagePersonalComputeMessageEnum string

const PersonalComputeMessagePersonalComputeMessageEnumDelegate PersonalComputeMessagePersonalComputeMessageEnum = `DELEGATE`

const PersonalComputeMessagePersonalComputeMessageEnumOn PersonalComputeMessagePersonalComputeMessageEnum = `ON`

// String representation for [fmt.Print]
func (f *PersonalComputeMessagePersonalComputeMessageEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalComputeMessagePersonalComputeMessageEnum) Set(v string) error {
	switch v {
	case `DELEGATE`, `ON`:
		*f = PersonalComputeMessagePersonalComputeMessageEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELEGATE", "ON"`, v)
	}
}

// Values returns all possible values for PersonalComputeMessagePersonalComputeMessageEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *PersonalComputeMessagePersonalComputeMessageEnum) Values() []PersonalComputeMessagePersonalComputeMessageEnum {
	return []PersonalComputeMessagePersonalComputeMessageEnum{
		PersonalComputeMessagePersonalComputeMessageEnumDelegate,
		PersonalComputeMessagePersonalComputeMessageEnumOn,
	}
}

// Type always returns PersonalComputeMessagePersonalComputeMessageEnum to satisfy [pflag.Value] interface
func (f *PersonalComputeMessagePersonalComputeMessageEnum) Type() string {
	return "PersonalComputeMessagePersonalComputeMessageEnum"
}

// Preview phase for settings that are feature previews. For settings that are
// not feature previews, the preview_phase field is left unset. Mirrors only the
// customer-facing phases surfaced in the UI; internal-only phases (DISABLED,
// DEV, UNDER_MIGRATION, LAUNCHED, etc.) are not exposed here.
type PreviewPhase string

const PreviewPhaseBeta PreviewPhase = `BETA`

const PreviewPhaseGa PreviewPhase = `GA`

const PreviewPhaseGaSoon PreviewPhase = `GA_SOON`

const PreviewPhasePrivatePreview PreviewPhase = `PRIVATE_PREVIEW`

const PreviewPhasePublicPreview PreviewPhase = `PUBLIC_PREVIEW`

// String representation for [fmt.Print]
func (f *PreviewPhase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PreviewPhase) Set(v string) error {
	switch v {
	case `BETA`, `GA`, `GA_SOON`, `PRIVATE_PREVIEW`, `PUBLIC_PREVIEW`:
		*f = PreviewPhase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BETA", "GA", "GA_SOON", "PRIVATE_PREVIEW", "PUBLIC_PREVIEW"`, v)
	}
}

// Values returns all possible values for PreviewPhase.
//
// There is no guarantee on the order of the values in the slice.
func (f *PreviewPhase) Values() []PreviewPhase {
	return []PreviewPhase{
		PreviewPhaseBeta,
		PreviewPhaseGa,
		PreviewPhaseGaSoon,
		PreviewPhasePrivatePreview,
		PreviewPhasePublicPreview,
	}
}

// Type always returns PreviewPhase to satisfy [pflag.Value] interface
func (f *PreviewPhase) Type() string {
	return "PreviewPhase"
}

type RestrictWorkspaceAdminsMessage struct {
	// When true, workspace admins cannot create governance tags. ALLOW_ALL
	// status does not override this; they are independent.
	DisableGovTagCreation bool `json:"disable_gov_tag_creation,omitempty"`
	// Controls what workspace admins can do with service principal tokens and
	// job identities. With ALLOW_ALL, workspace admins can create personal
	// access tokens on behalf of any service principal in the workspace, change
	// a job owner to any workspace user, and set a job's run_as identity to any
	// workspace user or to a service principal they have the Service Principal
	// User role on. With RESTRICT_TOKENS_AND_JOB_RUN_AS, they can only create
	// tokens for service principals they have the Service Principal User role
	// on, change a job owner to themselves, and set run_as to themselves or to
	// a service principal they have that role on.
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestrictWorkspaceAdminsMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestrictWorkspaceAdminsMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestrictWorkspaceAdminsMessageStatus string

const RestrictWorkspaceAdminsMessageStatusAllowAll RestrictWorkspaceAdminsMessageStatus = `ALLOW_ALL`

const RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs RestrictWorkspaceAdminsMessageStatus = `RESTRICT_TOKENS_AND_JOB_RUN_AS`

// String representation for [fmt.Print]
func (f *RestrictWorkspaceAdminsMessageStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RestrictWorkspaceAdminsMessageStatus) Set(v string) error {
	switch v {
	case `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`:
		*f = RestrictWorkspaceAdminsMessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL", "RESTRICT_TOKENS_AND_JOB_RUN_AS"`, v)
	}
}

// Values returns all possible values for RestrictWorkspaceAdminsMessageStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RestrictWorkspaceAdminsMessageStatus) Values() []RestrictWorkspaceAdminsMessageStatus {
	return []RestrictWorkspaceAdminsMessageStatus{
		RestrictWorkspaceAdminsMessageStatusAllowAll,
		RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs,
	}
}

// Type always returns RestrictWorkspaceAdminsMessageStatus to satisfy [pflag.Value] interface
func (f *RestrictWorkspaceAdminsMessageStatus) Type() string {
	return "RestrictWorkspaceAdminsMessageStatus"
}

// A setting and its value at the workspace or account scope. Each setting has
// one value type, so one value field applies to it: a primitive wrapper
// (boolean_val, string_val or integer_val) or a setting-specific message. The
// settings metadata endpoint reports the field each setting uses. When updating
// a setting, set only the value field that matches its type. Responses return
// that field with the value configured at this scope, if one is set, together
// with the matching effective_* field, which holds the value in effect after
// Databricks applies account policies and defaults.
type Setting struct {
	// Setting value for aibi_dashboard_embedding_access_policy setting. This is
	// the setting value set by consumers, check
	// effective_aibi_dashboard_embedding_access_policy for final setting value.
	AibiDashboardEmbeddingAccessPolicy *AibiDashboardEmbeddingAccessPolicy `json:"aibi_dashboard_embedding_access_policy,omitempty"`
	// Setting value for aibi_dashboard_embedding_approved_domains setting. This
	// is the setting value set by consumers, check
	// effective_aibi_dashboard_embedding_approved_domains for final setting
	// value.
	AibiDashboardEmbeddingApprovedDomains *AibiDashboardEmbeddingApprovedDomains `json:"aibi_dashboard_embedding_approved_domains,omitempty"`
	// Setting value for allowed_apps_user_api_scopes setting. This is the
	// setting value set by consumers, check
	// effective_allowed_apps_user_api_scopes for final setting value.
	AllowedAppsUserApiScopes *AllowedAppsUserApiScopesMessage `json:"allowed_apps_user_api_scopes,omitempty"`
	// Setting value for automatic_cluster_update_workspace setting. This is the
	// setting value set by consumers, check
	// effective_automatic_cluster_update_workspace for final setting value.
	AutomaticClusterUpdateWorkspace *ClusterAutoRestartMessage `json:"automatic_cluster_update_workspace,omitempty"`
	// Setting value for boolean type setting. This is the setting value set by
	// consumers, check effective_boolean_val for final setting value.
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Setting value for collaboration_platform_connectivity setting. This is
	// the setting value set by consumers, check
	// effective_collaboration_platform_connectivity for final setting value.
	CollaborationPlatformConnectivity *CollaborationPlatformConnectivityMessage `json:"collaboration_platform_connectivity,omitempty"`
	// Effective setting value for aibi_dashboard_embedding_access_policy
	// setting. This is the final effective value of setting. To set a value use
	// aibi_dashboard_embedding_access_policy.
	EffectiveAibiDashboardEmbeddingAccessPolicy *AibiDashboardEmbeddingAccessPolicy `json:"effective_aibi_dashboard_embedding_access_policy,omitempty"`
	// Effective setting value for aibi_dashboard_embedding_approved_domains
	// setting. This is the final effective value of setting. To set a value use
	// aibi_dashboard_embedding_approved_domains.
	EffectiveAibiDashboardEmbeddingApprovedDomains *AibiDashboardEmbeddingApprovedDomains `json:"effective_aibi_dashboard_embedding_approved_domains,omitempty"`
	// Effective setting value for allowed_apps_user_api_scopes setting. This is
	// the final effective value of setting. To set a value use
	// allowed_apps_user_api_scopes.
	EffectiveAllowedAppsUserApiScopes *AllowedAppsUserApiScopesMessage `json:"effective_allowed_apps_user_api_scopes,omitempty"`
	// Effective setting value for automatic_cluster_update_workspace setting.
	// This is the final effective value of setting. To set a value use
	// automatic_cluster_update_workspace.
	EffectiveAutomaticClusterUpdateWorkspace *ClusterAutoRestartMessage `json:"effective_automatic_cluster_update_workspace,omitempty"`
	// Effective setting value for boolean type setting. This is the final
	// effective value of setting. To set a value use boolean_val.
	EffectiveBooleanVal *BooleanMessage `json:"effective_boolean_val,omitempty"`
	// Effective setting value for collaboration_platform_connectivity setting.
	// This is the final effective value of setting. To set a value use
	// collaboration_platform_connectivity.
	EffectiveCollaborationPlatformConnectivity *CollaborationPlatformConnectivityMessage `json:"effective_collaboration_platform_connectivity,omitempty"`
	// Effective setting value for integer type setting. This is the final
	// effective value of setting. To set a value use integer_val.
	EffectiveIntegerVal *IntegerMessage `json:"effective_integer_val,omitempty"`
	// Effective setting value for operational_email_custom_recipient setting.
	// This is the final effective value of setting. To set a value use
	// operational_email_custom_recipient.
	EffectiveOperationalEmailCustomRecipient *OperationalEmailCustomRecipientMessage `json:"effective_operational_email_custom_recipient,omitempty"`
	// Effective setting value for personal_compute setting. This is the final
	// effective value of setting. To set a value use personal_compute.
	EffectivePersonalCompute *PersonalComputeMessage `json:"effective_personal_compute,omitempty"`
	// Effective setting value for restrict_workspace_admins setting. This is
	// the final effective value of setting. To set a value use
	// restrict_workspace_admins.
	EffectiveRestrictWorkspaceAdmins *RestrictWorkspaceAdminsMessage `json:"effective_restrict_workspace_admins,omitempty"`
	// Effective setting value for string type setting. This is the final
	// effective value of setting. To set a value use string_val.
	EffectiveStringVal *StringMessage `json:"effective_string_val,omitempty"`
	// Effective setting value for the workspace_dashboard_theme setting. This
	// is the final effective value of the setting. To set a value, use
	// workspace_dashboard_theme.
	EffectiveWorkspaceDashboardTheme *DashboardTheme `json:"effective_workspace_dashboard_theme,omitempty"`
	// Effective setting value for workspace_label setting. This is the final
	// effective value of setting. To set a value use workspace_label.
	EffectiveWorkspaceLabel *WorkspaceLabelMessage `json:"effective_workspace_label,omitempty"`
	// Setting value for integer type setting. This is the setting value set by
	// consumers, check effective_integer_val for final setting value.
	IntegerVal *IntegerMessage `json:"integer_val,omitempty"`
	// Name of the setting. The settings metadata endpoint lists the names
	// available at each scope.
	Name string `json:"name,omitempty"`
	// Setting value for operational_email_custom_recipient setting. This is the
	// setting value set by consumers, check
	// effective_operational_email_custom_recipient for final setting value.
	OperationalEmailCustomRecipient *OperationalEmailCustomRecipientMessage `json:"operational_email_custom_recipient,omitempty"`
	// Setting value for personal_compute setting. This is the setting value set
	// by consumers, check effective_personal_compute for final setting value.
	PersonalCompute *PersonalComputeMessage `json:"personal_compute,omitempty"`
	// Setting value for restrict_workspace_admins setting. This is the setting
	// value set by consumers, check effective_restrict_workspace_admins for
	// final setting value.
	RestrictWorkspaceAdmins *RestrictWorkspaceAdminsMessage `json:"restrict_workspace_admins,omitempty"`
	// Setting value for string type setting. This is the setting value set by
	// consumers, check effective_string_val for final setting value.
	StringVal *StringMessage `json:"string_val,omitempty"`
	// Setting value for the workspace_dashboard_theme setting. This is the
	// setting value set by consumers. Check effective_workspace_dashboard_theme
	// for the final setting value.
	WorkspaceDashboardTheme *DashboardTheme `json:"workspace_dashboard_theme,omitempty"`
	// Setting value for workspace_label setting. This is the setting value set
	// by consumers, check effective_workspace_label for final setting value.
	WorkspaceLabel *WorkspaceLabelMessage `json:"workspace_label,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Setting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Setting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SettingsMetadata struct {
	// Setting description for what this setting controls
	Description string `json:"description,omitempty"`
	// Human-readable display name for the setting or feature preview. This
	// field may be unset if no display name is available.
	DisplayName string `json:"display_name,omitempty"`
	// Link to databricks documentation for the setting
	DocsLink string `json:"docs_link,omitempty"`
	// Name of the setting.
	Name string `json:"name,omitempty"`
	// Preview phase for feature preview settings. This field is not set for
	// non-preview settings.
	PreviewPhase PreviewPhase `json:"preview_phase,omitempty"`
	// Sample message depicting the type of the setting. To set this setting,
	// the value sent must match this type.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SettingsMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SettingsMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StringMessage struct {
	// Represents a generic string value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StringMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StringMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// User Preference represents a user-specific setting scoped to an individual
// user within an account. Unlike workspace or account settings that apply to
// all users, user preferences allow personal customization (e.g., UI theme,
// editor preferences) without affecting other users.
type UserPreference struct {
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`

	EffectiveBooleanVal *BooleanMessage `json:"effective_boolean_val,omitempty"`

	EffectiveStringVal *StringMessage `json:"effective_string_val,omitempty"`
	// Name of the setting.
	Name string `json:"name,omitempty"`

	StringVal *StringMessage `json:"string_val,omitempty"`
	// User ID of the user.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UserPreference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UserPreference) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Workspace label and color for display (e.g. in account console).
type WorkspaceLabelMessage struct {
	// Color of the label badge. The supported colors are coral, lemon, lime,
	// turquoise, indigo, purple and pink. It is an empty string when no label
	// is set.
	Color string `json:"color,omitempty"`
	// Text of the label shown next to the workspace name in the top navigation
	// bar. An empty string means no label is set.
	Label string `json:"label,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceLabelMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceLabelMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
