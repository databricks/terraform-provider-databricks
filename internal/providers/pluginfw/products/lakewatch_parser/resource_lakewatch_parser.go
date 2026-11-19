// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_parser

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/lakewatch_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "lakewatch_parser"

var _ resource.ResourceWithConfigure = &ParserResource{}
var _ resource.ResourceWithModifyPlan = &ParserResource{}

func ResourceParser() resource.Resource {
	return &ParserResource{}
}

type ParserResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(ProviderConfigWorkspaceIDPlanModifier, "", ""))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	return attrs
}

// ProviderConfigWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfig
// only implements ToObjectValue() and Type().
func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// Parser extends the main model with additional fields.
type Parser struct {
	// Name of the parent UC catalog, set on create. Reads return the value as
	// supplied; the current UC catalog name is reported in
	// `effective_catalog_name`.
	CatalogName types.String `tfsdk:"catalog_name"`
	// When the parser was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The principal that created this parser. Each version records its own
	// `created_by`.
	CreatedBy types.Object `tfsdk:"created_by"`
	// Description of version 1. Accepted only by `CreateParser`; read a
	// version's description with `GetParserVersion`.
	Description types.String `tfsdk:"description"`
	// The current name of the parser's UC catalog. Tracks UC renames; when the
	// current name cannot be looked up, the name as of the parser's last write.
	EffectiveCatalogName types.String `tfsdk:"effective_catalog_name"`
	// The current name of the parser's UC schema. Tracks UC renames; when the
	// current name cannot be looked up, the name as of the parser's last write.
	EffectiveSchemaName types.String `tfsdk:"effective_schema_name"`
	// Newest version of this parser. Read its spec by passing this to
	// `GetParserVersion`, or by using the `latest` alias. Absent when every
	// version has been deleted.
	LatestVersion types.Int64 `tfsdk:"latest_version"`
	// Resource name. Format: `parsers/{parser_id}`, where the leaf is the
	// immutable server-assigned UUID. Create-time placement names are stored
	// separately and do not affect the resource name.
	Name types.String `tfsdk:"name"`
	// The principal that owns this parser. Server-set from the authenticated
	// creator, never supplied by the client. Beta reads by UUID perform no
	// authorization. Owner-gated write RPCs are outside this API's current
	// implementation scope.
	Owner types.Object `tfsdk:"owner"`
	// The parser's server-assigned UUID. It is not currently a UC securable ID,
	// is immutable for the life of the resource, is the leaf of `name`, and is
	// returned by CreateParser.
	ParserId types.String `tfsdk:"parser_id"`
	// Client-supplied parser name, stored as supplied and fixed at create time.
	// Must be a valid UC identifier. A custom parser is not currently a UC
	// securable and this field has no rename path.
	ParserName types.String `tfsdk:"parser_name"`
	// Name of the parent UC schema, set on create. Reads return the value as
	// supplied; the current UC schema name is reported in
	// `effective_schema_name`.
	SchemaName types.String `tfsdk:"schema_name"`
	// Spec of version 1, as a YAML document. Accepted only by `CreateParser`;
	// read a version's spec with `GetParserVersion`.
	Spec           types.String `tfsdk:"spec"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Parser struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Parser) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"created_by":      reflect.TypeOf(lakewatch_tf.Principal{}),
		"owner":           reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Parser
// only implements ToObjectValue() and Type().
func (m Parser) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"catalog_name": m.CatalogName,
			"create_time":            m.CreateTime,
			"created_by":             m.CreatedBy,
			"description":            m.Description,
			"effective_catalog_name": m.EffectiveCatalogName,
			"effective_schema_name":  m.EffectiveSchemaName,
			"latest_version":         m.LatestVersion,
			"name":                   m.Name,
			"owner":                  m.Owner,
			"parser_id":              m.ParserId,
			"parser_name":            m.ParserName,
			"schema_name":            m.SchemaName,
			"spec":                   m.Spec,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Parser) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"catalog_name": types.StringType,
			"create_time":            timetypes.RFC3339{}.Type(ctx),
			"created_by":             lakewatch_tf.Principal{}.Type(ctx),
			"description":            types.StringType,
			"effective_catalog_name": types.StringType,
			"effective_schema_name":  types.StringType,
			"latest_version":         types.Int64Type,
			"name":                   types.StringType,
			"owner":                  lakewatch_tf.Principal{}.Type(ctx),
			"parser_id":              types.StringType,
			"parser_name":            types.StringType,
			"schema_name":            types.StringType,
			"spec":                   types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Parser) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Parser) {
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				// Recursively sync the fields of CreatedBy
				toCreatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.Description.IsUnknown() && !from.Description.IsNull() {
		// Description is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Description = from.Description
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				// Recursively sync the fields of Owner
				toOwner.SyncFieldsDuringCreateOrUpdate(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Parser) SyncFieldsDuringRead(ctx context.Context, from Parser) {
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				toCreatedBy.SyncFieldsDuringRead(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.Description.IsUnknown() && !from.Description.IsNull() {
		// Description is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Description = from.Description
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				toOwner.SyncFieldsDuringRead(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Parser) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["catalog_name"] = attrs["catalog_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["created_by"] = attrs["created_by"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["description"] = attrs["description"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetComputed()
	attrs["description"] = attrs["description"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["latest_version"] = attrs["latest_version"].SetComputed()
	attrs["latest_version"] = attrs["latest_version"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["owner"] = attrs["owner"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["parser_id"] = attrs["parser_id"].SetComputed()
	attrs["parser_id"] = attrs["parser_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["parser_name"] = attrs["parser_name"].SetRequired()
	attrs["parser_name"] = attrs["parser_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].SetRequired()
	attrs["spec"] = attrs["spec"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetCreatedBy returns the value of the CreatedBy field in Parser as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Parser) GetCreatedBy(ctx context.Context) (lakewatch_tf.Principal, bool) {
	var e lakewatch_tf.Principal
	if m.CreatedBy.IsNull() || m.CreatedBy.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Principal
	d := m.CreatedBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCreatedBy sets the value of the CreatedBy field in Parser.
func (m *Parser) SetCreatedBy(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.CreatedBy = vs
}

// GetOwner returns the value of the Owner field in Parser as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Parser) GetOwner(ctx context.Context) (lakewatch_tf.Principal, bool) {
	var e lakewatch_tf.Principal
	if m.Owner.IsNull() || m.Owner.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Principal
	d := m.Owner.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOwner sets the value of the Owner field in Parser.
func (m *Parser) SetOwner(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.Owner = vs
}

func (r *ParserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ParserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Parser{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks lakewatch_parser",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ParserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ParserResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
}

func (r *ParserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Parser
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var parser lakewatch.Parser

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &parser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := lakewatch.CreateParserRequest{
		Parser: parser,
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	response, err := client.Lakewatch.CreateParser(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create lakewatch_parser", err.Error())
		return
	}

	var newState Parser

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, plan.ProviderConfig, &resp.State)...)
}

func (r *ParserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Parser
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetParserRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(existingState.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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
	response, err := client.Lakewatch.GetParser(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get lakewatch_parser", err.Error())
		return
	}

	var newState Parser
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, existingState.ProviderConfig, &resp.State)...)
}

func (r *ParserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// VARIANT_IMMUTABLE resources do not support updates - all changes require replacement.
	resp.Diagnostics.AddError("Update not supported", "This resource does not support updates. All changes require replacement.")
}

func (r *ParserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Parser
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest lakewatch.DeleteParserRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(state.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	err := client.Lakewatch.DeleteParser(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete lakewatch_parser", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ParserResource{}

func (r *ParserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
