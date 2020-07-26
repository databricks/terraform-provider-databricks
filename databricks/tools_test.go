package databricks

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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
		d, err := ResourceTester(t, []HTTPFixture{
			// read log output of test util for further stubs...
		}, resource{{.Name}}, nil, actionWithID("abc", resource{{.Name}}Read))
		assert.NoError(t, err, err)
		assert.Equal(t, "abc", d.Id(), "Id should not be empty")
		{{range $index, $element := .Resource.Schema}}assert.Equal(t, "...{{$index}}", d.Get("{{$index}}"))
		{{end}}
	}`)
	stub.stoobyDo(t, "Read_NotFound", `
	func TestResource{{.Name}}Read_NotFound(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			{   // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/...", 
				Response: service.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		}, resource{{.Name}}, nil, actionWithID("abc", resource{{.Name}}Read))
		assert.NoError(t, err, err)
		assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
	}`)
	stub.stoobyDo(t, "Read_Error", `
	func TestResource{{.Name}}Read_Error(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			{   // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/...", 
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		}, resource{{.Name}}, nil, actionWithID("abc", resource{{.Name}}Read))
		assert.Errorf(t, err, "Internal error happened")
		assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
	}`)
}

func (stub *resourceTestStub) Creates(t *testing.T) {
	stub.stoobyDo(t, "Create", `
	func TestResource{{.Name}}Create(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			// request #1 - most likely POST
			// request #2 - same as in TestResource{{.Name}}Read
		}, resource{{.Name}}, map[string]interface{}{
			{{range $key, $element := .Resource.Schema}}"{{$key}}": "...",
			{{end}}
		}, resource{{.Name}}Create)
		assert.NoError(t, err, err)
		assert.Equal(t, "...", d.Id())
	}`)
	stub.stoobyDo(t, "Create_Error", `
	func TestResource{{.Name}}Create_Error(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			{   // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/...", 
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		}, resource{{.Name}}, map[string]interface{}{
			{{range $key, $element := .Resource.Schema}}"{{$key}}": "...",
			{{end}}
		}, resource{{.Name}}Create)
		assert.Errorf(t, err, "Internal error happened")
		assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
	}`)
}

func (stub *resourceTestStub) Updates(t *testing.T) {
	stub.stoobyDo(t, "Update", `
	func TestResource{{.Name}}Update(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			// request #1 - most likely POST
			// request #2 - same as in TestResource{{.Name}}Read
		}, resource{{.Name}}, map[string]interface{}{
			{{range $key, $element := .Resource.Schema}}"{{$key}}": "...",
			{{end}}
		}, actionWithID("abc", resource{{.Name}}Update))
		assert.NoError(t, err, err)
		assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
	}`)
	stub.stoobyDo(t, "Update_Error", `
	func TestResource{{.Name}}Update_Error(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			{   // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/.../edit",
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		}, resource{{.Name}}, map[string]interface{}{
			{{range $key, $element := .Resource.Schema}}"{{$key}}": "...",
			{{end}}
		}, actionWithID("abc", resource{{.Name}}Update))
		assert.Errorf(t, err, "Internal error happened")
		assert.Equal(t, "abc", d.Id())
	}`)
}

func (stub *resourceTestStub) Deletes(t *testing.T) {
	stub.stoobyDo(t, "Delete", `
	func TestResource{{.Name}}Delete(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			{   // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/.../delete",
				ExpectedRequest: map[string]string{
					"...id": "abc",
				},
			},
		}, resource{{.Name}}, nil, actionWithID("abc", resource{{.Name}}Delete))
		assert.NoError(t, err, err)
		assert.Equal(t, "abc", d.Id())
	}`)
	stub.stoobyDo(t, "Delete_Error", `
	func TestResource{{.Name}}Delete_Error(t *testing.T) {
		d, err := ResourceTester(t, []HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/.../delete",
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		}, resource{{.Name}}, nil, actionWithID("abc", resource{{.Name}}Delete))
		assert.Errorf(t, err, "Internal error happened")
		assert.Equal(t, "abc", d.Id())
	}`)
}

func TestGenerateTestCodeStubs(t *testing.T) {
	funcs := getExistingUnitTests()
	for name, resource := range testAccProvider.ResourcesMap {
		if name != "databricks_notebook" {
			continue
		}
		stub := resourceTestStub{Resource: resource, others: &funcs}
		for i, part := range strings.Split(name, "_") {
			if i == 0 {
				// databricks_*
				continue
			}
			if len(part) < 4 {
				stub.Name += strings.ToUpper(part)
			} else {
				stub.Name += strings.Title(part)
			}
		}
		stub.Reads(t)
		stub.Creates(t)
		if resource.Update != nil {
			stub.Updates(t)
		}
		stub.Deletes(t)
	}
}

func getExistingUnitTests() []string {
	set := token.NewFileSet()
	packages, err := parser.ParseDir(set, ".", nil, 0)
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
