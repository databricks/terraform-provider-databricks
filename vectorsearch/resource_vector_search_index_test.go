package vectorsearch

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/databricks/databricks-sdk-go/service/vectorsearch"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVectorSearchIndexCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceVectorSearchIndex())
}

var indexResponse = &vectorsearch.VectorIndex{
	Name:         "abc",
	EndpointName: "test",
	PrimaryKey:   "id",
	IndexType:    "DELTA_SYNC",
	DeltaSyncIndexSpec: &vectorsearch.DeltaSyncVectorIndexSpecResponse{
		SourceTable:  "main.default.test",
		PipelineType: "TRIGGERED",
		EmbeddingSourceColumns: []vectorsearch.EmbeddingSourceColumn{
			{
				Name:                       "text",
				EmbeddingModelEndpointName: "e5_small_v2",
			},
		},
	},
	Status: &vectorsearch.VectorIndexStatus{
		Ready: true,
	},
}

var indexHcl = `
name          = "abc"
endpoint_name = "test"
primary_key   = "id"
index_type    = "DELTA_SYNC"
delta_sync_index_spec {
   source_table  = "main.default.test"
   pipeline_type = "TRIGGERED"
   embedding_source_columns {
   name                          = "text"
   embedding_model_endpoint_name = "e5_small_v2"
  }
}
`

func TestVectorSearchIndexCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockVectorSearchIndexesAPI().EXPECT()
			e.CreateIndex(mock.Anything, vectorsearch.CreateVectorIndexRequest{
				Name:         "abc",
				EndpointName: "test",
				PrimaryKey:   "id",
				IndexType:    "DELTA_SYNC",
				DeltaSyncIndexSpec: &vectorsearch.DeltaSyncVectorIndexSpecRequest{
					SourceTable:  "main.default.test",
					PipelineType: "TRIGGERED",
					EmbeddingSourceColumns: []vectorsearch.EmbeddingSourceColumn{
						{
							Name:                       "text",
							EmbeddingModelEndpointName: "e5_small_v2",
						},
					},
				},
			}).Return(&vectorsearch.CreateVectorIndexResponse{}, nil)
			e.GetIndexByIndexName(mock.Anything, "abc").Return(indexResponse, nil)
		},
		Resource: ResourceVectorSearchIndex(),
		HCL:      indexHcl,
		Create:   true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestVectorSearchIndexCreateNotReadyButIndexedRows(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockVectorSearchIndexesAPI().EXPECT()
			e.CreateIndex(mock.Anything, vectorsearch.CreateVectorIndexRequest{
				Name:         "abc",
				EndpointName: "test",
				PrimaryKey:   "id",
				IndexType:    "DELTA_SYNC",
				DeltaSyncIndexSpec: &vectorsearch.DeltaSyncVectorIndexSpecRequest{
					SourceTable:  "main.default.test",
					PipelineType: "TRIGGERED",
					EmbeddingSourceColumns: []vectorsearch.EmbeddingSourceColumn{
						{
							Name:                       "text",
							EmbeddingModelEndpointName: "e5_small_v2",
						},
					},
				},
			}).Return(&vectorsearch.CreateVectorIndexResponse{}, nil)
			e.GetIndexByIndexName(mock.Anything, "abc").Return(&vectorsearch.VectorIndex{
				Name:         "abc",
				EndpointName: "test",
				PrimaryKey:   "id",
				IndexType:    "DELTA_SYNC",
				DeltaSyncIndexSpec: &vectorsearch.DeltaSyncVectorIndexSpecResponse{
					SourceTable:  "main.default.test",
					PipelineType: "TRIGGERED",
					EmbeddingSourceColumns: []vectorsearch.EmbeddingSourceColumn{
						{
							Name:                       "text",
							EmbeddingModelEndpointName: "e5_small_v2",
						},
					},
				},
				Status: &vectorsearch.VectorIndexStatus{
					IndexedRowCount: 10,
				},
			}, nil)
		},
		Resource: ResourceVectorSearchIndex(),
		HCL:      indexHcl,
		Create:   true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestVectorSearchIndexRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockVectorSearchIndexesAPI().EXPECT()
			e.GetIndexByIndexName(mock.Anything, "abc").Return(indexResponse, nil)
		},
		Resource: ResourceVectorSearchIndex(),
		ID:       "abc",
		HCL:      indexHcl,
		Read:     true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestVectorSearchIndexDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockVectorSearchIndexesAPI().EXPECT()
			e.DeleteIndexByIndexName(mock.Anything, "abc").Return(nil)
			e.GetIndexByIndexName(mock.Anything, "abc").Return(nil, apierr.ErrResourceDoesNotExist)
		},
		Resource: ResourceVectorSearchIndex(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}
