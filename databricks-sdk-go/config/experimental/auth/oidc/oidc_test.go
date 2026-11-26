package oidc

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIDTokenSourceFn(t *testing.T) {
	wantToken := &IDToken{Value: "from-func"}
	wantErr := fmt.Errorf("test error")
	wantAud := "func-audience"
	wantCtx := context.Background()

	ts := IDTokenSourceFn(func(gotCtx context.Context, gotAud string) (*IDToken, error) {
		if gotCtx != wantCtx {
			t.Errorf("unexpected context: got %v, want %v", gotCtx, wantCtx)
		}
		if gotAud != wantAud {
			t.Errorf("unexpected audience: got %q, want %q", gotAud, wantAud)
		}
		return wantToken, wantErr
	})

	gotToken, gotErr := ts.IDToken(wantCtx, wantAud)

	if gotErr != wantErr {
		t.Errorf("IDToken() want error: %v, got error: %v", wantErr, gotErr)
	}
	if !cmp.Equal(gotToken, wantToken) {
		t.Errorf("IDToken() token = %v, want %v", gotToken, wantToken)
	}
}

func TestNewEnvIDTokenSource(t *testing.T) {
	testCases := []struct {
		desc     string
		envName  string
		envValue string
		want     *IDToken
		wantErr  bool
	}{
		{
			desc:     "success",
			envName:  "OIDC_TEST_TOKEN_SUCCESS",
			envValue: "test-token-123",
			want:     &IDToken{Value: "test-token-123"},
			wantErr:  false,
		},
		{
			desc:     "missing env var",
			envName:  "OIDC_TEST_TOKEN_MISSING",
			envValue: "",
			want:     nil,
			wantErr:  true,
		},
		{
			desc:     "empty env var",
			envName:  "OIDC_TEST_TOKEN_EMPTY",
			envValue: "",
			want:     nil,
			wantErr:  true,
		},
		{
			desc:     "different variable name",
			envName:  "ANOTHER_OIDC_TOKEN",
			envValue: "another-token-456",
			want:     &IDToken{Value: "another-token-456"},
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Setenv(tc.envName, tc.envValue)

			ts := NewEnvIDTokenSource(tc.envName)
			got, gotErr := ts.IDToken(context.Background(), "")

			if tc.wantErr && gotErr == nil {
				t.Fatalf("IDToken() want error, got none")
			}
			if !tc.wantErr && gotErr != nil {
				t.Fatalf("IDToken() want no error, got error: %v", gotErr)
			}
			if !cmp.Equal(got, tc.want) {
				t.Errorf("IDToken() token = %v, want %v", got, tc.want)
			}
		})
	}
}

type testFile struct {
	filename    string
	filecontent string
}

func TestNewFileTokenSource(t *testing.T) {
	testCases := []struct {
		desc     string
		file     *testFile // file to create
		filepath string
		want     *IDToken
		wantErr  bool
	}{
		{
			desc:     "missing filepath",
			file:     &testFile{filename: "token", filecontent: "content"},
			filepath: "",
			wantErr:  true,
		},
		{
			desc:     "empty file",
			file:     &testFile{filename: "token", filecontent: ""},
			filepath: "token",
			wantErr:  true,
		},
		{
			desc:     "file does not exist",
			filepath: "nonexistent-file",
			wantErr:  true,
		},
		{
			desc:     "file exists",
			file:     &testFile{filename: "token", filecontent: "content"},
			filepath: "token",
			want:     &IDToken{Value: "content"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tmpDir := t.TempDir()

			// Create the test file if any is given.
			if tc.file != nil {
				tokenFile := filepath.Join(tmpDir, tc.file.filename)
				if err := os.WriteFile(tokenFile, []byte(tc.file.filecontent), 0644); err != nil {
					t.Fatalf("failed to create token file: %v", err)
				}
			}

			// Only compute the fully qualified filepath if the relative
			// filepath is given.
			fp := tc.filepath
			if tc.filepath != "" {
				fp = filepath.Join(tmpDir, tc.filepath)
			}

			ts := NewFileTokenSource(fp)
			got, gotErr := ts.IDToken(context.Background(), "")

			if tc.wantErr && gotErr == nil {
				t.Fatalf("IDToken() want error, got none")
			}
			if !tc.wantErr && gotErr != nil {
				t.Fatalf("IDToken() want no error, got error: %v", gotErr)
			}
			if !cmp.Equal(got, tc.want) {
				t.Errorf("IDToken() token = %v, want %v", got, tc.want)
			}
		})
	}
}
