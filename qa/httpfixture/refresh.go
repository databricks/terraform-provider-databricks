package httpfixture

import (
	"fmt"
	"net/http"
	"net/url"
)

// PreserveMatchers carries author-selected volatility annotations to refreshed
// calls using the existing matching rules, never call indices. Interleaving and
// extra reads therefore cannot attach an annotation to another operation.
// Unmatched annotations fail refresh instead of being silently lost.
func PreserveMatchers(calls, previous []Call) error {
	compiled, err := compileCalls(previous)
	if err != nil {
		return err
	}
	retained := make([]bool, len(previous))
	for i := range calls {
		req, body, err := capturedRequest(calls[i].Request)
		if err != nil {
			return err
		}
		selected, keep := refreshedAnnotation(req, body, compiled, retained)
		if selected != -1 {
			calls[i].Request.Matchers = previous[selected].Request.Matchers
			retained[selected] = true
			for key := range previous[selected].Request.Matchers.Headers {
				keep[http.CanonicalHeaderKey(key)] = true
			}
		}
		for key := range calls[i].Request.Headers {
			if !keep[http.CanonicalHeaderKey(key)] {
				delete(calls[i].Request.Headers, key)
			}
		}
	}
	for i, call := range previous {
		if hasMatchers(call.Request.Matchers) && !retained[i] {
			return fmt.Errorf("call %d: could not preserve request matchers; adjust the old signature before recording", i+1)
		}
	}
	return nil
}

func refreshedAnnotation(req *http.Request, body requestBody, previous []exchange, retained []bool) (int, map[string]bool) {
	keep := map[string]bool{"Content-Type": true}
	selected := -1
	for j := range previous {
		old := &previous[j]
		if old.call.Request.Method != req.Method || old.call.Request.Path != req.URL.EscapedPath() {
			continue
		}
		for key := range old.call.Request.Headers {
			keep[http.CanonicalHeaderKey(key)] = true
		}
		if hasMatchers(old.call.Request.Matchers) && old.matches(req, body) &&
			(selected == -1 || retained[selected] && !retained[j]) {
			selected = j
		}
	}
	return selected, keep
}

func hasMatchers(m Matchers) bool { return len(m.Headers)+len(m.Query)+len(m.JSON) != 0 }

func capturedRequest(request Request) (*http.Request, requestBody, error) {
	u, err := url.Parse(request.Path)
	if err != nil {
		return nil, requestBody{}, err
	}
	u.RawQuery = url.Values(request.Query).Encode()
	req := &http.Request{Method: request.Method, URL: u, Header: http.Header{}}
	for name, value := range request.Headers {
		req.Header.Set(name, value)
	}
	var body requestBody
	if request.Body != nil {
		kind, err := request.Body.kind()
		if err != nil {
			return nil, body, err
		}
		if kind == "multipart" {
			body.parts = request.Body.Multipart
		} else {
			body.raw, err = request.Body.literalBytes()
			if err != nil {
				return nil, body, err
			}
			body.json, err = decodeJSON(body.raw)
			body.jsonValid = err == nil
		}
	}
	return req, body, nil
}
