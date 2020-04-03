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

// DBClientOption is used to configure the DataBricks Client
type DBClientOption struct {
	Host               string
	Token              string
	DefaultHeaders     map[string]string
	InsecureSkipVerify bool
	TimeoutSeconds     int
	client             http.Client
	cloudProvider      CloudServiceProvider
}

// Init initializes the client
func (o *DBClientOption) Init() {
	if o.TimeoutSeconds == 0 {
		o.TimeoutSeconds = 10
	}
	o.client = http.Client{
		Timeout: time.Duration(time.Duration(o.TimeoutSeconds) * time.Second),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: o.InsecureSkipVerify,
			},
		},
	}
}

func (o *DBClientOption) getHTTPClient() http.Client {
	return o.client
}

func (o *DBClientOption) getAuthHeader() map[string]string {
	auth := make(map[string]string)
	auth["Authorization"] = "Bearer " + o.Token
	auth["Content-Type"] = "application/json"
	return auth
}

func (o *DBClientOption) getUserAgentHeader() map[string]string {
	return map[string]string{
		"User-Agent": fmt.Sprintf("databricks-sdk-golang-%s", SdkVersion),
	}
}

func (o *DBClientOption) getDefaultHeaders() map[string]string {
	auth := o.getAuthHeader()
	userAgent := o.getUserAgentHeader()

	defaultHeaders := make(map[string]string)
	for k, v := range auth {
		defaultHeaders[k] = v
	}
	for k, v := range o.DefaultHeaders {
		defaultHeaders[k] = v
	}
	for k, v := range userAgent {
		defaultHeaders[k] = v
	}
	return defaultHeaders
}

func (o *DBClientOption) getRequestURI(path string, apiVersion string) (string, error) {
	var apiVersionString string
	if apiVersion == "" {
		apiVersionString = "2.0"
	} else {
		apiVersionString = apiVersion
	}

	parsedURI, err := url.Parse(o.Host)
	if err != nil {
		return "", err
	}
	requestURI := fmt.Sprintf("%s://%s/api/%s%s", parsedURI.Scheme, parsedURI.Host, apiVersionString, path)
	return requestURI, nil
}

// PerformQuery can be used in a client or directly
func PerformQuery(option DBClientOption, method, path string, apiVersion string, headers map[string]string, marshalJson bool, useRawPath bool, data interface{}) ([]byte, error) {

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
	} else {
		if marshalJson {
			bodyBytes, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}

			requestBody = bodyBytes
			log.Println(string(requestBody))
		} else {
			requestBody = []byte(data.(string))
		}
	}

	client := option.getHTTPClient()

	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	for k, v := range requestHeaders {
		request.Header.Set(k, v)
	}

	resp, err := client.Do(request)
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
