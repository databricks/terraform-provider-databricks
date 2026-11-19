// Package httpfixture replays HTTP exchanges without simulating remote state.
package httpfixture

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Request defines a complete signature; only listed headers are compared.
type Request struct {
	Method   string              `yaml:"method"`
	Path     string              `yaml:"path"`
	Query    map[string][]string `yaml:"query,omitempty"`
	Headers  map[string]string   `yaml:"headers,omitempty"`
	Body     *Body               `yaml:"body"`
	Matchers Matchers            `yaml:"matchers,omitempty"`
}

// Response defines the HTTP reply, with an optional JSON, text, or binary body.
type Response struct {
	Status  int               `yaml:"status,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Body    *Body             `yaml:"body"`
	// Delay is a Go duration, such as 2s, to wait before writing anything. The
	// wait ends early when the client closes the connection.
	Delay string `yaml:"delay,omitempty"`
	// Fault breaks the connection. Reset and close omit status, headers and body.
	Fault Fault `yaml:"fault,omitempty"`
	// StatusLine is the literal status line, such as HTTP/1.1 404 without a
	// reason phrase. Its code must equal Status.
	StatusLine string `yaml:"status_line,omitempty"`
}

// Call pairs a request signature with its response.
type Call struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
	// Times expects the call more than once in a row, each time with the same response.
	Times Times `yaml:"times,omitempty"`
	// MinGap and MaxGap are Go durations bounding the time since the previous
	// matched call's arrival, or with Times, since the call's previous repetition.
	MinGap string `yaml:"min_gap,omitempty"`
	MaxGap string `yaml:"max_gap,omitempty"`
}

type exchange struct {
	call         Call
	expectedBody any
	matchers     Matchers
	responseBody []byte
	delay        time.Duration
	bodyKind     string
	repeat
}

// Handler permits independent requests to interleave and consumes identical
// request signatures FIFO. A mutex keeps matching and consumption atomic.
type Handler struct {
	mu               sync.Mutex
	label            string
	calls            []exchange
	reusable         []exchange
	failures         []error
	enforceCallOrder bool
	// lastArrival and lastCall describe the step's previous matched call.
	lastArrival time.Time
	lastCall    string
}

// NewHandler returns an empty handler; fixture sets are configured separately.
func NewHandler() *Handler {
	return &Handler{}
}

// SetEnforceCallOrder requires consumable calls in fixture order when enabled.
// Reusable fallbacks remain unordered. Changing this setting preserves the
// current step's consumption and failures.
func (h *Handler) SetEnforceCallOrder(enabled bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.enforceCallOrder = enabled
}

// SetReusableFixtures replaces test-wide fallback exchanges, such as metadata.
// They survive step transitions, may be matched repeatedly, and need not be
// consumed. Invalid fixtures leave the previous set intact.
func (h *Handler) SetReusableFixtures(calls []Call) error {
	compiled, err := compileCalls(calls)
	if err != nil {
		return err
	}
	for i, e := range compiled {
		if e.repeated || e.hasGap() {
			return fmt.Errorf("reusable call %d: times, min_gap and max_gap apply only to step calls", i+1)
		}
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	h.reusable = compiled
	return nil
}

// StartStep reports the previous step's failures and installs the next exchanges.
// Installing them even after failure allows the runner to attempt cleanup.
func (h *Handler) StartStep(label string, calls []Call) error {
	compiled, err := compileCalls(calls)
	if err != nil {
		return fmt.Errorf("%s: %w", label, err)
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	err = h.verify()
	h.label, h.calls, h.failures = label, compiled, nil
	h.lastArrival, h.lastCall = time.Time{}, ""
	return err
}

// compileCalls validates and encodes the entire set before it can replace an
// active fixture set, so invalid input never installs only some of the calls.
func compileCalls(calls []Call) ([]exchange, error) {
	result := make([]exchange, len(calls))
	for i, call := range calls {
		if call.Request.Method == "" || !strings.HasPrefix(call.Request.Path, "/") {
			return nil, fmt.Errorf("call %d requires a method and absolute path", i+1)
		}
		if strings.ContainsAny(call.Request.Path, "?#") {
			return nil, fmt.Errorf("call %d: query parameters must use query", i+1)
		}
		if call.Response.writesStatus() && (call.Response.Status < 100 || call.Response.Status > 599) {
			return nil, fmt.Errorf("call %d: invalid response status %d", i+1, call.Response.Status)
		}
		e := exchange{call: call}
		var err error
		if call.Request.Body != nil {
			e.bodyKind, err = call.Request.Body.kind()
			if err == nil {
				switch e.bodyKind {
				case "json":
					e.expectedBody, err = normalizedValue(call.Request.Body.JSON)
				case "multipart":
					err = compileMultipart(call.Request.Body.Multipart)
				default:
					e.expectedBody, err = call.Request.Body.literalBytes()
				}
			}
			if err != nil {
				return nil, fmt.Errorf("call %d request: %w", i+1, err)
			}
		}
		e.matchers, e.expectedBody, err = compileMatchers(call.Request.Matchers, e.expectedBody, e.bodyKind == "json")
		if err != nil {
			return nil, fmt.Errorf("call %d request matchers: %w", i+1, err)
		}
		if err := compileResponse(&e, call.Response); err != nil {
			return nil, fmt.Errorf("call %d response: %w", i+1, err)
		}
		if e.repeat, err = compileRepeat(call); err != nil {
			return nil, fmt.Errorf("call %d: %w", i+1, err)
		}
		result[i] = e
	}
	return result, nil
}

// normalizedValue gives YAML fixture values the same representation as HTTP JSON.
func normalizedValue(value any) (any, error) {
	body, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return decodeJSON(body)
}

// decodeJSON preserves large numeric IDs and rejects trailing JSON values.
func decodeJSON(body []byte) (any, error) {
	decoder := json.NewDecoder(strings.NewReader(string(body)))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return nil, errors.New("body must contain exactly one JSON value")
	}
	return normalizeNumbers(value), nil
}

// normalizeNumbers equates spellings such as 1 and 1.0 without float64 rounding.
func normalizeNumbers(value any) any {
	switch v := value.(type) {
	case json.Number:
		n, _ := new(big.Rat).SetString(string(v))
		return n
	case map[string]any:
		for key, item := range v {
			v[key] = normalizeNumbers(item)
		}
	case []any:
		for i, item := range v {
			v[i] = normalizeNumbers(item)
		}
	}
	return value
}

// matches checks the whole signature, whether or not the exchange is still open.
func (e *exchange) matches(req *http.Request, body requestBody) bool {
	expected := e.call.Request
	if expected.Method != req.Method || expected.Path != req.URL.EscapedPath() {
		return false
	}
	if !matchesQuery(expected.Query, e.matchers.Query, req.URL.Query()) {
		return false
	}
	for key, header := range expected.Headers {
		if _, overridden := e.matchers.Headers[http.CanonicalHeaderKey(key)]; overridden {
			continue
		}
		if req.Header.Get(key) != header {
			return false
		}
	}
	for key := range e.matchers.Headers {
		if len(req.Header.Values(key)) != 1 {
			return false
		}
	}
	if expected.Body == nil {
		return len(body.raw) == 0
	}
	switch e.bodyKind {
	case "json":
		return body.jsonValid && matchesJSON(e.expectedBody, body.json)
	case "multipart":
		return matchesMultipart(expected.Body.Multipart, body.parts)
	default:
		return bytes.Equal(e.expectedBody.([]byte), body.raw)
	}
}

// ServeHTTP prefers consumable step fixtures over reusable fallbacks.
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Gaps are measured between arrivals, before reading the body or waiting for the lock.
	arrival := time.Now()
	e, label := h.consume(w, req, arrival)
	if e == nil {
		return
	}
	if err := respond(w, req, e); err != nil {
		h.mu.Lock()
		defer h.mu.Unlock()
		h.failures = append(h.failures, fmt.Errorf("%s: %s %s: %w", label, req.Method, req.URL.RequestURI(), err))
	}
}

// consume matches and consumes a fixture under the lock, so concurrent requests
// cannot consume a call twice. It returns nil after reporting a failure.
func (h *Handler) consume(w http.ResponseWriter, req *http.Request, arrival time.Time) (*exchange, string) {
	body, err := parseRequestBody(req)
	h.mu.Lock()
	defer h.mu.Unlock()
	if err != nil {
		h.fail(w, fmt.Errorf("%s: %s %s: invalid request body: %w", h.label, req.Method, req.URL.RequestURI(), err))
		return nil, ""
	}
	for i := range h.calls {
		if !h.calls[i].open() || !h.calls[i].matches(req, body) {
			continue
		}
		if next := h.nextRequired(); next != -1 && i > next {
			expected := h.calls[next].call.Request
			h.fail(w, fmt.Errorf("%s: out-of-order HTTP call: %s %s body=%s matched call %d; next expected call %d: %s %s", h.label, req.Method, req.URL.RequestURI(), body.raw, i+1, next+1, expected.Method, expected.Path))
			return nil, ""
		}
		h.record(i, req, arrival)
		return &h.calls[i], h.label
	}
	for i := range h.reusable {
		if h.reusable[i].matches(req, body) {
			return &h.reusable[i], h.label
		}
	}
	h.fail(w, fmt.Errorf("%s: unexpected HTTP call: %s %s body=%s", h.label, req.Method, req.URL.RequestURI(), body.raw))
	return nil, ""
}

// nextRequired returns the first open call matched fewer than its minimum
// number of times, which ordered replay cannot skip past. It returns -1 when
// nothing is required or calls are unordered.
func (h *Handler) nextRequired() int {
	if !h.enforceCallOrder {
		return -1
	}
	for i := range h.calls {
		if h.calls[i].open() && !h.calls[i].satisfied() {
			return i
		}
	}
	return -1
}

// record consumes call i at its arrival time. A gap out of bounds is reported
// but still answered, so the step fails with the gap rather than a later error.
// Ordered replay closes the earlier calls, which can no longer repeat.
func (h *Handler) record(i int, req *http.Request, arrival time.Time) {
	e := &h.calls[i]
	if err := e.checkGap(arrival, h.lastArrival, h.lastCall); err != nil {
		h.failures = append(h.failures, fmt.Errorf("%s: call %d: %s %s %w", h.label, i+1, req.Method, req.URL.RequestURI(), err))
	}
	if h.enforceCallOrder {
		for j := range h.calls[:i] {
			h.calls[j].closed = true
		}
	}
	e.count++
	e.lastArrival = arrival
	// Concurrent requests can take the lock out of arrival order; the latest
	// arrival stays the reference.
	if !arrival.Before(h.lastArrival) {
		h.lastArrival = arrival
		h.lastCall = fmt.Sprintf("call %d (%s %s)", i+1, req.Method, req.URL.RequestURI())
	}
}

func writeResponse(w http.ResponseWriter, e *exchange) {
	setResponseHeaders(w.Header(), e)
	w.WriteHeader(e.call.Response.Status)
	_, _ = w.Write(e.responseBody)
}

// setResponseHeaders applies the explicit headers, then the body's default type.
func setResponseHeaders(header http.Header, e *exchange) {
	for key, value := range e.call.Response.Headers {
		header.Set(key, value)
	}
	if e.call.Response.Body != nil && header.Get("Content-Type") == "" {
		kind, _ := e.call.Response.Body.kind()
		contentType := map[string]string{"json": "application/json", "text": "text/plain; charset=utf-8", "base64": "application/octet-stream"}
		header.Set("Content-Type", contentType[kind])
	}
}

func (h *Handler) fail(w http.ResponseWriter, err error) {
	h.failures = append(h.failures, err)
	// A non-retryable error keeps missing fixtures from becoming long SDK retries.
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// Verify reports unexpected requests and unconsumed step exchanges without
// clearing failures. Reusable fixtures are excluded from consumption checks.
func (h *Handler) Verify() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.verify()
}

// verify requires h.mu to be held.
func (h *Handler) verify() error {
	failures := append([]error(nil), h.failures...)
	for i, e := range h.calls {
		if !e.satisfied() {
			failures = append(failures, fmt.Errorf("%s: unconsumed call %d: %s %s%s", h.label, i+1, e.call.Request.Method, e.call.Request.Path, e.countDetail()))
		}
	}
	return errors.Join(failures...)
}
