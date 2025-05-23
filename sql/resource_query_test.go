package sql

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

var (
	queryResponse = sql.Query{
		Id:            "7890",
		WarehouseId:   "123456",
		DisplayName:   "TF new query",
		OwnerUserName: "user@domain.com",
		ParentPath:    "/Workspace/Shared/Querys",
		QueryText:     "select 42 as value",
	}
	createQueryHcl = `warehouse_id = "123456"
  query_text = "select 42 as value"
  display_name = "TF new query"
  parent_path = "/Shared/Querys"
  owner_user_name = "user@domain.com"
`
	createQueryRequest = sql.CreateQueryRequest{
		AutoResolveDisplayName: false,
		Query: &sql.CreateQueryRequestQuery{
			WarehouseId: "123456",
			QueryText:   "select 42 as value",
			DisplayName: "TF new query",
			ParentPath:  "/Shared/Querys",
		},
		ForceSendFields: []string{"AutoResolveDisplayName"},
	}
)

func TestQueryCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQueriesAPI().EXPECT()
			e.Create(mock.Anything, createQueryRequest).Return(&queryResponse, nil)
			e.Update(mock.Anything, sql.UpdateQueryRequest{
				Id:         "7890",
				UpdateMask: "owner_user_name",
				Query: &sql.UpdateQueryRequestQuery{
					OwnerUserName: "user@domain.com",
				},
			}).Return(&queryResponse, nil)
			e.GetById(mock.Anything, "7890").Return(&queryResponse, nil)
		},
		Resource: ResourceQuery(),
		Create:   true,
		HCL:      createQueryHcl,
	}.ApplyAndExpectData(t, map[string]any{
		"id":              "7890",
		"warehouse_id":    "123456",
		"display_name":    "TF new query",
		"owner_user_name": "user@domain.com",
	})
}

func TestQueryCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQueriesAPI().EXPECT()
			e.Create(mock.Anything, createQueryRequest).Return(nil, &apierr.APIError{
				StatusCode: http.StatusBadRequest,
				Message:    "Node named 'TF new query' already exists",
			})
		},
		Resource: ResourceQuery(),
		Create:   true,
		HCL:      createQueryHcl,
	}.ExpectError(t, "Node named 'TF new query' already exists")
}

func TestQueryRead_Import(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockQueriesAPI().EXPECT().GetById(mock.Anything, "7890").Return(&queryResponse, nil)
		},
		Resource: ResourceQuery(),
		Read:     true,
		ID:       "7890",
		New:      true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":              "7890",
		"warehouse_id":    "123456",
		"query_text":      "select 42 as value",
		"display_name":    "TF new query",
		"owner_user_name": "user@domain.com",
	})
}

func TestQueryRead_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockQueriesAPI().EXPECT().GetById(mock.Anything, "7890").Return(nil, &apierr.APIError{
				StatusCode: http.StatusBadRequest,
				Message:    "bad payload",
			})
		},
		Resource: ResourceQuery(),
		Read:     true,
		ID:       "7890",
		New:      true,
	}.ExpectError(t, "bad payload")
}

func TestQueryDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockQueriesAPI().EXPECT().DeleteById(mock.Anything, "7890").Return(nil)
		},
		Resource: ResourceQuery(),
		Delete:   true,
		ID:       "7890",
		New:      true,
	}.ApplyNoError(t)
}

func TestQueryUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockQueriesAPI().EXPECT()
			e.Update(mock.Anything, sql.UpdateQueryRequest{
				Id:                     "7890",
				UpdateMask:             "display_name,query_text,warehouse_id,parameters,owner_user_name",
				AutoResolveDisplayName: false,
				ForceSendFields:        []string{"AutoResolveDisplayName"},
				Query: &sql.UpdateQueryRequestQuery{
					WarehouseId:   "123456",
					DisplayName:   "TF new query",
					OwnerUserName: "user@domain.com",
					QueryText:     "select 42 as value",
				}}).Return(&queryResponse, nil)
			e.GetById(mock.Anything, "7890").Return(&queryResponse, nil)
		},
		Resource: ResourceQuery(),
		Update:   true,
		ID:       "7890",
		HCL: `warehouse_id = "123456"
  query_text = "select 42 as value"
  display_name = "TF new query"
  owner_user_name = "user@domain.com"
`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":              "7890",
		"warehouse_id":    "123456",
		"query_text":      "select 42 as value",
		"display_name":    "TF new query",
		"owner_user_name": "user@domain.com",
	})
}
