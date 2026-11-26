package fixtures_test

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
}

func TestSimpleGetErrorStub(t *testing.T) {
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{},
	})
	ctx := context.Background()
	err := client.Do(ctx, "GET", "/foo",
		httpclient.WithRequestData(map[string]any{
			"foo": "bar",
		}))
	// spacing in this test is very important
	assert.EqualError(t, err, `Get "/foo?foo=bar": missing stub, please add: "GET /foo?foo=bar": {
		Response: XXX {
			// fill in specific fields...
		},
	},`, err.Error())
}

func TestSimplePostErrorStubFilled(t *testing.T) {
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{
			"POST /foo": {
				Status:          203,
				ExpectedRequest: map[string]any{"foo": "bar"},
				Response:        `{"some": "reply"}`,
			},
		},
	})
	var out map[string]string
	ctx := context.Background()
	err := client.Do(ctx, "POST", "/foo",
		httpclient.WithResponseUnmarshal(&out),
		httpclient.WithRequestData(map[string]any{
			"foo": "bar",
		}))
	require.NoError(t, err)
	assert.Equal(t, "reply", out["some"])
}

func TestPassFile(t *testing.T) {
	client := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{
			"GET /some": {
				PassFile: "testdata/some.json",
			},
		},
	})
	var out struct {
		Some string `json:"some"`
	}
	ctx := context.Background()
	err := client.Do(ctx, "GET", "/some",
		httpclient.WithResponseUnmarshal(&out))
	require.NoError(t, err)
	assert.Equal(t, "data", out.Some)
}
