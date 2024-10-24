package user

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"
)

func TestValidateFilters(t *testing.T) {
	userInfo := UsersInfo{
		DisplayNameContains: "filter",
		UserNameContains:    "another_filter",
	}
	actualDiagnostics := validateFilters(&userInfo)
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic("Invalid configuration", "Exactly one of display_name_contains or user_name_contains should be specified, not both.")}
	assert.True(t, actualDiagnostics.HasError())
	assert.Equal(t, expectedDiagnostics, actualDiagnostics)
}
