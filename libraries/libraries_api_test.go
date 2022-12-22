package libraries

import (
	"context"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitForLibrariesInstalled(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=missing",
			ReuseRequest: true,
			Status:       404,
			Response:     common.NotFound("missing"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=error",
			ReuseRequest: true,
			Status:       500,
			Response: common.APIError{
				Message: "internal error",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=1005-abcd",
			ReuseRequest: true,
			Status:       400,
			Response: common.APIError{
				Message: "Cluster 1005-abcd does not exist",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=still-installing",
			ReuseRequest: true,
			Response: ClusterLibraryStatuses{
				ClusterID: "still-installing",
				LibraryStatuses: []LibraryStatus{
					{
						Status: "PENDING",
						Library: &Library{
							Jar: "a.jar",
						},
					},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=failed-wheel",
			ReuseRequest: true,
			Response: ClusterLibraryStatuses{
				ClusterID: "still-installing",
				LibraryStatuses: []LibraryStatus{
					{
						Status:   "FAILED",
						Messages: []string{"does not compute"},
						Library: &Library{
							Whl: "b.whl",
						},
					},
					{
						Status: "INSTALLED",
						Library: &Library{
							Jar: "a.jar",
						},
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/libraries/uninstall",
			ExpectedRequest: ClusterLibraryList{
				ClusterID: "failed-wheel",
				Libraries: []Library{
					{
						Whl: "b.whl",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		libs := NewLibrariesAPI(ctx, client)
		_, err := libs.WaitForLibrariesInstalled(Wait{
			"missing", 50 * time.Millisecond, true, false,
		})
		assert.EqualError(t, err, "missing")

		_, err = libs.WaitForLibrariesInstalled(Wait{
			"error", 50 * time.Millisecond, true, false,
		})
		assert.EqualError(t, err, "internal error")

		// cluster is not running
		_, err = libs.WaitForLibrariesInstalled(Wait{
			"still-installing", 50 * time.Millisecond, false, false,
		})
		assert.NoError(t, err)

		// cluster is running
		_, err = libs.WaitForLibrariesInstalled(Wait{
			"still-installing", 50 * time.Millisecond, true, false,
		})
		assert.EqualError(t, err, "0 libraries are ready, but there are still 1 pending")

		_, err = libs.WaitForLibrariesInstalled(Wait{
			"failed-wheel", 50 * time.Millisecond, true, false,
		})
		assert.EqualError(t, err, "whl:b.whl failed: does not compute")

		// uninstall b.whl and continue executing
		_, err = libs.WaitForLibrariesInstalled(Wait{
			"failed-wheel", 50 * time.Millisecond, true, true,
		})
		assert.NoError(t, err, "library should have been uninstalled and work proceeded")

		// Cluster not available or doesn't exist
		_, err = libs.WaitForLibrariesInstalled(Wait{
			"1005-abcd", 50 * time.Millisecond, false, false,
		})

		ae, _ := err.(common.APIError)
		assert.Equal(t, 404, ae.StatusCode)
		assert.Equal(t, "Cluster 1005-abcd does not exist", ae.Message)
	})
}

func TestClusterLibraryStatuses_UpdateLibraries(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/libraries/uninstall",
			ExpectedRequest: ClusterLibraryList{
				Libraries: []Library{
					{
						Jar: "remove.jar",
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/libraries/install",
			ExpectedRequest: ClusterLibraryList{
				Libraries: []Library{
					{
						Jar: "add.jar",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
			Response: ClusterLibraryStatuses{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		libsAPI := NewLibrariesAPI(ctx, client)
		err := libsAPI.UpdateLibraries("abc", ClusterLibraryList{
			Libraries: []Library{
				{
					Jar: "add.jar",
				},
			},
		}, ClusterLibraryList{
			Libraries: []Library{
				{
					Jar: "remove.jar",
				},
			},
		}, 1*time.Second)
		assert.NoError(t, err)
	})
}

func TestLibrariesDiff(t *testing.T) {
	install, uninstall := (&ClusterLibraryList{ // state
		ClusterID: "abc",
		Libraries: []Library{
			{ // kept
				Pypi: &PyPi{
					Package: "a",
				},
			},
			{ // added
				Maven: &Maven{
					Coordinates: "b",
				},
			},
			{ // added
				Cran: &Cran{
					Package: "c",
				},
			},
		},
	}).Diff(ClusterLibraryStatuses{ // remote
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{ // removed
				Library: &Library{
					Whl: "d",
				},
			},
			{ // removed
				Library: &Library{
					Jar: "d",
				},
			},
			{ // kept
				Library: &Library{
					Pypi: &PyPi{
						Package: "a",
					},
				},
			},
			{ // removed
				Library: &Library{
					Egg: "f",
				},
			},
		},
	})
	assert.Equal(t, "abc/cran:c,mvn:b", install.String())
	assert.Equal(t, "abc/egg:f,jar:d,whl:d", uninstall.String())
}

func TestClusterLibraryStatuses_ToLibraryList(t *testing.T) {
	cll := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				Library: &Library{
					Jar: "a",
				},
				Status: "INSTALLING",
			},
		},
	}.ToLibraryList()
	assert.Equal(t, "abc/jar:a", cll.String())
}

func TestClusterLibraryStatuses_NoNeedAllClusters(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				IsGlobal: true,
				Status:   "INSTALLING",
			},
		},
	}.IsRetryNeeded(false)
	require.NoError(t, err)
	assert.False(t, need)
}

func TestClusterLibraryStatuses_RetryingCodes(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				Status: "PENDING",
			},
			{
				Status: "RESOLVING",
			},
			{
				Status: "INSTALLING",
			},
			{
				Status: "INSTALLING",
			},
		},
	}.IsRetryNeeded(false)
	require.Error(t, err)
	assert.Equal(t, "0 libraries are ready, but there are still 4 pending", err.Error())
	assert.True(t, need)
}

func TestClusterLibraryStatuses_ReadyStatuses(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				Status: "INSTALLED",
			},
			{
				Status: "SKIPPED",
			},
			{
				Status: "UNINSTALL_ON_RESTART",
			},
		},
	}.IsRetryNeeded(false)
	require.NoError(t, err)
	assert.False(t, need)
}

func TestClusterLibraryStatuses_Errors(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				Status: "FAILED",
				Library: &Library{
					Whl: "a",
				},
				Messages: []string{"b"},
			},
			{
				Status: "FAILED",
				Library: &Library{
					Maven: &Maven{
						Coordinates: "a.b.c",
					},
				},
				Messages: []string{"b"},
			},
			{
				Status: "FAILED",
				Library: &Library{
					Cran: &Cran{
						Package: "a",
					},
				},
				Messages: []string{"b"},
			},
		},
	}.IsRetryNeeded(false)
	require.Error(t, err)
	assert.Equal(t, "whl:a failed: b\nmvn:a.b.c failed: b\ncran:a failed: b", err.Error())
	assert.False(t, need)
}

func TestNewLibraryFromInstanceState(t *testing.T) {
	tests := []struct {
		want string
		give any
	}{
		{"jar:a", map[string]any{"jar": "a"}},
		{"egg:b", map[string]any{"egg": "b"}},
		{"whl:c", map[string]any{"whl": "c"}},
		{"pypi:d", map[string]any{"pypi": []any{
			map[string]any{"package": "d"},
		}}},
		{"mvn:e", map[string]any{"maven": []any{
			map[string]any{"coordinates": "e"},
		}}},
		{"cran:f", map[string]any{"cran": []any{
			map[string]any{"package": "f"},
		}}},
		{"unknown", map[string]any{"bottle": "g"}},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := NewLibraryFromInstanceState(tt.give); got.String() != tt.want {
				t.Errorf("NewLibraryFromInstanceState() = %v, want %v", got, tt.want)
			}
		})
	}
}
