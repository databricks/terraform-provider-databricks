package sdkgoutility

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
	modularerr "github.com/databricks/sdk-go/core/apierr"
)

// ClassicError turns the error of a failed modular SDK call into the error
// that the classic client, github.com/databricks/databricks-sdk-go, returns
// for the same response. Code that reads classic errors then works on it as
// before: apierr.IsMissing, by which common.Resource drops a resource from
// state, errors.Is against the sentinel errors of the classic apierr package,
// and the message that users see in diagnostics.
//
// For a modular *apierr.APIError anywhere in the chain of err, ClassicError
// rebuilds the HTTP response from the status code, the header and the body
// that the error carries and parses it with the classic apierr.GetAPIError.
// It parses the body again, instead of mapping the status or the code of the
// modular error, because the classic parser doesn't decide by status alone: it
// also finds an error code in plain text and in HTML, and a 400 with
// RESOURCE_DOES_NOT_EXIST is a missing resource while a 404 with
// INVALID_PARAMETER_VALUE is not. The result is the classic *apierr.APIError
// by itself, without whatever wraps the modular error in err.
//
// ClassicError returns nil for nil and returns every other error unchanged:
// an error with no modular APIError in its chain, such as a classic
// *apierr.APIError, a network error or a context error, and a modular
// APIError with a status below 400, which the classic parser doesn't read as
// an error.
//
// The modular error carries neither the request nor the HTTP status text, so
// the classic parser gets stand-ins for them:
//
//   - The status text is http.StatusText of the status code, which is what
//     net/http itself reports over HTTP/2, where a response has no status
//     text. The classic parser makes the error code of an HTML response with
//     a <pre> block from it, so that error code, and with it the sentinel
//     error behind it, differs when a server sends another text or none over
//     HTTP/1.x.
//   - The request has no method and an empty URL. The classic parser has
//     overrides by method and path, which report a 400 with
//     INVALID_PARAMETER_VALUE from GET /api/2.x/clusters/get, jobs/get and
//     jobs/runs/get as a missing resource. ClassicError can't apply them, so
//     a caller of those APIs recognizes that response itself.
//
// ClassicError deliberately doesn't:
//
//   - retry. The classic client retries a 429, a 503 and some other errors
//     before it returns them; whether a modular call is retried is up to the
//     options of that call.
//   - recognize the redirect to the login page by which a workspace refuses a
//     request from outside its private link. The classic client reports it as
//     PRIVATE_LINK_VALIDATION_ERROR; a modular client follows the redirect
//     and fails to decode the page, which is no APIError.
//   - carry the request. ResponseWrapper of the result holds the rebuilt
//     response with the stand-in request and no request body.
func ClassicError(ctx context.Context, err error) error {
	modularErr, ok := errors.AsType[*modularerr.APIError](err)
	if !ok || modularErr == nil {
		return err
	}
	status, body := modularErr.HTTPStatusCode(), modularErr.HTTPBody()
	classicErr := apierr.GetAPIError(ctx, common.ResponseWrapper{
		ReadCloser: io.NopCloser(bytes.NewReader(body)),
		DebugBytes: body,
		Response: &http.Response{
			Status:     strconv.Itoa(status) + " " + http.StatusText(status),
			StatusCode: status,
			Header:     modularErr.HTTPHeader(),
			Body:       io.NopCloser(bytes.NewReader(body)),
			Request:    &http.Request{URL: &url.URL{}},
		},
	})
	if classicErr == nil {
		// A status below 400 is no error to the classic parser.
		return err
	}
	return classicErr
}
