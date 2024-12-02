package common

import (
	"context"
	"testing"

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
