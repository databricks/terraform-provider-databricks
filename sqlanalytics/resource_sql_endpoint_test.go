package sqlanalytics

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSQLEndpoints(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	ctx := context.Background()
	client := common.NewClientFromEnvironment()
	endpoitsAPI := NewSQLEndpointsAPI(ctx, client)
	se := SQLEndpoint{
		Name:            qa.RandomName("tf-"),
		ClusterSize:     "Small",
		AutoStopMinutes: 10,
		MaxNumClusters:  1,
		Tags: &Tags{
			CustomTags: []Tag{
				{"Country", "Netherlands"},
				{"City", "Amsterdam"},
			},
		},
	}
	err := endpoitsAPI.Create(&se, 20*time.Minute)
	require.NoError(t, err)
	defer func() {
		err = endpoitsAPI.Delete(se.ID)
		assert.NoError(t, err)
	}()

	se.Name = "renamed-" + se.Name
	err = endpoitsAPI.Edit(se)
	require.NoError(t, err)
}

func TestResourceSQLEndpointCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/sql/endpoints",
				ExpectedRequest: SQLEndpoint {
						Name: "foo",
						ClusterSize: "Small",
				},
				Response: SQLEndpoint {
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/endpoints/abc",
				ReuseRequest: true,
				Response: SQLEndpoint {
					Name: "foo",
					ClusterSize: "Small",
					ID: "abc",
					State: "RUNNING",
				},
			},
		},
		Resource: ResourceSQLEndpoint(),
		Create:   true,
		HCL: `
		name = "foo"
  		cluster_size = "Small"
		`,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
}
