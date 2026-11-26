package notificationdestinations

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"
)

func TestValidateType_InvalidType(t *testing.T) {
	actualDiagnostics := validateType("INVALID")
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic("Invalid type 'INVALID'; valid types are DATABRICKS_APP_SLACK, DATABRICKS_APP_TEAMS, EMAIL, MICROSOFT_TEAMS, PAGERDUTY, SLACK, WEBHOOK.", "")}
	assert.True(t, actualDiagnostics.HasError())
	assert.Equal(t, expectedDiagnostics, actualDiagnostics)
}

func TestValidateType_DatabricksAppTypes(t *testing.T) {
	assert.False(t, validateType("DATABRICKS_APP_SLACK").HasError())
	assert.False(t, validateType("DATABRICKS_APP_TEAMS").HasError())
}
