package qa

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/stretchr/testify/assert"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomLongName ...
func RandomLongName() string {
	return "Terraform Integration Test " + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
}

// RandomName gives random name with optional prefix. e.g. qa.RandomName("tf-")
func RandomName(prefix ...string) string {
	randLen := 12
	b := make([]byte, randLen)
	for i := range b {
		b[i] = charset[rand.Intn(randLen)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", strings.Join(prefix, ""), b)
	}
	return string(b)
}

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
	Fixtures      []HTTPFixture
	Resource      *schema.Resource
	RequiresNew   bool
	InstanceState map[string]string
	State         map[string]interface{}
	// HCL might be useful to test nested blocks
	HCL         string
	CommandMock common.CommandMock
	Create      bool
	Read        bool
	Update      bool
	Delete      bool
	Removed     bool
	ID          string
	NonWritable bool
	Azure       bool
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
	if f.Azure {
		client.AzureAuth.ResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
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
	pick := func(
		a func(*schema.ResourceData, interface{}) error,
		b func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics,
		d *schema.ResourceData, m interface{}) error {
		if b != nil {
			ctx := context.Background()
			diags := b(ctx, d, m)
			if diags != nil {
				return fmt.Errorf(diagsToString(diags))
			}
			return nil
		}
		return a(d, m)
	}
	resourceConfig := terraform.NewResourceConfigRaw(f.State)
	switch {
	case f.Create:
		// nolint should be a bigger context-aware refactor
		whatever = func(d *schema.ResourceData, m interface{}) error {
			return pick(f.Resource.Create, f.Resource.CreateContext, d, m)
		}
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
			return pick(f.Resource.Read, f.Resource.ReadContext, d, m)
		}
	case f.Update:
		if f.ID == "" {
			return nil, errors.New("ID must be set for Update")
		}
		if f.Resource.UpdateContext == nil && f.Resource.Update == nil {
			return nil, errors.New("Resource does not support Update")
		}
		whatever = func(d *schema.ResourceData, m interface{}) error {
			d.SetId(f.ID)
			return pick(f.Resource.Update, f.Resource.UpdateContext, d, m)
		}
	case f.Delete:
		if f.ID == "" {
			return nil, errors.New("ID must be set for Delete")
		}
		whatever = func(d *schema.ResourceData, m interface{}) error {
			d.SetId(f.ID)
			return pick(f.Resource.Delete, f.Resource.DeleteContext, d, m)
		}
	}

	if f.State != nil {
		diags := f.Resource.Validate(resourceConfig)
		if diags.HasError() {
			return nil, fmt.Errorf("Invalid config supplied. %s",
				strings.ReplaceAll(diagsToString(diags), "\"", ""))
		}
	}
	schemaMap := schema.InternalMap(f.Resource.Schema)
	is := &terraform.InstanceState{
		Attributes: f.InstanceState,
	}
	ctx := context.Background()
	diff, err := f.Resource.Diff(ctx, is, resourceConfig, client)
	// TODO: f.Resource.Data(is) - check why it doesn't work
	if err != nil {
		return nil, err
	}
	resourceData, err := schemaMap.Data(is, diff)
	if err != nil {
		return nil, err
	}
	err = f.Resource.InternalValidate(f.Resource.Schema, !f.NonWritable)
	if err != nil {
		return nil, err
	}
	err = whatever(resourceData, client)
	if err != nil {
		return resourceData, err
	}
	if resourceData.Id() == "" && !f.Removed {
		return resourceData, fmt.Errorf("Resource is not expected to be removed")
	}
	newState := resourceData.State()
	diff, err = schemaMap.Diff(ctx, newState, resourceConfig, nil, client, true)
	if err != nil {
		return nil, err
	}
	if diff == nil || f.InstanceState == nil {
		return resourceData, err
	}
	requireNew := []string{}
	for k, v := range diff.Attributes {
		if v.RequiresNew {
			log.Printf("[WARN] %s requires new: %#v %#v", k, v.Old, v.New)
			requireNew = append(requireNew, k)
		}
	}
	if len(requireNew) > 0 && !f.RequiresNew {
		err = fmt.Errorf("Changes from backend require new: %s", strings.Join(requireNew, ", "))
	}
	return resourceData, err
}

// ApplyNoError is a convenience method for no-data tests
func (f ResourceFixture) ApplyNoError(t *testing.T) {
	_, err := f.Apply(t)
	assert.NoError(t, err, err)
}

func diagsToString(diags diag.Diagnostics) string {
	if diags.HasError() {
		sort.Slice(diags, func(i, j int) bool {
			return diags[i].Detail < diags[j].Detail
		})
		issues := []string{}
		for _, diag := range diags {
			attributePath := ""
			if len(diag.AttributePath) > 0 {
				attributePath += "["
				for i, rs := range diag.AttributePath {
					if i > 0 {
						attributePath += "."
					}
					switch step := rs.(type) {
					case cty.GetAttrStep:
						attributePath += step.Name
					default:
						attributePath += "#"
					}
				}
				attributePath += "] "
			}
			if diag.Summary == "ConflictsWith" {
				issues = append(issues, diag.Detail)
			} else {
				issues = append(issues, fmt.Sprintf("%s%s", attributePath, diag.Summary))
			}
		}
		return strings.Join(issues, ". ")
	}
	return ""
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
					if alreadyJSON, ok := fixture.Response.(string); ok {
						_, err = rw.Write([]byte(alreadyJSON))
						assert.NoError(t, err, err)
					} else {
						responseBytes, err := json.Marshal(fixture.Response)
						if err != nil {
							assert.NoError(t, err, err)
							t.FailNow()
						}
						_, err = rw.Write(responseBytes)
						assert.NoError(t, err, err)
					}
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
		return "", errors.New("Cannot have more than one customer variable map")
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
		return "", fmt.Errorf("please set %d variables and restart", missing)
	}
	return internal.TrimLeadingWhitespace(template), nil
}

// EnvironmentTemplate asserts existence and fills in {env.VAR} & {var.RANDOM} placeholders in template
func EnvironmentTemplate(t *testing.T, template string, otherVars ...map[string]string) string {
	resp, err := environmentTemplate(t, template, otherVars...)
	if err != nil {
		t.Skipf(err.Error())
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

// AssertErrorStartsWith ..
func AssertErrorStartsWith(t *testing.T, err error, message string) bool {
	return assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
}

// GetCloudInstanceType gives common minimal instance type, depending on a cloud
func GetCloudInstanceType(c *common.DatabricksClient) string {
	if c.IsAzure() {
		return "Standard_DS3_v2"
	}
	// TODO: create a method on ClustersAPI to give
	// cloud specific delta-cache enabled instance by default.
	return "m4.large"
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

// GetEnvOrSkipTest proceeds with test only with that env variable
func GetEnvOrSkipTest(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		t.Skipf("Environment variable %s is missing", name)
	}
	return value
}
