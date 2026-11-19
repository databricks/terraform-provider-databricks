package catalog

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
)

// These tests pin how the classic SDK client handles responses whose body breaks off
// mid-read, on databricks_volume calls and on the token requests that authenticate
// them. Replay fixtures cannot break a body, so these use real servers.

const brokenBodyDiagnostic = "response body: unexpected EOF"

func TestVolumeCallsDoNotRetryBrokenResponseBody(t *testing.T) {
	// Every volume call sends Accept: application/json, so the SDK reads the whole body
	// before it looks at the status, and the read error replaces the status: a broken
	// 429 or 503 is not retried, and a broken 404 keeps the volume in state.
	truncated := func(status int, body string) volumeResponse {
		return volumeResponse{truncated: true, status: status, body: body}
	}
	notFound := truncated(http.StatusNotFound,
		`{"error_code": "NOT_FOUND", "message": "Volume does not exist"}`)
	tests := []struct {
		name       string
		run        volumeCall
		response   volumeResponse
		request    volumeRequest
		diagnostic string
		id         string
	}{
		{
			name:       "refresh GET 200",
			run:        refreshRetryVolume,
			response:   truncated(http.StatusOK, retryVolumeJSON),
			request:    refreshVolumeGet,
			diagnostic: "cannot read volume: " + brokenBodyDiagnostic,
			id:         retryVolumeID,
		},
		{
			name:       "refresh GET 404",
			run:        refreshRetryVolume,
			response:   notFound,
			request:    refreshVolumeGet,
			diagnostic: "cannot read volume: " + brokenBodyDiagnostic,
			id:         retryVolumeID,
		},
		{
			name:       "refresh GET 429",
			run:        refreshRetryVolume,
			response:   truncated(http.StatusTooManyRequests, tooManyRequestsJSON),
			request:    refreshVolumeGet,
			diagnostic: "cannot read volume: " + brokenBodyDiagnostic,
			id:         retryVolumeID,
		},
		{
			name: "refresh GET 503",
			run:  refreshRetryVolume,
			response: truncated(http.StatusServiceUnavailable, `{"error_code": `+
				`"TEMPORARILY_UNAVAILABLE", "message": "Try again later"}`),
			request:    refreshVolumeGet,
			diagnostic: "cannot read volume: " + brokenBodyDiagnostic,
			id:         retryVolumeID,
		},
		{
			// The server created the volume, but it is not saved in state.
			name:       "create POST 200",
			run:        createVolumeWithOwner,
			response:   truncated(http.StatusOK, retryVolumeJSON),
			request:    createVolumePost,
			diagnostic: "cannot create volume: " + brokenBodyDiagnostic,
			id:         "",
		},
		{
			// Unlike an intact 404, it fails the destroy.
			name:       "DELETE 404",
			run:        destroyVolume,
			response:   notFound,
			request:    deleteVolumeRequest,
			diagnostic: "cannot delete volume: " + brokenBodyDiagnostic,
			id:         retryVolumeID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// A retry would get a 200 and fail the request assertion.
			api, server := newScriptedVolumeServer(t, tt.response,
				volumeResponse{status: http.StatusOK, body: retryVolumeJSON})
			c := newVolumeClient(t, &config.Config{
				Host:                server.URL,
				RetryTimeoutSeconds: retryUntilContextEnds,
			})

			d, diags := tt.run(context.Background(), t, c)

			assertVolumeRequests(t, []volumeRequest{tt.request}, api)
			assert.Equal(t, tt.diagnostic, diagnosticSummary(t, diags))
			assert.Equal(t, tt.id, d.Id())
		})
	}
}

func TestVolumeRefreshTokenResponseCutShort(t *testing.T) {
	// A token request has no Accept header, so the SDK hands a gateway's HTML error
	// response to its error mapper unread, and the mapper's read fails with "failed to
	// read response".
	gatewayPage := "<html><body>Gateway error</body></html>"
	tests := []struct {
		name     string
		metadata string
		config   func(*testing.T) *config.Config
		script   []authExchange
		cutShort []int
		// summary is the refresh error, or empty if the refresh succeeds.
		summary string
	}{
		{
			// The request visitor doesn't retry, so the GET is never sent.
			name:     "m2m_401_fails_the_request_visitor",
			metadata: authAWSMetadata(),
			config:   func(*testing.T) *config.Config { return authM2MConfig("oauth-m2m") },
			script: []authExchange{
				authOIDCDiscovery,
				authM2MToken(http.StatusUnauthorized, gatewayPage),
			},
			cutShort: []int{1},
			summary: `cannot read volume: failed during request visitor: error ` +
				`getting token: Post "https://` + authAWSHost + `/oidc/v1/token": ` +
				`failed to read response`,
		},
		{
			// The SDK keeps the 502 on the error, and the token source retries it.
			name:     "m2m_502_is_retried_by_the_token_source",
			metadata: authAWSMetadata(),
			config:   func(*testing.T) *config.Config { return authM2MConfig("oauth-m2m") },
			script: []authExchange{
				authOIDCDiscovery,
				authM2MToken(http.StatusBadGateway, gatewayPage),
				authM2MToken(http.StatusOK, authBearerToken("m2m-token", 3600)),
				authBearerRead("m2m-token"),
			},
			cutShort: []int{1},
		},
		{
			// The SDK counts an Azure AD 429 as retriable, but it cannot rewind the
			// token request's body, so it gives up at once, and oauth2 resends the
			// request with the secret in the body, as after any first failure.
			name:     "azure_429_is_resent_by_oauth2",
			metadata: authCloudMetadata("Azure"),
			config:   authAzureClientSecretConfig,
			script: []authExchange{
				authAADClientCredentials(authAzureAppID, http.StatusTooManyRequests,
					gatewayPage),
				authAADClientCredentialsInBody(authAzureAppID, http.StatusOK,
					authBearerToken("aad-token", 3600)),
				authAADClientCredentials(authAzureMgmt, http.StatusOK,
					authBearerToken("management-token", 3600)),
				authAzureRead("aad-token", "management-token"),
			},
			cutShort: []int{0},
		},
		{
			// Likewise for a 504.
			name:     "azure_504_is_resent_by_oauth2",
			metadata: authCloudMetadata("Azure"),
			config:   authAzureClientSecretConfig,
			script: []authExchange{
				authAADClientCredentials(authAzureAppID, http.StatusGatewayTimeout,
					gatewayPage),
				authAADClientCredentialsInBody(authAzureAppID, http.StatusOK,
					authBearerToken("aad-token", 3600)),
				authAADClientCredentials(authAzureMgmt, http.StatusOK,
					authBearerToken("management-token", 3600)),
				authAzureRead("aad-token", "management-token"),
			},
			cutShort: []int{0},
		},
		{
			// The SDK cannot parse the Azure AD error, so it reports the HTTP error.
			name:     "azure_400_fails_the_request_visitor",
			metadata: authCloudMetadata("Azure"),
			config:   authAzureClientSecretConfig,
			script: []authExchange{
				authAADClientCredentials(authAzureAppID, http.StatusBadRequest, gatewayPage),
				authAADClientCredentialsInBody(authAzureAppID, http.StatusBadRequest,
					gatewayPage),
			},
			cutShort: []int{0, 1},
			summary: "cannot read volume: failed during request visitor: inner token: " +
				`Post "https://` + authAADToken + `": http 400: failed to read response`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateDefaultChain(t)
			s := &authServer{metadata: tt.metadata, script: tt.script, cutShort: tt.cutShort}

			a := refreshAuthVolume(t, s, tt.config(t))

			if tt.summary == "" {
				a.assertSucceeded(t)
			} else {
				a.assertFailed(t, tt.summary)
			}
		})
	}
}
