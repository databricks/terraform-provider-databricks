package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsClusterMissingTrueWhenClusterIdSpecifiedPresent(t *testing.T) {
	errorMessage := "{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Cluster 123 does not exist\"}"

	result := isClusterMissing(errorMessage, "123")

	assert.True(t, result)
}

func TestIsClusterMissingFalseWhenClusterIdSpecifiedNotPresent(t *testing.T) {
	errorMessage := "{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Cluster 123 does not exist\"}"

	result := isClusterMissing(errorMessage, "xyz")

	assert.False(t, result)
}

func TestIsClusterMissingFalseWhenErrorNotInCorrectFormat(t *testing.T) {
	errorMessage := "{\"error_code\":\"INVALID_PARAMETER_VALUE\",\"message\":\"Something random went bang xyz\"}"

	result := isClusterMissing(errorMessage, "xyz")

	assert.False(t, result)
}
