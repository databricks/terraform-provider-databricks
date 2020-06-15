package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/r3labs/diff"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Failed to load environment")
	}
	code := m.Run()
	os.Exit(code)
}

func DeserializeJSON(req *http.Request, m interface{}) error {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&m)
	return err
}

func compare(t *testing.T, a interface{}, b interface{}) {
	difference, err := diff.Diff(a, b)
	assert.NoError(t, err, err)
	jsonStr, err := json.Marshal(difference)
	assert.NoError(t, err, err)
	assert.True(t, reflect.DeepEqual(a, b), string(jsonStr))
}

func GetIntegrationDBAPIClient() *DBApiClient {
	var config DBApiClientConfig
	config.Token = os.Getenv("DATABRICKS_TOKEN")
	config.Host = os.Getenv("DATABRICKS_HOST")
	config.Setup()

	var c DBApiClient
	c.SetConfig(&config)
	return &c
}

func GetIntegrationMWSAPIClient() *DBApiClient {
	var config DBApiClientConfig
	tokenUnB64 := fmt.Sprintf("%s:%s", os.Getenv("DATABRICKS_USERNAME"), os.Getenv("DATABRICKS_PASSWORD"))
	config.AuthType = BasicAuth
	config.Token = base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
	config.Host = os.Getenv("DATABRICKS_MWS_HOST")

	var c DBApiClient
	c.SetConfig(&config)
	return &c
}

func GetCloudInstanceType(c *DBApiClient) string {
	if strings.Contains(c.Config.Host, "azure") {
		return "Standard_DS3_v2"
	}
	return "m4.large"
}

func AssertRequestWithMockServer(t *testing.T, rawPayloadArgs interface{}, requestMethod string, requestURI string, input interface{}, response string, responseStatus int, want interface{}, wantErr bool, apiCall func(client DBApiClient) (interface{}, error)) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, requestMethod, req.Method)
		assert.Equal(t, requestURI, req.RequestURI)
		if requestMethod == http.MethodPost || requestMethod == http.MethodPatch || requestMethod == http.MethodPut {
			err := DeserializeJSON(req, &input)
			assert.NoError(t, err, err)
			compare(t, rawPayloadArgs, input)
		}

		rw.WriteHeader(responseStatus)
		_, err := rw.Write([]byte(response))
		assert.NoError(t, err, err)
	}))
	// Close the server when test finishes
	defer server.Close()
	var config DBApiClientConfig
	config.Host = server.URL
	config.Setup()

	var dbClient DBApiClient
	dbClient.SetConfig(&config)

	output, err := apiCall(dbClient)

	assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(output), "Types are not equal between output of api call and expectiation!")
	if output != nil && !reflect.ValueOf(output).IsZero() {
		compare(t, want, output)
	}

	if wantErr {
		assert.Error(t, err, err)
	} else {
		assert.NoError(t, err, fmt.Sprintf("Expected no error but got: %v", err))
	}
}

func AssertMultipleRequestsWithMockServer(t *testing.T, rawPayloadArgs interface{}, requestMethod []string, requestURI []string, input interface{}, response []string, responseStatus []int, want interface{}, wantErr bool, apiCall func(client DBApiClient) (interface{}, error)) {
	counter := 0
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, requestMethod[counter], req.Method)
		assert.Equal(t, requestURI[counter], req.RequestURI)
		if requestMethod[counter] == http.MethodPost || requestMethod[counter] == http.MethodPatch || requestMethod[counter] == http.MethodPut {
			err := DeserializeJSON(req, &(input.([]interface{})[counter]))
			assert.NoError(t, err, err)
			compare(t, rawPayloadArgs.([]interface{})[counter], input.([]interface{})[counter])
		}

		rw.WriteHeader(responseStatus[counter])
		_, err := rw.Write([]byte(response[counter]))
		assert.NoError(t, err, err)
		counter++
	}))
	// Close the server when test finishes
	defer server.Close()
	var config DBApiClientConfig
	config.Host = server.URL
	config.Setup()

	var dbClient DBApiClient
	dbClient.SetConfig(&config)

	output, err := apiCall(dbClient)

	if output != nil {
		compare(t, want, output)
	}

	if wantErr {
		assert.Error(t, err, err)
	} else {
		assert.NoError(t, err, err)
	}
}
