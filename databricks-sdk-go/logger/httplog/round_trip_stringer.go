package httplog

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/exp/slices"
)

type RoundTripStringer struct {
	Request                  *http.Request
	Response                 *http.Response
	Err                      error
	RequestBody              []byte
	ResponseBody             []byte
	DebugHeaders             bool
	DebugAuthorizationHeader bool
	DebugTruncateBytes       int
}

var authorizationHeaders = map[string]bool{
	"Authorization":                          true,
	"X-Databricks-Azure-SP-Management-Token": true,
	"X-Databricks-GCP-SA-Access-Token":       true,
}

func (r RoundTripStringer) writeHeaders(sb *strings.Builder, prefix string, headers http.Header) {
	headerKeys := make([]string, 0, len(headers))
	for k := range headers {
		headerKeys = append(headerKeys, k)
	}
	slices.Sort(headerKeys)
	for i, k := range headerKeys {
		if i != 0 {
			sb.WriteString("\n")
		}
		v := headers[k]
		if _, ok := authorizationHeaders[k]; ok && !r.DebugAuthorizationHeader {
			v = []string{"REDACTED"}
		}
		trunc := onlyNBytes(strings.Join(v, ""), r.DebugTruncateBytes)
		sb.WriteString(fmt.Sprintf("%s* %s: %s", prefix, k, escapeNewLines(trunc)))
	}
}

func (r RoundTripStringer) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s %s", r.Request.Method,
		escapeNewLines(r.Request.URL.Path)))
	if r.Request.URL.RawQuery != "" {
		sb.WriteString("?")
		q, _ := url.QueryUnescape(r.Request.URL.RawQuery)
		sb.WriteString(q)
	}
	sb.WriteString("\n")
	if r.DebugHeaders {
		sb.WriteString("> * Host: ")
		sb.WriteString(escapeNewLines(r.Request.Host))
		sb.WriteString("\n")
		if len(r.Request.Header) > 0 {
			r.writeHeaders(&sb, "> ", r.Request.Header)
			sb.WriteString("\n")
		}
	}
	if len(r.RequestBody) > 0 {
		sb.WriteString(r.redactedDump("> ", r.RequestBody))
		sb.WriteString("\n")
	}
	sb.WriteString("< ")
	if r.Response == nil {
		sb.WriteString(fmt.Sprintf("Error: %s", r.Err))
		return sb.String()
	}

	sb.WriteString(fmt.Sprintf("%s %s", r.Response.Proto, r.Response.Status))
	// Only display error on this line if the response body is empty or the
	// client failed to read the response body.
	// Otherwise the response body will include details about the error.
	if len(r.ResponseBody) == 0 && r.Err != nil {
		sb.WriteString(fmt.Sprintf(" (Error: %s)", r.Err))
	}
	if r.DebugHeaders && len(r.Response.Header) > 0 {
		sb.WriteString("\n")
		r.writeHeaders(&sb, "< ", r.Response.Header)
	}
	if len(r.ResponseBody) > 0 {
		sb.WriteString("\n")
		sb.WriteString(r.redactedDump("< ", r.ResponseBody))
	}
	return sb.String()
}

func (r RoundTripStringer) redactedDump(prefix string, body []byte) (res string) {
	return bodyLogger{
		debugTruncateBytes: r.DebugTruncateBytes,
	}.redactedDump(prefix, body)
}

// CWE-117 prevention
func escapeNewLines(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	return in
}
