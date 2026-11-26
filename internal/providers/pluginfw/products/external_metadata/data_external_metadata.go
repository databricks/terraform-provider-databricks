// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package external_metadata

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "external_metadata"

var _ datasource.DataSourceWithConfigure = &ExternalMetadataDataSource{}

func DataSourceExternalMetadata() datasource.DataSource {
	return &ExternalMetadataDataSource{}
}

type ExternalMetadataDataSource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfigData contains the fields to configure the provider.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	return attrs
}

// ProviderConfigDataWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigDataWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfigData
// only implements ToObjectValue() and Type().
func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// ExternalMetadataData extends the main model with additional fields.
type ExternalMetadataData struct {
	// Names of the columns of the object, used for column-level mappings in
	// external lineage relationships. Column names must be unique within the
	// list, and the combined size of all column names is limited.
	Columns types.List `tfsdk:"columns"`
	// Time at which this external metadata object was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Username of external metadata object creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// User-provided free-form text description of the object. The description
	// must not exceed 5000 characters.
	Description types.String `tfsdk:"description"`
	// Type of the object within the external system, such as a table, a report,
	// or a dashboard. The value is free-form text of at most 255 characters.
	EntityType types.String `tfsdk:"entity_type"`
	// Unique identifier of the external metadata object. The ID is a UUID that
	// the system assigns on creation.
	Id types.String `tfsdk:"id"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of the external metadata object. The name is unique within the
	// metastore, must follow the [Unity Catalog naming requirements], and is
	// stored in lowercase. It cannot be changed after creation, and external
	// lineage relationships refer to the object by this name.
	//
	// [Unity Catalog naming requirements]: https://docs.databricks.com/aws/en/data-governance/unity-catalog/requirements#naming-requirements
	Name types.String `tfsdk:"name"`
	// Owner of the external metadata object. Defaults to the creator. To change
	// the owner, send an update whose update mask includes `owner`.
	Owner types.String `tfsdk:"owner"`
	// Key-value properties attached to the object. Keys and values must be
	// non-empty, and the combined length of all keys and values must not exceed
	// 5000 characters.
	Properties types.Map `tfsdk:"properties"`
	// Type of the external system that holds the object. Use `OTHER` for a
	// system that is not in the list.
	SystemType types.String `tfsdk:"system_type"`
	// Time at which this external metadata object was last modified.
	UpdateTime types.String `tfsdk:"update_time"`
	// Username of user who last modified external metadata object.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// URL of the object in the external system. When set, lineage graph viewers
	// can follow it to the external asset. The URL must not exceed 2048
	// characters.
	Url                types.String `tfsdk:"url"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ExternalMetadataData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ExternalMetadataData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":         reflect.TypeOf(types.String{}),
		"properties":      reflect.TypeOf(types.String{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalMetadataData
// only implements ToObjectValue() and Type().
func (m ExternalMetadataData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns":      m.Columns,
			"create_time":  m.CreateTime,
			"created_by":   m.CreatedBy,
			"description":  m.Description,
			"entity_type":  m.EntityType,
			"id":           m.Id,
			"metastore_id": m.MetastoreId,
			"name":         m.Name,
			"owner":        m.Owner,
			"properties":   m.Properties,
			"system_type":  m.SystemType,
			"update_time":  m.UpdateTime,
			"updated_by":   m.UpdatedBy,
			"url":          m.Url,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ExternalMetadataData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"create_time":  types.StringType,
			"created_by":   types.StringType,
			"description":  types.StringType,
			"entity_type":  types.StringType,
			"id":           types.StringType,
			"metastore_id": types.StringType,
			"name":         types.StringType,
			"owner":        types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"system_type": types.StringType,
			"update_time": types.StringType,
			"updated_by":  types.StringType,
			"url":         types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m ExternalMetadataData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["entity_type"] = attrs["entity_type"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["properties"] = attrs["properties"].SetComputed()
	attrs["system_type"] = attrs["system_type"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["url"] = attrs["url"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *ExternalMetadataDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ExternalMetadataDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ExternalMetadataData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ExternalMetadata",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ExternalMetadataDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ExternalMetadataDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ExternalMetadataData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetExternalMetadataRequest
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
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.ExternalMetadata.GetExternalMetadata(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get external_metadata", err.Error())
		return
	}

	var newState ExternalMetadataData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve provider_config from config so state.Set has the correct type info
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
