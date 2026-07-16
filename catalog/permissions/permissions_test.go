package permissions

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newTestAPI(t *testing.T) (UnityCatalogPermissionsAPI, *mocks.MockWorkspaceClient) {
	mw := mocks.NewMockWorkspaceClient(t)
	return NewUnityCatalogPermissionsAPI(context.Background(), mw.WorkspaceClient), mw
}

func TestUpdatePermissions_EmptyDiff(t *testing.T) {
	api, _ := newTestAPI(t)
	// no expectations — zero diff means zero API calls
	assert.NoError(t, api.UpdatePermissions(catalog.SecurableType("table"), "foo.bar", nil))
}

func TestUpdatePermissions_AddOnly(t *testing.T) {
	api, mw := newTestAPI(t)
	mw.GetMockGrantsAPI().EXPECT().
		Update(mock.Anything, catalog.UpdatePermissions{
			SecurableType: "table",
			FullName:      "foo.bar",
			Changes:       []catalog.PermissionsChange{{Principal: "me", Add: []catalog.Privilege{"SELECT"}}},
		}).Return(&catalog.UpdatePermissionsResponse{}, nil)

	err := api.UpdatePermissions(catalog.SecurableType("table"), "foo.bar", []catalog.PermissionsChange{
		{Principal: "me", Add: []catalog.Privilege{"SELECT"}},
	})
	assert.NoError(t, err)
}

func TestUpdatePermissions_RemoveOnly(t *testing.T) {
	api, mw := newTestAPI(t)
	mw.GetMockGrantsAPI().EXPECT().
		Update(mock.Anything, catalog.UpdatePermissions{
			SecurableType: "table",
			FullName:      "foo.bar",
			Changes:       []catalog.PermissionsChange{{Principal: "me", Remove: []catalog.Privilege{"SELECT"}}},
		}).Return(&catalog.UpdatePermissionsResponse{}, nil)

	err := api.UpdatePermissions(catalog.SecurableType("table"), "foo.bar", []catalog.PermissionsChange{
		{Principal: "me", Remove: []catalog.Privilege{"SELECT"}},
	})
	assert.NoError(t, err)
}

func TestUpdatePermissions_SwapPrivilege_RemovesBeforeAdds(t *testing.T) {
	// The Databricks API rejects a single call that both adds and removes privileges for the
	// same principal. Verify the fix sends removes in a first PATCH, then adds in a second.
	api, mw := newTestAPI(t)
	var callOrder []string

	mw.GetMockGrantsAPI().EXPECT().
		Update(mock.Anything, catalog.UpdatePermissions{
			SecurableType: "schema",
			FullName:      "main.myschema",
			Changes:       []catalog.PermissionsChange{{Principal: "me", Remove: []catalog.Privilege{"SELECT"}}},
		}).RunAndReturn(func(_ context.Context, _ catalog.UpdatePermissions) (*catalog.UpdatePermissionsResponse, error) {
		callOrder = append(callOrder, "remove")
		return &catalog.UpdatePermissionsResponse{}, nil
	})

	mw.GetMockGrantsAPI().EXPECT().
		Update(mock.Anything, catalog.UpdatePermissions{
			SecurableType: "schema",
			FullName:      "main.myschema",
			Changes:       []catalog.PermissionsChange{{Principal: "me", Add: []catalog.Privilege{"MODIFY"}}},
		}).RunAndReturn(func(_ context.Context, _ catalog.UpdatePermissions) (*catalog.UpdatePermissionsResponse, error) {
		callOrder = append(callOrder, "add")
		return &catalog.UpdatePermissionsResponse{}, nil
	})

	err := api.UpdatePermissions(catalog.SecurableType("schema"), "main.myschema", []catalog.PermissionsChange{
		{Principal: "me", Add: []catalog.Privilege{"MODIFY"}, Remove: []catalog.Privilege{"SELECT"}},
	})
	assert.NoError(t, err)
	assert.Equal(t, []string{"remove", "add"}, callOrder, "removes must be sent before adds")
}

func TestUpdatePermissions_MultiPrincipalBatching(t *testing.T) {
	// All removes across multiple principals are batched into one PATCH;
	// all adds (including from a principal with no removes) into another.
	api, mw := newTestAPI(t)
	e := mw.GetMockGrantsAPI().EXPECT()

	e.Update(mock.Anything, catalog.UpdatePermissions{
		SecurableType: "catalog",
		FullName:      "mycat",
		Changes: []catalog.PermissionsChange{
			{Principal: "alice", Remove: []catalog.Privilege{"SELECT"}},
			{Principal: "bob", Remove: []catalog.Privilege{"MODIFY"}},
		},
	}).Return(&catalog.UpdatePermissionsResponse{}, nil)

	e.Update(mock.Anything, catalog.UpdatePermissions{
		SecurableType: "catalog",
		FullName:      "mycat",
		Changes: []catalog.PermissionsChange{
			{Principal: "alice", Add: []catalog.Privilege{"MODIFY"}},
			{Principal: "bob", Add: []catalog.Privilege{"SELECT"}},
			{Principal: "carol", Add: []catalog.Privilege{"USE_CATALOG"}},
		},
	}).Return(&catalog.UpdatePermissionsResponse{}, nil)

	err := api.UpdatePermissions(catalog.SecurableType("catalog"), "mycat", []catalog.PermissionsChange{
		{Principal: "alice", Add: []catalog.Privilege{"MODIFY"}, Remove: []catalog.Privilege{"SELECT"}},
		{Principal: "bob", Add: []catalog.Privilege{"SELECT"}, Remove: []catalog.Privilege{"MODIFY"}},
		{Principal: "carol", Add: []catalog.Privilege{"USE_CATALOG"}},
	})
	assert.NoError(t, err)
}

func TestUpdatePermissions_RemoveError_AddNotCalled(t *testing.T) {
	// If the remove PATCH fails the error is returned immediately and the add PATCH is not sent.
	api, mw := newTestAPI(t)
	mw.GetMockGrantsAPI().EXPECT().
		Update(mock.Anything, catalog.UpdatePermissions{
			SecurableType: "table",
			FullName:      "foo.bar",
			Changes:       []catalog.PermissionsChange{{Principal: "me", Remove: []catalog.Privilege{"SELECT"}}},
		}).Return(nil, errors.New("remove failed"))
	// No expectation for the add call — testify will fail the test if it is called.

	err := api.UpdatePermissions(catalog.SecurableType("table"), "foo.bar", []catalog.PermissionsChange{
		{Principal: "me", Add: []catalog.Privilege{"MODIFY"}, Remove: []catalog.Privilege{"SELECT"}},
	})
	assert.ErrorContains(t, err, "remove failed")
}

func TestUpdatePermissions_AddError_Propagated(t *testing.T) {
	api, mw := newTestAPI(t)
	e := mw.GetMockGrantsAPI().EXPECT()

	e.Update(mock.Anything, catalog.UpdatePermissions{
		SecurableType: "table",
		FullName:      "foo.bar",
		Changes:       []catalog.PermissionsChange{{Principal: "me", Remove: []catalog.Privilege{"SELECT"}}},
	}).Return(&catalog.UpdatePermissionsResponse{}, nil)

	e.Update(mock.Anything, catalog.UpdatePermissions{
		SecurableType: "table",
		FullName:      "foo.bar",
		Changes:       []catalog.PermissionsChange{{Principal: "me", Add: []catalog.Privilege{"MODIFY"}}},
	}).Return(nil, errors.New("add failed"))

	err := api.UpdatePermissions(catalog.SecurableType("table"), "foo.bar", []catalog.PermissionsChange{
		{Principal: "me", Add: []catalog.Privilege{"MODIFY"}, Remove: []catalog.Privilege{"SELECT"}},
	})
	assert.ErrorContains(t, err, "add failed")
}
