package client

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
	"time"
)

type CloudServiceProvider string

const (
	AWS   CloudServiceProvider = "AmazonWebServices"
	Azure CloudServiceProvider = "Azure"
)

// DBApiClientConfig is used to configure the DataBricks Client
type DBApiClientConfig struct {
	Host               string
	Token              string
	DefaultHeaders     map[string]string
	InsecureSkipVerify bool
	TimeoutSeconds     int
	client             http.Client
	cloudProvider      CloudServiceProvider
}

// SetConfig initializes the client
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
	auth["Authorization"] = "Bearer " + c.Token
	auth["Content-Type"] = "application/json"
	return auth
}

func (c DBApiClientConfig) getUserAgentHeader() map[string]string {
	return map[string]string{
		"User-Agent": fmt.Sprintf("databricks-go-client-sdk/%s", ClientVersion),
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

func onlyNBytes(buf []byte, numBytes int64) []byte {
	if len(buf) > int(numBytes) {
		return buf[:numBytes]
	} else {
		return buf
	}
}

func PerformQuery(option *DBApiClientConfig, method, path string, apiVersion string, headers map[string]string, marshalJson bool, useRawPath bool, data interface{}) ([]byte, error) {

	var requestURL string
	var err error
	if useRawPath {
		requestURL = path
	} else {
		requestURL, err = option.getRequestURI(path, apiVersion)
		if err != nil {
			return nil, err
		}
	}
	requestHeaders := option.getDefaultHeaders()

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
		log.Println(string(requestURL))
	} else {
		if marshalJson {
			bodyBytes, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}

			requestBody = bodyBytes
			log.Println(string(onlyNBytes(requestBody, 1e3)))
		} else {
			requestBody = []byte(data.(string))
		}
	}

	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	for k, v := range requestHeaders {
		request.Header.Set(k, v)
	}

	resp, err := option.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Response from server (%d) %s", resp.StatusCode, string(body))
	}

	return body, nil
}
