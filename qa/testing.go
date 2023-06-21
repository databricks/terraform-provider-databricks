package qa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomLongName ...
func RandomLongName() string {
	return "Terraform Integration Test " + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
}

// RandomEmail generates random email
func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", RandomName("tf-"))
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
	Response        any
	Status          int
	ExpectedRequest any
	ReuseRequest    bool
	MatchAny        bool
}

// ResourceFixture helps testing resources and commands
type ResourceFixture struct {
	Fixtures      []HTTPFixture
	Resource      *schema.Resource
	RequiresNew   bool
	InstanceState map[string]string
	State         map[string]any
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
	AzureSPN    bool
	Gcp         bool
	AccountID   string
	Token       string
	// new resource
	New bool
}

// wrapper type for calling resource methords
type resourceCRUD func(context.Context, *schema.ResourceData, any) diag.Diagnostics

func (cb resourceCRUD) before(before func(d *schema.ResourceData)) resourceCRUD {
	return func(ctx context.Context, d *schema.ResourceData, i any) diag.Diagnostics {
		before(d)
		return cb(ctx, d, i)
	}
}

func (cb resourceCRUD) withId(id string) resourceCRUD {
	return cb.before(func(d *schema.ResourceData) {
		d.SetId(id)
	})
}

func (f ResourceFixture) prepareExecution() (resourceCRUD, error) {
	switch {
	case f.Create:
		if f.ID != "" {
			return nil, fmt.Errorf("ID is not available for Create")
		}
		return resourceCRUD(f.Resource.CreateContext).before(func(d *schema.ResourceData) {
			d.MarkNewResource()
		}), nil
	case f.Read:
		if f.ID == "" {
			return nil, fmt.Errorf("ID must be set for Read")
		}
		preRead := f.State
		f.State = nil
		return resourceCRUD(f.Resource.ReadContext).before(func(d *schema.ResourceData) {
			if f.New {
				d.MarkNewResource()
			}
			for k, v := range preRead {
				d.Set(k, v)
			}
		}).withId(f.ID), nil
	case f.Update:
		if f.ID == "" {
			return nil, fmt.Errorf("ID must be set for Update")
		}
		return resourceCRUD(f.Resource.UpdateContext).withId(f.ID), nil
	case f.Delete:
		if f.ID == "" {
			return nil, fmt.Errorf("ID must be set for Delete")
		}
		return resourceCRUD(f.Resource.DeleteContext).withId(f.ID), nil
	}
	return nil, fmt.Errorf("no `Create|Read|Update|Delete: true` specificed")
}

// Apply runs tests from fixture
func (f ResourceFixture) Apply(t *testing.T) (*schema.ResourceData, error) {
	token := "..."
	if f.Token != "" {
		token = f.Token
	}
	client, server, err := HttpFixtureClientWithToken(t, f.Fixtures, token)
	defer server.Close()
	if err != nil {
		return nil, err
	}
	client.Config.WithTesting()
	if f.CommandMock != nil {
		client.WithCommandMock(f.CommandMock)
	}
	if f.Azure {
		client.Config.AzureResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	}
	if f.AzureSPN {
		client.Config.AzureClientID = "a"
		client.Config.AzureClientSecret = "b"
		client.Config.AzureTenantID = "c"
	}
	if f.Gcp {
		client.Config.GoogleServiceAccount = "sa@prj.iam.gserviceaccount.com"
	}
	if f.AccountID != "" {
		client.Config.AccountID = f.AccountID
	}
	if len(f.HCL) > 0 {
		var out any
		// TODO: update to HCLv2 somehow, so that importer and this use the same stuff
		err = hcl.Decode(&out, f.HCL)
		if err != nil {
			return nil, err
		}
		f.State = fixHCL(out).(map[string]any)
	}
	resourceConfig := terraform.NewResourceConfigRaw(f.State)
	execute, err := f.prepareExecution()
	if err != nil {
		return nil, err
	}
	if f.State != nil {
		diags := f.Resource.Validate(resourceConfig)
		if diags.HasError() {
			return nil, fmt.Errorf("invalid config supplied. %s",
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
	if f.Update {
		err = f.requiresNew(diff)
		if err != nil {
			return nil, err
		}
	}
	resourceData, err := schemaMap.Data(is, diff)
	if err != nil {
		return nil, err
	}
	err = f.Resource.InternalValidate(f.Resource.Schema, !f.NonWritable)
	if err != nil {
		return nil, err
	}
	if execute != nil {
		// this is a bit strange, but we'll fix it later
		diags := execute(ctx, resourceData, client)
		if diags != nil {
			return resourceData, fmt.Errorf(diagsToString(diags))
		}
	}
	if resourceData.Id() == "" && !f.Removed {
		return resourceData, fmt.Errorf("resource is not expected to be removed")
	}
	newState := resourceData.State()
	diff, err = schemaMap.Diff(ctx, newState, resourceConfig, f.Resource.CustomizeDiff, client, true)
	if err != nil {
		return nil, err
	}
	if diff == nil || f.InstanceState == nil {
		return resourceData, nil
	}
	err = f.requiresNew(diff)
	return resourceData, err
}

func (f ResourceFixture) requiresNew(diff *terraform.InstanceDiff) error {
	requireNew := []string{}
	for k, v := range diff.Attributes {
		if v.Old == v.New {
			continue
		}
		if v.RequiresNew {
			log.Printf("[WARN] %s requires new: %#v -> %#v", k, v.Old, v.New)
			requireNew = append(requireNew, k)
		}
	}
	if len(requireNew) > 0 && !f.RequiresNew {
		return fmt.Errorf("changes require new: %s", strings.Join(requireNew, ", "))
	}
	return nil
}

// ApplyNoError is a convenience method for no-data tests
func (f ResourceFixture) ApplyNoError(t *testing.T) {
	_, err := f.Apply(t)
	assert.NoError(t, err)
}

// ApplyAndExpectData is a convenience method for tests that doesn't expect error, but want to check data
func (f ResourceFixture) ApplyAndExpectData(t *testing.T, data map[string]any) {
	d, err := f.Apply(t)
	require.NoError(t, err)
	for k, expected := range data {
		if k == "id" {
			assert.Equal(t, expected, d.Id())
		} else if that, ok := d.Get(k).(*schema.Set); ok {
			this := expected.([]string)
			assert.Equal(t, len(this), that.Len(), "set has different length")
			for _, item := range this {
				assert.True(t, that.Contains(item), "set does not contain %s", item)
			}
		} else {
			assert.Equal(t, expected, d.Get(k))
		}
	}
}

// ExpectError passes if there's an error
func (f ResourceFixture) ExpectError(t *testing.T, msg string) {
	_, err := f.Apply(t)
	assert.EqualError(t, err, msg)
}

type CornerCase struct {
	part  string
	value string
}

func CornerCaseID(id string) CornerCase {
	return CornerCase{"id", id}
}

func CornerCaseExpectError(msg string) CornerCase {
	return CornerCase{"expect_error", msg}
}

func CornerCaseSkipCRUD(method string) CornerCase {
	return CornerCase{"skip_crud", method}
}

func CornerCaseAccountID(id string) CornerCase {
	return CornerCase{"account_id", id}
}

var HTTPFailures = []HTTPFixture{
	{
		MatchAny:     true,
		ReuseRequest: true,
		Status:       418,
		Response: apierr.APIError{
			ErrorCode:  "NONSENSE",
			StatusCode: 418,
			Message:    "I'm a teapot",
		},
	},
}

// ResourceCornerCases checks for corner cases of error handling. Optional field name used to create error
func ResourceCornerCases(t *testing.T, resource *schema.Resource, cc ...CornerCase) {
	config := map[string]string{
		"id":           "x",
		"expect_error": "I'm a teapot",
		"account_id":   "",
	}
	m := map[string]func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics{
		"create": resource.CreateContext,
		"read":   resource.ReadContext,
		"update": resource.UpdateContext,
		"delete": resource.DeleteContext,
	}
	for _, corner := range cc {
		if corner.part == "skip_crud" {
			delete(m, corner.value)
		}
		config[corner.part] = corner.value
	}
	HTTPFixturesApply(t, HTTPFailures, func(ctx context.Context, client *common.DatabricksClient) {
		validData := resource.TestResourceData()
		client.Config.AccountID = config["account_id"]
		for n, v := range m {
			if v == nil {
				continue
			}
			validData.SetId(config["id"])
			diags := v(ctx, validData, client)
			if assert.Len(t, diags, 1) {
				assert.Equalf(t, config["expect_error"], diags[0].Summary,
					"%s didn't handle correct error on valid data", n)
			}
		}
	})
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
	return HttpFixtureClientWithToken(t, fixtures, "...")
}

// HttpFixtureClientWithToken creates client for emulated HTTP server
func HttpFixtureClientWithToken(t *testing.T, fixtures []HTTPFixture, token string) (*common.DatabricksClient, *httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		found := false
		for i, fixture := range fixtures {
			if (req.Method == fixture.Method && req.RequestURI == fixture.Resource) || fixture.MatchAny {
				if fixture.Status == 0 {
					rw.WriteHeader(200)
				} else {
					rw.WriteHeader(fixture.Status)
				}
				if fixture.ExpectedRequest != nil {
					buf := new(bytes.Buffer)
					_, err := buf.ReadFrom(req.Body)
					assert.NoError(t, err)
					jsonStr, err := json.Marshal(fixture.ExpectedRequest)
					assert.NoError(t, err)
					assert.JSONEq(t, string(jsonStr), buf.String(), "json strings do not match")
				}
				if fixture.Response != nil {
					if alreadyJSON, ok := fixture.Response.(string); ok {
						_, err := rw.Write([]byte(alreadyJSON))
						assert.NoError(t, err)
					} else {
						responseBytes, err := json.Marshal(fixture.Response)
						if err != nil {
							assert.NoError(t, err)
							t.FailNow()
						}
						_, err = rw.Write(responseBytes)
						assert.NoError(t, err)
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
			receivedRequest := map[string]any{}
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(req.Body)
			assert.NoError(t, err)
			err = json.Unmarshal(buf.Bytes(), &receivedRequest)
			assert.NoError(t, err)

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
	cfg := &config.Config{
		Host:             server.URL,
		Token:            token,
		AzureEnvironment: "PUBLIC",
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, nil, err
	}
	return &common.DatabricksClient{
		DatabricksClient: c,
	}, server, nil
}

// HTTPFixturesApply is a helper method
func HTTPFixturesApply(t *testing.T, fixtures []HTTPFixture, callback func(ctx context.Context, client *common.DatabricksClient)) {
	client, server, err := HttpFixtureClient(t, fixtures)
	defer server.Close()
	require.NoError(t, err)
	callback(context.Background(), client)
}

func fixHCL(v any) any {
	switch a := v.(type) {
	case []map[string]any:
		vals := []any{}
		for _, vv := range a {
			vals = append(vals, fixHCL(vv))
		}
		return vals
	case map[string]any:
		vals := map[string]any{}
		for k, ev := range a {
			vals[k] = fixHCL(ev)
		}
		return vals
	default:
		return v
	}
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
	return err != nil && assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
}

// GetEnvOrSkipTest proceeds with test only with that env variable
func GetEnvOrSkipTest(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		t.Skipf("Environment variable %s is missing", name)
	}
	return value
}

func RequireAnyCloudEnv(t *testing.T) {
	value := os.Getenv("CLOUD_ENV")
	if value == "" {
		t.Skip("CLOUD_ENV is required to run this test")
	}
}

func RequireCloudEnv(t *testing.T, cloudEnv string) {
	value := os.Getenv("CLOUD_ENV")
	if value != cloudEnv {
		t.Skipf("CLOUD_ENV=%s is required to run this test", cloudEnv)
	}
}
