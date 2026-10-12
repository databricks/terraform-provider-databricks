package terraformtest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const minimalDefinition = `version: 1
steps:
  - name: create
    terraform: |
      resource "databricks_group" "this" {
        display_name = "test"
      }
    calls: []
    state:
      resources:
        databricks_group.this:
          attributes: {id: group-id}
destroy:
  calls: []
  state: {resources: {}}
`

func TestDecodeDefinition(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(minimalDefinition))
	require.NoError(t, err)
	require.Equal(t, "create", d.Steps[0].Name)
	require.Equal(t, `resource "databricks_group" "this" {
  display_name = "test"
}
`, d.Steps[0].Terraform)
	require.Equal(t, "group-id", d.Steps[0].State.Resources["databricks_group.this"].Attributes["id"])
	require.Empty(t, d.Destroy.State.Resources)
}

func TestRejectInvalidDefinition(t *testing.T) {
	for _, tc := range []struct {
		name, yaml, message string
	}{
		{"version", strings.Replace(minimalDefinition, "version: 1", "version: 2", 1), "unsupported version 2"},
		{"no steps", "version: 1\nsteps: []", "at least one step"},
		{"unnamed step", strings.Replace(minimalDefinition, "name: create", "name: ''", 1), "name and terraform"},
		{"no HCL", "version: 1\nsteps: [{name: create}]", "name and terraform"},
		{"multiple documents", minimalDefinition + "---\nversion: 1\n", "single YAML document"},
		{"missing inventory", strings.Replace(minimalDefinition, "state: {resources: {}}", "state: {}", 1), "state.resources"},
		{"nonempty destroy", strings.Replace(minimalDefinition, "state: {resources: {}}", "state: {resources: {databricks_group.this: {}}}", 1), "destroy must expect empty state"},
		{"invalid second step", `version: 1
steps:
  - name: create
    terraform: |
      provider "databricks" {}
  - name: update
    terraform: |
      provider "databricks" {}
    calls:
      - request: {method: GET, path: /groups}
        response: {status: 0}
`, "step 2 (update): call 1: invalid response status"},
		{"invalid destroy", minimalDefinition + "  extra: true\n", "field extra not found"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := decodeDefinition(strings.NewReader(tc.yaml))
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestRejectDeferredFeatures(t *testing.T) {
	for _, field := range []string{"initial_state", "recording", "environment_template", "initial_state_template", "force_order", "strict"} {
		t.Run(field, func(t *testing.T) {
			_, err := decodeDefinition(strings.NewReader(minimalDefinition + field + ": {}\n"))
			require.ErrorContains(t, err, "field "+field+" not found")
		})
	}
	for _, field := range []string{"plan", "plan_only", "expect_error"} {
		t.Run(field, func(t *testing.T) {
			input := strings.Replace(minimalDefinition, "    calls: []", "    "+field+": {}\n    calls: []", 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, "field "+field+" not found")
		})
	}
}

func TestDecodeHTTPFixtureAndMetadata(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(`version: 1
host_metadata: {workspace_id: "12345"}
steps:
  - name: read
    terraform: |
      provider "databricks" {}
    calls:
      - request:
          method: GET
          path: /groups
          query: {attributes: [displayName]}
          body: null
        response:
          status: 200
          body:
            json: {id: group-id}
`))
	require.NoError(t, err)
	require.Equal(t, []string{"displayName"}, d.Steps[0].Calls[0].Request.Query["attributes"])
	require.Nil(t, d.Steps[0].Calls[0].Request.Body)
	require.Equal(t, map[string]any{"id": "group-id"}, d.Steps[0].Calls[0].Response.Body.JSON)
	metadata := d.metadataCalls()
	require.Len(t, metadata, 1)
	require.Equal(t, "/.well-known/databricks-config", metadata[0].Request.Path)
	require.Equal(t, map[string]any{"workspace_id": "12345"}, metadata[0].Response.Body.JSON)
}

func TestOptionalStateAndMalformedFixtureFields(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(`version: 1
steps:
  - name: read
    terraform: |
      provider "databricks" {}
`))
	require.NoError(t, err)
	require.Nil(t, d.Steps[0].State)
	_, err = decodeDefinition(strings.NewReader(`version: 1
steps:
  - name: read
    terraform: |
      provider "databricks" {}
    calls:
      - request: {method: GET, path: /groups, typo: true}
        response: {status: 200}
`))
	require.ErrorContains(t, err, "field typo not found")
}
