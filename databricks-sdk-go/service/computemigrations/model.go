// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package computemigrations

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// A request to approve migration execution at the MIGRATION_REVIEW gate.
type ApproveMigrationExecutionRequest struct {
	// Optional free-text feedback or instructions the user provides along with
	// the approval, carried into the next stage of the migration.
	AdditionalFeedback string `json:"additional_feedback,omitempty"`
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag string `json:"etag,omitempty"`
	// The ID of the migration to approve for execution.
	MigrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApproveMigrationExecutionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApproveMigrationExecutionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A request to approve the working copy at the TRANSLATION_READY gate.
type ApproveMigrationTranslationRequest struct {
	// Optional free-text feedback or instructions the user provides along with
	// the approval, carried into the next stage of the migration.
	AdditionalFeedback string `json:"additional_feedback,omitempty"`
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag string `json:"etag,omitempty"`
	// The ID of the migration to approve.
	MigrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApproveMigrationTranslationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApproveMigrationTranslationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for the baking window (post-execution production monitoring).
type BakingConfig struct {
	// Number of successful production runs required before the migration
	// reaches MIGRATION_COMPLETE. Default applied by the server (configurable).
	RunCount int `json:"run_count,omitempty"`
	// Maximum baking window duration in hours. If the required run_count has
	// not been reached within this window, the migration transitions to FAILED.
	// Default applied by the server (configurable).
	TimeWindowHours int `json:"time_window_hours,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BakingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BakingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Summary of baking metrics after the baking window completes.
type BakingReport struct {
	// Number of successful runs.
	SuccessfulRuns int `json:"successful_runs,omitempty"`
	// Number of runs observed during the baking window.
	TotalRuns int `json:"total_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BakingReport) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BakingReport) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A request to cancel a migration from any non-terminal state.
type CancelMigrationRequest struct {
	// Optional free-text reason for the cancellation.
	CancellationReason string `json:"cancellation_reason,omitempty"`
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag string `json:"etag,omitempty"`
	// The ID of the migration to cancel.
	MigrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CancelMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CancelMigrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateMigrationRequest struct {
	// The migration to create. The server generates `migration_id` and sets
	// `state` to PLANNING_PENDING. Required fields: `asset_type`, `asset_uri`.
	Migration Migration `json:"migration"`
	// Client-generated request ID for idempotent request deduplication. The
	// server de-duplicates requests with the same request_id for at least 24
	// hours, returning the existing migration instead of creating a duplicate.
	RequestId string `json:"-" url:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMigrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteMigrationRequest struct {
	// If nonempty, must match the current resource etag or the request fails
	// with ABORTED.
	Etag string `json:"-" url:"etag,omitempty"`
	// The ID of the migration to cancel.
	MigrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteMigrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetMigrationRequest struct {
	// The ID of the migration to retrieve.
	MigrationId string `json:"-" url:"-"`
}

func (s *GetMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListMigrationsRequest struct {
	// Filter by asset type.
	AssetType MigrationAssetType `json:"-" url:"asset_type,omitempty"`
	// Maximum number of migrations to return. Default: 50, maximum: 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous ListMigrations call. When paginating, all
	// other parameters must match the call that provided the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Filter by migration state. Returns migrations in any of the specified
	// states.
	States []MigrationState `json:"-" url:"states,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMigrationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMigrationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing a paginated list of migrations matching the request
// filters.
type ListMigrationsResponse struct {
	// The migrations matching the request filters.
	Migrations []Migration `json:"migrations,omitempty"`
	// Token for the next page. Omitted if there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMigrationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMigrationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A persistent object tracking the full lifecycle of migrating one Classic
// compute asset to Serverless. Each Migration has a well-defined state machine
// and is the unit of orchestration, approval, and rollback.
type Migration struct {
	// Type of Classic compute asset being migrated. Cannot be changed after
	// creation.
	AssetType MigrationAssetType `json:"asset_type"`
	// URI of the source asset (e.g., "jobs/12345"). Cannot be changed after
	// creation.
	AssetUri string `json:"asset_uri"`
	// When true, the server skips the translation review (TRANSLATION_READY)
	// instead of waiting for an explicit ApproveMigrationTranslation.
	// MIGRATION_REVIEW is never skipped. Unset/false means the translation
	// review requires explicit user action. Set at creation; cannot be changed
	// afterward. Does not suppress cancellation or terminal failures.
	AutoApprove bool `json:"auto_approve,omitempty"`
	// Baking window configuration. Mutable via UpdateMigration before baking
	// starts.
	BakingConfig *BakingConfig `json:"baking_config,omitempty"`
	// Summary of baking metrics after the baking window completes.
	BakingReport *BakingReport `json:"baking_report,omitempty"`
	// Timestamp when the migration was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Resource version for optimistic concurrency control. Clients read this
	// from GetMigration or ListMigrations and pass it back as the request
	// `etag` on update, delete, approve, reject, and cancel methods.
	Etag string `json:"etag,omitempty"`
	// The user who initiated this migration and their role. Set at creation.
	Initiator *MigrationInitiator `json:"initiator,omitempty"`
	// Server-generated unique identifier for this migration.
	MigrationId string `json:"migration_id,omitempty"`
	// Agent-produced planning report, presented to the user at the review
	// gates.
	MigrationReport *MigrationReport `json:"migration_report,omitempty"`
	// When true, Databricks runs the working copy to validate it after the
	// translation review. Unset/false skips the run, so the migration goes from
	// the translation review straight to execution review. Whether the
	// translation review waits for the user is controlled by `auto_approve`,
	// not this field. Set at creation; cannot be changed afterward.
	RunValidation bool `json:"run_validation,omitempty"`
	// Current state of the migration in the state machine. Managed by the
	// server.
	State MigrationState `json:"state,omitempty"`
	// History of state transitions, ordered chronologically. Each entry records
	// the state entered and the timestamp of the transition.
	StateHistory []MigrationStateTransition `json:"state_history,omitempty"`
	// Agent-produced validation results from sandbox execution.
	ValidationReport *ValidationReport `json:"validation_report,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Migration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Migration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A working copy (job or notebook) the migration agent created for the
// migration.
type MigrationAsset struct {
	// Type of asset this working copy represents.
	AssetType MigrationAssetType `json:"asset_type,omitempty"`
	// Identifier of the asset.
	Id string `json:"id,omitempty"`
	// Link to the asset, e.g. "jobs/12345".
	Uri string `json:"uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigrationAsset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrationAsset) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of a migration asset: the Classic asset being migrated, or an asset the
// migration agent created for it. Migrations can only be created for JOB
// assets.
type MigrationAssetType string

const MigrationAssetTypeFile MigrationAssetType = `FILE`

const MigrationAssetTypeInteractive MigrationAssetType = `INTERACTIVE`

const MigrationAssetTypeJob MigrationAssetType = `JOB`

// String representation for [fmt.Print]
func (f *MigrationAssetType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MigrationAssetType) Set(v string) error {
	switch v {
	case `FILE`, `INTERACTIVE`, `JOB`:
		*f = MigrationAssetType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE", "INTERACTIVE", "JOB"`, v)
	}
}

// Values returns all possible values for MigrationAssetType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MigrationAssetType) Values() []MigrationAssetType {
	return []MigrationAssetType{
		MigrationAssetTypeFile,
		MigrationAssetTypeInteractive,
		MigrationAssetTypeJob,
	}
}

// Type always returns MigrationAssetType to satisfy [pflag.Value] interface
func (f *MigrationAssetType) Type() string {
	return "MigrationAssetType"
}

// The initiator of a migration, including their role and user identity.
type MigrationInitiator struct {
	// The role of the initiator, determining approval authority. Set by the
	// server: the role is resolved from the caller's identity and permissions
	// on the asset, never taken from the request, since a caller must not be
	// able to assert its own approval authority.
	Role MigrationInitiatorRole `json:"role,omitempty"`
	// The user ID of the initiator. Set by the server from the request context.
	UserId int64 `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigrationInitiator) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrationInitiator) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Role of the migration initiator, determining approval authority.
type MigrationInitiatorRole string

const MigrationInitiatorRoleAccountAdmin MigrationInitiatorRole = `ACCOUNT_ADMIN`

const MigrationInitiatorRoleJobAdmin MigrationInitiatorRole = `JOB_ADMIN`

const MigrationInitiatorRoleSystem MigrationInitiatorRole = `SYSTEM`

const MigrationInitiatorRoleWorkspaceAdmin MigrationInitiatorRole = `WORKSPACE_ADMIN`

// String representation for [fmt.Print]
func (f *MigrationInitiatorRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MigrationInitiatorRole) Set(v string) error {
	switch v {
	case `ACCOUNT_ADMIN`, `JOB_ADMIN`, `SYSTEM`, `WORKSPACE_ADMIN`:
		*f = MigrationInitiatorRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCOUNT_ADMIN", "JOB_ADMIN", "SYSTEM", "WORKSPACE_ADMIN"`, v)
	}
}

// Values returns all possible values for MigrationInitiatorRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *MigrationInitiatorRole) Values() []MigrationInitiatorRole {
	return []MigrationInitiatorRole{
		MigrationInitiatorRoleAccountAdmin,
		MigrationInitiatorRoleJobAdmin,
		MigrationInitiatorRoleSystem,
		MigrationInitiatorRoleWorkspaceAdmin,
	}
}

// Type always returns MigrationInitiatorRole to satisfy [pflag.Value] interface
func (f *MigrationInitiatorRole) Type() string {
	return "MigrationInitiatorRole"
}

// Agent-produced planning report, presented to the user at the review gates.
type MigrationReport struct {
	// The primary asset the agent created — the entry point / validation
	// target.
	PrimaryAsset *MigrationAsset `json:"primary_asset,omitempty"`
	// Additional assets the agent created (e.g. notebooks or jobs the primary
	// asset invokes).
	SecondaryAssets []MigrationAsset `json:"secondary_assets,omitempty"`
}

func (s *MigrationReport) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The state of a migration in its lifecycle.
type MigrationState string

const MigrationStateCancelled MigrationState = `CANCELLED`

const MigrationStateFailed MigrationState = `FAILED`

const MigrationStateMigrationComplete MigrationState = `MIGRATION_COMPLETE`

const MigrationStateMigrationObservation MigrationState = `MIGRATION_OBSERVATION`

const MigrationStateMigrationPending MigrationState = `MIGRATION_PENDING`

const MigrationStateMigrationReview MigrationState = `MIGRATION_REVIEW`

const MigrationStateMigrationRunning MigrationState = `MIGRATION_RUNNING`

const MigrationStateNotSupported MigrationState = `NOT_SUPPORTED`

const MigrationStatePlanningPending MigrationState = `PLANNING_PENDING`

const MigrationStatePlanningReview MigrationState = `PLANNING_REVIEW`

const MigrationStatePlanningRunning MigrationState = `PLANNING_RUNNING`

const MigrationStateRevisionPending MigrationState = `REVISION_PENDING`

const MigrationStateTranslationCreating MigrationState = `TRANSLATION_CREATING`

const MigrationStateTranslationPending MigrationState = `TRANSLATION_PENDING`

const MigrationStateTranslationReady MigrationState = `TRANSLATION_READY`

const MigrationStateValidationRunning MigrationState = `VALIDATION_RUNNING`

const MigrationStateValidationScheduling MigrationState = `VALIDATION_SCHEDULING`

// String representation for [fmt.Print]
func (f *MigrationState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MigrationState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `MIGRATION_COMPLETE`, `MIGRATION_OBSERVATION`, `MIGRATION_PENDING`, `MIGRATION_REVIEW`, `MIGRATION_RUNNING`, `NOT_SUPPORTED`, `PLANNING_PENDING`, `PLANNING_REVIEW`, `PLANNING_RUNNING`, `REVISION_PENDING`, `TRANSLATION_CREATING`, `TRANSLATION_PENDING`, `TRANSLATION_READY`, `VALIDATION_RUNNING`, `VALIDATION_SCHEDULING`:
		*f = MigrationState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "MIGRATION_COMPLETE", "MIGRATION_OBSERVATION", "MIGRATION_PENDING", "MIGRATION_REVIEW", "MIGRATION_RUNNING", "NOT_SUPPORTED", "PLANNING_PENDING", "PLANNING_REVIEW", "PLANNING_RUNNING", "REVISION_PENDING", "TRANSLATION_CREATING", "TRANSLATION_PENDING", "TRANSLATION_READY", "VALIDATION_RUNNING", "VALIDATION_SCHEDULING"`, v)
	}
}

// Values returns all possible values for MigrationState.
//
// There is no guarantee on the order of the values in the slice.
func (f *MigrationState) Values() []MigrationState {
	return []MigrationState{
		MigrationStateCancelled,
		MigrationStateFailed,
		MigrationStateMigrationComplete,
		MigrationStateMigrationObservation,
		MigrationStateMigrationPending,
		MigrationStateMigrationReview,
		MigrationStateMigrationRunning,
		MigrationStateNotSupported,
		MigrationStatePlanningPending,
		MigrationStatePlanningReview,
		MigrationStatePlanningRunning,
		MigrationStateRevisionPending,
		MigrationStateTranslationCreating,
		MigrationStateTranslationPending,
		MigrationStateTranslationReady,
		MigrationStateValidationRunning,
		MigrationStateValidationScheduling,
	}
}

// Type always returns MigrationState to satisfy [pflag.Value] interface
func (f *MigrationState) Type() string {
	return "MigrationState"
}

// A record of a state transition in the migration lifecycle.
type MigrationStateTransition struct {
	// Free-text the service recorded for this transition — e.g. an
	// internally-determined cancellation or failure reason. Empty for
	// user-driven transitions.
	InternalComment string `json:"internal_comment,omitempty"`
	// The state that was entered.
	State MigrationState `json:"state,omitempty"`
	// Timestamp when this state was entered.
	TransitionTime *time.Time `json:"transition_time,omitempty"`
	// Free-text the initiating user supplied at a user-driven transition —
	// reviewer feedback on an approval, or the reason for a rejection or
	// cancellation. Empty for transitions the service drives internally.
	UserComment string `json:"user_comment,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigrationStateTransition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrationStateTransition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A request to reject migration execution at the MIGRATION_REVIEW gate, looping
// back to REVISION_PENDING.
type RejectMigrationExecutionRequest struct {
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag string `json:"etag,omitempty"`
	// The ID of the migration to reject.
	MigrationId string `json:"-" url:"-"`
	// Required free-text reason for the rejection, fed back into re-planning so
	// the next plan can address the user's concerns.
	RejectionReason string `json:"rejection_reason"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RejectMigrationExecutionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RejectMigrationExecutionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A request to reject the working copy at the TRANSLATION_READY gate, looping
// back to REVISION_PENDING.
type RejectMigrationTranslationRequest struct {
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag string `json:"etag,omitempty"`
	// The ID of the migration to reject.
	MigrationId string `json:"-" url:"-"`
	// Required free-text reason for the rejection, fed back into re-planning so
	// the next plan can address the user's concerns.
	RejectionReason string `json:"rejection_reason"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RejectMigrationTranslationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RejectMigrationTranslationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Validation results from running the translated workload in the sandbox.
type ValidationReport struct {
	// Duration of the sandbox validation run in milliseconds. Pair with
	// start_time to derive the end time.
	DurationMs int64 `json:"duration_ms,omitempty"`
	// Databricks-internal recommendation on whether to proceed with the
	// migration, based on the validation run. May differ from run_status: a run
	// can succeed while the recommendation is negative, or vice versa.
	Passed bool `json:"passed,omitempty"`
	// Status of the validation run itself (the objective run outcome), as
	// reported by the run system, e.g. "SUCCESS" / "FAILED" / "TIMED_OUT".
	RunStatus string `json:"run_status,omitempty"`
	// Link to the validation run (job or notebook).
	RunUrl string `json:"run_url,omitempty"`
	// Timestamp when validation started in the sandbox.
	StartTime *time.Time `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ValidationReport) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidationReport) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
