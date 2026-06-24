package privateendpointrule

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/retrier"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const resourceName = "mws_ncc_private_endpoint_rule"

func ResourcePrivateEndpointRule() resource.Resource {
	return &resourcePrivateEndpointRule{}
}

// apiClient lists the SDK methods this resource actually uses. Defined
// locally so tests can fake exactly what's needed and new SDK methods
// don't widen the contract.
type apiClient interface {
	CreatePrivateEndpointRule(ctx context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	GetPrivateEndpointRule(ctx context.Context, req settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	UpdatePrivateEndpointRule(ctx context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	DeletePrivateEndpointRule(ctx context.Context, req settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
}

type resourcePrivateEndpointRule struct {
	api apiClient

	// backoff controls the Create polling loop. It is exposed as a field
	// so in-package tests can override it.
	backoff retrier.BackoffPolicy
}

var (
	_ resource.Resource                = &resourcePrivateEndpointRule{}
	_ resource.ResourceWithConfigure   = &resourcePrivateEndpointRule{}
	_ resource.ResourceWithImportState = &resourcePrivateEndpointRule{}
)

func (r *resourcePrivateEndpointRule) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "databricks_" + resourceName
}

func (r *resourcePrivateEndpointRule) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resourceSchema()
}

// Configure resolves the account client from the provider data. The
// framework can call Configure more than once; we re-resolve every time to
// avoid serving a stale client.
func (r *resourcePrivateEndpointRule) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := pluginfwcommon.ConfigureResource(req, resp)
	if client == nil {
		return
	}
	ac, d := client.GetAccountClient()
	resp.Diagnostics.Append(d...)
	if d.HasError() {
		return
	}
	r.api = ac.NetworkConnectivity
}

// createTimeout caps how long Create waits for the rule to leave CREATING.
// Without it, a stuck server-side provisioning would make `terraform apply`
// hang. 30 minutes is well above typical transitions (a few minutes).
var createTimeout = 30 * time.Minute

func (r *resourcePrivateEndpointRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	var plan model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// enabled is Optional+Computed, but CreatePrivateEndpointRule has no enabled
	// field, so the server always creates with its own default. fromAPI would
	// then overwrite a known planned value (e.g. enabled = false) with that
	// default and fail Terraform's post-apply consistency check. Preserve the
	// configured value here; the next Update reconciles the server side. A
	// null/unknown plan (enabled omitted) takes the server value as-is.
	configuredEnabled := plan.Enabled
	applyConfiguredEnabled := func() {
		if !configuredEnabled.IsNull() && !configuredEnabled.IsUnknown() {
			plan.Enabled = configuredEnabled
		}
	}

	apiReq, d := plan.toCreateRequest(ctx)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	pendingRule, err := r.api.CreatePrivateEndpointRule(ctx, *apiReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to create private endpoint rule", err.Error())
		return
	}

	// Persist the created rule before polling so a polling failure still leaves
	// a readable resource that `terraform destroy` can clean up.
	plan.NetworkConnectivityConfigId = types.StringValue(apiReq.NetworkConnectivityConfigId)
	plan.ID = types.StringValue(packID(apiReq.NetworkConnectivityConfigId, pendingRule.RuleId))
	resp.Diagnostics.Append(plan.fromAPI(ctx, pendingRule)...)
	applyConfiguredEnabled()
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Poll until the rule leaves CREATING. PENDING counts as success: the
	// customer still has to approve the connection in their cloud console
	// (typically via the aws/azurerm/google provider in the same config),
	// so we don't wait for ESTABLISHED. The switch below handles the rest.
	rule, err := retrier.Run(ctx, retrier.RetryIf(r.backoff, isStillCreating), func(ctx context.Context) (*settings.NccPrivateEndpointRule, error) {
		return r.api.GetPrivateEndpointRule(ctx, settings.GetPrivateEndpointRuleRequest{
			// Use the caller-supplied NCC id, not the value echoed in the
			// create response, so polling does not depend on the server
			// round-tripping the field. The rule id is server-assigned, so it
			// must come from the create response.
			NetworkConnectivityConfigId: apiReq.NetworkConnectivityConfigId,
			PrivateEndpointRuleId:       pendingRule.RuleId,
		})
	})
	if err != nil {
		resp.Diagnostics.AddError("failed waiting for private endpoint rule provisioning", err.Error())
		return
	}

	switch rule.ConnectionState {
	case settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreateFailed:
		// Failure case: surface the server's error message.
		resp.Diagnostics.AddError("failed waiting for private endpoint rule provisioning", rule.ErrorMessage)
		return
	case settings.NccPrivateEndpointRulePrivateLinkConnectionStateRejected,
		settings.NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected,
		settings.NccPrivateEndpointRulePrivateLinkConnectionStateExpired:
		// The rule exists but is not usable. Surface it instead of pretending
		// the apply succeeded.
		resp.Diagnostics.AddError(
			"private endpoint rule reached unexpected state",
			fmt.Sprintf("expected PENDING or ESTABLISHED, got %q", rule.ConnectionState),
		)
		return
	default:
		// PENDING and ESTABLISHED are the expected success states (the rule is
		// usable or awaiting customer approval). The loop already excluded
		// CREATING, and a timed-out poll returned an error above, so anything
		// reaching here is one of those two or an enum value the vendored SDK
		// predates. Treat unknown future states as success rather than failing
		// the apply on a value we cannot classify; a later read surfaces it.
		resp.Diagnostics.Append(plan.fromAPI(ctx, rule)...)
		applyConfiguredEnabled()
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
		return
	}
}

// isStillCreating is the retrier predicate: keep polling only while the
// rule is in CREATING. Any error halts the loop; the SDK already retries
// transient network errors.
func isStillCreating(rule *settings.NccPrivateEndpointRule, err error) bool {
	return err == nil && rule.ConnectionState == settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreating
}

func (r *resourcePrivateEndpointRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nccId, ruleId, err := unpackID(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("invalid resource id", err.Error())
		return
	}

	rule, err := r.api.GetPrivateEndpointRule(ctx, settings.GetPrivateEndpointRuleRequest{
		NetworkConnectivityConfigId: nccId,
		PrivateEndpointRuleId:       ruleId,
	})
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to read private endpoint rule", err.Error())
		return
	}

	state.NetworkConnectivityConfigId = types.StringValue(nccId)
	state.ID = types.StringValue(packID(nccId, ruleId))
	resp.Diagnostics.Append(state.fromAPI(ctx, rule)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *resourcePrivateEndpointRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan, state model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq, d := plan.toUpdateRequest(ctx, state)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Empty update_mask means no updatable field changed; skip the API call.
	if apiReq.UpdateMask != "" {
		if _, err := r.api.UpdatePrivateEndpointRule(ctx, *apiReq); err != nil {
			resp.Diagnostics.AddError("failed to update private endpoint rule", err.Error())
			return
		}
	}
	// Mirror SDKv2: do not fold the update response back into state. Server-set
	// computed fields (e.g. updated_time) change on update, but the plan keeps
	// their prior values via UseStateForUnknown; overwriting them here would
	// trip Terraform's "provider produced inconsistent result after apply"
	// check. The next read reconciles them.
	plan.NetworkConnectivityConfigId = state.NetworkConnectivityConfigId
	plan.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourcePrivateEndpointRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nccId, ruleId, err := unpackID(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("invalid resource id", err.Error())
		return
	}

	_, err = r.api.DeletePrivateEndpointRule(ctx, settings.DeletePrivateEndpointRuleRequest{
		NetworkConnectivityConfigId: nccId,
		PrivateEndpointRuleId:       ruleId,
	})
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete private endpoint rule", err.Error())
	}
}

func (r *resourcePrivateEndpointRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	nccId, ruleId, err := unpackID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("invalid import id", err.Error())
		return
	}
	// Seed just the ID fields; the framework calls Read next to fill in
	// the rest. List fields need their element type set up front, which
	// is what emptyModel handles.
	m := emptyModel()
	m.ID = types.StringValue(packID(nccId, ruleId))
	m.NetworkConnectivityConfigId = types.StringValue(nccId)
	m.RuleId = types.StringValue(ruleId)
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
