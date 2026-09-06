package exporter

import (
	"bytes"
	"context"
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

func TestInteractivePrompts(t *testing.T) {
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
	services := ic.interactivePrompts()
	assert.Equal(t, "y", ic.match)
	assert.True(t, ic.mounts)
	assert.Equal(t, "a,mounts", services)
}

func TestRunSkipsInteractivePromptsWhenServicesOrListingIsConfigured(t *testing.T) {
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
