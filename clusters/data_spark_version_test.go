package clusters

import (
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSparkVersionLatest(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest: true,
				Scala:  "2.12",
			}).Return("7.4.x-scala2.12", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State:       map[string]any{},
		ID:          ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "7.4.x-scala2.12",
	})
}

func TestSparkVersionLTS(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest:          true,
				Scala:           "2.12",
				LongTermSupport: true,
			}).Return("7.3.x-scala2.12", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"long_term_support": true,
		},
		ID: ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "7.3.x-scala2.12",
	})
}

func TestSparkVersionESR(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest:          true,
				Scala:           "2.11",
				LongTermSupport: true,
				ML:              true,
			}).Return("5.5.x-cpu-esr-ml-scala2.11", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"long_term_support": true,
			"scala":             "2.11",
			"ml":                true,
		},
		ID: ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "5.5.x-cpu-esr-ml-scala2.11",
	})
}

func TestSparkVersionGpuMl(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest: true,
				Scala:  "2.12",
				GPU:    true,
				ML:     true,
			}).Return("7.1.x-gpu-ml-scala2.12", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"gpu": true,
			"ml":  true,
		},
		ID: ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "7.1.x-gpu-ml-scala2.12",
	})
}

func TestSparkVersionGenomics(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest:   true,
				Scala:    "2.12",
				Genomics: true,
			}).Return("7.3.x-hls-scala2.12", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"genomics": true,
		},
		ID: ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "7.3.x-hls-scala2.12",
	})
}

func TestSparkVersion300(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest:       true,
				Scala:        "2.12",
				SparkVersion: "3.0.0",
			}).Return("7.1.x-scala2.12", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"spark_version": "3.0.0",
		},
		ID: ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "7.1.x-scala2.12",
	})
}

func TestSparkVersionBeta(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest: true,
				Scala:  "2.12",
				Beta:   true,
			}).Return("7.5.x-scala2.12", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"beta": true,
		},
		ID: ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "7.5.x-scala2.12",
	})
}

func TestSparkVersionPhoton(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest: true,
				Scala:  "2.12",
				Photon: true,
			}).Return("8.3.x-photon-scala2.12", nil)
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"photon": true,
		},
		ID: ".",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "8.3.x-photon-scala2.12",
	})
}

func TestSparkVersionErrorNoResults(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest:          true,
				Scala:           "2.12",
				Beta:            true,
				LongTermSupport: true,
			}).Return("", fmt.Errorf("spark versions query returned no results. Please change your search criteria and try again"))
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"beta":              true,
			"long_term_support": true,
		},
		ID: ".",
	}.ExpectError(t, "spark versions query returned no results. Please change your search criteria and try again")
}

func TestSparkVersionErrorMultipleResults(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.SelectSparkVersion(mock.Anything, compute.SparkVersionRequest{
				Latest: false,
				Scala:  "2.12",
			}).Return("", fmt.Errorf("spark versions query returned multiple results. Please change your search criteria and try again"))
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"latest": false,
		},
		ID: ".",
	}.ExpectError(t, "spark versions query returned multiple results. Please change your search criteria and try again")
}

func TestSparkVersionErrorBadAnswer(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/spark-versions",
				Response: "{garbage....",
			},
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]any{
			"latest": false,
		},
		ID: ".",
	}.Apply(t)
	assert.Error(t, err)
	require.Equal(t, true, strings.Contains(err.Error(), "invalid character 'g' looking"))
}
