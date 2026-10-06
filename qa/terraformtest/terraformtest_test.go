package terraformtest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type recordingFixturePhases struct {
	phases [][]qa.HTTPFixture
}

func (r *recordingFixturePhases) SetFixtures(fixtures []qa.HTTPFixture) {
	r.phases = append(r.phases, fixtures)
}

func TestTerraformStepsKeepFixturesWithTheirPhase(t *testing.T) {
	firstFixtures := []qa.HTTPFixture{{Method: "POST", Resource: "/first"}}
	secondFixtures := []qa.HTTPFixture{{Method: "PATCH", Resource: "/second"}}
	destroyFixtures := []qa.HTTPFixture{{Method: "DELETE", Resource: "/second"}}
	testCase := TestCase{
		Steps: []Step{
			{Name: "create", Template: "first"},
			{Name: "update", Template: "second"},
		},
	}
	phases := &recordingFixturePhases{}

	steps := terraformSteps(
		testCase,
		[][]qa.HTTPFixture{firstFixtures, secondFixtures, destroyFixtures},
		phases,
		nil,
	)

	require.Len(t, steps, 3)
	assert.Equal(t, "first", steps[0].Config)
	assert.Equal(t, "second", steps[1].Config)
	assert.Equal(t, "second", steps[2].Config)
	assert.True(t, steps[2].Destroy)

	for _, step := range steps {
		step.PreConfig()
	}
	require.Len(t, phases.phases, 3)
	assert.Equal(t, firstFixtures, phases.phases[0])
	assert.Equal(t, secondFixtures, phases.phases[1])
	assert.Equal(t, destroyFixtures, phases.phases[2])
}

func TestTerraformStepsWithoutConfigurationsHaveNoDestroyStep(t *testing.T) {
	steps := terraformSteps(TestCase{}, nil, &recordingFixturePhases{}, nil)
	assert.Empty(t, steps)
}

func TestTerraformStepsForwardStateAndPlanChecks(t *testing.T) {
	check := resource.TestCheckResourceAttr("databricks_user.this", "display_name", "Fixture User")
	planChecks := resource.ConfigPlanChecks{
		PreApply: []plancheck.PlanCheck{
			plancheck.ExpectResourceAction("databricks_user.this", plancheck.ResourceActionNoop),
		},
	}

	steps := terraformSteps(TestCase{Steps: []Step{{
		Name:             "check",
		Template:         "resource {}",
		Check:            check,
		ConfigPlanChecks: planChecks,
	}}}, make([][]qa.HTTPFixture, 2), &recordingFixturePhases{}, nil)

	require.Len(t, steps, 2)
	require.NotNil(t, steps[0].Check)
	assert.Equal(t, planChecks, steps[0].ConfigPlanChecks)
}

func TestLoadFixtureFileGroupsExchangesByNamedStep(t *testing.T) {
	path := filepath.Join(t.TempDir(), "fixtures.jsonl")
	contents := []byte(
		`{"version":1,"step_index":0,"step_name":"create","sequence":0,` +
			`"request":{"method":"POST","path":"/users","headers":{"content-type":"application/json"},"body":{"name":"Fixture"}},` +
			`"response":{"status":201,"headers":{},"body":{"id":"user-id"}}}` + "\n" +
			`{"version":1,"step_index":1,"step_name":"destroy","sequence":0,` +
			`"request":{"method":"DELETE","path":"/users/user-id","headers":{},"body":null},` +
			`"response":{"status":204,"headers":{},"body":null}}` + "\n",
	)
	require.NoError(t, os.WriteFile(path, contents, 0o600))

	fixtures, err := loadFixtureFile(path, []Step{{Name: "create"}})

	require.NoError(t, err)
	require.Len(t, fixtures, 2)
	require.Len(t, fixtures[0], 1)
	assert.Equal(t, "POST", fixtures[0][0].Method)
	assert.Equal(t, "/users", fixtures[0][0].Resource)
	assert.JSONEq(t, `{"name":"Fixture"}`, string(fixtures[0][0].ExpectedRequest.(json.RawMessage)))
	require.Len(t, fixtures[1], 1)
	assert.Equal(t, "DELETE", fixtures[1][0].Method)
	assert.Nil(t, fixtures[1][0].ExpectedRequest)
}

func TestHostMetadataFixturesAreReusableInEveryPhase(t *testing.T) {
	resourceFixture := qa.HTTPFixture{Method: "POST", Resource: "/users"}
	phases := [][]qa.HTTPFixture{{resourceFixture}, nil}

	result := withHostMetadataFixtures(phases, WorkspaceHostMetadata("12345"), "http://fixture.test/oidc")

	require.Len(t, result, 2)
	for _, phase := range result {
		require.NotEmpty(t, phase)
		metadataFixture := phase[0]
		assert.Equal(t, "GET", metadataFixture.Method)
		assert.Equal(t, "/.well-known/databricks-config", metadataFixture.Resource)
		assert.True(t, metadataFixture.ReuseRequest)
		metadata := metadataFixture.Response.(config.HostMetadata)
		assert.Equal(t, "12345", metadata.WorkspaceID)
		assert.Equal(t, "http://fixture.test/oidc", metadata.OIDCEndpoint)
	}
	assert.Equal(t, resourceFixture, result[0][1])
}

func TestFixtureStepNamesRejectDuplicateAndReservedNames(t *testing.T) {
	_, err := fixtureStepNames([]Step{{Name: "same"}, {Name: "same"}})
	require.EqualError(t, err, `step 1 repeats name "same"`)

	_, err = fixtureStepNames([]Step{{Name: destroyStepName}})
	require.EqualError(t, err, `step 0 uses reserved name "destroy"`)
}
