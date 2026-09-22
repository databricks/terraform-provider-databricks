// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package sandbox_tf

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

type ComputeSpec_SdkV2 struct {
	// Idle duration after which the sandbox is automatically terminated.
	InactivityTimeout timetypes.GoDuration `tfsdk:"inactivity_timeout"`
}

func (to *ComputeSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComputeSpec_SdkV2) {
}

func (to *ComputeSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ComputeSpec_SdkV2) {
}

func (m ComputeSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inactivity_timeout"] = attrs["inactivity_timeout"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComputeSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ComputeSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComputeSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m ComputeSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inactivity_timeout": m.InactivityTimeout,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComputeSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inactivity_timeout": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

type CreateSandboxRequest_SdkV2 struct {
	// The sandbox to create.
	Sandbox types.List `tfsdk:"sandbox"`
	// Client-supplied ID that becomes the final path segment of the resource
	// name.
	SandboxId types.String `tfsdk:"-"`
}

func (to *CreateSandboxRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSandboxRequest_SdkV2) {
	if !from.Sandbox.IsNull() && !from.Sandbox.IsUnknown() {
		if toSandbox, ok := to.GetSandbox(ctx); ok {
			if fromSandbox, ok := from.GetSandbox(ctx); ok {
				// Recursively sync the fields of Sandbox
				toSandbox.SyncFieldsDuringCreateOrUpdate(ctx, fromSandbox)
				to.SetSandbox(ctx, toSandbox)
			}
		}
	}
}

func (to *CreateSandboxRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateSandboxRequest_SdkV2) {
	if !from.Sandbox.IsNull() && !from.Sandbox.IsUnknown() {
		if toSandbox, ok := to.GetSandbox(ctx); ok {
			if fromSandbox, ok := from.GetSandbox(ctx); ok {
				toSandbox.SyncFieldsDuringRead(ctx, fromSandbox)
				to.SetSandbox(ctx, toSandbox)
			}
		}
	}
}

func (m CreateSandboxRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sandbox"] = attrs["sandbox"].SetRequired()
	attrs["sandbox"] = attrs["sandbox"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sandbox_id"] = attrs["sandbox_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSandboxRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateSandboxRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sandbox": reflect.TypeOf(Sandbox_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSandboxRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateSandboxRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sandbox":    m.Sandbox,
			"sandbox_id": m.SandboxId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSandboxRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sandbox": basetypes.ListType{
				ElemType: Sandbox_SdkV2{}.Type(ctx),
			},
			"sandbox_id": types.StringType,
		},
	}
}

// GetSandbox returns the value of the Sandbox field in CreateSandboxRequest_SdkV2 as
// a Sandbox_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSandboxRequest_SdkV2) GetSandbox(ctx context.Context) (Sandbox_SdkV2, bool) {
	var e Sandbox_SdkV2
	if m.Sandbox.IsNull() || m.Sandbox.IsUnknown() {
		return e, false
	}
	var v []Sandbox_SdkV2
	d := m.Sandbox.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSandbox sets the value of the Sandbox field in CreateSandboxRequest_SdkV2.
func (m *CreateSandboxRequest_SdkV2) SetSandbox(ctx context.Context, v Sandbox_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sandbox"]
	m.Sandbox = types.ListValueMust(t, vs)
}

type DeleteSandboxRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSandboxRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSandboxRequest_SdkV2) {
}

func (to *DeleteSandboxRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSandboxRequest_SdkV2) {
}

func (m DeleteSandboxRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSandboxRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSandboxRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSandboxRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSandboxRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSandboxRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Request to run a command in the given sandbox and wait for it to finish.
type ExecuteCommandSyncRequest_SdkV2 struct {
	// Arguments passed to `cmd`.
	Args types.List `tfsdk:"args"`
	// Executable or command to run (e.g. `/bin/echo`, `python3`).
	Cmd types.String `tfsdk:"cmd"`
	// Extra environment variables for the command's process, merged over the
	// sandbox's default environment.
	Envs types.Map `tfsdk:"envs"`
	// Maximum time to wait for the command to finish. When it elapses the
	// command is terminated and the response carries status `TIMED_OUT`. The
	// server applies a default when unset and clamps to an upper bound;
	// negative or otherwise invalid durations are rejected with
	// `INVALID_ARGUMENT`.
	ExecutionTimeout timetypes.GoDuration `tfsdk:"execution_timeout"`
	// Resource name of the sandbox to run the command in, in the form
	// `sandboxes/{sandbox_id}`. Bound from the URL path.
	Name types.String `tfsdk:"-"`
}

func (to *ExecuteCommandSyncRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteCommandSyncRequest_SdkV2) {
	if !from.Args.IsNull() && !from.Args.IsUnknown() && to.Args.IsNull() && len(from.Args.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Args, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Args = from.Args
	}
}

func (to *ExecuteCommandSyncRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExecuteCommandSyncRequest_SdkV2) {
	if !from.Args.IsNull() && !from.Args.IsUnknown() && to.Args.IsNull() && len(from.Args.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Args, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Args = from.Args
	}
}

func (m ExecuteCommandSyncRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["args"] = attrs["args"].SetOptional()
	attrs["cmd"] = attrs["cmd"].SetRequired()
	attrs["envs"] = attrs["envs"].SetOptional()
	attrs["execution_timeout"] = attrs["execution_timeout"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExecuteCommandSyncRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExecuteCommandSyncRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"args": reflect.TypeOf(types.String{}),
		"envs": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteCommandSyncRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExecuteCommandSyncRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"args":              m.Args,
			"cmd":               m.Cmd,
			"envs":              m.Envs,
			"execution_timeout": m.ExecutionTimeout,
			"name":              m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExecuteCommandSyncRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"args": basetypes.ListType{
				ElemType: types.StringType,
			},
			"cmd": types.StringType,
			"envs": basetypes.MapType{
				ElemType: types.StringType,
			},
			"execution_timeout": timetypes.GoDuration{}.Type(ctx),
			"name":              types.StringType,
		},
	}
}

// GetArgs returns the value of the Args field in ExecuteCommandSyncRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteCommandSyncRequest_SdkV2) GetArgs(ctx context.Context) ([]types.String, bool) {
	if m.Args.IsNull() || m.Args.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Args.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetArgs sets the value of the Args field in ExecuteCommandSyncRequest_SdkV2.
func (m *ExecuteCommandSyncRequest_SdkV2) SetArgs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["args"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Args = types.ListValueMust(t, vs)
}

// GetEnvs returns the value of the Envs field in ExecuteCommandSyncRequest_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteCommandSyncRequest_SdkV2) GetEnvs(ctx context.Context) (map[string]types.String, bool) {
	if m.Envs.IsNull() || m.Envs.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Envs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvs sets the value of the Envs field in ExecuteCommandSyncRequest_SdkV2.
func (m *ExecuteCommandSyncRequest_SdkV2) SetEnvs(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["envs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Envs = types.MapValueMust(t, vs)
}

// Result of a completed unary command execution: captured output, exit code,
// and terminal status.
type ExecuteCommandSyncResponse_SdkV2 struct {
	// Daemon-generated identifier for this command execution, for correlation
	// (for example in `ListCommands`).
	CommandId types.String `tfsdk:"command_id"`
	// Process exit code. Unset when the process was terminated by a signal
	// (e.g. on `TIMED_OUT`) or never started (`FAILED`) rather than exiting
	// normally.
	ExitCode types.Int64 `tfsdk:"exit_code"`
	// Terminal status of the command execution. Always set on a successful
	// response; never `EXECUTE_COMMAND_STATUS_UNSPECIFIED`.
	Status types.String `tfsdk:"status"`
	// Captured standard error, with the same UTF-8 semantics as `stdout`.
	Stderr types.String `tfsdk:"stderr"`
	// Captured standard output as UTF-8 text. Invalid UTF-8 bytes are replaced
	// with the Unicode replacement character (U+FFFD).
	Stdout types.String `tfsdk:"stdout"`
	// True when `stdout` / `stderr` were truncated because the captured output
	// exceeded the server's per-response size cap. The dropped output is not
	// included in this response and is not recoverable through this unary API.
	Truncated types.Bool `tfsdk:"truncated"`
}

func (to *ExecuteCommandSyncResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteCommandSyncResponse_SdkV2) {
}

func (to *ExecuteCommandSyncResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExecuteCommandSyncResponse_SdkV2) {
}

func (m ExecuteCommandSyncResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["command_id"] = attrs["command_id"].SetComputed()
	attrs["exit_code"] = attrs["exit_code"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["stderr"] = attrs["stderr"].SetComputed()
	attrs["stdout"] = attrs["stdout"].SetComputed()
	attrs["truncated"] = attrs["truncated"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExecuteCommandSyncResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExecuteCommandSyncResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteCommandSyncResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ExecuteCommandSyncResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"command_id": m.CommandId,
			"exit_code":  m.ExitCode,
			"status":     m.Status,
			"stderr":     m.Stderr,
			"stdout":     m.Stdout,
			"truncated":  m.Truncated,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExecuteCommandSyncResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"command_id": types.StringType,
			"exit_code":  types.Int64Type,
			"status":     types.StringType,
			"stderr":     types.StringType,
			"stdout":     types.StringType,
			"truncated":  types.BoolType,
		},
	}
}

type GetSandboxRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetSandboxRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSandboxRequest_SdkV2) {
}

func (to *GetSandboxRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSandboxRequest_SdkV2) {
}

func (m GetSandboxRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSandboxRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSandboxRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSandboxRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSandboxRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSandboxRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListSandboxesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListSandboxesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSandboxesRequest_SdkV2) {
}

func (to *ListSandboxesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSandboxesRequest_SdkV2) {
}

func (m ListSandboxesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSandboxesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSandboxesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSandboxesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSandboxesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSandboxesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// A list of Sandboxes.
type ListSandboxesResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Sandboxes types.List `tfsdk:"sandboxes"`
}

func (to *ListSandboxesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSandboxesResponse_SdkV2) {
	if !from.Sandboxes.IsNull() && !from.Sandboxes.IsUnknown() && to.Sandboxes.IsNull() && len(from.Sandboxes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Sandboxes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Sandboxes = from.Sandboxes
	}
	if !from.Sandboxes.IsNull() && !from.Sandboxes.IsUnknown() {
		if toSandboxes, ok := to.GetSandboxes(ctx); ok {
			if fromSandboxes, ok := from.GetSandboxes(ctx); ok {
				// Recursively sync the fields of each Sandboxes element by position.
				for i := range toSandboxes {
					if i < len(fromSandboxes) {
						toSandboxes[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSandboxes[i])
					}
				}
				to.SetSandboxes(ctx, toSandboxes)
			}
		}
	}
}

func (to *ListSandboxesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSandboxesResponse_SdkV2) {
	if !from.Sandboxes.IsNull() && !from.Sandboxes.IsUnknown() && to.Sandboxes.IsNull() && len(from.Sandboxes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Sandboxes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Sandboxes = from.Sandboxes
	}
	if !from.Sandboxes.IsNull() && !from.Sandboxes.IsUnknown() {
		if toSandboxes, ok := to.GetSandboxes(ctx); ok {
			if fromSandboxes, ok := from.GetSandboxes(ctx); ok {
				for i := range toSandboxes {
					if i < len(fromSandboxes) {
						toSandboxes[i].SyncFieldsDuringRead(ctx, fromSandboxes[i])
					}
				}
				to.SetSandboxes(ctx, toSandboxes)
			}
		}
	}
}

func (m ListSandboxesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["sandboxes"] = attrs["sandboxes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSandboxesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSandboxesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sandboxes": reflect.TypeOf(Sandbox_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSandboxesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSandboxesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"sandboxes":       m.Sandboxes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSandboxesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"sandboxes": basetypes.ListType{
				ElemType: Sandbox_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSandboxes returns the value of the Sandboxes field in ListSandboxesResponse_SdkV2 as
// a slice of Sandbox_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSandboxesResponse_SdkV2) GetSandboxes(ctx context.Context) ([]Sandbox_SdkV2, bool) {
	if m.Sandboxes.IsNull() || m.Sandboxes.IsUnknown() {
		return nil, false
	}
	var v []Sandbox_SdkV2
	d := m.Sandboxes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSandboxes sets the value of the Sandboxes field in ListSandboxesResponse_SdkV2.
func (m *ListSandboxesResponse_SdkV2) SetSandboxes(ctx context.Context, v []Sandbox_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sandboxes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Sandboxes = types.ListValueMust(t, vs)
}

// A Sandbox resource representing an execution environment.
type Sandbox_SdkV2 struct {
	// Output only. The creation time of the sandbox.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Human-readable display label for the sandbox. At most 256 bytes.
	DisplayName types.String `tfsdk:"display_name"`
	// The AIP-compliant resource name, such as "sandboxes/my-sandbox".
	Name types.String `tfsdk:"name"`
	// The desired configuration of the sandbox, supplied by the caller at
	// creation time.
	Spec types.List `tfsdk:"spec"`
	// The observed runtime state of the sandbox, populated by the server.
	Status types.List `tfsdk:"status"`
	// Output only. The last update time of the sandbox metadata and spec.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Sandbox_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Sandbox_SdkV2) {
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *Sandbox_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Sandbox_SdkV2) {
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m Sandbox_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Sandbox.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Sandbox_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(SandboxSpec_SdkV2{}),
		"status": reflect.TypeOf(SandboxStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Sandbox_SdkV2
// only implements ToObjectValue() and Type().
func (m Sandbox_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":  m.CreateTime,
			"display_name": m.DisplayName,
			"name":         m.Name,
			"spec":         m.Spec,
			"status":       m.Status,
			"update_time":  m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Sandbox_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":  timetypes.RFC3339{}.Type(ctx),
			"display_name": types.StringType,
			"name":         types.StringType,
			"spec": basetypes.ListType{
				ElemType: SandboxSpec_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: SandboxStatus_SdkV2{}.Type(ctx),
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Sandbox_SdkV2 as
// a SandboxSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Sandbox_SdkV2) GetSpec(ctx context.Context) (SandboxSpec_SdkV2, bool) {
	var e SandboxSpec_SdkV2
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v []SandboxSpec_SdkV2
	d := m.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in Sandbox_SdkV2.
func (m *Sandbox_SdkV2) SetSpec(ctx context.Context, v SandboxSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	m.Spec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Sandbox_SdkV2 as
// a SandboxStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Sandbox_SdkV2) GetStatus(ctx context.Context) (SandboxStatus_SdkV2, bool) {
	var e SandboxStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []SandboxStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Sandbox_SdkV2.
func (m *Sandbox_SdkV2) SetStatus(ctx context.Context, v SandboxStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

type SandboxSpec_SdkV2 struct {
	// Compute configuration (size, inactivity timeout) requested for the
	// sandbox.
	Compute types.List `tfsdk:"compute"`
}

func (to *SandboxSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SandboxSpec_SdkV2) {
	if !from.Compute.IsNull() && !from.Compute.IsUnknown() {
		if toCompute, ok := to.GetCompute(ctx); ok {
			if fromCompute, ok := from.GetCompute(ctx); ok {
				// Recursively sync the fields of Compute
				toCompute.SyncFieldsDuringCreateOrUpdate(ctx, fromCompute)
				to.SetCompute(ctx, toCompute)
			}
		}
	}
}

func (to *SandboxSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SandboxSpec_SdkV2) {
	if !from.Compute.IsNull() && !from.Compute.IsUnknown() {
		if toCompute, ok := to.GetCompute(ctx); ok {
			if fromCompute, ok := from.GetCompute(ctx); ok {
				toCompute.SyncFieldsDuringRead(ctx, fromCompute)
				to.SetCompute(ctx, toCompute)
			}
		}
	}
}

func (m SandboxSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compute"] = attrs["compute"].SetOptional()
	attrs["compute"] = attrs["compute"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SandboxSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SandboxSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compute": reflect.TypeOf(ComputeSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SandboxSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m SandboxSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compute": m.Compute,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SandboxSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compute": basetypes.ListType{
				ElemType: ComputeSpec_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCompute returns the value of the Compute field in SandboxSpec_SdkV2 as
// a ComputeSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SandboxSpec_SdkV2) GetCompute(ctx context.Context) (ComputeSpec_SdkV2, bool) {
	var e ComputeSpec_SdkV2
	if m.Compute.IsNull() || m.Compute.IsUnknown() {
		return e, false
	}
	var v []ComputeSpec_SdkV2
	d := m.Compute.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCompute sets the value of the Compute field in SandboxSpec_SdkV2.
func (m *SandboxSpec_SdkV2) SetCompute(ctx context.Context, v ComputeSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compute"]
	m.Compute = types.ListValueMust(t, vs)
}

type SandboxStatus_SdkV2 struct {
	// Lifecycle state of the sandbox.
	State types.String `tfsdk:"state"`
}

func (to *SandboxStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SandboxStatus_SdkV2) {
}

func (to *SandboxStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SandboxStatus_SdkV2) {
}

func (m SandboxStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SandboxStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SandboxStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SandboxStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m SandboxStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"state": m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SandboxStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"state": types.StringType,
		},
	}
}

// A request to start a Sandbox.
type StartSandboxRequest_SdkV2 struct {
	// Resource name of the sandbox to start, in the form
	// `sandboxes/{sandbox_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *StartSandboxRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartSandboxRequest_SdkV2) {
}

func (to *StartSandboxRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StartSandboxRequest_SdkV2) {
}

func (m StartSandboxRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartSandboxRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StartSandboxRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartSandboxRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m StartSandboxRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartSandboxRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// A request to stop a Sandbox.
type StopSandboxRequest_SdkV2 struct {
	// Resource name of the sandbox to stop, in the form
	// `sandboxes/{sandbox_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *StopSandboxRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopSandboxRequest_SdkV2) {
}

func (to *StopSandboxRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StopSandboxRequest_SdkV2) {
}

func (m StopSandboxRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopSandboxRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StopSandboxRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopSandboxRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m StopSandboxRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StopSandboxRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateSandboxRequest_SdkV2 struct {
	// Resource name of the sandbox to update, in the form
	// `sandboxes/{sandbox_id}`.
	Name types.String `tfsdk:"-"`
	// The Sandbox resource carrying new field values. Only fields named in
	// `update_mask` are read; unmasked fields are ignored.
	Sandbox types.List `tfsdk:"sandbox"`
	// Field paths to update. Must be a non-empty subset of: - display_name -
	// spec.compute.inactivity_timeout Any other path returns
	// INVALID_PARAMETER_VALUE.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSandboxRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSandboxRequest_SdkV2) {
	if !from.Sandbox.IsNull() && !from.Sandbox.IsUnknown() {
		if toSandbox, ok := to.GetSandbox(ctx); ok {
			if fromSandbox, ok := from.GetSandbox(ctx); ok {
				// Recursively sync the fields of Sandbox
				toSandbox.SyncFieldsDuringCreateOrUpdate(ctx, fromSandbox)
				to.SetSandbox(ctx, toSandbox)
			}
		}
	}
}

func (to *UpdateSandboxRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSandboxRequest_SdkV2) {
	if !from.Sandbox.IsNull() && !from.Sandbox.IsUnknown() {
		if toSandbox, ok := to.GetSandbox(ctx); ok {
			if fromSandbox, ok := from.GetSandbox(ctx); ok {
				toSandbox.SyncFieldsDuringRead(ctx, fromSandbox)
				to.SetSandbox(ctx, toSandbox)
			}
		}
	}
}

func (m UpdateSandboxRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sandbox"] = attrs["sandbox"].SetRequired()
	attrs["sandbox"] = attrs["sandbox"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSandboxRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSandboxRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sandbox": reflect.TypeOf(Sandbox_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSandboxRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSandboxRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"sandbox":     m.Sandbox,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSandboxRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"sandbox": basetypes.ListType{
				ElemType: Sandbox_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetSandbox returns the value of the Sandbox field in UpdateSandboxRequest_SdkV2 as
// a Sandbox_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSandboxRequest_SdkV2) GetSandbox(ctx context.Context) (Sandbox_SdkV2, bool) {
	var e Sandbox_SdkV2
	if m.Sandbox.IsNull() || m.Sandbox.IsUnknown() {
		return e, false
	}
	var v []Sandbox_SdkV2
	d := m.Sandbox.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSandbox sets the value of the Sandbox field in UpdateSandboxRequest_SdkV2.
func (m *UpdateSandboxRequest_SdkV2) SetSandbox(ctx context.Context, v Sandbox_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sandbox"]
	m.Sandbox = types.ListValueMust(t, vs)
}
