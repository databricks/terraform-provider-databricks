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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ComputeSpec struct {
	// Idle duration after which the sandbox is automatically terminated.
	InactivityTimeout timetypes.GoDuration `tfsdk:"inactivity_timeout"`
}

func (to *ComputeSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComputeSpec) {
}

func (to *ComputeSpec) SyncFieldsDuringRead(ctx context.Context, from ComputeSpec) {
}

func (m ComputeSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ComputeSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComputeSpec
// only implements ToObjectValue() and Type().
func (m ComputeSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inactivity_timeout": m.InactivityTimeout,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComputeSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inactivity_timeout": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

type CreateSandboxRequest struct {
	// The sandbox to create.
	Sandbox types.Object `tfsdk:"sandbox"`
	// Client-supplied ID that becomes the final path segment of the resource
	// name.
	SandboxId types.String `tfsdk:"-"`
}

func (to *CreateSandboxRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSandboxRequest) {
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

func (to *CreateSandboxRequest) SyncFieldsDuringRead(ctx context.Context, from CreateSandboxRequest) {
	if !from.Sandbox.IsNull() && !from.Sandbox.IsUnknown() {
		if toSandbox, ok := to.GetSandbox(ctx); ok {
			if fromSandbox, ok := from.GetSandbox(ctx); ok {
				toSandbox.SyncFieldsDuringRead(ctx, fromSandbox)
				to.SetSandbox(ctx, toSandbox)
			}
		}
	}
}

func (m CreateSandboxRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sandbox"] = attrs["sandbox"].SetRequired()
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
func (m CreateSandboxRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sandbox": reflect.TypeOf(Sandbox{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSandboxRequest
// only implements ToObjectValue() and Type().
func (m CreateSandboxRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sandbox":    m.Sandbox,
			"sandbox_id": m.SandboxId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSandboxRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sandbox":    Sandbox{}.Type(ctx),
			"sandbox_id": types.StringType,
		},
	}
}

// GetSandbox returns the value of the Sandbox field in CreateSandboxRequest as
// a Sandbox value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSandboxRequest) GetSandbox(ctx context.Context) (Sandbox, bool) {
	var e Sandbox
	if m.Sandbox.IsNull() || m.Sandbox.IsUnknown() {
		return e, false
	}
	var v Sandbox
	d := m.Sandbox.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSandbox sets the value of the Sandbox field in CreateSandboxRequest.
func (m *CreateSandboxRequest) SetSandbox(ctx context.Context, v Sandbox) {
	vs := v.ToObjectValue(ctx)
	m.Sandbox = vs
}

type DeleteSandboxRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSandboxRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSandboxRequest) {
}

func (to *DeleteSandboxRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteSandboxRequest) {
}

func (m DeleteSandboxRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteSandboxRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSandboxRequest
// only implements ToObjectValue() and Type().
func (m DeleteSandboxRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSandboxRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Request to run a command in the given sandbox and wait for it to finish.
type ExecuteCommandSyncRequest struct {
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

func (to *ExecuteCommandSyncRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteCommandSyncRequest) {
	if !from.Args.IsNull() && !from.Args.IsUnknown() && to.Args.IsNull() && len(from.Args.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Args, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Args = from.Args
	}
}

func (to *ExecuteCommandSyncRequest) SyncFieldsDuringRead(ctx context.Context, from ExecuteCommandSyncRequest) {
	if !from.Args.IsNull() && !from.Args.IsUnknown() && to.Args.IsNull() && len(from.Args.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Args, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Args = from.Args
	}
}

func (m ExecuteCommandSyncRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExecuteCommandSyncRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"args": reflect.TypeOf(types.String{}),
		"envs": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteCommandSyncRequest
// only implements ToObjectValue() and Type().
func (m ExecuteCommandSyncRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExecuteCommandSyncRequest) Type(ctx context.Context) attr.Type {
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

// GetArgs returns the value of the Args field in ExecuteCommandSyncRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteCommandSyncRequest) GetArgs(ctx context.Context) ([]types.String, bool) {
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

// SetArgs sets the value of the Args field in ExecuteCommandSyncRequest.
func (m *ExecuteCommandSyncRequest) SetArgs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["args"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Args = types.ListValueMust(t, vs)
}

// GetEnvs returns the value of the Envs field in ExecuteCommandSyncRequest as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteCommandSyncRequest) GetEnvs(ctx context.Context) (map[string]types.String, bool) {
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

// SetEnvs sets the value of the Envs field in ExecuteCommandSyncRequest.
func (m *ExecuteCommandSyncRequest) SetEnvs(ctx context.Context, v map[string]types.String) {
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
type ExecuteCommandSyncResponse struct {
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

func (to *ExecuteCommandSyncResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteCommandSyncResponse) {
}

func (to *ExecuteCommandSyncResponse) SyncFieldsDuringRead(ctx context.Context, from ExecuteCommandSyncResponse) {
}

func (m ExecuteCommandSyncResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExecuteCommandSyncResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteCommandSyncResponse
// only implements ToObjectValue() and Type().
func (m ExecuteCommandSyncResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExecuteCommandSyncResponse) Type(ctx context.Context) attr.Type {
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

type GetSandboxRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetSandboxRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSandboxRequest) {
}

func (to *GetSandboxRequest) SyncFieldsDuringRead(ctx context.Context, from GetSandboxRequest) {
}

func (m GetSandboxRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSandboxRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSandboxRequest
// only implements ToObjectValue() and Type().
func (m GetSandboxRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSandboxRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListSandboxesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListSandboxesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSandboxesRequest) {
}

func (to *ListSandboxesRequest) SyncFieldsDuringRead(ctx context.Context, from ListSandboxesRequest) {
}

func (m ListSandboxesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSandboxesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSandboxesRequest
// only implements ToObjectValue() and Type().
func (m ListSandboxesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSandboxesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// A list of Sandboxes.
type ListSandboxesResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Sandboxes types.List `tfsdk:"sandboxes"`
}

func (to *ListSandboxesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSandboxesResponse) {
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

func (to *ListSandboxesResponse) SyncFieldsDuringRead(ctx context.Context, from ListSandboxesResponse) {
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

func (m ListSandboxesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSandboxesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sandboxes": reflect.TypeOf(Sandbox{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSandboxesResponse
// only implements ToObjectValue() and Type().
func (m ListSandboxesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"sandboxes":       m.Sandboxes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSandboxesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"sandboxes": basetypes.ListType{
				ElemType: Sandbox{}.Type(ctx),
			},
		},
	}
}

// GetSandboxes returns the value of the Sandboxes field in ListSandboxesResponse as
// a slice of Sandbox values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSandboxesResponse) GetSandboxes(ctx context.Context) ([]Sandbox, bool) {
	if m.Sandboxes.IsNull() || m.Sandboxes.IsUnknown() {
		return nil, false
	}
	var v []Sandbox
	d := m.Sandboxes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSandboxes sets the value of the Sandboxes field in ListSandboxesResponse.
func (m *ListSandboxesResponse) SetSandboxes(ctx context.Context, v []Sandbox) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sandboxes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Sandboxes = types.ListValueMust(t, vs)
}

// A Sandbox resource representing an execution environment.
type Sandbox struct {
	// Output only. The creation time of the sandbox.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Human-readable display label for the sandbox. At most 256 bytes.
	DisplayName types.String `tfsdk:"display_name"`
	// The AIP-compliant resource name, such as "sandboxes/my-sandbox".
	Name types.String `tfsdk:"name"`
	// The desired configuration of the sandbox, supplied by the caller at
	// creation time.
	Spec types.Object `tfsdk:"spec"`
	// The observed runtime state of the sandbox, populated by the server.
	Status types.Object `tfsdk:"status"`
	// Output only. The last update time of the sandbox metadata and spec.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Sandbox) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Sandbox) {
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

func (to *Sandbox) SyncFieldsDuringRead(ctx context.Context, from Sandbox) {
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

func (m Sandbox) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
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
func (m Sandbox) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(SandboxSpec{}),
		"status": reflect.TypeOf(SandboxStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Sandbox
// only implements ToObjectValue() and Type().
func (m Sandbox) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Sandbox) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":  timetypes.RFC3339{}.Type(ctx),
			"display_name": types.StringType,
			"name":         types.StringType,
			"spec":         SandboxSpec{}.Type(ctx),
			"status":       SandboxStatus{}.Type(ctx),
			"update_time":  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Sandbox as
// a SandboxSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Sandbox) GetSpec(ctx context.Context) (SandboxSpec, bool) {
	var e SandboxSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v SandboxSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Sandbox.
func (m *Sandbox) SetSpec(ctx context.Context, v SandboxSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Sandbox as
// a SandboxStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Sandbox) GetStatus(ctx context.Context) (SandboxStatus, bool) {
	var e SandboxStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v SandboxStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Sandbox.
func (m *Sandbox) SetStatus(ctx context.Context, v SandboxStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

type SandboxSpec struct {
	// Compute configuration (size, inactivity timeout) requested for the
	// sandbox.
	Compute types.Object `tfsdk:"compute"`
}

func (to *SandboxSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SandboxSpec) {
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

func (to *SandboxSpec) SyncFieldsDuringRead(ctx context.Context, from SandboxSpec) {
	if !from.Compute.IsNull() && !from.Compute.IsUnknown() {
		if toCompute, ok := to.GetCompute(ctx); ok {
			if fromCompute, ok := from.GetCompute(ctx); ok {
				toCompute.SyncFieldsDuringRead(ctx, fromCompute)
				to.SetCompute(ctx, toCompute)
			}
		}
	}
}

func (m SandboxSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compute"] = attrs["compute"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SandboxSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SandboxSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compute": reflect.TypeOf(ComputeSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SandboxSpec
// only implements ToObjectValue() and Type().
func (m SandboxSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compute": m.Compute,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SandboxSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compute": ComputeSpec{}.Type(ctx),
		},
	}
}

// GetCompute returns the value of the Compute field in SandboxSpec as
// a ComputeSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *SandboxSpec) GetCompute(ctx context.Context) (ComputeSpec, bool) {
	var e ComputeSpec
	if m.Compute.IsNull() || m.Compute.IsUnknown() {
		return e, false
	}
	var v ComputeSpec
	d := m.Compute.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCompute sets the value of the Compute field in SandboxSpec.
func (m *SandboxSpec) SetCompute(ctx context.Context, v ComputeSpec) {
	vs := v.ToObjectValue(ctx)
	m.Compute = vs
}

type SandboxStatus struct {
	// Lifecycle state of the sandbox.
	State types.String `tfsdk:"state"`
}

func (to *SandboxStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SandboxStatus) {
}

func (to *SandboxStatus) SyncFieldsDuringRead(ctx context.Context, from SandboxStatus) {
}

func (m SandboxStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SandboxStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SandboxStatus
// only implements ToObjectValue() and Type().
func (m SandboxStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"state": m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SandboxStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"state": types.StringType,
		},
	}
}

// A request to start a Sandbox.
type StartSandboxRequest struct {
	// Resource name of the sandbox to start, in the form
	// `sandboxes/{sandbox_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *StartSandboxRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartSandboxRequest) {
}

func (to *StartSandboxRequest) SyncFieldsDuringRead(ctx context.Context, from StartSandboxRequest) {
}

func (m StartSandboxRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StartSandboxRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartSandboxRequest
// only implements ToObjectValue() and Type().
func (m StartSandboxRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartSandboxRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// A request to stop a Sandbox.
type StopSandboxRequest struct {
	// Resource name of the sandbox to stop, in the form
	// `sandboxes/{sandbox_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *StopSandboxRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopSandboxRequest) {
}

func (to *StopSandboxRequest) SyncFieldsDuringRead(ctx context.Context, from StopSandboxRequest) {
}

func (m StopSandboxRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StopSandboxRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopSandboxRequest
// only implements ToObjectValue() and Type().
func (m StopSandboxRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StopSandboxRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateSandboxRequest struct {
	// Resource name of the sandbox to update, in the form
	// `sandboxes/{sandbox_id}`.
	Name types.String `tfsdk:"-"`
	// The Sandbox resource carrying new field values. Only fields named in
	// `update_mask` are read; unmasked fields are ignored.
	Sandbox types.Object `tfsdk:"sandbox"`
	// Field paths to update. Must be a non-empty subset of: - display_name -
	// spec.compute.inactivity_timeout Any other path returns
	// INVALID_PARAMETER_VALUE.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSandboxRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSandboxRequest) {
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

func (to *UpdateSandboxRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateSandboxRequest) {
	if !from.Sandbox.IsNull() && !from.Sandbox.IsUnknown() {
		if toSandbox, ok := to.GetSandbox(ctx); ok {
			if fromSandbox, ok := from.GetSandbox(ctx); ok {
				toSandbox.SyncFieldsDuringRead(ctx, fromSandbox)
				to.SetSandbox(ctx, toSandbox)
			}
		}
	}
}

func (m UpdateSandboxRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sandbox"] = attrs["sandbox"].SetRequired()
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
func (m UpdateSandboxRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sandbox": reflect.TypeOf(Sandbox{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSandboxRequest
// only implements ToObjectValue() and Type().
func (m UpdateSandboxRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"sandbox":     m.Sandbox,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSandboxRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"sandbox":     Sandbox{}.Type(ctx),
			"update_mask": types.StringType,
		},
	}
}

// GetSandbox returns the value of the Sandbox field in UpdateSandboxRequest as
// a Sandbox value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSandboxRequest) GetSandbox(ctx context.Context) (Sandbox, bool) {
	var e Sandbox
	if m.Sandbox.IsNull() || m.Sandbox.IsUnknown() {
		return e, false
	}
	var v Sandbox
	d := m.Sandbox.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSandbox sets the value of the Sandbox field in UpdateSandboxRequest.
func (m *UpdateSandboxRequest) SetSandbox(ctx context.Context, v Sandbox) {
	vs := v.ToObjectValue(ctx)
	m.Sandbox = vs
}
