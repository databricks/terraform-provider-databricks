package httpfixture

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// Matcher relaxes a selected request value without changing literal payloads.
type Matcher string

// AnyString requires a present string, including the empty string. It does not
// capture values or substitute them into later requests or responses.
const AnyString Matcher = "any_string"

// Matchers override only the selected request values. Headers and query keys
// require exactly one value; JSON targets must exist in the sample body.
// Keeping these annotations outside the payload avoids reserved JSON keys.
type Matchers struct {
	Headers map[string]Matcher `yaml:"headers,omitempty"`
	Query   map[string]Matcher `yaml:"query,omitempty"`
	JSON    []JSONMatcher      `yaml:"json,omitempty"`
}

// JSONMatcher selects a literal object key or zero-based array index at each
// path segment. An explicit empty path selects the whole JSON body.
type JSONMatcher struct {
	Path []any   `yaml:"path"`
	Type Matcher `yaml:"type"`
}

func compileMatchers(m Matchers, body any, hasBody bool) (Matchers, any, error) {
	headers := make(map[string]Matcher, len(m.Headers))
	for key, matcher := range m.Headers {
		canonical := http.CanonicalHeaderKey(key)
		if _, duplicate := headers[canonical]; duplicate {
			return Matchers{}, nil, fmt.Errorf("duplicate header matcher %q (case-insensitive)", key)
		}
		headers[canonical] = matcher
	}
	for location, values := range map[string]map[string]Matcher{"headers": headers, "query": m.Query} {
		for key, matcher := range values {
			if key == "" || matcher != AnyString {
				return Matchers{}, nil, fmt.Errorf("%s[%q]: expected a nonempty key and matcher any_string, got %q", location, key, matcher)
			}
		}
	}
	for i, matcher := range m.JSON {
		if !hasBody {
			return Matchers{}, nil, fmt.Errorf("json[%d]: requires body.json", i)
		}
		if matcher.Path == nil || matcher.Type != AnyString {
			return Matchers{}, nil, fmt.Errorf("json[%d]: requires an explicit path and type any_string", i)
		}
		var err error
		body, err = installJSONMatcher(body, matcher.Path)
		if err != nil {
			return Matchers{}, nil, fmt.Errorf("json[%d] path %v: %w", i, matcher.Path, err)
		}
	}
	return Matchers{Headers: headers, Query: m.Query}, body, nil
}

// installJSONMatcher edits only the normalized copy, never the fixture's sample.
// String/null leaves prevent a broad matcher from hiding an object or array.
func installJSONMatcher(value any, path []any) (any, error) {
	if len(path) == 0 {
		if _, duplicate := value.(Matcher); duplicate {
			return nil, fmt.Errorf("duplicate JSON matcher")
		}
		if _, isString := value.(string); !isString && value != nil {
			return nil, fmt.Errorf("any_string target must be a string or null sample")
		}
		return AnyString, nil
	}
	switch key := path[0].(type) {
	case string:
		object, ok := value.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("key %q requires an object", key)
		}
		child, exists := object[key]
		if !exists {
			return nil, fmt.Errorf("key %q is missing from sample", key)
		}
		replacement, err := installJSONMatcher(child, path[1:])
		object[key] = replacement
		return value, err
	case int:
		array, ok := value.([]any)
		if !ok || key < 0 || key >= len(array) {
			return nil, fmt.Errorf("index %d is outside sample array", key)
		}
		replacement, err := installJSONMatcher(array[key], path[1:])
		array[key] = replacement
		return value, err
	default:
		return nil, fmt.Errorf("path segments must be string keys or nonnegative integer indices")
	}
}

func matchesJSON(expected, actual any) bool {
	switch value := expected.(type) {
	case Matcher:
		_, isString := actual.(string)
		return value == AnyString && isString
	case map[string]any:
		object, ok := actual.(map[string]any)
		if !ok || len(value) != len(object) {
			return false
		}
		for key, expectedChild := range value {
			child, exists := object[key]
			if !exists || !matchesJSON(expectedChild, child) {
				return false
			}
		}
		return true
	case []any:
		array, ok := actual.([]any)
		if !ok || len(value) != len(array) {
			return false
		}
		for i, expectedChild := range value {
			if !matchesJSON(expectedChild, array[i]) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(expected, actual)
	}
}

func matchesQuery(expected map[string][]string, matchers map[string]Matcher, actual url.Values) bool {
	if len(matchers) == 0 {
		return url.Values(expected).Encode() == actual.Encode()
	}
	for key := range matchers {
		if len(actual[key]) != 1 {
			return false
		}
	}
	// Remove annotated keys from both sides so all remaining parameters retain
	// exact matching, including repeated-value order and unexpected parameters.
	exact := func(values map[string][]string) string {
		filtered := make(url.Values, len(values))
		for key, values := range values {
			if _, overridden := matchers[key]; !overridden {
				filtered[key] = values
			}
		}
		return filtered.Encode()
	}
	return exact(expected) == exact(actual)
}
