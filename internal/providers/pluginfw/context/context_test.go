package context

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/stretchr/testify/assert"
)

func TestSetUserAgentInResourceContext(t *testing.T) {
	ctx := context.Background()
	resourceKey := "resource"
	resourceName := "test-resource"
	actualContext := SetUserAgentInResourceContext(ctx, resourceName)
	expectedContext := useragent.InContext(ctx, "sdk", "pluginframework")
	expectedContext = useragent.InContext(expectedContext, resourceKey, resourceName)
	assert.Equal(t, expectedContext, actualContext)
}

func TestSetUserAgentInDataSourceContext(t *testing.T) {
	ctx := context.Background()
	dataSourceKey := "data"
	dataSourceName := "test-datasource"
	actualContext := SetUserAgentInDataSourceContext(ctx, dataSourceName)
	expectedContext := useragent.InContext(ctx, "sdk", "pluginframework")
	expectedContext = useragent.InContext(expectedContext, dataSourceKey, dataSourceName)
	assert.Equal(t, expectedContext, actualContext)
}
