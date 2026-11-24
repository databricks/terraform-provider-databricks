package httpfixture

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func volatileCall(result string) Call {
	c := call("POST", "/requests", map[string]any{
		"id": "recorded-id", "name": "fixed", "number": json.Number("9007199254740993"),
		"items": []any{map[string]any{"token": nil, "enabled": true}, "literal"},
	}, result)
	c.Request.Matchers = Matchers{JSON: []JSONMatcher{
		{Path: []any{"id"}, Type: AnyString},
		{Path: []any{"items", 0, "token"}, Type: AnyString},
	}}
	return c
}

func TestJSONMatchersRelaxOnlySelectedLeaves(t *testing.T) {
	valid := `{"id":"new-id","name":"fixed","number":9007199254740993,"items":[{"token":"new-token","enabled":true},"literal"]}`
	for _, tc := range []struct {
		name, body string
		matches    bool
	}{
		{"different strings", valid, true},
		{"empty string", strings.Replace(valid, `"new-id"`, `""`, 1), true},
		{"numeric equivalence", strings.Replace(valid, "9007199254740993", "9007199254740993.0", 1), true},
		{"missing value", strings.Replace(valid, `"id":"new-id",`, "", 1), false},
		{"null value", strings.Replace(valid, `"new-id"`, `null`, 1), false},
		{"number value", strings.Replace(valid, `"new-id"`, `123`, 1), false},
		{"boolean value", strings.Replace(valid, `"new-id"`, `true`, 1), false},
		{"object value", strings.Replace(valid, `"new-id"`, `{}`, 1), false},
		{"array value", strings.Replace(valid, `"new-id"`, `[]`, 1), false},
		{"other field", strings.Replace(valid, `"fixed"`, `"changed"`, 1), false},
		{"large number", strings.Replace(valid, "9007199254740993", "9007199254740992", 1), false},
		{"nested field", strings.Replace(valid, `true`, `false`, 1), false},
		{"missing nested key", strings.Replace(valid, `"token":"new-token",`, "", 1), false},
		{"extra key", strings.Replace(valid, `"name":`, `"extra":true,"name":`, 1), false},
		{"replaced key", strings.Replace(valid, `"id":`, `"other_id":`, 1), false},
		{"extra array item", strings.Replace(valid, `"literal"]`, `"literal",0]`, 1), false},
		{"reordered array", strings.Replace(valid, `{"token":"new-token","enabled":true},"literal"`, `"literal",{"token":"new-token","enabled":true}`, 1), false},
		{"missing body", "", false},
		{"null body", "null", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler()
			c := volatileCall("ok")
			require.NoError(t, h.StartStep("create", []Call{c}))
			w := send(h, "POST", "/requests", tc.body)
			if tc.matches {
				require.Equal(t, 200, w.Code, w.Body.String())
				require.NoError(t, h.Verify())
			} else {
				require.Equal(t, 400, w.Code)
				require.ErrorContains(t, h.Verify(), "create: unexpected HTTP call")
				require.ErrorContains(t, h.Verify(), "unconsumed call 1")
				require.Equal(t, 200, send(h, "POST", "/requests", valid).Code)
				require.NotContains(t, h.Verify().Error(), "unconsumed call")
			}
			// Compilation must not replace samples with internal matcher markers.
			require.Equal(t, "recorded-id", c.Request.Body.JSON.(map[string]any)["id"])
		})
	}
}

func TestJSONMatcherPathsAreLiteralAndCanSelectRoot(t *testing.T) {
	for _, tc := range []struct {
		name string
		body any
		path []any
		json string
	}{
		{"root", "recorded", []any{}, `"new"`},
		{"root null sample", nil, []any{}, `""`},
		{"root array", []any{"recorded"}, []any{0}, `["new"]`},
		{"dotted key", map[string]any{"a.b[0]": "recorded"}, []any{"a.b[0]"}, `{"a.b[0]":"new"}`},
		{"numeric map key", map[string]any{"0": "recorded"}, []any{"0"}, `{"0":"new"}`},
		{"empty key", map[string]any{"": "recorded"}, []any{""}, `{"":"new"}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := call("POST", "/value", nil, "ok")
			c.Request.Body = &Body{JSON: tc.body}
			c.Request.Matchers.JSON = []JSONMatcher{{Path: tc.path, Type: AnyString}}
			h := NewHandler()
			require.NoError(t, h.StartStep("create", []Call{c}))
			require.Equal(t, 200, send(h, "POST", "/value", tc.json).Code)
			require.NoError(t, h.Verify())
		})
	}
}

func TestHeaderAndQueryMatchersRequireOnePresentValue(t *testing.T) {
	for _, tc := range []struct {
		name, target string
		headers      http.Header
		matches      bool
	}{
		{"new values", "/requests?fixed=one&fixed=two&id=new", http.Header{"X-Request-Id": {"new"}, "X-Exact": {"fixed"}, "X-Extra": {"ignored"}}, true},
		{"empty values", "/requests?fixed=one&fixed=two&id=", http.Header{"X-Request-Id": {""}, "X-Exact": {"fixed"}}, true},
		{"absent header", "/requests?fixed=one&fixed=two&id=new", http.Header{"X-Exact": {"fixed"}}, false},
		{"repeated header", "/requests?fixed=one&fixed=two&id=new", http.Header{"X-Request-Id": {"a", "b"}, "X-Exact": {"fixed"}}, false},
		{"changed exact header", "/requests?fixed=one&fixed=two&id=new", http.Header{"X-Request-Id": {"new"}, "X-Exact": {"wrong"}}, false},
		{"absent query", "/requests?fixed=one&fixed=two", http.Header{"X-Request-Id": {"new"}, "X-Exact": {"fixed"}}, false},
		{"repeated query", "/requests?fixed=one&fixed=two&id=a&id=b", http.Header{"X-Request-Id": {"new"}, "X-Exact": {"fixed"}}, false},
		{"changed exact query", "/requests?fixed=two&fixed=one&id=new", http.Header{"X-Request-Id": {"new"}, "X-Exact": {"fixed"}}, false},
		{"extra query", "/requests?fixed=one&fixed=two&id=new&extra=x", http.Header{"X-Request-Id": {"new"}, "X-Exact": {"fixed"}}, false},
		{"missing exact query", "/requests?id=new", http.Header{"X-Request-Id": {"new"}, "X-Exact": {"fixed"}}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, withSample := range []bool{false, true} {
				c := call("GET", "/requests", nil, "ok")
				c.Request.Query = map[string][]string{"fixed": {"one", "two"}}
				c.Request.Headers = map[string]string{"X-Exact": "fixed"}
				if withSample {
					c.Request.Query["id"] = []string{"recorded"}
					c.Request.Headers["X-REQUEST-ID"] = "recorded"
				}
				c.Request.Matchers = Matchers{Headers: map[string]Matcher{"x-request-id": {Type: AnyString}}, Query: map[string]Matcher{"id": {Type: AnyString}}}
				h := NewHandler()
				require.NoError(t, h.StartStep("read", []Call{c}))
				req := httptest.NewRequest("GET", tc.target, nil)
				req.Header = tc.headers
				w := httptest.NewRecorder()
				h.ServeHTTP(w, req)
				if tc.matches {
					require.Equal(t, 200, w.Code)
					require.NoError(t, h.Verify())
				} else {
					require.Equal(t, 400, w.Code)
					require.ErrorContains(t, h.Verify(), "unconsumed call 1")
				}
			}
		})
	}
}

func TestMatcherSyntaxNeverChangesLiteralPayloads(t *testing.T) {
	h := NewHandler()
	body := map[string]any{"id": "any_string", "token": "****", "matchers": map[string]any{"type": "any_string"}}
	c := call("POST", "/literal", body, "any_string")
	c.Request.Query = map[string][]string{"id": {"****"}}
	c.Request.Headers = map[string]string{"X-Request-Id": "any_string"}
	require.NoError(t, h.StartStep("literal", []Call{c}))
	req := httptest.NewRequest("POST", "/literal?id=****", strings.NewReader(`{"id":"any_string","token":"****","matchers":{"type":"any_string"}}`))
	req.Header.Set("X-Request-Id", "any_string")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	require.Equal(t, `"any_string"`, w.Body.String())
	require.NoError(t, h.Verify())
	require.NoError(t, h.StartStep("not magic", []Call{call("POST", "/literal", body, "ok")}))
	require.Equal(t, 400, send(h, "POST", "/literal", `{"id":"changed","token":"anything","matchers":{"type":"any_string"}}`).Code)
}

func TestRejectInvalidMatchersWithoutReplacingFixtures(t *testing.T) {
	for _, tc := range []struct {
		name     string
		matchers Matchers
		message  string
	}{
		{"unknown header", Matchers{Headers: map[string]Matcher{"X-Id": {Type: "any"}}}, "expected a nonempty key and matcher any_string"},
		{"empty header", Matchers{Headers: map[string]Matcher{"": {Type: AnyString}}}, "expected a nonempty key"},
		{"duplicate header", Matchers{Headers: map[string]Matcher{"X-Id": {Type: AnyString}, "x-id": {Type: AnyString}}}, "duplicate header matcher"},
		{"unknown query", Matchers{Query: map[string]Matcher{"id": {Type: "*"}}}, "expected a nonempty key and matcher any_string"},
		{"empty query", Matchers{Query: map[string]Matcher{"": {Type: AnyString}}}, "expected a nonempty key"},
		{"unknown JSON", Matchers{JSON: []JSONMatcher{{Path: []any{"id"}, Type: "any"}}}, "explicit path and type any_string"},
		{"missing path", Matchers{JSON: []JSONMatcher{{Type: AnyString}}}, "explicit path and type any_string"},
		{"missing key", Matchers{JSON: []JSONMatcher{{Path: []any{"missing"}, Type: AnyString}}}, "is missing from sample"},
		{"negative index", Matchers{JSON: []JSONMatcher{{Path: []any{"items", -1}, Type: AnyString}}}, "outside sample array"},
		{"large index", Matchers{JSON: []JSONMatcher{{Path: []any{"items", 2}, Type: AnyString}}}, "outside sample array"},
		{"fraction index", Matchers{JSON: []JSONMatcher{{Path: []any{"items", 0.5}, Type: AnyString}}}, "path segments must be"},
		{"null segment", Matchers{JSON: []JSONMatcher{{Path: []any{nil}, Type: AnyString}}}, "path segments must be"},
		{"wrong container", Matchers{JSON: []JSONMatcher{{Path: []any{"id", "child"}, Type: AnyString}}}, "requires an object"},
		{"object target", Matchers{JSON: []JSONMatcher{{Path: []any{"items", 0}, Type: AnyString}}}, "string or null sample"},
		{"array target", Matchers{JSON: []JSONMatcher{{Path: []any{"items"}, Type: AnyString}}}, "string or null sample"},
		{"number target", Matchers{JSON: []JSONMatcher{{Path: []any{"number"}, Type: AnyString}}}, "string or null sample"},
		{"duplicate path", Matchers{JSON: []JSONMatcher{{Path: []any{"id"}, Type: AnyString}, {Path: []any{"id"}, Type: AnyString}}}, "duplicate JSON matcher"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler()
			original := call("GET", "/original", nil, "original")
			require.NoError(t, h.StartStep("original", []Call{original}))
			require.NoError(t, h.SetReusableFixtures([]Call{original}))
			invalid := volatileCall("invalid")
			invalid.Request.Matchers = tc.matchers
			require.ErrorContains(t, h.StartStep("invalid", []Call{original, invalid}), "invalid: call 2 request matchers:")
			require.ErrorContains(t, h.SetReusableFixtures([]Call{invalid}), tc.message)
			for range 2 {
				require.Equal(t, `"original"`, send(h, "GET", "/original", "").Body.String())
			}
			require.NoError(t, h.Verify())
		})
	}
	h := NewHandler()
	c := call("GET", "/no-body", nil, "ok")
	c.Request.Matchers.JSON = []JSONMatcher{{Path: []any{}, Type: AnyString}}
	require.ErrorContains(t, h.StartStep("invalid", []Call{c}), "requires body.json")
}

func TestMatchersKeepFIFOOrderingAndReusableFallback(t *testing.T) {
	for _, ordered := range []bool{false, true} {
		h := NewHandler()
		h.SetEnforceCallOrder(ordered)
		first, second := call("POST", "/id", "captured-one", "first"), call("POST", "/id", "captured-two", "second")
		first.Request.Matchers.JSON = []JSONMatcher{{Path: []any{}, Type: AnyString}}
		second.Request.Matchers.JSON = first.Request.Matchers.JSON
		require.NoError(t, h.SetReusableFixtures([]Call{first}))
		require.NoError(t, h.StartStep("create", []Call{first, call("GET", "/middle", nil, "middle"), second}))
		require.Equal(t, `"first"`, send(h, "POST", "/id", `"random-one"`).Body.String())
		if ordered {
			require.Equal(t, 400, send(h, "POST", "/id", `"random-two"`).Code)
			require.ErrorContains(t, h.Verify(), "out-of-order HTTP call")
		} else {
			require.Equal(t, `"second"`, send(h, "POST", "/id", `"random-two"`).Body.String())
		}
		require.Equal(t, `"middle"`, send(h, "GET", "/middle", "").Body.String())
		if ordered {
			require.Equal(t, `"second"`, send(h, "POST", "/id", `"random-three"`).Body.String())
		} else {
			require.NoError(t, h.Verify())
		}
		for range 2 {
			require.Equal(t, `"first"`, send(h, "POST", "/id", `"random-four"`).Body.String())
		}
		if ordered {
			require.NotContains(t, h.Verify().Error(), "unconsumed call")
		}
	}
}

func TestAbsentHeaderMatcherRequiresMissingHeader(t *testing.T) {
	for _, tc := range []struct {
		name    string
		headers http.Header
		matches bool
	}{
		{"missing", http.Header{"X-Other": {"ignored"}}, true},
		{"empty value", http.Header{"X-Workspace-Id": {""}}, false},
		{"value", http.Header{"X-Workspace-Id": {"12345"}}, false},
		{"repeated", http.Header{"X-Workspace-Id": {"1", "2"}}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := call("GET", "/items", nil, "ok")
			absent := Matcher{Type: Absent}
			c.Request.Matchers.Headers = map[string]Matcher{"x-workspace-id": absent}
			h := NewHandler()
			require.NoError(t, h.StartStep("read", []Call{c}))
			req := httptest.NewRequest("GET", "/items", nil)
			req.Header = tc.headers
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
			require.Equal(t, tc.matches, w.Code == 200, w.Body.String())
		})
	}
}

// patternRequest sends value in one location of a request whose sample holds
// "captured" there, with a pattern matcher on that location.
func patternRequest(location, pattern, value string) (Call, *http.Request) {
	c := call("POST", "/items", map[string]any{"trace": "captured"}, "ok")
	body := `{"trace":"captured"}`
	switch location {
	case "header":
		c.Request.Headers = map[string]string{"Traceparent": "captured"}
		c.Request.Matchers.Headers = map[string]Matcher{"traceparent": {Pattern: pattern}}
	case "query":
		c.Request.Query = map[string][]string{"trace": {"captured"}}
		c.Request.Matchers.Query = map[string]Matcher{"trace": {Pattern: pattern}}
	case "json":
		c.Request.Matchers.JSON = []JSONMatcher{{Path: []any{"trace"}, Pattern: pattern}}
		body = `{"trace":"` + value + `"}`
	}
	req := httptest.NewRequest("POST", "/items", strings.NewReader(body))
	switch location {
	case "header":
		req.Header.Set("traceparent", value)
	case "query":
		req.URL.RawQuery = url.Values{"trace": {value}}.Encode()
	}
	return c, req
}

func TestPatternMatchersRequireWholeValueMatch(t *testing.T) {
	for _, tc := range []struct {
		name, value string
		matches     bool
	}{
		{"whole value", "00-0af7-01", true},
		{"prefix only", "00-0af7-01-extra", false},
		{"suffix only", "x00-0af7-01", false},
		{"wrong format", "00-XYZW-01", false},
		{"empty", "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, location := range []string{"header", "query", "json"} {
				c, req := patternRequest(location, "00-[0-9a-f]{4}-01", tc.value)
				h := NewHandler()
				require.NoError(t, h.StartStep("create", []Call{c}))
				w := httptest.NewRecorder()
				h.ServeHTTP(w, req)
				require.Equal(t, tc.matches, w.Code == 200, "%s: %s",
					location, w.Body.String())
			}
		})
	}
}

func TestPatternMatchersRequireOnePresentString(t *testing.T) {
	header := call("GET", "/items", nil, "ok")
	header.Request.Matchers.Headers = map[string]Matcher{"X-Id": {Pattern: ".*"}}
	for _, headers := range []http.Header{{}, {"X-Id": {"a", "b"}}} {
		h := NewHandler()
		require.NoError(t, h.StartStep("read", []Call{header}))
		req := httptest.NewRequest("GET", "/items", nil)
		req.Header = headers
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		require.Equal(t, 400, w.Code)
	}
	for _, body := range []string{`{"id":null}`, `{"id":1}`, `{}`} {
		c := call("POST", "/items", map[string]any{"id": "captured"}, "ok")
		c.Request.Matchers.JSON = []JSONMatcher{{Path: []any{"id"}, Pattern: ".*"}}
		h := NewHandler()
		require.NoError(t, h.StartStep("create", []Call{c}))
		require.Equal(t, 400, send(h, "POST", "/items", body).Code, body)
	}
}

func TestRejectInvalidAbsentAndPatternMatchers(t *testing.T) {
	headers := func(m Matcher) Matchers {
		return Matchers{Headers: map[string]Matcher{"X-Id": m}}
	}
	query := func(m Matcher) Matchers { return Matchers{Query: map[string]Matcher{"id": m}} }
	json := func(m JSONMatcher) Matchers { return Matchers{JSON: []JSONMatcher{m}} }
	for _, tc := range []struct {
		name     string
		samples  map[string]string
		matchers Matchers
		message  string
	}{
		{"absent query", nil, query(Matcher{Type: Absent}),
			`query["id"]: absent applies only to headers`},
		{"absent JSON", nil, json(JSONMatcher{Path: []any{"id"}, Type: Absent}),
			"json[0]: absent applies only to headers"},
		{"absent Transfer-Encoding", nil, Matchers{Headers: map[string]Matcher{
			"transfer-encoding": {Type: Absent}}}, `headers["transfer-encoding"]: net/http ` +
			"keeps Host, Transfer-Encoding and Trailer outside the request headers"},
		{"absent Host", nil, Matchers{Headers: map[string]Matcher{"Host": {Type: Absent}}},
			`headers["Host"]: net/http keeps Host`},
		{"Trailer pattern", nil, Matchers{Headers: map[string]Matcher{
			"Trailer": {Pattern: ".*"}}}, `headers["Trailer"]: net/http keeps Host`},
		{"absent with sample", map[string]string{"x-id": "1"}, headers(Matcher{Type: Absent}),
			`headers["X-Id"]: remove the header's sample from headers; it must be absent`},
		{"invalid header pattern", nil, headers(Matcher{Pattern: "("}),
			`headers["X-Id"]: invalid pattern "("`},
		// Wrapping this in ^(?:...)$ would parse as ^(?:foo)|(?:bar)$.
		{"pattern valid only when wrapped", nil, headers(Matcher{Pattern: "foo)|(?:bar"}),
			`headers["X-Id"]: invalid pattern "foo)|(?:bar": error parsing regexp: unexpected )`},
		{"invalid JSON pattern", nil, json(JSONMatcher{Path: []any{"id"}, Pattern: "["}),
			`json[0]: invalid pattern "["`},
		{"type and pattern", nil, query(Matcher{Type: AnyString, Pattern: ".*"}),
			`query["id"]: set either a matcher type or a pattern, not both`},
		{"JSON type and pattern", nil,
			json(JSONMatcher{Path: []any{"id"}, Type: AnyString, Pattern: ".*"}),
			"json[0]: set either a matcher type or a pattern, not both"},
		{"unknown header type", nil, headers(Matcher{Type: "missing"}),
			`headers["X-Id"]: expected a nonempty key and matcher any_string, absent or ` +
				`a pattern, got "missing"`},
		{"pattern on object", nil, json(JSONMatcher{Path: []any{"items", 0}, Pattern: "."}),
			"json[0] path [items 0]: matcher target must be a string or null sample"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := volatileCall("invalid")
			c.Request.Headers = tc.samples
			c.Request.Matchers = tc.matchers
			require.ErrorContains(t, NewHandler().StartStep("create", []Call{c}),
				"create: call 1 request matchers: "+tc.message)
		})
	}
}

func TestMatcherYAMLForms(t *testing.T) {
	var m Matchers
	input := "headers: {A: any_string, B: absent, C: {pattern: '^x$'}}\n" +
		"json: [{path: [id], pattern: '[0-9]+'}, {path: [name], type: any_string}]\n"
	require.NoError(t, yaml.Unmarshal([]byte(input), &m))
	require.Equal(t, Matchers{
		Headers: map[string]Matcher{
			"A": {Type: AnyString}, "B": {Type: Absent}, "C": {Pattern: "^x$"},
		},
		JSON: []JSONMatcher{
			{Path: []any{"id"}, Pattern: "[0-9]+"},
			{Path: []any{"name"}, Type: AnyString},
		},
	}, m)
	output, err := yaml.Marshal(m)
	require.NoError(t, err)
	var roundTrip Matchers
	require.NoError(t, yaml.Unmarshal(output, &roundTrip))
	require.Equal(t, m, roundTrip)
	require.Contains(t, string(output), "C:\n        pattern: ^x$")
	for _, invalid := range []string{
		"{pattern: ''}", "{regex: x}", "{pattern: x, type: any_string}", "[any_string]",
	} {
		err := yaml.Unmarshal([]byte("headers: {A: "+invalid+"}"), &m)
		require.ErrorContains(t, err, "line 1:", invalid)
	}
}
