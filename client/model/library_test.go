package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClusterLibraryStatuses_Apply(t *testing.T) {
	tests := []struct {
		Name     string
		Statuses []LibraryStatus
		Callback func(key string, value interface{}) error
		Amount   int
		Retry    bool
		Error    string
	}{
		{
			Name: "egg",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Egg: "dbfs://path/to/egg",
					},
					IsLibraryInstalledOnAllClusters: false,
					Messages: []string{
						"first message",
						"second message",
					},
					Status: "RESOLVING",
				},
			},
			Callback: func(key string, value interface{}) error {
				assert.Equal(t, "library_egg", key)
				assert.Equal(t, []map[string]interface{}{
					{
						"path":   "dbfs://path/to/egg",
						"status": "RESOLVING",
					},
				}, value)
				return nil
			},
			Amount: 1,
			Retry:  true,
		},
		{
			Name: "wheel",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Whl: "dbfs://wheel/path",
					},
					Status: "INSTALLED",
				},
			},
			Callback: func(key string, value interface{}) error {
				assert.Equal(t, "library_whl", key)
				assert.Equal(t, []map[string]interface{}{
					{
						"path":   "dbfs://wheel/path",
						"status": "INSTALLED",
					},
				}, value)
				return nil
			},
			Amount: 1,
			Retry:  false,
		},
		{
			Name: "jar",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Jar: "s3a://path/to/file.jar",
					},
					Status: "INSTALLED",
				},
			},
			Callback: func(key string, value interface{}) error {
				assert.Equal(t, "library_jar", key)
				assert.Equal(t, []map[string]interface{}{
					{
						"path":   "s3a://path/to/file.jar",
						"status": "INSTALLED",
					},
				}, value)
				return nil
			},
			Amount: 1,
			Retry:  false,
		},
		{
			Name: "maven",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Maven: &Maven{
							Coordinates: "foo:bar:0.0.1",
							Exclusions: []string{
								"foo:baz:0.0.1",
							},
							Repo: "http://something.com",
						},
					},
					Status: "INSTALLED",
				},
			},
			Callback: func(key string, value interface{}) error {
				assert.Equal(t, "library_maven", key)
				assert.Equal(t, []map[string]interface{}{{
					"coordinates": "foo:bar:0.0.1",
					"exclusions":  []string{"foo:baz:0.0.1"},
					"repo":        "http://something.com",
					"status":      "INSTALLED"}}, value)
				return nil
			},
			Amount: 1,
			Retry:  false,
		},
		{
			Name: "pypi",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Pypi: &PyPi{
							Package: "airflow==0.0.1",
							Repo:    "https://foo.com",
						},
					},
					Status: "INSTALLED",
				},
			},
			Callback: func(key string, value interface{}) error {
				assert.Equal(t, "library_pypi", key)
				assert.Equal(t, []map[string]interface{}{{
					"package": "airflow==0.0.1",
					"repo":    "https://foo.com",
					"status":  "INSTALLED"}}, value)
				return nil
			},
			Amount: 1,
			Retry:  false,
		},
		{
			Name: "cran",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Cran: &Cran{
							Package: "dplyr",
							Repo:    "some",
						},
					},
					Status: "INSTALLED",
				},
			},
			Callback: func(key string, value interface{}) error {
				assert.Equal(t, "library_cran", key)
				assert.Equal(t, []map[string]interface{}{{
					"package": "dplyr",
					"repo":    "some",
					"status":  "INSTALLED"}}, value)
				return nil
			},
			Amount: 1,
			Retry:  false,
		},
		{
			Name: "jar failed",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Jar: "s3a://path/to/file.jar",
					},
					Status: "FAILED",
					Messages: []string{
						"First message here",
						"Second message here",
					},
				},
				{
					Library: &Library{
						Maven: &Maven{
							Coordinates: "foo:bar:0.0.1",
							Exclusions: []string{
								"foo:baz:0.0.1",
							},
							Repo: "http://something.com",
						},
					},
					Status:   "FAILED",
					Messages: []string{"Does not compute"},
				},
			},
			Amount: 0,
			Error: "library_jar[s3a://path/to/file.jar] failed: First message here, Second message here\n" +
				"library_maven[foo:bar:0.0.1http://something.comfoo:baz:0.0.1] failed: Does not compute",
			Retry: false,
		},
		{
			Name: "jars installing",
			Statuses: []LibraryStatus{
				{
					Library: &Library{
						Jar: "s3a://path/to/file.jar",
					},
					Status: "PENDING",
				},
				{
					Library: &Library{
						Maven: &Maven{
							Coordinates: "foo:bar:0.0.1",
							Exclusions: []string{
								"foo:baz:0.0.1",
							},
							Repo: "http://something.com",
						},
					},
					Status: "INSTALLED",
				},
				{
					Library: &Library{
						Maven: &Maven{
							Coordinates: "org:bar:0.0.1",
						},
					},
					Status: "SKIPPED",
				},
			},
			Amount: 3,
			Error:  "2 libraries are ready, but there are still 1 pending",
			Retry:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			cls := ClusterLibraryStatuses{
				LibraryStatuses: tt.Statuses,
			}
			libs := 0
			cb := func(k string, v interface{}) error {
				if lm, ok := v.([]map[string]interface{}); ok {
					libs += len(lm)
				}
				if tt.Callback != nil {
					return tt.Callback(k, v)
				}
				return nil
			}
			if err := cls.Apply(cb); err != nil {
				assert.EqualError(t, err, tt.Error, tt.Name)
			}
			assert.Equal(t, tt.Amount, libs, "Actual lib amount doesn't match")
			retry, err := cls.IsRetryNeeded()
			assert.Equal(t, tt.Retry, retry, "Is there a need for retry?")
			if len(tt.Error) > 0 {
				assert.EqualError(t, err, tt.Error, tt.Name)
			}
		})
	}
}
