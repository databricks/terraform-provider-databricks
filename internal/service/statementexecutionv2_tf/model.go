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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Metadata of one result set chunk, as listed in the result manifest.
type BaseChunkInfo struct {
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

func (to *BaseChunkInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BaseChunkInfo) {
}

func (to *BaseChunkInfo) SyncFieldsDuringRead(ctx context.Context, from BaseChunkInfo) {
}

func (m BaseChunkInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BaseChunkInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BaseChunkInfo
// only implements ToObjectValue() and Type().
func (m BaseChunkInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m BaseChunkInfo) Type(ctx context.Context) attr.Type {
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
type CancelStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *CancelStatementRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelStatementRequest) {
}

func (to *CancelStatementRequest) SyncFieldsDuringRead(ctx context.Context, from CancelStatementRequest) {
}

func (m CancelStatementRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CancelStatementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelStatementRequest
// only implements ToObjectValue() and Type().
func (m CancelStatementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": m.StatementId,
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelStatementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_id": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// Empty response to a cancel request. Receiving it means the cancel request was
// accepted.
type CancelStatementResponse struct {
}

func (to *CancelStatementResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelStatementResponse) {
}

func (to *CancelStatementResponse) SyncFieldsDuringRead(ctx context.Context, from CancelStatementResponse) {
}

func (m CancelStatementResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelStatementResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelStatementResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelStatementResponse
// only implements ToObjectValue() and Type().
func (m CancelStatementResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m CancelStatementResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Describes one column of the result set: its name, SQL type, and position.
type ColumnInfo struct {
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

func (to *ColumnInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo) {
}

func (to *ColumnInfo) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo) {
}

func (m ColumnInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ColumnInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo
// only implements ToObjectValue() and Type().
func (m ColumnInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ColumnInfo) Type(ctx context.Context) attr.Type {
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
type ExecuteStatementRequest struct {
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
	// ``` { ..., "statement": "SELECT * FROM my_table WHERE name = :my_name AND
	// date = :my_date", "parameters": [ { "name": "my_name", "value": "the
	// name" }, { "name": "my_date", "value": "2020-01-01", "type": "DATE" } ] }
	// ```
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

func (to *ExecuteStatementRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExecuteStatementRequest) {
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

func (to *ExecuteStatementRequest) SyncFieldsDuringRead(ctx context.Context, from ExecuteStatementRequest) {
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

func (m ExecuteStatementRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExecuteStatementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(StatementParameter{}),
		"query_tags": reflect.TypeOf(QueryTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExecuteStatementRequest
// only implements ToObjectValue() and Type().
func (m ExecuteStatementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExecuteStatementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"byte_limit":      types.Int64Type,
			"catalog":         types.StringType,
			"disposition":     types.StringType,
			"format":          types.StringType,
			"on_wait_timeout": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: StatementParameter{}.Type(ctx),
			},
			"query_tags": basetypes.ListType{
				ElemType: QueryTag{}.Type(ctx),
			},
			"row_limit":    types.Int64Type,
			"schema":       types.StringType,
			"statement":    types.StringType,
			"wait_timeout": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

// GetParameters returns the value of the Parameters field in ExecuteStatementRequest as
// a slice of StatementParameter values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteStatementRequest) GetParameters(ctx context.Context) ([]StatementParameter, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []StatementParameter
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ExecuteStatementRequest.
func (m *ExecuteStatementRequest) SetParameters(ctx context.Context, v []StatementParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

// GetQueryTags returns the value of the QueryTags field in ExecuteStatementRequest as
// a slice of QueryTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExecuteStatementRequest) GetQueryTags(ctx context.Context) ([]QueryTag, bool) {
	if m.QueryTags.IsNull() || m.QueryTags.IsUnknown() {
		return nil, false
	}
	var v []QueryTag
	d := m.QueryTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryTags sets the value of the QueryTags field in ExecuteStatementRequest.
func (m *ExecuteStatementRequest) SetQueryTags(ctx context.Context, v []QueryTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QueryTags = types.ListValueMust(t, vs)
}

// A short-lived cloud-storage link to one chunk of result data.
type ExternalLink struct {
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

func (to *ExternalLink) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalLink) {
	if !from.HttpHeaders.IsNull() && !from.HttpHeaders.IsUnknown() && to.HttpHeaders.IsNull() && len(from.HttpHeaders.Elements()) == 0 {
		// The default representation of an empty map for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty map for HttpHeaders, and the deserialized field value is Null,
		// set the resulting resource state to the empty map to match the planned value.
		to.HttpHeaders = from.HttpHeaders
	}
}

func (to *ExternalLink) SyncFieldsDuringRead(ctx context.Context, from ExternalLink) {
	if !from.HttpHeaders.IsNull() && !from.HttpHeaders.IsUnknown() && to.HttpHeaders.IsNull() && len(from.HttpHeaders.Elements()) == 0 {
		// The default representation of an empty map for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty map for HttpHeaders, and the deserialized field value is Null,
		// set the resulting resource state to the empty map to match the planned value.
		to.HttpHeaders = from.HttpHeaders
	}
}

func (m ExternalLink) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExternalLink) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"http_headers": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalLink
// only implements ToObjectValue() and Type().
func (m ExternalLink) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExternalLink) Type(ctx context.Context) attr.Type {
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

// GetHttpHeaders returns the value of the HttpHeaders field in ExternalLink as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalLink) GetHttpHeaders(ctx context.Context) (map[string]types.String, bool) {
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

// SetHttpHeaders sets the value of the HttpHeaders field in ExternalLink.
func (m *ExternalLink) SetHttpHeaders(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["http_headers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.HttpHeaders = types.MapValueMust(t, vs)
}

type GetResultDataRequest struct {
	// The zero-based index of the result chunk to fetch.
	ChunkIndex types.Int64 `tfsdk:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *GetResultDataRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetResultDataRequest) {
}

func (to *GetResultDataRequest) SyncFieldsDuringRead(ctx context.Context, from GetResultDataRequest) {
}

func (m GetResultDataRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetResultDataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetResultDataRequest
// only implements ToObjectValue() and Type().
func (m GetResultDataRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"chunk_index":  m.ChunkIndex,
			"statement_id": m.StatementId,
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetResultDataRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunk_index":  types.Int64Type,
			"statement_id": types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

type GetStatementResultRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId types.String `tfsdk:"-"`
}

func (to *GetStatementResultRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatementResultRequest) {
}

func (to *GetStatementResultRequest) SyncFieldsDuringRead(ctx context.Context, from GetStatementResultRequest) {
}

func (m GetStatementResultRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetStatementResultRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatementResultRequest
// only implements ToObjectValue() and Type().
func (m GetStatementResultRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"statement_id": m.StatementId,
			"warehouse_id": m.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatementResultRequest) Type(ctx context.Context) attr.Type {
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
type QueryTag struct {
	// The key of the tag.
	Key types.String `tfsdk:"key"`
	// The value of the tag. Key-only tags have no value.
	Value types.String `tfsdk:"value"`
}

func (to *QueryTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryTag) {
}

func (to *QueryTag) SyncFieldsDuringRead(ctx context.Context, from QueryTag) {
}

func (m QueryTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryTag
// only implements ToObjectValue() and Type().
func (m QueryTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryTag) Type(ctx context.Context) attr.Type {
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
type ResultData struct {
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

func (to *ResultData) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultData) {
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

func (to *ResultData) SyncFieldsDuringRead(ctx context.Context, from ResultData) {
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

func (m ResultData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResultData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array":     reflect.TypeOf(types.Object{}),
		"external_links": reflect.TypeOf(ExternalLink{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData
// only implements ToObjectValue() and Type().
func (m ResultData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ResultData) Type(ctx context.Context) attr.Type {
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
				ElemType: ExternalLink{}.Type(ctx),
			},
			"next_chunk_index":         types.Int64Type,
			"next_chunk_internal_link": types.StringType,
			"row_count":                types.Int64Type,
			"row_offset":               types.Int64Type,
		},
	}
}

// GetDataArray returns the value of the DataArray field in ResultData as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData) GetDataArray(ctx context.Context) ([]types.Object, bool) {
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

// SetDataArray sets the value of the DataArray field in ResultData.
func (m *ResultData) SetDataArray(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataArray = types.ListValueMust(t, vs)
}

// GetExternalLinks returns the value of the ExternalLinks field in ResultData as
// a slice of ExternalLink values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData) GetExternalLinks(ctx context.Context) ([]ExternalLink, bool) {
	if m.ExternalLinks.IsNull() || m.ExternalLinks.IsUnknown() {
		return nil, false
	}
	var v []ExternalLink
	d := m.ExternalLinks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalLinks sets the value of the ExternalLinks field in ResultData.
func (m *ResultData) SetExternalLinks(ctx context.Context, v []ExternalLink) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["external_links"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExternalLinks = types.ListValueMust(t, vs)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	Chunks types.List `tfsdk:"chunks"`
	// The format of the result data, as requested with `format`.
	Format types.String `tfsdk:"format"`
	// The schema of the result set: the ordered list of column descriptions.
	Schema types.Object `tfsdk:"schema"`
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

func (to *ResultManifest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest) {
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

func (to *ResultManifest) SyncFieldsDuringRead(ctx context.Context, from ResultManifest) {
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

func (m ResultManifest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["chunks"] = attrs["chunks"].SetOptional()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
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
func (m ResultManifest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"chunks": reflect.TypeOf(BaseChunkInfo{}),
		"schema": reflect.TypeOf(Schema{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest
// only implements ToObjectValue() and Type().
func (m ResultManifest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ResultManifest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"chunks": basetypes.ListType{
				ElemType: BaseChunkInfo{}.Type(ctx),
			},
			"format":            types.StringType,
			"schema":            Schema{}.Type(ctx),
			"total_byte_count":  types.Int64Type,
			"total_chunk_count": types.Int64Type,
			"total_row_count":   types.Int64Type,
			"truncated":         types.BoolType,
		},
	}
}

// GetChunks returns the value of the Chunks field in ResultManifest as
// a slice of BaseChunkInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest) GetChunks(ctx context.Context) ([]BaseChunkInfo, bool) {
	if m.Chunks.IsNull() || m.Chunks.IsUnknown() {
		return nil, false
	}
	var v []BaseChunkInfo
	d := m.Chunks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChunks sets the value of the Chunks field in ResultManifest.
func (m *ResultManifest) SetChunks(ctx context.Context, v []BaseChunkInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["chunks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Chunks = types.ListValueMust(t, vs)
}

// GetSchema returns the value of the Schema field in ResultManifest as
// a Schema value.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest) GetSchema(ctx context.Context) (Schema, bool) {
	var e Schema
	if m.Schema.IsNull() || m.Schema.IsUnknown() {
		return e, false
	}
	var v Schema
	d := m.Schema.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchema sets the value of the Schema field in ResultManifest.
func (m *ResultManifest) SetSchema(ctx context.Context, v Schema) {
	vs := v.ToObjectValue(ctx)
	m.Schema = vs
}

// The schema is an ordered list of column descriptions.
type Schema struct {
	// The number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count"`
	// The column descriptions, in result set order.
	Columns types.List `tfsdk:"columns"`
}

func (to *Schema) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Schema) {
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

func (to *Schema) SyncFieldsDuringRead(ctx context.Context, from Schema) {
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

func (m Schema) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Schema) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Schema
// only implements ToObjectValue() and Type().
func (m Schema) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": m.ColumnCount,
			"columns":      m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Schema) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in Schema as
// a slice of ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *Schema) GetColumns(ctx context.Context) ([]ColumnInfo, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in Schema.
func (m *Schema) SetColumns(ctx context.Context, v []ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// The error that caused a statement execution to fail.
type ServiceError struct {
	// The error code that classifies the failure.
	ErrorCode types.String `tfsdk:"error_code"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message"`
}

func (to *ServiceError) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServiceError) {
}

func (to *ServiceError) SyncFieldsDuringRead(ctx context.Context, from ServiceError) {
}

func (m ServiceError) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ServiceError) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServiceError
// only implements ToObjectValue() and Type().
func (m ServiceError) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_code": m.ErrorCode,
			"message":    m.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServiceError) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_code": types.StringType,
			"message":    types.StringType,
		},
	}
}

// A named parameter to substitute for a parameter marker in the statement.
type StatementParameter struct {
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

func (to *StatementParameter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementParameter) {
}

func (to *StatementParameter) SyncFieldsDuringRead(ctx context.Context, from StatementParameter) {
}

func (m StatementParameter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StatementParameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementParameter
// only implements ToObjectValue() and Type().
func (m StatementParameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  m.Name,
			"type":  m.Type_,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementParameter) Type(ctx context.Context) attr.Type {
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
type StatementResponse struct {
	// The schema and metadata of the result set. Present when `status.state` is
	// `SUCCEEDED`.
	Manifest types.Object `tfsdk:"manifest"`
	// The first chunk of result data. Present when `status.state` is
	// `SUCCEEDED`.
	Result types.Object `tfsdk:"result"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id"`
	// The execution state of the statement and, if it failed, the error
	// details.
	Status types.Object `tfsdk:"status"`
}

func (to *StatementResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementResponse) {
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

func (to *StatementResponse) SyncFieldsDuringRead(ctx context.Context, from StatementResponse) {
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

func (m StatementResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["manifest"] = attrs["manifest"].SetOptional()
	attrs["result"] = attrs["result"].SetOptional()
	attrs["statement_id"] = attrs["statement_id"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StatementResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StatementResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest{}),
		"result":   reflect.TypeOf(ResultData{}),
		"status":   reflect.TypeOf(StatementStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementResponse
// only implements ToObjectValue() and Type().
func (m StatementResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m StatementResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"manifest":     ResultManifest{}.Type(ctx),
			"result":       ResultData{}.Type(ctx),
			"statement_id": types.StringType,
			"status":       StatementStatus{}.Type(ctx),
		},
	}
}

// GetManifest returns the value of the Manifest field in StatementResponse as
// a ResultManifest value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse) GetManifest(ctx context.Context) (ResultManifest, bool) {
	var e ResultManifest
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v ResultManifest
	d := m.Manifest.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManifest sets the value of the Manifest field in StatementResponse.
func (m *StatementResponse) SetManifest(ctx context.Context, v ResultManifest) {
	vs := v.ToObjectValue(ctx)
	m.Manifest = vs
}

// GetResult returns the value of the Result field in StatementResponse as
// a ResultData value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse) GetResult(ctx context.Context) (ResultData, bool) {
	var e ResultData
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v ResultData
	d := m.Result.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResult sets the value of the Result field in StatementResponse.
func (m *StatementResponse) SetResult(ctx context.Context, v ResultData) {
	vs := v.ToObjectValue(ctx)
	m.Result = vs
}

// GetStatus returns the value of the Status field in StatementResponse as
// a StatementStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementResponse) GetStatus(ctx context.Context) (StatementStatus, bool) {
	var e StatementStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v StatementStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in StatementResponse.
func (m *StatementResponse) SetStatus(ctx context.Context, v StatementStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {
	// The error details. Only populated when `state` is `FAILED`.
	Error types.Object `tfsdk:"error"`
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

func (to *StatementStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StatementStatus) {
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

func (to *StatementStatus) SyncFieldsDuringRead(ctx context.Context, from StatementStatus) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m StatementStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error"] = attrs["error"].SetOptional()
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
func (m StatementStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(ServiceError{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StatementStatus
// only implements ToObjectValue() and Type().
func (m StatementStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error":     m.Error,
			"sql_state": m.SqlState,
			"state":     m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StatementStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error":     ServiceError{}.Type(ctx),
			"sql_state": types.StringType,
			"state":     types.StringType,
		},
	}
}

// GetError returns the value of the Error field in StatementStatus as
// a ServiceError value.
// If the field is unknown or null, the boolean return value is false.
func (m *StatementStatus) GetError(ctx context.Context) (ServiceError, bool) {
	var e ServiceError
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v ServiceError
	d := m.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in StatementStatus.
func (m *StatementStatus) SetError(ctx context.Context, v ServiceError) {
	vs := v.ToObjectValue(ctx)
	m.Error = vs
}
