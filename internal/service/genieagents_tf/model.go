// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package genieagents_tf

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

type CreateGenieAgentRequest struct {
	// The Genie Agent to create.
	GenieAgent types.Object `tfsdk:"genie_agent"`
	// The ID for the new Genie Agent: 32 lowercase hexadecimal characters (a
	// UUID without dashes), not all zeros. Assigned by the server if omitted.
	// Fails with ALREADY_EXISTS if the ID is in use in the workspace, including
	// by a Genie Agent in the trash, or PERMISSION_DENIED if the caller cannot
	// view the Genie Agent using it.
	GenieAgentId types.String `tfsdk:"-"`
}

func (to *CreateGenieAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGenieAgentRequest) {
	if !from.GenieAgent.IsNull() && !from.GenieAgent.IsUnknown() {
		if toGenieAgent, ok := to.GetGenieAgent(ctx); ok {
			if fromGenieAgent, ok := from.GetGenieAgent(ctx); ok {
				// Recursively sync the fields of GenieAgent
				toGenieAgent.SyncFieldsDuringCreateOrUpdate(ctx, fromGenieAgent)
				to.SetGenieAgent(ctx, toGenieAgent)
			}
		}
	}
}

func (to *CreateGenieAgentRequest) SyncFieldsDuringRead(ctx context.Context, from CreateGenieAgentRequest) {
	if !from.GenieAgent.IsNull() && !from.GenieAgent.IsUnknown() {
		if toGenieAgent, ok := to.GetGenieAgent(ctx); ok {
			if fromGenieAgent, ok := from.GetGenieAgent(ctx); ok {
				toGenieAgent.SyncFieldsDuringRead(ctx, fromGenieAgent)
				to.SetGenieAgent(ctx, toGenieAgent)
			}
		}
	}
}

func (m CreateGenieAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["genie_agent"] = attrs["genie_agent"].SetRequired()
	attrs["genie_agent_id"] = attrs["genie_agent_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateGenieAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateGenieAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"genie_agent": reflect.TypeOf(GenieAgent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGenieAgentRequest
// only implements ToObjectValue() and Type().
func (m CreateGenieAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"genie_agent":    m.GenieAgent,
			"genie_agent_id": m.GenieAgentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGenieAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"genie_agent":    GenieAgent{}.Type(ctx),
			"genie_agent_id": types.StringType,
		},
	}
}

// GetGenieAgent returns the value of the GenieAgent field in CreateGenieAgentRequest as
// a GenieAgent value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateGenieAgentRequest) GetGenieAgent(ctx context.Context) (GenieAgent, bool) {
	var e GenieAgent
	if m.GenieAgent.IsNull() || m.GenieAgent.IsUnknown() {
		return e, false
	}
	var v GenieAgent
	d := m.GenieAgent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGenieAgent sets the value of the GenieAgent field in CreateGenieAgentRequest.
func (m *CreateGenieAgentRequest) SetGenieAgent(ctx context.Context, v GenieAgent) {
	vs := v.ToObjectValue(ctx)
	m.GenieAgent = vs
}

type DeleteGenieAgentRequest struct {
	// Etag from a previous read. On mismatch the delete fails with ABORTED and
	// nothing is deleted.
	Etag types.String `tfsdk:"-"`
	// The resource name of the Genie Agent to delete, in the form
	// `agents/{genie_agent_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteGenieAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteGenieAgentRequest) {
}

func (to *DeleteGenieAgentRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteGenieAgentRequest) {
}

func (m DeleteGenieAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGenieAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteGenieAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteGenieAgentRequest
// only implements ToObjectValue() and Type().
func (m DeleteGenieAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteGenieAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
		},
	}
}

// A Genie Agent resource.
type GenieAgent struct {
	// Time when the Genie Agent was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The definition of the Genie Agent: its sample questions, data sources,
	// instructions, and benchmarks. Must set at least one field. Returned on
	// get and on list with `GENIE_AGENT_VIEW_FULL`. Create and update return a
	// submitted definition unchanged; one the service would store differently,
	// or with a column config the workspace can't keep (unknown column, or a
	// display name while those are disabled), is rejected with
	// INVALID_PARAMETER_VALUE.
	Definition types.Object `tfsdk:"definition"`
	// Description of the Genie Agent.
	//
	// The maximum length is 4,096 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	Description types.String `tfsdk:"description"`
	// Display name of the Genie Agent. Must be unique within its folder
	// (effective_parent_path): a create, rename, or move that would collide
	// with an existing Genie Agent's name in the target folder is rejected with
	// ALREADY_EXISTS.
	//
	// The maximum length is 63 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	DisplayName types.String `tfsdk:"display_name"`
	// The workspace folder that contains the Genie Agent.
	EffectiveParentPath types.String `tfsdk:"effective_parent_path"`
	// Weak etag of `definition` only; metadata changes do not affect it. Send
	// it on update or delete to detect concurrent definition changes.
	Etag types.String `tfsdk:"etag"`
	// The resource name of the Genie Agent, in the form
	// `agents/{genie_agent_id}`. Ignored on create.
	Name types.String `tfsdk:"name"`
	// The workspace folder to create the Genie Agent in. If omitted, the
	// caller's home folder is used. Returned exactly as submitted; see
	// `effective_parent_path` for the current folder. On update, setting it
	// moves the Genie Agent and clearing it does not. At most 4,096 characters.
	ParentPath types.String `tfsdk:"parent_path"`
	// Time when the Genie Agent was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Warehouse associated with the Genie Agent.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (to *GenieAgent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgent) {
	if !from.Definition.IsNull() && !from.Definition.IsUnknown() {
		if toDefinition, ok := to.GetDefinition(ctx); ok {
			if fromDefinition, ok := from.GetDefinition(ctx); ok {
				// Recursively sync the fields of Definition
				toDefinition.SyncFieldsDuringCreateOrUpdate(ctx, fromDefinition)
				to.SetDefinition(ctx, toDefinition)
			}
		}
	}
}

func (to *GenieAgent) SyncFieldsDuringRead(ctx context.Context, from GenieAgent) {
	if !from.Definition.IsNull() && !from.Definition.IsUnknown() {
		if toDefinition, ok := to.GetDefinition(ctx); ok {
			if fromDefinition, ok := from.GetDefinition(ctx); ok {
				toDefinition.SyncFieldsDuringRead(ctx, fromDefinition)
				to.SetDefinition(ctx, toDefinition)
			}
		}
	}
}

func (m GenieAgent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["definition"] = attrs["definition"].SetRequired()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["effective_parent_path"] = attrs["effective_parent_path"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent_path"] = attrs["parent_path"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"definition": reflect.TypeOf(GenieAgentDefinition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgent
// only implements ToObjectValue() and Type().
func (m GenieAgent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":           m.CreateTime,
			"definition":            m.Definition,
			"description":           m.Description,
			"display_name":          m.DisplayName,
			"effective_parent_path": m.EffectiveParentPath,
			"etag":                  m.Etag,
			"name":                  m.Name,
			"parent_path":           m.ParentPath,
			"update_time":           m.UpdateTime,
			"warehouse_id":          m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":           timetypes.RFC3339{}.Type(ctx),
			"definition":            GenieAgentDefinition{}.Type(ctx),
			"description":           types.StringType,
			"display_name":          types.StringType,
			"effective_parent_path": types.StringType,
			"etag":                  types.StringType,
			"name":                  types.StringType,
			"parent_path":           types.StringType,
			"update_time":           timetypes.RFC3339{}.Type(ctx),
			"warehouse_id":          types.StringType,
		},
	}
}

// GetDefinition returns the value of the Definition field in GenieAgent as
// a GenieAgentDefinition value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgent) GetDefinition(ctx context.Context) (GenieAgentDefinition, bool) {
	var e GenieAgentDefinition
	if m.Definition.IsNull() || m.Definition.IsUnknown() {
		return e, false
	}
	var v GenieAgentDefinition
	d := m.Definition.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefinition sets the value of the Definition field in GenieAgent.
func (m *GenieAgent) SetDefinition(ctx context.Context, v GenieAgentDefinition) {
	vs := v.ToObjectValue(ctx)
	m.Definition = vs
}

// The expected answer to a benchmark question.
type GenieAgentBenchmarkAnswer struct {
	// The answer.
	Content types.String `tfsdk:"content"`
	// The format of the answer.
	Format types.String `tfsdk:"format"`
}

func (to *GenieAgentBenchmarkAnswer) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentBenchmarkAnswer) {
}

func (to *GenieAgentBenchmarkAnswer) SyncFieldsDuringRead(ctx context.Context, from GenieAgentBenchmarkAnswer) {
}

func (m GenieAgentBenchmarkAnswer) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetRequired()
	attrs["format"] = attrs["format"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentBenchmarkAnswer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentBenchmarkAnswer) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentBenchmarkAnswer
// only implements ToObjectValue() and Type().
func (m GenieAgentBenchmarkAnswer) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": m.Content,
			"format":  m.Format,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentBenchmarkAnswer) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"format":  types.StringType,
		},
	}
}

// A benchmark question with its expected answer.
type GenieAgentBenchmarkQuestion struct {
	// The expected answer. At most one answer is supported; omit it when the
	// question has no expected answer.
	Answer types.List `tfsdk:"answer"`
	// Notes on how to evaluate answers to the question.
	EvaluationNote types.String `tfsdk:"evaluation_note"`
	// Identifier of the question: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// The question.
	Question types.String `tfsdk:"question"`
}

func (to *GenieAgentBenchmarkQuestion) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentBenchmarkQuestion) {
	if !from.Answer.IsNull() && !from.Answer.IsUnknown() && to.Answer.IsNull() && len(from.Answer.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Answer, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Answer = from.Answer
	}
	if !from.Answer.IsNull() && !from.Answer.IsUnknown() {
		if toAnswer, ok := to.GetAnswer(ctx); ok {
			if fromAnswer, ok := from.GetAnswer(ctx); ok {
				// Recursively sync the fields of each Answer element by position.
				for i := range toAnswer {
					if i < len(fromAnswer) {
						toAnswer[i].SyncFieldsDuringCreateOrUpdate(ctx, fromAnswer[i])
					}
				}
				to.SetAnswer(ctx, toAnswer)
			}
		}
	}
}

func (to *GenieAgentBenchmarkQuestion) SyncFieldsDuringRead(ctx context.Context, from GenieAgentBenchmarkQuestion) {
	if !from.Answer.IsNull() && !from.Answer.IsUnknown() && to.Answer.IsNull() && len(from.Answer.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Answer, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Answer = from.Answer
	}
	if !from.Answer.IsNull() && !from.Answer.IsUnknown() {
		if toAnswer, ok := to.GetAnswer(ctx); ok {
			if fromAnswer, ok := from.GetAnswer(ctx); ok {
				for i := range toAnswer {
					if i < len(fromAnswer) {
						toAnswer[i].SyncFieldsDuringRead(ctx, fromAnswer[i])
					}
				}
				to.SetAnswer(ctx, toAnswer)
			}
		}
	}
}

func (m GenieAgentBenchmarkQuestion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["answer"] = attrs["answer"].SetOptional()
	attrs["evaluation_note"] = attrs["evaluation_note"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["question"] = attrs["question"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentBenchmarkQuestion.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentBenchmarkQuestion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"answer": reflect.TypeOf(GenieAgentBenchmarkAnswer{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentBenchmarkQuestion
// only implements ToObjectValue() and Type().
func (m GenieAgentBenchmarkQuestion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"answer":          m.Answer,
			"evaluation_note": m.EvaluationNote,
			"id":              m.Id,
			"question":        m.Question,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentBenchmarkQuestion) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"answer": basetypes.ListType{
				ElemType: GenieAgentBenchmarkAnswer{}.Type(ctx),
			},
			"evaluation_note": types.StringType,
			"id":              types.StringType,
			"question":        types.StringType,
		},
	}
}

// GetAnswer returns the value of the Answer field in GenieAgentBenchmarkQuestion as
// a slice of GenieAgentBenchmarkAnswer values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentBenchmarkQuestion) GetAnswer(ctx context.Context) ([]GenieAgentBenchmarkAnswer, bool) {
	if m.Answer.IsNull() || m.Answer.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentBenchmarkAnswer
	d := m.Answer.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnswer sets the value of the Answer field in GenieAgentBenchmarkQuestion.
func (m *GenieAgentBenchmarkQuestion) SetAnswer(ctx context.Context, v []GenieAgentBenchmarkAnswer) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["answer"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Answer = types.ListValueMust(t, vs)
}

// Benchmarks for evaluating a Genie Agent.
type GenieAgentBenchmarks struct {
	// The benchmark questions, sorted by `id` in descending order.
	Questions types.List `tfsdk:"questions"`
}

func (to *GenieAgentBenchmarks) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentBenchmarks) {
	if !from.Questions.IsNull() && !from.Questions.IsUnknown() && to.Questions.IsNull() && len(from.Questions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Questions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Questions = from.Questions
	}
	if !from.Questions.IsNull() && !from.Questions.IsUnknown() {
		if toQuestions, ok := to.GetQuestions(ctx); ok {
			if fromQuestions, ok := from.GetQuestions(ctx); ok {
				// Recursively sync the fields of each Questions element by position.
				for i := range toQuestions {
					if i < len(fromQuestions) {
						toQuestions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromQuestions[i])
					}
				}
				to.SetQuestions(ctx, toQuestions)
			}
		}
	}
}

func (to *GenieAgentBenchmarks) SyncFieldsDuringRead(ctx context.Context, from GenieAgentBenchmarks) {
	if !from.Questions.IsNull() && !from.Questions.IsUnknown() && to.Questions.IsNull() && len(from.Questions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Questions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Questions = from.Questions
	}
	if !from.Questions.IsNull() && !from.Questions.IsUnknown() {
		if toQuestions, ok := to.GetQuestions(ctx); ok {
			if fromQuestions, ok := from.GetQuestions(ctx); ok {
				for i := range toQuestions {
					if i < len(fromQuestions) {
						toQuestions[i].SyncFieldsDuringRead(ctx, fromQuestions[i])
					}
				}
				to.SetQuestions(ctx, toQuestions)
			}
		}
	}
}

func (m GenieAgentBenchmarks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["questions"] = attrs["questions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentBenchmarks.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentBenchmarks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"questions": reflect.TypeOf(GenieAgentBenchmarkQuestion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentBenchmarks
// only implements ToObjectValue() and Type().
func (m GenieAgentBenchmarks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"questions": m.Questions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentBenchmarks) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"questions": basetypes.ListType{
				ElemType: GenieAgentBenchmarkQuestion{}.Type(ctx),
			},
		},
	}
}

// GetQuestions returns the value of the Questions field in GenieAgentBenchmarks as
// a slice of GenieAgentBenchmarkQuestion values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentBenchmarks) GetQuestions(ctx context.Context) ([]GenieAgentBenchmarkQuestion, bool) {
	if m.Questions.IsNull() || m.Questions.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentBenchmarkQuestion
	d := m.Questions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuestions sets the value of the Questions field in GenieAgentBenchmarks.
func (m *GenieAgentBenchmarks) SetQuestions(ctx context.Context, v []GenieAgentBenchmarkQuestion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["questions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Questions = types.ListValueMust(t, vs)
}

// Settings for one column of a table or metric view.
type GenieAgentColumnConfig struct {
	// The name of the column. Must exist in the table or metric view.
	Column types.String `tfsdk:"column"`
	// Description of the column, overriding its Unity Catalog comment.
	Description types.String `tfsdk:"description"`
	// Display name of the column in the Genie Agent. Accepted only where column
	// display names are enabled for the workspace.
	DisplayName types.String `tfsdk:"display_name"`
	// Whether the Genie Agent indexes the column's values to match them in
	// questions.
	EnableEntityMatching types.Bool `tfsdk:"enable_entity_matching"`
	// Whether the Genie Agent samples example values from the column to learn
	// its format.
	EnableFormatAssistance types.Bool `tfsdk:"enable_format_assistance"`
	// Whether the column is hidden from the Genie Agent.
	Exclude types.Bool `tfsdk:"exclude"`
	// Synonyms for the column name.
	Synonyms types.List `tfsdk:"synonyms"`
}

func (to *GenieAgentColumnConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentColumnConfig) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (to *GenieAgentColumnConfig) SyncFieldsDuringRead(ctx context.Context, from GenieAgentColumnConfig) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (m GenieAgentColumnConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column"] = attrs["column"].SetRequired()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["enable_entity_matching"] = attrs["enable_entity_matching"].SetOptional()
	attrs["enable_format_assistance"] = attrs["enable_format_assistance"].SetOptional()
	attrs["exclude"] = attrs["exclude"].SetOptional()
	attrs["synonyms"] = attrs["synonyms"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentColumnConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentColumnConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synonyms": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentColumnConfig
// only implements ToObjectValue() and Type().
func (m GenieAgentColumnConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column":                   m.Column,
			"description":              m.Description,
			"display_name":             m.DisplayName,
			"enable_entity_matching":   m.EnableEntityMatching,
			"enable_format_assistance": m.EnableFormatAssistance,
			"exclude":                  m.Exclude,
			"synonyms":                 m.Synonyms,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentColumnConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column":                   types.StringType,
			"description":              types.StringType,
			"display_name":             types.StringType,
			"enable_entity_matching":   types.BoolType,
			"enable_format_assistance": types.BoolType,
			"exclude":                  types.BoolType,
			"synonyms": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetSynonyms returns the value of the Synonyms field in GenieAgentColumnConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentColumnConfig) GetSynonyms(ctx context.Context) ([]types.String, bool) {
	if m.Synonyms.IsNull() || m.Synonyms.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Synonyms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSynonyms sets the value of the Synonyms field in GenieAgentColumnConfig.
func (m *GenieAgentColumnConfig) SetSynonyms(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["synonyms"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Synonyms = types.ListValueMust(t, vs)
}

// General settings of a Genie Agent. The warehouse, display name, and
// description are separate fields of the Genie Agent.
type GenieAgentConfig struct {
	// Sample questions shown to users of the Genie Agent, sorted by `id`.
	SampleQuestions types.List `tfsdk:"sample_questions"`
}

func (to *GenieAgentConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentConfig) {
	if !from.SampleQuestions.IsNull() && !from.SampleQuestions.IsUnknown() && to.SampleQuestions.IsNull() && len(from.SampleQuestions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SampleQuestions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SampleQuestions = from.SampleQuestions
	}
	if !from.SampleQuestions.IsNull() && !from.SampleQuestions.IsUnknown() {
		if toSampleQuestions, ok := to.GetSampleQuestions(ctx); ok {
			if fromSampleQuestions, ok := from.GetSampleQuestions(ctx); ok {
				// Recursively sync the fields of each SampleQuestions element by position.
				for i := range toSampleQuestions {
					if i < len(fromSampleQuestions) {
						toSampleQuestions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSampleQuestions[i])
					}
				}
				to.SetSampleQuestions(ctx, toSampleQuestions)
			}
		}
	}
}

func (to *GenieAgentConfig) SyncFieldsDuringRead(ctx context.Context, from GenieAgentConfig) {
	if !from.SampleQuestions.IsNull() && !from.SampleQuestions.IsUnknown() && to.SampleQuestions.IsNull() && len(from.SampleQuestions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SampleQuestions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SampleQuestions = from.SampleQuestions
	}
	if !from.SampleQuestions.IsNull() && !from.SampleQuestions.IsUnknown() {
		if toSampleQuestions, ok := to.GetSampleQuestions(ctx); ok {
			if fromSampleQuestions, ok := from.GetSampleQuestions(ctx); ok {
				for i := range toSampleQuestions {
					if i < len(fromSampleQuestions) {
						toSampleQuestions[i].SyncFieldsDuringRead(ctx, fromSampleQuestions[i])
					}
				}
				to.SetSampleQuestions(ctx, toSampleQuestions)
			}
		}
	}
}

func (m GenieAgentConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sample_questions"] = attrs["sample_questions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sample_questions": reflect.TypeOf(GenieAgentSampleQuestion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentConfig
// only implements ToObjectValue() and Type().
func (m GenieAgentConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sample_questions": m.SampleQuestions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sample_questions": basetypes.ListType{
				ElemType: GenieAgentSampleQuestion{}.Type(ctx),
			},
		},
	}
}

// GetSampleQuestions returns the value of the SampleQuestions field in GenieAgentConfig as
// a slice of GenieAgentSampleQuestion values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentConfig) GetSampleQuestions(ctx context.Context) ([]GenieAgentSampleQuestion, bool) {
	if m.SampleQuestions.IsNull() || m.SampleQuestions.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentSampleQuestion
	d := m.SampleQuestions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSampleQuestions sets the value of the SampleQuestions field in GenieAgentConfig.
func (m *GenieAgentConfig) SetSampleQuestions(ctx context.Context, v []GenieAgentSampleQuestion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sample_questions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SampleQuestions = types.ListValueMust(t, vs)
}

// The data a Genie Agent can access.
type GenieAgentDataSources struct {
	// Unity Catalog MCP connections, sorted by `mcp_connection`.
	McpConnections types.List `tfsdk:"mcp_connections"`
	// Unity Catalog tables, views, and metric views, in the order they were
	// added.
	Tables types.List `tfsdk:"tables"`
	// Unity Catalog volumes, in the order they were added.
	Volumes types.List `tfsdk:"volumes"`
}

func (to *GenieAgentDataSources) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentDataSources) {
	if !from.McpConnections.IsNull() && !from.McpConnections.IsUnknown() && to.McpConnections.IsNull() && len(from.McpConnections.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for McpConnections, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.McpConnections = from.McpConnections
	}
	if !from.McpConnections.IsNull() && !from.McpConnections.IsUnknown() {
		if toMcpConnections, ok := to.GetMcpConnections(ctx); ok {
			if fromMcpConnections, ok := from.GetMcpConnections(ctx); ok {
				// Recursively sync the fields of each McpConnections element by position.
				for i := range toMcpConnections {
					if i < len(fromMcpConnections) {
						toMcpConnections[i].SyncFieldsDuringCreateOrUpdate(ctx, fromMcpConnections[i])
					}
				}
				to.SetMcpConnections(ctx, toMcpConnections)
			}
		}
	}
	if !from.Tables.IsNull() && !from.Tables.IsUnknown() && to.Tables.IsNull() && len(from.Tables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tables = from.Tables
	}
	if !from.Tables.IsNull() && !from.Tables.IsUnknown() {
		if toTables, ok := to.GetTables(ctx); ok {
			if fromTables, ok := from.GetTables(ctx); ok {
				// Recursively sync the fields of each Tables element by position.
				for i := range toTables {
					if i < len(fromTables) {
						toTables[i].SyncFieldsDuringCreateOrUpdate(ctx, fromTables[i])
					}
				}
				to.SetTables(ctx, toTables)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() && to.Volumes.IsNull() && len(from.Volumes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Volumes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Volumes = from.Volumes
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				// Recursively sync the fields of each Volumes element by position.
				for i := range toVolumes {
					if i < len(fromVolumes) {
						toVolumes[i].SyncFieldsDuringCreateOrUpdate(ctx, fromVolumes[i])
					}
				}
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
}

func (to *GenieAgentDataSources) SyncFieldsDuringRead(ctx context.Context, from GenieAgentDataSources) {
	if !from.McpConnections.IsNull() && !from.McpConnections.IsUnknown() && to.McpConnections.IsNull() && len(from.McpConnections.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for McpConnections, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.McpConnections = from.McpConnections
	}
	if !from.McpConnections.IsNull() && !from.McpConnections.IsUnknown() {
		if toMcpConnections, ok := to.GetMcpConnections(ctx); ok {
			if fromMcpConnections, ok := from.GetMcpConnections(ctx); ok {
				for i := range toMcpConnections {
					if i < len(fromMcpConnections) {
						toMcpConnections[i].SyncFieldsDuringRead(ctx, fromMcpConnections[i])
					}
				}
				to.SetMcpConnections(ctx, toMcpConnections)
			}
		}
	}
	if !from.Tables.IsNull() && !from.Tables.IsUnknown() && to.Tables.IsNull() && len(from.Tables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tables = from.Tables
	}
	if !from.Tables.IsNull() && !from.Tables.IsUnknown() {
		if toTables, ok := to.GetTables(ctx); ok {
			if fromTables, ok := from.GetTables(ctx); ok {
				for i := range toTables {
					if i < len(fromTables) {
						toTables[i].SyncFieldsDuringRead(ctx, fromTables[i])
					}
				}
				to.SetTables(ctx, toTables)
			}
		}
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() && to.Volumes.IsNull() && len(from.Volumes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Volumes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Volumes = from.Volumes
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() {
		if toVolumes, ok := to.GetVolumes(ctx); ok {
			if fromVolumes, ok := from.GetVolumes(ctx); ok {
				for i := range toVolumes {
					if i < len(fromVolumes) {
						toVolumes[i].SyncFieldsDuringRead(ctx, fromVolumes[i])
					}
				}
				to.SetVolumes(ctx, toVolumes)
			}
		}
	}
}

func (m GenieAgentDataSources) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["mcp_connections"] = attrs["mcp_connections"].SetOptional()
	attrs["tables"] = attrs["tables"].SetOptional()
	attrs["volumes"] = attrs["volumes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentDataSources.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentDataSources) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"mcp_connections": reflect.TypeOf(GenieAgentMcpConnection{}),
		"tables":          reflect.TypeOf(GenieAgentTable{}),
		"volumes":         reflect.TypeOf(GenieAgentVolume{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentDataSources
// only implements ToObjectValue() and Type().
func (m GenieAgentDataSources) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"mcp_connections": m.McpConnections,
			"tables":          m.Tables,
			"volumes":         m.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentDataSources) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"mcp_connections": basetypes.ListType{
				ElemType: GenieAgentMcpConnection{}.Type(ctx),
			},
			"tables": basetypes.ListType{
				ElemType: GenieAgentTable{}.Type(ctx),
			},
			"volumes": basetypes.ListType{
				ElemType: GenieAgentVolume{}.Type(ctx),
			},
		},
	}
}

// GetMcpConnections returns the value of the McpConnections field in GenieAgentDataSources as
// a slice of GenieAgentMcpConnection values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentDataSources) GetMcpConnections(ctx context.Context) ([]GenieAgentMcpConnection, bool) {
	if m.McpConnections.IsNull() || m.McpConnections.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentMcpConnection
	d := m.McpConnections.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMcpConnections sets the value of the McpConnections field in GenieAgentDataSources.
func (m *GenieAgentDataSources) SetMcpConnections(ctx context.Context, v []GenieAgentMcpConnection) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["mcp_connections"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.McpConnections = types.ListValueMust(t, vs)
}

// GetTables returns the value of the Tables field in GenieAgentDataSources as
// a slice of GenieAgentTable values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentDataSources) GetTables(ctx context.Context) ([]GenieAgentTable, bool) {
	if m.Tables.IsNull() || m.Tables.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentTable
	d := m.Tables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTables sets the value of the Tables field in GenieAgentDataSources.
func (m *GenieAgentDataSources) SetTables(ctx context.Context, v []GenieAgentTable) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tables = types.ListValueMust(t, vs)
}

// GetVolumes returns the value of the Volumes field in GenieAgentDataSources as
// a slice of GenieAgentVolume values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentDataSources) GetVolumes(ctx context.Context) ([]GenieAgentVolume, bool) {
	if m.Volumes.IsNull() || m.Volumes.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentVolume
	d := m.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in GenieAgentDataSources.
func (m *GenieAgentDataSources) SetVolumes(ctx context.Context, v []GenieAgentVolume) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Volumes = types.ListValueMust(t, vs)
}

// The definition of a Genie Agent.
type GenieAgentDefinition struct {
	// Benchmarks for evaluating the quality of the Genie Agent's answers.
	Benchmarks types.Object `tfsdk:"benchmarks"`
	// General settings of the Genie Agent.
	Config types.Object `tfsdk:"config"`
	// The data the Genie Agent can access.
	DataSources types.Object `tfsdk:"data_sources"`
	// Instructions and SQL that guide how the Genie Agent answers questions.
	Instructions types.Object `tfsdk:"instructions"`
}

func (to *GenieAgentDefinition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentDefinition) {
	if !from.Benchmarks.IsNull() && !from.Benchmarks.IsUnknown() {
		if toBenchmarks, ok := to.GetBenchmarks(ctx); ok {
			if fromBenchmarks, ok := from.GetBenchmarks(ctx); ok {
				// Recursively sync the fields of Benchmarks
				toBenchmarks.SyncFieldsDuringCreateOrUpdate(ctx, fromBenchmarks)
				to.SetBenchmarks(ctx, toBenchmarks)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.DataSources.IsNull() && !from.DataSources.IsUnknown() {
		if toDataSources, ok := to.GetDataSources(ctx); ok {
			if fromDataSources, ok := from.GetDataSources(ctx); ok {
				// Recursively sync the fields of DataSources
				toDataSources.SyncFieldsDuringCreateOrUpdate(ctx, fromDataSources)
				to.SetDataSources(ctx, toDataSources)
			}
		}
	}
	if !from.Instructions.IsNull() && !from.Instructions.IsUnknown() {
		if toInstructions, ok := to.GetInstructions(ctx); ok {
			if fromInstructions, ok := from.GetInstructions(ctx); ok {
				// Recursively sync the fields of Instructions
				toInstructions.SyncFieldsDuringCreateOrUpdate(ctx, fromInstructions)
				to.SetInstructions(ctx, toInstructions)
			}
		}
	}
}

func (to *GenieAgentDefinition) SyncFieldsDuringRead(ctx context.Context, from GenieAgentDefinition) {
	if !from.Benchmarks.IsNull() && !from.Benchmarks.IsUnknown() {
		if toBenchmarks, ok := to.GetBenchmarks(ctx); ok {
			if fromBenchmarks, ok := from.GetBenchmarks(ctx); ok {
				toBenchmarks.SyncFieldsDuringRead(ctx, fromBenchmarks)
				to.SetBenchmarks(ctx, toBenchmarks)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.DataSources.IsNull() && !from.DataSources.IsUnknown() {
		if toDataSources, ok := to.GetDataSources(ctx); ok {
			if fromDataSources, ok := from.GetDataSources(ctx); ok {
				toDataSources.SyncFieldsDuringRead(ctx, fromDataSources)
				to.SetDataSources(ctx, toDataSources)
			}
		}
	}
	if !from.Instructions.IsNull() && !from.Instructions.IsUnknown() {
		if toInstructions, ok := to.GetInstructions(ctx); ok {
			if fromInstructions, ok := from.GetInstructions(ctx); ok {
				toInstructions.SyncFieldsDuringRead(ctx, fromInstructions)
				to.SetInstructions(ctx, toInstructions)
			}
		}
	}
}

func (m GenieAgentDefinition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["benchmarks"] = attrs["benchmarks"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["data_sources"] = attrs["data_sources"].SetOptional()
	attrs["instructions"] = attrs["instructions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentDefinition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentDefinition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"benchmarks":   reflect.TypeOf(GenieAgentBenchmarks{}),
		"config":       reflect.TypeOf(GenieAgentConfig{}),
		"data_sources": reflect.TypeOf(GenieAgentDataSources{}),
		"instructions": reflect.TypeOf(GenieAgentInstructions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentDefinition
// only implements ToObjectValue() and Type().
func (m GenieAgentDefinition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"benchmarks":   m.Benchmarks,
			"config":       m.Config,
			"data_sources": m.DataSources,
			"instructions": m.Instructions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentDefinition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"benchmarks":   GenieAgentBenchmarks{}.Type(ctx),
			"config":       GenieAgentConfig{}.Type(ctx),
			"data_sources": GenieAgentDataSources{}.Type(ctx),
			"instructions": GenieAgentInstructions{}.Type(ctx),
		},
	}
}

// GetBenchmarks returns the value of the Benchmarks field in GenieAgentDefinition as
// a GenieAgentBenchmarks value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentDefinition) GetBenchmarks(ctx context.Context) (GenieAgentBenchmarks, bool) {
	var e GenieAgentBenchmarks
	if m.Benchmarks.IsNull() || m.Benchmarks.IsUnknown() {
		return e, false
	}
	var v GenieAgentBenchmarks
	d := m.Benchmarks.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBenchmarks sets the value of the Benchmarks field in GenieAgentDefinition.
func (m *GenieAgentDefinition) SetBenchmarks(ctx context.Context, v GenieAgentBenchmarks) {
	vs := v.ToObjectValue(ctx)
	m.Benchmarks = vs
}

// GetConfig returns the value of the Config field in GenieAgentDefinition as
// a GenieAgentConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentDefinition) GetConfig(ctx context.Context) (GenieAgentConfig, bool) {
	var e GenieAgentConfig
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v GenieAgentConfig
	d := m.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in GenieAgentDefinition.
func (m *GenieAgentDefinition) SetConfig(ctx context.Context, v GenieAgentConfig) {
	vs := v.ToObjectValue(ctx)
	m.Config = vs
}

// GetDataSources returns the value of the DataSources field in GenieAgentDefinition as
// a GenieAgentDataSources value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentDefinition) GetDataSources(ctx context.Context) (GenieAgentDataSources, bool) {
	var e GenieAgentDataSources
	if m.DataSources.IsNull() || m.DataSources.IsUnknown() {
		return e, false
	}
	var v GenieAgentDataSources
	d := m.DataSources.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataSources sets the value of the DataSources field in GenieAgentDefinition.
func (m *GenieAgentDefinition) SetDataSources(ctx context.Context, v GenieAgentDataSources) {
	vs := v.ToObjectValue(ctx)
	m.DataSources = vs
}

// GetInstructions returns the value of the Instructions field in GenieAgentDefinition as
// a GenieAgentInstructions value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentDefinition) GetInstructions(ctx context.Context) (GenieAgentInstructions, bool) {
	var e GenieAgentInstructions
	if m.Instructions.IsNull() || m.Instructions.IsUnknown() {
		return e, false
	}
	var v GenieAgentInstructions
	d := m.Instructions.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstructions sets the value of the Instructions field in GenieAgentDefinition.
func (m *GenieAgentDefinition) SetInstructions(ctx context.Context, v GenieAgentInstructions) {
	vs := v.ToObjectValue(ctx)
	m.Instructions = vs
}

// An example question paired with the SQL that answers it.
type GenieAgentExampleQuestionSql struct {
	// Identifier of the example: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// Parameters used in the SQL, in the order given.
	Parameters types.List `tfsdk:"parameters"`
	// The question.
	Question types.String `tfsdk:"question"`
	// The SQL that answers the question.
	Sql types.String `tfsdk:"sql"`
	// Guidance on when to use this example.
	UsageGuidance types.String `tfsdk:"usage_guidance"`
}

func (to *GenieAgentExampleQuestionSql) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentExampleQuestionSql) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() {
		if toParameters, ok := to.GetParameters(ctx); ok {
			if fromParameters, ok := from.GetParameters(ctx); ok {
				// Recursively sync the fields of each Parameters element by position.
				for i := range toParameters {
					if i < len(fromParameters) {
						toParameters[i].SyncFieldsDuringCreateOrUpdate(ctx, fromParameters[i])
					}
				}
				to.SetParameters(ctx, toParameters)
			}
		}
	}
}

func (to *GenieAgentExampleQuestionSql) SyncFieldsDuringRead(ctx context.Context, from GenieAgentExampleQuestionSql) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() {
		if toParameters, ok := to.GetParameters(ctx); ok {
			if fromParameters, ok := from.GetParameters(ctx); ok {
				for i := range toParameters {
					if i < len(fromParameters) {
						toParameters[i].SyncFieldsDuringRead(ctx, fromParameters[i])
					}
				}
				to.SetParameters(ctx, toParameters)
			}
		}
	}
}

func (m GenieAgentExampleQuestionSql) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["question"] = attrs["question"].SetRequired()
	attrs["sql"] = attrs["sql"].SetRequired()
	attrs["usage_guidance"] = attrs["usage_guidance"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentExampleQuestionSql.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentExampleQuestionSql) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(GenieAgentSqlParameter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentExampleQuestionSql
// only implements ToObjectValue() and Type().
func (m GenieAgentExampleQuestionSql) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":             m.Id,
			"parameters":     m.Parameters,
			"question":       m.Question,
			"sql":            m.Sql,
			"usage_guidance": m.UsageGuidance,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentExampleQuestionSql) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: GenieAgentSqlParameter{}.Type(ctx),
			},
			"question":       types.StringType,
			"sql":            types.StringType,
			"usage_guidance": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in GenieAgentExampleQuestionSql as
// a slice of GenieAgentSqlParameter values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentExampleQuestionSql) GetParameters(ctx context.Context) ([]GenieAgentSqlParameter, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentSqlParameter
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in GenieAgentExampleQuestionSql.
func (m *GenieAgentExampleQuestionSql) SetParameters(ctx context.Context, v []GenieAgentSqlParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

// Instructions and SQL that guide how a Genie Agent answers questions.
type GenieAgentInstructions struct {
	// Example questions paired with the SQL that answers them.
	ExampleQuestionSqls types.List `tfsdk:"example_question_sqls"`
	// Predefined joins between tables or metric views.
	JoinSpecs types.List `tfsdk:"join_specs"`
	// Unity Catalog SQL functions the generated SQL may call.
	SqlFunctions types.List `tfsdk:"sql_functions"`
	// Reusable SQL filters, expressions, and measures.
	SqlSnippets types.Object `tfsdk:"sql_snippets"`
	// General text instructions. At most one text instruction is supported.
	TextInstructions types.List `tfsdk:"text_instructions"`
}

func (to *GenieAgentInstructions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentInstructions) {
	if !from.ExampleQuestionSqls.IsNull() && !from.ExampleQuestionSqls.IsUnknown() && to.ExampleQuestionSqls.IsNull() && len(from.ExampleQuestionSqls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExampleQuestionSqls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExampleQuestionSqls = from.ExampleQuestionSqls
	}
	if !from.ExampleQuestionSqls.IsNull() && !from.ExampleQuestionSqls.IsUnknown() {
		if toExampleQuestionSqls, ok := to.GetExampleQuestionSqls(ctx); ok {
			if fromExampleQuestionSqls, ok := from.GetExampleQuestionSqls(ctx); ok {
				// Recursively sync the fields of each ExampleQuestionSqls element by position.
				for i := range toExampleQuestionSqls {
					if i < len(fromExampleQuestionSqls) {
						toExampleQuestionSqls[i].SyncFieldsDuringCreateOrUpdate(ctx, fromExampleQuestionSqls[i])
					}
				}
				to.SetExampleQuestionSqls(ctx, toExampleQuestionSqls)
			}
		}
	}
	if !from.JoinSpecs.IsNull() && !from.JoinSpecs.IsUnknown() && to.JoinSpecs.IsNull() && len(from.JoinSpecs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JoinSpecs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JoinSpecs = from.JoinSpecs
	}
	if !from.JoinSpecs.IsNull() && !from.JoinSpecs.IsUnknown() {
		if toJoinSpecs, ok := to.GetJoinSpecs(ctx); ok {
			if fromJoinSpecs, ok := from.GetJoinSpecs(ctx); ok {
				// Recursively sync the fields of each JoinSpecs element by position.
				for i := range toJoinSpecs {
					if i < len(fromJoinSpecs) {
						toJoinSpecs[i].SyncFieldsDuringCreateOrUpdate(ctx, fromJoinSpecs[i])
					}
				}
				to.SetJoinSpecs(ctx, toJoinSpecs)
			}
		}
	}
	if !from.SqlFunctions.IsNull() && !from.SqlFunctions.IsUnknown() && to.SqlFunctions.IsNull() && len(from.SqlFunctions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SqlFunctions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SqlFunctions = from.SqlFunctions
	}
	if !from.SqlFunctions.IsNull() && !from.SqlFunctions.IsUnknown() {
		if toSqlFunctions, ok := to.GetSqlFunctions(ctx); ok {
			if fromSqlFunctions, ok := from.GetSqlFunctions(ctx); ok {
				// Recursively sync the fields of each SqlFunctions element by position.
				for i := range toSqlFunctions {
					if i < len(fromSqlFunctions) {
						toSqlFunctions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSqlFunctions[i])
					}
				}
				to.SetSqlFunctions(ctx, toSqlFunctions)
			}
		}
	}
	if !from.SqlSnippets.IsNull() && !from.SqlSnippets.IsUnknown() {
		if toSqlSnippets, ok := to.GetSqlSnippets(ctx); ok {
			if fromSqlSnippets, ok := from.GetSqlSnippets(ctx); ok {
				// Recursively sync the fields of SqlSnippets
				toSqlSnippets.SyncFieldsDuringCreateOrUpdate(ctx, fromSqlSnippets)
				to.SetSqlSnippets(ctx, toSqlSnippets)
			}
		}
	}
	if !from.TextInstructions.IsNull() && !from.TextInstructions.IsUnknown() && to.TextInstructions.IsNull() && len(from.TextInstructions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TextInstructions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TextInstructions = from.TextInstructions
	}
	if !from.TextInstructions.IsNull() && !from.TextInstructions.IsUnknown() {
		if toTextInstructions, ok := to.GetTextInstructions(ctx); ok {
			if fromTextInstructions, ok := from.GetTextInstructions(ctx); ok {
				// Recursively sync the fields of each TextInstructions element by position.
				for i := range toTextInstructions {
					if i < len(fromTextInstructions) {
						toTextInstructions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromTextInstructions[i])
					}
				}
				to.SetTextInstructions(ctx, toTextInstructions)
			}
		}
	}
}

func (to *GenieAgentInstructions) SyncFieldsDuringRead(ctx context.Context, from GenieAgentInstructions) {
	if !from.ExampleQuestionSqls.IsNull() && !from.ExampleQuestionSqls.IsUnknown() && to.ExampleQuestionSqls.IsNull() && len(from.ExampleQuestionSqls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExampleQuestionSqls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExampleQuestionSqls = from.ExampleQuestionSqls
	}
	if !from.ExampleQuestionSqls.IsNull() && !from.ExampleQuestionSqls.IsUnknown() {
		if toExampleQuestionSqls, ok := to.GetExampleQuestionSqls(ctx); ok {
			if fromExampleQuestionSqls, ok := from.GetExampleQuestionSqls(ctx); ok {
				for i := range toExampleQuestionSqls {
					if i < len(fromExampleQuestionSqls) {
						toExampleQuestionSqls[i].SyncFieldsDuringRead(ctx, fromExampleQuestionSqls[i])
					}
				}
				to.SetExampleQuestionSqls(ctx, toExampleQuestionSqls)
			}
		}
	}
	if !from.JoinSpecs.IsNull() && !from.JoinSpecs.IsUnknown() && to.JoinSpecs.IsNull() && len(from.JoinSpecs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JoinSpecs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JoinSpecs = from.JoinSpecs
	}
	if !from.JoinSpecs.IsNull() && !from.JoinSpecs.IsUnknown() {
		if toJoinSpecs, ok := to.GetJoinSpecs(ctx); ok {
			if fromJoinSpecs, ok := from.GetJoinSpecs(ctx); ok {
				for i := range toJoinSpecs {
					if i < len(fromJoinSpecs) {
						toJoinSpecs[i].SyncFieldsDuringRead(ctx, fromJoinSpecs[i])
					}
				}
				to.SetJoinSpecs(ctx, toJoinSpecs)
			}
		}
	}
	if !from.SqlFunctions.IsNull() && !from.SqlFunctions.IsUnknown() && to.SqlFunctions.IsNull() && len(from.SqlFunctions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SqlFunctions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SqlFunctions = from.SqlFunctions
	}
	if !from.SqlFunctions.IsNull() && !from.SqlFunctions.IsUnknown() {
		if toSqlFunctions, ok := to.GetSqlFunctions(ctx); ok {
			if fromSqlFunctions, ok := from.GetSqlFunctions(ctx); ok {
				for i := range toSqlFunctions {
					if i < len(fromSqlFunctions) {
						toSqlFunctions[i].SyncFieldsDuringRead(ctx, fromSqlFunctions[i])
					}
				}
				to.SetSqlFunctions(ctx, toSqlFunctions)
			}
		}
	}
	if !from.SqlSnippets.IsNull() && !from.SqlSnippets.IsUnknown() {
		if toSqlSnippets, ok := to.GetSqlSnippets(ctx); ok {
			if fromSqlSnippets, ok := from.GetSqlSnippets(ctx); ok {
				toSqlSnippets.SyncFieldsDuringRead(ctx, fromSqlSnippets)
				to.SetSqlSnippets(ctx, toSqlSnippets)
			}
		}
	}
	if !from.TextInstructions.IsNull() && !from.TextInstructions.IsUnknown() && to.TextInstructions.IsNull() && len(from.TextInstructions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TextInstructions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TextInstructions = from.TextInstructions
	}
	if !from.TextInstructions.IsNull() && !from.TextInstructions.IsUnknown() {
		if toTextInstructions, ok := to.GetTextInstructions(ctx); ok {
			if fromTextInstructions, ok := from.GetTextInstructions(ctx); ok {
				for i := range toTextInstructions {
					if i < len(fromTextInstructions) {
						toTextInstructions[i].SyncFieldsDuringRead(ctx, fromTextInstructions[i])
					}
				}
				to.SetTextInstructions(ctx, toTextInstructions)
			}
		}
	}
}

func (m GenieAgentInstructions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["example_question_sqls"] = attrs["example_question_sqls"].SetOptional()
	attrs["join_specs"] = attrs["join_specs"].SetOptional()
	attrs["sql_functions"] = attrs["sql_functions"].SetOptional()
	attrs["sql_snippets"] = attrs["sql_snippets"].SetOptional()
	attrs["text_instructions"] = attrs["text_instructions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentInstructions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentInstructions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"example_question_sqls": reflect.TypeOf(GenieAgentExampleQuestionSql{}),
		"join_specs":            reflect.TypeOf(GenieAgentJoinSpec{}),
		"sql_functions":         reflect.TypeOf(GenieAgentSqlFunction{}),
		"sql_snippets":          reflect.TypeOf(GenieAgentSqlSnippets{}),
		"text_instructions":     reflect.TypeOf(GenieAgentTextInstruction{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentInstructions
// only implements ToObjectValue() and Type().
func (m GenieAgentInstructions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"example_question_sqls": m.ExampleQuestionSqls,
			"join_specs":            m.JoinSpecs,
			"sql_functions":         m.SqlFunctions,
			"sql_snippets":          m.SqlSnippets,
			"text_instructions":     m.TextInstructions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentInstructions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"example_question_sqls": basetypes.ListType{
				ElemType: GenieAgentExampleQuestionSql{}.Type(ctx),
			},
			"join_specs": basetypes.ListType{
				ElemType: GenieAgentJoinSpec{}.Type(ctx),
			},
			"sql_functions": basetypes.ListType{
				ElemType: GenieAgentSqlFunction{}.Type(ctx),
			},
			"sql_snippets": GenieAgentSqlSnippets{}.Type(ctx),
			"text_instructions": basetypes.ListType{
				ElemType: GenieAgentTextInstruction{}.Type(ctx),
			},
		},
	}
}

// GetExampleQuestionSqls returns the value of the ExampleQuestionSqls field in GenieAgentInstructions as
// a slice of GenieAgentExampleQuestionSql values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentInstructions) GetExampleQuestionSqls(ctx context.Context) ([]GenieAgentExampleQuestionSql, bool) {
	if m.ExampleQuestionSqls.IsNull() || m.ExampleQuestionSqls.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentExampleQuestionSql
	d := m.ExampleQuestionSqls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExampleQuestionSqls sets the value of the ExampleQuestionSqls field in GenieAgentInstructions.
func (m *GenieAgentInstructions) SetExampleQuestionSqls(ctx context.Context, v []GenieAgentExampleQuestionSql) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["example_question_sqls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExampleQuestionSqls = types.ListValueMust(t, vs)
}

// GetJoinSpecs returns the value of the JoinSpecs field in GenieAgentInstructions as
// a slice of GenieAgentJoinSpec values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentInstructions) GetJoinSpecs(ctx context.Context) ([]GenieAgentJoinSpec, bool) {
	if m.JoinSpecs.IsNull() || m.JoinSpecs.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentJoinSpec
	d := m.JoinSpecs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJoinSpecs sets the value of the JoinSpecs field in GenieAgentInstructions.
func (m *GenieAgentInstructions) SetJoinSpecs(ctx context.Context, v []GenieAgentJoinSpec) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["join_specs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.JoinSpecs = types.ListValueMust(t, vs)
}

// GetSqlFunctions returns the value of the SqlFunctions field in GenieAgentInstructions as
// a slice of GenieAgentSqlFunction values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentInstructions) GetSqlFunctions(ctx context.Context) ([]GenieAgentSqlFunction, bool) {
	if m.SqlFunctions.IsNull() || m.SqlFunctions.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentSqlFunction
	d := m.SqlFunctions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlFunctions sets the value of the SqlFunctions field in GenieAgentInstructions.
func (m *GenieAgentInstructions) SetSqlFunctions(ctx context.Context, v []GenieAgentSqlFunction) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_functions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SqlFunctions = types.ListValueMust(t, vs)
}

// GetSqlSnippets returns the value of the SqlSnippets field in GenieAgentInstructions as
// a GenieAgentSqlSnippets value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentInstructions) GetSqlSnippets(ctx context.Context) (GenieAgentSqlSnippets, bool) {
	var e GenieAgentSqlSnippets
	if m.SqlSnippets.IsNull() || m.SqlSnippets.IsUnknown() {
		return e, false
	}
	var v GenieAgentSqlSnippets
	d := m.SqlSnippets.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlSnippets sets the value of the SqlSnippets field in GenieAgentInstructions.
func (m *GenieAgentInstructions) SetSqlSnippets(ctx context.Context, v GenieAgentSqlSnippets) {
	vs := v.ToObjectValue(ctx)
	m.SqlSnippets = vs
}

// GetTextInstructions returns the value of the TextInstructions field in GenieAgentInstructions as
// a slice of GenieAgentTextInstruction values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentInstructions) GetTextInstructions(ctx context.Context) ([]GenieAgentTextInstruction, bool) {
	if m.TextInstructions.IsNull() || m.TextInstructions.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentTextInstruction
	d := m.TextInstructions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTextInstructions sets the value of the TextInstructions field in GenieAgentInstructions.
func (m *GenieAgentInstructions) SetTextInstructions(ctx context.Context, v []GenieAgentTextInstruction) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["text_instructions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TextInstructions = types.ListValueMust(t, vs)
}

// One side of a join.
type GenieAgentJoinSource struct {
	// The alias the join condition uses for this side.
	Alias types.String `tfsdk:"alias"`
	// The full three-level name of the table or metric view.
	Identifier types.String `tfsdk:"identifier"`
}

func (to *GenieAgentJoinSource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentJoinSource) {
}

func (to *GenieAgentJoinSource) SyncFieldsDuringRead(ctx context.Context, from GenieAgentJoinSource) {
}

func (m GenieAgentJoinSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alias"] = attrs["alias"].SetOptional()
	attrs["identifier"] = attrs["identifier"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentJoinSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentJoinSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentJoinSource
// only implements ToObjectValue() and Type().
func (m GenieAgentJoinSource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":      m.Alias,
			"identifier": m.Identifier,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentJoinSource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":      types.StringType,
			"identifier": types.StringType,
		},
	}
}

// A predefined join between two tables or metric views.
type GenieAgentJoinSpec struct {
	// Description of the join's purpose.
	Comment types.String `tfsdk:"comment"`
	// Identifier of the join: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// Guidance on when to use the join.
	Instruction types.String `tfsdk:"instruction"`
	// The left side of the join.
	Left types.Object `tfsdk:"left"`
	// How the rows of the two sides relate.
	RelationshipType types.String `tfsdk:"relationship_type"`
	// The right side of the join.
	Right types.Object `tfsdk:"right"`
	// The join condition, for example `orders.customer_id = customers.id`.
	Sql types.String `tfsdk:"sql"`
}

func (to *GenieAgentJoinSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentJoinSpec) {
	if !from.Left.IsNull() && !from.Left.IsUnknown() {
		if toLeft, ok := to.GetLeft(ctx); ok {
			if fromLeft, ok := from.GetLeft(ctx); ok {
				// Recursively sync the fields of Left
				toLeft.SyncFieldsDuringCreateOrUpdate(ctx, fromLeft)
				to.SetLeft(ctx, toLeft)
			}
		}
	}
	if !from.Right.IsNull() && !from.Right.IsUnknown() {
		if toRight, ok := to.GetRight(ctx); ok {
			if fromRight, ok := from.GetRight(ctx); ok {
				// Recursively sync the fields of Right
				toRight.SyncFieldsDuringCreateOrUpdate(ctx, fromRight)
				to.SetRight(ctx, toRight)
			}
		}
	}
}

func (to *GenieAgentJoinSpec) SyncFieldsDuringRead(ctx context.Context, from GenieAgentJoinSpec) {
	if !from.Left.IsNull() && !from.Left.IsUnknown() {
		if toLeft, ok := to.GetLeft(ctx); ok {
			if fromLeft, ok := from.GetLeft(ctx); ok {
				toLeft.SyncFieldsDuringRead(ctx, fromLeft)
				to.SetLeft(ctx, toLeft)
			}
		}
	}
	if !from.Right.IsNull() && !from.Right.IsUnknown() {
		if toRight, ok := to.GetRight(ctx); ok {
			if fromRight, ok := from.GetRight(ctx); ok {
				toRight.SyncFieldsDuringRead(ctx, fromRight)
				to.SetRight(ctx, toRight)
			}
		}
	}
}

func (m GenieAgentJoinSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["instruction"] = attrs["instruction"].SetOptional()
	attrs["left"] = attrs["left"].SetRequired()
	attrs["relationship_type"] = attrs["relationship_type"].SetOptional()
	attrs["right"] = attrs["right"].SetRequired()
	attrs["sql"] = attrs["sql"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentJoinSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentJoinSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"left":  reflect.TypeOf(GenieAgentJoinSource{}),
		"right": reflect.TypeOf(GenieAgentJoinSource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentJoinSpec
// only implements ToObjectValue() and Type().
func (m GenieAgentJoinSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":           m.Comment,
			"id":                m.Id,
			"instruction":       m.Instruction,
			"left":              m.Left,
			"relationship_type": m.RelationshipType,
			"right":             m.Right,
			"sql":               m.Sql,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentJoinSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":           types.StringType,
			"id":                types.StringType,
			"instruction":       types.StringType,
			"left":              GenieAgentJoinSource{}.Type(ctx),
			"relationship_type": types.StringType,
			"right":             GenieAgentJoinSource{}.Type(ctx),
			"sql":               types.StringType,
		},
	}
}

// GetLeft returns the value of the Left field in GenieAgentJoinSpec as
// a GenieAgentJoinSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentJoinSpec) GetLeft(ctx context.Context) (GenieAgentJoinSource, bool) {
	var e GenieAgentJoinSource
	if m.Left.IsNull() || m.Left.IsUnknown() {
		return e, false
	}
	var v GenieAgentJoinSource
	d := m.Left.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLeft sets the value of the Left field in GenieAgentJoinSpec.
func (m *GenieAgentJoinSpec) SetLeft(ctx context.Context, v GenieAgentJoinSource) {
	vs := v.ToObjectValue(ctx)
	m.Left = vs
}

// GetRight returns the value of the Right field in GenieAgentJoinSpec as
// a GenieAgentJoinSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentJoinSpec) GetRight(ctx context.Context) (GenieAgentJoinSource, bool) {
	var e GenieAgentJoinSource
	if m.Right.IsNull() || m.Right.IsUnknown() {
		return e, false
	}
	var v GenieAgentJoinSource
	d := m.Right.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRight sets the value of the Right field in GenieAgentJoinSpec.
func (m *GenieAgentJoinSpec) SetRight(ctx context.Context, v GenieAgentJoinSource) {
	vs := v.ToObjectValue(ctx)
	m.Right = vs
}

// A Unity Catalog MCP connection the Genie Agent can use.
type GenieAgentMcpConnection struct {
	// The name of the Unity Catalog connection, for example
	// `system.ai.google_calendar`.
	McpConnection types.String `tfsdk:"mcp_connection"`
}

func (to *GenieAgentMcpConnection) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentMcpConnection) {
}

func (to *GenieAgentMcpConnection) SyncFieldsDuringRead(ctx context.Context, from GenieAgentMcpConnection) {
}

func (m GenieAgentMcpConnection) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["mcp_connection"] = attrs["mcp_connection"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentMcpConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentMcpConnection) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentMcpConnection
// only implements ToObjectValue() and Type().
func (m GenieAgentMcpConnection) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"mcp_connection": m.McpConnection,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentMcpConnection) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"mcp_connection": types.StringType,
		},
	}
}

// A sample question shown to users of the Genie Agent.
type GenieAgentSampleQuestion struct {
	// Identifier of the sample question: 32 lowercase hexadecimal characters,
	// unique within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// The text of the question.
	Question types.String `tfsdk:"question"`
}

func (to *GenieAgentSampleQuestion) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSampleQuestion) {
}

func (to *GenieAgentSampleQuestion) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSampleQuestion) {
}

func (m GenieAgentSampleQuestion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["question"] = attrs["question"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSampleQuestion.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSampleQuestion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSampleQuestion
// only implements ToObjectValue() and Type().
func (m GenieAgentSampleQuestion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":       m.Id,
			"question": m.Question,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSampleQuestion) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":       types.StringType,
			"question": types.StringType,
		},
	}
}

// A reusable non-aggregated SQL expression (dimension).
type GenieAgentSqlExpression struct {
	// The alias the Genie Agent uses for the expression in generated SQL.
	Alias types.String `tfsdk:"alias"`
	// Display name of the expression.
	DisplayName types.String `tfsdk:"display_name"`
	// Identifier of the expression: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// Guidance on when to use the expression.
	Instruction types.String `tfsdk:"instruction"`
	// The SQL of the expression.
	Sql types.String `tfsdk:"sql"`
	// Synonyms for the expression.
	Synonyms types.List `tfsdk:"synonyms"`
}

func (to *GenieAgentSqlExpression) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlExpression) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (to *GenieAgentSqlExpression) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlExpression) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (m GenieAgentSqlExpression) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alias"] = attrs["alias"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["instruction"] = attrs["instruction"].SetOptional()
	attrs["sql"] = attrs["sql"].SetRequired()
	attrs["synonyms"] = attrs["synonyms"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlExpression.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlExpression) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synonyms": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlExpression
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlExpression) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":        m.Alias,
			"display_name": m.DisplayName,
			"id":           m.Id,
			"instruction":  m.Instruction,
			"sql":          m.Sql,
			"synonyms":     m.Synonyms,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlExpression) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":        types.StringType,
			"display_name": types.StringType,
			"id":           types.StringType,
			"instruction":  types.StringType,
			"sql":          types.StringType,
			"synonyms": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetSynonyms returns the value of the Synonyms field in GenieAgentSqlExpression as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlExpression) GetSynonyms(ctx context.Context) ([]types.String, bool) {
	if m.Synonyms.IsNull() || m.Synonyms.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Synonyms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSynonyms sets the value of the Synonyms field in GenieAgentSqlExpression.
func (m *GenieAgentSqlExpression) SetSynonyms(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["synonyms"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Synonyms = types.ListValueMust(t, vs)
}

// A reusable SQL filter.
type GenieAgentSqlFilter struct {
	// Display name of the filter.
	DisplayName types.String `tfsdk:"display_name"`
	// Identifier of the filter: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// Guidance on when to use the filter.
	Instruction types.String `tfsdk:"instruction"`
	// The SQL of the filter.
	Sql types.String `tfsdk:"sql"`
	// Synonyms for the filter.
	Synonyms types.List `tfsdk:"synonyms"`
}

func (to *GenieAgentSqlFilter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlFilter) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (to *GenieAgentSqlFilter) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlFilter) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (m GenieAgentSqlFilter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["instruction"] = attrs["instruction"].SetOptional()
	attrs["sql"] = attrs["sql"].SetRequired()
	attrs["synonyms"] = attrs["synonyms"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlFilter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synonyms": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlFilter
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlFilter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name": m.DisplayName,
			"id":           m.Id,
			"instruction":  m.Instruction,
			"sql":          m.Sql,
			"synonyms":     m.Synonyms,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlFilter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name": types.StringType,
			"id":           types.StringType,
			"instruction":  types.StringType,
			"sql":          types.StringType,
			"synonyms": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetSynonyms returns the value of the Synonyms field in GenieAgentSqlFilter as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlFilter) GetSynonyms(ctx context.Context) ([]types.String, bool) {
	if m.Synonyms.IsNull() || m.Synonyms.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Synonyms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSynonyms sets the value of the Synonyms field in GenieAgentSqlFilter.
func (m *GenieAgentSqlFilter) SetSynonyms(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["synonyms"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Synonyms = types.ListValueMust(t, vs)
}

// A Unity Catalog SQL function the generated SQL may call.
type GenieAgentSqlFunction struct {
	// Identifier of the entry: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// The full three-level name of the function, for example
	// `catalog.schema.function`.
	Identifier types.String `tfsdk:"identifier"`
}

func (to *GenieAgentSqlFunction) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlFunction) {
}

func (to *GenieAgentSqlFunction) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlFunction) {
}

func (m GenieAgentSqlFunction) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["identifier"] = attrs["identifier"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlFunction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlFunction) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlFunction
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlFunction) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         m.Id,
			"identifier": m.Identifier,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlFunction) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"identifier": types.StringType,
		},
	}
}

// A reusable aggregated SQL measure.
type GenieAgentSqlMeasure struct {
	// The alias the Genie Agent uses for the measure in generated SQL.
	Alias types.String `tfsdk:"alias"`
	// Display name of the measure.
	DisplayName types.String `tfsdk:"display_name"`
	// Identifier of the measure: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id types.String `tfsdk:"id"`
	// Guidance on when to use the measure.
	Instruction types.String `tfsdk:"instruction"`
	// The SQL of the measure.
	Sql types.String `tfsdk:"sql"`
	// Synonyms for the measure.
	Synonyms types.List `tfsdk:"synonyms"`
}

func (to *GenieAgentSqlMeasure) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlMeasure) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (to *GenieAgentSqlMeasure) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlMeasure) {
	if !from.Synonyms.IsNull() && !from.Synonyms.IsUnknown() && to.Synonyms.IsNull() && len(from.Synonyms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Synonyms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Synonyms = from.Synonyms
	}
}

func (m GenieAgentSqlMeasure) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alias"] = attrs["alias"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["instruction"] = attrs["instruction"].SetOptional()
	attrs["sql"] = attrs["sql"].SetRequired()
	attrs["synonyms"] = attrs["synonyms"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlMeasure.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlMeasure) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synonyms": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlMeasure
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlMeasure) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":        m.Alias,
			"display_name": m.DisplayName,
			"id":           m.Id,
			"instruction":  m.Instruction,
			"sql":          m.Sql,
			"synonyms":     m.Synonyms,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlMeasure) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":        types.StringType,
			"display_name": types.StringType,
			"id":           types.StringType,
			"instruction":  types.StringType,
			"sql":          types.StringType,
			"synonyms": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetSynonyms returns the value of the Synonyms field in GenieAgentSqlMeasure as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlMeasure) GetSynonyms(ctx context.Context) ([]types.String, bool) {
	if m.Synonyms.IsNull() || m.Synonyms.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Synonyms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSynonyms sets the value of the Synonyms field in GenieAgentSqlMeasure.
func (m *GenieAgentSqlMeasure) SetSynonyms(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["synonyms"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Synonyms = types.ListValueMust(t, vs)
}

// A parameter of a parameterized SQL example.
type GenieAgentSqlParameter struct {
	// The parameter's default value.
	DefaultValue types.Object `tfsdk:"default_value"`
	// Description of the parameter.
	Description types.String `tfsdk:"description"`
	// The name of the parameter in the SQL.
	Name types.String `tfsdk:"name"`
	// The parameter's data type, for example `STRING` or `INTEGER`.
	TypeHint types.String `tfsdk:"type_hint"`
}

func (to *GenieAgentSqlParameter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlParameter) {
	if !from.DefaultValue.IsNull() && !from.DefaultValue.IsUnknown() {
		if toDefaultValue, ok := to.GetDefaultValue(ctx); ok {
			if fromDefaultValue, ok := from.GetDefaultValue(ctx); ok {
				// Recursively sync the fields of DefaultValue
				toDefaultValue.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultValue)
				to.SetDefaultValue(ctx, toDefaultValue)
			}
		}
	}
}

func (to *GenieAgentSqlParameter) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlParameter) {
	if !from.DefaultValue.IsNull() && !from.DefaultValue.IsUnknown() {
		if toDefaultValue, ok := to.GetDefaultValue(ctx); ok {
			if fromDefaultValue, ok := from.GetDefaultValue(ctx); ok {
				toDefaultValue.SyncFieldsDuringRead(ctx, fromDefaultValue)
				to.SetDefaultValue(ctx, toDefaultValue)
			}
		}
	}
}

func (m GenieAgentSqlParameter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_value"] = attrs["default_value"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["type_hint"] = attrs["type_hint"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlParameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_value": reflect.TypeOf(GenieAgentSqlParameterDefaultValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlParameter
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlParameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_value": m.DefaultValue,
			"description":   m.Description,
			"name":          m.Name,
			"type_hint":     m.TypeHint,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlParameter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_value": GenieAgentSqlParameterDefaultValue{}.Type(ctx),
			"description":   types.StringType,
			"name":          types.StringType,
			"type_hint":     types.StringType,
		},
	}
}

// GetDefaultValue returns the value of the DefaultValue field in GenieAgentSqlParameter as
// a GenieAgentSqlParameterDefaultValue value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlParameter) GetDefaultValue(ctx context.Context) (GenieAgentSqlParameterDefaultValue, bool) {
	var e GenieAgentSqlParameterDefaultValue
	if m.DefaultValue.IsNull() || m.DefaultValue.IsUnknown() {
		return e, false
	}
	var v GenieAgentSqlParameterDefaultValue
	d := m.DefaultValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultValue sets the value of the DefaultValue field in GenieAgentSqlParameter.
func (m *GenieAgentSqlParameter) SetDefaultValue(ctx context.Context, v GenieAgentSqlParameterDefaultValue) {
	vs := v.ToObjectValue(ctx)
	m.DefaultValue = vs
}

// A range of default values for a SQL parameter.
type GenieAgentSqlParameterDefaultRange struct {
	// The upper end of the range.
	Max types.String `tfsdk:"max"`
	// The lower end of the range.
	Min types.String `tfsdk:"min"`
}

func (to *GenieAgentSqlParameterDefaultRange) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlParameterDefaultRange) {
}

func (to *GenieAgentSqlParameterDefaultRange) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlParameterDefaultRange) {
}

func (m GenieAgentSqlParameterDefaultRange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["max"] = attrs["max"].SetOptional()
	attrs["min"] = attrs["min"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlParameterDefaultRange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlParameterDefaultRange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlParameterDefaultRange
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlParameterDefaultRange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max": m.Max,
			"min": m.Min,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlParameterDefaultRange) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max": types.StringType,
			"min": types.StringType,
		},
	}
}

// The default value of a SQL parameter: a list of values or a range.
type GenieAgentSqlParameterDefaultValue struct {
	// The default range.
	ValueRange types.Object `tfsdk:"value_range"`
	// The default values. A single default is a one-element list.
	Values types.List `tfsdk:"values"`
}

func (to *GenieAgentSqlParameterDefaultValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlParameterDefaultValue) {
	if !from.ValueRange.IsNull() && !from.ValueRange.IsUnknown() {
		if toValueRange, ok := to.GetValueRange(ctx); ok {
			if fromValueRange, ok := from.GetValueRange(ctx); ok {
				// Recursively sync the fields of ValueRange
				toValueRange.SyncFieldsDuringCreateOrUpdate(ctx, fromValueRange)
				to.SetValueRange(ctx, toValueRange)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *GenieAgentSqlParameterDefaultValue) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlParameterDefaultValue) {
	if !from.ValueRange.IsNull() && !from.ValueRange.IsUnknown() {
		if toValueRange, ok := to.GetValueRange(ctx); ok {
			if fromValueRange, ok := from.GetValueRange(ctx); ok {
				toValueRange.SyncFieldsDuringRead(ctx, fromValueRange)
				to.SetValueRange(ctx, toValueRange)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m GenieAgentSqlParameterDefaultValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value_range"] = attrs["value_range"].SetOptional()
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlParameterDefaultValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlParameterDefaultValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value_range": reflect.TypeOf(GenieAgentSqlParameterDefaultRange{}),
		"values":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlParameterDefaultValue
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlParameterDefaultValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value_range": m.ValueRange,
			"values":      m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlParameterDefaultValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value_range": GenieAgentSqlParameterDefaultRange{}.Type(ctx),
			"values": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetValueRange returns the value of the ValueRange field in GenieAgentSqlParameterDefaultValue as
// a GenieAgentSqlParameterDefaultRange value.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlParameterDefaultValue) GetValueRange(ctx context.Context) (GenieAgentSqlParameterDefaultRange, bool) {
	var e GenieAgentSqlParameterDefaultRange
	if m.ValueRange.IsNull() || m.ValueRange.IsUnknown() {
		return e, false
	}
	var v GenieAgentSqlParameterDefaultRange
	d := m.ValueRange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValueRange sets the value of the ValueRange field in GenieAgentSqlParameterDefaultValue.
func (m *GenieAgentSqlParameterDefaultValue) SetValueRange(ctx context.Context, v GenieAgentSqlParameterDefaultRange) {
	vs := v.ToObjectValue(ctx)
	m.ValueRange = vs
}

// GetValues returns the value of the Values field in GenieAgentSqlParameterDefaultValue as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlParameterDefaultValue) GetValues(ctx context.Context) ([]types.String, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in GenieAgentSqlParameterDefaultValue.
func (m *GenieAgentSqlParameterDefaultValue) SetValues(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

// Reusable SQL snippets.
type GenieAgentSqlSnippets struct {
	// Non-aggregated expressions (dimensions), such as `price * units`.
	Expressions types.List `tfsdk:"expressions"`
	// Filters, such as `state = 'CA'`.
	Filters types.List `tfsdk:"filters"`
	// Aggregated measures, such as `SUM(price * units)`.
	Measures types.List `tfsdk:"measures"`
}

func (to *GenieAgentSqlSnippets) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentSqlSnippets) {
	if !from.Expressions.IsNull() && !from.Expressions.IsUnknown() && to.Expressions.IsNull() && len(from.Expressions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Expressions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Expressions = from.Expressions
	}
	if !from.Expressions.IsNull() && !from.Expressions.IsUnknown() {
		if toExpressions, ok := to.GetExpressions(ctx); ok {
			if fromExpressions, ok := from.GetExpressions(ctx); ok {
				// Recursively sync the fields of each Expressions element by position.
				for i := range toExpressions {
					if i < len(fromExpressions) {
						toExpressions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromExpressions[i])
					}
				}
				to.SetExpressions(ctx, toExpressions)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				// Recursively sync the fields of each Filters element by position.
				for i := range toFilters {
					if i < len(fromFilters) {
						toFilters[i].SyncFieldsDuringCreateOrUpdate(ctx, fromFilters[i])
					}
				}
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.Measures.IsNull() && !from.Measures.IsUnknown() && to.Measures.IsNull() && len(from.Measures.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Measures, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Measures = from.Measures
	}
	if !from.Measures.IsNull() && !from.Measures.IsUnknown() {
		if toMeasures, ok := to.GetMeasures(ctx); ok {
			if fromMeasures, ok := from.GetMeasures(ctx); ok {
				// Recursively sync the fields of each Measures element by position.
				for i := range toMeasures {
					if i < len(fromMeasures) {
						toMeasures[i].SyncFieldsDuringCreateOrUpdate(ctx, fromMeasures[i])
					}
				}
				to.SetMeasures(ctx, toMeasures)
			}
		}
	}
}

func (to *GenieAgentSqlSnippets) SyncFieldsDuringRead(ctx context.Context, from GenieAgentSqlSnippets) {
	if !from.Expressions.IsNull() && !from.Expressions.IsUnknown() && to.Expressions.IsNull() && len(from.Expressions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Expressions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Expressions = from.Expressions
	}
	if !from.Expressions.IsNull() && !from.Expressions.IsUnknown() {
		if toExpressions, ok := to.GetExpressions(ctx); ok {
			if fromExpressions, ok := from.GetExpressions(ctx); ok {
				for i := range toExpressions {
					if i < len(fromExpressions) {
						toExpressions[i].SyncFieldsDuringRead(ctx, fromExpressions[i])
					}
				}
				to.SetExpressions(ctx, toExpressions)
			}
		}
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() && to.Filters.IsNull() && len(from.Filters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Filters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Filters = from.Filters
	}
	if !from.Filters.IsNull() && !from.Filters.IsUnknown() {
		if toFilters, ok := to.GetFilters(ctx); ok {
			if fromFilters, ok := from.GetFilters(ctx); ok {
				for i := range toFilters {
					if i < len(fromFilters) {
						toFilters[i].SyncFieldsDuringRead(ctx, fromFilters[i])
					}
				}
				to.SetFilters(ctx, toFilters)
			}
		}
	}
	if !from.Measures.IsNull() && !from.Measures.IsUnknown() && to.Measures.IsNull() && len(from.Measures.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Measures, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Measures = from.Measures
	}
	if !from.Measures.IsNull() && !from.Measures.IsUnknown() {
		if toMeasures, ok := to.GetMeasures(ctx); ok {
			if fromMeasures, ok := from.GetMeasures(ctx); ok {
				for i := range toMeasures {
					if i < len(fromMeasures) {
						toMeasures[i].SyncFieldsDuringRead(ctx, fromMeasures[i])
					}
				}
				to.SetMeasures(ctx, toMeasures)
			}
		}
	}
}

func (m GenieAgentSqlSnippets) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["expressions"] = attrs["expressions"].SetOptional()
	attrs["filters"] = attrs["filters"].SetOptional()
	attrs["measures"] = attrs["measures"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentSqlSnippets.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentSqlSnippets) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"expressions": reflect.TypeOf(GenieAgentSqlExpression{}),
		"filters":     reflect.TypeOf(GenieAgentSqlFilter{}),
		"measures":    reflect.TypeOf(GenieAgentSqlMeasure{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentSqlSnippets
// only implements ToObjectValue() and Type().
func (m GenieAgentSqlSnippets) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expressions": m.Expressions,
			"filters":     m.Filters,
			"measures":    m.Measures,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentSqlSnippets) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expressions": basetypes.ListType{
				ElemType: GenieAgentSqlExpression{}.Type(ctx),
			},
			"filters": basetypes.ListType{
				ElemType: GenieAgentSqlFilter{}.Type(ctx),
			},
			"measures": basetypes.ListType{
				ElemType: GenieAgentSqlMeasure{}.Type(ctx),
			},
		},
	}
}

// GetExpressions returns the value of the Expressions field in GenieAgentSqlSnippets as
// a slice of GenieAgentSqlExpression values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlSnippets) GetExpressions(ctx context.Context) ([]GenieAgentSqlExpression, bool) {
	if m.Expressions.IsNull() || m.Expressions.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentSqlExpression
	d := m.Expressions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExpressions sets the value of the Expressions field in GenieAgentSqlSnippets.
func (m *GenieAgentSqlSnippets) SetExpressions(ctx context.Context, v []GenieAgentSqlExpression) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["expressions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Expressions = types.ListValueMust(t, vs)
}

// GetFilters returns the value of the Filters field in GenieAgentSqlSnippets as
// a slice of GenieAgentSqlFilter values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlSnippets) GetFilters(ctx context.Context) ([]GenieAgentSqlFilter, bool) {
	if m.Filters.IsNull() || m.Filters.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentSqlFilter
	d := m.Filters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFilters sets the value of the Filters field in GenieAgentSqlSnippets.
func (m *GenieAgentSqlSnippets) SetFilters(ctx context.Context, v []GenieAgentSqlFilter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["filters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Filters = types.ListValueMust(t, vs)
}

// GetMeasures returns the value of the Measures field in GenieAgentSqlSnippets as
// a slice of GenieAgentSqlMeasure values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentSqlSnippets) GetMeasures(ctx context.Context) ([]GenieAgentSqlMeasure, bool) {
	if m.Measures.IsNull() || m.Measures.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentSqlMeasure
	d := m.Measures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMeasures sets the value of the Measures field in GenieAgentSqlSnippets.
func (m *GenieAgentSqlSnippets) SetMeasures(ctx context.Context, v []GenieAgentSqlMeasure) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["measures"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Measures = types.ListValueMust(t, vs)
}

// A Unity Catalog table, view, or metric view the Genie Agent can access.
type GenieAgentTable struct {
	// Column settings for the table, sorted by `column`.
	ColumnConfigs types.List `tfsdk:"column_configs"`
	// Description of the table.
	Description types.String `tfsdk:"description"`
	// The full three-level name, for example `catalog.schema.table`.
	Identifier types.String `tfsdk:"identifier"`
}

func (to *GenieAgentTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentTable) {
	if !from.ColumnConfigs.IsNull() && !from.ColumnConfigs.IsUnknown() && to.ColumnConfigs.IsNull() && len(from.ColumnConfigs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnConfigs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnConfigs = from.ColumnConfigs
	}
	if !from.ColumnConfigs.IsNull() && !from.ColumnConfigs.IsUnknown() {
		if toColumnConfigs, ok := to.GetColumnConfigs(ctx); ok {
			if fromColumnConfigs, ok := from.GetColumnConfigs(ctx); ok {
				// Recursively sync the fields of each ColumnConfigs element by position.
				for i := range toColumnConfigs {
					if i < len(fromColumnConfigs) {
						toColumnConfigs[i].SyncFieldsDuringCreateOrUpdate(ctx, fromColumnConfigs[i])
					}
				}
				to.SetColumnConfigs(ctx, toColumnConfigs)
			}
		}
	}
}

func (to *GenieAgentTable) SyncFieldsDuringRead(ctx context.Context, from GenieAgentTable) {
	if !from.ColumnConfigs.IsNull() && !from.ColumnConfigs.IsUnknown() && to.ColumnConfigs.IsNull() && len(from.ColumnConfigs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnConfigs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnConfigs = from.ColumnConfigs
	}
	if !from.ColumnConfigs.IsNull() && !from.ColumnConfigs.IsUnknown() {
		if toColumnConfigs, ok := to.GetColumnConfigs(ctx); ok {
			if fromColumnConfigs, ok := from.GetColumnConfigs(ctx); ok {
				for i := range toColumnConfigs {
					if i < len(fromColumnConfigs) {
						toColumnConfigs[i].SyncFieldsDuringRead(ctx, fromColumnConfigs[i])
					}
				}
				to.SetColumnConfigs(ctx, toColumnConfigs)
			}
		}
	}
}

func (m GenieAgentTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column_configs"] = attrs["column_configs"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["identifier"] = attrs["identifier"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"column_configs": reflect.TypeOf(GenieAgentColumnConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentTable
// only implements ToObjectValue() and Type().
func (m GenieAgentTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_configs": m.ColumnConfigs,
			"description":    m.Description,
			"identifier":     m.Identifier,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_configs": basetypes.ListType{
				ElemType: GenieAgentColumnConfig{}.Type(ctx),
			},
			"description": types.StringType,
			"identifier":  types.StringType,
		},
	}
}

// GetColumnConfigs returns the value of the ColumnConfigs field in GenieAgentTable as
// a slice of GenieAgentColumnConfig values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenieAgentTable) GetColumnConfigs(ctx context.Context) ([]GenieAgentColumnConfig, bool) {
	if m.ColumnConfigs.IsNull() || m.ColumnConfigs.IsUnknown() {
		return nil, false
	}
	var v []GenieAgentColumnConfig
	d := m.ColumnConfigs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnConfigs sets the value of the ColumnConfigs field in GenieAgentTable.
func (m *GenieAgentTable) SetColumnConfigs(ctx context.Context, v []GenieAgentColumnConfig) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["column_configs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnConfigs = types.ListValueMust(t, vs)
}

// A general text instruction for the Genie Agent.
type GenieAgentTextInstruction struct {
	// The text of the instruction.
	Content types.String `tfsdk:"content"`
	// Identifier of the instruction: 32 lowercase hexadecimal characters,
	// unique within the Genie Agent.
	Id types.String `tfsdk:"id"`
}

func (to *GenieAgentTextInstruction) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentTextInstruction) {
}

func (to *GenieAgentTextInstruction) SyncFieldsDuringRead(ctx context.Context, from GenieAgentTextInstruction) {
}

func (m GenieAgentTextInstruction) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentTextInstruction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentTextInstruction) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentTextInstruction
// only implements ToObjectValue() and Type().
func (m GenieAgentTextInstruction) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": m.Content,
			"id":      m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentTextInstruction) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"id":      types.StringType,
		},
	}
}

// A Unity Catalog volume the Genie Agent can access.
type GenieAgentVolume struct {
	// Description of the volume.
	Description types.String `tfsdk:"description"`
	// The volume path, for example `/Volumes/catalog/schema/volume`. May name a
	// directory or file within the volume.
	Path types.String `tfsdk:"path"`
}

func (to *GenieAgentVolume) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieAgentVolume) {
}

func (to *GenieAgentVolume) SyncFieldsDuringRead(ctx context.Context, from GenieAgentVolume) {
}

func (m GenieAgentVolume) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAgentVolume.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieAgentVolume) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieAgentVolume
// only implements ToObjectValue() and Type().
func (m GenieAgentVolume) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"path":        m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieAgentVolume) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"path":        types.StringType,
		},
	}
}

type GetGenieAgentRequest struct {
	// The resource name of the Genie Agent to retrieve, in the form
	// `agents/{genie_agent_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetGenieAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetGenieAgentRequest) {
}

func (to *GetGenieAgentRequest) SyncFieldsDuringRead(ctx context.Context, from GetGenieAgentRequest) {
}

func (m GetGenieAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGenieAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetGenieAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGenieAgentRequest
// only implements ToObjectValue() and Type().
func (m GetGenieAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetGenieAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListGenieAgentsRequest struct {
	// Maximum number of Genie Agents to return per page. Defaults to 20 when
	// unset or 0; values above 100 are coerced to 100. Negative values are
	// rejected with INVALID_PARAMETER_VALUE. With `GENIE_AGENT_VIEW_FULL`, a
	// smaller per-page maximum applies.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token from a previous response's `next_page_token`. A token
	// that was not returned by this method is rejected with
	// INVALID_PARAMETER_VALUE.
	PageToken types.String `tfsdk:"-"`
	// The fields to return for each Genie Agent. Defaults to
	// `GENIE_AGENT_VIEW_BASIC`. Use `GENIE_AGENT_VIEW_FULL` only when the agent
	// bodies are needed, since it returns smaller pages.
	View types.String `tfsdk:"-"`
}

func (to *ListGenieAgentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGenieAgentsRequest) {
}

func (to *ListGenieAgentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListGenieAgentsRequest) {
}

func (m ListGenieAgentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGenieAgentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGenieAgentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGenieAgentsRequest
// only implements ToObjectValue() and Type().
func (m ListGenieAgentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"view":       m.View,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGenieAgentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"view":       types.StringType,
		},
	}
}

// Response for a list request: a page of Genie Agents plus the token for the
// next page.
type ListGenieAgentsResponse struct {
	// The Genie Agents, with the fields selected by the request's `view`.
	GenieAgents types.List `tfsdk:"genie_agents"`
	// Token to get the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Resource names, in the form `agents/{genie_agent_id}`, of Genie Agents in
	// this page that could not be read and are omitted from `genie_agents`.
	// Populated only for `GENIE_AGENT_VIEW_FULL`. To see why, get the Genie
	// Agent by name; to retry, repeat the request. Other kinds of resource
	// names may be added to this field in the future.
	Unreachable types.Set `tfsdk:"unreachable"`
}

func (to *ListGenieAgentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListGenieAgentsResponse) {
	if !from.GenieAgents.IsNull() && !from.GenieAgents.IsUnknown() && to.GenieAgents.IsNull() && len(from.GenieAgents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for GenieAgents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.GenieAgents = from.GenieAgents
	}
	if !from.GenieAgents.IsNull() && !from.GenieAgents.IsUnknown() {
		if toGenieAgents, ok := to.GetGenieAgents(ctx); ok {
			if fromGenieAgents, ok := from.GetGenieAgents(ctx); ok {
				// Recursively sync the fields of each GenieAgents element by position.
				for i := range toGenieAgents {
					if i < len(fromGenieAgents) {
						toGenieAgents[i].SyncFieldsDuringCreateOrUpdate(ctx, fromGenieAgents[i])
					}
				}
				to.SetGenieAgents(ctx, toGenieAgents)
			}
		}
	}
	if !from.Unreachable.IsNull() && !from.Unreachable.IsUnknown() && to.Unreachable.IsNull() && len(from.Unreachable.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Unreachable, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Unreachable = from.Unreachable
	}
}

func (to *ListGenieAgentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListGenieAgentsResponse) {
	if !from.GenieAgents.IsNull() && !from.GenieAgents.IsUnknown() && to.GenieAgents.IsNull() && len(from.GenieAgents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for GenieAgents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.GenieAgents = from.GenieAgents
	}
	if !from.GenieAgents.IsNull() && !from.GenieAgents.IsUnknown() {
		if toGenieAgents, ok := to.GetGenieAgents(ctx); ok {
			if fromGenieAgents, ok := from.GetGenieAgents(ctx); ok {
				for i := range toGenieAgents {
					if i < len(fromGenieAgents) {
						toGenieAgents[i].SyncFieldsDuringRead(ctx, fromGenieAgents[i])
					}
				}
				to.SetGenieAgents(ctx, toGenieAgents)
			}
		}
	}
	if !from.Unreachable.IsNull() && !from.Unreachable.IsUnknown() && to.Unreachable.IsNull() && len(from.Unreachable.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Unreachable, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Unreachable = from.Unreachable
	}
}

func (m ListGenieAgentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["genie_agents"] = attrs["genie_agents"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["unreachable"] = attrs["unreachable"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGenieAgentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListGenieAgentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"genie_agents": reflect.TypeOf(GenieAgent{}),
		"unreachable":  reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListGenieAgentsResponse
// only implements ToObjectValue() and Type().
func (m ListGenieAgentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"genie_agents":    m.GenieAgents,
			"next_page_token": m.NextPageToken,
			"unreachable":     m.Unreachable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListGenieAgentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"genie_agents": basetypes.ListType{
				ElemType: GenieAgent{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"unreachable": basetypes.SetType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetGenieAgents returns the value of the GenieAgents field in ListGenieAgentsResponse as
// a slice of GenieAgent values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListGenieAgentsResponse) GetGenieAgents(ctx context.Context) ([]GenieAgent, bool) {
	if m.GenieAgents.IsNull() || m.GenieAgents.IsUnknown() {
		return nil, false
	}
	var v []GenieAgent
	d := m.GenieAgents.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGenieAgents sets the value of the GenieAgents field in ListGenieAgentsResponse.
func (m *ListGenieAgentsResponse) SetGenieAgents(ctx context.Context, v []GenieAgent) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["genie_agents"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.GenieAgents = types.ListValueMust(t, vs)
}

// GetUnreachable returns the value of the Unreachable field in ListGenieAgentsResponse as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListGenieAgentsResponse) GetUnreachable(ctx context.Context) ([]types.String, bool) {
	if m.Unreachable.IsNull() || m.Unreachable.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Unreachable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUnreachable sets the value of the Unreachable field in ListGenieAgentsResponse.
func (m *ListGenieAgentsResponse) SetUnreachable(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["unreachable"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Unreachable = types.SetValueMust(t, vs)
}

type UpdateGenieAgentRequest struct {
	// Etag from a previous read. On mismatch the update fails with ABORTED
	// (metadata changes may already be applied); get the agent and retry. May
	// return TEMPORARILY_UNAVAILABLE.
	Etag types.String `tfsdk:"-"`
	// The Genie Agent to update. Its `name` identifies the resource; fields
	// named in update_mask are applied.
	GenieAgent types.Object `tfsdk:"genie_agent"`
	// The resource name of the Genie Agent, in the form
	// `agents/{genie_agent_id}`. Ignored on create.
	Name types.String `tfsdk:"-"`
	// The set of fields to update. The mask is required and must name at least
	// one field; an empty or absent mask is rejected with
	// INVALID_PARAMETER_VALUE, as is a path that is not an updatable field.
	//
	// A field named in the mask takes the value from `genie_agent`, even when
	// it is unset there: an unset `description` is cleared; an unset
	// `display_name`, `warehouse_id`, or `definition` is rejected with
	// INVALID_PARAMETER_VALUE; an unset `parent_path` is cleared without moving
	// the Genie Agent. `definition` is replaced as a whole; paths within it are
	// not supported.
	//
	// The special value `*` replaces the whole Genie Agent, as if every
	// updatable field were named. It must be the only path in the mask. Prefer
	// naming fields explicitly: with `*`, any field added to Genie Agent later
	// is also replaced, so a client unaware of it may clear it.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateGenieAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateGenieAgentRequest) {
	if !from.GenieAgent.IsNull() && !from.GenieAgent.IsUnknown() {
		if toGenieAgent, ok := to.GetGenieAgent(ctx); ok {
			if fromGenieAgent, ok := from.GetGenieAgent(ctx); ok {
				// Recursively sync the fields of GenieAgent
				toGenieAgent.SyncFieldsDuringCreateOrUpdate(ctx, fromGenieAgent)
				to.SetGenieAgent(ctx, toGenieAgent)
			}
		}
	}
}

func (to *UpdateGenieAgentRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateGenieAgentRequest) {
	if !from.GenieAgent.IsNull() && !from.GenieAgent.IsUnknown() {
		if toGenieAgent, ok := to.GetGenieAgent(ctx); ok {
			if fromGenieAgent, ok := from.GetGenieAgent(ctx); ok {
				toGenieAgent.SyncFieldsDuringRead(ctx, fromGenieAgent)
				to.SetGenieAgent(ctx, toGenieAgent)
			}
		}
	}
}

func (m UpdateGenieAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["genie_agent"] = attrs["genie_agent"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateGenieAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateGenieAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"genie_agent": reflect.TypeOf(GenieAgent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateGenieAgentRequest
// only implements ToObjectValue() and Type().
func (m UpdateGenieAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":        m.Etag,
			"genie_agent": m.GenieAgent,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateGenieAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":        types.StringType,
			"genie_agent": GenieAgent{}.Type(ctx),
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetGenieAgent returns the value of the GenieAgent field in UpdateGenieAgentRequest as
// a GenieAgent value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateGenieAgentRequest) GetGenieAgent(ctx context.Context) (GenieAgent, bool) {
	var e GenieAgent
	if m.GenieAgent.IsNull() || m.GenieAgent.IsUnknown() {
		return e, false
	}
	var v GenieAgent
	d := m.GenieAgent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGenieAgent sets the value of the GenieAgent field in UpdateGenieAgentRequest.
func (m *UpdateGenieAgentRequest) SetGenieAgent(ctx context.Context, v GenieAgent) {
	vs := v.ToObjectValue(ctx)
	m.GenieAgent = vs
}
