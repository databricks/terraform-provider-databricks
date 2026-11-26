package apierr

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type applyOverridesTestCase struct {
	name          string
	initialError  *APIError
	path          string
	method        string
	expectedError error
}

func (a applyOverridesTestCase) testCase(t *testing.T) {
	resp := &http.Response{
		Request: &http.Request{
			URL: &url.URL{
				Path: a.path,
			},
			Method: a.method,
		},
		StatusCode: a.initialError.StatusCode,
	}
	applyOverrides(context.Background(), a.initialError, resp)
	assert.ErrorIs(t, a.initialError, a.expectedError)
}

var testCases = []applyOverridesTestCase{
	{
		initialError: &APIError{
			ErrorCode:  "INVALID_PARAMETER_VALUE",
			StatusCode: 400,
			Message:    "Cluster abc does not exist",
		},
		path:          "/api/2.1/clusters/get",
		method:        "GET",
		expectedError: ErrResourceDoesNotExist,
	},
	{
		initialError: &APIError{
			ErrorCode:  "INVALID_PARAMETER_VALUE",
			StatusCode: 400,
			Message:    "Job abc does not exist",
		},
		path:          "/api/2.0/jobs/get",
		method:        "GET",
		expectedError: ErrResourceDoesNotExist,
	},
	{
		initialError: &APIError{
			ErrorCode:  "INVALID_PARAMETER_VALUE",
			StatusCode: 400,
			Message:    "Job abc does not exist",
		},
		path:          "/api/2.1/jobs/get",
		method:        "GET",
		expectedError: ErrResourceDoesNotExist,
	},
	{
		initialError: &APIError{
			ErrorCode:  "INVALID_PARAMETER_VALUE",
			StatusCode: 400,
			Message:    "Job abc does not exist",
		},
		path:          "/api/2.2/jobs/get",
		method:        "GET",
		expectedError: ErrResourceDoesNotExist,
	},
	{
		initialError: &APIError{
			ErrorCode:  "INVALID_PARAMETER_VALUE",
			StatusCode: 400,
			Message:    "Notebook abc does not exist",
		},
		path:          "/api/2.0/notebooks/get",
		method:        "GET",
		expectedError: ErrBadRequest,
	},
	{
		initialError: &APIError{
			ErrorCode:  "INVALID_PARAMETER_VALUE",
			StatusCode: 400,
			Message:    "Invalid job attribute",
		},
		path:          "/api/2.0/jobs/get",
		method:        "GET",
		expectedError: ErrBadRequest,
	},
	{
		initialError: &APIError{
			ErrorCode:  "INVALID_PARAMETER_VALUE",
			StatusCode: 404,
			Message:    "Job abc does not exist",
		},
		path:          "/api/2.0/jobs/get",
		method:        "GET",
		expectedError: ErrBadRequest,
	},
}

func TestApplyOverrides(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}
