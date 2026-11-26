package config

import (
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/require"
)

func init() {
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
}

func someValidToken(bearer string) any {
	return map[string]any{
		"token_type":   "Bearer",
		"access_token": bearer,
		"expires_on":   time.Now().Add(5 * time.Minute).Unix(),
	}
}

func authenticateRequest(cfg *Config) (*http.Request, error) {
	cfg.ConfigFile = "/dev/null"
	cfg.DebugHeaders = true
	if cfg.HTTPTransport == nil {
		// Credential-selection tests do not contact discovery or token endpoints.
		cfg.HTTPTransport = fixtures.MappingTransport{}
	}
	req, _ := http.NewRequest("GET", "http://localhost", nil)
	err := cfg.Authenticate(req)
	return req, err
}

func assertHeaders(t *testing.T, cfg *Config, expectedHeaders map[string]string) {
	req, err := authenticateRequest(cfg)
	require.NoError(t, err)
	actualHeaders := map[string]string{}
	for k := range req.Header {
		actualHeaders[k] = req.Header.Get(k)
	}
	require.Equal(t, expectedHeaders, actualHeaders)
}

func TestMsiHappyFlow(t *testing.T) {
	assertHeaders(t, &Config{
		AzureUseMSI:     true,
		AzureResourceID: "/a/b/c",
		AuthType:        "azure-msi",
		HTTPTransport: fixtures.MappingTransport{
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.azure.com%2F": {
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
				Response: someValidToken("bcd"),
			},
			"GET /a/b/c?api-version=2018-04-01": {
				Response: `{"properties": {
					"workspaceUrl": "https://abc"
				}}`,
			},
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=2ff814a6-3304-4ab8-85cb-cd0e6f879c1d": {
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
				Response: someValidToken("cde"),
			},
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.core.windows.net%2F": {
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
				Response: someValidToken("def"),
			},
		},
	}, map[string]string{
		"Authorization":                            "Bearer cde",
		"X-Databricks-Azure-Sp-Management-Token":   "def",
		"X-Databricks-Azure-Workspace-Resource-Id": "/a/b/c",
	})
}

func TestMsiFailsOnResolveWorkspace(t *testing.T) {
	_, err := authenticateRequest(&Config{
		AzureUseMSI:     true,
		AzureResourceID: "/a/b/c",
		AuthType:        "azure-msi",
		HTTPTransport: fixtures.MappingTransport{
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.azure.com%2F": {
				Response: someValidToken("bcd"),
			},
			"GET /a/b/c?api-version=2018-04-01": {
				Status: 404,
				Response: azureResourceManagerErrorResponse{
					Error: azureResourceManagerErrorError{
						Message: "nope",
					},
				},
			},
		},
	})
	require.ErrorIs(t, err, apierr.ErrNotFound)
}

func TestMsiTokenNotFound(t *testing.T) {
	_, err := authenticateRequest(&Config{
		AzureUseMSI:     true,
		AzureClientID:   "abc",
		AzureResourceID: "/a/b/c",
		AuthType:        "azure-msi",
		HTTPTransport: fixtures.MappingTransport{
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&client_id=abc&resource=https%3A%2F%2Fmanagement.azure.com%2F": {
				Status:   404,
				Response: `...`,
			},
		},
	})
	require.ErrorIs(t, err, apierr.ErrNotFound)
}

func TestMsiHappyFlowWithHostAndNoResourceID(t *testing.T) {
	assertHeaders(t, &Config{
		Host:        "https://adb-123.4.azuredatabricks.net",
		AzureUseMSI: true,
		AuthType:    "azure-msi",
		HTTPTransport: fixtures.MappingTransport{
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=2ff814a6-3304-4ab8-85cb-cd0e6f879c1d": {
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
				Response: someValidToken("cde"),
			},
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.core.windows.net%2F": {
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
				Response: someValidToken("def"),
			},
		},
	}, map[string]string{
		"Authorization":                          "Bearer cde",
		"X-Databricks-Azure-Sp-Management-Token": "def",
	})
}

func TestMsiInvalidTokenExpiry(t *testing.T) {
	_, err := authenticateRequest(&Config{
		AzureUseMSI:     true,
		AzureResourceID: "/a/b/c",
		AuthType:        "azure-msi",
		HTTPTransport: fixtures.MappingTransport{
			"GET /metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.azure.com%2F": {
				Response: map[string]any{
					"token_type":   "Bearer",
					"access_token": "abc",
					"expires_on":   "12345678912345678901234567890123456789123456789",
				},
			},
		},
	})
	require.ErrorIs(t, err, errInvalidTokenExpiry)
}
