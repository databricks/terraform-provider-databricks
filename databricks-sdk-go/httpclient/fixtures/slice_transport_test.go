package fixtures_test

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestGetErrorStub(t *testing.T) {
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{},
	})
	ctx := context.Background()
	err := client.Do(ctx, "GET", "/foo", httpclient.WithRequestData(map[string]any{
		"foo": "bar",
	}))
	// spacing in this test is very important
	assert.EqualError(t, err, `Get "/foo?foo=bar": missing stub, please add: {
		Method:   "GET",
		Resource: "/foo?foo=bar",
		Response: XXX {
			// fill in specific fields...
		},
	},`, err.Error())
}

func TestGetErrorStubFilled(t *testing.T) {
	var out struct {
		Foo string `json:"foo"`
	}
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/foo?foo=bar",
				Response: map[string]any{
					"foo": "bar",
				},
			},
		},
	})
	ctx := context.Background()
	err := client.Do(ctx, "GET", "/foo",
		httpclient.WithResponseUnmarshal(&out),
		httpclient.WithRequestData(map[string]any{
			"foo": "bar",
		}))
	assert.NoError(t, err)
	assert.Equal(t, "bar", out.Foo)
}

func TestReplyWithString(t *testing.T) {
	var out struct {
		Foo string `json:"foo"`
	}
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/foo?foo=bar",
				Response: `{"foo": "bar"}`,
			},
		},
	})
	ctx := context.Background()
	err := client.Do(ctx, "GET", "/foo",
		httpclient.WithResponseUnmarshal(&out),
		httpclient.WithRequestData(map[string]any{
			"foo": "bar",
		}))
	assert.NoError(t, err)
	assert.Equal(t, "bar", out.Foo)
}

func TestGetErrorStubWithHost(t *testing.T) {
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{},
	})
	ctx := context.Background()
	err := client.Do(ctx, "GET", "http://localhost:1234/foo", httpclient.WithRequestData(map[string]any{
		"foo": "bar",
	}))
	// spacing in this test is very important
	assert.EqualError(t, err, `Get "http://localhost:1234/foo?foo=bar": missing stub, please add: {
		Method:   "GET",
		Resource: "/foo?foo=bar",
		Response: XXX {
			// fill in specific fields...
		},
	},`, err.Error())
}

func TestPostErrorStub(t *testing.T) {
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{},
	})
	ctx := context.Background()
	err := client.Do(ctx, "POST", "/foo", httpclient.WithRequestData(map[string]any{
		"foo": "bar",
	}))
	// spacing in this test is very important
	assert.EqualError(t, err, `Post "/foo": missing stub, please add: {
		Method:   "POST",
		Resource: "/foo",
		ExpectedRequest: XXX {
			Foo: "bar",
		},
		// ExpectedRequest: map[string]interface {}{"foo":"bar"},
		Response: XXX {
			// fill in specific fields...
		},
	},`)
}

func TestPostErrorStubFilled(t *testing.T) {
	type op struct {
		Foo string `json:"foo"`
	}
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{
			{
				Method:   "POST",
				Resource: "/foo",
				ExpectedRequest: op{
					Foo: "bar",
				},
			},
		},
	})
	ctx := context.Background()
	err := client.Do(ctx, "POST", "/foo", httpclient.WithRequestData(map[string]any{
		"foo": "bar",
	}))
	assert.NoError(t, err)
}

func TestPostUnexpectedRequest(t *testing.T) {
	type op struct {
		Foo string `json:"foo"`
	}
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{
			{
				Method:   "POST",
				Resource: "/foo",
				ExpectedRequest: op{
					Foo: "ba..r",
				},
			},
		},
	})
	ctx := context.Background()
	err := client.Do(ctx, "POST", "/foo", httpclient.WithRequestData(map[string]any{
		"foo": "bar",
	}))
	assert.EqualError(t, err, `Post "/foo": expected: want {
  "foo": "ba..r"
}, got {
  "foo": "bar"
}`, err.Error())
}
