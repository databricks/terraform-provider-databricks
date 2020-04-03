package service

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/r3labs/diff"
	"github.com/databrickslabs/databricks-terraform/client"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Failed to load environment")
	}
	code := m.Run()
	os.Exit(code)
}

func DeserializeJson(req *http.Request, m interface{}) error {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&m)
	return err
}

func compare(t *testing.T, a interface{}, b interface{}) {
	difference, err := diff.Diff(a, b)
	assert.NoError(t, err, err)
	jsonStr, _ := json.Marshal(difference)
	assert.True(t, reflect.DeepEqual(a, b), string(jsonStr))
}

func GetIntegrationDBAPIClient() *DBApiClient {
	var o client.DBClientOption
	o.Token = os.Getenv("TOKEN")
	o.Host = os.Getenv("HOST")

	var c DBApiClient
	c.Init(o)
	return &c
}

func GetCloudInstanceType(c *DBApiClient) string {
	if strings.Contains(c.Option.Host, "azure") {
		return "Standard_DS3_v2"
	} else {
		return "i3.xlarge"
	}
}

func AssertRequestWithMockServer(t *testing.T, rawPayloadArgs interface{}, requestMethod string, requestURI string, input interface{}, response string, want interface{}, wantErr bool, apiCall func(client DBApiClient) (interface{}, error)) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.Method, requestMethod)
		assert.Equal(t, req.RequestURI, requestURI)
		if requestMethod == http.MethodPost {
			err := DeserializeJson(req, &input)
			assert.NoError(t, err, err)
			t.Log(input)
			t.Log(rawPayloadArgs)
			compare(t, rawPayloadArgs, input)
		}

		_, err := rw.Write([]byte(response))
		assert.NoError(t, err, err)

	}))
	// Close the server when test finishes
	defer server.Close()
	var o client.DBClientOption
	o.Host = server.URL

	dbClient := DBApiClient{Option: o}

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
