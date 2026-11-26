package fixtures

import (
	"fmt"
	"net/http"
)

var Failures = SliceTransport{
	{
		MatchAny:     true,
		ReuseRequest: true,
		Status:       418,
		Response:     `{"error_code": "NONSENSE", "message": "I'm a teapot"}`,
	},
}

type SliceTransport []HTTPFixture

func (fixtures SliceTransport) SkipRetryOnIO() bool {
	return true
}

func (fixtures SliceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resource := resourceFromRequest(req)
	for i, f := range fixtures {
		if !f.Match(req.Method, resource) {
			continue
		}
		if f.Status == 0 {
			f.Status = 200
		}
		err := f.AssertRequest(req)
		if err != nil {
			return nil, fmt.Errorf("expected: %w", err)
		}
		// Reset the request if it is already used
		if !f.ReuseRequest {
			fixtures[i] = HTTPFixture{}
		}
		return f.Reply(req)
	}
	expectedRequest, err := bodyStub(req)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	// whitespace in this string is very important for unit tests
	stub := fmt.Sprintf(`{
		Method:   "%s",
		Resource: "%s",
		%sResponse: XXX {
			// fill in specific fields...
		},
	},`, req.Method, resource, expectedRequest)
	return nil, fmt.Errorf("missing stub, please add: %s", stub)
}
