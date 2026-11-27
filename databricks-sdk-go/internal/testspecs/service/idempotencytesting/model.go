// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package idempotencytesting

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateTestResourceRequest struct {
	RequestId string `json:"-" url:"request_id,omitempty"`

	TestResource TestResource `json:"test_resource"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateTestResourceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTestResourceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TestResource struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TestResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TestResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
