package tests

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"
	"testing"
	"text/template"

	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

type resourceTestStub struct {
	Name     string
	Resource *schema.Resource
	others   *[]string
}

func (stub *resourceTestStub) stoobyDo(t *testing.T, suffix, tpl string) bool {
	testName := "TestResource" + stub.Name + suffix
	for _, test := range *stub.others {
		if test == testName {
			return true
		}
	}
	tp := template.Must(template.New(suffix).Parse(tpl))
	err := tp.Execute(os.Stdout, stub)
	assert.NoError(t, err)
	t.Logf("Please add missing test %s, ", testName)
	t.Fail()
	return false
}

func (stub *resourceTestStub) Reads(t *testing.T) {
	stub.stoobyDo(t, "Read", `
	func TestResource{{.Name}}Read(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				// read log output of test util for further stubs...
			},
			Resource: Resource{{.Name}}(),
			Read: true,
			ID: "abc",
		}.Apply(t)
		assert.NoError(t, err)
		assert.Equal(t, "abc", d.Id(), "Id should not be empty")
		{{range $index, $element := .Resource.Schema}}assert.Equal(t, "...{{$index}}", d.Get("{{$index}}"))
		{{end}}
	}`)
	stub.stoobyDo(t, "Read_NotFound", `
	func TestResource{{.Name}}Read_NotFound(t *testing.T) {
		qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{   // read log output for correct url...
					Method:   "GET",
					Resource: "/api/2.0/...", 
					Response: common.APIErrorBody{
						ErrorCode: "NOT_FOUND",
						Message:   "Item not found",
					},
					Status: 404,
				},
			},
			Resource: Resource{{.Name}}(),
			Read: true,
			Removed: true,
			ID: "abc",
		}.ApplyNoError(t)
	}`)
	stub.stoobyDo(t, "Read_Error", `
	func TestResource{{.Name}}Read_Error(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{   // read log output for correct url...
					Method:   "GET",
					Resource: "/api/2.0/...", 
					Response: common.APIErrorBody{
						ErrorCode: "INVALID_REQUEST",
						Message:   "Internal error happened",
					},
					Status: 400,
				},
			},
			Resource: Resource{{.Name}}(),
			Read: true,
			ID: "abc",
		}.Apply(t)
		qa.AssertErrorStartsWith(t, err, "Internal error happened")
		assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
	}`)
}

func (stub *resourceTestStub) Creates(t *testing.T) {
	stub.stoobyDo(t, "Create", `
	func TestResource{{.Name}}Create(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				// request #1 - most likely POST
				// request #2 - same as in TestResource{{.Name}}Read
			},
			Resource: Resource{{.Name}}(),
			Create: true,
			HCL: `+"`"+`
			{{range $key, $element := .Resource.Schema}}{{$key}} = "..."
			{{end}}
			`+"`"+`,
		}.Apply(t)
		assert.NoError(t, err)
		assert.Equal(t, "...", d.Id())
	}`)
	stub.stoobyDo(t, "Create_Error", `
	func TestResource{{.Name}}Create_Error(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{   // read log output for better stub url...
					Method:   "POST",
					Resource: "/api/2.0/...", 
					Response: common.APIErrorBody{
						ErrorCode: "INVALID_REQUEST",
						Message:   "Internal error happened",
					},
					Status: 400,
				},
			},
			Resource: Resource{{.Name}}(),
			Create: true,
			HCL: `+"`"+`
			{{range $key, $element := .Resource.Schema}}{{$key}} = "..."
			{{end}}
			`+"`"+`,
		}.Apply(t)
		qa.AssertErrorStartsWith(t, err, "Internal error happened")
		assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
	}`)
}

func (stub *resourceTestStub) Updates(t *testing.T) {
	stub.stoobyDo(t, "Update", `
	func TestResource{{.Name}}Update(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				// request #1 - most likely POST
				// request #2 - same as in TestResource{{.Name}}Read
			},
			Resource: Resource{{.Name}}(),
			Update: true,
			ID: "abc",
			HCL: `+"`"+`
			{{range $key, $element := .Resource.Schema}}{{$key}} = "..."
			{{end}}
			`+"`"+`,
		}.Apply(t)
		assert.NoError(t, err)
		assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
	}`)
	stub.stoobyDo(t, "Update_Error", `
	func TestResource{{.Name}}Update_Error(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{   // read log output for better stub url...
					Method:   "POST",
					Resource: "/api/2.0/.../edit",
					Response: common.APIErrorBody{
						ErrorCode: "INVALID_REQUEST",
						Message:   "Internal error happened",
					},
					Status: 400,
				},
			}, 
			Resource: Resource{{.Name}}(),
			Update: true,
			ID: "abc",
			HCL: `+"`"+`
			{{range $key, $element := .Resource.Schema}}{{$key}} = "..."
			{{end}}
			`+"`"+`,
		}.Apply(t)
		qa.AssertErrorStartsWith(t, err, "Internal error happened")
		assert.Equal(t, "abc", d.Id())
	}`)
}

func (stub *resourceTestStub) Deletes(t *testing.T) {
	stub.stoobyDo(t, "Delete", `
	func TestResource{{.Name}}Delete(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{   // read log output for better stub url...
					Method:   "POST",
					Resource: "/api/2.0/.../delete",
					ExpectedRequest: map[string]string{
						"...id": "abc",
					},
				},
			},
			Resource: Resource{{.Name}}(),
			Delete: true,
			ID: "abc",
		}.Apply(t)
		assert.NoError(t, err)
		assert.Equal(t, "abc", d.Id())
	}`)
	stub.stoobyDo(t, "Delete_Error", `
	func TestResource{{.Name}}Delete_Error(t *testing.T) {
		d, err := qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{
					Method:   "POST",
					Resource: "/api/2.0/.../delete",
					Response: common.APIErrorBody{
						ErrorCode: "INVALID_REQUEST",
						Message:   "Internal error happened",
					},
					Status: 400,
				},
			},
			Resource: Resource{{.Name}}(),
			Delete: true,
			ID: "abc",
		}.Apply(t)
		qa.AssertErrorStartsWith(t, err, "Internal error happened")
		assert.Equal(t, "abc", d.Id())
	}`)
}

func TestGenerateTestCodeStubs(t *testing.T) {
	funcs := getExistingUnitTests()
	t.Logf("Got %d unit tests in total. %v",
		len(funcs), resourceTestStub{})
	t.Skip()
	p := sdkv2.DatabricksProvider()
	for name, resource := range p.ResourcesMap {
		if name != "databricks_group_instance_profile" {
			continue
		}
		stub := resourceTestStub{Resource: resource, others: &funcs}
		for i, part := range strings.Split(name, "_") {
			if i == 0 {
				// databricks_*
				continue
			}
			if part == "acl" || part == "mws" {
				stub.Name += strings.ToUpper(part)
			} else {
				stub.Name += strings.Title(part)
			}
		}
		stub.Creates(t)
		stub.Reads(t)
		stub.Deletes(t)
	}
}

func getExistingUnitTests() []string {
	set := token.NewFileSet()
	packages, err := parser.ParseDir(set, "..", nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}
	funcs := []string{}
	for _, p := range packages {
		for _, f := range p.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					fnName := fn.Name.Name
					if !strings.HasPrefix(fnName, "Test") {
						continue
					}
					if strings.HasPrefix(fnName, "TestAcc") ||
						strings.HasPrefix(fnName, "TestAzure") ||
						strings.HasPrefix(fnName, "TestAws") {
						continue
					}
					funcs = append(funcs, fnName)
				}
			}
		}
	}
	sort.Strings(funcs)
	return funcs
}
