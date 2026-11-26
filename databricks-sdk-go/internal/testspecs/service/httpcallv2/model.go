// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package httpcallv2

import (
	"encoding/json"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type ComplexQueryParam struct {
	NestedOptionalQueryParam string `json:"nested_optional_query_param,omitempty" url:"nested_optional_query_param,omitempty"`

	NestedRepeatedQueryParam []string `json:"nested_repeated_query_param,omitempty" url:"nested_repeated_query_param,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComplexQueryParam) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplexQueryParam) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This mimics "old" style post requests which have the resource inlined.
type CreateResourceRequest struct {
	// Body element
	BodyField string `json:"body_field,omitempty"`

	PathParamBool bool `json:"-" url:"-"`

	PathParamInt int `json:"-" url:"-"`

	PathParamString string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateResourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateResourceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetResourceRequest struct {
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	FieldMask *fieldmask.FieldMask `json:"-" url:"field_mask,omitempty"`

	OptionalComplexQueryParam *ComplexQueryParam `json:"-" url:"optional_complex_query_param,omitempty"`

	PathParamBool bool `json:"-" url:"-"`

	PathParamInt int `json:"-" url:"-"`

	PathParamString string `json:"-" url:"-"`

	QueryParamBool bool `json:"-" url:"query_param_bool,omitempty"`

	QueryParamInt int `json:"-" url:"query_param_int,omitempty"`

	QueryParamString string `json:"-" url:"query_param_string,omitempty"`

	RepeatedComplexQueryParam []ComplexQueryParam `json:"-" url:"repeated_complex_query_param,omitempty"`

	RepeatedQueryParam []string `json:"-" url:"repeated_query_param,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetResourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetResourceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Resource struct {
	AnyField json.RawMessage `json:"any_field,omitempty"`

	BodyField string `json:"body_field,omitempty"`

	NestedPathParamBool bool `json:"nested_path_param_bool,omitempty"`

	NestedPathParamInt int `json:"nested_path_param_int,omitempty"`

	NestedPathParamString string `json:"nested_path_param_string,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Resource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Resource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This mimics a parameterless custom method's request: all fields are
// path-bound, so the serialized JSON body is the empty object {}.
type SyncResourceRequest struct {
	// Boolean path segment identifying the resource to sync.
	PathParamBool bool `json:"-" url:"-"`
	// Integer path segment identifying the resource to sync.
	PathParamInt int `json:"-" url:"-"`
	// String path segment identifying the resource to sync.
	PathParamString string `json:"-" url:"-"`
}

func (s *SyncResourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateResourceRequest struct {
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	FieldMask *fieldmask.FieldMask `json:"-" url:"field_mask,omitempty"`

	NestedPathParamBool bool `json:"-" url:"-"`

	NestedPathParamInt int `json:"-" url:"-"`

	NestedPathParamString string `json:"-" url:"-"`

	OptionalComplexQueryParam *ComplexQueryParam `json:"-" url:"optional_complex_query_param,omitempty"`

	QueryParamBool bool `json:"-" url:"query_param_bool,omitempty"`

	QueryParamInt int `json:"-" url:"query_param_int,omitempty"`

	QueryParamString string `json:"-" url:"query_param_string,omitempty"`

	RepeatedComplexQueryParam []ComplexQueryParam `json:"-" url:"repeated_complex_query_param,omitempty"`

	RepeatedQueryParam []string `json:"-" url:"repeated_query_param,omitempty"`
	// Body element
	Resource Resource `json:"resource"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateResourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateResourceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
