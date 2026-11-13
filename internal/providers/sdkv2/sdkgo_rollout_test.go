package sdkv2

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

func TestParseSdkGoResources(t *testing.T) {
	for value, want := range map[string][]string{
		"":                                  nil,
		" , ":                               nil,
		"databricks_service_principal":      {"databricks_service_principal"},
		" databricks_service_principal , ,": {"databricks_service_principal"},
	} {
		got, err := parseSdkGoResources(value)
		require.NoError(t, err, value)
		assert.ElementsMatch(t, want, got, value)
	}
}

func TestParseSdkGoResourcesRejectsUnknownResources(t *testing.T) {
	_, err := parseSdkGoResources("databricks_service_principal,databricks_user")
	assert.EqualError(t, err, `USE_SDK_GO_RESOURCES: resource "databricks_user" has no `+
		`sdk-go implementation; supported resources: databricks_service_principal`)
}

func TestSdkGoResourcesOptionReplacesEnvironment(t *testing.T) {
	options := &sdkV2ProviderOptions{}
	WithSdkGoResources([]string{"databricks_service_principal"})(options)
	got, err := options.selectedSdkGoResources()
	require.NoError(t, err)
	assert.Equal(t, []string{"databricks_service_principal"}, got)

	WithSdkGoResources(nil)(options)
	got, err = options.selectedSdkGoResources()
	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestSdkGoResourceNamesCoverRegisteredResources(t *testing.T) {
	resources := Resources()
	for _, name := range SdkGoResourceNames() {
		assert.Contains(t, resources, name)
	}
}

func TestEnableSdkGoResourcesMarksOnlySelectedResources(t *testing.T) {
	seen := map[string]bool{}
	read := func(name string) *schema.Resource {
		return &schema.Resource{ReadContext: func(ctx context.Context,
			_ *schema.ResourceData, _ any) diag.Diagnostics {
			seen[name] = common.UsesSdkGo(ctx)
			return nil
		}}
	}
	resources := map[string]*schema.Resource{
		"selected": read("selected"),
		"other":    read("other"),
	}
	enableSdkGoResources(resources, []string{"selected", "missing"})
	for _, r := range resources {
		r.ReadContext(context.Background(), nil, nil)
	}
	assert.Equal(t, map[string]bool{"selected": true, "other": false}, seen)
}

func TestDatabricksProviderRejectsUnknownSdkGoResource(t *testing.T) {
	p := DatabricksProvider(WithSdkGoResources([]string{"databricks_user"}))
	config := terraform.NewResourceConfigRaw(map[string]any{})
	diags := p.Configure(context.Background(), config)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary,
		`resource "databricks_user" has no sdk-go implementation`)
}
