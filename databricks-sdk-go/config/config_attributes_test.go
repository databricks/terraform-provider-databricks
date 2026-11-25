package config

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestConfigFile_Configure_ListParsing tests that comma-separated list values
// in configuration files are correctly parsed into slices.
func TestConfigFile_Configure_ListParsing(t *testing.T) {
	testCases := []struct {
		name    string
		profile string
		want    []string
	}{
		{
			name:    "single item",
			profile: "single-item",
			want:    []string{"clusters"},
		},
		{
			name:    "multiple items",
			profile: "multiple-items",
			want:    []string{"alpha", "beta", "gamma"},
		},
		{
			name:    "whitespace around items is trimmed",
			profile: "whitespace-around-items",
			want:    []string{"alpha", "beta", "gamma"},
		},
		{
			name:    "empty items are filtered out",
			profile: "empty-items-filtered",
			want:    []string{"alpha", "beta"},
		},
		{
			name:    "whitespace-only items are filtered out",
			profile: "whitespace-only-items-filtered",
			want:    []string{"alpha", "beta"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			withMockEnv(t, map[string]string{})

			cfg := &Config{
				Profile:    tc.profile,
				ConfigFile: "testdata/list-parsing/.databrickscfg",
			}
			err := ConfigFile.Configure(cfg)
			if err != nil {
				t.Fatalf("Configure failed: %v", err)
			}
			if diff := cmp.Diff(tc.want, cfg.Scopes); diff != "" {
				t.Errorf("list mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConfigAttribute_GetString(t *testing.T) {
	findAttr := func(name string) *ConfigAttribute {
		for i := range ConfigAttributes {
			if ConfigAttributes[i].Name == name {
				return &ConfigAttributes[i]
			}
		}
		t.Fatalf("attribute %q not found", name)
		return nil
	}

	testCases := []struct {
		name     string
		attrName string
		wantKind reflect.Kind
		cfg      *Config
		want     string
	}{
		{
			name:     "string type returns string value",
			attrName: "host",
			wantKind: reflect.String,
			cfg:      &Config{Host: "https://example.databricks.com"},
			want:     "https://example.databricks.com",
		},
		{
			name:     "string type returns empty string when not set",
			attrName: "host",
			wantKind: reflect.String,
			cfg:      &Config{},
			want:     "",
		},
		{
			name:     "bool type returns true",
			attrName: "skip_verify",
			wantKind: reflect.Bool,
			cfg:      &Config{InsecureSkipVerify: true},
			want:     "true",
		},
		{
			name:     "bool type returns false",
			attrName: "skip_verify",
			wantKind: reflect.Bool,
			cfg:      &Config{InsecureSkipVerify: false},
			want:     "false",
		},
		{
			name:     "int type returns numeric string",
			attrName: "http_timeout_seconds",
			wantKind: reflect.Int,
			cfg:      &Config{HTTPTimeoutSeconds: 120},
			want:     "120",
		},
		{
			name:     "int type returns zero",
			attrName: "http_timeout_seconds",
			wantKind: reflect.Int,
			cfg:      &Config{HTTPTimeoutSeconds: 0},
			want:     "0",
		},
		{
			name:     "slice type returns comma-joined string",
			attrName: "scopes",
			wantKind: reflect.Slice,
			cfg:      &Config{Scopes: []string{"alpha", "beta", "gamma"}},
			want:     "alpha,beta,gamma",
		},
		{
			name:     "slice type returns single item",
			attrName: "scopes",
			wantKind: reflect.Slice,
			cfg:      &Config{Scopes: []string{"all-apis"}},
			want:     "all-apis",
		},
		{
			name:     "slice type returns empty string for nil slice",
			attrName: "scopes",
			wantKind: reflect.Slice,
			cfg:      &Config{Scopes: nil},
			want:     "",
		},
		{
			name:     "slice type returns empty string for empty slice",
			attrName: "scopes",
			wantKind: reflect.Slice,
			cfg:      &Config{Scopes: []string{}},
			want:     "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attr := findAttr(tc.attrName)
			if attr.Kind != tc.wantKind {
				t.Fatalf("expected %v kind, got %v", tc.wantKind, attr.Kind)
			}
			got := attr.GetString(tc.cfg)
			if got != tc.want {
				t.Errorf("GetString() = %q, want %q", got, tc.want)
			}
		})
	}
}
