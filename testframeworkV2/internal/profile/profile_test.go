package profile

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// writeCfg drops a fixture .databrickscfg into t.TempDir() and returns
// the absolute path. Tests should use it for path-based loading; for
// the inference helpers (which are the bulk of the suite) we exercise
// inferCloudLevel directly.
func writeCfg(t *testing.T, body string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, ".databrickscfg")
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatalf("write cfg: %v", err)
	}
	return path
}

// TestInferCloudLevel locks in DESIGN.md §10 G9's mapping rules — these
// are the strings the runner compares requires.cloud / requires.level
// against, so a regression here silently breaks skip-checks.
func TestInferCloudLevel(t *testing.T) {
	for _, tc := range []struct {
		name      string
		host      string
		accountID string
		wantCloud Cloud
		wantLevel Level
	}{
		// AWS — workspace and account.
		{
			"AWS workspace via dbc-* host",
			"https://dbc-abcd1234.cloud.databricks.com",
			"",
			CloudAWS, LevelWorkspace,
		},
		{
			"AWS account via accounts.cloud.databricks.com",
			"https://accounts.cloud.databricks.com",
			"abc-account-uuid",
			CloudAWS, LevelAccount,
		},
		{
			"AWS account via account_id alone (no accounts. prefix)",
			"https://my-workspace.cloud.databricks.com",
			"abc-account-uuid",
			CloudAWS, LevelAccount,
		},

		// Azure — workspace and account.
		{
			"Azure workspace via .azuredatabricks.net",
			"https://adb-1234567890.0.azuredatabricks.net",
			"",
			CloudAzure, LevelWorkspace,
		},
		{
			"Azure account via accounts.azuredatabricks.net",
			"https://accounts.azuredatabricks.net",
			"abc-account-uuid",
			CloudAzure, LevelAccount,
		},

		// GCP — workspace and account.
		{
			"GCP workspace via .gcp.databricks.com",
			"https://1234567890.5.gcp.databricks.com",
			"",
			CloudGCP, LevelWorkspace,
		},
		{
			"GCP account via accounts.gcp.databricks.com",
			"https://accounts.gcp.databricks.com",
			"abc-account-uuid",
			CloudGCP, LevelAccount,
		},

		// Edge cases.
		{
			"empty host stays Unknown",
			"",
			"",
			CloudUnknown, LevelWorkspace,
		},
		{
			"unrecognized host stays Unknown",
			"https://my-self-managed.example.com",
			"",
			CloudUnknown, LevelWorkspace,
		},
		{
			"hostname without scheme",
			"accounts.cloud.databricks.com",
			"",
			CloudAWS, LevelAccount,
		},
		{
			"trailing path is stripped",
			"https://dbc-abcd1234.cloud.databricks.com/api/2.0/foo",
			"",
			CloudAWS, LevelWorkspace,
		},
		{
			"port is stripped",
			"https://dbc-abcd1234.cloud.databricks.com:443",
			"",
			CloudAWS, LevelWorkspace,
		},
		{
			"uppercase host is normalized",
			"HTTPS://DBC-ABCD1234.CLOUD.DATABRICKS.COM",
			"",
			CloudAWS, LevelWorkspace,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c, l := inferCloudLevel(tc.host, tc.accountID)
			if c != tc.wantCloud || l != tc.wantLevel {
				t.Errorf("inferCloudLevel(%q, %q) = (%q, %q), want (%q, %q)",
					tc.host, tc.accountID, c, l, tc.wantCloud, tc.wantLevel)
			}
		})
	}
}

func TestLoadFromPath_AccountAWS(t *testing.T) {
	path := writeCfg(t, `
[ACCOUNT_AWS]
host          = https://accounts.cloud.databricks.com
account_id    = abc-account-uuid
client_id     = some-client
client_secret = some-secret
auth_type     = oauth-m2m

[WORKSPACE_AWS]
host  = https://dbc-1234.cloud.databricks.com
token = pat-redacted
`)

	p, err := LoadFromPath(path, "ACCOUNT_AWS")
	if err != nil {
		t.Fatalf("LoadFromPath: %v", err)
	}
	if p.Name != "ACCOUNT_AWS" {
		t.Errorf("Name: got %q want ACCOUNT_AWS", p.Name)
	}
	if p.Host != "https://accounts.cloud.databricks.com" {
		t.Errorf("Host: got %q", p.Host)
	}
	if p.AccountID != "abc-account-uuid" {
		t.Errorf("AccountID: got %q", p.AccountID)
	}
	if p.AuthType != "oauth-m2m" {
		t.Errorf("AuthType: got %q", p.AuthType)
	}
	if p.Cloud != CloudAWS {
		t.Errorf("Cloud: got %q want aws", p.Cloud)
	}
	if p.Level != LevelAccount {
		t.Errorf("Level: got %q want account", p.Level)
	}
	if p.Raw["client_id"] != "some-client" {
		t.Errorf("Raw.client_id: got %q", p.Raw["client_id"])
	}
}

func TestLoadFromPath_WorkspaceGCP(t *testing.T) {
	path := writeCfg(t, `
[WS_GCP]
host  = https://1234567890.5.gcp.databricks.com
token = pat-redacted
`)
	p, err := LoadFromPath(path, "WS_GCP")
	if err != nil {
		t.Fatalf("LoadFromPath: %v", err)
	}
	if p.Cloud != CloudGCP {
		t.Errorf("Cloud: got %q want gcp", p.Cloud)
	}
	if p.Level != LevelWorkspace {
		t.Errorf("Level: got %q want workspace", p.Level)
	}
}

func TestLoadFromPath_HandlesCommentsAndBlanks(t *testing.T) {
	path := writeCfg(t, `
# This is a comment
; This too

[A]
host = https://dbc-aaa.cloud.databricks.com
# inline comments at column 0 between sections are fine

[B]
host = https://dbc-bbb.cloud.databricks.com
`)
	a, err := LoadFromPath(path, "A")
	if err != nil {
		t.Fatalf("Load A: %v", err)
	}
	if a.Host != "https://dbc-aaa.cloud.databricks.com" {
		t.Errorf("A.Host: got %q", a.Host)
	}
	b, err := LoadFromPath(path, "B")
	if err != nil {
		t.Fatalf("Load B: %v", err)
	}
	if b.Host != "https://dbc-bbb.cloud.databricks.com" {
		t.Errorf("B.Host: got %q", b.Host)
	}
}

func TestLoadFromPath_SectionNotFound(t *testing.T) {
	path := writeCfg(t, "[A]\nhost = https://x.cloud.databricks.com\n")
	_, err := LoadFromPath(path, "B")
	if err == nil {
		t.Fatal("expected ErrSectionNotFound, got nil")
	}
	if !errors.Is(err, ErrSectionNotFound) {
		t.Errorf("expected error to wrap ErrSectionNotFound, got: %v", err)
	}
}

func TestLoadFromPath_FileMissing(t *testing.T) {
	_, err := LoadFromPath(filepath.Join(t.TempDir(), "no-such-file"), "X")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if errors.Is(err, ErrSectionNotFound) {
		t.Errorf("missing-file error should NOT be ErrSectionNotFound: %v", err)
	}
}

func TestLoadFromPath_EmptyName(t *testing.T) {
	path := writeCfg(t, "[A]\nhost = https://x.cloud.databricks.com\n")
	_, err := LoadFromPath(path, "")
	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}
}

func TestLoadFromPath_RejectsMalformedLine(t *testing.T) {
	path := writeCfg(t, "[A]\nhost https://no-equals.cloud.databricks.com\n")
	_, err := LoadFromPath(path, "A")
	if err == nil {
		t.Fatal("expected error for malformed line, got nil")
	}
}

func TestSectionExists(t *testing.T) {
	path := writeCfg(t, "[A]\nhost = x\n[B]\nhost = y\n")
	for _, tc := range []struct {
		name string
		want bool
	}{
		{"A", true},
		{"B", true},
		{"C", false},
	} {
		got, err := SectionExists(path, tc.name)
		if err != nil {
			t.Errorf("SectionExists(%q): %v", tc.name, err)
			continue
		}
		if got != tc.want {
			t.Errorf("SectionExists(%q) = %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestSectionExists_RejectsEmptyName(t *testing.T) {
	path := writeCfg(t, "[A]\n")
	_, err := SectionExists(path, "")
	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}
}

func TestDefaultPath_HonorsEnvVar(t *testing.T) {
	t.Setenv("DATABRICKS_CONFIG_FILE", "/custom/path/databrickscfg")
	if got := DefaultPath(); got != "/custom/path/databrickscfg" {
		t.Errorf("DefaultPath: got %q", got)
	}
}

func TestDefaultPath_FallsBackToHome(t *testing.T) {
	t.Setenv("DATABRICKS_CONFIG_FILE", "")
	t.Setenv("HOME", "/home/testuser")
	got := DefaultPath()
	want := filepath.Join("/home/testuser", ".databrickscfg")
	if got != want {
		t.Errorf("DefaultPath: got %q want %q", got, want)
	}
}

// TestLoadFromPath_ImplicitDefaultSection documents the lenient parser
// behaviour: keys before any [SECTION] header land in an implicit
// "DEFAULT" section. The Databricks CLI itself accepts this shape.
func TestLoadFromPath_ImplicitDefaultSection(t *testing.T) {
	path := writeCfg(t, "host = https://dbc-default.cloud.databricks.com\n[A]\nhost = https://dbc-a.cloud.databricks.com\n")
	p, err := LoadFromPath(path, "DEFAULT")
	if err != nil {
		t.Fatalf("Load DEFAULT: %v", err)
	}
	if p.Host != "https://dbc-default.cloud.databricks.com" {
		t.Errorf("DEFAULT.Host: got %q", p.Host)
	}
}

func TestLoad_UsesDefaultPath(t *testing.T) {
	path := writeCfg(t, "[X]\nhost = https://dbc-x.cloud.databricks.com\n")
	t.Setenv("DATABRICKS_CONFIG_FILE", path)
	p, err := Load("X")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if p.Host != "https://dbc-x.cloud.databricks.com" {
		t.Errorf("Host: got %q", p.Host)
	}
}
