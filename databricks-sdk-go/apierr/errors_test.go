package apierr

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAPIError_IsRetriable_transientRegexMatches(t *testing.T) {
	err := APIError{
		Message: "worker env WorkerEnvId(workerenv-XXXXX) not found",
	}

	if !err.IsRetriable(context.Background()) {
		t.Errorf("expected error to be retriable")
	}
}

func makeTestReponseWrapper(statusCode int, resp string) common.ResponseWrapper {
	return common.ResponseWrapper{
		Response: &http.Response{
			Status:     fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)),
			StatusCode: statusCode,
			Request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/api/2.0/myservice",
				},
			},
		},
		DebugBytes: []byte{},
		ReadCloser: io.NopCloser(bytes.NewReader([]byte(resp))),
	}
}

func TestGetAPIError_ErrorDetails(t *testing.T) {
	want := ErrorDetails{ErrorInfo: &ErrorInfo{
		Reason: "reason",
		Domain: "domain",
	}}
	APIError := &APIError{errorDetails: want}

	got := APIError.ErrorDetails()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected error details (-want +got):\n%s", diff)
	}
}

func TestGetAPIError(t *testing.T) {
	testCases := []struct {
		name        string
		resp        common.ResponseWrapper
		want        *APIError
		wantErrorIs error
	}{
		{
			name: "empty response",
			resp: makeTestReponseWrapper(http.StatusNotFound, ""),
			want: &APIError{
				ErrorCode:  "",
				Message:    "Not Found",
				StatusCode: http.StatusNotFound,
			},
			wantErrorIs: ErrNotFound,
		},
		{
			name: "happy path",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST",
				"message": "Cluster abc does not exist"
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
			},
			wantErrorIs: ErrResourceDoesNotExist,
		},
		{
			name: "string error response",
			resp: makeTestReponseWrapper(http.StatusBadRequest, `MALFORMED_REQUEST: vpc_endpoints malformed parameters: VPC Endpoint ... with use_case ... cannot be attached in ... list`),
			want: &APIError{
				ErrorCode:  "MALFORMED_REQUEST",
				Message:    "vpc_endpoints malformed parameters: VPC Endpoint ... with use_case ... cannot be attached in ... list",
				StatusCode: http.StatusBadRequest,
			},
			wantErrorIs: ErrBadRequest,
		},
		{
			name: "numeric error code",
			resp: makeTestReponseWrapper(http.StatusBadRequest, `{
				"error_code": 500,
				"message": "Cluster abc does not exist"}
			`),
			want: &APIError{
				ErrorCode:  "500",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusBadRequest,
			},
			wantErrorIs: ErrBadRequest,
		},
		{
			name: "private link redirect",
			resp: common.ResponseWrapper{
				Response: &http.Response{
					Request: &http.Request{
						URL: &url.URL{
							Host:     "adb-12345678910.1.azuredatabricks.net",
							Path:     "/login.html",
							RawQuery: "error=private-link-validation-error",
						},
					},
				},
			},
			want: &APIError{
				ErrorCode:  "PRIVATE_LINK_VALIDATION_ERROR",
				Message:    "The requested workspace has Azure Private Link enabled and is not accessible from the current network. Ensure that Azure Private Link is properly configured and that your device has access to the Azure Private Link endpoint. For more information, see https://learn.microsoft.com/en-us/azure/databricks/security/network/classic/private-link-standard#authentication-troubleshooting.",
				StatusCode: http.StatusForbidden,
			},
			wantErrorIs: ErrPermissionDenied,
		},
		{
			name: "applies overrides",
			resp: common.ResponseWrapper{
				Response: &http.Response{
					StatusCode: http.StatusBadRequest,
					Request: &http.Request{
						Method: "GET",
						URL: &url.URL{
							Path: "/api/2.1/clusters/get",
						},
					},
				},
				DebugBytes: []byte{},
				ReadCloser: io.NopCloser(bytes.NewReader([]byte(`{
					"error_code": "INVALID_PARAMETER_VALUE",
					"message": "Cluster abc does not exist"}
				`))),
			},
			want: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusBadRequest,
			},
			wantErrorIs: ErrResourceDoesNotExist,
		},
		{
			name: "unknown error message",
			resp: makeTestReponseWrapper(http.StatusInternalServerError, "unknown error message"),
			want: &APIError{
				ErrorCode:  "UNKNOWN", // default error code
				Message:    "unknown error message",
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "all error details type",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": "reason",
						"domain": "domain",
						"metadata": {"k1": "v1", "k2": "v2"}
					},
					{
						"@type": "type.googleapis.com/google.rpc.RequestInfo",
						"request_id": "req42",
						"serving_data": "data"
					},
					{
						"@type": "type.googleapis.com/google.rpc.RetryInfo",
						"retry_delay": "42.0000000012s"
					},
					{
						"@type": "type.googleapis.com/google.rpc.DebugInfo",
						"stack_entries": ["entry1", "entry2"],
						"detail": "detail"
					},
					{
						"@type": "type.googleapis.com/google.rpc.QuotaFailure",
						"violations": [{"subject": "subject", "description": "description"}]
					},
					{
						"@type": "type.googleapis.com/google.rpc.PreconditionFailure",
						"violations": [{"type": "type", "subject": "subject", "description": "description"}]
					},
					{
						"@type": "type.googleapis.com/google.rpc.BadRequest",
						"field_violations": [{"field": "field", "description": "description"}]
					},
					{
						"@type": "type.googleapis.com/google.rpc.ResourceInfo",
						"resource_type": "resource_type",
						"resource_name": "resource_name",
						"owner": "owner",
						"description": "description"
					},
					{
						"@type": "type.googleapis.com/google.rpc.Help",
						"links": [{"description": "description", "url": "url"}]
					}
				]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{
					{
						Type:   "type.googleapis.com/google.rpc.ErrorInfo",
						Reason: "reason",
						Domain: "domain",
						Metadata: map[string]string{
							"k1": "v1",
							"k2": "v2",
						},
					},
					{
						Type: "type.googleapis.com/google.rpc.RequestInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.RetryInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.DebugInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.QuotaFailure",
					},
					{
						Type: "type.googleapis.com/google.rpc.PreconditionFailure",
					},
					{
						Type: "type.googleapis.com/google.rpc.BadRequest",
					},
					{
						Type: "type.googleapis.com/google.rpc.ResourceInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.Help",
					},
				},
				errorDetails: ErrorDetails{
					ErrorInfo: &ErrorInfo{
						Reason:   "reason",
						Domain:   "domain",
						Metadata: map[string]string{"k1": "v1", "k2": "v2"},
					},
					RequestInfo: &RequestInfo{
						RequestID:   "req42",
						ServingData: "data",
					},
					RetryInfo: &RetryInfo{
						RetryDelay: 42*time.Second + time.Nanosecond,
					},
					DebugInfo: &DebugInfo{
						StackEntries: []string{"entry1", "entry2"},
						Detail:       "detail",
					},
					QuotaFailure: &QuotaFailure{
						Violations: []QuotaFailureViolation{{Subject: "subject", Description: "description"}},
					},
					PreconditionFailure: &PreconditionFailure{
						Violations: []PreconditionFailureViolation{{Type: "type", Subject: "subject", Description: "description"}},
					},
					BadRequest: &BadRequest{
						FieldViolations: []BadRequestFieldViolation{{Field: "field", Description: "description"}},
					},
					ResourceInfo: &ResourceInfo{
						ResourceType: "resource_type",
						ResourceName: "resource_name",
						Owner:        "owner",
						Description:  "description",
					},
					Help: &Help{
						Links: []HelpLink{{Description: "description", URL: "url"}},
					},
				},
			},
		},
		{
			name: "unknown error details type",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [
					{
						"@type": "foo", 
						"reason": "reason"
					}
				]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{
					{
						Type:   "foo",
						Reason: "reason",
					},
				},
				errorDetails: ErrorDetails{
					UnknownDetails: []any{
						map[string]interface{}{
							"@type":  "foo",
							"reason": "reason",
						},
					},
				},
			},
		},
		{
			name: "invalid error details",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [
					42, 
					"foobar",
					{
						"foo": "bar"
					},
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.RequestInfo",
						"request_id": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.RetryInfo",
						"retry_delay": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.DebugInfo",
						"stack_entries": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.QuotaFailure",
						"violations": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.PreconditionFailure",
						"violations": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.BadRequest",
						"field_violations": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.ResourceInfo",
						"resource_type": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.Help",
						"links": 0
					}
				]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{
					{},
					// No ErrorInfo because it fails to unmarshal.
					{Type: "type.googleapis.com/google.rpc.RequestInfo"},
					{Type: "type.googleapis.com/google.rpc.RetryInfo"},
					{Type: "type.googleapis.com/google.rpc.DebugInfo"},
					{Type: "type.googleapis.com/google.rpc.QuotaFailure"},
					{Type: "type.googleapis.com/google.rpc.PreconditionFailure"},
					{Type: "type.googleapis.com/google.rpc.BadRequest"},
					{Type: "type.googleapis.com/google.rpc.ResourceInfo"},
					{Type: "type.googleapis.com/google.rpc.Help"},
				},
				errorDetails: ErrorDetails{
					UnknownDetails: []any{
						42.0,
						"foobar",
						map[string]interface{}{
							"foo": "bar",
						},
						map[string]interface{}{
							"@type":  "type.googleapis.com/google.rpc.ErrorInfo",
							"reason": 0.0,
						},
						map[string]interface{}{
							"@type":      "type.googleapis.com/google.rpc.RequestInfo",
							"request_id": 0.0,
						},
						map[string]interface{}{
							"@type":       "type.googleapis.com/google.rpc.RetryInfo",
							"retry_delay": 0.0,
						},
						map[string]interface{}{
							"@type":         "type.googleapis.com/google.rpc.DebugInfo",
							"stack_entries": 0.0,
						},
						map[string]interface{}{
							"@type":      "type.googleapis.com/google.rpc.QuotaFailure",
							"violations": 0.0,
						},
						map[string]interface{}{
							"@type":      "type.googleapis.com/google.rpc.PreconditionFailure",
							"violations": 0.0,
						},
						map[string]interface{}{
							"@type":            "type.googleapis.com/google.rpc.BadRequest",
							"field_violations": 0.0,
						},
						map[string]interface{}{
							"@type":         "type.googleapis.com/google.rpc.ResourceInfo",
							"resource_type": 0.0,
						},
						map[string]interface{}{
							"@type": "type.googleapis.com/google.rpc.Help",
							"links": 0.0,
						},
					},
				},
			},
		},
		{
			name: "only keep the last error details of a type",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": "first"
					},
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": "second"
					}
				]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{
					{
						Type:   "type.googleapis.com/google.rpc.ErrorInfo",
						Reason: "first",
					},
					{
						Type:   "type.googleapis.com/google.rpc.ErrorInfo",
						Reason: "second",
					},
				},
				errorDetails: ErrorDetails{
					ErrorInfo: &ErrorInfo{
						Reason: "second",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.want.ResponseWrapper = &tc.resp

			got := GetAPIError(context.Background(), tc.resp)

			opts := cmp.Options{
				cmp.AllowUnexported(APIError{}),            // to check ErrorDetails
				cmpopts.IgnoreFields(APIError{}, "unwrap"), // tested via wantErrorIs

				//  Skip values that implement io.ReadCloser (Body, etc.)
				cmpopts.IgnoreInterfaces(struct{ io.ReadCloser }{}),

				// Ignore unexported fields of http.Response and http.Request
				// so we don’t walk into ctx and other private internals.
				cmpopts.IgnoreUnexported(http.Response{}, http.Request{}),
			}
			if diff := cmp.Diff(tc.want, got, opts); diff != "" {
				t.Errorf("unexpected error (-want +got):\n%s", diff)
			}
			if tc.wantErrorIs != nil && !errors.Is(got, tc.wantErrorIs) {
				t.Errorf("errors.Is(%q, %q) failed", got, tc.wantErrorIs)
			}
		})
	}
}
