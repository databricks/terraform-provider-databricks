package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

// CloudServiceProvider is a custom type for different types of cloud service providers
type CloudServiceProvider string

// List of CloudServiceProviders Databricks is available on
const (
	AWS   CloudServiceProvider = "AmazonWebServices"
	Azure CloudServiceProvider = "Azure"
)

// APIErrorBody maps "proper" databricks rest api errors to a struct
type APIErrorBody struct {
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	// The following two are for scim api only for RFC 7644 Section 3.7.3 https://tools.ietf.org/html/rfc7644#section-3.7.3
	ScimDetail string `json:"detail,omitempty"`
	ScimStatus string `json:"status,omitempty"`
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

var clientAuthorizerMutex sync.Mutex

var transientErrorStringMatches []string = []string{ // TODO: Should we make these regexes to match more of the message or is this sufficient?
	"com.databricks.backend.manager.util.UnknownWorkerEnvironmentException",
	"does not have any associated worker environments",
	"There is no worker environment with id",
}

// checkHTTPRetry inspects HTTP errors from the Databricks API for known transient errors on Workspace creation
func checkHTTPRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if resp == nil {
		// If response is nil we can't make retry choices.
		// In this case don't retry and return the original error from httpclient
		return false, err
	}
	if resp.StatusCode >= 400 {
		log.Printf("Failed request detected. Status Code: %v\n", resp.StatusCode)
		// reading the body means that the caller cannot read it themselves
		// But that's ok because we've hit an error case
		// Our job now is to
		//  - capture the error and return it
		//  - determine if the error is retryable

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}
		var errorBody APIErrorBody
		err = json.Unmarshal(body, &errorBody)
		// this is most likely HTML... since un-marshalling JSON failed
		if err != nil {
			// Status parts first in case html message is not as expected
			statusParts := strings.SplitN(resp.Status, " ", 2)
			if len(statusParts) < 2 {
				errorBody.ErrorCode = "UNKNOWN"
			} else {
				errorBody.ErrorCode = strings.ReplaceAll(strings.ToUpper(strings.Trim(statusParts[1], " .")), " ", "_")
			}
			stringBody := string(body)
			messageRE := regexp.MustCompile(`<pre>(.*)</pre>`)
			messageMatches := messageRE.FindStringSubmatch(stringBody)
			// No messages with <pre> </pre> format found so return a default APIError
			if len(messageMatches) < 2 {
				return false, APIError{
					Message:    fmt.Sprintf("Response from server (%d) %s: %v", resp.StatusCode, stringBody, err),
					ErrorCode:  errorBody.ErrorCode,
					StatusCode: resp.StatusCode,
					Resource:   resp.Request.URL.Path,
				}
			}
			errorBody.Message = strings.Trim(messageMatches[1], " .")
		}
		dbAPIError := APIError{
			Message:    errorBody.Message,
			ErrorCode:  errorBody.ErrorCode,
			StatusCode: resp.StatusCode,
			Resource:   resp.Request.URL.Path,
		}
		// Handle scim error message details
		if dbAPIError.Message == "" && errorBody.ScimDetail != "" {
			if errorBody.ScimDetail == "null" {
				dbAPIError.Message = "SCIM API Internal Error"
			} else {
				dbAPIError.Message = errorBody.ScimDetail
			}
			dbAPIError.ErrorCode = fmt.Sprintf("SCIM_%s", errorBody.ScimStatus)
		}
		// Handle transient errors for retries
		for _, substring := range transientErrorStringMatches {
			if strings.Contains(errorBody.Message, substring) {
				log.Println("Failed request detected: Retryable type found. Attempting retry...")
				return true, dbAPIError
			}
		}
		return false, dbAPIError
	}
	return false, nil
}

func (c *DatabricksClient) getOrCreateToken() error {
	if c.customAuthorizer != nil {
		// Lock incase terraform tries to getOrCreateToken from multiple go routines on the same client ptr.
		clientAuthorizerMutex.Lock()
		defer clientAuthorizerMutex.Unlock()
		// todo: introduce a property?....
		if reflect.ValueOf(c.Token).IsZero() {
			log.Println("NOT AUTHORIZED SO ATTEMPTING TO AUTHORIZE")
			return c.customAuthorizer()
		}
		log.Println("ALREADY AUTHORIZED")
		return nil
	}
	return nil
}

func (c *DatabricksClient) performQuery(method string, path string, apiVersion string, headers map[string]string, data interface{}) ([]byte, error) {
	return c.genericQuery(method, path, apiVersion, headers, true, false, data)
}

func (c *DatabricksClient) performRawQuery(method, path string, apiVersion string, headers map[string]string, marshalJSON bool, data interface{}) (body []byte, err error) {
	return c.genericQuery(method, path, apiVersion, headers, marshalJSON, true, data)
}

func makeRequestBody(method string, requestURL *string, data interface{}, marshalJSON bool) ([]byte, error) {
	var requestBody []byte
	if method == "GET" {
		params, err := query.Values(data)
		if err != nil {
			return nil, err
		}
		*requestURL += "?" + params.Encode()
		auditGetPayload(*requestURL)
	} else {
		if marshalJSON {
			bodyBytes, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			requestBody = bodyBytes
		} else {
			requestBody = []byte(data.(string))
		}
		auditNonGetPayload(method, *requestURL, data)
	}
	return requestBody, nil
}

// PerformQuery is a generic function that accepts a config, method, path, apiversion, headers,
// and some flags to perform query against the Databricks api
func (c *DatabricksClient) genericQuery(method, path, apiVersion string, headers map[string]string,
	marshalJSON bool, useRawPath bool, data interface{}) (body []byte, err error) {
	if c.httpClient == nil {
		return []byte(""), fmt.Errorf("DatabricksClient is not configured")
	}
	err = c.getOrCreateToken()
	if err != nil {
		return []byte(""), err
	}
	if apiVersion == "" {
		apiVersion = "2.0"
	}
	requestURL := path
	if !useRawPath {
		requestURL = fmt.Sprintf("%s/%s%s", c.uriPrefix, apiVersion, path)
	}
	requestBody, err := makeRequestBody(method, &requestURL, data, marshalJSON)
	request, err := retryablehttp.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", c.userAgent)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", c.authType, c.Token))
	if len(headers) > 0 {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}
	resp, err := c.httpClient.Do(request)
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
	// Don't need to check the status code here as the RetryCheck for
	// retryablehttp.Client is doing that and returning an error
	return body, nil
}

func onlyNBytes(j string, numBytes int64) string {
	if len([]byte(j)) > int(numBytes) {
		return string([]byte(j)[:numBytes])
	}
	return j
}

func auditNonGetPayload(method string, uri string, object interface{}) {
	logStmt := struct {
		Method  string
		URI     string
		Payload interface{}
	}{
		Method:  method,
		URI:     uri,
		Payload: object,
	}
	jsonStr, _ := json.Marshal(Mask(logStmt))
	log.Println(onlyNBytes(string(jsonStr), 1e3))
}

func auditGetPayload(uri string) {
	logStmt := struct {
		Method string
		URI    string
	}{
		Method: "GET",
		URI:    uri,
	}
	jsonStr, _ := json.Marshal(Mask(logStmt))
	log.Println(onlyNBytes(string(jsonStr), 1e3))
}
