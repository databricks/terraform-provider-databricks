package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/logger/httplog"
)

// ErrHTMLContent is returned when the response body is HTML instead of JSON.
//
// This almost always indicates an issue with your authentication configuration.
// If you encounter this error, please verify the following:
//
//   - Databricks Host: Ensure your host is set correctly.
//   - Permissions: Confirm that the authentication method has the required
//     permissions for the API operation you are trying to perform.
//   - Network/Proxy: If you are behind a corporate firewall, ensure it is not
//     blocking or redirecting API traffic.
//
// A common cause of this error is Private Link redirecting the SDK to a login
// page, which the SDK cannot process. This usually happens when trying to
// access a Private Link-enabled workspace configured with no public internet
// access from a different network than the VPC endpoint belongs to.
//
// For more details, please refer to the [Unified Auth] documentation and
// [Private Link Authentication Troubleshooting].
//
// [Unified Auth]: https://docs.databricks.com/aws/en/dev-tools/auth/unified-auth
// [Private Link Authentication Troubleshooting]: https://learn.microsoft.com/en-us/azure/databricks/security/network/classic/private-link-standard#authentication-troubleshooting
var ErrHTMLContent = errors.New("received HTML response instead of JSON")

func WithResponseHeader(key string, value *string) DoOption {
	return DoOption{
		out: func(body *common.ResponseWrapper) error {
			*value = body.Response.Header.Get(key)
			return nil
		},
	}
}

// WithResponseUnmarshal unmarshals the response body into response. The
// supported response types are the following:
//   - *bytes.Buffer,
//   - *io.ReadCloser,
//   - *[]byte,
//   - a pointer to a struct with a Contents io.ReadCloser field,
//   - a pointer to a struct representing a JSON object.
//
// If response is a pointer to a io.ReadCloser or a struct with a io.ReadCloser
// field name "Contents", then the response io.ReadCloser is set to the value of
// the body's reader without actually reading it.
func WithResponseUnmarshal(response any) DoOption {
	return DoOption{
		in: func(r *http.Request) error {
			if r.Header.Get("Accept") != "" {
				return nil
			}
			switch response.(type) {
			case *bytes.Buffer, *io.ReadCloser, *[]byte:
				// don't send Accept header for raw types, even though we have
				// openapi/code/method.go:440#IsResponseByteStream() setting the
				// Accept header explicitly.
				return nil
			default:
				r.Header.Set("Accept", "application/json")
				return nil
			}
		},
		out: func(body *common.ResponseWrapper) error {
			if response == nil {
				return nil
			}
			err := injectHeaders(response, body)
			if err != nil {
				return err
			}

			if field, ok := findContentsField(response); ok {
				field.Set(reflect.ValueOf(body.ReadCloser))
				return nil
			}
			if reader, ok := response.(*io.ReadCloser); ok {
				*reader = body.ReadCloser
				return nil
			}
			if buffer, ok := response.(*bytes.Buffer); ok {
				defer body.ReadCloser.Close()
				_, err := buffer.ReadFrom(body.ReadCloser)
				return err
			}

			// At this point, fully read the content of the body and use it
			// to populate the response object (whether it is a slice of bytes
			// or a JSON object).
			defer body.ReadCloser.Close()
			bodyBytes, err := io.ReadAll(body.ReadCloser)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}
			if len(bodyBytes) == 0 {
				return nil
			}
			if bs, ok := response.(*[]byte); ok {
				*bs = bodyBytes
				return nil
			}
			if err := json.Unmarshal(bodyBytes, &response); err != nil {
				if _, ok := err.(*json.SyntaxError); ok && isHTMLContent(bodyBytes) {
					return ErrHTMLContent
				}
				return fmt.Errorf("failed to unmarshal response body: %w. %s", err, makeUnexpectedResponse(body.Response, body.RequestBody.DebugBytes, bodyBytes))
			}
			return nil
		},
	}
}

func isHTMLContent(bodyBytes []byte) bool {
	return strings.HasPrefix(string(bodyBytes), "<")
}

func findContentsField(response any) (*reflect.Value, bool) {
	value := reflect.ValueOf(response)
	value = reflect.Indirect(value)
	if value.Kind() != reflect.Struct {
		return nil, false
	}

	// Check if there is a "Contents" of type io.ReadCloser
	// This is by internal convention with the teams.
	ioType := reflect.TypeOf((*io.ReadCloser)(nil)).Elem()
	contentField := value.FieldByName("Contents")
	if !contentField.IsValid() || !contentField.CanSet() || contentField.Type() != ioType {
		return nil, false
	}
	return &contentField, true
}

func injectHeaders(response any, body *common.ResponseWrapper) error {
	value := reflect.ValueOf(response)
	value = reflect.Indirect(value)
	objectType := value.Type()

	if objectType.Kind() != reflect.Struct {
		return nil
	}
	headers := body.Response.Header

	for i := 0; i < objectType.NumField(); i++ {
		field := objectType.Field(i)
		headerTag := parseHeaderTag(field)
		if headerTag.ignore {
			continue
		}
		// We only support a single value for a header, and of types int and string.
		headerValue := headers.Get(headerTag.name)
		if headerValue == "" {
			continue
		}
		switch field.Type.Kind() {
		case reflect.String:
			value.Field(i).Set(reflect.ValueOf(headerValue))
		case reflect.Int:
			intValue, err := strconv.Atoi(headerValue)
			if err != nil {
				return fmt.Errorf("failed to parse header %s as int: %w", headerTag.name, err)
			}
			value.Field(i).Set(reflect.ValueOf(intValue))
		case reflect.Int64:
			intValue, err := strconv.ParseInt(headerValue, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse header %s as int: %w", headerTag.name, err)
			}
			value.Field(i).Set(reflect.ValueOf(intValue))
		default:
			// Do not fail the request if the header is not supported to avoid breaking changes.
			logger.Warnf(context.Background(), "unsupported header type %s for field %s", field.Type.Kind(), field.Name)
		}

	}
	return nil
}

type headerTag struct {
	name      string
	omitempty bool
	ignore    bool
}

func parseHeaderTag(field reflect.StructField) headerTag {
	raw := field.Tag.Get("header")
	name := field.Name

	if raw == "-" || raw == "" {
		return headerTag{ignore: true}
	}

	parts := strings.Split(raw, ",")

	if parts[0] != "" {
		name = parts[0]
	}

	headerTag := headerTag{
		name: name,
	}

	for _, v := range parts {
		switch v {
		case "omitempty":
			headerTag.omitempty = true
		}
	}
	return headerTag
}

func makeUnexpectedResponse(resp *http.Response, requestBody, responseBody []byte) string {
	var req *http.Request
	if resp != nil {
		req = resp.Request
	}
	rts := httplog.RoundTripStringer{
		Request:                  req,
		Response:                 resp,
		RequestBody:              requestBody,
		ResponseBody:             responseBody,
		DebugHeaders:             true,
		DebugTruncateBytes:       10 * 1024,
		DebugAuthorizationHeader: false,
	}
	return fmt.Sprintf("This is likely a bug in the Databricks SDK for Go or the underlying REST API. Please report this issue with the following debugging information to the SDK issue tracker at https://github.com/databricks/databricks-sdk-go/issues. Request log:\n```\n%s\n```", rts.String())
}
