package fixtures

import (
	"fmt"
	"net/http"
)

// MappingTransport stubs are 2 lines shorter than [SliceTransport]
type MappingTransport map[string]HTTPFixture

func (fixtures MappingTransport) SkipRetryOnIO() bool {
	return true
}

func (fixtures MappingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	key := fmt.Sprintf("%s %s", req.Method, resourceFromRequest(req))
	f, ok := fixtures[key]
	if ok {
		err := f.AssertHeaders(req)
		if err != nil {
			return nil, fmt.Errorf("headers: %w", err)
		}
		err = f.AssertRequest(req)
		if err != nil {
			return nil, fmt.Errorf("body: %w", err)
		}
		return f.Reply(req)
	}
	expectedRequest, err := bodyStub(req)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	// whitespace in this string is very important for unit tests
	stub := fmt.Sprintf(`"%s": {
		%sResponse: XXX {
			// fill in specific fields...
		},
	},`, key, expectedRequest)
	return nil, fmt.Errorf("missing stub, please add: %s", stub)
}
