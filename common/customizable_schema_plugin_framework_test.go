package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/stretchr/testify/assert"
)

func TestCustomizeSchema(t *testing.T) {
	scm := pluginFrameworkStructToSchema(DummyTfSdk{}, func(c CustomizableSchemaPluginFramework) CustomizableSchemaPluginFramework {
		c.AddNewField("new_field", schema.StringAttribute{Required: true})
		c.SchemaPath("nested").AddNewField("new_field", schema.StringAttribute{Required: true})
		c.SchemaPath("nested").AddNewField("to_be_removed", schema.StringAttribute{Required: true})
		c.SchemaPath("nested").RemoveField("to_be_removed")
		return c
	})
	assert.True(t, scm.Attributes["new_field"].IsRequired())
	assert.True(t, MustSchemaAttributePath(scm.Attributes, "nested", "new_field").IsRequired())
	attr := MustSchemaAttributePath(scm.Attributes, "nested").(schema.SingleNestedAttribute).Attributes
	_, ok := attr["to_be_removed"]
	assert.True(t, !ok)
}
