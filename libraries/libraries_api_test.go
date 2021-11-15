package libraries

import (
	"context"
	"testing"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
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
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=still-installing",
			ReuseRequest: true,
			Response: ClusterLibraryStatuses{
				ClusterID: "still-installing",
				LibraryStatuses: []LibraryStatus{
					{
						Status: "INSTALLING",
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
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		libs := NewLibrariesAPI(ctx, client)
		_, err := libs.WaitForLibrariesInstalled("missing", 50*time.Millisecond)
		assert.EqualError(t, err, "missing")

		_, err = libs.WaitForLibrariesInstalled("error", 50*time.Millisecond)
		assert.EqualError(t, err, "internal error")

		_, err = libs.WaitForLibrariesInstalled("still-installing", 50*time.Millisecond)
		assert.EqualError(t, err, "0 libraries are ready, but there are still 1 pending")

		_, err = libs.WaitForLibrariesInstalled("failed-wheel", 50*time.Millisecond)
		assert.EqualError(t, err, "whl:b.whl failed: does not compute")

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
		})
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
				IsLibraryInstalledOnAllClusters: true,
				Status:                          "INSTALLING",
			},
		},
	}.IsRetryNeeded()
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
	}.IsRetryNeeded()
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
	}.IsRetryNeeded()
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
	}.IsRetryNeeded()
	require.Error(t, err)
	assert.Equal(t, "whl:a failed: b\nmvn:a.b.c failed: b\ncran:a failed: b", err.Error())
	assert.False(t, need)
}

func TestNewLibraryFromInstanceState(t *testing.T) {
	tests := []struct {
		want string
		give interface{}
	}{
		{"jar:a", map[string]interface{}{"jar": "a"}},
		{"egg:b", map[string]interface{}{"egg": "b"}},
		{"whl:c", map[string]interface{}{"whl": "c"}},
		{"pypi:d", map[string]interface{}{"pypi": []interface{}{
			map[string]interface{}{"package": "d"},
		}}},
		{"mvn:e", map[string]interface{}{"maven": []interface{}{
			map[string]interface{}{"coordinates": "e"},
		}}},
		{"cran:f", map[string]interface{}{"cran": []interface{}{
			map[string]interface{}{"package": "f"},
		}}},
		{"unknown", map[string]interface{}{"bottle": "g"}},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := NewLibraryFromInstanceState(tt.give); got.String() != tt.want {
				t.Errorf("NewLibraryFromInstanceState() = %v, want %v", got, tt.want)
			}
		})
	}
}
