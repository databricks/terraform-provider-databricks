// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package aifunctions_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ai_classify
type AiClassifyOptions_SdkV2 struct {
	// When true, includes a per-label confidence score in the response.
	EnableConfidenceScores types.Bool `tfsdk:"enable_confidence_scores"`
	// When true, includes a rationale explaining each classification in the
	// response.
	EnableRationales types.Bool `tfsdk:"enable_rationales"`
	// Natural-language guidance that steers how the text is classified (up to
	// 20,000 characters).
	Instructions types.String `tfsdk:"instructions"`
	// When true, allows more than one label to be returned per input.
	Multilabel types.Bool `tfsdk:"multilabel"`
	// The function version to invoke. Defaults to the latest version. Supported
	// versions: ["2.1"].
	Version types.String `tfsdk:"version"`
}

func (to *AiClassifyOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiClassifyOptions_SdkV2) {
}

func (to *AiClassifyOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiClassifyOptions_SdkV2) {
}

func (m AiClassifyOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enable_confidence_scores"] = attrs["enable_confidence_scores"].SetOptional()
	attrs["enable_rationales"] = attrs["enable_rationales"].SetOptional()
	attrs["instructions"] = attrs["instructions"].SetOptional()
	attrs["multilabel"] = attrs["multilabel"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiClassifyOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiClassifyOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiClassifyOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m AiClassifyOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_confidence_scores": m.EnableConfidenceScores,
			"enable_rationales":        m.EnableRationales,
			"instructions":             m.Instructions,
			"multilabel":               m.Multilabel,
			"version":                  m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiClassifyOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enable_confidence_scores": types.BoolType,
			"enable_rationales":        types.BoolType,
			"instructions":             types.StringType,
			"multilabel":               types.BoolType,
			"version":                  types.StringType,
		},
	}
}

type AiClassifyRequest_SdkV2 struct {
	// The content to classify. It accepts a plain string or the response object
	// of [ai_parse_document](:method:AiFunctions/AiParseDocument).
	Content jsontypes.Normalized `tfsdk:"content"`
	// The label set to classify as. Either a JSON array of label strings (e.g.
	// ["spam", "not_spam"]), or a JSON object mapping each label to a
	// description (e.g. {"spam": "unsolicited bulk message", "not_spam": "a
	// legitimate message"}). Accepts 2 to 500 labels, each 1 to 100 characters.
	Labels jsontypes.Normalized `tfsdk:"labels"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options types.List `tfsdk:"options"`
}

func (to *AiClassifyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiClassifyRequest_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (to *AiClassifyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiClassifyRequest_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (m AiClassifyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetRequired()
	attrs["labels"] = attrs["labels"].SetRequired()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiClassifyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiClassifyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AiClassifyOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiClassifyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m AiClassifyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": m.Content,
			"labels":  m.Labels,
			"options": m.Options,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiClassifyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": jsontypes.NormalizedType{},
			"labels":  jsontypes.NormalizedType{},
			"options": basetypes.ListType{
				ElemType: AiClassifyOptions_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOptions returns the value of the Options field in AiClassifyRequest_SdkV2 as
// a AiClassifyOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiClassifyRequest_SdkV2) GetOptions(ctx context.Context) (AiClassifyOptions_SdkV2, bool) {
	var e AiClassifyOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []AiClassifyOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in AiClassifyRequest_SdkV2.
func (m *AiClassifyRequest_SdkV2) SetOptions(ctx context.Context, v AiClassifyOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

type AiClassifyResponse_SdkV2 struct {
	// Additional metadata returned by AI Classify.
	Metadata types.List `tfsdk:"metadata"`
	// The function result as a JSON value. An array of per-label objects: one
	// element in single-label mode (the default), or multiple elements when
	// `multilabel` is true. When `enable_confidence_scores` and
	// `enable_rationales` are true, `confidence_score` and `rationale` are
	// included in each response value, respectively.
	Response jsontypes.Normalized `tfsdk:"response"`
}

func (to *AiClassifyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiClassifyResponse_SdkV2) {
	if !from.Metadata.IsNull() && !from.Metadata.IsUnknown() {
		if toMetadata, ok := to.GetMetadata(ctx); ok {
			if fromMetadata, ok := from.GetMetadata(ctx); ok {
				// Recursively sync the fields of Metadata
				toMetadata.SyncFieldsDuringCreateOrUpdate(ctx, fromMetadata)
				to.SetMetadata(ctx, toMetadata)
			}
		}
	}
}

func (to *AiClassifyResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiClassifyResponse_SdkV2) {
	if !from.Metadata.IsNull() && !from.Metadata.IsUnknown() {
		if toMetadata, ok := to.GetMetadata(ctx); ok {
			if fromMetadata, ok := from.GetMetadata(ctx); ok {
				toMetadata.SyncFieldsDuringRead(ctx, fromMetadata)
				to.SetMetadata(ctx, toMetadata)
			}
		}
	}
}

func (m AiClassifyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["metadata"] = attrs["metadata"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["response"] = attrs["response"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiClassifyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiClassifyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metadata": reflect.TypeOf(AiClassifyResponseMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiClassifyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m AiClassifyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metadata": m.Metadata,
			"response": m.Response,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiClassifyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metadata": basetypes.ListType{
				ElemType: AiClassifyResponseMetadata_SdkV2{}.Type(ctx),
			},
			"response": jsontypes.NormalizedType{},
		},
	}
}

// GetMetadata returns the value of the Metadata field in AiClassifyResponse_SdkV2 as
// a AiClassifyResponseMetadata_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiClassifyResponse_SdkV2) GetMetadata(ctx context.Context) (AiClassifyResponseMetadata_SdkV2, bool) {
	var e AiClassifyResponseMetadata_SdkV2
	if m.Metadata.IsNull() || m.Metadata.IsUnknown() {
		return e, false
	}
	var v []AiClassifyResponseMetadata_SdkV2
	d := m.Metadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetadata sets the value of the Metadata field in AiClassifyResponse_SdkV2.
func (m *AiClassifyResponse_SdkV2) SetMetadata(ctx context.Context, v AiClassifyResponseMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metadata"]
	m.Metadata = types.ListValueMust(t, vs)
}

type AiClassifyResponseMetadata_SdkV2 struct {
	// The resolved function version.
	Version types.String `tfsdk:"version"`
}

func (to *AiClassifyResponseMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiClassifyResponseMetadata_SdkV2) {
}

func (to *AiClassifyResponseMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiClassifyResponseMetadata_SdkV2) {
}

func (m AiClassifyResponseMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiClassifyResponseMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiClassifyResponseMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiClassifyResponseMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m AiClassifyResponseMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiClassifyResponseMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"version": types.StringType,
		},
	}
}

// A bounding box on a source page; used by bbox-input citations.
type AiExtractBbox_SdkV2 struct {
	// Pixel coordinates on the page image as [x0, y0, x1, y1].
	Coord types.List `tfsdk:"coord"`
	// 0-based page index the box is on.
	PageId types.Int64 `tfsdk:"page_id"`
}

func (to *AiExtractBbox_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiExtractBbox_SdkV2) {
	if !from.Coord.IsNull() && !from.Coord.IsUnknown() && to.Coord.IsNull() && len(from.Coord.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Coord, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Coord = from.Coord
	}
}

func (to *AiExtractBbox_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiExtractBbox_SdkV2) {
	if !from.Coord.IsNull() && !from.Coord.IsUnknown() && to.Coord.IsNull() && len(from.Coord.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Coord, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Coord = from.Coord
	}
}

func (m AiExtractBbox_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["coord"] = attrs["coord"].SetOptional()
	attrs["page_id"] = attrs["page_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiExtractBbox.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiExtractBbox_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"coord": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiExtractBbox_SdkV2
// only implements ToObjectValue() and Type().
func (m AiExtractBbox_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"coord":   m.Coord,
			"page_id": m.PageId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiExtractBbox_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"coord": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"page_id": types.Int64Type,
		},
	}
}

// GetCoord returns the value of the Coord field in AiExtractBbox_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AiExtractBbox_SdkV2) GetCoord(ctx context.Context) ([]types.Int64, bool) {
	if m.Coord.IsNull() || m.Coord.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.Coord.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCoord sets the value of the Coord field in AiExtractBbox_SdkV2.
func (m *AiExtractBbox_SdkV2) SetCoord(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["coord"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Coord = types.ListValueMust(t, vs)
}

// A citation locating an extracted value in the source. start/stop are set for
// span (STRING input) citations, bbox for bbox (parsed-document input)
// citations.
type AiExtractCitation_SdkV2 struct {
	// Bounding boxes locating the citation on the source pages; set for bbox
	// citations.
	Bbox types.List `tfsdk:"bbox"`
	// Integer matching a citation_ids entry on an extracted field.
	Id types.Int64 `tfsdk:"id"`
	// Inclusive 0-based character offset into the input string; set for span
	// citations.
	Start types.Int64 `tfsdk:"start"`
	// Exclusive 0-based character offset into the input string; set for span
	// citations.
	Stop types.Int64 `tfsdk:"stop"`
}

func (to *AiExtractCitation_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiExtractCitation_SdkV2) {
	if !from.Bbox.IsNull() && !from.Bbox.IsUnknown() && to.Bbox.IsNull() && len(from.Bbox.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Bbox, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Bbox = from.Bbox
	}
	if !from.Bbox.IsNull() && !from.Bbox.IsUnknown() {
		if toBbox, ok := to.GetBbox(ctx); ok {
			if fromBbox, ok := from.GetBbox(ctx); ok {
				// Recursively sync the fields of each Bbox element by position.
				for i := range toBbox {
					if i < len(fromBbox) {
						toBbox[i].SyncFieldsDuringCreateOrUpdate(ctx, fromBbox[i])
					}
				}
				to.SetBbox(ctx, toBbox)
			}
		}
	}
}

func (to *AiExtractCitation_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiExtractCitation_SdkV2) {
	if !from.Bbox.IsNull() && !from.Bbox.IsUnknown() && to.Bbox.IsNull() && len(from.Bbox.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Bbox, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Bbox = from.Bbox
	}
	if !from.Bbox.IsNull() && !from.Bbox.IsUnknown() {
		if toBbox, ok := to.GetBbox(ctx); ok {
			if fromBbox, ok := from.GetBbox(ctx); ok {
				for i := range toBbox {
					if i < len(fromBbox) {
						toBbox[i].SyncFieldsDuringRead(ctx, fromBbox[i])
					}
				}
				to.SetBbox(ctx, toBbox)
			}
		}
	}
}

func (m AiExtractCitation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bbox"] = attrs["bbox"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["start"] = attrs["start"].SetOptional()
	attrs["stop"] = attrs["stop"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiExtractCitation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiExtractCitation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"bbox": reflect.TypeOf(AiExtractBbox_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiExtractCitation_SdkV2
// only implements ToObjectValue() and Type().
func (m AiExtractCitation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bbox":  m.Bbox,
			"id":    m.Id,
			"start": m.Start,
			"stop":  m.Stop,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiExtractCitation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bbox": basetypes.ListType{
				ElemType: AiExtractBbox_SdkV2{}.Type(ctx),
			},
			"id":    types.Int64Type,
			"start": types.Int64Type,
			"stop":  types.Int64Type,
		},
	}
}

// GetBbox returns the value of the Bbox field in AiExtractCitation_SdkV2 as
// a slice of AiExtractBbox_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AiExtractCitation_SdkV2) GetBbox(ctx context.Context) ([]AiExtractBbox_SdkV2, bool) {
	if m.Bbox.IsNull() || m.Bbox.IsUnknown() {
		return nil, false
	}
	var v []AiExtractBbox_SdkV2
	d := m.Bbox.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBbox sets the value of the Bbox field in AiExtractCitation_SdkV2.
func (m *AiExtractCitation_SdkV2) SetBbox(ctx context.Context, v []AiExtractBbox_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["bbox"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Bbox = types.ListValueMust(t, vs)
}

// ai_extract
type AiExtractOptions_SdkV2 struct {
	// When true, includes citation metadata locating each extracted value in
	// the source. Depending on the type of input, citations can be one of two
	// types:
	//
	// For raw text (STRING) inputs, a citation is a span of text in the
	// original input. Each object in `metadata.citations` has an `id` (integer
	// matching a `citation_ids` entry on a field), a `start` (inclusive 0-based
	// character offset into the input string), and a `stop` (exclusive 0-based
	// character offset into the input string).
	//
	// For PDF documents and images (when using ai_extract downstream of
	// ai_parse_document), a citation is a bounding box in the original input.
	// Each object in `metadata.citations` has an `id` (integer matching a
	// `citation_ids` entry on a field) and a `bbox` (array of {coord, page_id}
	// objects, identical in shape to element.bbox in ai_parse_document output;
	// coord is pixel coordinates on the page image as [x0, y0, x1, y1], and
	// page_id is a 0-based page index).
	EnableCitations types.Bool `tfsdk:"enable_citations"`
	// When true, includes a per-field confidence score in the response.
	EnableConfidenceScores types.Bool `tfsdk:"enable_confidence_scores"`
	// Natural-language guidance that steers how data is extracted (up to 20,000
	// characters).
	Instructions types.String `tfsdk:"instructions"`
	// Extraction mode. Supported modes: "precision" — more powerful
	// extraction for complex schemas, long documents, and reasoning-heavy
	// extractions. Defaults to none (standard extraction).
	Mode types.String `tfsdk:"mode"`
	// The function version to invoke. Defaults to the latest version. Supported
	// versions: ["2.1"].
	Version types.String `tfsdk:"version"`
}

func (to *AiExtractOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiExtractOptions_SdkV2) {
}

func (to *AiExtractOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiExtractOptions_SdkV2) {
}

func (m AiExtractOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enable_citations"] = attrs["enable_citations"].SetOptional()
	attrs["enable_confidence_scores"] = attrs["enable_confidence_scores"].SetOptional()
	attrs["instructions"] = attrs["instructions"].SetOptional()
	attrs["mode"] = attrs["mode"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiExtractOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiExtractOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiExtractOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m AiExtractOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_citations":         m.EnableCitations,
			"enable_confidence_scores": m.EnableConfidenceScores,
			"instructions":             m.Instructions,
			"mode":                     m.Mode,
			"version":                  m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiExtractOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enable_citations":         types.BoolType,
			"enable_confidence_scores": types.BoolType,
			"instructions":             types.StringType,
			"mode":                     types.StringType,
			"version":                  types.StringType,
		},
	}
}

type AiExtractRequest_SdkV2 struct {
	// The text to extract from. It accepts a plain string or the response
	// object of [ai_parse_document](:method:AiFunctions/AiParseDocument).
	Content jsontypes.Normalized `tfsdk:"content"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options types.List `tfsdk:"options"`
	// The extraction schema defining the fields to extract. Either a JSON array
	// of field names, assumed to be strings (e.g. ["company", "valuation"]), or
	// a JSON object mapping each field to its type/description/nullability
	// (e.g. {"company": {"type": "string", "description": "the company
	// name"}}). Accepts up to 256 fields, 12 levels of nesting, and 500 enum
	// values. Supported field types are string, integer, number, boolean, and
	// enum.
	Schema jsontypes.Normalized `tfsdk:"schema"`
}

func (to *AiExtractRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiExtractRequest_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (to *AiExtractRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiExtractRequest_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (m AiExtractRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetRequired()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["schema"] = attrs["schema"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiExtractRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiExtractRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AiExtractOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiExtractRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m AiExtractRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": m.Content,
			"options": m.Options,
			"schema":  m.Schema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiExtractRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": jsontypes.NormalizedType{},
			"options": basetypes.ListType{
				ElemType: AiExtractOptions_SdkV2{}.Type(ctx),
			},
			"schema": jsontypes.NormalizedType{},
		},
	}
}

// GetOptions returns the value of the Options field in AiExtractRequest_SdkV2 as
// a AiExtractOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiExtractRequest_SdkV2) GetOptions(ctx context.Context) (AiExtractOptions_SdkV2, bool) {
	var e AiExtractOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []AiExtractOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in AiExtractRequest_SdkV2.
func (m *AiExtractRequest_SdkV2) SetOptions(ctx context.Context, v AiExtractOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

type AiExtractResponse_SdkV2 struct {
	// Additional metadata returned by AI Extract.
	Metadata types.List `tfsdk:"metadata"`
	// The function result as a JSON value. When `enable_confidence_scores` and
	// `enable_citations` are true, `confidence` and `citation_ids` are included
	// in each response field, respectively.
	Response jsontypes.Normalized `tfsdk:"response"`
}

func (to *AiExtractResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiExtractResponse_SdkV2) {
	if !from.Metadata.IsNull() && !from.Metadata.IsUnknown() {
		if toMetadata, ok := to.GetMetadata(ctx); ok {
			if fromMetadata, ok := from.GetMetadata(ctx); ok {
				// Recursively sync the fields of Metadata
				toMetadata.SyncFieldsDuringCreateOrUpdate(ctx, fromMetadata)
				to.SetMetadata(ctx, toMetadata)
			}
		}
	}
}

func (to *AiExtractResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiExtractResponse_SdkV2) {
	if !from.Metadata.IsNull() && !from.Metadata.IsUnknown() {
		if toMetadata, ok := to.GetMetadata(ctx); ok {
			if fromMetadata, ok := from.GetMetadata(ctx); ok {
				toMetadata.SyncFieldsDuringRead(ctx, fromMetadata)
				to.SetMetadata(ctx, toMetadata)
			}
		}
	}
}

func (m AiExtractResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["metadata"] = attrs["metadata"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["response"] = attrs["response"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiExtractResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiExtractResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metadata": reflect.TypeOf(AiExtractResponseMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiExtractResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m AiExtractResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metadata": m.Metadata,
			"response": m.Response,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiExtractResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metadata": basetypes.ListType{
				ElemType: AiExtractResponseMetadata_SdkV2{}.Type(ctx),
			},
			"response": jsontypes.NormalizedType{},
		},
	}
}

// GetMetadata returns the value of the Metadata field in AiExtractResponse_SdkV2 as
// a AiExtractResponseMetadata_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiExtractResponse_SdkV2) GetMetadata(ctx context.Context) (AiExtractResponseMetadata_SdkV2, bool) {
	var e AiExtractResponseMetadata_SdkV2
	if m.Metadata.IsNull() || m.Metadata.IsUnknown() {
		return e, false
	}
	var v []AiExtractResponseMetadata_SdkV2
	d := m.Metadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetadata sets the value of the Metadata field in AiExtractResponse_SdkV2.
func (m *AiExtractResponse_SdkV2) SetMetadata(ctx context.Context, v AiExtractResponseMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metadata"]
	m.Metadata = types.ListValueMust(t, vs)
}

type AiExtractResponseMetadata_SdkV2 struct {
	// How the source was chunked for citation offsets (span for text input,
	// bbox for parsed-document input); present when citations are enabled.
	ChunkType types.String `tfsdk:"chunk_type"`
	// Citation objects locating each result in the source; present when
	// citations are enabled.
	Citations types.List `tfsdk:"citations"`
	// The resolved extraction mode; present when a non-default mode was used.
	Mode types.String `tfsdk:"mode"`
	// The resolved function version.
	Version types.String `tfsdk:"version"`
}

func (to *AiExtractResponseMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiExtractResponseMetadata_SdkV2) {
	if !from.Citations.IsNull() && !from.Citations.IsUnknown() && to.Citations.IsNull() && len(from.Citations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Citations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Citations = from.Citations
	}
	if !from.Citations.IsNull() && !from.Citations.IsUnknown() {
		if toCitations, ok := to.GetCitations(ctx); ok {
			if fromCitations, ok := from.GetCitations(ctx); ok {
				// Recursively sync the fields of each Citations element by position.
				for i := range toCitations {
					if i < len(fromCitations) {
						toCitations[i].SyncFieldsDuringCreateOrUpdate(ctx, fromCitations[i])
					}
				}
				to.SetCitations(ctx, toCitations)
			}
		}
	}
}

func (to *AiExtractResponseMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiExtractResponseMetadata_SdkV2) {
	if !from.Citations.IsNull() && !from.Citations.IsUnknown() && to.Citations.IsNull() && len(from.Citations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Citations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Citations = from.Citations
	}
	if !from.Citations.IsNull() && !from.Citations.IsUnknown() {
		if toCitations, ok := to.GetCitations(ctx); ok {
			if fromCitations, ok := from.GetCitations(ctx); ok {
				for i := range toCitations {
					if i < len(fromCitations) {
						toCitations[i].SyncFieldsDuringRead(ctx, fromCitations[i])
					}
				}
				to.SetCitations(ctx, toCitations)
			}
		}
	}
}

func (m AiExtractResponseMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["chunk_type"] = attrs["chunk_type"].SetOptional()
	attrs["citations"] = attrs["citations"].SetOptional()
	attrs["mode"] = attrs["mode"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiExtractResponseMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiExtractResponseMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"citations": reflect.TypeOf(AiExtractCitation_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiExtractResponseMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m AiExtractResponseMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunk_type": m.ChunkType,
			"citations":  m.Citations,
			"mode":       m.Mode,
			"version":    m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiExtractResponseMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunk_type": types.StringType,
			"citations": basetypes.ListType{
				ElemType: AiExtractCitation_SdkV2{}.Type(ctx),
			},
			"mode":    types.StringType,
			"version": types.StringType,
		},
	}
}

// GetCitations returns the value of the Citations field in AiExtractResponseMetadata_SdkV2 as
// a slice of AiExtractCitation_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AiExtractResponseMetadata_SdkV2) GetCitations(ctx context.Context) ([]AiExtractCitation_SdkV2, bool) {
	if m.Citations.IsNull() || m.Citations.IsUnknown() {
		return nil, false
	}
	var v []AiExtractCitation_SdkV2
	d := m.Citations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCitations sets the value of the Citations field in AiExtractResponseMetadata_SdkV2.
func (m *AiExtractResponseMetadata_SdkV2) SetCitations(ctx context.Context, v []AiExtractCitation_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["citations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Citations = types.ListValueMust(t, vs)
}

// Metadata about the source file; present only for file-path input.
type AiParseDocumentFileMetadata_SdkV2 struct {
	// Last-modified timestamp of the source file, as an HTTP date string.
	FileModificationTime types.String `tfsdk:"file_modification_time"`
	// Base name of the source file.
	FileName types.String `tfsdk:"file_name"`
	// Unity Catalog volume path of the source file.
	FilePath types.String `tfsdk:"file_path"`
	// Size of the source file in bytes.
	FileSize types.Int64 `tfsdk:"file_size"`
}

func (to *AiParseDocumentFileMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiParseDocumentFileMetadata_SdkV2) {
}

func (to *AiParseDocumentFileMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiParseDocumentFileMetadata_SdkV2) {
}

func (m AiParseDocumentFileMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_modification_time"] = attrs["file_modification_time"].SetOptional()
	attrs["file_name"] = attrs["file_name"].SetOptional()
	attrs["file_path"] = attrs["file_path"].SetOptional()
	attrs["file_size"] = attrs["file_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiParseDocumentFileMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiParseDocumentFileMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiParseDocumentFileMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m AiParseDocumentFileMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_modification_time": m.FileModificationTime,
			"file_name":              m.FileName,
			"file_path":              m.FilePath,
			"file_size":              m.FileSize,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiParseDocumentFileMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_modification_time": types.StringType,
			"file_name":              types.StringType,
			"file_path":              types.StringType,
			"file_size":              types.Int64Type,
		},
	}
}

// ai_parse_document
type AiParseDocumentOptions_SdkV2 struct {
	// Element types for which an AI-generated description is produced. Use "*"
	// (default) to generate descriptions for all supported element types,
	// "figure" to generate them for figures only, or "" (empty string) to
	// generate none. Only figure descriptions are supported for version "2.0",
	// so "*" and "figure" produce the same behavior.
	DescriptionElementTypes types.String `tfsdk:"description_element_types"`
	// Unity Catalog volume path where rendered page and element images are
	// written.
	ImageOutputPath types.String `tfsdk:"image_output_path"`
	// Pages to parse (1-indexed), as a comma-separated list of page numbers or
	// ranges (e.g. "1,3,5-10").
	PageRange types.String `tfsdk:"page_range"`
	// The ai_parse_document output schema version. Supported value: "2.0".
	Version types.String `tfsdk:"version"`
}

func (to *AiParseDocumentOptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiParseDocumentOptions_SdkV2) {
}

func (to *AiParseDocumentOptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiParseDocumentOptions_SdkV2) {
}

func (m AiParseDocumentOptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description_element_types"] = attrs["description_element_types"].SetOptional()
	attrs["image_output_path"] = attrs["image_output_path"].SetOptional()
	attrs["page_range"] = attrs["page_range"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiParseDocumentOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiParseDocumentOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiParseDocumentOptions_SdkV2
// only implements ToObjectValue() and Type().
func (m AiParseDocumentOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description_element_types": m.DescriptionElementTypes,
			"image_output_path":         m.ImageOutputPath,
			"page_range":                m.PageRange,
			"version":                   m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiParseDocumentOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description_element_types": types.StringType,
			"image_output_path":         types.StringType,
			"page_range":                types.StringType,
			"version":                   types.StringType,
		},
	}
}

// A single page that failed to parse while the overall request succeeded.
type AiParseDocumentPageError_SdkV2 struct {
	// Message describing why the page failed.
	ErrorMessage types.String `tfsdk:"error_message"`
	// 0-based index of the page that failed.
	PageId types.Int64 `tfsdk:"page_id"`
}

func (to *AiParseDocumentPageError_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiParseDocumentPageError_SdkV2) {
}

func (to *AiParseDocumentPageError_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiParseDocumentPageError_SdkV2) {
}

func (m AiParseDocumentPageError_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["page_id"] = attrs["page_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiParseDocumentPageError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiParseDocumentPageError_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiParseDocumentPageError_SdkV2
// only implements ToObjectValue() and Type().
func (m AiParseDocumentPageError_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_message": m.ErrorMessage,
			"page_id":       m.PageId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiParseDocumentPageError_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message": types.StringType,
			"page_id":       types.Int64Type,
		},
	}
}

type AiParseDocumentRequest_SdkV2 struct {
	// The document to parse, given as a Unity Catalog volume path to the source
	// file (the REST API accepts only a UC volume path, not inline binary
	// data). Supported formats: PDF, DOCX, DOC, PPTX, PPT, JPG, JPEG, PNG,
	// TIFF. Accepts up to 100 pages and 100 MB per document.
	Content types.String `tfsdk:"content"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options types.List `tfsdk:"options"`
}

func (to *AiParseDocumentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiParseDocumentRequest_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				// Recursively sync the fields of Options
				toOptions.SyncFieldsDuringCreateOrUpdate(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (to *AiParseDocumentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiParseDocumentRequest_SdkV2) {
	if !from.Options.IsNull() && !from.Options.IsUnknown() {
		if toOptions, ok := to.GetOptions(ctx); ok {
			if fromOptions, ok := from.GetOptions(ctx); ok {
				toOptions.SyncFieldsDuringRead(ctx, fromOptions)
				to.SetOptions(ctx, toOptions)
			}
		}
	}
}

func (m AiParseDocumentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetRequired()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["options"] = attrs["options"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiParseDocumentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiParseDocumentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(AiParseDocumentOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiParseDocumentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m AiParseDocumentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": m.Content,
			"options": m.Options,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiParseDocumentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"options": basetypes.ListType{
				ElemType: AiParseDocumentOptions_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOptions returns the value of the Options field in AiParseDocumentRequest_SdkV2 as
// a AiParseDocumentOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiParseDocumentRequest_SdkV2) GetOptions(ctx context.Context) (AiParseDocumentOptions_SdkV2, bool) {
	var e AiParseDocumentOptions_SdkV2
	if m.Options.IsNull() || m.Options.IsUnknown() {
		return e, false
	}
	var v []AiParseDocumentOptions_SdkV2
	d := m.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOptions sets the value of the Options field in AiParseDocumentRequest_SdkV2.
func (m *AiParseDocumentRequest_SdkV2) SetOptions(ctx context.Context, v AiParseDocumentOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	m.Options = types.ListValueMust(t, vs)
}

type AiParseDocumentResponse_SdkV2 struct {
	// The parsed document as a JSON value, containing the extracted pages and
	// elements.
	Document jsontypes.Normalized `tfsdk:"document"`
	// Per-page partial-failure details; present when the request succeeds (2xx)
	// but individual pages fail.
	ErrorStatus types.List `tfsdk:"error_status"`
	// Additional metadata returned by AI Parse Document.
	Metadata types.List `tfsdk:"metadata"`
}

func (to *AiParseDocumentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiParseDocumentResponse_SdkV2) {
	if !from.ErrorStatus.IsNull() && !from.ErrorStatus.IsUnknown() && to.ErrorStatus.IsNull() && len(from.ErrorStatus.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ErrorStatus, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ErrorStatus = from.ErrorStatus
	}
	if !from.ErrorStatus.IsNull() && !from.ErrorStatus.IsUnknown() {
		if toErrorStatus, ok := to.GetErrorStatus(ctx); ok {
			if fromErrorStatus, ok := from.GetErrorStatus(ctx); ok {
				// Recursively sync the fields of each ErrorStatus element by position.
				for i := range toErrorStatus {
					if i < len(fromErrorStatus) {
						toErrorStatus[i].SyncFieldsDuringCreateOrUpdate(ctx, fromErrorStatus[i])
					}
				}
				to.SetErrorStatus(ctx, toErrorStatus)
			}
		}
	}
	if !from.Metadata.IsNull() && !from.Metadata.IsUnknown() {
		if toMetadata, ok := to.GetMetadata(ctx); ok {
			if fromMetadata, ok := from.GetMetadata(ctx); ok {
				// Recursively sync the fields of Metadata
				toMetadata.SyncFieldsDuringCreateOrUpdate(ctx, fromMetadata)
				to.SetMetadata(ctx, toMetadata)
			}
		}
	}
}

func (to *AiParseDocumentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiParseDocumentResponse_SdkV2) {
	if !from.ErrorStatus.IsNull() && !from.ErrorStatus.IsUnknown() && to.ErrorStatus.IsNull() && len(from.ErrorStatus.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ErrorStatus, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ErrorStatus = from.ErrorStatus
	}
	if !from.ErrorStatus.IsNull() && !from.ErrorStatus.IsUnknown() {
		if toErrorStatus, ok := to.GetErrorStatus(ctx); ok {
			if fromErrorStatus, ok := from.GetErrorStatus(ctx); ok {
				for i := range toErrorStatus {
					if i < len(fromErrorStatus) {
						toErrorStatus[i].SyncFieldsDuringRead(ctx, fromErrorStatus[i])
					}
				}
				to.SetErrorStatus(ctx, toErrorStatus)
			}
		}
	}
	if !from.Metadata.IsNull() && !from.Metadata.IsUnknown() {
		if toMetadata, ok := to.GetMetadata(ctx); ok {
			if fromMetadata, ok := from.GetMetadata(ctx); ok {
				toMetadata.SyncFieldsDuringRead(ctx, fromMetadata)
				to.SetMetadata(ctx, toMetadata)
			}
		}
	}
}

func (m AiParseDocumentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["document"] = attrs["document"].SetOptional()
	attrs["error_status"] = attrs["error_status"].SetOptional()
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["metadata"] = attrs["metadata"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiParseDocumentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiParseDocumentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error_status": reflect.TypeOf(AiParseDocumentPageError_SdkV2{}),
		"metadata":     reflect.TypeOf(AiParseDocumentResponseMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiParseDocumentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m AiParseDocumentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"document":     m.Document,
			"error_status": m.ErrorStatus,
			"metadata":     m.Metadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiParseDocumentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"document": jsontypes.NormalizedType{},
			"error_status": basetypes.ListType{
				ElemType: AiParseDocumentPageError_SdkV2{}.Type(ctx),
			},
			"metadata": basetypes.ListType{
				ElemType: AiParseDocumentResponseMetadata_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetErrorStatus returns the value of the ErrorStatus field in AiParseDocumentResponse_SdkV2 as
// a slice of AiParseDocumentPageError_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AiParseDocumentResponse_SdkV2) GetErrorStatus(ctx context.Context) ([]AiParseDocumentPageError_SdkV2, bool) {
	if m.ErrorStatus.IsNull() || m.ErrorStatus.IsUnknown() {
		return nil, false
	}
	var v []AiParseDocumentPageError_SdkV2
	d := m.ErrorStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetErrorStatus sets the value of the ErrorStatus field in AiParseDocumentResponse_SdkV2.
func (m *AiParseDocumentResponse_SdkV2) SetErrorStatus(ctx context.Context, v []AiParseDocumentPageError_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["error_status"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ErrorStatus = types.ListValueMust(t, vs)
}

// GetMetadata returns the value of the Metadata field in AiParseDocumentResponse_SdkV2 as
// a AiParseDocumentResponseMetadata_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiParseDocumentResponse_SdkV2) GetMetadata(ctx context.Context) (AiParseDocumentResponseMetadata_SdkV2, bool) {
	var e AiParseDocumentResponseMetadata_SdkV2
	if m.Metadata.IsNull() || m.Metadata.IsUnknown() {
		return e, false
	}
	var v []AiParseDocumentResponseMetadata_SdkV2
	d := m.Metadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetadata sets the value of the Metadata field in AiParseDocumentResponse_SdkV2.
func (m *AiParseDocumentResponse_SdkV2) SetMetadata(ctx context.Context, v AiParseDocumentResponseMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metadata"]
	m.Metadata = types.ListValueMust(t, vs)
}

type AiParseDocumentResponseMetadata_SdkV2 struct {
	// Describes the source file; present only for file-path input.
	FileMetadata types.List `tfsdk:"file_metadata"`
	// Unique identifier for the parse request.
	Id types.String `tfsdk:"id"`
	// The resolved function version.
	Version types.String `tfsdk:"version"`
}

func (to *AiParseDocumentResponseMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiParseDocumentResponseMetadata_SdkV2) {
	if !from.FileMetadata.IsNull() && !from.FileMetadata.IsUnknown() {
		if toFileMetadata, ok := to.GetFileMetadata(ctx); ok {
			if fromFileMetadata, ok := from.GetFileMetadata(ctx); ok {
				// Recursively sync the fields of FileMetadata
				toFileMetadata.SyncFieldsDuringCreateOrUpdate(ctx, fromFileMetadata)
				to.SetFileMetadata(ctx, toFileMetadata)
			}
		}
	}
}

func (to *AiParseDocumentResponseMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiParseDocumentResponseMetadata_SdkV2) {
	if !from.FileMetadata.IsNull() && !from.FileMetadata.IsUnknown() {
		if toFileMetadata, ok := to.GetFileMetadata(ctx); ok {
			if fromFileMetadata, ok := from.GetFileMetadata(ctx); ok {
				toFileMetadata.SyncFieldsDuringRead(ctx, fromFileMetadata)
				to.SetFileMetadata(ctx, toFileMetadata)
			}
		}
	}
}

func (m AiParseDocumentResponseMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_metadata"] = attrs["file_metadata"].SetOptional()
	attrs["file_metadata"] = attrs["file_metadata"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiParseDocumentResponseMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiParseDocumentResponseMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_metadata": reflect.TypeOf(AiParseDocumentFileMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiParseDocumentResponseMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m AiParseDocumentResponseMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_metadata": m.FileMetadata,
			"id":            m.Id,
			"version":       m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiParseDocumentResponseMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_metadata": basetypes.ListType{
				ElemType: AiParseDocumentFileMetadata_SdkV2{}.Type(ctx),
			},
			"id":      types.StringType,
			"version": types.StringType,
		},
	}
}

// GetFileMetadata returns the value of the FileMetadata field in AiParseDocumentResponseMetadata_SdkV2 as
// a AiParseDocumentFileMetadata_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiParseDocumentResponseMetadata_SdkV2) GetFileMetadata(ctx context.Context) (AiParseDocumentFileMetadata_SdkV2, bool) {
	var e AiParseDocumentFileMetadata_SdkV2
	if m.FileMetadata.IsNull() || m.FileMetadata.IsUnknown() {
		return e, false
	}
	var v []AiParseDocumentFileMetadata_SdkV2
	d := m.FileMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileMetadata sets the value of the FileMetadata field in AiParseDocumentResponseMetadata_SdkV2.
func (m *AiParseDocumentResponseMetadata_SdkV2) SetFileMetadata(ctx context.Context, v AiParseDocumentFileMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["file_metadata"]
	m.FileMetadata = types.ListValueMust(t, vs)
}
