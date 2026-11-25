package apierr

import (
	"encoding/json"
	"time"
)

// ErrorDetails contains the error details of an API error. It is the union of
// known error details types and unknown details.
type ErrorDetails struct {
	ErrorInfo           *ErrorInfo
	RequestInfo         *RequestInfo
	RetryInfo           *RetryInfo
	DebugInfo           *DebugInfo
	QuotaFailure        *QuotaFailure
	PreconditionFailure *PreconditionFailure
	BadRequest          *BadRequest
	ResourceInfo        *ResourceInfo
	Help                *Help

	// UnknownDetails contains error details that cannot be unmarshalled into
	// one of the known types above.
	UnknownDetails []any
}

// ErrorInfo describes the cause of the error with structured details.
type ErrorInfo struct {
	// The reason of the error. This is a constant value that identifies the
	// proximate cause of the error.
	Reason string

	// The logical grouping to which the "reason" belongs.
	Domain string

	// Additional structured details about this error.
	Metadata map[string]string
}

// RequestInfo Contains metadata about the request that clients can attach when
// filing a bug or providing other forms of feedback.
type RequestInfo struct {
	// An opaque string that should only be interpreted by the service that
	// generated it. For example, it can be used to identify requests in the
	// service's logs.
	RequestID string

	// Any data that was used to serve this request. For example, an encrypted
	// stack trace that can be sent back to the service provider for debugging.
	ServingData string
}

// RetryInfo describes when the clients can retry a failed request. Clients
// could ignore the recommendation here or retry when this information is
// missing from error responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retries have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay time.Duration
}

// Describes additional debugging info.
type DebugInfo struct {
	// The stack trace entries indicating where the error occurred.
	StackEntries []string

	// Additional debugging information provided by the server.
	Detail string
}

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the project
// id and the description of the quota limit that was exceeded.  If the
// calling project hasn't enabled the service in the developer console, then
// a service could respond with the project id and set `service_disabled`
// to true.
//
// Also see RetryInfo and Help types for other details about handling a
// quota failure.
type QuotaFailure struct {
	// Describes all quota violations.
	Violations []QuotaFailureViolation
}

type QuotaFailureViolation struct {
	// The subject on which the quota check failed.
	Subject string

	// A description of how the quota check failed. Clients can use this
	// description to find more about the quota configuration in the service's
	// public documentation, or find the relevant quota limit to adjust through
	// developer console.
	//
	// For example: "Service disabled" or "Daily Limit for read operations
	// exceeded".
	Description string
}

// Describes what preconditions have failed.
type PreconditionFailure struct {
	// Describes all precondition violations.
	Violations []PreconditionFailureViolation
}

type PreconditionFailureViolation struct {
	// The type of PreconditionFailure.
	Type string

	// The subject, relative to the type, that failed.
	Subject string

	// A description of how the precondition failed. Developers can use this
	// description to understand how to fix the failure.
	//
	// For example: "Terms of service not accepted".
	Description string
}

// Describes violations in a client request. This error type focuses on the
// syntactic aspects of the request.
type BadRequest struct {
	FieldViolations []BadRequestFieldViolation
}

type BadRequestFieldViolation struct {
	// A path leading to a field in the request body.
	Field string

	// A description of why the request element is bad.
	Description string
}

// Describes the resource that is being accessed.
type ResourceInfo struct {
	// A name for the type of resource being accessed.
	ResourceType string

	// The name of the resource being accessed.
	ResourceName string

	// The owner of the resource (optional).
	Owner string

	// Describes what error is encountered when accessing this resource.
	Description string
}

// Provides links to documentation or for performing an out of band action.
//
// For example, if a quota check failed with an error indicating the calling
// project hasn't enabled the accessed service, this can contain a URL pointing
// directly to the right place in the developer console to flip the bit.
type Help struct {
	// URL(s) pointing to additional information on handling the current error.
	Links []HelpLink
}

type HelpLink struct {
	// Describes what the link offers.
	Description string

	// The URL of the link.
	URL string
}

const (
	errorInfoType           string = "type.googleapis.com/google.rpc.ErrorInfo"
	requestInfoType         string = "type.googleapis.com/google.rpc.RequestInfo"
	retryInfoType           string = "type.googleapis.com/google.rpc.RetryInfo"
	debugInfoType           string = "type.googleapis.com/google.rpc.DebugInfo"
	quotaFailureType        string = "type.googleapis.com/google.rpc.QuotaFailure"
	preconditionFailureType string = "type.googleapis.com/google.rpc.PreconditionFailure"
	badRequestType          string = "type.googleapis.com/google.rpc.BadRequest"
	resourceInfoType        string = "type.googleapis.com/google.rpc.ResourceInfo"
	helpType                string = "type.googleapis.com/google.rpc.Help"
)

// errorInfoPb is the wire-format representation of ErrorInfo. It is used
// internally to unmarshal ErrorInfo from JSON.
type errorInfoPb struct {
	Reason   string            `json:"reason"`
	Domain   string            `json:"domain"`
	Metadata map[string]string `json:"metadata"`
}

// requestInfoPb is the wire-format representation of RequestInfo. It is used
// internally to unmarshal RequestInfo from JSON.
type requestInfoPb struct {
	RequestID   string `json:"request_id"`
	ServingData string `json:"serving_data"`
}

// retryInfoPb is the wire-format representation of RetryInfo. It is used
// internally to unmarshal RetryInfo from JSON.
type retryInfoPb struct {
	// The duration type is encoded as a string rather than an where the string
	// ends in the suffix "s" (indicating seconds) and is preceded by a decimal
	// number of seconds. For example, "3.000000001s", represents a duration of
	// 3 seconds and 1 nanosecond.
	RetryDelay string `json:"retry_delay"`
}

// debugInfoPb is the wire-format representation of DebugInfo. It is used
// internally to unmarshal DebugInfo from JSON.
type debugInfoPb struct {
	StackEntries []string `json:"stack_entries"`
	Detail       string   `json:"detail"`
}

// quotaFailurePb is the wire-format representation of QuotaFailure. It is used
// internally to unmarshal QuotaFailure from JSON.
type quotaFailurePb struct {
	Violations []quotaFailureViolationPb `json:"violations"`
}

type quotaFailureViolationPb struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

// preconditionFailurePb is the wire-format representation of PreconditionFailure.
// It is used internally to unmarshal PreconditionFailure from JSON.
type preconditionFailurePb struct {
	Violations []preconditionFailureViolationPb `json:"violations"`
}

type preconditionFailureViolationPb struct {
	Type        string `json:"type"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

// badRequestPb is the wire-format representation of BadRequest. It is used
// internally to unmarshal BadRequest from JSON.
type badRequestPb struct {
	FieldViolations []badRequestFieldViolationPb `json:"field_violations"`
}

type badRequestFieldViolationPb struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

// resourceInfoPb is the wire-format representation of ResourceInfo. It is used
// internally to unmarshal ResourceInfo from JSON.
type resourceInfoPb struct {
	ResourceType string `json:"resource_type"`
	ResourceName string `json:"resource_name"`
	Owner        string `json:"owner"`
	Description  string `json:"description"`
}

// helpPb is the wire-format representation of Help. It is used internally to
// unmarshal Help from JSON.
type helpPb struct {
	Links []helpLinkPb `json:"links"`
}

type helpLinkPb struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

func parseErrorDetails(details []any) ErrorDetails {
	var ed ErrorDetails
	for _, d := range details {
		switch d := d.(type) {
		case ErrorInfo:
			ed.ErrorInfo = &d
		case RequestInfo:
			ed.RequestInfo = &d
		case RetryInfo:
			ed.RetryInfo = &d
		case DebugInfo:
			ed.DebugInfo = &d
		case QuotaFailure:
			ed.QuotaFailure = &d
		case PreconditionFailure:
			ed.PreconditionFailure = &d
		case BadRequest:
			ed.BadRequest = &d
		case ResourceInfo:
			ed.ResourceInfo = &d
		case Help:
			ed.Help = &d
		default:
			ed.UnknownDetails = append(ed.UnknownDetails, d)
		}
	}
	return ed
}

// unmarshalDetails attempts to unmarshal the given slice of bytes into a known
// error details type. It works as follows:
//
//   - If the message is a known type, it unmarshals the message into that type.
//   - If the message is not a known type, it returns the results of calling
//     [json.Unmarshal] on the raw message.
//   - If [json.Unmarshal] fails, it returns the input as is.
func unmarshalDetails(d []byte) any {
	var a any
	if err := json.Unmarshal(d, &a); err != nil {
		return d // not a valid JSON message
	}
	m, ok := a.(map[string]any)
	if !ok {
		return a // not a JSON object
	}
	t, ok := m["@type"].(string)
	if !ok {
		return a // JSON object with no @type field
	}

	switch t {
	case errorInfoType:
		var pb errorInfoPb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		return ErrorInfo(pb)
	case requestInfoType:
		var pb requestInfoPb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		return RequestInfo(pb)
	case retryInfoType:
		var pb retryInfoPb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		d, err := time.ParseDuration(pb.RetryDelay)
		if err != nil {
			return m // not a valid known type
		}
		return RetryInfo{RetryDelay: d}
	case debugInfoType:
		var pb debugInfoPb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		return DebugInfo(pb)
	case quotaFailureType:
		var pb quotaFailurePb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		qf := QuotaFailure{}
		for _, v := range pb.Violations {
			qf.Violations = append(qf.Violations, QuotaFailureViolation(v))
		}
		return qf
	case preconditionFailureType:
		var pb preconditionFailurePb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		pf := PreconditionFailure{}
		for _, v := range pb.Violations {
			pf.Violations = append(pf.Violations, PreconditionFailureViolation(v))
		}
		return pf
	case badRequestType:
		var pb badRequestPb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		br := BadRequest{}
		for _, v := range pb.FieldViolations {
			br.FieldViolations = append(br.FieldViolations, BadRequestFieldViolation(v))
		}
		return br
	case resourceInfoType:
		var pb resourceInfoPb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		return ResourceInfo(pb)
	case helpType:
		var pb helpPb
		if err := json.Unmarshal(d, &pb); err != nil {
			return m // not a valid known type
		}
		h := Help{}
		for _, l := range pb.Links {
			h.Links = append(h.Links, HelpLink(l))
		}
		return h
	default:
		return m // unknown type
	}
}
