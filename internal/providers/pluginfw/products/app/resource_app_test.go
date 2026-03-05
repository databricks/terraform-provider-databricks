package app

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestUppercasePlanModifier(t *testing.T) {
	tests := []struct {
		name          string
		planValue     types.String
		expectedValue types.String
	}{
		{
			name:          "lowercase value is uppercased",
			planValue:     types.StringValue("medium"),
			expectedValue: types.StringValue("MEDIUM"),
		},
		{
			name:          "mixed case value is uppercased",
			planValue:     types.StringValue("Medium"),
			expectedValue: types.StringValue("MEDIUM"),
		},
		{
			name:          "uppercase value remains unchanged",
			planValue:     types.StringValue("LARGE"),
			expectedValue: types.StringValue("LARGE"),
		},
		{
			name:          "null value is not modified",
			planValue:     types.StringNull(),
			expectedValue: types.StringNull(),
		},
		{
			name:          "unknown value is not modified",
			planValue:     types.StringUnknown(),
			expectedValue: types.StringUnknown(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := planmodifier.StringRequest{
				PlanValue: tt.planValue,
			}
			resp := &planmodifier.StringResponse{
				PlanValue: tt.planValue,
			}

			m := uppercasePlanModifier{}
			m.PlanModifyString(context.Background(), req, resp)

			assert.Equal(t, tt.expectedValue, resp.PlanValue)
		})
	}
}

func TestComputeSizeValidation(t *testing.T) {
	var computeSize apps.ComputeSize
	values := make([]string, 0)
	for _, v := range computeSize.Values() {
		values = append(values, string(v))
	}
	v := stringvalidator.OneOfCaseInsensitive(values...)

	tests := []struct {
		name        string
		value       types.String
		expectError bool
	}{
		{
			name:        "valid uppercase MEDIUM",
			value:       types.StringValue("MEDIUM"),
			expectError: false,
		},
		{
			name:        "valid uppercase LARGE",
			value:       types.StringValue("LARGE"),
			expectError: false,
		},
		{
			name:        "valid mixed case Medium",
			value:       types.StringValue("Medium"),
			expectError: false,
		},
		{
			name:        "valid lowercase large",
			value:       types.StringValue("large"),
			expectError: false,
		},
		{
			name:        "invalid value FooBar is rejected",
			value:       types.StringValue("FooBar"),
			expectError: true,
		},
		{
			name:        "invalid value test is rejected",
			value:       types.StringValue("test"),
			expectError: true,
		},
		{
			name:        "null value is accepted",
			value:       types.StringNull(),
			expectError: false,
		},
		{
			name:        "unknown value is accepted",
			value:       types.StringUnknown(),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.StringRequest{
				ConfigValue: tt.value,
			}
			resp := &validator.StringResponse{}

			v.ValidateString(context.Background(), req, resp)

			assert.Equal(t, tt.expectError, resp.Diagnostics.HasError())
		})
	}
}
