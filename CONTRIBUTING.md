# Contributing to Databricks Terraform Provider

---

- [Issues for new contributors](#issues-for-new-contributors)
- [Contribution Workflow](#contribution-workflow)
- [Contributing to Databricks Terraform Provider](#contributing-to-databricks-terraform-provider)
- [Installing from source](#installing-from-source)
- [Contributing documentation](#contributing-documentation)
- [Developing provider](#developing-provider)
- [Debugging](#debugging)
- [Adding a new resource](#adding-a-new-resource)
- [Testing](#testing)
- [Code conventions](#code-conventions)
- [Linting](#linting)
- [Developing with Visual Studio Code Devcontainers](#developing-with-visual-studio-code-devcontainers)

We happily welcome contributions to the Databricks Terraform Provider. We use GitHub Issues to track community reported issues and GitHub Pull Requests for accepting changes.

## Issues for new contributors

New contributors should look for the following tags when searching for a first contribution to the Databricks code base. We strongly recommend that new contributors tackle "good first issue" projects first; this helps the contributor become familiar with the contribution workflow, and for the core devs to become acquainted with the contributor.

[good first issue](https://github.com/databricks/terraform-provider-databricks/labels/good%20first%20issue)

## Contribution Workflow

Code contributions—bug fixes, new development, test improvement — all follow a GitHub-centered workflow. To participate in Databricks Terraform provider development, set up a GitHub account. Then:

1. Fork the repo you plan to work on. Go to the project repo page and use the Fork button. This will create a copy of the repo, under your username. (For more details on how to fork a repository see [this guide](https://help.github.com/articles/fork-a-repo/).)

1. Clone down the repo to your local system.

    ```bash
    git clone git@github.com:YOUR_USER_NAME/terraform-provider-databricks.git
    ```

1. Create a new branch to hold your work.

    ```bash
    git checkout -b new-branch-name
    ```

1. Work on your new code. Write and run tests.

1. Commit your changes.

    ```bash
    git add -A

    git commit -m "commit message here"
    ```

1. Push your changes to your GitHub repo.

    ```bash
    git push origin branch-name
    ```

1. Open a Pull Request (PR). Go to the original project repo on GitHub. There will be a message about your recently pushed branch, asking if you would like to open a pull request. Follow the prompts, compare across repositories, and submit the PR. This will send an email to the committers. You may want to consider sending an email to the mailing list for more visibility. (For more details, see the [GitHub guide on PRs](https://help.github.com/articles/creating-a-pull-request-from-a-fork).)

Maintainers and other contributors will review your PR. Please participate in the conversation, and try to make any requested changes. Once the PR is approved, the code will be merged.

Additional git and GitHub resources:

[Git documentation](https://git-scm.com/documentation)
[Git development workflow](https://docs.scipy.org/doc/numpy/dev/development_workflow.html)
[Resolving merge conflicts](https://help.github.com/articles/resolving-a-merge-conflict-using-the-command-line/)

## Installing from source

On MacOS X, you can install GoLang through `brew install go`, on Debian-based Linux, you can install it by `sudo apt-get install golang -y`.

```bash
git clone https://github.com/databricks/terraform-provider-databricks.git
cd terraform-provider-databricks
make install
```

Most likely, `terraform init -upgrade -verify-plugins=false -lock=false` would be a very great command to use.

## Contributing documentation

Make sure you have [`terrafmt`](https://github.com/katbyte/terrafmt) installed to be able to run `make fmt-docs`

```bash
go install github.com/katbyte/terrafmt@latest
```

All documentation contributions should be as detailed as possible and follow the [required format](https://www.terraform.io/registry/providers/docs). The following additional checks must also be valid:

- `make fmt-docs` to make sure code examples are consistent
- Correct rendering with Terraform Registry Doc Preview Tool - <https://registry.terraform.io/tools/doc-preview>
- Cross-linking integrity between markdown files. Pay special attention, when resource doc refers to data doc or guide.

## Developing provider

In order to simplify development workflow, you should use [dev_overrides](https://www.terraform.io/cli/config/config-file#development-overrides-for-provider-developers) section in your `~/.terraformrc` file. Please run `make build` and replace "provider-binary" with the path to `terraform-provider-databricks` executable in your current working directory:

```bash
$ cat ~/.terraformrc
provider_installation {
   dev_overrides {
     "databricks/databricks" = "provider-binary"
   }
   direct {}
}
```

After installing the necessary software for building provider from sources, you should be able to run `make coverage` to run the tests and see the coverage.

## Developing Resources or Data Sources using Plugin Framework 

### Package organization for Providers
We are migrating the resource from SDKv2 to Plugin Framework provider and hence both of them exist in the codebase. For uniform code convention, readability and development, they are organized in the `internal/providers` directory under root as follows: 
- `providers`: Contains the changes that `depends` on both internal/providers/sdkv2 and internal/providers/pluginfw packages, eg: `GetProviderServer`.
- `common`: Contains the changes `used by` both internal/providers/sdkv2 and internal/providers/pluginfw packages, eg: `ProviderName`.
- `pluginfw`: Contains the changes specific to Plugin Framework. This package shouldn't depend on sdkv2 or common.
- `sdkv2`: Contains the changes specific to SDKv2. This package shouldn't depend on pluginfw or common.

### Migrating resource to plugin framework
Ideally there shouldn't be any behaviour change when migrating a resource or data source to either Go SDk or Plugin Framework.
- Please make sure there are no breaking differences due to changes in schema by running: `make diff-schema`. 
- Integration tests shouldn't require any major changes.   


### Code Organization
Each resource and data source should be defined in package `internal/providers/pluginfw/resources/<resource>`, e.g.: `internal/providers/pluginfw/resources/volume` package will contain both resource, data sources and other utils specific to volumes. Tests (both unit and integration tests) will also remain in this package. 

Note: Only Docs will stay under root docs/ directory.


### Code Conventions
1. Make sure the resource or data source implemented is of the right type:
    ```golang
    var _ resource.ResourceWithConfigure = &QualityMonitorResource{}
    var _ datasource.DataSourceWithConfigure = &VolumesDataSource{}
    ```
2. To get the databricks client, `func (*common.DatabricksClient).GetWorkspaceClient()` or `func (*common.DatabricksClient).GetAccountClient()` will be used instead of directly using the underlying `WorkspaceClient()`, `AccountClient()` functions respectively.  
3. Any method that returns only diagnostics should be called inline while appending diagnostics in response. Example:
    ```golang
    resp.Diagnostics.Append(req.Plan.Get(ctx, &monitorInfoTfSDK)...)
    if resp.Diagnostics.HasError() {
        return
    }
    ```
    is preferred over the following:
    ```golang
    diags := req.Plan.Get(ctx, &monitorInfoTfSDK)
    if diags.HasError() {
        resp.Diagnostics.Append(diags...)
        return
    }
    ```
4. Any method returning an error should directly be followed by appending that to the diagnostics. 
    ```golang
    err := method()
    if err != nil { 
        resp.Diagnostics.AddError("message", err.Error())
        return 
    }
    ```
5. Any method returning a value alongside Diagnostics should also directly be followed by appending that to the diagnostics.


## Debugging

**TF_LOG_PROVIDER=DEBUG terraform apply** allows you to see the internal logs from `terraform apply`.

You can [run provider in a debug mode](https://www.terraform.io/plugin/sdkv2/debugging#running-terraform-with-a-provider-in-debug-mode) from VScode IDE by launching `Debug Provider` run configuration and invoking `terraform apply` with `TF_REATTACH_PROVIDERS` environment variable.

## Adding a new resource

Boilerplate for data sources could be generated via `go run provider/gen/main.go -name mws_workspaces -package mws -is-data -dry-run=false`.

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

- Use the `json` tag to determine the serde properties of the field. The allowed tags are defined here: <https://go.googlesource.com/go/+/go1.16/src/encoding/json/encode.go#158>
- Use the custom `tf` tag indicates properties to be annotated on the Terraform schema for this struct. Supported values are:
  - `optional` for optional fields
  - `computed` for computed fields
  - `alias:X` to use a custom name in HCL for a field
  - `default:X` to set a default value for a field
  - `max_items:N` to set the maximum number of items for a multi-valued parameter
  - `slice_set` to indicate that a the parameter should accept a set instead of a list
  - `sensitive` to mark a field as sensitive and prevent Terraform from showing its value in the plan or apply output
  - `force_new` to indicate a change in this value requires the replacement (destroy and create) of the resource
  - `suppress_diff` to allow comparison based on something other than primitive, list or map equality, either via a `CustomizeDiffFunc`, or the default diff for the type of the schema
- Do not use bare references to structs in the model; rather, use pointers to structs. Maps and slices are permitted, as well as the following primitive types: int, int32, int64, float64, bool, string.
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
  Schema: exampleSchema,
  Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
   var e Example
   common.DataToStructPointer(d, exampleSchema, &e)
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
   common.DataToStructPointer(d, exampleSchema, &e)
   return NewExampleApi(ctx, c).Update(d.Id(), e)
  },
  Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
   return NewExampleApi(ctx, c).Delete(d.Id())
  },
 }.ToResource()
}
```

*Add the resource to the top-level provider.* Simply add the resource to the provider definition in `provider/provider.go`.

*Write unit tests for your resource.* To write your unit tests, you can make use of `ResourceFixture` and `HTTPFixture` structs defined in the `qa` package. This starts a fake HTTP server, asserting that your resource provider generates the correct request for a given HCL template body for your resource. Update tests should have `InstanceState` field in order to test various corner-cases, like `ForceNew` schemas. It's possible to expect fixture to require new resource by specifying `RequiresNew` field. With the help of `qa.ResourceCornerCases` and `qa.ResourceFixture` one can achieve 100% code coverage for all of the new code.

A simple example:

```go
func TestExampleCornerCases(t *testing.T) {
 qa.ResourceCornerCases(t, ResourceExample())
}

func TestExampleResourceCreate(t *testing.T) {
 qa.ResourceFixture{
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
 }.ApplyNoError(t)
}
```

*Write acceptance tests.* These are E2E tests which run terraform against the live cloud and Databricks APIs. For these, you can use the `Step` helpers defined in the `internal/acceptance` package. An example:

```go
func TestAccSecretAclResource(t *testing.T) {
 WorkspaceLevel(t, Step{
  Template: `
  resource "databricks_group" "ds" {
   display_name = "data-scientists-{var.RANDOM}"
  }
  resource "databricks_secret_scope" "app" {
   name = "app-{var.RANDOM}"
  }
  resource "databricks_secret_acl" "ds_can_read_app" {
   principal = databricks_group.ds.display_name
   permission = "READ"
   scope = databricks_secret_scope.app.name
  }`,
  Check: func(s *terraform.State) error {
   w := databricks.Must(databricks.NewWorkspaceClient())

   ctx := context.Background()
   me, err := w.CurrentUser.Me(ctx)
   require.NoError(t, err)

   scope := s.RootModule().Resources["databricks_secret_scope.app"].Primary.ID
   acls, err := w.Secrets.ListAclsByScope(ctx, scope)
   require.NoError(t, err)
   assert.Equal(t, 2, len(acls.Items))
   m := map[string]string{}
   for _, acl := range acls.Items {
    m[acl.Principal] = string(acl.Permission)
   }

   group := s.RootModule().Resources["databricks_group.ds"].Primary.Attributes["display_name"]
   require.Contains(t, m, group)
   assert.Equal(t, "READ", m[group])
   assert.Equal(t, "MANAGE", m[me.UserName])
   return nil
  },
 })
}
```

## Testing

- [Integration tests](scripts/README.md) should be run at a client level against both azure and aws to maintain sdk parity against both apis.
- Terraform acceptance tests should be run against both aws and azure to maintain parity of provider between both cloud services
- Consider test functions as scenarios, that you are debugging from IDE when specific issues arise. Test tables are discouraged. Single-use functions in tests are discouraged, unless resource definitions they make are longer than 80 lines.
- All tests should be capable of repeatedly running on "dirty" environment, which means not requiring a new clean environment every time the test runs.
- All tests should re-use compute resources whenever possible.
- Prefer `require.NoError` (stops the test on error) to `assert.NoError` (continues the test on error) when checking the results.

## Code conventions

- Files should not be larger than 600 lines
- Single function should fit to be seen on 13" screen: no more than 40 lines of code. Only exception to this rule is `*_test.go` files.
- There should be no unnecessary package exports: no structs & types with leading capital letter, unless they are of value outside of the package.
- `fmt.Sprintf` with more than 4 placeholders is considered too complex to maintain. Should be avoided at all cost. Use `qa.EnvironmentTemplate(t, "This is {env.DATABRICKS_HOST} with {var.RANDOM} name.")` instead
- Import statements should all be first ordered by "GoLang internal", "Vendor packages" and then "current provider packages". Within those sections imports must follow alphabetical order.

## Linting

Please use makefile for linting. If you run `staticcheck` by itself it will fail due to different tags containing same functions.
So please run `make lint` instead.

## Developing with Visual Studio Code Devcontainers

NOTE: This use of devcontainers for terraform-provider-databricks development is **experimental** and not officially supported by Databricks

This project has configuration for working with [Visual Studio Code Devcontainers](https://code.visualstudio.com/docs/remote/containers) - this allows you to containerize your development prerequisites (e.g. golang, terraform). To use this you will need [Visual Studio Code](https://code.visualstudio.com/) and [Docker](https://www.docker.com/products/docker-desktop).

To get started, clone this repo and open the folder with Visual Studio Code. If you don't have the [Remote Development extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) then you should be prompted to install it.

Once the folder is loaded and the extension is installed you should be prompted to re-open the folder in a devcontainer. This will built and run the container image with the correct tools (and versions) ready to start working on and building the code. The in-built terminal will launch a shell inside the container for running `make` commands etc.

See the docs for more details on working with [devcontainers](https://code.visualstudio.com/docs/remote/containers).
