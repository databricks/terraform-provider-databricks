// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Statement Execution API v2 can be used to execute SQL statements on a SQL
// warehouse and fetch the result. It works the same as the Statement Execution
// API (v1), except that `warehouse_id` is a path parameter on every call,
// including polling, fetching result chunks, and canceling. Routing on the
// warehouse ID gives lower latency and higher concurrency.
//
// Version 2 currently supports [Lakehouse//RT] warehouses only. For other
// warehouse types, use the Statement Execution API (v1). You can also run
// statements on Lakehouse//RT warehouses with version 1, but version 2 is
// optimized for them.
//
// [Lakehouse//RT]: https://docs.databricks.com/compute/sql-warehouse/real-time.html
package statementexecutionv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type StatementExecutionV2Interface interface {

	// Requests that an executing statement be canceled. Callers must poll for
	// status to see the terminal state. The cancel response is empty; receiving a
	// response indicates successful receipt.
	CancelStatement(ctx context.Context, request CancelStatementRequest) (*CancelStatementResponse, error)

	// Execute a SQL statement and optionally await its results for a specified
	// time.
	//
	// **Use case: small result sets with INLINE + JSON_ARRAY**
	//
	// For flows that generate small and predictable result sets (<= 25 MiB),
	// `INLINE` responses of `JSON_ARRAY` result data are typically the simplest way
	// to execute and fetch result data.
	//
	// **Use case: large result sets with EXTERNAL_LINKS**
	//
	// Using `EXTERNAL_LINKS` to fetch result data allows you to fetch large result
	// sets efficiently. The main differences from using `INLINE` disposition are
	// that the result data is accessed with short-lived cloud-storage URLs, and
	// that there are 3 supported formats: `JSON_ARRAY`, `ARROW_STREAM` and `CSV`
	// compared to only `JSON_ARRAY` with `INLINE`.
	//
	// **External-link URLs**
	//
	// External links point to data stored within your workspace's internal cloud
	// storage. The URLs are valid for only a short period, <= 15 minutes. Alongside
	// each `external_link` is an expiration field indicating the time at which the
	// URL is no longer valid. In `EXTERNAL_LINKS` mode, chunks can be resolved and
	// fetched multiple times and in parallel.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs that
	// are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived cloud-storage
	// URL is generated to download the results. The URL contains temporary access
	// credentials, so protect it and do not set an `Authorization` header in the
	// download request.
	//
	// The `EXTERNAL_LINKS` disposition can be disabled upon request by creating a
	// [support case].
	//
	// See also [Security best practices].
	//
	// ----
	//
	// StatementResponse contains `statement_id` and `status`; other fields might be
	// absent or present depending on context. If the SQL warehouse fails to execute
	// the provided statement, a 200 response is returned with `status.state` set to
	// `FAILED` (in contrast to a failure when accepting the request, which results
	// in a non-200 response). Details of the error can be found at `status.error`
	// in case of execution failures.
	//
	// [Security best practices]: https://docs.databricks.com/sql/admin/sql-execution-tutorial.html#security-best-practices
	// [support case]: https://docs.databricks.com/resources/support.html
	ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error)

	// After the statement execution has `SUCCEEDED`, this request can be used to
	// fetch any chunk by index. Whereas the first chunk with `chunk_index=0` is
	// typically fetched with the execute or the get-status request, this request
	// can be used to fetch subsequent chunks. The response structure is identical
	// to the nested `result` element of those responses, and similarly includes the
	// `next_chunk_index` field for simple iteration through the result set.
	// Depending on `disposition`, the response returns chunks of data either
	// inline, or as links.
	GetResultData(ctx context.Context, request GetResultDataRequest) (*ResultData, error)

	// This request can be used to poll for the statement's status.
	// StatementResponse contains `statement_id` and `status`; other fields might be
	// absent or present depending on context. When the `status.state` field is
	// `SUCCEEDED` it will also return the result manifest and the first chunk of
	// the result data. When the statement is in the terminal states `CANCELED`,
	// `CLOSED` or `FAILED`, it returns HTTP 200 with the state set. After at least
	// 12 hours in a terminal state, the statement is removed from the warehouse and
	// further calls will receive an HTTP 404 response.
	//
	// **NOTE** This call currently might take up to 5 seconds to get the latest
	// status and result.
	GetStatementResult(ctx context.Context, request GetStatementResultRequest) (*StatementResponse, error)
}

func NewStatementExecutionV2(client *client.DatabricksClient) *StatementExecutionV2API {
	return &StatementExecutionV2API{
		statementExecutionV2Impl: statementExecutionV2Impl{
			client: client,
		},
	}
}

// The Statement Execution API v2 can be used to execute SQL statements on a SQL
// warehouse and fetch the result. It works the same as the Statement Execution
// API (v1), except that `warehouse_id` is a path parameter on every call,
// including polling, fetching result chunks, and canceling. Routing on the
// warehouse ID gives lower latency and higher concurrency.
//
// Version 2 currently supports [Lakehouse//RT] warehouses only. For other
// warehouse types, use the Statement Execution API (v1). You can also run
// statements on Lakehouse//RT warehouses with version 1, but version 2 is
// optimized for them.
//
// [Lakehouse//RT]: https://docs.databricks.com/compute/sql-warehouse/real-time.html
type StatementExecutionV2API struct {
	statementExecutionV2Impl
}
