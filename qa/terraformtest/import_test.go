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
		{"missing persistence", "import: {resource: databricks_group.this, id: group-id}", "import.persist is required"},
		{"null persistence", "import: {resource: databricks_group.this, id: group-id, persist: null}", "import.persist is required"},
		{"invalid persistence", "import: {resource: databricks_group.this, id: group-id, persist: maybe}", "cannot unmarshal"},
		{"verify persistent", "import: {resource: databricks_group.this, id: group-id, persist: true, verify: true}", "import.verify requires persist: false"},
		{"missing ID", "import: {resource: databricks_group.this, persist: true}", "import.id is required"},
		{"blank ID", "import: {resource: databricks_group.this, persist: true, id: ' '}", "import.id is required"},
		{"plan", "import: {resource: databricks_group.this, persist: true, id: group-id}\n    plan: {expect: empty}", "import cannot be combined"},
		{"plan only", "import: {resource: databricks_group.this, persist: true, id: group-id}\n    plan_only: true", "import cannot be combined"},
		{"state", "import: {resource: databricks_group.this, persist: true, id: group-id}\n    state: {}", "import cannot be combined"},
		{"ignored without verification", "import: {resource: databricks_group.this, persist: false, ignore: [name]}", "require import.verify: true"},
		{"identifier without verification", "import: {resource: databricks_group.this, persist: false, identifier_attribute: name}", "require import.verify: true"},
		{"empty ignore", "import: {resource: databricks_group.this, persist: false, verify: true, ignore: ['']}", "nonempty attribute prefixes"},
		{"verify expected failure", "import: {resource: databricks_group.this, persist: false, verify: true}\n    expect_error: rejected", "import.verify cannot be combined"},
		{"unknown field", "import: {resource: databricks_group.this, persist: true, id: group-id, nonsense: true}", "field nonsense not found"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Replace(minimalDefinition, "destroy:", "  - name: import\n    "+tc.extra+"\ndestroy:", 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.message)
		})
	}
	_, err := decodeDefinition(strings.NewReader("version: 1\nsteps: [{name: import, import: {resource: databricks_group.this, persist: true, id: group-id}}]"))
	require.ErrorContains(t, err, "config is required for the first import step")
}

func TestImportRetainsConfigurationForLaterImportsAndCleanup(t *testing.T) {
	persist := true
	d := &testDefinition{Steps: []step{
		{Name: "create", Config: `provider "databricks" {} # applied`},
		{Name: "import", Config: `provider "databricks" {} # import config`, Import: &importStep{Persist: &persist, Resource: "databricks_group.this", ID: "group-id"}},
		{Name: "import again", Import: &importStep{Persist: &persist, Resource: "databricks_group.other", ID: "other-id"}},
	}}
	native := testSteps(t, d, httpfixture.NewHandler())
	require.Contains(t, native[1].Config, "# import config")
	require.Contains(t, native[2].Config, "# import config")
	require.Equal(t, d.Steps[1].Config, native[3].Config)
	require.True(t, native[1].ImportStatePersist)
	require.True(t, native[3].Destroy)
}

func TestImportChecksFixturesWithoutPostApply(t *testing.T) {
	d, err := loadDefinition("testdata/imports/group.yaml")
	require.NoError(t, err)
	native := testSteps(t, d, httpfixture.NewHandler())
	native[0].PreConfig()
	require.Nil(t, native[0].PostApplyFunc)
	require.ErrorContains(t, native[0].ImportStateCheck([]*terraform.InstanceState{}), "step 1 (import): unconsumed call")
}

func TestImportCannotHideMissingOrUnusedFixtures(t *testing.T) {
	for _, missing := range []bool{true, false} {
		name := "unused"
		if missing {
			name = "missing"
		}
		t.Run(name, func(t *testing.T) {
			d, err := loadDefinition("testdata/imports/app_by_name.yaml")
			require.NoError(t, err)
			if missing {
				d.Steps[0].Calls = nil
				pattern := ".*"
				d.Steps[0].ExpectError = &pattern
			} else {
				d.Steps[0].Calls = append(d.Steps[0].Calls, httpfixture.Call{
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
			require.Contains(t, output, "step 1 (import_by_name): "+message)
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

func TestImportRejectsAlreadyManagedAddress(t *testing.T) {
	d, err := loadDefinition("testdata/imports/app_by_name.yaml")
	require.NoError(t, err)
	repeated := d.Steps[0]
	repeated.Name, repeated.Calls = "already_managed", nil
	pattern := "Resource already managed by Terraform"
	repeated.ExpectError = &pattern
	d.Steps = append(d.Steps, repeated)
	run(t, d)
}

func TestImportExpectedAPIError(t *testing.T) {
	d, err := loadDefinition("testdata/imports/app_by_name.yaml")
	require.NoError(t, err)
	d.Steps = d.Steps[:1]
	d.Destroy.Calls = nil
	pattern := "import rejected by fixture"
	d.Steps[0].ExpectError = &pattern
	d.Steps[0].Calls[0].Response = httpfixture.Response{
		Status: 400,
		Body: &httpfixture.Body{JSON: map[string]any{
			"error_code": "INVALID_PARAMETER_VALUE", "message": pattern,
		}},
	}
	run(t, d)
}

func importWithChangedDescription(t *testing.T) *testDefinition {
	t.Helper()
	d, err := loadDefinition("testdata/imports/app_verification.yaml")
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

func TestImportDiscardPreservesLifecycleOutputs(t *testing.T) {
	d, err := loadDefinition("testdata/imports/app_verification.yaml")
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

func TestImportPreservesCleanupConfigurationWhenDiscarded(t *testing.T) {
	persist := false
	d := &testDefinition{Steps: []step{
		{Name: "create", Config: `provider "databricks" {} # applied`},
		{Name: "import", Config: `provider "databricks" {} # import only`, Import: &importStep{Resource: "databricks_group.this", Persist: &persist}},
		{Name: "import again", Import: &importStep{Resource: "databricks_group.this", Persist: &persist}},
	}}
	native := testSteps(t, d, httpfixture.NewHandler())
	require.Contains(t, native[1].Config, "# import only")
	require.Contains(t, native[2].Config, "# applied")
	require.Equal(t, d.Steps[0].Config, native[3].Config)
	require.False(t, native[1].ImportStatePersist)
}

func TestImportDiscardFromInitialCheckpoint(t *testing.T) {
	d, err := loadDefinition("testdata/initial_state/group_update.yaml")
	require.NoError(t, err)
	read := d.Steps[0].Calls[0]
	persist := false
	d.Steps = []step{{
		Name: "import_existing_group", Config: d.Steps[0].Config,
		Import: &importStep{Resource: "databricks_group.this", Persist: &persist},
		Calls:  []httpfixture.Call{read, read},
	}}
	d.Destroy.Calls[0] = read
	require.NoError(t, d.validate())
	run(t, d)
}

func TestImportPreservesExistingResources(t *testing.T) {
	for _, checkpoint := range []bool{false, true} {
		name := "after_apply"
		if checkpoint {
			name = "from_checkpoint"
		}
		t.Run(name, func(t *testing.T) {
			d, err := loadDefinition("testdata/initial_state/group_update.yaml")
			require.NoError(t, err)
			first := d.Steps[0].Calls[0]
			second := httpfixture.Call{
				Request:  httpfixture.Request{Method: "GET", Path: "/api/2.0/preview/scim/v2/Groups/imported-id", Query: first.Request.Query},
				Response: httpfixture.Response{Status: 200, Body: &httpfixture.Body{JSON: map[string]any{"id": "imported-id", "displayName": "Imported group"}}},
			}
			existingConfig := `resource "databricks_group" "this" {
                display_name = "Example group"
            }`
			if !checkpoint {
				existingConfig += `
output "sentinel" { value = 42 }
`
			}
			bothConfig := existingConfig + `
resource "databricks_group" "other" {
    display_name = "Imported group"
}`
			persist := true
			imported := step{Name: "import_another_group", Config: bothConfig,
				Import: &importStep{Resource: "databricks_group.other", ID: "imported-id", Persist: &persist},
				Calls:  []httpfixture.Call{second, second},
			}
			expected := &expectedState{Resources: map[string]expectedResource{}}
			for address, id := range map[string]string{"databricks_group.this": "group-id", "databricks_group.other": "imported-id"} {
				var value yaml.Node
				require.NoError(t, value.Encode(id))
				expected.Resources[address] = expectedResource{Attributes: []expectedAttribute{{Path: []any{"id"}, Equals: value}}}
			}
			// An empty plan also proves that import did not discard prior outputs.
			d.Steps = []step{imported, {Name: "check_both_groups", Config: bothConfig, PlanOnly: true,
				Plan: &expectedPlan{Expect: "empty"}, State: expected, Calls: []httpfixture.Call{first, second},
			}}
			if !checkpoint {
				d.InitialState, d.checkpoint = nil, nil
				create := httpfixture.Call{
					Request: httpfixture.Request{Method: "POST", Path: "/api/2.0/preview/scim/v2/Groups", Body: &httpfixture.Body{JSON: map[string]any{
						"displayName": "Example group", "schemas": []any{"urn:ietf:params:scim:schemas:core:2.0:Group"}, "entitlements": []any{map[string]any{}},
					}}},
					Response: first.Response,
				}
				d.Steps = append([]step{{Name: "create_first_group", Config: existingConfig, Calls: []httpfixture.Call{create, first, first}}}, d.Steps...)
			}
			remove := d.Destroy.Calls[1]
			removeOther := remove
			removeOther.Request.Path = "/api/2.0/preview/scim/v2/Groups/imported-id"
			d.Destroy.Calls = []httpfixture.Call{first, second, remove, removeOther}
			require.NoError(t, d.validate())
			run(t, d)
		})
	}
}

func TestImportCanDiscardWithoutPriorState(t *testing.T) {
	d, err := loadDefinition("testdata/imports/app_by_name.yaml")
	require.NoError(t, err)
	d.Steps = d.Steps[:1]
	persist := false
	d.Steps[0].Import.Persist = &persist
	// The explicitly discarded import adds nothing to the journey's cleanup state.
	d.Destroy.Calls = nil
	run(t, d)
}
