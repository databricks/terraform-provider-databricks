package util

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportingCallsRead(t *testing.T) {
	r := CommonResource{
		Read: func(ctx context.Context,
			d *schema.ResourceData,
			c *common.DatabricksClient) error {
			d.SetId("abc")
			return d.Set("foo", 1)
		},
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}.ToResource()

	d := r.TestResourceData()
	datas, err := r.Importer.StateContext(
		context.Background(), d,
		&common.DatabricksClient{})
	require.NoError(t, err)
	assert.Len(t, datas, 1)
	assert.True(t, r.Schema["foo"].ForceNew)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, 1, d.Get("foo"))
}

func TestUpdate(t *testing.T) {
	r := CommonResource{
		Update: func(ctx context.Context,
			d *schema.ResourceData,
			c *common.DatabricksClient) error {
			return d.Set("foo", 1)
		},
		Read: func(ctx context.Context,
			d *schema.ResourceData,
			c *common.DatabricksClient) error {
			return common.NotFound("nope")
		},
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}.ToResource()

	d := r.TestResourceData()
	datas, err := r.Importer.StateContext(
		context.Background(), d,
		&common.DatabricksClient{})
	require.NoError(t, err)
	assert.Len(t, datas, 1)
	assert.False(t, r.Schema["foo"].ForceNew)
	assert.Equal(t, "", d.Id())
}
