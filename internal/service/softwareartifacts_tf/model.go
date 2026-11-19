// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package softwareartifacts_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GetSoftwareArtifactRequest struct {
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetSoftwareArtifactRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSoftwareArtifactRequest) {
}

func (to *GetSoftwareArtifactRequest) SyncFieldsDuringRead(ctx context.Context, from GetSoftwareArtifactRequest) {
}

func (m GetSoftwareArtifactRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSoftwareArtifactRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSoftwareArtifactRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSoftwareArtifactRequest
// only implements ToObjectValue() and Type().
func (m GetSoftwareArtifactRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSoftwareArtifactRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSoftwareArtifactVersionRequest struct {
	// Resource name of the version, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}/versions/{version}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetSoftwareArtifactVersionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSoftwareArtifactVersionRequest) {
}

func (to *GetSoftwareArtifactVersionRequest) SyncFieldsDuringRead(ctx context.Context, from GetSoftwareArtifactVersionRequest) {
}

func (m GetSoftwareArtifactVersionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSoftwareArtifactVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSoftwareArtifactVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSoftwareArtifactVersionRequest
// only implements ToObjectValue() and Type().
func (m GetSoftwareArtifactVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSoftwareArtifactVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Image acceleration information for a software artifact version.
type ImageAccelerationDetail struct {
	// Image acceleration status for the version.
	Status types.String `tfsdk:"status"`
}

func (to *ImageAccelerationDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ImageAccelerationDetail) {
}

func (to *ImageAccelerationDetail) SyncFieldsDuringRead(ctx context.Context, from ImageAccelerationDetail) {
}

func (m ImageAccelerationDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ImageAccelerationDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ImageAccelerationDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ImageAccelerationDetail
// only implements ToObjectValue() and Type().
func (m ImageAccelerationDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ImageAccelerationDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type ListSoftwareArtifactVersionsRequest struct {
	// Maximum number of versions to return. If omitted or set to 0, the page
	// size defaults to 50. The maximum value is 50; values above 50 or below 0
	// are rejected.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque continuation token from a previous response's next_page_token.
	// Omit for the first page.
	PageToken types.String `tfsdk:"-"`
	// Resource name of the parent artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListSoftwareArtifactVersionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactVersionsRequest) {
}

func (to *ListSoftwareArtifactVersionsRequest) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactVersionsRequest) {
}

func (m ListSoftwareArtifactVersionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSoftwareArtifactVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSoftwareArtifactVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactVersionsRequest
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactVersionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response containing a page of software artifact versions.
type ListSoftwareArtifactVersionsResponse struct {
	// Opaque token to pass as the next request's page_token. Empty when there
	// are no more pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The versions of the requested software artifact.
	SoftwareArtifactVersions types.List `tfsdk:"software_artifact_versions"`
}

func (to *ListSoftwareArtifactVersionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactVersionsResponse) {
	if !from.SoftwareArtifactVersions.IsNull() && !from.SoftwareArtifactVersions.IsUnknown() && to.SoftwareArtifactVersions.IsNull() && len(from.SoftwareArtifactVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SoftwareArtifactVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SoftwareArtifactVersions = from.SoftwareArtifactVersions
	}
	if !from.SoftwareArtifactVersions.IsNull() && !from.SoftwareArtifactVersions.IsUnknown() {
		if toSoftwareArtifactVersions, ok := to.GetSoftwareArtifactVersions(ctx); ok {
			if fromSoftwareArtifactVersions, ok := from.GetSoftwareArtifactVersions(ctx); ok {
				// Recursively sync the fields of each SoftwareArtifactVersions element by position.
				for i := range toSoftwareArtifactVersions {
					if i < len(fromSoftwareArtifactVersions) {
						toSoftwareArtifactVersions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSoftwareArtifactVersions[i])
					}
				}
				to.SetSoftwareArtifactVersions(ctx, toSoftwareArtifactVersions)
			}
		}
	}
}

func (to *ListSoftwareArtifactVersionsResponse) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactVersionsResponse) {
	if !from.SoftwareArtifactVersions.IsNull() && !from.SoftwareArtifactVersions.IsUnknown() && to.SoftwareArtifactVersions.IsNull() && len(from.SoftwareArtifactVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SoftwareArtifactVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SoftwareArtifactVersions = from.SoftwareArtifactVersions
	}
	if !from.SoftwareArtifactVersions.IsNull() && !from.SoftwareArtifactVersions.IsUnknown() {
		if toSoftwareArtifactVersions, ok := to.GetSoftwareArtifactVersions(ctx); ok {
			if fromSoftwareArtifactVersions, ok := from.GetSoftwareArtifactVersions(ctx); ok {
				for i := range toSoftwareArtifactVersions {
					if i < len(fromSoftwareArtifactVersions) {
						toSoftwareArtifactVersions[i].SyncFieldsDuringRead(ctx, fromSoftwareArtifactVersions[i])
					}
				}
				to.SetSoftwareArtifactVersions(ctx, toSoftwareArtifactVersions)
			}
		}
	}
}

func (m ListSoftwareArtifactVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["software_artifact_versions"] = attrs["software_artifact_versions"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSoftwareArtifactVersionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSoftwareArtifactVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifact_versions": reflect.TypeOf(SoftwareArtifactVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactVersionsResponse
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":            m.NextPageToken,
			"software_artifact_versions": m.SoftwareArtifactVersions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactVersionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"software_artifact_versions": basetypes.ListType{
				ElemType: SoftwareArtifactVersion{}.Type(ctx),
			},
		},
	}
}

// GetSoftwareArtifactVersions returns the value of the SoftwareArtifactVersions field in ListSoftwareArtifactVersionsResponse as
// a slice of SoftwareArtifactVersion values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSoftwareArtifactVersionsResponse) GetSoftwareArtifactVersions(ctx context.Context) ([]SoftwareArtifactVersion, bool) {
	if m.SoftwareArtifactVersions.IsNull() || m.SoftwareArtifactVersions.IsUnknown() {
		return nil, false
	}
	var v []SoftwareArtifactVersion
	d := m.SoftwareArtifactVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSoftwareArtifactVersions sets the value of the SoftwareArtifactVersions field in ListSoftwareArtifactVersionsResponse.
func (m *ListSoftwareArtifactVersionsResponse) SetSoftwareArtifactVersions(ctx context.Context, v []SoftwareArtifactVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["software_artifact_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SoftwareArtifactVersions = types.ListValueMust(t, vs)
}

type ListSoftwareArtifactsRequest struct {
	// Maximum number of software artifacts to return. If omitted or set to 0,
	// the page size defaults to 50. The maximum value is 50; values above 50 or
	// below 0 are rejected.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque continuation token from a previous response's next_page_token.
	// Omit for the first page.
	PageToken types.String `tfsdk:"-"`
	// Resource name of the parent schema, of the form
	// `schemas/{catalog}.{schema}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListSoftwareArtifactsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactsRequest) {
}

func (to *ListSoftwareArtifactsRequest) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactsRequest) {
}

func (m ListSoftwareArtifactsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSoftwareArtifactsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSoftwareArtifactsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactsRequest
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response containing a page of software artifacts.
type ListSoftwareArtifactsResponse struct {
	// Opaque token to pass as the next request's page_token. Empty when there
	// are no more pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The software artifacts under the requested schema.
	SoftwareArtifacts types.List `tfsdk:"software_artifacts"`
}

func (to *ListSoftwareArtifactsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactsResponse) {
	if !from.SoftwareArtifacts.IsNull() && !from.SoftwareArtifacts.IsUnknown() && to.SoftwareArtifacts.IsNull() && len(from.SoftwareArtifacts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SoftwareArtifacts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SoftwareArtifacts = from.SoftwareArtifacts
	}
	if !from.SoftwareArtifacts.IsNull() && !from.SoftwareArtifacts.IsUnknown() {
		if toSoftwareArtifacts, ok := to.GetSoftwareArtifacts(ctx); ok {
			if fromSoftwareArtifacts, ok := from.GetSoftwareArtifacts(ctx); ok {
				// Recursively sync the fields of each SoftwareArtifacts element by position.
				for i := range toSoftwareArtifacts {
					if i < len(fromSoftwareArtifacts) {
						toSoftwareArtifacts[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSoftwareArtifacts[i])
					}
				}
				to.SetSoftwareArtifacts(ctx, toSoftwareArtifacts)
			}
		}
	}
}

func (to *ListSoftwareArtifactsResponse) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactsResponse) {
	if !from.SoftwareArtifacts.IsNull() && !from.SoftwareArtifacts.IsUnknown() && to.SoftwareArtifacts.IsNull() && len(from.SoftwareArtifacts.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SoftwareArtifacts, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SoftwareArtifacts = from.SoftwareArtifacts
	}
	if !from.SoftwareArtifacts.IsNull() && !from.SoftwareArtifacts.IsUnknown() {
		if toSoftwareArtifacts, ok := to.GetSoftwareArtifacts(ctx); ok {
			if fromSoftwareArtifacts, ok := from.GetSoftwareArtifacts(ctx); ok {
				for i := range toSoftwareArtifacts {
					if i < len(fromSoftwareArtifacts) {
						toSoftwareArtifacts[i].SyncFieldsDuringRead(ctx, fromSoftwareArtifacts[i])
					}
				}
				to.SetSoftwareArtifacts(ctx, toSoftwareArtifacts)
			}
		}
	}
}

func (m ListSoftwareArtifactsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["software_artifacts"] = attrs["software_artifacts"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSoftwareArtifactsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSoftwareArtifactsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifacts": reflect.TypeOf(SoftwareArtifact{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactsResponse
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":    m.NextPageToken,
			"software_artifacts": m.SoftwareArtifacts,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"software_artifacts": basetypes.ListType{
				ElemType: SoftwareArtifact{}.Type(ctx),
			},
		},
	}
}

// GetSoftwareArtifacts returns the value of the SoftwareArtifacts field in ListSoftwareArtifactsResponse as
// a slice of SoftwareArtifact values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSoftwareArtifactsResponse) GetSoftwareArtifacts(ctx context.Context) ([]SoftwareArtifact, bool) {
	if m.SoftwareArtifacts.IsNull() || m.SoftwareArtifacts.IsUnknown() {
		return nil, false
	}
	var v []SoftwareArtifact
	d := m.SoftwareArtifacts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSoftwareArtifacts sets the value of the SoftwareArtifacts field in ListSoftwareArtifactsResponse.
func (m *ListSoftwareArtifactsResponse) SetSoftwareArtifacts(ctx context.Context, v []SoftwareArtifact) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["software_artifacts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SoftwareArtifacts = types.ListValueMust(t, vs)
}

// Request to resolve a software artifact version from its tag alias.
type ResolveSoftwareArtifactVersionRequest struct {
	// Image tag alias selecting the version.
	Alias types.String `tfsdk:"alias"`
	// Resource name of the software artifact containing the alias.
	Parent types.String `tfsdk:"-"`
}

func (to *ResolveSoftwareArtifactVersionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveSoftwareArtifactVersionRequest) {
}

func (to *ResolveSoftwareArtifactVersionRequest) SyncFieldsDuringRead(ctx context.Context, from ResolveSoftwareArtifactVersionRequest) {
}

func (m ResolveSoftwareArtifactVersionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alias"] = attrs["alias"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveSoftwareArtifactVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveSoftwareArtifactVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveSoftwareArtifactVersionRequest
// only implements ToObjectValue() and Type().
func (m ResolveSoftwareArtifactVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":  m.Alias,
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveSoftwareArtifactVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":  types.StringType,
			"parent": types.StringType,
		},
	}
}

// Response containing the resolved software artifact version.
type ResolveSoftwareArtifactVersionResponse struct {
	// The software artifact version selected by the alias.
	SoftwareArtifactVersion types.Object `tfsdk:"software_artifact_version"`
}

func (to *ResolveSoftwareArtifactVersionResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveSoftwareArtifactVersionResponse) {
	if !from.SoftwareArtifactVersion.IsNull() && !from.SoftwareArtifactVersion.IsUnknown() {
		if toSoftwareArtifactVersion, ok := to.GetSoftwareArtifactVersion(ctx); ok {
			if fromSoftwareArtifactVersion, ok := from.GetSoftwareArtifactVersion(ctx); ok {
				// Recursively sync the fields of SoftwareArtifactVersion
				toSoftwareArtifactVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromSoftwareArtifactVersion)
				to.SetSoftwareArtifactVersion(ctx, toSoftwareArtifactVersion)
			}
		}
	}
}

func (to *ResolveSoftwareArtifactVersionResponse) SyncFieldsDuringRead(ctx context.Context, from ResolveSoftwareArtifactVersionResponse) {
	if !from.SoftwareArtifactVersion.IsNull() && !from.SoftwareArtifactVersion.IsUnknown() {
		if toSoftwareArtifactVersion, ok := to.GetSoftwareArtifactVersion(ctx); ok {
			if fromSoftwareArtifactVersion, ok := from.GetSoftwareArtifactVersion(ctx); ok {
				toSoftwareArtifactVersion.SyncFieldsDuringRead(ctx, fromSoftwareArtifactVersion)
				to.SetSoftwareArtifactVersion(ctx, toSoftwareArtifactVersion)
			}
		}
	}
}

func (m ResolveSoftwareArtifactVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["software_artifact_version"] = attrs["software_artifact_version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveSoftwareArtifactVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveSoftwareArtifactVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifact_version": reflect.TypeOf(SoftwareArtifactVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveSoftwareArtifactVersionResponse
// only implements ToObjectValue() and Type().
func (m ResolveSoftwareArtifactVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"software_artifact_version": m.SoftwareArtifactVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveSoftwareArtifactVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"software_artifact_version": SoftwareArtifactVersion{}.Type(ctx),
		},
	}
}

// GetSoftwareArtifactVersion returns the value of the SoftwareArtifactVersion field in ResolveSoftwareArtifactVersionResponse as
// a SoftwareArtifactVersion value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResolveSoftwareArtifactVersionResponse) GetSoftwareArtifactVersion(ctx context.Context) (SoftwareArtifactVersion, bool) {
	var e SoftwareArtifactVersion
	if m.SoftwareArtifactVersion.IsNull() || m.SoftwareArtifactVersion.IsUnknown() {
		return e, false
	}
	var v SoftwareArtifactVersion
	d := m.SoftwareArtifactVersion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSoftwareArtifactVersion sets the value of the SoftwareArtifactVersion field in ResolveSoftwareArtifactVersionResponse.
func (m *ResolveSoftwareArtifactVersionResponse) SetSoftwareArtifactVersion(ctx context.Context, v SoftwareArtifactVersion) {
	vs := v.ToObjectValue(ctx)
	m.SoftwareArtifactVersion = vs
}

// A software artifact in a Unity Catalog schema.
type SoftwareArtifact struct {
	// Time at which this software artifact was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Principal name of the software artifact creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name types.String `tfsdk:"name"`
	// Principal name (a user, service principal, or group) that owns this
	// artifact.
	Owner types.String `tfsdk:"owner"`
	// Time at which this software artifact was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Principal name that last updated this software artifact.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *SoftwareArtifact) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SoftwareArtifact) {
}

func (to *SoftwareArtifact) SyncFieldsDuringRead(ctx context.Context, from SoftwareArtifact) {
}

func (m SoftwareArtifact) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SoftwareArtifact.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SoftwareArtifact) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SoftwareArtifact
// only implements ToObjectValue() and Type().
func (m SoftwareArtifact) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"created_by":  m.CreatedBy,
			"name":        m.Name,
			"owner":       m.Owner,
			"update_time": m.UpdateTime,
			"updated_by":  m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SoftwareArtifact) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"created_by":  types.StringType,
			"name":        types.StringType,
			"owner":       types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by":  types.StringType,
		},
	}
}

// An immutable version of a software artifact, identified by its content
// digest.
type SoftwareArtifactVersion struct {
	// Time at which this software artifact version was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Principal name of the software artifact version creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Content digest of the version, for example `sha256:...`.
	Digest types.String `tfsdk:"digest"`
	// Image acceleration information for this version.
	ImageAccelerationDetails types.Object `tfsdk:"image_acceleration_details"`
	// Resource name of the version, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}/versions/{version}`.
	Name types.String `tfsdk:"name"`
	// Tags (aliases) that currently point at this version, for example `latest`
	// or `v1.2.3`. Bounded at 10 per version; a version pointed at by more tags
	// returns an arbitrary 10 of them.
	Tags types.List `tfsdk:"tags"`
	// Time at which this software artifact version was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Principal name that last updated this software artifact version.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *SoftwareArtifactVersion) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SoftwareArtifactVersion) {
	if !from.ImageAccelerationDetails.IsNull() && !from.ImageAccelerationDetails.IsUnknown() {
		if toImageAccelerationDetails, ok := to.GetImageAccelerationDetails(ctx); ok {
			if fromImageAccelerationDetails, ok := from.GetImageAccelerationDetails(ctx); ok {
				// Recursively sync the fields of ImageAccelerationDetails
				toImageAccelerationDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromImageAccelerationDetails)
				to.SetImageAccelerationDetails(ctx, toImageAccelerationDetails)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *SoftwareArtifactVersion) SyncFieldsDuringRead(ctx context.Context, from SoftwareArtifactVersion) {
	if !from.ImageAccelerationDetails.IsNull() && !from.ImageAccelerationDetails.IsUnknown() {
		if toImageAccelerationDetails, ok := to.GetImageAccelerationDetails(ctx); ok {
			if fromImageAccelerationDetails, ok := from.GetImageAccelerationDetails(ctx); ok {
				toImageAccelerationDetails.SyncFieldsDuringRead(ctx, fromImageAccelerationDetails)
				to.SetImageAccelerationDetails(ctx, toImageAccelerationDetails)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m SoftwareArtifactVersion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["digest"] = attrs["digest"].SetComputed()
	attrs["image_acceleration_details"] = attrs["image_acceleration_details"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["tags"] = attrs["tags"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SoftwareArtifactVersion.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SoftwareArtifactVersion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"image_acceleration_details": reflect.TypeOf(ImageAccelerationDetail{}),
		"tags":                       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SoftwareArtifactVersion
// only implements ToObjectValue() and Type().
func (m SoftwareArtifactVersion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":                m.CreateTime,
			"created_by":                 m.CreatedBy,
			"digest":                     m.Digest,
			"image_acceleration_details": m.ImageAccelerationDetails,
			"name":                       m.Name,
			"tags":                       m.Tags,
			"update_time":                m.UpdateTime,
			"updated_by":                 m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SoftwareArtifactVersion) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":                timetypes.RFC3339{}.Type(ctx),
			"created_by":                 types.StringType,
			"digest":                     types.StringType,
			"image_acceleration_details": ImageAccelerationDetail{}.Type(ctx),
			"name":                       types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by":  types.StringType,
		},
	}
}

// GetImageAccelerationDetails returns the value of the ImageAccelerationDetails field in SoftwareArtifactVersion as
// a ImageAccelerationDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *SoftwareArtifactVersion) GetImageAccelerationDetails(ctx context.Context) (ImageAccelerationDetail, bool) {
	var e ImageAccelerationDetail
	if m.ImageAccelerationDetails.IsNull() || m.ImageAccelerationDetails.IsUnknown() {
		return e, false
	}
	var v ImageAccelerationDetail
	d := m.ImageAccelerationDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetImageAccelerationDetails sets the value of the ImageAccelerationDetails field in SoftwareArtifactVersion.
func (m *SoftwareArtifactVersion) SetImageAccelerationDetails(ctx context.Context, v ImageAccelerationDetail) {
	vs := v.ToObjectValue(ctx)
	m.ImageAccelerationDetails = vs
}

// GetTags returns the value of the Tags field in SoftwareArtifactVersion as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SoftwareArtifactVersion) GetTags(ctx context.Context) ([]types.String, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in SoftwareArtifactVersion.
func (m *SoftwareArtifactVersion) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type UpdateSoftwareArtifactRequest struct {
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name types.String `tfsdk:"-"`
	// The software artifact to update. Its name identifies the existing
	// artifact; only owner may be changed.
	SoftwareArtifact types.Object `tfsdk:"software_artifact"`
	// Fields to update, relative to `software_artifact`. Accepts `owner` or
	// `*`; owner is the only mutable field, so both select the same update.
	// Output-only values in the resource are ignored.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSoftwareArtifactRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSoftwareArtifactRequest) {
	if !from.SoftwareArtifact.IsNull() && !from.SoftwareArtifact.IsUnknown() {
		if toSoftwareArtifact, ok := to.GetSoftwareArtifact(ctx); ok {
			if fromSoftwareArtifact, ok := from.GetSoftwareArtifact(ctx); ok {
				// Recursively sync the fields of SoftwareArtifact
				toSoftwareArtifact.SyncFieldsDuringCreateOrUpdate(ctx, fromSoftwareArtifact)
				to.SetSoftwareArtifact(ctx, toSoftwareArtifact)
			}
		}
	}
}

func (to *UpdateSoftwareArtifactRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateSoftwareArtifactRequest) {
	if !from.SoftwareArtifact.IsNull() && !from.SoftwareArtifact.IsUnknown() {
		if toSoftwareArtifact, ok := to.GetSoftwareArtifact(ctx); ok {
			if fromSoftwareArtifact, ok := from.GetSoftwareArtifact(ctx); ok {
				toSoftwareArtifact.SyncFieldsDuringRead(ctx, fromSoftwareArtifact)
				to.SetSoftwareArtifact(ctx, toSoftwareArtifact)
			}
		}
	}
}

func (m UpdateSoftwareArtifactRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["software_artifact"] = attrs["software_artifact"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSoftwareArtifactRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSoftwareArtifactRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifact": reflect.TypeOf(SoftwareArtifact{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSoftwareArtifactRequest
// only implements ToObjectValue() and Type().
func (m UpdateSoftwareArtifactRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              m.Name,
			"software_artifact": m.SoftwareArtifact,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSoftwareArtifactRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"software_artifact": SoftwareArtifact{}.Type(ctx),
			"update_mask":       types.StringType,
		},
	}
}

// GetSoftwareArtifact returns the value of the SoftwareArtifact field in UpdateSoftwareArtifactRequest as
// a SoftwareArtifact value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSoftwareArtifactRequest) GetSoftwareArtifact(ctx context.Context) (SoftwareArtifact, bool) {
	var e SoftwareArtifact
	if m.SoftwareArtifact.IsNull() || m.SoftwareArtifact.IsUnknown() {
		return e, false
	}
	var v SoftwareArtifact
	d := m.SoftwareArtifact.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSoftwareArtifact sets the value of the SoftwareArtifact field in UpdateSoftwareArtifactRequest.
func (m *UpdateSoftwareArtifactRequest) SetSoftwareArtifact(ctx context.Context, v SoftwareArtifact) {
	vs := v.ToObjectValue(ctx)
	m.SoftwareArtifact = vs
}
