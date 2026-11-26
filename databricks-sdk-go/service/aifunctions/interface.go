// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aifunctions

import (
	"context"
)

// Transform and enrich data with AI.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AiFunctionsService interface {

	// Classifies content according to a set of provided labels. The AI Classify
	// REST API has a default rate limit of 1,200 requests per minute per
	// workspace. Contact your account team to request a higher limit.
	AiClassify(ctx context.Context, request AiClassifyRequest) (*AiClassifyResponse, error)

	// Turn text and structured data into decisions your application can use.
	// Define questions and criteria to choose an option, estimate a
	// probability, or assign a score given a provided state. The AI Decide REST
	// API has a default rate limit of 1,200 requests per minute per workspace.
	// Contact your account team to request a higher limit.
	//
	// > **Note:** The v1 AI Decide syntax supporting strings and objects in
	// `state` and objects and arrays in `instructions` will no longer be
	// supported starting February 1, 2027. To use AI Decide, you must set
	// `options.version` to `"2.0"`, pass an array of typed content objects to
	// `state`, and a nonblank string to `instructions`.
	AiDecide(ctx context.Context, request AiDecideRequest) (*AiDecideResponse, error)

	// Extracts structured data from text and documents according to a provided
	// schema. The AI Extract REST API has a default rate limit of 120 requests
	// per minute per workspace. Contact your account team to request a higher
	// limit.
	AiExtract(ctx context.Context, request AiExtractRequest) (*AiExtractResponse, error)

	// Parse structured content from unstructured documents.
	AiParseDocument(ctx context.Context, request AiParseDocumentRequest) (*AiParseDocumentResponse, error)

	// Starts the asynchronous cancellation of a batch. Requests in progress are
	// stopped, and the results of the requests that finished are kept.
	// Cancellation is best effort, so the batch can still complete.
	CancelBatch(ctx context.Context, request CancelBatchRequest) (*CancelBatchResponse, error)

	// Creates a batch that runs an AI function on every request in a JSONL file
	// on a Unity Catalog volume. The batch runs asynchronously. To follow its
	// progress, get the batch.
	CreateBatch(ctx context.Context, request CreateBatchRequest) (*Batch, error)

	// Gets the latest status and progress of a batch.
	GetBatch(ctx context.Context, request GetBatchRequest) (*Batch, error)

	// Lists the batches that the caller created in this workspace.
	ListBatches(ctx context.Context, request ListBatchesRequest) (*ListBatchesResponse, error)
}
