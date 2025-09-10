// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package alert_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "alert_v2"

var _ datasource.DataSourceWithConfigure = &AlertV2DataSource{}

func DataSourceAlertV2() datasource.DataSource {
	return &AlertV2DataSource{}
}

type AlertV2DataSource struct {
	Client *autogen.DatabricksClient
}

// AlertV2Data extends the main model with additional fields.
type AlertV2Data struct {
	sql_tf.AlertV2
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AlertV2Data struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AlertV2Data) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.AlertV2.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AlertV2Data
// only implements ToObjectValue() and Type().
func (m AlertV2Data) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.AlertV2.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()
	embeddedAttrs["workspace_id"] = m.WorkspaceID

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AlertV2Data) Type(ctx context.Context) attr.Type {
	embeddedType := m.AlertV2.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()
	attrTypes["workspace_id"] = types.StringType

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *AlertV2Data) SyncFieldsDuringRead(ctx context.Context, existingState AlertV2Data) {
	m.AlertV2.SyncFieldsDuringRead(ctx, existingState.AlertV2)
}

func (r *AlertV2DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *AlertV2DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AlertV2Data{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AlertV2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AlertV2DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AlertV2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config AlertV2Data
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest sql.GetAlertV2Request
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.AlertsV2.GetAlert(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get alert_v2", err.Error())
		return
	}

	var newState AlertV2Data
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
