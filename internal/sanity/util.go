package sanity

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/databrickslabs/databricks-terraform/common"
)

type MissingResourceCheck struct {
	Name     string
	ReadFunc func() error
}

type MissingResourceChecks []MissingResourceCheck

func (tests MissingResourceChecks) Verify(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.ReadFunc()
			assert.NotNil(t, err, "err should not be nil")
			assert.IsType(t, common.APIError{}, err, fmt.Sprintf("error: %s is not type api error", err.Error()))
			if apiError, ok := err.(common.APIError); ok {
				assert.True(t, apiError.IsMissing(), apiError)
			}
		})
	}
}