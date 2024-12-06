package catalog

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var onlineTableHcl = `
name = "main.default.online_table"
spec {
	source_table_full_name = "main.default.test"
	primary_key_columns = [
	  "id"
	]
	run_triggered {
	}
  }
`

func TestOnlineTableCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceOnlineTable())
}

func TestOnlineTableCreate(t *testing.T) {
	otStatusOnline := &catalog.OnlineTable{
		Name: "main.default.online_table",
		Spec: &catalog.OnlineTableSpec{
			RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
			SourceTableFullName: "main.default.test",
			PrimaryKeyColumns:   []string{"id"},
		},
		Status: &catalog.OnlineTableStatus{DetailedState: catalog.OnlineTableStateOnline},
	}
	otStatusNotSet := &catalog.OnlineTable{
		Name: "main.default.online_table",
		Spec: &catalog.OnlineTableSpec{
			RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
			SourceTableFullName: "main.default.test",
			PrimaryKeyColumns:   []string{"id"},
		},
	}
	otStatusNotSetWait := &catalog.WaitGetOnlineTableActive[catalog.OnlineTable]{
		Response: otStatusNotSet,
		Name:     "main.default.online_table",
		Poll: func(d time.Duration, f func(*catalog.OnlineTable)) (*catalog.OnlineTable, error) {
			return otStatusOnline, nil
		},
	}
	// otStatusUnknown := &catalog.OnlineTable{
	// 	Name: "main.default.online_table",
	// 	Spec: &catalog.OnlineTableSpec{
	// 		RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
	// 		SourceTableFullName: "main.default.test",
	// 		PrimaryKeyColumns:   []string{"id"},
	// 	},
	// 	Status: &catalog.OnlineTableStatus{},
	// }
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockOnlineTablesAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateOnlineTableRequest{
				Table: &catalog.OnlineTable{
					Name: "main.default.online_table",
					Spec: &catalog.OnlineTableSpec{
						RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
						SourceTableFullName: "main.default.test",
						PrimaryKeyColumns:   []string{"id"},
					},
				},
			}).Return(otStatusNotSetWait, nil)
			e.GetByName(mock.Anything, "main.default.online_table").Return(otStatusOnline, nil)
		},
		Resource: ResourceOnlineTable(),
		HCL:      onlineTableHcl,
		Create:   true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": "main.default.online_table",
	})
}

func TestOnlineTableCreate_ErrorImmediately(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockOnlineTablesAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateOnlineTableRequest{
				Table: &catalog.OnlineTable{
					Name: "main.default.online_table",
					Spec: &catalog.OnlineTableSpec{
						RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
						SourceTableFullName: "main.default.test",
						PrimaryKeyColumns:   []string{"id"},
					},
				},
			}).Return(nil, fmt.Errorf("error!"))
		},
		Resource: ResourceOnlineTable(),
		HCL:      onlineTableHcl,
		Create:   true,
	}.ExpectError(t, "error!")
}

func TestOnlineTableCreate_ErrorInWait(t *testing.T) {
	otStatusProvisioning := &catalog.OnlineTable{
		Name: "main.default.online_table",
		Spec: &catalog.OnlineTableSpec{
			RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
			SourceTableFullName: "main.default.test",
			PrimaryKeyColumns:   []string{"id"},
		},
		Status: &catalog.OnlineTableStatus{DetailedState: catalog.OnlineTableStateProvisioning},
	}
	otStatusErrorWait := &catalog.WaitGetOnlineTableActive[catalog.OnlineTable]{
		Response: otStatusProvisioning,
		Name:     "main.default.online_table",
		Poll: func(d time.Duration, f func(*catalog.OnlineTable)) (*catalog.OnlineTable, error) {
			return nil, errors.New("failed to reach ACTIVE, got OFFLINE_FAILED: error!")
		},
	}
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockOnlineTablesAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateOnlineTableRequest{
				Table: &catalog.OnlineTable{
					Name: "main.default.online_table",
					Spec: &catalog.OnlineTableSpec{
						RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
						SourceTableFullName: "main.default.test",
						PrimaryKeyColumns:   []string{"id"},
					},
				},
			}).Return(otStatusErrorWait, nil)
		},
		Resource: ResourceOnlineTable(),
		HCL:      onlineTableHcl,
		Create:   true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "failed to reach ACTIVE, got OFFLINE_FAILED: error!")
	assert.Equal(t, "main.default.online_table", d.Id())
}

func TestOnlineTableRead(t *testing.T) {
	ot := &catalog.OnlineTable{
		Name: "main.default.online_table",
		Spec: &catalog.OnlineTableSpec{
			RunTriggered:        &catalog.OnlineTableSpecTriggeredSchedulingPolicy{},
			SourceTableFullName: "main.default.test",
			PrimaryKeyColumns:   []string{"id"},
		},
		Status: &catalog.OnlineTableStatus{DetailedState: catalog.OnlineTableStateOnline},
	}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockOnlineTablesAPI().EXPECT()
			e.GetByName(mock.Anything, "main.default.online_table").Return(ot, nil)
		},
		Resource: ResourceOnlineTable(),
		ID:       "main.default.online_table",
		HCL:      onlineTableHcl,
		Read:     true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": "main.default.online_table",
	})
}

func TestOnlineTableRead_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockOnlineTablesAPI().EXPECT()
			e.GetByName(mock.Anything, "main.default.online_table").Return(nil, fmt.Errorf("error!"))
		},
		Resource: ResourceOnlineTable(),
		ID:       "main.default.online_table",
		HCL:      onlineTableHcl,
		Read:     true,
	}.ExpectError(t, "error!")
}

func TestOnlineTableDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockOnlineTablesAPI().EXPECT()
			e.DeleteByName(mock.Anything, "main.default.online_table").Return(nil)
			e.GetByName(mock.Anything, "main.default.online_table").Return(nil, apierr.ErrResourceDoesNotExist)
		},
		Resource: ResourceOnlineTable(),
		ID:       "main.default.online_table",
		HCL:      onlineTableHcl,
		Delete:   true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": "main.default.online_table",
	})
}
