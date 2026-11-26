// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package geniemigrate_tf

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

// The request to add workspace files to a migration project.
type AddProjectFilesRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name types.String `tfsdk:"-"`
	// The files to add, each identified by `workspace.path` or
	// `workspace.object_id`. Send at most 500 files per request. Duplicate
	// files, and requests that exceed a file-count limit, return
	// `INVALID_PARAMETER_VALUE`. Files already in the project aren't
	// overwritten. An empty list changes nothing.
	Sources types.List `tfsdk:"sources"`
}

func (to *AddProjectFilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddProjectFilesRequest) {
	if !from.Sources.IsNull() && !from.Sources.IsUnknown() {
		if toSources, ok := to.GetSources(ctx); ok {
			if fromSources, ok := from.GetSources(ctx); ok {
				// Recursively sync the fields of each Sources element by position.
				for i := range toSources {
					if i < len(fromSources) {
						toSources[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSources[i])
					}
				}
				to.SetSources(ctx, toSources)
			}
		}
	}
}

func (to *AddProjectFilesRequest) SyncFieldsDuringRead(ctx context.Context, from AddProjectFilesRequest) {
	if !from.Sources.IsNull() && !from.Sources.IsUnknown() {
		if toSources, ok := to.GetSources(ctx); ok {
			if fromSources, ok := from.GetSources(ctx); ok {
				for i := range toSources {
					if i < len(fromSources) {
						toSources[i].SyncFieldsDuringRead(ctx, fromSources[i])
					}
				}
				to.SetSources(ctx, toSources)
			}
		}
	}
}

func (m AddProjectFilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sources"] = attrs["sources"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddProjectFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AddProjectFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sources": reflect.TypeOf(SourceFileRef{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddProjectFilesRequest
// only implements ToObjectValue() and Type().
func (m AddProjectFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"sources": m.Sources,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AddProjectFilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"sources": basetypes.ListType{
				ElemType: SourceFileRef{}.Type(ctx),
			},
		},
	}
}

// GetSources returns the value of the Sources field in AddProjectFilesRequest as
// a slice of SourceFileRef values.
// If the field is unknown or null, the boolean return value is false.
func (m *AddProjectFilesRequest) GetSources(ctx context.Context) ([]SourceFileRef, bool) {
	if m.Sources.IsNull() || m.Sources.IsUnknown() {
		return nil, false
	}
	var v []SourceFileRef
	d := m.Sources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSources sets the value of the Sources field in AddProjectFilesRequest.
func (m *AddProjectFilesRequest) SetSources(ctx context.Context, v []SourceFileRef) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Sources = types.ListValueMust(t, vs)
}

// Lists the files from an `AddProjectFiles` request, including files that were
// already in the project.
type AddProjectFilesResponse struct {
	// The requested files, in request order, without conversion results. Call
	// `ListProjectFiles` or `GetProjectFile` to read their state.
	Files types.List `tfsdk:"files"`
}

func (to *AddProjectFilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddProjectFilesResponse) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				// Recursively sync the fields of each Files element by position.
				for i := range toFiles {
					if i < len(fromFiles) {
						toFiles[i].SyncFieldsDuringCreateOrUpdate(ctx, fromFiles[i])
					}
				}
				to.SetFiles(ctx, toFiles)
			}
		}
	}
}

func (to *AddProjectFilesResponse) SyncFieldsDuringRead(ctx context.Context, from AddProjectFilesResponse) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				for i := range toFiles {
					if i < len(fromFiles) {
						toFiles[i].SyncFieldsDuringRead(ctx, fromFiles[i])
					}
				}
				to.SetFiles(ctx, toFiles)
			}
		}
	}
}

func (m AddProjectFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["files"] = attrs["files"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddProjectFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AddProjectFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(ProjectFile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddProjectFilesResponse
// only implements ToObjectValue() and Type().
func (m AddProjectFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files": m.Files,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AddProjectFilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"files": basetypes.ListType{
				ElemType: ProjectFile{}.Type(ctx),
			},
		},
	}
}

// GetFiles returns the value of the Files field in AddProjectFilesResponse as
// a slice of ProjectFile values.
// If the field is unknown or null, the boolean return value is false.
func (m *AddProjectFilesResponse) GetFiles(ctx context.Context) ([]ProjectFile, bool) {
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return nil, false
	}
	var v []ProjectFile
	d := m.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in AddProjectFilesResponse.
func (m *AddProjectFilesResponse) SetFiles(ctx context.Context, v []ProjectFile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Files = types.ListValueMust(t, vs)
}

// The request message for `CancelOperation` method.
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name types.String `tfsdk:"-"`
}

func (to *CancelOperationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelOperationRequest) {
}

func (to *CancelOperationRequest) SyncFieldsDuringRead(ctx context.Context, from CancelOperationRequest) {
}

func (m CancelOperationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelOperationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelOperationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelOperationRequest
// only implements ToObjectValue() and Type().
func (m CancelOperationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelOperationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// A conversion of a migration project's files, with its state and the status of
// each file.
type ConversionRun struct {
	// The time the run was accepted, before its background job started.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The time the run ended. Absent while the run is active.
	EndTime timetypes.RFC3339 `tfsdk:"end_time"`
	// Why the run failed. Set only when `state` is
	// `CONVERSION_RUN_STATE_FAILED`.
	ErrorCode types.String `tfsdk:"error_code"`
	// The files in this run, at most 500, each with its status in this run.
	// Later conversions and reviews can change `ProjectFile.status` but don't
	// change this list.
	Files types.List `tfsdk:"files"`
	// The resource name of the project being converted, in the format
	// `projects/{project_id}`.
	Project types.String `tfsdk:"project"`
	// The skill folders that the run was started with. Empty if the run
	// discovers skills automatically.
	SkillPaths types.List `tfsdk:"skill_paths"`
	// The state of the run.
	State types.String `tfsdk:"state"`
}

func (to *ConversionRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConversionRun) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				// Recursively sync the fields of each Files element by position.
				for i := range toFiles {
					if i < len(fromFiles) {
						toFiles[i].SyncFieldsDuringCreateOrUpdate(ctx, fromFiles[i])
					}
				}
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.SkillPaths.IsNull() && !from.SkillPaths.IsUnknown() && to.SkillPaths.IsNull() && len(from.SkillPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SkillPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SkillPaths = from.SkillPaths
	}
}

func (to *ConversionRun) SyncFieldsDuringRead(ctx context.Context, from ConversionRun) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				for i := range toFiles {
					if i < len(fromFiles) {
						toFiles[i].SyncFieldsDuringRead(ctx, fromFiles[i])
					}
				}
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.SkillPaths.IsNull() && !from.SkillPaths.IsUnknown() && to.SkillPaths.IsNull() && len(from.SkillPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SkillPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SkillPaths = from.SkillPaths
	}
}

func (m ConversionRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["end_time"] = attrs["end_time"].SetComputed()
	attrs["error_code"] = attrs["error_code"].SetComputed()
	attrs["files"] = attrs["files"].SetComputed()
	attrs["project"] = attrs["project"].SetComputed()
	attrs["skill_paths"] = attrs["skill_paths"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConversionRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ConversionRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files":       reflect.TypeOf(ConversionRunFile{}),
		"skill_paths": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConversionRun
// only implements ToObjectValue() and Type().
func (m ConversionRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"end_time":    m.EndTime,
			"error_code":  m.ErrorCode,
			"files":       m.Files,
			"project":     m.Project,
			"skill_paths": m.SkillPaths,
			"state":       m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ConversionRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"end_time":    timetypes.RFC3339{}.Type(ctx),
			"error_code":  types.StringType,
			"files": basetypes.ListType{
				ElemType: ConversionRunFile{}.Type(ctx),
			},
			"project": types.StringType,
			"skill_paths": basetypes.ListType{
				ElemType: types.StringType,
			},
			"state": types.StringType,
		},
	}
}

// GetFiles returns the value of the Files field in ConversionRun as
// a slice of ConversionRunFile values.
// If the field is unknown or null, the boolean return value is false.
func (m *ConversionRun) GetFiles(ctx context.Context) ([]ConversionRunFile, bool) {
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return nil, false
	}
	var v []ConversionRunFile
	d := m.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in ConversionRun.
func (m *ConversionRun) SetFiles(ctx context.Context, v []ConversionRunFile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Files = types.ListValueMust(t, vs)
}

// GetSkillPaths returns the value of the SkillPaths field in ConversionRun as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ConversionRun) GetSkillPaths(ctx context.Context) ([]types.String, bool) {
	if m.SkillPaths.IsNull() || m.SkillPaths.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SkillPaths.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSkillPaths sets the value of the SkillPaths field in ConversionRun.
func (m *ConversionRun) SetSkillPaths(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["skill_paths"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SkillPaths = types.ListValueMust(t, vs)
}

// A file in a conversion run.
type ConversionRunFile struct {
	// The resource name of the file, in the format
	// `projects/{project_id}/files/{file_id}`.
	Name types.String `tfsdk:"name"`
	// The status of the file in this run. `ProjectFile.status` has the file's
	// latest status across all runs.
	Status types.String `tfsdk:"status"`
}

func (to *ConversionRunFile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConversionRunFile) {
}

func (to *ConversionRunFile) SyncFieldsDuringRead(ctx context.Context, from ConversionRunFile) {
}

func (m ConversionRunFile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConversionRunFile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ConversionRunFile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConversionRunFile
// only implements ToObjectValue() and Type().
func (m ConversionRunFile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   m.Name,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ConversionRunFile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"status": types.StringType,
		},
	}
}

// The request to convert files of a migration project.
type ConvertProjectFilesRequest struct {
	// The files to convert, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`. All files must belong to `parent`. At most 500
	// files; a workspace can set a lower limit, and requests over it return
	// `INVALID_PARAMETER_VALUE`.
	Names types.List `tfsdk:"names"`
	// The resource name of the project whose files to convert, in the format
	// `projects/{project_id}`. For example, `projects/123`.
	Parent types.String `tfsdk:"-"`
	// Custom skills that the conversion can use, as workspace paths of skill
	// folders. Each folder must contain a `SKILL.md` file and be directly
	// inside a skills folder: `<home>/.assistant/skills`,
	// `/Workspace/.assistant/skills`, or another configured skills folder. At
	// most 10 paths, which together must fit in 1000 bytes when encoded as a
	// JSON array. Larger values return `INVALID_PARAMETER_VALUE`. If empty, the
	// conversion discovers skills automatically.
	SkillPaths types.List `tfsdk:"skill_paths"`
	// Whether to convert each selected file from scratch instead of continuing
	// from its previous results (changes, to-do items, blockers, and notes).
	// Supported only for projects whose target framework is
	// `TARGET_FRAMEWORK_SQL`; other projects return `INVALID_PARAMETER_VALUE`.
	// Defaults to `false`.
	StartOver types.Bool `tfsdk:"start_over"`
}

func (to *ConvertProjectFilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConvertProjectFilesRequest) {
	if !from.Names.IsNull() && !from.Names.IsUnknown() && to.Names.IsNull() && len(from.Names.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Names, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Names = from.Names
	}
	if !from.SkillPaths.IsNull() && !from.SkillPaths.IsUnknown() && to.SkillPaths.IsNull() && len(from.SkillPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SkillPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SkillPaths = from.SkillPaths
	}
}

func (to *ConvertProjectFilesRequest) SyncFieldsDuringRead(ctx context.Context, from ConvertProjectFilesRequest) {
	if !from.Names.IsNull() && !from.Names.IsUnknown() && to.Names.IsNull() && len(from.Names.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Names, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Names = from.Names
	}
	if !from.SkillPaths.IsNull() && !from.SkillPaths.IsUnknown() && to.SkillPaths.IsNull() && len(from.SkillPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SkillPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SkillPaths = from.SkillPaths
	}
}

func (m ConvertProjectFilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["names"] = attrs["names"].SetOptional()
	attrs["skill_paths"] = attrs["skill_paths"].SetOptional()
	attrs["start_over"] = attrs["start_over"].SetOptional()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConvertProjectFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ConvertProjectFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"names":       reflect.TypeOf(types.String{}),
		"skill_paths": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConvertProjectFilesRequest
// only implements ToObjectValue() and Type().
func (m ConvertProjectFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"names":       m.Names,
			"parent":      m.Parent,
			"skill_paths": m.SkillPaths,
			"start_over":  m.StartOver,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ConvertProjectFilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"parent": types.StringType,
			"skill_paths": basetypes.ListType{
				ElemType: types.StringType,
			},
			"start_over": types.BoolType,
		},
	}
}

// GetNames returns the value of the Names field in ConvertProjectFilesRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ConvertProjectFilesRequest) GetNames(ctx context.Context) ([]types.String, bool) {
	if m.Names.IsNull() || m.Names.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Names.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNames sets the value of the Names field in ConvertProjectFilesRequest.
func (m *ConvertProjectFilesRequest) SetNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Names = types.ListValueMust(t, vs)
}

// GetSkillPaths returns the value of the SkillPaths field in ConvertProjectFilesRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ConvertProjectFilesRequest) GetSkillPaths(ctx context.Context) ([]types.String, bool) {
	if m.SkillPaths.IsNull() || m.SkillPaths.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SkillPaths.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSkillPaths sets the value of the SkillPaths field in ConvertProjectFilesRequest.
func (m *ConvertProjectFilesRequest) SetSkillPaths(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["skill_paths"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SkillPaths = types.ListValueMust(t, vs)
}

// The result of a completed conversion.
type ConvertProjectFilesResponse struct {
	// The completed conversion run. Each file's outcome is in `run.files`; a
	// completed run can include files that failed to convert.
	Run types.Object `tfsdk:"run"`
}

func (to *ConvertProjectFilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ConvertProjectFilesResponse) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				// Recursively sync the fields of Run
				toRun.SyncFieldsDuringCreateOrUpdate(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (to *ConvertProjectFilesResponse) SyncFieldsDuringRead(ctx context.Context, from ConvertProjectFilesResponse) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				toRun.SyncFieldsDuringRead(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (m ConvertProjectFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run"] = attrs["run"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConvertProjectFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ConvertProjectFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(ConversionRun{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConvertProjectFilesResponse
// only implements ToObjectValue() and Type().
func (m ConvertProjectFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": m.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ConvertProjectFilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": ConversionRun{}.Type(ctx),
		},
	}
}

// GetRun returns the value of the Run field in ConvertProjectFilesResponse as
// a ConversionRun value.
// If the field is unknown or null, the boolean return value is false.
func (m *ConvertProjectFilesResponse) GetRun(ctx context.Context) (ConversionRun, bool) {
	var e ConversionRun
	if m.Run.IsNull() || m.Run.IsUnknown() {
		return e, false
	}
	var v ConversionRun
	d := m.Run.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRun sets the value of the Run field in ConvertProjectFilesResponse.
func (m *ConvertProjectFilesResponse) SetRun(ctx context.Context, v ConversionRun) {
	vs := v.ToObjectValue(ctx)
	m.Run = vs
}

type CreateProjectRequest struct {
	// The project to create. The server sets `name` in the response.
	Project types.Object `tfsdk:"project"`
}

func (to *CreateProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				// Recursively sync the fields of Project
				toProject.SyncFieldsDuringCreateOrUpdate(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (to *CreateProjectRequest) SyncFieldsDuringRead(ctx context.Context, from CreateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m CreateProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProjectRequest
// only implements ToObjectValue() and Type().
func (m CreateProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project": m.Project,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project": Project{}.Type(ctx),
		},
	}
}

// GetProject returns the value of the Project field in CreateProjectRequest as
// a Project value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateProjectRequest) GetProject(ctx context.Context) (Project, bool) {
	var e Project
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v Project
	d := m.Project.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProject sets the value of the Project field in CreateProjectRequest.
func (m *CreateProjectRequest) SetProject(ctx context.Context, v Project) {
	vs := v.ToObjectValue(ctx)
	m.Project = vs
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details types.List `tfsdk:"details"`

	ErrorCode types.String `tfsdk:"error_code"`

	Message types.String `tfsdk:"message"`

	StackTrace types.String `tfsdk:"stack_trace"`
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringRead(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (m DatabricksServiceExceptionWithDetailsProto) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["error_code"] = attrs["error_code"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["stack_trace"] = attrs["stack_trace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksServiceExceptionWithDetailsProto.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabricksServiceExceptionWithDetailsProto) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"details": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksServiceExceptionWithDetailsProto
// only implements ToObjectValue() and Type().
func (m DatabricksServiceExceptionWithDetailsProto) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"details":     m.Details,
			"error_code":  m.ErrorCode,
			"message":     m.Message,
			"stack_trace": m.StackTrace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabricksServiceExceptionWithDetailsProto) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"details": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"error_code":  types.StringType,
			"message":     types.StringType,
			"stack_trace": types.StringType,
		},
	}
}

// GetDetails returns the value of the Details field in DatabricksServiceExceptionWithDetailsProto as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabricksServiceExceptionWithDetailsProto) GetDetails(ctx context.Context) ([]types.Object, bool) {
	if m.Details.IsNull() || m.Details.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Details.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetails sets the value of the Details field in DatabricksServiceExceptionWithDetailsProto.
func (m *DatabricksServiceExceptionWithDetailsProto) SetDetails(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Details = types.ListValueMust(t, vs)
}

type DeleteProjectRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProjectRequest) {
}

func (to *DeleteProjectRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteProjectRequest) {
}

func (m DeleteProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProjectRequest
// only implements ToObjectValue() and Type().
func (m DeleteProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// A Databricks job or pipeline created during conversion.
type DestinationResource struct {
	// The ID of the job that conversion created, if any.
	JobId types.Int64 `tfsdk:"job_id"`
	// The ID (a UUID) of the pipeline that conversion created, if any.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (to *DestinationResource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DestinationResource) {
}

func (to *DestinationResource) SyncFieldsDuringRead(ctx context.Context, from DestinationResource) {
}

func (m DestinationResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["job_id"] = attrs["job_id"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DestinationResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DestinationResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DestinationResource
// only implements ToObjectValue() and Type().
func (m DestinationResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":      m.JobId,
			"pipeline_id": m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DestinationResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":      types.Int64Type,
			"pipeline_id": types.StringType,
		},
	}
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name types.String `tfsdk:"-"`
}

func (to *GetOperationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOperationRequest) {
}

func (to *GetOperationRequest) SyncFieldsDuringRead(ctx context.Context, from GetOperationRequest) {
}

func (m GetOperationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOperationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOperationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOperationRequest
// only implements ToObjectValue() and Type().
func (m GetOperationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOperationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetProjectFileRequest struct {
	// The resource name of the file, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`.
	Name types.String `tfsdk:"-"`
}

func (to *GetProjectFileRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProjectFileRequest) {
}

func (to *GetProjectFileRequest) SyncFieldsDuringRead(ctx context.Context, from GetProjectFileRequest) {
}

func (m GetProjectFileRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProjectFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetProjectFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProjectFileRequest
// only implements ToObjectValue() and Type().
func (m GetProjectFileRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProjectFileRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetProjectRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name types.String `tfsdk:"-"`
}

func (to *GetProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProjectRequest) {
}

func (to *GetProjectRequest) SyncFieldsDuringRead(ctx context.Context, from GetProjectRequest) {
}

func (m GetProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProjectRequest
// only implements ToObjectValue() and Type().
func (m GetProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListProjectFilesRequest struct {
	// The maximum number of files to return, at most 500. If omitted, zero, or
	// negative, the server chooses. The server can return a different number of
	// files than requested, so keep requesting pages until `next_page_token` is
	// empty.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response. Omit it for the first
	// page. Use it only with the same `parent`.
	PageToken types.String `tfsdk:"-"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListProjectFilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectFilesRequest) {
}

func (to *ListProjectFilesRequest) SyncFieldsDuringRead(ctx context.Context, from ListProjectFilesRequest) {
}

func (m ListProjectFilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectFilesRequest
// only implements ToObjectValue() and Type().
func (m ListProjectFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectFilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// The response to `ListProjectFiles`.
type ListProjectFilesResponse struct {
	// The token for the next page. Empty when there are no more pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The files in this page.
	ProjectFiles types.List `tfsdk:"project_files"`
}

func (to *ListProjectFilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectFilesResponse) {
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() && to.ProjectFiles.IsNull() && len(from.ProjectFiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProjectFiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProjectFiles = from.ProjectFiles
	}
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() {
		if toProjectFiles, ok := to.GetProjectFiles(ctx); ok {
			if fromProjectFiles, ok := from.GetProjectFiles(ctx); ok {
				// Recursively sync the fields of each ProjectFiles element by position.
				for i := range toProjectFiles {
					if i < len(fromProjectFiles) {
						toProjectFiles[i].SyncFieldsDuringCreateOrUpdate(ctx, fromProjectFiles[i])
					}
				}
				to.SetProjectFiles(ctx, toProjectFiles)
			}
		}
	}
}

func (to *ListProjectFilesResponse) SyncFieldsDuringRead(ctx context.Context, from ListProjectFilesResponse) {
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() && to.ProjectFiles.IsNull() && len(from.ProjectFiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProjectFiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProjectFiles = from.ProjectFiles
	}
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() {
		if toProjectFiles, ok := to.GetProjectFiles(ctx); ok {
			if fromProjectFiles, ok := from.GetProjectFiles(ctx); ok {
				for i := range toProjectFiles {
					if i < len(fromProjectFiles) {
						toProjectFiles[i].SyncFieldsDuringRead(ctx, fromProjectFiles[i])
					}
				}
				to.SetProjectFiles(ctx, toProjectFiles)
			}
		}
	}
}

func (m ListProjectFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["project_files"] = attrs["project_files"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project_files": reflect.TypeOf(ProjectFile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectFilesResponse
// only implements ToObjectValue() and Type().
func (m ListProjectFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"project_files":   m.ProjectFiles,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectFilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"project_files": basetypes.ListType{
				ElemType: ProjectFile{}.Type(ctx),
			},
		},
	}
}

// GetProjectFiles returns the value of the ProjectFiles field in ListProjectFilesResponse as
// a slice of ProjectFile values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProjectFilesResponse) GetProjectFiles(ctx context.Context) ([]ProjectFile, bool) {
	if m.ProjectFiles.IsNull() || m.ProjectFiles.IsUnknown() {
		return nil, false
	}
	var v []ProjectFile
	d := m.ProjectFiles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjectFiles sets the value of the ProjectFiles field in ListProjectFilesResponse.
func (m *ListProjectFilesResponse) SetProjectFiles(ctx context.Context, v []ProjectFile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["project_files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ProjectFiles = types.ListValueMust(t, vs)
}

type ListProjectsRequest struct {
	// The folder to list, as `folder = "folders/{folder_id}"`. For example,
	// `folder = "folders/456"`. No other filter is supported. Defaults to the
	// caller's home folder.
	Filter types.String `tfsdk:"-"`
	// The maximum number of projects to return, at most 500. Defaults to 100 if
	// omitted, zero, or negative. A page can have fewer projects than
	// requested, or none, even when more pages follow.
	PageSize types.Int64 `tfsdk:"-"`
	// The `next_page_token` from a previous response. Omit it for the first
	// page. Use the same `filter` for every page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListProjectsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsRequest) {
}

func (to *ListProjectsRequest) SyncFieldsDuringRead(ctx context.Context, from ListProjectsRequest) {
}

func (m ListProjectsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsRequest
// only implements ToObjectValue() and Type().
func (m ListProjectsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// The response to `ListProjects`.
type ListProjectsResponse struct {
	// The token for the next page. Empty when there are no more pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The projects in this page, without `project_config` or `changes_locked`.
	// Call `GetProject` for the full project.
	Projects types.List `tfsdk:"projects"`
}

func (to *ListProjectsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsResponse) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() {
		if toProjects, ok := to.GetProjects(ctx); ok {
			if fromProjects, ok := from.GetProjects(ctx); ok {
				// Recursively sync the fields of each Projects element by position.
				for i := range toProjects {
					if i < len(fromProjects) {
						toProjects[i].SyncFieldsDuringCreateOrUpdate(ctx, fromProjects[i])
					}
				}
				to.SetProjects(ctx, toProjects)
			}
		}
	}
}

func (to *ListProjectsResponse) SyncFieldsDuringRead(ctx context.Context, from ListProjectsResponse) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() {
		if toProjects, ok := to.GetProjects(ctx); ok {
			if fromProjects, ok := from.GetProjects(ctx); ok {
				for i := range toProjects {
					if i < len(fromProjects) {
						toProjects[i].SyncFieldsDuringRead(ctx, fromProjects[i])
					}
				}
				to.SetProjects(ctx, toProjects)
			}
		}
	}
}

func (m ListProjectsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["projects"] = attrs["projects"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"projects": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsResponse
// only implements ToObjectValue() and Type().
func (m ListProjectsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"projects":        m.Projects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"projects": basetypes.ListType{
				ElemType: Project{}.Type(ctx),
			},
		},
	}
}

// GetProjects returns the value of the Projects field in ListProjectsResponse as
// a slice of Project values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProjectsResponse) GetProjects(ctx context.Context) ([]Project, bool) {
	if m.Projects.IsNull() || m.Projects.IsUnknown() {
		return nil, false
	}
	var v []Project
	d := m.Projects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjects sets the value of the Projects field in ListProjectsResponse.
func (m *ListProjectsResponse) SetProjects(ctx context.Context, v []Project) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["projects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Projects = types.ListValueMust(t, vs)
}

// The location of a file or folder. Exactly one location type is set. More
// location types can be added later, so clients should ignore types they don't
// recognize.
type Location struct {
	// A file or folder in the workspace.
	Workspace types.Object `tfsdk:"workspace"`
}

func (to *Location) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Location) {
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				// Recursively sync the fields of Workspace
				toWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (to *Location) SyncFieldsDuringRead(ctx context.Context, from Location) {
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				toWorkspace.SyncFieldsDuringRead(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (m Location) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace"] = attrs["workspace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Location.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Location) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace": reflect.TypeOf(WorkspaceLocation{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Location
// only implements ToObjectValue() and Type().
func (m Location) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace": m.Workspace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Location) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace": WorkspaceLocation{}.Type(ctx),
		},
	}
}

// GetWorkspace returns the value of the Workspace field in Location as
// a WorkspaceLocation value.
// If the field is unknown or null, the boolean return value is false.
func (m *Location) GetWorkspace(ctx context.Context) (WorkspaceLocation, bool) {
	var e WorkspaceLocation
	if m.Workspace.IsNull() || m.Workspace.IsUnknown() {
		return e, false
	}
	var v WorkspaceLocation
	d := m.Workspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspace sets the value of the Workspace field in Location.
func (m *Location) SetWorkspace(ctx context.Context, v WorkspaceLocation) {
	vs := v.ToObjectValue(ctx)
	m.Workspace = vs
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done types.Bool `tfsdk:"done"`
	// The error result of the operation in case of failure or cancellation.
	Error types.Object `tfsdk:"error"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata types.Object `tfsdk:"metadata"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name types.String `tfsdk:"name"`
	// The normal, successful response of the operation.
	Response types.Object `tfsdk:"response"`
}

func (to *Operation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				// Recursively sync the fields of Error
				toError.SyncFieldsDuringCreateOrUpdate(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (to *Operation) SyncFieldsDuringRead(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m Operation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["done"] = attrs["done"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["response"] = attrs["response"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Operation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Operation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(DatabricksServiceExceptionWithDetailsProto{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation
// only implements ToObjectValue() and Type().
func (m Operation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"done":     m.Done,
			"error":    m.Error,
			"metadata": m.Metadata,
			"name":     m.Name,
			"response": m.Response,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Operation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"done":     types.BoolType,
			"error":    DatabricksServiceExceptionWithDetailsProto{}.Type(ctx),
			"metadata": types.ObjectType{},
			"name":     types.StringType,
			"response": types.ObjectType{},
		},
	}
}

// GetError returns the value of the Error field in Operation as
// a DatabricksServiceExceptionWithDetailsProto value.
// If the field is unknown or null, the boolean return value is false.
func (m *Operation) GetError(ctx context.Context) (DatabricksServiceExceptionWithDetailsProto, bool) {
	var e DatabricksServiceExceptionWithDetailsProto
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v DatabricksServiceExceptionWithDetailsProto
	d := m.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in Operation.
func (m *Operation) SetError(ctx context.Context, v DatabricksServiceExceptionWithDetailsProto) {
	vs := v.ToObjectValue(ctx)
	m.Error = vs
}

// A migration project: the conversion settings and source files of one
// migration.
type Project struct {
	// Whether the source dialect and target framework are locked. They lock
	// after a conversion or analysis runs on the project. Other settings stay
	// editable.
	ChangesLocked types.Bool `tfsdk:"changes_locked"`
	// The time the project was created, if available. Can be absent from
	// `CreateProject` and `UpdateProject` responses.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The display name of the project. Required on create, and can't be blank.
	// Changing it renames the project.
	DisplayName types.String `tfsdk:"display_name"`
	// The workspace folder that contains the project, in the format
	// `folders/{folder_id}`. For example, `folders/456`. If omitted on create,
	// defaults to the caller's home folder. `UpdateProject` can't change it,
	// but moving the project in the workspace does. Returned when available.
	// Converted files are written to `project_config.destination_location`,
	// which can be a different folder.
	Folder types.String `tfsdk:"folder"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. The server assigns it when the project is
	// created.
	Name types.String `tfsdk:"name"`
	// The conversion settings of the project. Required on create. Returned by
	// `CreateProject`, `GetProject`, and `UpdateProject`, but not by
	// `ListProjects`.
	ProjectConfig types.Object `tfsdk:"project_config"`
	// The time the project was last modified, if available. Can be absent from
	// `CreateProject` and `UpdateProject` responses.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Project) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Project) {
	if !from.ProjectConfig.IsNull() && !from.ProjectConfig.IsUnknown() {
		if toProjectConfig, ok := to.GetProjectConfig(ctx); ok {
			if fromProjectConfig, ok := from.GetProjectConfig(ctx); ok {
				// Recursively sync the fields of ProjectConfig
				toProjectConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromProjectConfig)
				to.SetProjectConfig(ctx, toProjectConfig)
			}
		}
	}
}

func (to *Project) SyncFieldsDuringRead(ctx context.Context, from Project) {
	if !from.ProjectConfig.IsNull() && !from.ProjectConfig.IsUnknown() {
		if toProjectConfig, ok := to.GetProjectConfig(ctx); ok {
			if fromProjectConfig, ok := from.GetProjectConfig(ctx); ok {
				toProjectConfig.SyncFieldsDuringRead(ctx, fromProjectConfig)
				to.SetProjectConfig(ctx, toProjectConfig)
			}
		}
	}
}

func (m Project) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["changes_locked"] = attrs["changes_locked"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["folder"] = attrs["folder"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["project_config"] = attrs["project_config"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Project.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Project) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project_config": reflect.TypeOf(ProjectConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Project
// only implements ToObjectValue() and Type().
func (m Project) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes_locked": m.ChangesLocked,
			"create_time":    m.CreateTime,
			"display_name":   m.DisplayName,
			"folder":         m.Folder,
			"name":           m.Name,
			"project_config": m.ProjectConfig,
			"update_time":    m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Project) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes_locked": types.BoolType,
			"create_time":    timetypes.RFC3339{}.Type(ctx),
			"display_name":   types.StringType,
			"folder":         types.StringType,
			"name":           types.StringType,
			"project_config": ProjectConfig{}.Type(ctx),
			"update_time":    timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetProjectConfig returns the value of the ProjectConfig field in Project as
// a ProjectConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetProjectConfig(ctx context.Context) (ProjectConfig, bool) {
	var e ProjectConfig
	if m.ProjectConfig.IsNull() || m.ProjectConfig.IsUnknown() {
		return e, false
	}
	var v ProjectConfig
	d := m.ProjectConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjectConfig sets the value of the ProjectConfig field in Project.
func (m *Project) SetProjectConfig(ctx context.Context, v ProjectConfig) {
	vs := v.ToObjectValue(ctx)
	m.ProjectConfig = vs
}

// Conversion settings of a migration project.
type ProjectConfig struct {
	// The workspace folder where converted files are written. Required on
	// create; set `workspace.path` or `workspace.object_id`. `UpdateProject`
	// can't change it.
	DestinationLocation types.Object `tfsdk:"destination_location"`
	// The workspace JSON file that holds the project's lineage report.
	// Optional. To clear it, include `project_config.lineage_location` in
	// `update_mask` and leave this field unset.
	LineageLocation types.Object `tfsdk:"lineage_location"`
	// The dialect of the project's source files. Can't be changed when
	// `Project.changes_locked` is `true`.
	SourceDialect types.String `tfsdk:"source_dialect"`
	// The framework that source files are converted to. Required on create, and
	// can't be `TARGET_FRAMEWORK_UNSPECIFIED`. Can't be changed when
	// `Project.changes_locked` is `true`.
	TargetFramework types.String `tfsdk:"target_framework"`
}

func (to *ProjectConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectConfig) {
	if !from.DestinationLocation.IsNull() && !from.DestinationLocation.IsUnknown() {
		if toDestinationLocation, ok := to.GetDestinationLocation(ctx); ok {
			if fromDestinationLocation, ok := from.GetDestinationLocation(ctx); ok {
				// Recursively sync the fields of DestinationLocation
				toDestinationLocation.SyncFieldsDuringCreateOrUpdate(ctx, fromDestinationLocation)
				to.SetDestinationLocation(ctx, toDestinationLocation)
			}
		}
	}
	if !from.LineageLocation.IsNull() && !from.LineageLocation.IsUnknown() {
		if toLineageLocation, ok := to.GetLineageLocation(ctx); ok {
			if fromLineageLocation, ok := from.GetLineageLocation(ctx); ok {
				// Recursively sync the fields of LineageLocation
				toLineageLocation.SyncFieldsDuringCreateOrUpdate(ctx, fromLineageLocation)
				to.SetLineageLocation(ctx, toLineageLocation)
			}
		}
	}
}

func (to *ProjectConfig) SyncFieldsDuringRead(ctx context.Context, from ProjectConfig) {
	if !from.DestinationLocation.IsNull() && !from.DestinationLocation.IsUnknown() {
		if toDestinationLocation, ok := to.GetDestinationLocation(ctx); ok {
			if fromDestinationLocation, ok := from.GetDestinationLocation(ctx); ok {
				toDestinationLocation.SyncFieldsDuringRead(ctx, fromDestinationLocation)
				to.SetDestinationLocation(ctx, toDestinationLocation)
			}
		}
	}
	if !from.LineageLocation.IsNull() && !from.LineageLocation.IsUnknown() {
		if toLineageLocation, ok := to.GetLineageLocation(ctx); ok {
			if fromLineageLocation, ok := from.GetLineageLocation(ctx); ok {
				toLineageLocation.SyncFieldsDuringRead(ctx, fromLineageLocation)
				to.SetLineageLocation(ctx, toLineageLocation)
			}
		}
	}
}

func (m ProjectConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_location"] = attrs["destination_location"].SetOptional()
	attrs["lineage_location"] = attrs["lineage_location"].SetOptional()
	attrs["source_dialect"] = attrs["source_dialect"].SetOptional()
	attrs["target_framework"] = attrs["target_framework"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"destination_location": reflect.TypeOf(Location{}),
		"lineage_location":     reflect.TypeOf(Location{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectConfig
// only implements ToObjectValue() and Type().
func (m ProjectConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_location": m.DestinationLocation,
			"lineage_location":     m.LineageLocation,
			"source_dialect":       m.SourceDialect,
			"target_framework":     m.TargetFramework,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_location": Location{}.Type(ctx),
			"lineage_location":     Location{}.Type(ctx),
			"source_dialect":       types.StringType,
			"target_framework":     types.StringType,
		},
	}
}

// GetDestinationLocation returns the value of the DestinationLocation field in ProjectConfig as
// a Location value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectConfig) GetDestinationLocation(ctx context.Context) (Location, bool) {
	var e Location
	if m.DestinationLocation.IsNull() || m.DestinationLocation.IsUnknown() {
		return e, false
	}
	var v Location
	d := m.DestinationLocation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestinationLocation sets the value of the DestinationLocation field in ProjectConfig.
func (m *ProjectConfig) SetDestinationLocation(ctx context.Context, v Location) {
	vs := v.ToObjectValue(ctx)
	m.DestinationLocation = vs
}

// GetLineageLocation returns the value of the LineageLocation field in ProjectConfig as
// a Location value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectConfig) GetLineageLocation(ctx context.Context) (Location, bool) {
	var e Location
	if m.LineageLocation.IsNull() || m.LineageLocation.IsUnknown() {
		return e, false
	}
	var v Location
	d := m.LineageLocation.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLineageLocation sets the value of the LineageLocation field in ProjectConfig.
func (m *ProjectConfig) SetLineageLocation(ctx context.Context, v Location) {
	vs := v.ToObjectValue(ctx)
	m.LineageLocation = vs
}

// A source file in a migration project, with its latest conversion results. The
// service sets the conversion results. To accept the converted output, call
// `ReviewProjectFiles`; to undo that, call `UnreviewProjectFiles`. Neither
// method changes the conversion results.
type ProjectFile struct {
	// Issues that block the conversion. Returned only by `GetProjectFile`,
	// which reads them from the converted file.
	Blockers types.List `tfsdk:"blockers"`
	// The changes made during conversion. Returned only by `GetProjectFile`,
	// which reads them from the converted file.
	Changes types.List `tfsdk:"changes"`
	// The location of the asset that this file was derived from. Absent if the
	// file wasn't derived from another asset.
	DerivedFrom types.Object `tfsdk:"derived_from"`
	// The location of the converted file. Absent until a conversion writes it.
	Destination types.Object `tfsdk:"destination"`
	// The kind of the converted file, when recorded.
	DestinationKind types.String `tfsdk:"destination_kind"`
	// The Databricks job or pipeline that conversion created for this file, if
	// any.
	DestinationResource types.Object `tfsdk:"destination_resource"`
	// The error reported for this file, if any. `GetProjectFile` returns the
	// error stored with the converted file when available, and otherwise the
	// latest error recorded for the project file.
	ErrorMessage types.String `tfsdk:"error_message"`
	// Whether the file has conversion notes. Returned by `ListProjectFiles`;
	// call `GetProjectFile` to read the notes.
	HasNotes types.Bool `tfsdk:"has_notes"`
	// The resource name of the file, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`. The server assigns `file_id`, an opaque value
	// of at most 63 characters.
	Name types.String `tfsdk:"name"`
	// Informational notes about the conversion that need no decision. Returned
	// only by `GetProjectFile`, which reads them from the converted file.
	Notes types.List `tfsdk:"notes"`
	// The location of the source file.
	Source types.Object `tfsdk:"source"`
	// The file's latest conversion status across all runs. For its status in a
	// specific run, read `ConversionRun.files` in that run's operation
	// metadata.
	Status types.String `tfsdk:"status"`
	// Follow-up work (to-do items) identified during conversion. Returned only
	// by `GetProjectFile`, which reads them from the converted file.
	Todos types.List `tfsdk:"todos"`
}

func (to *ProjectFile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectFile) {
	if !from.Blockers.IsNull() && !from.Blockers.IsUnknown() && to.Blockers.IsNull() && len(from.Blockers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Blockers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Blockers = from.Blockers
	}
	if !from.Changes.IsNull() && !from.Changes.IsUnknown() && to.Changes.IsNull() && len(from.Changes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Changes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Changes = from.Changes
	}
	if !from.DerivedFrom.IsNull() && !from.DerivedFrom.IsUnknown() {
		if toDerivedFrom, ok := to.GetDerivedFrom(ctx); ok {
			if fromDerivedFrom, ok := from.GetDerivedFrom(ctx); ok {
				// Recursively sync the fields of DerivedFrom
				toDerivedFrom.SyncFieldsDuringCreateOrUpdate(ctx, fromDerivedFrom)
				to.SetDerivedFrom(ctx, toDerivedFrom)
			}
		}
	}
	if !from.Destination.IsNull() && !from.Destination.IsUnknown() {
		if toDestination, ok := to.GetDestination(ctx); ok {
			if fromDestination, ok := from.GetDestination(ctx); ok {
				// Recursively sync the fields of Destination
				toDestination.SyncFieldsDuringCreateOrUpdate(ctx, fromDestination)
				to.SetDestination(ctx, toDestination)
			}
		}
	}
	if !from.DestinationResource.IsNull() && !from.DestinationResource.IsUnknown() {
		if toDestinationResource, ok := to.GetDestinationResource(ctx); ok {
			if fromDestinationResource, ok := from.GetDestinationResource(ctx); ok {
				// Recursively sync the fields of DestinationResource
				toDestinationResource.SyncFieldsDuringCreateOrUpdate(ctx, fromDestinationResource)
				to.SetDestinationResource(ctx, toDestinationResource)
			}
		}
	}
	if !from.Notes.IsNull() && !from.Notes.IsUnknown() && to.Notes.IsNull() && len(from.Notes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notes = from.Notes
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				// Recursively sync the fields of Source
				toSource.SyncFieldsDuringCreateOrUpdate(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.Todos.IsNull() && !from.Todos.IsUnknown() && to.Todos.IsNull() && len(from.Todos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Todos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Todos = from.Todos
	}
}

func (to *ProjectFile) SyncFieldsDuringRead(ctx context.Context, from ProjectFile) {
	if !from.Blockers.IsNull() && !from.Blockers.IsUnknown() && to.Blockers.IsNull() && len(from.Blockers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Blockers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Blockers = from.Blockers
	}
	if !from.Changes.IsNull() && !from.Changes.IsUnknown() && to.Changes.IsNull() && len(from.Changes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Changes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Changes = from.Changes
	}
	if !from.DerivedFrom.IsNull() && !from.DerivedFrom.IsUnknown() {
		if toDerivedFrom, ok := to.GetDerivedFrom(ctx); ok {
			if fromDerivedFrom, ok := from.GetDerivedFrom(ctx); ok {
				toDerivedFrom.SyncFieldsDuringRead(ctx, fromDerivedFrom)
				to.SetDerivedFrom(ctx, toDerivedFrom)
			}
		}
	}
	if !from.Destination.IsNull() && !from.Destination.IsUnknown() {
		if toDestination, ok := to.GetDestination(ctx); ok {
			if fromDestination, ok := from.GetDestination(ctx); ok {
				toDestination.SyncFieldsDuringRead(ctx, fromDestination)
				to.SetDestination(ctx, toDestination)
			}
		}
	}
	if !from.DestinationResource.IsNull() && !from.DestinationResource.IsUnknown() {
		if toDestinationResource, ok := to.GetDestinationResource(ctx); ok {
			if fromDestinationResource, ok := from.GetDestinationResource(ctx); ok {
				toDestinationResource.SyncFieldsDuringRead(ctx, fromDestinationResource)
				to.SetDestinationResource(ctx, toDestinationResource)
			}
		}
	}
	if !from.Notes.IsNull() && !from.Notes.IsUnknown() && to.Notes.IsNull() && len(from.Notes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notes = from.Notes
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				toSource.SyncFieldsDuringRead(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.Todos.IsNull() && !from.Todos.IsUnknown() && to.Todos.IsNull() && len(from.Todos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Todos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Todos = from.Todos
	}
}

func (m ProjectFile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["blockers"] = attrs["blockers"].SetOptional()
	attrs["changes"] = attrs["changes"].SetOptional()
	attrs["derived_from"] = attrs["derived_from"].SetComputed()
	attrs["destination"] = attrs["destination"].SetComputed()
	attrs["destination_kind"] = attrs["destination_kind"].SetComputed()
	attrs["destination_resource"] = attrs["destination_resource"].SetOptional()
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["has_notes"] = attrs["has_notes"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["notes"] = attrs["notes"].SetOptional()
	attrs["source"] = attrs["source"].SetComputed()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["todos"] = attrs["todos"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectFile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectFile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"blockers":             reflect.TypeOf(types.String{}),
		"changes":              reflect.TypeOf(types.String{}),
		"derived_from":         reflect.TypeOf(Location{}),
		"destination":          reflect.TypeOf(Location{}),
		"destination_resource": reflect.TypeOf(DestinationResource{}),
		"notes":                reflect.TypeOf(types.String{}),
		"source":               reflect.TypeOf(Location{}),
		"todos":                reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectFile
// only implements ToObjectValue() and Type().
func (m ProjectFile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"blockers":             m.Blockers,
			"changes":              m.Changes,
			"derived_from":         m.DerivedFrom,
			"destination":          m.Destination,
			"destination_kind":     m.DestinationKind,
			"destination_resource": m.DestinationResource,
			"error_message":        m.ErrorMessage,
			"has_notes":            m.HasNotes,
			"name":                 m.Name,
			"notes":                m.Notes,
			"source":               m.Source,
			"status":               m.Status,
			"todos":                m.Todos,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectFile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"blockers": basetypes.ListType{
				ElemType: types.StringType,
			},
			"changes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"derived_from":         Location{}.Type(ctx),
			"destination":          Location{}.Type(ctx),
			"destination_kind":     types.StringType,
			"destination_resource": DestinationResource{}.Type(ctx),
			"error_message":        types.StringType,
			"has_notes":            types.BoolType,
			"name":                 types.StringType,
			"notes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"source": Location{}.Type(ctx),
			"status": types.StringType,
			"todos": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetBlockers returns the value of the Blockers field in ProjectFile as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetBlockers(ctx context.Context) ([]types.String, bool) {
	if m.Blockers.IsNull() || m.Blockers.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Blockers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBlockers sets the value of the Blockers field in ProjectFile.
func (m *ProjectFile) SetBlockers(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["blockers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Blockers = types.ListValueMust(t, vs)
}

// GetChanges returns the value of the Changes field in ProjectFile as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetChanges(ctx context.Context) ([]types.String, bool) {
	if m.Changes.IsNull() || m.Changes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in ProjectFile.
func (m *ProjectFile) SetChanges(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Changes = types.ListValueMust(t, vs)
}

// GetDerivedFrom returns the value of the DerivedFrom field in ProjectFile as
// a Location value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetDerivedFrom(ctx context.Context) (Location, bool) {
	var e Location
	if m.DerivedFrom.IsNull() || m.DerivedFrom.IsUnknown() {
		return e, false
	}
	var v Location
	d := m.DerivedFrom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDerivedFrom sets the value of the DerivedFrom field in ProjectFile.
func (m *ProjectFile) SetDerivedFrom(ctx context.Context, v Location) {
	vs := v.ToObjectValue(ctx)
	m.DerivedFrom = vs
}

// GetDestination returns the value of the Destination field in ProjectFile as
// a Location value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetDestination(ctx context.Context) (Location, bool) {
	var e Location
	if m.Destination.IsNull() || m.Destination.IsUnknown() {
		return e, false
	}
	var v Location
	d := m.Destination.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestination sets the value of the Destination field in ProjectFile.
func (m *ProjectFile) SetDestination(ctx context.Context, v Location) {
	vs := v.ToObjectValue(ctx)
	m.Destination = vs
}

// GetDestinationResource returns the value of the DestinationResource field in ProjectFile as
// a DestinationResource value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetDestinationResource(ctx context.Context) (DestinationResource, bool) {
	var e DestinationResource
	if m.DestinationResource.IsNull() || m.DestinationResource.IsUnknown() {
		return e, false
	}
	var v DestinationResource
	d := m.DestinationResource.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestinationResource sets the value of the DestinationResource field in ProjectFile.
func (m *ProjectFile) SetDestinationResource(ctx context.Context, v DestinationResource) {
	vs := v.ToObjectValue(ctx)
	m.DestinationResource = vs
}

// GetNotes returns the value of the Notes field in ProjectFile as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetNotes(ctx context.Context) ([]types.String, bool) {
	if m.Notes.IsNull() || m.Notes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Notes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotes sets the value of the Notes field in ProjectFile.
func (m *ProjectFile) SetNotes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notes = types.ListValueMust(t, vs)
}

// GetSource returns the value of the Source field in ProjectFile as
// a Location value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetSource(ctx context.Context) (Location, bool) {
	var e Location
	if m.Source.IsNull() || m.Source.IsUnknown() {
		return e, false
	}
	var v Location
	d := m.Source.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSource sets the value of the Source field in ProjectFile.
func (m *ProjectFile) SetSource(ctx context.Context, v Location) {
	vs := v.ToObjectValue(ctx)
	m.Source = vs
}

// GetTodos returns the value of the Todos field in ProjectFile as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectFile) GetTodos(ctx context.Context) ([]types.String, bool) {
	if m.Todos.IsNull() || m.Todos.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Todos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTodos sets the value of the Todos field in ProjectFile.
func (m *ProjectFile) SetTodos(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["todos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Todos = types.ListValueMust(t, vs)
}

// The request to remove files from a migration project.
type RemoveProjectFilesRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`.
	Name types.String `tfsdk:"-"`
	// The files to remove, by `ProjectFile.name`. For example,
	// `projects/123/files/789`. Every file must belong to the project in
	// `name`. An empty list changes nothing. Send at most 10 files per request,
	// the default limit. Requests over a project's limit return
	// `INVALID_PARAMETER_VALUE`.
	Names types.List `tfsdk:"names"`
}

func (to *RemoveProjectFilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveProjectFilesRequest) {
	if !from.Names.IsNull() && !from.Names.IsUnknown() && to.Names.IsNull() && len(from.Names.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Names, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Names = from.Names
	}
}

func (to *RemoveProjectFilesRequest) SyncFieldsDuringRead(ctx context.Context, from RemoveProjectFilesRequest) {
	if !from.Names.IsNull() && !from.Names.IsUnknown() && to.Names.IsNull() && len(from.Names.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Names, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Names = from.Names
	}
}

func (m RemoveProjectFilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["names"] = attrs["names"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveProjectFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RemoveProjectFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveProjectFilesRequest
// only implements ToObjectValue() and Type().
func (m RemoveProjectFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  m.Name,
			"names": m.Names,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveProjectFilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetNames returns the value of the Names field in RemoveProjectFilesRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RemoveProjectFilesRequest) GetNames(ctx context.Context) ([]types.String, bool) {
	if m.Names.IsNull() || m.Names.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Names.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNames sets the value of the Names field in RemoveProjectFilesRequest.
func (m *RemoveProjectFilesRequest) SetNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Names = types.ListValueMust(t, vs)
}

// Lists the files from a `RemoveProjectFiles` request.
type RemoveProjectFilesResponse struct {
	// The requested files, in request order. Each entry identifies a source
	// file and has no conversion results.
	Files types.List `tfsdk:"files"`
}

func (to *RemoveProjectFilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveProjectFilesResponse) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				// Recursively sync the fields of each Files element by position.
				for i := range toFiles {
					if i < len(fromFiles) {
						toFiles[i].SyncFieldsDuringCreateOrUpdate(ctx, fromFiles[i])
					}
				}
				to.SetFiles(ctx, toFiles)
			}
		}
	}
}

func (to *RemoveProjectFilesResponse) SyncFieldsDuringRead(ctx context.Context, from RemoveProjectFilesResponse) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				for i := range toFiles {
					if i < len(fromFiles) {
						toFiles[i].SyncFieldsDuringRead(ctx, fromFiles[i])
					}
				}
				to.SetFiles(ctx, toFiles)
			}
		}
	}
}

func (m RemoveProjectFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["files"] = attrs["files"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveProjectFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RemoveProjectFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(ProjectFile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveProjectFilesResponse
// only implements ToObjectValue() and Type().
func (m RemoveProjectFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files": m.Files,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveProjectFilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"files": basetypes.ListType{
				ElemType: ProjectFile{}.Type(ctx),
			},
		},
	}
}

// GetFiles returns the value of the Files field in RemoveProjectFilesResponse as
// a slice of ProjectFile values.
// If the field is unknown or null, the boolean return value is false.
func (m *RemoveProjectFilesResponse) GetFiles(ctx context.Context) ([]ProjectFile, bool) {
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return nil, false
	}
	var v []ProjectFile
	d := m.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in RemoveProjectFilesResponse.
func (m *RemoveProjectFilesResponse) SetFiles(ctx context.Context, v []ProjectFile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Files = types.ListValueMust(t, vs)
}

// The request to mark files of a migration project as reviewed.
type ReviewProjectFilesRequest struct {
	// The files to review: 1 to 100 distinct file resource names, in the format
	// `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`.
	Names types.List `tfsdk:"names"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. Every file in `names` must belong to this
	// project.
	Parent types.String `tfsdk:"-"`
}

func (to *ReviewProjectFilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReviewProjectFilesRequest) {
}

func (to *ReviewProjectFilesRequest) SyncFieldsDuringRead(ctx context.Context, from ReviewProjectFilesRequest) {
}

func (m ReviewProjectFilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["names"] = attrs["names"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReviewProjectFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ReviewProjectFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReviewProjectFilesRequest
// only implements ToObjectValue() and Type().
func (m ReviewProjectFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"names":  m.Names,
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ReviewProjectFilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"parent": types.StringType,
		},
	}
}

// GetNames returns the value of the Names field in ReviewProjectFilesRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ReviewProjectFilesRequest) GetNames(ctx context.Context) ([]types.String, bool) {
	if m.Names.IsNull() || m.Names.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Names.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNames sets the value of the Names field in ReviewProjectFilesRequest.
func (m *ReviewProjectFilesRequest) SetNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Names = types.ListValueMust(t, vs)
}

// The files reviewed by `ReviewProjectFiles`, in request order.
type ReviewProjectFilesResponse struct {
	// The reviewed files, in request order. To read a file's changes, to-do
	// items, blockers, and notes, call `GetProjectFile`.
	ProjectFiles types.List `tfsdk:"project_files"`
}

func (to *ReviewProjectFilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReviewProjectFilesResponse) {
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() && to.ProjectFiles.IsNull() && len(from.ProjectFiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProjectFiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProjectFiles = from.ProjectFiles
	}
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() {
		if toProjectFiles, ok := to.GetProjectFiles(ctx); ok {
			if fromProjectFiles, ok := from.GetProjectFiles(ctx); ok {
				// Recursively sync the fields of each ProjectFiles element by position.
				for i := range toProjectFiles {
					if i < len(fromProjectFiles) {
						toProjectFiles[i].SyncFieldsDuringCreateOrUpdate(ctx, fromProjectFiles[i])
					}
				}
				to.SetProjectFiles(ctx, toProjectFiles)
			}
		}
	}
}

func (to *ReviewProjectFilesResponse) SyncFieldsDuringRead(ctx context.Context, from ReviewProjectFilesResponse) {
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() && to.ProjectFiles.IsNull() && len(from.ProjectFiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProjectFiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProjectFiles = from.ProjectFiles
	}
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() {
		if toProjectFiles, ok := to.GetProjectFiles(ctx); ok {
			if fromProjectFiles, ok := from.GetProjectFiles(ctx); ok {
				for i := range toProjectFiles {
					if i < len(fromProjectFiles) {
						toProjectFiles[i].SyncFieldsDuringRead(ctx, fromProjectFiles[i])
					}
				}
				to.SetProjectFiles(ctx, toProjectFiles)
			}
		}
	}
}

func (m ReviewProjectFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_files"] = attrs["project_files"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReviewProjectFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ReviewProjectFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project_files": reflect.TypeOf(ProjectFile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReviewProjectFilesResponse
// only implements ToObjectValue() and Type().
func (m ReviewProjectFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_files": m.ProjectFiles,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ReviewProjectFilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_files": basetypes.ListType{
				ElemType: ProjectFile{}.Type(ctx),
			},
		},
	}
}

// GetProjectFiles returns the value of the ProjectFiles field in ReviewProjectFilesResponse as
// a slice of ProjectFile values.
// If the field is unknown or null, the boolean return value is false.
func (m *ReviewProjectFilesResponse) GetProjectFiles(ctx context.Context) ([]ProjectFile, bool) {
	if m.ProjectFiles.IsNull() || m.ProjectFiles.IsUnknown() {
		return nil, false
	}
	var v []ProjectFile
	d := m.ProjectFiles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjectFiles sets the value of the ProjectFiles field in ReviewProjectFilesResponse.
func (m *ReviewProjectFilesResponse) SetProjectFiles(ctx context.Context, v []ProjectFile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["project_files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ProjectFiles = types.ListValueMust(t, vs)
}

// A reference to a source file in the workspace. Set `workspace.path` or
// `workspace.object_id`. For example, `{"workspace": {"path":
// "/Workspace/Shared/migration/source.sql"}}`.
type SourceFileRef struct {
	// The location of the asset that this file was derived from, if any. Used
	// only by `AddProjectFiles`, and can't be changed after the file is added.
	DerivedFrom types.Object `tfsdk:"derived_from"`
	// The source file in the workspace. Must identify an existing file the
	// caller can read.
	Workspace types.Object `tfsdk:"workspace"`
}

func (to *SourceFileRef) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SourceFileRef) {
	if !from.DerivedFrom.IsNull() && !from.DerivedFrom.IsUnknown() {
		if toDerivedFrom, ok := to.GetDerivedFrom(ctx); ok {
			if fromDerivedFrom, ok := from.GetDerivedFrom(ctx); ok {
				// Recursively sync the fields of DerivedFrom
				toDerivedFrom.SyncFieldsDuringCreateOrUpdate(ctx, fromDerivedFrom)
				to.SetDerivedFrom(ctx, toDerivedFrom)
			}
		}
	}
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				// Recursively sync the fields of Workspace
				toWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (to *SourceFileRef) SyncFieldsDuringRead(ctx context.Context, from SourceFileRef) {
	if !from.DerivedFrom.IsNull() && !from.DerivedFrom.IsUnknown() {
		if toDerivedFrom, ok := to.GetDerivedFrom(ctx); ok {
			if fromDerivedFrom, ok := from.GetDerivedFrom(ctx); ok {
				toDerivedFrom.SyncFieldsDuringRead(ctx, fromDerivedFrom)
				to.SetDerivedFrom(ctx, toDerivedFrom)
			}
		}
	}
	if !from.Workspace.IsNull() && !from.Workspace.IsUnknown() {
		if toWorkspace, ok := to.GetWorkspace(ctx); ok {
			if fromWorkspace, ok := from.GetWorkspace(ctx); ok {
				toWorkspace.SyncFieldsDuringRead(ctx, fromWorkspace)
				to.SetWorkspace(ctx, toWorkspace)
			}
		}
	}
}

func (m SourceFileRef) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["derived_from"] = attrs["derived_from"].SetOptional()
	attrs["workspace"] = attrs["workspace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SourceFileRef.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SourceFileRef) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"derived_from": reflect.TypeOf(Location{}),
		"workspace":    reflect.TypeOf(WorkspaceLocation{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SourceFileRef
// only implements ToObjectValue() and Type().
func (m SourceFileRef) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"derived_from": m.DerivedFrom,
			"workspace":    m.Workspace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SourceFileRef) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"derived_from": Location{}.Type(ctx),
			"workspace":    WorkspaceLocation{}.Type(ctx),
		},
	}
}

// GetDerivedFrom returns the value of the DerivedFrom field in SourceFileRef as
// a Location value.
// If the field is unknown or null, the boolean return value is false.
func (m *SourceFileRef) GetDerivedFrom(ctx context.Context) (Location, bool) {
	var e Location
	if m.DerivedFrom.IsNull() || m.DerivedFrom.IsUnknown() {
		return e, false
	}
	var v Location
	d := m.DerivedFrom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDerivedFrom sets the value of the DerivedFrom field in SourceFileRef.
func (m *SourceFileRef) SetDerivedFrom(ctx context.Context, v Location) {
	vs := v.ToObjectValue(ctx)
	m.DerivedFrom = vs
}

// GetWorkspace returns the value of the Workspace field in SourceFileRef as
// a WorkspaceLocation value.
// If the field is unknown or null, the boolean return value is false.
func (m *SourceFileRef) GetWorkspace(ctx context.Context) (WorkspaceLocation, bool) {
	var e WorkspaceLocation
	if m.Workspace.IsNull() || m.Workspace.IsUnknown() {
		return e, false
	}
	var v WorkspaceLocation
	d := m.Workspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspace sets the value of the Workspace field in SourceFileRef.
func (m *SourceFileRef) SetWorkspace(ctx context.Context, v WorkspaceLocation) {
	vs := v.ToObjectValue(ctx)
	m.Workspace = vs
}

// The request to remove the reviewed status from files of a migration project.
type UnreviewProjectFilesRequest struct {
	// The files to unreview: 1 to 100 distinct file resource names, in the
	// format `projects/{project_id}/files/{file_id}`. For example,
	// `projects/123/files/789`.
	Names types.List `tfsdk:"names"`
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. Every file in `names` must belong to this
	// project.
	Parent types.String `tfsdk:"-"`
}

func (to *UnreviewProjectFilesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UnreviewProjectFilesRequest) {
}

func (to *UnreviewProjectFilesRequest) SyncFieldsDuringRead(ctx context.Context, from UnreviewProjectFilesRequest) {
}

func (m UnreviewProjectFilesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["names"] = attrs["names"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnreviewProjectFilesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UnreviewProjectFilesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnreviewProjectFilesRequest
// only implements ToObjectValue() and Type().
func (m UnreviewProjectFilesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"names":  m.Names,
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UnreviewProjectFilesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"parent": types.StringType,
		},
	}
}

// GetNames returns the value of the Names field in UnreviewProjectFilesRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UnreviewProjectFilesRequest) GetNames(ctx context.Context) ([]types.String, bool) {
	if m.Names.IsNull() || m.Names.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Names.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNames sets the value of the Names field in UnreviewProjectFilesRequest.
func (m *UnreviewProjectFilesRequest) SetNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Names = types.ListValueMust(t, vs)
}

// The files unreviewed by `UnreviewProjectFiles`, in request order.
type UnreviewProjectFilesResponse struct {
	// The unreviewed files, in request order. To read a file's changes, to-do
	// items, blockers, and notes, call `GetProjectFile`.
	ProjectFiles types.List `tfsdk:"project_files"`
}

func (to *UnreviewProjectFilesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UnreviewProjectFilesResponse) {
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() && to.ProjectFiles.IsNull() && len(from.ProjectFiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProjectFiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProjectFiles = from.ProjectFiles
	}
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() {
		if toProjectFiles, ok := to.GetProjectFiles(ctx); ok {
			if fromProjectFiles, ok := from.GetProjectFiles(ctx); ok {
				// Recursively sync the fields of each ProjectFiles element by position.
				for i := range toProjectFiles {
					if i < len(fromProjectFiles) {
						toProjectFiles[i].SyncFieldsDuringCreateOrUpdate(ctx, fromProjectFiles[i])
					}
				}
				to.SetProjectFiles(ctx, toProjectFiles)
			}
		}
	}
}

func (to *UnreviewProjectFilesResponse) SyncFieldsDuringRead(ctx context.Context, from UnreviewProjectFilesResponse) {
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() && to.ProjectFiles.IsNull() && len(from.ProjectFiles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ProjectFiles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ProjectFiles = from.ProjectFiles
	}
	if !from.ProjectFiles.IsNull() && !from.ProjectFiles.IsUnknown() {
		if toProjectFiles, ok := to.GetProjectFiles(ctx); ok {
			if fromProjectFiles, ok := from.GetProjectFiles(ctx); ok {
				for i := range toProjectFiles {
					if i < len(fromProjectFiles) {
						toProjectFiles[i].SyncFieldsDuringRead(ctx, fromProjectFiles[i])
					}
				}
				to.SetProjectFiles(ctx, toProjectFiles)
			}
		}
	}
}

func (m UnreviewProjectFilesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_files"] = attrs["project_files"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnreviewProjectFilesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UnreviewProjectFilesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project_files": reflect.TypeOf(ProjectFile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnreviewProjectFilesResponse
// only implements ToObjectValue() and Type().
func (m UnreviewProjectFilesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_files": m.ProjectFiles,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UnreviewProjectFilesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_files": basetypes.ListType{
				ElemType: ProjectFile{}.Type(ctx),
			},
		},
	}
}

// GetProjectFiles returns the value of the ProjectFiles field in UnreviewProjectFilesResponse as
// a slice of ProjectFile values.
// If the field is unknown or null, the boolean return value is false.
func (m *UnreviewProjectFilesResponse) GetProjectFiles(ctx context.Context) ([]ProjectFile, bool) {
	if m.ProjectFiles.IsNull() || m.ProjectFiles.IsUnknown() {
		return nil, false
	}
	var v []ProjectFile
	d := m.ProjectFiles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjectFiles sets the value of the ProjectFiles field in UnreviewProjectFilesResponse.
func (m *UnreviewProjectFilesResponse) SetProjectFiles(ctx context.Context, v []ProjectFile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["project_files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ProjectFiles = types.ListValueMust(t, vs)
}

type UpdateProjectRequest struct {
	// The resource name of the project, in the format `projects/{project_id}`.
	// For example, `projects/123`. The server assigns it when the project is
	// created.
	Name types.String `tfsdk:"-"`
	// The project to update, identified by `name`.
	Project types.Object `tfsdk:"project"`
	// The fields to update. Supported fields are `display_name`,
	// `project_config.source_dialect`, `project_config.target_framework`, and
	// `project_config.lineage_location`. `project_config` and `*` update all
	// three `project_config` fields and can't be combined with `display_name`.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				// Recursively sync the fields of Project
				toProject.SyncFieldsDuringCreateOrUpdate(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (to *UpdateProjectRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m UpdateProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProjectRequest
// only implements ToObjectValue() and Type().
func (m UpdateProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"project":     m.Project,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"project":     Project{}.Type(ctx),
			"update_mask": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in UpdateProjectRequest as
// a Project value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateProjectRequest) GetProject(ctx context.Context) (Project, bool) {
	var e Project
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v Project
	d := m.Project.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProject sets the value of the Project field in UpdateProjectRequest.
func (m *UpdateProjectRequest) SetProject(ctx context.Context, v Project) {
	vs := v.ToObjectValue(ctx)
	m.Project = vs
}

// A file or folder in the workspace. On input, set `path` or `object_id`; if
// both are set, `object_id` is used. Responses set both when known.
type WorkspaceLocation struct {
	// The workspace object ID, as returned in `object_id` by the Workspace API.
	// For example, `12345`.
	ObjectId types.Int64 `tfsdk:"object_id"`
	// The absolute workspace path. For example,
	// `/Workspace/Shared/migration/source.sql`.
	Path types.String `tfsdk:"path"`
}

func (to *WorkspaceLocation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceLocation) {
}

func (to *WorkspaceLocation) SyncFieldsDuringRead(ctx context.Context, from WorkspaceLocation) {
}

func (m WorkspaceLocation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceLocation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceLocation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceLocation
// only implements ToObjectValue() and Type().
func (m WorkspaceLocation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"object_id": m.ObjectId,
			"path":      m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceLocation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"object_id": types.Int64Type,
			"path":      types.StringType,
		},
	}
}
