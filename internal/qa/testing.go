package qa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/r3labs/diff"
	"github.com/stretchr/testify/assert"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomLongName ...
func RandomLongName() string {
	return "Terraform Integration Test " + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
}

// RandomName is what it is
func RandomName() string {
	randLen := 12
	b := make([]byte, randLen)
	for i := range b {
		b[i] = charset[rand.Intn(randLen)]
	}
	return string(b)
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

// ResourceFixture helps testing resources and commands
type ResourceFixture struct {
	Fixtures []HTTPFixture
	Resource *schema.Resource
	State    map[string]interface{}
	// HCL might be useful to test nested blocks
	HCL         string
	CommandMock common.CommandMock
	Create      bool
	Read        bool
	Update      bool
	Delete      bool
	ID          string
	NonWritable bool
	// new resource
	New bool
}

// Apply runs tests from fixture
func (f ResourceFixture) Apply(t *testing.T) (*schema.ResourceData, error) {
	client, server, err := HttpFixtureClient(t, f.Fixtures)
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
		preRead := f.State
		f.State = nil
		whatever = func(d *schema.ResourceData, m interface{}) error {
			d.SetId(f.ID)
			if f.New {
				d.MarkNewResource()
			}
			for k, v := range preRead {
				err = d.Set(k, v)
				assert.NoError(t, err)
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
	if f.State != nil {
		resourceConfig := terraform.NewResourceConfigRaw(f.State)
		warns, errs := f.Resource.Validate(resourceConfig)
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
	resourceData := schema.TestResourceDataRaw(t, f.Resource.Schema, f.State)
	err = f.Resource.InternalValidate(f.Resource.Schema, !f.NonWritable)
	if err != nil {
		return nil, err
	}

	// warns, errs := schemaMap(r.Schema).Validate(c)
	return resourceData, whatever(resourceData, client)
}

// UnionFixturesLists merges two HTTP fixture lists together
func UnionFixturesLists(fixturesLists ...[]HTTPFixture) (fixtureList []HTTPFixture) {
	for _, v := range fixturesLists {
		fixtureList = append(fixtureList, v...)
	}
	return
}

// HttpFixtureClient creates client for emulated HTTP server
func HttpFixtureClient(t *testing.T, fixtures []HTTPFixture) (client *common.DatabricksClient, server *httptest.Server, err error) {
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
				expectedRequest += "ExpectedRequest: XXX {\n"
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
				Response: XXX {
					// fill in specific fields...
				},
			},`, req.Method, req.RequestURI, expectedRequest)
			assert.Fail(t, fmt.Sprintf("Missing stub, please add: %s", stub))
			t.FailNow()
		}
	}))
	client = &common.DatabricksClient{
		Host:  server.URL,
		Token: "...",
	}
	err = client.Configure()
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

// For writing a unit test to intercept the errors (t.Fatalf literally ends the test in failure)
func environmentTemplate(t *testing.T, template string, otherVars ...map[string]string) (string, error) {
	vars := map[string]string{
		"RANDOM": acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum),
	}
	if len(otherVars) > 1 {
		return "", errors.New("Cannot have more than one customer variable map!")
	}
	if len(otherVars) == 1 {
		for k, v := range otherVars[0] {
			vars[k] = v
		}
	}
	// pullAll otherVars
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
		return "", fmt.Errorf("please set %d variables and restart.", missing)
	}
	return internal.TrimLeadingWhitespace(template), nil
}

// EnvironmentTemplate asserts existence and fills in {env.VAR} & {var.RANDOM} placeholders in template
func EnvironmentTemplate(t *testing.T, template string, otherVars ...map[string]string) string {
	resp, err := environmentTemplate(t, template, otherVars...)
	if err != nil {
		t.Fatal(err.Error())
	}
	return resp
}

// FirstKeyValue gets it from HCL string
func FirstKeyValue(t *testing.T, str, key string) string {
	r := regexp.MustCompile(key + `\s+=\s+"([^"]*)"`)
	match := r.FindStringSubmatch(str)
	if len(match) != 2 {
		t.Fatalf("Cannot find %s in given string", key)
	}
	return match[1]
}

func AssertErrorStartsWith(t *testing.T, err error, message string) bool {
	return assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
}

// func TestMain(m *testing.M) {
// 	//TODO: is this needed at all?...
// 	err := godotenv.Load("../../.env") // TODO: make teardowns work here as well
// 	// TODO: add common instance pool & cluster for libs & stuff
// 	log.SetFlags(log.Lshortfile | log.Ltime)
// 	if err != nil {
// 		log.Println("Failed to load environment")
// 	}
// 	code := m.Run()
// 	os.Exit(code)
// }

func DeserializeJSON(req *http.Request, m interface{}) error {
	// TODO: remove it!!!
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&m)
	return err
}

func compare(t *testing.T, a interface{}, b interface{}) {
	// TODO: remove diff package because of license
	difference, err := diff.Diff(a, b)
	assert.NoError(t, err, err)
	jsonStr, err := json.Marshal(difference)
	assert.NoError(t, err, err)
	assert.True(t, reflect.DeepEqual(a, b), string(jsonStr))
}

// GetCloudInstanceType gives common minimal instance type, depending on a cloud
func GetCloudInstanceType(c *common.DatabricksClient) string {
	if c.IsUsingAzureAuth() {
		return "Standard_DS3_v2"
	}
	// TODO: create a method on ClustersAPI to give
	// cloud specific delta-cache enabled instance by default.
	return "m4.large"
}

func AssertRequestWithMockServer(t *testing.T, rawPayloadArgs interface{}, requestMethod string, requestURI string, input interface{}, response string, responseStatus int, want interface{}, wantErr bool, apiCall func(client common.DatabricksClient) (interface{}, error)) {
	t.Log("[DEPRECATED] Please rewrite the code to use ResourceFixture")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, requestMethod, req.Method, "HTTP method doesn't match")
		assert.Equal(t, requestURI, req.RequestURI, "URL doesn't match")
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
	client := common.DatabricksClient{
		Host:  server.URL,
		Token: "...",
	}
	err := client.Configure()
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

func AssertMultipleRequestsWithMockServer(t *testing.T, rawPayloadArgs interface{}, requestMethod []string, requestURI []string, input interface{}, response []string, responseStatus []int, want interface{}, wantErr bool, apiCall func(client common.DatabricksClient) (interface{}, error)) {
	t.Log("[DEPRECATED] Please rewrite the code to use ResourceFixture")
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
	client := common.DatabricksClient{
		Host:  server.URL,
		Token: "...",
	}
	err := client.Configure()
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

// TestCreateTempFile  ...
func TestCreateTempFile(t *testing.T, data string) string {
	tmpFile, err := ioutil.TempFile("", "tf-test-create-dbfs-file")
	if err != nil {
		t.Fatal(err)
	}
	filename := tmpFile.Name()

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		os.Remove(filename)
		t.Fatal(err)
	}

	return filename
}
