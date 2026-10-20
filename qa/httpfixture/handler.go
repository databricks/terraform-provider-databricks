// Package httpfixture replays HTTP exchanges without simulating remote state.
package httpfixture

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
)

// Body distinguishes an absent HTTP body from a JSON value, including JSON null.
type Body struct {
	JSON any `yaml:"json"`
}

// Request defines a complete signature; only listed headers are compared.
type Request struct {
	Method   string              `yaml:"method"`
	Path     string              `yaml:"path"`
	Query    map[string][]string `yaml:"query,omitempty"`
	Headers  map[string]string   `yaml:"headers,omitempty"`
	Body     *Body               `yaml:"body"`
	Matchers Matchers            `yaml:"matchers,omitempty"`
}

// Response defines the HTTP reply, with an optional JSON body.
type Response struct {
	Status  int               `yaml:"status"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Body    *Body             `yaml:"body"`
}

// Call pairs a request signature with its response.
type Call struct {
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

type exchange struct {
	call         Call
	expectedBody any
	matchers     Matchers
	responseBody []byte
	used         bool
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
		if call.Response.Status < 100 || call.Response.Status > 599 {
			return nil, fmt.Errorf("call %d: invalid response status %d", i+1, call.Response.Status)
		}
		e := exchange{call: call}
		var err error
		if call.Request.Body != nil {
			e.expectedBody, err = normalizedValue(call.Request.Body.JSON)
			if err != nil {
				return nil, fmt.Errorf("call %d request: %w", i+1, err)
			}
		}
		e.matchers, e.expectedBody, err = compileMatchers(call.Request.Matchers, e.expectedBody, call.Request.Body != nil)
		if err != nil {
			return nil, fmt.Errorf("call %d request matchers: %w", i+1, err)
		}
		if call.Response.Body != nil {
			e.responseBody, err = json.Marshal(call.Response.Body.JSON)
			if err != nil {
				return nil, fmt.Errorf("call %d response: %w", i+1, err)
			}
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

// matches checks the whole signature without consuming the exchange.
func (e *exchange) matches(req *http.Request, body []byte, value any) bool {
	expected := e.call.Request
	if e.used || expected.Method != req.Method || expected.Path != req.URL.EscapedPath() {
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
		return len(body) == 0
	}
	return len(body) > 0 && matchesJSON(e.expectedBody, value)
}

// ServeHTTP prefers consumable step fixtures over reusable fallbacks. Matching
// and consumption share a lock so concurrent requests cannot consume a call twice.
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	var value any
	if err == nil && len(body) > 0 {
		value, err = decodeJSON(body)
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if err != nil {
		h.fail(w, fmt.Errorf("%s: %s %s: invalid request body: %w", h.label, req.Method, req.URL.RequestURI(), err))
		return
	}
	next := -1
	for i := range h.calls {
		if next == -1 && !h.calls[i].used {
			next = i
		}
		if h.calls[i].matches(req, body, value) {
			if h.enforceCallOrder && i != next {
				expected := h.calls[next].call.Request
				h.fail(w, fmt.Errorf("%s: out-of-order HTTP call: %s %s body=%s matched call %d; next expected call %d: %s %s", h.label, req.Method, req.URL.RequestURI(), body, i+1, next+1, expected.Method, expected.Path))
				return
			}
			h.calls[i].used = true
			writeResponse(w, &h.calls[i])
			return
		}
	}
	for i := range h.reusable {
		if h.reusable[i].matches(req, body, value) {
			writeResponse(w, &h.reusable[i])
			return
		}
	}
	h.fail(w, fmt.Errorf("%s: unexpected HTTP call: %s %s body=%s", h.label, req.Method, req.URL.RequestURI(), body))
}

func writeResponse(w http.ResponseWriter, e *exchange) {
	for key, value := range e.call.Response.Headers {
		w.Header().Set(key, value)
	}
	if e.call.Response.Body != nil && w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(e.call.Response.Status)
	_, _ = w.Write(e.responseBody)
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
		if !e.used {
			failures = append(failures, fmt.Errorf("%s: unconsumed call %d: %s %s", h.label, i+1, e.call.Request.Method, e.call.Request.Path))
		}
	}
	return errors.Join(failures...)
}
