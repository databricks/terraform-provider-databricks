package mws_private_access_settings

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional().SetComputed()
	attrs["private_access_settings_name"] = attrs["private_access_settings_name"].SetOptional().SetComputed()
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

	pasName := data.PrivateAccessSettingsName.ValueString()
	pasId := data.PrivateAccessSettingsId.ValueString()
	pas, diags := r.getPrivateAccessSettings(ctx, c, pasName, pasId)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var newState MwsPrivateAccessSettingsData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, pas, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *MwsPrivateAccessSettingsDataSource) getPrivateAccessSettings(ctx context.Context, c *databricks.AccountClient, pasName, pasId string) (p provisioning.PrivateAccessSettings, dd diag.Diagnostics) {
	if pasName != "" {
		all, err := c.PrivateAccess.List(ctx)
		if err != nil {
			dd.AddError("failed to list private access settings", err.Error())
			return
		}
		var matches []provisioning.PrivateAccessSettings
		for _, pas := range all {
			if pas.PrivateAccessSettingsName == pasName {
				matches = append(matches, pas)
			}
		}
		dd.Append(validatePrivateAccessSettingsList(ctx, matches, pasName)...)
		if dd.HasError() {
			return
		}
		return matches[0], dd
	}
	if pasId != "" {
		pas, err := c.PrivateAccess.GetByPrivateAccessSettingsId(ctx, pasId)
		if err != nil {
			dd.AddError(fmt.Sprintf("failed to get private access settings with id: %s", pasId), err.Error())
			return
		}
		return *pas, dd
	}

	dd.AddError("you need to specify either `private_access_settings_name` or `private_access_settings_id`", "")
	return
}

func validatePrivateAccessSettingsList(_ context.Context, matches []provisioning.PrivateAccessSettings, name string) diag.Diagnostics {
	if len(matches) == 0 {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is no private access settings with name '%s'", name), "")}
	}
	if len(matches) > 1 {
		ids := make([]string, len(matches))
		for i, m := range matches {
			ids[i] = m.PrivateAccessSettingsId
		}
		return diag.Diagnostics{diag.NewErrorDiagnostic(
			fmt.Sprintf("there is more than one private access settings with name '%s'", name),
			fmt.Sprintf("The IDs of those settings are: %s. Specify the exact one using private_access_settings_id.", strings.Join(ids, ", ")),
		)}
	}
	return nil
}
