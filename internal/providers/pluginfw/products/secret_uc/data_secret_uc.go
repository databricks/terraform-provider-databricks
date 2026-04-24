// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package secret_uc

import (
	"context"
	"reflect"
	"regexp"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
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

const dataSourceName = "secret_uc"

var _ datasource.DataSourceWithConfigure = &SecretDataSource{}

func DataSourceSecret() datasource.DataSource {
	return &SecretDataSource{}
}

type SecretDataSource struct {
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
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
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

// SecretData extends the main model with additional fields.
type SecretData struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the **BROWSE** privilege when
	// **include_browse** is enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// The name of the catalog where the schema and the secret reside.
	CatalogName types.String `tfsdk:"catalog_name"`
	// User-provided free-form text description of the secret.
	Comment types.String `tfsdk:"comment"`
	// The time at which this secret was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The principal that created the secret.
	CreatedBy types.String `tfsdk:"created_by"`
	// The effective owner of the secret, which may differ from the directly-set
	// **owner** due to inheritance.
	EffectiveOwner types.String `tfsdk:"effective_owner"`
	// The secret value. Only populated in responses when you have the
	// **READ_SECRET** privilege and **include_value** is set to true in the
	// request. The maximum size is 60 KiB.
	EffectiveValue types.String `tfsdk:"effective_value"`
	// User-provided expiration time of the secret. This field indicates when
	// the secret should no longer be used and may be displayed as a warning in
	// the UI. It is purely informational and does not trigger any automatic
	// actions or affect the secret's lifecycle.
	ExpireTime timetypes.RFC3339 `tfsdk:"expire_time"`

	ExternalSecretId types.String `tfsdk:"external_secret_id"`
	// The three-level (fully qualified) name of the secret, in the form of
	// **catalog_name.schema_name.secret_name**.
	FullName types.String `tfsdk:"full_name"`
	// Unique identifier of the metastore hosting the secret.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the secret, relative to its parent schema.
	Name types.String `tfsdk:"name"`
	// The owner of the secret. Defaults to the creating principal on creation.
	// Can be updated to transfer ownership of the secret to another principal.
	Owner types.String `tfsdk:"owner"`
	// The name of the schema where the secret resides.
	SchemaName types.String `tfsdk:"schema_name"`
	// The time at which this secret was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// The principal that last updated the secret.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// The secret value to store. This field is input-only and is not returned
	// in responses — use the **effective_value** field (via GetSecret with
	// **include_value** set to true) to read the secret value. The maximum size
	// is 60 KiB (pre-encryption). Accepted content includes passwords, tokens,
	// keys, and other sensitive credential data.
	Value              types.String `tfsdk:"value"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// SecretData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m SecretData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretData
// only implements ToObjectValue() and Type().
func (m SecretData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"browse_only":        m.BrowseOnly,
			"catalog_name":       m.CatalogName,
			"comment":            m.Comment,
			"create_time":        m.CreateTime,
			"created_by":         m.CreatedBy,
			"effective_owner":    m.EffectiveOwner,
			"effective_value":    m.EffectiveValue,
			"expire_time":        m.ExpireTime,
			"external_secret_id": m.ExternalSecretId,
			"full_name":          m.FullName,
			"metastore_id":       m.MetastoreId,
			"name":               m.Name,
			"owner":              m.Owner,
			"schema_name":        m.SchemaName,
			"update_time":        m.UpdateTime,
			"updated_by":         m.UpdatedBy,
			"value":              m.Value,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m SecretData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":        types.BoolType,
			"catalog_name":       types.StringType,
			"comment":            types.StringType,
			"create_time":        timetypes.RFC3339{}.Type(ctx),
			"created_by":         types.StringType,
			"effective_owner":    types.StringType,
			"effective_value":    types.StringType,
			"expire_time":        timetypes.RFC3339{}.Type(ctx),
			"external_secret_id": types.StringType,
			"full_name":          types.StringType,
			"metastore_id":       types.StringType,
			"name":               types.StringType,
			"owner":              types.StringType,
			"schema_name":        types.StringType,
			"update_time":        timetypes.RFC3339{}.Type(ctx),
			"updated_by":         types.StringType,
			"value":              types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m SecretData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["browse_only"] = attrs["browse_only"].SetComputed()
	attrs["catalog_name"] = attrs["catalog_name"].SetComputed()
	attrs["comment"] = attrs["comment"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["effective_owner"] = attrs["effective_owner"].SetComputed()
	attrs["effective_value"] = attrs["effective_value"].SetComputed()
	attrs["expire_time"] = attrs["expire_time"].SetComputed()
	attrs["external_secret_id"] = attrs["external_secret_id"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["schema_name"] = attrs["schema_name"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["value"] = attrs["value"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *SecretDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *SecretDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SecretData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Secret",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SecretDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SecretDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config SecretData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetSecretRequest
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

	response, err := client.SecretsUc.GetSecret(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get secret_uc", err.Error())
		return
	}

	var newState SecretData
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
