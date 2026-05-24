package telemetry

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
)

// The SDK formats each (key, value) pair as "key/value" in the user-agent
// string; both tests assert on token substrings rather than the full string
// so adding fixed SDK dimensions (go version, os, etc.) doesn't break them.

func TestWithResource(t *testing.T) {
	tests := []struct {
		name       string
		ctx        context.Context
		resource   string
		wantTokens []string
	}{
		{
			name:       "tags sdk and resource on a fresh context",
			ctx:        context.Background(),
			resource:   "mws_ncc_private_endpoint_rule",
			wantTokens: []string{"sdk/pluginframework", "resource/mws_ncc_private_endpoint_rule"},
		},
		{
			name:       "preserves earlier user-agent values on an enriched context",
			ctx:        useragent.InContext(context.Background(), "caller", "upstream"),
			resource:   "cluster",
			wantTokens: []string{"caller/upstream", "sdk/pluginframework", "resource/cluster"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := useragent.FromContext(WithResource(tt.ctx, tt.resource))
			for _, want := range tt.wantTokens {
				if !strings.Contains(ua, want) {
					t.Errorf("user-agent %q is missing token %q", ua, want)
				}
			}
		})
	}
}

func TestWithDataSource(t *testing.T) {
	tests := []struct {
		name       string
		ctx        context.Context
		dataSource string
		wantTokens []string
	}{
		{
			name:       "tags sdk and data on a fresh context",
			ctx:        context.Background(),
			dataSource: "cluster",
			wantTokens: []string{"sdk/pluginframework", "data/cluster"},
		},
		{
			name:       "preserves earlier user-agent values on an enriched context",
			ctx:        useragent.InContext(context.Background(), "caller", "upstream"),
			dataSource: "catalog",
			wantTokens: []string{"caller/upstream", "sdk/pluginframework", "data/catalog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := useragent.FromContext(WithDataSource(tt.ctx, tt.dataSource))
			for _, want := range tt.wantTokens {
				if !strings.Contains(ua, want) {
					t.Errorf("user-agent %q is missing token %q", ua, want)
				}
			}
		})
	}
}
