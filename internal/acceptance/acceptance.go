package acceptance

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/databrickslabs/databricks-terraform/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Step ...
type Step struct {
	Template string
	Callback func(ctx context.Context, client *common.DatabricksClient, id string) error
	Check    func(*terraform.State) error

	Destroy                   bool
	ExpectNonEmptyPlan        bool
	ExpectError               *regexp.Regexp
	PlanOnly                  bool
	PreventDiskCleanup        bool
	PreventPostDestroyRefresh bool
	ImportState               bool
	ImportStateVerify         bool
}

// Test wrapper over terraform testing framework
func Test(t *testing.T, steps []Step, otherVars ...map[string]string) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	provider := provider.DatabricksProvider()
	cwd, err := os.Getwd()
	if err != nil {
		t.Skip(err.Error())
	}
	awsAttrs := ""
	if cloudEnv == "AWS" {
		awsAttrs = "aws_attributes {}"
	}
	instancePoolID := ""
	if cloudEnv != "MWS" {
		instancePoolID = compute.CommonInstancePoolID()
	}
	vars := map[string]string{
		"CWD":                     cwd,
		"AWS_ATTRIBUTES":          awsAttrs,
		"COMMON_INSTANCE_POOL_ID": instancePoolID,
	}
	ts := []resource.TestStep{}
	ctx := context.Background()
	client := common.CommonEnvironmentClient()

	type testResource struct {
		ID       string
		Name     string
		Resource *schema.Resource
	}

	resourceAndName := regexp.MustCompile(`resource\s+"([^"]*)"\s+"([^"]*)"`)
	resourcesEverCreated := map[testResource]bool{}
	stepConfig := ""
	for _, s := range steps {
		if s.Template != "" {
			stepConfig = qa.EnvironmentTemplate(t, s.Template, vars)
		}
		ts = append(ts, resource.TestStep{
			Config:                    stepConfig,
			Destroy:                   s.Destroy,
			ExpectNonEmptyPlan:        s.ExpectNonEmptyPlan,
			PlanOnly:                  s.PlanOnly,
			PreventDiskCleanup:        s.PreventDiskCleanup,
			PreventPostDestroyRefresh: s.PreventPostDestroyRefresh,
			ImportState:               s.ImportState,
			ImportStateVerify:         s.ImportStateVerify,
			Check: func(state *terraform.State) error {
				for n, is := range state.RootModule().Resources {
					p := strings.Split(n, ".")
					if p[0] == "data" {
						continue
					}
					r := provider.ResourcesMap[p[0]]
					resourcesEverCreated[testResource{
						ID:       is.Primary.ID,
						Name:     p[1],
						Resource: r,
					}] = true
					dia := r.ReadContext(ctx, r.Data(&terraform.InstanceState{
						ID: is.Primary.ID,
					}), client)
					if dia != nil {
						return fmt.Errorf("%v", dia)
					}
				}
				if s.Callback != nil {
					match := resourceAndName.FindStringSubmatch(stepConfig)
					rootModule := state.RootModule()
					res := rootModule.Resources[match[1]+"."+match[2]]
					id := res.Primary.ID
					return s.Callback(ctx, client, id)
				}
				if s.Check != nil {
					return s.Check(state)
				}
				return nil
			},
		})
	}
	AccTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"databricks": func() (*schema.Provider, error) {
				return provider, nil
			},
		},
		Steps: ts,
		CheckDestroy: func(t *terraform.State) error {
			// TODO: generically check if all of ID's are removed.
			return nil
		},
	})
}

// AccTest wrapper for acceptance tests
func AccTest(t *testing.T, tc resource.TestCase) {
	tc.ProviderFactories = map[string]func() (*schema.Provider, error){
		"databricks": func() (*schema.Provider, error) {
			return provider.DatabricksProvider(), nil
		},
	}
	// this allows to debug from VSCode if it's launched with CLOUD_ENV var
	cloudEnv := os.Getenv("CLOUD_ENV")
	tc.IsUnitTest = cloudEnv != ""
	// TODO: all tests must: create, edit API, edit local
	// TODO: generic resource destroy check
	// TODO: fix tmpdir issue
	if cloudEnv != "" {
		// let's be more chatty in integration test logs
		for i, s := range tc.Steps {
			if s.Config != "" {
				t.Logf("Test %s (%s) step %d config is:\n%s",
					t.Name(), cloudEnv, i,
					internal.TrimLeadingWhitespace(s.Config))
			}
		}
	}

	resource.Test(t, tc)
}

// ResourceCheck calls back a function with client and resource id
func ResourceCheck(name string,
	cb func(ctx context.Context, client *common.DatabricksClient, id string) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		client := common.CommonEnvironmentClient()
		return cb(context.Background(), client, rs.Primary.ID)
	}
}
