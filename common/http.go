package common

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	e2example                   = "https://registry.terraform.io/providers/databricks/databricks/latest/docs/guides/aws-workspace"
	accountsHost                = "accounts.cloud.databricks.com"
	transientErrorStringMatches = []string{
		"com.databricks.backend.manager.util.UnknownWorkerEnvironmentException",
		"does not have any associated worker environments",
		"There is no worker environment with id",
		"Unknown worker environment",
		"ClusterNotReadyException",
		"connection reset by peer",
		"TLS handshake timeout",
		"connection refused",
		"Unexpected error",
		"i/o timeout",
	}
)

// APIErrorBody maps "proper" databricks rest api errors to a struct
type APIErrorBody struct {
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	// The following two are for scim api only
	// for RFC 7644 Section 3.7.3 https://tools.ietf.org/html/rfc7644#section-3.7.3
	ScimDetail string `json:"detail,omitempty"`
	ScimStatus string `json:"status,omitempty"`
	ScimType   string `json:"scimType,omitempty"`
	API12Error string `json:"error,omitempty"`
}

// APIError is a generic struct for an api error on databricks
type APIError struct {
	ErrorCode  string
	Message    string
	Resource   string
	StatusCode int
}

// Error returns error message string instead of
func (apiError APIError) Error() string {
	if apiError.StatusCode != 404 {
		docs := apiError.DocumentationURL()
		log.Printf("[WARN] %s:%d - %s %s", apiError.Resource, apiError.StatusCode, apiError.Message, docs)
	}
	return apiError.Message
}

// IsMissing tells if error is about missing resource
func IsMissing(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(APIError)
	return ok && e.IsMissing()
}

// IsMissing tells if it is missing resource
func (apiError APIError) IsMissing() bool {
	return apiError.StatusCode == http.StatusNotFound
}

// IsTooManyRequests shows rate exceeded limits
func (apiError APIError) IsTooManyRequests() bool {
	return apiError.StatusCode == http.StatusTooManyRequests
}

// DocumentationURL guesses doc link
func (apiError APIError) DocumentationURL() string {
	endpointRE := regexp.MustCompile(`/api/2.0/([^/]+)/([^/]+)$`)
	endpointMatches := endpointRE.FindStringSubmatch(apiError.Resource)
	if len(endpointMatches) < 3 {
		return ""
	}
	return fmt.Sprintf("https://docs.databricks.com/dev-tools/api/latest/%s.html#%s",
		endpointMatches[1], endpointMatches[2])
}

// IsRetriable returns true if error is retriable
func (apiError APIError) IsRetriable() bool {
	// Handle transient errors for retries
	for _, substring := range transientErrorStringMatches {
		if strings.Contains(apiError.Message, substring) {
			log.Printf("[INFO] Attempting retry because of %#v", substring)
			return true
		}
	}
	// some API's recommend retries on HTTP 500, but we'll add that later
	return false
}

// NotFound returns properly formatted Not Found error
func NotFound(message string) APIError {
	return APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    message,
	}
}

func (c *DatabricksClient) parseUnknownError(
	status string, body []byte, err error) (errorBody APIErrorBody) {
	// this is most likely HTML... since un-marshalling JSON failed
	// Status parts first in case html message is not as expected
	statusParts := strings.SplitN(status, " ", 2)
	if len(statusParts) < 2 {
		errorBody.ErrorCode = "UNKNOWN"
	} else {
		errorBody.ErrorCode = strings.ReplaceAll(
			strings.ToUpper(strings.Trim(statusParts[1], " .")),
			" ", "_")
	}
	stringBody := string(body)
	messageRE := regexp.MustCompile(`<pre>(.*)</pre>`)
	messageMatches := messageRE.FindStringSubmatch(stringBody)
	// No messages with <pre> </pre> format found so return a APIError
	if len(messageMatches) < 2 {
		errorBody.Message = fmt.Sprintf("Response from server (%s) %s: %v",
			status, stringBody, err)
		return
	}
	errorBody.Message = strings.Trim(messageMatches[1], " .")
	return
}

func (c *DatabricksClient) isAccountsClient() bool {
	return c.Config.IsAccountClient()
}

func (c *DatabricksClient) commonErrorClarity(resp *http.Response) *APIError {
	isAccountsAPI := strings.HasPrefix(resp.Request.URL.Path, "/api/2.0/accounts") || strings.HasPrefix(resp.Request.URL.Path, "/api/2.0/preview/accounts")
	isAccountsClient := c.isAccountsClient()
	isTesting := strings.HasPrefix(resp.Request.URL.Host, "127.0.0.1")
	if !isTesting && isAccountsClient && !isAccountsAPI {
		return &APIError{
			ErrorCode: "INCORRECT_CONFIGURATION",
			Message: fmt.Sprintf("Databricks API (%s) requires you to set `host` property "+
				"(or DATABRICKS_HOST env variable) to result of `databricks_mws_workspaces.this.workspace_url`. "+
				"This error may happen if you're using provider in both normal and multiworkspace mode. Please "+
				"refactor your code into different modules. Runnable example that we use for integration testing "+
				"can be found in this repository at %s", resp.Request.URL.Path, e2example),
			StatusCode: resp.StatusCode,
			Resource:   resp.Request.URL.Path,
		}
	}
	// common confusion with this provider: calling workspace apis on accounts host
	if !isTesting && isAccountsAPI && !isAccountsClient {
		return &APIError{
			ErrorCode: "INCORRECT_CONFIGURATION",
			Message: fmt.Sprintf("Accounts API (%s) requires you to set %s as DATABRICKS_HOST, but you have "+
				"specified %s instead. This error may happen if you're using provider in both "+
				"normal and multiworkspace mode. Please refactor your code into different modules. "+
				"Runnable example that we use for integration testing can be found in this "+
				"repository at %s", resp.Request.URL.Path, accountsHost, c.Host, e2example),
			StatusCode: resp.StatusCode,
			Resource:   resp.Request.URL.Path,
		}
	}
	return nil
}

func (c *DatabricksClient) parseError(resp *http.Response) APIError {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIError{
			Message:    err.Error(),
			ErrorCode:  "IO_READ",
			StatusCode: resp.StatusCode,
			Resource:   resp.Request.URL.Path,
		}
	}
	log.Printf("[DEBUG] %s %v", resp.Status, c.redactedDump(body))
	mwsError := c.commonErrorClarity(resp)
	if mwsError != nil {
		return *mwsError
	}
	// try to read in nicely formatted API error response
	var errorBody APIErrorBody
	err = json.Unmarshal(body, &errorBody)
	if err != nil {
		errorBody = c.parseUnknownError(resp.Status, body, err)
	}
	if errorBody.API12Error != "" {
		// API 1.2 has different response format, let's adapt
		errorBody.Message = errorBody.API12Error
	}
	// Handle SCIM error message details
	if errorBody.Message == "" && errorBody.ScimDetail != "" {
		if errorBody.ScimDetail == "null" {
			errorBody.Message = "SCIM API Internal Error"
		} else {
			errorBody.Message = errorBody.ScimDetail
		}
		// add more context from SCIM responses
		errorBody.Message = fmt.Sprintf("%s %s", errorBody.ScimType, errorBody.Message)
		errorBody.Message = strings.Trim(errorBody.Message, " ")
		errorBody.ErrorCode = fmt.Sprintf("SCIM_%s", errorBody.ScimStatus)
	}
	if resp.StatusCode == 403 {
		errorBody.Message = fmt.Sprintf("%s. Using %s auth: %s",
			strings.Trim(errorBody.Message, "."), c.AuthType,
			c.configDebugString())
	}
	return APIError{
		Message:    errorBody.Message,
		ErrorCode:  errorBody.ErrorCode,
		StatusCode: resp.StatusCode,
		Resource:   resp.Request.URL.Path,
	}
}

// checkHTTPRetry inspects HTTP errors from the Databricks API for known transient errors on Workspace creation
func (c *DatabricksClient) checkHTTPRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ue, ok := err.(*url.Error); ok {
		apiError := APIError{
			ErrorCode:  "IO_ERROR",
			StatusCode: 523,
			Message:    ue.Error(),
		}
		return apiError.IsRetriable(), apiError
	}
	if resp == nil {
		// If response is nil we can't make retry choices.
		// In this case don't retry and return the original error from httpclient
		return false, err
	}
	if resp.StatusCode == 429 {
		return true, APIError{
			ErrorCode:  "TOO_MANY_REQUESTS",
			Message:    "Current request has to be retried",
			StatusCode: 429,
		}
	}
	if resp.StatusCode >= 400 {
		apiError := c.parseError(resp)
		return apiError.IsRetriable(), apiError
	}
	return false, nil
}

// Get on path
func (c *DatabricksClient) Get(ctx context.Context, path string, request any, response any) error {
	return c.Client.Do(ctx, http.MethodGet, path, request, response, c.addApiPrefix)
}

// Post on path
func (c *DatabricksClient) Post(ctx context.Context, path string, request any, response any) error {
	return c.Client.Do(ctx, http.MethodPost, path, request, response, c.addApiPrefix)
}

// Delete on path
func (c *DatabricksClient) Delete(ctx context.Context, path string, request any) error {
	return c.Client.Do(ctx, http.MethodDelete, path, request, nil, c.addApiPrefix)
}

// Patch on path
func (c *DatabricksClient) Patch(ctx context.Context, path string, request any) error {
	return c.Client.Do(ctx, http.MethodPatch, path, request, nil, c.addApiPrefix)
}

// Put on path
func (c *DatabricksClient) Put(ctx context.Context, path string, request any) error {
	return c.Client.Do(ctx, http.MethodPut, path, request, nil, c.addApiPrefix)
}

type ApiVersion string

const (
	API_1_2 ApiVersion = "1.2"
	API_2_0 ApiVersion = "2.0"
	API_2_1 ApiVersion = "2.1"
)

func (c *DatabricksClient) addApiPrefix(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("no URL found in request")
	}
	ctx := r.Context()
	av, ok := ctx.Value(Api).(ApiVersion)
	if !ok {
		av = API_2_0
	}
	r.URL.Path = fmt.Sprintf("/api/%s%s", av, r.URL.Path)
	return nil
}

// scimPathVisitorFactory is a separate method for the sake of unit tests
func (c *DatabricksClient) scimVisitor(r *http.Request) error {
	r.Header.Set("Content-Type", "application/scim+json; charset=utf-8")
	if c.Config.IsAccountClient() && c.Config.AccountID != "" {
		// until `/preview` is there for workspace scim,
		// `/api/2.0` is added by completeUrl visitor
		r.URL.Path = strings.ReplaceAll(r.URL.Path, "/api/2.0/preview",
			fmt.Sprintf("/api/2.0/accounts/%s", c.Config.AccountID))
	}
	return nil
}

// Scim sets SCIM headers
func (c *DatabricksClient) Scim(ctx context.Context, method, path string, request any, response any) error {
	return c.Client.Do(ctx, http.MethodPut, path, request, nil, c.addApiPrefix, c.scimVisitor)
}

func (c *DatabricksClient) recursiveMask(requestMap map[string]any) any {
	for k, v := range requestMap {
		if k == "string_value" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if k == "token_value" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if k == "secret" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if k == "content" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if m, ok := v.(map[string]any); ok {
			requestMap[k] = c.recursiveMask(m)
			continue
		}
		// todo: dapi...
		// TODO: just redact any dapiXXX & "secret": "...."...
		if s, ok := v.(string); ok {
			requestMap[k] = onlyNBytes(s, c.DebugTruncateBytes)
		}
	}
	return requestMap
}

func (c *DatabricksClient) redactedDump(body []byte) (res string) {
	if len(body) == 0 {
		return
	}
	if body[0] != '{' {
		return fmt.Sprintf("[non-JSON document of %d bytes]", len(body))
	}
	var requestMap map[string]any
	err := json.Unmarshal(body, &requestMap)
	if err != nil {
		// error in this case is not much relevant
		return
	}
	rePacked, err := json.MarshalIndent(c.recursiveMask(requestMap), "", "  ")
	if err != nil {
		// error in this case is not much relevant
		return
	}
	maxBytes := 1024
	if c.DebugTruncateBytes > maxBytes {
		maxBytes = c.DebugTruncateBytes
	}
	return onlyNBytes(string(rePacked), maxBytes)
}

func (c *DatabricksClient) userAgent(ctx context.Context) string {
	resource := "unknown"
	terraformVersion := "unknown"
	if rn, ok := ctx.Value(ResourceName).(string); ok {
		resource = rn
	}
	if c.Provider != nil {
		terraformVersion = c.Provider.TerraformVersion
	}
	return fmt.Sprintf("databricks-tf-provider/%s (+%s) terraform/%s",
		Version(), resource, terraformVersion)
}

// CWE-117 prevention
func escapeNewLines(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	return in
}

func (c *DatabricksClient) createDebugHeaders(header http.Header, host string) string {
	headers := ""
	if c.DebugHeaders {
		if host != "" {
			headers += fmt.Sprintf("\n * Host: %s", escapeNewLines(host))
		}
		for k, v := range header {
			trunc := onlyNBytes(strings.Join(v, ""), c.DebugTruncateBytes)
			headers += fmt.Sprintf("\n * %s: %s", k, escapeNewLines(trunc))
		}
		if len(headers) > 0 {
			headers += "\n"
		}
	}
	return headers
}

// todo: do is better name
func (c *DatabricksClient) genericQuery(ctx context.Context, method, requestURL string, data any,
	visitors ...func(*http.Request) error) (body []byte, err error) {
	if c.httpClient == nil {
		return nil, fmt.Errorf("DatabricksClient is not configured")
	}
	if err = c.rateLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limited: %w", err)
	}
	requestBody, err := makeRequestBody(method, &requestURL, data)
	if err != nil {
		return nil, fmt.Errorf("request marshal: %w", err)
	}
	request, err := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	request.Header.Set("User-Agent", c.userAgent(ctx))
	for _, requestVisitor := range visitors {
		err = requestVisitor(request)
		if err != nil {
			return nil, fmt.Errorf("failed visitor: %w", err)
		}
	}
	headers := c.createDebugHeaders(request.Header, c.Host)
	log.Printf("[DEBUG] %s %s %s%v", method, escapeNewLines(request.URL.Path),
		headers, c.redactedDump(requestBody)) // lgtm [go/log-injection] lgtm [go/clear-text-logging]

	r, err := retryablehttp.FromRequest(request)
	if err != nil {
		return nil, err // no error invariants possible because of `makeRequestBody`
	}
	resp, err := c.httpClient.Do(r)
	// retryablehttp library now returns only wrapped errors
	var ae APIError
	if errors.As(err, &ae) {
		// don't re-wrap, as upper layers may depend on handling common.APIError
		return nil, ae
	}
	if err != nil {
		// i don't even know which errors in the real world would end up here.
		// `retryablehttp` package nicely wraps _everything_ to `url.Error`.
		return nil, fmt.Errorf("failed request: %w", err)
	}
	defer func() {
		if ferr := resp.Body.Close(); ferr != nil {
			err = fmt.Errorf("failed to close: %w", ferr)
		}
	}()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body: %w", err)
	}
	headers = c.createDebugHeaders(resp.Header, "")
	log.Printf("[DEBUG] %s %s %s <- %s %s", resp.Status, headers, c.redactedDump(body), method, strings.ReplaceAll(request.URL.Path, "\n", ""))
	return body, nil
}

func makeQueryString(data any) (string, error) {
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
		return "?" + strings.Join(s, "&"), nil
	}
	if inputType.Kind() == reflect.Struct {
		params, err := query.Values(data)
		if err != nil {
			return "", fmt.Errorf("cannot create query string: %w", err)
		}
		return "?" + params.Encode(), nil
	}
	return "", fmt.Errorf("unsupported query string data: %#v", data)
}

func makeRequestBody(method string, requestURL *string, data any) ([]byte, error) {
	var requestBody []byte
	if data == nil && (method == "DELETE" || method == "GET") {
		return requestBody, nil
	}
	if method == "GET" {
		qs, err := makeQueryString(data)
		if err != nil {
			return nil, err
		}
		*requestURL += qs
		return requestBody, nil
	}
	if reader, ok := data.(io.Reader); ok {
		raw, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("failed to read from reader: %w", err)
		}
		return raw, nil
	}
	if str, ok := data.(string); ok {
		return []byte(str), nil
	}
	bodyBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("request marshal failure: %w", err)
	}
	return bodyBytes, nil
}

func onlyNBytes(j string, numBytes int) string {
	diff := len([]byte(j)) - numBytes
	if diff > 0 {
		return fmt.Sprintf("%s... (%d more bytes)", j[:numBytes], diff)
	}
	return j
}
