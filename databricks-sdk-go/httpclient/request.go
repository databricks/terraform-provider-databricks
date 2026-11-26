package httpclient

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const UrlEncodedContentType = "application/x-www-form-urlencoded"

// WithRequestHeader adds a request visitor, that sets a header on a request
func WithRequestHeader(k, v string) DoOption {
	return WithRequestVisitor(func(r *http.Request) error {
		r.Header.Set(k, v)
		return nil
	})
}

// WithRequestHeaders adds a request visitor, that set all headers from a map
func WithRequestHeaders(headers map[string]string) DoOption {
	return WithRequestVisitor(func(r *http.Request) error {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
		return nil
	})
}

// WithToken uses the specified golang.org/x/oauth2 token on a request
func WithToken(token *oauth2.Token) DoOption {
	visitor := WithTokenSource(oauth2.StaticTokenSource(token))
	visitor.isAuthOption = true
	return visitor
}

// WithAuthTokenSource uses the specified auth.TokenSource on a request. The
// request's context is passed to the token source when fetching the token.
func WithAuthTokenSource(ts auth.TokenSource) DoOption {
	return WithRequestVisitor(func(r *http.Request) error {
		token, err := ts.Token(r.Context())
		if err != nil {
			return fmt.Errorf("token: %w", err)
		}
		authHeader := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
		r.Header.Set("Authorization", authHeader)
		return nil
	})
}

// WithTokenSource uses the specified golang.org/x/oauth2 token source on a
// request.
func WithTokenSource(ts oauth2.TokenSource) DoOption {
	return WithAuthTokenSource(authconv.AuthTokenSource(ts))
}

// WithRequestVisitor applies given function on a request
func WithRequestVisitor(visitor func(r *http.Request) error) DoOption {
	return DoOption{
		in: visitor,
	}
}

// WithRequestData takes either a struct instance, map, string, bytes, or io.Reader
// and sends it either as query string for GET and DELETE calls, or as request body
// for POST, PUT, and PATCH calls.
//
// Experimental: this method may eventually be split into more granular options.
func WithRequestData(body any) DoOption {
	// refactor this, so that we split JSON/query string serialization and make
	// separate request visitors internally.
	return DoOption{
		body: body,
	}
}

// WithQueryParameters takes a map and sends it as query string for non GET/DELETE/HEAD calls.
// This is ignored for GET/DELETE/HEAD calls, as the query parameters are serialized from the body instead.
//
// Experimental: this method may eventually be split into more granular options.
func WithQueryParameters(queryParams map[string]any) DoOption {
	// refactor this, so that we split JSON/query string serialization and make
	// separate request visitors internally.
	return DoOption{
		queryParams: queryParams,
	}
}

// WithUrlEncodedData takes either a struct instance, map, string, bytes, or io.Reader plus
// a content type, and sends it either as query string for GET and DELETE calls, or as request body
// for POST, PUT, and PATCH calls. The content type is set to "application/x-www-form-urlencoded"
// and the body is encoded as a query string.
//
// Experimental: this method may eventually be split into more granular options.
func WithUrlEncodedData(body any) DoOption {
	return DoOption{
		in: func(r *http.Request) error {
			r.Header.Set("Content-Type", UrlEncodedContentType)
			return nil
		},
		body:        body,
		contentType: UrlEncodedContentType,
	}
}

func makeQueryString(data interface{}) (string, error) {
	inputVal := reflect.ValueOf(data)
	inputType := reflect.TypeOf(data)
	if inputType.Kind() == reflect.Map {
		s := []string{}
		keys := inputVal.MapKeys()
		// sort map keys by their string repr, so that tests can be deterministic
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})
		for _, k := range keys {
			v := inputVal.MapIndex(k)
			if v.IsZero() {
				continue
			}
			s = append(s, fmt.Sprintf("%s=%s",
				strings.Replace(url.QueryEscape(fmt.Sprintf("%v", k.Interface())), "+", "%20", -1),
				strings.Replace(url.QueryEscape(fmt.Sprintf("%v", v.Interface())), "+", "%20", -1)))
		}
		return strings.Join(s, "&"), nil
	}
	if inputType.Kind() == reflect.Struct {
		params, err := query.Values(data)
		if err != nil {
			return "", fmt.Errorf("cannot create query string: %w", err)
		}
		// go-querystring honors the `url` tag's omitempty option but knows
		// nothing about ForceSendFields, so a zero-valued field (e.g. a false
		// bool) listed in ForceSendFields would be dropped. Re-add those
		// fields so query parameters follow the same ForceSendFields semantics
		// already used for JSON bodies.
		addForceSendQueryParams(inputVal, params, "")
		// Query parameters may be nested, but the keys generated by
		// query.Values use the "[" and "]" characters to represent nesting.
		// Replace all instances of "[" with "." and "]" with empty string
		// to make the query string compatible with the proto API.
		// See the following for more information:
		// https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule
		protoCompatibleParams := make(url.Values)
		for k, vs := range params {
			newK := strings.Replace(k, "[", ".", -1)
			newK = strings.Replace(newK, "]", "", -1)
			for _, v := range vs {
				protoCompatibleParams.Add(newK, v)
			}
		}
		return protoCompatibleParams.Encode(), nil
	}
	return "", fmt.Errorf("unsupported query string data: %#v", data)
}

// addForceSendQueryParams re-adds query parameters that go-querystring dropped
// because of the `url` tag's omitempty option but which the struct's
// ForceSendFields asks to send anyway. This mirrors the ForceSendFields
// handling for JSON bodies in the marshal package, including for nested
// structs: each struct level honors its own ForceSendFields.
//
// prefix is the go-querystring key prefix for the fields of structVal ("" at
// the top level). Nested keys use the same "parent[child]" bracket notation
// that go-querystring emits, so the caller's proto-compat rewrite ("[" -> ".",
// "]" -> "") applies uniformly to forced and non-forced parameters alike.
//
// Only basic (scalar) types are forced, matching marshal. Nil pointers and
// slices are never forced (a nil pointer means "not set"; marshal also skips
// nil/empty slices). Slice elements are not recursed into.
func addForceSendQueryParams(structVal reflect.Value, params url.Values, prefix string) {
	structVal = reflect.Indirect(structVal)
	if structVal.Kind() != reflect.Struct {
		return
	}
	structType := structVal.Type()

	// Collect this struct level's ForceSendFields (Go field names).
	forced := map[string]bool{}
	if f := structVal.FieldByName("ForceSendFields"); f.IsValid() {
		if names, ok := f.Interface().([]string); ok {
			for _, name := range names {
				forced[name] = true
			}
		}
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if !field.IsExported() {
			continue
		}
		// Reuse the first segment of the `url` tag as the parameter name;
		// skip fields that are not query parameters (no tag or "-").
		paramName, _, _ := strings.Cut(field.Tag.Get("url"), ",")
		if paramName == "" || paramName == "-" {
			continue
		}
		key := paramName
		if prefix != "" {
			key = prefix + "[" + paramName + "]"
		}
		value := structVal.Field(i)

		// Recurse into nested structs (and non-nil pointers to structs) so a
		// nested struct's own ForceSendFields is honored too. reflect.Indirect
		// of a nil pointer yields an invalid Value, so nil pointers stop here.
		if reflect.Indirect(value).Kind() == reflect.Struct {
			addForceSendQueryParams(value, params, key)
			continue
		}

		// Force-send a dropped zero-valued scalar, leaving any value
		// go-querystring already encoded untouched.
		if !forced[field.Name] || !isForceableQueryKind(value.Kind()) {
			continue
		}
		if _, exists := params[key]; exists {
			continue
		}
		// go-querystring formats these kinds with fmt.Sprint, so this keeps
		// the encoding identical to a non-zero value.
		params.Set(key, fmt.Sprint(value.Interface()))
	}
}

// isForceableQueryKind reports whether a value of the given kind may be
// force-sent as a query parameter. Mirrors the basic-type rule used by the
// marshal package for JSON bodies.
func isForceableQueryKind(k reflect.Kind) bool {
	switch k {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}

func EncodeMultiSegmentPathParameter(p string) string {
	segments := strings.Split(p, "/")
	b := strings.Builder{}
	b.Grow(len(p))
	for i, s := range segments {
		if i > 0 {
			b.WriteString("/")
		}
		b.WriteString(url.PathEscape(s))
	}
	return b.String()
}

// We used to not send any query parameters for non GET/DELETE/HEAD requests.
// Moreover, serialization for query paramters in GET/DELETE/HEAD requests depends on the `url` tag.
// This tag is wrongly generated and fixing it will have an unknown inpact on the SDK.
// So:
// * GET/DELETE/HEAD requests are sent with query parameters serialized from data using the `url` tag as before (no change).
// * The rest of the requests are sent with query parameters serialized from explicitQueryParams, which does not use the `url` tag.
// TODO: For SDK-Mod, refactor this and remove the `url` tag completely.
func makeRequestBody(method string, requestURL *string, data interface{}, contentType string, explicitQueryParams map[string]any) (common.RequestBody, error) {
	if data == nil && len(explicitQueryParams) == 0 {
		return common.RequestBody{}, nil
	}
	if data != nil {
		if method == "GET" || method == "DELETE" || method == "HEAD" {
			qs, err := makeQueryString(data)
			if err != nil {
				return common.RequestBody{}, err
			}
			*requestURL += "?" + qs
			return common.NewRequestBody([]byte{})
		}
		if contentType == UrlEncodedContentType {
			qs, err := makeQueryString(data)
			if err != nil {
				return common.RequestBody{}, err
			}
			return common.NewRequestBody(qs)
		}
	}
	if len(explicitQueryParams) > 0 {
		qs, err := makeQueryString(explicitQueryParams)
		if err != nil {
			return common.RequestBody{}, err
		}
		*requestURL += "?" + qs
		return common.NewRequestBody(data)
	}
	return common.NewRequestBody(data)
}
