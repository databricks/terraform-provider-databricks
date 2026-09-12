// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package domain

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/domains"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/domains_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "domain"

var _ resource.ResourceWithConfigure = &DomainResource{}
var _ resource.ResourceWithModifyPlan = &DomainResource{}

func ResourceDomain() resource.Resource {
	return &DomainResource{}
}

type DomainResource struct {
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

// Domain extends the main model with additional fields.
type Domain struct {
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
	UpdateTime     timetypes.RFC3339 `tfsdk:"update_time"`
	ProviderConfig types.Object      `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Domain struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Domain) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"business_owner_ids":  reflect.TypeOf(types.Int64{}),
		"icon":                reflect.TypeOf(domains_tf.DomainIcon{}),
		"technical_owner_ids": reflect.TypeOf(types.Int64{}),
		"provider_config":     reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Domain
// only implements ToObjectValue() and Type().
func (m Domain) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"business_owner_ids": m.BusinessOwnerIds,
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

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Domain) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"business_owner_ids": basetypes.ListType{
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

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Domain) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Domain) {
	if !from.BusinessOwnerIds.IsNull() && !from.BusinessOwnerIds.IsUnknown() && to.BusinessOwnerIds.IsNull() && len(from.BusinessOwnerIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for BusinessOwnerIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.BusinessOwnerIds = from.BusinessOwnerIds
	}
	if !from.Draft.IsUnknown() && !from.Draft.IsNull() {
		// Draft is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Draft = from.Draft
	}
	if !from.Icon.IsNull() && !from.Icon.IsUnknown() {
		if toIcon, ok := to.GetIcon(ctx); ok {
			if fromIcon, ok := from.GetIcon(ctx); ok {
				// Recursively sync the fields of Icon
				toIcon.SyncFieldsDuringCreateOrUpdate(ctx, fromIcon)
				to.SetIcon(ctx, toIcon)
			}
		}
	}
	if !from.TechnicalOwnerIds.IsNull() && !from.TechnicalOwnerIds.IsUnknown() && to.TechnicalOwnerIds.IsNull() && len(from.TechnicalOwnerIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TechnicalOwnerIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TechnicalOwnerIds = from.TechnicalOwnerIds
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Domain) SyncFieldsDuringRead(ctx context.Context, from Domain) {
	if !from.BusinessOwnerIds.IsNull() && !from.BusinessOwnerIds.IsUnknown() && to.BusinessOwnerIds.IsNull() && len(from.BusinessOwnerIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for BusinessOwnerIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.BusinessOwnerIds = from.BusinessOwnerIds
	}
	if !from.Draft.IsUnknown() && !from.Draft.IsNull() {
		// Draft is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Draft = from.Draft
	}
	if !from.Icon.IsNull() && !from.Icon.IsUnknown() {
		if toIcon, ok := to.GetIcon(ctx); ok {
			if fromIcon, ok := from.GetIcon(ctx); ok {
				toIcon.SyncFieldsDuringRead(ctx, fromIcon)
				to.SetIcon(ctx, toIcon)
			}
		}
	}
	if !from.TechnicalOwnerIds.IsNull() && !from.TechnicalOwnerIds.IsUnknown() && to.TechnicalOwnerIds.IsNull() && len(from.TechnicalOwnerIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TechnicalOwnerIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TechnicalOwnerIds = from.TechnicalOwnerIds
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Domain) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["business_owner_ids"] = attrs["business_owner_ids"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["domain_id"] = attrs["domain_id"].SetComputed()
	attrs["domain_id"] = attrs["domain_id"].SetOptional()
	attrs["domain_id"] = attrs["domain_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["domain_id"] = attrs["domain_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)
	attrs["draft"] = attrs["draft"].SetOptional()
	attrs["draft"] = attrs["draft"].SetComputed()
	attrs["draft"] = attrs["draft"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_draft"] = attrs["effective_draft"].SetComputed()
	attrs["icon"] = attrs["icon"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["parent_domain_id"] = attrs["parent_domain_id"].SetOptional()
	attrs["parent_domain_id"] = attrs["parent_domain_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["subtitle"] = attrs["subtitle"].SetOptional()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["technical_owner_ids"] = attrs["technical_owner_ids"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetBusinessOwnerIds returns the value of the BusinessOwnerIds field in Domain as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Domain) GetBusinessOwnerIds(ctx context.Context) ([]types.Int64, bool) {
	if m.BusinessOwnerIds.IsNull() || m.BusinessOwnerIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.BusinessOwnerIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBusinessOwnerIds sets the value of the BusinessOwnerIds field in Domain.
func (m *Domain) SetBusinessOwnerIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["business_owner_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.BusinessOwnerIds = types.ListValueMust(t, vs)
}

// GetIcon returns the value of the Icon field in Domain as
// a domains_tf.DomainIcon value.
// If the field is unknown or null, the boolean return value is false.
func (m *Domain) GetIcon(ctx context.Context) (domains_tf.DomainIcon, bool) {
	var e domains_tf.DomainIcon
	if m.Icon.IsNull() || m.Icon.IsUnknown() {
		return e, false
	}
	var v domains_tf.DomainIcon
	d := m.Icon.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIcon sets the value of the Icon field in Domain.
func (m *Domain) SetIcon(ctx context.Context, v domains_tf.DomainIcon) {
	vs := v.ToObjectValue(ctx)
	m.Icon = vs
}

// GetTechnicalOwnerIds returns the value of the TechnicalOwnerIds field in Domain as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Domain) GetTechnicalOwnerIds(ctx context.Context) ([]types.Int64, bool) {
	if m.TechnicalOwnerIds.IsNull() || m.TechnicalOwnerIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.TechnicalOwnerIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTechnicalOwnerIds sets the value of the TechnicalOwnerIds field in Domain.
func (m *Domain) SetTechnicalOwnerIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["technical_owner_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TechnicalOwnerIds = types.ListValueMust(t, vs)
}

func (r *DomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *DomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Domain{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks domain",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *DomainResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	tfschema.ValidateWorkspaceID(ctx, r.Client, req, resp)
}

func (r *DomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Domain
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var domain domains.Domain

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &domain)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := domains.CreateDomainRequest{
		Domain:   domain,
		DomainId: plan.DomainId.ValueString(),
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

	response, err := client.Domains.CreateDomain(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create domain", err.Error())
		return
	}

	var newState Domain

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

func (r *DomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Domain
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest domains.GetDomainRequest
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
	response, err := client.Domains.GetDomain(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get domain", err.Error())
		return
	}

	var newState Domain
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

func (r *DomainResource) update(ctx context.Context, plan Domain, diags *diag.Diagnostics, state *tfsdk.State) {
	var domain domains.Domain

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &domain)...)
	if diags.HasError() {
		return
	}

	updateRequest := domains.UpdateDomainRequest{
		Domain:     domain,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("business_owner_ids,description,draft,icon,subtitle,technical_owner_ids", ",")),
	}

	var namespace ProviderConfig
	diags.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Domains.UpdateDomain(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update domain", err.Error())
		return
	}

	var newState Domain

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *DomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Domain
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *DomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Domain
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest domains.DeleteDomainRequest
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

	err := client.Domains.DeleteDomain(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete domain", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &DomainResource{}

func (r *DomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
