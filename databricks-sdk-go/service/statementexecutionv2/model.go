// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package statementexecutionv2

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// Metadata of one result set chunk, as listed in the result manifest.
type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BaseChunkInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BaseChunkInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to cancel the statement given in the path on the warehouse given in
// the path.
type CancelStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId string `json:"-" url:"-"`
}

func (s *CancelStatementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Empty response to a cancel request. Receiving it means the cancel request was
// accepted.
type CancelStatementResponse struct {
}

func (s *CancelStatementResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Describes one column of the result set: its name, SQL type, and position.
type ColumnInfo struct {
	// The name of the column.
	Name string `json:"name,omitempty"`
	// The ordinal position of the column (starting at position 0).
	Position int `json:"position,omitempty"`
	// The format of the interval type.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision int `json:"type_precision,omitempty"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale int `json:"type_scale,omitempty"`
	// The full SQL type specification.
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The name of the base data type. This doesn't include details for complex
// types such as STRUCT, MAP or ARRAY.
type ColumnTypeName string

const ColumnTypeNameArray ColumnTypeName = `ARRAY`

const ColumnTypeNameBinary ColumnTypeName = `BINARY`

const ColumnTypeNameBoolean ColumnTypeName = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeName = `BYTE`

const ColumnTypeNameChar ColumnTypeName = `CHAR`

const ColumnTypeNameDate ColumnTypeName = `DATE`

const ColumnTypeNameDecimal ColumnTypeName = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeName = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeName = `FLOAT`

const ColumnTypeNameInt ColumnTypeName = `INT`

const ColumnTypeNameInterval ColumnTypeName = `INTERVAL`

const ColumnTypeNameLong ColumnTypeName = `LONG`

const ColumnTypeNameMap ColumnTypeName = `MAP`

const ColumnTypeNameNull ColumnTypeName = `NULL`

const ColumnTypeNameShort ColumnTypeName = `SHORT`

const ColumnTypeNameString ColumnTypeName = `STRING`

const ColumnTypeNameStruct ColumnTypeName = `STRUCT`

const ColumnTypeNameTimestamp ColumnTypeName = `TIMESTAMP`

const ColumnTypeNameUserDefinedType ColumnTypeName = `USER_DEFINED_TYPE`

// String representation for [fmt.Print]
func (f *ColumnTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TIMESTAMP`, `USER_DEFINED_TYPE`:
		*f = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TIMESTAMP", "USER_DEFINED_TYPE"`, v)
	}
}

// Values returns all possible values for ColumnTypeName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ColumnTypeName) Values() []ColumnTypeName {
	return []ColumnTypeName{
		ColumnTypeNameArray,
		ColumnTypeNameBinary,
		ColumnTypeNameBoolean,
		ColumnTypeNameByte,
		ColumnTypeNameChar,
		ColumnTypeNameDate,
		ColumnTypeNameDecimal,
		ColumnTypeNameDouble,
		ColumnTypeNameFloat,
		ColumnTypeNameInt,
		ColumnTypeNameInterval,
		ColumnTypeNameLong,
		ColumnTypeNameMap,
		ColumnTypeNameNull,
		ColumnTypeNameShort,
		ColumnTypeNameString,
		ColumnTypeNameStruct,
		ColumnTypeNameTimestamp,
		ColumnTypeNameUserDefinedType,
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

// Request to execute a SQL statement on the warehouse given in the path.
type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explicitly set.
	ByteLimit int64 `json:"byte_limit,omitempty"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog string `json:"catalog,omitempty"`
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
	Disposition FetchDisposition `json:"disposition,omitempty"`
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
	Format Format `json:"format,omitempty"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling the statement status. When set to `CANCEL`, the statement
	// execution is canceled and the call returns with a `CANCELED` state.
	// Lakehouse//RT warehouses currently don't support `CANCEL`: the statement
	// keeps running as with `CONTINUE`.
	OnWaitTimeout TimeoutAction `json:"on_wait_timeout,omitempty"`
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
	Parameters []StatementParameter `json:"parameters,omitempty"`
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
	QueryTags []QueryTag `json:"query_tags,omitempty"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit int64 `json:"row_limit,omitempty"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema string `json:"schema,omitempty"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`. The maximum query text size is 16 MiB.
	Statement string `json:"statement"`
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
	WaitTimeout string `json:"wait_timeout,omitempty"`
	// The ID of the SQL warehouse to run the statement on. Currently, only
	// [Lakehouse//RT] warehouses are supported. See also [What are SQL
	// warehouses?].
	//
	// [Lakehouse//RT]: https://docs.databricks.com/compute/sql-warehouse/real-time.html
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExecuteStatementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExecuteStatementRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A short-lived cloud-storage link to one chunk of result data.
type ExternalLink struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// Indicates the date-time that the given external link will expire and
	// become invalid, after which point a new `external_link` must be
	// requested.
	Expiration string `json:"expiration,omitempty"`
	// A short-lived cloud-storage URL pointing to a chunk of result data,
	// hosted by an external service, with a short expiration time (<= 15
	// minutes). As this URL contains a temporary credential, it should be
	// considered sensitive and the client should not expose this URL in a log.
	ExternalLink string `json:"external_link,omitempty"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders map[string]string `json:"http_headers,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a get-result-chunk request for that index.
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`. Lakehouse//RT
	// warehouses currently don't return this field; use `next_chunk_index`
	// instead.
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLink) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLink) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FetchDisposition string

const FetchDispositionExternalLinks FetchDisposition = `EXTERNAL_LINKS`

const FetchDispositionInline FetchDisposition = `INLINE`

// String representation for [fmt.Print]
func (f *FetchDisposition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FetchDisposition) Set(v string) error {
	switch v {
	case `EXTERNAL_LINKS`, `INLINE`:
		*f = FetchDisposition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL_LINKS", "INLINE"`, v)
	}
}

// Values returns all possible values for FetchDisposition.
//
// There is no guarantee on the order of the values in the slice.
func (f *FetchDisposition) Values() []FetchDisposition {
	return []FetchDisposition{
		FetchDispositionExternalLinks,
		FetchDispositionInline,
	}
}

// Type always returns FetchDisposition to satisfy [pflag.Value] interface
func (f *FetchDisposition) Type() string {
	return "FetchDisposition"
}

// The format of the result data.
type Format string

const FormatArrowStream Format = `ARROW_STREAM`

const FormatCsv Format = `CSV`

const FormatJsonArray Format = `JSON_ARRAY`

// String representation for [fmt.Print]
func (f *Format) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Format) Set(v string) error {
	switch v {
	case `ARROW_STREAM`, `CSV`, `JSON_ARRAY`:
		*f = Format(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARROW_STREAM", "CSV", "JSON_ARRAY"`, v)
	}
}

// Values returns all possible values for Format.
//
// There is no guarantee on the order of the values in the slice.
func (f *Format) Values() []Format {
	return []Format{
		FormatArrowStream,
		FormatCsv,
		FormatJsonArray,
	}
}

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

type GetResultDataRequest struct {
	// The zero-based index of the result chunk to fetch.
	ChunkIndex int `json:"-" url:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId string `json:"-" url:"-"`
}

func (s *GetResultDataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetStatementResultRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
	// The ID of the SQL warehouse that ran the statement.
	WarehouseId string `json:"-" url:"-"`
}

func (s *GetStatementResultRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// * A query execution can be annotated with an optional key-value pair to allow
// users to attribute the executions by key and optional value to filter by.
// QueryTag is the user-facing representation.
type QueryTag struct {
	// The key of the tag.
	Key string `json:"key"`
	// The value of the tag. Key-only tags have no value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray [][]json.RawMessage `json:"data_array,omitempty"`
	// Short-lived cloud-storage links to the result data, used with the
	// `EXTERNAL_LINKS` disposition.
	ExternalLinks []ExternalLink `json:"external_links,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a get-result-chunk request for that index.
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`. Lakehouse//RT
	// warehouses currently don't return this field; use `next_chunk_index`
	// instead.
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultData) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	Chunks []BaseChunkInfo `json:"chunks,omitempty"`
	// The format of the result data, as requested with `format`.
	Format Format `json:"format,omitempty"`
	// The schema of the result set: the ordered list of column descriptions.
	Schema *Schema `json:"schema,omitempty"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount int64 `json:"total_byte_count,omitempty"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount int `json:"total_chunk_count,omitempty"`
	// The total number of rows in the result set.
	TotalRowCount int64 `json:"total_row_count,omitempty"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The schema is an ordered list of column descriptions.
type Schema struct {
	// The number of columns in the result set.
	ColumnCount int `json:"column_count,omitempty"`
	// The column descriptions, in result set order.
	Columns []ColumnInfo `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Schema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Schema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The error that caused a statement execution to fail.
type ServiceError struct {
	// The error code that classifies the failure.
	ErrorCode ServiceErrorCode `json:"error_code,omitempty"`
	// A brief summary of the error condition.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServiceError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServiceError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Error codes that classify a failed statement execution.
type ServiceErrorCode string

const ServiceErrorCodeAborted ServiceErrorCode = `ABORTED`

const ServiceErrorCodeAlreadyExists ServiceErrorCode = `ALREADY_EXISTS`

const ServiceErrorCodeBadRequest ServiceErrorCode = `BAD_REQUEST`

const ServiceErrorCodeCancelled ServiceErrorCode = `CANCELLED`

const ServiceErrorCodeDeadlineExceeded ServiceErrorCode = `DEADLINE_EXCEEDED`

const ServiceErrorCodeInternalError ServiceErrorCode = `INTERNAL_ERROR`

const ServiceErrorCodeIoError ServiceErrorCode = `IO_ERROR`

const ServiceErrorCodeNotFound ServiceErrorCode = `NOT_FOUND`

const ServiceErrorCodeResourceExhausted ServiceErrorCode = `RESOURCE_EXHAUSTED`

const ServiceErrorCodeServiceUnderMaintenance ServiceErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ServiceErrorCodeTemporarilyUnavailable ServiceErrorCode = `TEMPORARILY_UNAVAILABLE`

const ServiceErrorCodeUnauthenticated ServiceErrorCode = `UNAUTHENTICATED`

const ServiceErrorCodeUnknown ServiceErrorCode = `UNKNOWN`

const ServiceErrorCodeWorkspaceTemporarilyUnavailable ServiceErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ServiceErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServiceErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `DEADLINE_EXCEEDED`, `INTERNAL_ERROR`, `IO_ERROR`, `NOT_FOUND`, `RESOURCE_EXHAUSTED`, `SERVICE_UNDER_MAINTENANCE`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNKNOWN`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ServiceErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "DEADLINE_EXCEEDED", "INTERNAL_ERROR", "IO_ERROR", "NOT_FOUND", "RESOURCE_EXHAUSTED", "SERVICE_UNDER_MAINTENANCE", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNKNOWN", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ServiceErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServiceErrorCode) Values() []ServiceErrorCode {
	return []ServiceErrorCode{
		ServiceErrorCodeAborted,
		ServiceErrorCodeAlreadyExists,
		ServiceErrorCodeBadRequest,
		ServiceErrorCodeCancelled,
		ServiceErrorCodeDeadlineExceeded,
		ServiceErrorCodeInternalError,
		ServiceErrorCodeIoError,
		ServiceErrorCodeNotFound,
		ServiceErrorCodeResourceExhausted,
		ServiceErrorCodeServiceUnderMaintenance,
		ServiceErrorCodeTemporarilyUnavailable,
		ServiceErrorCodeUnauthenticated,
		ServiceErrorCodeUnknown,
		ServiceErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ServiceErrorCode to satisfy [pflag.Value] interface
func (f *ServiceErrorCode) Type() string {
	return "ServiceErrorCode"
}

// A named parameter to substitute for a parameter marker in the statement.
type StatementParameter struct {
	// The name of a parameter marker to be substituted in the statement.
	Name string `json:"name"`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type string `json:"type,omitempty"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StatementParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Manifest *ResultManifest `json:"manifest,omitempty"`
	// The first chunk of result data. Present when `status.state` is
	// `SUCCEEDED`.
	Result *ResultData `json:"result,omitempty"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"statement_id,omitempty"`
	// The execution state of the statement and, if it failed, the error
	// details.
	Status *StatementStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StatementResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {
	// The error details. Only populated when `state` is `FAILED`.
	Error *ServiceError `json:"error,omitempty"`
	// SQLSTATE error code returned when the statement execution fails. Only
	// populated when the statement status is `FAILED`.
	SqlState string `json:"sql_state,omitempty"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accompanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State StatementStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StatementStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Statement execution state.
type StatementStatusState string

const StatementStatusStateCanceled StatementStatusState = `CANCELED`

const StatementStatusStateClosed StatementStatusState = `CLOSED`

const StatementStatusStateFailed StatementStatusState = `FAILED`

const StatementStatusStatePending StatementStatusState = `PENDING`

const StatementStatusStateRunning StatementStatusState = `RUNNING`

const StatementStatusStateSucceeded StatementStatusState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *StatementStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StatementStatusState) Set(v string) error {
	switch v {
	case `CANCELED`, `CLOSED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = StatementStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "CLOSED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for StatementStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *StatementStatusState) Values() []StatementStatusState {
	return []StatementStatusState{
		StatementStatusStateCanceled,
		StatementStatusStateClosed,
		StatementStatusStateFailed,
		StatementStatusStatePending,
		StatementStatusStateRunning,
		StatementStatusStateSucceeded,
	}
}

// Type always returns StatementStatusState to satisfy [pflag.Value] interface
func (f *StatementStatusState) Type() string {
	return "StatementStatusState"
}

// What happens when the statement execution outlives `wait_timeout`.
type TimeoutAction string

const TimeoutActionCancel TimeoutAction = `CANCEL`

const TimeoutActionContinue TimeoutAction = `CONTINUE`

// String representation for [fmt.Print]
func (f *TimeoutAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TimeoutAction) Set(v string) error {
	switch v {
	case `CANCEL`, `CONTINUE`:
		*f = TimeoutAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCEL", "CONTINUE"`, v)
	}
}

// Values returns all possible values for TimeoutAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *TimeoutAction) Values() []TimeoutAction {
	return []TimeoutAction{
		TimeoutActionCancel,
		TimeoutActionContinue,
	}
}

// Type always returns TimeoutAction to satisfy [pflag.Value] interface
func (f *TimeoutAction) Type() string {
	return "TimeoutAction"
}
