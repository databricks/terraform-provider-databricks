package httpfixture

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
				c.Request.Matchers = Matchers{Headers: map[string]Matcher{"x-request-id": AnyString}, Query: map[string]Matcher{"id": AnyString}}
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
		{"unknown header", Matchers{Headers: map[string]Matcher{"X-Id": "any"}}, "expected a nonempty key and matcher any_string"},
		{"empty header", Matchers{Headers: map[string]Matcher{"": AnyString}}, "expected a nonempty key"},
		{"duplicate header", Matchers{Headers: map[string]Matcher{"X-Id": AnyString, "x-id": AnyString}}, "duplicate header matcher"},
		{"unknown query", Matchers{Query: map[string]Matcher{"id": "*"}}, "expected a nonempty key and matcher any_string"},
		{"empty query", Matchers{Query: map[string]Matcher{"": AnyString}}, "expected a nonempty key"},
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
