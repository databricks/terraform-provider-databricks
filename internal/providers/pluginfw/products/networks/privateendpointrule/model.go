package privateendpointrule

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// model is the Terraform-side representation of an NCC private endpoint rule.
// It is intentionally distinct from the SDK's NccPrivateEndpointRule: the
// translation layer below is the anti-corruption boundary between Terraform's
// Null/Unknown-aware types and the SDK's native Go types.
type model struct {
	ID                          types.String       `tfsdk:"id"`
	NetworkConnectivityConfigId types.String       `tfsdk:"network_connectivity_config_id"`
	RuleId                      types.String       `tfsdk:"rule_id"`
	AccountId                   types.String       `tfsdk:"account_id"`
	EndpointService             types.String       `tfsdk:"endpoint_service"`
	GroupId                     types.String       `tfsdk:"group_id"`
	ResourceId                  types.String       `tfsdk:"resource_id"`
	DomainNames                 types.List         `tfsdk:"domain_names"`
	ResourceNames               types.List         `tfsdk:"resource_names"`
	Enabled                     types.Bool         `tfsdk:"enabled"`
	EndpointName                types.String       `tfsdk:"endpoint_name"`
	VpcEndpointId               types.String       `tfsdk:"vpc_endpoint_id"`
	ConnectionState             types.String       `tfsdk:"connection_state"`
	CreationTime                types.Int64        `tfsdk:"creation_time"`
	UpdatedTime                 types.Int64        `tfsdk:"updated_time"`
	Deactivated                 types.Bool         `tfsdk:"deactivated"`
	DeactivatedAt               types.Int64        `tfsdk:"deactivated_at"`
	ErrorMessage                types.String       `tfsdk:"error_message"`
	GcpEndpoint                 []gcpEndpointModel `tfsdk:"gcp_endpoint"`
}

type gcpEndpointModel struct {
	PscEndpointUri    types.String `tfsdk:"psc_endpoint_uri"`
	ServiceAttachment types.String `tfsdk:"service_attachment"`
}

// emptyModel returns a model with collection-typed fields initialized to
// their typed null value. PF's reflection requires lists to declare their
// element type at conversion time; the zero types.List has DynamicPseudoType
// and fails the round-trip through tfsdk.State.Set / tfsdk.Plan.Set. Code
// paths that construct a model from scratch (rather than reading it from
// req.Plan or req.State, which populate all fields from the schema) must
// start here.
func emptyModel() model {
	return model{
		DomainNames:   types.ListNull(types.StringType),
		ResourceNames: types.ListNull(types.StringType),
	}
}

func packID(nccId, ruleId string) string {
	return nccId + "/" + ruleId
}

func unpackID(composite string) (nccId, ruleId string, err error) {
	parts := strings.SplitN(composite, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("expected ID in the form <network_connectivity_config_id>/<rule_id>, got %q", composite)
	}
	return parts[0], parts[1], nil
}

// stringsFromList decodes a Plugin Framework list of strings into a Go slice.
// Null and Unknown lists become a nil slice; the SDK treats nil and an empty
// list identically (omitted from the JSON body) so callers don't need to
// distinguish.
func stringsFromList(ctx context.Context, list types.List) ([]string, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return nil, nil
	}
	var out []string
	d := list.ElementsAs(ctx, &out, false)
	return out, d
}

// toCreateRequest builds the SDK create request from the plan model. Only
// user-settable fields are included; computed fields (rule_id, vpc_endpoint_id,
// connection_state, ...) are populated by the server and applied in fromAPI.
func (m *model) toCreateRequest(ctx context.Context) (*settings.CreatePrivateEndpointRuleRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	domains, d := stringsFromList(ctx, m.DomainNames)
	diags.Append(d...)
	resources, d := stringsFromList(ctx, m.ResourceNames)
	diags.Append(d...)
	return &settings.CreatePrivateEndpointRuleRequest{
		NetworkConnectivityConfigId: m.NetworkConnectivityConfigId.ValueString(),
		PrivateEndpointRule: settings.CreatePrivateEndpointRule{
			EndpointService: m.EndpointService.ValueString(),
			GroupId:         m.GroupId.ValueString(),
			ResourceId:      m.ResourceId.ValueString(),
			DomainNames:     domains,
			ResourceNames:   resources,
			GcpEndpoint:     gcpEndpointToAPI(m.GcpEndpoint),
		},
	}, diags
}

// gcpEndpointToAPI is the inverse of gcpEndpointFromAPI: it converts the
// size-0-or-1 slice the schema exposes back into the SDK's pointer-to-struct.
// The schema enforces SizeAtMost(1), so callers never see more than one
// element.
func gcpEndpointToAPI(gcp []gcpEndpointModel) *settings.GcpEndpoint {
	if len(gcp) == 0 {
		return nil
	}
	return &settings.GcpEndpoint{
		PscEndpointUri:    gcp[0].PscEndpointUri.ValueString(),
		ServiceAttachment: gcp[0].ServiceAttachment.ValueString(),
	}
}

// toUpdateRequest builds the SDK update request with an update_mask of fields
// that changed between prev (state) and m (plan). An empty UpdateMask means
// nothing updatable changed; the caller can skip the API call entirely.
func (m *model) toUpdateRequest(ctx context.Context, prev model) (*settings.UpdateNccPrivateEndpointRuleRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	nccId, ruleId, err := unpackID(m.ID.ValueString())
	if err != nil {
		diags.AddError("invalid resource id", err.Error())
		return nil, diags
	}
	var mask []string
	rule := settings.UpdatePrivateEndpointRule{}
	if !m.Enabled.Equal(prev.Enabled) {
		mask = append(mask, "enabled")
		rule.Enabled = m.Enabled.ValueBool()
	}
	if !m.DomainNames.Equal(prev.DomainNames) {
		mask = append(mask, "domain_names")
		names, d := stringsFromList(ctx, m.DomainNames)
		diags.Append(d...)
		rule.DomainNames = names
	}
	if !m.ResourceNames.Equal(prev.ResourceNames) {
		mask = append(mask, "resource_names")
		names, d := stringsFromList(ctx, m.ResourceNames)
		diags.Append(d...)
		rule.ResourceNames = names
	}
	return &settings.UpdateNccPrivateEndpointRuleRequest{
		NetworkConnectivityConfigId: nccId,
		PrivateEndpointRuleId:       ruleId,
		PrivateEndpointRule:         rule,
		UpdateMask:                  strings.Join(mask, ","),
	}, diags
}

// fromAPI populates the model from an SDK NccPrivateEndpointRule response.
// network_connectivity_config_id is a path parameter, not on the response
// struct; the caller is expected to set it.
func (m *model) fromAPI(ctx context.Context, rule *settings.NccPrivateEndpointRule) diag.Diagnostics {
	var diags diag.Diagnostics
	m.RuleId = types.StringValue(rule.RuleId)
	m.AccountId = types.StringValue(rule.AccountId)
	m.EndpointService = types.StringValue(rule.EndpointService)
	m.GroupId = types.StringValue(rule.GroupId)
	m.ResourceId = types.StringValue(rule.ResourceId)
	m.Enabled = types.BoolValue(rule.Enabled)
	m.EndpointName = types.StringValue(rule.EndpointName)
	m.VpcEndpointId = types.StringValue(rule.VpcEndpointId)
	m.ConnectionState = types.StringValue(string(rule.ConnectionState))
	m.CreationTime = types.Int64Value(rule.CreationTime)
	m.UpdatedTime = types.Int64Value(rule.UpdatedTime)
	m.Deactivated = types.BoolValue(rule.Deactivated)
	m.DeactivatedAt = types.Int64Value(rule.DeactivatedAt)
	m.ErrorMessage = types.StringValue(rule.ErrorMessage)

	domainList, d := types.ListValueFrom(ctx, types.StringType, rule.DomainNames)
	diags.Append(d...)
	m.DomainNames = domainList

	resourceList, d := types.ListValueFrom(ctx, types.StringType, rule.ResourceNames)
	diags.Append(d...)
	m.ResourceNames = resourceList

	m.GcpEndpoint = gcpEndpointFromAPI(rule.GcpEndpoint)
	return diags
}

// gcpEndpointFromAPI converts the SDK's pointer-to-struct into the
// size-0-or-1 slice the schema expects. Nil here causes the framework to omit
// the block from state, matching how the SDKv2 implementation treated the
// field.
func gcpEndpointFromAPI(gcp *settings.GcpEndpoint) []gcpEndpointModel {
	if gcp == nil {
		return nil
	}
	return []gcpEndpointModel{{
		PscEndpointUri:    types.StringValue(gcp.PscEndpointUri),
		ServiceAttachment: types.StringValue(gcp.ServiceAttachment),
	}}
}
