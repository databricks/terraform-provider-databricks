package clusters

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
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

func TestZones_404(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list-zones",
				Status:   404,
				Response: common.NotFound("missing"),
			},
		},
		Read:        true,
		Resource:    DataSourceClusterZones(),
		NonWritable: true,
		ID:          ".",
	}.ExpectError(t, "missing")
}
