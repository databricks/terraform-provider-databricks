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

func TestValidateInstanceProfileARN(t *testing.T) {
	testCases := []struct {
		instanceProfileARN string
		errorCount         int
	}{
		{"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", 0},
		{"arn:aws:iam::999999999999:role/not-an-instance-profile", 1},
		{"", 1},
		{"invalid-profile", 1},
	}
	for _, tc := range testCases {
		_, errs := ValidateInstanceProfileARN(tc.instanceProfileARN, "key")

		assert.Lenf(t, errs, tc.errorCount, "directory '%s' does not generate the expected error count", tc.instanceProfileARN)
	}
}
