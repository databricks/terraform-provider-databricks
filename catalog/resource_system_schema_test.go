package catalog

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestSystemSchemaCreate(t *testing.T) {
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
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/billing",
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
		HCL:      `system_schema = ["access", "billing"]`,
		Create:   true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
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
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		HCL:      `system_schema = ["access", "billing"]`,
		Create:   true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
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
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/billing",
				Status:   200,
			},
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/information_schema",
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
			"system_schema.#": "1",
			"system_schema.0": "information_schema",
		},
		HCL:    `system_schema = ["access", "billing"]`,
		Update: true,
		ID:     "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, 2, d.Get("system_schema.#"))
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
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		InstanceState: map[string]string{
			"system_schema.#": "1",
			"system_schema.0": "information_schema",
		},
		HCL:    `system_schema = ["access", "billing"]`,
		Update: true,
		ID:     "abc",
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
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"system_schema": []any{"access", "billing"},
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
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
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
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access",
				Status:   200,
			},
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/billing",
				Status:   200,
			},
		},
		HCL:      `system_schema = ["access", "billing"]`,
		Resource: ResourceSystemSchema(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
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
				Resource: "/api/2.1/unity-catalog/metastores/abc/systemschemas/access",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSystemSchema(),
		HCL:      `system_schema = ["access", "billing"]`,
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
