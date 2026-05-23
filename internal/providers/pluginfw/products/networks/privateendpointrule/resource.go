package privateendpointrule

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	resourceName                            = "mws_ncc_private_endpoint_rule"
	defaultPrivateEndpointRuleCreateTimeout = 30 * time.Minute
)

func ResourcePrivateEndpointRule() resource.Resource {
	return &resourcePrivateEndpointRule{}
}

type resourcePrivateEndpointRule struct {
	client *common.DatabricksClient
}

var (
	_ resource.ResourceWithConfigure   = &resourcePrivateEndpointRule{}
	_ resource.ResourceWithImportState = &resourcePrivateEndpointRule{}
)

func (r *resourcePrivateEndpointRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (r *resourcePrivateEndpointRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resourceSchema()
}

func (r *resourcePrivateEndpointRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if r.client == nil {
		r.client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *resourcePrivateEndpointRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	var plan model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, d := plan.toCreateRequest(ctx)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	acc, err := r.client.AccountClient()
	if err != nil {
		resp.Diagnostics.AddError("failed to get account client", err.Error())
		return
	}

	rule, err := acc.NetworkConnectivity.CreatePrivateEndpointRule(ctx, *createReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to create private endpoint rule", err.Error())
		return
	}

	// Pack the ID and persist the initial state before polling so that a
	// polling failure or context cancellation still leaves a Read-able
	// resource that `terraform destroy` can clean up.
	resp.Diagnostics.Append(plan.fromAPI(ctx, rule)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.NetworkConnectivityConfigId = types.StringValue(createReq.NetworkConnectivityConfigId)
	plan.ID = types.StringValue(packID(createReq.NetworkConnectivityConfigId, rule.RuleId))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	final, err := waitForPrivateEndpointRuleCreate(ctx, acc, createReq.NetworkConnectivityConfigId, rule.RuleId, defaultPrivateEndpointRuleCreateTimeout)
	if err != nil {
		resp.Diagnostics.AddError("failed waiting for private endpoint rule provisioning", err.Error())
		return
	}
	resp.Diagnostics.Append(plan.fromAPI(ctx, final)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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

	acc, err := r.client.AccountClient()
	if err != nil {
		resp.Diagnostics.AddError("failed to get account client", err.Error())
		return
	}

	rule, err := acc.NetworkConnectivity.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to read private endpoint rule", err.Error())
		return
	}

	resp.Diagnostics.Append(state.fromAPI(ctx, rule)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.NetworkConnectivityConfigId = types.StringValue(nccId)
	state.ID = types.StringValue(packID(nccId, ruleId))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *resourcePrivateEndpointRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	var plan model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, mask, d := plan.toUpdateRequest(ctx, state)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	acc, err := r.client.AccountClient()
	if err != nil {
		resp.Diagnostics.AddError("failed to get account client", err.Error())
		return
	}

	// No update_mask means no updatable field actually changed; nothing to do.
	// The server's update handler does not touch connection_state, so no
	// polling is needed regardless of which fields were changed.
	if len(mask) > 0 {
		updated, err := acc.NetworkConnectivity.UpdatePrivateEndpointRule(ctx, *updateReq)
		if err != nil {
			resp.Diagnostics.AddError("failed to update private endpoint rule", err.Error())
			return
		}
		resp.Diagnostics.Append(plan.fromAPI(ctx, updated)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
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

	acc, err := r.client.AccountClient()
	if err != nil {
		resp.Diagnostics.AddError("failed to get account client", err.Error())
		return
	}

	_, err = acc.NetworkConnectivity.DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete private endpoint rule", err.Error())
		return
	}
}

func (r *resourcePrivateEndpointRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	nccId, ruleId, err := unpackID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("invalid import id", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), packID(nccId, ruleId))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_connectivity_config_id"), nccId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("rule_id"), ruleId)...)
}

// shouldRetry mirrors the helper in app/resource_app.go: continue while the
// retrier's wrapped error is non-halt.
func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(*retries.Err)
	if !ok || e == nil {
		return false
	}
	return !e.Halt
}

// waitForPrivateEndpointRuleCreate polls Get until the rule leaves the CREATING
// state. PENDING and ESTABLISHED are success terminal states; at that point
// vpc_endpoint_id (AWS) / endpoint_name (Azure) / gcp_endpoint (GCP) are
// populated. CREATE_FAILED is terminal failure and surfaces ErrorMessage.
// REJECTED / DISCONNECTED / EXPIRED are not expected on a fresh Create but are
// treated as terminal failures defensively.
func waitForPrivateEndpointRuleCreate(ctx context.Context, acc *databricks.AccountClient, nccId, ruleId string, timeout time.Duration) (*settings.NccPrivateEndpointRule, error) {
	retrier := retries.New[settings.NccPrivateEndpointRule](retries.WithTimeout(timeout), retries.WithRetryFunc(shouldRetry))
	return retrier.Run(ctx, func(ctx context.Context) (*settings.NccPrivateEndpointRule, error) {
		rule, err := acc.NetworkConnectivity.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
		if err != nil {
			return nil, retries.Halt(err)
		}
		switch rule.ConnectionState {
		case settings.NccPrivateEndpointRulePrivateLinkConnectionStatePending,
			settings.NccPrivateEndpointRulePrivateLinkConnectionStateEstablished:
			return rule, nil
		case settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreateFailed:
			return nil, retries.Halt(fmt.Errorf("private endpoint rule %s creation failed: %s", ruleId, rule.ErrorMessage))
		case settings.NccPrivateEndpointRulePrivateLinkConnectionStateRejected,
			settings.NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected,
			settings.NccPrivateEndpointRulePrivateLinkConnectionStateExpired:
			return nil, retries.Halt(fmt.Errorf("private endpoint rule %s reached unexpected terminal state %s", ruleId, rule.ConnectionState))
		default:
			return nil, retries.Continues(fmt.Sprintf("state %s", rule.ConnectionState))
		}
	})
}
