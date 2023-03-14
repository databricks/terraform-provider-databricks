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
	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
	databricks.WithProduct("tf-integration-tests", common.Version())
	if isInDebug() {
		// Terraform SDK v2 intercepts default logger
		// that Go SDK SimpleLogger is using, so we have
		// to re-implement one again.
		logger.DefaultLogger = stdErrLogger{}
	}
}

func workspaceLevel(t *testing.T, steps ...step) {
	loadDebugEnvIfRunsFromIDE(t, "workspace")
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	if os.Getenv("DATABRICKS_ACCOUNT_ID") != "" {
		skipf(t)("Skipping workspace test on account level")
	}
	t.Parallel()
	run(t, steps)
}

func accountLevel(t *testing.T, steps ...step) {
	loadDebugEnvIfRunsFromIDE(t, "account")
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
	loadDebugEnvIfRunsFromIDE(t, "ucws")
	GetEnvOrSkipTest(t, "TEST_METASTORE_ID")
	if os.Getenv("DATABRICKS_ACCOUNT_ID") != "" {
		skipf(t)("Skipping workspace test on account level")
	}
	t.Parallel()
	run(t, steps)
}

func unityAccountLevel(t *testing.T, steps ...step) {
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
	GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	GetEnvOrSkipTest(t, "TEST_METASTORE_ID")
	t.Parallel()
	run(t, steps)
}

type step struct {
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

// Test wrapper over terraform testing framework
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
		"AWS_ATTRIBUTES": awsAttrs,
	}
	ts := []resource.TestStep{}
	ctx := context.Background()
	type testResource struct {
		ID       string
		Name     string
		Resource *schema.Resource
	}

	resourceAndName := regexp.MustCompile(`resource\s+"([^"]*)"\s+"([^"]*)"`)
	resourcesEverCreated := map[testResource]bool{}
	stepConfig := ""
	for i, s := range steps {
		if s.Template != "" {
			stepConfig = environmentTemplate(t, s.Template, vars)
		}
		stepNum := i
		thisStep := s
		stepCallback := thisStep.Callback
		stepCheck := thisStep.Check
		ts = append(ts, resource.TestStep{
			PreConfig: func() {
				if stepConfig == "" {
					return
				}
				logger.Infof(ctx, "Test %s (%s) step %d config is:\n%s",
					t.Name(), cloudEnv, stepNum,
					commands.TrimLeadingWhitespace(stepConfig))
			},
			Config:                    stepConfig,
			Destroy:                   s.Destroy,
			ExpectNonEmptyPlan:        s.ExpectNonEmptyPlan,
			PlanOnly:                  s.PlanOnly,
			PreventDiskCleanup:        s.PreventDiskCleanup,
			PreventPostDestroyRefresh: s.PreventPostDestroyRefresh,
			ImportState:               s.ImportState,
			ImportStateVerify:         s.ImportStateVerify,
			Check: func(state *terraform.State) error {
				// get configured client from provider
				client := provider.Meta().(*common.DatabricksClient)
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
				if stepCallback != nil {
					match := resourceAndName.FindStringSubmatch(stepConfig)
					rootModule := state.RootModule()
					res := rootModule.Resources[match[1]+"."+match[2]]
					id := res.Primary.ID
					return stepCallback(ctx, client, id)
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
	return path.Base(ex) == "__debug_bin"
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

type stdErrLogger struct {
	traceEnabled bool
}

func (l stdErrLogger) Enabled(_ context.Context, level logger.Level) bool {
	return true
}

func (l stdErrLogger) Tracef(_ context.Context, format string, v ...interface{}) {
	if l.traceEnabled {
		fmt.Fprintf(os.Stderr, "[TRACE] "+format+"\n", v...)
	}
}

func (l stdErrLogger) Debugf(_ context.Context, format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "\n[DEBUG] "+format+"\n", v...)
}

func (l stdErrLogger) Infof(_ context.Context, format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "\n[INFO] "+format+"\n", v...)
}

func (l stdErrLogger) Warnf(_ context.Context, format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "\n[WARN] "+format+"\n", v...)
}

func (l stdErrLogger) Errorf(_ context.Context, format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "[ERROR] "+format+"\n", v...)
}
