// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package domain

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/domains"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/domains_tf"
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

const dataSourceName = "domain"

var _ datasource.DataSourceWithConfigure = &DomainDataSource{}

func DataSourceDomain() datasource.DataSource {
	return &DomainDataSource{}
}

type DomainDataSource struct {
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

// DomainData extends the main model with additional fields.
type DomainData struct {
	// Principal IDs of the business owners (users, groups, or service
	// principals).
	BusinessOwnerIds types.List `tfsdk:"business_owner_ids"`
	// Timestamp when the domain was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Full description (max 4096 chars)
	Description types.String `tfsdk:"description"`
	// Unique identifier for the domain. If omitted at Create, the server
	// generates one.
	DomainId types.String `tfsdk:"domain_id"`
	// Whether to mark the domain as a draft. If omitted on Create, the server
	// applies a default; the resolved value is returned in `effective_draft`.
	Draft types.Bool `tfsdk:"draft"`
	// Resolved draft state of the domain.
	EffectiveDraft types.Bool `tfsdk:"effective_draft"`
	// Icon to display for the domain.
	Icon types.Object `tfsdk:"icon"`
	// Full resource name of the domain. The primary identifier for this
	// resource. Format: `domains/{domain_id}` Identifies the domain on get,
	// update, and delete. Not an input on create — to choose the id, set
	// `CreateDomainRequest.domain_id`.
	Name types.String `tfsdk:"name"`
	// Domain ID of the parent. If absent, this is a top-level domain. If
	// present, this domain is a subdomain of the specified parent.
	ParentDomainId types.String `tfsdk:"parent_domain_id"`
	// Short description (max 280 chars)
	Subtitle types.String `tfsdk:"subtitle"`
	// Governed tag key associated with this domain.
	TagKey types.String `tfsdk:"tag_key"`
	// Principal IDs of the technical owners (users, groups, or service
	// principals).
	TechnicalOwnerIds types.List `tfsdk:"technical_owner_ids"`
	// Timestamp when the domain was last updated.
	UpdateTime         timetypes.RFC3339 `tfsdk:"update_time"`
	ProviderConfigData types.Object      `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DomainData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DomainData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"business_owner_ids":  reflect.TypeOf(types.Int64{}),
		"icon":                reflect.TypeOf(domains_tf.DomainIcon{}),
		"technical_owner_ids": reflect.TypeOf(types.Int64{}),
		"provider_config":     reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DomainData
// only implements ToObjectValue() and Type().
func (m DomainData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"business_owner_ids":  m.BusinessOwnerIds,
			"create_time":         m.CreateTime,
			"description":         m.Description,
			"domain_id":           m.DomainId,
			"draft":               m.Draft,
			"effective_draft":     m.EffectiveDraft,
			"icon":                m.Icon,
			"name":                m.Name,
			"parent_domain_id":    m.ParentDomainId,
			"subtitle":            m.Subtitle,
			"tag_key":             m.TagKey,
			"technical_owner_ids": m.TechnicalOwnerIds,
			"update_time":         m.UpdateTime,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DomainData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"business_owner_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"create_time":      timetypes.RFC3339{}.Type(ctx),
			"description":      types.StringType,
			"domain_id":        types.StringType,
			"draft":            types.BoolType,
			"effective_draft":  types.BoolType,
			"icon":             domains_tf.DomainIcon{}.Type(ctx),
			"name":             types.StringType,
			"parent_domain_id": types.StringType,
			"subtitle":         types.StringType,
			"tag_key":          types.StringType,
			"technical_owner_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m DomainData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["business_owner_ids"] = attrs["business_owner_ids"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["domain_id"] = attrs["domain_id"].SetComputed()
	attrs["draft"] = attrs["draft"].SetComputed()
	attrs["effective_draft"] = attrs["effective_draft"].SetComputed()
	attrs["icon"] = attrs["icon"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parent_domain_id"] = attrs["parent_domain_id"].SetComputed()
	attrs["subtitle"] = attrs["subtitle"].SetComputed()
	attrs["tag_key"] = attrs["tag_key"].SetComputed()
	attrs["technical_owner_ids"] = attrs["technical_owner_ids"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *DomainDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *DomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DomainData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Domain",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DomainDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DomainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config DomainData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest domains.GetDomainRequest
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

	response, err := client.Domains.GetDomain(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get domain", err.Error())
		return
	}

	var newState DomainData
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
