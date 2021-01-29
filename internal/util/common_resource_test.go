package util

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
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
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, 1, d.Get("foo"))
}
