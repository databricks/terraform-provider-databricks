package identity

import (
	"testing"

	
	"github.com/stretchr/testify/assert"
)



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
