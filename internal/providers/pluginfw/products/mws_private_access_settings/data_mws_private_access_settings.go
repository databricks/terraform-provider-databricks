package mws_private_access_settings

import (
	"context"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "mws_private_access_settings"

func DataSourceMwsPrivateAccessSettings() datasource.DataSource {
	return &MwsPrivateAccessSettingsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &MwsPrivateAccessSettingsDataSource{}

type MwsPrivateAccessSettingsDataSource struct {
	Client *common.DatabricksClient
}

type MwsPrivateAccessSettingsData struct {
	AccountId                 types.String `tfsdk:"account_id"`
	AllowedVpcEndpointIds     types.List   `tfsdk:"allowed_vpc_endpoint_ids"`
	PrivateAccessLevel        types.String `tfsdk:"private_access_level"`
	PrivateAccessSettingsId   types.String `tfsdk:"private_access_settings_id"`
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	PublicAccessEnabled       types.Bool   `tfsdk:"public_access_enabled"`
	Region                    types.String `tfsdk:"region"`
}

func (m MwsPrivateAccessSettingsData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

func (m MwsPrivateAccessSettingsData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                   m.AccountId,
			"allowed_vpc_endpoint_ids":     m.AllowedVpcEndpointIds,
			"private_access_level":         m.PrivateAccessLevel,
			"private_access_settings_id":   m.PrivateAccessSettingsId,
			"private_access_settings_name": m.PrivateAccessSettingsName,
			"public_access_enabled":        m.PublicAccessEnabled,
			"region":                       m.Region,
		},
	)
}

func (m MwsPrivateAccessSettingsData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                   types.StringType,
			"allowed_vpc_endpoint_ids":     basetypes.ListType{ElemType: types.StringType},
			"private_access_level":         types.StringType,
			"private_access_settings_id":   types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

func (m MwsPrivateAccessSettingsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["allowed_vpc_endpoint_ids"] = attrs["allowed_vpc_endpoint_ids"].SetComputed()
	attrs["private_access_level"] = attrs["private_access_level"].SetComputed()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetRequired()
	attrs["private_access_settings_name"] = attrs["private_access_settings_name"].SetComputed()
	attrs["public_access_enabled"] = attrs["public_access_enabled"].SetComputed()
	attrs["region"] = attrs["region"].SetComputed()

	return attrs
}

func (r *MwsPrivateAccessSettingsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (r *MwsPrivateAccessSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, MwsPrivateAccessSettingsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks MwsPrivateAccessSettings",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *MwsPrivateAccessSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = pluginfwcommon.ConfigureDataSource(req, resp)
}

func (r *MwsPrivateAccessSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var data MwsPrivateAccessSettingsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := c.PrivateAccess.GetByPrivateAccessSettingsId(ctx, data.PrivateAccessSettingsId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to get mws_private_access_settings", err.Error())
		return
	}

	var newState MwsPrivateAccessSettingsData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
