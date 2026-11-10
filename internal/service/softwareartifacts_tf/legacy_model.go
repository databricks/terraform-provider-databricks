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
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GetSoftwareArtifactRequest_SdkV2 struct {
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetSoftwareArtifactRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSoftwareArtifactRequest_SdkV2) {
}

func (to *GetSoftwareArtifactRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSoftwareArtifactRequest_SdkV2) {
}

func (m GetSoftwareArtifactRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSoftwareArtifactRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSoftwareArtifactRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSoftwareArtifactRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSoftwareArtifactRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSoftwareArtifactVersionRequest_SdkV2 struct {
	// Resource name of the version, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}/versions/{version}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetSoftwareArtifactVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSoftwareArtifactVersionRequest_SdkV2) {
}

func (to *GetSoftwareArtifactVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSoftwareArtifactVersionRequest_SdkV2) {
}

func (m GetSoftwareArtifactVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSoftwareArtifactVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSoftwareArtifactVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSoftwareArtifactVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSoftwareArtifactVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Image acceleration information for a software artifact version.
type ImageAccelerationDetail_SdkV2 struct {
	// Image acceleration status for the version.
	Status types.String `tfsdk:"status"`
}

func (to *ImageAccelerationDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ImageAccelerationDetail_SdkV2) {
}

func (to *ImageAccelerationDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ImageAccelerationDetail_SdkV2) {
}

func (m ImageAccelerationDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ImageAccelerationDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ImageAccelerationDetail_SdkV2
// only implements ToObjectValue() and Type().
func (m ImageAccelerationDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ImageAccelerationDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type ListSoftwareArtifactVersionsRequest_SdkV2 struct {
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

func (to *ListSoftwareArtifactVersionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactVersionsRequest_SdkV2) {
}

func (to *ListSoftwareArtifactVersionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactVersionsRequest_SdkV2) {
}

func (m ListSoftwareArtifactVersionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSoftwareArtifactVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response containing a page of software artifact versions.
type ListSoftwareArtifactVersionsResponse_SdkV2 struct {
	// Opaque token to pass as the next request's page_token. Empty when there
	// are no more pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The versions of the requested software artifact.
	SoftwareArtifactVersions types.List `tfsdk:"software_artifact_versions"`
}

func (to *ListSoftwareArtifactVersionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactVersionsResponse_SdkV2) {
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

func (to *ListSoftwareArtifactVersionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactVersionsResponse_SdkV2) {
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

func (m ListSoftwareArtifactVersionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSoftwareArtifactVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifact_versions": reflect.TypeOf(SoftwareArtifactVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":            m.NextPageToken,
			"software_artifact_versions": m.SoftwareArtifactVersions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"software_artifact_versions": basetypes.ListType{
				ElemType: SoftwareArtifactVersion_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSoftwareArtifactVersions returns the value of the SoftwareArtifactVersions field in ListSoftwareArtifactVersionsResponse_SdkV2 as
// a slice of SoftwareArtifactVersion_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSoftwareArtifactVersionsResponse_SdkV2) GetSoftwareArtifactVersions(ctx context.Context) ([]SoftwareArtifactVersion_SdkV2, bool) {
	if m.SoftwareArtifactVersions.IsNull() || m.SoftwareArtifactVersions.IsUnknown() {
		return nil, false
	}
	var v []SoftwareArtifactVersion_SdkV2
	d := m.SoftwareArtifactVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSoftwareArtifactVersions sets the value of the SoftwareArtifactVersions field in ListSoftwareArtifactVersionsResponse_SdkV2.
func (m *ListSoftwareArtifactVersionsResponse_SdkV2) SetSoftwareArtifactVersions(ctx context.Context, v []SoftwareArtifactVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["software_artifact_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SoftwareArtifactVersions = types.ListValueMust(t, vs)
}

type ListSoftwareArtifactsRequest_SdkV2 struct {
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

func (to *ListSoftwareArtifactsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactsRequest_SdkV2) {
}

func (to *ListSoftwareArtifactsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactsRequest_SdkV2) {
}

func (m ListSoftwareArtifactsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSoftwareArtifactsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response containing a page of software artifacts.
type ListSoftwareArtifactsResponse_SdkV2 struct {
	// Opaque token to pass as the next request's page_token. Empty when there
	// are no more pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The software artifacts under the requested schema.
	SoftwareArtifacts types.List `tfsdk:"software_artifacts"`
}

func (to *ListSoftwareArtifactsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSoftwareArtifactsResponse_SdkV2) {
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

func (to *ListSoftwareArtifactsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSoftwareArtifactsResponse_SdkV2) {
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

func (m ListSoftwareArtifactsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSoftwareArtifactsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifacts": reflect.TypeOf(SoftwareArtifact_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSoftwareArtifactsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSoftwareArtifactsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":    m.NextPageToken,
			"software_artifacts": m.SoftwareArtifacts,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSoftwareArtifactsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"software_artifacts": basetypes.ListType{
				ElemType: SoftwareArtifact_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSoftwareArtifacts returns the value of the SoftwareArtifacts field in ListSoftwareArtifactsResponse_SdkV2 as
// a slice of SoftwareArtifact_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSoftwareArtifactsResponse_SdkV2) GetSoftwareArtifacts(ctx context.Context) ([]SoftwareArtifact_SdkV2, bool) {
	if m.SoftwareArtifacts.IsNull() || m.SoftwareArtifacts.IsUnknown() {
		return nil, false
	}
	var v []SoftwareArtifact_SdkV2
	d := m.SoftwareArtifacts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSoftwareArtifacts sets the value of the SoftwareArtifacts field in ListSoftwareArtifactsResponse_SdkV2.
func (m *ListSoftwareArtifactsResponse_SdkV2) SetSoftwareArtifacts(ctx context.Context, v []SoftwareArtifact_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["software_artifacts"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SoftwareArtifacts = types.ListValueMust(t, vs)
}

// Request to resolve a software artifact version from its tag alias.
type ResolveSoftwareArtifactVersionRequest_SdkV2 struct {
	// Image tag alias selecting the version.
	Alias types.String `tfsdk:"alias"`
	// Resource name of the software artifact containing the alias.
	Parent types.String `tfsdk:"-"`
}

func (to *ResolveSoftwareArtifactVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveSoftwareArtifactVersionRequest_SdkV2) {
}

func (to *ResolveSoftwareArtifactVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveSoftwareArtifactVersionRequest_SdkV2) {
}

func (m ResolveSoftwareArtifactVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResolveSoftwareArtifactVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveSoftwareArtifactVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveSoftwareArtifactVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":  m.Alias,
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveSoftwareArtifactVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":  types.StringType,
			"parent": types.StringType,
		},
	}
}

// Response containing the resolved software artifact version.
type ResolveSoftwareArtifactVersionResponse_SdkV2 struct {
	// The software artifact version selected by the alias.
	SoftwareArtifactVersion types.List `tfsdk:"software_artifact_version"`
}

func (to *ResolveSoftwareArtifactVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResolveSoftwareArtifactVersionResponse_SdkV2) {
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

func (to *ResolveSoftwareArtifactVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResolveSoftwareArtifactVersionResponse_SdkV2) {
	if !from.SoftwareArtifactVersion.IsNull() && !from.SoftwareArtifactVersion.IsUnknown() {
		if toSoftwareArtifactVersion, ok := to.GetSoftwareArtifactVersion(ctx); ok {
			if fromSoftwareArtifactVersion, ok := from.GetSoftwareArtifactVersion(ctx); ok {
				toSoftwareArtifactVersion.SyncFieldsDuringRead(ctx, fromSoftwareArtifactVersion)
				to.SetSoftwareArtifactVersion(ctx, toSoftwareArtifactVersion)
			}
		}
	}
}

func (m ResolveSoftwareArtifactVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["software_artifact_version"] = attrs["software_artifact_version"].SetComputed()
	attrs["software_artifact_version"] = attrs["software_artifact_version"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolveSoftwareArtifactVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResolveSoftwareArtifactVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifact_version": reflect.TypeOf(SoftwareArtifactVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolveSoftwareArtifactVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ResolveSoftwareArtifactVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"software_artifact_version": m.SoftwareArtifactVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResolveSoftwareArtifactVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"software_artifact_version": basetypes.ListType{
				ElemType: SoftwareArtifactVersion_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSoftwareArtifactVersion returns the value of the SoftwareArtifactVersion field in ResolveSoftwareArtifactVersionResponse_SdkV2 as
// a SoftwareArtifactVersion_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResolveSoftwareArtifactVersionResponse_SdkV2) GetSoftwareArtifactVersion(ctx context.Context) (SoftwareArtifactVersion_SdkV2, bool) {
	var e SoftwareArtifactVersion_SdkV2
	if m.SoftwareArtifactVersion.IsNull() || m.SoftwareArtifactVersion.IsUnknown() {
		return e, false
	}
	var v []SoftwareArtifactVersion_SdkV2
	d := m.SoftwareArtifactVersion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSoftwareArtifactVersion sets the value of the SoftwareArtifactVersion field in ResolveSoftwareArtifactVersionResponse_SdkV2.
func (m *ResolveSoftwareArtifactVersionResponse_SdkV2) SetSoftwareArtifactVersion(ctx context.Context, v SoftwareArtifactVersion_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["software_artifact_version"]
	m.SoftwareArtifactVersion = types.ListValueMust(t, vs)
}

// A software artifact in a Unity Catalog schema.
type SoftwareArtifact_SdkV2 struct {
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

func (to *SoftwareArtifact_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SoftwareArtifact_SdkV2) {
}

func (to *SoftwareArtifact_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SoftwareArtifact_SdkV2) {
}

func (m SoftwareArtifact_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SoftwareArtifact_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SoftwareArtifact_SdkV2
// only implements ToObjectValue() and Type().
func (m SoftwareArtifact_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SoftwareArtifact_SdkV2) Type(ctx context.Context) attr.Type {
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
type SoftwareArtifactVersion_SdkV2 struct {
	// Time at which this software artifact version was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Principal name of the software artifact version creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Content digest of the version, for example `sha256:...`.
	Digest types.String `tfsdk:"digest"`
	// Image acceleration information for this version.
	ImageAccelerationDetails types.List `tfsdk:"image_acceleration_details"`
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

func (to *SoftwareArtifactVersion_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SoftwareArtifactVersion_SdkV2) {
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

func (to *SoftwareArtifactVersion_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SoftwareArtifactVersion_SdkV2) {
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

func (m SoftwareArtifactVersion_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["digest"] = attrs["digest"].SetComputed()
	attrs["image_acceleration_details"] = attrs["image_acceleration_details"].SetComputed()
	attrs["image_acceleration_details"] = attrs["image_acceleration_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m SoftwareArtifactVersion_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"image_acceleration_details": reflect.TypeOf(ImageAccelerationDetail_SdkV2{}),
		"tags":                       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SoftwareArtifactVersion_SdkV2
// only implements ToObjectValue() and Type().
func (m SoftwareArtifactVersion_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SoftwareArtifactVersion_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"created_by":  types.StringType,
			"digest":      types.StringType,
			"image_acceleration_details": basetypes.ListType{
				ElemType: ImageAccelerationDetail_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
			"updated_by":  types.StringType,
		},
	}
}

// GetImageAccelerationDetails returns the value of the ImageAccelerationDetails field in SoftwareArtifactVersion_SdkV2 as
// a ImageAccelerationDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SoftwareArtifactVersion_SdkV2) GetImageAccelerationDetails(ctx context.Context) (ImageAccelerationDetail_SdkV2, bool) {
	var e ImageAccelerationDetail_SdkV2
	if m.ImageAccelerationDetails.IsNull() || m.ImageAccelerationDetails.IsUnknown() {
		return e, false
	}
	var v []ImageAccelerationDetail_SdkV2
	d := m.ImageAccelerationDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetImageAccelerationDetails sets the value of the ImageAccelerationDetails field in SoftwareArtifactVersion_SdkV2.
func (m *SoftwareArtifactVersion_SdkV2) SetImageAccelerationDetails(ctx context.Context, v ImageAccelerationDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["image_acceleration_details"]
	m.ImageAccelerationDetails = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in SoftwareArtifactVersion_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SoftwareArtifactVersion_SdkV2) GetTags(ctx context.Context) ([]types.String, bool) {
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

// SetTags sets the value of the Tags field in SoftwareArtifactVersion_SdkV2.
func (m *SoftwareArtifactVersion_SdkV2) SetTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type UpdateSoftwareArtifactRequest_SdkV2 struct {
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name types.String `tfsdk:"-"`
	// The software artifact to update. Its name identifies the existing
	// artifact; only owner may be changed.
	SoftwareArtifact types.List `tfsdk:"software_artifact"`
	// Fields to update, relative to `software_artifact`. Accepts `owner` or
	// `*`; owner is the only mutable field, so both select the same update.
	// Output-only values in the resource are ignored.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSoftwareArtifactRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSoftwareArtifactRequest_SdkV2) {
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

func (to *UpdateSoftwareArtifactRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSoftwareArtifactRequest_SdkV2) {
	if !from.SoftwareArtifact.IsNull() && !from.SoftwareArtifact.IsUnknown() {
		if toSoftwareArtifact, ok := to.GetSoftwareArtifact(ctx); ok {
			if fromSoftwareArtifact, ok := from.GetSoftwareArtifact(ctx); ok {
				toSoftwareArtifact.SyncFieldsDuringRead(ctx, fromSoftwareArtifact)
				to.SetSoftwareArtifact(ctx, toSoftwareArtifact)
			}
		}
	}
}

func (m UpdateSoftwareArtifactRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["software_artifact"] = attrs["software_artifact"].SetRequired()
	attrs["software_artifact"] = attrs["software_artifact"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateSoftwareArtifactRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"software_artifact": reflect.TypeOf(SoftwareArtifact_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSoftwareArtifactRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSoftwareArtifactRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              m.Name,
			"software_artifact": m.SoftwareArtifact,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSoftwareArtifactRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"software_artifact": basetypes.ListType{
				ElemType: SoftwareArtifact_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetSoftwareArtifact returns the value of the SoftwareArtifact field in UpdateSoftwareArtifactRequest_SdkV2 as
// a SoftwareArtifact_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSoftwareArtifactRequest_SdkV2) GetSoftwareArtifact(ctx context.Context) (SoftwareArtifact_SdkV2, bool) {
	var e SoftwareArtifact_SdkV2
	if m.SoftwareArtifact.IsNull() || m.SoftwareArtifact.IsUnknown() {
		return e, false
	}
	var v []SoftwareArtifact_SdkV2
	d := m.SoftwareArtifact.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSoftwareArtifact sets the value of the SoftwareArtifact field in UpdateSoftwareArtifactRequest_SdkV2.
func (m *UpdateSoftwareArtifactRequest_SdkV2) SetSoftwareArtifact(ctx context.Context, v SoftwareArtifact_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["software_artifact"]
	m.SoftwareArtifact = types.ListValueMust(t, vs)
}
