// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jsonmarshallv2

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type GetResourceRequest struct {
	Name string `json:"-" url:"-"`
	// Description.
	Resource Resource `json:"-" url:"resource"`
}

func (s *GetResourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type NestedMessage struct {
	OptionalDuration *duration.Duration `json:"optional_duration,omitempty" url:"optional_duration,omitempty"`

	OptionalString string `json:"optional_string,omitempty" url:"optional_string,omitempty"`

	OptionalTimestamp *time.Time `json:"optional_timestamp,omitempty" url:"optional_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NestedMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NestedMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type OptionalFields struct {
	Duration *duration.Duration `json:"duration,omitempty" url:"duration,omitempty"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	FieldMask *fieldmask.FieldMask `json:"field_mask,omitempty" url:"field_mask,omitempty"`
	// Legacy Well Known types
	LegacyDuration string `json:"legacy_duration,omitempty" url:"legacy_duration,omitempty"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	LegacyFieldMask string `json:"legacy_field_mask,omitempty" url:"legacy_field_mask,omitempty"`

	LegacyTimestamp string `json:"legacy_timestamp,omitempty" url:"legacy_timestamp,omitempty"`

	ListValue []json.RawMessage `json:"list_value,omitempty" url:"list_value,omitempty"`
	// Lint disable reason: This is a dummy field used to test SDK Generation
	// logic.
	Map map[string]string `json:"map,omitempty" url:"map,omitempty"`

	OptionalBool bool `json:"optional_bool,omitempty" url:"optional_bool,omitempty"`

	OptionalInt32 int `json:"optional_int32,omitempty" url:"optional_int32,omitempty"`

	OptionalInt64 int64 `json:"optional_int64,omitempty" url:"optional_int64,omitempty"`

	OptionalMessage *NestedMessage `json:"optional_message,omitempty" url:"optional_message,omitempty"`

	OptionalString string `json:"optional_string,omitempty" url:"optional_string,omitempty"`

	Struct map[string]json.RawMessage `json:"struct,omitempty" url:"struct,omitempty"`

	TestEnum TestEnum `json:"test_enum,omitempty" url:"test_enum,omitempty"`

	Timestamp *time.Time `json:"timestamp,omitempty" url:"timestamp,omitempty"`

	Value *json.RawMessage `json:"value,omitempty" url:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OptionalFields) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OptionalFields) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepeatedFields struct {
	RepeatedBool []bool `json:"repeated_bool,omitempty" url:"repeated_bool,omitempty"`

	RepeatedDuration []duration.Duration `json:"repeated_duration,omitempty" url:"repeated_duration,omitempty"`

	RepeatedFieldMask []fieldmask.FieldMask `json:"repeated_field_mask,omitempty" url:"repeated_field_mask,omitempty"`

	RepeatedInt32 []int `json:"repeated_int32,omitempty" url:"repeated_int32,omitempty"`

	RepeatedInt64 []int64 `json:"repeated_int64,omitempty" url:"repeated_int64,omitempty"`

	RepeatedListValue [][]json.RawMessage `json:"repeated_list_value,omitempty" url:"repeated_list_value,omitempty"`

	RepeatedMessage []NestedMessage `json:"repeated_message,omitempty" url:"repeated_message,omitempty"`

	RepeatedString []string `json:"repeated_string,omitempty" url:"repeated_string,omitempty"`

	RepeatedStruct []map[string]json.RawMessage `json:"repeated_struct,omitempty" url:"repeated_struct,omitempty"`

	RepeatedTimestamp []time.Time `json:"repeated_timestamp,omitempty" url:"repeated_timestamp,omitempty"`

	RepeatedValue []json.RawMessage `json:"repeated_value,omitempty" url:"repeated_value,omitempty"`

	TestRepeatedEnum []TestEnum `json:"test_repeated_enum,omitempty" url:"test_repeated_enum,omitempty"`
}

func (s *RepeatedFields) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RequiredFields struct {
	RequiredBool bool `json:"required_bool" url:"required_bool"`

	RequiredDuration duration.Duration `json:"required_duration" url:"required_duration"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	RequiredFieldMask fieldmask.FieldMask `json:"required_field_mask" url:"required_field_mask"`

	RequiredInt32 int `json:"required_int32" url:"required_int32"`

	RequiredInt64 int64 `json:"required_int64" url:"required_int64"`

	RequiredListValue []json.RawMessage `json:"required_list_value" url:"required_list_value"`

	RequiredMessage NestedMessage `json:"required_message" url:"required_message"`

	RequiredString string `json:"required_string" url:"required_string"`

	RequiredStruct map[string]json.RawMessage `json:"required_struct" url:"required_struct"`

	RequiredTimestamp time.Time `json:"required_timestamp" url:"required_timestamp"`

	RequiredValue json.RawMessage `json:"required_value" url:"required_value"`

	TestRequiredEnum TestEnum `json:"test_required_enum" url:"test_required_enum"`
}

func (s *RequiredFields) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// We separate this into 3 submessages to simplify test cases. E.g., any
// required top level field needs to be included in the expected json for each
// test case.
type Resource struct {
	OptionalFields *OptionalFields `json:"optional_fields,omitempty" url:"optional_fields,omitempty"`

	RepeatedFields *RepeatedFields `json:"repeated_fields,omitempty" url:"repeated_fields,omitempty"`

	RequiredFields *RequiredFields `json:"required_fields,omitempty" url:"required_fields,omitempty"`
}

func (s *Resource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TestEnum string

const TestEnumTestEnumOne TestEnum = `TEST_ENUM_ONE`

const TestEnumTestEnumTwo TestEnum = `TEST_ENUM_TWO`

// String representation for [fmt.Print]
func (f *TestEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TestEnum) Set(v string) error {
	switch v {
	case `TEST_ENUM_ONE`, `TEST_ENUM_TWO`:
		*f = TestEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TEST_ENUM_ONE", "TEST_ENUM_TWO"`, v)
	}
}

// Values returns all possible values for TestEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *TestEnum) Values() []TestEnum {
	return []TestEnum{
		TestEnumTestEnumOne,
		TestEnumTestEnumTwo,
	}
}

// Type always returns TestEnum to satisfy [pflag.Value] interface
func (f *TestEnum) Type() string {
	return "TestEnum"
}
