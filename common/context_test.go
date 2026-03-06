package common

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestAddContextToAllResources(t *testing.T) {
	check := func(ctx context.Context, rd *schema.ResourceData, i any) diag.Diagnostics {
		assert.Equal(t, "bar", ResourceName.GetOrUnknown(ctx))
		assert.Equal(t, "sdkv2", Sdk.GetOrUnknown(ctx))
		return nil
	}
	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"foo_bar": {
				CreateContext: check,
				ReadContext:   check,
				UpdateContext: check,
				DeleteContext: check,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"foo_bar": {
				ReadContext: check,
			},
		},
	}
	AddContextToAllResources(p, "foo")
	p.ResourcesMap["foo_bar"].CreateContext(context.Background(), nil, nil)
}

func TestAddUnifiedProviderContext_WithProviderConfig(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"provider_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workspace_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	check := func(ctx context.Context, rd *schema.ResourceData, i any) diag.Diagnostics {
		ua := useragent.FromContext(ctx)
		assert.True(t, strings.Contains(ua, "unifiedprovider/true"),
			"User-Agent should contain unifiedprovider/true, got: %s", ua)
		return nil
	}

	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"foo_bar": {
				Schema:        testSchema,
				CreateContext: check,
			},
		},
	}
	AddContextToAllResources(p, "foo")

	d := schema.TestResourceDataRaw(t, testSchema, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "123456",
			},
		},
	})

	p.ResourcesMap["foo_bar"].CreateContext(context.Background(), d, nil)
}

func TestAddUnifiedProviderContext_WithoutProviderConfig(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	check := func(ctx context.Context, rd *schema.ResourceData, i any) diag.Diagnostics {
		ua := useragent.FromContext(ctx)
		assert.False(t, strings.Contains(ua, "unifiedprovider/true"),
			"User-Agent should NOT contain unifiedprovider/true, got: %s", ua)
		return nil
	}

	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"foo_bar": {
				Schema:        testSchema,
				CreateContext: check,
			},
		},
	}
	AddContextToAllResources(p, "foo")

	d := schema.TestResourceDataRaw(t, testSchema, map[string]interface{}{
		"name": "test",
	})

	p.ResourcesMap["foo_bar"].CreateContext(context.Background(), d, nil)
}

func TestAddUnifiedProviderContext_WithEmptyWorkspaceID(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"provider_config": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workspace_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	check := func(ctx context.Context, rd *schema.ResourceData, i any) diag.Diagnostics {
		ua := useragent.FromContext(ctx)
		assert.False(t, strings.Contains(ua, "unifiedprovider/true"),
			"User-Agent should NOT contain unifiedprovider/true when workspace_id is empty, got: %s", ua)
		return nil
	}

	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"foo_bar": {
				Schema:        testSchema,
				CreateContext: check,
			},
		},
	}
	AddContextToAllResources(p, "foo")

	d := schema.TestResourceDataRaw(t, testSchema, map[string]interface{}{
		"name": "test",
		"provider_config": []interface{}{
			map[string]interface{}{
				"workspace_id": "",
			},
		},
	})

	p.ResourcesMap["foo_bar"].CreateContext(context.Background(), d, nil)
}
