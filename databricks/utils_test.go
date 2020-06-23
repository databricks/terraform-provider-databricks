package databricks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

// HTTPFixture defines request structure for test
type HTTPFixture struct {
	Method          string
	Resource        string
	Response        interface{}
	Status          int
	ExpectedRequest interface{}
}

// ResourceTester helps testing HTTP resources with fixtures
func ResourceTester(t *testing.T,
	fixtures []HTTPFixture,
	resouceFunc func() *schema.Resource,
	state map[string]interface{},
	whatever func(d *schema.ResourceData, c interface{}) error) (*schema.ResourceData, error) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		found := false
		for _, fixture := range fixtures {
			if req.Method == fixture.Method && req.RequestURI == fixture.Resource {
				if fixture.Status == 0 {
					rw.WriteHeader(200)
				} else {
					rw.WriteHeader(fixture.Status)
				}
				if fixture.ExpectedRequest != nil {
					buf := new(bytes.Buffer)
					_, err := buf.ReadFrom(req.Body)
					assert.NoError(t, err, err)
					jsonStr, err := json.Marshal(fixture.ExpectedRequest)
					assert.NoError(t, err, err)
					assert.JSONEq(t, string(jsonStr), buf.String())
				}
				if fixture.Response != nil {
					responseBytes, err := json.Marshal(fixture.Response)
					if err != nil {
						assert.NoError(t, err, err)
						t.FailNow()
					}
					_, err = rw.Write(responseBytes)
					assert.NoError(t, err, err)
				}
				found = true
				break
			}
		}
		if !found {
			assert.Fail(t, fmt.Sprintf("Received unexpected call: %s %s", req.Method, req.RequestURI))
			t.FailNow()
		}
	}))

	defer server.Close()
	var config service.DBApiClientConfig
	config.Host = server.URL
	config.Setup()

	var client service.DBApiClient
	client.SetConfig(&config)

	resource := resouceFunc()
	resourceData := schema.TestResourceDataRaw(t, resource.Schema, state)
	return resourceData, whatever(resourceData, &client)
}

func TestIsClusterMissingTrueWhenClusterIdSpecifiedPresent(t *testing.T) {
	errorMessage := "{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Cluster 123 does not exist\"}"

	result := isClusterMissing(errorMessage, "123")

	assert.True(t, result)
}

func TestIsClusterMissingFalseWhenClusterIdSpecifiedNotPresent(t *testing.T) {
	errorMessage := "{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Cluster 123 does not exist\"}"

	result := isClusterMissing(errorMessage, "xyz")

	assert.False(t, result)
}

func TestIsClusterMissingFalseWhenErrorNotInCorrectFormat(t *testing.T) {
	errorMessage := "{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Something random went bang xyz\"}"

	result := isClusterMissing(errorMessage, "xyz")

	assert.False(t, result)
}

func TestValidateInstanceProfileARN(t *testing.T) {
	testCases := []struct {
		instanceProfileARN string
		errorCount         int
	}{
		{"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", 0},
		{"arn:aws:iam::999999999999:role/not-an-instance-profile", 1},
		{"", 1},
		{"invalid-profile", 1},
	}
	for _, tc := range testCases {
		_, errs := ValidateInstanceProfileARN(tc.instanceProfileARN, "key")

		assert.Lenf(t, errs, tc.errorCount, "directory '%s' does not generate the expected error count", tc.instanceProfileARN)
	}
}
