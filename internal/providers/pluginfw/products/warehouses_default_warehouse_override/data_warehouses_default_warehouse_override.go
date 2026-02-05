// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses_default_warehouse_override

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "warehouses_default_warehouse_override"

var _ datasource.DataSourceWithConfigure = &DefaultWarehouseOverrideDataSource{}

func DataSourceDefaultWarehouseOverride() datasource.DataSource {
	return &DefaultWarehouseOverrideDataSource{}
}

type DefaultWarehouseOverrideDataSource struct {
	Client *autogen.DatabricksClient
}

// DefaultWarehouseOverrideData extends the main model with additional fields.
type DefaultWarehouseOverrideData struct {
	// The ID component of the resource name (user ID).
	DefaultWarehouseOverrideId types.String `tfsdk:"default_warehouse_override_id"`
	// The resource name of the default warehouse override. Format:
	// default-warehouse-overrides/{default_warehouse_override_id}
	Name types.String `tfsdk:"name"`
	// The type of override behavior.
	Type_ types.String `tfsdk:"type"`
	// The specific warehouse ID when type is CUSTOM. Not set for LAST_SELECTED
	// type.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DefaultWarehouseOverrideData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DefaultWarehouseOverrideData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultWarehouseOverrideData
// only implements ToObjectValue() and Type().
func (m DefaultWarehouseOverrideData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_warehouse_override_id": m.DefaultWarehouseOverrideId,
			"name":                          m.Name,
			"type":                          m.Type_,
			"warehouse_id":                  m.WarehouseId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DefaultWarehouseOverrideData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_warehouse_override_id": types.StringType,
			"name":                          types.StringType,
			"type":                          types.StringType,
			"warehouse_id":                  types.StringType,
		},
	}
}

func (m DefaultWarehouseOverrideData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_warehouse_override_id"] = attrs["default_warehouse_override_id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["type"] = attrs["type"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetComputed()

	return attrs
}

func (r *DefaultWarehouseOverrideDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *DefaultWarehouseOverrideDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DefaultWarehouseOverrideData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks DefaultWarehouseOverride",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DefaultWarehouseOverrideDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DefaultWarehouseOverrideDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config DefaultWarehouseOverrideData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest sql.GetDefaultWarehouseOverrideRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Warehouses.GetDefaultWarehouseOverride(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get warehouses_default_warehouse_override", err.Error())
		return
	}

	var newState DefaultWarehouseOverrideData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
