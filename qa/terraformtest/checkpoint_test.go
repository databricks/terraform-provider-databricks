package terraformtest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"
)

func TestInitialStateReplay(t *testing.T) {
	Run(t, "testdata/initial_state")
}

func TestInitialCheckpointPaths(t *testing.T) {
	for _, tc := range []struct {
		name, value string
	}{
		{"missing", "{}"},
		{"empty", "{checkpoint: ''}"},
		{"absolute", "{checkpoint: /tmp/snapshot.tfstate}"},
		{"old state model", "{resources: {}}"},
		{"unknown field", "{checkpoint: group.tfstate, typo: true}"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := decodeDefinition(strings.NewReader(minimalDefinition + "initial_state: " + tc.value + "\n"))
			require.Error(t, err)
		})
	}
}

func TestLoadInitialCheckpointRelativeToYAML(t *testing.T) {
	directory := t.TempDir()
	contents, err := os.ReadFile("testdata/initial_state/checkpoints/group.tfstate")
	require.NoError(t, err)
	writeDiscoveryFile(t, directory, "nested/snapshots/group.tfstate", string(contents))
	writeDiscoveryFile(t, directory, "nested/test.yaml", minimalDefinition+"initial_state: {checkpoint: snapshots/group.tfstate}\n")
	d, err := loadDefinition(filepath.Join(directory, "nested/test.yaml"))
	require.NoError(t, err)
	require.Equal(t, contents, d.checkpoint)
	require.Equal(t, "snapshots/group.tfstate", d.InitialState.Checkpoint)
}

func TestInitialCheckpointMustBeReadableAndNonempty(t *testing.T) {
	for _, tc := range []struct {
		name, contents string
	}{
		{"missing", ""},
		{"empty", ""},
		{"whitespace", " \n\t"},
		{"directory", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, "test.yaml", minimalDefinition+"initial_state: {checkpoint: group.tfstate}\n")
			switch tc.name {
			case "directory":
				require.NoError(t, os.Mkdir(filepath.Join(directory, "group.tfstate"), 0700))
			case "missing":
			default:
				writeDiscoveryFile(t, directory, "group.tfstate", tc.contents)
			}
			_, err := loadDefinition(filepath.Join(directory, "test.yaml"))
			require.ErrorContains(t, err, "initial_state checkpoint")
		})
	}
}

func TestCheckpointSeedPreservesBytesAndOnlyRunsOnFirstStep(t *testing.T) {
	d, err := loadDefinition("testdata/initial_state/group_update.yaml")
	require.NoError(t, err)
	testCase := resource.TestCase{
		ProtoV6ProviderFactories: providerFactories("http://127.0.0.1"),
		Steps: []resource.TestStep{
			{Config: d.Steps[0].Config},
			{Config: `provider "databricks" {}`},
			{Config: `provider "databricks" {}`, Destroy: true},
		},
	}
	seedInitialCheckpoint(t, d, &testCase)
	first := testCase.Steps[0]
	require.Empty(t, first.Config)
	directory := first.ConfigDirectory.Exec(config.TestStepConfigRequest{})
	contents, err := os.ReadFile(filepath.Join(directory, "terraform.tfstate"))
	require.NoError(t, err)
	require.Equal(t, d.checkpoint, contents)
	for _, later := range testCase.Steps[1:] {
		require.Nil(t, later.ConfigDirectory)
		require.NotEmpty(t, later.Config)
	}
	for _, s := range testCase.Steps {
		require.Contains(t, s.ProtoV6ProviderFactories, "databricks")
	}
	require.Empty(t, testCase.ProtoV6ProviderFactories)
}

func TestCheckpointConfigKeepsExplicitProviderRequirements(t *testing.T) {
	for _, source := range []string{
		`resource "databricks_group" "this" { display_name = "test" }`,
		`provider "databricks" { account_id = "account" }`,
		`terraform {
  required_providers {
    databricks = { source = "databricks/databricks" }
  }
}`,
	} {
		hcl, err := checkpointConfig(source)
		require.NoError(t, err)
		require.Contains(t, hcl, source)
		require.Equal(t, 1, strings.Count(hcl, "terraform {"))
	}
	_, err := checkpointConfig(`resource "databricks_group" {`)
	require.ErrorContains(t, err, "initial_state step configuration")
}

func TestInitialCheckpointRejectsAlternateBackends(t *testing.T) {
	for _, block := range []string{`backend "local" {}`, `backend "remote" {}`, `cloud {}`} {
		_, err := checkpointConfig("terraform {\n" + block + "\n}\n")
		require.ErrorContains(t, err, "initial_state requires the default local backend")
	}
	_, err := checkpointConfig("terraform {}\nterraform {\n  backend \"remote\" {}\n}\n")
	require.ErrorContains(t, err, "initial_state requires the default local backend")
}

func TestInitialStateRequiresDestroyFixtures(t *testing.T) {
	directory := t.TempDir()
	contents, err := os.ReadFile("testdata/initial_state/rejected_update.yaml")
	require.NoError(t, err)
	checkpoint, err := os.ReadFile("testdata/initial_state/checkpoints/group.tfstate")
	require.NoError(t, err)
	writeDiscoveryFile(t, directory, "checkpoints/group.tfstate", string(checkpoint))
	fixture := strings.Split(string(contents), "destroy:\n")[0] + "destroy: {calls: []}\n"
	writeDiscoveryFile(t, directory, "MissingCleanup.yaml", fixture)
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, "step 2 (destroy): unexpected HTTP call")
}
