package exporter

import (
	"bytes"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

type dummyReader string

func (d dummyReader) Read(p []byte) (int, error) {
	n := copy(p, []byte(d))
	return n, nil
}

func TestInteractivePrompts(t *testing.T) {
	cliInput = dummyReader("y\n")
	cliOutput = &bytes.Buffer{}
	ic := &importContext{
		Client: &common.DatabricksClient{},
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
	ic.interactivePrompts()
	assert.Equal(t, "a,mounts", ic.listing)
	assert.Equal(t, "y", ic.match)
	assert.True(t, ic.mounts)
}
