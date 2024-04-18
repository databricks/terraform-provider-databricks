package qa

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestRandomEmail(t *testing.T) {
	email := RandomEmail()
	assert.NotEmpty(t, email)
}

func TestRandomLongName(t *testing.T) {
	n := RandomLongName()
	assert.Equal(t, 37, len(n))
}

func TestRandomName(t *testing.T) {
	n := RandomName("x", "y")
	assert.Equal(t, 14, len(n))
}

func TestResourceFixture(t *testing.T) {
	client, server, err := HttpFixtureClient(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/a/b/c",
			Response: HTTPFixture{
				Method: "SOME",
			},
			ExpectedRequest: map[string]bool{
				"check": true,
			},
		},
	})
	defer server.Close()
	assert.NoError(t, err)

	var a HTTPFixture
	err = client.Post(context.Background(), "/a/b/c", map[string]bool{
		"check": true,
	}, &a)
	assert.NoError(t, err)
	assert.Equal(t, "SOME", a.Method)
}

func TestResourceFixture_Hint(t *testing.T) {
	t2 := testing.T{}
	client, server, err := HttpFixtureClient(&t2, []HTTPFixture{})
	defer server.Close()
	assert.NoError(t, err)

	var a HTTPFixture
	err = client.Post(context.Background(), "/a/b/c", map[string]bool{
		"check": true,
	}, &a)
	assert.Error(t, err)
	assert.True(t, t2.Failed())
}

func noopContext(_ context.Context, _ *schema.ResourceData, _ *common.DatabricksClient) error {
	return nil
}

var noopResource = common.Resource{
	Schema: map[string]*schema.Schema{
		"dummy": {
			Type:     schema.TypeBool,
			Required: true,
		},
	},
	Read:   noopContext,
	Create: noopContext,
	Update: noopContext,
	Delete: noopContext,
}

var noopContextResource = common.Resource{
	Schema: map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"dummy": {
			Type:     schema.TypeBool,
			Required: true,
		},
	},
	Read:   noopContext,
	Create: noopContext,
	Update: func(_ context.Context, d *schema.ResourceData, _ *common.DatabricksClient) error {
		// nolint
		d.Set("trigger", "corrupt")
		return nil
	},
	Delete: noopContext,
}

func TestResourceFixture_ID(t *testing.T) {
	_, err := ResourceFixture{}.prepareExecution(nil)
	assert.EqualError(t, err, "no `Create|Read|Update|Delete: true` or `ExpectedDiff` specified")

	f := ResourceFixture{
		Resource: noopResource,
		Read:     true,
		AzureSPN: true,
		Gcp:      true,
		HCL:      `dummy = true`,
	}
	_, err = f.Apply(t)
	assert.EqualError(t, err, "ID must be set for Read")

	f.Read = false
	f.Delete = true
	_, err = f.Apply(t)
	assert.EqualError(t, err, "ID must be set for Delete")

	f.Delete = false
	f.Update = true
	_, err = f.Apply(t)
	assert.EqualError(t, err, "ID must be set for Update")

	f.ID = "_"
	f.Create = true
	_, err = f.Apply(t)
	assert.EqualError(t, err, "ID is not available for Create")

	f.ID = ""
	_, err = f.Apply(t)
	assert.EqualError(t, err, "resource is not expected to be removed")

	f.Removed = true
	_, err = f.Apply(t)
	assert.NoError(t, err)
	f.ApplyNoError(t)
}

func TestResourceFixture_Apply(t *testing.T) {
	ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "text",
				Data:       "yes",
			}
		},
		Azure:    true,
		Resource: noopResource,
		ID:       "x",
		New:      true,
		Read:     true,
		HCL:      `dummy = true`,
	}.ApplyNoError(t)
}

func TestResourceFixture_Apply_Token(t *testing.T) {
	ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "text",
				Data:       "yes",
			}
		},
		Azure:    true,
		Resource: noopResource,
		ID:       "x",
		New:      true,
		Read:     true,
		HCL:      `dummy = true`,
		Token:    "test",
	}.ApplyNoError(t)
}

func TestResourceFixture_ApplyDelete(t *testing.T) {
	d, err := ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "text",
				Data:       "yes",
			}
		},
		Azure:    true,
		Resource: noopContextResource,
		ID:       "x",
		Delete:   true,
		HCL: `
		dummy = true
		trigger = "now"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, true, d.Get("dummy"))
}

func TestResourceFixture_ApplyAndExpectData(t *testing.T) {
	ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "text",
				Data:       "yes",
			}
		},
		Azure:    true,
		Resource: noopContextResource,
		ID:       "x",
		Delete:   true,
		HCL: `
		dummy = true
		trigger = "now"
		`,
	}.ApplyAndExpectData(t, map[string]any{"id": "x", "dummy": true, "trigger": "now"})
}

func TestResourceFixture_InstanceState(t *testing.T) {
	ResourceFixture{
		Resource: noopContextResource,
		ID:       "x",
		Update:   true,
		InstanceState: map[string]string{
			"dummy":   "false",
			"trigger": "y",
		},
		State: map[string]any{
			"dummy":   "true",
			"trigger": "x",
		},
	}.ExpectError(t, "changes require new: trigger")
}

func TestResourceFixture_Apply_Fail(t *testing.T) {
	ResourceFixture{
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "text",
				Data:       "yes",
			}
		},
		Resource: noopResource,
		Create:   true,
		State: map[string]any{
			"dummy": true,
			"check": false,
		},
	}.ExpectError(t, "invalid config supplied. [check] Invalid or unknown key")
}

func TestUnionFixturesLists(t *testing.T) {
	x := UnionFixturesLists([]HTTPFixture{
		{Method: "GET"},
		{Method: "POST"},
	}, []HTTPFixture{
		{Method: "DELETE"},
	})
	assert.Len(t, x, 3)
}

func TestFixHCL_CornerCase(t *testing.T) {
	x := fixHCL([]map[string]any{
		{
			"x": true,
		},
	})
	assert.Len(t, x, 1)
}

func TestGetEnvOrSkipTest(t *testing.T) {
	u := GetEnvOrSkipTest(t, "HOME")
	assert.NotEmpty(t, u)
}

func TestGetEnvOrSkipTest_Skip(t *testing.T) {
	GetEnvOrSkipTest(t, "")
}

func TestDiagsToString(t *testing.T) {
	var d diag.Diagnostics
	assert.Empty(t, diagsToString(d))

	assert.Equal(t, "[d] c. b", diagsToString(diag.Diagnostics{
		diag.Diagnostic{
			Detail:        "a",
			Summary:       "c",
			AttributePath: cty.GetAttrPath("d"),
		},
		diag.Diagnostic{
			Detail:        "b",
			Summary:       "ConflictsWith",
			AttributePath: cty.IndexIntPath(2),
		},
	}))
}

func TestResourceCornerCases(t *testing.T) {
	type dummy struct{}
	x := map[string]string{
		"foo": "bar",
	}
	ResourceCornerCases(t, common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var b dummy
			return c.Post(ctx, "/dummy", x, &b)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var b dummy
			return c.Get(ctx, "/dummy", x, &b)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return c.Put(ctx, "/dummy", x)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return c.Delete(ctx, "/dummy", x)
		},
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	},
		CornerCaseID("x"),
		CornerCaseExpectError("i'm a teapot"),
		CornerCaseSkipCRUD("head"))
}

func TestAssertErrorStartsWith(t *testing.T) {
	AssertErrorStartsWith(t, fmt.Errorf("abc"), "a")
}
