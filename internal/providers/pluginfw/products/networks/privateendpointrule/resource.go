package privateendpointrule

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/provider"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/telemetry"
	"github.com/databricks/terraform-provider-databricks/internal/retrier"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const resourceName = "mws_ncc_private_endpoint_rule"

func ResourcePrivateEndpointRule() resource.Resource {
	return &resourcePrivateEndpointRule{}
}

// api names exactly the SDK calls this resource depends on. Defining it
// locally (rather than importing settings.NetworkConnectivityInterface) follows
// Go's "interfaces belong to the consumer" idiom: future additions to the SDK
// can't leak into this resource's contract, and the fake in tests needs to
// implement only what's actually used.
type apiClient interface {
	CreatePrivateEndpointRule(ctx context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	GetPrivateEndpointRule(ctx context.Context, req settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	UpdatePrivateEndpointRule(ctx context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	DeletePrivateEndpointRule(ctx context.Context, req settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
}

type resourcePrivateEndpointRule struct {
	api apiClient

	// backoff is the retry backoff policy for the resource.
	// It is exposed as a field to ease in-package testing.
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

// Configure resolves the account client from the provider data. PF may call
// this more than once on the same resource value if the underlying provider
// data is updated; we re-resolve on every call so we never serve a stale
// client.
func (r *resourcePrivateEndpointRule) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	ac, d := provider.AccountClient(req.ProviderData)
	resp.Diagnostics.Append(d...)
	if ac == nil {
		// ProviderData not yet wired (pre-Configure phase) or a real error
		// already appended above. Either way, bail.
		return
	}
	r.api = ac.NetworkConnectivity
}

// createTimeout caps how long Create waits for the rule to leave CREATING.
// Plugin Framework gives a request ctx with no deadline (unlike SDKv2, which
// defaulted CRUD operations to 20 minutes), so without this ceiling a stalled
// reconciler would make `terraform apply` hang until SIGINT. 30 minutes is
// generous against observed transitions from CREATING to PENDING, which are
// typically under a few minutes.
var createTimeout = 30 * time.Minute

func (r *resourcePrivateEndpointRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = telemetry.WithResource(ctx, resourceName)

	// Bound the operation; without this the polling loop below could hang
	// indefinitely (PF's request ctx carries no deadline). The rule is
	// persisted to state before polling starts, so a deadline hit still
	// leaves a destroyable resource behind.
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	var plan model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq, d := plan.toCreateRequest(ctx)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	// CreatePrivateEndpointRule is not idempotent; retrying on transient errors
	// would create duplicate rules. Let the SDK's transport-level retry handle
	// connection-level flakes; surface anything else as a hard failure.
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Poll until the rule leaves CREATING. PENDING is the success terminal
	// for the server's side of the handshake; the transition to ESTABLISHED
	// needs the customer to approve the connection in their cloud console
	// (out of band; usually wired via the aws/azurerm/google provider in
	// the same config), so the loop intentionally does not wait for it.
	// The post-loop switch classifies the terminal state.
	rule, err := retrier.Run(ctx, retrier.RetryIf(r.backoff, isStillCreating), func(ctx context.Context) (*settings.NccPrivateEndpointRule, error) {
		return r.api.GetPrivateEndpointRule(ctx, settings.GetPrivateEndpointRuleRequest{
			NetworkConnectivityConfigId: apiReq.NetworkConnectivityConfigId,
			PrivateEndpointRuleId:       pendingRule.RuleId,
		})
	})
	if err != nil {
		resp.Diagnostics.AddError("failed waiting for private endpoint rule provisioning", err.Error())
		return
	}
	switch rule.ConnectionState {
	case settings.NccPrivateEndpointRulePrivateLinkConnectionStatePending,
		settings.NccPrivateEndpointRulePrivateLinkConnectionStateEstablished:
		// Server-side handshake done; rule is usable (ESTABLISHED) or
		// waiting on customer approval (PENDING). Fall through to persist.
	case settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreateFailed:
		resp.Diagnostics.AddError("failed waiting for private endpoint rule provisioning", rule.ErrorMessage)
		return
	default:
		// REJECTED, DISCONNECTED, EXPIRED, or any future enum value: the
		// rule exists server-side but is not usable. Surfacing this beats
		// a silent green apply on a broken resource.
		resp.Diagnostics.AddError(
			"private endpoint rule reached unexpected state",
			fmt.Sprintf("expected PENDING or ESTABLISHED, got %q", rule.ConnectionState),
		)
		return
	}
	resp.Diagnostics.Append(plan.fromAPI(ctx, rule)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// isStillCreating is the retrier predicate: retry only while the rule is in
// the single transient state CREATING. Any error from Get halts the retry,
// transport-level errors are meant to be handled by the SDK.
func isStillCreating(rule *settings.NccPrivateEndpointRule, err error) bool {
	return err == nil && rule.ConnectionState == settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreating
}

func (r *resourcePrivateEndpointRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = telemetry.WithResource(ctx, resourceName)

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
	ctx = telemetry.WithResource(ctx, resourceName)

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
		updated, err := r.api.UpdatePrivateEndpointRule(ctx, *apiReq)
		if err != nil {
			resp.Diagnostics.AddError("failed to update private endpoint rule", err.Error())
			return
		}
		resp.Diagnostics.Append(plan.fromAPI(ctx, updated)...)
	}
	plan.NetworkConnectivityConfigId = state.NetworkConnectivityConfigId
	plan.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourcePrivateEndpointRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = telemetry.WithResource(ctx, resourceName)

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
	// Set just enough state for the post-import Read to take over. Other
	// scalar fields default to their null PF value; list-typed fields need
	// their element type explicitly via emptyModel().
	m := emptyModel()
	m.ID = types.StringValue(packID(nccId, ruleId))
	m.NetworkConnectivityConfigId = types.StringValue(nccId)
	m.RuleId = types.StringValue(ruleId)
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
