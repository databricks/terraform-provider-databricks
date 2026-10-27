package terraformtest

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestImportReplay(t *testing.T) {
	Run(t, "testdata/imports")
}

func TestImportValidation(t *testing.T) {
	for _, tc := range []struct{ name, extra, message string }{
		{"missing resource", "import: {}", "import.resource is required"},
		{"plan", "import: {resource: databricks_group.this}\n    plan: {expect: empty}", "import cannot be combined"},
		{"plan only", "import: {resource: databricks_group.this}\n    plan_only: true", "import cannot be combined"},
		{"state", "import: {resource: databricks_group.this}\n    state: {}", "import cannot be combined"},
		{"ignored without verification", "import: {resource: databricks_group.this, ignore: [name]}", "require import.verify: true"},
		{"identifier without verification", "import: {resource: databricks_group.this, identifier_attribute: name}", "require import.verify: true"},
		{"empty ignore", "import: {resource: databricks_group.this, verify: true, ignore: ['']}", "nonempty attribute prefixes"},
		{"verify expected failure", "import: {resource: databricks_group.this, verify: true}\n    expect_error: rejected", "import.verify cannot be combined"},
		{"unknown field", "import: {resource: databricks_group.this, verfy: true}", "field verfy not found"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Replace(minimalDefinition, "destroy:", "  - name: import\n    "+tc.extra+"\ndestroy:", 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.message)
		})
	}
	_, err := decodeDefinition(strings.NewReader("version: 1\nsteps: [{name: import, import: {resource: databricks_group.this}}]"))
	require.ErrorContains(t, err, "prior successful apply step or initial_state")
}

func TestImportPreservesCleanupConfiguration(t *testing.T) {
	d := &testDefinition{Steps: []step{
		{Name: "create", Config: `provider "databricks" {} # applied`},
		{Name: "import", Config: `provider "databricks" {} # import only`, Import: &importStep{Resource: "databricks_group.this"}},
		{Name: "import again", Import: &importStep{Resource: "databricks_group.this"}},
	}}
	native := testSteps(t, d, httpfixture.NewHandler())
	require.Contains(t, native[1].Config, "# import only")
	require.Contains(t, native[2].Config, "# applied")
	require.Equal(t, d.Steps[0].Config, native[3].Config)
	require.False(t, native[1].ImportStatePersist)
	require.True(t, native[3].Destroy)
}

func TestImportChecksFixturesWithoutPostApply(t *testing.T) {
	d, err := loadDefinition("testdata/imports/group.yaml")
	require.NoError(t, err)
	native := testSteps(t, d, httpfixture.NewHandler())
	native[1].PreConfig()
	require.Nil(t, native[1].PostApplyFunc)
	require.ErrorContains(t, native[1].ImportStateCheck([]*terraform.InstanceState{}), "step 2 (import): unconsumed call")
}

func importWithChangedDescription(t *testing.T) *testDefinition {
	t.Helper()
	d, err := loadDefinition("testdata/imports/app_by_name.yaml")
	require.NoError(t, err)
	body := d.Steps[1].Calls[0].Response.Body.JSON.(map[string]any)
	body["description"] = "Different description"
	return d
}

func TestImportVerificationRejectsChangedState(t *testing.T) {
	d := importWithChangedDescription(t)
	directory := t.TempDir()
	contents, err := yaml.Marshal(d)
	require.NoError(t, err)
	writeDiscoveryFile(t, directory, "Mismatch.yaml", string(contents))
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, "ImportStateVerify")
	require.Contains(t, output, "description")
}

func TestImportCanIgnoreSelectedAttributes(t *testing.T) {
	d := importWithChangedDescription(t)
	d.Steps[1].Import.Ignore = append(d.Steps[1].Import.Ignore, "description")
	run(t, d)
}

func TestImportCannotHideMissingOrUnusedFixtures(t *testing.T) {
	for _, missing := range []bool{true, false} {
		name := "unused"
		if missing {
			name = "missing"
		}
		t.Run(name, func(t *testing.T) {
			d, err := loadDefinition("testdata/imports/group.yaml")
			require.NoError(t, err)
			if missing {
				d.Steps[1].Calls = nil
				d.Steps[1].Import.Verify = false
				pattern := ".*"
				d.Steps[1].ExpectError = &pattern
			} else {
				d.Steps[1].Calls = append(d.Steps[1].Calls, httpfixture.Call{
					Request:  httpfixture.Request{Method: "GET", Path: "/unused"},
					Response: httpfixture.Response{Status: 200},
				})
			}
			directory := t.TempDir()
			contents, err := yaml.Marshal(d)
			require.NoError(t, err)
			writeDiscoveryFile(t, directory, "Fixtures.yaml", string(contents))
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			require.Error(t, err, output)
			message := "unconsumed call"
			if missing {
				message = "unexpected HTTP call"
			}
			require.Contains(t, output, "step 2 (import): "+message)
		})
	}
}

func TestImportWithLocalFiles(t *testing.T) {
	d, err := loadDefinition("testdata/imports/group.yaml")
	require.NoError(t, err)
	d.Steps[0].Config = `resource "databricks_group" "this" {
        display_name = trimspace(file("${path.module}/name.txt"))
    }`
	d.Steps[0].fileContents = map[string][]byte{"name.txt": []byte("Example group\n")}
	run(t, d)
}

func TestImportFromInitialCheckpoint(t *testing.T) {
	d, err := loadDefinition("testdata/initial_state/group_update.yaml")
	require.NoError(t, err)
	read := d.Steps[0].Calls[0]
	d.Steps = []step{{
		Name: "import_existing_group", Config: d.Steps[0].Config,
		Import: &importStep{Resource: "databricks_group.this"},
		Calls:  []httpfixture.Call{read, read},
	}}
	d.Destroy.Calls[0] = read
	require.NoError(t, d.validate())
	run(t, d)
}

func TestImportPreservesLifecycleOutputs(t *testing.T) {
	d, err := loadDefinition("testdata/imports/group.yaml")
	require.NoError(t, err)
	d.Steps[0].Config += `
output "sentinel" { value = 42 }
`
	d.Steps = append(d.Steps, step{
		Name: "unchanged_after_import", Config: d.Steps[0].Config,
		PlanOnly: true, Plan: &expectedPlan{Expect: "empty"},
		Calls: []httpfixture.Call{d.Steps[0].Calls[1]},
	})
	run(t, d)
}

func TestImportExpectedAPIError(t *testing.T) {
	d, err := loadDefinition("testdata/imports/app_by_name.yaml")
	require.NoError(t, err)
	pattern := "import rejected by fixture"
	d.Steps[1].ExpectError = &pattern
	d.Steps[1].Import.Verify = false
	d.Steps[1].Import.Ignore = nil
	d.Steps[1].Import.IdentifierAttribute = ""
	d.Steps[1].Calls[0].Response = httpfixture.Response{
		Status: 400,
		Body: &httpfixture.Body{JSON: map[string]any{
			"error_code": "INVALID_PARAMETER_VALUE", "message": pattern,
		}},
	}
	run(t, d)
}
