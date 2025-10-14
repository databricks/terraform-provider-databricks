// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package rfa_access_request_destinations

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "rfa_access_request_destinations"

var _ datasource.DataSourceWithConfigure = &AccessRequestDestinationDataSource{}

func DataSourceAccessRequestDestination() datasource.DataSource {
	return &AccessRequestDestinationDataSource{}
}

type AccessRequestDestinationDataSource struct {
	Client *autogen.DatabricksClient
}

// AccessRequestDestinationsData extends the main model with additional fields.
type AccessRequestDestinationsData struct {
	catalog_tf.AccessRequestDestinations
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AccessRequestDestinationsData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AccessRequestDestinationsData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.AccessRequestDestinations.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessRequestDestinationsData
// only implements ToObjectValue() and Type().
func (m AccessRequestDestinationsData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.AccessRequestDestinations.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccessRequestDestinationsData) Type(ctx context.Context) attr.Type {
	embeddedType := m.AccessRequestDestinations.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *AccessRequestDestinationsData) SyncFieldsDuringRead(ctx context.Context, existingState AccessRequestDestinationsData) {
	m.AccessRequestDestinations.SyncFieldsDuringRead(ctx, existingState.AccessRequestDestinations)
}

func (r *AccessRequestDestinationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *AccessRequestDestinationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AccessRequestDestinationsData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AccessRequestDestinations",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AccessRequestDestinationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AccessRequestDestinationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config AccessRequestDestinationsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetAccessRequestDestinationsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Rfa.GetAccessRequestDestinations(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get rfa_access_request_destinations", err.Error())
		return
	}

	var newState AccessRequestDestinationsData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
