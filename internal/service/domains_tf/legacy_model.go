// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package domains_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateDomainRequest_SdkV2 struct {
	Domain types.List `tfsdk:"domain"`
	// Client-supplied resource ID for the new domain. If omitted, the server
	// generates one.
	DomainId types.String `tfsdk:"-"`
}

func (to *CreateDomainRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDomainRequest_SdkV2) {
	if !from.Domain.IsNull() && !from.Domain.IsUnknown() {
		if toDomain, ok := to.GetDomain(ctx); ok {
			if fromDomain, ok := from.GetDomain(ctx); ok {
				// Recursively sync the fields of Domain
				toDomain.SyncFieldsDuringCreateOrUpdate(ctx, fromDomain)
				to.SetDomain(ctx, toDomain)
			}
		}
	}
}

func (to *CreateDomainRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDomainRequest_SdkV2) {
	if !from.Domain.IsNull() && !from.Domain.IsUnknown() {
		if toDomain, ok := to.GetDomain(ctx); ok {
			if fromDomain, ok := from.GetDomain(ctx); ok {
				toDomain.SyncFieldsDuringRead(ctx, fromDomain)
				to.SetDomain(ctx, toDomain)
			}
		}
	}
}

func (m CreateDomainRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["domain"] = attrs["domain"].SetRequired()
	attrs["domain"] = attrs["domain"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["domain_id"] = attrs["domain_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDomainRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDomainRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain": reflect.TypeOf(Domain_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDomainRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDomainRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domain":    m.Domain,
			"domain_id": m.DomainId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDomainRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"domain": basetypes.ListType{
				ElemType: Domain_SdkV2{}.Type(ctx),
			},
			"domain_id": types.StringType,
		},
	}
}

// GetDomain returns the value of the Domain field in CreateDomainRequest_SdkV2 as
// a Domain_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDomainRequest_SdkV2) GetDomain(ctx context.Context) (Domain_SdkV2, bool) {
	var e Domain_SdkV2
	if m.Domain.IsNull() || m.Domain.IsUnknown() {
		return e, false
	}
	var v []Domain_SdkV2
	d := m.Domain.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDomain sets the value of the Domain field in CreateDomainRequest_SdkV2.
func (m *CreateDomainRequest_SdkV2) SetDomain(ctx context.Context, v Domain_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain"]
	m.Domain = types.ListValueMust(t, vs)
}

type DeleteDomainRequest_SdkV2 struct {
	// When false (default), DeleteDomain is rejected with FAILED_PRECONDITION
	// if the domain still has Glossary pages. When true, those pages are
	// deleted first and then the domain is removed.
	Force types.Bool `tfsdk:"-"`
	// Full resource name of the domain to delete. Format: `domains/{domain_id}`
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDomainRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDomainRequest_SdkV2) {
	if !from.Force.IsUnknown() && !from.Force.IsNull() {
		// Force is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Force = from.Force
	}
}

func (to *DeleteDomainRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDomainRequest_SdkV2) {
	if !from.Force.IsUnknown() && !from.Force.IsNull() {
		// Force is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Force = from.Force
	}
}

func (m DeleteDomainRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["force"] = attrs["force"].SetOptional()
	attrs["force"] = attrs["force"].SetComputed()
	attrs["force"] = attrs["force"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDomainRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDomainRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDomainRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDomainRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": m.Force,
			"name":  m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDomainRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

type Domain_SdkV2 struct {
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
	Icon types.List `tfsdk:"icon"`
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
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Domain_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Domain_SdkV2) {
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
}

func (to *Domain_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Domain_SdkV2) {
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
}

func (m Domain_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["business_owner_ids"] = attrs["business_owner_ids"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["domain_id"] = attrs["domain_id"].SetComputed()
	attrs["draft"] = attrs["draft"].SetOptional()
	attrs["draft"] = attrs["draft"].SetComputed()
	attrs["draft"] = attrs["draft"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_draft"] = attrs["effective_draft"].SetComputed()
	attrs["icon"] = attrs["icon"].SetOptional()
	attrs["icon"] = attrs["icon"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["parent_domain_id"] = attrs["parent_domain_id"].SetOptional()
	attrs["parent_domain_id"] = attrs["parent_domain_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["subtitle"] = attrs["subtitle"].SetOptional()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["technical_owner_ids"] = attrs["technical_owner_ids"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Domain.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Domain_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"business_owner_ids":  reflect.TypeOf(types.Int64{}),
		"icon":                reflect.TypeOf(DomainIcon_SdkV2{}),
		"technical_owner_ids": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Domain_SdkV2
// only implements ToObjectValue() and Type().
func (m Domain_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
		})
}

// Type implements basetypes.ObjectValuable.
func (m Domain_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"business_owner_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"create_time":     timetypes.RFC3339{}.Type(ctx),
			"description":     types.StringType,
			"domain_id":       types.StringType,
			"draft":           types.BoolType,
			"effective_draft": types.BoolType,
			"icon": basetypes.ListType{
				ElemType: DomainIcon_SdkV2{}.Type(ctx),
			},
			"name":             types.StringType,
			"parent_domain_id": types.StringType,
			"subtitle":         types.StringType,
			"tag_key":          types.StringType,
			"technical_owner_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetBusinessOwnerIds returns the value of the BusinessOwnerIds field in Domain_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Domain_SdkV2) GetBusinessOwnerIds(ctx context.Context) ([]types.Int64, bool) {
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

// SetBusinessOwnerIds sets the value of the BusinessOwnerIds field in Domain_SdkV2.
func (m *Domain_SdkV2) SetBusinessOwnerIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["business_owner_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.BusinessOwnerIds = types.ListValueMust(t, vs)
}

// GetIcon returns the value of the Icon field in Domain_SdkV2 as
// a DomainIcon_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Domain_SdkV2) GetIcon(ctx context.Context) (DomainIcon_SdkV2, bool) {
	var e DomainIcon_SdkV2
	if m.Icon.IsNull() || m.Icon.IsUnknown() {
		return e, false
	}
	var v []DomainIcon_SdkV2
	d := m.Icon.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIcon sets the value of the Icon field in Domain_SdkV2.
func (m *Domain_SdkV2) SetIcon(ctx context.Context, v DomainIcon_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["icon"]
	m.Icon = types.ListValueMust(t, vs)
}

// GetTechnicalOwnerIds returns the value of the TechnicalOwnerIds field in Domain_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Domain_SdkV2) GetTechnicalOwnerIds(ctx context.Context) ([]types.Int64, bool) {
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

// SetTechnicalOwnerIds sets the value of the TechnicalOwnerIds field in Domain_SdkV2.
func (m *Domain_SdkV2) SetTechnicalOwnerIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["technical_owner_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TechnicalOwnerIds = types.ListValueMust(t, vs)
}

// Icon configuration for a domain.
type DomainIcon_SdkV2 struct {
	// Hex color code with # prefix (e.g., "#FF5733").
	Color types.String `tfsdk:"color"`

	Name types.String `tfsdk:"name"`
}

func (to *DomainIcon_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DomainIcon_SdkV2) {
}

func (to *DomainIcon_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DomainIcon_SdkV2) {
}

func (m DomainIcon_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["color"] = attrs["color"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DomainIcon.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DomainIcon_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DomainIcon_SdkV2
// only implements ToObjectValue() and Type().
func (m DomainIcon_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"color": m.Color,
			"name":  m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DomainIcon_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"color": types.StringType,
			"name":  types.StringType,
		},
	}
}

type GetDomainRequest_SdkV2 struct {
	// Full resource name of the domain to retrieve. Format:
	// `domains/{domain_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetDomainRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDomainRequest_SdkV2) {
}

func (to *GetDomainRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDomainRequest_SdkV2) {
}

func (m GetDomainRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDomainRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDomainRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDomainRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDomainRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDomainRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListDomainsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Filter by parent domain. - Absent: return all domains regardless of
	// hierarchy. - Present: return only direct children of the specified
	// domain.
	ParentDomainId types.String `tfsdk:"-"`
}

func (to *ListDomainsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDomainsRequest_SdkV2) {
}

func (to *ListDomainsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDomainsRequest_SdkV2) {
}

func (m ListDomainsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["parent_domain_id"] = attrs["parent_domain_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDomainsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDomainsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDomainsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDomainsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":        m.PageSize,
			"page_token":       m.PageToken,
			"parent_domain_id": m.ParentDomainId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDomainsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":        types.Int64Type,
			"page_token":       types.StringType,
			"parent_domain_id": types.StringType,
		},
	}
}

type ListDomainsResponse_SdkV2 struct {
	Domains types.List `tfsdk:"domains"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDomainsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDomainsResponse_SdkV2) {
	if !from.Domains.IsNull() && !from.Domains.IsUnknown() && to.Domains.IsNull() && len(from.Domains.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Domains, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Domains = from.Domains
	}
	if !from.Domains.IsNull() && !from.Domains.IsUnknown() {
		if toDomains, ok := to.GetDomains(ctx); ok {
			if fromDomains, ok := from.GetDomains(ctx); ok {
				// Recursively sync the fields of each Domains element by position.
				for i := range toDomains {
					if i < len(fromDomains) {
						toDomains[i].SyncFieldsDuringCreateOrUpdate(ctx, fromDomains[i])
					}
				}
				to.SetDomains(ctx, toDomains)
			}
		}
	}
}

func (to *ListDomainsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDomainsResponse_SdkV2) {
	if !from.Domains.IsNull() && !from.Domains.IsUnknown() && to.Domains.IsNull() && len(from.Domains.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Domains, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Domains = from.Domains
	}
	if !from.Domains.IsNull() && !from.Domains.IsUnknown() {
		if toDomains, ok := to.GetDomains(ctx); ok {
			if fromDomains, ok := from.GetDomains(ctx); ok {
				for i := range toDomains {
					if i < len(fromDomains) {
						toDomains[i].SyncFieldsDuringRead(ctx, fromDomains[i])
					}
				}
				to.SetDomains(ctx, toDomains)
			}
		}
	}
}

func (m ListDomainsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["domains"] = attrs["domains"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDomainsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDomainsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domains": reflect.TypeOf(Domain_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDomainsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDomainsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domains":         m.Domains,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDomainsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"domains": basetypes.ListType{
				ElemType: Domain_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDomains returns the value of the Domains field in ListDomainsResponse_SdkV2 as
// a slice of Domain_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDomainsResponse_SdkV2) GetDomains(ctx context.Context) ([]Domain_SdkV2, bool) {
	if m.Domains.IsNull() || m.Domains.IsUnknown() {
		return nil, false
	}
	var v []Domain_SdkV2
	d := m.Domains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomains sets the value of the Domains field in ListDomainsResponse_SdkV2.
func (m *ListDomainsResponse_SdkV2) SetDomains(ctx context.Context, v []Domain_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domains"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Domains = types.ListValueMust(t, vs)
}

type UpdateDomainRequest_SdkV2 struct {
	Domain types.List `tfsdk:"domain"`
	// Full resource name of the domain. The primary identifier for this
	// resource. Format: `domains/{domain_id}` Identifies the domain on get,
	// update, and delete. Not an input on create — to choose the id, set
	// `CreateDomainRequest.domain_id`.
	Name types.String `tfsdk:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDomainRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDomainRequest_SdkV2) {
	if !from.Domain.IsNull() && !from.Domain.IsUnknown() {
		if toDomain, ok := to.GetDomain(ctx); ok {
			if fromDomain, ok := from.GetDomain(ctx); ok {
				// Recursively sync the fields of Domain
				toDomain.SyncFieldsDuringCreateOrUpdate(ctx, fromDomain)
				to.SetDomain(ctx, toDomain)
			}
		}
	}
}

func (to *UpdateDomainRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDomainRequest_SdkV2) {
	if !from.Domain.IsNull() && !from.Domain.IsUnknown() {
		if toDomain, ok := to.GetDomain(ctx); ok {
			if fromDomain, ok := from.GetDomain(ctx); ok {
				toDomain.SyncFieldsDuringRead(ctx, fromDomain)
				to.SetDomain(ctx, toDomain)
			}
		}
	}
}

func (m UpdateDomainRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["domain"] = attrs["domain"].SetRequired()
	attrs["domain"] = attrs["domain"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDomainRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDomainRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain": reflect.TypeOf(Domain_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDomainRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDomainRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domain":      m.Domain,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDomainRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"domain": basetypes.ListType{
				ElemType: Domain_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetDomain returns the value of the Domain field in UpdateDomainRequest_SdkV2 as
// a Domain_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDomainRequest_SdkV2) GetDomain(ctx context.Context) (Domain_SdkV2, bool) {
	var e Domain_SdkV2
	if m.Domain.IsNull() || m.Domain.IsUnknown() {
		return e, false
	}
	var v []Domain_SdkV2
	d := m.Domain.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDomain sets the value of the Domain field in UpdateDomainRequest_SdkV2.
func (m *UpdateDomainRequest_SdkV2) SetDomain(ctx context.Context, v Domain_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain"]
	m.Domain = types.ListValueMust(t, vs)
}
