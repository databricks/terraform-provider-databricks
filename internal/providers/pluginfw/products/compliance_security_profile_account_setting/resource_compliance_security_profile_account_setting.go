package compliance_security_profile_account_setting

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const (
	resourceName = "compliance_security_profile_account_setting"
	fieldMask    = "csp_enablement_account.is_enforced,csp_enablement_account.compliance_standards"
)

var _ resource.ResourceWithConfigure = &ComplianceSecurityProfileAccountSettingResource{}
var _ resource.ResourceWithImportState = &ComplianceSecurityProfileAccountSettingResource{}

type ComplianceSecurityProfileAccountSettingResource struct {
	Client *autogen.DatabricksClient
}

type ComplianceSecurityProfileAccountSetting struct {
	settings_tf.CspEnablementAccountSetting
}

func ResourceComplianceSecurityProfileAccountSetting() resource.Resource {
	return &ComplianceSecurityProfileAccountSettingResource{}
}

func (m ComplianceSecurityProfileAccountSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.CspEnablementAccountSetting.GetComplexFieldTypes(ctx)
}

func (m ComplianceSecurityProfileAccountSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embedded := m.CspEnablementAccountSetting.ToObjectValue(ctx)
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embedded.Attributes(),
	)
}

func (m ComplianceSecurityProfileAccountSetting) Type(ctx context.Context) attr.Type {
	embedded := m.CspEnablementAccountSetting.Type(ctx).(basetypes.ObjectType)
	return types.ObjectType{AttrTypes: embedded.AttributeTypes()}
}

func (m *ComplianceSecurityProfileAccountSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan ComplianceSecurityProfileAccountSetting) {
	m.CspEnablementAccountSetting.SyncFieldsDuringCreateOrUpdate(ctx, plan.CspEnablementAccountSetting)
}

func (m *ComplianceSecurityProfileAccountSetting) SyncFieldsDuringRead(ctx context.Context, state ComplianceSecurityProfileAccountSetting) {
	m.CspEnablementAccountSetting.SyncFieldsDuringRead(ctx, state.CspEnablementAccountSetting)
}

func (r *ComplianceSecurityProfileAccountSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ComplianceSecurityProfileAccountSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ComplianceSecurityProfileAccountSetting{}, func(cs tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		cs.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "etag")
		cs.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "setting_name")
		cs.SetRequired("csp_enablement_account")
		cs.SetRequired("csp_enablement_account", "compliance_standards")
		cs.SetRequired("csp_enablement_account", "is_enforced")
		return cs
	})

	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks compliance_security_profile_account_setting",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ComplianceSecurityProfileAccountSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ComplianceSecurityProfileAccountSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ComplianceSecurityProfileAccountSetting
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.SettingName = types.StringValue("default")
	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ComplianceSecurityProfileAccountSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ComplianceSecurityProfileAccountSetting
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readReq settings.GetCspEnablementAccountSettingRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &readReq)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := client.Settings.CspEnablementAccount().Get(ctx, readReq)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get compliance_security_profile_account_setting", err.Error())
		return
	}

	var newState ComplianceSecurityProfileAccountSetting
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, res, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, state)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *ComplianceSecurityProfileAccountSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ComplianceSecurityProfileAccountSetting
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.SettingName = types.StringValue("default")
	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ComplianceSecurityProfileAccountSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Settings.CspEnablementAccount().Update(ctx, settings.UpdateCspEnablementAccountSettingRequest{
		AllowMissing: true,
		FieldMask:    fieldMask,
		Setting: settings.CspEnablementAccountSetting{
			SettingName: "default",
			CspEnablementAccount: settings.CspEnablementAccount{
				IsEnforced:          false,
				ComplianceStandards: []settings.ComplianceStandard{},
				ForceSendFields:     []string{"IsEnforced"},
			},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to disable compliance_security_profile_account_setting", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *ComplianceSecurityProfileAccountSettingResource) update(ctx context.Context, plan ComplianceSecurityProfileAccountSetting, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetAccountClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var setting settings.CspEnablementAccountSetting
	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &setting)...)
	if diags.HasError() {
		return
	}

	setting.SettingName = "default"
	setting.CspEnablementAccount.ForceSendFields = []string{"IsEnforced"}

	res, err := client.Settings.CspEnablementAccount().Update(ctx, settings.UpdateCspEnablementAccountSettingRequest{
		AllowMissing: true,
		FieldMask:    fieldMask,
		Setting:      setting,
	})
	if err != nil {
		diags.AddError("failed to update compliance_security_profile_account_setting", err.Error())
		return
	}

	var newState ComplianceSecurityProfileAccountSetting
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, res, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ComplianceSecurityProfileAccountSettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "global" {
		resp.Diagnostics.AddError(
			"Unexpected import identifier",
			fmt.Sprintf("Expected identifier `global`, got %q", req.ID),
		)
		return
	}

	resource.ImportStatePassthroughID(ctx, path.Root("setting_name"), req, resp)
}
