package compute

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestZones(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list-zones",
				Response: ZonesInfo{
					DefaultZone: "a",
					Zones:       []string{"a", "b"},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceClusterZones(),
		NonWritable: true,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "a", d.Get("default_zone"))
	assert.Equal(t, 2, d.Get("zones.#"))
}
