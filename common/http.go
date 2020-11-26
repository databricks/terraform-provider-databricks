package common

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	e2example                   = "https://github.com/databrickslabs/terraform-provider-databricks/blob/master/scripts/awsmt-integration/main.tf"
	accountsHost                = "accounts.cloud.databricks.com"
	transientErrorStringMatches = []string{
		"com.databricks.backend.manager.util.UnknownWorkerEnvironmentException",
		"does not have any associated worker environments",
		"There is no worker environment with id",
		"ClusterNotReadyException",
		"connection reset by peer",
		"connection refused",
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
	docs := apiError.DocumentationURL()
	if docs == "" {
		return fmt.Sprintf("%s\n(%d on %s)", apiError.Message, apiError.StatusCode, apiError.Resource)
	}
	return fmt.Sprintf("%s\nPlease consult API docs at %s for details.",
		apiError.Message,
		docs)
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

func (c *DatabricksClient) commonErrorClarity(resp *http.Response) *APIError {
	isAccountsAPI := strings.HasPrefix(resp.Request.URL.Path, "/api/2.0/accounts")
	isAccountsClient := strings.Contains(c.Host, accountsHost)
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIError{
			Message:    err.Error(),
			ErrorCode:  "IO_READ",
			StatusCode: resp.StatusCode,
			Resource:   resp.Request.URL.Path,
		}
	}
	log.Printf("[INFO] %s %v", resp.Status, c.redactedDump(body))
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
		errorBody.ErrorCode = fmt.Sprintf("SCIM_%s", errorBody.ScimStatus)
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
		apiError := APIError{ErrorCode: "IO_ERROR", Message: ue.Error()}
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
func (c *DatabricksClient) Get(ctx context.Context, path string, request interface{}, response interface{}) error {
	body, err := c.authenticatedQuery(ctx, http.MethodGet, path, request, c.api2)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

// Post on path
func (c *DatabricksClient) Post(ctx context.Context, path string, request interface{}, response interface{}) error {
	body, err := c.authenticatedQuery(ctx, http.MethodPost, path, request, c.api2)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

// Delete on path
func (c *DatabricksClient) Delete(ctx context.Context, path string, request interface{}) error {
	_, err := c.authenticatedQuery(ctx, http.MethodDelete, path, request, c.api2)
	return err
}

// Patch on path
func (c *DatabricksClient) Patch(ctx context.Context, path string, request interface{}) error {
	_, err := c.authenticatedQuery(ctx, http.MethodPatch, path, request, c.api2)
	return err
}

// Put on path
func (c *DatabricksClient) Put(ctx context.Context, path string, request interface{}) error {
	_, err := c.authenticatedQuery(ctx, http.MethodPut, path, request, c.api2)
	return err
}

func (c *DatabricksClient) unmarshall(path string, body []byte, response interface{}) error {
	if response == nil {
		return nil
	}
	if len(body) == 0 {
		return nil
	}
	err := json.Unmarshal(body, &response)
	if err == nil {
		return nil
	}
	return APIError{
		ErrorCode:  "UNKNOWN",
		StatusCode: 200,
		Resource:   "..." + path,
		Message: fmt.Sprintf("Invalid JSON received (%d bytes): %v",
			len(body), string(body)),
	}
}

func (c *DatabricksClient) api2(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("No URL found in request")
	}
	r.URL.Path = fmt.Sprintf("/api/2.0%s", r.URL.Path)
	r.Header.Set("Content-Type", "application/json")

	url, err := url.Parse(c.Host)
	if err != nil {
		return err
	}
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme

	return nil
}

func (c *DatabricksClient) api12(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("No URL found in request")
	}
	r.URL.Path = fmt.Sprintf("/api/1.2%s", r.URL.Path)
	r.Header.Set("Content-Type", "application/json")

	url, err := url.Parse(c.Host)
	if err != nil {
		return err
	}
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme

	return nil
}

// Scim sets SCIM headers
func (c *DatabricksClient) Scim(ctx context.Context, method, path string, request interface{}, response interface{}) error {
	body, err := c.authenticatedQuery(ctx, method, path, request, c.api2, func(r *http.Request) error {
		r.Header.Set("Content-Type", "application/scim+json")
		return nil
	})
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

// OldAPI performs call on context api
func (c *DatabricksClient) OldAPI(ctx context.Context, method, path string, request interface{}, response interface{}) error {
	body, err := c.authenticatedQuery(ctx, method, path, request, c.api12)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

func (c *DatabricksClient) authenticatedQuery(ctx context.Context, method, requestURL string,
	data interface{}, visitors ...func(*http.Request) error) (body []byte, err error) {
	err = c.Authenticate()
	if err != nil {
		return
	}
	visitors = append([]func(*http.Request) error{c.authVisitor}, visitors...)
	return c.genericQuery(ctx, method, requestURL, data, visitors...)
}

func (c *DatabricksClient) recursiveMask(requestMap map[string]interface{}) interface{} {
	for k, v := range requestMap {
		if k == "string_value" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if k == "token_value" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if k == "content" {
			requestMap[k] = "**REDACTED**"
			continue
		}
		if m, ok := v.(map[string]interface{}); ok {
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
	var requestMap map[string]interface{}
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
	return onlyNBytes(string(rePacked), 1024)
}

func (c *DatabricksClient) userAgent(ctx context.Context) string {
	resource := "unknown"
	terraformVersion := "unknown"
	if rn, ok := ctx.Value(ResourceName).(string); ok {
		resource = rn
	}
	if tv, ok := ctx.Value(TerraformVersion).(string); ok {
		terraformVersion = tv
	}
	return fmt.Sprintf("databricks-tf-provider/%s (+%s) terraform/%s", Version(), resource, terraformVersion)
}

// todo: do is better name
func (c *DatabricksClient) genericQuery(ctx context.Context, method, requestURL string, data interface{},
	visitors ...func(*http.Request) error) (body []byte, err error) {
	if c.httpClient == nil {
		return nil, fmt.Errorf("DatabricksClient is not configured")
	}
	requestBody, err := makeRequestBody(method, &requestURL, data, true)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", c.userAgent(ctx))
	for _, requestVisitor := range visitors {
		err = requestVisitor(request)
		if err != nil {
			return nil, err
		}
	}
	headers := ""
	if c.DebugHeaders {
		for k, v := range request.Header {
			headers += fmt.Sprintf("\n * %s: %s", k, onlyNBytes(strings.Join(v, ""), 16))
		}
		if len(headers) > 0 {
			headers += "\n"
		}
	}
	log.Printf("[DEBUG] %s %s %s%v", method, requestURL, headers, c.redactedDump(requestBody))

	r, err := retryablehttp.FromRequest(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(r)
	// retryablehttp library now returns only wrapped errors
	var ae APIError
	if errors.As(err, &ae) {
		return nil, ae
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		if ferr := resp.Body.Close(); ferr != nil {
			err = ferr
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] %s %v <- %s %s", resp.Status, c.redactedDump(body), method, requestURL)
	return body, nil
}

func makeRequestBody(method string, requestURL *string, data interface{}, marshalJSON bool) ([]byte, error) {
	var requestBody []byte
	if method == "GET" {
		if data == nil {
			return requestBody, nil
		}
		inputVal := reflect.ValueOf(data)
		inputType := reflect.TypeOf(data)
		switch inputType.Kind() {
		case reflect.Map:
			s := []string{}
			for _, k := range inputVal.MapKeys() {
				v := inputVal.MapIndex(k)
				if v.IsZero() {
					continue
				}
				s = append(s, fmt.Sprintf("%v=%s", k.Interface(),
					url.PathEscape(fmt.Sprintf("%v", v.Interface()))))
			}
			*requestURL += "?" + strings.Join(s, "&")
		case reflect.Struct:
			params, err := query.Values(data)
			if err != nil {
				return nil, err
			}
			*requestURL += "?" + params.Encode()
		default:
			return requestBody, fmt.Errorf("Unsupported request data: %#v", data)
		}
	} else {
		if marshalJSON {
			bodyBytes, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return nil, err
			}
			requestBody = bodyBytes
		} else {
			requestBody = []byte(data.(string))
		}
	}
	return requestBody, nil
}

func onlyNBytes(j string, numBytes int) string {
	diff := len([]byte(j)) - numBytes
	if diff > 0 {
		return fmt.Sprintf("%s... (%d more bytes)", j[:numBytes], diff)
	}
	return j
}
