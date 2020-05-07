package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

// CloudServiceProvider is a custom type for different types of cloud service providers
type CloudServiceProvider string

// List of CloudServiceProviders Databricks is available on
const (
	AWS   CloudServiceProvider = "AmazonWebServices"
	Azure CloudServiceProvider = "Azure"
)

// DBApiErrorBody is a struct for a custom api error for all the services on databrickss.
type DBApiErrorBody struct {
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	// The following two are for scim api only for RFC 7644 Section 3.7.3 https://tools.ietf.org/html/rfc7644#section-3.7.3
	ScimDetail string `json:"detail,omitempty"`
	ScimStatus string `json:"status,omitempty"`
}

// DBApiError is a generic struct for an api error on databricks
type DBApiError struct {
	ErrorBody  *DBApiErrorBody
	StatusCode int
	Err        error
}

// Error is a interface implementation of the error interface.
func (r DBApiError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

// AuthType is a custom type for a type of authentication allowed on Databricks
type AuthType string

// List of AuthTypes supported by this go sdk.
const (
	BasicAuth AuthType = "BASIC"
)

// DBApiClientConfig is used to configure the DataBricks Client
type DBApiClientConfig struct {
	Host               string
	Token              string
	AuthType           AuthType
	UserAgent          string
	DefaultHeaders     map[string]string
	InsecureSkipVerify bool
	TimeoutSeconds     int
	client             http.Client
}

// Setup initializes the client
func (c *DBApiClientConfig) Setup() {
	if c.TimeoutSeconds == 0 {
		c.TimeoutSeconds = 60
	}
	c.client = http.Client{
		Timeout: time.Duration(time.Duration(c.TimeoutSeconds) * time.Second),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.InsecureSkipVerify,
			},
		},
	}
}

func (c DBApiClientConfig) getAuthHeader() map[string]string {
	auth := make(map[string]string)
	if c.AuthType == BasicAuth {
		auth["Authorization"] = "Basic " + c.Token
	} else {
		auth["Authorization"] = "Bearer " + c.Token
	}
	auth["Content-Type"] = "application/json"
	return auth
}

func (c DBApiClientConfig) getUserAgentHeader() map[string]string {
	if reflect.ValueOf(c.UserAgent).IsZero() {
		return map[string]string{
			"User-Agent": "databricks-go-client-sdk",
		}
	}
	return map[string]string{
		"User-Agent": c.UserAgent,
	}
}

func (c DBApiClientConfig) getDefaultHeaders() map[string]string {
	auth := c.getAuthHeader()
	userAgent := c.getUserAgentHeader()

	defaultHeaders := make(map[string]string)
	for k, v := range auth {
		defaultHeaders[k] = v
	}
	for k, v := range c.DefaultHeaders {
		defaultHeaders[k] = v
	}
	for k, v := range userAgent {
		defaultHeaders[k] = v
	}
	return defaultHeaders
}

func (c DBApiClientConfig) getRequestURI(path string, apiVersion string) (string, error) {
	var apiVersionString string
	if apiVersion == "" {
		apiVersionString = "2.0"
	} else {
		apiVersionString = apiVersion
	}

	parsedURI, err := url.Parse(c.Host)
	if err != nil {
		return "", err
	}
	requestURI := fmt.Sprintf("%s://%s/api/%s%s", parsedURI.Scheme, parsedURI.Host, apiVersionString, path)
	return requestURI, nil
}

func onlyNBytes(j string, numBytes int64) string {
	if len([]byte(j)) > int(numBytes) {
		return string([]byte(j)[:numBytes])
	}
	return j
}

func auditNonGetPayload(method string, uri string, object interface{}, mask *SecretsMask) {
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
	if mask != nil {
		log.Println(onlyNBytes(mask.MaskString(string(jsonStr)), 1e3))
	} else {
		log.Println(onlyNBytes(string(jsonStr), 1e3))
	}
}

func auditGetPayload(uri string, mask *SecretsMask) {
	logStmt := struct {
		Method string
		URI    string
	}{
		Method: "GET",
		URI:    uri,
	}
	jsonStr, _ := json.Marshal(Mask(logStmt))
	if mask != nil {
		log.Println(onlyNBytes(mask.MaskString(string(jsonStr)), 1e3))
	} else {
		log.Println(onlyNBytes(string(jsonStr), 1e3))
	}
}

// PerformQuery is a generic function that accepts a config, method, path, apiversion, headers,
// and some flags to perform query against the Databricks api
func PerformQuery(config *DBApiClientConfig, method, path string, apiVersion string, headers map[string]string, marshalJSON bool, useRawPath bool, data interface{}, secretsMask *SecretsMask) (body []byte, err error) {
	var requestURL string
	if useRawPath {
		requestURL = path
	} else {
		requestURL, err = config.getRequestURI(path, apiVersion)
		if err != nil {
			return nil, err
		}
	}
	requestHeaders := config.getDefaultHeaders()

	if len(headers) > 0 {
		for k, v := range headers {
			requestHeaders[k] = v
		}
	}

	var requestBody []byte
	if method == "GET" {
		params, err := query.Values(data)

		if err != nil {
			return nil, err
		}
		requestURL += "?" + params.Encode()
		auditGetPayload(requestURL, secretsMask)

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
		auditNonGetPayload(method, requestURL, data, secretsMask)
	}

	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	for k, v := range requestHeaders {
		request.Header.Set(k, v)
	}

	resp, err := config.client.Do(request)
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

	if resp.StatusCode >= 400 {
		var errorBody DBApiErrorBody
		err = json.Unmarshal(body, &errorBody)
		if err != nil {
			return nil, fmt.Errorf("Response from server (%d) %s", resp.StatusCode, string(body))
		}
		return nil, DBApiError{
			ErrorBody:  &errorBody,
			StatusCode: resp.StatusCode,
			Err:        fmt.Errorf("Response from server %s", string(body)),
		}
	}

	return body, nil
}
