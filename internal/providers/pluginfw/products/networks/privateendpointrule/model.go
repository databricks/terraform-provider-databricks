package privateendpointrule

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type model struct {
	ID                          types.String      `tfsdk:"id"`
	NetworkConnectivityConfigId types.String      `tfsdk:"network_connectivity_config_id"`
	RuleId                      types.String      `tfsdk:"rule_id"`
	AccountId                   types.String      `tfsdk:"account_id"`
	EndpointService             types.String      `tfsdk:"endpoint_service"`
	GroupId                     types.String      `tfsdk:"group_id"`
	ResourceId                  types.String      `tfsdk:"resource_id"`
	DomainNames                 types.List        `tfsdk:"domain_names"`
	ResourceNames               types.List        `tfsdk:"resource_names"`
	Enabled                     types.Bool        `tfsdk:"enabled"`
	EndpointName                types.String      `tfsdk:"endpoint_name"`
	VpcEndpointId               types.String      `tfsdk:"vpc_endpoint_id"`
	ConnectionState             types.String      `tfsdk:"connection_state"`
	CreationTime                types.Int64       `tfsdk:"creation_time"`
	UpdatedTime                 types.Int64       `tfsdk:"updated_time"`
	Deactivated                 types.Bool        `tfsdk:"deactivated"`
	DeactivatedAt               types.Int64       `tfsdk:"deactivated_at"`
	ErrorMessage                types.String      `tfsdk:"error_message"`
	GcpEndpoint                 []gcpEndpointModel `tfsdk:"gcp_endpoint"`
}

type gcpEndpointModel struct {
	PscEndpointUri    types.String `tfsdk:"psc_endpoint_uri"`
	ServiceAttachment types.String `tfsdk:"service_attachment"`
}

const idSeparator = "/"

func packID(nccId, ruleId string) string {
	return nccId + idSeparator + ruleId
}

func unpackID(composite string) (nccId, ruleId string, err error) {
	parts := strings.SplitN(composite, idSeparator, 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("expected ID in the form %q, got %q", "<network_connectivity_config_id>/<rule_id>", composite)
	}
	return parts[0], parts[1], nil
}

// toCreateRequest converts the plan model into the SDK CreatePrivateEndpointRuleRequest.
// Only user-settable fields are included; computed fields (rule_id, vpc_endpoint_id,
// connection_state, etc.) are populated from the server response in fromAPI.
func (m *model) toCreateRequest(ctx context.Context) (*settings.CreatePrivateEndpointRuleRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	req := &settings.CreatePrivateEndpointRuleRequest{
		NetworkConnectivityConfigId: m.NetworkConnectivityConfigId.ValueString(),
		PrivateEndpointRule: settings.CreatePrivateEndpointRule{
			EndpointService: m.EndpointService.ValueString(),
			GroupId:         m.GroupId.ValueString(),
			ResourceId:      m.ResourceId.ValueString(),
		},
	}
	if !m.DomainNames.IsNull() && !m.DomainNames.IsUnknown() {
		d := m.DomainNames.ElementsAs(ctx, &req.PrivateEndpointRule.DomainNames, false)
		diags.Append(d...)
	}
	if !m.ResourceNames.IsNull() && !m.ResourceNames.IsUnknown() {
		d := m.ResourceNames.ElementsAs(ctx, &req.PrivateEndpointRule.ResourceNames, false)
		diags.Append(d...)
	}
	return req, diags
}

// toUpdateRequest produces the SDK UpdateNccPrivateEndpointRuleRequest along with the
// update_mask of fields that changed between prev (state) and m (plan). Only the
// updatable fields (enabled, domain_names, resource_names) are considered; ForceNew
// fields would have triggered destroy+create at the plan phase.
func (m *model) toUpdateRequest(ctx context.Context, prev model) (*settings.UpdateNccPrivateEndpointRuleRequest, []string, diag.Diagnostics) {
	var diags diag.Diagnostics
	nccId, ruleId, err := unpackID(m.ID.ValueString())
	if err != nil {
		diags.AddError("invalid resource id", err.Error())
		return nil, nil, diags
	}
	var mask []string
	rule := settings.UpdatePrivateEndpointRule{}
	if !m.Enabled.Equal(prev.Enabled) {
		mask = append(mask, "enabled")
		rule.Enabled = m.Enabled.ValueBool()
	}
	if !m.DomainNames.Equal(prev.DomainNames) {
		mask = append(mask, "domain_names")
		var names []string
		if !m.DomainNames.IsNull() && !m.DomainNames.IsUnknown() {
			d := m.DomainNames.ElementsAs(ctx, &names, false)
			diags.Append(d...)
		}
		rule.DomainNames = names
	}
	if !m.ResourceNames.Equal(prev.ResourceNames) {
		mask = append(mask, "resource_names")
		var names []string
		if !m.ResourceNames.IsNull() && !m.ResourceNames.IsUnknown() {
			d := m.ResourceNames.ElementsAs(ctx, &names, false)
			diags.Append(d...)
		}
		rule.ResourceNames = names
	}
	return &settings.UpdateNccPrivateEndpointRuleRequest{
		NetworkConnectivityConfigId: nccId,
		PrivateEndpointRuleId:       ruleId,
		PrivateEndpointRule:         rule,
		UpdateMask:                  strings.Join(mask, ","),
	}, mask, diags
}

// fromAPI populates the model from an SDK NccPrivateEndpointRule response. The
// network_connectivity_config_id is not present on the response struct itself
// (it's a path parameter), so the caller is expected to set it before or after.
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

	if rule.GcpEndpoint != nil {
		g := gcpEndpointModel{}
		g.fromAPI(rule.GcpEndpoint)
		m.GcpEndpoint = []gcpEndpointModel{g}
	} else {
		m.GcpEndpoint = nil
	}
	return diags
}

func (g *gcpEndpointModel) fromAPI(in *settings.GcpEndpoint) {
	g.PscEndpointUri = types.StringValue(in.PscEndpointUri)
	g.ServiceAttachment = types.StringValue(in.ServiceAttachment)
}
