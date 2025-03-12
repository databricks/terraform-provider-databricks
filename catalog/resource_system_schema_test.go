package catalog

import (
	"errors"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSystemSchemaCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access",
				Status:   200,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas?",
				Response: catalog.ListSystemSchemasResponse{
					Schemas: []catalog.SystemSchemaInfo{
						{
							Schema: "access",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
						{
							Schema: "billing",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
					},
				},
			},
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		HCL:      `schema = "access"`,
		Create:   true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
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
						State:  catalog.SystemSchemaInfoStateEnableCompleted,
					},
					{
						Schema: "billing",
						State:  catalog.SystemSchemaInfoStateEnableCompleted,
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access",
				Status:   200,
			},
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/information_schema?",
				Status:   200,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas?",
				Response: catalog.ListSystemSchemasResponse{
					Schemas: []catalog.SystemSchemaInfo{
						{
							Schema: "access",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
						{
							Schema: "billing",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
					},
				},
			},
		},
		Resource: ResourceSystemSchema(),
		InstanceState: map[string]string{
			"schema": "information_schema",
		},
		HCL:    `schema = "access"`,
		Update: true,
		ID:     "abc|information_schema",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|access", d.Id())
	assert.Equal(t, "access", d.Get("schema"))
}

func TestSystemSchemaUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		InstanceState: map[string]string{
			"schema": "information_schema",
		},
		HCL:    `schema = "access"`,
		Update: true,
		ID:     "abc|information_schema",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
}

func TestSystemSchemaRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas?",
				Response: catalog.ListSystemSchemasResponse{
					Schemas: []catalog.SystemSchemaInfo{
						{
							Schema: "access",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
						{
							Schema: "billing",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
					},
				},
			},
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		ID:       "abc|access",
	}.ApplyAndExpectData(t, map[string]any{
		"schema": "access",
		"state":  string(catalog.SystemSchemaInfoStateEnableCompleted),
	})
}

func TestSystemSchemaRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		ID:       "abc|access",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|access", d.Id(), "Id should not be empty for error reads")
}

func TestSystemSchemaRead_NotEnabled(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas?",
				Response: catalog.ListSystemSchemasResponse{
					Schemas: []catalog.SystemSchemaInfo{
						{
							Schema: "access",
							State:  catalog.SystemSchemaInfoStateAvailable,
						},
						{
							Schema: "billing",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
					},
				},
			},
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		Removed:  true,
		ID:       "abc|access",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "", d.Id(), "Id should be empty if a schema is not enabled")
}

func TestSystemSchemaRead_NotExists(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas?",
				Response: catalog.ListSystemSchemasResponse{
					Schemas: []catalog.SystemSchemaInfo{
						{
							Schema: "billing",
							State:  catalog.SystemSchemaInfoStateEnableCompleted,
						},
					},
				},
			},
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		Removed:  true,
		ID:       "abc|access",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "", d.Id(), "Id should be empty if a schema does not exist")
}

func TestSystemSchemaDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access?",
				Status:   200,
			},
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/billing?",
				Status:   200,
			},
		},
		HCL:      `schema = "access"`,
		Resource: ResourceSystemSchema(),
		Delete:   true,
		ID:       "abc|access",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|access", d.Id())
}

func TestSystemSchemaDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.GetMetastoreSummaryResponse{
					MetastoreId: "abc",
				},
			},
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access?",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		HCL:      `schema = "access"`,
		Delete:   true,
		ID:       "abc|access",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|access", d.Id())
}
