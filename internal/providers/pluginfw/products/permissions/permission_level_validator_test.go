package permissions

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
)

func TestPermissionLevelValidator_KnownObjectType(t *testing.T) {
	v := permissionLevelValidator{}

	tests := []struct {
		name            string
		objectType      string
		permissionLevel string
		expectError     bool
		errorContains   string
	}{
		{
			name:            "valid cluster permission",
			objectType:      "clusters",
			permissionLevel: "CAN_ATTACH_TO",
			expectError:     false,
		},
		{
			name:            "valid cluster permission - CAN_RESTART",
			objectType:      "clusters",
			permissionLevel: "CAN_RESTART",
			expectError:     false,
		},
		{
			name:            "invalid cluster permission",
			objectType:      "clusters",
			permissionLevel: "INVALID_PERMISSION",
			expectError:     true,
			errorContains:   "not valid for object type",
		},
		{
			name:            "valid job permission",
			objectType:      "jobs",
			permissionLevel: "CAN_VIEW",
			expectError:     false,
		},
		{
			name:            "valid authorization permission",
			objectType:      "authorization",
			permissionLevel: "CAN_USE",
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test config
			configValues := map[string]tftypes.Value{
				"object_type":      tftypes.NewValue(tftypes.String, tt.objectType),
				"permission_level": tftypes.NewValue(tftypes.String, tt.permissionLevel),
			}

			config := tfsdk.Config{
				Schema: schema.Schema{
					Attributes: map[string]schema.Attribute{
						"object_type": schema.StringAttribute{
							Required: true,
						},
						"permission_level": schema.StringAttribute{
							Required: true,
						},
					},
				},
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"object_type":      tftypes.String,
						"permission_level": tftypes.String,
					},
				}, configValues),
			}

			req := validator.StringRequest{
				Path:        path.Root("permission_level"),
				ConfigValue: types.StringValue(tt.permissionLevel),
				Config:      config,
			}

			resp := &validator.StringResponse{}

			v.ValidateString(context.Background(), req, resp)

			if tt.expectError {
				assert.True(t, resp.Diagnostics.HasError(), "Expected error but got none")
				if tt.errorContains != "" {
					assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), tt.errorContains)
				}
			} else {
				assert.False(t, resp.Diagnostics.HasError(), "Expected no error but got: %v", resp.Diagnostics.Errors())
			}
		})
	}
}

func TestPermissionLevelValidator_UnknownObjectType(t *testing.T) {
	v := permissionLevelValidator{}

	// Test with an unknown object type - should not error (let API handle it)
	configValues := map[string]tftypes.Value{
		"object_type":      tftypes.NewValue(tftypes.String, "new-unknown-type"),
		"permission_level": tftypes.NewValue(tftypes.String, "SOME_PERMISSION"),
	}

	config := tfsdk.Config{
		Schema: schema.Schema{
			Attributes: map[string]schema.Attribute{
				"object_type": schema.StringAttribute{
					Required: true,
				},
				"permission_level": schema.StringAttribute{
					Required: true,
				},
			},
		},
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"object_type":      tftypes.String,
				"permission_level": tftypes.String,
			},
		}, configValues),
	}

	req := validator.StringRequest{
		Path:        path.Root("permission_level"),
		ConfigValue: types.StringValue("SOME_PERMISSION"),
		Config:      config,
	}

	resp := &validator.StringResponse{}

	v.ValidateString(context.Background(), req, resp)

	// Should NOT error - let API handle validation for unknown types
	assert.False(t, resp.Diagnostics.HasError(), "Should not error for unknown object type")
}

func TestPermissionLevelValidator_MissingObjectType(t *testing.T) {
	v := permissionLevelValidator{}

	// Test with null object_type - should not error (let API handle it)
	configValues := map[string]tftypes.Value{
		"object_type":      tftypes.NewValue(tftypes.String, nil),
		"permission_level": tftypes.NewValue(tftypes.String, "CAN_USE"),
	}

	config := tfsdk.Config{
		Schema: schema.Schema{
			Attributes: map[string]schema.Attribute{
				"object_type": schema.StringAttribute{
					Required: true,
				},
				"permission_level": schema.StringAttribute{
					Required: true,
				},
			},
		},
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"object_type":      tftypes.String,
				"permission_level": tftypes.String,
			},
		}, configValues),
	}

	req := validator.StringRequest{
		Path:        path.Root("permission_level"),
		ConfigValue: types.StringValue("CAN_USE"),
		Config:      config,
	}

	resp := &validator.StringResponse{}

	v.ValidateString(context.Background(), req, resp)

	// Should NOT error - let API handle validation when object_type is missing
	assert.False(t, resp.Diagnostics.HasError(), "Should not error when object_type is null")
}
