package terraformtest

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const minimalDefinition = `version: 1
steps:
  - name: create
    config: |
      resource "databricks_group" "this" {
        display_name = "test"
      }
    calls: []
    state:
      resources:
        databricks_group.this:
          attributes:
            - path: [id]
              equals: group-id
destroy:
  calls: []
`

func TestDecodeDefinition(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(minimalDefinition))
	require.NoError(t, err)
	require.Equal(t, "create", d.Steps[0].Name)
	require.Equal(t, `resource "databricks_group" "this" {
  display_name = "test"
}
`, d.Steps[0].Config)
	attributes := d.Steps[0].State.Resources["databricks_group.this"].Attributes
	require.Len(t, attributes, 1)
	require.Equal(t, []any{"id"}, attributes[0].Path)
	require.Equal(t, "group-id", attributes[0].Equals.Value)
	require.Empty(t, d.Destroy.Calls)
}

func definitionWithAttributes(attributes string) string {
	return strings.Replace(minimalDefinition, "          attributes:\n            - path: [id]\n              equals: group-id\n",
		"          attributes: "+attributes+"\n", 1)
}

func TestDecodeNestedPrimitiveStateChecks(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(definitionWithAttributes(`[
            {path: [settings, enabled], equals: true},
            {path: [items, 0, name], equals: one},
            {path: [optional], equals: null},
            {path: [number], equals: 42},
            {path: [labels, "0"], equals: zero}
          ]`)))
	require.NoError(t, err)
	attributes := d.Steps[0].State.Resources["databricks_group.this"].Attributes
	require.Equal(t, []any{"items", 0, "name"}, attributes[1].Path)
	require.Equal(t, []any{"labels", "0"}, attributes[4].Path)
	check, err := attributes[2].valueCheck()
	require.NoError(t, err)
	require.NoError(t, check.CheckValue(nil))
	require.Error(t, check.CheckValue("set"))
	require.NoError(t, (stateAssertion{expected: *d.Steps[0].State}).check(groupState()))
}

func TestRejectInvalidStateAttributeChecks(t *testing.T) {
	for _, tc := range []struct {
		name, attributes, message string
	}{
		{"missing path", "[{equals: true}]", "path must start with an attribute name"},
		{"empty path", "[{path: [], equals: true}]", "path must start with an attribute name"},
		{"empty root key", "[{path: [''], equals: true}]", "path must start with an attribute name"},
		{"root index", "[{path: [0], equals: true}]", "path must start with an attribute name"},
		{"negative index", "[{path: [items, -1], equals: true}]", "path indices must be nonnegative"},
		{"fractional index", "[{path: [items, 0.5], equals: true}]", "path segments must be string keys"},
		{"boolean index", "[{path: [items, true], equals: true}]", "path segments must be string keys"},
		{"null index", "[{path: [items, null], equals: true}]", "path segments must be string keys"},
		{"overflow index", "[{path: [items, 18446744073709551615], equals: true}]", "path segments must be string keys"},
		{"object segment", "[{path: [items, {}], equals: true}]", "path segments must be string keys"},
		{"missing equals", "[{path: [id]}]", "equals is required"},
		{"object expectation", "[{path: [settings], equals: {enabled: true}}]", "collection assertions are not supported"},
		{"list expectation", "[{path: [members], equals: [one, two]}]", "collection assertions are not supported"},
		{"infinity", "[{path: [number], equals: .inf}]", "equals must be a JSON primitive or null"},
		{"NaN", "[{path: [number], equals: .nan}]", "equals must be a JSON primitive or null"},
		{"timestamp", "[{path: [id], equals: 2026-10-12}]", "equals must be a JSON primitive or null"},
		{"misspelled equals", "[{path: [id], equal: group-id}]", "field equal not found"},
		{"misspelled path", "[{paths: [id], equals: group-id}]", "field paths not found"},
		{"old attribute map", "{id: group-id}", "cannot unmarshal"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := decodeDefinition(strings.NewReader(definitionWithAttributes(tc.attributes)))
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestStateChecksValidateLaterStepsBeforeTerraform(t *testing.T) {
	input := strings.Replace(minimalDefinition, "destroy:\n", `  - name: update
    config: |
      provider "databricks" {}
    state:
      resources:
        databricks_group.this:
          attributes:
            - path: [members, -1]
              equals: one
destroy:
`, 1)
	_, err := decodeDefinition(strings.NewReader(input))
	require.ErrorContains(t, err, "step 2 (update): databricks_group.this: attribute check 1: path indices must be nonnegative")
}

func TestDecodeStateNumbersPreservesPrecisionAndTypes(t *testing.T) {
	for _, tc := range []struct {
		name, expected, actual string
	}{
		{"large integer", "9007199254740993", "9007199254740993"},
		{"uint64", "18446744073709551615", "18446744073709551615"},
		{"fraction", "1.2", "1.2"},
		{"exponent", "1e20", "100000000000000000000"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := decodeDefinition(strings.NewReader(definitionWithAttributes("[{path: [number], equals: " + tc.expected + "}]")))
			require.NoError(t, err)
			state := groupState()
			state.Values.RootModule.Resources[0].AttributeValues["number"] = json.Number(tc.actual)
			require.NoError(t, (stateAssertion{expected: *d.Steps[0].State}).check(state))
			state.Values.RootModule.Resources[0].AttributeValues["number"] = tc.actual
			require.ErrorContains(t, (stateAssertion{expected: *d.Steps[0].State}).check(state), "expected json.Number value")
		})
	}
}

func TestRejectInvalidDefinition(t *testing.T) {
	for _, tc := range []struct {
		name, yaml, message string
	}{
		{"version", strings.Replace(minimalDefinition, "version: 1", "version: 2", 1), "unsupported version 2"},
		{"no steps", "version: 1\nsteps: []", "at least one step"},
		{"unnamed step", strings.Replace(minimalDefinition, "name: create", "name: ''", 1), "name and config"},
		{"no HCL", "version: 1\nsteps: [{name: create}]", "name and config"},
		{"old HCL field", strings.Replace(minimalDefinition, "config:", "terraform:", 1), "field terraform not found"},
		{"multiple documents", minimalDefinition + "---\nversion: 1\n", "single YAML document"},
		{"empty resource address", strings.Replace(minimalDefinition, "databricks_group.this:", "'':", 1), "state resource address must not be empty"},
		{"destroy state override", minimalDefinition + "  state: {resources: {}}\n", "field state not found"},
		{"nonempty destroy override", minimalDefinition + "  state: {resources: {databricks_group.this: {}}}\n", "field state not found"},
		{"invalid second step", `version: 1
steps:
  - name: create
    config: |
      provider "databricks" {}
  - name: update
    config: |
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

func TestDecodeOptionalStateChecks(t *testing.T) {
	for _, tc := range []struct {
		name, expectation string
	}{
		{"omitted", ""},
		{"empty state", "    state: {}\n"},
		{"empty resources", "    state: {resources: {}}\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := decodeDefinition(strings.NewReader(`version: 1
steps:
  - name: read
    config: |
      provider "databricks" {}
` + tc.expectation))
			require.NoError(t, err)
			if tc.expectation == "" {
				require.Nil(t, d.Steps[0].State)
			} else {
				require.NotNil(t, d.Steps[0].State)
				require.Empty(t, d.Steps[0].State.Resources)
			}
		})
	}
}

func TestRejectDeferredFeatures(t *testing.T) {
	for _, field := range []string{"recording", "environment_template", "initial_state_template", "force_order", "strict"} {
		t.Run(field, func(t *testing.T) {
			_, err := decodeDefinition(strings.NewReader(minimalDefinition + field + ": {}\n"))
			require.ErrorContains(t, err, "field "+field+" not found")
		})
	}
}

func TestDecodeHTTPFixtureAndMetadata(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(`version: 1
host_metadata: {workspace_id: "12345"}
steps:
  - name: read
    config: |
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
    config: |
      provider "databricks" {}
`))
	require.NoError(t, err)
	require.Nil(t, d.Steps[0].State)
	_, err = decodeDefinition(strings.NewReader(`version: 1
steps:
  - name: read
    config: |
      provider "databricks" {}
    calls:
      - request: {method: GET, path: /groups, typo: true}
        response: {status: 200}
`))
	require.ErrorContains(t, err, "field typo not found")
}
