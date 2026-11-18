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
)

// Request defines a complete signature; only listed headers are compared. Host
// names the host the request must reach; empty means the handler's default host.
type Request struct {
	Method   string              `yaml:"method"`
	Host     string              `yaml:"host,omitempty"`
	Path     string              `yaml:"path"`
	Query    map[string][]string `yaml:"query,omitempty"`
	Headers  map[string]string   `yaml:"headers,omitempty"`
	Body     *Body               `yaml:"body"`
	Matchers Matchers            `yaml:"matchers,omitempty"`
}

// Response defines the HTTP reply, with an optional JSON, text, or binary body.
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
	host         string
	expectedBody any
	matchers     compiledMatchers
	responseBody []byte
	bodyKind     string
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
	hosts            hostNames
}

// hostNames maps request Host headers to host names. Calls without a host
// expect defaultName, which is empty for requests to an unnamed host.
type hostNames struct {
	byAuthority map[string]string
	defaultName string
}

// SetHosts names the hosts requests may reach: hosts maps each name to the
// authority (host or host:port) a request carries in its Host header. Requests
// to other authorities belong to no named host. Calls without a host match
// defaultHost, or requests to an unnamed host when defaultHost is empty. Set
// hosts before installing fixtures that name them; invalid hosts leave the
// previous set intact.
func (h *Handler) SetHosts(hosts map[string]string, defaultHost string) error {
	names := hostNames{byAuthority: map[string]string{}, defaultName: defaultHost}
	for name, authority := range hosts {
		if name == "" || authority == "" {
			return fmt.Errorf("host %q: expected a nonempty name and authority", name)
		}
		if other, duplicate := names.byAuthority[authority]; duplicate {
			return fmt.Errorf("hosts %q and %q have the same authority %s", min(name, other), max(name, other), authority)
		}
		names.byAuthority[authority] = name
	}
	if _, known := hosts[defaultHost]; defaultHost != "" && !known {
		return fmt.Errorf("default host %q is not a named host", defaultHost)
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	h.hosts = names
	return nil
}

// callHost resolves a call's host name, rejecting names SetHosts did not define.
func (n hostNames) callHost(name string) (string, error) {
	if name == "" {
		return n.defaultName, nil
	}
	for _, known := range n.byAuthority {
		if known == name {
			return name, nil
		}
	}
	return "", fmt.Errorf("host %q is not a named host", name)
}

// onHost describes a named host in messages; unnamed hosts add nothing.
func onHost(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf(" on host %q", name)
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
	h.mu.Lock()
	defer h.mu.Unlock()
	compiled, err := compileCalls(calls, h.hosts)
	if err != nil {
		return err
	}
	h.reusable = compiled
	return nil
}

// StartStep reports the previous step's failures and installs the next exchanges.
// Installing them even after failure allows the runner to attempt cleanup.
func (h *Handler) StartStep(label string, calls []Call) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	compiled, err := compileCalls(calls, h.hosts)
	if err != nil {
		return fmt.Errorf("%s: %w", label, err)
	}
	err = h.verify()
	h.label, h.calls, h.failures = label, compiled, nil
	return err
}

// compileCalls validates and encodes the entire set before it can replace an
// active fixture set, so invalid input never installs only some of the calls.
func compileCalls(calls []Call, hosts hostNames) ([]exchange, error) {
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
		host, err := hosts.callHost(call.Request.Host)
		if err != nil {
			return nil, fmt.Errorf("call %d: %w", i+1, err)
		}
		e := exchange{call: call, host: host}
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
		e.matchers, e.expectedBody, err = compileMatchers(call.Request.Matchers, call.Request.Headers, e.expectedBody, e.bodyKind == "json")
		if err != nil {
			return nil, fmt.Errorf("call %d request matchers: %w", i+1, err)
		}
		if call.Response.Body != nil {
			e.responseBody, err = call.Response.Body.literalBytes()
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

// matches checks the whole signature, including the request's host name,
// without consuming the exchange.
func (e *exchange) matches(req *http.Request, host string, body requestBody) bool {
	expected := e.call.Request
	if e.used || e.host != host || expected.Method != req.Method || expected.Path != req.URL.EscapedPath() {
		return false
	}
	if !matchesQuery(expected.Query, e.matchers.query, req.URL.Query()) {
		return false
	}
	if !matchesHeaders(expected.Headers, e.matchers.headers, req.Header) {
		return false
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

// ServeHTTP prefers consumable step fixtures over reusable fallbacks. Matching
// and consumption share a lock so concurrent requests cannot consume a call twice.
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, err := parseRequestBody(req)
	h.mu.Lock()
	defer h.mu.Unlock()
	host := h.hosts.byAuthority[req.Host]
	if err != nil {
		h.fail(w, fmt.Errorf("%s: %s %s%s: invalid request body: %w", h.label, req.Method, req.URL.RequestURI(), onHost(host), err))
		return
	}
	next := -1
	for i := range h.calls {
		if next == -1 && !h.calls[i].used {
			next = i
		}
		if h.calls[i].matches(req, host, body) {
			if h.enforceCallOrder && i != next {
				expected := h.calls[next]
				h.fail(w, fmt.Errorf("%s: out-of-order HTTP call: %s %s%s body=%s matched call %d; next expected call %d: %s %s%s", h.label, req.Method, req.URL.RequestURI(), onHost(host), body.raw, i+1, next+1, expected.call.Request.Method, expected.call.Request.Path, onHost(expected.host)))
				return
			}
			h.calls[i].used = true
			writeResponse(w, &h.calls[i])
			return
		}
	}
	for i := range h.reusable {
		if h.reusable[i].matches(req, host, body) {
			writeResponse(w, &h.reusable[i])
			return
		}
	}
	h.fail(w, fmt.Errorf("%s: unexpected HTTP call: %s %s%s body=%s", h.label, req.Method, req.URL.RequestURI(), onHost(host), body.raw))
}

func writeResponse(w http.ResponseWriter, e *exchange) {
	for key, value := range e.call.Response.Headers {
		w.Header().Set(key, value)
	}
	if e.call.Response.Body != nil && w.Header().Get("Content-Type") == "" {
		kind, _ := e.call.Response.Body.kind()
		contentType := map[string]string{"json": "application/json", "text": "text/plain; charset=utf-8", "base64": "application/octet-stream"}
		w.Header().Set("Content-Type", contentType[kind])
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
			failures = append(failures, fmt.Errorf("%s: unconsumed call %d: %s %s%s", h.label, i+1, e.call.Request.Method, e.call.Request.Path, onHost(e.host)))
		}
	}
	return errors.Join(failures...)
}
