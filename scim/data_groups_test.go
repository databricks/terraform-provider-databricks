package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceGroups_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceGroups(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestDataSourceGroups(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?count=100&startIndex=1`,
				Response: GroupList{
					TotalResults: 0,
					ItemsPerPage: 0,
					StartIndex:   1,
					Resources: []Group{
						{
							DisplayName: "ds",
							ID:          "eerste",
						},
						{
							DisplayName: "product",
							ID:          "abc",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?count=100&startIndex=3`,
				Response: GroupList{
					Resources: []Group{},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceGroups(),
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"display_names": []string{
			"ds",
			"product",
		},
	})
}

func TestDataSourceGroups_Filter(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?count=100&filter=display_name+sw+%22prod%22&startIndex=1`,
				Response: GroupList{
					TotalResults: 0,
					ItemsPerPage: 0,
					StartIndex:   1,
					Resources: []Group{
						{
							DisplayName: "product",
							ID:          "abc",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: `/api/2.0/preview/scim/v2/Groups?count=100&filter=display_name+sw+%22prod%22&startIndex=2`,
				Response: GroupList{
					Resources: []Group{},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceGroups(),
		HCL:         `filter="display_name sw \"prod\""`,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"display_names": []string{
			"product",
		},
	})
}
