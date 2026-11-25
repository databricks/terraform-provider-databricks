package credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPATCredentials(t *testing.T) {
	testCases := []struct {
		desc        string
		token       string
		wantErr     error
		wantHeaders map[string]string
	}{
		{
			desc:  "valid token",
			token: "dapi1234567890abcdef",
			wantHeaders: map[string]string{
				"Authorization": "Bearer dapi1234567890abcdef",
			},
		},
		{
			desc:    "empty token",
			token:   "",
			wantErr: errTokenRequired,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			gotCreds, gotErr := NewPATCredentials(tc.token)

			if tc.wantErr != nil {
				if !errors.Is(gotErr, tc.wantErr) {
					t.Fatalf("NewPATCredentials() want error %v, got error: %v", tc.wantErr, gotErr)
				}
				return // no need to test the rest
			}
			if gotErr != nil {
				t.Fatalf("NewPATCredentials() want no error, got error: %v", gotErr)
			}
			if gotCreds == nil {
				t.Fatalf("NewPATCredentials() returned nil credentials")
			}

			gotHeaders, err := gotCreds.AuthHeaders(context.Background())
			if err != nil {
				t.Fatalf("AuthHeaders() error = %v", err)
			}

			if !cmp.Equal(gotHeaders, tc.wantHeaders) {
				t.Errorf("AuthHeaders() = %v, want %v", gotHeaders, tc.wantHeaders)
			}
		})
	}
}
