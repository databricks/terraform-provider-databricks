// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package entity_tag_assignment

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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "entity_tag_assignment"

var _ datasource.DataSourceWithConfigure = &EntityTagAssignmentDataSource{}

func DataSourceEntityTagAssignment() datasource.DataSource {
	return &EntityTagAssignmentDataSource{}
}

type EntityTagAssignmentDataSource struct {
	Client *autogen.DatabricksClient
}

// EntityTagAssignmentDataExtended extends the main model with additional fields.
type EntityTagAssignmentDataExtended struct {
	catalog_tf.EntityTagAssignment
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// EntityTagAssignmentDataExtended struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m EntityTagAssignmentDataExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.EntityTagAssignment.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EntityTagAssignmentDataExtended
// only implements ToObjectValue() and Type().
func (m EntityTagAssignmentDataExtended) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return m.EntityTagAssignment.ToObjectValue(ctx)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m EntityTagAssignmentDataExtended) Type(ctx context.Context) attr.Type {
	return m.EntityTagAssignment.Type(ctx)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *EntityTagAssignmentDataExtended) SyncFieldsDuringRead(ctx context.Context, existingState EntityTagAssignmentDataExtended) {
	m.EntityTagAssignment.SyncFieldsDuringRead(ctx, existingState.EntityTagAssignment)
}

func (r *EntityTagAssignmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *EntityTagAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, EntityTagAssignmentDataExtended{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks EntityTagAssignment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EntityTagAssignmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *EntityTagAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config EntityTagAssignmentDataExtended
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetEntityTagAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.EntityTagAssignments.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get entity_tag_assignment", err.Error())
		return
	}

	var newState EntityTagAssignmentDataExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
