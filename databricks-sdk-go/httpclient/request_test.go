package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestMakeRequestBody(t *testing.T) {
	type x struct {
		Scope string `json:"scope" url:"scope"`
	}
	requestURL := "/a/b/c"
	body, err := makeRequestBody("GET", &requestURL, x{"test"}, "", nil)
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))

	requestURL = "/a/b/c"
	body, err = makeRequestBody("POST", &requestURL, x{"test"}, "", nil)
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", requestURL)
	x1 := `{"scope":"test"}`
	require.Equal(t, []byte(x1), bodyBytes)

	requestURL = "/a/b/c"
	body, err = makeRequestBody("HEAD", &requestURL, x{"test"}, "", nil)
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))
}

func TestMakeRequestBodyFromReader(t *testing.T) {
	requestURL := "/a/b/c"
	body, err := makeRequestBody("PUT", &requestURL, strings.NewReader("abc"), "", nil)
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, []byte("abc"), bodyBytes)
}

func TestUrlEncoding(t *testing.T) {
	data := getOAuthTokenRequest{
		Assertion:            "assertion",
		AuthorizationDetails: "[{\"a\":\"b\"}]",
		GrantType:            "grant",
	}
	requestURL := "/a/b/c"
	body, err := makeRequestBody("POST", &requestURL, data, UrlEncodedContentType, nil)
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", requestURL)
	require.Equal(t, string(bodyBytes), "assertion=assertion&authorization_details=%5B%7B%22a%22%3A%22b%22%7D%5D&grant_type=grant")
}

func TestMakeRequestBodyReaderError(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("POST", &requestURL, errReader(false), "", nil)
	// The request body is only read once the request is sent, so no error
	// should be returned until then.
	require.NoError(t, err, "request body reader error should be ignored")
}

func TestMakeRequestBodyJsonError(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo chan string `json:"foo"`
	}
	_, err := makeRequestBody("POST", &requestURL, x{make(chan string)}, "", nil)
	require.EqualError(t, err, "request marshal failure: json: unsupported type: chan string")
}

type failingUrlEncode string

func (fue failingUrlEncode) EncodeValues(key string, v *url.Values) error {
	return fmt.Errorf("%s", string(fue))
}

func TestMakeRequestBodyQueryFailingEncode(t *testing.T) {
	requestURL := "/a/b/c"
	type x struct {
		Foo failingUrlEncode `url:"foo"`
	}
	_, err := makeRequestBody("GET", &requestURL, x{failingUrlEncode("always failing")}, "", nil)
	require.EqualError(t, err, "cannot create query string: always failing")
}

func TestMakeRequestBodyQueryUnsupported(t *testing.T) {
	requestURL := "/a/b/c"
	_, err := makeRequestBody("GET", &requestURL, true, "", nil)
	require.EqualError(t, err, "unsupported query string data: true")
}

func TestMakeQueryStringForceSendFields(t *testing.T) {
	type req struct {
		Cascade bool   `json:"-" url:"cascade,omitempty"`
		Force   bool   `json:"-" url:"force,omitempty"`
		Name    string `json:"-" url:"name,omitempty"`

		ForceSendFields []string `json:"-" url:"-"`
	}

	cases := []struct {
		name string
		in   req
		want string
	}{
		{
			name: "zero value dropped without ForceSendFields",
			in:   req{Cascade: false},
			want: "",
		},
		{
			name: "false bool sent when forced",
			in:   req{Cascade: false, ForceSendFields: []string{"Cascade"}},
			want: "cascade=false",
		},
		{
			name: "non-zero value still sent without forcing",
			in:   req{Cascade: true},
			want: "cascade=true",
		},
		{
			name: "multiple forced fields, keys sorted by Encode",
			in:   req{ForceSendFields: []string{"Cascade", "Force"}},
			want: "cascade=false&force=false",
		},
		{
			name: "forcing does not override an explicitly set value",
			in:   req{Cascade: true, ForceSendFields: []string{"Cascade"}},
			want: "cascade=true",
		},
		{
			name: "unknown forced field is ignored",
			in:   req{ForceSendFields: []string{"DoesNotExist"}},
			want: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := makeQueryString(tc.in)
			if err != nil {
				t.Fatalf("makeQueryString returned error: %v", err)
			}
			if got != tc.want {
				t.Errorf("makeQueryString() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestMakeQueryStringForceSendFieldsNested(t *testing.T) {
	type inner struct {
		Flag bool  `json:"-" url:"flag,omitempty"`
		Num  int64 `json:"-" url:"num,omitempty"`

		ForceSendFields []string `json:"-" url:"-"`
	}
	type deep struct {
		Inner inner `json:"-" url:"inner"`

		ForceSendFields []string `json:"-" url:"-"`
	}
	type outer struct {
		Inner    inner  `json:"-" url:"inner"`
		PtrInner *inner `json:"-" url:"ptr_inner,omitempty"`
		Deep     deep   `json:"-" url:"deep"`

		ForceSendFields []string `json:"-" url:"-"`
	}

	cases := []struct {
		name string
		in   outer
		want string
	}{
		{
			name: "nested struct honors its own ForceSendFields",
			in:   outer{Inner: inner{ForceSendFields: []string{"Flag"}}},
			want: "inner.flag=false",
		},
		{
			name: "nested zero field dropped without forcing",
			in:   outer{Inner: inner{Flag: false}},
			want: "",
		},
		{
			name: "nested non-zero value still sent",
			in:   outer{Inner: inner{Num: 7}},
			want: "inner.num=7",
		},
		{
			name: "non-nil pointer to nested struct is honored",
			in:   outer{PtrInner: &inner{ForceSendFields: []string{"Flag"}}},
			want: "ptr_inner.flag=false",
		},
		{
			name: "nil pointer to nested struct adds nothing",
			in:   outer{ForceSendFields: []string{"PtrInner"}},
			want: "",
		},
		{
			name: "doubly nested struct honors deepest ForceSendFields",
			in:   outer{Deep: deep{Inner: inner{ForceSendFields: []string{"Num"}}}},
			want: "deep.inner.num=0",
		},
		{
			name: "mixed top-level and nested forcing",
			in: outer{
				Inner:           inner{ForceSendFields: []string{"Flag", "Num"}},
				ForceSendFields: nil,
			},
			want: "inner.flag=false&inner.num=0",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := makeQueryString(tc.in)
			if err != nil {
				t.Fatalf("makeQueryString returned error: %v", err)
			}
			if got != tc.want {
				t.Errorf("makeQueryString() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestWithTokenSource(t *testing.T) {
	client := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			foo := r.Header.Get("Foo")
			require.Equal(t, "bar", foo)
			token := r.Header.Get("Authorization")
			reader := strings.NewReader(token)
			return &http.Response{
				StatusCode: 204,
				Request:    r,
				Body:       io.NopCloser(reader),
			}, nil
		}),
	})

	var buf bytes.Buffer
	ctx := context.Background()
	err := client.Do(ctx, "GET", "abc",
		WithResponseUnmarshal(&buf),
		WithRequestHeader("Foo", "bar"),
		WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
			TokenType:   "awesome",
			AccessToken: "token",
		})))
	require.NoError(t, err)
	require.Equal(t, "awesome token", buf.String())
}

func TestEncodeMultiSegmentPathParameter(t *testing.T) {
	// Slashes should not be encoded.
	assert.Equal(t, "a/b/c", EncodeMultiSegmentPathParameter("a/b/c"))
	// # and ? should be encoded.
	assert.Equal(t, "a%23b%3Fc", EncodeMultiSegmentPathParameter("a#b?c"))
}

func TestMakeRequestBodyExplicitQueryParams(t *testing.T) {
	type x struct {
		Scope string `json:"scope" url:"scope"`
	}
	requestURL := "/a/b/c"
	// For GET, it should be ignored.
	body, err := makeRequestBody("GET", &requestURL, x{"test"}, "", map[string]any{"foo": "bar"})
	require.NoError(t, err)
	bodyBytes, err := io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?scope=test", requestURL)
	require.Equal(t, 0, len(bodyBytes))

	requestURL = "/a/b/c"
	body, err = makeRequestBody("POST", &requestURL, x{"test"}, "", map[string]any{"foo": "bar"})
	require.NoError(t, err)
	bodyBytes, err = io.ReadAll(body.Reader)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c?foo=bar", requestURL)
	x1 := `{"scope":"test"}`
	require.Equal(t, []byte(x1), bodyBytes)
}
