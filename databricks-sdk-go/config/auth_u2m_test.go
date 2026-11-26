package config

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestU2MCredentials_Name(t *testing.T) {
	if got := (u2mCredentials{}).Name(); got != "databricks-cli" {
		t.Errorf("Name() = %q, want %q", got, "databricks-cli")
	}
}

func TestU2MCredentials_Configure(t *testing.T) {
	testCases := []struct {
		name string
		// skip returns true on platforms where the test cannot run.
		// Tests that create mock CLI shell scripts are skipped on Windows.
		skip         func() bool
		cfg          func(t *testing.T) *Config
		wantErr      string
		wantProvider bool
	}{
		{
			name:    "missing host returns error",
			cfg:     func(t *testing.T) *Config { return &Config{} },
			wantErr: "host is required",
		},
		{
			name: "CLI not found returns error",
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: "/nonexistent/path/to/databricks",
				}
			},
			wantErr: "databricks CLI not found",
		},
		{
			name: "legacy CLI detected returns error",
			cfg: func(t *testing.T) *Config {
				// Only a CLI found in PATH is checked; an explicit path is trusted.
				t.Setenv("PATH", filepath.Dir(createLegacyCli(t)))
				return &Config{
					Host: "https://workspace.cloud.databricks.com",
				}
			},
			wantErr: "legacy databricks CLI detected",
		},
		{
			name: "not logged in returns error",
			skip: func() bool { return runtime.GOOS == "windows" },
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: createMockCli(t, "#!/bin/sh\necho 'databricks OAuth is not configured' >&2; exit 1"),
				}
			},
			wantErr: "databricks OAuth is not configured",
		},
		{
			name: "token error returns error",
			skip: func() bool { return runtime.GOOS == "windows" },
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: createMockCli(t, "#!/bin/sh\necho 'refresh token is invalid' >&2; exit 1"),
				}
			},
			wantErr: "refresh token is invalid",
		},
		{
			name: "successful authentication",
			skip: func() bool { return runtime.GOOS == "windows" },
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: createMockCli(t, mockCliTokenResponse(t)),
				}
			},
			wantProvider: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.skip != nil && tc.skip() {
				t.Skip("Skipping on this platform")
			}
			cp, err := (u2mCredentials{}).Configure(context.Background(), tc.cfg(t))

			switch {
			case tc.wantErr != "":
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("got error %v, want error containing %q", err, tc.wantErr)
				}
			case tc.wantProvider:
				if err != nil || cp == nil {
					t.Errorf("got (%v, %v), want (provider, nil)", cp, err)
				}
			}
		})
	}
}

func TestU2MCredentials_Configure_CustomScopesError(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         func(t *testing.T) *Config
		scopeSource *Source
		wantErr     string
	}{
		{
			name: "error when scopes are explicitly set",
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:   "https://workspace.cloud.databricks.com",
					Scopes: []string{"clusters", "jobs"},
				}
			},
			wantErr: ErrCustomScopesNotSupported,
		},
		{
			name: "no error when scopes are empty",
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					Scopes:            []string{},
					DatabricksCliPath: createMockCli(t, mockCliTokenResponse(t)),
				}
			},
		},
		{
			name: "no error when scopes are nil",
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: createMockCli(t, mockCliTokenResponse(t)),
				}
			},
		},
		{
			name: "no error when scopes come from config file",
			cfg: func(t *testing.T) *Config {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					Scopes:            []string{"sql"},
					DatabricksCliPath: createMockCli(t, mockCliTokenResponse(t)),
				}
			},
			scopeSource: &Source{Type: SourceFile, Name: "~/.databrickscfg"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if runtime.GOOS == "windows" {
				t.Skip("Skipping on Windows")
			}
			cfg := tc.cfg(t)
			if tc.scopeSource != nil {
				for i := range ConfigAttributes {
					if ConfigAttributes[i].Name == "scopes" {
						cfg.SetAttrSource(&ConfigAttributes[i], *tc.scopeSource)
						break
					}
				}
			}

			_, err := (u2mCredentials{}).Configure(context.Background(), cfg)
			var gotErr string
			if err != nil {
				gotErr = err.Error()
			}
			if gotErr != tc.wantErr {
				t.Errorf("Configure() error = %q, want %q", gotErr, tc.wantErr)
			}
		})
	}
}

func createMockCli(t *testing.T, script string) string {
	t.Helper()
	cli := filepath.Join(t.TempDir(), "databricks")
	if err := os.WriteFile(cli, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	// Pad to pass size check that distinguishes new CLI from legacy Python CLI
	if err := os.Truncate(cli, databricksCliMinSize+1); err != nil {
		t.Fatal(err)
	}
	return cli
}

func createLegacyCli(t *testing.T) string {
	t.Helper()
	name := "databricks"
	if runtime.GOOS == "windows" {
		name = "databricks.exe"
	}
	cli := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(cli, []byte(legacyCliEntryPoint), 0755); err != nil {
		t.Fatal(err)
	}
	return cli
}

func mockCliTokenResponse(t *testing.T) string {
	t.Helper()
	resp, _ := json.Marshal(cliTokenResponse{
		AccessToken: "test-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(time.Hour).Format(time.RFC3339),
	})
	return "#!/bin/sh\necho '" + string(resp) + "'"
}
