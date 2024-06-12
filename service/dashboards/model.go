// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDashboardRequest struct {
	// The display name of the dashboard.
	DisplayName string `tfsdk:"display_name"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash.
	ParentPath string `tfsdk:"parent_path"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard string `tfsdk:"serialized_dashboard"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `tfsdk:"warehouse_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime string `tfsdk:"create_time"`
	// UUID identifying the dashboard.
	DashboardId string `tfsdk:"dashboard_id"`
	// The display name of the dashboard.
	DisplayName string `tfsdk:"display_name"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read.
	Etag string `tfsdk:"etag"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState LifecycleState `tfsdk:"lifecycle_state"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash.
	ParentPath string `tfsdk:"parent_path"`
	// The workspace path of the dashboard asset, including the file name.
	Path string `tfsdk:"path"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard string `tfsdk:"serialized_dashboard"`
	// The timestamp of when the dashboard was last updated by the user.
	UpdateTime string `tfsdk:"update_time"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `tfsdk:"warehouse_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Dashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `tfsdk:"-" url:"-"`
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string `tfsdk:"-" url:"-"`
}

type LifecycleState string

const LifecycleStateActive LifecycleState = `ACTIVE`

const LifecycleStateTrashed LifecycleState = `TRASHED`

// String representation for [fmt.Print]
func (f *LifecycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LifecycleState) Set(v string) error {
	switch v {
	case `ACTIVE`, `TRASHED`:
		*f = LifecycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "TRASHED"`, v)
	}
}

// Type always returns LifecycleState to satisfy [pflag.Value] interface
func (f *LifecycleState) Type() string {
	return "LifecycleState"
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	DisplayName string `tfsdk:"display_name"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath string `tfsdk:"parent_path"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId string `tfsdk:"source_dashboard_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MigrateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string `tfsdk:"-" url:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials bool `tfsdk:"embed_credentials"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId string `tfsdk:"warehouse_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PublishRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName string `tfsdk:"display_name"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials bool `tfsdk:"embed_credentials"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime string `tfsdk:"revision_create_time"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId string `tfsdk:"warehouse_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PublishedDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishedDashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `tfsdk:"-" url:"-"`
}

type TrashDashboardResponse struct {
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string `tfsdk:"-" url:"-"`
}

type UnpublishDashboardResponse struct {
}

type UpdateDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `tfsdk:"-" url:"-"`
	// The display name of the dashboard.
	DisplayName string `tfsdk:"display_name"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read.
	Etag string `tfsdk:"etag"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard string `tfsdk:"serialized_dashboard"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `tfsdk:"warehouse_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
