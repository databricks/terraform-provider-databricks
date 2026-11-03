package terraformtest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestOrderedLifecycleReplay(t *testing.T) {
	d, err := loadDefinition("testdata/initial_state/group_update.yaml")
	require.NoError(t, err)
	d.EnforceCallOrder = true
	run(t, d)
}

func TestDecodeCallOrder(t *testing.T) {
	for _, setting := range []string{"", "enforce_call_order: false\n", "enforce_call_order: true\n"} {
		d, err := decodeDefinition(strings.NewReader(minimalDefinition + setting))
		require.NoError(t, err)
		require.Equal(t, strings.Contains(setting, "true"), d.EnforceCallOrder)
	}
	_, err := decodeDefinition(strings.NewReader(minimalDefinition + "enforce_call_order: {}\n"))
	require.Error(t, err)
}

func TestCallOrderIsOptionalAndIncludesDestroy(t *testing.T) {
	for _, tc := range []struct {
		name    string
		ordered bool
		destroy bool
	}{
		{"unordered independent reads", false, false},
		{"ordered independent reads", true, false},
		{"ordered destroy", true, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := loadDefinition("testdata/initial_state/group_update.yaml")
			require.NoError(t, err)
			d.EnforceCallOrder = tc.ordered
			if tc.destroy {
				d.Destroy.Calls[0], d.Destroy.Calls[1] = d.Destroy.Calls[1], d.Destroy.Calls[0]
			} else {
				d.Steps[0].Calls[0], d.Steps[0].Calls[1] = d.Steps[0].Calls[1], d.Steps[0].Calls[0]
			}
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, d.InitialState.Checkpoint, string(d.checkpoint))
			contents, err := yaml.Marshal(d)
			require.NoError(t, err)
			writeDiscoveryFile(t, directory, "Order.yaml", string(contents))
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			if !tc.ordered {
				require.NoError(t, err, output)
				return
			}
			require.Error(t, err, output)
			require.Contains(t, output, "out-of-order HTTP call")
			if tc.destroy {
				require.Contains(t, output, "step 2 (destroy): out-of-order HTTP call")
			}
		})
	}
}
