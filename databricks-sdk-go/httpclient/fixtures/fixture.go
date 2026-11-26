package fixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/httpclient/traceparent"
)

// HTTPFixture defines request structure for test
type HTTPFixture struct {
	Method       string
	Resource     string
	MatchAny     bool
	ReuseRequest bool

	Response        any
	Status          int
	ResponseHeaders map[string][]string
	ExpectedRequest any
	ExpectedHeaders map[string]string
	PassFile        string
}

func (f HTTPFixture) Match(method, resource string) bool {
	if f.MatchAny {
		return true
	}
	return method == f.Method && resource == f.Resource
}

func (f HTTPFixture) AssertHeaders(req *http.Request) error {
	if f.ExpectedHeaders == nil {
		return nil
	}
	actualHeaders := map[string]string{}
	// remove user agent & traceparent from comparison, as it'll make fixtures too difficult
	// to maintain in the long term
	toSkip := map[string]struct{}{
		textproto.CanonicalMIMEHeaderKey("User-Agent"):                   {},
		textproto.CanonicalMIMEHeaderKey(traceparent.TRACEPARENT_HEADER): {},
	}
	for k := range req.Header {
		if _, skip := toSkip[k]; !skip {
			actualHeaders[k] = req.Header.Get(k)
		}
	}
	if !reflect.DeepEqual(f.ExpectedHeaders, actualHeaders) {
		expectedJSON, _ := json.MarshalIndent(f.ExpectedHeaders, "", "  ")
		actualJSON, _ := json.MarshalIndent(actualHeaders, "", "  ")
		return fmt.Errorf("want %s, got %s", expectedJSON, actualJSON)
	}
	return nil
}

func (f HTTPFixture) AssertRequest(req *http.Request) error {
	if f.ExpectedRequest == nil {
		return nil
	}
	actualRequestBuf := new(bytes.Buffer)
	_, err := actualRequestBuf.ReadFrom(req.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}
	qs, ok := f.ExpectedRequest.(url.Values)
	if ok {
		expectedQS := qs.Encode()
		actualQS := actualRequestBuf.String()
		if expectedQS != actualQS {
			return fmt.Errorf("want %s, got %s", expectedQS, actualQS)
		}
		return nil
	}
	// do pretty-priting, so that it's easier to debug failures
	expectedRequestJSON, err := json.MarshalIndent(f.ExpectedRequest, "", "  ")
	if err != nil {
		return fmt.Errorf("unmarshal expected: %w", err)
	}
	var expectedANY, actualANY any
	err = json.Unmarshal(expectedRequestJSON, &expectedANY)
	if err != nil {
		return fmt.Errorf("marshal expected to any: %w", err)
	}
	err = json.Unmarshal(actualRequestBuf.Bytes(), &actualANY)
	if err != nil {
		return fmt.Errorf("marshal actual to any: %w", err)
	}
	if !reflect.DeepEqual(expectedANY, actualANY) {
		actualRequestJSON, _ := json.MarshalIndent(actualANY, "", "  ")
		return fmt.Errorf("want %s, got %s", expectedRequestJSON, actualRequestJSON)
	}
	return nil
}

func (f HTTPFixture) replyWith(req *http.Request, body string) (*http.Response, error) {
	return &http.Response{
		Request:    req,
		Proto:      "HTTP/1.1",
		StatusCode: f.Status,
		Status:     http.StatusText(f.Status),
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     f.ResponseHeaders,
	}, nil
}

func (f HTTPFixture) Reply(req *http.Request) (*http.Response, error) {
	if f.Status == 0 {
		f.Status = 200
	}
	if f.PassFile != "" {
		raw, err := os.ReadFile(f.PassFile)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f.PassFile, err)
		}
		return f.replyWith(req, string(raw))
	}
	if f.Response == nil {
		return f.replyWith(req, "")
	}
	if alreadyJSON, ok := f.Response.(string); ok {
		return f.replyWith(req, alreadyJSON)
	}
	responseBytes, err := json.Marshal(f.Response)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	return f.replyWith(req, string(responseBytes))
}
