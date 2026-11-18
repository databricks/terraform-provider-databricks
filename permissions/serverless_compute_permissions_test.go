package permissions

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	testInteractiveComputeId = "708f8c89-dcbc-3cc4-b8da-4c638a21994c"
	testAutomatedComputeId   = "42d14549-3223-342b-a48b-ec54edf27df2"
)

// serverlessComputeAcl is what the Permissions API returns for a serverless compute after the
// provider sets group CAN_USE: the group, the applying user, and the inherited admins grant.
func serverlessComputeAcl(computeId string) *iam.ObjectPermissions {
	return &iam.ObjectPermissions{
		ObjectId:   "/serverless-compute/" + computeId,
		ObjectType: "serverless-compute",
		AccessControlList: []iam.AccessControlResponse{
			{
				GroupName:      "data-science",
				AllPermissions: []iam.Permission{{PermissionLevel: "CAN_USE"}},
			},
			{
				UserName:       TestingAdminUser,
				AllPermissions: []iam.Permission{{PermissionLevel: "CAN_MANAGE"}},
			},
			{
				GroupName: "admins",
				AllPermissions: []iam.Permission{{
					PermissionLevel:     "CAN_MANAGE",
					Inherited:           true,
					InheritedFromObject: []string{"/serverless-compute"},
				}},
			},
		},
	}
}

// expectSetAndGet mocks the Permissions API calls made by create: a full-replace Set of the group
// grant plus the applying user as CAN_MANAGE, then a Get to read the result back.
func expectSetAndGet(mwc *mocks.MockWorkspaceClient, computeId string) {
	mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
	e := mwc.GetMockPermissionsAPI().EXPECT()
	e.Set(mock.Anything, iam.SetObjectPermissions{
		RequestObjectId:   computeId,
		RequestObjectType: "serverless-compute",
		AccessControlList: []iam.AccessControlRequest{
			{GroupName: "data-science", PermissionLevel: "CAN_USE"},
			{UserName: TestingAdminUser, PermissionLevel: "CAN_MANAGE"},
		},
	}).Return(nil, nil)
	e.Get(mock.Anything, iam.GetPermissionRequest{
		RequestObjectId:   computeId,
		RequestObjectType: "serverless-compute",
	}).Return(serverlessComputeAcl(computeId), nil)
}

func groupCanUse() []any {
	return []any{
		map[string]any{
			"group_name":       "data-science",
			"permission_level": "CAN_USE",
		},
	}
}

func assertGroupCanUseOnly(t *testing.T, d *schema.ResourceData) {
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	elem := ac.List()[0].(map[string]any)
	assert.Equal(t, "data-science", elem["group_name"])
	assert.Equal(t, "CAN_USE", elem["permission_level"])
}

func TestResourcePermissionsCreate_ServerlessComputeId(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			expectSetAndGet(mwc, testAutomatedComputeId)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"serverless_compute_id": testAutomatedComputeId,
			"access_control":        groupCanUse(),
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/serverless-compute/"+testAutomatedComputeId, d.Id())
	assert.Equal(t, testAutomatedComputeId, d.Get("serverless_compute_id"))
	assert.Equal(t, "serverless-compute", d.Get("object_type"))
	// The applying user and the inherited admins grant stay out of the state.
	assertGroupCanUseOnly(t, d)
}

func TestResourcePermissionsCreate_DefaultServerlessComputeInteractive(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockServerlessComputeServiceAPI().EXPECT().ListAll(mock.Anything, compute.ListServerlessComputesRequest{}).Return([]compute.ServerlessCompute{
				{ServerlessComputeId: testInteractiveComputeId, Name: "Default Interactive Compute", IsDefault: true},
			}, nil)
			expectSetAndGet(mwc, testInteractiveComputeId)
		},
		Resource: ResourcePermissions(),
		HCL: `
		default_serverless_compute = "INTERACTIVE"
		access_control {
			group_name       = "data-science"
			permission_level = "CAN_USE"
		}
		`,
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/serverless-compute/"+testInteractiveComputeId, d.Id())
	assert.Equal(t, "INTERACTIVE", d.Get("default_serverless_compute"))
	// Configured by type, so the ID attribute must stay empty to avoid a diff on the next plan.
	assert.Equal(t, "", d.Get("serverless_compute_id"))
	assertGroupCanUseOnly(t, d)
}

func TestResourcePermissionsCreate_DefaultServerlessComputeAutomated(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockServerlessComputeServiceAPI().EXPECT().ListComputePoliciesAll(mock.Anything, compute.ListComputePoliciesRequest{}).Return([]compute.ComputePolicy{
				// The default Lakebase policy is also a default, but doesn't allow automated compute.
				{
					ComputePolicyId:        "lakebase-default",
					IsDefault:              true,
					AllowedComputeProducts: &compute.AllowedComputeProducts{Lakebase: &compute.LakebaseConstraints{}},
				},
				{
					ComputePolicyId:        "custom-automated",
					AllowedComputeProducts: &compute.AllowedComputeProducts{Automated: &compute.AutomatedConstraints{}},
				},
				{
					ComputePolicyId:        testAutomatedComputeId,
					DisplayName:            "Default Automated Compute",
					IsDefault:              true,
					AllowedComputeProducts: &compute.AllowedComputeProducts{Automated: &compute.AutomatedConstraints{}},
				},
			}, nil)
			expectSetAndGet(mwc, testAutomatedComputeId)
		},
		Resource: ResourcePermissions(),
		HCL: `
		default_serverless_compute = "AUTOMATED"
		access_control {
			group_name       = "data-science"
			permission_level = "CAN_USE"
		}
		`,
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/serverless-compute/"+testAutomatedComputeId, d.Id())
	assert.Equal(t, "", d.Get("serverless_compute_id"))
	assertGroupCanUseOnly(t, d)
}

func TestResourcePermissionsCreate_DefaultServerlessComputeNotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockServerlessComputeServiceAPI().EXPECT().ListAll(mock.Anything, compute.ListServerlessComputesRequest{}).Return([]compute.ServerlessCompute{
				{ServerlessComputeId: "custom", Name: "custom"},
			}, nil)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"default_serverless_compute": "INTERACTIVE",
			"access_control":             groupCanUse(),
		},
		Create: true,
	}.ExpectError(t, "default interactive serverless compute not found in this workspace")
}

func TestResourcePermissionsCreate_DefaultServerlessComputeListError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockServerlessComputeServiceAPI().EXPECT().ListComputePoliciesAll(mock.Anything, compute.ListComputePoliciesRequest{}).Return(nil, &apierr.APIError{
				ErrorCode:  "PERMISSION_DENIED",
				Message:    "serverless compute is not enabled",
				StatusCode: 403,
			})
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"default_serverless_compute": "AUTOMATED",
			"access_control":             groupCanUse(),
		},
		Create: true,
	}.ExpectError(t, "cannot list compute policies: serverless compute is not enabled")
}

func TestResourcePermissions_DefaultServerlessComputeInvalidValue(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourcePermissions(),
		HCL: `
		default_serverless_compute = "interactive"
		access_control {
			group_name       = "data-science"
			permission_level = "CAN_USE"
		}
		`,
		Create: true,
	}.Apply(t)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "expected default_serverless_compute to be one of")
}

func TestGetDefaultServerlessComputeID_UnsupportedType(t *testing.T) {
	_, err := getDefaultServerlessComputeID(context.Background(), nil, "SQL")
	assert.EqualError(t, err, `unsupported default_serverless_compute "SQL"; allowed values: INTERACTIVE, AUTOMATED`)
}

func TestResourcePermissionsRead_ServerlessCompute(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   testInteractiveComputeId,
				RequestObjectType: "serverless-compute",
			}).Return(serverlessComputeAcl(testInteractiveComputeId), nil)
		},
		Resource: ResourcePermissions(),
		Read:     true,
		New:      true,
		ID:       "/serverless-compute/" + testInteractiveComputeId,
	}.Apply(t)
	require.NoError(t, err)
	// Imports resolve to serverless_compute_id.
	assert.Equal(t, testInteractiveComputeId, d.Get("serverless_compute_id"))
	assertGroupCanUseOnly(t, d)
}

func TestResourcePermissionsDelete_ServerlessCompute(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   testInteractiveComputeId,
				RequestObjectType: "serverless-compute",
			}).Return(serverlessComputeAcl(testInteractiveComputeId), nil)
			// Deleting leaves only the inherited admins grant and the applying user.
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   testInteractiveComputeId,
				RequestObjectType: "serverless-compute",
				AccessControlList: []iam.AccessControlRequest{
					{UserName: TestingAdminUser, PermissionLevel: "CAN_MANAGE"},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/serverless-compute/" + testInteractiveComputeId,
	}.ApplyNoError(t)
}

func TestResourcePermissionsUpdate_DefaultServerlessCompute(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			// Update works from the stored ID, so the default compute isn't looked up again.
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   testInteractiveComputeId,
				RequestObjectType: "serverless-compute",
				AccessControlList: []iam.AccessControlRequest{
					{GroupName: "users", PermissionLevel: "CAN_USE"},
					{UserName: TestingAdminUser, PermissionLevel: "CAN_MANAGE"},
				},
			}).Return(nil, nil)
			acl := serverlessComputeAcl(testInteractiveComputeId)
			acl.AccessControlList[0].GroupName = "users"
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   testInteractiveComputeId,
				RequestObjectType: "serverless-compute",
			}).Return(acl, nil)
		},
		InstanceState: map[string]string{
			"default_serverless_compute": "INTERACTIVE",
		},
		HCL: `
		default_serverless_compute = "INTERACTIVE"
		access_control {
			group_name       = "users"
			permission_level = "CAN_USE"
		}
		`,
		Resource: ResourcePermissions(),
		Update:   true,
		ID:       "/serverless-compute/" + testInteractiveComputeId,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/serverless-compute/"+testInteractiveComputeId, d.Id())
	assert.Equal(t, "", d.Get("serverless_compute_id"))
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	assert.Equal(t, "users", ac.List()[0].(map[string]any)["group_name"])
}

func TestResourcePermissionsRead_DefaultServerlessComputeInState(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   testAutomatedComputeId,
				RequestObjectType: "serverless-compute",
			}).Return(serverlessComputeAcl(testAutomatedComputeId), nil)
		},
		InstanceState: map[string]string{
			"default_serverless_compute": "AUTOMATED",
		},
		HCL: `
		default_serverless_compute = "AUTOMATED"
		access_control {
			group_name       = "data-science"
			permission_level = "CAN_USE"
		}
		`,
		Resource: ResourcePermissions(),
		Read:     true,
		ID:       "/serverless-compute/" + testAutomatedComputeId,
	}.Apply(t)
	require.NoError(t, err)
	// A refresh of a resource configured by type keeps the type and doesn't add the ID.
	assert.Equal(t, "AUTOMATED", d.Get("default_serverless_compute"))
	assert.Equal(t, "", d.Get("serverless_compute_id"))
	assertGroupCanUseOnly(t, d)
}
