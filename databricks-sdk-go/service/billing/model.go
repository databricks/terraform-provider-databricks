// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

type ActionConfiguration struct {
	// Databricks action configuration ID.
	ActionConfigurationId string `json:"action_configuration_id,omitempty"`
	// The action to take when the alert enters a triggered state.
	// `EMAIL_NOTIFICATION` sends an email and `BLOCK_USAGE` blocks further
	// usage. `BLOCK_USAGE` is only supported on Unity AI Gateway budgets, with
	// `resource_type` set to `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`.
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target of the action. For `EMAIL_NOTIFICATION`, the email address that
	// receives the alert, up to 256 characters. Must be omitted for
	// `BLOCK_USAGE`.
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ActionConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ActionConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of action that a budget alert executes when its threshold is crossed.
type ActionConfigurationType string

const ActionConfigurationTypeBlockUsage ActionConfigurationType = `BLOCK_USAGE`

const ActionConfigurationTypeEmailNotification ActionConfigurationType = `EMAIL_NOTIFICATION`

// String representation for [fmt.Print]
func (f *ActionConfigurationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActionConfigurationType) Set(v string) error {
	switch v {
	case `BLOCK_USAGE`, `EMAIL_NOTIFICATION`:
		*f = ActionConfigurationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK_USAGE", "EMAIL_NOTIFICATION"`, v)
	}
}

// Values returns all possible values for ActionConfigurationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActionConfigurationType) Values() []ActionConfigurationType {
	return []ActionConfigurationType{
		ActionConfigurationTypeBlockUsage,
		ActionConfigurationTypeEmailNotification,
	}
}

// Type always returns ActionConfigurationType to satisfy [pflag.Value] interface
func (f *ActionConfigurationType) Type() string {
	return "ActionConfigurationType"
}

type AlertConfiguration struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state. An alert can have up to 10 actions.
	ActionConfigurations []ActionConfiguration `json:"action_configurations,omitempty"`
	// Databricks alert configuration ID.
	AlertConfigurationId string `json:"alert_configuration_id,omitempty"`
	// Per-principal threshold overrides for this alert. Only applies to
	// per-user alerts (`scope_type` =
	// `ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER`); ignored for shared alerts.
	PrincipalOverrides []PrincipalOverride `json:"principal_overrides,omitempty"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`. A decimal number
	// passed as a string, from 0 to 1,000,000,000.
	QuantityThreshold string `json:"quantity_threshold,omitempty"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
	// How the alert threshold is evaluated. Determines whether spend is tracked
	// in aggregate or per individual user.
	ScopeType AlertConfigurationScopeType `json:"scope_type,omitempty"`
	// The time window of usage data for the budget.
	TimePeriod AlertConfigurationTimePeriod `json:"time_period,omitempty"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType AlertConfigurationTriggerType `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertConfigurationQuantityType string

const AlertConfigurationQuantityTypeListPriceDollarsUsd AlertConfigurationQuantityType = `LIST_PRICE_DOLLARS_USD`

// String representation for [fmt.Print]
func (f *AlertConfigurationQuantityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertConfigurationQuantityType) Set(v string) error {
	switch v {
	case `LIST_PRICE_DOLLARS_USD`:
		*f = AlertConfigurationQuantityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LIST_PRICE_DOLLARS_USD"`, v)
	}
}

// Values returns all possible values for AlertConfigurationQuantityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertConfigurationQuantityType) Values() []AlertConfigurationQuantityType {
	return []AlertConfigurationQuantityType{
		AlertConfigurationQuantityTypeListPriceDollarsUsd,
	}
}

// Type always returns AlertConfigurationQuantityType to satisfy [pflag.Value] interface
func (f *AlertConfigurationQuantityType) Type() string {
	return "AlertConfigurationQuantityType"
}

// Evaluation scope for an alert configuration.
type AlertConfigurationScopeType string

const AlertConfigurationScopeTypeAlertConfigurationScopeTypePerUser AlertConfigurationScopeType = `ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER`

const AlertConfigurationScopeTypeAlertConfigurationScopeTypeShared AlertConfigurationScopeType = `ALERT_CONFIGURATION_SCOPE_TYPE_SHARED`

// String representation for [fmt.Print]
func (f *AlertConfigurationScopeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertConfigurationScopeType) Set(v string) error {
	switch v {
	case `ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER`, `ALERT_CONFIGURATION_SCOPE_TYPE_SHARED`:
		*f = AlertConfigurationScopeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER", "ALERT_CONFIGURATION_SCOPE_TYPE_SHARED"`, v)
	}
}

// Values returns all possible values for AlertConfigurationScopeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertConfigurationScopeType) Values() []AlertConfigurationScopeType {
	return []AlertConfigurationScopeType{
		AlertConfigurationScopeTypeAlertConfigurationScopeTypePerUser,
		AlertConfigurationScopeTypeAlertConfigurationScopeTypeShared,
	}
}

// Type always returns AlertConfigurationScopeType to satisfy [pflag.Value] interface
func (f *AlertConfigurationScopeType) Type() string {
	return "AlertConfigurationScopeType"
}

type AlertConfigurationTimePeriod string

const AlertConfigurationTimePeriodDay AlertConfigurationTimePeriod = `DAY`

const AlertConfigurationTimePeriodMonth AlertConfigurationTimePeriod = `MONTH`

// String representation for [fmt.Print]
func (f *AlertConfigurationTimePeriod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertConfigurationTimePeriod) Set(v string) error {
	switch v {
	case `DAY`, `MONTH`:
		*f = AlertConfigurationTimePeriod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAY", "MONTH"`, v)
	}
}

// Values returns all possible values for AlertConfigurationTimePeriod.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertConfigurationTimePeriod) Values() []AlertConfigurationTimePeriod {
	return []AlertConfigurationTimePeriod{
		AlertConfigurationTimePeriodDay,
		AlertConfigurationTimePeriodMonth,
	}
}

// Type always returns AlertConfigurationTimePeriod to satisfy [pflag.Value] interface
func (f *AlertConfigurationTimePeriod) Type() string {
	return "AlertConfigurationTimePeriod"
}

type AlertConfigurationTriggerType string

const AlertConfigurationTriggerTypeCumulativeSpendingExceeded AlertConfigurationTriggerType = `CUMULATIVE_SPENDING_EXCEEDED`

// String representation for [fmt.Print]
func (f *AlertConfigurationTriggerType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertConfigurationTriggerType) Set(v string) error {
	switch v {
	case `CUMULATIVE_SPENDING_EXCEEDED`:
		*f = AlertConfigurationTriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUMULATIVE_SPENDING_EXCEEDED"`, v)
	}
}

// Values returns all possible values for AlertConfigurationTriggerType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertConfigurationTriggerType) Values() []AlertConfigurationTriggerType {
	return []AlertConfigurationTriggerType{
		AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
	}
}

// Type always returns AlertConfigurationTriggerType to satisfy [pflag.Value] interface
func (f *AlertConfigurationTriggerType) Type() string {
	return "AlertConfigurationTriggerType"
}

// The calling user's spend against a single blocking threshold. Values may be
// omitted if live spend or threshold data is temporarily unavailable.
type AlertSpendStatus struct {
	// Percentage of the threshold consumed, as a decimal string rounded down to
	// one decimal place (for example, "68.4" means 68.4%). May exceed 100 when
	// spend is over the threshold. Returns "0.0" if there's no spend, except
	// that a $0 threshold always returns "100.0". For shared thresholds, this
	// is the only value returned; the absolute `spend` and `threshold` dollar
	// amounts are withheld.
	PercentConsumed string `json:"percent_consumed,omitempty"`
	// Whether the threshold is evaluated per individual user
	// (`ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER`) or on the shared spend across
	// all users (`ALERT_CONFIGURATION_SCOPE_TYPE_SHARED`).
	ScopeType AlertConfigurationScopeType `json:"scope_type,omitempty"`
	// The calling user's live spend against this threshold, in USD at list
	// price. Returns "0" if the caller has no usage. Omitted for shared
	// thresholds, where only `percent_consumed` is returned.
	Spend string `json:"spend,omitempty"`
	// The calling user's effective threshold, in USD at list price. The
	// effective threshold is the actual threshold amount that applies to the
	// user, after adjusting for any user or group overrides. Omitted for shared
	// thresholds, where only `percent_consumed` is returned.
	Threshold string `json:"threshold,omitempty"`
	// The time window the threshold applies to, for example `MONTH`.
	TimePeriod AlertConfigurationTimePeriod `json:"time_period,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertSpendStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertSpendStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A budget tracks the spend of your account, or of the workspaces and custom
// tags selected by its filter, against one or more alert thresholds. Spend is
// measured in US dollars at list price. For an overview, see
// https://docs.databricks.com/en/admin/account-settings/budgets.html.
type BudgetConfiguration struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have at least one alert configuration. A budget can have up to four
	// shared alerts per time period, each with a unique `quantity_threshold`.
	// Budgets with `resource_type` set to
	// `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY` can also use per-user alerts,
	// with optional `principal_overrides`. These budgets can have at most one
	// per-user alert per time period, and can use shared alerts, a per-user
	// alert, or both. The `BLOCK_USAGE` action is only supported on budgets
	// with `resource_type` set to `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`. At
	// most one shared alert per time period can have a `BLOCK_USAGE` action,
	// and it must be the shared alert with the highest threshold in that
	// period.
	AlertConfigurations []AlertConfiguration `json:"alert_configurations,omitempty"`
	// Databricks budget configuration ID.
	BudgetConfigurationId string `json:"budget_configuration_id,omitempty"`
	// Creation time of this budget configuration, in milliseconds since the
	// Unix epoch.
	CreateTime int64 `json:"create_time,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included. The filter cannot be changed after the
	// budget is created. To change it, delete the budget and create a new one.
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`
	// Whether this budget tracks external pass-through spend (third-party costs
	// billed through a Databricks product) in addition to internal Databricks
	// spend. Only applicable to Unity Gateway budgets, i.e. when
	// `resource_type` is `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`. This field is
	// immutable after create.
	IncludeExternalSpend bool `json:"include_external_spend,omitempty"`
	// The resource scope for this budget. Determines whether the budget tracks
	// all resources or a specific resource.
	ResourceType BudgetResourceType `json:"resource_type,omitempty"`
	// Time of the most recent update to this budget configuration, in
	// milliseconds since the Unix epoch.
	UpdateTime int64 `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BudgetConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BudgetConfigurationFilter struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	Tags []BudgetConfigurationFilterTagClause `json:"tags,omitempty"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	WorkspaceId *BudgetConfigurationFilterWorkspaceIdClause `json:"workspace_id,omitempty"`
}

func (s *BudgetConfigurationFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BudgetConfigurationFilterClause struct {
	// How usage is compared with `values`. `IN` matches usage whose value is
	// one of `values`.
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`
	// Values to match. A clause can list up to 30 values.
	Values []string `json:"values,omitempty"`
}

func (s *BudgetConfigurationFilterClause) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BudgetConfigurationFilterOperator string

const BudgetConfigurationFilterOperatorIn BudgetConfigurationFilterOperator = `IN`

// String representation for [fmt.Print]
func (f *BudgetConfigurationFilterOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BudgetConfigurationFilterOperator) Set(v string) error {
	switch v {
	case `IN`:
		*f = BudgetConfigurationFilterOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN"`, v)
	}
}

// Values returns all possible values for BudgetConfigurationFilterOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *BudgetConfigurationFilterOperator) Values() []BudgetConfigurationFilterOperator {
	return []BudgetConfigurationFilterOperator{
		BudgetConfigurationFilterOperatorIn,
	}
}

// Type always returns BudgetConfigurationFilterOperator to satisfy [pflag.Value] interface
func (f *BudgetConfigurationFilterOperator) Type() string {
	return "BudgetConfigurationFilterOperator"
}

type BudgetConfigurationFilterTagClause struct {
	// The custom tag key to match, exactly as it appears in your usage data. A
	// key can appear in at most one clause per operator.
	Key string `json:"key,omitempty"`
	// The tag values to match for `key`. Usage is included when its value for
	// `key` is one of the listed values.
	Value *BudgetConfigurationFilterClause `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BudgetConfigurationFilterTagClause) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetConfigurationFilterTagClause) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BudgetConfigurationFilterWorkspaceIdClause struct {
	// How usage is compared with `values`. `IN` matches usage from any of the
	// listed workspaces.
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`
	// IDs of the workspaces whose usage is included. A clause can list up to 30
	// workspaces, all of which must belong to this account.
	Values []int64 `json:"values,omitempty"`
}

func (s *BudgetConfigurationFilterWorkspaceIdClause) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A serverless usage policy is a set of tags that are applied to the serverless
// usage of the workloads that use the policy. Each usage record in the billable
// usage system table carries the policy ID and these tags, so that serverless
// usage can be attributed to the policy.
type BudgetPolicy struct {
	// List of workspaces that this budget policy will be exclusively bound to.
	// An empty binding implies that this budget policy is open to any workspace
	// in the account.
	BindingWorkspaceIds []int64 `json:"binding_workspace_ids,omitempty"`
	// A list of tags defined by the customer. At most 20 entries are allowed
	// per policy.
	CustomTags []compute.CustomPolicyTag `json:"custom_tags,omitempty"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"policy_id,omitempty"`
	// The name of the policy. - Must be unique among active policies. - Can
	// contain only characters from the ISO 8859-1 (latin1) set. - Can't start
	// with reserved keywords such as `databricks:default-policy`.
	PolicyName string `json:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BudgetPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Resource scope for a budget configuration. Determines whether the budget
// tracks all resources or a specific resource.
type BudgetResourceType string

const BudgetResourceTypeBudgetResourceTypeAllResources BudgetResourceType = `BUDGET_RESOURCE_TYPE_ALL_RESOURCES`

const BudgetResourceTypeBudgetResourceTypeUnityAiGateway BudgetResourceType = `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`

// String representation for [fmt.Print]
func (f *BudgetResourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BudgetResourceType) Set(v string) error {
	switch v {
	case `BUDGET_RESOURCE_TYPE_ALL_RESOURCES`, `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`:
		*f = BudgetResourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUDGET_RESOURCE_TYPE_ALL_RESOURCES", "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY"`, v)
	}
}

// Values returns all possible values for BudgetResourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *BudgetResourceType) Values() []BudgetResourceType {
	return []BudgetResourceType{
		BudgetResourceTypeBudgetResourceTypeAllResources,
		BudgetResourceTypeBudgetResourceTypeUnityAiGateway,
	}
}

// Type always returns BudgetResourceType to satisfy [pflag.Value] interface
func (f *BudgetResourceType) Type() string {
	return "BudgetResourceType"
}

// The calling user's live spend against the blocking thresholds of a single
// budget.
type BudgetSpendStatus struct {
	// The calling user's spend against each blocking threshold on this budget,
	// one entry per scope type and time period. A blocking threshold is a
	// threshold that, once reached, blocks further requests through Unity
	// Gateway.
	AlertSpendStatuses []AlertSpendStatus `json:"alert_spend_statuses,omitempty"`
	// The resource name of this budget spend status, of the form
	// `accounts/{account_id}/budget-spend-statuses/{budget_id}`.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BudgetSpendStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetSpendStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account. The request fails if the type is not set.
	DashboardType UsageDashboardType `json:"dashboard_type,omitempty"`
	// The major version of the usage dashboard template to import. Version 2
	// adds cost forecasting and object-level cost drill-down.
	MajorVersion UsageDashboardMajorVersion `json:"major_version,omitempty"`
	// The ID of the workspace to import the usage dashboard into. The workspace
	// must be running, belong to the account, and not be a disaster recovery
	// secondary.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBillingUsageDashboardResponse struct {
	// The ID of the AI/BI dashboard that was created.
	DashboardId string `json:"dashboard_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A budget tracks the spend of your account, or of the workspaces and custom
// tags selected by its filter, against one or more alert thresholds. Spend is
// measured in US dollars at list price. For an overview, see
// https://docs.databricks.com/en/admin/account-settings/budgets.html.
type CreateBudgetConfigurationBudget struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have at least one alert configuration. A budget can have up to four
	// shared alerts per time period, each with a unique `quantity_threshold`.
	// Budgets with `resource_type` set to
	// `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY` can also use per-user alerts,
	// with optional `principal_overrides`. These budgets can have at most one
	// per-user alert per time period, and can use shared alerts, a per-user
	// alert, or both. The `BLOCK_USAGE` action is only supported on budgets
	// with `resource_type` set to `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`. At
	// most one shared alert per time period can have a `BLOCK_USAGE` action,
	// and it must be the shared alert with the highest threshold in that
	// period.
	AlertConfigurations []CreateBudgetConfigurationBudgetAlertConfigurations `json:"alert_configurations,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included. The filter cannot be changed after the
	// budget is created. To change it, delete the budget and create a new one.
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`
	// Whether this budget tracks external pass-through spend (third-party costs
	// billed through a Databricks product) in addition to internal Databricks
	// spend. Only applicable to Unity Gateway budgets, i.e. when
	// `resource_type` is `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`. This field is
	// immutable after create.
	IncludeExternalSpend bool `json:"include_external_spend,omitempty"`
	// The resource scope for this budget. Determines whether the budget tracks
	// all resources or a specific resource.
	ResourceType BudgetResourceType `json:"resource_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBudgetConfigurationBudgetActionConfigurations struct {
	// The action to take when the alert enters a triggered state.
	// `EMAIL_NOTIFICATION` sends an email and `BLOCK_USAGE` blocks further
	// usage. `BLOCK_USAGE` is only supported on Unity AI Gateway budgets, with
	// `resource_type` set to `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`.
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target of the action. For `EMAIL_NOTIFICATION`, the email address that
	// receives the alert, up to 256 characters. Must be omitted for
	// `BLOCK_USAGE`.
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBudgetConfigurationBudgetActionConfigurations) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetConfigurationBudgetActionConfigurations) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBudgetConfigurationBudgetAlertConfigurations struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state. An alert can have up to 10 actions.
	ActionConfigurations []CreateBudgetConfigurationBudgetActionConfigurations `json:"action_configurations,omitempty"`
	// Per-principal threshold overrides for this alert. Only applies to
	// per-user alerts (`scope_type` =
	// `ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER`); ignored for shared alerts.
	PrincipalOverrides []PrincipalOverride `json:"principal_overrides,omitempty"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`. A decimal number
	// passed as a string, from 0 to 1,000,000,000.
	QuantityThreshold string `json:"quantity_threshold,omitempty"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
	// How the alert threshold is evaluated. Determines whether spend is tracked
	// in aggregate or per individual user.
	ScopeType AlertConfigurationScopeType `json:"scope_type,omitempty"`
	// The time window of usage data for the budget.
	TimePeriod AlertConfigurationTimePeriod `json:"time_period,omitempty"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType AlertConfigurationTriggerType `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBudgetConfigurationBudgetAlertConfigurations) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetConfigurationBudgetAlertConfigurations) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBudgetConfigurationRequest struct {
	// Properties of the new budget configuration.
	Budget CreateBudgetConfigurationBudget `json:"budget"`
}

func (s *CreateBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateBudgetConfigurationResponse struct {
	// The created budget configuration.
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func (s *CreateBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A request to create a BudgetPolicy.
type CreateBudgetPolicyRequest struct {
	// The policy to create. `policy_name` must be provided, and `custom_tags`
	// may be required depending on the cloud provider. Leave `policy_id` empty
	// because Databricks generates it. All other fields are optional.
	Policy *BudgetPolicy `json:"policy,omitempty"`
	// A unique identifier for this request. Restricted to 36 ASCII characters.
	// A random UUID is recommended. This request is only idempotent if a
	// `request_id` is provided.
	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * Log Delivery Configuration
type CreateLogDeliveryConfigurationParams struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `json:"config_name,omitempty"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId string `json:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if log_type is BILLABLE_USAGE. This is the
	// optional start month and year for delivery, specified in YYYY-MM format.
	// Defaults to current year and month. BILLABLE_USAGE logs are not available
	// for usage before March 2019 (2019-03).
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// Log delivery type. Supported values are: * `BILLABLE_USAGE` — Configure
	// [billable usage log delivery]. For the CSV schema, see the [View billable
	// usage]. * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON
	// schema, see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType LogType `json:"log_type"`
	// The file type of log delivery. * If `log_type` is `BILLABLE_USAGE`, this
	// value must be `CSV`. Only the CSV (comma-separated values) format is
	// supported. For the schema, see the [View billable usage] * If `log_type`
	// is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON (JavaScript
	// Object Notation) format is supported. For the schema, see the
	// [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat OutputFormat `json:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId string `json:"storage_configuration_id"`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateLogDeliveryConfigurationParams) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateLogDeliveryConfigurationParams) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteBudgetConfigurationRequest struct {
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" url:"-"`
}

func (s *DeleteBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" url:"-"`
}

func (s *DeleteBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// * The status string for log delivery. Possible values are: `CREATED`: There
// were no log delivery attempts since the config was created. `SUCCEEDED`: The
// latest attempt of log delivery has succeeded completely. `USER_FAILURE`: The
// latest attempt of log delivery failed because of misconfiguration of customer
// provided permissions on role or storage. `SYSTEM_FAILURE`: The latest attempt
// of log delivery failed because of an Databricks internal error. Contact
// support if it doesn't go away soon. `NOT_FOUND`: The log delivery status as
// the configuration has been disabled since the release of this feature or
// there are no workspaces in the account.
type DeliveryStatus string

const DeliveryStatusCreated DeliveryStatus = `CREATED`

const DeliveryStatusNotFound DeliveryStatus = `NOT_FOUND`

const DeliveryStatusSucceeded DeliveryStatus = `SUCCEEDED`

const DeliveryStatusSystemFailure DeliveryStatus = `SYSTEM_FAILURE`

const DeliveryStatusUserFailure DeliveryStatus = `USER_FAILURE`

// String representation for [fmt.Print]
func (f *DeliveryStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeliveryStatus) Set(v string) error {
	switch v {
	case `CREATED`, `NOT_FOUND`, `SUCCEEDED`, `SYSTEM_FAILURE`, `USER_FAILURE`:
		*f = DeliveryStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATED", "NOT_FOUND", "SUCCEEDED", "SYSTEM_FAILURE", "USER_FAILURE"`, v)
	}
}

// Values returns all possible values for DeliveryStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeliveryStatus) Values() []DeliveryStatus {
	return []DeliveryStatus{
		DeliveryStatusCreated,
		DeliveryStatusNotFound,
		DeliveryStatusSucceeded,
		DeliveryStatusSystemFailure,
		DeliveryStatusUserFailure,
	}
}

// Type always returns DeliveryStatus to satisfy [pflag.Value] interface
func (f *DeliveryStatus) Type() string {
	return "DeliveryStatus"
}

type DownloadRequest struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth string `json:"-" url:"end_month"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData bool `json:"-" url:"personal_data,omitempty"`
	// Format specification for month in the format `YYYY-MM`. This is used to
	// specify billable usage `start_month` and `end_month` properties.
	// **Note**: Billable usage logs are unavailable before March 2019
	// (`2019-03`).
	StartMonth string `json:"-" url:"start_month"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DownloadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DownloadRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DownloadResponse struct {
	Contents io.ReadCloser `json:"-"`
}

func (s *DownloadResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Structured representation of a filter to be applied to a list of policies.
// All specified filters will be applied in conjunction.
type Filter struct {
	// Deprecated: This field is ignored. Creator filtering will be removed in a
	// future version.
	CreatorUserId int64 `json:"creator_user_id,omitempty" url:"creator_user_id,omitempty"`
	// Deprecated: Do not use this field in new integrations. Creator filtering
	// will be removed in a future version. The policy creator user name to be
	// filtered on. If unspecified, all policies will be returned.
	CreatorUserName string `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	// Only policies whose name contains this string are returned. Matching is
	// case-insensitive. If unspecified, all policies are returned.
	PolicyName string `json:"policy_name,omitempty" url:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Filter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Filter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account. The request fails if the type is not set.
	DashboardType UsageDashboardType `json:"-" url:"dashboard_type,omitempty"`
	// The ID of the workspace that contains the usage dashboard. If omitted,
	// the most recently created dashboard of the given type across all
	// workspaces in the account is returned.
	WorkspaceId int64 `json:"-" url:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetBillingUsageDashboardResponse struct {
	// The ID of the AI/BI dashboard. The ID is empty if no matching dashboard
	// is found.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The URL of the dashboard in its workspace. The URL is empty if no
	// matching dashboard is found.
	DashboardUrl string `json:"dashboard_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetBudgetConfigurationRequest struct {
	// The budget configuration ID
	BudgetId string `json:"-" url:"-"`
}

func (s *GetBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBudgetConfigurationResponse struct {
	// The requested budget configuration.
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func (s *GetBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" url:"-"`
}

func (s *GetBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBudgetSpendStatusRequest struct {
	// The resource name of the budget spend status to read, of the form
	// `accounts/{account_id}/budget-spend-statuses/{budget_id}`. A `budget_id`
	// that isn't a valid UUID fails with INVALID_PARAMETER_VALUE.
	Name string `json:"-" url:"-"`
}

func (s *GetBudgetSpendStatusRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLogDeliveryConfigurationResponse struct {
	// The fetched log delivery configuration
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

func (s *GetLogDeliveryConfigurationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLogDeliveryRequest struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId string `json:"-" url:"-"`
}

func (s *GetLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The limit configuration of the policy. Limit configuration provide a budget
// policy level cost control by enforcing the limit.
type LimitConfig struct {
}

func (s *LimitConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListBudgetConfigurationsRequest struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetConfigurationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetConfigurationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBudgetConfigurationsResponse struct {
	Budgets []BudgetConfiguration `json:"budgets,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetConfigurationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBudgetPoliciesRequest struct {
	// A filter to apply to the list of policies.
	FilterBy *Filter `json:"-" url:"filter_by,omitempty"`
	// The maximum number of budget policies to return. If unspecified, at most
	// 100 budget policies will be returned. The maximum value is 1000; values
	// above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListBudgetPolicies` call. Provide
	// this to retrieve the subsequent page. If unspecified, the first page will
	// be returned.
	//
	// When paginating, all other parameters provided to `ListBudgetPolicies`
	// must match the call that provided the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// How to sort the returned policies. If unspecified, policies are sorted by
	// name in ascending order.
	SortSpec *SortSpec `json:"-" url:"sort_spec,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of policies.
type ListBudgetPoliciesResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	Policies []BudgetPolicy `json:"policies,omitempty"`
	// A token that can be sent as `page_token` to retrieve the previous page.
	// In this field is omitted, there are no previous pages.
	PreviousPageToken string `json:"previous_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBudgetSpendStatusesRequest struct {
	// A filter expression that narrows the returned budgets by workspace. Only
	// `workspace_id` is supported, in lower case with numeric workspace IDs,
	// for example `workspace_id = 123` or `workspace_id IN (123, 456)`. Matches
	// budgets that apply to any of the listed workspaces, including budgets
	// that apply to all workspaces. Any other expression, such as an empty `IN
	// ()` list, fails with INVALID_PARAMETER_VALUE. If empty, returns every
	// budget in the account that has a blocking threshold.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of spend statuses to return per page. If unspecified
	// or 0, at most 100 are returned. The maximum value is 100; values above
	// 100 are coerced to 100. Negative values fail with
	// INVALID_PARAMETER_VALUE. The server may return fewer results than
	// requested, including zero, even when more pages remain; page until
	// `next_page_token` is absent.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous `ListBudgetSpendStatuses` response, used to
	// fetch the next page. Requests the first page if absent. All other request
	// parameters except `page_size` must match the request that returned the
	// token; otherwise the request fails with INVALID_PARAMETER_VALUE.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent account whose budget spend statuses to list, of the form
	// `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetSpendStatusesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetSpendStatusesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ListBudgetSpendStatuses.
type ListBudgetSpendStatusesResponse struct {
	// The calling user's spend status, one entry per returned budget.
	BudgetSpendStatuses []BudgetSpendStatus `json:"budget_spend_statuses,omitempty"`
	// Token to pass as `page_token` to retrieve the next page. Absent when
	// there are no more results. A page can contain fewer results than
	// `page_size`, or none, while this token is still set, so keep paging until
	// it's absent.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetSpendStatusesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetSpendStatusesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListLogDeliveryRequest struct {
	// The Credentials id to filter the search results with
	CredentialsId string `json:"-" url:"credentials_id,omitempty"`
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The log delivery status to filter the search results with
	Status LogDeliveryConfigStatus `json:"-" url:"status,omitempty"`
	// The Storage Configuration id to filter the search results with
	StorageConfigurationId string `json:"-" url:"storage_configuration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListLogDeliveryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * Log Delivery Status
//
// `ENABLED`: All dependencies have executed and succeeded `DISABLED`: At least
// one dependency has succeeded
type LogDeliveryConfigStatus string

const LogDeliveryConfigStatusDisabled LogDeliveryConfigStatus = `DISABLED`

const LogDeliveryConfigStatusEnabled LogDeliveryConfigStatus = `ENABLED`

// String representation for [fmt.Print]
func (f *LogDeliveryConfigStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LogDeliveryConfigStatus) Set(v string) error {
	switch v {
	case `DISABLED`, `ENABLED`:
		*f = LogDeliveryConfigStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLED", "ENABLED"`, v)
	}
}

// Values returns all possible values for LogDeliveryConfigStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *LogDeliveryConfigStatus) Values() []LogDeliveryConfigStatus {
	return []LogDeliveryConfigStatus{
		LogDeliveryConfigStatusDisabled,
		LogDeliveryConfigStatusEnabled,
	}
}

// Type always returns LogDeliveryConfigStatus to satisfy [pflag.Value] interface
func (f *LogDeliveryConfigStatus) Type() string {
	return "LogDeliveryConfigStatus"
}

// * Log Delivery Configuration
type LogDeliveryConfiguration struct {
	// Databricks account ID.
	AccountId string `json:"account_id"`
	// The unique UUID of log delivery configuration
	ConfigId string `json:"config_id,omitempty"`
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `json:"config_name,omitempty"`
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId string `json:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if log_type is BILLABLE_USAGE. This is the
	// optional start month and year for delivery, specified in YYYY-MM format.
	// Defaults to current year and month. BILLABLE_USAGE logs are not available
	// for usage before March 2019 (2019-03).
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// The LogDeliveryStatus of this log delivery configuration
	LogDeliveryStatus *LogDeliveryStatus `json:"log_delivery_status,omitempty"`
	// Log delivery type. Supported values are: * `BILLABLE_USAGE` — Configure
	// [billable usage log delivery]. For the CSV schema, see the [View billable
	// usage]. * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON
	// schema, see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType LogType `json:"log_type"`
	// The file type of log delivery. * If `log_type` is `BILLABLE_USAGE`, this
	// value must be `CSV`. Only the CSV (comma-separated values) format is
	// supported. For the schema, see the [View billable usage] * If `log_type`
	// is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON (JavaScript
	// Object Notation) format is supported. For the schema, see the
	// [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat OutputFormat `json:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId string `json:"storage_configuration_id"`
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	UpdateTime int64 `json:"update_time,omitempty"`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogDeliveryStatus struct {
	// The UTC time for the latest log delivery attempt.
	LastAttemptTime string `json:"last_attempt_time,omitempty"`
	// The UTC time for the latest successful log delivery.
	LastSuccessfulAttemptTime string `json:"last_successful_attempt_time,omitempty"`
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	Message string `json:"message"`
	// Enum that describes the status. Possible values are: * `CREATED`: There
	// were no log delivery attempts since the config was created. *
	// `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
	// * `USER_FAILURE`: The latest attempt of log delivery failed because of
	// misconfiguration of customer provided permissions on role or storage. *
	// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
	// Databricks internal error. Contact support if it doesn't go away soon. *
	// `NOT_FOUND`: The log delivery status as the configuration has been
	// disabled since the release of this feature or there are no workspaces in
	// the account.
	Status DeliveryStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogDeliveryStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogDeliveryStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * Log Delivery Type
type LogType string

const LogTypeAuditLogs LogType = `AUDIT_LOGS`

const LogTypeBillableUsage LogType = `BILLABLE_USAGE`

// String representation for [fmt.Print]
func (f *LogType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LogType) Set(v string) error {
	switch v {
	case `AUDIT_LOGS`, `BILLABLE_USAGE`:
		*f = LogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUDIT_LOGS", "BILLABLE_USAGE"`, v)
	}
}

// Values returns all possible values for LogType.
//
// There is no guarantee on the order of the values in the slice.
func (f *LogType) Values() []LogType {
	return []LogType{
		LogTypeAuditLogs,
		LogTypeBillableUsage,
	}
}

// Type always returns LogType to satisfy [pflag.Value] interface
func (f *LogType) Type() string {
	return "LogType"
}

// * Log Delivery Output Format
type OutputFormat string

const OutputFormatCsv OutputFormat = `CSV`

const OutputFormatJson OutputFormat = `JSON`

// String representation for [fmt.Print]
func (f *OutputFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OutputFormat) Set(v string) error {
	switch v {
	case `CSV`, `JSON`:
		*f = OutputFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CSV", "JSON"`, v)
	}
}

// Values returns all possible values for OutputFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		OutputFormatCsv,
		OutputFormatJson,
	}
}

// Type always returns OutputFormat to satisfy [pflag.Value] interface
func (f *OutputFormat) Type() string {
	return "OutputFormat"
}

// Per-principal threshold override on a PER_USER alert: bumps the alert's
// quantity_threshold for one principal_id.
type PrincipalOverride struct {
	// Dollar amount that overrides the parent alert's quantity_threshold for
	// this principal.
	OverrideThreshold string `json:"override_threshold,omitempty"`
	// Account-level principal id (user, group, or service principal).
	PrincipalId int64 `json:"principal_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrincipalOverride) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrincipalOverride) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SortSpec struct {
	// Whether to sort in descending order.
	Descending bool `json:"descending,omitempty" url:"descending,omitempty"`
	// The field to sort by. If unspecified, policies are sorted by name.
	Field SortSpecField `json:"field,omitempty" url:"field,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SortSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SortSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SortSpecField string

const SortSpecFieldPolicyName SortSpecField = `POLICY_NAME`

// String representation for [fmt.Print]
func (f *SortSpecField) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SortSpecField) Set(v string) error {
	switch v {
	case `POLICY_NAME`:
		*f = SortSpecField(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "POLICY_NAME"`, v)
	}
}

// Values returns all possible values for SortSpecField.
//
// There is no guarantee on the order of the values in the slice.
func (f *SortSpecField) Values() []SortSpecField {
	return []SortSpecField{
		SortSpecFieldPolicyName,
	}
}

// Type always returns SortSpecField to satisfy [pflag.Value] interface
func (f *SortSpecField) Type() string {
	return "SortSpecField"
}

// A budget tracks the spend of your account, or of the workspaces and custom
// tags selected by its filter, against one or more alert thresholds. Spend is
// measured in US dollars at list price. For an overview, see
// https://docs.databricks.com/en/admin/account-settings/budgets.html.
type UpdateBudgetConfigurationBudget struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have at least one alert configuration. A budget can have up to four
	// shared alerts per time period, each with a unique `quantity_threshold`.
	// Budgets with `resource_type` set to
	// `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY` can also use per-user alerts,
	// with optional `principal_overrides`. These budgets can have at most one
	// per-user alert per time period, and can use shared alerts, a per-user
	// alert, or both. The `BLOCK_USAGE` action is only supported on budgets
	// with `resource_type` set to `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`. At
	// most one shared alert per time period can have a `BLOCK_USAGE` action,
	// and it must be the shared alert with the highest threshold in that
	// period.
	AlertConfigurations []AlertConfiguration `json:"alert_configurations,omitempty"`
	// Databricks budget configuration ID.
	BudgetConfigurationId string `json:"budget_configuration_id,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included. The filter cannot be changed after the
	// budget is created. To change it, delete the budget and create a new one.
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`
	// Whether this budget tracks external pass-through spend (third-party costs
	// billed through a Databricks product) in addition to internal Databricks
	// spend. Only applicable to Unity Gateway budgets, i.e. when
	// `resource_type` is `BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY`. This field is
	// immutable after create.
	IncludeExternalSpend bool `json:"include_external_spend,omitempty"`
	// The resource scope for this budget. Determines whether the budget tracks
	// all resources or a specific resource.
	ResourceType BudgetResourceType `json:"resource_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateBudgetConfigurationRequest struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID. Alerts and actions are matched by ID: those omitted from the
	// request are removed, and those without an ID are created.
	Budget UpdateBudgetConfigurationBudget `json:"budget"`
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" url:"-"`
}

func (s *UpdateBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateBudgetConfigurationResponse struct {
	// The updated budget.
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func (s *UpdateBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateBudgetPolicyRequest struct {
	// Deprecated: This field is ignored.
	LimitConfig *LimitConfig `json:"-" url:"limit_config,omitempty"`
	// The policy to update. `creator_user_id` cannot be specified in the
	// request. All other fields must be specified even if not changed. The
	// `policy_id` is used to identify the policy to update.
	Policy BudgetPolicy `json:"policy"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"-" url:"-"`
}

func (s *UpdateBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// * Update Log Delivery Configuration
type UpdateLogDeliveryConfigurationStatusRequest struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId string `json:"-" url:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status"`
}

func (s *UpdateLogDeliveryConfigurationStatusRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UsageDashboardMajorVersion string

const UsageDashboardMajorVersionUsageDashboardMajorVersion1 UsageDashboardMajorVersion = `USAGE_DASHBOARD_MAJOR_VERSION_1`

const UsageDashboardMajorVersionUsageDashboardMajorVersion2 UsageDashboardMajorVersion = `USAGE_DASHBOARD_MAJOR_VERSION_2`

// String representation for [fmt.Print]
func (f *UsageDashboardMajorVersion) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UsageDashboardMajorVersion) Set(v string) error {
	switch v {
	case `USAGE_DASHBOARD_MAJOR_VERSION_1`, `USAGE_DASHBOARD_MAJOR_VERSION_2`:
		*f = UsageDashboardMajorVersion(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "USAGE_DASHBOARD_MAJOR_VERSION_1", "USAGE_DASHBOARD_MAJOR_VERSION_2"`, v)
	}
}

// Values returns all possible values for UsageDashboardMajorVersion.
//
// There is no guarantee on the order of the values in the slice.
func (f *UsageDashboardMajorVersion) Values() []UsageDashboardMajorVersion {
	return []UsageDashboardMajorVersion{
		UsageDashboardMajorVersionUsageDashboardMajorVersion1,
		UsageDashboardMajorVersionUsageDashboardMajorVersion2,
	}
}

// Type always returns UsageDashboardMajorVersion to satisfy [pflag.Value] interface
func (f *UsageDashboardMajorVersion) Type() string {
	return "UsageDashboardMajorVersion"
}

type UsageDashboardType string

const UsageDashboardTypeUsageDashboardTypeGlobal UsageDashboardType = `USAGE_DASHBOARD_TYPE_GLOBAL`

const UsageDashboardTypeUsageDashboardTypeWorkspace UsageDashboardType = `USAGE_DASHBOARD_TYPE_WORKSPACE`

// String representation for [fmt.Print]
func (f *UsageDashboardType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UsageDashboardType) Set(v string) error {
	switch v {
	case `USAGE_DASHBOARD_TYPE_GLOBAL`, `USAGE_DASHBOARD_TYPE_WORKSPACE`:
		*f = UsageDashboardType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "USAGE_DASHBOARD_TYPE_GLOBAL", "USAGE_DASHBOARD_TYPE_WORKSPACE"`, v)
	}
}

// Values returns all possible values for UsageDashboardType.
//
// There is no guarantee on the order of the values in the slice.
func (f *UsageDashboardType) Values() []UsageDashboardType {
	return []UsageDashboardType{
		UsageDashboardTypeUsageDashboardTypeGlobal,
		UsageDashboardTypeUsageDashboardTypeWorkspace,
	}
}

// Type always returns UsageDashboardType to satisfy [pflag.Value] interface
func (f *UsageDashboardType) Type() string {
	return "UsageDashboardType"
}

// * Properties of the new log delivery configuration.
type WrappedCreateLogDeliveryConfiguration struct {
	LogDeliveryConfiguration CreateLogDeliveryConfigurationParams `json:"log_delivery_configuration"`
}

func (s *WrappedCreateLogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type WrappedLogDeliveryConfiguration struct {
	// The created log delivery configuration
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

func (s *WrappedLogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type WrappedLogDeliveryConfigurations struct {
	// Log delivery configurations were returned successfully.
	LogDeliveryConfigurations []LogDeliveryConfiguration `json:"log_delivery_configurations,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WrappedLogDeliveryConfigurations) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WrappedLogDeliveryConfigurations) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
