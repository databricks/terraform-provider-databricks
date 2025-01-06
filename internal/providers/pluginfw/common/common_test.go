package common_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/stretchr/testify/assert"
)

func TestGetDatabricksStagingName(t *testing.T) {
	resourceName := "test"
	expected := "databricks_test_pluginframework"
	result := common.GetDatabricksStagingName(resourceName)
	assert.Equal(t, expected, result, "GetDatabricksStagingName should return the expected staging name")
}

func TestGetDatabricksProductionName(t *testing.T) {
	resourceName := "test"
	expected := "databricks_test"
	result := common.GetDatabricksProductionName(resourceName)
	assert.Equal(t, expected, result, "GetDatabricksProductionName should return the expected production name")
}
