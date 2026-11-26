// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Lorem Ipsum
package jsonmarshallv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type JsonMarshallV2Interface interface {
	GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error)
}

func NewJsonMarshallV2(client *client.DatabricksClient) *JsonMarshallV2API {
	return &JsonMarshallV2API{
		jsonMarshallV2Impl: jsonMarshallV2Impl{
			client: client,
		},
	}
}

// Lorem Ipsum
type JsonMarshallV2API struct {
	jsonMarshallV2Impl
}
