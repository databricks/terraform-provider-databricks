package databricks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/client/service"
)

// EnvironmentTemplate asserts existance and fills in {env.VAR} & {var.RANDOM} placeholders in template
func EnvironmentTemplate(t *testing.T, template string) string {
	vars := map[string]string{
		"RANDOM": acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum),
	}
	missing := 0
	var varType, varName, value string
	r := regexp.MustCompile(`{(env|var).([^{}]*)}`)
	for _, variableMatch := range r.FindAllStringSubmatch(template, -1) {
		value = ""
		varType = variableMatch[1]
		varName = variableMatch[2]
		switch varType {
		case "env":
			value = os.Getenv(varName)
		case "var":
			value = vars[varName]
		}
		if value == "" {
			t.Logf("Missing %s %s variable.", varType, varName)
			missing++
			continue
		}
		template = strings.ReplaceAll(template, `{`+varType+`.`+varName+`}`, value)
	}
	if missing > 0 {
		t.Fatalf("Please set %d variables and restart.", missing)
	}
	return service.TrimLeadingWhitespace(template)
}

func FirstKeyValue(t *testing.T, str, key string) string {
	r := regexp.MustCompile(key + `\s+=\s+"([^"]*)"`)
	match := r.FindStringSubmatch(str)
	if len(match) != 2 {
		t.Fatalf("Cannot find %s in given string", key)
	}
	return match[1]
}

func TestEnvironmentTemplate(t *testing.T) {
	res := EnvironmentTemplate(t, `
	resource "user" "me" {
		name  = "{env.USER}"
		email = "{env.USER}+{var.RANDOM}@example.com"
	}`)
	assert.True(t, strings.Contains(res, fmt.Sprintf(`name = "%s"`, os.Getenv("USER"))), res)
	assert.Equal(t, os.Getenv("USER"), FirstKeyValue(t, res, "name"))
}

func assertErrorStartsWith(t *testing.T, err error, message string) bool {
	return assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
}

type MissingResourceCheck struct {
	name          string
	readFunc      func() error
	isCustomCheck bool
	resourceID    string
}

func TestMwsAccMissingResources(t *testing.T) {
	if cloudEnv, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloudEnv != "MWS" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV=MWS' is set.")
	}

	mwsAcctID := os.Getenv("DATABRICKS_ACCOUNT_ID")
	if mwsAcctID == "" {
		t.Skip("Must have DATABRICKS_ACCOUNT_ID environment variable set.")
	}
	randStringID := acctest.RandString(10)
	randIntID := 2000000 + acctest.RandIntRange(100000, 20000000)

	client := service.CommonEnvironmentClient()
	tests := []MissingResourceCheck{
		{
			name: "Credential",
			readFunc: func() error {
				_, err := client.MWSCredentials().Read(mwsAcctID, randStringID)
				return err
			},
		},
		{
			name: "Network",
			readFunc: func() error {
				_, err := client.MWSNetworks().Read(mwsAcctID, randStringID)
				return err
			},
		},
		{
			name: "Storage",
			readFunc: func() error {
				_, err := client.MWSStorageConfigurations().Read(mwsAcctID, randStringID)
				return err
			},
		},
		{
			name: "Workspace",
			readFunc: func() error {
				_, err := client.MWSWorkspaces().Read(mwsAcctID, int64(randIntID))
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testVerifyResourceIsMissing(t, tt.readFunc)
		})
	}
}

// Capture this test for aws
func TestAccMissingResourcesInWorkspace(t *testing.T) {
	cloudENV, ok := os.LookupEnv("CLOUD_ENV")
	if !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' set")
	}
	randIntID := 2000000 + acctest.RandIntRange(100000, 20000000)
	randStringID := acctest.RandString(10)
	// example 405E7E8E4A000024
	randomClusterPolicyID := fmt.Sprintf("400E9E9E9A%d",
		acctest.RandIntRange(100000, 999999),
	)
	// example 0101-120000-brick1-pool-ABCD1234
	randomInstancePoolID := fmt.Sprintf(
		"%v-%v-%s-pool-%s",
		acctest.RandIntRange(1000, 9999),
		acctest.RandIntRange(100000, 999999),
		acctest.RandString(6),
		acctest.RandString(8),
	)
	client := service.NewClientFromEnvironment()

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
				_, err := client.Tokens().Read(randStringID)
				return err
			},
		},
		{
			name: "CheckIfSecretScopesAreMissing",
			readFunc: func() error {
				_, err := client.SecretScopes().Read(randStringID)
				return err
			},
		},
		{
			name: "CheckIfSecretsAreMissing",
			readFunc: func() error {
				_, err := client.Secrets().Read(randStringID, randStringID)
				return err
			},
		},
		{
			name: "CheckIfSecretsACLsAreMissing",
			readFunc: func() error {
				_, err := client.SecretAcls().Read(randStringID, randStringID)
				return err
			},
		},
		{
			name: "CheckIfSecretsACLsAreMissing",
			readFunc: func() error {
				_, err := client.SecretAcls().Read(randStringID, randStringID)
				return err
			},
		},
		{
			name: "CheckIfNotebooksAreMissing",
			readFunc: func() error {
				// ID must start with a /
				_, err := client.Notebooks().Read("/" + randStringID)
				return err
			},
		},
		{
			name: "CheckIfInstancePoolsAreMissing",
			readFunc: func() error {
				_, err := client.InstancePools().Read(randomInstancePoolID)
				return err
			},
		},
		{
			name: "CheckIfClustersAreMissing",
			readFunc: func() error {
				_, err := client.Clusters().Get(randStringID)
				return err
			},
			resourceID: randStringID,
		},
		{
			name: "CheckIfDBFSFilesAreMissing",
			readFunc: func() error {
				_, err := client.DBFS().Read("/" + randStringID)
				return err
			},
		},
		{
			name: "CheckIfGroupsAreMissing",
			readFunc: func() error {
				_, err := client.Groups().Read(randStringID)
				t.Log(err)
				return err
			},
		},
		{
			name: "CheckIfUsersAreMissing",
			readFunc: func() error {
				_, err := client.Users().Read(randStringID)
				t.Log(err)
				return err
			},
		},
		{
			name: "CheckIfClusterPoliciesAreMissing",
			readFunc: func() error {
				_, err := client.ClusterPolicies().Get(randomClusterPolicyID)
				t.Log(err)
				return err
			},
		},
		{
			name: "CheckIfJobsAreMissing",
			readFunc: func() error {
				_, err := client.Jobs().Read(strconv.Itoa(randIntID))
				return err
			},
			resourceID: strconv.Itoa(randIntID),
		},
	}
	if cloudENV == "AWS" {
		// Handle aws only tests where instance profiles only exist on aws
		awsOnlyTests := []testTable{
			{
				name: "CheckIfInstanceProfilesAreMissing",
				readFunc: func() error {
					_, err := client.InstanceProfiles().Read(randStringID)
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

func testVerifyResourceIsMissingCustomVerification(t *testing.T, resourceID string, readFunc func() error,
	customCheck func(err error, rId string) bool) {
	err := readFunc()
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, err, service.APIError{}, fmt.Sprintf("error: %s is not type api error", err.Error()))
	if apiError, ok := err.(service.APIError); ok {
		assert.True(t, customCheck(err, resourceID), fmt.Sprintf("error: %v is not missing;"+
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
}

func TestGetParentDirPath(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		expectedDirPath string
		expectedError   error
	}{
		{
			name:            "basic_path",
			path:            "/test/abc/file.py",
			expectedDirPath: "/test/abc",
			expectedError:   nil,
		},
		{
			name:            "root_path",
			path:            "/file.py",
			expectedDirPath: "",
			expectedError:   DirPathRootDirError,
		},
		{
			name:            "empty_path",
			path:            "",
			expectedDirPath: "",
			expectedError:   PathEmptyError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dirPath, err := GetParentDirPath(tt.path)
			assert.Equal(t, tt.expectedDirPath, dirPath, "dirPath values should match")
			assert.Equal(t, tt.expectedError, err, "err values should match")
		})
	}
}

type errorSlice []error

func (a errorSlice) Len() int           { return len(a) }
func (a errorSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a errorSlice) Less(i, j int) bool { return a[i].Error() < a[j].Error() }

// HTTPFixture defines request structure for test
type HTTPFixture struct {
	Method          string
	Resource        string
	Response        interface{}
	Status          int
	ExpectedRequest interface{}
	ReuseRequest    bool
}

type ResourceFixture struct {
	Fixtures []HTTPFixture
	Resource *schema.Resource
	State    map[string]interface{}
	// HCL might be useful to test nested blocks
	HCL         string
	CommandMock service.CommandMock
	Create      bool
	Read        bool
	Update      bool
	Delete      bool
	ID          string
	// new resource
	New bool
}

func (f ResourceFixture) Apply(t *testing.T) (*schema.ResourceData, error) {
	client, server, err := httpFixtureClient(t, f.Fixtures)
	defer server.Close()
	if err != nil {
		return nil, err
	}
	if f.CommandMock != nil {
		client.WithCommandMock(f.CommandMock)
	}
	if len(f.HCL) > 0 {
		var out interface{}
		err = hcl.Decode(&out, f.HCL)
		if err != nil {
			return nil, err
		}
		f.State = fixHCL(out).(map[string]interface{})
	}
	var whatever func(d *schema.ResourceData, c interface{}) error
	switch {
	case f.Create:
		whatever = f.Resource.Create
		if f.ID != "" {
			return nil, errors.New("ID is not available for Create")
		}
	case f.Read:
		if f.ID == "" {
			return nil, errors.New("ID must be set for Read")
		}
		whatever = func(d *schema.ResourceData, m interface{}) error {
			d.SetId(f.ID)
			if f.New {
				d.MarkNewResource()
			}
			return f.Resource.Read(d, m)
		}
	case f.Update:
		if f.ID == "" {
			return nil, errors.New("ID must be set for Update")
		}
		whatever = func(d *schema.ResourceData, m interface{}) error {
			d.SetId(f.ID)
			return f.Resource.Update(d, m)
		}
	case f.Delete:
		if f.ID == "" {
			return nil, errors.New("ID must be set for Delete")
		}
		whatever = func(d *schema.ResourceData, m interface{}) error {
			d.SetId(f.ID)
			return f.Resource.Delete(d, m)
		}
	}
	return resourceTesterScaffolding(t, f.Resource, f.State, client, whatever)
}

func UnionFixturesLists(fixturesLists ...[]HTTPFixture) (fixtureList []HTTPFixture) {
	for _, v := range fixturesLists {
		fixtureList = append(fixtureList, v...)
	}
	return
}

// ResourceTester helps testing HTTP resources with fixtures
func ResourceTester(t *testing.T,
	fixtures []HTTPFixture,
	resourceFunc func() *schema.Resource,
	state map[string]interface{},
	whatever func(d *schema.ResourceData, c interface{}) error) (*schema.ResourceData, error) {
	client, server, err := httpFixtureClient(t, fixtures)
	defer server.Close()
	if err != nil {
		return nil, err
	}
	return resourceTesterScaffolding(t, resourceFunc(), state, client, whatever)
}

func httpFixtureClient(t *testing.T, fixtures []HTTPFixture) (client *service.DatabricksClient, server *httptest.Server, err error) {
	server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		found := false
		for i, fixture := range fixtures {
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
					assert.JSONEq(t, string(jsonStr), buf.String(), "json strings do not match")
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
				// Reset the request if it is already used
				if !fixture.ReuseRequest {
					fixtures[i] = HTTPFixture{}
				}
				break
			}
		}
		if !found {
			receivedRequest := map[string]interface{}{}
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(req.Body)
			assert.NoError(t, err, err)
			err = json.Unmarshal(buf.Bytes(), &receivedRequest)
			assert.NoError(t, err, err)

			expectedRequest := ""
			if len(receivedRequest) > 0 {
				// guessing model name would require going over AST,
				// which is not something i'm willing to write on my weekend
				expectedRequest += "ExpectedRequest: model.XXX {\n"
				for key, value := range receivedRequest {
					camel := ""
					for _, part := range strings.Split(key, "_") {
						if len(key) < 4 {
							// golang styles, meh...
							camel += strings.ToUpper(key)
						} else {
							camel += strings.Title(part)
						}
					}
					// best effort prediction of what struct should look like...
					expectedRequest += fmt.Sprintf("					%s: %#v,\n", camel, value)
				}
				expectedRequest += "				},\n"
				expectedRequest += fmt.Sprintf("				// ExpectedRequest: %#v,\n", receivedRequest)
			}
			stub := fmt.Sprintf(`{
				Method:   "%s",
				Resource: "%s",
				%s
				Response: model.XXX {
					// fill in specific fields...
				},
			},`, req.Method, req.RequestURI, expectedRequest)
			assert.Fail(t, fmt.Sprintf("Missing stub, please add: %s", stub))
			t.FailNow()
		}
	}))
	client = &service.DatabricksClient{
		Host:  server.URL,
		Token: "...",
	}
	err = client.Configure("dev")
	return client, server, err
}

func fixHCL(v interface{}) interface{} {
	switch a := v.(type) {
	case []map[string]interface{}:
		vals := []interface{}{}
		for _, vv := range a {
			vals = append(vals, fixHCL(vv))
		}
		return vals
	case map[string]interface{}:
		vals := map[string]interface{}{}
		for k, ev := range a {
			vals[k] = fixHCL(ev)
		}
		return vals
	default:
		return v
	}
}

func resourceTesterScaffolding(t *testing.T, res *schema.Resource,
	state map[string]interface{}, client *service.DatabricksClient,
	whatever func(d *schema.ResourceData, c interface{}) error) (*schema.ResourceData, error) {
	if res == nil {
		return nil, errors.New("Resource is not set")
	}
	if state != nil {
		resourceConfig := terraform.NewResourceConfigRaw(state)
		warns, errs := res.Validate(resourceConfig)
		if len(warns) > 0 || len(errs) > 0 {
			var issues string
			if len(warns) > 0 {
				sort.Strings(warns)
				issues += ". " + strings.Join(warns, ". ")
			}
			if len(errs) > 0 {
				sort.Sort(errorSlice(errs))
				for _, err := range errs {
					issues += ". " + err.Error()
				}
			}
			// remove characters that need escaping, it's only tests...
			issues = strings.ReplaceAll(issues, "\"", "")
			return nil, fmt.Errorf("Invalid config supplied%s", issues)
		}
	}
	resourceData := schema.TestResourceDataRaw(t, res.Schema, state)
	err := res.InternalValidate(res.Schema, true)
	if err != nil {
		return nil, err
	}

	// warns, errs := schemaMap(r.Schema).Validate(c)
	return resourceData, whatever(resourceData, client)
}

func actionWithID(id string, w schema.CreateFunc) schema.CreateFunc {
	return func(d *schema.ResourceData, c interface{}) error {
		d.SetId(id)
		return w(d, c)
	}
}

func debugIfCloudEnvSet() bool {
	return os.Getenv("CLOUD_ENV") != ""
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
