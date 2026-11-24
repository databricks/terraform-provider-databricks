// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aifunctions

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// ai_classify
type AiClassifyOptions struct {
	// When true, includes a per-label confidence score in the response.
	EnableConfidenceScores bool `json:"enable_confidence_scores,omitempty"`
	// When true, includes a rationale explaining each classification in the
	// response.
	EnableRationales bool `json:"enable_rationales,omitempty"`
	// Natural-language guidance that steers how the text is classified (up to
	// 20,000 characters).
	Instructions string `json:"instructions,omitempty"`
	// When true, allows more than one label to be returned per input.
	Multilabel bool `json:"multilabel,omitempty"`
	// The function version to invoke. Defaults to the latest version. Supported
	// versions: ["2.1"].
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiClassifyOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiClassifyOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiClassifyRequest struct {
	// The content to classify. It accepts a plain string or the response object
	// of [ai_parse_document](:method:AiFunctions/AiParseDocument).
	Content json.RawMessage `json:"content"`
	// The label set to classify as. Either a JSON array of label strings (e.g.
	// ["spam", "not_spam"]), or a JSON object mapping each label to a
	// description (e.g. {"spam": "unsolicited bulk message", "not_spam": "a
	// legitimate message"}). Accepts 2 to 5000 labels, each 1 to 100
	// characters. When multilabel is true, a single label is also accepted.
	Labels json.RawMessage `json:"labels"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options *AiClassifyOptions `json:"options,omitempty"`
}

func (s *AiClassifyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiClassifyResponse struct {
	// Additional metadata returned by AI Classify.
	Metadata *AiClassifyResponseMetadata `json:"metadata,omitempty"`
	// The function result as a JSON value. An array of per-label objects: one
	// element in single-label mode (the default), or multiple elements when
	// `multilabel` is true. When `enable_confidence_scores` and
	// `enable_rationales` are true, `confidence_score` and `rationale` are
	// included in each response value, respectively.
	Response *json.RawMessage `json:"response,omitempty"`
}

func (s *AiClassifyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiClassifyResponseMetadata struct {
	// The resolved function version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiClassifyResponseMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiClassifyResponseMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ai_decide
type AiDecideOptions struct {
	// Set to `"2.0"`. If unset, the deprecated version 1.0 is used during the
	// [migration period](/api/ai-functions/v1#aidecide).
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiDecideOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiDecideOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiDecideRequest struct {
	// Function options. Omitted fields fall back to their documented defaults.
	Options *AiDecideOptions `json:"options,omitempty"`
	// A JSON object mapping up to 256 question IDs to their definitions. Each
	// ID must be nonblank; its answer is returned with the same ID in
	// `response.answers`.
	//
	// Each definition is an object with the required fields `type` and
	// `instructions`. The `criteria` field is optional for the type `noul` but
	// is required for the types `choice` and `score`.
	//
	// In AI Decide v2, `instructions` and each criteria description must be
	// nonblank strings. Each question may contain at most 20,000 characters
	// across its ID, instructions, and criteria.
	//
	// The `type` can be one of:
	//
	// - `choice`: Selects one option from a defined set. Requires `criteria` to
	// be an object mapping 1 to 255 nonblank option names to nonblank string
	// descriptions. For example:
	//
	// ```json { "team": { "type": "choice", "instructions": "Which team should
	// handle this ticket?", "criteria": { "billing": "Payments, charges, and
	// refunds", "technical_support": "Product errors and troubleshooting" } } }
	// ```
	//
	// - `noul`: Estimates the probability that the answer to a true-or-false
	// question is true. `criteria` can take the fields `true` or `false`, or
	// both, with nonblank string descriptions. Omit `criteria` to use the
	// question alone. For example, both of the following are valid:
	//
	// ```json { "escalate": { "type": "noul", "instructions": "Does this ticket
	// need escalation?", "criteria": { "true": "Suspected fraud or an exception
	// to standard policy", "false": "A routine issue frontline support can
	// resolve" } } } ```
	//
	// or
	//
	// ```json { "escalate": { "type": "noul", "instructions": "Does this ticket
	// need escalation?" } } ```
	//
	// - `score`: Rates the state on an ordered scale. Requires `criteria` to be
	// an array of 2 to 10 unique, nonblank string descriptions, ordered from
	// low to high. Array positions define levels starting at 0. For example:
	//
	// ```json { "urgency": { "type": "score", "instructions": "How urgent is
	// this ticket?", "criteria": [ "Routine: can wait a few days",
	// "Time-sensitive: needs attention today", "Critical: needs immediate
	// action" ] } } ```
	Questions json.RawMessage `json:"questions"`
	// A nonempty array of typed content objects shared by all questions. Set
	// `options.version` to `"2.0"` to use this format. Text and JSON content
	// plus questions must fit within 32 KiB as compact UTF-8 JSON, excluding
	// image parts.
	//
	// | type | Content field | Value | | --- | --- | --- | | `text` | `text` |
	// A nonblank string. | | `json` | `data` | A JSON object or array. | |
	// `image` | `image_url` | An inline base64 PNG or JPEG data URL:
	// `data:image/png;base64,<base64-encoded image bytes>` or
	// `data:image/jpeg;base64,<base64-encoded image bytes>`. HTTP URLs and file
	// paths are not supported. Each image must be at most 4 MiB after base64
	// decoding, with a resolution up to 1,568 × 1,568 pixels (or an equivalent
	// pixel count). A request supports up to five images. Images are only
	// supported in [select regions]. |
	//
	// [select regions]: https://docs.databricks.com/aws/en/resources/feature-region-support#ai-function-availability
	State json.RawMessage `json:"state"`
}

func (s *AiDecideRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiDecideResponse struct {
	// Metadata identifying the function API version used for this request.
	Metadata *AiDecideResponseMetadata `json:"metadata,omitempty"`
	// A JSON object containing an `answers` map with one answer for each
	// question, keyed by the same IDs as `questions`.
	//
	// Each answer has a `type` matching its question and the following fields:
	//
	// - `choice`: choice is the highest-probability option from `criteria`.
	// `probabilities` maps every option name to its probability. `confidence`
	// is a number from 0 to 1 indicating how well the state supports the
	// assessment. The `probabilities` values sum to 1. For example:
	//
	// ```json { "answers": { "team": { "type": "choice", "choice": "billing",
	// "probabilities": { "billing": 0.85, "technical_support": 0.15 },
	// "confidence": 0.9 } } } ```
	//
	// - `noul`: `probability` is a number from 0 to 1 estimating the
	// probability that the answer is true. For example:
	//
	// ```json { "answers": { "escalate": { "type": "noul", "probability": 0.8 }
	// } } ```
	//
	// - `score`: score is a numeric rating on the zero-based scale defined by
	// `criteria`. It can fall between levels, from 0 to the number of levels
	// minus 1. `probabilities` maps each level index to its probability, and
	// `legend` maps each index to its original description. Both maps use
	// string keys such as "0", "1", and "2". `confidence` is a number from 0 to
	// 1 indicating how well the state supports the assessment. The
	// `probabilities` values sum to 1. Depending on the model, `score` may be
	// truncated and `probabilities` rounded to 2 decimal places while
	// preserving a sum of 1, so the score can differ slightly from the
	// probability-weighted mean computed from the returned probabilities. For
	// example:
	//
	// ```json { "answers": { "urgency": { "type": "score", "score": 1.6,
	// "probabilities": { "0": 0.1, "1": 0.2, "2": 0.7 }, "legend": { "0":
	// "Routine: can wait a few days", "1": "Time-sensitive: needs attention
	// today", "2": "Critical: needs immediate action" }, "confidence": 0.85 } }
	// } ```
	Response *json.RawMessage `json:"response,omitempty"`
}

func (s *AiDecideResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiDecideResponseMetadata struct {
	// The function API version used to evaluate the request.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiDecideResponseMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiDecideResponseMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A bounding box on a source page; used by bbox-input citations.
type AiExtractBbox struct {
	// Pixel coordinates on the page image as [x0, y0, x1, y1].
	Coord []int64 `json:"coord,omitempty"`
	// 0-based page index the box is on.
	PageId int64 `json:"page_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractBbox) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractBbox) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A citation locating an extracted value in the source. start/stop are set for
// span (STRING input) citations, bbox for bbox (parsed-document input)
// citations.
type AiExtractCitation struct {
	// Bounding boxes locating the citation on the source pages; set for bbox
	// citations.
	Bbox []AiExtractBbox `json:"bbox,omitempty"`
	// Integer matching a citation_ids entry on an extracted field.
	Id int64 `json:"id,omitempty"`
	// Inclusive 0-based character offset into the input string; set for span
	// citations.
	Start int64 `json:"start,omitempty"`
	// Exclusive 0-based character offset into the input string; set for span
	// citations.
	Stop int64 `json:"stop,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractCitation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractCitation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ai_extract
type AiExtractOptions struct {
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
	EnableCitations bool `json:"enable_citations,omitempty"`
	// When true, includes a per-field confidence score in the response.
	EnableConfidenceScores bool `json:"enable_confidence_scores,omitempty"`
	// Natural-language guidance that steers how data is extracted (up to 20,000
	// characters).
	Instructions string `json:"instructions,omitempty"`
	// Extraction mode. Supported modes: "precision" — more powerful
	// extraction for complex schemas, long documents, and reasoning-heavy
	// extractions. Defaults to none (standard extraction).
	Mode string `json:"mode,omitempty"`
	// The function version to invoke. Defaults to the latest version. Supported
	// versions: ["2.1"].
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiExtractRequest struct {
	// The text to extract from. It accepts a plain string or the response
	// object of [ai_parse_document](:method:AiFunctions/AiParseDocument).
	Content json.RawMessage `json:"content"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options *AiExtractOptions `json:"options,omitempty"`
	// The extraction schema defining the fields to extract. Either a JSON array
	// of field names, assumed to be strings (e.g. ["company", "valuation"]), or
	// a JSON object mapping each field to its type/description (e.g.
	// {"company": {"type": "string", "description": "the company name"}}). Enum
	// fields list their values in "labels" (e.g. {"currency": {"type": "enum",
	// "labels": ["USD", "EUR"]}}). Primitive and enum fields are nullable by
	// default. Accepts up to 256 fields, 12 levels of nesting, and 500 enum
	// values. Supported field types are string, integer, number, boolean, and
	// enum.
	Schema json.RawMessage `json:"schema"`
}

func (s *AiExtractRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiExtractResponse struct {
	// Additional metadata returned by AI Extract.
	Metadata *AiExtractResponseMetadata `json:"metadata,omitempty"`
	// The function result as a JSON value. When `enable_confidence_scores` and
	// `enable_citations` are true, `confidence` and `citation_ids` are included
	// in each response field, respectively.
	Response *json.RawMessage `json:"response,omitempty"`
}

func (s *AiExtractResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiExtractResponseMetadata struct {
	// How the source was chunked for citation offsets (span for text input,
	// bbox for parsed-document input); present when citations are enabled.
	ChunkType string `json:"chunk_type,omitempty"`
	// Citation objects locating each result in the source; present when
	// citations are enabled.
	Citations []AiExtractCitation `json:"citations,omitempty"`
	// The resolved extraction mode; present when a non-default mode was used.
	Mode string `json:"mode,omitempty"`
	// The resolved function version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractResponseMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractResponseMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata about the source file; present only for file-path input.
type AiParseDocumentFileMetadata struct {
	// Last-modified timestamp of the source file, as an HTTP date string.
	FileModificationTime string `json:"file_modification_time,omitempty"`
	// Base name of the source file.
	FileName string `json:"file_name,omitempty"`
	// Unity Catalog volume path of the source file.
	FilePath string `json:"file_path,omitempty"`
	// Size of the source file in bytes.
	FileSize int64 `json:"file_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentFileMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentFileMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ai_parse_document
type AiParseDocumentOptions struct {
	// Element types for which an AI-generated description is produced. Use "*"
	// (default) to generate descriptions for all supported element types,
	// "figure" to generate them for figures only, or "" (empty string) to
	// generate none. Only figure descriptions are supported for version "2.0",
	// so "*" and "figure" produce the same behavior.
	DescriptionElementTypes string `json:"description_element_types,omitempty"`
	// Unity Catalog volume path where rendered page and element images are
	// written.
	ImageOutputPath string `json:"image_output_path,omitempty"`
	// Pages to parse (1-indexed), as a comma-separated list of page numbers or
	// ranges (e.g. "1,3,5-10").
	PageRange string `json:"page_range,omitempty"`
	// The ai_parse_document output schema version. Supported value: "2.0".
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A single page that failed to parse while the overall request succeeded.
type AiParseDocumentPageError struct {
	// Message describing why the page failed.
	ErrorMessage string `json:"error_message,omitempty"`
	// 0-based index of the page that failed.
	PageId int64 `json:"page_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentPageError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentPageError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiParseDocumentRequest struct {
	// The document to parse. The REST API accepts a Unity Catalog volume path
	// to the source file or a base64 data URL containing the file bytes (for
	// example, data:application/pdf;base64,<encoded-file-bytes> or
	// data:image/jpeg;base64,<encoded-file-bytes>). Supported formats: PDF,
	// DOCX, DOC, PPTX, PPT, JPG, JPEG, PNG, TIFF. Accepts up to 100 pages and
	// 100 MB per document.
	Content string `json:"content"`
	// Function options. Omitted fields fall back to their documented defaults.
	Options *AiParseDocumentOptions `json:"options,omitempty"`
}

func (s *AiParseDocumentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiParseDocumentResponse struct {
	// The parsed document as a JSON value, containing the extracted pages and
	// elements.
	Document *json.RawMessage `json:"document,omitempty"`
	// Per-page partial-failure details; present when the request succeeds (2xx)
	// but individual pages fail.
	ErrorStatus []AiParseDocumentPageError `json:"error_status,omitempty"`
	// Additional metadata returned by AI Parse Document.
	Metadata *AiParseDocumentResponseMetadata `json:"metadata,omitempty"`
}

func (s *AiParseDocumentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiParseDocumentResponseMetadata struct {
	// Describes the source file; present only for file-path input.
	FileMetadata *AiParseDocumentFileMetadata `json:"file_metadata,omitempty"`
	// Unique identifier for the parse request.
	Id string `json:"id,omitempty"`
	// The resolved function version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentResponseMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentResponseMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The AI Functions Asynchronous API lets you process large volumes of requests
// as a batch in a single call. It is in Beta: a workspace admin must turn on
// the **AI Functions Asynchronous API** preview on the workspace Previews page
// before you can call it.
//
// For examples, see
// [`ai_parse_document`](/agents/agent-bricks/document-parsing#asynchronous-api),
// [`ai_extract`](/agents/agent-bricks/info-extraction#asynchronous-api), and
// [`ai_classify`](/agents/agent-bricks/classification#asynchronous-api).
type Batch struct {
	// Output only. The time when the batch was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The AI function, for example `ai_parse_document`.
	Endpoint string `json:"endpoint"`
	// Unity Catalog volume path of the JSONL input file. Each line is one
	// request, in the form `{"custom_id": ..., "body": ...}`. The file can hold
	// up to 50,000 requests and be up to 200 MiB.
	InputRef string `json:"input_ref"`
	// Output only. Unique identifier of the batch, returned by create. Format:
	// `batch_<id>`.
	Name string `json:"name,omitempty"`
	// Unity Catalog volume path to write the JSONL output file to. If the path
	// is a directory, the file is `<path>/<name>_output.jsonl`. Each line is
	// the result of one request, in the form `{"custom_id": ..., "response":
	// ...}`, or `{"custom_id": ..., "error": {"type": ..., "message": ...}}`
	// for a request that failed or did not finish.
	OutputRef string `json:"output_ref"`
	// Parameters of the function other than its input, applied to every request
	// in the batch, for example `schema` for `ai_extract`. Put function
	// options, such as `version`, in a nested `options` object. For the
	// parameters of each function, see [AI
	// Functions](/large-language-models/ai-functions).
	Params map[string]json.RawMessage `json:"params,omitempty"`
	// Output only. The percentage of requests that are processed, from 0 to
	// 100. Returned when you get a batch, not when you list batches.
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// Output only. The number of requests in the batch, by result. Returned
	// when you get a batch, not when you list batches.
	RequestCounts *BatchRequestCounts `json:"request_counts,omitempty"`
	// Output only. The status of the batch. A batch starts as `PENDING`, runs
	// as `IN_PROGRESS`, is `FINALIZING` while the output file is written, and
	// ends as `COMPLETED`. A running batch that is cancelled is `CANCELLING`,
	// then `CANCELLED`. A `COMPLETED` batch can include requests that failed,
	// and a `FAILED` batch has no `output_ref`.
	Status BatchStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Batch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Batch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Counts of the requests in a batch.
type BatchRequestCounts struct {
	// Output only. The number of requests that failed. The service does not
	// retry them.
	Failed int64 `json:"failed,omitempty"`
	// Output only. The number of requests that were processed successfully.
	Succeeded int64 `json:"succeeded,omitempty"`
	// Output only. The number of requests in the batch.
	Total int64 `json:"total,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BatchRequestCounts) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BatchRequestCounts) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of a batch.
type BatchStatus string

const BatchStatusCancelled BatchStatus = `CANCELLED`

const BatchStatusCancelling BatchStatus = `CANCELLING`

const BatchStatusCompleted BatchStatus = `COMPLETED`

const BatchStatusFailed BatchStatus = `FAILED`

const BatchStatusFinalizing BatchStatus = `FINALIZING`

const BatchStatusInProgress BatchStatus = `IN_PROGRESS`

const BatchStatusPending BatchStatus = `PENDING`

// String representation for [fmt.Print]
func (f *BatchStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BatchStatus) Set(v string) error {
	switch v {
	case `CANCELLED`, `CANCELLING`, `COMPLETED`, `FAILED`, `FINALIZING`, `IN_PROGRESS`, `PENDING`:
		*f = BatchStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "CANCELLING", "COMPLETED", "FAILED", "FINALIZING", "IN_PROGRESS", "PENDING"`, v)
	}
}

// Values returns all possible values for BatchStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *BatchStatus) Values() []BatchStatus {
	return []BatchStatus{
		BatchStatusCancelled,
		BatchStatusCancelling,
		BatchStatusCompleted,
		BatchStatusFailed,
		BatchStatusFinalizing,
		BatchStatusInProgress,
		BatchStatusPending,
	}
}

// Type always returns BatchStatus to satisfy [pflag.Value] interface
func (f *BatchStatus) Type() string {
	return "BatchStatus"
}

type CancelBatchRequest struct {
	// Unique identifier of the batch. Format: `batch_<id>`.
	Name string `json:"-" url:"-"`
}

func (s *CancelBatchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CancelBatchResponse struct {
}

func (s *CancelBatchResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateBatchRequest struct {
	// The batch to create.
	Batch Batch `json:"batch"`
}

func (s *CreateBatchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetBatchRequest struct {
	// Unique identifier of the batch. Format: `batch_<id>`.
	Name string `json:"-" url:"-"`
}

func (s *GetBatchRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListBatchesRequest struct {
	// Maximum number of batches to return in one page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous call to list batches. Omit it to get the
	// first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBatchesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBatchesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBatchesResponse struct {
	// The batches in this page.
	Batches []Batch `json:"batches,omitempty"`
	// A token to get the next page. It is empty on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBatchesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBatchesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
