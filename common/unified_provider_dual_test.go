package common

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/go-cty/cty"
	ctymsgpack "github.com/hashicorp/go-cty/cty/msgpack"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// dualGrpcHarness drives a dual resource (no Update, like databricks_group_member)
// through the real SDKv2 gRPC server so that plans are produced exactly as
// Terraform core would see them.
type dualGrpcHarness struct {
	t   *testing.T
	srv tfprotov5.ProviderServer
	ty  cty.Type
}

func newDualGrpcHarness(t *testing.T, customizeDiff func(context.Context, *schema.ResourceDiff, *DatabricksClient) error) *dualGrpcHarness {
	r := Resource{
		IsDual:        true,
		Schema:        AddApiField(map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}}),
		CustomizeDiff: customizeDiff,
		Read:          func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error { return nil },
		Delete:        func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error { return nil },
		Create: func(ctx context.Context, d *schema.ResourceData, c *DatabricksClient) error {
			d.SetId("abc")
			return nil
		},
	}
	AddNamespaceInSchema(r.Schema)
	NamespaceCustomizeSchemaMap(r.Schema)
	res := r.ToResource()
	c := &DatabricksClient{DatabricksClient: &client.DatabricksClient{Config: &config.Config{
		Host:      "https://accounts.cloud.databricks.com",
		AccountID: "abc",
		Token:     "x",
	}}}
	p := &schema.Provider{ResourcesMap: map[string]*schema.Resource{"databricks_dual": res}}
	p.SetMeta(c)
	return &dualGrpcHarness{t: t, srv: schema.NewGRPCProviderServer(p), ty: res.CoreConfigSchema().ImpliedType()}
}

func (h *dualGrpcHarness) enc(v cty.Value) *tfprotov5.DynamicValue {
	b, err := ctymsgpack.Marshal(v, h.ty)
	require.NoError(h.t, err)
	return &tfprotov5.DynamicValue{MsgPack: b}
}

func (h *dualGrpcHarness) dec(dv *tfprotov5.DynamicValue) cty.Value {
	v, err := ctymsgpack.Unmarshal(dv.MsgPack, h.ty)
	require.NoError(h.t, err)
	return v
}

func (h *dualGrpcHarness) providerConfigType() cty.Type {
	return h.ty.AttributeType("provider_config")
}

// config builds the resource config as Terraform core sends it: no
// provider_config block yields an empty list, not null.
func (h *dualGrpcHarness) config() cty.Value {
	m := map[string]cty.Value{}
	for k, at := range h.ty.AttributeTypes() {
		m[k] = cty.NullVal(at)
	}
	m["name"] = cty.StringVal("n")
	m["provider_config"] = cty.ListValEmpty(h.providerConfigType().ElementType())
	return cty.ObjectVal(m)
}

func (h *dualGrpcHarness) plan(prior, proposed cty.Value) cty.Value {
	resp, err := h.srv.PlanResourceChange(context.Background(), &tfprotov5.PlanResourceChangeRequest{
		TypeName:         "databricks_dual",
		PriorState:       h.enc(prior),
		ProposedNewState: h.enc(proposed),
		Config:           h.enc(h.config()),
	})
	require.NoError(h.t, err)
	require.Empty(h.t, resp.Diagnostics)
	return h.dec(resp.PlannedState)
}

func (h *dualGrpcHarness) apply(planned cty.Value) cty.Value {
	resp, err := h.srv.ApplyResourceChange(context.Background(), &tfprotov5.ApplyResourceChangeRequest{
		TypeName:     "databricks_dual",
		PriorState:   h.enc(cty.NullVal(h.ty)),
		PlannedState: h.enc(planned),
		Config:       h.enc(h.config()),
	})
	require.NoError(h.t, err)
	require.Empty(h.t, resp.Diagnostics)
	return h.dec(resp.NewState)
}

// proposedFrom mimics Terraform core's ProposedNew for a resource whose config
// is unchanged: prior state with the (empty) provider_config block from config.
func (h *dualGrpcHarness) proposedFrom(prior cty.Value) cty.Value {
	m := prior.AsValueMap()
	m["provider_config"] = h.config().GetAttr("provider_config")
	return cty.ObjectVal(m)
}

// Regression test for https://github.com/databricks/terraform-provider-databricks/issues/6169:
// after creating an account-level dual resource, the next plan must be empty.
func TestCustomizeDiffDualResources_AccountLevel_NoPerpetualDiff(t *testing.T) {
	for name, cd := range map[string]func(context.Context, *schema.ResourceDiff, *DatabricksClient) error{
		"ForceNew":   CustomizeDiffDualResources,
		"NoForceNew": CustomizeDiffDualResourcesNoForceNew,
	} {
		t.Run(name, func(t *testing.T) {
			h := newDualGrpcHarness(t, cd)
			null := cty.NullVal(h.ty)
			state := h.apply(h.plan(null, h.config()))
			planned := h.plan(state, h.proposedFrom(state))
			assert.True(t, planned.GetAttr("provider_config").IsKnown(),
				"provider_config must not be planned as known after apply")
			assert.True(t, planned.RawEquals(state), "expected empty plan, got %#v", planned)
		})
	}
}

// State written by earlier provider versions holds a null provider_config.
// Refresh plus plan must not produce a diff for it either.
func TestCustomizeDiffDualResources_AccountLevel_LegacyNullState(t *testing.T) {
	h := newDualGrpcHarness(t, CustomizeDiffDualResources)
	m := h.config().AsValueMap()
	m["id"] = cty.StringVal("abc")
	m["provider_config"] = cty.NullVal(h.providerConfigType())
	legacy := cty.ObjectVal(m)
	planned := h.plan(legacy, h.proposedFrom(legacy))
	assert.True(t, planned.RawEquals(legacy), "expected empty plan, got %#v", planned)
}

func dualDiffWithRawConfig(t *testing.T, c *DatabricksClient, is map[string]string, pc map[string]any) *terraform.InstanceDiff {
	res := newDualResourceForCustomizeDiff()
	cfg := map[string]any{"name": "test"}
	if pc != nil {
		cfg["provider_config"] = []any{pc}
	}
	ctyVal := configMapToCtyValue(cfg)
	m := ctyVal.AsValueMap()
	m["api"] = cty.NullVal(cty.String)
	ctyVal = cty.ObjectVal(m)
	rc := terraform.NewResourceConfigShimmed(ctyVal, res.CoreConfigSchema())
	state := &terraform.InstanceState{ID: "abc", Attributes: is, RawConfig: ctyVal}
	diff, err := res.Diff(context.Background(), state, rc, c)
	require.NoError(t, err)
	return diff
}

// An explicitly configured provider_config at account level is still planned.
func TestCustomizeDiffDualResources_AccountLevel_KeepsConfiguredProviderConfig(t *testing.T) {
	c := &DatabricksClient{DatabricksClient: &client.DatabricksClient{Config: &config.Config{
		Host: "https://accounts.cloud.databricks.com", AccountID: "abc", Token: "x",
	}}}
	diff := dualDiffWithRawConfig(t, c, map[string]string{"id": "abc", "name": "test"},
		map[string]any{"workspace_id": "123"})
	require.NotNil(t, diff)
	assert.Equal(t, "123", diff.Attributes["provider_config.0.workspace_id"].New)
}

// Workspace-level behavior is unchanged: a state without provider_config still
// plans it as computed so the post-Read hook can populate it.
func TestCustomizeDiffDualResources_WorkspaceLevel_Unchanged(t *testing.T) {
	c := &DatabricksClient{DatabricksClient: &client.DatabricksClient{Config: &config.Config{
		Host: "https://workspace.cloud.databricks.com", Token: "x",
	}}}
	diff := dualDiffWithRawConfig(t, c, map[string]string{"id": "abc", "name": "test"}, nil)
	require.NotNil(t, diff)
	assert.True(t, diff.Attributes["provider_config.#"].NewComputed)
}
