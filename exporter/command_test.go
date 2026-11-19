package exporter

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

type dummyReader string

func (d dummyReader) Read(p []byte) (int, error) {
	n := copy(p, []byte(d))
	return n, nil
}

// isolateDatabricksEnv makes sure that the test isn't affected by DATABRICKS_* environment
// variables or ~/.databrickscfg of the developer running it.
func isolateDatabricksEnv(t *testing.T) {
	for _, kv := range os.Environ() {
		if k, _, _ := strings.Cut(kv, "="); strings.HasPrefix(k, "DATABRICKS_") {
			t.Setenv(k, "")
		}
	}
	t.Setenv("HOME", t.TempDir())
}

func TestInteractivePrompts(t *testing.T) {
	isolateDatabricksEnv(t)
	originalInput := cliInput
	originalOutput := cliOutput
	t.Cleanup(func() {
		cliInput = originalInput
		cliOutput = originalOutput
	})

	cliInput = dummyReader("y\n")
	cliOutput = &bytes.Buffer{}
	ic := &importContext{
		Client: &common.DatabricksClient{
			DatabricksClient: &client.DatabricksClient{
				Config: &config.Config{},
			},
		},
		Context: context.Background(),
		Importables: map[string]importable{
			"x": {
				Service: "a",
				List: func(_ *importContext) error {
					return nil
				},
			},
			"y": {
				Service: "mounts",
				List: func(_ *importContext) error {
					return nil
				},
			},
		},
	}
	services, err := ic.interactivePrompts()
	assert.NoError(t, err)
	assert.Equal(t, "y", ic.match)
	assert.True(t, ic.mounts)
	assert.Equal(t, "a,mounts", services)
}

func TestInteractivePromptsStopsAfterFailedAuthAttempts(t *testing.T) {
	isolateDatabricksEnv(t)
	// a missing profile can't be fixed by entering host & token
	cfgFile := filepath.Join(t.TempDir(), ".databrickscfg")
	assert.NoError(t, os.WriteFile(cfgFile, []byte("[DEFAULT]\n"), 0600))
	t.Setenv("DATABRICKS_CONFIG_FILE", cfgFile)
	t.Setenv("DATABRICKS_CONFIG_PROFILE", "doesnotexist")
	originalInput := cliInput
	originalOutput := cliOutput
	t.Cleanup(func() {
		cliInput = originalInput
		cliOutput = originalOutput
	})

	cliInput = dummyReader("y\n")
	cliOutput = &bytes.Buffer{}
	ic := &importContext{
		Client: &common.DatabricksClient{
			DatabricksClient: &client.DatabricksClient{
				Config: &config.Config{},
			},
		},
		Context: context.Background(),
	}
	_, err := ic.interactivePrompts()
	assert.ErrorContains(t, err, "can't authenticate after 3 attempts")
}

func TestRunSkipsInteractivePromptsWhenServicesOrListingIsConfigured(t *testing.T) {
	isolateDatabricksEnv(t)
	originalInput := cliInput
	originalOutput := cliOutput
	t.Cleanup(func() {
		cliInput = originalInput
		cliOutput = originalOutput
	})

	for _, tc := range []struct {
		name string
		args []string
	}{
		{
			name: "services",
			args: []string{"-services", "groups,users"},
		},
		{
			name: "listing",
			args: []string{"-listing", "groups"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			cliInput = dummyReader("https://example.com\n")
			cliOutput = output

			args := []string{"-directory", t.TempDir(), "-targetCloud", "invalid"}
			args = append(args, tc.args...)

			err := Run(args...)

			assert.EqualError(t, err, "invalid targetCloud value: invalid. Must be one of: aws, azure, gcp")
			assert.Empty(t, output.String())
		})
	}
}
