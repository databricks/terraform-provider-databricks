package common

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestHTTP404TriggersResourceRemovalForReadAndDelete(t *testing.T) {
	nope := func(ctx context.Context,
		d *schema.ResourceData,
		c *DatabricksClient) error {
		return apierr.NotFound("nope")
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
			return apierr.NotFound("nope")
		},
		Delete: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			return apierr.NotFound("nope")
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
		Create: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			return d.Set("foo", 1)
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return fmt.Errorf("nope")
		},
	}.ToResource()

	ctx := context.Background()
	ctx = context.WithValue(ctx, ResourceName, "sample")

	err := r.CustomizeDiff(ctx, nil, nil)
	assert.EqualError(t, err, "cannot customize diff for sample: nope")

	r = Resource{
		Create: func(ctx context.Context,
			d *schema.ResourceData,
			c *DatabricksClient) error {
			return d.Set("foo", 1)
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			panic("oops")
		},
	}.ToResource()

	err = r.CustomizeDiff(ctx, nil, nil)
	assert.EqualError(t, err, "cannot customize diff for sample: panic: oops")
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

// Tests:
// - workspace-level provider with workspace-level resource with no workspace id
// - workspace-level provider with workspace-level resource with correct workspace ID
// - workspace-level provider with workspace-level resource with correct management workspace ID
// - workspace-level provider with workspace-level resource with incorrect workspace ID
// - workspace-level provider with account-level resource
// - account-level provider with workspace-level resource with no workspace id
// - account-level provider with workspace-level resource with workspace ID
// - account-level provider with workspace-level resource with management workspace ID
// - account-level provider with account-level resource
type resourceFuncToTest int

const (
	testRead resourceFuncToTest = iota
	testCreate
	testDelete
)

type assertion func(*testing.T, *schema.ResourceData, diag.Diagnostics)

type workspaceIdFieldTest struct {
	mockAccountClient   func(*mocks.MockAccountClient)
	mockWorkspaceClient func(*mocks.MockWorkspaceClient)
	workspaceIdField    WorkspaceIdField
	resourceData        map[string]any
	cb                  func(*testing.T) func(context.Context, *schema.ResourceData, *DatabricksClient) error
	funcToTest          resourceFuncToTest
	assertions          []assertion
}

func (test workspaceIdFieldTest) makeProviderResourceTest(t *testing.T) (string, func(*testing.T)) {
	var a *databricks.AccountClient
	var w *databricks.WorkspaceClient
	var testName string
	config := &config.Config{
		DatabricksEnvironment: &config.DatabricksEnvironment{
			Cloud:   config.CloudAWS,
			DnsZone: ".dev.databricks.com",
		},
	}
	config.WithTesting()
	if test.mockAccountClient != nil {
		mac := mocks.NewMockAccountClient(t)
		test.mockAccountClient(mac)
		a = mac.AccountClient
		a.Config = config
		config.AccountID = "test-account-id"
		testName = "AccountProvider"
	}
	if test.mockWorkspaceClient != nil {
		mwc := mocks.NewMockWorkspaceClient(t)
		test.mockWorkspaceClient(mwc)
		w = mwc.WorkspaceClient
		w.Config = config
		testName = "WorkspaceProvider"
	}
	if test.workspaceIdField == NoWorkspaceId {
		testName += "/No workspace_id field"
	} else {
		testName += fmt.Sprintf("/%s", test.workspaceIdField.Field())
	}
	c := &DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: config,
		},
		cachedAccountClient: a,
		cachedWorkspaceClients: map[int64]*cachedWorkspaceClient{
			defaultWorkspaceId: {
				client: w,
			},
		},
	}
	r := Resource{
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
		Create:           test.cb(t),
		Read:             test.cb(t),
		Update:           test.cb(t),
		Delete:           test.cb(t),
		WorkspaceIdField: test.workspaceIdField,
	}.ToResource()
	var f func(context.Context, *schema.ResourceData, any) diag.Diagnostics
	switch test.funcToTest {
	case testRead:
		f = r.ReadContext
		testName += "/ReadContext"
	case testCreate:
		f = r.CreateContext
		testName += "/CreateContext"
	case testDelete:
		f = r.DeleteContext
		testName += "/DeleteContext"
	default:
		t.Fatalf("unknown function to test: %v", test.funcToTest)
	}
	d := r.TestResourceData()
	for k, v := range test.resourceData {
		testName += fmt.Sprintf("/%s=%v", k, v)
		d.Set(k, v)
	}
	return testName, func(t *testing.T) {
		res := f(context.Background(), d, c)
		if test.assertions != nil {
			for _, assertion := range test.assertions {
				assertion(t, d, res)
			}
		} else {
			noError(t, d, res)
		}
	}
}

func noError(t *testing.T, _ *schema.ResourceData, diags diag.Diagnostics) {
	assert.Empty(t, diags)
}

func expectOriginalWorkspaceId(id int) assertion {
	return func(t *testing.T, rd *schema.ResourceData, d diag.Diagnostics) {
		assert.Equal(t, id, rd.Get("original_workspace_id"))
	}
}

func readWithWorkspace(t *testing.T) func(_ context.Context, d *schema.ResourceData, c *DatabricksClient) error {
	return func(_ context.Context, d *schema.ResourceData, c *DatabricksClient) error {
		_, err := c.WorkspaceClient()
		assert.NoError(t, err)
		return nil
	}
}

func readWithAccount(t *testing.T) func(_ context.Context, d *schema.ResourceData, c *DatabricksClient) error {
	return func(_ context.Context, d *schema.ResourceData, c *DatabricksClient) error {
		_, err := c.AccountClient()
		assert.NoError(t, err)
		return nil
	}
}

var testCases = []workspaceIdFieldTest{
	{
		mockWorkspaceClient: func(*mocks.MockWorkspaceClient) {},
		workspaceIdField:    WorkspaceId,
		cb:                  readWithWorkspace,
	},
	{
		mockWorkspaceClient: func(*mocks.MockWorkspaceClient) {},
		resourceData:        map[string]any{"workspace_id": 123},
		workspaceIdField:    WorkspaceId,
		cb:                  readWithWorkspace,
	},
	{
		mockWorkspaceClient: func(*mocks.MockWorkspaceClient) {},
		resourceData:        map[string]any{"workspace_id": 123},
		workspaceIdField:    WorkspaceId,
		funcToTest:          testCreate,
		cb:                  readWithWorkspace,
		assertions: []assertion{
			noError,
			expectOriginalWorkspaceId(123),
		},
	},
	{
		mockWorkspaceClient: func(*mocks.MockWorkspaceClient) {},
		resourceData:        map[string]any{"original_workspace_id": 123},
		workspaceIdField:    WorkspaceId,
		funcToTest:          testDelete,
		cb:                  readWithWorkspace,
	},
	{
		mockWorkspaceClient: func(*mocks.MockWorkspaceClient) {},
		workspaceIdField:    ManagementWorkspaceId,
		cb:                  readWithWorkspace,
	},
	{
		mockWorkspaceClient: func(*mocks.MockWorkspaceClient) {},
		workspaceIdField:    ManagementWorkspaceId,
		resourceData:        map[string]any{"management_workspace_id": 123},
		cb:                  readWithWorkspace,
	},
	// This case shouldn't happen in practice, as any workspace resource should support either workspace_id or
	// management_workspace_id. It still shouldn't error.
	{
		mockWorkspaceClient: func(*mocks.MockWorkspaceClient) {},
		workspaceIdField:    NoWorkspaceId,
		cb:                  readWithWorkspace,
	},
	{
		mockAccountClient: func(mac *mocks.MockAccountClient) {
			mac.GetMockWorkspacesAPI().EXPECT().GetByWorkspaceId(mock.Anything, int64(123)).Return(&provisioning.Workspace{
				WorkspaceId:    123,
				DeploymentName: "123",
			}, nil)
		},
		workspaceIdField: WorkspaceId,
		resourceData:     map[string]any{"workspace_id": 123},
		cb:               readWithWorkspace,
	},
	{
		mockAccountClient: func(mac *mocks.MockAccountClient) {},
		workspaceIdField:  WorkspaceId,
		cb:                readWithAccount,
	},
	{
		mockAccountClient: func(mac *mocks.MockAccountClient) {
			mac.GetMockWorkspacesAPI().EXPECT().GetByWorkspaceId(mock.Anything, int64(123)).Return(&provisioning.Workspace{
				WorkspaceId:    123,
				DeploymentName: "123",
			}, nil)
		},
		workspaceIdField: ManagementWorkspaceId,
		resourceData:     map[string]any{"management_workspace_id": 123},
		cb:               readWithWorkspace,
	},
}

func TestProviderResourceCombinations(t *testing.T) {
	for _, test := range testCases {
		name, testCase := test.makeProviderResourceTest(t)
		t.Run(name, testCase)
	}
}
