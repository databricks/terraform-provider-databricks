package common

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stretchr/testify/assert"
)

type stringLengthBetweenValidator struct {
	Max int
	Min int
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v stringLengthBetweenValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("string length must be between %d and %d", v.Min, v.Max)
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v stringLengthBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("string length must be between `%d` and `%d`", v.Min, v.Max)
}

// Validate runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v stringLengthBetweenValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	strLen := len(req.ConfigValue.ValueString())

	if strLen < v.Min || strLen > v.Max {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid String Length",
			fmt.Sprintf("String length must be between %d and %d, got: %d.", v.Min, v.Max, strLen),
		)

		return
	}
}

func TestCustomizeSchema(t *testing.T) {
	scm := PluginFrameworkResourceStructToSchema(DummyTfSdk{}, func(c CustomizableSchemaPluginFramework) CustomizableSchemaPluginFramework {
		c.AddNewField("new_field", schema.StringAttribute{Required: true})
		c.AddNewField("new_field", schema.StringAttribute{Required: true}, "nested")
		c.AddNewField("to_be_removed", schema.StringAttribute{Required: true}, "nested")
		c.RemoveField("to_be_removed", "nested")
		c.SetRequired("nested", "enabled")
		c.SetOptional("description")
		c.SetSensitive("nested", "name")
		c.SetDeprecated("deprecated", "map")
		c.SetReadOnly("map")
		c.AddValidator(stringLengthBetweenValidator{}, "description")
		return c
	})
	assert.True(t, scm.Attributes["new_field"].IsRequired())
	assert.True(t, MustSchemaAttributePath(scm.Attributes, "nested", "new_field").IsRequired())
	assert.True(t, MustSchemaAttributePath(scm.Attributes, "nested", "enabled").IsRequired())
	assert.True(t, MustSchemaAttributePath(scm.Attributes, "nested", "name").IsSensitive())
	assert.True(t, MustSchemaAttributePath(scm.Attributes, "map").GetDeprecationMessage() == "deprecated")
	assert.True(t, scm.Attributes["description"].IsOptional())
	assert.True(t, !scm.Attributes["map"].IsOptional())
	assert.True(t, !scm.Attributes["map"].IsRequired())
	assert.True(t, scm.Attributes["map"].IsComputed())
	attr := MustSchemaAttributePath(scm.Attributes, "nested").(schema.SingleNestedAttribute).Attributes
	_, ok := attr["to_be_removed"]
	assert.True(t, len(MustSchemaAttributePath(scm.Attributes, "description").(schema.StringAttribute).Validators) == 1)
	assert.True(t, !ok)
}
