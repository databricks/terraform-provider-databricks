// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package generated_tests

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/httpcallv2"
	"github.com/databricks/databricks-sdk-go/qa"
)

func TestHttpCall_LegacyHttpPostNoQueryParamsNoBody(t *testing.T) {
	input := httpcallv2.CreateResourceRequest{
		PathParamString: "string_val",
		PathParamInt:    123,
		PathParamBool:   true,
	}
	qa.HTTPFixtures{
		{
			Method:   "POST",
			Resource: "/api/2.0/http-call/string_val/123/true",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.CreateResource(ctx, input)
	})
}

func TestHttpCall_LegacyHttpPostWithBody(t *testing.T) {
	input := httpcallv2.CreateResourceRequest{
		BodyField:       "request_body_content",
		PathParamString: "test_string",
		PathParamInt:    456,
		PathParamBool:   false,
	}
	qa.HTTPFixtures{
		{
			Method:   "POST",
			Resource: "/api/2.0/http-call/test_string/456/false",
			ExpectedRequest: httpcallv2.CreateResourceRequest{
				BodyField:       "request_body_content",
				PathParamString: "test_string",
				PathParamInt:    456,
				PathParamBool:   false,
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.CreateResource(ctx, input)
	})
}

func TestHttpCall_UpdateResourceNoQueryParamsNoBody(t *testing.T) {
	input := httpcallv2.UpdateResourceRequest{
		Resource: httpcallv2.Resource{
			AnyField:              json.RawMessage("{\"key\": \"value\"}"),
			NestedPathParamBool:   true,
			NestedPathParamInt:    789,
			NestedPathParamString: "update_string",
		},
		NestedPathParamString: "update_string",
		NestedPathParamInt:    789,
		NestedPathParamBool:   true,
	}
	qa.HTTPFixtures{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/http-call/update_string/789/true",
			ExpectedRequest: httpcallv2.Resource{
				AnyField:              json.RawMessage("{\"key\": \"value\"}"),
				NestedPathParamBool:   true,
				NestedPathParamInt:    789,
				NestedPathParamString: "update_string",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.UpdateResource(ctx, input)
	})
}

func TestHttpCall_UpdateResourceWithBody(t *testing.T) {
	input := httpcallv2.UpdateResourceRequest{
		Resource: httpcallv2.Resource{
			BodyField:             "request_body_content",
			NestedPathParamBool:   true,
			NestedPathParamInt:    789,
			NestedPathParamString: "update_string",
		},
		NestedPathParamString: "update_string",
		NestedPathParamInt:    789,
		NestedPathParamBool:   true,
	}
	qa.HTTPFixtures{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/http-call/update_string/789/true",
			ExpectedRequest: httpcallv2.Resource{
				BodyField:             "request_body_content",
				NestedPathParamBool:   true,
				NestedPathParamInt:    789,
				NestedPathParamString: "update_string",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.UpdateResource(ctx, input)
	})
}

func TestHttpCall_UpdateResourceWithSimpleQueryParams(t *testing.T) {
	input := httpcallv2.UpdateResourceRequest{
		Resource: httpcallv2.Resource{
			NestedPathParamBool:   true,
			NestedPathParamInt:    789,
			NestedPathParamString: "update_string",
		},
		NestedPathParamString: "update_string",
		NestedPathParamInt:    789,
		NestedPathParamBool:   true,
		QueryParamString:      "query_string_val",
		QueryParamInt:         999,
		QueryParamBool:        true,
		FieldMask:             &fieldmask.FieldMask{Paths: []string{"field.mask.value"}},
	}
	qa.HTTPFixtures{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/http-call/update_string/789/true?field_mask=field.mask.value&query_param_bool=true&query_param_int=999&query_param_string=query_string_val",
			ExpectedRequest: httpcallv2.Resource{
				NestedPathParamBool:   true,
				NestedPathParamInt:    789,
				NestedPathParamString: "update_string",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.UpdateResource(ctx, input)
	})
}

func TestHttpCall_UpdateResourceWithOneNestedQueryParam(t *testing.T) {
	input := httpcallv2.UpdateResourceRequest{
		Resource: httpcallv2.Resource{
			NestedPathParamBool:   true,
			NestedPathParamInt:    789,
			NestedPathParamString: "update_string",
		},
		NestedPathParamString: "update_string",
		NestedPathParamInt:    789,
		NestedPathParamBool:   true,
		OptionalComplexQueryParam: &httpcallv2.ComplexQueryParam{
			NestedOptionalQueryParam: "nested_optional",
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/http-call/update_string/789/true?optional_complex_query_param=%26%7Bnested_optional%20%5B%5D%20%5B%5D%7D",
			ExpectedRequest: httpcallv2.Resource{
				NestedPathParamBool:   true,
				NestedPathParamInt:    789,
				NestedPathParamString: "update_string",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.UpdateResource(ctx, input)
	})
}

func TestHttpCall_UpdateResourceWithRepeatedQueryParam(t *testing.T) {
	input := httpcallv2.UpdateResourceRequest{
		Resource: httpcallv2.Resource{
			NestedPathParamBool:   true,
			NestedPathParamInt:    789,
			NestedPathParamString: "update_string",
		},
		NestedPathParamString: "update_string",
		NestedPathParamInt:    789,
		NestedPathParamBool:   true,
		RepeatedQueryParam: []string{
			"item1",
			"item2",
			"item3",
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/http-call/update_string/789/true?repeated_query_param=%5Bitem1%20item2%20item3%5D",
			ExpectedRequest: httpcallv2.Resource{
				NestedPathParamBool:   true,
				NestedPathParamInt:    789,
				NestedPathParamString: "update_string",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.UpdateResource(ctx, input)
	})
}

func TestHttpCall_UpdateResourceWithRepeatedNestedQueryParam(t *testing.T) {
	input := httpcallv2.UpdateResourceRequest{
		Resource: httpcallv2.Resource{
			NestedPathParamBool:   true,
			NestedPathParamInt:    789,
			NestedPathParamString: "update_string",
		},
		NestedPathParamString: "update_string",
		NestedPathParamInt:    789,
		NestedPathParamBool:   true,
		OptionalComplexQueryParam: &httpcallv2.ComplexQueryParam{
			NestedRepeatedQueryParam: []string{
				"item1",
				"item2",
				"item3",
			},
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/http-call/update_string/789/true?optional_complex_query_param=%26%7B%20%5Bitem1%20item2%20item3%5D%20%5B%5D%7D",
			ExpectedRequest: httpcallv2.Resource{
				NestedPathParamBool:   true,
				NestedPathParamInt:    789,
				NestedPathParamString: "update_string",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.UpdateResource(ctx, input)
	})
}

func TestHttpCall_UpdateResourceWithDoubleRepeatedNestedQueryParam(t *testing.T) {
	input := httpcallv2.UpdateResourceRequest{
		Resource: httpcallv2.Resource{
			NestedPathParamBool:   true,
			NestedPathParamInt:    789,
			NestedPathParamString: "update_string",
		},
		NestedPathParamString: "update_string",
		NestedPathParamInt:    789,
		NestedPathParamBool:   true,
		RepeatedComplexQueryParam: []httpcallv2.ComplexQueryParam{
			httpcallv2.ComplexQueryParam{
				NestedRepeatedQueryParam: []string{
					"item1",
					"item2",
					"item3",
				},
			},
			httpcallv2.ComplexQueryParam{
				NestedRepeatedQueryParam: []string{
					"item4",
					"item5",
					"item6",
				},
			},
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "PATCH",
			Resource: "/api/2.0/http-call/update_string/789/true?repeated_complex_query_param=%5B%7B%20%5Bitem1%20item2%20item3%5D%20%5B%5D%7D%20%7B%20%5Bitem4%20item5%20item6%5D%20%5B%5D%7D%5D",
			ExpectedRequest: httpcallv2.Resource{
				NestedPathParamBool:   true,
				NestedPathParamInt:    789,
				NestedPathParamString: "update_string",
			},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.UpdateResource(ctx, input)
	})
}

func TestHttpCall_GetResourceNoQueryParams(t *testing.T) {
	input := httpcallv2.GetResourceRequest{
		PathParamString: "get_string",
		PathParamInt:    123,
		PathParamBool:   true,
	}
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/http-call/get_string/123/true?",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.GetResource(ctx, input)
	})
}

func TestHttpCall_GetResourceWithSimpleQueryParams(t *testing.T) {
	input := httpcallv2.GetResourceRequest{
		PathParamString:  "get_string",
		PathParamInt:     456,
		PathParamBool:    false,
		QueryParamString: "query_string_val",
		QueryParamInt:    999,
		QueryParamBool:   true,
		FieldMask:        &fieldmask.FieldMask{Paths: []string{"field.mask.value"}},
	}
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/http-call/get_string/456/false?field_mask=field.mask.value&query_param_bool=true&query_param_int=999&query_param_string=query_string_val",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.GetResource(ctx, input)
	})
}

func TestHttpCall_GetResourceWithOneNestedQueryParam(t *testing.T) {
	input := httpcallv2.GetResourceRequest{
		PathParamString: "get_string",
		PathParamInt:    789,
		PathParamBool:   true,
		OptionalComplexQueryParam: &httpcallv2.ComplexQueryParam{
			NestedOptionalQueryParam: "nested_optional",
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/http-call/get_string/789/true?optional_complex_query_param.nested_optional_query_param=nested_optional",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.GetResource(ctx, input)
	})
}

func TestHttpCall_GetResourceWithRepeatedQueryParam(t *testing.T) {
	input := httpcallv2.GetResourceRequest{
		PathParamString: "get_string",
		PathParamInt:    101,
		PathParamBool:   false,
		RepeatedQueryParam: []string{
			"item1",
			"item2",
			"item3",
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/http-call/get_string/101/false?repeated_query_param=item1&repeated_query_param=item2&repeated_query_param=item3",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.GetResource(ctx, input)
	})
}

func TestHttpCall_GetResourceWithRepeatedNestedQueryParam(t *testing.T) {
	input := httpcallv2.GetResourceRequest{
		PathParamString: "get_string",
		PathParamInt:    202,
		PathParamBool:   true,
		OptionalComplexQueryParam: &httpcallv2.ComplexQueryParam{
			NestedRepeatedQueryParam: []string{
				"item1",
				"item2",
				"item3",
			},
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/http-call/get_string/202/true?optional_complex_query_param.nested_repeated_query_param=item1&optional_complex_query_param.nested_repeated_query_param=item2&optional_complex_query_param.nested_repeated_query_param=item3",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.GetResource(ctx, input)
	})
}

func TestHttpCall_GetResourceWithDoubleRepeatedNestedQueryParam(t *testing.T) {
	input := httpcallv2.GetResourceRequest{
		PathParamString: "get_string",
		PathParamInt:    303,
		PathParamBool:   false,
		RepeatedComplexQueryParam: []httpcallv2.ComplexQueryParam{
			httpcallv2.ComplexQueryParam{
				NestedRepeatedQueryParam: []string{
					"item1",
					"item2",
					"item3",
				},
			},
			httpcallv2.ComplexQueryParam{
				NestedRepeatedQueryParam: []string{
					"item4",
					"item5",
					"item6",
				},
			},
		},
	}
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/http-call/get_string/303/false?repeated_complex_query_param=%7B+%5Bitem1+item2+item3%5D+%5B%5D%7D&repeated_complex_query_param=%7B+%5Bitem4+item5+item6%5D+%5B%5D%7D",
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.GetResource(ctx, input)
	})
}

func TestHttpCall_SyncResourceNoBody(t *testing.T) {
	input := httpcallv2.SyncResourceRequest{
		PathParamString: "sync_string",
		PathParamInt:    123,
		PathParamBool:   true,
	}
	qa.HTTPFixtures{
		{
			Method:          "POST",
			Resource:        "/api/2.0/http-call/sync_string/123/true/state:sync",
			ExpectedRequest: httpcallv2.SyncResourceRequest{},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		commands := httpcallv2.NewHttpCallV2(client)
		commands.SyncResource(ctx, input)
	})
}
