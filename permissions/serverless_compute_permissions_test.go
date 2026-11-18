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

	serverlessComputeType = "serverless-compute"
	computePolicyType     = "compute-policy"
)

// computeAcl is what the Permissions API returns for a serverless compute or compute policy after
// the provider grants a group CAN_USE: the group, the applying user, and the inherited admins grant.
func computeAcl(objectType, id, group string) *iam.ObjectPermissions {
	return &iam.ObjectPermissions{
		ObjectId:   "/" + objectType + "/" + id,
		ObjectType: objectType,
		AccessControlList: []iam.AccessControlResponse{
			{
				GroupName:      group,
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
					InheritedFromObject: []string{"/" + objectType},
				}},
			},
		},
	}
}

// expectSetAndGet mocks the Permissions API calls made by create or update: a full-replace Set of
// the group grant plus the applying user as CAN_MANAGE, then a Get to read the result back.
func expectSetAndGet(mwc *mocks.MockWorkspaceClient, objectType, id, group string) {
	mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
	e := mwc.GetMockPermissionsAPI().EXPECT()
	e.Set(mock.Anything, iam.SetObjectPermissions{
		RequestObjectId:   id,
		RequestObjectType: objectType,
		AccessControlList: []iam.AccessControlRequest{
			{GroupName: group, PermissionLevel: "CAN_USE"},
			{UserName: TestingAdminUser, PermissionLevel: "CAN_MANAGE"},
		},
	}).Return(nil, nil)
	e.Get(mock.Anything, iam.GetPermissionRequest{
		RequestObjectId:   id,
		RequestObjectType: objectType,
	}).Return(computeAcl(objectType, id, group), nil)
}

func expectListServerlessComputes(mwc *mocks.MockWorkspaceClient, computes ...compute.ServerlessCompute) {
	mwc.GetMockServerlessComputeServiceAPI().EXPECT().ListAll(mock.Anything, compute.ListServerlessComputesRequest{}).Return(computes, nil)
}

func expectListComputePolicies(mwc *mocks.MockWorkspaceClient, policies ...compute.ComputePolicy) {
	mwc.GetMockServerlessComputeServiceAPI().EXPECT().ListComputePoliciesAll(mock.Anything, compute.ListComputePoliciesRequest{}).Return(policies, nil)
}

var (
	defaultInteractiveCompute = compute.ServerlessCompute{
		ServerlessComputeId: testInteractiveComputeId, Name: "Default Interactive Compute", IsDefault: true,
	}
	// The default Lakebase policy is also a default, but doesn't allow automated compute.
	defaultLakebasePolicy = compute.ComputePolicy{
		ComputePolicyId:        "lakebase-default",
		IsDefault:              true,
		AllowedComputeProducts: &compute.AllowedComputeProducts{Lakebase: &compute.LakebaseConstraints{}},
	}
	customAutomatedPolicy = compute.ComputePolicy{
		ComputePolicyId:        "custom-automated",
		AllowedComputeProducts: &compute.AllowedComputeProducts{Automated: &compute.AutomatedConstraints{}},
	}
	defaultAutomatedPolicy = compute.ComputePolicy{
		ComputePolicyId:        testAutomatedComputeId,
		DisplayName:            "Default Automated Compute",
		IsDefault:              true,
		AllowedComputeProducts: &compute.AllowedComputeProducts{Automated: &compute.AutomatedConstraints{}},
	}
)

func groupCanUse(group string) []any {
	return []any{
		map[string]any{
			"group_name":       group,
			"permission_level": "CAN_USE",
		},
	}
}

func assertGroupCanUseOnly(t *testing.T, d *schema.ResourceData, group string) {
	ac := d.Get("access_control").(*schema.Set)
	require.Equal(t, 1, len(ac.List()))
	elem := ac.List()[0].(map[string]any)
	assert.Equal(t, group, elem["group_name"])
	assert.Equal(t, "CAN_USE", elem["permission_level"])
}

func TestResourcePermissionsCreate_ServerlessComputeId(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			expectSetAndGet(mwc, serverlessComputeType, testInteractiveComputeId, "data-science")
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"serverless_compute_id": testInteractiveComputeId,
			"access_control":        groupCanUse("data-science"),
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/serverless-compute/"+testInteractiveComputeId, d.Id())
	assert.Equal(t, testInteractiveComputeId, d.Get("serverless_compute_id"))
	assert.Equal(t, serverlessComputeType, d.Get("object_type"))
	// The applying user and the inherited admins grant stay out of the state.
	assertGroupCanUseOnly(t, d, "data-science")
}

func TestResourcePermissionsCreate_ComputePolicyId(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			expectSetAndGet(mwc, computePolicyType, testAutomatedComputeId, "data-science")
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"compute_policy_id": testAutomatedComputeId,
			"access_control":    groupCanUse("data-science"),
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/compute-policy/"+testAutomatedComputeId, d.Id())
	assert.Equal(t, testAutomatedComputeId, d.Get("compute_policy_id"))
	assert.Equal(t, computePolicyType, d.Get("object_type"))
	assertGroupCanUseOnly(t, d, "data-science")
}

func TestResourcePermissionsCreate_DefaultServerlessComputeInteractive(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			expectListServerlessComputes(mwc, defaultInteractiveCompute)
			expectSetAndGet(mwc, serverlessComputeType, testInteractiveComputeId, "data-science")
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
	assertGroupCanUseOnly(t, d, "data-science")
}

func TestResourcePermissionsCreate_DefaultComputePolicyAutomated(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			expectListComputePolicies(mwc, defaultLakebasePolicy, customAutomatedPolicy, defaultAutomatedPolicy)
			expectSetAndGet(mwc, computePolicyType, testAutomatedComputeId, "data-science")
		},
		Resource: ResourcePermissions(),
		HCL: `
		default_compute_policy = "AUTOMATED"
		access_control {
			group_name       = "data-science"
			permission_level = "CAN_USE"
		}
		`,
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/compute-policy/"+testAutomatedComputeId, d.Id())
	assert.Equal(t, "AUTOMATED", d.Get("default_compute_policy"))
	assert.Equal(t, "", d.Get("compute_policy_id"))
	assertGroupCanUseOnly(t, d, "data-science")
}

func TestResourcePermissionsCreate_DefaultServerlessComputeNotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			expectListServerlessComputes(mwc, compute.ServerlessCompute{ServerlessComputeId: "custom", Name: "custom"})
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"default_serverless_compute": "INTERACTIVE",
			"access_control":             groupCanUse("data-science"),
		},
		Create: true,
	}.ExpectError(t, "default interactive serverless compute not found in this workspace")
}

func TestResourcePermissionsCreate_DefaultComputePolicyNotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			expectListComputePolicies(mwc, defaultLakebasePolicy, customAutomatedPolicy)
		},
		Resource: ResourcePermissions(),
		State: map[string]any{
			"default_compute_policy": "AUTOMATED",
			"access_control":         groupCanUse("data-science"),
		},
		Create: true,
	}.ExpectError(t, "default automated compute policy not found in this workspace")
}

func TestResourcePermissionsCreate_DefaultComputePolicyListError(t *testing.T) {
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
			"default_compute_policy": "AUTOMATED",
			"access_control":         groupCanUse("data-science"),
		},
		Create: true,
	}.ExpectError(t, "cannot list compute policies: serverless compute is not enabled")
}

func TestResourcePermissions_DefaultLookupInvalidValues(t *testing.T) {
	for field, value := range map[string]string{
		"default_serverless_compute": "AUTOMATED",
		"default_compute_policy":     "automated",
	} {
		t.Run(field, func(t *testing.T) {
			_, err := qa.ResourceFixture{
				Resource: ResourcePermissions(),
				HCL: field + ` = "` + value + `"
				access_control {
					group_name       = "data-science"
					permission_level = "CAN_USE"
				}
				`,
				Create: true,
			}.Apply(t)
			require.Error(t, err)
			assert.Contains(t, err.Error(), "expected "+field+" to be one of")
		})
	}
}

func TestGetDefaultIDs_UnsupportedValues(t *testing.T) {
	_, err := getDefaultServerlessComputeID(context.Background(), nil, "AUTOMATED")
	assert.EqualError(t, err, `unsupported default_serverless_compute "AUTOMATED"; allowed values: INTERACTIVE`)
	_, err = getDefaultComputePolicyID(context.Background(), nil, "LAKEBASE")
	assert.EqualError(t, err, `unsupported default_compute_policy "LAKEBASE"; allowed values: AUTOMATED`)
}

func TestResourcePermissionsUpdate_DefaultComputePolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			// Update works from the stored ID, so the default policy isn't looked up again.
			expectSetAndGet(mwc, computePolicyType, testAutomatedComputeId, "users")
		},
		InstanceState: map[string]string{
			"default_compute_policy": "AUTOMATED",
		},
		HCL: `
		default_compute_policy = "AUTOMATED"
		access_control {
			group_name       = "users"
			permission_level = "CAN_USE"
		}
		`,
		Resource: ResourcePermissions(),
		Update:   true,
		ID:       "/compute-policy/" + testAutomatedComputeId,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/compute-policy/"+testAutomatedComputeId, d.Id())
	assert.Equal(t, "", d.Get("compute_policy_id"))
	assertGroupCanUseOnly(t, d, "users")
}

func TestResourcePermissionsRead_DefaultServerlessComputeInState(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   testInteractiveComputeId,
				RequestObjectType: serverlessComputeType,
			}).Return(computeAcl(serverlessComputeType, testInteractiveComputeId, "data-science"), nil)
		},
		InstanceState: map[string]string{
			"default_serverless_compute": "INTERACTIVE",
		},
		HCL: `
		default_serverless_compute = "INTERACTIVE"
		access_control {
			group_name       = "data-science"
			permission_level = "CAN_USE"
		}
		`,
		Resource: ResourcePermissions(),
		Read:     true,
		ID:       "/serverless-compute/" + testInteractiveComputeId,
	}.Apply(t)
	require.NoError(t, err)
	// A refresh of a resource configured by type keeps the type and doesn't add the ID.
	assert.Equal(t, "INTERACTIVE", d.Get("default_serverless_compute"))
	assert.Equal(t, "", d.Get("serverless_compute_id"))
	assertGroupCanUseOnly(t, d, "data-science")
}

func TestResourcePermissionsRead_Import(t *testing.T) {
	for objectType, field := range map[string]string{
		serverlessComputeType: "serverless_compute_id",
		computePolicyType:     "compute_policy_id",
	} {
		t.Run(objectType, func(t *testing.T) {
			d, err := qa.ResourceFixture{
				MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
					mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
					mwc.GetMockPermissionsAPI().EXPECT().Get(mock.Anything, iam.GetPermissionRequest{
						RequestObjectId:   "abc",
						RequestObjectType: objectType,
					}).Return(computeAcl(objectType, "abc", "data-science"), nil)
				},
				Resource: ResourcePermissions(),
				Read:     true,
				New:      true,
				ID:       "/" + objectType + "/abc",
			}.Apply(t)
			require.NoError(t, err)
			// Imports resolve to the ID field for the object type.
			assert.Equal(t, "abc", d.Get(field))
			assertGroupCanUseOnly(t, d, "data-science")
		})
	}
}

func TestResourcePermissionsDelete_ComputePolicy(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockCurrentUserAPI().EXPECT().Me(mock.Anything, mock.Anything).Return(&iam.User{UserName: TestingAdminUser}, nil)
			e := mwc.GetMockPermissionsAPI().EXPECT()
			e.Get(mock.Anything, iam.GetPermissionRequest{
				RequestObjectId:   testAutomatedComputeId,
				RequestObjectType: computePolicyType,
			}).Return(computeAcl(computePolicyType, testAutomatedComputeId, "data-science"), nil)
			// Deleting leaves only the inherited admins grant and the applying user.
			e.Set(mock.Anything, iam.SetObjectPermissions{
				RequestObjectId:   testAutomatedComputeId,
				RequestObjectType: computePolicyType,
				AccessControlList: []iam.AccessControlRequest{
					{UserName: TestingAdminUser, PermissionLevel: "CAN_MANAGE"},
				},
			}).Return(nil, nil)
		},
		Resource: ResourcePermissions(),
		Delete:   true,
		ID:       "/compute-policy/" + testAutomatedComputeId,
	}.ApplyNoError(t)
}
