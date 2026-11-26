package config

import (
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/require"
)

func TestAuthServerCheckHost(t *testing.T) {
	assertHeaders(t, &Config{
		Host:               "YYY",
		AuthType:           "metadata-service",
		MetadataServiceURL: "http://localhost:1234/metadata/token",
		HTTPTransport: fixtures.MappingTransport{
			"GET /metadata/token": {
				ExpectedHeaders: map[string]string{
					"Accept":                        "application/json",
					"X-Databricks-Host":             "https://YYY",
					"X-Databricks-Metadata-Version": "1",
				},
				Response: someValidToken("abc"),
			},
		},
	}, map[string]string{
		"Authorization": "Bearer abc",
	})
}

func TestAuthServerRefresh(t *testing.T) {
	assertHeaders(t, &Config{
		Host:               "YYY",
		AuthType:           "metadata-service",
		MetadataServiceURL: "http://localhost:1234/metadata/token",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/metadata/token",
				Response: map[string]any{
					"token_type":   "Bearer",
					"access_token": "abc",
					"expires_on":   time.Now().Add(1 * time.Second).Unix(),
				},
			},
			{
				Method:   "GET",
				Resource: "/metadata/token",
				Response: someValidToken("bcd"),
			},
		},
	}, map[string]string{
		"Authorization": "Bearer bcd",
	})
}

func TestAuthServerNotLocalhost(t *testing.T) {
	_, err := authenticateRequest(&Config{
		Host:               "YYY",
		AuthType:           "metadata-service",
		MetadataServiceURL: "http://otherhost/metadata/token",
		HTTPTransport:      fixtures.Failures,
	})
	require.ErrorIs(t, err, errMetadataServiceNotLocalhost)
}

func TestAuthServerMalformed(t *testing.T) {
	_, err := authenticateRequest(&Config{
		Host:               "YYY",
		AuthType:           "metadata-service",
		MetadataServiceURL: "#$%^&*",
		HTTPTransport:      fixtures.Failures,
	})
	require.ErrorIs(t, err, errMetadataServiceMalformed)
}

func TestAuthServerNoContent(t *testing.T) {
	_, err := authenticateRequest(&Config{
		Host:               "YYY",
		AuthType:           "metadata-service",
		MetadataServiceURL: "http://localhost:1234/metadata/token",
		HTTPTransport: fixtures.MappingTransport{
			"GET /metadata/token": {
				Response: ``,
			},
		},
	})
	require.ErrorIs(t, err, errInvalidToken)
}

func TestAuthServerFailures(t *testing.T) {
	_, err := authenticateRequest(&Config{
		Host:               "YYY",
		AuthType:           "metadata-service",
		MetadataServiceURL: "http://localhost:1234/metadata/token",
		HTTPTransport:      fixtures.Failures,
	})
	var httpError *httpclient.HttpError
	require.ErrorAs(t, err, &httpError)
	require.Equal(t, 418, httpError.StatusCode)
}
