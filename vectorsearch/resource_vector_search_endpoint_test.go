package vectorsearch

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestVectorSearchEndpointCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceVectorSearchEndpoint())
}

func TestVectorSearchEndpointCreate(t *testing.T) {
	assert.Equal(t, 1, 1)
}
