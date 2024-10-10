package common

import (
	"testing"

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
