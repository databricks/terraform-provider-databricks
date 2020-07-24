package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
	"github.com/r3labs/diff"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env") // TODO: make teardowns work here as well
	// TODO: add common instance pool & cluster for libs & stuff
	log.SetFlags(log.Lshortfile | log.Ltime)
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

func GetIntegrationMWSAPIClient() *DatabricksClient {
	client := DatabricksClient{
		Host:     os.Getenv("DATABRICKS_MWS_HOST"),
		Username: os.Getenv("DATABRICKS_USERNAME"),
		Password: os.Getenv("DATABRICKS_PASSWORD"),
	}
	err := client.Configure("dev")
	if err != nil {
		panic(err)
	}
	return &client
}

func GetCloudInstanceType(c *DatabricksClient) string {
	if c.IsAzure() {
		return "Standard_DS3_v2"
	}
	// TODO: create a method on ClustersAPI to give
	// cloud specific delta-cache enabled instance by default.
	return "m4.large"
}

func AssertRequestWithMockServer(t *testing.T, rawPayloadArgs interface{}, requestMethod string, requestURI string, input interface{}, response string, responseStatus int, want interface{}, wantErr bool, apiCall func(client DatabricksClient) (interface{}, error)) {
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
	client := DatabricksClient{
		Host:  server.URL,
		Token: "...",
	}
	err := client.Configure("dev")
	assert.NoError(t, err, fmt.Sprintf("Expected no error but got: %v", err))
	output, err := apiCall(client)

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

func AssertMultipleRequestsWithMockServer(t *testing.T, rawPayloadArgs interface{}, requestMethod []string, requestURI []string, input interface{}, response []string, responseStatus []int, want interface{}, wantErr bool, apiCall func(client DatabricksClient) (interface{}, error)) {
	counter := 0
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		if counter == len(requestMethod) {
			t.Fatalf("Received more requests than expected")
			return
		}
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
	client := DatabricksClient{
		Host:  server.URL,
		Token: "...",
	}
	err := client.Configure("dev")
	assert.NoError(t, err, fmt.Sprintf("Expected no error but got: %v", err))
	output, err := apiCall(client)

	if output != nil {
		compare(t, want, output)
	}

	if wantErr {
		assert.Error(t, err, err)
	} else {
		assert.NoError(t, err, err)
	}
}
