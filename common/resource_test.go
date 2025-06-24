package common

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
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

func createTestResourceForSkipRead(skipRead bool) Resource {
	res := Resource{
		Create: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			log.Println("[DEBUG] Create called")
			return d.Set("foo", 1)
		},
		Read: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			log.Println("[DEBUG] Read called")
			d.Set("foo", 2)
			return nil
		},
		Update: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			log.Println("[DEBUG] Update called")
			return d.Set("foo", 3)
		},
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
	if skipRead {
		res.CanSkipReadAfterCreateAndUpdate = func(d *schema.ResourceData) bool {
			return true
		}
	}
	return res
}

func TestCreateSkipRead(t *testing.T) {
	client := &DatabricksClient{}
	ctx := context.Background()
	r := createTestResourceForSkipRead(true).ToResource()
	d := r.TestResourceData()
	diags := r.CreateContext(ctx, d, client)
	assert.False(t, diags.HasError())
	assert.Equal(t, 1, d.Get("foo"))
}

func TestCreateDontSkipRead(t *testing.T) {
	client := &DatabricksClient{}
	ctx := context.Background()
	r := createTestResourceForSkipRead(false).ToResource()
	d := r.TestResourceData()
	diags := r.CreateContext(ctx, d, client)
	assert.False(t, diags.HasError())
	assert.Equal(t, 2, d.Get("foo"))
}

func TestUpdateSkipRead(t *testing.T) {
	client := &DatabricksClient{}
	ctx := context.Background()
	r := createTestResourceForSkipRead(true).ToResource()
	d := r.TestResourceData()
	datas, err := r.Importer.StateContext(ctx, d, client)
	require.NoError(t, err)
	assert.Len(t, datas, 1)
	assert.False(t, r.Schema["foo"].ForceNew)
	assert.Equal(t, "", d.Id())

	diags := r.UpdateContext(ctx, d, client)
	assert.False(t, diags.HasError())
	assert.Equal(t, 3, d.Get("foo"))
}

func TestUpdateDontSkipRead(t *testing.T) {
	client := &DatabricksClient{}
	ctx := context.Background()
	r := createTestResourceForSkipRead(false).ToResource()
	d := r.TestResourceData()
	datas, err := r.Importer.StateContext(ctx, d, client)
	require.NoError(t, err)
	assert.Len(t, datas, 1)
	assert.False(t, r.Schema["foo"].ForceNew)
	assert.Equal(t, "", d.Id())

	diags := r.UpdateContext(ctx, d, client)
	assert.False(t, diags.HasError())
	assert.Equal(t, 2, d.Get("foo"))
}

func TestHTTP404TriggersResourceRemovalForReadAndDelete(t *testing.T) {
	nope := func(ctx context.Context,
		d *schema.ResourceData,
		c *DatabricksClient) error {
		return &apierr.APIError{
			ErrorCode:  "NOT_FOUND",
			StatusCode: 404,
			Message:    "nope",
		}
	}
	r := Resource{
		Create: nope,
		Read:   nope,
		Update: nope,
		Delete: nope,
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}.ToResource()

	client := &DatabricksClient{}
	ctx := context.Background()
	d := r.TestResourceData()

	// Create propagates 404 error
	diags := r.CreateContext(ctx, d, client)
	assert.True(t, diags.HasError())
	assert.Equal(t, "nope", diags[0].Summary)

	// Read removes the resource and swallows 404 error (clears ID)
	d.SetId("a")
	diags = r.ReadContext(ctx, d, client)
	assert.False(t, diags.HasError())
	assert.Equal(t, "", d.Id())

	// Update propagates 404 error (keeps ID)
	d.SetId("b")
	diags = r.UpdateContext(ctx, d, client)
	assert.True(t, diags.HasError())
	assert.Equal(t, "nope", diags[0].Summary)
	assert.Equal(t, "b", d.Id())

	// Delete removes the resource and swallows 404 error
	// if it was removed on backend (clears ID)
	d.SetId("c")
	diags = r.DeleteContext(ctx, d, client)
	assert.False(t, diags.HasError())
	assert.Equal(t, "", d.Id())
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
			return &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			}
		},
		Delete: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			return &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			}
		},
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}.ToResource()

	client := &DatabricksClient{}
	ctx := context.Background()
	d := r.TestResourceData()
	datas, err := r.Importer.StateContext(ctx, d, client)
	require.NoError(t, err)
	assert.Len(t, datas, 1)
	assert.False(t, r.Schema["foo"].ForceNew)
	assert.Equal(t, "", d.Id())

	diags := r.UpdateContext(ctx, d, client)
	assert.True(t, diags.HasError())
	assert.Equal(t, "nope", diags[0].Summary)
}

func TestRecoverableFromPanic(t *testing.T) {
	r := Resource{
		Update: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			return d.Set("foo", 1)
		},
		Read: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			panic(fmt.Errorf("what to do?..."))
		},
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}.ToResource()

	client := &DatabricksClient{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, ResourceName, "sample")
	d := r.TestResourceData()

	diags := r.UpdateContext(ctx, d, client)
	assert.True(t, diags.HasError())
	assert.Equal(t, "cannot read sample: panic: what to do?...", diags[0].Summary)
}

func TestCustomizeDiffRobustness(t *testing.T) {
	r := Resource{
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return fmt.Errorf("nope")
		},
	}.ToResource()

	ctx := context.Background()
	ctx = context.WithValue(ctx, ResourceName, "sample")

	err := r.CustomizeDiff(ctx, nil, nil)
	assert.EqualError(t, err, "cannot customize diff for sample: nope")

	r = Resource{
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			panic("oops")
		},
	}.ToResource()

	err = r.CustomizeDiff(ctx, nil, nil)
	assert.EqualError(t, err, "cannot customize diff for sample: panic: oops")
}

func TestWorkspacePathPrefixDiffSuppress(t *testing.T) {
	assert.True(t, WorkspacePathPrefixDiffSuppress("k", "/Workspace/foo/bar", "/Workspace/foo/bar", nil))
	assert.True(t, WorkspacePathPrefixDiffSuppress("k", "/Workspace/foo/bar", "/foo/bar", nil))
	assert.True(t, WorkspacePathPrefixDiffSuppress("k", "/foo/bar", "/Workspace/foo/bar", nil))
	assert.True(t, WorkspacePathPrefixDiffSuppress("k", "/foo/bar", "/foo/bar", nil))
	assert.False(t, WorkspacePathPrefixDiffSuppress("k", "/Workspace/1", "/Workspace/2", nil))
}

func TestWorkspaceOrEmptyPathPrefixDiffSuppress(t *testing.T) {
	assert.True(t, WorkspaceOrEmptyPathPrefixDiffSuppress("k", "/Workspace/foo/bar", "/Workspace/foo/bar", nil))
	assert.True(t, WorkspaceOrEmptyPathPrefixDiffSuppress("k", "/Workspace/foo/bar", "/foo/bar", nil))
	assert.True(t, WorkspaceOrEmptyPathPrefixDiffSuppress("k", "/foo/bar", "/Workspace/foo/bar", nil))
	assert.True(t, WorkspaceOrEmptyPathPrefixDiffSuppress("k", "/foo/bar", "/foo/bar", nil))
	assert.True(t, WorkspaceOrEmptyPathPrefixDiffSuppress("k", "/foo/bar", "", nil))
	assert.False(t, WorkspaceOrEmptyPathPrefixDiffSuppress("k", "/Workspace/1", "/Workspace/2", nil))
}

func TestEqualFoldDiffSuppress(t *testing.T) {
	assert.True(t, EqualFoldDiffSuppress("k", "A", "a", nil))
	assert.False(t, EqualFoldDiffSuppress("k", "A", "A2", nil))
}

func TestNoCustomize(t *testing.T) {
	dummySchema := map[string]*schema.Schema{
		"dummy": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}
	assert.Equal(t, dummySchema, NoCustomize(dummySchema))
}

func fakeDataSource() Resource {
	type Test struct {
		SomeId string `json:"some_id" tf:"computed"`
		Id     string `json:"id" tf:"computed"`
	}
	return NoClientData(func(ctx context.Context, data *Test) error {
		data.SomeId = "abc"
		data.Id = "def"
		return nil
	})
}

func TestFakeDataSource(t *testing.T) {
	r := fakeDataSource().ToResource()
	d := r.TestResourceData()
	client := &DatabricksClient{}
	diags := r.ReadContext(context.Background(), d, client)
	assert.False(t, diags.HasError())
	assert.Equal(t, "abc", d.Get("some_id"))
	assert.Equal(t, "def", d.Id())
}
