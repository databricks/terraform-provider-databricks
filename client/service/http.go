package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
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

// IsTooManyRequests ...
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

var transientErrorStringMatches []string = []string{
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

func (c *DatabricksClient) get(path string, request interface{}, response interface{}) error {
	if c.authVisitor == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	body, err := c.genericQuery2(http.MethodGet, path, request, c.authVisitor, c.api2)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

func (c *DatabricksClient) post(path string, request interface{}, response interface{}) error {
	if c.authVisitor == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	body, err := c.genericQuery2(http.MethodPost, path, request, c.authVisitor, c.api2)
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

func (c *DatabricksClient) delete(path string, request interface{}) error {
	if c.authVisitor == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	_, err := c.genericQuery2(http.MethodDelete, path, request, c.authVisitor, c.api2)
	return err
}

func (c *DatabricksClient) patch(path string, request interface{}) error {
	if c.authVisitor == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	_, err := c.genericQuery2(http.MethodPatch, path, request, c.authVisitor, c.api2)
	return err
}

func (c *DatabricksClient) put(path string, request interface{}) error {
	if c.authVisitor == nil {
		return fmt.Errorf("Authentication not initialized")
	}
	_, err := c.genericQuery2(http.MethodPut, path, request, c.authVisitor, c.api2)
	return err
}

// TODO: rename to internal or something...
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

// TODO: rename to internal or something...
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

func (c *DatabricksClient) performScim(method, path string, request interface{}, response interface{}) error {
	body, err := c.genericQuery2(method, path, request, c.authVisitor,
		c.api2, func(r *http.Request) error {
			r.Header.Set("Content-Type", "application/scim+json")
			return nil
		})
	if err != nil {
		return err
	}
	return c.unmarshall(path, body, &response)
}

// todo: do is better name
func (c *DatabricksClient) genericQuery2(method, requestURL string, data interface{},
	visitors ...func(*http.Request) error) (body []byte, err error) {
	if c.httpClient == nil {
		return nil, fmt.Errorf("DatabricksClient is not configured")
	}
	requestBody, err := makeRequestBody(method, &requestURL, data, true)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	// TODO: add masking **REDACTED** + subscription + mws acc id -> perhaps re-parse json and truncate only some fields?...
	log.Printf("[INFO] %s %s %v", method, requestURL, onlyNBytes(string(requestBody), 512))
	request.Header.Set("User-Agent", c.userAgent)
	for _, requestVisitor := range visitors {
		err = requestVisitor(request)
		if err != nil {
			return nil, err
		}
	}
	r, err := retryablehttp.FromRequest(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(r)
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

func makeRequestBody(method string, requestURL *string, data interface{}, marshalJSON bool) ([]byte, error) {
	var requestBody []byte
	if method == "GET" {
		if m, ok := data.(map[string]interface{}); ok {
			s := []string{}
			for k, v := range m {
				if v == "" {
					continue
				}
				s = append(s, fmt.Sprintf("%s=%s", k, url.PathEscape(fmt.Sprintf("%v", v))))
			}
			*requestURL += "?" + strings.Join(s, "&")
		} else if m, ok := data.(map[string]int64); ok {
			s := []string{}
			for k, v := range m {
				if v == 0 {
					continue
				}
				s = append(s, fmt.Sprintf("%s=%s", k, url.PathEscape(fmt.Sprintf("%v", v))))
			}
			*requestURL += "?" + strings.Join(s, "&")
		} else if m, ok := data.(map[string]string); ok {
			s := []string{}
			for k, v := range m {
				if v == "" {
					continue
				}
				s = append(s, fmt.Sprintf("%s=%s", k, url.PathEscape(v)))
			}
			*requestURL += "?" + strings.Join(s, "&")
		} else {
			params, err := query.Values(data)
			if err != nil {
				return nil, err
			}
			*requestURL += "?" + params.Encode()
		}
		// auditGetPayload(*requestURL)
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
		// auditNonGetPayload(method, *requestURL, data)
	}
	return requestBody, nil
}

func onlyNBytes(j string, numBytes int64) string {
	if len([]byte(j)) > int(numBytes) {
		return string([]byte(j)[:numBytes])
	}
	return j
}

// TODO: port to new way of request logging
// func auditNonGetPayload(method string, uri string, object interface{}) {
// 	logStmt := struct {
// 		Method  string
// 		URI     string
// 		Payload interface{}
// 	}{
// 		Method:  method,
// 		URI:     uri,
// 		Payload: object,
// 	}
// 	jsonStr, _ := json.Marshal(Mask(logStmt))
// 	log.Println(onlyNBytes(string(jsonStr), 1e3))
// }

// func auditGetPayload(uri string) {
// 	logStmt := struct {
// 		Method string
// 		URI    string
// 	}{
// 		Method: "GET",
// 		URI:    uri,
// 	}
// 	jsonStr, _ := json.Marshal(Mask(logStmt))
// 	log.Println(onlyNBytes(string(jsonStr), 1e3))
// }
