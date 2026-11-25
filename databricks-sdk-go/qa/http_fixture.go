package qa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
)

// HTTPFixture defines request structure for test
type HTTPFixture struct {
	Method          string
	Resource        string
	Response        interface{}
	Status          int
	ExpectedRequest interface{}
	ReuseRequest    bool
	MatchAny        bool
}

var HTTPFailures = []HTTPFixture{
	{
		MatchAny:     true,
		ReuseRequest: true,
		Status:       418,
		Response: apierr.APIError{
			ErrorCode:  "NONSENSE",
			StatusCode: 418,
			Message:    "I'm a teapot",
		},
	},
}

type HTTPFixtures []HTTPFixture

// hostMetadataFixture is auto-injected by the qa framework to prevent tests
// from making real HTTP calls to the /.well-known/databricks-config endpoint
// during config resolution. Tests that need to mock host metadata should
// provide their own fixture for this endpoint.
var hostMetadataFixture = HTTPFixture{
	Method:       "GET",
	Resource:     "/.well-known/databricks-config",
	ReuseRequest: true,
	Status:       404,
	Response: apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    "host metadata endpoint auto-stubbed by qa framework",
	},
}

func (fixtures HTTPFixtures) withHostMetadata() HTTPFixtures {
	for _, f := range fixtures {
		if f.Method == "GET" && f.Resource == "/.well-known/databricks-config" {
			return fixtures
		}
	}
	return append(HTTPFixtures{hostMetadataFixture}, fixtures...)
}

// Client creates DatabricksClient for emulated HTTP server
func (fixtures HTTPFixtures) Config(t *testing.T) (*config.Config, *httptest.Server) {
	fixtures = fixtures.withHostMetadata()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		found := false
		for i, fixture := range fixtures {
			if (req.Method == fixture.Method && req.RequestURI == fixture.Resource) || fixture.MatchAny {
				if fixture.Status == 0 {
					rw.WriteHeader(200)
				} else {
					rw.WriteHeader(fixture.Status)
				}
				if fixture.ExpectedRequest != nil {
					buf := new(bytes.Buffer)
					_, err := buf.ReadFrom(req.Body)
					assert.NoError(t, err, err)
					jsonStr, err := json.Marshal(fixture.ExpectedRequest)
					assert.NoError(t, err, err)
					assert.JSONEq(t, string(jsonStr), buf.String(), "json strings do not match")
				}
				if fixture.Response != nil {
					if alreadyJSON, ok := fixture.Response.(string); ok {
						_, err := rw.Write([]byte(alreadyJSON))
						assert.NoError(t, err, err)
					} else {
						responseBytes, err := json.Marshal(fixture.Response)
						if err != nil {
							assert.NoError(t, err, err)
							t.FailNow()
						}
						_, err = rw.Write(responseBytes)
						assert.NoError(t, err, err)
					}
				}
				found = true
				// Reset the request if it is already used
				if !fixture.ReuseRequest {
					fixtures[i] = HTTPFixture{}
				}
				break
			}
		}
		if !found {
			receivedRequest := map[string]interface{}{}
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(req.Body)
			assert.NoError(t, err, err)
			err = json.Unmarshal(buf.Bytes(), &receivedRequest)
			assert.NoError(t, err, err)

			expectedRequest := ""
			if len(receivedRequest) > 0 {
				// guessing model name would require going over AST,
				// which is not something i'm willing to write on my weekend
				expectedRequest += "ExpectedRequest: XXX {\n"
				for key, value := range receivedRequest {
					camel := ""
					for _, part := range strings.Split(key, "_") {
						if len(key) < 4 {
							// golang styles, meh...
							camel += strings.ToUpper(key)
						} else {
							camel += strings.Title(part)
						}
					}
					// best effort prediction of what struct should look like...
					expectedRequest += fmt.Sprintf("					%s: %#v,\n", camel, value)
				}
				expectedRequest += "				},\n"
				expectedRequest += fmt.Sprintf("				// ExpectedRequest: %#v,\n", receivedRequest)
			}
			stub := fmt.Sprintf(`{
				Method:   "%s",
				Resource: "%s",
				%s
				Response: XXX {
					// fill in specific fields...
				},
			},`, req.Method, req.RequestURI, expectedRequest)
			assert.Fail(t, fmt.Sprintf("Missing stub, please add: %s", stub))
			t.FailNow()
		}
	}))
	return &config.Config{
		Credentials:      config.PatCredentials{},
		Host:             server.URL,
		Token:            "x",
		AzureEnvironment: "public",
	}, server
}

// HTTPFixturesApply is a helper method for executing a callback and closing emulated server after
func (fixtures HTTPFixtures) Apply(t *testing.T, callback func(ctx context.Context, cfg *config.Config)) {
	cfg, server := fixtures.Config(t)
	defer server.Close()
	callback(context.Background(), cfg)
}

func (fixtures HTTPFixtures) Client(t *testing.T) (*client.DatabricksClient, *httptest.Server) {
	cfg, server := fixtures.Config(t)
	client, err := client.New(cfg)
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	return client, server
}

func (fixtures HTTPFixtures) ApplyClient(t *testing.T, callback func(ctx context.Context, client *client.DatabricksClient)) {
	client, server := fixtures.Client(t)
	defer server.Close()
	callback(context.Background(), client)
}
