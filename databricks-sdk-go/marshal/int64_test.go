package marshal

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"
)

func TestUnmarshal_int64NumberOrString(t *testing.T) {
	type nested struct {
		ID int64 `json:"id"`
	}
	type response struct {
		ID     int64            `json:"id"`
		IDs    []int64          `json:"ids"`
		Nested nested           `json:"nested"`
		ByName map[string]int64 `json:"by_name"`
		Name   string           `json:"name"`
	}

	input := []byte(`{
		"id": "9223372036854775807",
		"ids": [1, "-9223372036854775808"],
		"nested": {"id": "3"},
		"by_name": {"four": "4"},
		"name": "5"
	}`)
	want := response{
		ID:     math.MaxInt64,
		IDs:    []int64{1, math.MinInt64},
		Nested: nested{ID: 3},
		ByName: map[string]int64{"four": 4},
		Name:   "5",
	}

	var got response
	if err := Unmarshal(input, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unmarshal() = %#v, want %#v", got, want)
	}
}

func TestUnmarshal_invalidInt64String(t *testing.T) {
	type response struct {
		ID int64 `json:"id"`
	}

	err := Unmarshal([]byte(`{"id":"not-an-int"}`), &response{})
	if err == nil {
		t.Fatal("Unmarshal() error = nil, want non-nil")
	}
}

func TestUnmarshal_int64NullPreservesValue(t *testing.T) {
	type response struct {
		ID int64 `json:"id"`
	}

	got := response{ID: 7}
	if err := Unmarshal([]byte(`{"id":null}`), &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.ID != 7 {
		t.Errorf("Unmarshal() ID = %d, want 7", got.ID)
	}
}

func TestUnmarshal_int64StringPreservesRawMessage(t *testing.T) {
	type response struct {
		Created int64             `json:"created"`
		Outputs []json.RawMessage `json:"outputs"`
	}

	input := []byte(`{
		"created": "7",
		"outputs": [{ "escaped": "<value>", "numbers" : [1, 2] }]
	}`)
	var baseline struct {
		Outputs []json.RawMessage `json:"outputs"`
	}
	if err := json.Unmarshal(input, &baseline); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	var got response
	if err := Unmarshal(input, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.Created != 7 {
		t.Errorf("Unmarshal() Created = %d, want 7", got.Created)
	}
	if !reflect.DeepEqual(got.Outputs, baseline.Outputs) {
		t.Errorf("Unmarshal() Outputs = %q, want %q", got.Outputs, baseline.Outputs)
	}
}

func TestUnmarshal_int64StringInAnonymousField(t *testing.T) {
	type Embedded struct {
		ID int64 `json:"id"`
	}
	type response struct {
		Embedded
	}

	var got response
	if err := Unmarshal([]byte(`{"id":"7"}`), &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.ID != 7 {
		t.Errorf("Unmarshal() ID = %d, want 7", got.ID)
	}
}

func TestUnmarshal_customUnmarshalerInt64String(t *testing.T) {
	got := int64FieldsWithForceSend{}
	if err := json.Unmarshal([]byte(`{"id":"7"}`), &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if got.ID != 7 {
		t.Errorf("json.Unmarshal() ID = %d, want 7", got.ID)
	}
}

func TestUnmarshal_caseInsensitiveFieldName(t *testing.T) {
	type response struct {
		Value string
	}

	var got response
	if err := Unmarshal([]byte(`{"value":"test"}`), &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.Value != "test" {
		t.Errorf("Unmarshal() Value = %q, want %q", got.Value, "test")
	}
}

type int64FieldsWithForceSend struct {
	ID int64 `json:"id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *int64FieldsWithForceSend) UnmarshalJSON(data []byte) error {
	return Unmarshal(data, s)
}
