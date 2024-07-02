package clusters

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestZones(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.ListZones(mock.Anything).Return(&compute.ListAvailableZonesResponse{
				DefaultZone: "a",
				Zones:       []string{"a", "b"},
			}, nil)
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
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.ListZones(mock.Anything).Return(&compute.ListAvailableZonesResponse{}, fmt.Errorf("missing"))
		},
		Read:        true,
		Resource:    DataSourceClusterZones(),
		NonWritable: true,
		ID:          ".",
	}.ExpectError(t, "missing")
}
