package common

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
)

func TestRetryOnTimeout(t *testing.T) {
	timeoutErr := errors.New("request failed: request timed out after 1m0s of inactivity")
	otherErr := errors.New("non-retriable")

	testCases := []struct {
		name      string
		callErrs  []error
		wantErr   error
		wantCalls int
	}{
		{name: "success on first call", callErrs: []error{nil}, wantCalls: 1},
		{name: "timeout then succeed", callErrs: []error{timeoutErr, nil}, wantCalls: 2},
		{name: "non-timeout halts", callErrs: []error{otherErr}, wantErr: otherErr, wantCalls: 1},
		{name: "timeout then non-timeout halts", callErrs: []error{timeoutErr, otherErr}, wantErr: otherErr, wantCalls: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			_, err := RetryOnTimeout(context.Background(), func(ctx context.Context) (*struct{}, error) {
				e := tc.callErrs[calls]
				calls++
				return nil, e
			})
			if calls != tc.wantCalls {
				t.Errorf("call count = %d, want %d", calls, tc.wantCalls)
			}
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("err = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestRetryOn504(t *testing.T) {
	otherErr := errors.New("not 504")

	testCases := []struct {
		name      string
		callErrs  []error
		wantErr   error
		wantCalls int
	}{
		{name: "success on first call", callErrs: []error{nil}, wantCalls: 1},
		{name: "504 then succeed", callErrs: []error{apierr.ErrDeadlineExceeded, nil}, wantCalls: 2},
		{name: "wrapped 504 then succeed", callErrs: []error{fmt.Errorf("got 504: %w", apierr.ErrDeadlineExceeded), nil}, wantCalls: 2},
		{name: "non-504 halts", callErrs: []error{otherErr}, wantErr: otherErr, wantCalls: 1},
		{name: "504 then non-504 halts", callErrs: []error{apierr.ErrDeadlineExceeded, otherErr}, wantErr: otherErr, wantCalls: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			_, err := RetryOn504(context.Background(), func(ctx context.Context) (*struct{}, error) {
				e := tc.callErrs[calls]
				calls++
				return nil, e
			})
			if calls != tc.wantCalls {
				t.Errorf("call count = %d, want %d", calls, tc.wantCalls)
			}
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("err = %v, want %v", err, tc.wantErr)
			}
		})
	}
}
