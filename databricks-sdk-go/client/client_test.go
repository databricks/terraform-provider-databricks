package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type hc func(r *http.Request) (*http.Response, error)

func (cb hc) RoundTrip(r *http.Request) (*http.Response, error) {
	// Return 404 for host metadata endpoint to prevent config resolution
	// from interfering with test assertions.
	if r.Method == "GET" && r.URL.Path == "/.well-known/databricks-config" {
		return &http.Response{
			StatusCode: 404,
			Body:       io.NopCloser(strings.NewReader(`{"error_code":"NOT_FOUND","message":"host metadata endpoint auto-stubbed by test framework"}`)),
			Request:    r,
		}, nil
	}
	return cb(r)
}

func TestSimpleRequestFailsURLError(t *testing.T) {
	c, err := New(&config.Config{
		Host:                "some",
		Token:               "token",
		ConfigFile:          "/dev/null",
		RetryTimeoutSeconds: 1,
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, "GET", r.Method)
			require.Equal(t, "/a/b", r.URL.Path)
			require.Equal(t, "c=d", r.URL.RawQuery)
			require.Equal(t, "f", r.Header.Get("e"))
			auth := r.Header.Get("Authorization")
			require.Equal(t, "Bearer token", auth)
			return nil, fmt.Errorf("nope")
		}),
	})
	require.NoError(t, err)
	err = c.Do(
		context.Background(),
		"GET",
		"/a/b",
		map[string]string{"e": "f"},
		nil,
		map[string]string{"c": "d"},
		nil,
	)
	require.EqualError(t, err, `Get "https://some/a/b?c=d": nope`)
}

func TestSimpleRequestFailsAPIError(t *testing.T) {
	c, err := New(&config.Config{
		Host:                "some",
		Token:               "token",
		ConfigFile:          "/dev/null",
		RetryTimeoutSeconds: 1,
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Request:    r,
				Body:       io.NopCloser(strings.NewReader(`{"error_code": "INVALID_PARAMETER_VALUE", "message": "nope"}`)),
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(
		context.Background(),
		"GET",
		"/a/b",
		map[string]string{"e": "f"},
		nil,
		map[string]string{"c": "d"},
		nil,
	)
	require.EqualError(t, err, "nope")
	require.ErrorIs(t, err, apierr.ErrInvalidParameterValue)
}

func TestETag(t *testing.T) {
	reason := "some_reason"
	domain := "a_domain"
	eTag := "sample_etag"
	c, err := New(&config.Config{
		Host:                "some",
		Token:               "token",
		ConfigFile:          "/dev/null",
		RetryTimeoutSeconds: 1,
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Request:    r,
				Body: io.NopCloser(strings.NewReader(fmt.Sprintf(`{
						"error_code": "RESOURCE_CONFLICT",
						"message": "test_public_workspace_setting",
						"stack_trace": "java.io.PrintWriter@329e4ed3",
						"details": [
						  {
							"@type": "%s",
							"reason": "%s",
							"domain": "%s",
							"metadata": {
							  "etag": "%s"
							}
						  },
						  {
							"@type": "anotherType",
							"reason": "",
							"domain": "",
							"metadata": {
							  "etag": "anotherTag"
							}
						  }
						]
					  }`, "type.googleapis.com/google.rpc.ErrorInfo", reason, domain, eTag))),
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(
		context.Background(),
		"GET",
		"/a/b",
		map[string]string{"e": "f"},
		nil,
		map[string]string{"c": "d"},
		nil,
	)
	details := apierr.GetErrorInfo(err)
	require.Equal(t, 1, len(details))
	errorDetails := details[0]
	require.Equal(t, reason, errorDetails.Reason)
	require.Equal(t, domain, errorDetails.Domain)
	require.Equal(t, map[string]string{
		"etag": eTag,
	}, errorDetails.Metadata)
}

func TestSimpleRequestSucceeds(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	var resp Dummy
	err = c.Do(
		context.Background(),
		"POST",
		"/c",
		nil,
		nil,
		Dummy{1},
		&resp,
	)
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
}

func TestQueryParamsRequestSucceeds(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			if r.URL.RawQuery != "a=b&c=1" {
				return nil, fmt.Errorf("unexpected query params: %s", r.URL.RawQuery)
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	var resp Dummy
	err = c.Do(
		context.Background(),
		"POST",
		"/c",
		nil,
		map[string]any{
			"a": "b",
			"c": 1,
		},
		Dummy{1},
		&resp,
	)
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
}

func TestSimpleRequestRetried(t *testing.T) {
	type Dummy struct {
		Foo int `json:"foo"`
	}
	var retried [1]bool
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			if !retried[0] {
				retried[0] = true
				return nil, &url.Error{
					Op:  "open",
					URL: "/a/b",
					Err: fmt.Errorf("connection refused"),
				}
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"foo": 2}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	var resp Dummy
	err = c.Do(
		context.Background(),
		"PATCH",
		"/a",
		nil,
		nil,
		Dummy{1},
		&resp,
	)
	require.NoError(t, err)
	require.Equal(t, 2, resp.Foo)
	require.True(t, retried[0], "request was not retried")
}

func TestSimpleRequestAPIError(t *testing.T) {
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body: io.NopCloser(strings.NewReader(`{
						"error_code": "NOT_FOUND",
						"message": "Something was not found"
					}`)),
				Request: r,
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(
		context.Background(),
		"PATCH",
		"/a",
		nil,
		nil,
		map[string]any{},
		nil,
	)
	var aerr *apierr.APIError
	require.ErrorAs(t, err, &aerr)
	require.Equal(t, "NOT_FOUND", aerr.ErrorCode)
	require.ErrorIs(t, err, apierr.ErrNotFound)
}

func TestHttpTransport(t *testing.T) {
	calledMock := false
	cfg := &config.Config{
		Host:  "localhost",
		Token: "...",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			calledMock = true
			return &http.Response{Request: r}, nil
		}),
	}
	client, err := New(cfg)
	require.NoError(t, err)

	err = client.Do(
		context.Background(),
		"GET",
		"/a",
		nil,
		nil,
		nil,
		bytes.Buffer{},
	)
	require.NoError(t, err)
	require.True(t, calledMock)
}

func TestDoRemovesDoubleSlashesFromFilesAPI(t *testing.T) {
	i := 0
	expectedPaths := []string{
		"/api/2.0/fs/files/Volumes/abc/def/ghi",
		"/api/2.0/anotherservice//test",
	}
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, expectedPaths[i], r.URL.Path)
			i++
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(
		context.Background(),
		"GET",
		"/api/2.0/fs/files//Volumes/abc/def/ghi",
		nil,
		nil,
		map[string]any{},
		nil,
	)
	require.NoError(t, err)
	err = c.Do(
		context.Background(),
		"GET",
		"/api/2.0/anotherservice//test",
		nil,
		nil,
		map[string]any{},
		nil,
	)
	require.NoError(t, err)
}

func TestNonJSONResponseIncludedInError(t *testing.T) {
	type testCase struct {
		statusCode   int
		status       string
		errorMessage string
	}
	cases := []testCase{
		{
			statusCode:   400,
			status:       "Bad Request",
			errorMessage: httpclient.ErrHTMLContent.Error(),
		},
		{
			statusCode:   500,
			status:       "Internal Server Error",
			errorMessage: httpclient.ErrHTMLContent.Error(),
		},
		{
			statusCode:   200,
			status:       "OK",
			errorMessage: httpclient.ErrHTMLContent.Error(),
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d %s", tc.statusCode, tc.status), func(t *testing.T) {
			c, err := New(&config.Config{
				Host:       "some",
				Token:      "token",
				ConfigFile: "/dev/null",
				HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
					r.Header.Del("traceparent") // clear nondeterministic traceparent header
					return &http.Response{
						Proto:   "HTTP/2.0",
						Status:  tc.status,
						Body:    io.NopCloser(strings.NewReader(`<html><body>hello</body></html>`)),
						Request: r,
						Header: http.Header{
							"Content-Type": []string{"text/html"},
						},
					}, nil
				}),
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			var m map[string]string
			gotErr := c.Do(context.Background(), "GET", "/a", nil, nil, nil, &m)

			if gotErr == nil {
				t.Fatalf("expected error, got nil")
			}
			if !strings.Contains(gotErr.Error(), tc.errorMessage) {
				t.Fatalf("expected error to contain %q, got %q", tc.errorMessage, gotErr.Error())
			}
		})
	}
}

func captureUserAgent(t *testing.T) string {
	var userAgent string
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			// Capture the user agent via the round tripper.
			userAgent = r.UserAgent()

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)

	err = c.Do(
		context.Background(),
		"GET",
		"/a",
		nil,
		nil,
		nil,
		nil,
	)
	require.NoError(t, err)

	return userAgent
}

func TestUserAgentForDBR(t *testing.T) {
	for v, sv := range map[string]string{
		// DBR versions that don't need to be sanitized.
		"client.0": "client.0",
		"client.1": "client.1",
		"15.5":     "15.5",
		"15.5.0":   "15.5.0",
		"13.3":     "13.3",

		// DBR versions that need to be sanitized.
		"foo🧟bar": "foo-bar",
		"foo/bar": "foo-bar",
		"foo bar": "foo-bar",
	} {
		t.Run(v, func(t *testing.T) {
			env.CleanupEnvironment(t)
			useragent.ClearCache()

			t.Setenv("DATABRICKS_RUNTIME_VERSION", v)
			userAgent := captureUserAgent(t)

			// The user agent should contain the runtime version, with the value
			// sanitized if necessary.
			assert.Contains(t, userAgent, "runtime/"+sv)
		})
	}
}

func TestUserAgentForCiCd(t *testing.T) {
	ciToEnv := map[string]map[string]string{
		"github": {
			"GITHUB_ACTIONS": "true",
		},
		"gitlab": {
			"GITLAB_CI": "true",
		},
		"jenkins": {
			"JENKINS_URL": "https://jenkins.example.com",
		},
		"azure-devops": {
			"TF_BUILD": "True",
		},
		"circle": {
			"CIRCLECI": "true",
		},
		"travis": {
			"TRAVIS": "true",
		},
		"bitbucket": {
			"BITBUCKET_BUILD_NUMBER": "123",
		},
		"google-cloud-build": {
			"PROJECT_ID":     "",
			"BUILD_ID":       "",
			"PROJECT_NUMBER": "",
			"LOCATION":       "",
		},
		"aws-code-build": {
			"CODEBUILD_BUILD_ARN": "",
		},
		"tf-cloud": {
			"TFC_RUN_ID": "",
		},
	}

	for ci, envVars := range ciToEnv {
		t.Run(ci, func(t *testing.T) {
			env.CleanupEnvironment(t)
			useragent.ClearCache()

			for k, v := range envVars {
				t.Setenv(k, v)
			}

			userAgent := captureUserAgent(t)

			// The user agent should contain the CI/CD provider.
			assert.Contains(t, userAgent, "cicd/"+ci)
		})
	}

}

func TestUserAgentForAgent(t *testing.T) {
	agentToEnv := map[string]map[string]string{
		"claude-code": {
			"CLAUDECODE": "1",
		},
		"cursor": {
			"CURSOR_AGENT": "1",
		},
		"codex": {
			"CODEX_CI": "1",
		},
	}

	for agent, envVars := range agentToEnv {
		t.Run(agent, func(t *testing.T) {
			env.CleanupEnvironment(t)
			useragent.ClearCache()

			for k, v := range envVars {
				t.Setenv(k, v)
			}

			userAgent := captureUserAgent(t)

			// The user agent should contain the agent provider.
			assert.Contains(t, userAgent, "agent/"+agent)

			// There should be exactly one agent entry.
			assert.Equal(t, 1, strings.Count(userAgent, "agent/"))
		})
	}
}

func TestUserAgentForMultipleAgents(t *testing.T) {
	env.CleanupEnvironment(t)
	useragent.ClearCache()

	t.Setenv("CLAUDECODE", "1")
	t.Setenv("CURSOR_AGENT", "1")

	userAgent := captureUserAgent(t)

	// Nested agents (e.g. Claude Code spawning a Cursor CLI subagent) set
	// multiple explicit matchers on the same process. Report "multiple" so
	// the stacked case is visible in telemetry.
	assert.Contains(t, userAgent, "agent/multiple")
}

func TestUserAgentForAgentNotSet(t *testing.T) {
	env.CleanupEnvironment(t)
	useragent.ClearCache()

	userAgent := captureUserAgent(t)

	// When no agent env vars are set, no agent/ entry should appear.
	assert.NotContains(t, userAgent, "agent/")
}

func TestRetryOn503(t *testing.T) {
	var requested bool
	c, err := New(&config.Config{
		Host:       "some",
		Token:      "token",
		ConfigFile: "/dev/null",
		HTTPTransport: hc(func(r *http.Request) (*http.Response, error) {
			if !requested {
				requested = true
				return &http.Response{
					StatusCode: 503,
					Body:       io.NopCloser(strings.NewReader(`{}`)),
					Request:    r,
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Request:    r,
			}, nil
		}),
	})
	require.NoError(t, err)
	err = c.Do(
		context.Background(),
		"GET",
		"/a/b",
		nil,
		nil,
		map[string]any{},
		nil,
	)
	require.NoError(t, err)
}
