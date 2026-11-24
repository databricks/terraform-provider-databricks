// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDatabaseCatalogRequest struct {
	Catalog DatabaseCatalog `json:"catalog"`
}

func (s *CreateDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateDatabaseInstanceRequest struct {
	// Instance to create.
	DatabaseInstance DatabaseInstance `json:"database_instance"`
}

func (s *CreateDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateDatabaseInstanceRoleRequest struct {
	DatabaseInstanceName string `json:"-" url:"database_instance_name,omitempty"`

	DatabaseInstanceRole DatabaseInstanceRole `json:"database_instance_role"`

	InstanceName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDatabaseTableRequest struct {
	Table DatabaseTable `json:"table"`
}

func (s *CreateDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateSyncedDatabaseTableRequest struct {
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
}

func (s *CreateSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomTag struct {
	// The key of the custom tag.
	Key string `json:"key,omitempty"`
	// The value of the custom tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Database Catalog registers a Postgres database that lives in a Lakebase
// database instance as a catalog in Unity Catalog, so that the database's
// schemas and tables are governed and discoverable through Unity Catalog. For
// an overview, see
// https://docs.databricks.com/aws/en/oltp/projects/register-uc.
type DatabaseCatalog struct {
	// If true, the database named in database_name is created in the instance
	// when it does not already exist, with the calling user as its owner.
	// Defaults to false.
	CreateDatabaseIfNotExists bool `json:"create_database_if_not_exists,omitempty"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName string `json:"database_instance_name"`
	// The name of the Postgres database in the instance that the catalog
	// registers. The database must already exist unless
	// create_database_if_not_exists is true. The system databases postgres,
	// databricks_system, template0 and template1 cannot be registered.
	DatabaseName string `json:"database_name"`
	// The name of the catalog in Unity Catalog.
	Name string `json:"name"`
	// An immutable UUID identifier for the catalog, assigned when the catalog
	// is created.
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A short-lived credential for connecting to database instances with the
// caller's Databricks identity.
type DatabaseCredential struct {
	// The time at which the token expires. Tokens are valid for one hour after
	// they are issued, and a new token is needed to open connections after that
	// time.
	ExpirationTime string `json:"expiration_time,omitempty"`
	// The OAuth token to use as the Postgres password when connecting to a
	// database instance.
	Token string `json:"token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	// Must be set when creating an instance.
	Capacity string `json:"capacity,omitempty"`
	// The refs of the child instances. This is only available if the instance
	// is parent instance.
	ChildInstanceRefs []DatabaseInstanceRef `json:"child_instance_refs,omitempty"`
	// The timestamp when the instance was created.
	CreationTime string `json:"creation_time,omitempty"`
	// The email of the creator of the instance.
	Creator string `json:"creator,omitempty"`
	// Custom tags associated with the instance. This field is only included on
	// create and update responses. An instance can have at most 10 tags; keys
	// must be unique ignoring case.
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// Deprecated: this field always matches the value of capacity; read
	// capacity instead. This is an output only field that contains the value
	// computed from the input field combined with server side defaults. Use the
	// field without the effective_ prefix to set the value.
	EffectiveCapacity string `json:"effective_capacity,omitempty"`
	// The recorded custom tags associated with the instance. This is an output
	// only field that contains the value computed from the input field combined
	// with server side defaults. Use the field without the effective_ prefix to
	// set the value.
	EffectiveCustomTags []CustomTag `json:"effective_custom_tags,omitempty"`
	// Whether the instance has PG native password login enabled. This is an
	// output only field that contains the value computed from the input field
	// combined with server side defaults. Use the field without the effective_
	// prefix to set the value.
	EffectiveEnablePgNativeLogin bool `json:"effective_enable_pg_native_login,omitempty"`
	// Whether secondaries serving read-only traffic are enabled. Defaults to
	// false. This is an output only field that contains the value computed from
	// the input field combined with server side defaults. Use the field without
	// the effective_ prefix to set the value.
	EffectiveEnableReadableSecondaries bool `json:"effective_enable_readable_secondaries,omitempty"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries. This is an output
	// only field that contains the value computed from the input field combined
	// with server side defaults. Use the field without the effective_ prefix to
	// set the value.
	EffectiveNodeCount int `json:"effective_node_count,omitempty"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained. This is an output only field
	// that contains the value computed from the input field combined with
	// server side defaults. Use the field without the effective_ prefix to set
	// the value.
	EffectiveRetentionWindowInDays int `json:"effective_retention_window_in_days,omitempty"`
	// Whether the instance is stopped. This is an output only field that
	// contains the value computed from the input field combined with server
	// side defaults. Use the field without the effective_ prefix to set the
	// value.
	EffectiveStopped bool `json:"effective_stopped,omitempty"`
	// The policy that is applied to the instance. This is an output only field
	// that contains the value computed from the input field combined with
	// server side defaults. Use the field without the effective_ prefix to set
	// the value.
	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// Whether to enable PG native password login on the instance. Defaults to
	// false.
	EnablePgNativeLogin bool `json:"enable_pg_native_login,omitempty"`
	// Whether to enable secondaries to serve read-only traffic. Defaults to
	// false. Requires a node_count greater than 1.
	EnableReadableSecondaries bool `json:"enable_readable_secondaries,omitempty"`
	// The name of the instance. This is the unique identifier for the instance.
	// Must be 1 to 63 characters long, contain only letters, digits and
	// hyphens, start with a letter, and not contain consecutive hyphens.
	Name string `json:"name"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries. This field is input
	// only, see effective_node_count for the output.
	NodeCount int `json:"node_count,omitempty"`
	// The ref of the parent instance. This is only available if the instance is
	// a child instance. A child instance inherits custom_tags and
	// usage_policy_id from its parent and cannot set them. Input: For
	// specifying the parent instance to create a child instance. Optional.
	// Output: Only populated if provided as input to create a child instance.
	ParentInstanceRef *DatabaseInstanceRef `json:"parent_instance_ref,omitempty"`
	// The version of Postgres running on the instance.
	PgVersion string `json:"pg_version,omitempty"`
	// The DNS endpoint to connect to the instance for read only access. This is
	// only available if enable_readable_secondaries is true.
	ReadOnlyDns string `json:"read_only_dns,omitempty"`
	// The DNS endpoint to connect to the instance for read+write access.
	ReadWriteDns string `json:"read_write_dns,omitempty"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained. The default value is 7 days.
	// Valid values are 2 to 35 days.
	RetentionWindowInDays int `json:"retention_window_in_days,omitempty"`
	// The current state of the instance.
	State DatabaseInstanceState `json:"state,omitempty"`
	// Whether to stop the instance. An input only param, see effective_stopped
	// for the output. An instance cannot be created in the stopped state; set
	// this field on update to stop or start the instance.
	Stopped bool `json:"stopped,omitempty"`
	// An immutable UUID identifier for the instance.
	Uid string `json:"uid,omitempty"`
	// The desired usage policy to associate with the instance.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// DatabaseInstanceRef is a reference to a database instance. It is used in the
// DatabaseInstance object to refer to the parent instance of an instance and to
// refer the child instances of an instance. To specify as a parent instance
// during creation of an instance, the lsn and branch_time fields are optional.
// If not specified, the child instance will be created from the latest lsn of
// the parent. If both lsn and branch_time are specified, the lsn will be used
// to create the child instance.
type DatabaseInstanceRef struct {
	// Branch time of the ref database instance. For a parent ref instance, this
	// is the point in time on the parent instance from which the instance was
	// created. For a child ref instance, this is the point in time on the
	// instance from which the child instance was created. Input: For specifying
	// the point in time to create a child instance. Optional. Output: Only
	// populated if provided as input to create a child instance.
	BranchTime string `json:"branch_time,omitempty"`
	// For a parent ref instance, this is the LSN on the parent instance from
	// which the instance was created. For a child ref instance, this is the LSN
	// on the instance from which the child instance was created. This is an
	// output only field that contains the value computed from the input field
	// combined with server side defaults. Use the field without the effective_
	// prefix to set the value.
	EffectiveLsn string `json:"effective_lsn,omitempty"`
	// User-specified WAL LSN of the ref database instance.
	//
	// Input: For specifying the WAL LSN to create a child instance. Optional.
	// Output: Only populated if provided as input to create a child instance.
	Lsn string `json:"lsn,omitempty"`
	// Name of the ref database instance.
	Name string `json:"name,omitempty"`
	// Id of the ref database instance.
	Uid string `json:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstanceRef) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRef) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A DatabaseInstanceRole represents a Postgres role in a database instance.
type DatabaseInstanceRole struct {
	// The desired API-exposed Postgres role attribute to associate with the
	// role. Optional.
	Attributes *DatabaseInstanceRoleAttributes `json:"attributes,omitempty"`
	// The attributes that are applied to the role. This is an output only field
	// that contains the value computed from the input field combined with
	// server side defaults. Use the field without the effective_ prefix to set
	// the value.
	EffectiveAttributes *DatabaseInstanceRoleAttributes `json:"effective_attributes,omitempty"`
	// The type of the role.
	IdentityType DatabaseInstanceRoleIdentityType `json:"identity_type,omitempty"`

	InstanceName string `json:"instance_name,omitempty"`
	// An enum value for a standard role that this role is a member of.
	MembershipRole DatabaseInstanceRoleMembershipRole `json:"membership_role,omitempty"`
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstanceRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Attributes that can be granted to a Postgres role. We are only implementing a
// subset for now, see xref:
// https://www.postgresql.org/docs/16/sql-createrole.html The values follow
// Postgres keyword naming e.g. CREATEDB, BYPASSRLS, etc. which is why they
// don't include typical underscores between words. We were requested to make
// this a nested object/struct representation since these are knobs from an
// external spec.
type DatabaseInstanceRoleAttributes struct {
	Bypassrls bool `json:"bypassrls,omitempty"`

	Createdb bool `json:"createdb,omitempty"`

	Createrole bool `json:"createrole,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseInstanceRoleAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseInstanceRoleAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabaseInstanceRoleIdentityType string

const DatabaseInstanceRoleIdentityTypeGroup DatabaseInstanceRoleIdentityType = `GROUP`

const DatabaseInstanceRoleIdentityTypePgOnly DatabaseInstanceRoleIdentityType = `PG_ONLY`

const DatabaseInstanceRoleIdentityTypeServicePrincipal DatabaseInstanceRoleIdentityType = `SERVICE_PRINCIPAL`

const DatabaseInstanceRoleIdentityTypeUser DatabaseInstanceRoleIdentityType = `USER`

// String representation for [fmt.Print]
func (f *DatabaseInstanceRoleIdentityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseInstanceRoleIdentityType) Set(v string) error {
	switch v {
	case `GROUP`, `PG_ONLY`, `SERVICE_PRINCIPAL`, `USER`:
		*f = DatabaseInstanceRoleIdentityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GROUP", "PG_ONLY", "SERVICE_PRINCIPAL", "USER"`, v)
	}
}

// Values returns all possible values for DatabaseInstanceRoleIdentityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseInstanceRoleIdentityType) Values() []DatabaseInstanceRoleIdentityType {
	return []DatabaseInstanceRoleIdentityType{
		DatabaseInstanceRoleIdentityTypeGroup,
		DatabaseInstanceRoleIdentityTypePgOnly,
		DatabaseInstanceRoleIdentityTypeServicePrincipal,
		DatabaseInstanceRoleIdentityTypeUser,
	}
}

// Type always returns DatabaseInstanceRoleIdentityType to satisfy [pflag.Value] interface
func (f *DatabaseInstanceRoleIdentityType) Type() string {
	return "DatabaseInstanceRoleIdentityType"
}

// Roles that the DatabaseInstanceRole can be a member of.
type DatabaseInstanceRoleMembershipRole string

const DatabaseInstanceRoleMembershipRoleDatabricksSuperuser DatabaseInstanceRoleMembershipRole = `DATABRICKS_SUPERUSER`

// String representation for [fmt.Print]
func (f *DatabaseInstanceRoleMembershipRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseInstanceRoleMembershipRole) Set(v string) error {
	switch v {
	case `DATABRICKS_SUPERUSER`:
		*f = DatabaseInstanceRoleMembershipRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS_SUPERUSER"`, v)
	}
}

// Values returns all possible values for DatabaseInstanceRoleMembershipRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseInstanceRoleMembershipRole) Values() []DatabaseInstanceRoleMembershipRole {
	return []DatabaseInstanceRoleMembershipRole{
		DatabaseInstanceRoleMembershipRoleDatabricksSuperuser,
	}
}

// Type always returns DatabaseInstanceRoleMembershipRole to satisfy [pflag.Value] interface
func (f *DatabaseInstanceRoleMembershipRole) Type() string {
	return "DatabaseInstanceRoleMembershipRole"
}

type DatabaseInstanceState string

const DatabaseInstanceStateAvailable DatabaseInstanceState = `AVAILABLE`

const DatabaseInstanceStateDeleting DatabaseInstanceState = `DELETING`

const DatabaseInstanceStateFailingOver DatabaseInstanceState = `FAILING_OVER`

const DatabaseInstanceStateStarting DatabaseInstanceState = `STARTING`

const DatabaseInstanceStateStopped DatabaseInstanceState = `STOPPED`

const DatabaseInstanceStateUpdating DatabaseInstanceState = `UPDATING`

// String representation for [fmt.Print]
func (f *DatabaseInstanceState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatabaseInstanceState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`:
		*f = DatabaseInstanceState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DELETING", "FAILING_OVER", "STARTING", "STOPPED", "UPDATING"`, v)
	}
}

// Values returns all possible values for DatabaseInstanceState.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatabaseInstanceState) Values() []DatabaseInstanceState {
	return []DatabaseInstanceState{
		DatabaseInstanceStateAvailable,
		DatabaseInstanceStateDeleting,
		DatabaseInstanceStateFailingOver,
		DatabaseInstanceStateStarting,
		DatabaseInstanceStateStopped,
		DatabaseInstanceStateUpdating,
	}
}

// Type always returns DatabaseInstanceState to satisfy [pflag.Value] interface
func (f *DatabaseInstanceState) Type() string {
	return "DatabaseInstanceState"
}

// A database table is a table that already exists in a Postgres database of a
// Lakebase database instance and is registered in Unity Catalog, so that it is
// governed and discoverable there while the Postgres table remains the source
// of truth. Unlike a synced table, no data is copied from a Unity Catalog
// source table into Postgres.
type DatabaseTable struct {
	// The name of the database instance that hosts the Postgres table. It must
	// be set when creating a database table.
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
	// The name of the Postgres database in the instance that contains the
	// table. It must be set when creating a database table. Tables in the
	// system databases postgres, databricks_system, template0 and template1,
	// and tables in the schemas pg_catalog and information_schema, cannot be
	// registered.
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table in Unity
	// Catalog. The catalog must be a standard or default storage catalog;
	// database tables cannot be registered in a database catalog. The schema
	// and table parts must match the schema and table names of the Postgres
	// table that is registered.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabaseTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabaseTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseCatalogRequest struct {
	// The Unity Catalog name of the catalog to delete.
	Name string `json:"-" url:"-"`
}

func (s *DeleteDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDatabaseInstanceRequest struct {
	// By default, an instance cannot be deleted if it has descendant instances
	// created via PITR. Setting this flag to true is not supported and the
	// request is rejected; delete the child instances first.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the instance to delete.
	Name string `json:"-" url:"-"`
	// Deprecated: omitting the field or setting it to true hard deletes the
	// instance. Setting it to false returns a bad request.
	Purge bool `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDatabaseInstanceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseInstanceRoleRequest struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	AllowMissing bool `json:"-" url:"allow_missing,omitempty"`

	InstanceName string `json:"-" url:"-"`

	Name string `json:"-" url:"-"`

	ReassignOwnedTo string `json:"-" url:"reassign_owned_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDatabaseInstanceRoleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDatabaseTableRequest struct {
	// Specifies the database table by its full three-part (catalog, schema,
	// table) name.
	Name string `json:"-" url:"-"`
}

func (s *DeleteDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSyncedDatabaseTableRequest struct {
	// Specifies the synced table by its full three-part (catalog, schema,
	// table) name.
	Name string `json:"-" url:"-"`
	// When true, the Postgres table is also dropped from the database. When
	// false, the synced table is unregistered from Unity Catalog and its
	// pipeline stops syncing it, but the Postgres table and its data are kept.
	PurgeData bool `json:"-" url:"purge_data,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSyncedDatabaseTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeltaTableSyncInfo struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	DeltaCommitTimestamp string `json:"delta_commit_timestamp,omitempty"`
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

type FindDatabaseInstanceByUidRequest struct {
	// UID of the instance to get, as returned in the uid field of the instance.
	Uid string `json:"-" url:"uid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FindDatabaseInstanceByUidRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FindDatabaseInstanceByUidRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Generates a credential that can be used to access database instances
type GenerateDatabaseCredentialRequest struct {
	// Unity Catalog permissions to embed in the credential. Databricks verifies
	// that the caller holds each permission in Unity Catalog and records a
	// reference to it in the token. When the user connects, Postgres grants
	// them the matching privileges on the Postgres objects that back those
	// Unity Catalog resources. Each resource must have a Postgres counterpart,
	// such as a synced table or a table in a database catalog.
	Claims []RequestedClaims `json:"claims,omitempty"`
	// Names of the database instances to connect to with the credential. Every
	// named instance must exist in the workspace, or the request fails. At
	// least one of instance_names or claims must be specified.
	InstanceNames []string `json:"instance_names,omitempty"`
	// A client-generated identifier for this request, such as a UUID. The
	// request fails when it is omitted.
	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenerateDatabaseCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateDatabaseCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDatabaseCatalogRequest struct {
	// The Unity Catalog name of the catalog to retrieve.
	Name string `json:"-" url:"-"`
}

func (s *GetDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDatabaseInstanceRequest struct {
	// Name of the instance to get.
	Name string `json:"-" url:"-"`
}

func (s *GetDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDatabaseInstanceRoleRequest struct {
	InstanceName string `json:"-" url:"-"`

	Name string `json:"-" url:"-"`
}

func (s *GetDatabaseInstanceRoleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDatabaseTableRequest struct {
	// Specifies the database table by its full three-part (catalog, schema,
	// table) name.
	Name string `json:"-" url:"-"`
}

func (s *GetDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSyncedDatabaseTableRequest struct {
	// Specifies the synced table by its full three-part (catalog, schema,
	// table) name.
	Name string `json:"-" url:"-"`
}

func (s *GetSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListDatabaseCatalogsRequest struct {
	// Name of the instance to get database catalogs for.
	InstanceName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseCatalogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseCatalogsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseCatalogsResponse struct {
	DatabaseCatalogs []DatabaseCatalog `json:"database_catalogs,omitempty"`
	// Pagination token to request the next page of database catalogs.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseCatalogsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseCatalogsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseInstanceRolesRequest struct {
	InstanceName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstanceRolesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstanceRolesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseInstanceRolesResponse struct {
	// List of database instance roles.
	DatabaseInstanceRoles []DatabaseInstanceRole `json:"database_instance_roles,omitempty"`
	// Pagination token to request the next page of instances.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstanceRolesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstanceRolesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseInstancesRequest struct {
	// Upper bound for items returned. Defaults to 100, which is also the
	// maximum value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstancesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstancesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDatabaseInstancesResponse struct {
	// List of instances.
	DatabaseInstances []DatabaseInstance `json:"database_instances,omitempty"`
	// Pagination token to request the next page of instances.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDatabaseInstancesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDatabaseInstancesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSyncedDatabaseTablesRequest struct {
	// Name of the instance to get synced tables for.
	InstanceName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSyncedDatabaseTablesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSyncedDatabaseTablesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSyncedDatabaseTablesResponse struct {
	// Pagination token to request the next page of synced tables.
	NextPageToken string `json:"next_page_token,omitempty"`

	SyncedTables []SyncedDatabaseTable `json:"synced_tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSyncedDatabaseTablesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSyncedDatabaseTablesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec struct {
	// Budget policy to set on the newly created pipeline.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog string `json:"storage_catalog,omitempty"`
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
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
	// The level of access that the credential grants on the listed resources.
	// READ_ONLY covers SELECT on each table, along with USE SCHEMA on its
	// schema and USE CATALOG on its catalog.
	PermissionSet RequestedClaimsPermissionSet `json:"permission_set,omitempty"`
	// The Unity Catalog resources that the permission set applies to.
	Resources []RequestedResource `json:"resources,omitempty"`
}

func (s *RequestedClaims) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The level of access that a claim grants on its resources.
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
	// The full three-level Unity Catalog name of the database table or synced
	// table, in the form catalog.schema.table.
	TableName string `json:"table_name,omitempty"`

	UnspecifiedResourceName string `json:"unspecified_resource_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RequestedResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RequestedResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A synced table is a Postgres table in a Lakebase database instance that
// Databricks keeps in sync with a source table in Unity Catalog. A data
// synchronization pipeline loads the source data and then applies later changes
// according to the scheduling policy, so the table can serve low-latency
// queries without a hand-built ETL process. For an overview, see
// https://docs.databricks.com/aws/en/oltp/projects/sync-tables.
type SyncedDatabaseTable struct {
	// Reports the status of the data synchronization pipeline that keeps this
	// table in sync with its source, including the pipeline ID and the position
	// of the last successful sync.
	DataSynchronizationStatus *SyncedTableStatus `json:"data_synchronization_status,omitempty"`
	// Name of the target database instance. This is required when creating
	// synced database tables in standard catalogs. This is optional when
	// creating synced database tables in registered catalogs. If this field is
	// specified when creating synced database tables in registered catalogs,
	// the database instance name MUST match that of the registered catalog (or
	// the request will be rejected).
	DatabaseInstanceName string `json:"database_instance_name,omitempty"`
	// The name of the database instance that this table is registered to. This
	// field is always returned; for tables inside database catalogs it is the
	// database instance associated with the catalog. This is an output only
	// field that contains the value computed from the input field combined with
	// server side defaults. Use the field without the effective_ prefix to set
	// the value.
	EffectiveDatabaseInstanceName string `json:"effective_database_instance_name,omitempty"`
	// The name of the logical database that this table is registered to. This
	// is an output only field that contains the value computed from the input
	// field combined with server side defaults. Use the field without the
	// effective_ prefix to set the value.
	EffectiveLogicalDatabaseName string `json:"effective_logical_database_name,omitempty"`
	// Target Postgres database object (logical database) name for this table.
	//
	// When creating a synced table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a synced table in a standard catalog, this field is
	// required. In this scenario, specifying this field will allow targeting an
	// arbitrary postgres database. If this database does not exist yet, it is
	// created when `create_database_objects_if_missing` in `spec` is true.
	LogicalDatabaseName string `json:"logical_database_name,omitempty"`
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name"`
	// Specifies the source table, the key columns and how the data
	// synchronization pipeline runs.
	Spec *SyncedTableSpec `json:"spec,omitempty"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedDatabaseTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedDatabaseTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
type SyncedTableContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus struct {
	// The last source table Delta version that was successfully synced to the
	// synced table. Only populated if the table is still synced and available
	// for serving.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. Only populated if the table is still
	// synced and available for serving.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableFailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableFailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Progress information of the Synced Table data synchronization pipeline.
type SyncedTablePipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds float64 `json:"estimated_completion_time_seconds,omitempty"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing int64 `json:"latest_version_currently_processing,omitempty"`
	// The current phase of the data synchronization pipeline.
	ProvisioningPhase ProvisioningPhase `json:"provisioning_phase,omitempty"`
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
	// Identifies the position in the source Delta table at the time of the last
	// successful sync.
	DeltaTableSyncInfo *DeltaTableSyncInfo `json:"delta_table_sync_info,omitempty"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	SyncEndTimestamp string `json:"sync_end_timestamp,omitempty"`
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	SyncStartTimestamp string `json:"sync_start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTablePosition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTablePosition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress *SyncedTablePipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

func (s *SyncedTableProvisioningStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SyncedTableSchedulingPolicy string

const SyncedTableSchedulingPolicyContinuous SyncedTableSchedulingPolicy = `CONTINUOUS`

const SyncedTableSchedulingPolicySnapshot SyncedTableSchedulingPolicy = `SNAPSHOT`

const SyncedTableSchedulingPolicyTriggered SyncedTableSchedulingPolicy = `TRIGGERED`

// String representation for [fmt.Print]
func (f *SyncedTableSchedulingPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableSchedulingPolicy) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`:
		*f = SyncedTableSchedulingPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "SNAPSHOT", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for SyncedTableSchedulingPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableSchedulingPolicy) Values() []SyncedTableSchedulingPolicy {
	return []SyncedTableSchedulingPolicy{
		SyncedTableSchedulingPolicyContinuous,
		SyncedTableSchedulingPolicySnapshot,
		SyncedTableSchedulingPolicyTriggered,
	}
}

// Type always returns SyncedTableSchedulingPolicy to satisfy [pflag.Value] interface
func (f *SyncedTableSchedulingPolicy) Type() string {
	return "SyncedTableSchedulingPolicy"
}

// Specification of a synced database table.
type SyncedTableSpec struct {
	// When true, enables accelerated sync mode for the initial data load. This
	// significantly improves performance for large tables. Requires
	// workspace-level enablement.
	AcceleratedSync bool `json:"accelerated_sync,omitempty"`
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist. The request fails if this
	// is false and the database or schema does not exist.
	CreateDatabaseObjectsIfMissing bool `json:"create_database_objects_if_missing,omitempty"`
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If existing_pipeline_id is defined, the synced table will be bin packed
	// into the existing pipeline referenced. This avoids creating a new
	// pipeline and allows sharing existing compute. In this case, the
	// scheduling_policy of this synced table must match the scheduling policy
	// of the existing pipeline.
	ExistingPipelineId string `json:"existing_pipeline_id,omitempty"`
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If new_pipeline_spec is defined, a new pipeline is created for this
	// synced table. The location pointed to is used to store intermediate files
	// (checkpoints, event logs etc). The caller must have write permissions to
	// create Delta tables in the specified catalog and schema. Again, note this
	// requires write permissions, whereas the source table only requires read
	// permissions.
	NewPipelineSpec *NewPipelineSpec `json:"new_pipeline_spec,omitempty"`
	// The primary key columns are used to insert and update rows in the
	// destination table. At least one column is required. Rows whose primary
	// key columns contain null values are not synced, and the primary key
	// cannot be changed after the table is created.
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Specifies the scheduling policy of the underlying pipeline. With
	// CONTINUOUS, the pipeline keeps running and applies source changes as they
	// happen. With TRIGGERED and SNAPSHOT, the pipeline stops after the initial
	// load and syncs again only when it is run (manually, on a schedule or from
	// a job); TRIGGERED applies source changes incrementally, while SNAPSHOT
	// copies the full source table on every run. CONTINUOUS and TRIGGERED
	// require the source table to provide a change data feed.
	SchedulingPolicy SyncedTableSchedulingPolicy `json:"scheduling_policy,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// The time series key deduplicates (tie-breaks) rows with the same primary
	// key. When primary key values can repeat in the source table, the synced
	// table keeps only the row with the latest value of this column for each
	// primary key. The key cannot be changed after the table is created.
	TimeseriesKey string `json:"timeseries_key,omitempty"`
	// Override the default Delta->PG type mapping for specific columns. A
	// TypeOverride with PG_SPECIFIC_TYPE_UNSPECIFIED is rejected; a valid
	// pg_type must be set.
	TypeOverrides []SyncedTableSpecTypeOverride `json:"type_overrides,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// PostgreSQL-specific target types that can override the default Delta-to-PG
// mapping.
type SyncedTableSpecPgSpecificType string

const SyncedTableSpecPgSpecificTypePgSpecificTypeHalfvec SyncedTableSpecPgSpecificType = `PG_SPECIFIC_TYPE_HALFVEC`

const SyncedTableSpecPgSpecificTypePgSpecificTypeVarchar SyncedTableSpecPgSpecificType = `PG_SPECIFIC_TYPE_VARCHAR`

const SyncedTableSpecPgSpecificTypePgSpecificTypeVector SyncedTableSpecPgSpecificType = `PG_SPECIFIC_TYPE_VECTOR`

// String representation for [fmt.Print]
func (f *SyncedTableSpecPgSpecificType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SyncedTableSpecPgSpecificType) Set(v string) error {
	switch v {
	case `PG_SPECIFIC_TYPE_HALFVEC`, `PG_SPECIFIC_TYPE_VARCHAR`, `PG_SPECIFIC_TYPE_VECTOR`:
		*f = SyncedTableSpecPgSpecificType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PG_SPECIFIC_TYPE_HALFVEC", "PG_SPECIFIC_TYPE_VARCHAR", "PG_SPECIFIC_TYPE_VECTOR"`, v)
	}
}

// Values returns all possible values for SyncedTableSpecPgSpecificType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableSpecPgSpecificType) Values() []SyncedTableSpecPgSpecificType {
	return []SyncedTableSpecPgSpecificType{
		SyncedTableSpecPgSpecificTypePgSpecificTypeHalfvec,
		SyncedTableSpecPgSpecificTypePgSpecificTypeVarchar,
		SyncedTableSpecPgSpecificTypePgSpecificTypeVector,
	}
}

// Type always returns SyncedTableSpecPgSpecificType to satisfy [pflag.Value] interface
func (f *SyncedTableSpecPgSpecificType) Type() string {
	return "SyncedTableSpecPgSpecificType"
}

// Overrides the default Delta-to-PostgreSQL type mapping for a single column.
type SyncedTableSpecTypeOverride struct {
	// Name of the source column whose target PostgreSQL type should be
	// overridden.
	ColumnName string `json:"column_name"`
	// PostgreSQL-specific target type to use for the column.
	PgType SyncedTableSpecPgSpecificType `json:"pg_type"`
	// Size parameter for the target type, for types that take one (e.g. vector
	// dimension, varchar length). Required when the chosen pg_type needs a
	// size.
	Size int `json:"size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableSpecTypeOverride) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableSpecTypeOverride) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of a synced table.
type SyncedTableState string

const SyncedTableStateSyncedTabledOffline SyncedTableState = `SYNCED_TABLED_OFFLINE`

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
	case `SYNCED_TABLED_OFFLINE`, `SYNCED_TABLE_OFFLINE_FAILED`, `SYNCED_TABLE_ONLINE`, `SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE`, `SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE`, `SYNCED_TABLE_ONLINE_PIPELINE_FAILED`, `SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE`, `SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES`, `SYNCED_TABLE_PROVISIONING`, `SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT`, `SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES`:
		*f = SyncedTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SYNCED_TABLED_OFFLINE", "SYNCED_TABLE_OFFLINE_FAILED", "SYNCED_TABLE_ONLINE", "SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE", "SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE", "SYNCED_TABLE_ONLINE_PIPELINE_FAILED", "SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE", "SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES", "SYNCED_TABLE_PROVISIONING", "SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT", "SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Values returns all possible values for SyncedTableState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SyncedTableState) Values() []SyncedTableState {
	return []SyncedTableState{
		SyncedTableStateSyncedTabledOffline,
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

// Status of a synced table.
type SyncedTableStatus struct {
	// Provides detailed status for an online table with the CONTINUOUS
	// scheduling policy. It is populated when detailed_state is
	// SYNCED_TABLE_ONLINE_CONTINUOUS_UPDATE or
	// SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES.
	ContinuousUpdateStatus *SyncedTableContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	// The state of the synced table.
	DetailedState SyncedTableState `json:"detailed_state,omitempty"`
	// Provides detailed status for a table whose data synchronization pipeline
	// has failed. It is populated when detailed_state is
	// SYNCED_TABLE_OFFLINE_FAILED or SYNCED_TABLE_ONLINE_PIPELINE_FAILED.
	FailedStatus *SyncedTableFailedStatus `json:"failed_status,omitempty"`
	// Summary of the last successful synchronization from source to
	// destination.
	//
	// Will always be present if there has been a successful sync. Even if the
	// most recent syncs have failed.
	//
	// Limitation: The only exception is if the synced table is doing a FULL
	// REFRESH, then the last sync information will not be available until the
	// full refresh is complete. This limitation will be addressed in a future
	// version.
	//
	// This top-level field is a convenience for consumers who want easy access
	// to last sync information without having to traverse detailed_status.
	LastSync *SyncedTablePosition `json:"last_sync,omitempty"`
	// A text description of the current state of the synced table.
	Message string `json:"message,omitempty"`
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	PipelineId string `json:"pipeline_id,omitempty"`
	// Provides detailed status while the table is being provisioned. It is
	// populated when detailed_state is
	// SYNCED_TABLE_PROVISIONING_PIPELINE_RESOURCES or
	// SYNCED_TABLE_PROVISIONING_INITIAL_SNAPSHOT.
	ProvisioningStatus *SyncedTableProvisioningStatus `json:"provisioning_status,omitempty"`
	// Provides detailed status for an online table with the TRIGGERED or
	// SNAPSHOT scheduling policy. It is populated when detailed_state is
	// SYNCED_TABLE_ONLINE_TRIGGERED_UPDATE,
	// SYNCED_TABLE_ONLINE_NO_PENDING_UPDATE or
	// SYNCED_TABLE_ONLINE_UPDATING_PIPELINE_RESOURCES.
	TriggeredUpdateStatus *SyncedTableTriggeredUpdateStatus `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus struct {
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	Timestamp string `json:"timestamp,omitempty"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress *SyncedTablePipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SyncedTableTriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SyncedTableTriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateDatabaseCatalogRequest struct {
	// Note that updating a database catalog is not yet supported.
	DatabaseCatalog DatabaseCatalog `json:"database_catalog"`
	// The name of the catalog in Unity Catalog.
	Name string `json:"-" url:"-"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateDatabaseCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateDatabaseInstanceRequest struct {
	// Specifies the instance to update, identified by its name.
	DatabaseInstance DatabaseInstance `json:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	// Must be 1 to 63 characters long, contain only letters, digits and
	// hyphens, start with a letter, and not contain consecutive hyphens.
	Name string `json:"-" url:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible. Use "*" to update every field. To wipe out custom_tags,
	// specify custom_tags in the update_mask with an empty custom_tags list.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateDatabaseInstanceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSyncedDatabaseTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
	// Note that updating a synced database table is not yet supported.
	SyncedTable SyncedDatabaseTable `json:"synced_table"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateSyncedDatabaseTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
