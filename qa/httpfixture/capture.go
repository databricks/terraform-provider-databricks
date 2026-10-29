package httpfixture

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// CapturedStep retains dispatch order, even when concurrent replies finish in
// another order. Recording that order does not opt a test into ordered replay.
type CapturedStep struct {
	Label string
	Calls []Call
}

// CaptureTransport observes HTTP at the same boundary used by replay. It does
// not authenticate, retry, choose hosts, or write files; those belong to callers.
// Response bodies are buffered and restored before the SDK sees them.
type CaptureTransport struct {
	base     http.RoundTripper
	mu       sync.Mutex
	steps    []CapturedStep
	inFlight int
	failures []error
}

// NewCaptureTransport wraps the caller's transport. Snapshots contain API
// payloads in memory; the owning recorder must sanitize them before persistence.
func NewCaptureTransport(base http.RoundTripper) *CaptureTransport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &CaptureTransport{base: base}
}

// StartStep refuses transitions during a request so a late reply cannot be
// attributed to the next step. Capture failures are permanent for this session.
func (c *CaptureTransport) StartStep(label string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := c.verify(); err != nil {
		return err
	}
	c.steps = append(c.steps, CapturedStep{Label: label})
	return nil
}

func (c *CaptureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	step, index, err := c.reserve()
	if err != nil {
		return nil, err
	}
	call, response, err := c.exchange(req)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.inFlight--
	if err != nil {
		// Transport errors can contain credential-bearing URLs. Keep them out of
		// the collected diagnostics; the SDK still receives the original error.
		c.failures = append(c.failures, fmt.Errorf("%s: could not capture call %d", c.steps[step].Label, index+1))
	} else {
		c.steps[step].Calls[index] = call
	}
	return response, err
}

func (c *CaptureTransport) reserve() (int, int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.steps) == 0 {
		return 0, 0, errors.New("HTTP capture requires StartStep before requests")
	}
	step := len(c.steps) - 1
	index := len(c.steps[step].Calls)
	c.steps[step].Calls = append(c.steps[step].Calls, Call{})
	c.inFlight++
	return step, index, nil
}

func (c *CaptureTransport) exchange(req *http.Request) (Call, *http.Response, error) {
	request, err := captureRequest(req)
	if err != nil {
		return Call{}, nil, err
	}
	response, err := c.base.RoundTrip(req)
	if err != nil {
		return Call{}, response, err
	}
	body, err := readAndRestore(&response.Body)
	if err != nil {
		return Call{}, response, err
	}
	return Call{Request: request, Response: Response{
		Status: response.StatusCode, Headers: captureHeaders(response.Header), Body: captureBody(body),
	}}, response, nil
}

// Snapshot returns an independent copy so normalization cannot race with capture
// or alter a later snapshot. An incomplete exchange never becomes a fixture.
func (c *CaptureTransport) Snapshot() ([]CapturedStep, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := c.verify(); err != nil {
		return nil, err
	}
	// YAML is also the persistence representation, including binary bodies and
	// exact large integers; use it instead of a shallow copy of nested maps.
	return cloneCapturedSteps(c.steps)
}

func (c *CaptureTransport) Verify() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.verify()
}

func (c *CaptureTransport) verify() error {
	if c.inFlight != 0 {
		return fmt.Errorf("HTTP capture has %d requests still in flight", c.inFlight)
	}
	return errors.Join(c.failures...)
}

func readAndRestore(body *io.ReadCloser) ([]byte, error) {
	if *body == nil {
		return nil, nil
	}
	data, err := io.ReadAll(*body)
	closeErr := (*body).Close()
	*body = io.NopCloser(bytes.NewReader(data))
	return data, errors.Join(err, closeErr)
}
