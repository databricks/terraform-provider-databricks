package httpfixture

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"

	"gopkg.in/yaml.v3"
)

// Body selects semantic JSON, exact UTF-8 text, exact binary bytes encoded as
// base64, or multipart form parts. A nil *Body means no HTTP body.
type Body struct {
	JSON      any             `yaml:"json,omitempty"`
	Text      *string         `yaml:"text,omitempty"`
	Base64    *string         `yaml:"base64,omitempty"`
	Multipart []MultipartPart `yaml:"multipart,omitempty"`
}

// MultipartPart describes one form field or file. Part order and the transport
// boundary are ignored; names, filenames, bytes and listed headers remain exact.
type MultipartPart struct {
	Name     string            `yaml:"name"`
	Filename string            `yaml:"filename,omitempty"`
	Headers  map[string]string `yaml:"headers,omitempty"`
	Body     *Body             `yaml:"body"`
}

// UnmarshalYAML requires an explicit representation, including json: null.
// Decode nodes directly so references to other YAML anchors remain valid.
func (b *Body) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode || len(node.Content) != 2 {
		return fmt.Errorf("body requires exactly one of json, text, base64, multipart")
	}
	*b = Body{}
	key, value := node.Content[0].Value, node.Content[1]
	if value.Kind == yaml.AliasNode {
		value = value.Alias
	}
	switch key {
	case "json":
		return value.Decode(&b.JSON)
	case "text", "base64":
		if value.Tag != "!!str" {
			return fmt.Errorf("body.%s must be a string", key)
		}
		if key == "text" {
			b.Text = &value.Value
		} else {
			b.Base64 = &value.Value
		}
	case "multipart":
		if err := value.Decode(&b.Multipart); err != nil {
			return err
		}
		if len(b.Multipart) == 0 {
			return fmt.Errorf("body.multipart must contain at least one part")
		}
	default:
		return fmt.Errorf("field %s not found in httpfixture.Body", key)
	}
	return nil
}

// UnmarshalYAML keeps strict field checking inside the custom Body decoder.
func (p *MultipartPart) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("multipart part must be a mapping")
	}
	for i := 0; i < len(node.Content); i += 2 {
		switch key := node.Content[i].Value; key {
		case "name", "filename", "headers", "body":
		default:
			return fmt.Errorf("field %s not found in httpfixture.MultipartPart", key)
		}
	}
	type plain MultipartPart
	return node.Decode((*plain)(p))
}

// MarshalYAML preserves an explicit JSON null instead of emitting an empty body
// mapping. The other representations also retain intentionally empty bytes.
func (b Body) MarshalYAML() (any, error) {
	kind, err := b.kind()
	if err != nil {
		return nil, err
	}
	switch kind {
	case "text":
		return map[string]any{kind: *b.Text}, nil
	case "base64":
		return map[string]any{kind: *b.Base64}, nil
	case "multipart":
		return map[string]any{kind: b.Multipart}, nil
	default:
		return map[string]any{kind: b.JSON}, nil
	}
}

func (b *Body) kind() (string, error) {
	kind, count := "json", 0
	if b.JSON != nil {
		count++
	}
	if b.Text != nil {
		kind, count = "text", count+1
	}
	if b.Base64 != nil {
		kind, count = "base64", count+1
	}
	if b.Multipart != nil {
		kind, count = "multipart", count+1
	}
	if count > 1 {
		return "", fmt.Errorf("body requires exactly one of json, text, base64, multipart")
	}
	return kind, nil
}

// literalBytes also validates binary samples before a test starts.
func (b *Body) literalBytes() ([]byte, error) {
	kind, err := b.kind()
	if err != nil {
		return nil, err
	}
	switch kind {
	case "json":
		return json.Marshal(b.JSON)
	case "text":
		return []byte(*b.Text), nil
	case "base64":
		return base64.StdEncoding.DecodeString(*b.Base64)
	default:
		return nil, fmt.Errorf("multipart is supported only for request bodies")
	}
}

func compileMultipart(parts []MultipartPart) error {
	if len(parts) == 0 {
		return fmt.Errorf("body.multipart must contain at least one part")
	}
	for i, part := range parts {
		if part.Name == "" || part.Body == nil {
			return fmt.Errorf("multipart part %d requires name and body", i+1)
		}
		kind, err := part.Body.kind()
		if err != nil {
			return err
		}
		if kind != "text" && kind != "base64" {
			return fmt.Errorf("multipart part %d requires body.text or body.base64", i+1)
		}
		if _, err := part.Body.literalBytes(); err != nil {
			return fmt.Errorf("multipart part %d: %w", i+1, err)
		}
	}
	return nil
}

type requestBody struct {
	raw       []byte
	json      any
	jsonValid bool
	parts     []MultipartPart
}

// parseRequestBody never spools uploads to disk. A non-JSON body may still match
// text/base64 fixtures; it cannot accidentally match a JSON-null fixture.
func parseRequestBody(req *http.Request) (requestBody, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return requestBody{}, err
	}
	value, jsonErr := decodeJSON(body)
	result := requestBody{raw: body, json: value, jsonValid: jsonErr == nil}
	contentType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err == nil && contentType == "multipart/form-data" {
		result.parts, err = readMultipart(body, params["boundary"])
		if err != nil {
			return requestBody{}, fmt.Errorf("invalid multipart body: %w", err)
		}
	}
	return result, nil
}

func readMultipart(body []byte, boundary string) ([]MultipartPart, error) {
	reader := multipart.NewReader(bytes.NewReader(body), boundary)
	var parts []MultipartPart
	for {
		part, err := reader.NextRawPart()
		if err == io.EOF {
			return parts, nil
		}
		if err != nil {
			return nil, err
		}
		content, err := io.ReadAll(part)
		if err != nil {
			return nil, err
		}
		disposition, params, err := mime.ParseMediaType(part.Header.Get("Content-Disposition"))
		if err != nil || disposition != "form-data" || params["name"] == "" {
			return nil, fmt.Errorf("every multipart part requires a form-data name")
		}
		headers := make(map[string]string, len(part.Header))
		for key := range part.Header {
			headers[key] = part.Header.Get(key)
		}
		text := string(content)
		parts = append(parts, MultipartPart{Name: params["name"], Filename: params["filename"], Headers: headers, Body: &Body{Text: &text}})
	}
}

func matchesMultipart(expected, actual []MultipartPart) bool {
	if len(expected) != len(actual) {
		return false
	}
	assigned := make([]int, len(actual))
	for i := range assigned {
		assigned[i] = -1
	}
	// Reassign duplicate parts when optional headers make signatures overlap.
	// A greedy match could otherwise depend on wire order.
	var assign func(int, []bool) bool
	assign = func(want int, seen []bool) bool {
		for got := range actual {
			if seen[got] || !matchesPart(expected[want], actual[got]) {
				continue
			}
			seen[got] = true
			if assigned[got] == -1 || assign(assigned[got], seen) {
				assigned[got] = want
				return true
			}
		}
		return false
	}
	for i := range expected {
		if !assign(i, make([]bool, len(actual))) {
			return false
		}
	}
	return true
}

func matchesPart(expected, actual MultipartPart) bool {
	if expected.Name != actual.Name || expected.Filename != actual.Filename {
		return false
	}
	want, _ := expected.Body.literalBytes()
	got, _ := actual.Body.literalBytes()
	return bytes.Equal(want, got) && matchesPartHeaders(expected.Headers, actual.Headers)
}

func matchesPartHeaders(expected, actual map[string]string) bool {
	for key, value := range expected {
		if actual[http.CanonicalHeaderKey(key)] != value {
			return false
		}
	}
	return true
}
