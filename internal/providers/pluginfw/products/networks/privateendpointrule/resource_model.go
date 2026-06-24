package privateendpointrule

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// model is the Terraform-side representation. It uses Terraform's
// Null/Unknown-aware types; the translation functions below convert to and
// from the SDK's plain Go types.
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

// emptyModel returns a model with list fields set to typed-null. The
// framework rejects a list that has no element type, so code paths that
// build a model by hand (rather than via req.Plan or req.State) must start
// from here.
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

// stringsFromList decodes a list-of-strings attribute into a Go slice.
// Null/unknown become nil, which the SDK serializes the same as an empty
// list.
func stringsFromList(ctx context.Context, list types.List) ([]string, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return nil, nil
	}
	var out []string
	d := list.ElementsAs(ctx, &out, false)
	return out, d
}

// toCreateRequest builds the SDK create body. Only user-settable fields go
// in; server-set fields (rule_id, vpc_endpoint_id, etc.) come back via
// fromAPI.
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

// gcpEndpointToAPI converts the size-0-or-1 slice (schema enforces
// SizeAtMost(1)) back to the SDK's pointer-to-struct.
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

// fromAPI copies server-returned fields into m. The caller separately sets
// NetworkConnectivityConfigId from the URL path; we don't trust it from the
// response so the value stays consistent with how the resource was addressed.
func (m *model) fromAPI(ctx context.Context, rule *settings.NccPrivateEndpointRule) diag.Diagnostics {
	var diags diag.Diagnostics
	m.RuleId = types.StringValue(rule.RuleId)
	m.AccountId = types.StringValue(rule.AccountId)
	// endpoint_service, group_id, and resource_id are Optional and not Computed,
	// and each applies to a single cloud (AWS sets endpoint_service; Azure sets
	// group_id/resource_id), so the others are always unset and the server
	// returns them as "". The plan for an omitted Optional attribute is null, so
	// storing a known "" against it fails Terraform's post-apply consistency
	// check. An empty string is never a valid value for these fields, so "" and
	// null denote the same "unset" state; collapsing "" to null keeps state
	// matching the null plan, mirroring how stringsToList maps empty lists.
	//
	// Caveat: an explicit `group_id = ""` in config plans as a known "" yet
	// still collapses to null here, so it surfaces as a post-apply inconsistency
	// rather than a clean plan-time error. That value is invalid regardless and
	// no validator rejects it earlier; this is the accepted edge of treating ""
	// as null.
	m.EndpointService = stringOrNull(rule.EndpointService)
	m.GroupId = stringOrNull(rule.GroupId)
	m.ResourceId = stringOrNull(rule.ResourceId)
	m.Enabled = types.BoolValue(rule.Enabled)
	m.EndpointName = types.StringValue(rule.EndpointName)
	m.VpcEndpointId = types.StringValue(rule.VpcEndpointId)
	m.ConnectionState = types.StringValue(string(rule.ConnectionState))
	m.CreationTime = types.Int64Value(rule.CreationTime)
	m.UpdatedTime = types.Int64Value(rule.UpdatedTime)
	m.Deactivated = types.BoolValue(rule.Deactivated)
	m.DeactivatedAt = types.Int64Value(rule.DeactivatedAt)
	m.ErrorMessage = types.StringValue(rule.ErrorMessage)

	// An absent or empty server list maps to a typed null, the same shape an
	// unset HCL list takes. Feeding an empty/nil slice to ListValueFrom yields
	// a known empty list, which would diff against a null config and churn the
	// plan on every refresh.
	m.DomainNames = stringsToList(ctx, rule.DomainNames, &diags)
	m.ResourceNames = stringsToList(ctx, rule.ResourceNames, &diags)

	m.GcpEndpoint = gcpEndpointFromAPI(rule.GcpEndpoint)
	return diags
}

// stringOrNull maps a server scalar to a Terraform string, collapsing "" to
// null. Use it only for attributes where the empty string is not a meaningful
// value, so that "" and null denote the same "unset" state.
func stringOrNull(s string) types.String {
	if s == "" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

// stringsToList converts a server-returned string slice into a Terraform list,
// mapping nil or empty to a typed null so it matches an unset HCL list.
func stringsToList(ctx context.Context, vals []string, diags *diag.Diagnostics) types.List {
	if len(vals) == 0 {
		return types.ListNull(types.StringType)
	}
	list, d := types.ListValueFrom(ctx, types.StringType, vals)
	diags.Append(d...)
	return list
}

// gcpEndpointFromAPI converts the SDK's pointer-to-struct into the
// size-0-or-1 slice the schema expects. Returning nil omits the block from
// state.
func gcpEndpointFromAPI(gcp *settings.GcpEndpoint) []gcpEndpointModel {
	if gcp == nil {
		return nil
	}
	return []gcpEndpointModel{{
		PscEndpointUri:    types.StringValue(gcp.PscEndpointUri),
		ServiceAttachment: types.StringValue(gcp.ServiceAttachment),
	}}
}
