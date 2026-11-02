package httpfixture

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func bodyText(value string) *Body { return &Body{Text: &value} }

func bodyBase64(value string) *Body { return &Body{Base64: &value} }

func TestRawBodiesAndResponses(t *testing.T) {
	for _, tc := range []struct {
		name, content, contentType string
		body                       *Body
	}{
		{"text", "# comment\n", "text/plain; charset=utf-8", bodyText("# comment\n")},
		{"binary", "\x00\xff\n", "application/octet-stream", bodyBase64("AP8K")},
		{"empty text", "", "text/plain; charset=utf-8", bodyText("")},
		{"empty binary", "", "application/octet-stream", bodyBase64("")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler()
			require.NoError(t, h.StartStep("upload", []Call{{
				Request:  Request{Method: "PUT", Path: "/file", Body: tc.body},
				Response: Response{Status: 200, Body: tc.body},
			}}))
			w := send(h, "PUT", "/file", tc.content)
			require.Equal(t, 200, w.Code)
			require.Equal(t, tc.content, w.Body.String())
			require.Equal(t, tc.contentType, w.Header().Get("Content-Type"))
			require.NoError(t, h.Verify())
		})
	}
}

func TestRawBodiesRetainExactMatchingAndFIFO(t *testing.T) {
	h := NewHandler()
	for _, body := range []string{"a", " a", "a\n", "null", "not-json"} {
		c := Call{Request: Request{Method: "PUT", Path: "/file", Body: bodyText(body)}, Response: Response{Status: 200}}
		require.NoError(t, h.StartStep("exact", []Call{c}))
		require.Equal(t, 400, send(h, "PUT", "/file", body+" ").Code)
		require.ErrorContains(t, h.Verify(), "unconsumed call 1")
		// Start a fresh handler because unexpected calls remain test failures.
		h = NewHandler()
		c.Response.Body = bodyText("first")
		second := c
		second.Response.Body = bodyText("second")
		require.NoError(t, h.StartStep("fifo", []Call{c, second}))
		require.Equal(t, "first", send(h, "PUT", "/file", body).Body.String())
		require.Equal(t, "second", send(h, "PUT", "/file", body).Body.String())
		require.NoError(t, h.Verify())
	}
}

func multipartRequest(t *testing.T, handler http.Handler, boundary string, parts []MultipartPart) *httptest.ResponseRecorder {
	t.Helper()
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	require.NoError(t, writer.SetBoundary(boundary))
	for _, part := range parts {
		headers := make(textproto.MIMEHeader)
		disposition := `form-data; name="` + part.Name + `"`
		if part.Filename != "" {
			disposition += `; filename="` + part.Filename + `"`
		}
		headers.Set("Content-Disposition", disposition)
		for key, value := range part.Headers {
			headers.Set(key, value)
		}
		w, err := writer.CreatePart(headers)
		require.NoError(t, err)
		body, err := part.Body.literalBytes()
		require.NoError(t, err)
		_, err = w.Write(body)
		require.NoError(t, err)
	}
	require.NoError(t, writer.Close())
	req := httptest.NewRequest("POST", "/upload", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}

func TestMultipartIgnoresBoundaryAndPartOrder(t *testing.T) {
	parts := []MultipartPart{
		{Name: "path", Body: bodyText("/Shared/script.py")},
		{Name: "content", Filename: "content", Body: bodyBase64("AP8K"), Headers: map[string]string{"content-type": "application/octet-stream"}},
		{Name: "empty", Filename: "empty", Body: bodyText("")},
	}
	h := NewHandler()
	call := Call{Request: Request{Method: "POST", Path: "/upload", Body: &Body{Multipart: parts}}, Response: Response{Status: 204}}
	require.NoError(t, h.StartStep("upload", []Call{call, call}))
	require.Equal(t, 204, multipartRequest(t, h, "first-boundary", parts).Code)
	reversed := []MultipartPart{parts[2], parts[1], parts[0]}
	require.Equal(t, 204, multipartRequest(t, h, "another-boundary", reversed).Code)
	require.NoError(t, h.Verify())
}

func TestMultipartRejectsChangedPartsWithoutConsumption(t *testing.T) {
	for _, tc := range []struct {
		name  string
		parts []MultipartPart
	}{
		{"value", []MultipartPart{{Name: "path", Body: bodyText("wrong")}, {Name: "content", Filename: "content", Body: bodyText("# comment\n")}}},
		{"file bytes", []MultipartPart{{Name: "path", Body: bodyText("/Shared/script.py")}, {Name: "content", Filename: "content", Body: bodyText("# changed\n")}}},
		{"filename", []MultipartPart{{Name: "path", Body: bodyText("/Shared/script.py")}, {Name: "content", Filename: "other", Body: bodyText("# comment\n")}}},
		{"missing", []MultipartPart{{Name: "path", Body: bodyText("/Shared/script.py")}}},
		{"duplicate", []MultipartPart{{Name: "path", Body: bodyText("/Shared/script.py")}, {Name: "path", Body: bodyText("/Shared/script.py")}}},
		{"extra", []MultipartPart{{Name: "path", Body: bodyText("/Shared/script.py")}, {Name: "content", Filename: "content", Body: bodyText("# comment\n")}, {Name: "extra", Body: bodyText("")}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			parts := []MultipartPart{{Name: "path", Body: bodyText("/Shared/script.py")}, {Name: "content", Filename: "content", Body: bodyText("# comment\n")}}
			h := NewHandler()
			require.NoError(t, h.StartStep("upload", []Call{{Request: Request{Method: "POST", Path: "/upload", Body: &Body{Multipart: parts}}, Response: Response{Status: 204}}}))
			w := multipartRequest(t, h, "boundary", tc.parts)
			require.Equal(t, 400, w.Code)
			require.ErrorContains(t, h.Verify(), "unconsumed call 1")
			require.Equal(t, 204, multipartRequest(t, h, "new-boundary", parts).Code)
			require.ErrorContains(t, h.Verify(), "unexpected HTTP call")
			require.NotContains(t, h.Verify().Error(), "unconsumed call")
		})
	}
}

func TestMultipartListedHeadersRemainExact(t *testing.T) {
	h := NewHandler()
	parts := []MultipartPart{{Name: "content", Body: bodyText("# comment"), Headers: map[string]string{"Content-Type": "text/plain"}}}
	require.NoError(t, h.StartStep("upload", []Call{{Request: Request{Method: "POST", Path: "/upload", Body: &Body{Multipart: parts}}, Response: Response{Status: 204}}}))
	wrong := []MultipartPart{{Name: "content", Body: bodyText("# comment"), Headers: map[string]string{"Content-Type": "application/octet-stream"}}}
	require.Equal(t, 400, multipartRequest(t, h, "boundary", wrong).Code)
	require.ErrorContains(t, h.Verify(), "unconsumed call 1")
}

func TestMalformedMultipartReportsStep(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.StartStep("upload", nil))
	req := httptest.NewRequest("POST", "/upload", strings.NewReader("broken"))
	req.Header.Set("Content-Type", "multipart/form-data; boundary=missing")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	require.Equal(t, 400, w.Code)
	require.ErrorContains(t, h.Verify(), "upload: POST /upload: invalid request body: invalid multipart body")
}

func TestBodyYAMLValidation(t *testing.T) {
	for _, tc := range []struct{ body, message string }{
		{"{}", "exactly one"},
		{"{json: {}, text: x}", "exactly one"},
		{"{typo: x}", "field typo not found"},
		{"{text: 42}", "body.text must be a string"},
		{"{base64: null}", "body.base64 must be a string"},
		{"{base64: invalid!}", "illegal base64"},
		{"{multipart: []}", "at least one part"},
		{"{multipart: [{body: {text: x}}]}", "requires name and body"},
		{"{multipart: [{name: file}]}", "requires name and body"},
		{"{multipart: [{name: file, body: {json: {}}}]}", "requires body.text or body.base64"},
		{"{multipart: [{name: file, body: {text: x}, typo: true}]}", "field typo not found"},
	} {
		t.Run(tc.body, func(t *testing.T) {
			var body Body
			err := yaml.Unmarshal([]byte(tc.body), &body)
			if err == nil {
				_, err = compileCalls([]Call{{Request: Request{Method: "POST", Path: "/upload", Body: &body}, Response: Response{Status: 200}}})
			}
			require.ErrorContains(t, err, tc.message)
		})
	}
	for _, input := range []string{"{json: null}", "{text: ''}", "{base64: ''}"} {
		var body Body
		require.NoError(t, yaml.Unmarshal([]byte(input), &body))
		_, err := body.literalBytes()
		require.NoError(t, err)
	}
}

func TestRawBodyRejectsJSONMatchersAndMultipartResponse(t *testing.T) {
	call := Call{Request: Request{Method: "POST", Path: "/file", Body: bodyText("x"), Matchers: Matchers{JSON: []JSONMatcher{{Path: []any{}, Type: AnyString}}}}, Response: Response{Status: 200}}
	_, err := compileCalls([]Call{call})
	require.Error(t, err)
	call.Request.Matchers = Matchers{}
	call.Response.Body = &Body{Multipart: []MultipartPart{{Name: "file", Body: bodyText("x")}}}
	_, err = compileCalls([]Call{call})
	require.ErrorContains(t, err, "multipart is supported only for request bodies")
}

func TestBodyYAMLRoundTripPreservesRepresentation(t *testing.T) {
	for _, body := range []*Body{{JSON: nil}, {JSON: map[string]any{"id": "file"}}, bodyText(""), bodyBase64(""), {Multipart: []MultipartPart{{Name: "file", Body: bodyText("# comment")}}}} {
		encoded, err := yaml.Marshal(body)
		require.NoError(t, err)
		var decoded Body
		require.NoError(t, yaml.Unmarshal(encoded, &decoded))
		require.Equal(t, body, &decoded)
	}
}

func TestMultipartRequiresEveryDuplicatePart(t *testing.T) {
	h := NewHandler()
	part := MultipartPart{Name: "tag", Body: bodyText("same")}
	parts := []MultipartPart{part, part}
	require.NoError(t, h.StartStep("upload", []Call{{Request: Request{Method: "POST", Path: "/upload", Body: &Body{Multipart: parts}}, Response: Response{Status: 204}}}))
	require.Equal(t, 400, multipartRequest(t, h, "boundary", parts[:1]).Code)
	require.ErrorContains(t, h.Verify(), "unconsumed call 1")
	require.Equal(t, 204, multipartRequest(t, h, "boundary", parts).Code)
}

func TestMultipartOverlappingPartHeadersDoNotImposeOrder(t *testing.T) {
	parts := []MultipartPart{
		{Name: "tag", Body: bodyText("same")},
		{Name: "tag", Body: bodyText("same"), Headers: map[string]string{"Content-Type": "text/plain"}},
	}
	h := NewHandler()
	require.NoError(t, h.StartStep("upload", []Call{{Request: Request{Method: "POST", Path: "/upload", Body: &Body{Multipart: parts}}, Response: Response{Status: 204}}}))
	require.Equal(t, 204, multipartRequest(t, h, "boundary", []MultipartPart{parts[1], parts[0]}).Code)
	require.NoError(t, h.Verify())
}

func TestBodyYAMLPreservesExternalAnchors(t *testing.T) {
	var document struct {
		Sample any  `yaml:"sample"`
		Body   Body `yaml:"body"`
	}
	require.NoError(t, yaml.Unmarshal([]byte("sample: &sample {name: script}\nbody: {json: *sample}\n"), &document))
	require.Equal(t, map[string]any{"name": "script"}, document.Body.JSON)
	require.NoError(t, yaml.Unmarshal([]byte("sample: &sample '# comment'\nbody: {multipart: [{name: content, body: {text: *sample}}]}\n"), &document))
	require.Equal(t, "# comment", *document.Body.Multipart[0].Body.Text)
}
