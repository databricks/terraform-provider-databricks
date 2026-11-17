// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package statementexecutionv2_tf

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

// Metadata of one result set chunk, as listed in the result manifest.
type BaseChunkInfo_SdkV2 struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

func (to *BaseChunkInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BaseChunkInfo_SdkV2) {
}

func (to *BaseChunkInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BaseChunkInfo_SdkV2) {
}

func (m BaseChunkInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_count"] = attrs["byte_count"].SetOptional()
	attrs["chunk_index"] = attrs["chunk_index"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["row_offset"] = attrs["row_offset"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BaseChunkInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BaseChunkInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BaseChunkInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m BaseChunkInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":  m.ByteCount,
			"chunk_index": m.ChunkIndex,
			"row_count":   m.RowCount,
			"row_offset":  m.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BaseChunkInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_count":  types.Int64Type,
			"chunk_index": types.Int64Type,
			"row_count":   types.Int64Type,
			"row_offset":  types.Int64Type,
		},
	}
}

// Request to cancel the statement given in the path on the warehouse given in
// the path.
type CancelStatementRequest_SdkV2 struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *CancelStatementRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelStatementRequest_SdkV2) {
}

func (to *CancelStatementRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CancelStatementRequest_SdkV2) {
}

func (m CancelStatementRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()
	attrs["statement_id"] = attrs["statement_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelStatementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelStatementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelStatementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CancelStatementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": m.StatementId,
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelStatementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// Empty response to a cancel request. Receiving it means the cancel request was
// accepted.
type CancelStatementResponse_SdkV2 struct {
}

func (to *CancelStatementResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelStatementResponse_SdkV2) {
}

func (to *CancelStatementResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CancelStatementResponse_SdkV2) {
}

func (m CancelStatementResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelStatementResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelStatementResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelStatementResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CancelStatementResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m CancelStatementResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Describes one column of the result set: its name, SQL type, and position.
type ColumnInfo_SdkV2 struct {
	// The name of the column.
	Name types.String `tfsdk:"name"`
	// The ordinal position of the column (starting at position 0).
	Position types.Int64 `tfsdk:"position"`
	// The format of the interval type.
	TypeIntervalType types.String `tfsdk:"type_interval_type"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName types.String `tfsdk:"type_name"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision types.Int64 `tfsdk:"type_precision"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale types.Int64 `tfsdk:"type_scale"`
	// The full SQL type specification.
	TypeText types.String `tfsdk:"type_text"`
}

func (to *ColumnInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (to *ColumnInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (m ColumnInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["type_interval_type"] = attrs["type_interval_type"].SetOptional()
	attrs["type_name"] = attrs["type_name"].SetOptional()
	attrs["type_precision"] = attrs["type_precision"].SetOptional()
	attrs["type_scale"] = attrs["type_scale"].SetOptional()
	attrs["type_text"] = attrs["type_text"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ColumnInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ColumnInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":               m.Name,
			"position":           m.Position,
			"type_interval_type": m.TypeIntervalType,
			"type_name":          m.TypeName,
			"type_precision":     m.TypePrecision,
			"type_scale":         m.TypeScale,
			"type_text":          m.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ColumnInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":               types.StringType,
			"position":           types.Int64Type,
			"type_interval_type": types.StringType,
			"type_name":          types.StringType,
			"type_precision":     types.Int64Type,
			"type_scale":         types.Int64Type,
			"type_text":          types.StringType,
		},
	}
}

// Request to execute a SQL statement on the warehouse given in the path.
type ExecuteStatementRequest_SdkV2 struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explicitly set.
	ByteLimit types.Int64 `tfsdk:"byte_limit"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog types.String `tfsdk:"catalog"`
	// The fetch disposition provides two modes of fetching results: `INLINE`
	// and `EXTERNAL_LINKS`.
	//
	// Statements executed with `INLINE` disposition will return result data
	// inline, in `JSON_ARRAY` format, in a series of chunks. If a given
	// statement produces a result set with a size larger than 25 MiB, that
	// statement execution is aborted, and no result set will be available.
	//
	// **NOTE** Byte limits are computed based upon internal representations of
	// the result set data, and might not match the sizes visible in JSON
	// responses.
	//
	// Statements executed with `EXTERNAL_LINKS` disposition will return result
	// data as external links: URLs that point to cloud storage internal to the
	// workspace. Using `EXTERNAL_LINKS` disposition allows statements to
	// generate arbitrarily sized result sets for fetching up to 100 GiB. The
	// resulting links have two important properties:
	//
	// 1. They point to resources _external_ to the Databricks compute;
	// therefore any associated authentication information (typically a personal
	// access token, OAuth token, or similar) _must be removed_ when fetching
	// from these links.
	//
	// 2. These are short-lived cloud-storage URLs with a specific expiration,
	// indicated in the response. The behavior when attempting to use an expired
	// link is cloud specific.
	Disposition types.String `tfsdk:"disposition"`
	// Statement execution supports three result formats: `JSON_ARRAY`,
	// `ARROW_STREAM`, and `CSV`. Set `format` explicitly: on Lakehouse//RT
	// warehouses, an omitted `format` with the `INLINE` disposition currently
	// returns Arrow data rather than `JSON_ARRAY`.
	//
	// Important: The formats `ARROW_STREAM` and `CSV` are supported only with
	// `EXTERNAL_LINKS` disposition. `JSON_ARRAY` is supported in `INLINE` and
	// `EXTERNAL_LINKS` disposition.
	//
	// When specifying `format=JSON_ARRAY`, result data will be formatted as an
	// array of arrays of values, where each value is either the *string
	// representation* of a value, or `null`. For example, the output of `SELECT
	// concat('id-', id) AS strCol, id AS intCol, null AS nullCol FROM range(3)`
	// would look like this:
	//
	// ``` [ [ "id-1", "1", null ], [ "id-2", "2", null ], [ "id-3", "3", null
	// ], ] ```
	//
	// When specifying `format=JSON_ARRAY` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result contains compact JSON with no indentation or
	// extra whitespace.
	//
	// When specifying `format=ARROW_STREAM` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result will be formatted as Apache Arrow Stream. See
	// the [Apache Arrow streaming format].
	//
	// When specifying `format=CSV` and `disposition=EXTERNAL_LINKS`, each chunk
	// in the result will be a CSV according to [RFC 4180] standard. All column
	// values will have *string representation* similar to the `JSON_ARRAY`
	// format, and `null` values will be encoded as “null”. Only the first
	// chunk in the result would contain a header row with column names. For
	// example, the output of `SELECT concat('id-', id) AS strCol, id AS intCol,
	// null as nullCol FROM range(3)` would look like this:
	//
	// ``` strCol,intCol,nullCol id-1,1,null id-2,2,null id-3,3,null ```
	//
	// [Apache Arrow streaming format]: https://arrow.apache.org/docs/format/Columnar.html#ipc-streaming-format
	// [RFC 4180]: https://www.rfc-editor.org/rfc/rfc4180
	Format types.String `tfsdk:"format"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling the statement status. When set to `CANCEL`, the statement
	// execution is canceled and the call returns with a `CANCELED` state.
	// Lakehouse//RT warehouses currently don't support `CANCEL`: the statement
	// keeps running as with `CONTINUE`.
	OnWaitTimeout types.String `tfsdk:"on_wait_timeout"`
	// A list of parameters to pass into a SQL statement containing parameter
	// markers. A parameter consists of a name, a value, and optionally a type.
	// To represent a NULL value, the `value` field may be omitted or set to
	// `null` explicitly. If the `type` field is omitted, the value is
	// interpreted as a string.
	//
	// If the type is given, parameters will be checked for type correctness
	// according to the given type. A value is correct if the provided string
	// can be converted to the requested type using the `cast` function. The
	// exact semantics are described in the section [`cast` function] of the SQL
	// language reference.
	//
	// For example, the following statement contains two parameters, `my_name`
	// and `my_date`:
	//
	// ``` SELECT * FROM my_table WHERE name = :my_name AND date = :my_date ```
	//
	// The parameters can be passed in the request body as follows:
	//
	// ` { ..., "statement": "SELECT * FROM my_table WHERE name = :my_name AND
	// date = :my_date", "parameters": [ { "name": "my_name", "value": "the
	// name" }, { "name": "my_date", "value": "2020-01-01", "type": "DATE" } ] }
	// `
	//
	// Currently, positional parameters denoted by a `?` marker are not
	// supported by the Databricks SQL Statement Execution API.
	//
	// Also see the section [Parameter markers] of the SQL language reference.
	//
	// [Parameter markers]: https://docs.databricks.com/sql/language-manual/sql-ref-parameter-marker.html
	// [`cast` function]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Parameters types.List `tfsdk:"parameters"`
	// An array of query tags to annotate a SQL statement. A query tag consists
	// of a non-empty key and, optionally, a value. To represent a NULL value,
	// either omit the `value` field or manually set it to `null` or white
	// space. Refer to the SQL language reference for the format specification
	// of query tags. There's no significance to the order of tags. Only one
	// value per key will be recorded. A sequence in excess of 20 query tags
	// will be coerced to 20. Example:
	//
	// { ..., "query_tags": [ { "key": "team", "value": "eng" }, { "key": "some
	// key only tag" } ] }
	QueryTags types.List `tfsdk:"query_tags"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit types.Int64 `tfsdk:"row_limit"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema types.String `tfsdk:"schema"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`. The maximum query text size is 16 MiB.
	Statement types.String `tfsdk:"statement"`
	// The time in seconds the call will wait for the statement's result set as
	// `Ns`, where `N` can be set to 0 or to a value between 5 and 50.
	//
	// When set to `0s`, the statement will execute in asynchronous mode and the
	// call will not wait for the execution to finish. In this case, the call
	// returns directly with `PENDING` state and a statement ID which can be
	// used for polling the statement status.
	//
	// When set between 5 and 50 seconds, the call will behave synchronously up
	// to this timeout and wait for the statement execution to finish. If the
	// execution finishes within this time, the call returns immediately with a
	// manifest and result data (or a `FAILED` state in case of an execution
	// error). If the statement takes longer to execute, `on_wait_timeout`
	// determines what should happen after the timeout is reached.
	WaitTimeout types.String `tfsdk:"wait_timeout"`
	// The ID of the SQL warehouse to run the statement on. Currently, only
	// [Lakehouse//RT] warehouses are supported. See also [What are SQL
	// warehouses?].
	//
	// [Lakehouse//RT]: https://docs.databricks.com/compute/sql-warehouse/real-time.html
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId types.String `tfsdk:"-"`
}

func (to *ExecuteStatementRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteStatementRequest_SdkV2) {
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
	if !from.QueryTags.IsNull() && !from.QueryTags.IsUnknown() && to.QueryTags.IsNull() && len(from.QueryTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryTags = from.QueryTags
	}
	if !from.QueryTags.IsNull() && !from.QueryTags.IsUnknown() {
		if toQueryTags, ok := to.GetQueryTags(ctx); ok {
			if fromQueryTags, ok := from.GetQueryTags(ctx); ok {
				// Recursively sync the fields of each QueryTags element by position.
				for i := range toQueryTags {
					if i < len(fromQueryTags) {
						toQueryTags[i].SyncFieldsDuringCreateOrUpdate(ctx, fromQueryTags[i])
					}
				}
				to.SetQueryTags(ctx, toQueryTags)
			}
		}
	}
}

func (to *ExecuteStatementRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExecuteStatementRequest_SdkV2) {
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
	if !from.QueryTags.IsNull() && !from.QueryTags.IsUnknown() && to.QueryTags.IsNull() && len(from.QueryTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryTags = from.QueryTags
	}
	if !from.QueryTags.IsNull() && !from.QueryTags.IsUnknown() {
		if toQueryTags, ok := to.GetQueryTags(ctx); ok {
			if fromQueryTags, ok := from.GetQueryTags(ctx); ok {
				for i := range toQueryTags {
					if i < len(fromQueryTags) {
						toQueryTags[i].SyncFieldsDuringRead(ctx, fromQueryTags[i])
					}
				}
				to.SetQueryTags(ctx, toQueryTags)
			}
		}
	}
}

func (m ExecuteStatementRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_limit"] = attrs["byte_limit"].SetOptional()
	attrs["catalog"] = attrs["catalog"].SetOptional()
	attrs["disposition"] = attrs["disposition"].SetOptional()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["on_wait_timeout"] = attrs["on_wait_timeout"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["query_tags"] = attrs["query_tags"].SetOptional()
	attrs["row_limit"] = attrs["row_limit"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["statement"] = attrs["statement"].SetRequired()
	attrs["wait_timeout"] = attrs["wait_timeout"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExecuteStatementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExecuteStatementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(StatementParameter_SdkV2{}),
		"query_tags": reflect.TypeOf(QueryTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteStatementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExecuteStatementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_limit":      m.ByteLimit,
			"catalog":         m.Catalog,
			"disposition":     m.Disposition,
			"format":          m.Format,
			"on_wait_timeout": m.OnWaitTimeout,
			"parameters":      m.Parameters,
			"query_tags":      m.QueryTags,
			"row_limit":       m.RowLimit,
			"schema":          m.Schema,
			"statement":       m.Statement,
			"wait_timeout":    m.WaitTimeout,
			"warehouse_id":    m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExecuteStatementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_limit":      types.Int64Type,
			"catalog":         types.StringType,
			"disposition":     types.StringType,
			"format":          types.StringType,
			"on_wait_timeout": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: StatementParameter_SdkV2{}.Type(ctx),
			},
			"query_tags": basetypes.ListType{
				ElemType: QueryTag_SdkV2{}.Type(ctx),
			},
			"row_limit":    types.Int64Type,
			"schema":       types.StringType,
			"statement":    types.StringType,
			"wait_timeout": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in ExecuteStatementRequest_SdkV2 as
// a slice of StatementParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteStatementRequest_SdkV2) GetParameters(ctx context.Context) ([]StatementParameter_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []StatementParameter_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ExecuteStatementRequest_SdkV2.
func (m *ExecuteStatementRequest_SdkV2) SetParameters(ctx context.Context, v []StatementParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

// GetQueryTags returns the value of the QueryTags field in ExecuteStatementRequest_SdkV2 as
// a slice of QueryTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteStatementRequest_SdkV2) GetQueryTags(ctx context.Context) ([]QueryTag_SdkV2, bool) {
	if m.QueryTags.IsNull() || m.QueryTags.IsUnknown() {
		return nil, false
	}
	var v []QueryTag_SdkV2
	d := m.QueryTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryTags sets the value of the QueryTags field in ExecuteStatementRequest_SdkV2.
func (m *ExecuteStatementRequest_SdkV2) SetQueryTags(ctx context.Context, v []QueryTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QueryTags = types.ListValueMust(t, vs)
}

// A short-lived cloud-storage link to one chunk of result data.
type ExternalLink_SdkV2 struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// Indicates the date-time that the given external link will expire and
	// become invalid, after which point a new `external_link` must be
	// requested.
	Expiration types.String `tfsdk:"expiration"`
	// A short-lived cloud-storage URL pointing to a chunk of result data,
	// hosted by an external service, with a short expiration time (<= 15
	// minutes). As this URL contains a temporary credential, it should be
	// considered sensitive and the client should not expose this URL in a log.
	ExternalLink types.String `tfsdk:"external_link"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders types.Map `tfsdk:"http_headers"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a get-result-chunk request for that index.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`. Lakehouse//RT
	// warehouses currently don't return this field; use `next_chunk_index`
	// instead.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

func (to *ExternalLink_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalLink_SdkV2) {
}

func (to *ExternalLink_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalLink_SdkV2) {
}

func (m ExternalLink_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_count"] = attrs["byte_count"].SetOptional()
	attrs["chunk_index"] = attrs["chunk_index"].SetOptional()
	attrs["expiration"] = attrs["expiration"].SetOptional()
	attrs["external_link"] = attrs["external_link"].SetOptional()
	attrs["http_headers"] = attrs["http_headers"].SetOptional()
	attrs["next_chunk_index"] = attrs["next_chunk_index"].SetOptional()
	attrs["next_chunk_internal_link"] = attrs["next_chunk_internal_link"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["row_offset"] = attrs["row_offset"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalLink.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExternalLink_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"http_headers": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalLink_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalLink_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":               m.ByteCount,
			"chunk_index":              m.ChunkIndex,
			"expiration":               m.Expiration,
			"external_link":            m.ExternalLink,
			"http_headers":             m.HttpHeaders,
			"next_chunk_index":         m.NextChunkIndex,
			"next_chunk_internal_link": m.NextChunkInternalLink,
			"row_count":                m.RowCount,
			"row_offset":               m.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalLink_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_count":    types.Int64Type,
			"chunk_index":   types.Int64Type,
			"expiration":    types.StringType,
			"external_link": types.StringType,
			"http_headers": basetypes.MapType{
				ElemType: types.StringType,
			},
			"next_chunk_index":         types.Int64Type,
			"next_chunk_internal_link": types.StringType,
			"row_count":                types.Int64Type,
			"row_offset":               types.Int64Type,
		},
	}
}

// GetHttpHeaders returns the value of the HttpHeaders field in ExternalLink_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalLink_SdkV2) GetHttpHeaders(ctx context.Context) (map[string]types.String, bool) {
	if m.HttpHeaders.IsNull() || m.HttpHeaders.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.HttpHeaders.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHttpHeaders sets the value of the HttpHeaders field in ExternalLink_SdkV2.
func (m *ExternalLink_SdkV2) SetHttpHeaders(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["http_headers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.HttpHeaders = types.MapValueMust(t, vs)
}

type GetResultDataRequest_SdkV2 struct {
	// The zero-based index of the result chunk to fetch.
	ChunkIndex types.Int64 `tfsdk:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *GetResultDataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetResultDataRequest_SdkV2) {
}

func (to *GetResultDataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetResultDataRequest_SdkV2) {
}

func (m GetResultDataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()
	attrs["statement_id"] = attrs["statement_id"].SetRequired()
	attrs["chunk_index"] = attrs["chunk_index"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetResultDataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetResultDataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetResultDataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetResultDataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunk_index":  m.ChunkIndex,
			"statement_id": m.StatementId,
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetResultDataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunk_index":  types.Int64Type,
			"statement_id": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

type GetStatementResultRequest_SdkV2 struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *GetStatementResultRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatementResultRequest_SdkV2) {
}

func (to *GetStatementResultRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetStatementResultRequest_SdkV2) {
}

func (m GetStatementResultRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()
	attrs["statement_id"] = attrs["statement_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatementResultRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetStatementResultRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementResultRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetStatementResultRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": m.StatementId,
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatementResultRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// * A query execution can be annotated with an optional key-value pair to allow
// users to attribute the executions by key and optional value to filter by.
// QueryTag is the user-facing representation.
type QueryTag_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (to *QueryTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryTag_SdkV2) {
}

func (to *QueryTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryTag_SdkV2) {
}

func (m QueryTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryTag_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Contains the result data of a single chunk when using `INLINE` disposition.
// When using `EXTERNAL_LINKS` disposition, the array `external_links` is used
// instead to provide short-lived cloud-storage URLs to the result data in cloud
// storage. Exactly one of these alternatives is used. (While the
// `external_links` array prepares the API to return multiple links in a single
// response, currently only a single link is returned.)
type ResultData_SdkV2 struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray types.List `tfsdk:"data_array"`
	// Short-lived cloud-storage links to the result data, used with the
	// `EXTERNAL_LINKS` disposition.
	ExternalLinks types.List `tfsdk:"external_links"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a get-result-chunk request for that index.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`. Lakehouse//RT
	// warehouses currently don't return this field; use `next_chunk_index`
	// instead.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

func (to *ResultData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultData_SdkV2) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
	if !from.ExternalLinks.IsNull() && !from.ExternalLinks.IsUnknown() && to.ExternalLinks.IsNull() && len(from.ExternalLinks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExternalLinks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExternalLinks = from.ExternalLinks
	}
	if !from.ExternalLinks.IsNull() && !from.ExternalLinks.IsUnknown() {
		if toExternalLinks, ok := to.GetExternalLinks(ctx); ok {
			if fromExternalLinks, ok := from.GetExternalLinks(ctx); ok {
				// Recursively sync the fields of each ExternalLinks element by position.
				for i := range toExternalLinks {
					if i < len(fromExternalLinks) {
						toExternalLinks[i].SyncFieldsDuringCreateOrUpdate(ctx, fromExternalLinks[i])
					}
				}
				to.SetExternalLinks(ctx, toExternalLinks)
			}
		}
	}
}

func (to *ResultData_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultData_SdkV2) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
	if !from.ExternalLinks.IsNull() && !from.ExternalLinks.IsUnknown() && to.ExternalLinks.IsNull() && len(from.ExternalLinks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExternalLinks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExternalLinks = from.ExternalLinks
	}
	if !from.ExternalLinks.IsNull() && !from.ExternalLinks.IsUnknown() {
		if toExternalLinks, ok := to.GetExternalLinks(ctx); ok {
			if fromExternalLinks, ok := from.GetExternalLinks(ctx); ok {
				for i := range toExternalLinks {
					if i < len(fromExternalLinks) {
						toExternalLinks[i].SyncFieldsDuringRead(ctx, fromExternalLinks[i])
					}
				}
				to.SetExternalLinks(ctx, toExternalLinks)
			}
		}
	}
}

func (m ResultData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["byte_count"] = attrs["byte_count"].SetOptional()
	attrs["chunk_index"] = attrs["chunk_index"].SetOptional()
	attrs["data_array"] = attrs["data_array"].SetOptional()
	attrs["external_links"] = attrs["external_links"].SetOptional()
	attrs["next_chunk_index"] = attrs["next_chunk_index"].SetOptional()
	attrs["next_chunk_internal_link"] = attrs["next_chunk_internal_link"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()
	attrs["row_offset"] = attrs["row_offset"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResultData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array":     reflect.TypeOf(types.Object{}),
		"external_links": reflect.TypeOf(ExternalLink_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"byte_count":               m.ByteCount,
			"chunk_index":              m.ChunkIndex,
			"data_array":               m.DataArray,
			"external_links":           m.ExternalLinks,
			"next_chunk_index":         m.NextChunkIndex,
			"next_chunk_internal_link": m.NextChunkInternalLink,
			"row_count":                m.RowCount,
			"row_offset":               m.RowOffset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultData_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_count":  types.Int64Type,
			"chunk_index": types.Int64Type,
			"data_array": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"external_links": basetypes.ListType{
				ElemType: ExternalLink_SdkV2{}.Type(ctx),
			},
			"next_chunk_index":         types.Int64Type,
			"next_chunk_internal_link": types.StringType,
			"row_count":                types.Int64Type,
			"row_offset":               types.Int64Type,
		},
	}
}

// GetDataArray returns the value of the DataArray field in ResultData_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData_SdkV2) GetDataArray(ctx context.Context) ([]types.Object, bool) {
	if m.DataArray.IsNull() || m.DataArray.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.DataArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataArray sets the value of the DataArray field in ResultData_SdkV2.
func (m *ResultData_SdkV2) SetDataArray(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataArray = types.ListValueMust(t, vs)
}

// GetExternalLinks returns the value of the ExternalLinks field in ResultData_SdkV2 as
// a slice of ExternalLink_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData_SdkV2) GetExternalLinks(ctx context.Context) ([]ExternalLink_SdkV2, bool) {
	if m.ExternalLinks.IsNull() || m.ExternalLinks.IsUnknown() {
		return nil, false
	}
	var v []ExternalLink_SdkV2
	d := m.ExternalLinks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalLinks sets the value of the ExternalLinks field in ResultData_SdkV2.
func (m *ResultData_SdkV2) SetExternalLinks(ctx context.Context, v []ExternalLink_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["external_links"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExternalLinks = types.ListValueMust(t, vs)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest_SdkV2 struct {
	// Array of result set chunk metadata.
	Chunks types.List `tfsdk:"chunks"`
	// The format of the result data, as requested with `format`.
	Format types.String `tfsdk:"format"`
	// The schema of the result set: the ordered list of column descriptions.
	Schema types.List `tfsdk:"schema"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount types.Int64 `tfsdk:"total_byte_count"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount types.Int64 `tfsdk:"total_chunk_count"`
	// The total number of rows in the result set.
	TotalRowCount types.Int64 `tfsdk:"total_row_count"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated types.Bool `tfsdk:"truncated"`
}

func (to *ResultManifest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest_SdkV2) {
	if !from.Chunks.IsNull() && !from.Chunks.IsUnknown() && to.Chunks.IsNull() && len(from.Chunks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Chunks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Chunks = from.Chunks
	}
	if !from.Chunks.IsNull() && !from.Chunks.IsUnknown() {
		if toChunks, ok := to.GetChunks(ctx); ok {
			if fromChunks, ok := from.GetChunks(ctx); ok {
				// Recursively sync the fields of each Chunks element by position.
				for i := range toChunks {
					if i < len(fromChunks) {
						toChunks[i].SyncFieldsDuringCreateOrUpdate(ctx, fromChunks[i])
					}
				}
				to.SetChunks(ctx, toChunks)
			}
		}
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() {
		if toSchema, ok := to.GetSchema(ctx); ok {
			if fromSchema, ok := from.GetSchema(ctx); ok {
				// Recursively sync the fields of Schema
				toSchema.SyncFieldsDuringCreateOrUpdate(ctx, fromSchema)
				to.SetSchema(ctx, toSchema)
			}
		}
	}
}

func (to *ResultManifest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultManifest_SdkV2) {
	if !from.Chunks.IsNull() && !from.Chunks.IsUnknown() && to.Chunks.IsNull() && len(from.Chunks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Chunks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Chunks = from.Chunks
	}
	if !from.Chunks.IsNull() && !from.Chunks.IsUnknown() {
		if toChunks, ok := to.GetChunks(ctx); ok {
			if fromChunks, ok := from.GetChunks(ctx); ok {
				for i := range toChunks {
					if i < len(fromChunks) {
						toChunks[i].SyncFieldsDuringRead(ctx, fromChunks[i])
					}
				}
				to.SetChunks(ctx, toChunks)
			}
		}
	}
	if !from.Schema.IsNull() && !from.Schema.IsUnknown() {
		if toSchema, ok := to.GetSchema(ctx); ok {
			if fromSchema, ok := from.GetSchema(ctx); ok {
				toSchema.SyncFieldsDuringRead(ctx, fromSchema)
				to.SetSchema(ctx, toSchema)
			}
		}
	}
}

func (m ResultManifest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["chunks"] = attrs["chunks"].SetOptional()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["schema"] = attrs["schema"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["total_byte_count"] = attrs["total_byte_count"].SetOptional()
	attrs["total_chunk_count"] = attrs["total_chunk_count"].SetOptional()
	attrs["total_row_count"] = attrs["total_row_count"].SetOptional()
	attrs["truncated"] = attrs["truncated"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResultManifest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"chunks": reflect.TypeOf(BaseChunkInfo_SdkV2{}),
		"schema": reflect.TypeOf(Schema_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultManifest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunks":            m.Chunks,
			"format":            m.Format,
			"schema":            m.Schema,
			"total_byte_count":  m.TotalByteCount,
			"total_chunk_count": m.TotalChunkCount,
			"total_row_count":   m.TotalRowCount,
			"truncated":         m.Truncated,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultManifest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunks": basetypes.ListType{
				ElemType: BaseChunkInfo_SdkV2{}.Type(ctx),
			},
			"format": types.StringType,
			"schema": basetypes.ListType{
				ElemType: Schema_SdkV2{}.Type(ctx),
			},
			"total_byte_count":  types.Int64Type,
			"total_chunk_count": types.Int64Type,
			"total_row_count":   types.Int64Type,
			"truncated":         types.BoolType,
		},
	}
}

// GetChunks returns the value of the Chunks field in ResultManifest_SdkV2 as
// a slice of BaseChunkInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest_SdkV2) GetChunks(ctx context.Context) ([]BaseChunkInfo_SdkV2, bool) {
	if m.Chunks.IsNull() || m.Chunks.IsUnknown() {
		return nil, false
	}
	var v []BaseChunkInfo_SdkV2
	d := m.Chunks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChunks sets the value of the Chunks field in ResultManifest_SdkV2.
func (m *ResultManifest_SdkV2) SetChunks(ctx context.Context, v []BaseChunkInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["chunks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Chunks = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in ResultManifest_SdkV2 as
// a Schema_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest_SdkV2) GetSchema(ctx context.Context) (Schema_SdkV2, bool) {
	var e Schema_SdkV2
	if m.Schema.IsNull() || m.Schema.IsUnknown() {
		return e, false
	}
	var v []Schema_SdkV2
	d := m.Schema.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchema sets the value of the Schema field in ResultManifest_SdkV2.
func (m *ResultManifest_SdkV2) SetSchema(ctx context.Context, v Schema_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["schema"]
	m.Schema = types.ListValueMust(t, vs)
}

// The schema is an ordered list of column descriptions.
type Schema_SdkV2 struct {
	// The number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count"`
	// The column descriptions, in result set order.
	Columns types.List `tfsdk:"columns"`
}

func (to *Schema_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Schema_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() {
		if toColumns, ok := to.GetColumns(ctx); ok {
			if fromColumns, ok := from.GetColumns(ctx); ok {
				// Recursively sync the fields of each Columns element by position.
				for i := range toColumns {
					if i < len(fromColumns) {
						toColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromColumns[i])
					}
				}
				to.SetColumns(ctx, toColumns)
			}
		}
	}
}

func (to *Schema_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Schema_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() {
		if toColumns, ok := to.GetColumns(ctx); ok {
			if fromColumns, ok := from.GetColumns(ctx); ok {
				for i := range toColumns {
					if i < len(fromColumns) {
						toColumns[i].SyncFieldsDuringRead(ctx, fromColumns[i])
					}
				}
				to.SetColumns(ctx, toColumns)
			}
		}
	}
}

func (m Schema_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column_count"] = attrs["column_count"].SetOptional()
	attrs["columns"] = attrs["columns"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Schema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Schema_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Schema_SdkV2
// only implements ToObjectValue() and Type().
func (m Schema_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": m.ColumnCount,
			"columns":      m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Schema_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in Schema_SdkV2 as
// a slice of ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Schema_SdkV2) GetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in Schema_SdkV2.
func (m *Schema_SdkV2) SetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// The error that caused a statement execution to fail.
type ServiceError_SdkV2 struct {
	// The error code that classifies the failure.
	ErrorCode types.String `tfsdk:"error_code"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message"`
}

func (to *ServiceError_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServiceError_SdkV2) {
}

func (to *ServiceError_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServiceError_SdkV2) {
}

func (m ServiceError_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error_code"] = attrs["error_code"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServiceError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServiceError_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServiceError_SdkV2
// only implements ToObjectValue() and Type().
func (m ServiceError_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_code": m.ErrorCode,
			"message":    m.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServiceError_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_code": types.StringType,
			"message":    types.StringType,
		},
	}
}

// A named parameter to substitute for a parameter marker in the statement.
type StatementParameter_SdkV2 struct {
	// The name of a parameter marker to be substituted in the statement.
	Name types.String `tfsdk:"name"`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type_ types.String `tfsdk:"type"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value types.String `tfsdk:"value"`
}

func (to *StatementParameter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementParameter_SdkV2) {
}

func (to *StatementParameter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StatementParameter_SdkV2) {
}

func (m StatementParameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StatementParameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementParameter_SdkV2
// only implements ToObjectValue() and Type().
func (m StatementParameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  m.Name,
			"type":  m.Type_,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementParameter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":  types.StringType,
			"type":  types.StringType,
			"value": types.StringType,
		},
	}
}

// Statement Execution (Lakehouse//RT) is version 2 of the Statement Execution
// API. It is optimized for Lakehouse//RT warehouses, with lower latency and
// higher concurrency, and takes `warehouse_id` as a path parameter on every
// call. It currently supports Lakehouse//RT warehouses only. The Statement
// Execution API (version 1) works with all SQL warehouse types, including
// Lakehouse//RT.
type StatementResponse_SdkV2 struct {
	// The schema and metadata of the result set. Present when `status.state` is
	// `SUCCEEDED`.
	Manifest types.List `tfsdk:"manifest"`
	// The first chunk of result data. Present when `status.state` is
	// `SUCCEEDED`.
	Result types.List `tfsdk:"result"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id"`
	// The execution state of the statement and, if it failed, the error
	// details.
	Status types.List `tfsdk:"status"`
}

func (to *StatementResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementResponse_SdkV2) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				// Recursively sync the fields of Manifest
				toManifest.SyncFieldsDuringCreateOrUpdate(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
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

func (to *StatementResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StatementResponse_SdkV2) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				toManifest.SyncFieldsDuringRead(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
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

func (m StatementResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["manifest"] = attrs["manifest"].SetOptional()
	attrs["manifest"] = attrs["manifest"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["result"] = attrs["result"].SetOptional()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["statement_id"] = attrs["statement_id"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StatementResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest_SdkV2{}),
		"result":   reflect.TypeOf(ResultData_SdkV2{}),
		"status":   reflect.TypeOf(StatementStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m StatementResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"manifest":     m.Manifest,
			"result":       m.Result,
			"statement_id": m.StatementId,
			"status":       m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"manifest": basetypes.ListType{
				ElemType: ResultManifest_SdkV2{}.Type(ctx),
			},
			"result": basetypes.ListType{
				ElemType: ResultData_SdkV2{}.Type(ctx),
			},
			"statement_id": types.StringType,
			"status": basetypes.ListType{
				ElemType: StatementStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetManifest returns the value of the Manifest field in StatementResponse_SdkV2 as
// a ResultManifest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse_SdkV2) GetManifest(ctx context.Context) (ResultManifest_SdkV2, bool) {
	var e ResultManifest_SdkV2
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v []ResultManifest_SdkV2
	d := m.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in StatementResponse_SdkV2.
func (m *StatementResponse_SdkV2) SetManifest(ctx context.Context, v ResultManifest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	m.Manifest = types.ListValueMust(t, vs)
}

// GetResult returns the value of the Result field in StatementResponse_SdkV2 as
// a ResultData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse_SdkV2) GetResult(ctx context.Context) (ResultData_SdkV2, bool) {
	var e ResultData_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []ResultData_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in StatementResponse_SdkV2.
func (m *StatementResponse_SdkV2) SetResult(ctx context.Context, v ResultData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in StatementResponse_SdkV2 as
// a StatementStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse_SdkV2) GetStatus(ctx context.Context) (StatementStatus_SdkV2, bool) {
	var e StatementStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []StatementStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in StatementResponse_SdkV2.
func (m *StatementResponse_SdkV2) SetStatus(ctx context.Context, v StatementStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus_SdkV2 struct {
	// The error details. Only populated when `state` is `FAILED`.
	Error types.List `tfsdk:"error"`
	// SQLSTATE error code returned when the statement execution fails. Only
	// populated when the statement status is `FAILED`.
	SqlState types.String `tfsdk:"sql_state"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accompanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State types.String `tfsdk:"state"`
}

func (to *StatementStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementStatus_SdkV2) {
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

func (to *StatementStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StatementStatus_SdkV2) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m StatementStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
	attrs["error"] = attrs["error"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sql_state"] = attrs["sql_state"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StatementStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(ServiceError_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m StatementStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error":     m.Error,
			"sql_state": m.SqlState,
			"state":     m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": basetypes.ListType{
				ElemType: ServiceError_SdkV2{}.Type(ctx),
			},
			"sql_state": types.StringType,
			"state":     types.StringType,
		},
	}
}

// GetError returns the value of the Error field in StatementStatus_SdkV2 as
// a ServiceError_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementStatus_SdkV2) GetError(ctx context.Context) (ServiceError_SdkV2, bool) {
	var e ServiceError_SdkV2
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v []ServiceError_SdkV2
	d := m.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in StatementStatus_SdkV2.
func (m *StatementStatus_SdkV2) SetError(ctx context.Context, v ServiceError_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	m.Error = types.ListValueMust(t, vs)
}
