package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddApiField_ValidValues(t *testing.T) {
	s := AddApiField(map[string]*schema.Schema{})
	apiSchema := s["api"]
	require.NotNil(t, apiSchema)
	require.NotNil(t, apiSchema.ValidateFunc)

	// "account" is valid
	_, errs := apiSchema.ValidateFunc("account", "api")
	assert.Empty(t, errs)

	// "workspace" is valid
	_, errs = apiSchema.ValidateFunc("workspace", "api")
	assert.Empty(t, errs)
}

func TestAddApiField_InvalidValues(t *testing.T) {
	s := AddApiField(map[string]*schema.Schema{})
	validateFunc := s["api"].ValidateFunc

	for _, invalid := range []string{"", "foo", "ACCOUNT", "Workspace", "acct", "ws"} {
		_, errs := validateFunc(invalid, "api")
		assert.NotEmpty(t, errs, "expected error for value %q", invalid)
	}
}
