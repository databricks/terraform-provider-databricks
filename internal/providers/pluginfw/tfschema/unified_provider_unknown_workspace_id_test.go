package tfschema

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// accountProviderClient stands in for an account-level provider without workspace_id,
// whose fallback workspace ID lookup fails. It counts the lookups.
type accountProviderClient struct {
	lookups int
}

func (c *accountProviderClient) GetProviderWorkspaceID() string {
	return ""
}

func (c *accountProviderClient) CurrentWorkspaceID(ctx context.Context) (int64, error) {
	c.lookups++
	return 0, errors.New("an account host has no current workspace")
}

func (c *accountProviderClient) ValidateWorkspaceAccess(ctx context.Context, workspaceID string) diag.Diagnostics {
	return nil
}

var unknownTestProviderConfigType = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{"workspace_id": tftypes.String},
}

var unknownTestRootType = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{"provider_config": unknownTestProviderConfigType},
}

var unknownTestSchema = resourceschema.Schema{
	Attributes: map[string]resourceschema.Attribute{
		"provider_config": resourceschema.SingleNestedAttribute{
			Optional: true,
			Computed: true,
			Attributes: map[string]resourceschema.Attribute{
				"workspace_id": resourceschema.StringAttribute{Optional: true, Computed: true},
			},
		},
	},
}

func unknownTestResource(providerConfig tftypes.Value) tftypes.Value {
	return tftypes.NewValue(unknownTestRootType, map[string]tftypes.Value{
		"provider_config": providerConfig,
	})
}

func unknownTestProviderConfig(workspaceID any) tftypes.Value {
	return tftypes.NewValue(unknownTestProviderConfigType, map[string]tftypes.Value{
		"workspace_id": tftypes.NewValue(tftypes.String, workspaceID),
	})
}

func TestWorkspaceDriftDetection_UnknownWorkspaceID(t *testing.T) {
	existing := unknownTestResource(unknownTestProviderConfig("12345"))
	noState := tftypes.NewValue(unknownTestRootType, nil)
	unknownWorkspaceID := unknownTestResource(unknownTestProviderConfig(tftypes.UnknownValue))
	unknownProviderConfig := unknownTestResource(
		tftypes.NewValue(unknownTestProviderConfigType, tftypes.UnknownValue))
	tests := []struct {
		name            string
		state           tftypes.Value
		config          tftypes.Value
		wantReplacement bool
	}{
		{
			name:   "create with unknown workspace_id",
			state:  noState,
			config: unknownWorkspaceID,
		},
		{
			name:   "create with unknown provider_config",
			state:  noState,
			config: unknownProviderConfig,
		},
		{
			name:            "unknown workspace_id replaces an existing resource",
			state:           existing,
			config:          unknownWorkspaceID,
			wantReplacement: true,
		},
		{
			name:            "unknown provider_config replaces an existing resource",
			state:           existing,
			config:          unknownProviderConfig,
			wantReplacement: true,
		},
		{
			name:   "unknown workspace_id without a prior workspace ID",
			state:  unknownTestResource(tftypes.NewValue(unknownTestProviderConfigType, nil)),
			config: unknownWorkspaceID,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			client := &accountProviderClient{}
			// The plan keeps the configured value, as ProviderConfigPlanModifier does.
			plan := tfsdk.Plan{Schema: unknownTestSchema, Raw: tc.config}
			req := resource.ModifyPlanRequest{
				State:  tfsdk.State{Schema: unknownTestSchema, Raw: tc.state},
				Plan:   plan,
				Config: tfsdk.Config{Schema: unknownTestSchema, Raw: tc.config},
			}
			resp := &resource.ModifyPlanResponse{Plan: plan}

			WorkspaceDriftDetection(ctx, client, req, resp)

			require.False(t, resp.Diagnostics.HasError(), "diagnostics: %v", resp.Diagnostics)
			assert.Zero(t, client.lookups, "no fallback workspace ID should be looked up")
			wsIDPath := path.Root("provider_config").AtName("workspace_id")
			assert.Equal(t, tc.wantReplacement, resp.RequiresReplace.Contains(wsIDPath),
				"RequiresReplace = %v", resp.RequiresReplace)
			assert.True(t, resp.Plan.Raw.Equal(tc.config),
				"the planned value should stay unknown, got %v", resp.Plan.Raw)
		})
	}
}

func TestProviderConfigPlanModifier_UnknownConfig(t *testing.T) {
	ctx := context.Background()
	attrTypes := ProviderConfig{}.Type(ctx).(types.ObjectType).AttrTypes
	stored := ProviderConfig{WorkspaceID: types.StringValue("12345")}.ToObjectValue(ctx)
	unknown := types.ObjectUnknown(attrTypes)
	unknownWorkspaceID := ProviderConfig{WorkspaceID: types.StringUnknown()}.ToObjectValue(ctx)
	null := types.ObjectNull(attrTypes)
	tests := []struct {
		name   string
		config types.Object
		state  types.Object
		want   types.Object
	}{
		{name: "create with unknown provider_config", config: unknown, state: null, want: unknown},
		{name: "update with unknown provider_config", config: unknown, state: stored, want: unknown},
		{
			name:   "update with unknown workspace_id",
			config: unknownWorkspaceID,
			state:  stored,
			want:   unknownWorkspaceID,
		},
		{name: "create without provider_config", config: null, state: null, want: unknown},
		{name: "update without provider_config", config: null, state: stored, want: stored},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stateRaw := tftypes.NewValue(unknownTestRootType, nil)
			if !tc.state.IsNull() {
				stateRaw = unknownTestResource(unknownTestProviderConfig("12345"))
			}
			req := planmodifier.ObjectRequest{
				Path:        path.Root("provider_config"),
				ConfigValue: tc.config,
				StateValue:  tc.state,
				PlanValue:   tc.config,
				State:       tfsdk.State{Schema: unknownTestSchema, Raw: stateRaw},
			}
			resp := &planmodifier.ObjectResponse{PlanValue: req.PlanValue}

			ProviderConfigPlanModifier{}.PlanModifyObject(ctx, req, resp)

			assert.True(t, tc.want.Equal(resp.PlanValue), "planned %v, want %v", resp.PlanValue, tc.want)
		})
	}
}
