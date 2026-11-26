// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tags

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateTagAssignmentRequest struct {
	TagAssignment TagAssignment `json:"tag_assignment"`
}

func (s *CreateTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateTagPolicyRequest struct {
	TagPolicy TagPolicy `json:"tag_policy"`
}

func (s *CreateTagPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned. For apps, the
	// entity_id is the app name
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed values are apps,
	// dashboards, geniespaces, notebooks
	EntityType string `json:"-" url:"-"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"-" url:"-"`
}

func (s *DeleteTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteTagPolicyRequest struct {
	// The key of the governed tag whose policy to delete.
	TagKey string `json:"-" url:"-"`
}

func (s *DeleteTagPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned. For apps, the
	// entity_id is the app name
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed values are apps,
	// dashboards, geniespaces, notebooks
	EntityType string `json:"-" url:"-"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"-" url:"-"`
}

func (s *GetTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetTagPolicyRequest struct {
	// The key of the governed tag whose policy to get.
	TagKey string `json:"-" url:"-"`
}

func (s *GetTagPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListTagAssignmentsRequest struct {
	// The identifier of the entity to which the tag is assigned. For apps, the
	// entity_id is the app name
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed values are apps,
	// dashboards, geniespaces, notebooks
	EntityType string `json:"-" url:"-"`
	// Optional. Maximum number of tag assignments to return in a single page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of tag assignments. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTagAssignmentsResponse struct {
	// Pagination token to request the next page of tag assignments
	NextPageToken string `json:"next_page_token,omitempty"`

	TagAssignments []TagAssignment `json:"tag_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTagPoliciesRequest struct {
	// The maximum number of results to return in this request. Fewer results
	// may be returned than requested. If unspecified or set to 0, this defaults
	// to 1000. The maximum value is 1000; values above 1000 will be coerced
	// down to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// An optional page token received from a previous list tag policies call.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTagPoliciesResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	TagPolicies []TagPolicy `json:"tag_policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A tag assignment attaches a tag, which is a key with an optional value, to a
// workspace object such as a dashboard, notebook, Genie space or app. An object
// can have at most one assignment for each tag key. If the key belongs to a
// governed tag, the caller needs the ASSIGN permission on that governed tag,
// and the value must be one of its allowed values. See
// https://docs.databricks.com/aws/en/admin/governed-tags/.
type TagAssignment struct {
	// The identifier of the entity to which the tag is assigned. For apps, the
	// entity_id is the app name
	EntityId string `json:"entity_id"`
	// The type of entity to which the tag is assigned. Allowed values are apps,
	// dashboards, geniespaces, notebooks
	EntityType string `json:"entity_type"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"tag_key"`
	// The value of the tag. If the value is omitted, the tag is assigned with
	// its key only.
	TagValue string `json:"tag_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TagAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TagAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A tag policy makes a tag key governed and defines the values that can be
// assigned to it. Governed tags are enforced across all workspaces in the
// account, and only users with the ASSIGN permission on the governed tag can
// assign it. See https://docs.databricks.com/aws/en/admin/governed-tags/.
type TagPolicy struct {
	// Timestamp when the tag policy was created
	CreateTime string `json:"create_time,omitempty"`
	// A description of the governed tag, such as its purpose or the meaning of
	// its values.
	Description string `json:"description,omitempty"`
	// The unique identifier of the tag policy, generated when the policy is
	// created.
	Id string `json:"id,omitempty"`
	// The tag key that this policy governs. Tag keys are case-sensitive.
	TagKey string `json:"tag_key"`
	// Timestamp when the tag policy was last updated
	UpdateTime string `json:"update_time,omitempty"`
	// The allowed values of the governed tag. Only these values can be assigned
	// to the tag key, and each value must be unique. A policy without values
	// defines a key-only governed tag.
	Values []Value `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TagPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TagPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned. For apps, the
	// entity_id is the app name
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed values are apps,
	// dashboards, geniespaces, notebooks
	EntityType string `json:"-" url:"-"`

	TagAssignment TagAssignment `json:"tag_assignment"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"-" url:"-"`
	// The comma-separated list of fields to update. Only `tag_value` can be
	// updated, because the entity type, entity ID and tag key identify the tag
	// assignment.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateTagAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateTagPolicyRequest struct {
	// The tag key that this policy governs. Tag keys are case-sensitive.
	TagKey string `json:"-" url:"-"`

	TagPolicy TagPolicy `json:"tag_policy"`
	// The comma-separated list of tag policy fields to update. When the mask
	// includes `values`, the values in the request replace all existing allowed
	// values.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateTagPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An allowed value of a governed tag.
type Value struct {
	// The value that can be assigned to the tag key. Values are case-sensitive.
	Name string `json:"name"`
}

func (s *Value) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
