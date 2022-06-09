package common

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIError(t *testing.T) {
	ae := NotFound("ClusterNotReadyException: test")
	ae.Resource = "c"
	assert.Equal(t, "ClusterNotReadyException: test", ae.Error())
	assert.True(t, ae.IsMissing())
	assert.True(t, ae.IsRetriable())

	ae.StatusCode = http.StatusTooManyRequests
	assert.True(t, ae.IsTooManyRequests())
}

func TestCommonErrorFromWorkspaceClientToE2(t *testing.T) {
	ws := DatabricksClient{
		Host: "https://qwerty.cloud.databricks.com/",
	}
	accountsAPIForWorkspaceClient := ws.commonErrorClarity(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://accounts.cloud.databricks.com/api/2.0/accounts/a/log-delivery",
			nil),
	})
	require.Error(t, accountsAPIForWorkspaceClient)
	assert.True(t, strings.HasPrefix(accountsAPIForWorkspaceClient.Error(),
		"Accounts API (/api/2.0/accounts/a/log-delivery) requires you to set accounts.cloud.databricks.com"),
		"Actual message: %s", accountsAPIForWorkspaceClient.Error())

	workspaceAPIFromWorkspaceClient := ws.commonErrorClarity(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://qwerty.cloud.databricks.com/api/2.0/clusters/list",
			nil),
	})
	assert.Nil(t, workspaceAPIFromWorkspaceClient)
}

func TestCommonErrorFromE2ClientToWorkspace(t *testing.T) {
	ws := DatabricksClient{
		Host: "accounts.cloud.databricks.com",
	}
	ws.fixHost()
	accountsAPIForWorkspaceClient := ws.commonErrorClarity(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://querty.cloud.databricks.com/api/2.0/clusters/list",
			nil),
	})
	require.Error(t, accountsAPIForWorkspaceClient)
	assert.True(t, strings.HasPrefix(accountsAPIForWorkspaceClient.Error(),
		"Databricks API (/api/2.0/clusters/list) requires you to set `host` property (or DATABRICKS_HOST env variable)"),
		"Actual message: %s", accountsAPIForWorkspaceClient.Error())

	e2APIFromE2Client := ws.commonErrorClarity(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://accounts.cloud.databricks.com/api/2.0/accounts/a/log-delivery",
			nil),
	})
	assert.Nil(t, e2APIFromE2Client)
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("test error")
}

func (i errReader) Close() error {
	if int(i) > 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func TestParseError_IO(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	var body errReader
	err := ws.parseError(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://querty.cloud.databricks.com/api/2.0/clusters/list",
			nil),
		Body: body,
	})
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "test error"),
		"Actual message: %s", err.Error())
}

func TestParseError_MWS(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	err := ws.parseError(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://accounts.cloud.databricks.com/api/2.0/accounts/a/log-delivery",
			nil),
		Body:       http.NoBody,
		StatusCode: 400,
	})
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(),
		"Accounts API (/api/2.0/accounts/a/log-delivery) requires you to set accounts.cloud.databricks.com"),
		"Actual message: %s", err.Error())
}

func TestParseError_API12(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	err := ws.parseError(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://querty.cloud.databricks.com/api/2.0/clusters/list",
			nil),
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"error": "Error from API 1.2"
		}`))),
	})
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Error from API 1.2"),
		"Actual message: %s", err.Error())
}

func TestParseError_Enhance403(t *testing.T) {
	ws := DatabricksClient{
		Host:  "qwerty.cloud.databricks.com",
		Token: "x",
	}
	assert.NoError(t, ws.Authenticate(context.Background()))
	err := ws.parseError(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://querty.cloud.databricks.com/api/2.0/clusters/list",
			nil),
		StatusCode: 403,
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"error_code": "PERMISSION_DENIED",
			"message": "You are not authorized."
		}`))),
	})
	assert.EqualError(t, err, "You are not authorized. Using pat auth: "+
		"host=https://qwerty.cloud.databricks.com, token=***REDACTED***")
}

func TestParseError_SCIM(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	err := ws.parseError(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://querty.cloud.databricks.com/api/2.0/clusters/list",
			nil),
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"detail": "Detailed SCIM message",
			"status": "MALFUNCTION",
			"string_value": "sensitive",
			"token_value": "sensitive",
			"content": "sensitive"
		}`))),
	})
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Detailed SCIM message"),
		"Actual message: %s", err.Error())
}

func TestParseError_SCIMNull(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	err := ws.parseError(&http.Response{
		Request: httptest.NewRequest(
			"GET", "https://querty.cloud.databricks.com/api/2.0/clusters/list",
			nil),
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"detail": "null"
		}`))),
	})
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "SCIM API Internal Error"),
		"Actual message: %s", err.Error())
}

func TestCheckHTTPRetry_Connection(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	retry, err := ws.checkHTTPRetry(context.Background(), nil, &url.Error{
		Err: fmt.Errorf("connection refused"),
		URL: "xyz",
	})
	assert.True(t, retry)
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "connection refused"),
		"Actual message: %s", err.Error())
}

func TestCheckHTTPRetry_NilResp(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	retry, _ := ws.checkHTTPRetry(context.Background(), nil, fmt.Errorf("test error"))
	assert.False(t, retry)
}

func TestCheckHTTPRetry_429(t *testing.T) {
	ws := DatabricksClient{
		Host: "qwerty.cloud.databricks.com",
	}
	retry, err := ws.checkHTTPRetry(context.Background(), &http.Response{
		StatusCode: 429,
	}, fmt.Errorf("test error"))
	assert.True(t, retry)
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Current request has to be retried"),
		"Actual message: %s", err.Error())
}

func singleRequestServer(t *testing.T, method, url, response string) (*DatabricksClient, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			if req.Method == method && req.RequestURI == url {
				_, err := rw.Write([]byte(response))
				assert.NoError(t, err)
				return
			}
			assert.Fail(t, fmt.Sprintf("Received unexpected call: %s %s",
				req.Method, req.RequestURI))
		}))
	client := &DatabricksClient{
		Host:               server.URL + "/",
		Token:              "..",
		InsecureSkipVerify: true,
		DebugHeaders:       true,
	}
	err := client.Configure()
	assert.NoError(t, err)
	return client, server
}

func TestGet_Error(t *testing.T) {
	defer CleanupEnvironment()()
	ws := DatabricksClient{}
	ws.configureHTTPCLient()
	err := ws.Get(context.Background(), "/imaginary/endpoint", nil, nil)
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "authentication is not configured"),
		"Actual message: %s", err.Error())
}

func TestPost_Error(t *testing.T) {
	ws, server := singleRequestServer(t, "POST", "/api/2.0/imaginary/endpoint", `{corrupt: "json"`)
	defer server.Close()

	var resp map[string]string
	err := ws.Post(context.Background(), "/imaginary/endpoint", APIErrorBody{
		ScimDetail: "some",
	}, &resp)
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "Invalid JSON received (16 bytes)"),
		"Actual message: %s", err.Error())
}

func TestDelete(t *testing.T) {
	ws, server := singleRequestServer(t, "DELETE", "/api/2.0/imaginary/endpoint", ``)
	defer server.Close()

	err := ws.Delete(context.Background(), "/imaginary/endpoint", APIErrorBody{
		ScimDetail: "some",
	})
	require.NoError(t, err)
}

func TestPatch(t *testing.T) {
	ws, server := singleRequestServer(t, "PATCH", "/api/2.0/imaginary/endpoint", ``)
	defer server.Close()

	err := ws.Patch(context.Background(), "/imaginary/endpoint", APIErrorBody{
		ScimDetail: "some",
	})
	require.NoError(t, err)
}

func TestPut(t *testing.T) {
	ws, server := singleRequestServer(t, "PUT", "/api/2.0/imaginary/endpoint", ``)
	defer server.Close()

	err := ws.Put(context.Background(), "/imaginary/endpoint", APIErrorBody{
		ScimDetail: "some",
	})
	require.NoError(t, err)
}

func TestUnmarshall(t *testing.T) {
	ws := DatabricksClient{}
	err := ws.unmarshall("/a/b/c", nil, nil)
	require.NoError(t, err)
	err = ws.unmarshall("/a/b/c", nil, "abc")
	require.NoError(t, err)
}

func TestAPI2(t *testing.T) {
	ws := DatabricksClient{Host: "ht_tp://example.com/"}
	err := ws.completeUrl(&http.Request{})
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "no URL found in request"),
		"Actual message: %s", err.Error())

	err = ws.completeUrl(&http.Request{
		Header: http.Header{},
		URL: &url.URL{
			Path: "/x/y/x",
		},
	})
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(),
		`parse "ht_tp://example.com/": first path segment in URL cannot contain colon`),
		"Actual message: %s", err.Error())
}

func TestScim(t *testing.T) {
	ws, server := singleRequestServer(t, "GET", "/api/2.0/imaginary/endpoint", `{"a": "b"}`)
	defer server.Close()

	var resp map[string]string
	err := ws.Scim(context.Background(), "GET", "/imaginary/endpoint", nil, &resp)
	require.NoError(t, err)
}

func TestScimFailingQuery(t *testing.T) {
	err := (&DatabricksClient{
		Host:  "https://localhost",
		Token: "..",
	}).Scim(context.Background(), "GET", "/foo", nil, nil)
	assert.EqualError(t, err, "DatabricksClient is not configured")
}

func TestScimVisitorForAccounts(t *testing.T) {
	request := &http.Request{
		Header: http.Header{},
		URL: &url.URL{
			Path: "/api/2.0/preview/scim/v2/Users/abc",
		},
	}
	err := (&DatabricksClient{
		Host:      "https://accounts.everywhere",
		AccountID: "uuid",
	}).scimVisitor(request)
	assert.NoError(t, err)
	assert.Equal(t, "application/scim+json; charset=utf-8", request.Header.Get("Content-Type"))
	assert.Equal(t, "/api/2.0/accounts/uuid/scim/v2/Users/abc", request.URL.Path)
}

func TestMakeRequestBody(t *testing.T) {
	type x struct {
		Scope string `json:"scope" url:"scope"`
	}
	requestURL := "/a/b/c"
	_, err := makeRequestBody("GET", &requestURL, x{"test"})
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c?scope=test", requestURL)

	body, _ := makeRequestBody("POST", &requestURL, "abc")
	assert.Equal(t, []byte("abc"), body)
}

func TestMakeRequestBodyFromReader(t *testing.T) {
	requestURL := "/a/b/c"
	body, err := makeRequestBody("PUT", &requestURL, strings.NewReader("abc"))
	require.NoError(t, err)
	assert.Equal(t, []byte("abc"), body)
}

func TestMakeRequestBodyReaderError(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("POST", &requestURL, errReader(1))
	assert.EqualError(t, err, "failed to read from reader: test error")
}

func TestMakeRequestBodyJsonError(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo chan string `json:"foo"`
	}
	_, err := makeRequestBody("POST", &requestURL, x{make(chan string)})
	assert.EqualError(t, err, "request marshal failure: json: unsupported type: chan string")
}

type failingUrlEncode string

func (fue failingUrlEncode) EncodeValues(key string, v *url.Values) error {
	return fmt.Errorf(string(fue))
}

func TestMakeRequestBodyQueryFailingEncode(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo failingUrlEncode `url:"foo"`
	}
	_, err := makeRequestBody("GET", &requestURL, x{failingUrlEncode("always failing")})
	assert.EqualError(t, err, "cannot create query string: always failing")
}

func TestMakeRequestBodyQueryUnsupported(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("GET", &requestURL, true)
	assert.EqualError(t, err, "unsupported query string data: true")
}

func TestReaderBodyIsNotDumped(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			raw, err := ioutil.ReadAll(req.Body)
			assert.NoError(t, err)
			assert.Equal(t, "abc", string(raw))
			rw.WriteHeader(200)
		}))
	defer server.Close()
	client := &DatabricksClient{
		Host:               server.URL + "/",
		Token:              "..",
		InsecureSkipVerify: true,
		DebugHeaders:       true,
	}
	err := client.Configure()
	assert.NoError(t, err)
	ctx := context.Background()
	err = client.Post(ctx, "/workspace/import-file", strings.NewReader("abc"), nil)
	assert.NoError(t, err)
}

func TestRedactedDumpMalformedJsonReturnsEmptyString(t *testing.T) {
	client := &DatabricksClient{}
	res := client.redactedDump([]byte("{..}"))
	assert.Equal(t, "", res)
}

func TestRedactedDumpOverridesMaxBytes(t *testing.T) {
	client := &DatabricksClient{
		DebugTruncateBytes: 1300,
	}
	res := client.redactedDump([]byte(`{"foo":"` + strings.Repeat("x", 1500) + `"}`))
	assert.Len(t, res, 1319)
	assert.True(t, strings.HasSuffix(res, "... (35 more bytes)"))
}

func TestMakeRequestBodyForMap(t *testing.T) {
	requestURL := "/a"
	_, err := makeRequestBody("GET", &requestURL, map[string]int{
		// i hope this will not trigger false positives too often
		"e": 1,
		"a": 2,
		"f": 3,
		"g": 4,
		"c": 5,
		"b": 6,
		"d": 7,
	})
	require.NoError(t, err)
	assert.Equal(t, "/a?a=2&b=6&c=5&d=7&e=1&f=3&g=4", requestURL)
}

func TestClient_HandleErrors(t *testing.T) {
	tests := []struct {
		name               string
		response           string
		responseStatus     int
		expectedErrorCode  string
		expectedMessage    string
		expectedResource   string
		expectedStatusCode int
		apiCall            func(client *DatabricksClient) error
	}{
		{
			name: "Status 404",
			response: `{
							"error_code": "RESOURCE_DOES_NOT_EXIST",
							"message": "Token ... does not exist!"
						}`,
			responseStatus:     http.StatusNotFound,
			expectedErrorCode:  "RESOURCE_DOES_NOT_EXIST",
			expectedMessage:    "Token ... does not exist!",
			expectedResource:   "/api/2.0/token/create",
			expectedStatusCode: 404,
			apiCall: func(client *DatabricksClient) error {
				return client.Post(context.Background(), "/token/create", map[string]string{
					"foo": "bar",
				}, nil)
			},
		},
		{
			name:               "HTML Status 404",
			response:           `<pre> Hello world </pre>`,
			responseStatus:     http.StatusNotFound,
			expectedErrorCode:  "NOT_FOUND",
			expectedMessage:    "Hello world",
			expectedResource:   "/api/2.0/token/create",
			expectedStatusCode: 404,
			apiCall: func(client *DatabricksClient) error {
				return client.Post(context.Background(), "/token/create", map[string]string{
					"foo": "bar",
				}, nil)
			},
		},
		{
			name:               "Invalid HTML Status 404",
			response:           `<html> Hello world </html>`,
			responseStatus:     http.StatusNotFound,
			expectedErrorCode:  "NOT_FOUND",
			expectedMessage:    "Response from server (404 Not Found) <html> Hello world </html>: invalid character '<' looking for beginning of value",
			expectedResource:   "/api/2.0/token/create",
			expectedStatusCode: 404,
			apiCall: func(client *DatabricksClient) error {
				return client.Post(context.Background(), "/token/create", map[string]string{
					"foo": "bar",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(tt.responseStatus)
				_, err := rw.Write([]byte(tt.response))
				assert.NoError(t, err, err)
			}))
			// Close the server when test finishes
			defer server.Close()
			client := DatabricksClient{
				Host:  server.URL,
				Token: "...",
			}
			err := client.Configure()
			assert.NoError(t, err)

			err = tt.apiCall(&client)
			t.Log(err)
			assert.IsType(t, APIError{}, err)
			assert.Equal(t, tt.expectedErrorCode, err.(APIError).ErrorCode, "error code is not the same")
			assert.Equal(t, tt.expectedMessage, err.(APIError).Message, "message is not the same")
			assert.Equal(t, tt.expectedResource, err.(APIError).Resource, "resource is not the same")
			assert.Equal(t, tt.expectedStatusCode, err.(APIError).StatusCode, "status code is not the same")
		})
	}
}

func TestGenericQueryNotConfigured(t *testing.T) {
	_, err := (&DatabricksClient{}).genericQuery(context.Background(), "GET", "/foo", true)
	assert.EqualError(t, err, "DatabricksClient is not configured")
}

func TestGenericQueryStoppedContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	client := &DatabricksClient{Host: "https://localhost", Token: ".."}
	err := client.Configure()
	assert.NoError(t, err)
	_, err = client.genericQuery(ctx, "GET", "/foo", true)
	assert.EqualError(t, err, "rate limited: context canceled")
}

func TestGenericQueryMarshalError(t *testing.T) {
	ctx := context.Background()
	client := &DatabricksClient{Host: "https://localhost", Token: ".."}
	err := client.Configure()
	assert.NoError(t, err)
	_, err = client.genericQuery(ctx, "POST", "/foo", errReader(1))
	assert.EqualError(t, err, "request marshal: failed to read from reader: test error")
}

func TestGenericQueryInvalidMethod(t *testing.T) {
	ctx := context.Background()
	client := &DatabricksClient{Host: "https://localhost", Token: ".."}
	err := client.Configure()
	assert.NoError(t, err)
	_, err = client.genericQuery(ctx, "😃", "/foo", strings.NewReader("abc"))
	assert.EqualError(t, err, `new request: net/http: invalid method "😃"`)
}

func TestGenericQueryFailingVisitor(t *testing.T) {
	ctx := context.Background()
	client := &DatabricksClient{Host: "https://localhost", Token: ".."}
	err := client.Configure()
	assert.NoError(t, err)
	_, err = client.genericQuery(ctx, "POST", "/foo", strings.NewReader("abc"),
		func(r *http.Request) error {
			return fmt.Errorf("😃")
		})
	assert.EqualError(t, err, `failed visitor: 😃`)
}

func TestGenericQueryFailingRequest(t *testing.T) {
	ctx := context.Background()
	client := &DatabricksClient{Host: "https://localhost", Token: ".."}
	err := client.Configure()
	assert.NoError(t, err)
	client.httpClient.RetryMax = 0
	client.httpClient.ErrorHandler = func(_ *http.Response, _ error, _ int) (*http.Response, error) {
		return nil, fmt.Errorf("😃")
	}
	_, err = client.genericQuery(ctx, "PUT", "https://127.0.0.1/foo", strings.NewReader("abc"))
	assert.EqualError(t, err, `failed request: 😃`)
}

func TestGenericQueryFailingResponseBodyRead(t *testing.T) {
	client, server := singleRequestServer(t, "GET", "/api/2.0/imaginary/endpoint", `{"a": "b"}`)
	defer server.Close()
	client.httpClient.RetryMax = 0
	client.httpClient.ResponseLogHook = func(_ retryablehttp.Logger, r *http.Response) {
		r.Body = errReader(1)
	}
	ctx := context.Background()
	_, err := client.genericQuery(ctx, "GET", fmt.Sprintf("%s/api/2.0/imaginary/endpoint", server.URL), nil)
	assert.EqualError(t, err, "response body: test error")
}

func TestGenericQueryFailingResponseBodyClose(t *testing.T) {
	client, server := singleRequestServer(t, "GET", "/api/2.0/imaginary/endpoint", `{"a": "b"}`)
	defer server.Close()
	client.httpClient.RetryMax = 0
	client.httpClient.ResponseLogHook = func(_ retryablehttp.Logger, r *http.Response) {
		r.Body = errReader(1000)
	}
	ctx := context.Background()
	_, err := client.genericQuery(ctx, "GET", fmt.Sprintf("%s/api/2.0/imaginary/endpoint", server.URL), nil)
	assert.EqualError(t, err, "failed to close: test error")
}

func TestParseUnknownErrorStatusMalformed(t *testing.T) {
	eb := (&DatabricksClient{}).parseUnknownError("malformed", nil, fmt.Errorf("test"))
	assert.Equal(t, "UNKNOWN", eb.ErrorCode)
}
