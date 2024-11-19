package notificationdestinations

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"
)

func TestValidateType_InvalidType(t *testing.T) {
	actualDiagnostics := validateType("INVALID")
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic("Invalid type 'INVALID'; valid types are EMAIL, MICROSOFT_TEAMS, PAGERDUTY, SLACK, WEBHOOK.", "")}
	assert.True(t, actualDiagnostics.HasError())
	assert.Equal(t, expectedDiagnostics, actualDiagnostics)
}
