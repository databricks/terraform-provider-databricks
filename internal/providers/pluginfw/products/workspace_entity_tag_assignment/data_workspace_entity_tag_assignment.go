// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_entity_tag_assignment

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/tags"
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

const dataSourceName = "workspace_entity_tag_assignment"

var _ datasource.DataSourceWithConfigure = &TagAssignmentDataSource{}

func DataSourceTagAssignment() datasource.DataSource {
	return &TagAssignmentDataSource{}
}

type TagAssignmentDataSource struct {
	Client *autogen.DatabricksClient
}

// TagAssignmentData extends the main model with additional fields.
type TagAssignmentData struct {
	// The identifier of the entity to which the tag is assigned
	EntityId types.String `tfsdk:"entity_id"`
	// The type of entity to which the tag is assigned. Allowed values are
	// dashboards, geniespaces
	EntityType types.String `tfsdk:"entity_type"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey types.String `tfsdk:"tag_key"`
	// The value of the tag
	TagValue types.String `tfsdk:"tag_value"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// TagAssignmentData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m TagAssignmentData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagAssignmentData
// only implements ToObjectValue() and Type().
func (m TagAssignmentData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   m.EntityId,
			"entity_type": m.EntityType,
			"tag_key":     m.TagKey,
			"tag_value":   m.TagValue,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m TagAssignmentData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
			"tag_value":   types.StringType,
		},
	}
}

func (m TagAssignmentData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_value"] = attrs["tag_value"].SetComputed()

	return attrs
}

func (r *TagAssignmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *TagAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, TagAssignmentData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks TagAssignment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *TagAssignmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *TagAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config TagAssignmentData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest tags.GetTagAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.WorkspaceEntityTagAssignments.GetTagAssignment(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get workspace_entity_tag_assignment", err.Error())
		return
	}

	var newState TagAssignmentData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
