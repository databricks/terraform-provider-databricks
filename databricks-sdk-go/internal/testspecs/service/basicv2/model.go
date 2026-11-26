// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package basicv2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type Basic struct {
	// Name of the Resource. Must contain only: - alphanumeric characters -
	// underscores - hyphens
	//
	// Note: The name must be unique within the scope of the resource.
	Name string `json:"name"`
}

type CreateBasicRequest struct {
	// Basic information about the resource
	Basic Basic `json:"basic"`
}

type DeleteBasicRequest struct {
	// The name of the resource
	Name string `json:"-" url:"-"`
}

type GetBasicRequest struct {
	// The name of the resource
	Name string `json:"-" url:"-"`
}

type ListBasicsRequest struct {
	// Page size for pagination
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBasicsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBasicsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBasicsResponse struct {
	// List of resources
	Basics []Basic `json:"basics,omitempty"`
	// The next page token for pagination
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBasicsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBasicsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateBasicRequest struct {
	// Basic information about the resource
	Basic Basic `json:"basic"`
	// The name of the resource
	Name string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	UpdateMask string `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateBasicRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateBasicRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
