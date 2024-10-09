package common

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/stretchr/testify/assert"
)

func TestGetDatabricksStagingName(t *testing.T) {
	resourceName := "test"
	expected := "databricks_test_pluginframework"
	result := GetDatabricksStagingName(resourceName)
	assert.Equal(t, expected, result, "GetDatabricksStagingName should return the expected staging name")
}

func TestGetDatabricksProductionName(t *testing.T) {
	resourceName := "test"
	expected := "databricks_test"
	result := GetDatabricksProductionName(resourceName)
	assert.Equal(t, expected, result, "GetDatabricksProductionName should return the expected production name")
}

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
