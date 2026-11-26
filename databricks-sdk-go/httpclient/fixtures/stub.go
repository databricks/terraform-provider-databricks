package fixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"unicode"
)

// search returns the value of cond evaluated at the nearest letter to index i
// in name. dir determines the direction of search: if true, search forwards, if
// false, search backwards.
func search(name string, cond func(rune) bool, dir bool, i int) bool {
	nameLen := len(name)
	incr := 1
	if !dir {
		incr = -1
	}
	for j := i; j >= 0 && j < nameLen; j += incr {
		if unicode.IsLetter(rune(name[j])) {
			return cond(rune(name[j]))
		}
	}
	return false
}

// checkCondAtNearestLetters returns the value of cond evaluated on the rune at
// index i in name. If that rune is not a letter, search in both directions for
// the nearest letter and return the result of cond on those letters.
func checkCondAtNearestLetters(name string, cond func(rune) bool, i int) bool {
	r := rune(name[i])
	if unicode.IsLetter(r) {
		return cond(r)
	}
	return search(name, cond, true, i) && search(name, cond, false, i)
}

// splitASCII emulates positive lookaheads from JVM regex:
// (?<=[a-z])(?=[A-Z])|(?<=[A-Z])(?=[A-Z][a-z])|([-_\s])
// and converts all words to lower case.
func splitASCII(name string) (w []string) {
	var current []rune
	nameLen := len(name)
	var isPrevUpper, isCurrentUpper, isNextLower, isNextUpper, isNotLastChar bool
	for i := 0; i < nameLen; i++ {
		r := rune(name[i])
		if r == '$' {
			continue
		}
		isCurrentUpper = checkCondAtNearestLetters(name, unicode.IsUpper, i)
		r = unicode.ToLower(r)
		isNextLower = false
		isNextUpper = false
		isNotLastChar = i+1 < nameLen
		if isNotLastChar {
			isNextLower = checkCondAtNearestLetters(name, unicode.IsLower, i+1)
			isNextUpper = checkCondAtNearestLetters(name, unicode.IsUpper, i+1)
		}
		split, before, after := false, false, true
		if isPrevUpper && isCurrentUpper && isNextLower && isNotLastChar {
			split = true
			before = false
			after = true
		}
		if !isCurrentUpper && isNextUpper {
			split = true
			before = true
			after = false
		}
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			split = true
			before = false
			after = false
		}
		if before {
			current = append(current, r)
		}
		if split && len(current) > 0 {
			w = append(w, string(current))
			current = []rune{}
		}
		if after {
			current = append(current, r)
		}
		isPrevUpper = isCurrentUpper
	}
	if len(current) > 0 {
		w = append(w, string(current))
	}
	return w
}

func toPascalCase(s string) string {
	var sb strings.Builder
	for _, w := range splitASCII(s) {
		sb.WriteString(strings.Title(w))
	}
	return sb.String()
}

func resourceFromRequest(req *http.Request) string {
	resource := req.RequestURI
	if resource == "" {
		query := ""
		if req.URL.RawQuery != "" {
			query = "?" + req.URL.RawQuery
		}
		resource = req.URL.Path + query
	}
	return resource
}

func bodyStub(req *http.Request) (string, error) {
	if req.Body == nil {
		return "", nil
	}
	receivedRequest := map[string]any{}
	actualRequestJSON := new(bytes.Buffer)
	_, err := actualRequestJSON.ReadFrom(req.Body)
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}
	expectedRequest := ""
	if actualRequestJSON.Len() == 0 {
		return "", nil
	}
	contentType := req.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		values, err := url.ParseQuery(actualRequestJSON.String())
		if err != nil {
			return "", fmt.Errorf("form: %w", err)
		}
		expectedRequest += "ExpectedRequest: url.Values{\n"
		for k, v := range values {
			expectedRequest += fmt.Sprintf("\t\t\t\"%s\": %#v,\n", k, v)
		}
		expectedRequest += "\t\t},\n\t\t"
		return expectedRequest, nil
	}
	err = json.Unmarshal(actualRequestJSON.Bytes(), &receivedRequest)
	if err != nil {
		return "", fmt.Errorf("unmarshal: %w", err)
	}
	// guessing model name would require going over AST,
	// which is not something i'm willing to write on my weekend
	expectedRequest += "ExpectedRequest: XXX {\n"
	for key, value := range receivedRequest {
		camel := toPascalCase(key)
		expectedRequest += fmt.Sprintf("\t\t\t%s: %#v,\n", camel, value)
	}
	expectedRequest += "\t\t},\n"
	expectedRequest += fmt.Sprintf("\t\t// ExpectedRequest: %#v,\n\t\t", receivedRequest)
	return expectedRequest, nil
}
