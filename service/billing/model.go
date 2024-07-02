// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// Budget configuration to be created.
type Budget struct {
	Alerts []BudgetAlert `tfsdk:"alerts"`
	// Optional end date of the budget.
	EndDate string `tfsdk:"end_date"`
	// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
	// account that matches this expression will be counted in this budget.
	//
	// Supported properties on left-hand side of comparison: * `workspaceId` -
	// the ID of the workspace * `sku` - SKU of the cluster, e.g.
	// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.'tag name'` - tag of
	// the cluster
	//
	// Supported comparison operators: * `=` - equal * `!=` - not equal
	//
	// Supported logical operators: `AND`, `OR`.
	//
	// Examples: * `workspaceId=123 OR (sku='STANDARD_ALL_PURPOSE_COMPUTE' AND
	// tag.'my tag'='my value')` * `workspaceId!=456` *
	// `sku='STANDARD_ALL_PURPOSE_COMPUTE' OR sku='PREMIUM_ALL_PURPOSE_COMPUTE'`
	// * `tag.name1='value1' AND tag.name2='value2'`
	Filter string `tfsdk:"filter"`
	// Human-readable name of the budget.
	Name string `tfsdk:"name"`
	// Period length in years, months, weeks and/or days. Examples: `1 month`,
	// `30 days`, `1 year, 2 months, 1 week, 2 days`
	Period string `tfsdk:"period"`
	// Start date of the budget period calculation.
	StartDate string `tfsdk:"start_date"`
	// Target amount of the budget per period in USD.
	TargetAmount string `tfsdk:"target_amount"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Budget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Budget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BudgetAlert struct {
	// List of email addresses to be notified when budget percentage is exceeded
	// in the given period.
	EmailNotifications []string `tfsdk:"email_notifications"`
	// Percentage of the target amount used in the currect period that will
	// trigger a notification.
	MinPercentage int `tfsdk:"min_percentage"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *BudgetAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List of budgets.
type BudgetList struct {
	Budgets []BudgetWithStatus `tfsdk:"budgets"`
}

// Budget configuration with daily status.
type BudgetWithStatus struct {
	Alerts []BudgetAlert `tfsdk:"alerts"`

	BudgetId string `tfsdk:"budget_id"`

	CreationTime string `tfsdk:"creation_time"`
	// Optional end date of the budget.
	EndDate string `tfsdk:"end_date"`
	// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
	// account that matches this expression will be counted in this budget.
	//
	// Supported properties on left-hand side of comparison: * `workspaceId` -
	// the ID of the workspace * `sku` - SKU of the cluster, e.g.
	// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.'tag name'` - tag of
	// the cluster
	//
	// Supported comparison operators: * `=` - equal * `!=` - not equal
	//
	// Supported logical operators: `AND`, `OR`.
	//
	// Examples: * `workspaceId=123 OR (sku='STANDARD_ALL_PURPOSE_COMPUTE' AND
	// tag.'my tag'='my value')` * `workspaceId!=456` *
	// `sku='STANDARD_ALL_PURPOSE_COMPUTE' OR sku='PREMIUM_ALL_PURPOSE_COMPUTE'`
	// * `tag.name1='value1' AND tag.name2='value2'`
	Filter string `tfsdk:"filter"`
	// Human-readable name of the budget.
	Name string `tfsdk:"name"`
	// Period length in years, months, weeks and/or days. Examples: `1 month`,
	// `30 days`, `1 year, 2 months, 1 week, 2 days`
	Period string `tfsdk:"period"`
	// Start date of the budget period calculation.
	StartDate string `tfsdk:"start_date"`
	// Amount used in the budget for each day (noncumulative).
	StatusDaily []BudgetWithStatusStatusDailyItem `tfsdk:"status_daily"`
	// Target amount of the budget per period in USD.
	TargetAmount string `tfsdk:"target_amount"`

	UpdateTime string `tfsdk:"update_time"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *BudgetWithStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetWithStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BudgetWithStatusStatusDailyItem struct {
	// Amount used in this day in USD.
	Amount string `tfsdk:"amount"`

	Date string `tfsdk:"date"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *BudgetWithStatusStatusDailyItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetWithStatusStatusDailyItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateLogDeliveryConfigurationParams struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `tfsdk:"config_name"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId string `tfsdk:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `tfsdk:"delivery_path_prefix"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime string `tfsdk:"delivery_start_time"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType LogType `tfsdk:"log_type"`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat OutputFormat `tfsdk:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `tfsdk:"status"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId string `tfsdk:"storage_configuration_id"`
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
	WorkspaceIdsFilter []int64 `tfsdk:"workspace_ids_filter"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateLogDeliveryConfigurationParams) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateLogDeliveryConfigurationParams) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete budget
type DeleteBudgetRequest struct {
	// Budget ID
	BudgetId string `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// The status string for log delivery. Possible values are: * `CREATED`: There
// were no log delivery attempts since the config was created. * `SUCCEEDED`:
// The latest attempt of log delivery has succeeded completely. *
// `USER_FAILURE`: The latest attempt of log delivery failed because of
// misconfiguration of customer provided permissions on role or storage. *
// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
// Databricks internal error. Contact support if it doesn't go away soon. *
// `NOT_FOUND`: The log delivery status as the configuration has been disabled
// since the release of this feature or there are no workspaces in the account.
type DeliveryStatus string

// There were no log delivery attempts since the config was created.
const DeliveryStatusCreated DeliveryStatus = `CREATED`

// The log delivery status as the configuration has been disabled since the
// release of this feature or there are no workspaces in the account.
const DeliveryStatusNotFound DeliveryStatus = `NOT_FOUND`

// The latest attempt of log delivery has succeeded completely.
const DeliveryStatusSucceeded DeliveryStatus = `SUCCEEDED`

// The latest attempt of log delivery failed because of an <Databricks> internal
// error. Contact support if it doesn't go away soon.
const DeliveryStatusSystemFailure DeliveryStatus = `SYSTEM_FAILURE`

// The latest attempt of log delivery failed because of misconfiguration of
// customer provided permissions on role or storage.
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

// Type always returns DeliveryStatus to satisfy [pflag.Value] interface
func (f *DeliveryStatus) Type() string {
	return "DeliveryStatus"
}

// Return billable usage logs
type DownloadRequest struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth string `tfsdk:"-" url:"end_month"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData bool `tfsdk:"-" url:"personal_data,omitempty"`
	// Format: `YYYY-MM`. First month to return billable usage logs for. This
	// field is required.
	StartMonth string `tfsdk:"-" url:"start_month"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DownloadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DownloadRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DownloadResponse struct {
	Contents io.ReadCloser `tfsdk:"-"`
}

// Get budget and its status
type GetBudgetRequest struct {
	// Budget ID
	BudgetId string `tfsdk:"-" url:"-"`
}

// Get log delivery configuration
type GetLogDeliveryRequest struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId string `tfsdk:"-" url:"-"`
}

// Get all log delivery configurations
type ListLogDeliveryRequest struct {
	// Filter by credential configuration ID.
	CredentialsId string `tfsdk:"-" url:"credentials_id,omitempty"`
	// Filter by status `ENABLED` or `DISABLED`.
	Status LogDeliveryConfigStatus `tfsdk:"-" url:"status,omitempty"`
	// Filter by storage configuration ID.
	StorageConfigurationId string `tfsdk:"-" url:"storage_configuration_id,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListLogDeliveryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of log delivery configuration. Set to `ENABLED` (enabled) or
// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable the
// configuration](#operation/patch-log-delivery-config-status) later. Deletion
// of a configuration is not supported, so disable a log delivery configuration
// that is no longer needed.
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

// Type always returns LogDeliveryConfigStatus to satisfy [pflag.Value] interface
func (f *LogDeliveryConfigStatus) Type() string {
	return "LogDeliveryConfigStatus"
}

type LogDeliveryConfiguration struct {
	// The Databricks account ID that hosts the log delivery configuration.
	AccountId string `tfsdk:"account_id"`
	// Databricks log delivery configuration ID.
	ConfigId string `tfsdk:"config_id"`
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `tfsdk:"config_name"`
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	CreationTime int64 `tfsdk:"creation_time"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId string `tfsdk:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `tfsdk:"delivery_path_prefix"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime string `tfsdk:"delivery_start_time"`
	// Databricks log delivery status.
	LogDeliveryStatus *LogDeliveryStatus `tfsdk:"log_delivery_status"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType LogType `tfsdk:"log_type"`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat OutputFormat `tfsdk:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `tfsdk:"status"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId string `tfsdk:"storage_configuration_id"`
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	UpdateTime int64 `tfsdk:"update_time"`
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
	WorkspaceIdsFilter []int64 `tfsdk:"workspace_ids_filter"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *LogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks log delivery status.
type LogDeliveryStatus struct {
	// The UTC time for the latest log delivery attempt.
	LastAttemptTime string `tfsdk:"last_attempt_time"`
	// The UTC time for the latest successful log delivery.
	LastSuccessfulAttemptTime string `tfsdk:"last_successful_attempt_time"`
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	Message string `tfsdk:"message"`
	// The status string for log delivery. Possible values are: * `CREATED`:
	// There were no log delivery attempts since the config was created. *
	// `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
	// * `USER_FAILURE`: The latest attempt of log delivery failed because of
	// misconfiguration of customer provided permissions on role or storage. *
	// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
	// Databricks internal error. Contact support if it doesn't go away soon. *
	// `NOT_FOUND`: The log delivery status as the configuration has been
	// disabled since the release of this feature or there are no workspaces in
	// the account.
	Status DeliveryStatus `tfsdk:"status"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *LogDeliveryStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogDeliveryStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Log delivery type. Supported values are:
//
// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the CSV
// schema, see the [View billable usage].
//
// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema, see
// [Configure audit logging]
//
// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
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

// Type always returns LogType to satisfy [pflag.Value] interface
func (f *LogType) Type() string {
	return "LogType"
}

// The file type of log delivery.
//
// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the CSV
// (comma-separated values) format is supported. For the schema, see the [View
// billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be `JSON`.
// Only the JSON (JavaScript Object Notation) format is supported. For the
// schema, see the [Configuring audit logs].
//
// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
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

// Type always returns OutputFormat to satisfy [pflag.Value] interface
func (f *OutputFormat) Type() string {
	return "OutputFormat"
}

type PatchStatusResponse struct {
}

type UpdateLogDeliveryConfigurationStatusRequest struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId string `tfsdk:"-" url:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `tfsdk:"status"`
}

type UpdateResponse struct {
}

type WrappedBudget struct {
	// Budget configuration to be created.
	Budget Budget `tfsdk:"budget"`
	// Budget ID
	BudgetId string `tfsdk:"-" url:"-"`
}

type WrappedBudgetWithStatus struct {
	// Budget configuration with daily status.
	Budget BudgetWithStatus `tfsdk:"budget"`
}

type WrappedCreateLogDeliveryConfiguration struct {
	LogDeliveryConfiguration *CreateLogDeliveryConfigurationParams `tfsdk:"log_delivery_configuration"`
}

type WrappedLogDeliveryConfiguration struct {
	LogDeliveryConfiguration *LogDeliveryConfiguration `tfsdk:"log_delivery_configuration"`
}

type WrappedLogDeliveryConfigurations struct {
	LogDeliveryConfigurations []LogDeliveryConfiguration `tfsdk:"log_delivery_configurations"`
}
