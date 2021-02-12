package common

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportingCallsRead(t *testing.T) {
	r := Resource{
		Read: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
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
		&DatabricksClient{})
	require.NoError(t, err)
	assert.Len(t, datas, 1)
	assert.True(t, r.Schema["foo"].ForceNew)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, 1, d.Get("foo"))
}

func TestUpdate(t *testing.T) {
	r := Resource{
		Update: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			return d.Set("foo", 1)
		},
		Read: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			return NotFound("nope")
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
		&DatabricksClient{})
	require.NoError(t, err)
	assert.Len(t, datas, 1)
	assert.False(t, r.Schema["foo"].ForceNew)
	assert.Equal(t, "", d.Id())
}
