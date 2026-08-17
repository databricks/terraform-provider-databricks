package secret_uc

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// secretValueAttrs are the attributes that carry secret material: `value` is the
// write path and `effective_value` the read path. Both must be masked in
// plan/apply output.
var secretValueAttrs = []string{"value", "effective_value"}

func TestResourceSecret_SecretValuesAreSensitive(t *testing.T) {
	r := ResourceSecret()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)

	for _, name := range secretValueAttrs {
		attr, ok := resp.Schema.Attributes[name]
		require.True(t, ok, "%s attribute must exist", name)
		strAttr, ok := attr.(schema.StringAttribute)
		require.True(t, ok, "%s must be a string attribute", name)
		assert.True(t, strAttr.Sensitive, "%s must be sensitive so the secret is not shown in plan/apply output", name)
	}
}

func TestDataSourceSecret_SecretValuesAreSensitive(t *testing.T) {
	d := DataSourceSecret()
	resp := &datasource.SchemaResponse{}
	d.Schema(context.Background(), datasource.SchemaRequest{}, resp)

	for _, name := range secretValueAttrs {
		attr, ok := resp.Schema.Attributes[name]
		require.True(t, ok, "%s attribute must exist", name)
		strAttr, ok := attr.(dsschema.StringAttribute)
		require.True(t, ok, "%s must be a string attribute", name)
		assert.True(t, strAttr.Sensitive, "%s must be sensitive so the secret is not shown in plan/apply output", name)
	}
}

func TestDataSourceSecrets_SecretValuesAreSensitive(t *testing.T) {
	d := DataSourceSecrets()
	resp := &datasource.SchemaResponse{}
	d.Schema(context.Background(), datasource.SchemaRequest{}, resp)

	secrets, ok := resp.Schema.Attributes["secrets"]
	require.True(t, ok, "secrets attribute must exist")
	list, ok := secrets.(dsschema.ListNestedAttribute)
	require.True(t, ok, "secrets must be a list nested attribute")

	// Every listed secret is rendered, so masking must apply to the nested attributes.
	for _, name := range secretValueAttrs {
		attr, ok := list.NestedObject.Attributes[name]
		require.True(t, ok, "secrets.%s attribute must exist", name)
		strAttr, ok := attr.(dsschema.StringAttribute)
		require.True(t, ok, "secrets.%s must be a string attribute", name)
		assert.True(t, strAttr.Sensitive, "secrets.%s must be sensitive so the secret is not shown in plan/apply output", name)
	}
}
