// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_gateway_model_provider_service

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "ai_gateway_model_provider_service"

var _ datasource.DataSourceWithConfigure = &ModelProviderServiceDataSource{}

func DataSourceModelProviderService() datasource.DataSource {
	return &ModelProviderServiceDataSource{}
}

type ModelProviderServiceDataSource struct {
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

// ModelProviderServiceData extends the main model with additional fields.
type ModelProviderServiceData struct {
	// Whether the caller sees only metadata available through the BROWSE
	// privilege.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// User-provided description.
	Comment types.String `tfsdk:"comment"`
	// Behavioral configuration: provider connection, model catalog, and
	// passthrough policy. See `ModelProviderServiceConfig` for the per-field
	// contract. Required on CreateModelProviderService; on Update it is
	// required only when `config` (or a `config.*` subpath) appears in
	// `update_mask`.
	Config types.Object `tfsdk:"config"`
	// When the provider service was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Creator identity.
	CreatedBy types.String `tfsdk:"created_by"`
	// The resolved owner of the model provider service. Falls back to the
	// caller's identity when `owner` is not explicitly set on creation.
	EffectiveOwner types.String `tfsdk:"effective_owner"`
	// Optimistic concurrency control token. Server-generated from the entity's
	// state and returned on every read. To use it as an if-match precondition
	// on a mutation, echo the last-read value back via the dedicated `etag`
	// field on the Update / Delete request; the server rejects the mutation if
	// the stored etag differs.
	Etag types.String `tfsdk:"etag"`
	// Metastore hosting the provider service.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Resource name of the provider service. Format:
	// `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
	// Each `{...}` component is capped at 255 characters individually.
	// Server-derived on Create from `parent` + `model_provider_service_id`;
	// required and immutable on Update/Get/Delete.
	Name types.String `tfsdk:"name"`
	// The owner of the model provider service. Write-only; read owner via
	// effective_owner.
	Owner types.String `tfsdk:"owner"`
	// When the provider service was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Identity of the last updater.
	UpdatedBy          types.String `tfsdk:"updated_by"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ModelProviderServiceData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ModelProviderServiceData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config":          reflect.TypeOf(catalog_tf.ModelProviderServiceConfig{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelProviderServiceData
// only implements ToObjectValue() and Type().
func (m ModelProviderServiceData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"browse_only":     m.BrowseOnly,
			"comment":         m.Comment,
			"config":          m.Config,
			"create_time":     m.CreateTime,
			"created_by":      m.CreatedBy,
			"effective_owner": m.EffectiveOwner,
			"etag":            m.Etag,
			"metastore_id":    m.MetastoreId,
			"name":            m.Name,
			"owner":           m.Owner,
			"update_time":     m.UpdateTime,
			"updated_by":      m.UpdatedBy,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ModelProviderServiceData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":     types.BoolType,
			"comment":         types.StringType,
			"config":          catalog_tf.ModelProviderServiceConfig{}.Type(ctx),
			"create_time":     timetypes.RFC3339{}.Type(ctx),
			"created_by":      types.StringType,
			"effective_owner": types.StringType,
			"etag":            types.StringType,
			"metastore_id":    types.StringType,
			"name":            types.StringType,
			"owner":           types.StringType,
			"update_time":     timetypes.RFC3339{}.Type(ctx),
			"updated_by":      types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m ModelProviderServiceData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["browse_only"] = attrs["browse_only"].SetComputed()
	attrs["comment"] = attrs["comment"].SetComputed()
	attrs["config"] = attrs["config"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["effective_owner"] = attrs["effective_owner"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *ModelProviderServiceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ModelProviderServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ModelProviderServiceData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ModelProviderService",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ModelProviderServiceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ModelProviderServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ModelProviderServiceData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetModelProviderServiceRequest
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

	response, err := client.AiGateway.GetModelProviderService(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get ai_gateway_model_provider_service", err.Error())
		return
	}

	var newState ModelProviderServiceData
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
