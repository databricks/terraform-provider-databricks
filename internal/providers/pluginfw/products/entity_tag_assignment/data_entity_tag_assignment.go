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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	return attrs
}

func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// EntityTagAssignmentData extends the main model with additional fields.
type EntityTagAssignmentData struct {
	// The fully qualified name of the entity to which the tag is assigned
	EntityName types.String `tfsdk:"entity_name"`
	// The type of the entity to which the tag is assigned. Allowed values are:
	// catalogs, schemas, tables, columns, volumes.
	EntityType types.String `tfsdk:"entity_type"`
	// The key of the tag
	TagKey types.String `tfsdk:"tag_key"`
	// The value of the tag
	TagValue           types.String `tfsdk:"tag_value"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// EntityTagAssignmentData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m EntityTagAssignmentData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EntityTagAssignmentData
// only implements ToObjectValue() and Type().
func (m EntityTagAssignmentData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"entity_name": m.EntityName,
			"entity_type": m.EntityType,
			"tag_key":     m.TagKey,
			"tag_value":   m.TagValue,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m EntityTagAssignmentData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"entity_name": types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
			"tag_value":   types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *EntityTagAssignmentData) SyncFieldsDuringRead(ctx context.Context, from EntityTagAssignmentData) {
	to.ProviderConfigData = from.ProviderConfigData

}

func (m EntityTagAssignmentData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_name"] = attrs["entity_name"].SetRequired()
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_value"] = attrs["tag_value"].SetOptional()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *EntityTagAssignmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *EntityTagAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, EntityTagAssignmentData{}, nil)
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

	var config EntityTagAssignmentData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetEntityTagAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
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

	var newState EntityTagAssignmentData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
