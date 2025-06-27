package catalog

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestSystemSchemaCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			e := w.GetMockSystemSchemasAPI().EXPECT()
			e.Enable(mock.Anything, catalog.EnableRequest{
				MetastoreId: "abc",
				SchemaName:  "access",
			}).Return(nil)
			e.ListByMetastoreId(mock.Anything, "abc").Return(&catalog.ListSystemSchemasResponse{
				Schemas: []catalog.SystemSchemaInfo{
					{
						Schema: "access",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
					{
						Schema: "billing",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
				},
			}, nil)
		},
		Resource: ResourceSystemSchema(),
		HCL:      `schema = "access"`,
		Create:   true,
	}.ApplyAndExpectData(t, map[string]any{
		"schema":    "access",
		"id":        "abc|access",
		"full_name": "system.access",
	})
}

func TestSystemSchemaCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().Enable(mock.Anything, catalog.EnableRequest{
				MetastoreId: "abc",
				SchemaName:  "access",
			}).Return(errors.New("Internal error happened"))
		},
		Resource: ResourceSystemSchema(),
		HCL:      `schema = "access"`,
		Create:   true,
	}.ExpectError(t, "Internal error happened")
}

func TestSystemSchemaCreateAlreadyEnabled(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			e := w.GetMockSystemSchemasAPI().EXPECT()
			e.Enable(mock.Anything, catalog.EnableRequest{
				MetastoreId: "abc",
				SchemaName:  "billing",
			}).Return(errors.New("billing system schema can only be enabled by Databricks"))
			e.ListByMetastoreId(mock.Anything, "abc").Return(&catalog.ListSystemSchemasResponse{
				Schemas: []catalog.SystemSchemaInfo{
					{
						Schema: "access",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
					{
						Schema: "billing",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
				},
			}, nil)
		},
		Resource: ResourceSystemSchema(),
		HCL:      `schema = "billing"`,
		Create:   true,
	}.ApplyAndExpectData(t, map[string]any{
		"schema":    "billing",
		"id":        "abc|billing",
		"full_name": "system.billing",
	})
}

func TestSystemSchemaUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			e := w.GetMockSystemSchemasAPI().EXPECT()
			e.Enable(mock.Anything, catalog.EnableRequest{
				MetastoreId: "abc",
				SchemaName:  "access",
			}).Return(nil)
			e.Disable(mock.Anything, catalog.DisableRequest{
				MetastoreId: "abc",
				SchemaName:  "information_schema",
			}).Return(nil)
			e.ListByMetastoreId(mock.Anything, "abc").Return(&catalog.ListSystemSchemasResponse{
				Schemas: []catalog.SystemSchemaInfo{
					{
						Schema: "access",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
					{
						Schema: "billing",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
				},
			}, nil)
		},
		Resource: ResourceSystemSchema(),
		InstanceState: map[string]string{
			"schema": "information_schema",
		},
		HCL:    `schema = "access"`,
		Update: true,
		ID:     "abc|information_schema",
	}.ApplyAndExpectData(t, map[string]any{
		"schema": "access",
		"id":     "abc|access",
	})
}

func TestSystemSchemaUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().Enable(mock.Anything, catalog.EnableRequest{
				MetastoreId: "abc",
				SchemaName:  "access",
			}).Return(errors.New("Internal error happened"))
		},
		Resource: ResourceSystemSchema(),
		InstanceState: map[string]string{
			"schema": "information_schema",
		},
		HCL:    `schema = "access"`,
		Update: true,
		ID:     "abc|information_schema",
	}.ExpectError(t, "Internal error happened")
}

func TestSystemSchemaRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().ListByMetastoreId(mock.Anything, "abc").Return(&catalog.ListSystemSchemasResponse{
				Schemas: []catalog.SystemSchemaInfo{
					{
						Schema: "access",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
					{
						Schema: "billing",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
				},
			}, nil)
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		ID:       "abc|access",
	}.ApplyAndExpectData(t, map[string]any{
		"schema": "access",
		"state":  string(SystemSchemaInfoStateEnableCompleted),
	})
}

func TestSystemSchemaRead_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().ListByMetastoreId(mock.Anything, "abc").Return(nil, errors.New("Internal error happened"))
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		ID:       "abc|access",
	}.ExpectError(t, "Internal error happened")
}

func TestSystemSchemaRead_NotEnabled(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().ListByMetastoreId(mock.Anything, "abc").Return(&catalog.ListSystemSchemasResponse{
				Schemas: []catalog.SystemSchemaInfo{
					{
						Schema: "access",
						State:  string(SystemSchemaInfoStateAvailable),
					},
					{
						Schema: "billing",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
				},
			}, nil)
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		Removed:  true,
		ID:       "abc|access",
	}.ApplyAndExpectData(t, map[string]any{})
}

func TestSystemSchemaRead_NotExists(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().ListByMetastoreId(mock.Anything, "abc").Return(&catalog.ListSystemSchemasResponse{
				Schemas: []catalog.SystemSchemaInfo{
					{
						Schema: "billing",
						State:  string(SystemSchemaInfoStateEnableCompleted),
					},
				},
			}, nil)
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		Removed:  true,
		ID:       "abc|access",
	}.ApplyAndExpectData(t, map[string]any{"id": ""})
}

func TestSystemSchemaDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			e := w.GetMockSystemSchemasAPI().EXPECT()
			e.Disable(mock.Anything, catalog.DisableRequest{
				MetastoreId: "abc",
				SchemaName:  "access",
			}).Return(nil)
		},
		HCL:      `schema = "access"`,
		Resource: ResourceSystemSchema(),
		Delete:   true,
		ID:       "abc|access",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "abc|access",
	})
}

func TestSystemSchemaDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().Disable(mock.Anything, catalog.DisableRequest{
				MetastoreId: "abc",
				SchemaName:  "access",
			}).Return(errors.New("Internal error happened"))
		},
		Resource: ResourceSystemSchema(),
		HCL:      `schema = "access"`,
		Delete:   true,
		ID:       "abc|access",
	}.ExpectError(t, "Internal error happened")
}

func TestSystemSchemaDelete_DisabledByDatabricks(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Summary(mock.Anything).Return(&catalog.GetMetastoreSummaryResponse{
				MetastoreId: "abc",
			}, nil)
			w.GetMockSystemSchemasAPI().EXPECT().Disable(mock.Anything, catalog.DisableRequest{
				MetastoreId: "abc",
				SchemaName:  "access",
			}).Return(errors.New("access system schema can only be disabled by Databricks"))
		},
		Resource: ResourceSystemSchema(),
		HCL:      `schema = "access"`,
		Delete:   true,
		ID:       "abc|access",
	}.ApplyAndExpectData(t, map[string]any{})
}
