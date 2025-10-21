// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_quality_refresh

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/dataquality"
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

const dataSourceName = "data_quality_refresh"

var _ datasource.DataSourceWithConfigure = &RefreshDataSource{}

func DataSourceRefresh() datasource.DataSource {
	return &RefreshDataSource{}
}

type RefreshDataSource struct {
	Client *autogen.DatabricksClient
}

// RefreshData extends the main model with additional fields.
type RefreshData struct {
	// Time when the refresh ended (milliseconds since 1/1/1970 UTC).
	EndTimeMs types.Int64 `tfsdk:"end_time_ms"`
	// An optional message to give insight into the current state of the refresh
	// (e.g. FAILURE messages).
	Message types.String `tfsdk:"message"`
	// The UUID of the request object. For example, table id.
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`or
	// `table`.
	ObjectType types.String `tfsdk:"object_type"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"refresh_id"`
	// Time when the refresh started (milliseconds since 1/1/1970 UTC).
	StartTimeMs types.Int64 `tfsdk:"start_time_ms"`
	// The current state of the refresh.
	State types.String `tfsdk:"state"`
	// What triggered the refresh.
	Trigger types.String `tfsdk:"trigger"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// RefreshData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m RefreshData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RefreshData
// only implements ToObjectValue() and Type().
func (m RefreshData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time_ms":   m.EndTimeMs,
			"message":       m.Message,
			"object_id":     m.ObjectId,
			"object_type":   m.ObjectType,
			"refresh_id":    m.RefreshId,
			"start_time_ms": m.StartTimeMs,
			"state":         m.State,
			"trigger":       m.Trigger,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m RefreshData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"end_time_ms": types.Int64Type,
			"message":       types.StringType,
			"object_id":     types.StringType,
			"object_type":   types.StringType,
			"refresh_id":    types.Int64Type,
			"start_time_ms": types.Int64Type,
			"state":         types.StringType,
			"trigger":       types.StringType,
		},
	}
}

func (m RefreshData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end_time_ms"] = attrs["end_time_ms"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["refresh_id"] = attrs["refresh_id"].SetComputed()
	attrs["start_time_ms"] = attrs["start_time_ms"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["trigger"] = attrs["trigger"].SetComputed()

	return attrs
}

func (r *RefreshDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *RefreshDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RefreshData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Refresh",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RefreshDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *RefreshDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config RefreshData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest dataquality.GetRefreshRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DataQuality.GetRefresh(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get data_quality_refresh", err.Error())
		return
	}

	var newState RefreshData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
