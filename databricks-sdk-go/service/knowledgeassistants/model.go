// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledgeassistants

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateExampleRequest struct {
	// The example to create under the parent Knowledge Assistant.
	Example Example `json:"example"`
	// Parent resource where this example will be created. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent string `json:"-" url:"-"`
}

func (s *CreateExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateKnowledgeAssistantRequest struct {
	// The Knowledge Assistant to create.
	KnowledgeAssistant KnowledgeAssistant `json:"knowledge_assistant"`
}

func (s *CreateKnowledgeAssistantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateKnowledgeSourceRequest struct {
	KnowledgeSource KnowledgeSource `json:"knowledge_source"`
	// Parent resource where this source will be created. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent string `json:"-" url:"-"`
}

func (s *CreateKnowledgeSourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteExampleRequest struct {
	// The resource name of the example to delete. Format:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteKnowledgeAssistantRequest struct {
	// The resource name of the knowledge assistant to be deleted. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteKnowledgeAssistantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteKnowledgeSourceRequest struct {
	// The resource name of the Knowledge Source to delete. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteKnowledgeSourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An example associated with a Knowledge Assistant. Contains a question and
// guidelines for how the assistant should respond.
type Example struct {
	// Timestamp when this example was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The universally unique identifier (UUID) of the example.
	ExampleId string `json:"example_id,omitempty"`
	// Guidelines for answering the question. Optional — examples may be
	// created with just a question; the front-end form does not require
	// guidelines.
	Guidelines []string `json:"guidelines,omitempty"`
	// Full resource name:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name string `json:"name,omitempty"`
	// The example question.
	Question string `json:"question"`
	// Timestamp when this example was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Example) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Example) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// FileTableSpec specifies a file table source configuration.
type FileTableSpec struct {
	// The name of the column containing BINARY file content to be indexed.
	FileCol string `json:"file_col"`
	// Full UC name of the table, in the format of
	// {CATALOG}.{SCHEMA}.{TABLE_NAME}.
	TableName string `json:"table_name"`
}

func (s *FileTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// FilesSpec specifies a files source configuration.
type FilesSpec struct {
	// A UC volume path that includes a list of files.
	Path string `json:"path"`
}

func (s *FilesSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExampleRequest struct {
	// The resource name of the example. Format:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
}

func (s *GetExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetKnowledgeAssistantPermissionLevelsRequest struct {
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId string `json:"-" url:"-"`
}

func (s *GetKnowledgeAssistantPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetKnowledgeAssistantPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []KnowledgeAssistantPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetKnowledgeAssistantPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetKnowledgeAssistantPermissionsRequest struct {
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId string `json:"-" url:"-"`
}

func (s *GetKnowledgeAssistantPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetKnowledgeAssistantRequest struct {
	// The resource name of the knowledge assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
}

func (s *GetKnowledgeAssistantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetKnowledgeSourceRequest struct {
	// The resource name of the Knowledge Source. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"-" url:"-"`
}

func (s *GetKnowledgeSourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// IndexSpec specifies a vector search index source configuration.
type IndexSpec struct {
	// The column that specifies a link or reference to where the information
	// came from.
	DocUriCol string `json:"doc_uri_col"`
	// Full UC name of the vector search index, in the format of
	// {CATALOG}.{SCHEMA}.{INDEX_NAME}.
	IndexName string `json:"index_name"`
	// The column that includes the document text for retrieval.
	TextCol string `json:"text_col"`
}

func (s *IndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Entity message that represents a knowledge assistant. Note: REQUIRED
// annotations below represent create-time requirements. For updates, required
// fields are determined by the update mask.
type KnowledgeAssistant struct {
	// Creation timestamp.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The creator of the Knowledge Assistant.
	Creator string `json:"creator,omitempty"`
	// Description of what this agent can do (user-facing). Required when
	// creating a Knowledge Assistant. When updating a Knowledge Assistant,
	// optional unless included in update_mask.
	Description string `json:"description"`
	// The display name of the Knowledge Assistant, unique at workspace level.
	// Required when creating a Knowledge Assistant. When updating a Knowledge
	// Assistant, optional unless included in update_mask.
	DisplayName string `json:"display_name"`
	// The name of the knowledge assistant agent endpoint.
	EndpointName string `json:"endpoint_name,omitempty"`
	// Error details when the Knowledge Assistant is in FAILED state.
	ErrorInfo string `json:"error_info,omitempty"`
	// The MLflow experiment ID.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Deprecated: use knowledge_assistant_id instead.
	Id string `json:"id,omitempty"`
	// Additional global instructions on how the agent should generate answers.
	// Optional on create and update. When updating a Knowledge Assistant,
	// include this field in update_mask to modify it.
	Instructions string `json:"instructions,omitempty"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"name,omitempty"`
	// State of the Knowledge Assistant. Not returned in List responses.
	State KnowledgeAssistantState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistant) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistant) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeAssistantAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel KnowledgeAssistantPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistantAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistantAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeAssistantAccessControlResponse struct {
	// All permissions.
	AllPermissions []KnowledgeAssistantPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistantAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistantAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeAssistantPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel KnowledgeAssistantPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistantPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistantPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type KnowledgeAssistantPermissionLevel string

const KnowledgeAssistantPermissionLevelCanManage KnowledgeAssistantPermissionLevel = `CAN_MANAGE`

const KnowledgeAssistantPermissionLevelCanQuery KnowledgeAssistantPermissionLevel = `CAN_QUERY`

// String representation for [fmt.Print]
func (f *KnowledgeAssistantPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KnowledgeAssistantPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`:
		*f = KnowledgeAssistantPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY"`, v)
	}
}

// Values returns all possible values for KnowledgeAssistantPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *KnowledgeAssistantPermissionLevel) Values() []KnowledgeAssistantPermissionLevel {
	return []KnowledgeAssistantPermissionLevel{
		KnowledgeAssistantPermissionLevelCanManage,
		KnowledgeAssistantPermissionLevelCanQuery,
	}
}

// Type always returns KnowledgeAssistantPermissionLevel to satisfy [pflag.Value] interface
func (f *KnowledgeAssistantPermissionLevel) Type() string {
	return "KnowledgeAssistantPermissionLevel"
}

type KnowledgeAssistantPermissions struct {
	AccessControlList []KnowledgeAssistantAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistantPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistantPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeAssistantPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel KnowledgeAssistantPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistantPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistantPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeAssistantPermissionsRequest struct {
	AccessControlList []KnowledgeAssistantAccessControlRequest `json:"access_control_list,omitempty"`
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId string `json:"-" url:"-"`
}

func (s *KnowledgeAssistantPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type KnowledgeAssistantState string

const KnowledgeAssistantStateActive KnowledgeAssistantState = `ACTIVE`

const KnowledgeAssistantStateCreating KnowledgeAssistantState = `CREATING`

const KnowledgeAssistantStateFailed KnowledgeAssistantState = `FAILED`

// String representation for [fmt.Print]
func (f *KnowledgeAssistantState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KnowledgeAssistantState) Set(v string) error {
	switch v {
	case `ACTIVE`, `CREATING`, `FAILED`:
		*f = KnowledgeAssistantState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "CREATING", "FAILED"`, v)
	}
}

// Values returns all possible values for KnowledgeAssistantState.
//
// There is no guarantee on the order of the values in the slice.
func (f *KnowledgeAssistantState) Values() []KnowledgeAssistantState {
	return []KnowledgeAssistantState{
		KnowledgeAssistantStateActive,
		KnowledgeAssistantStateCreating,
		KnowledgeAssistantStateFailed,
	}
}

// Type always returns KnowledgeAssistantState to satisfy [pflag.Value] interface
func (f *KnowledgeAssistantState) Type() string {
	return "KnowledgeAssistantState"
}

// KnowledgeSource represents a source of knowledge for the KnowledgeAssistant.
// Used in create/update requests and returned in Get/List responses. Note:
// REQUIRED annotations below represent create-time requirements. For updates,
// required fields are determined by the update mask.
type KnowledgeSource struct {
	// Timestamp when this knowledge source was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Description of the knowledge source. Required when creating a Knowledge
	// Source. When updating a Knowledge Source, optional unless included in
	// update_mask.
	Description string `json:"description"`
	// Human-readable display name of the knowledge source. Required when
	// creating a Knowledge Source. When updating a Knowledge Source, optional
	// unless included in update_mask.
	DisplayName string `json:"display_name"`

	FileTable *FileTableSpec `json:"file_table,omitempty"`

	Files *FilesSpec `json:"files,omitempty"`

	Id string `json:"id,omitempty"`

	Index *IndexSpec `json:"index,omitempty"`
	// Timestamp representing the cutoff before which content in this knowledge
	// source is being ingested.
	KnowledgeCutoffTime *time.Time `json:"knowledge_cutoff_time,omitempty"`
	// Full resource name:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"name,omitempty"`
	// The type of the source: "index", "files", or "file_table". Required when
	// creating a Knowledge Source. When updating a Knowledge Source, this field
	// is ignored.
	SourceType string `json:"source_type"`

	State KnowledgeSourceState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeSourceState string

const KnowledgeSourceStateFailedUpdate KnowledgeSourceState = `FAILED_UPDATE`

const KnowledgeSourceStateUpdated KnowledgeSourceState = `UPDATED`

const KnowledgeSourceStateUpdating KnowledgeSourceState = `UPDATING`

// String representation for [fmt.Print]
func (f *KnowledgeSourceState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KnowledgeSourceState) Set(v string) error {
	switch v {
	case `FAILED_UPDATE`, `UPDATED`, `UPDATING`:
		*f = KnowledgeSourceState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_UPDATE", "UPDATED", "UPDATING"`, v)
	}
}

// Values returns all possible values for KnowledgeSourceState.
//
// There is no guarantee on the order of the values in the slice.
func (f *KnowledgeSourceState) Values() []KnowledgeSourceState {
	return []KnowledgeSourceState{
		KnowledgeSourceStateFailedUpdate,
		KnowledgeSourceStateUpdated,
		KnowledgeSourceStateUpdating,
	}
}

// Type always returns KnowledgeSourceState to satisfy [pflag.Value] interface
func (f *KnowledgeSourceState) Type() string {
	return "KnowledgeSourceState"
}

type ListExamplesRequest struct {
	// The maximum number of examples to return. If unspecified, at most 100
	// examples will be returned. The maximum value is 100; values above 100
	// will be coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListExamples` call. Provide this
	// to retrieve the subsequent page. If unspecified, the first page will be
	// returned.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent resource to list from. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExamplesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExamplesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Knowledge Assistant examples.
type ListExamplesResponse struct {
	Examples []Example `json:"examples,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExamplesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExamplesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListKnowledgeAssistantsRequest struct {
	// The maximum number of knowledge assistants to return. If unspecified, at
	// most 100 knowledge assistants will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListKnowledgeAssistants` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeAssistantsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeAssistantsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Knowledge Assistants.
type ListKnowledgeAssistantsResponse struct {
	KnowledgeAssistants []KnowledgeAssistant `json:"knowledge_assistants,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeAssistantsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeAssistantsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListKnowledgeSourcesRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent resource to list from. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeSourcesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeSourcesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListKnowledgeSourcesResponse struct {
	KnowledgeSources []KnowledgeSource `json:"knowledge_sources,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeSourcesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeSourcesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SyncKnowledgeSourcesRequest struct {
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
}

func (s *SyncKnowledgeSourcesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExampleRequest struct {
	Example Example `json:"example"`
	// The resource name of the example to update. Format:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
	// Comma-delimited list of fields to update on the example. Allowed values:
	// `question`, `guidelines`. Examples: - `question` - `question,guidelines`
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateKnowledgeAssistantRequest struct {
	// The Knowledge Assistant update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Assistant fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeAssistant KnowledgeAssistant `json:"knowledge_assistant"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
	// Comma-delimited list of fields to update on the Knowledge Assistant.
	// Allowed values: `display_name`, `description`, `instructions`. Examples:
	// - `display_name` - `description,instructions`
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateKnowledgeAssistantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateKnowledgeSourceRequest struct {
	// The Knowledge Source update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Source fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeSource KnowledgeSource `json:"knowledge_source"`
	// The resource name of the Knowledge Source to update. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"-" url:"-"`
	// Comma-delimited list of fields to update on the Knowledge Source. Allowed
	// values: `display_name`, `description`. Examples: - `display_name` -
	// `display_name,description`
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateKnowledgeSourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
