// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package online_store

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/ml_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "online_store"

var _ datasource.DataSourceWithConfigure = &OnlineStoreDataSource{}

func DataSourceOnlineStore() datasource.DataSource {
	return &OnlineStoreDataSource{}
}

type OnlineStoreDataSource struct {
	Client *autogen.DatabricksClient
}

// OnlineStoreData extends the main model with additional fields.
type OnlineStoreData struct {
	ml_tf.OnlineStore
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// OnlineStoreData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m OnlineStoreData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.OnlineStore.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineStoreData
// only implements ToObjectValue() and Type().
func (m OnlineStoreData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.OnlineStore.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()
	embeddedAttrs["workspace_id"] = m.WorkspaceID

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m OnlineStoreData) Type(ctx context.Context) attr.Type {
	embeddedType := m.OnlineStore.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()
	attrTypes["workspace_id"] = types.StringType

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *OnlineStoreData) SyncFieldsDuringRead(ctx context.Context, existingState OnlineStoreData) {
	m.OnlineStore.SyncFieldsDuringRead(ctx, existingState.OnlineStore)
}

func (r *OnlineStoreDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *OnlineStoreDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, OnlineStoreData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks OnlineStore",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *OnlineStoreDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *OnlineStoreDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config OnlineStoreData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest ml.GetOnlineStoreRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureStore.GetOnlineStore(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get online_store", err.Error())
		return
	}

	var newState OnlineStoreData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
