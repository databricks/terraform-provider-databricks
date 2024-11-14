package framework

import (
	"context"

	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func WrapResource(r resource.Resource) resource.Resource {
	var wrapped resource.Resource = r
	if c, ok := wrapped.(resource.ResourceWithConfigure); ok {
		wrapped = Configurer{c}
	}
	if c, ok := wrapped.(resource.ResourceWithConfigValidators); ok {
		wrapped = ConfigValidatorer{c}
	}
	if c, ok := wrapped.(resource.ResourceWithImportState); ok {
		wrapped = StateImporter{c}
	}
	if c, ok := wrapped.(resource.ResourceWithModifyPlan); ok {
		wrapped = PlanModifier{c}
	}
	if c, ok := wrapped.(resource.ResourceWithMoveState); ok {
		wrapped = StateMover{c}
	}
	if c, ok := wrapped.(resource.ResourceWithUpgradeState); ok {
		wrapped = StateUpgrader{c}
	}
	if c, ok := wrapped.(resource.ResourceWithValidateConfig); ok {
		wrapped = ConfigValidator{c}
	}
	return wrapped
}

type Resource struct {
	resource.Resource
}

func configureContext(ctx context.Context, r resource.Resource) context.Context {
	resp := &resource.MetadataResponse{}
	r.Metadata(ctx, resource.MetadataRequest{}, resp)
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resp.TypeName)
	return ctx
}

func (r Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.Resource.Metadata(ctx, req, resp)
}

func (r Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.Resource.Schema(ctx, req, resp)
}

func (r Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = configureContext(ctx, r)
	r.Resource.Create(ctx, req, resp)
}

func (r Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = configureContext(ctx, r)
	r.Resource.Read(ctx, req, resp)
}

func (r Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = configureContext(ctx, r)
	r.Resource.Update(ctx, req, resp)
}

func (r Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = configureContext(ctx, r)
	r.Resource.Delete(ctx, req, resp)
}

type Configurer struct {
	resource.ResourceWithConfigure
}

func (c Configurer) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	ctx = configureContext(ctx, c)
	c.ResourceWithConfigure.Configure(ctx, req, resp)
}

type ConfigValidatorer struct {
	resource.ResourceWithConfigValidators
}

func (c ConfigValidatorer) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	ctx = configureContext(ctx, c)
	return c.ResourceWithConfigValidators.ConfigValidators(ctx)
}

type StateImporter struct {
	resource.ResourceWithImportState
}

func (s StateImporter) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	ctx = configureContext(ctx, s)
	s.ResourceWithImportState.ImportState(ctx, req, resp)
}

type PlanModifier struct {
	resource.ResourceWithModifyPlan
}

func (p PlanModifier) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	ctx = configureContext(ctx, p)
	p.ResourceWithModifyPlan.ModifyPlan(ctx, req, resp)
}

type StateMover struct {
	resource.ResourceWithMoveState
}

func (s StateMover) MoveState(ctx context.Context) []resource.StateMover {
	ctx = configureContext(ctx, s)
	return s.ResourceWithMoveState.MoveState(ctx)
}

type StateUpgrader struct {
	resource.ResourceWithUpgradeState
}

func (s StateUpgrader) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	ctx = configureContext(ctx, s)
	return s.ResourceWithUpgradeState.UpgradeState(ctx)
}

type ConfigValidator struct {
	resource.ResourceWithValidateConfig
}

func (c ConfigValidator) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	ctx = configureContext(ctx, c)
	c.ResourceWithValidateConfig.ValidateConfig(ctx, req, resp)
}
