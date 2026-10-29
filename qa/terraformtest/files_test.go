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

func TestFilePayloadReplay(t *testing.T) {
	Run(t, "testdata/files")
}

func TestLoadStepFilesRelativeToYAML(t *testing.T) {
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "nested/assets/script.py", "# comment\n")
	writeDiscoveryFile(t, directory, "nested/assets/binary", "\x00\xff\n")
	writeDiscoveryFile(t, directory, "nested/assets/empty", "")
	definition := strings.Replace(minimalDefinition, "    calls: []", "    files: {script.py: assets/script.py, binary: assets/binary, empty: assets/empty}\n    calls: []", 1)
	writeDiscoveryFile(t, directory, "nested/test.yaml", definition)
	d, err := loadDefinition(filepath.Join(directory, "nested/test.yaml"))
	require.NoError(t, err)
	require.Equal(t, []byte("# comment\n"), d.Steps[0].fileContents["script.py"])
	require.Equal(t, []byte{0, 255, '\n'}, d.Steps[0].fileContents["binary"])
	require.Empty(t, d.Steps[0].fileContents["empty"])
	require.Contains(t, d.Steps[0].fileContents, "empty")
}

func TestMissingLaterStepFileFailsBeforeReplay(t *testing.T) {
	directory := t.TempDir()
	definition := strings.Replace(minimalDefinition, "destroy:", `  - name: update
    config: |
      provider "databricks" {}
    files: {script.py: missing.py}
    calls: []
destroy:`, 1)
	writeDiscoveryFile(t, directory, "test.yaml", definition)
	_, err := loadDefinition(filepath.Join(directory, "test.yaml"))
	require.ErrorContains(t, err, `step 2 (update) files["script.py"]`)
	require.ErrorContains(t, err, "missing.py")
}

func TestStepFilesRejectUnsafeDestinations(t *testing.T) {
	for _, name := range []string{"", ".", "..", "nested/script.py", "../script.py", "/script.py", `nested\script.py`, "main.tf", "config.tf.json", "values.tfvars", "values.tfvars.json", "terraform.tfstate", ".terraform.lock.hcl", "terraform_plugin_test.tf"} {
		t.Run(name, func(t *testing.T) {
			require.Error(t, validateFiles(map[string]string{name: "assets/script.py"}))
		})
	}
	for _, source := range []string{"", " ", "/tmp/script.py"} {
		require.ErrorContains(t, validateFiles(map[string]string{"script.py": source}), "relative to the YAML")
	}
	require.NoError(t, validateFiles(map[string]string{"script.py": "assets/script.py", "data.json": "../shared/data.json"}))
}

func TestStepFilesPersistAndCanBeReplaced(t *testing.T) {
	definition := &testDefinition{Steps: []step{
		{fileContents: map[string][]byte{"script.py": []byte("# comment\n"), "data.json": []byte("{}")}},
		{fileContents: map[string][]byte{"script.py": []byte("# updated\n")}},
	}}
	testCase := resource.TestCase{ProtoV6ProviderFactories: ProviderFactories(fixtureConfig("http://127.0.0.1")), Steps: []resource.TestStep{
		{Config: `provider "databricks" {}`},
		{Config: `provider "databricks" {}`},
		{Config: `provider "databricks" {}`, Destroy: true},
	}}
	configureStepFiles(t, definition, &testCase)
	for i, s := range testCase.Steps {
		directory := s.ConfigDirectory.Exec(config.TestStepConfigRequest{})
		script, err := os.ReadFile(filepath.Join(directory, "script.py"))
		require.NoError(t, err)
		expected := "# updated\n"
		if i == 0 {
			expected = "# comment\n"
		}
		require.Equal(t, expected, string(script))
		data, err := os.ReadFile(filepath.Join(directory, "data.json"))
		require.NoError(t, err)
		require.Equal(t, "{}", string(data))
		configuration, err := os.ReadFile(filepath.Join(directory, "main.tf"))
		require.NoError(t, err)
		require.Contains(t, string(configuration), `databricks/databricks`)
		require.Empty(t, s.Config)
		require.Contains(t, s.ProtoV6ProviderFactories, "databricks")
	}
	require.Nil(t, testCase.ProtoV6ProviderFactories)
}

func TestFilesPreserveInitialCheckpoint(t *testing.T) {
	d, err := loadDefinition("testdata/initial_state/group_update.yaml")
	require.NoError(t, err)
	d.Steps[0].fileContents = map[string][]byte{"script.py": []byte("# comment\n")}
	testCase := resource.TestCase{ProtoV6ProviderFactories: ProviderFactories(fixtureConfig("http://127.0.0.1")), Steps: []resource.TestStep{
		{Config: d.Steps[0].Config},
		{Config: d.Steps[0].Config, Destroy: true},
	}}
	seedInitialCheckpoint(t, d, &testCase)
	configureStepFiles(t, d, &testCase)
	for i, s := range testCase.Steps {
		directory := s.ConfigDirectory.Exec(config.TestStepConfigRequest{})
		script, err := os.ReadFile(filepath.Join(directory, "script.py"))
		require.NoError(t, err)
		require.Equal(t, "# comment\n", string(script))
		checkpoint, err := os.ReadFile(filepath.Join(directory, "terraform.tfstate"))
		if i == 0 {
			require.NoError(t, err)
			require.Equal(t, d.checkpoint, checkpoint)
		} else {
			require.True(t, os.IsNotExist(err))
		}
		require.Contains(t, s.ProtoV6ProviderFactories, "databricks")
	}
}
