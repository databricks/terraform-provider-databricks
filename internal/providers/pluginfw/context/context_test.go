package common

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/stretchr/testify/assert"
)

func TestSetResourceNameInContext(t *testing.T) {
	ctx := context.Background()
	resourceKey := "resource"
	resourceName := "test-resource"
	actualContext := SetResourceNameInContext(ctx, resourceName)
	expectedContext := useragent.InContext(ctx, resourceKey, resourceName)
	assert.Equal(t, expectedContext, actualContext)
}

func TestSetDataSourceNameInContext(t *testing.T) {
	ctx := context.Background()
	dataSourceKey := "data"
	dataSourceName := "test-datasource"
	actualContext := SetDataSourceNameInContext(ctx, dataSourceName)
	expectedContext := useragent.InContext(ctx, dataSourceKey, dataSourceName)
	assert.Equal(t, expectedContext, actualContext)
}
