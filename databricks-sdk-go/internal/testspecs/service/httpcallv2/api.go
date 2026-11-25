// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Lorem Ipsum
package httpcallv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type HttpCallV2Interface interface {

	// This mimics "old" style post requests which have the resource inlined.
	//
	// Set the `path_param_string` value before calling. The _resource_ is sent as
	// the request **body**. See the [API overview] for details.
	//
	// Supported body styles:
	//
	// - [inline] - referenced
	//
	// Resolution order:
	//
	// 1. inline body 2. referenced resource
	//
	// [API overview]: https://docs.databricks.com/api
	// [inline]: https://docs.databricks.com/api/inline
	CreateResource(ctx context.Context, request CreateResourceRequest) (*Resource, error)

	GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error)

	// This mimics a parameterless custom method: a body:"*" request whose fields
	// are all path-bound, so the request carries Content-Type: application/json but
	// has no JSON body fields to serialize. The canonical body for such a request
	// is the empty object {}. The :sync verb sits after a literal path segment
	// because Databricks rejects a custom verb placed directly after a path
	// variable as ambiguous.
	SyncResource(ctx context.Context, request SyncResourceRequest) (*Resource, error)

	// This mimics "new" style post requests which have a body field.
	//
	// > Prefer this over the inline form.
	UpdateResource(ctx context.Context, request UpdateResourceRequest) (*Resource, error)
}

func NewHttpCallV2(client *client.DatabricksClient) *HttpCallV2API {
	return &HttpCallV2API{
		httpCallV2Impl: httpCallV2Impl{
			client: client,
		},
	}
}

// Lorem Ipsum
type HttpCallV2API struct {
	httpCallV2Impl
}
