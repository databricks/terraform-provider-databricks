package tfschema

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type TestTfSdk struct {
	Description types.String            `tfsdk:"description" tf:""`
	Nested      *NestedTfSdk            `tfsdk:"nested" tf:"optional"`
	Map         map[string]types.String `tfsdk:"map" tf:"optional"`
}

type NestedTfSdk struct {
	Name    types.String `tfsdk:"name" tf:"optional"`
	Enabled types.Bool   `tfsdk:"enabled" tf:"optional"`
}

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

func TestCustomizeSchemaSetRequired(t *testing.T) {
	scm := ResourceStructToSchema(TestTfSdk{}, func(c CustomizableSchema) CustomizableSchema {
		c.SetRequired("nested", "enabled")
		return c
	})

	assert.True(t, scm.Attributes["nested"].(schema.SingleNestedAttribute).Attributes["enabled"].IsRequired())
}

func TestCustomizeSchemaSetOptional(t *testing.T) {
	scm := ResourceStructToSchema(TestTfSdk{}, func(c CustomizableSchema) CustomizableSchema {
		c.SetOptional("description")
		return c
	})

	assert.True(t, scm.Attributes["description"].IsOptional())
}

func TestCustomizeSchemaSetSensitive(t *testing.T) {
	scm := ResourceStructToSchema(TestTfSdk{}, func(c CustomizableSchema) CustomizableSchema {
		c.SetSensitive("nested", "name")
		return c
	})

	assert.True(t, scm.Attributes["nested"].(schema.SingleNestedAttribute).Attributes["name"].IsSensitive())
}

func TestCustomizeSchemaSetDeprecated(t *testing.T) {
	scm := ResourceStructToSchema(TestTfSdk{}, func(c CustomizableSchema) CustomizableSchema {
		c.SetDeprecated("deprecated", "map")
		return c
	})

	assert.True(t, scm.Attributes["map"].GetDeprecationMessage() == "deprecated")
}

func TestCustomizeSchemaSetReadOnly(t *testing.T) {
	scm := ResourceStructToSchema(TestTfSdk{}, func(c CustomizableSchema) CustomizableSchema {
		c.SetReadOnly("map")
		return c
	})
	assert.True(t, !scm.Attributes["map"].IsOptional())
	assert.True(t, !scm.Attributes["map"].IsRequired())
	assert.True(t, scm.Attributes["map"].IsComputed())
}

func TestCustomizeSchemaAddValidator(t *testing.T) {
	scm := ResourceStructToSchema(TestTfSdk{}, func(c CustomizableSchema) CustomizableSchema {
		c.AddValidator(stringLengthBetweenValidator{}, "description")
		return c
	})

	assert.True(t, len(scm.Attributes["description"].(schema.StringAttribute).Validators) == 1)
}

func TestCustomizeSchemaAddPlanModifier(t *testing.T) {
	scm := ResourceStructToSchema(TestTfSdk{}, func(c CustomizableSchema) CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "description")
		return c
	})

	assert.True(t, len(scm.Attributes["description"].(schema.StringAttribute).PlanModifiers) == 1)
}
