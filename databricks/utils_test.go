package databricks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/stretchr/testify/assert"
)

func TestMissingMWSResources(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	mwsAcctId := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	randStringId := acctest.RandString(10)
	randIntId := 2000000 + acctest.RandIntRange(100000, 20000000)

	client := getMWSClient()
	tests := []struct {
		name            string
		readFunc        func() error
		isCustomCheck   bool
		resourceID      string
		customCheckFunc func(err error, rId string) bool
	}{
		{
			name: "CheckIfMWSCredentialsAreMissing",
			readFunc: func() error {
				_, err := client.MWSCredentials().Read(mwsAcctId, randStringId)
				return err
			},
		},
		{
			name: "CheckIfMWSNetworksAreMissing",
			readFunc: func() error {
				_, err := client.MWSNetworks().Read(mwsAcctId, randStringId)
				return err
			},
		},
		{
			name: "CheckIfMWSStorageConfigurationsAreMissing",
			readFunc: func() error {
				_, err := client.MWSStorageConfigurations().Read(mwsAcctId, randStringId)
				return err
			},
		},
		{
			name: "CheckIfMWSWorkspacesAreMissing",
			readFunc: func() error {
				_, err := client.MWSWorkspaces().Read(mwsAcctId, int64(randIntId))
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCustomCheck {
				// Test custom check because api call does not return 404 not found if the resource does not exist
				testVerifyResourceIsMissingCustomVerification(t, tt.resourceID, tt.readFunc, tt.customCheckFunc)
			} else {
				testVerifyResourceIsMissing(t, tt.readFunc)
			}
		})
	}
}

// Capture this test for aws
func TestAccAwsMissingWorkspaceResources(t *testing.T) {
	testMissingWorkspaceResources(t, service.AWS)
}

// Capture this test for azure
func TestAccAzureMissingWorkspaceResources(t *testing.T) {
	testMissingWorkspaceResources(t, service.Azure)
}

func testMissingWorkspaceResources(t *testing.T, cloud service.CloudServiceProvider) {
	if _, ok := os.LookupEnv("TF_ACC"); !ok {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	randIntId := 2000000 + acctest.RandIntRange(100000, 20000000)
	randStringId := acctest.RandString(10)
	// example 405E7E8E4A000024
	randomClusterPolicyId := fmt.Sprintf("400E9E9E9A%d",
		acctest.RandIntRange(100000, 999999),
	)
	// example 0101-120000-brick1-pool-ABCD1234
	randomInstancePoolId := fmt.Sprintf(
		"%v-%v-%s-pool-%s",
		acctest.RandIntRange(1000, 9999),
		acctest.RandIntRange(100000, 999999),
		acctest.RandString(6),
		acctest.RandString(8),
	)
	client := getIntegrationDBAPIClient(t)

	type testTable struct {
		name            string
		readFunc        func() error
		isCustomCheck   bool
		resourceID      string
		customCheckFunc func(err error, rId string) bool
	}
	tests := []testTable{
		{
			name: "CheckIfTokensAreMissing",
			readFunc: func() error {
				_, err := client.Tokens().Read(randStringId)
				return err
			},
		},
		{
			name: "CheckIfSecretScopesAreMissing",
			readFunc: func() error {
				_, err := client.SecretScopes().Read(randStringId)
				return err
			},
		},
		{
			name: "CheckIfSecretsAreMissing",
			readFunc: func() error {
				_, err := client.Secrets().Read(randStringId, randStringId)
				return err
			},
		},
		{
			name: "CheckIfSecretsACLsAreMissing",
			readFunc: func() error {
				_, err := client.SecretAcls().Read(randStringId, randStringId)
				return err
			},
		},
		{
			name: "CheckIfSecretsACLsAreMissing",
			readFunc: func() error {
				_, err := client.SecretAcls().Read(randStringId, randStringId)
				return err
			},
		},
		{
			name: "CheckIfNotebooksAreMissing",
			readFunc: func() error {
				// ID must start with a /
				_, err := client.Notebooks().Read("/" + randStringId)
				return err
			},
		},
		{
			name: "CheckIfInstancePoolsAreMissing",
			readFunc: func() error {
				_, err := client.InstancePools().Read(randomInstancePoolId)
				return err
			},
		},
		{
			name: "CheckIfClustersAreMissing",
			readFunc: func() error {
				_, err := client.Clusters().Get(randStringId)
				return err
			},
			isCustomCheck:   true,
			customCheckFunc: isClusterMissing,
			resourceID:      randStringId,
		},
		{
			name: "CheckIfDBFSFilesAreMissing",
			readFunc: func() error {
				_, err := client.DBFS().Read("/" + randStringId)
				return err
			},
		},
		{
			name: "CheckIfGroupsAreMissing",
			readFunc: func() error {
				_, err := client.Groups().Read(randStringId)
				t.Log(err)
				return err
			},
		},
		{
			name: "CheckIfUsersAreMissing",
			readFunc: func() error {
				_, err := client.Users().Read(randStringId)
				t.Log(err)
				return err
			},
		},
		{
			name: "CheckIfClusterPoliciesAreMissing",
			readFunc: func() error {
				_, err := client.ClusterPolicies().Get(randomClusterPolicyId)
				t.Log(err)
				return err
			},
		},
		{
			name: "CheckIfJobsAreMissing",
			readFunc: func() error {
				_, err := client.Jobs().Read(int64(randIntId))
				return err
			},
			isCustomCheck:   true,
			customCheckFunc: isJobMissing,
			resourceID:      strconv.Itoa(randIntId),
		},
	}
	// Handle aws only tests where instance profiles only exist on aws
	if cloud == service.AWS {
		awsOnlyTests := []testTable{
			{
				name: "CheckIfInstanceProfilesAreMissing",
				readFunc: func() error {
					_, err := client.InstanceProfiles().Read(randStringId)
					return err
				},
			},
		}
		tests = append(tests, awsOnlyTests...)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCustomCheck {
				// Test custom check because api call does not return 404 not found if the resource does not exist
				testVerifyResourceIsMissingCustomVerification(t, tt.resourceID, tt.readFunc, tt.customCheckFunc)
			} else {
				testVerifyResourceIsMissing(t, tt.readFunc)
			}
		})
	}
}

func testVerifyResourceIsMissingCustomVerification(t *testing.T, resourceId string, readFunc func() error,
	customCheck func(err error, rId string) bool) {
	err := readFunc()
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, err, service.APIError{}, fmt.Sprintf("error: %s is not type api error", err.Error()))
	if apiError, ok := err.(service.APIError); ok {
		assert.True(t, customCheck(err, resourceId), fmt.Sprintf("error: %v is not missing;"+
			"\nstatus code: %v;"+
			"\nerror code: %s",
			apiError, apiError.StatusCode, apiError.ErrorCode))
	}
}

func testVerifyResourceIsMissing(t *testing.T, readFunc func() error) {
	err := readFunc()
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, err, service.APIError{}, fmt.Sprintf("error: %s is not type api error", err.Error()))
	if apiError, ok := err.(service.APIError); ok {
		assert.True(t, apiError.IsMissing(), fmt.Sprintf("error: %v is not missing;"+
			"\nstatus code: %v;"+
			"\nerror code: %s",
			apiError, apiError.StatusCode, apiError.ErrorCode))
	}
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
	err := errors.New("{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Cluster 123 does not exist\"}")

	result := isClusterMissing(err, "123")

	assert.True(t, result)
}

func TestIsClusterMissingFalseWhenClusterIdSpecifiedNotPresent(t *testing.T) {
	err := errors.New("{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Cluster 123 does not exist\"}")

	result := isClusterMissing(err, "xyz")

	assert.False(t, result)
}

func TestIsClusterMissingFalseWhenErrorNotInCorrectFormat(t *testing.T) {
	err := errors.New("{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Something random went bang xyz\"}")

	result := isClusterMissing(err, "xyz")

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
