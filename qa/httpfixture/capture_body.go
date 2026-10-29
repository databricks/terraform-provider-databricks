package httpfixture

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

func captureRequest(req *http.Request) (Request, error) {
	body, err := readAndRestore(&req.Body)
	if err != nil {
		return Request{}, err
	}
	result := Request{Method: req.Method, Path: req.URL.EscapedPath(), Query: req.URL.Query(),
		Headers: captureRequestHeaders(req.Header), Body: captureBody(body)}
	mediaType, params, _ := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if mediaType == "multipart/form-data" {
		parts, err := readMultipart(body, params["boundary"])
		if err != nil {
			return Request{}, err
		}
		for i := range parts {
			// Boundary and Content-Disposition are transport details; names and
			// filenames already carry their meaning in the fixture format.
			delete(parts[i].Headers, "Content-Disposition")
			parts[i].Body = literalCaptureBody([]byte(*parts[i].Body.Text))
		}
		result.Body = &Body{Multipart: parts}
		delete(result.Headers, "Content-Type")
	}
	return result, nil
}

// Keep request headers available for author-selected matcher annotations. The
// recorder decides which noncredential headers belong in the persisted fixture.
func captureRequestHeaders(headers http.Header) map[string]string {
	result := map[string]string{}
	for name := range headers {
		lower := strings.ToLower(name)
		if strings.Contains(lower, "auth") || strings.Contains(lower, "cookie") ||
			strings.Contains(lower, "token") || strings.Contains(lower, "secret") ||
			strings.Contains(lower, "api-key") || strings.Contains(lower, "apikey") {
			continue
		}
		result[name] = headers.Get(name)
	}
	return result
}

// Responses retain only headers that affect decoding, caching, or retry timing.
func captureHeaders(headers http.Header) map[string]string {
	result := map[string]string{}
	for _, name := range []string{"Content-Type", "Etag", "Last-Modified", "Retry-After"} {
		if value := headers.Get(name); value != "" {
			result[name] = value
		}
	}
	return result
}

func captureBody(data []byte) *Body {
	if len(data) == 0 {
		return nil
	}
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.UseNumber()
	var value any
	if json.Valid(data) && decoder.Decode(&value) == nil {
		return &Body{JSON: captureNumbers(value)}
	}
	return literalCaptureBody(data)
}

func literalCaptureBody(data []byte) *Body {
	if utf8.Valid(data) {
		text := string(data)
		return &Body{Text: &text}
	}
	binary := base64.StdEncoding.EncodeToString(data)
	return &Body{Base64: &binary}
}

// YAML treats json.Number as a string. Native integers keep IDs exact and
// numeric through serialization, rather than rounding them through float64.
func captureNumbers(value any) any {
	switch v := value.(type) {
	case json.Number:
		if n, err := strconv.ParseInt(string(v), 10, 64); err == nil {
			return n
		}
		if n, err := strconv.ParseUint(string(v), 10, 64); err == nil {
			return n
		}
		if n, err := strconv.ParseFloat(string(v), 64); err == nil {
			return n
		}
	case map[string]any:
		for key, child := range v {
			v[key] = captureNumbers(child)
		}
	case []any:
		for i, child := range v {
			v[i] = captureNumbers(child)
		}
	}
	return value
}

func cloneCapturedSteps(steps []CapturedStep) ([]CapturedStep, error) {
	data, err := yaml.Marshal(steps)
	if err != nil {
		return nil, fmt.Errorf("encode captured HTTP calls: %w", err)
	}
	var clone []CapturedStep
	if err := yaml.Unmarshal(data, &clone); err != nil {
		return nil, fmt.Errorf("decode captured HTTP calls: %w", err)
	}
	return clone, nil
}
