package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type CloudServiceProvider string

const (
	AWS   CloudServiceProvider = "AmazonWebServices"
	Azure CloudServiceProvider = "Azure"
)

type DBApiErrorBody struct {
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	// The following two are for scim api only for RFC 7644 Section 3.7.3 https://tools.ietf.org/html/rfc7644#section-3.7.3
	ScimDetail string `json:"detail,omitempty"`
	ScimStatus string `json:"status,omitempty"`
}

type DBApiError struct {
	ErrorBody  *DBApiErrorBody
	StatusCode int
	Err        error
}

func (r DBApiError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

type AuthType string

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
	cloudProvider      CloudServiceProvider
}

// Setup initializes the client
func (c *DBApiClientConfig) Setup() {
	if c.TimeoutSeconds == 0 {
		c.TimeoutSeconds = 10
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
			"User-Agent": fmt.Sprintf("databricks-go-client-sdk"),
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
	} else {
		return j
	}
}

func auditNonGetPayload(method string, uri string, object interface{}, mask *SecretsMask) {
	logStmt := struct {
		Method  string
		Uri     string
		Payload interface{}
	}{
		Method:  method,
		Uri:     uri,
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
		Uri    string
	}{
		Method: "GET",
		Uri:    uri,
	}
	jsonStr, _ := json.Marshal(Mask(logStmt))
	if mask != nil {
		log.Println(onlyNBytes(mask.MaskString(string(jsonStr)), 1e3))
	} else {
		log.Println(onlyNBytes(string(jsonStr), 1e3))
	}
}

func PerformQuery(config *DBApiClientConfig, method, path string, apiVersion string, headers map[string]string, marshalJson bool, useRawPath bool, data interface{}, secretsMask *SecretsMask) (body []byte, err error) {
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

	if headers != nil && len(headers) > 0 {
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
		if marshalJson {
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
