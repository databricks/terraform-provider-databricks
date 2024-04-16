package acceptance

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	dbproviderlogger "github.com/databricks/terraform-provider-databricks/logger"
	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
	databricks.WithProduct("tf-integration-tests", common.Version())
	os.Setenv("TF_LOG", "DEBUG")
	dbproviderlogger.SetLogger()
}

func workspaceLevel(t *testing.T, steps ...step) {
	loadWorkspaceEnv(t)
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	if os.Getenv("DATABRICKS_ACCOUNT_ID") != "" {
		skipf(t)("Skipping workspace test on account level")
	}
	t.Parallel()
	run(t, steps)
}

func accountLevel(t *testing.T, steps ...step) {
	loadAccountEnv(t)
	cfg := &config.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	}
	err := cfg.EnsureResolved()
	if err != nil {
		skipf(t)("error: %s", err)
	}
	if !cfg.IsAccountClient() {
		skipf(t)("Not in account env: %s/%s", cfg.AccountID, cfg.Host)
	}
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()
	run(t, steps)
}

func unityWorkspaceLevel(t *testing.T, steps ...step) {
	loadUcwsEnv(t)
	GetEnvOrSkipTest(t, "TEST_METASTORE_ID")
	if os.Getenv("DATABRICKS_ACCOUNT_ID") != "" {
		skipf(t)("Skipping workspace test on account level")
	}
	t.Parallel()
	run(t, steps)
}

func unityAccountLevel(t *testing.T, steps ...step) {
	loadUcacctEnv(t)
	GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	GetEnvOrSkipTest(t, "TEST_METASTORE_ID")
	t.Parallel()
	run(t, steps)
}

// A step in a terraform acceptance test
type step struct {
	// Terraform HCL for resources to materialize in this test step.
	Template string

	// This function is called after the template is applied. Useful for making assertions
	// or doing cleanup.
	Check func(*terraform.State) error

	// Setup function called before the template is materialized.
	PreConfig func()

	Destroy                   bool
	ExpectNonEmptyPlan        bool
	ExpectError               *regexp.Regexp
	PlanOnly                  bool
	PreventDiskCleanup        bool
	PreventPostDestroyRefresh bool
	ImportState               bool
	ImportStateVerify         bool
}

func createUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "10000000-2000-3000-4000-500000000000"
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// environmentTemplate asserts existence and fills in {env.VAR} & {var.RANDOM} placeholders in template.
// For writing a unit test to intercept the errors (t.Fatalf literally ends the test in failure)
func environmentTemplate(t *testing.T, template string, otherVars ...map[string]string) string {
	vars := map[string]string{
		"RANDOM":      qa.RandomName("t"),
		"RANDOM_UUID": createUuid(),
	}
	if len(otherVars) > 1 {
		skipf(t)("cannot have more than one custom variable map")
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
			skipf(t)("Missing %s %s variable.", varType, varName)
			missing++
			continue
		}
		template = strings.ReplaceAll(template, `{`+varType+`.`+varName+`}`, value)
	}
	if missing > 0 {
		skipf(t)("please set %d variables and restart", missing)
	}
	return commands.TrimLeadingWhitespace(template)
}

// Test wrapper over terraform testing framework. Multiple steps share the same
// terraform state context.
func run(t *testing.T, steps []step) {
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
	if cloudEnv == "aws" {
		awsAttrs = "aws_attributes {}"
	}
	vars := map[string]string{
		"CWD":            cwd,
		"STICKY_RANDOM":  qa.RandomName("s"),
		"AWS_ATTRIBUTES": awsAttrs,
	}
	ts := []resource.TestStep{}
	ctx := context.Background()

	for i, s := range steps {
		stepConfig := ""
		if s.Template != "" {
			stepConfig = environmentTemplate(t, s.Template, vars)
		}
		stepNum := i
		thisStep := s
		stepCheck := thisStep.Check
		stepPreConfig := s.PreConfig
		ts = append(ts, resource.TestStep{
			PreConfig: func() {
				if stepConfig == "" {
					return
				}
				logger.Infof(ctx, "Test %s (%s) step %d config is:\n%s",
					t.Name(), cloudEnv, stepNum,
					commands.TrimLeadingWhitespace(stepConfig))

				if stepPreConfig != nil {
					stepPreConfig()
				}
			},
			Config:                    stepConfig,
			Destroy:                   s.Destroy,
			ExpectNonEmptyPlan:        s.ExpectNonEmptyPlan,
			PlanOnly:                  s.PlanOnly,
			PreventDiskCleanup:        s.PreventDiskCleanup,
			PreventPostDestroyRefresh: s.PreventPostDestroyRefresh,
			ImportState:               s.ImportState,
			ImportStateVerify:         s.ImportStateVerify,
			ExpectError:               s.ExpectError,
			Check: func(state *terraform.State) error {
				// get configured client from provider
				client := provider.Meta().(*common.DatabricksClient)

				// Default check for all runs. Asserts that the read operation succeeds.
				for n, is := range state.RootModule().Resources {
					p := strings.Split(n, ".")

					// Skip data resources.
					if p[0] == "data" {
						continue
					}
					r := provider.ResourcesMap[p[0]]
					dia := r.ReadContext(ctx, r.Data(is.Primary), client)
					if dia != nil {
						return fmt.Errorf("%v", dia)
					}
				}
				if stepCheck != nil {
					return stepCheck(state)
				}
				return nil
			},
		})
	}
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
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

// resourceCheck calls back a function with client and resource id
func resourceCheck(name string,
	cb func(ctx context.Context, client *common.DatabricksClient, id string) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		client, err := client.New(&config.Config{})
		if err != nil {
			panic(err)
		}
		return cb(context.Background(), &common.DatabricksClient{
			DatabricksClient: client,
		}, rs.Primary.ID)
	}
}

// resourceCheckWithState calls back a function with client and resource instance state
func resourceCheckWithState(name string,
	cb func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}
		client, err := client.New(&config.Config{})
		if err != nil {
			panic(err)
		}
		return cb(context.Background(), &common.DatabricksClient{
			DatabricksClient: client,
		}, rs.Primary)
	}
}

const fullCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const hexCharset = "0123456789abcdef"

// GetEnvOrSkipTest proceeds with test only with that env variable
func GetEnvOrSkipTest(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		skipf(t)("Environment variable %s is missing", name)
	}
	return value
}

func GetEnvInt64OrSkipTest(t *testing.T, name string) int64 {
	v := GetEnvOrSkipTest(t, name)
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		skipf(t)("`%s` is not int64: %s", v, err)
	}
	return i
}

// RandomEmail generates random email
func RandomEmail(prefix ...string) string {
	return fmt.Sprintf("%s@example.com", RandomName(
		append([]string{"sdk-go-"}, prefix...)...))
}

// RandomName gives random name with optional prefix. e.g. qa.RandomName("tf-")
func RandomName(prefix ...string) string {
	rand.Seed(time.Now().UnixNano())
	randLen := 12
	b := make([]byte, randLen)
	for i := range b {
		b[i] = fullCharset[rand.Intn(randLen)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", strings.Join(prefix, ""), b)
	}
	return string(b)
}

func RandomHex(prefix string, randLen int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, randLen)
	for i := range b {
		b[i] = hexCharset[rand.Intn(randLen)%len(hexCharset)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", prefix, b)
	}
	return string(b)
}

func skipf(t *testing.T) func(format string, args ...any) {
	if isInDebug() {
		// VSCode "debug test" feature doesn't show dlv logs,
		// so that we fail here for maintainer productivity.
		return t.Fatalf
	}
	return t.Skipf
}

// detects if test is run from "debug test" feature in VSCode
func isInDebug() bool {
	ex, _ := os.Executable()
	return strings.HasPrefix(path.Base(ex), "__debug_bin")
}

func setDebugLogger() {
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
}

func loadWorkspaceEnv(t *testing.T) {
	setDebugLogger()
	loadDebugEnvIfRunsFromIDE(t, "workspace")
}

func loadAccountEnv(t *testing.T) {
	setDebugLogger()
	loadDebugEnvIfRunsFromIDE(t, "account")
}

func loadUcwsEnv(t *testing.T) {
	setDebugLogger()
	loadDebugEnvIfRunsFromIDE(t, "ucws")
}

func loadUcacctEnv(t *testing.T) {
	setDebugLogger()
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
}

func isAws(t *testing.T) bool {
	awsCloudEnvs := []string{"MWS", "aws", "ucws", "ucacct"}
	return isCloudEnvInList(t, awsCloudEnvs)
}

func isAzure(t *testing.T) bool {
	azureCloudEnvs := []string{"azure", "azure-ucacct"}
	return isCloudEnvInList(t, azureCloudEnvs)
}

func isGcp(t *testing.T) bool {
	gcpCloudEnvs := []string{"gcp-accounts", "gcp-ucacct", "gcp-ucws", "gcp"}
	return isCloudEnvInList(t, gcpCloudEnvs)
}

func isCloudEnvInList(t *testing.T, cloudEnvs []string) bool {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		skipf(t)("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	return slices.Contains(cloudEnvs, cloudEnv)
}

func isAuthedAsWorkspaceServicePrincipal(ctx context.Context) (bool, error) {
	w := databricks.Must(databricks.NewWorkspaceClient())
	user, err := w.CurrentUser.Me(ctx)
	if err != nil {
		return false, err
	}
	for _, emailValue := range user.Emails {
		if emailValue.Primary && strings.Contains(emailValue.Value, "@") {
			return false, nil
		}
	}
	return true, nil
}

// loads debug environment from ~/.databricks/debug-env.json
func loadDebugEnvIfRunsFromIDE(t *testing.T, key string) {
	if !isInDebug() {
		return
	}
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("cannot find user home: %s", err)
	}
	raw, err := os.ReadFile(filepath.Join(home, ".databricks/debug-env.json"))
	if err != nil {
		t.Fatalf("cannot load ~/.databricks/debug-env.json: %s", err)
	}
	var conf map[string]map[string]string
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		t.Fatalf("cannot parse ~/.databricks/debug-env.json: %s", err)
	}
	vars, ok := conf[key]
	if !ok {
		t.Fatalf("~/.databricks/debug-env.json#%s not configured", key)
	}
	for k, v := range vars {
		os.Setenv(k, v)
	}
}
