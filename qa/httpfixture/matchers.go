package httpfixture

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"

	"gopkg.in/yaml.v3"
)

// MatcherType names a matcher that takes no argument.
type MatcherType string

const (
	// AnyString requires a present string, including the empty string. It does not
	// capture values or substitute them into later requests or responses.
	AnyString MatcherType = "any_string"
	// Absent requires a header to be missing. Unlisted query parameters and JSON
	// fields are already rejected, so it applies only to headers.
	Absent MatcherType = "absent"
)

// Matcher replaces the exact comparison of one request value. Set either Type or
// Pattern, a regular expression the whole value must match. In YAML, a type is a
// scalar and a pattern is a {pattern: <regular expression>} mapping.
type Matcher struct {
	Type    MatcherType
	Pattern string
}

// UnmarshalYAML rejects mapping keys other than pattern, because yaml.v3 does not
// apply the decoder's KnownFields setting inside custom unmarshalers.
func (m *Matcher) UnmarshalYAML(node *yaml.Node) error {
	switch {
	case node.Kind == yaml.ScalarNode:
		m.Type = MatcherType(node.Value)
		return nil
	case node.Kind == yaml.MappingNode && len(node.Content) == 2 &&
		node.Content[0].Value == "pattern":
		if err := node.Content[1].Decode(&m.Pattern); err != nil || m.Pattern == "" {
			return fmt.Errorf("line %d: pattern must be a nonempty regular expression",
				node.Line)
		}
		return nil
	}
	return fmt.Errorf("line %d: expected a matcher name or "+
		"{pattern: <regular expression>}", node.Line)
}

// MarshalYAML writes the forms UnmarshalYAML reads, so refreshed fixtures keep them.
func (m Matcher) MarshalYAML() (any, error) {
	if m.Pattern != "" {
		return map[string]string{"pattern": m.Pattern}, nil
	}
	return string(m.Type), nil
}

// Matchers override only the selected request values. Headers and query keys
// require exactly one value unless the header must be absent; JSON targets must
// exist in the sample body. Keeping these annotations outside the payload avoids
// reserved JSON keys.
type Matchers struct {
	Headers map[string]Matcher `yaml:"headers,omitempty"`
	Query   map[string]Matcher `yaml:"query,omitempty"`
	JSON    []JSONMatcher      `yaml:"json,omitempty"`
}

// JSONMatcher selects a literal object key or zero-based array index at each
// path segment. An explicit empty path selects the whole JSON body. Set either
// Type or Pattern.
type JSONMatcher struct {
	Path    []any       `yaml:"path"`
	Type    MatcherType `yaml:"type,omitempty"`
	Pattern string      `yaml:"pattern,omitempty"`
}

// compiledMatcher is a validated Matcher. A nil pattern accepts every string.
type compiledMatcher struct {
	absent  bool
	pattern *regexp.Regexp
}

func (m compiledMatcher) matches(value string) bool {
	return m.pattern == nil || m.pattern.MatchString(value)
}

// compiledMatchers keys header matchers by canonical header name.
type compiledMatchers struct {
	headers map[string]compiledMatcher
	query   map[string]compiledMatcher
}

// errMatcherType reports a type the location does not support; callers name the
// location and the types it allows.
var errMatcherType = errors.New("unsupported matcher type")

func compileMatcher(m Matcher, allowAbsent bool) (compiledMatcher, error) {
	switch {
	case m.Pattern != "" && m.Type != "":
		return compiledMatcher{}, errors.New(
			"set either a matcher type or a pattern, not both")
	case m.Pattern != "":
		// Anchoring makes the pattern constrain the whole value, as an exact sample does.
		pattern, err := regexp.Compile("^(?:" + m.Pattern + ")$")
		if err != nil {
			return compiledMatcher{}, fmt.Errorf("invalid pattern %q: %w", m.Pattern, err)
		}
		return compiledMatcher{pattern: pattern}, nil
	case m.Type == AnyString:
		return compiledMatcher{}, nil
	case m.Type == Absent && allowAbsent:
		return compiledMatcher{absent: true}, nil
	case m.Type == Absent:
		return compiledMatcher{}, errors.New("absent applies only to headers; " +
			"unlisted query parameters and JSON fields already fail to match")
	}
	return compiledMatcher{}, errMatcherType
}

func compileMatchers(m Matchers, headerSamples map[string]string, body any,
	hasBody bool) (compiledMatchers, any, error) {
	headers, err := compileHeaderMatchers(m.Headers, headerSamples)
	if err != nil {
		return compiledMatchers{}, nil, err
	}
	result := compiledMatchers{headers: headers, query: map[string]compiledMatcher{}}
	for key, matcher := range m.Query {
		if result.query[key], err = compileKeyMatcher("query", key, matcher); err != nil {
			return compiledMatchers{}, nil, err
		}
	}
	for i, matcher := range m.JSON {
		if !hasBody {
			return compiledMatchers{}, nil, fmt.Errorf("json[%d]: requires body.json", i)
		}
		compiled, err := compileMatcher(Matcher{matcher.Type, matcher.Pattern}, false)
		if matcher.Path == nil || errors.Is(err, errMatcherType) {
			return compiledMatchers{}, nil, fmt.Errorf(
				"json[%d]: requires an explicit path and type any_string or a pattern", i)
		}
		if err != nil {
			return compiledMatchers{}, nil, fmt.Errorf("json[%d]: %w", i, err)
		}
		body, err = installJSONMatcher(body, matcher.Path, compiled)
		if err != nil {
			return compiledMatchers{}, nil, fmt.Errorf("json[%d] path %v: %w",
				i, matcher.Path, err)
		}
	}
	return result, body, nil
}

func compileHeaderMatchers(matchers map[string]Matcher,
	samples map[string]string) (map[string]compiledMatcher, error) {
	sampled := map[string]bool{}
	for key := range samples {
		sampled[http.CanonicalHeaderKey(key)] = true
	}
	result := map[string]compiledMatcher{}
	for key, matcher := range matchers {
		canonical := http.CanonicalHeaderKey(key)
		if _, duplicate := result[canonical]; duplicate {
			return nil, fmt.Errorf("duplicate header matcher %q (case-insensitive)", key)
		}
		compiled, err := compileKeyMatcher("headers", key, matcher)
		if err != nil {
			return nil, err
		}
		if compiled.absent && sampled[canonical] {
			return nil, fmt.Errorf("headers[%q]: remove the header's sample from "+
				"headers; it must be absent", key)
		}
		result[canonical] = compiled
	}
	return result, nil
}

func compileKeyMatcher(location, key string, matcher Matcher) (compiledMatcher, error) {
	allowed := "any_string or a pattern"
	if location == "headers" {
		allowed = "any_string, absent or a pattern"
	}
	compiled, err := compileMatcher(matcher, location == "headers")
	if key == "" || errors.Is(err, errMatcherType) {
		return compiledMatcher{}, fmt.Errorf(
			"%s[%q]: expected a nonempty key and matcher %s, got %q",
			location, key, allowed, matcher.Type)
	}
	if err != nil {
		return compiledMatcher{}, fmt.Errorf("%s[%q]: %w", location, key, err)
	}
	return compiled, nil
}

// installJSONMatcher edits only the normalized copy, never the fixture's sample.
// String/null leaves prevent a broad matcher from hiding an object or array.
func installJSONMatcher(value any, path []any, matcher compiledMatcher) (any, error) {
	if len(path) == 0 {
		if _, duplicate := value.(compiledMatcher); duplicate {
			return nil, fmt.Errorf("duplicate JSON matcher")
		}
		if _, isString := value.(string); !isString && value != nil {
			return nil, fmt.Errorf("matcher target must be a string or null sample")
		}
		return matcher, nil
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
		replacement, err := installJSONMatcher(child, path[1:], matcher)
		object[key] = replacement
		return value, err
	case int:
		array, ok := value.([]any)
		if !ok || key < 0 || key >= len(array) {
			return nil, fmt.Errorf("index %d is outside sample array", key)
		}
		replacement, err := installJSONMatcher(array[key], path[1:], matcher)
		array[key] = replacement
		return value, err
	default:
		return nil, fmt.Errorf("path segments must be string keys or nonnegative integer indices")
	}
}

func matchesJSON(expected, actual any) bool {
	switch value := expected.(type) {
	case compiledMatcher:
		s, isString := actual.(string)
		return isString && value.matches(s)
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

func matchesHeaders(expected map[string]string, matchers map[string]compiledMatcher,
	actual http.Header) bool {
	for key, header := range expected {
		if _, overridden := matchers[http.CanonicalHeaderKey(key)]; overridden {
			continue
		}
		if actual.Get(key) != header {
			return false
		}
	}
	for key, matcher := range matchers {
		values := actual.Values(key)
		if matcher.absent {
			if len(values) != 0 {
				return false
			}
		} else if len(values) != 1 || !matcher.matches(values[0]) {
			return false
		}
	}
	return true
}

func matchesQuery(expected map[string][]string, matchers map[string]compiledMatcher,
	actual url.Values) bool {
	if len(matchers) == 0 {
		return url.Values(expected).Encode() == actual.Encode()
	}
	for key, matcher := range matchers {
		if len(actual[key]) != 1 || !matcher.matches(actual[key][0]) {
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
