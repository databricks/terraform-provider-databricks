Contributing to Databricks Terraform Provider
---

- [Contributing to Databricks Terraform Provider](#contributing-to-databricks-terraform-provider)
- [Installing from source](#installing-from-source)
- [Developing provider](#developing-provider)
- [Developing with Visual Studio Code Devcontainers](#developing-with-visual-studio-code-devcontainers)
- [Building and Installing with Docker](#building-and-installing-with-docker)
- [Testing](#testing)
- [Code conventions](#code-conventions)
- [Linting](#linting)
- [Unit testing resources](#unit-testing-resources)
- [Generating asserts for the first time in test](#generating-asserts-for-the-first-time-in-test)
- [Random naming anywhere](#random-naming-anywhere)
- [Integration Testing](#integration-testing)
- [Pre-release procedure](#pre-release-procedure)

We happily welcome contributions to databricks-terraform. We use GitHub Issues to track community reported issues and GitHub Pull Requests for accepting changes.

## Installing for Terraform 0.12

If you use Terraform 0.12, please execute the following curl command in your shell:

```bash
curl https://raw.githubusercontent.com/databrickslabs/databricks-terraform/master/godownloader-databricks-provider.sh | bash -s -- -b $HOME/.terraform.d/plugins
```

## Installing from source

The following command (tested on Ubuntu 20.04) will install `make`, `golang`, `git` with all of the dependent packages as well as Databricks Terrafrom provider from sources. Required version of GoLang is at least 1.13. Required version of terraform is at least 0.12. 

On MacOS X, you can install GoLang through `brew install go`, on Debian-based Linux, you can install it by `sudo apt-get install golang -y`.

```bash
git clone https://github.com/databrickslabs/terraform-provider-databricks.git
cd terraform-provider-databricks
make install
```

Most likely, `terraform init -upgrade -verify-plugins=false -lock=false` would be a very great command to use.

## Developing provider

After installing necessary software for building provider from sources, you should install `golangci-lint` and `gotestsum` in order to run `make test`.

Make sure you have `$GOPATH/bin` in your `$PATH`:
```
echo "export PATH=\$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
```

Installing `golangci-lint`:
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
$(go env GOPATH)/bin/golangci-lint
```

Installing `gotestsum`:
```bash
go get gotest.tools/gotestsum
```

Installing `goimports`:
```bash
go get golang.org/x/tools/cmd/goimports
```

After this, you should be able to run `make test`.

## Developing with Visual Studio Code Devcontainers

This project has configuration for working with [Visual Studio Code Devcontainers](https://code.visualstudio.com/docs/remote/containers) - this allows you to containerise your development prerequisites (e.g. golang, terraform). To use this you will need [Visual Studio Code](https://code.visualstudio.com/) and [Docker](https://www.docker.com/products/docker-desktop).

To get started, clone this repo and open the folder with Visual Studio Code. If you don't have the [Remote Development extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) then you should be prompted to install it.

Once the folder is loaded and the extension is installed you should be prompted to re-open the folder in a devcontainer. This will built and run the container image with the correct tools (and versions) ready to start working on and building the code. The in-built terminal will launch a shell inside the container for running `make` commands etc.

See the docs for more details on working with [devcontainers](https://code.visualstudio.com/docs/remote/containers).

## Building and Installing with Docker

To install and build the code if you dont want to install golang, terraform, etc. All you need is docker and git.

First make sure you clone the repository and you are in the directory.

Then build the docker image with this command:

```bash
$ docker build -t databricks-terraform . 
```

Then run the execute the terraform binary via the following command and volume mount. Make sure that you are in the directory
 with the terraform code. The following command you can execute the following commands and additional ones as part of 
 the terraform binary.
 
```bash
$ docker run -it -v $(pwd):/workpace -w /workpace databricks-terraform init
$ docker run -it -v $(pwd):/workpace -w /workpace databricks-terraform plan
$ docker run -it -v $(pwd):/workpace -w /workpace databricks-terraform apply
```

## Adding a new resource

The general process for adding a new resource is:

*Define the resource models.* The models for a resource are `struct`s defining the schemas of the objects in the Databricks REST API. Define structures used for multiple resources in a common `models.go` file; otherwise, you can define these directly in your resource file. An example model:
```go
type Field struct {
 A string `json:"a,omitempty"`
 AMoreComplicatedName int `json:"a_more_complicated_name,omitempty"`
}

type Example struct {
 ID string `json:"id"`
 TheField *Field `json:"the_field"`
 AnotherField bool `json:"another_field"`
 Filters []string `json:"filters" tf:"optional"`
}
```

Some interesting points to note here:
* Use the `json` tag to determine the serde properties of the field. The allowed tags are defined here: https://go.googlesource.com/go/+/go1.16/src/encoding/json/encode.go#158
* Use the custom `tf` tag indicates properties to be annotated on the Terraform schema for this struct. Supported values are:
  * `optional` for optional fields
  * `computed` for computed fields
  * `alias:X` to use a custom name in HCL for a field
  * `default:X` to set a default value for a field
  * `max_items:N` to set the maximum number of items for a multi-valued parameter
  * `slice_set` to indicate that a the parameter should accept a set instead of a list
* Do not use bare references to structs in the model; rather, use pointers to structs. Maps and slices are permitted, as well as the following primitive types: int, int32, int64, float64, bool, string.
See `typeToSchema` in `common/reflect_resource.go` for the up-to-date list of all supported field types and values for the `tf` tag.

*Define the Terraform schema.* This is made easy for you by the `StructToSchema` method in the `common` package, which converts your struct automatically to a Terraform schema, accepting also a function allowing the user to post-process the automatically generated schema, if needed.
```go
var exampleSchema = common.StructToSchema(Example{}, func(m map[string]*schema.Schema) map[string]*schema.Schema { return m })
```

*Define the API client for the resource.* You will need to implement create, read, update, and delete functions.
```go
type ExampleApi struct {
	client *common.DatabricksClient
	ctx    context.Context
}

func NewExampleApi(ctx context.Context, m interface{}) ExampleApi {
	return ExampleApi{m.(*common.DatabricksClient), ctx}
}

func (a ExampleApi) Create(e Example) (string, error) {
	var id string
	err := a.client.Post(a.ctx, "/example", e, &id)
	return id, err
}

func (a ExampleApi) Read(id string) (e Example, err error) {
	err = a.client.Get(a.ctx, "/example/"+id, nil, &e)
	return
}

func (a ExampleApi) Update(id string, e Example) error {
	return a.client.Put(a.ctx, "/example/"+string(id), e)
}

func (a ExampleApi) Delete(id string) error {
	return a.client.Delete(a.ctx, "/pipelines/"+id, nil)
}
```

*Define the Resource object itself.* This is made quite simple by using the `toResource` function defined on the `Resource` type in the `common` package. A simple example:
```go
func ResourceExample() *schema.Resource {
	return common.Resource{
		Schema:        exampleSchema,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var e Example
			err := common.DataToStructPointer(d, exampleSchema, &e)
			if err != nil {
				return err
			}
			id, err := NewExampleApi(ctx, c).Create(e)
			if err != nil {
				return err
			}
			d.SetId(string(id))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			i, err := NewExampleApi(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(i.Spec, exampleSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var e Example
			err := common.DataToStructPointer(d, exampleSchema, &e)
			if err != nil {
				return err
			}
			return NewExampleApi(ctx, c).Update(d.Id(), e)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewExampleApi(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
```

*Add the resource to the top-level provider.* Simply add the resource to the provider definition in `provider/provider.go`.

*Write unit tests for your resource.* To write your unit tests, you can make use of `ResourceFixture` and `HTTPFixture` structs defined in the `qa` package. This starts a fake HTTP server, asserting that your resource provdier generates the correct request for a given HCL template body for your resource. An example:

```go
func TestExampleResourceCreate(t *testing.T) {
		d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/example",
				ExpectedRequest: Example{
					TheField: Field{
						A: "test",
					},
				},
				Response: map[string]interface{} {
					"id": "abcd",
					"the_field": map[string]interface{} {
						"a": "test",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/example/abcd",
				Response: map[string]interface{}{
					"id":    "abcd",
					"the_field": map[string]interface{} {
						"a": "test",
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceExample(),
		HCL: `the_field {
			a = "test"
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abcd", d.Id())
}
```

*Write acceptance tests.* These are E2E tests which run terraform against the live cloud and Databricks APIs. For these, you can use the `Test` and `Step` structs defined in the `acceptance` package. An example:

```go
func TestPreviewAccPipelineResource_CreatePipeline(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_example" "this" {
				the_field {
					a = "test"
					a_more_complicated_name = 3
				}
				another_field = true
				filters = [
					"a",
					"b"
				]
			}
			`,
		},
	})
}
```

## Debugging

**TF_LOG=DEBUG terraform apply** allows you to see the internal logs from `terraform apply`.

## Testing

* [Integration tests](scripts/README.md) should be run at a client level against both azure and aws to maintain sdk parity against both apis.
* Terraform acceptance tests should be run against both aws and azure to maintain parity of provider between both cloud services
* Consider test functions as scenarios, that you are debugging from IDE when specific issues arise. Test tables are discouraged. Single-use functions in tests are discouraged, unless resource definitions they make are longer than 80 lines.
* All tests should be capable of repeatedly running on "dirty" environment, which means not requiring a new clean environment every time the test runs.
* All tests should re-use compute resources whenever possible.
* Prefer `require.NoError` (stops the test on error) to `assert.NoError` (continues the test on error) when checking the results.

## Code conventions

* Files should not be larger than 600 lines
* Single function should fit to be seen on 13" screen: no more than 40 lines of code. Only exception to this rule is `*_test.go` files. 
* There should be no unnecessary package exports: no structs & types with leading capital letter, unless they are of value outside of the package.
* `fmt.Sprintf` with more than 4 placeholders is considered too complex to maintain. Should be avoided at all cost. Use `qa.EnvironmentTemplate(t, "This is {env.DATABRICKS_HOST} with {var.RANDOM} name.")` instead
* Import statements should all be first ordered by "GoLang internal", "Vendor packages" and then "current provider packages". Within those sections imports must follow alphabetical order.

## Linting

Please use makefile for linting. If you run `golangci-lint` by itself it will fail due to different tags containing same functions. 
So please run `make lint` instead.

## Unit testing resources

Eventually, all of resources would be automatically checked for a unit test presence. `TestGenerateTestCodeStubs` is going to fail, when resource has certain test cases missing. Until all existing resources have tests, you can generate stub code, which will be logged to stdout by changing these lines of `generate_test.go` with name of resource you're creating:

```go
for name, resource := range p.ResourcesMap {
	if name != "databricks_user" {
		continue
	}
	//...
```

In order to unit test a resource, which runs fast and could be included in code coverage, one should use `ResourceTester`, that launches embedded HTTP server with `HTTPFixture`'s containing all calls that should have been made in given scenario. Some may argue that this is not a pure unit test, because it creates a side effect in form of embedded server, though it's always on different random port, making it possible to execute these tests in parallel. Therefore comments about non-pure unit tests will be ignored, if they use `ResourceTester` helper.

```go
func TestPermissionsCreate(t *testing.T) {
	_, err := internal.ResourceTester(t, []qa.HTTPFixture{
		{
            Method:   http.MethodPatch,
            // requires full URI
            Resource: "/api/2.0/preview/permissions/clusters/abc",
            // works with entities, not JSON. Diff is displayed in case of missmatch 
			ExpectedRequest: AccessControlChangeList{
				AccessControlList: []*AccessControlChange{
					{
						UserName:        &TestingUser,
						PermissionLevel: "CAN_USE",
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/permissions/clusters/abc",
			Response: AccessControlChangeList{
				AccessControlList: []*AccessControlChange{
					{
						UserName:        &TestingUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/scim/v2/Me",
			Response: User{
				UserName: "chuck.norris",
			},
		},
    }, 
    // next argument is function, that creates resource (to make schema for ResourceData)
    resourcePermissions, 
    // state represented as native structure (though a bit clunky)
    map[string]interface{}{
		"cluster_id": "abc",
		"access_control": []interface{}{
			map[string]interface{}{
				"user_name":        TestingUser,
				"permission_level": "CAN_USE",
			},
		},
    },
    // the last argument is a function, that performs a stage on resource (Create/update/delete/read)
    resourcePermissionsCreate)
	assert.NoError(t, err, err)
}
```

Each resource should have both unit and integration tests. 

## Generating asserts for the first time in test

```go
for k, v := range d.State().Attributes {
	fmt.Printf("assert.Equal(t, %#v, d.Get(%#v))\n", v, k)
}
```

## Random naming anywhere

Terraform SDK provides `randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)` for convenient random names generation.

## Integration Testing

Currently Databricks supports two cloud providers `azure` and `aws` thus integration testing with the correct cloud service provider is 
crucial for making sure that the provider behaves as expected on all supported clouds. Please read [dedicated instructions](scripts/README.md) for details.

## Pre-release procedure

1. `make test-azure` 
2. `make test-mws` if MWS related code changed given release.
3. Create release notes.
4. Perfrom backwards-compatibility checks and make proper notes. 