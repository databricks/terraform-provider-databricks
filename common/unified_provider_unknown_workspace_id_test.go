package common

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type namespaceCustomizeDiffFunc func(context.Context, *schema.ResourceDiff, *DatabricksClient) error

// diffWithUnknownWorkspaceID runs the plan diff, as PlanResourceChange does, of a resource
// whose config sets provider_config.workspace_id to a value that is unknown at plan time.
// A nil state plans a create. The provider's server rejects every request, so a lookup of
// a fallback workspace ID fails the diff.
func diffWithUnknownWorkspaceID(t *testing.T, hostType config.HostType, providerWorkspaceID string,
	customizeDiff namespaceCustomizeDiffFunc, state map[string]string) (*terraform.InstanceDiff, error) {
	t.Helper()
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		w.WriteHeader(http.StatusNotFound)
	}))
	t.Cleanup(func() {
		server.Close()
		assert.Zero(t, requests.Load(), "the diff should not call the API")
	})
	cfg := &config.Config{
		Host:        server.URL,
		Token:       "test-token",
		AccountID:   "test-account",
		WorkspaceID: providerWorkspaceID,
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
			return &config.HostMetadata{HostType: hostType}, nil
		},
	}
	require.NoError(t, cfg.EnsureResolved())
	require.Equal(t, hostType, cfg.HostType())
	c := &DatabricksClient{DatabricksClient: &client.DatabricksClient{Config: cfg}}

	testSchema := map[string]*schema.Schema{
		"name": {Type: schema.TypeString, Required: true},
	}
	AddNamespaceInSchema(testSchema)
	r := &schema.Resource{
		Schema: testSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m any) error {
			return customizeDiff(ctx, d, m.(*DatabricksClient))
		},
	}
	configVal := cty.ObjectVal(map[string]cty.Value{
		"name": cty.StringVal("test"),
		"provider_config": cty.ListVal([]cty.Value{cty.ObjectVal(map[string]cty.Value{
			"workspace_id": cty.UnknownVal(cty.String),
		})}),
	})
	rc := terraform.NewResourceConfigShimmed(configVal, schema.InternalMap(testSchema).CoreConfigSchema())
	// As in PlanResourceChange, the prior state carries the raw config that
	// GetRawConfigAt reads.
	is := &terraform.InstanceState{Attributes: state, RawConfig: configVal}
	if state != nil {
		is.ID = "test-resource-id"
	}
	return r.SimpleDiff(context.Background(), is, rc, c)
}

func TestNamespaceCustomizeDiff_UnknownWorkspaceID(t *testing.T) {
	existing := map[string]string{
		"name":                           "test",
		"provider_config.#":              "1",
		"provider_config.0.workspace_id": "100",
	}
	tests := []struct {
		name                string
		hostType            config.HostType
		providerWorkspaceID string
		customizeDiff       namespaceCustomizeDiffFunc
		state               map[string]string
		wantReplace         bool
	}{
		{
			name:          "create",
			hostType:      config.AccountHost,
			customizeDiff: NamespaceCustomizeDiff,
		},
		{
			name:          "workspace-scoped resource on an account host is replaced",
			hostType:      config.AccountHost,
			customizeDiff: NamespaceCustomizeDiff,
			state:         existing,
			wantReplace:   true,
		},
		{
			name:                "workspace-scoped resource on an account host with the provider's workspace_id is replaced",
			hostType:            config.AccountHost,
			providerWorkspaceID: "100",
			customizeDiff:       NamespaceCustomizeDiff,
			state:               existing,
			wantReplace:         true,
		},
		{
			name:          "workspace-scoped resource on a unified host is replaced",
			hostType:      config.UnifiedHost,
			customizeDiff: NamespaceCustomizeDiff,
			state:         existing,
			wantReplace:   true,
		},
		{
			name:          "workspace-scoped resource on a workspace host is updated in place",
			hostType:      config.WorkspaceHost,
			customizeDiff: NamespaceCustomizeDiff,
			state:         existing,
		},
		{
			name:          "metastore-scoped resource on an account host is updated in place",
			hostType:      config.AccountHost,
			customizeDiff: NamespaceCustomizeDiffNoForceNew,
			state:         existing,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diff, err := diffWithUnknownWorkspaceID(t, tc.hostType, tc.providerWorkspaceID,
				tc.customizeDiff, tc.state)
			require.NoError(t, err)
			require.NotNil(t, diff)
			assert.Equal(t, tc.wantReplace, diff.RequiresNew())
			if tc.wantReplace {
				// Terraform plans the replacement's create separately, which leaves
				// workspace_id unknown.
				attr, ok := diff.Attributes["provider_config.#"]
				require.True(t, ok, "provider_config should be in the diff")
				assert.True(t, attr.NewComputed, "provider_config should be unknown")
				assert.True(t, attr.RequiresNew, "provider_config should force a replacement")
				return
			}
			attr, ok := diff.Attributes[workspaceIDSchemaKey]
			require.True(t, ok, "workspace_id should be in the diff")
			assert.True(t, attr.NewComputed, "workspace_id should stay unknown")
			assert.False(t, attr.RequiresNew)
		})
	}
}
