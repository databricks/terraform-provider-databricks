// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Request to acknowledge assessment warnings.
type AcknowledgeAssessmentWarningsRequest struct {
	// Required. The assessment resource name, in the format
	// `migrations/{migration}/assessments/{assessment}`.
	Name string `json:"-" url:"-"`
}

func (s *AcknowledgeAssessmentWarningsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response produced after assessment warnings are acknowledged.
type AcknowledgeAssessmentWarningsResponse struct {
	// The assessment after its warnings are acknowledged.
	Assessment *Assessment `json:"assessment,omitempty"`
}

func (s *AcknowledgeAssessmentWarningsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to assess a migration.
type AssessMigrationRequest struct {
	// Required. The migration resource name, in the format
	// `migrations/{migration}`.
	Name string `json:"-" url:"-"`
}

func (s *AssessMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response produced after an assessment run is created.
type AssessMigrationResponse struct {
	// The assessment run created for the migration.
	Assessment *Assessment `json:"assessment,omitempty"`
}

func (s *AssessMigrationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A server-managed assessment of a migration's source configuration. Completed
// check results are immutable.
type Assessment struct {
	// Output-only. Results of the checks performed by this assessment.
	Checks []AssessmentCheck `json:"checks,omitempty"`
	// Output-only. The time when the assessment completed.
	CompleteTime *time.Time `json:"complete_time,omitempty"`
	// Output-only. The time when the assessment started.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output-only. The time when the assessment expires.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// The assessment resource name, in the format
	// `migrations/{migration}/assessments/{assessment}`.
	Name string `json:"name,omitempty"`
	// Output-only. The aggregate assessment state.
	State AssessmentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Assessment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Assessment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The result of one assessment check.
type AssessmentCheck struct {
	// Output-only. Whether this result blocks migration from starting.
	BlocksMigration bool `json:"blocks_migration,omitempty"`
	// Output-only. Stable identifier of the check.
	CheckId string `json:"check_id,omitempty"`
	// Output-only. The source database the check evaluated. Unset for checks
	// that assess the whole source server rather than a single database.
	Database string `json:"database,omitempty"`
	// Output-only. Measurements observed by this check. Absent when
	// unavailable.
	Details *AssessmentCheckDetails `json:"details,omitempty"`
	// Output-only. Explanation of the result.
	Reason string `json:"reason,omitempty"`
	// Output-only. The check verdict.
	Verdict AssessmentCheckVerdict `json:"verdict,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AssessmentCheck) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AssessmentCheck) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Measurements observed when an assessment check ran. Unset for older
// assessments.
type AssessmentCheckDetails struct {
	// Output-only. Physical size of the source database observed by this check.
	DatabaseSize *DatabaseSizeDetails `json:"database_size,omitempty"`
	// Output-only. Age of the oldest source transaction observed by this check.
	LongTransaction *LongTransactionDetails `json:"long_transaction,omitempty"`
}

func (s *AssessmentCheckDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The verdict of a single assessment check.
type AssessmentCheckVerdict string

const AssessmentCheckVerdictAssessmentCheckVerdictActionRequired AssessmentCheckVerdict = `ASSESSMENT_CHECK_VERDICT_ACTION_REQUIRED`

const AssessmentCheckVerdictAssessmentCheckVerdictError AssessmentCheckVerdict = `ASSESSMENT_CHECK_VERDICT_ERROR`

const AssessmentCheckVerdictAssessmentCheckVerdictPass AssessmentCheckVerdict = `ASSESSMENT_CHECK_VERDICT_PASS`

const AssessmentCheckVerdictAssessmentCheckVerdictUnsupported AssessmentCheckVerdict = `ASSESSMENT_CHECK_VERDICT_UNSUPPORTED`

const AssessmentCheckVerdictAssessmentCheckVerdictWarning AssessmentCheckVerdict = `ASSESSMENT_CHECK_VERDICT_WARNING`

// String representation for [fmt.Print]
func (f *AssessmentCheckVerdict) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AssessmentCheckVerdict) Set(v string) error {
	switch v {
	case `ASSESSMENT_CHECK_VERDICT_ACTION_REQUIRED`, `ASSESSMENT_CHECK_VERDICT_ERROR`, `ASSESSMENT_CHECK_VERDICT_PASS`, `ASSESSMENT_CHECK_VERDICT_UNSUPPORTED`, `ASSESSMENT_CHECK_VERDICT_WARNING`:
		*f = AssessmentCheckVerdict(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASSESSMENT_CHECK_VERDICT_ACTION_REQUIRED", "ASSESSMENT_CHECK_VERDICT_ERROR", "ASSESSMENT_CHECK_VERDICT_PASS", "ASSESSMENT_CHECK_VERDICT_UNSUPPORTED", "ASSESSMENT_CHECK_VERDICT_WARNING"`, v)
	}
}

// Values returns all possible values for AssessmentCheckVerdict.
//
// There is no guarantee on the order of the values in the slice.
func (f *AssessmentCheckVerdict) Values() []AssessmentCheckVerdict {
	return []AssessmentCheckVerdict{
		AssessmentCheckVerdictAssessmentCheckVerdictActionRequired,
		AssessmentCheckVerdictAssessmentCheckVerdictError,
		AssessmentCheckVerdictAssessmentCheckVerdictPass,
		AssessmentCheckVerdictAssessmentCheckVerdictUnsupported,
		AssessmentCheckVerdictAssessmentCheckVerdictWarning,
	}
}

// Type always returns AssessmentCheckVerdict to satisfy [pflag.Value] interface
func (f *AssessmentCheckVerdict) Type() string {
	return "AssessmentCheckVerdict"
}

// The aggregate state of an assessment run.
type AssessmentState string

const AssessmentStateAssessmentStateAssessing AssessmentState = `ASSESSMENT_STATE_ASSESSING`

const AssessmentStateAssessmentStateAwaitingAck AssessmentState = `ASSESSMENT_STATE_AWAITING_ACK`

const AssessmentStateAssessmentStateBlocked AssessmentState = `ASSESSMENT_STATE_BLOCKED`

const AssessmentStateAssessmentStateExpired AssessmentState = `ASSESSMENT_STATE_EXPIRED`

const AssessmentStateAssessmentStateFailed AssessmentState = `ASSESSMENT_STATE_FAILED`

const AssessmentStateAssessmentStateNeedsAction AssessmentState = `ASSESSMENT_STATE_NEEDS_ACTION`

const AssessmentStateAssessmentStateReady AssessmentState = `ASSESSMENT_STATE_READY`

// String representation for [fmt.Print]
func (f *AssessmentState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AssessmentState) Set(v string) error {
	switch v {
	case `ASSESSMENT_STATE_ASSESSING`, `ASSESSMENT_STATE_AWAITING_ACK`, `ASSESSMENT_STATE_BLOCKED`, `ASSESSMENT_STATE_EXPIRED`, `ASSESSMENT_STATE_FAILED`, `ASSESSMENT_STATE_NEEDS_ACTION`, `ASSESSMENT_STATE_READY`:
		*f = AssessmentState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASSESSMENT_STATE_ASSESSING", "ASSESSMENT_STATE_AWAITING_ACK", "ASSESSMENT_STATE_BLOCKED", "ASSESSMENT_STATE_EXPIRED", "ASSESSMENT_STATE_FAILED", "ASSESSMENT_STATE_NEEDS_ACTION", "ASSESSMENT_STATE_READY"`, v)
	}
}

// Values returns all possible values for AssessmentState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AssessmentState) Values() []AssessmentState {
	return []AssessmentState{
		AssessmentStateAssessmentStateAssessing,
		AssessmentStateAssessmentStateAwaitingAck,
		AssessmentStateAssessmentStateBlocked,
		AssessmentStateAssessmentStateExpired,
		AssessmentStateAssessmentStateFailed,
		AssessmentStateAssessmentStateNeedsAction,
		AssessmentStateAssessmentStateReady,
	}
}

// Type always returns AssessmentState to satisfy [pflag.Value] interface
func (f *AssessmentState) Type() string {
	return "AssessmentState"
}

type Branch struct {
	// The part of the name, chosen by the user when the resource was created.
	BranchId string `json:"branch_id,omitempty"`
	// A timestamp indicating when the branch was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The full resource path of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"name,omitempty"`
	// The project containing this branch (API resource hierarchy). Format:
	// projects/{project_id}
	//
	// Note: This field indicates where the branch exists in the resource
	// hierarchy. For point-in-time branching from another branch, see
	// `status.source_branch`.
	Parent string `json:"parent,omitempty"`
	// The spec contains the branch configuration.
	//
	// Read the resource's current values from `status`, rather than `spec`.
	// `spec` tracks field ownership for infrastructure-as-code tools and may
	// omit unmanaged fields.
	Spec *BranchSpec `json:"spec,omitempty"`
	// The current status of a Branch.
	Status *BranchStatus `json:"status,omitempty"`
	// System-generated unique ID for the branch.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Branch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Branch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BranchOperationMetadata struct {
}

func (s *BranchOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BranchSpec struct {
	// Absolute expiration timestamp. When set, the branch will expire at this
	// time. Mutually exclusive with `ttl` and `no_expiry`.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// When set to true, protects the branch from deletion and reset. Associated
	// compute endpoints and the project cannot be deleted while the branch is
	// protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// Explicitly disable expiration. When set to true, the branch will not
	// expire. If set to false, the request is invalid; provide either ttl or
	// expire_time instead. Mutually exclusive with `expire_time` and `ttl`.
	NoExpiry bool `json:"no_expiry,omitempty"`
	// The name of the source branch from which this branch was created (data
	// lineage for point-in-time recovery). If not specified, defaults to the
	// project's default branch. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch string `json:"source_branch,omitempty"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn string `json:"source_branch_lsn,omitempty"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime *time.Time `json:"source_branch_time,omitempty"`
	// The snapshot this branch was created from. When set, the branch's data
	// comes from the snapshot rather than a source branch, so source_branch,
	// source_branch_lsn, and source_branch_time must be empty. The snapshot
	// must be AVAILABLE and belong to this branch's project. Format:
	// projects/{project_id}/snapshots/{snapshot_id}
	SourceSnapshot string `json:"source_snapshot,omitempty"`
	// Relative time-to-live duration. When set, the branch will expire at
	// creation_time + ttl. Mutually exclusive with `expire_time` and
	// `no_expiry`.
	Ttl *duration.Duration `json:"ttl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BranchSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BranchSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BranchStatus struct {
	// Part of the resource name.
	BranchId string `json:"branch_id,omitempty"`
	// The branch's state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState BranchStatusState `json:"current_state,omitempty"`
	// Whether the branch is the project's default branch.
	Default bool `json:"default,omitempty"`
	// A timestamp indicating when the branch was deleted. Empty if the branch
	// is not deleted.
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// Absolute expiration time for the branch. Empty if expiration is disabled.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// Whether the branch is protected.
	IsProtected bool `json:"is_protected,omitempty"`
	// The logical size of the branch.
	LogicalSizeBytes int64 `json:"logical_size_bytes,omitempty"`
	// The pending state of the branch, if a state transition is in progress.
	PendingState BranchStatusState `json:"pending_state,omitempty"`
	// A timestamp indicating when the branch is scheduled to be purged. Empty
	// if the branch is not deleted, otherwise set to a timestamp in the future.
	PurgeTime *time.Time `json:"purge_time,omitempty"`
	// The name of the source branch from which this branch was created. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch string `json:"source_branch,omitempty"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn string `json:"source_branch_lsn,omitempty"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime *time.Time `json:"source_branch_time,omitempty"`
	// The snapshot this branch was restored from. Set only for branches created
	// by restoring a snapshot; unset for all other branches. Format:
	// projects/{project_id}/snapshots/{snapshot_id}
	SourceSnapshot string `json:"source_snapshot,omitempty"`
	// A timestamp indicating when the `current_state` began.
	StateChangeTime *time.Time `json:"state_change_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BranchStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BranchStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the branch.
type BranchStatusState string

const BranchStatusStateArchived BranchStatusState = `ARCHIVED`

const BranchStatusStateDeleted BranchStatusState = `DELETED`

const BranchStatusStateImporting BranchStatusState = `IMPORTING`

const BranchStatusStateInit BranchStatusState = `INIT`

const BranchStatusStateReady BranchStatusState = `READY`

const BranchStatusStateResetting BranchStatusState = `RESETTING`

// String representation for [fmt.Print]
func (f *BranchStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BranchStatusState) Set(v string) error {
	switch v {
	case `ARCHIVED`, `DELETED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`:
		*f = BranchStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARCHIVED", "DELETED", "IMPORTING", "INIT", "READY", "RESETTING"`, v)
	}
}

// Values returns all possible values for BranchStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *BranchStatusState) Values() []BranchStatusState {
	return []BranchStatusState{
		BranchStatusStateArchived,
		BranchStatusStateDeleted,
		BranchStatusStateImporting,
		BranchStatusStateInit,
		BranchStatusStateReady,
		BranchStatusStateResetting,
	}
}

// Type always returns BranchStatusState to satisfy [pflag.Value] interface
func (f *BranchStatusState) Type() string {
	return "BranchStatusState"
}

type Catalog struct {
	// The ID chosen when the catalog was created, which is also its name in
	// Unity Catalog. It is the last segment of the name field.
	CatalogId string `json:"catalog_id,omitempty"`
	// A timestamp indicating when the catalog was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The full resource path of the catalog.
	//
	// Format: "catalogs/{catalog_id}".
	Name string `json:"name,omitempty"`
	// The desired state of the Catalog. The catalog has no update method, so
	// the spec is fixed when the catalog is created.
	Spec *CatalogCatalogSpec `json:"spec,omitempty"`
	// The observed state of the Catalog.
	Status *CatalogCatalogStatus `json:"status,omitempty"`
	// The immutable UUID that identifies the catalog, assigned when the catalog
	// is created.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the catalog was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Catalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Catalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The desired state of the Catalog.
type CatalogCatalogSpec struct {
	// The resource path of the branch that holds the Postgres database. The
	// branch must have a read-write compute endpoint.
	//
	// Format: projects/{project_id}/branches/{branch_id}.
	Branch string `json:"branch,omitempty"`
	// If set to true, the specified postgres_database is created on behalf of
	// the calling user if it does not already exist. In this case, the calling
	// user has a role created for them in Postgres if they do not already have
	// one.
	//
	// Defaults to false, meaning that the request fails if the specified
	// postgres_database does not already exist.
	CreateDatabaseIfMissing bool `json:"create_database_if_missing,omitempty"`
	// The name of the Postgres database inside the specified Lakebase project
	// and branch to be associated with the UC catalog. This database must
	// already exist, unless create_database_if_missing is set to true on
	// creation. The system databases postgres, databricks_system, template0 and
	// template1 cannot be registered.
	//
	// A database can only be registered with one UC catalog at a time. To
	// re-register a database with a different catalog, the existing catalog
	// must be deleted first.
	//
	// A child branch inherits the fact of parent's registration. This means the
	// same-named database in a child branch cannot be registered with a second
	// catalog while the parent's registration exists. To allow registering the
	// database of a child branch, drop and recreate the database on the child
	// branch. This removes the fact of parent's registration from this branch
	// only.
	//
	// Doing Point In Time Restore (PITR) prior to the moment before the
	// Postgres DB was registered in the Catalog drops the fact of registration
	// of the database. So the user should avoid doing so.
	PostgresDatabase string `json:"postgres_database"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogCatalogSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogCatalogSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The observed state of the Catalog.
type CatalogCatalogStatus struct {
	// The resource path of the branch associated with the catalog.
	//
	// Format: projects/{project_id}/branches/{branch_id}.
	Branch string `json:"branch,omitempty"`
	// The name of the Postgres database associated with the catalog.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// The resource path of the project associated with the catalog.
	//
	// Format: projects/{project_id}.
	Project string `json:"project,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogCatalogStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogCatalogStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CatalogOperationMetadata struct {
}

func (s *CatalogOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Lakebase CDF configuration (CdfConfig): one per Postgres schema per
// database, replicating that schema's tables into a Unity Catalog schema.
// Immutable once created.
type CdfConfig struct {
	// The Unity Catalog catalog that replicated tables are written into. Set at
	// creation; the CdfConfig is immutable.
	Catalog string `json:"catalog"`
	// The user-specified id; equals the final segment of `name`. Defaults to
	// the Postgres schema name for configs without an explicit id.
	CdfConfigId string `json:"cdf_config_id,omitempty"`
	// When the CdfConfig was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The full resource name of the CdfConfig. Format:
	// projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
	Name string `json:"name,omitempty"`
	// The Postgres schema this CdfConfig replicates from. Unique within the
	// parent database. Set at creation; the CdfConfig is immutable.
	PostgresSchema string `json:"postgres_schema"`
	// The Unity Catalog schema that replicated tables are written into. Set at
	// creation; the CdfConfig is immutable.
	Schema string `json:"schema"`
	// The service principal (application id) whose identity authorizes this
	// change data feed's ongoing writes to Unity Catalog. When unset, writes
	// are authorized as the creating identity. Changing it requires a new
	// CdfConfig, which writes to new Delta tables.
	ServicePrincipal string `json:"service_principal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CdfConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CdfConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for CdfConfig long-running operations. Intentionally empty today;
// fields (e.g. progress) may be added as the operation contract grows.
type CdfConfigOperationMetadata struct {
}

func (s *CdfConfigOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The replication state of a single replicated table (CdfStatus).
type CdfState string

const CdfStateCdfStateSkipped CdfState = `CDF_STATE_SKIPPED`

const CdfStateCdfStateSnapshotting CdfState = `CDF_STATE_SNAPSHOTTING`

const CdfStateCdfStateStreaming CdfState = `CDF_STATE_STREAMING`

const CdfStateCdfStateTerminated CdfState = `CDF_STATE_TERMINATED`

// String representation for [fmt.Print]
func (f *CdfState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CdfState) Set(v string) error {
	switch v {
	case `CDF_STATE_SKIPPED`, `CDF_STATE_SNAPSHOTTING`, `CDF_STATE_STREAMING`, `CDF_STATE_TERMINATED`:
		*f = CdfState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CDF_STATE_SKIPPED", "CDF_STATE_SNAPSHOTTING", "CDF_STATE_STREAMING", "CDF_STATE_TERMINATED"`, v)
	}
}

// Values returns all possible values for CdfState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CdfState) Values() []CdfState {
	return []CdfState{
		CdfStateCdfStateSkipped,
		CdfStateCdfStateSnapshotting,
		CdfStateCdfStateStreaming,
		CdfStateCdfStateTerminated,
	}
}

// Type always returns CdfState to satisfy [pflag.Value] interface
func (f *CdfState) Type() string {
	return "CdfState"
}

// The read-only replication status of a single Postgres table replicated under
// a CdfConfig. One status exists per replicated table. It is created
// automatically and cannot be modified.
type CdfStatus struct {
	// The high-watermark Log Sequence Number (LSN) committed to Delta Lake.
	CommittedLsn string `json:"committed_lsn,omitempty"`
	// When replication for this table was first established.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The last time changes for this table were written to Delta Lake.
	LastSyncTime *time.Time `json:"last_sync_time,omitempty"`
	// Output only. The full resource name of the CdfStatus. Format:
	// projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}/cdf-statuses/{cdf_status}
	// The {cdf_status} segment is the Postgres table name.
	Name string `json:"name,omitempty"`
	// The Postgres table being replicated.
	PostgresTable string `json:"postgres_table,omitempty"`
	// The current replication state of this table.
	State CdfState `json:"state,omitempty"`
	// Human-readable detail for the current state (e.g. the skip/error reason).
	// Empty for healthy states.
	StatusDetail string `json:"status_detail,omitempty"`
	// The Unity Catalog table receiving replicated data.
	UcTable string `json:"uc_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CdfStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CdfStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBranchRequest struct {
	// The Branch to create.
	Branch Branch `json:"branch"`
	// The ID to use for the Branch. This becomes the final component of the
	// branch's resource name. The ID is required and must be 1-63 characters
	// long, start with a lowercase letter, and contain only lowercase letters,
	// numbers, and hyphens. For example, `development` becomes
	// `projects/my-app/branches/development`.
	BranchId string `json:"-" url:"branch_id"`
	// The Project where this Branch will be created. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`
	// If true, update the branch if it already exists instead of returning an
	// error.
	ReplaceExisting bool `json:"-" url:"replace_existing,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBranchRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCatalogRequest struct {
	Catalog Catalog `json:"catalog"`
	// The ID of the catalog to create, which is also its name in Unity Catalog.
	// It becomes the last segment of the catalog's resource name,
	// catalogs/{catalog_id}.
	CatalogId string `json:"-" url:"catalog_id"`
}

func (s *CreateCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateCdfConfigRequest struct {
	// The CdfConfig to create. The catalog, schema, and postgres_schema fields
	// are required; service_principal is optional. All other fields are output
	// only and ignored on input.
	CdfConfig CdfConfig `json:"cdf_config"`
	// The user-specified id for the CdfConfig, forming the final segment of its
	// resource name. Must match the pattern `[a-z][a-z0-9_]{0,62}`. Defaults to
	// the Postgres schema name when omitted.
	CdfConfigId string `json:"-" url:"cdf_config_id,omitempty"`
	// The parent database under which to create the CdfConfig. Format:
	// projects/{project}/branches/{branch}/databases/{database}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCdfConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCdfConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDataApiRequest struct {
	// The Data API configuration to create.
	DataApi DataApi `json:"data_api"`
	// Parent database:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Parent string `json:"-" url:"-"`
}

func (s *CreateDataApiRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateDatabaseRequest struct {
	// The desired specification of a Database.
	Database Database `json:"database"`
	// The ID used to identify and manage the Database resource through the API.
	// The Postgres database name is specified in `spec.postgres_database`.
	//
	// This value should be 3-63 characters, start with a lowercase letter, end
	// with a lowercase letter or digit, and contain only lowercase letters,
	// digits, and hyphens.
	//
	// If database_id is not specified in the request, it is generated
	// automatically.
	DatabaseId string `json:"-" url:"database_id,omitempty"`
	// The Branch where this Database will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`
	// If true, update the database if it already exists instead of returning an
	// error. This requires `database_id` to be set. Without `database_id`, the
	// request follows the normal create path and fails if the database name
	// already exists.
	ReplaceExisting bool `json:"-" url:"replace_existing,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDatabaseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDatabaseRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpointRequest struct {
	// The endpoint to create. Only `spec` is read from this message; the system
	// sets the other fields.
	Endpoint Endpoint `json:"endpoint"`
	// The ID to use for the endpoint. It becomes the final segment of the
	// endpoint's resource name,
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`. The
	// ID must be 3 to 63 characters long, start with a lowercase letter, end
	// with a lowercase letter or digit, and contain only lowercase letters,
	// digits, and hyphens.
	EndpointId string `json:"-" url:"endpoint_id"`
	// The Branch where this Endpoint will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`
	// If true, update the endpoint if it already exists instead of returning an
	// error. The request's `spec.endpoint_type` must match the existing
	// endpoint's type, and the other `spec` fields are applied as an update.
	// Use this to configure the read-write endpoint that was created together
	// with the branch.
	ReplaceExisting bool `json:"-" url:"replace_existing,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateMigrationRequest struct {
	// Required. The migration to create.
	Migration Migration `json:"migration"`
	// Required. The ID to use for the migration. It becomes the final component
	// of the resource name (`migrations/{migration_id}`). It must be 1-63
	// characters long, start with a lowercase letter, and contain only
	// lowercase letters, numbers, and hyphens.
	MigrationId string `json:"-" url:"migration_id"`
}

func (s *CreateMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateProjectRequest struct {
	// The Project to create.
	Project Project `json:"project"`
	// The ID to use for the project. It becomes the final segment of the
	// project's resource name, `projects/{project_id}`. The ID must be 3 to 63
	// characters long, start with a lowercase letter, end with a lowercase
	// letter or digit, and contain only lowercase letters, digits, and hyphens.
	// The ID of a soft-deleted project stays reserved until that project is
	// permanently deleted.
	ProjectId string `json:"-" url:"project_id"`
}

func (s *CreateProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateRoleRequest struct {
	// The Branch where this Role is created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`
	// If true, update the role if it already exists instead of returning an
	// error. This requires `role_id` to be set. Without `role_id`, the request
	// follows the normal create path and fails if the role name already exists.
	//
	// When the role already exists, only mutable fields are updated. Attributes
	// omitted from the request retain their current values. Set an attribute
	// explicitly to false to disable it. `identity_type`, `auth_method`, and
	// `postgres_role` retain their existing values. Omit these fields or supply
	// their current values. A differing supplied value returns
	// INVALID_PARAMETER_VALUE and rejects the entire request, including changes
	// to mutable fields. `membership_roles` is overwritten, not merged. Leaving
	// `membership_roles` empty clears all of the role's existing memberships,
	// including `DATABRICKS_SUPERUSER`. Always send the complete desired list
	// of memberships when using this field.
	ReplaceExisting bool `json:"-" url:"replace_existing,omitempty"`
	// The desired specification of a Role.
	Role Role `json:"role"`
	// The ID used to identify and manage the Role resource through the API. The
	// Postgres role name is specified in `spec.postgres_role`.
	//
	// This value should be 3-63 characters, start with a lowercase letter, end
	// with a lowercase letter or digit, and contain only lowercase letters,
	// digits, and hyphens.
	//
	// If role_id is not specified in the request, it is generated
	// automatically.
	RoleId string `json:"-" url:"role_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateSnapshotRequest struct {
	// The project in which to create the snapshot. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`
	// The snapshot to create. Only `spec` is read from this message; the system
	// sets the other fields.
	Snapshot Snapshot `json:"snapshot"`
	// Client-chosen ID for the snapshot. It becomes the final segment of the
	// snapshot resource name and cannot be changed after creation. The ID must
	// be 3 to 63 characters long, start with a lowercase letter, end with a
	// lowercase letter or digit, and contain only lowercase letters, digits,
	// and hyphens.
	SnapshotId string `json:"-" url:"snapshot_id"`
}

func (s *CreateSnapshotRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateSyncedTableRequest struct {
	// The synced table to create. Its configuration is read from spec, and its
	// name comes from synced_table_id.
	SyncedTable SyncedTable `json:"synced_table"`
	// The ID to use for the Synced Table. This becomes the final component of
	// the SyncedTable's resource name. ID is required and is the synced table
	// name, containing (catalog, schema, table) tuple. Elements of the tuple
	// are the UC entity names.
	//
	// Example: "{catalog}.{schema}.{table}"
	//
	// synced_table_id represents both of the following:
	//
	// 1. An online VIEW virtual table in the Unity Catalog accessible via the
	// Lakehouse Federation. 2. Postgres table named "{table}" in schema
	// "{schema}" in the connected Postgres database
	SyncedTableId string `json:"-" url:"synced_table_id"`
}

func (s *CreateSyncedTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Take a snapshot once per day, at the configured hour.
type DailySchedule struct {
	// The hour of the day, in UTC, at which to take the snapshot, in [0, 23].
	Hour int `json:"hour,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DailySchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DailySchedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// DataApi represents the Data API (PostgREST) configuration for a Database. At
// most one DataApi per database. Create enables Data API, Delete disables it.
type DataApi struct {
	// A timestamp indicating when the Data API was first enabled.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Resource name:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}/data-api
	Name string `json:"name,omitempty"`
	// The database containing this Data API configuration. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Parent string `json:"parent,omitempty"`
	// The desired Data API configuration.
	Spec *DataApiDataApiSpec `json:"spec,omitempty"`
	// The observed Data API state (read-only).
	Status *DataApiDataApiStatus `json:"status,omitempty"`
	// A timestamp indicating when the Data API configuration was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataApi) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataApi) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Desired PostgREST configuration (input).
type DataApiDataApiSpec struct {
	// Enable aggregate functions (count, sum, avg, etc.) in Data API responses.
	// Default: true.
	DbAggregatesEnabled bool `json:"db_aggregates_enabled,omitempty"`
	// Additional schemas to include in the PostgreSQL search path. Each entry
	// must be a valid PostgreSQL schema name.
	DbExtraSearchPath []string `json:"db_extra_search_path,omitempty"`
	// Maximum number of rows returned in a single Data API response. Must be a
	// positive integer.
	DbMaxRows int `json:"db_max_rows,omitempty"`
	// Database schemas exposed through the Data API. Each entry must be a valid
	// PostgreSQL schema name (1-63 chars, [a-zA-Z_][a-zA-Z0-9_$]*). Maximum 100
	// entries. Default: ["public"].
	DbSchemas []string `json:"db_schemas,omitempty"`
	// Maximum lifetime for cached JWT tokens. Zero duration disables caching.
	JwtCacheMaxLifetime *duration.Duration `json:"jwt_cache_max_lifetime,omitempty"`
	// JSON path to the role claim in JWT tokens (e.g., ".sub"). Default:
	// ".sub".
	JwtRoleClaimKey string `json:"jwt_role_claim_key,omitempty"`
	// OpenAPI documentation mode for the Data API endpoint.
	OpenapiMode OpenApiMode `json:"openapi_mode,omitempty"`
	// Allowed origins for CORS requests. Each entry should be a valid origin
	// URL, or use "*" to allow all origins.
	ServerCorsAllowedOrigins []string `json:"server_cors_allowed_origins,omitempty"`
	// Enable the Server-Timing header in Data API responses.
	ServerTimingEnabled bool `json:"server_timing_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataApiDataApiSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataApiDataApiSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Observed state (output-only).
type DataApiDataApiStatus struct {
	// Schemas available in the database (for reference when configuring
	// db_schemas).
	AvailableSchemas []string `json:"available_schemas,omitempty"`
	// Actual aggregate function setting read from the database.
	DbAggregatesEnabled bool `json:"db_aggregates_enabled,omitempty"`
	// Actual extra search path schemas read from the database.
	DbExtraSearchPath []string `json:"db_extra_search_path,omitempty"`
	// Actual max rows setting read from the database.
	DbMaxRows int `json:"db_max_rows,omitempty"`
	// Actual exposed schemas read from the database.
	DbSchemas []string `json:"db_schemas,omitempty"`
	// Actual JWT cache max lifetime read from the database.
	JwtCacheMaxLifetime *duration.Duration `json:"jwt_cache_max_lifetime,omitempty"`
	// Actual JWT role claim key read from the database.
	JwtRoleClaimKey string `json:"jwt_role_claim_key,omitempty"`
	// Actual OpenAPI mode read from the database.
	OpenapiMode OpenApiMode `json:"openapi_mode,omitempty"`
	// Actual CORS allowed origins read from the database.
	ServerCorsAllowedOrigins []string `json:"server_cors_allowed_origins,omitempty"`
	// Actual Server-Timing header setting read from the database.
	ServerTimingEnabled bool `json:"server_timing_enabled,omitempty"`
	// Data API endpoint URL.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataApiDataApiStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataApiDataApiStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DataApiOperationMetadata struct {
}

func (s *DataApiOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Database represents a Postgres database within a Branch.
type Database struct {
	// A timestamp indicating when the database was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The part of the name, chosen by the user when the resource was created.
	DatabaseId string `json:"database_id,omitempty"`
	// The resource name of the database. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Name string `json:"name,omitempty"`
	// The branch containing this database. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The desired state of the Database.
	//
	// Read the resource's current values from `status`, rather than `spec`.
	// `spec` tracks field ownership for infrastructure-as-code tools and may
	// omit unmanaged fields.
	Spec *DatabaseDatabaseSpec `json:"spec,omitempty"`
	// The observed state of the Database.
	Status *DatabaseDatabaseStatus `json:"status,omitempty"`
	// A timestamp indicating when the database was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Database) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Database) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseCredential struct {
	// Timestamp in UTC of when this credential expires.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// The OAuth token that can be used as a password when connecting to a
	// database.
	Token string `json:"token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseDatabaseSpec struct {
	// The name of the Postgres database.
	//
	// This expects a valid Postgres identifier as specified in the link below.
	// https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
	// Required when creating the Database.
	//
	// To rename, pass a valid postgres identifier when updating the Database.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// The name of the role that owns the database. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	//
	// The role must be in the same branch as the database. To change the owner,
	// pass valid existing Role name when updating the Database
	//
	// Required when creating the Database.
	Role string `json:"role"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseDatabaseSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseDatabaseSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseDatabaseStatus struct {
	// Part of the resource name.
	DatabaseId string `json:"database_id,omitempty"`
	// The name of the Postgres database.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// The name of the role that owns the database. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Role string `json:"role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseDatabaseStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseDatabaseStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A mapping between one source and target database.
type DatabaseMapping struct {
	// Output-only. The resolved target PostgreSQL database name. Equal to
	// `target_database` when provided; otherwise equal to `source_database`.
	EffectiveTargetDatabase string `json:"effective_target_database,omitempty"`
	// Required. The exact PostgreSQL database name on the source server.
	SourceDatabase string `json:"source_database"`
	// The PostgreSQL database name to create on the target branch. When
	// omitted, the source database name is preserved.
	TargetDatabase string `json:"target_database,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseMapping) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseMapping) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseOperationMetadata struct {
}

func (s *DatabaseOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Physical database size, including indexes and bloat; not a dump archive size
// estimate.
type DatabaseSizeDetails struct {
	// Size reported by pg_database_size, in bytes. Zero is a measured empty
	// size.
	SizeBytes int64 `json:"size_bytes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseSizeDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseSizeDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details []json.RawMessage `json:"details,omitempty"`

	ErrorCode ErrorCode `json:"error_code,omitempty"`

	Message string `json:"message,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksServiceExceptionWithDetailsProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksServiceExceptionWithDetailsProto) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The day of the week on which a weekly snapshot is taken.
type DayOfWeek string

const DayOfWeekFriday DayOfWeek = `FRIDAY`

const DayOfWeekMonday DayOfWeek = `MONDAY`

const DayOfWeekSaturday DayOfWeek = `SATURDAY`

const DayOfWeekSunday DayOfWeek = `SUNDAY`

const DayOfWeekThursday DayOfWeek = `THURSDAY`

const DayOfWeekTuesday DayOfWeek = `TUESDAY`

const DayOfWeekWednesday DayOfWeek = `WEDNESDAY`

// String representation for [fmt.Print]
func (f *DayOfWeek) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DayOfWeek) Set(v string) error {
	switch v {
	case `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = DayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Values returns all possible values for DayOfWeek.
//
// There is no guarantee on the order of the values in the slice.
func (f *DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// Type always returns DayOfWeek to satisfy [pflag.Value] interface
func (f *DayOfWeek) Type() string {
	return "DayOfWeek"
}

type DeleteBranchRequest struct {
	// The full resource path of the branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
	// If true, permanently delete the branch; if false, soft delete.
	Purge bool `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteBranchRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteCatalogRequest struct {
	// The full resource path of the catalog to delete.
	//
	// Format: "catalogs/{catalog_id}".
	Name string `json:"-" url:"-"`
}

func (s *DeleteCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteCdfConfigRequest struct {
	// When true, also drops the replicated Delta tables in Unity Catalog. When
	// false (the default), the replicated tables are preserved at their last
	// synced state.
	Force bool `json:"-" url:"force,omitempty"`
	// The resource name of the CdfConfig to delete. Format:
	// projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteCdfConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCdfConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDataApiRequest struct {
	// Resource name:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}/data-api
	Name string `json:"-" url:"-"`
}

func (s *DeleteDataApiRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDatabaseRequest struct {
	// The resource name of the postgres database. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteDatabaseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteEndpointRequest struct {
	// The full resource path of the endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteMigrationRequest struct {
	// Required. The migration resource name, in the format
	// `migrations/{migration}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteProjectRequest struct {
	// The full resource path of the project to delete. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
	// If true, permanently deletes the project and all of its data immediately.
	// If false or unset, the project is soft-deleted: it stops accepting
	// connections and is retained for 7 days, during which it can be restored
	// with Undelete, and is then permanently deleted. Deleting an already
	// soft-deleted project with `purge` set to true removes it immediately. See
	// https://docs.databricks.com/aws/en/oltp/projects/manage-projects for
	// details.
	Purge bool `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteProjectRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteRoleRequest struct {
	// The full resource path of the role to delete. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
	// Reassign objects. If this is set, all objects owned by the role are
	// reassigned to the role specified in this parameter. The role must already
	// exist. Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}
	//
	// NOTE: setting this requires spinning up a compute to succeed, since it
	// involves running SQL queries.
	ReassignOwnedTo string `json:"-" url:"reassign_owned_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSnapshotRequest struct {
	// The resource name of the snapshot to delete. Format:
	// projects/{project_id}/snapshots/{snapshot_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteSnapshotRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSyncedTableRequest struct {
	// The Full resource name of the synced table, of the format
	// "synced_tables/{catalog}.{schema}.{table}", where (catalog, schema,
	// table) are the UC entity names.
	Name string `json:"-" url:"-"`
}

func (s *DeleteSyncedTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeltaTableSyncInfo struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	DeltaCommitTime *time.Time `json:"delta_commit_time,omitempty"`
	// The Delta Lake commit version that was last successfully synced.
	DeltaCommitVersion int64 `json:"delta_commit_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaTableSyncInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaTableSyncInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A compute endpoint of a Lakebase Postgres branch. An endpoint provides the
// compute that runs Postgres for the branch's data and the hostnames that
// clients connect to. Each branch has one read-write endpoint, created together
// with the branch, and can have additional read-only endpoints (read replicas).
type Endpoint struct {
	// A timestamp indicating when the compute endpoint was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The endpoint ID chosen when the endpoint was created. It is the final
	// segment of `name`.
	EndpointId string `json:"endpoint_id,omitempty"`
	// Output only. The full resource path of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"name,omitempty"`
	// The branch containing this endpoint (API resource hierarchy). Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The desired configuration of the endpoint. Every field under `spec`
	// except `endpoint_type` can be changed after creation. See `status` for
	// the configuration that is currently in effect.
	//
	// Read the resource's current values from `status`, rather than `spec`.
	// `spec` tracks field ownership for infrastructure-as-code tools and may
	// omit unmanaged fields.
	Spec *EndpointSpec `json:"spec,omitempty"`
	// The configuration and state of the endpoint that are currently in effect,
	// as reported by the system, including the hostnames to connect to.
	Status *EndpointStatus `json:"status,omitempty"`
	// A system-generated ID that identifies the endpoint. Unlike `name`, it is
	// not derived from the user-chosen endpoint ID. The endpoint's connection
	// hostnames are based on this ID.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Endpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Endpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointGroupSpec struct {
	// Whether to allow read-only connections to read-write endpoints. Only
	// relevant for read-write endpoints where `max` is greater than 1. When
	// enabled, the endpoint also reports a read-only host in `status.hosts`.
	EnableReadableSecondaries bool `json:"enable_readable_secondaries,omitempty"`
	// The maximum number of computes in the endpoint group. Currently, this
	// must be equal to min. Set to 1 for single compute endpoints, to disable
	// HA. To manually suspend all computes in an endpoint group, set disabled
	// to true on the EndpointSpec.
	Max int `json:"max"`
	// The minimum number of computes in the endpoint group. Currently, this
	// must be equal to max. This must be greater than or equal to 1.
	Min int `json:"min"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointGroupSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointGroupSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointGroupStatus struct {
	// Whether read-only connections to read-write endpoints are allowed. Only
	// relevant if the endpoint has more than one compute (`max` greater than
	// 1).
	EnableReadableSecondaries bool `json:"enable_readable_secondaries,omitempty"`
	// The maximum number of computes in the endpoint group. Currently, this
	// must be equal to min. Set to 1 for single compute endpoints, to disable
	// HA. To manually suspend all computes in an endpoint group, set disabled
	// to true on the EndpointSpec.
	Max int `json:"max"`
	// The minimum number of computes in the endpoint group. Currently, this
	// must be equal to max. This must be greater than or equal to 1.
	Min int `json:"min"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointGroupStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointGroupStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Encapsulates various hostnames (r/w or r/o, pooled or not) for an endpoint.
type EndpointHosts struct {
	// The hostname to connect to this endpoint. For read-write endpoints, this
	// is a read-write hostname which connects to the primary compute. For
	// read-only endpoints, this is a read-only hostname which allows read-only
	// operations.
	Host string `json:"host,omitempty"`
	// An optionally defined read-only host for the endpoint, without pooling.
	// For read-only endpoints, this attribute is always defined and is
	// equivalent to host. For read-write endpoints, this attribute is defined
	// if the enclosing endpoint is a group with greater than 1 computes
	// configured, and has readable secondaries enabled.
	ReadOnlyHost string `json:"read_only_host,omitempty"`
	// The read-only hostname of the compute endpoint, with pooling. This
	// attribute is always defined for read-only endpoints, and may be defined
	// for read-write endpoints if configured with read replicas and allow
	// read-only connections.
	ReadOnlyPooledHost string `json:"read_only_pooled_host,omitempty"`
	// The read-write hostname of the compute endpoint, with pooling. This
	// attribute is only defined for read-write endpoints.
	ReadWritePooledHost string `json:"read_write_pooled_host,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointHosts) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointHosts) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointOperationMetadata struct {
}

func (s *EndpointOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A collection of settings for a compute endpoint.
type EndpointSettings struct {
	// Postgres server parameters to set on the compute endpoint, keyed by
	// parameter name. Each value is given as a string.
	PgSettings map[string]string `json:"pg_settings,omitempty"`
}

func (s *EndpointSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EndpointSpec struct {
	// The maximum number of Compute Units. The maximum value is 64. The
	// difference between the minimum and maximum Compute Units (max - min) must
	// not exceed 16. Set it equal to `autoscaling_limit_min_cu` for a
	// fixed-size compute. When omitted at creation, the project's default
	// endpoint settings apply.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units. Minimum value is 0.5, and it must
	// not exceed `autoscaling_limit_max_cu`. When omitted at creation, the
	// project's default endpoint settings apply.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action. To accept
	// connections again, set this field back to false.
	Disabled bool `json:"disabled,omitempty"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType EndpointType `json:"endpoint_type"`
	// Settings for optional HA configuration of the endpoint. If unspecified,
	// the endpoint defaults to non HA settings, with a single compute backing
	// the endpoint (and no readable secondaries for Read/Write endpoints).
	// Automatic suspension is not available for an endpoint with more than one
	// compute, so set `no_suspension` to true when configuring one.
	Group *EndpointGroupSpec `json:"group,omitempty"`
	// When set to true, explicitly disables automatic suspension (never
	// suspend). The value must be true when provided. Mutually exclusive with
	// `suspend_timeout_duration`.
	NoSuspension bool `json:"no_suspension,omitempty"`
	// Postgres server parameters for the compute endpoint. On update, an empty
	// `pg_settings` map leaves the current parameters unchanged.
	Settings *EndpointSettings `json:"settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. If specified, the value must be between 60s and 604800s (1
	// minute to 1 week). A suspended endpoint keeps its hostnames and
	// configuration, and the next connection attempt restarts it. When neither
	// this field nor `no_suspension` is set at creation, the project's default
	// endpoint settings apply. Mutually exclusive with `no_suspension`.
	SuspendTimeoutDuration *duration.Duration `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointStatus struct {
	// The maximum number of Compute Units. The maximum value is 64. The
	// difference between the minimum and maximum Compute Units (max - min) must
	// not exceed 16.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// The state the compute endpoint is in now.
	CurrentState EndpointStatusState `json:"current_state,omitempty"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled bool `json:"disabled,omitempty"`
	// The endpoint ID, which is the final segment of the endpoint's resource
	// name.
	EndpointId string `json:"endpoint_id,omitempty"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType EndpointType `json:"endpoint_type,omitempty"`
	// Details on the HA configuration of the endpoint.
	Group *EndpointGroupStatus `json:"group,omitempty"`
	// Contains host information for connecting to the endpoint.
	Hosts *EndpointHosts `json:"hosts,omitempty"`
	// A timestamp indicating when the compute endpoint was last active.
	LastActiveTime *time.Time `json:"last_active_time,omitempty"`
	// The state the compute endpoint is transitioning to, if a transition is in
	// progress.
	PendingState EndpointStatusState `json:"pending_state,omitempty"`
	// The Postgres server parameters currently set on the compute endpoint.
	Settings *EndpointSettings `json:"settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. Absent when automatic suspension is disabled.
	SuspendTimeoutDuration *duration.Duration `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the compute endpoint.
type EndpointStatusState string

const EndpointStatusStateActive EndpointStatusState = `ACTIVE`

const EndpointStatusStateDegraded EndpointStatusState = `DEGRADED`

const EndpointStatusStateIdle EndpointStatusState = `IDLE`

const EndpointStatusStateInit EndpointStatusState = `INIT`

// String representation for [fmt.Print]
func (f *EndpointStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStatusState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DEGRADED`, `IDLE`, `INIT`:
		*f = EndpointStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DEGRADED", "IDLE", "INIT"`, v)
	}
}

// Values returns all possible values for EndpointStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStatusState) Values() []EndpointStatusState {
	return []EndpointStatusState{
		EndpointStatusStateActive,
		EndpointStatusStateDegraded,
		EndpointStatusStateIdle,
		EndpointStatusStateInit,
	}
}

// Type always returns EndpointStatusState to satisfy [pflag.Value] interface
func (f *EndpointStatusState) Type() string {
	return "EndpointStatusState"
}

// The compute endpoint type. Either `read_write` or `read_only`.
type EndpointType string

const EndpointTypeEndpointTypeReadOnly EndpointType = `ENDPOINT_TYPE_READ_ONLY`

const EndpointTypeEndpointTypeReadWrite EndpointType = `ENDPOINT_TYPE_READ_WRITE`

// String representation for [fmt.Print]
func (f *EndpointType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointType) Set(v string) error {
	switch v {
	case `ENDPOINT_TYPE_READ_ONLY`, `ENDPOINT_TYPE_READ_WRITE`:
		*f = EndpointType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ENDPOINT_TYPE_READ_ONLY", "ENDPOINT_TYPE_READ_WRITE"`, v)
	}
}

// Values returns all possible values for EndpointType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointType) Values() []EndpointType {
	return []EndpointType{
		EndpointTypeEndpointTypeReadOnly,
		EndpointTypeEndpointTypeReadWrite,
	}
}

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
}

// Error codes returned by Databricks APIs to indicate specific failure
// conditions.
type ErrorCode string

const ErrorCodeAborted ErrorCode = `ABORTED`

const ErrorCodeAlreadyExists ErrorCode = `ALREADY_EXISTS`

const ErrorCodeBadRequest ErrorCode = `BAD_REQUEST`

const ErrorCodeCancelled ErrorCode = `CANCELLED`

const ErrorCodeCatalogAlreadyExists ErrorCode = `CATALOG_ALREADY_EXISTS`

const ErrorCodeCatalogDoesNotExist ErrorCode = `CATALOG_DOES_NOT_EXIST`

const ErrorCodeCatalogNotEmpty ErrorCode = `CATALOG_NOT_EMPTY`

const ErrorCodeCouldNotAcquireLock ErrorCode = `COULD_NOT_ACQUIRE_LOCK`

const ErrorCodeCustomerUnauthorized ErrorCode = `CUSTOMER_UNAUTHORIZED`

const ErrorCodeDacAlreadyExists ErrorCode = `DAC_ALREADY_EXISTS`

const ErrorCodeDacDoesNotExist ErrorCode = `DAC_DOES_NOT_EXIST`

const ErrorCodeDataLoss ErrorCode = `DATA_LOSS`

const ErrorCodeDeadlineExceeded ErrorCode = `DEADLINE_EXCEEDED`

const ErrorCodeDeploymentTimeout ErrorCode = `DEPLOYMENT_TIMEOUT`

const ErrorCodeDirectoryNotEmpty ErrorCode = `DIRECTORY_NOT_EMPTY`

const ErrorCodeDirectoryProtected ErrorCode = `DIRECTORY_PROTECTED`

const ErrorCodeDryRunFailed ErrorCode = `DRY_RUN_FAILED`

const ErrorCodeEndpointNotFound ErrorCode = `ENDPOINT_NOT_FOUND`

const ErrorCodeExternalLocationAlreadyExists ErrorCode = `EXTERNAL_LOCATION_ALREADY_EXISTS`

const ErrorCodeExternalLocationDoesNotExist ErrorCode = `EXTERNAL_LOCATION_DOES_NOT_EXIST`

const ErrorCodeFeatureDisabled ErrorCode = `FEATURE_DISABLED`

const ErrorCodeGitConflict ErrorCode = `GIT_CONFLICT`

const ErrorCodeGitRemoteError ErrorCode = `GIT_REMOTE_ERROR`

const ErrorCodeGitSensitiveTokenDetected ErrorCode = `GIT_SENSITIVE_TOKEN_DETECTED`

const ErrorCodeGitUnknownRef ErrorCode = `GIT_UNKNOWN_REF`

const ErrorCodeGitUrlNotOnAllowList ErrorCode = `GIT_URL_NOT_ON_ALLOW_LIST`

const ErrorCodeInsecurePartnerResponse ErrorCode = `INSECURE_PARTNER_RESPONSE`

const ErrorCodeInternalError ErrorCode = `INTERNAL_ERROR`

const ErrorCodeInvalidParameterValue ErrorCode = `INVALID_PARAMETER_VALUE`

const ErrorCodeInvalidState ErrorCode = `INVALID_STATE`

const ErrorCodeInvalidStateTransition ErrorCode = `INVALID_STATE_TRANSITION`

const ErrorCodeIoError ErrorCode = `IO_ERROR`

const ErrorCodeIpynbFileInRepo ErrorCode = `IPYNB_FILE_IN_REPO`

const ErrorCodeMalformedPartnerResponse ErrorCode = `MALFORMED_PARTNER_RESPONSE`

const ErrorCodeMalformedRequest ErrorCode = `MALFORMED_REQUEST`

const ErrorCodeManagedResourceGroupDoesNotExist ErrorCode = `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`

const ErrorCodeMaxBlockSizeExceeded ErrorCode = `MAX_BLOCK_SIZE_EXCEEDED`

const ErrorCodeMaxChildNodeSizeExceeded ErrorCode = `MAX_CHILD_NODE_SIZE_EXCEEDED`

const ErrorCodeMaxListSizeExceeded ErrorCode = `MAX_LIST_SIZE_EXCEEDED`

const ErrorCodeMaxNotebookSizeExceeded ErrorCode = `MAX_NOTEBOOK_SIZE_EXCEEDED`

const ErrorCodeMaxReadSizeExceeded ErrorCode = `MAX_READ_SIZE_EXCEEDED`

const ErrorCodeMetastoreAlreadyExists ErrorCode = `METASTORE_ALREADY_EXISTS`

const ErrorCodeMetastoreDoesNotExist ErrorCode = `METASTORE_DOES_NOT_EXIST`

const ErrorCodeMetastoreNotEmpty ErrorCode = `METASTORE_NOT_EMPTY`

const ErrorCodeNotFound ErrorCode = `NOT_FOUND`

const ErrorCodeNotImplemented ErrorCode = `NOT_IMPLEMENTED`

const ErrorCodePartialDelete ErrorCode = `PARTIAL_DELETE`

const ErrorCodePermissionDenied ErrorCode = `PERMISSION_DENIED`

const ErrorCodePermissionNotPropagated ErrorCode = `PERMISSION_NOT_PROPAGATED`

const ErrorCodePrincipalDoesNotExist ErrorCode = `PRINCIPAL_DOES_NOT_EXIST`

const ErrorCodeProjectsOperationTimeout ErrorCode = `PROJECTS_OPERATION_TIMEOUT`

const ErrorCodeProviderAlreadyExists ErrorCode = `PROVIDER_ALREADY_EXISTS`

const ErrorCodeProviderDoesNotExist ErrorCode = `PROVIDER_DOES_NOT_EXIST`

const ErrorCodeProviderShareNotAccessible ErrorCode = `PROVIDER_SHARE_NOT_ACCESSIBLE`

const ErrorCodeQuotaExceeded ErrorCode = `QUOTA_EXCEEDED`

const ErrorCodeRecipientAlreadyExists ErrorCode = `RECIPIENT_ALREADY_EXISTS`

const ErrorCodeRecipientDoesNotExist ErrorCode = `RECIPIENT_DOES_NOT_EXIST`

const ErrorCodeRequestLimitExceeded ErrorCode = `REQUEST_LIMIT_EXCEEDED`

const ErrorCodeResourceAlreadyExists ErrorCode = `RESOURCE_ALREADY_EXISTS`

const ErrorCodeResourceConflict ErrorCode = `RESOURCE_CONFLICT`

const ErrorCodeResourceDoesNotExist ErrorCode = `RESOURCE_DOES_NOT_EXIST`

const ErrorCodeResourceExhausted ErrorCode = `RESOURCE_EXHAUSTED`

const ErrorCodeResourceLimitExceeded ErrorCode = `RESOURCE_LIMIT_EXCEEDED`

const ErrorCodeSchemaAlreadyExists ErrorCode = `SCHEMA_ALREADY_EXISTS`

const ErrorCodeSchemaDoesNotExist ErrorCode = `SCHEMA_DOES_NOT_EXIST`

const ErrorCodeSchemaNotEmpty ErrorCode = `SCHEMA_NOT_EMPTY`

const ErrorCodeSearchQueryTooLong ErrorCode = `SEARCH_QUERY_TOO_LONG`

const ErrorCodeSearchQueryTooShort ErrorCode = `SEARCH_QUERY_TOO_SHORT`

const ErrorCodeServiceUnderMaintenance ErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ErrorCodeShareAlreadyExists ErrorCode = `SHARE_ALREADY_EXISTS`

const ErrorCodeShareDoesNotExist ErrorCode = `SHARE_DOES_NOT_EXIST`

const ErrorCodeStorageCredentialAlreadyExists ErrorCode = `STORAGE_CREDENTIAL_ALREADY_EXISTS`

const ErrorCodeStorageCredentialDoesNotExist ErrorCode = `STORAGE_CREDENTIAL_DOES_NOT_EXIST`

const ErrorCodeTableAlreadyExists ErrorCode = `TABLE_ALREADY_EXISTS`

const ErrorCodeTableDoesNotExist ErrorCode = `TABLE_DOES_NOT_EXIST`

const ErrorCodeTemporarilyUnavailable ErrorCode = `TEMPORARILY_UNAVAILABLE`

const ErrorCodeUnauthenticated ErrorCode = `UNAUTHENTICATED`

const ErrorCodeUnavailable ErrorCode = `UNAVAILABLE`

const ErrorCodeUnknown ErrorCode = `UNKNOWN`

const ErrorCodeUnparseableHttpError ErrorCode = `UNPARSEABLE_HTTP_ERROR`

const ErrorCodeWorkspaceTemporarilyUnavailable ErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `CATALOG_ALREADY_EXISTS`, `CATALOG_DOES_NOT_EXIST`, `CATALOG_NOT_EMPTY`, `COULD_NOT_ACQUIRE_LOCK`, `CUSTOMER_UNAUTHORIZED`, `DAC_ALREADY_EXISTS`, `DAC_DOES_NOT_EXIST`, `DATA_LOSS`, `DEADLINE_EXCEEDED`, `DEPLOYMENT_TIMEOUT`, `DIRECTORY_NOT_EMPTY`, `DIRECTORY_PROTECTED`, `DRY_RUN_FAILED`, `ENDPOINT_NOT_FOUND`, `EXTERNAL_LOCATION_ALREADY_EXISTS`, `EXTERNAL_LOCATION_DOES_NOT_EXIST`, `FEATURE_DISABLED`, `GIT_CONFLICT`, `GIT_REMOTE_ERROR`, `GIT_SENSITIVE_TOKEN_DETECTED`, `GIT_UNKNOWN_REF`, `GIT_URL_NOT_ON_ALLOW_LIST`, `INSECURE_PARTNER_RESPONSE`, `INTERNAL_ERROR`, `INVALID_PARAMETER_VALUE`, `INVALID_STATE`, `INVALID_STATE_TRANSITION`, `IO_ERROR`, `IPYNB_FILE_IN_REPO`, `MALFORMED_PARTNER_RESPONSE`, `MALFORMED_REQUEST`, `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`, `MAX_BLOCK_SIZE_EXCEEDED`, `MAX_CHILD_NODE_SIZE_EXCEEDED`, `MAX_LIST_SIZE_EXCEEDED`, `MAX_NOTEBOOK_SIZE_EXCEEDED`, `MAX_READ_SIZE_EXCEEDED`, `METASTORE_ALREADY_EXISTS`, `METASTORE_DOES_NOT_EXIST`, `METASTORE_NOT_EMPTY`, `NOT_FOUND`, `NOT_IMPLEMENTED`, `PARTIAL_DELETE`, `PERMISSION_DENIED`, `PERMISSION_NOT_PROPAGATED`, `PRINCIPAL_DOES_NOT_EXIST`, `PROJECTS_OPERATION_TIMEOUT`, `PROVIDER_ALREADY_EXISTS`, `PROVIDER_DOES_NOT_EXIST`, `PROVIDER_SHARE_NOT_ACCESSIBLE`, `QUOTA_EXCEEDED`, `RECIPIENT_ALREADY_EXISTS`, `RECIPIENT_DOES_NOT_EXIST`, `REQUEST_LIMIT_EXCEEDED`, `RESOURCE_ALREADY_EXISTS`, `RESOURCE_CONFLICT`, `RESOURCE_DOES_NOT_EXIST`, `RESOURCE_EXHAUSTED`, `RESOURCE_LIMIT_EXCEEDED`, `SCHEMA_ALREADY_EXISTS`, `SCHEMA_DOES_NOT_EXIST`, `SCHEMA_NOT_EMPTY`, `SEARCH_QUERY_TOO_LONG`, `SEARCH_QUERY_TOO_SHORT`, `SERVICE_UNDER_MAINTENANCE`, `SHARE_ALREADY_EXISTS`, `SHARE_DOES_NOT_EXIST`, `STORAGE_CREDENTIAL_ALREADY_EXISTS`, `STORAGE_CREDENTIAL_DOES_NOT_EXIST`, `TABLE_ALREADY_EXISTS`, `TABLE_DOES_NOT_EXIST`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNAVAILABLE`, `UNKNOWN`, `UNPARSEABLE_HTTP_ERROR`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "CATALOG_ALREADY_EXISTS", "CATALOG_DOES_NOT_EXIST", "CATALOG_NOT_EMPTY", "COULD_NOT_ACQUIRE_LOCK", "CUSTOMER_UNAUTHORIZED", "DAC_ALREADY_EXISTS", "DAC_DOES_NOT_EXIST", "DATA_LOSS", "DEADLINE_EXCEEDED", "DEPLOYMENT_TIMEOUT", "DIRECTORY_NOT_EMPTY", "DIRECTORY_PROTECTED", "DRY_RUN_FAILED", "ENDPOINT_NOT_FOUND", "EXTERNAL_LOCATION_ALREADY_EXISTS", "EXTERNAL_LOCATION_DOES_NOT_EXIST", "FEATURE_DISABLED", "GIT_CONFLICT", "GIT_REMOTE_ERROR", "GIT_SENSITIVE_TOKEN_DETECTED", "GIT_UNKNOWN_REF", "GIT_URL_NOT_ON_ALLOW_LIST", "INSECURE_PARTNER_RESPONSE", "INTERNAL_ERROR", "INVALID_PARAMETER_VALUE", "INVALID_STATE", "INVALID_STATE_TRANSITION", "IO_ERROR", "IPYNB_FILE_IN_REPO", "MALFORMED_PARTNER_RESPONSE", "MALFORMED_REQUEST", "MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST", "MAX_BLOCK_SIZE_EXCEEDED", "MAX_CHILD_NODE_SIZE_EXCEEDED", "MAX_LIST_SIZE_EXCEEDED", "MAX_NOTEBOOK_SIZE_EXCEEDED", "MAX_READ_SIZE_EXCEEDED", "METASTORE_ALREADY_EXISTS", "METASTORE_DOES_NOT_EXIST", "METASTORE_NOT_EMPTY", "NOT_FOUND", "NOT_IMPLEMENTED", "PARTIAL_DELETE", "PERMISSION_DENIED", "PERMISSION_NOT_PROPAGATED", "PRINCIPAL_DOES_NOT_EXIST", "PROJECTS_OPERATION_TIMEOUT", "PROVIDER_ALREADY_EXISTS", "PROVIDER_DOES_NOT_EXIST", "PROVIDER_SHARE_NOT_ACCESSIBLE", "QUOTA_EXCEEDED", "RECIPIENT_ALREADY_EXISTS", "RECIPIENT_DOES_NOT_EXIST", "REQUEST_LIMIT_EXCEEDED", "RESOURCE_ALREADY_EXISTS", "RESOURCE_CONFLICT", "RESOURCE_DOES_NOT_EXIST", "RESOURCE_EXHAUSTED", "RESOURCE_LIMIT_EXCEEDED", "SCHEMA_ALREADY_EXISTS", "SCHEMA_DOES_NOT_EXIST", "SCHEMA_NOT_EMPTY", "SEARCH_QUERY_TOO_LONG", "SEARCH_QUERY_TOO_SHORT", "SERVICE_UNDER_MAINTENANCE", "SHARE_ALREADY_EXISTS", "SHARE_DOES_NOT_EXIST", "STORAGE_CREDENTIAL_ALREADY_EXISTS", "STORAGE_CREDENTIAL_DOES_NOT_EXIST", "TABLE_ALREADY_EXISTS", "TABLE_DOES_NOT_EXIST", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNAVAILABLE", "UNKNOWN", "UNPARSEABLE_HTTP_ERROR", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		ErrorCodeAborted,
		ErrorCodeAlreadyExists,
		ErrorCodeBadRequest,
		ErrorCodeCancelled,
		ErrorCodeCatalogAlreadyExists,
		ErrorCodeCatalogDoesNotExist,
		ErrorCodeCatalogNotEmpty,
		ErrorCodeCouldNotAcquireLock,
		ErrorCodeCustomerUnauthorized,
		ErrorCodeDacAlreadyExists,
		ErrorCodeDacDoesNotExist,
		ErrorCodeDataLoss,
		ErrorCodeDeadlineExceeded,
		ErrorCodeDeploymentTimeout,
		ErrorCodeDirectoryNotEmpty,
		ErrorCodeDirectoryProtected,
		ErrorCodeDryRunFailed,
		ErrorCodeEndpointNotFound,
		ErrorCodeExternalLocationAlreadyExists,
		ErrorCodeExternalLocationDoesNotExist,
		ErrorCodeFeatureDisabled,
		ErrorCodeGitConflict,
		ErrorCodeGitRemoteError,
		ErrorCodeGitSensitiveTokenDetected,
		ErrorCodeGitUnknownRef,
		ErrorCodeGitUrlNotOnAllowList,
		ErrorCodeInsecurePartnerResponse,
		ErrorCodeInternalError,
		ErrorCodeInvalidParameterValue,
		ErrorCodeInvalidState,
		ErrorCodeInvalidStateTransition,
		ErrorCodeIoError,
		ErrorCodeIpynbFileInRepo,
		ErrorCodeMalformedPartnerResponse,
		ErrorCodeMalformedRequest,
		ErrorCodeManagedResourceGroupDoesNotExist,
		ErrorCodeMaxBlockSizeExceeded,
		ErrorCodeMaxChildNodeSizeExceeded,
		ErrorCodeMaxListSizeExceeded,
		ErrorCodeMaxNotebookSizeExceeded,
		ErrorCodeMaxReadSizeExceeded,
		ErrorCodeMetastoreAlreadyExists,
		ErrorCodeMetastoreDoesNotExist,
		ErrorCodeMetastoreNotEmpty,
		ErrorCodeNotFound,
		ErrorCodeNotImplemented,
		ErrorCodePartialDelete,
		ErrorCodePermissionDenied,
		ErrorCodePermissionNotPropagated,
		ErrorCodePrincipalDoesNotExist,
		ErrorCodeProjectsOperationTimeout,
		ErrorCodeProviderAlreadyExists,
		ErrorCodeProviderDoesNotExist,
		ErrorCodeProviderShareNotAccessible,
		ErrorCodeQuotaExceeded,
		ErrorCodeRecipientAlreadyExists,
		ErrorCodeRecipientDoesNotExist,
		ErrorCodeRequestLimitExceeded,
		ErrorCodeResourceAlreadyExists,
		ErrorCodeResourceConflict,
		ErrorCodeResourceDoesNotExist,
		ErrorCodeResourceExhausted,
		ErrorCodeResourceLimitExceeded,
		ErrorCodeSchemaAlreadyExists,
		ErrorCodeSchemaDoesNotExist,
		ErrorCodeSchemaNotEmpty,
		ErrorCodeSearchQueryTooLong,
		ErrorCodeSearchQueryTooShort,
		ErrorCodeServiceUnderMaintenance,
		ErrorCodeShareAlreadyExists,
		ErrorCodeShareDoesNotExist,
		ErrorCodeStorageCredentialAlreadyExists,
		ErrorCodeStorageCredentialDoesNotExist,
		ErrorCodeTableAlreadyExists,
		ErrorCodeTableDoesNotExist,
		ErrorCodeTemporarilyUnavailable,
		ErrorCodeUnauthenticated,
		ErrorCodeUnavailable,
		ErrorCodeUnknown,
		ErrorCodeUnparseableHttpError,
		ErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ErrorCode to satisfy [pflag.Value] interface
func (f *ErrorCode) Type() string {
	return "ErrorCode"
}

type GenerateDatabaseCredentialRequest struct {
	// The returned token will be scoped to UC tables with the specified
	// permissions.
	Claims []RequestedClaims `json:"claims,omitempty"`
	// The endpoint resource name for which this credential will be generated.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint string `json:"endpoint"`
	// Timestamp in UTC of when this credential should expire. Must be at least
	// 300 seconds (5 minutes) and at most 1 hour from the current time.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// The display name of a `Databricks` workspace group. When set, the
	// returned credential is scoped to this group, so the caller connects
	// directly as the group's Postgres role. The caller must be a member of the
	// group. When omitted, the credential is scoped to the caller's own
	// identity.
	GroupName string `json:"group_name,omitempty"`
	// The requested time-to-live for the generated credential token. Must be at
	// least 300 seconds (5 minutes) and at most 3600 seconds (1 hour).
	Ttl *duration.Duration `json:"ttl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateDatabaseCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateDatabaseCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAssessmentRequest struct {
	// Required. The assessment resource name, in the format
	// `migrations/{migration}/assessments/{assessment}`.
	Name string `json:"-" url:"-"`
}

func (s *GetAssessmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBranchRequest struct {
	// The full resource path of the branch to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

func (s *GetBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetCatalogRequest struct {
	// The full resource path of the catalog to retrieve.
	//
	// Format: "catalogs/{catalog_id}".
	Name string `json:"-" url:"-"`
}

func (s *GetCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetCdfConfigRequest struct {
	// The resource name of the CdfConfig to retrieve. Format:
	// projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
	Name string `json:"-" url:"-"`
}

func (s *GetCdfConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetCdfStatusRequest struct {
	// The resource name of the CdfStatus to retrieve. Format:
	// projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}/cdf-statuses/{cdf_status}
	Name string `json:"-" url:"-"`
}

func (s *GetCdfStatusRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDataApiRequest struct {
	// Resource name:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}/data-api
	Name string `json:"-" url:"-"`
}

func (s *GetDataApiRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDatabaseRequest struct {
	// The name of the Database to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Name string `json:"-" url:"-"`
}

func (s *GetDatabaseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetEndpointRequest struct {
	// The full resource path of the endpoint to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
}

func (s *GetEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMigrationRequest struct {
	// Required. The migration resource name, in the format
	// `migrations/{migration}`.
	Name string `json:"-" url:"-"`
}

func (s *GetMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

func (s *GetOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetProjectRequest struct {
	// The full resource path of the project to retrieve. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
}

func (s *GetProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRoleRequest struct {
	// The full resource path of the role to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
}

func (s *GetRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSnapshotRequest struct {
	// The resource name of the snapshot to retrieve. Format:
	// projects/{project_id}/snapshots/{snapshot_id}
	Name string `json:"-" url:"-"`
}

func (s *GetSnapshotRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSnapshotScheduleRequest struct {
	// The resource name of the branch's snapshot schedule. Format:
	// projects/{project_id}/branches/{branch_id}/snapshot-schedule
	Name string `json:"-" url:"-"`
}

func (s *GetSnapshotScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSyncedTableRequest struct {
	// The Full resource name of the synced table. Format:
	// "synced_tables/{catalog}.{schema}.{table}", where (catalog, schema,
	// table) are the entity names in the Unity Catalog.
	Name string `json:"-" url:"-"`
}

func (s *GetSyncedTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configuration for the initial default branch created during project creation.
type InitialBranchSpec struct {
	// Whether the initial default branch should be protected from deletion.
	IsProtected bool `json:"is_protected,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InitialBranchSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InitialBranchSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for the initial Read/Write endpoint created during project
// creation.
type InitialEndpointSpec struct {
	// The maximum number of Compute Units for the initial endpoint. Minimum
	// value is 0.5.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units for the initial endpoint. Minimum
	// value is 0.5, and it must not exceed autoscaling_limit_max_cu.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// Settings for HA configuration of the endpoint.
	Group *EndpointGroupSpec `json:"group,omitempty"`
	// When set to true, explicitly disables automatic suspension (never
	// suspend). Should be set to true when provided. Mutually exclusive with
	// `suspend_timeout_duration`.
	NoSuspension bool `json:"no_suspension,omitempty"`
	// Duration of inactivity after which the initial endpoint is automatically
	// suspended. If specified, should be between 60s and 604800s (1 minute to 1
	// week). Mutually exclusive with `no_suspension`.
	SuspendTimeoutDuration *duration.Duration `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InitialEndpointSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InitialEndpointSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAssessmentsRequest struct {
	// The maximum number of assessments to return. The server default is 100;
	// values above 1000 are capped to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous call.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Required. The parent migration resource name, in the format
	// `migrations/{migration}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAssessmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAssessmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to list assessments.
type ListAssessmentsResponse struct {
	// Assessment runs ordered from newest to oldest.
	Assessments []Assessment `json:"assessments,omitempty"`
	// A token for retrieving the next page. Empty when there are no more
	// results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAssessmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAssessmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBranchesRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Project that owns this collection of branches. Format:
	// projects/{project_id}
	Parent string `json:"-" url:"-"`
	// Whether to include soft-deleted branches in the response. When true,
	// deleted branches are included alongside active branches. Purged branches
	// are never returned.
	ShowDeleted bool `json:"-" url:"show_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBranchesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBranchesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBranchesResponse struct {
	// List of branches in the project.
	Branches []Branch `json:"branches,omitempty"`
	// Token to request the next page of branches.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBranchesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBranchesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCdfConfigsRequest struct {
	// Maximum number of CdfConfigs to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token returned by a previous ListCdfConfigs call. Empty on the
	// first page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent database to list CdfConfigs for. Format:
	// projects/{project}/branches/{branch}/databases/{database}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCdfConfigsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCdfConfigsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to a ListCdfConfigs request, containing a page of CdfConfigs and a
// token for fetching the next page.
type ListCdfConfigsResponse struct {
	// The CdfConfigs under the parent database.
	CdfConfigs []CdfConfig `json:"cdf_configs,omitempty"`
	// Token to retrieve the next page of results; empty when there are no more.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCdfConfigsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCdfConfigsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCdfStatusesRequest struct {
	// Maximum number of CdfStatuses to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token returned by a previous ListCdfStatuses call. Empty on
	// the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent CdfConfig to list CdfStatuses for. Format:
	// projects/{project}/branches/{branch}/databases/{database}/cdf-configs/{cdf_config}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCdfStatusesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCdfStatusesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to a ListCdfStatuses request, containing a page of replicated table
// statuses and a token for fetching the next page.
type ListCdfStatusesResponse struct {
	// The replicated tables under the parent CdfConfig.
	CdfStatuses []CdfStatus `json:"cdf_statuses,omitempty"`
	// Token to retrieve the next page of results; empty when there are no more.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCdfStatusesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCdfStatusesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabasesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Databases. Requests first page
	// if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Branch that owns this collection of databases. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabasesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabasesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabasesResponse struct {
	// List of databases.
	Databases []Database `json:"databases,omitempty"`
	// Pagination token to request the next page of databases.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabasesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabasesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsRequest struct {
	// Upper bound for the number of endpoints returned. If set, the value must
	// be at least 10.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Branch that owns this collection of endpoints. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsResponse struct {
	// List of compute endpoints in the branch.
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// Token to request the next page of compute endpoints.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMigrationsRequest struct {
	// The maximum number of migrations to return. The server default is 100;
	// values above 1000 are capped to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous call.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Whether to include soft-deleted migrations in the response. Hard-deleted
	// and already-purged migrations are never returned.
	ShowDeleted bool `json:"-" url:"show_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMigrationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMigrationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to list migrations.
type ListMigrationsResponse struct {
	// Migrations visible to the caller.
	Migrations []Migration `json:"migrations,omitempty"`
	// A token for retrieving the next page. Empty when there are no more
	// results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMigrationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMigrationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProjectsRequest struct {
	// Upper bound for the number of projects returned. Cannot be negative.
	// Values above the maximum page size are reduced to that maximum.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Whether to include soft-deleted projects in the response. When true,
	// soft-deleted projects are included alongside active projects.
	// Hard-deleted and already-purged projects are never returned.
	ShowDeleted bool `json:"-" url:"show_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListProjectsResponse struct {
	// Token to request the next page of projects.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all projects in the workspace that the user has permission to
	// access.
	Projects []Project `json:"projects,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListProjectsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProjectsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRolesRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Branch that owns this collection of roles. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRolesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRolesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRolesResponse struct {
	// Token to request the next page of Postgres roles.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Postgres roles in the branch.
	Roles []Role `json:"roles,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRolesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRolesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSnapshotsRequest struct {
	// Maximum number of snapshots to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response; omit for the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The project that owns the snapshots. Format: projects/{project_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSnapshotsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSnapshotsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSnapshotsResponse struct {
	// Token to retrieve the next page; empty if there are no more pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The snapshots in the project.
	Snapshots []Snapshot `json:"snapshots,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSnapshotsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSnapshotsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Oldest active source transaction at assessment time.
type LongTransactionDetails struct {
	// Age of the oldest active transaction when assessed. Zero when none is
	// active.
	OldestTransactionAge *duration.Duration `json:"oldest_transaction_age,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LongTransactionDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LongTransactionDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A managed migration of PostgreSQL databases into a Lakebase branch.
type Migration struct {
	// Output-only. The time when the migration was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Required. Mappings from source databases to target databases.
	DatabaseMappings []DatabaseMapping `json:"database_mappings"`
	// Output-only. The time when the migration was soft-deleted. Empty if the
	// migration is active.
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// Output-only. The current error preventing migration progress, if any.
	Error *MigrationError `json:"error,omitempty"`
	// Required. Immutable. The intended migration mode, selected when the
	// migration is created.
	MigrationType MigrationType `json:"migration_type"`
	// The migration resource name, in the format `migrations/{migration}`.
	Name string `json:"name,omitempty"`
	// Output-only. The time when the soft-deleted migration is scheduled for
	// permanent deletion. Empty if the migration is active or permanent
	// deletion is not scheduled.
	PurgeTime *time.Time `json:"purge_time,omitempty"`
	// Required. The name of the Unity Catalog connection (a UC securable name,
	// such as `my_postgres_connection`) used to connect to the source
	// PostgreSQL server.
	SourceConnection string `json:"source_connection"`
	// Output-only. The authoritative migration lifecycle state.
	State MigrationState `json:"state,omitempty"`
	// The Lakebase branch receiving the migrated databases, in the format
	// `projects/{project_id}/branches/{branch_id}`. May be omitted during
	// assessment, but must be set before the migration starts, and may be
	// changed until then.
	TargetBranch string `json:"target_branch,omitempty"`
	// Output-only. The time when the migration was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Migration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Migration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An error that prevents migration progress.
type MigrationError struct {
	// A stable machine-readable error code.
	Code string `json:"code,omitempty"`
	// A concise explanation of the failure.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigrationError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrationError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for a long-running operation that applies an action to a migration.
type MigrationOperationMetadata struct {
}

func (s *MigrationOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The lifecycle state of a migration.
type MigrationState string

const MigrationStateMigrationStateAssessing MigrationState = `MIGRATION_STATE_ASSESSING`

const MigrationStateMigrationStateAssessmentNeedsAction MigrationState = `MIGRATION_STATE_ASSESSMENT_NEEDS_ACTION`

const MigrationStateMigrationStateCompleted MigrationState = `MIGRATION_STATE_COMPLETED`

const MigrationStateMigrationStateDraft MigrationState = `MIGRATION_STATE_DRAFT`

const MigrationStateMigrationStateFailed MigrationState = `MIGRATION_STATE_FAILED`

const MigrationStateMigrationStateInitialDataLoad MigrationState = `MIGRATION_STATE_INITIAL_DATA_LOAD`

const MigrationStateMigrationStateReadyToStart MigrationState = `MIGRATION_STATE_READY_TO_START`

const MigrationStateMigrationStateSchemaLoad MigrationState = `MIGRATION_STATE_SCHEMA_LOAD`

// String representation for [fmt.Print]
func (f *MigrationState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MigrationState) Set(v string) error {
	switch v {
	case `MIGRATION_STATE_ASSESSING`, `MIGRATION_STATE_ASSESSMENT_NEEDS_ACTION`, `MIGRATION_STATE_COMPLETED`, `MIGRATION_STATE_DRAFT`, `MIGRATION_STATE_FAILED`, `MIGRATION_STATE_INITIAL_DATA_LOAD`, `MIGRATION_STATE_READY_TO_START`, `MIGRATION_STATE_SCHEMA_LOAD`:
		*f = MigrationState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MIGRATION_STATE_ASSESSING", "MIGRATION_STATE_ASSESSMENT_NEEDS_ACTION", "MIGRATION_STATE_COMPLETED", "MIGRATION_STATE_DRAFT", "MIGRATION_STATE_FAILED", "MIGRATION_STATE_INITIAL_DATA_LOAD", "MIGRATION_STATE_READY_TO_START", "MIGRATION_STATE_SCHEMA_LOAD"`, v)
	}
}

// Values returns all possible values for MigrationState.
//
// There is no guarantee on the order of the values in the slice.
func (f *MigrationState) Values() []MigrationState {
	return []MigrationState{
		MigrationStateMigrationStateAssessing,
		MigrationStateMigrationStateAssessmentNeedsAction,
		MigrationStateMigrationStateCompleted,
		MigrationStateMigrationStateDraft,
		MigrationStateMigrationStateFailed,
		MigrationStateMigrationStateInitialDataLoad,
		MigrationStateMigrationStateReadyToStart,
		MigrationStateMigrationStateSchemaLoad,
	}
}

// Type always returns MigrationState to satisfy [pflag.Value] interface
func (f *MigrationState) Type() string {
	return "MigrationState"
}

// The intended data transfer mode of a migration.
type MigrationType string

const MigrationTypeMigrationTypeOneTimeCopy MigrationType = `MIGRATION_TYPE_ONE_TIME_COPY`

// String representation for [fmt.Print]
func (f *MigrationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MigrationType) Set(v string) error {
	switch v {
	case `MIGRATION_TYPE_ONE_TIME_COPY`:
		*f = MigrationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MIGRATION_TYPE_ONE_TIME_COPY"`, v)
	}
}

// Values returns all possible values for MigrationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MigrationType) Values() []MigrationType {
	return []MigrationType{
		MigrationTypeMigrationTypeOneTimeCopy,
	}
}

// Type always returns MigrationType to satisfy [pflag.Value] interface
func (f *MigrationType) Type() string {
	return "MigrationType"
}

// Take a snapshot once per month, on the configured day at the configured hour.
type MonthlySchedule struct {
	// The day of the month on which to take the snapshot, in [1, 31]. In
	// shorter months the snapshot is taken on the last day instead (day 31 runs
	// on Feb 28 or 29, and on Apr 30), so every month gets exactly one
	// snapshot.
	Day int `json:"day"`
	// The hour of the day, in UTC, at which to take the snapshot, in [0, 23].
	Hour int `json:"hour,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MonthlySchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonthlySchedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NewPipelineSpec struct {
	// Budget policy to set on the newly created pipeline.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Release channel of the underlying pipeline's runtime. Some source table
	// configurations (e.g., read-time CDF) require PREVIEW. Defaults to CURRENT
	// if not specified.
	PipelineChannel NewPipelineSpecPipelineChannel `json:"pipeline_channel,omitempty"`
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog string `json:"storage_catalog,omitempty"`
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	StorageSchema string `json:"storage_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NewPipelineSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NewPipelineSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Release channel of the underlying pipeline's runtime. PREVIEW provides early
// access to the latest features but may be less stable. Some source table
// configurations (e.g., read-time CDF) require PREVIEW. Defaults to CURRENT if
// not specified.
type NewPipelineSpecPipelineChannel string

const NewPipelineSpecPipelineChannelCurrent NewPipelineSpecPipelineChannel = `CURRENT`

const NewPipelineSpecPipelineChannelPreview NewPipelineSpecPipelineChannel = `PREVIEW`

// String representation for [fmt.Print]
func (f *NewPipelineSpecPipelineChannel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NewPipelineSpecPipelineChannel) Set(v string) error {
	switch v {
	case `CURRENT`, `PREVIEW`:
		*f = NewPipelineSpecPipelineChannel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CURRENT", "PREVIEW"`, v)
	}
}

// Values returns all possible values for NewPipelineSpecPipelineChannel.
//
// There is no guarantee on the order of the values in the slice.
func (f *NewPipelineSpecPipelineChannel) Values() []NewPipelineSpecPipelineChannel {
	return []NewPipelineSpecPipelineChannel{
		NewPipelineSpecPipelineChannelCurrent,
		NewPipelineSpecPipelineChannelPreview,
	}
}

// Type always returns NewPipelineSpecPipelineChannel to satisfy [pflag.Value] interface
func (f *NewPipelineSpecPipelineChannel) Type() string {
	return "NewPipelineSpecPipelineChannel"
}

// Controls how the Data API exposes the OpenAPI documentation endpoint. Only
// IGNORE_PRIVILEGES and DISABLED are supported today; "follow-privileges" is
// not implemented yet (it may be added later as value 3 — adding new enum
// values is backward-compatible).
type OpenApiMode string

const OpenApiModeOpenApiModeDisabled OpenApiMode = `OPEN_API_MODE_DISABLED`

const OpenApiModeOpenApiModeIgnorePrivileges OpenApiMode = `OPEN_API_MODE_IGNORE_PRIVILEGES`

// String representation for [fmt.Print]
func (f *OpenApiMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OpenApiMode) Set(v string) error {
	switch v {
	case `OPEN_API_MODE_DISABLED`, `OPEN_API_MODE_IGNORE_PRIVILEGES`:
		*f = OpenApiMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OPEN_API_MODE_DISABLED", "OPEN_API_MODE_IGNORE_PRIVILEGES"`, v)
	}
}

// Values returns all possible values for OpenApiMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *OpenApiMode) Values() []OpenApiMode {
	return []OpenApiMode{
		OpenApiModeOpenApiModeDisabled,
		OpenApiModeOpenApiModeIgnorePrivileges,
	}
}

// Type always returns OpenApiMode to satisfy [pflag.Value] interface
func (f *OpenApiMode) Type() string {
	return "OpenApiMode"
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done,omitempty"`
	// The error result of the operation in case of failure or cancellation.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name string `json:"name,omitempty"`
	// The normal, successful response of the operation.
	Response json.RawMessage `json:"response,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Project struct {
	// A timestamp indicating when the project was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// A timestamp indicating when the project was soft-deleted. Empty if the
	// project is not deleted, otherwise set to a timestamp in the past.
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// Configuration for the initial default branch created as part of project
	// creation. Allows overriding branch protection. These settings only apply
	// at creation time and do not affect resources created after project
	// creation.
	InitialBranchSpec *InitialBranchSpec `json:"initial_branch_spec,omitempty"`
	// Configuration settings for the initial Read/Write endpoint created inside
	// the initial branch for a newly created project. If omitted, the initial
	// endpoint created will have default settings, without high availability
	// configured. This field does not apply to any endpoints created after
	// project creation. Use spec.default_endpoint_settings to configure default
	// settings for endpoints created after project creation.
	InitialEndpointSpec *InitialEndpointSpec `json:"initial_endpoint_spec,omitempty"`
	// Output only. The full resource path of the project. Format:
	// projects/{project_id}
	Name string `json:"name,omitempty"`
	// The project ID chosen when the project was created. It is the final
	// segment of `name`.
	ProjectId string `json:"project_id,omitempty"`
	// A timestamp indicating when the project is scheduled for permanent
	// deletion. Empty if the project is not deleted, otherwise set to a
	// timestamp in the future.
	PurgeTime *time.Time `json:"purge_time,omitempty"`
	// The desired configuration of the project. The fields under `spec` are the
	// ones that can be changed after creation, except `pg_version`, which is
	// fixed when the project is created. See `status` for the configuration
	// that is currently in effect.
	//
	// Read the resource's current values from `status`, rather than `spec`.
	// `spec` tracks field ownership for infrastructure-as-code tools and may
	// omit unmanaged fields.
	Spec *ProjectSpec `json:"spec,omitempty"`
	// The configuration and state of the project that are currently in effect,
	// as reported by the system.
	Status *ProjectStatus `json:"status,omitempty"`
	// A system-generated UUID that identifies the project. Unlike `name`, it is
	// not derived from the user-chosen project ID.
	Uid string `json:"uid,omitempty"`
	// A timestamp indicating when the project was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Project) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Project) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProjectCustomTag struct {
	// The key of the custom tag.
	Key string `json:"key,omitempty"`
	// The value of the custom tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectCustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectCustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A collection of settings for a compute endpoint.
type ProjectDefaultEndpointSettings struct {
	// The maximum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu,omitempty"`
	// The minimum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// When set to true, explicitly disables automatic suspension (never
	// suspend). Should be set to true when provided. Mutually exclusive with
	// `suspend_timeout_duration`.
	NoSuspension bool `json:"no_suspension,omitempty"`
	// A raw representation of Postgres settings.
	PgSettings map[string]string `json:"pg_settings,omitempty"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. If specified should be between 60s and 604800s (1 minute to 1
	// week). Mutually exclusive with `no_suspension`.
	SuspendTimeoutDuration *duration.Duration `json:"suspend_timeout_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectDefaultEndpointSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectDefaultEndpointSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProjectOperationMetadata struct {
}

func (s *ProjectOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ProjectSpec struct {
	// The desired budget policy to associate with the project. See
	// status.budget_policy_id for the policy that is actually applied to the
	// project.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Custom tags to associate with the project, used for billing and cost
	// tracking. To update tags, provide the new tag list and include
	// "spec.custom_tags" in the update_mask. To clear all tags, provide an
	// empty list and include "spec.custom_tags" in the update_mask. To preserve
	// existing tags, omit this field from the update_mask (or use wildcard "*"
	// which auto-excludes empty tags).
	CustomTags []ProjectCustomTag `json:"custom_tags,omitempty"`
	// The full resource path of the default branch of the project. Format:
	// projects/{project_id}/branches/{branch_id} At creation time, the branch
	// ID in this path names the initial branch, which is called `production`
	// when the field is omitted. On update, the path must name an existing
	// branch of this project, which then becomes the default branch.
	DefaultBranch string `json:"default_branch,omitempty"`
	// Settings applied by default to compute endpoints created in this project.
	// The initial endpoint also takes these values for any setting that
	// initial_endpoint_spec does not provide.
	DefaultEndpointSettings *ProjectDefaultEndpointSettings `json:"default_endpoint_settings,omitempty"`
	// Human-readable project name. Length should be between 1 and 256
	// characters.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to enable PG native password login on all endpoints in this
	// project. Defaults to false.
	EnablePgNativeLogin bool `json:"enable_pg_native_login,omitempty"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project. Value should be between
	// 172800s (2 days) and 3024000s (35 days). Defaults to 604800s (7 days)
	// when omitted.
	HistoryRetentionDuration *duration.Duration `json:"history_retention_duration,omitempty"`
	// The major Postgres version number. Defaults to 17 when omitted. The set
	// of versions that new projects can use may change over time.
	PgVersion int `json:"pg_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProjectStatus struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes int64 `json:"branch_logical_size_limit_bytes,omitempty"`
	// The budget policy that is applied to the project.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The most recent time when any endpoint of this project was active.
	ComputeLastActiveTime *time.Time `json:"compute_last_active_time,omitempty"`
	// The compute policy that is applied to the project.
	ComputePolicyId string `json:"compute_policy_id,omitempty"`
	// The effective custom tags associated with the project.
	CustomTags []ProjectCustomTag `json:"custom_tags,omitempty"`
	// The full resource path of the default branch of the project.
	DefaultBranch string `json:"default_branch,omitempty"`
	// The effective default endpoint settings.
	DefaultEndpointSettings *ProjectDefaultEndpointSettings `json:"default_endpoint_settings,omitempty"`
	// The effective human-readable project name.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to enable PG native password login on all endpoints in this
	// project.
	EnablePgNativeLogin bool `json:"enable_pg_native_login,omitempty"`
	// The effective number of seconds to retain the shared history for point in
	// time recovery.
	HistoryRetentionDuration *duration.Duration `json:"history_retention_duration,omitempty"`
	// The email of the project owner.
	Owner string `json:"owner,omitempty"`
	// The effective major Postgres version number.
	PgVersion int `json:"pg_version,omitempty"`
	// The project ID, which is the final segment of the project's resource
	// name.
	ProjectId string `json:"project_id,omitempty"`
	// The current space occupied by the project in storage.
	SyntheticStorageSizeBytes int64 `json:"synthetic_storage_size_bytes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ProjectStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProjectStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ProvisioningInfoState string

const ProvisioningInfoStateActive ProvisioningInfoState = `ACTIVE`

const ProvisioningInfoStateDegraded ProvisioningInfoState = `DEGRADED`

const ProvisioningInfoStateDeleting ProvisioningInfoState = `DELETING`

const ProvisioningInfoStateFailed ProvisioningInfoState = `FAILED`

const ProvisioningInfoStateProvisioning ProvisioningInfoState = `PROVISIONING`

const ProvisioningInfoStateUpdating ProvisioningInfoState = `UPDATING`

// String representation for [fmt.Print]
func (f *ProvisioningInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProvisioningInfoState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DEGRADED`, `DELETING`, `FAILED`, `PROVISIONING`, `UPDATING`:
		*f = ProvisioningInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DEGRADED", "DELETING", "FAILED", "PROVISIONING", "UPDATING"`, v)
	}
}

// Values returns all possible values for ProvisioningInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ProvisioningInfoState) Values() []ProvisioningInfoState {
	return []ProvisioningInfoState{
		ProvisioningInfoStateActive,
		ProvisioningInfoStateDegraded,
		ProvisioningInfoStateDeleting,
		ProvisioningInfoStateFailed,
		ProvisioningInfoStateProvisioning,
		ProvisioningInfoStateUpdating,
	}
}

// Type always returns ProvisioningInfoState to satisfy [pflag.Value] interface
func (f *ProvisioningInfoState) Type() string {
	return "ProvisioningInfoState"
}

// The current phase of the data synchronization pipeline.
type ProvisioningPhase string

const ProvisioningPhaseProvisioningPhaseIndexScan ProvisioningPhase = `PROVISIONING_PHASE_INDEX_SCAN`

const ProvisioningPhaseProvisioningPhaseIndexSort ProvisioningPhase = `PROVISIONING_PHASE_INDEX_SORT`

const ProvisioningPhaseProvisioningPhaseMain ProvisioningPhase = `PROVISIONING_PHASE_MAIN`

// String representation for [fmt.Print]
func (f *ProvisioningPhase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProvisioningPhase) Set(v string) error {
	switch v {
	case `PROVISIONING_PHASE_INDEX_SCAN`, `PROVISIONING_PHASE_INDEX_SORT`, `PROVISIONING_PHASE_MAIN`:
		*f = ProvisioningPhase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PROVISIONING_PHASE_INDEX_SCAN", "PROVISIONING_PHASE_INDEX_SORT", "PROVISIONING_PHASE_MAIN"`, v)
	}
}

// Values returns all possible values for ProvisioningPhase.
//
// There is no guarantee on the order of the values in the slice.
func (f *ProvisioningPhase) Values() []ProvisioningPhase {
	return []ProvisioningPhase{
		ProvisioningPhaseProvisioningPhaseIndexScan,
		ProvisioningPhaseProvisioningPhaseIndexSort,
		ProvisioningPhaseProvisioningPhaseMain,
	}
}

// Type always returns ProvisioningPhase to satisfy [pflag.Value] interface
func (f *ProvisioningPhase) Type() string {
	return "ProvisioningPhase"
}

type RequestedClaims struct {
	PermissionSet RequestedClaimsPermissionSet `json:"permission_set,omitempty"`

	Resources []RequestedResource `json:"resources,omitempty"`
}

func (s *RequestedClaims) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RequestedClaimsPermissionSet string

const RequestedClaimsPermissionSetReadOnly RequestedClaimsPermissionSet = `READ_ONLY`

// String representation for [fmt.Print]
func (f *RequestedClaimsPermissionSet) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RequestedClaimsPermissionSet) Set(v string) error {
	switch v {
	case `READ_ONLY`:
		*f = RequestedClaimsPermissionSet(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ_ONLY"`, v)
	}
}

// Values returns all possible values for RequestedClaimsPermissionSet.
//
// There is no guarantee on the order of the values in the slice.
func (f *RequestedClaimsPermissionSet) Values() []RequestedClaimsPermissionSet {
	return []RequestedClaimsPermissionSet{
		RequestedClaimsPermissionSetReadOnly,
	}
}

// Type always returns RequestedClaimsPermissionSet to satisfy [pflag.Value] interface
func (f *RequestedClaimsPermissionSet) Type() string {
	return "RequestedClaimsPermissionSet"
}

type RequestedResource struct {
	// The full Unity Catalog table name.
	TableName string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RequestedResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RequestedResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to reveal the current password for a password-authenticated Postgres
// role.
type RevealRolePasswordRequest struct {
	// The full resource path of the role whose password should be revealed.
	// Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
}

func (s *RevealRolePasswordRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response containing the password for a password-authenticated Postgres role.
type RevealRolePasswordResponse struct {
	// The current password for the Postgres role.
	Password string `json:"password,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RevealRolePasswordResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RevealRolePasswordResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Role represents a Postgres role within a Branch.
type Role struct {
	// A timestamp indicating when the role was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. The full resource path of the role. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"name,omitempty"`
	// The Branch where this Role exists. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent string `json:"parent,omitempty"`
	// The part of the name, chosen by the user when the resource was created.
	RoleId string `json:"role_id,omitempty"`
	// The spec contains the role configuration, including identity type,
	// authentication method, and role attributes.
	//
	// Read the resource's current values from `status`, rather than `spec`.
	// `spec` tracks field ownership for infrastructure-as-code tools and may
	// omit unmanaged fields.
	Spec *RoleRoleSpec `json:"spec,omitempty"`
	// Current status of the role, including its identity type, authentication
	// method, and role attributes.
	Status *RoleRoleStatus `json:"status,omitempty"`
	// A timestamp indicating when the role was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Role) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Role) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Attributes that can be granted to a Postgres role. We are only implementing a
// subset for now, see xref:
// https://www.postgresql.org/docs/16/sql-createrole.html The values follow
// Postgres keyword naming e.g. CREATEDB, BYPASSRLS, etc. which is why they
// don't include typical underscores between words.
type RoleAttributes struct {
	// Whether the role bypasses every row-level security policy (the Postgres
	// BYPASSRLS attribute).
	Bypassrls bool `json:"bypassrls,omitempty"`
	// Whether the role can create databases (the Postgres CREATEDB attribute).
	Createdb bool `json:"createdb,omitempty"`
	// Whether the role can create, alter and drop other roles (the Postgres
	// CREATEROLE attribute).
	Createrole bool `json:"createrole,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RoleAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RoleAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// How the role is authenticated when connecting to Postgres.
type RoleAuthMethod string

const RoleAuthMethodLakebaseOauthV1 RoleAuthMethod = `LAKEBASE_OAUTH_V1`

const RoleAuthMethodNoLogin RoleAuthMethod = `NO_LOGIN`

const RoleAuthMethodPgPasswordScramSha256 RoleAuthMethod = `PG_PASSWORD_SCRAM_SHA_256`

// String representation for [fmt.Print]
func (f *RoleAuthMethod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RoleAuthMethod) Set(v string) error {
	switch v {
	case `LAKEBASE_OAUTH_V1`, `NO_LOGIN`, `PG_PASSWORD_SCRAM_SHA_256`:
		*f = RoleAuthMethod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LAKEBASE_OAUTH_V1", "NO_LOGIN", "PG_PASSWORD_SCRAM_SHA_256"`, v)
	}
}

// Values returns all possible values for RoleAuthMethod.
//
// There is no guarantee on the order of the values in the slice.
func (f *RoleAuthMethod) Values() []RoleAuthMethod {
	return []RoleAuthMethod{
		RoleAuthMethodLakebaseOauthV1,
		RoleAuthMethodNoLogin,
		RoleAuthMethodPgPasswordScramSha256,
	}
}

// Type always returns RoleAuthMethod to satisfy [pflag.Value] interface
func (f *RoleAuthMethod) Type() string {
	return "RoleAuthMethod"
}

// The type of the Databricks managed identity that this Role represents. Leave
// empty if you wish to create a regular Postgres role not associated with a
// Databricks identity.
type RoleIdentityType string

const RoleIdentityTypeGroup RoleIdentityType = `GROUP`

const RoleIdentityTypeServicePrincipal RoleIdentityType = `SERVICE_PRINCIPAL`

const RoleIdentityTypeUser RoleIdentityType = `USER`

// String representation for [fmt.Print]
func (f *RoleIdentityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RoleIdentityType) Set(v string) error {
	switch v {
	case `GROUP`, `SERVICE_PRINCIPAL`, `USER`:
		*f = RoleIdentityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GROUP", "SERVICE_PRINCIPAL", "USER"`, v)
	}
}

// Values returns all possible values for RoleIdentityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RoleIdentityType) Values() []RoleIdentityType {
	return []RoleIdentityType{
		RoleIdentityTypeGroup,
		RoleIdentityTypeServicePrincipal,
		RoleIdentityTypeUser,
	}
}

// Type always returns RoleIdentityType to satisfy [pflag.Value] interface
func (f *RoleIdentityType) Type() string {
	return "RoleIdentityType"
}

// Roles that the DatabaseInstanceRole can be a member of.
type RoleMembershipRole string

const RoleMembershipRoleDatabricksSuperuser RoleMembershipRole = `DATABRICKS_SUPERUSER`

// String representation for [fmt.Print]
func (f *RoleMembershipRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RoleMembershipRole) Set(v string) error {
	switch v {
	case `DATABRICKS_SUPERUSER`:
		*f = RoleMembershipRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS_SUPERUSER"`, v)
	}
}

// Values returns all possible values for RoleMembershipRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *RoleMembershipRole) Values() []RoleMembershipRole {
	return []RoleMembershipRole{
		RoleMembershipRoleDatabricksSuperuser,
	}
}

// Type always returns RoleMembershipRole to satisfy [pflag.Value] interface
func (f *RoleMembershipRole) Type() string {
	return "RoleMembershipRole"
}

type RoleOperationMetadata struct {
}

func (s *RoleOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RoleRoleSpec struct {
	// The Postgres role attributes to grant to the role.
	Attributes *RoleAttributes `json:"attributes,omitempty"`
	// Controls how the Postgres role authenticates when a client opens a
	// database connection. Supported values:
	//
	// * LAKEBASE_OAUTH_V1: the role authenticates by presenting a Databricks
	// OAuth access token derived from the backing managed identity (the
	// Databricks user, service principal, or group named by the role's
	// `postgres_role`). No static password exists for roles using this method.
	// * PG_PASSWORD_SCRAM_SHA_256: the role authenticates with a Postgres
	// password verified server-side using the SCRAM-SHA-256 mechanism. Lakebase
	// generates a password for the role.
	//
	// NO_LOGIN is not currently supported for Role creation.
	//
	// If auth_method is left unspecified, a meaningful authentication method is
	// derived from the identity_type: * For the managed identities, OAUTH is
	// used. * For the regular postgres roles, authentication based on postgres
	// passwords is used.
	//
	// NOTE: for the Databricks identity type GROUP, LAKEBASE_OAUTH_V1 is the
	// default auth method (group can login as well).
	//
	// This value cannot be changed after creation.
	AuthMethod RoleAuthMethod `json:"auth_method,omitempty"`
	// The type of Databricks identity that backs the role. Leave unset to
	// create a native Postgres role that is not associated with a Databricks
	// identity. When specifying a managed-identity, the chosen postgres_role
	// must be a valid:
	//
	// * application ID for SERVICE_PRINCIPAL * user email for USER * group name
	// for GROUP
	//
	// This value cannot be changed after creation.
	IdentityType RoleIdentityType `json:"identity_type,omitempty"`
	// An enum value for a standard role that this role is a member of. When
	// selected in update_mask, this field replaces all existing memberships. An
	// empty or omitted list clears all memberships, including
	// DATABRICKS_SUPERUSER.
	MembershipRoles []RoleMembershipRole `json:"membership_roles,omitempty"`
	// The name of the Postgres role.
	//
	// This expects a valid Postgres identifier as specified in the link below.
	// https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
	//
	// Required when creating the Role.
	//
	// If you wish to create a Postgres Role backed by a managed Databricks
	// identity, then postgres_role must be one of the following:
	//
	// 1. user email for IdentityType.USER 2. app ID for
	// IdentityType.SERVICE_PRINCIPAL 3. group name for IdentityType.GROUP
	//
	// This value cannot be changed after creation.
	PostgresRole string `json:"postgres_role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RoleRoleSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RoleRoleSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RoleRoleStatus struct {
	// The Postgres role attributes currently granted to the role.
	Attributes *RoleAttributes `json:"attributes,omitempty"`
	// How the role authenticates when it connects to Postgres.
	AuthMethod RoleAuthMethod `json:"auth_method,omitempty"`
	// The type of Databricks identity that backs the role, if any.
	IdentityType RoleIdentityType `json:"identity_type,omitempty"`
	// The standard roles that the role is currently a member of.
	MembershipRoles []RoleMembershipRole `json:"membership_roles,omitempty"`
	// The name of the Postgres role.
	PostgresRole string `json:"postgres_role,omitempty"`
	// Part of the resource name.
	RoleId string `json:"role_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RoleRoleStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RoleRoleStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// One cadence at which automatic snapshots are taken.
type ScheduleCadence struct {
	// Take a snapshot once per day.
	DailySchedule *DailySchedule `json:"daily_schedule,omitempty"`
	// Take a snapshot once per month.
	MonthlySchedule *MonthlySchedule `json:"monthly_schedule,omitempty"`
	// How long snapshots from this cadence are kept before automatic deletion.
	// Must be at least 1 hour. Applied when a snapshot is taken; not
	// retroactive, so changing it affects only later snapshots.
	Retention duration.Duration `json:"retention"`
	// Take a snapshot once per week.
	WeeklySchedule *WeeklySchedule `json:"weekly_schedule,omitempty"`
}

func (s *ScheduleCadence) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An immutable, point-in-time copy of a branch's data within a project. It
// remains available after the source branch is deleted. To restore a snapshot,
// create a new branch with `spec.source_snapshot` set to the snapshot's
// resource name. For an overview, see
// https://docs.databricks.com/aws/en/oltp/projects/snapshots.
type Snapshot struct {
	// When the snapshot was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The resource name of the snapshot. Format:
	// projects/{project_id}/snapshots/{snapshot_id}
	Name string `json:"name,omitempty"`
	// The user-chosen ID; the final segment of `name`.
	SnapshotId string `json:"snapshot_id,omitempty"`
	// Client-provided configuration of the snapshot. When you create a
	// snapshot, set exactly one of `expire_time`, `ttl`, or `no_expiry`.
	Spec *SnapshotSpec `json:"spec,omitempty"`
	// Server-observed state of the snapshot.
	Status *SnapshotStatus `json:"status,omitempty"`
	// Unique system-generated ID for the snapshot.
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Snapshot) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Snapshot) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for the long-running snapshot Create and Delete operations.
type SnapshotOperationMetadata struct {
}

func (s *SnapshotOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The automatic snapshot cadences for a branch. There is exactly one schedule
// per branch (singleton); it is configured in place, not created or deleted.
//
// Name: projects/{project_id}/branches/{branch_id}/snapshot-schedule
type SnapshotSchedule struct {
	// The resource name of the branch's snapshot schedule. Format:
	// projects/{project_id}/branches/{branch_id}/snapshot-schedule
	Name string `json:"name,omitempty"`
	// The cadences at which automatic snapshots are taken. Update replaces the
	// whole set; an empty set disables automatic snapshots. Order is not
	// significant. When several cadences fire together, one snapshot is taken,
	// retained for the longest of their retentions.
	Schedule []ScheduleCadence `json:"schedule,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SnapshotSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SnapshotSchedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for the long-running snapshot schedule Update operation.
type SnapshotScheduleOperationMetadata struct {
}

func (s *SnapshotScheduleOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Client-provided configuration of the snapshot.
type SnapshotSpec struct {
	// Absolute time at which the snapshot is deleted. It must not be in the
	// past. Mutually exclusive with `ttl` and `no_expiry`.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// If true, the snapshot never expires. The value `false` is rejected; set
	// `expire_time` or `ttl` instead. Mutually exclusive with `ttl` and
	// `expire_time`.
	NoExpiry bool `json:"no_expiry,omitempty"`
	// The source branch to snapshot. The branch must belong to the same project
	// as the snapshot. Format: projects/{project_id}/branches/{branch_id}
	SourceBranch string `json:"source_branch"`
	// Postgres log sequence number (LSN) to snapshot from. It must fall within
	// the project's history retention window. If neither this field nor
	// `source_branch_time` is set, the snapshot captures the current head of
	// the source branch. Mutually exclusive with `source_branch_time`.
	SourceBranchLsn string `json:"source_branch_lsn,omitempty"`
	// Timestamp to snapshot from. It must fall within the project's history
	// retention window. Mutually exclusive with `source_branch_lsn`.
	SourceBranchTime *time.Time `json:"source_branch_time,omitempty"`
	// Time-to-live, as a positive duration in seconds with an `s` suffix. The
	// snapshot expires this long after it is created. Mutually exclusive with
	// `expire_time` and `no_expiry`. Reads report the resolved absolute
	// `expire_time` instead.
	Ttl *duration.Duration `json:"ttl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SnapshotSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SnapshotSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Server-observed state of a snapshot.
type SnapshotStatus struct {
	// Incremental storage size in bytes since the previous snapshot. Unset when
	// the snapshot is not billed on incremental usage.
	DiffSizeBytes int64 `json:"diff_size_bytes,omitempty"`
	// Absolute time at which the snapshot is deleted.
	ExpireTime *time.Time `json:"expire_time,omitempty"`
	// Full logical size of the snapshot, in bytes.
	FullSizeBytes int64 `json:"full_size_bytes,omitempty"`
	// True if the snapshot never expires.
	NoExpiry bool `json:"no_expiry,omitempty"`
	// The source branch the snapshot was taken from. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch string `json:"source_branch,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SnapshotStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SnapshotStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to start a migration.
type StartMigrationRequest struct {
	// Required. The migration resource name, in the format
	// `migrations/{migration}`.
	Name string `json:"-" url:"-"`
}

func (s *StartMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response produced after migration preparation has started.
type StartMigrationResponse struct {
	// The migration after the start action is applied.
	Migration *Migration `json:"migration,omitempty"`
}

func (s *StartMigrationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A synced table is a read-only Postgres table that Databricks keeps up to date
// from a Unity Catalog source table through a managed pipeline. It is also
// registered in Unity Catalog under its three-part name. See
// https://docs.databricks.com/aws/en/oltp/projects/sync-tables.
type SyncedTable struct {
	// The time when the synced table was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The full resource name of the synced table, where (catalog, schema,
	// table) are the Unity Catalog names of the synced table.
	//
	// Format: "synced_tables/{catalog}.{schema}.{table}".
	//
	// For the corresponding source table in Unity Catalog, look for the
	// "source_table_full_name" attribute.
	Name string `json:"name,omitempty"`
	// Configuration details of the synced table, such as the source table,
	// scheduling policy, etc.
	Spec *SyncedTableSyncedTableSpec `json:"spec,omitempty"`
	// Synced Table data synchronization status.
	Status *SyncedTableSyncedTableStatus `json:"status,omitempty"`
	// The three-part Unity Catalog name of the synced table, in the format
	// "{catalog}.{schema}.{table}". It is the synced_table_id that was used to
	// create the synced table and the last component of name.
	SyncedTableId string `json:"synced_table_id,omitempty"`
	// The Unity Catalog table ID for this synced table.
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for SyncedTable long-running operations.
type SyncedTableOperationMetadata struct {
}

func (s *SyncedTableOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Progress information of the Synced Table data synchronization pipeline.
type SyncedTablePipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds float64 `json:"estimated_completion_time_seconds,omitempty"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing int64 `json:"latest_version_currently_processing,omitempty"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion float64 `json:"sync_progress_completion,omitempty"`
	// The number of rows that have been synced in this update.
	SyncedRowCount int64 `json:"synced_row_count,omitempty"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount int64 `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTablePipelineProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePipelineProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SyncedTablePosition struct {
	// The source Delta table version that the last successful sync read.
	DeltaTableSyncInfo *DeltaTableSyncInfo `json:"delta_table_sync_info,omitempty"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	SyncEndTime *time.Time `json:"sync_end_time,omitempty"`
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	SyncStartTime *time.Time `json:"sync_start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTablePosition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePosition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of a synced table.
type SyncedTableState string

const SyncedTableStateSyncedTableOffline SyncedTableState = `SYNCED_TABLE_OFFLINE`

const SyncedTableStateSyncedTableOfflineFailed SyncedTableState = `SYNCED_TABLE_OFFLINE_FAILED`

const SyncedTableStateSyncedTableOnline SyncedTableState = `SYNCED_TABLE_ONLINE`

const SyncedTableStateSyncedTableOnlineContinuousUpdate SyncedTableState = `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`

const SyncedTableStateSyncedTableOnlineNoPendingUpdate SyncedTableState = `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`

const SyncedTableStateSyncedTableOnlinePipelineFailed SyncedTableState = `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`

const SyncedTableStateSyncedTableOnlineTriggeredUpdate SyncedTableState = `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`

const SyncedTableStateSyncedTableOnlineUpdatingPipelineResources SyncedTableState = `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`

const SyncedTableStateSyncedTableProvisioning SyncedTableState = `SYNCED_TABLE_PROVISIONING`

const SyncedTableStateSyncedTableProvisioningInitialSnapshot SyncedTableState = `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`

const SyncedTableStateSyncedTableProvisioningPipelineResources SyncedTableState = `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`

// String representation for [fmt.Print]
func (f *SyncedTableState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableState) Set(v string) error {
	switch v {
	case `SYNCED_TABLE_OFFLINE`, `SYNCED_TABLE_OFFLINE_FAILED`, `SYNCED_TABLE_ONLINE`, `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`, `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`, `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`, `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`, `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`, `SYNCED_TABLE_PROVISIONING`, `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`, `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`:
		*f = SyncedTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SYNCED_TABLE_OFFLINE", "SYNCED_TABLE_OFFLINE_FAILED", "SYNCED_TABLE_ONLINE", "SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE", "SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE", "SYNCED_TABLE_ONLINE_PIPELINE_FAILED", "SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE", "SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES", "SYNCED_TABLE_PROVISIONING", "SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT", "SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Values returns all possible values for SyncedTableState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableState) Values() []SyncedTableState {
	return []SyncedTableState{
		SyncedTableStateSyncedTableOffline,
		SyncedTableStateSyncedTableOfflineFailed,
		SyncedTableStateSyncedTableOnline,
		SyncedTableStateSyncedTableOnlineContinuousUpdate,
		SyncedTableStateSyncedTableOnlineNoPendingUpdate,
		SyncedTableStateSyncedTableOnlinePipelineFailed,
		SyncedTableStateSyncedTableOnlineTriggeredUpdate,
		SyncedTableStateSyncedTableOnlineUpdatingPipelineResources,
		SyncedTableStateSyncedTableProvisioning,
		SyncedTableStateSyncedTableProvisioningInitialSnapshot,
		SyncedTableStateSyncedTableProvisioningPipelineResources,
	}
}

// Type always returns SyncedTableState to satisfy [pflag.Value] interface
func (f *SyncedTableState) Type() string {
	return "SyncedTableState"
}

type SyncedTableSyncedTableSpec struct {
	// When true, enables accelerated sync mode for the initial data load. This
	// significantly improves performance for large tables. Requires
	// workspace-level enablement through Lakebase Accelerated Sync preview.
	AcceleratedSync bool `json:"accelerated_sync,omitempty"`
	// The full resource name of the branch that contains the synced table.
	//
	// Format: "projects/{project_id}/branches/{branch_id}".
	//
	// A value must be specified when creating a synced table inside a standard
	// catalog. Inside a Lakebase catalog it can be omitted, and if specified it
	// must match the branch of that catalog.
	Branch string `json:"branch,omitempty"`
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist. The request will fail if
	// this is false and the database/schema do not exist.
	//
	// Defaults to true if omitted.
	CreateDatabaseObjectsIfMissing bool `json:"create_database_objects_if_missing,omitempty"`
	// ID of an existing pipeline to bin-pack this synced table into. The
	// pipeline must be a synced table pipeline that runs continuously if and
	// only if scheduling_policy is CONTINUOUS. At most one of
	// existing_pipeline_id and new_pipeline_spec should be defined.
	//
	// The pipeline used for the synced table is returned in status.pipeline_id.
	ExistingPipelineId string `json:"existing_pipeline_id,omitempty"`
	// Extra PostgreSQL-only columns to add to the synced table.
	ExtraColumns []SyncedTableSyncedTableSpecExtraColumn `json:"extra_columns,omitempty"`
	// Specification for creating a new pipeline. At most one of
	// existing_pipeline_id and new_pipeline_spec should be defined.
	//
	// The pipeline used for the synced table is returned in status.pipeline_id.
	NewPipelineSpec *NewPipelineSpec `json:"new_pipeline_spec,omitempty"`
	// The name of the Postgres database where the synced table is created. A
	// value must be specified on creation. If the synced table is created
	// inside a Lakebase catalog, the value must match the Postgres database of
	// that catalog.
	PostgresDatabase string `json:"postgres_database,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	// At least one column must be specified.
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Scheduling policy of the underlying pipeline. A value must be specified
	// on creation. See
	// https://docs.databricks.com/aws/en/oltp/projects/sync-tables for how to
	// choose a sync mode.
	SchedulingPolicy SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy `json:"scheduling_policy,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	//
	// For the corresponding destination table, use any of the two:
	//
	// * synced_table_id used at the creation of the SyncedTable * "name"
	// consisting of "synced_tables/" prefix and the full name of the
	// destination table.
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string `json:"timeseries_key,omitempty"`
	// Override the default Delta->PG type mapping for specific columns. A
	// TypeOverride with PG_SPECIFIC_TYPE_UNSPECIFIED is rejected; a valid
	// pg_type must be set.
	TypeOverrides []SyncedTableSyncedTableSpecTypeOverride `json:"type_overrides,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSyncedTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSyncedTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An extra PostgreSQL column to add to the synced table.
type SyncedTableSyncedTableSpecExtraColumn struct {
	// Name of the column.
	ColumnName string `json:"column_name"`
	// PostgreSQL type of the column, for example "tsvector" or "vector(1024)".
	ColumnType string `json:"column_type"`
	// SQL expression used to compute the column's value, for example
	// "to_tsvector('english', content)".
	Compute string `json:"compute,omitempty"`

	Maintenance SyncedTableSyncedTableSpecExtraColumnMaintenance `json:"maintenance,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSyncedTableSpecExtraColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSyncedTableSpecExtraColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// How the column's value is populated and kept up to date.
type SyncedTableSyncedTableSpecExtraColumnMaintenance string

const SyncedTableSyncedTableSpecExtraColumnMaintenanceStoredGenerated SyncedTableSyncedTableSpecExtraColumnMaintenance = `STORED_GENERATED`

// String representation for [fmt.Print]
func (f *SyncedTableSyncedTableSpecExtraColumnMaintenance) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableSyncedTableSpecExtraColumnMaintenance) Set(v string) error {
	switch v {
	case `STORED_GENERATED`:
		*f = SyncedTableSyncedTableSpecExtraColumnMaintenance(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "STORED_GENERATED"`, v)
	}
}

// Values returns all possible values for SyncedTableSyncedTableSpecExtraColumnMaintenance.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableSyncedTableSpecExtraColumnMaintenance) Values() []SyncedTableSyncedTableSpecExtraColumnMaintenance {
	return []SyncedTableSyncedTableSpecExtraColumnMaintenance{
		SyncedTableSyncedTableSpecExtraColumnMaintenanceStoredGenerated,
	}
}

// Type always returns SyncedTableSyncedTableSpecExtraColumnMaintenance to satisfy [pflag.Value] interface
func (f *SyncedTableSyncedTableSpecExtraColumnMaintenance) Type() string {
	return "SyncedTableSyncedTableSpecExtraColumnMaintenance"
}

// PostgreSQL-specific target types that can override the default Delta-to-PG
// mapping.
type SyncedTableSyncedTableSpecPgSpecificType string

const SyncedTableSyncedTableSpecPgSpecificTypePgSpecificTypeHalfvec SyncedTableSyncedTableSpecPgSpecificType = `PG_SPECIFIC_TYPE_HALFVEC`

const SyncedTableSyncedTableSpecPgSpecificTypePgSpecificTypeVarchar SyncedTableSyncedTableSpecPgSpecificType = `PG_SPECIFIC_TYPE_VARCHAR`

const SyncedTableSyncedTableSpecPgSpecificTypePgSpecificTypeVector SyncedTableSyncedTableSpecPgSpecificType = `PG_SPECIFIC_TYPE_VECTOR`

// String representation for [fmt.Print]
func (f *SyncedTableSyncedTableSpecPgSpecificType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableSyncedTableSpecPgSpecificType) Set(v string) error {
	switch v {
	case `PG_SPECIFIC_TYPE_HALFVEC`, `PG_SPECIFIC_TYPE_VARCHAR`, `PG_SPECIFIC_TYPE_VECTOR`:
		*f = SyncedTableSyncedTableSpecPgSpecificType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PG_SPECIFIC_TYPE_HALFVEC", "PG_SPECIFIC_TYPE_VARCHAR", "PG_SPECIFIC_TYPE_VECTOR"`, v)
	}
}

// Values returns all possible values for SyncedTableSyncedTableSpecPgSpecificType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableSyncedTableSpecPgSpecificType) Values() []SyncedTableSyncedTableSpecPgSpecificType {
	return []SyncedTableSyncedTableSpecPgSpecificType{
		SyncedTableSyncedTableSpecPgSpecificTypePgSpecificTypeHalfvec,
		SyncedTableSyncedTableSpecPgSpecificTypePgSpecificTypeVarchar,
		SyncedTableSyncedTableSpecPgSpecificTypePgSpecificTypeVector,
	}
}

// Type always returns SyncedTableSyncedTableSpecPgSpecificType to satisfy [pflag.Value] interface
func (f *SyncedTableSyncedTableSpecPgSpecificType) Type() string {
	return "SyncedTableSyncedTableSpecPgSpecificType"
}

// Scheduling policy of the synced table's underlying pipeline.
type SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy string

const SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyContinuous SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy = `CONTINUOUS`

const SyncedTableSyncedTableSpecSyncedTableSchedulingPolicySnapshot SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy = `SNAPSHOT`

const SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyTriggered SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy = `TRIGGERED`

// String representation for [fmt.Print]
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`:
		*f = SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "SNAPSHOT", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) Values() []SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy {
	return []SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy{
		SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyContinuous,
		SyncedTableSyncedTableSpecSyncedTableSchedulingPolicySnapshot,
		SyncedTableSyncedTableSpecSyncedTableSchedulingPolicyTriggered,
	}
}

// Type always returns SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy to satisfy [pflag.Value] interface
func (f *SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy) Type() string {
	return "SyncedTableSyncedTableSpecSyncedTableSchedulingPolicy"
}

// Overrides the default Delta-to-PostgreSQL type mapping for a single column.
type SyncedTableSyncedTableSpecTypeOverride struct {
	// Name of the source column whose target PostgreSQL type should be
	// overridden.
	ColumnName string `json:"column_name"`
	// PostgreSQL-specific target type to use for the column.
	PgType SyncedTableSyncedTableSpecPgSpecificType `json:"pg_type"`
	// Size parameter for the target type, for types that take one (e.g. vector
	// dimension, varchar length). Required when the chosen pg_type needs a
	// size.
	Size int `json:"size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSyncedTableSpecTypeOverride) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSyncedTableSpecTypeOverride) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SyncedTableSyncedTableStatus struct {
	// The state of the synced table.
	DetailedState SyncedTableState `json:"detailed_state,omitempty"`
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// Summary of the last successful synchronization from source to
	// destination.
	LastSync *SyncedTablePosition `json:"last_sync,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	LastSyncTime *time.Time `json:"last_sync_time,omitempty"`
	// A text description of the current state of the synced table.
	Message string `json:"message,omitempty"`
	// Progress of the sync that is currently running, either the initial data
	// load or an update of a triggered or snapshot synced table.
	OngoingSyncProgress *SyncedTablePipelineProgress `json:"ongoing_sync_progress,omitempty"`
	// ID of the associated pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The full resource name of the project associated with the table.
	//
	// Format: "projects/{project_id}".
	Project string `json:"project,omitempty"`
	// The current phase of the data synchronization pipeline.
	ProvisioningPhase ProvisioningPhase `json:"provisioning_phase,omitempty"`
	// The provisioning state of the synced table entity in Unity Catalog.
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSyncedTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSyncedTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UndeleteBranchRequest struct {
	// The full resource path of the branch to undelete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
}

func (s *UndeleteBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to restore a soft-deleted migration before it is permanently purged.
type UndeleteMigrationRequest struct {
	// Required. The migration resource name, in the format
	// `migrations/{migration}`.
	Name string `json:"-" url:"-"`
}

func (s *UndeleteMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response produced after a soft-deleted migration is restored.
type UndeleteMigrationResponse struct {
	// The restored migration.
	Migration *Migration `json:"migration,omitempty"`
}

func (s *UndeleteMigrationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to restore a soft-deleted project within its retention period.
type UndeleteProjectRequest struct {
	// The full resource path of the project to undelete. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
}

func (s *UndeleteProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateBranchRequest struct {
	// The Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	Branch Branch `json:"branch"`
	// Output only. The full resource path of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateBranchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateDataApiRequest struct {
	// The Data API configuration to update. The data_api's `name` field
	// identifies the resource.
	DataApi DataApi `json:"data_api"`
	// Resource name:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}/data-api
	Name string `json:"-" url:"-"`
	// The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateDataApiRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateDatabaseRequest struct {
	// The Database to update.
	//
	// The database's `name` field is used to identify the database to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Database Database `json:"database"`
	// The resource name of the database. Format:
	// projects/{project_id}/branches/{branch_id}/databases/{database_id}
	Name string `json:"-" url:"-"`
	// The list of fields to update. The updatable fields are `spec.role` and
	// `spec.postgres_database`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateDatabaseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateEndpointRequest struct {
	// The Endpoint to update.
	//
	// The endpoint's `name` field is used to identify the endpoint to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint Endpoint `json:"endpoint"`
	// Output only. The full resource path of the endpoint. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name string `json:"-" url:"-"`
	// The fields to update, as paths relative to the endpoint, for example
	// `spec.autoscaling_limit_max_cu`. Pass `spec` to update every field that
	// is set in `spec`. Only fields under `spec` can be updated. To change the
	// suspension setting, include `spec.suspension` or the specific suspension
	// field.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateMigrationRequest struct {
	// Required. The migration containing the user-owned fields to update.
	Migration Migration `json:"migration"`
	// The migration resource name, in the format `migrations/{migration}`.
	Name string `json:"-" url:"-"`
	// Required. The fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateProjectRequest struct {
	// Output only. The full resource path of the project. Format:
	// projects/{project_id}
	Name string `json:"-" url:"-"`
	// The Project to update.
	//
	// The project's `name` field is used to identify the project to update.
	// Format: projects/{project_id}
	Project Project `json:"project"`
	// The fields to update, as paths relative to the project, for example
	// `spec.display_name`. Pass `*` to update every field that is set in the
	// request. Only fields under `spec` can be updated.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateProjectRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateRoleRequest struct {
	// Output only. The full resource path of the role. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name string `json:"-" url:"-"`
	// The Postgres Role to update.
	//
	// The role's `name` field is used to identify the role to update. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Role Role `json:"role"`
	// The list of fields to update. Supported fields are
	// `spec.membership_roles`, `spec.attributes.createdb`,
	// `spec.attributes.createrole`, and `spec.attributes.bypassrls`.
	//
	// `spec.identity_type`, `spec.auth_method`, and `spec.postgres_role` cannot
	// be changed. Including an immutable field as an explicit path in
	// update_mask returns INVALID_PARAMETER_VALUE. With `*` or masks containing
	// only mutable fields, supplied values for these immutable fields are
	// ignored. `*` by itself selects all supported mutable fields. For this
	// mask, all three attributes (`createdb`, `createrole`, and `bypassrls`)
	// must be supplied; omitting any returns an INVALID_PARAMETER_VALUE error.
	// The `*` mask replaces `membership_roles`. Omitting it or supplying an
	// empty list clears all memberships, including `DATABRICKS_SUPERUSER`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSnapshotScheduleRequest struct {
	// The resource name of the branch's snapshot schedule. Format:
	// projects/{project_id}/branches/{branch_id}/snapshot-schedule
	Name string `json:"-" url:"-"`
	// The snapshot schedule to set. Its `name` identifies the branch. Format:
	// projects/{project_id}/branches/{branch_id}/snapshot-schedule
	SnapshotSchedule SnapshotSchedule `json:"snapshot_schedule"`
	// Fields to update. The only updatable path is `schedule`, which replaces
	// the entire set of cadences.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSnapshotScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Take a snapshot once per week, on the configured day at the configured hour.
type WeeklySchedule struct {
	// The day of the week on which to take the snapshot.
	DayOfWeek DayOfWeek `json:"day_of_week"`
	// The hour of the day, in UTC, at which to take the snapshot, in [0, 23].
	Hour int `json:"hour,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WeeklySchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WeeklySchedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
