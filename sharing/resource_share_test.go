package sharing

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/databricks-sdk-go/service/sharing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDiffShareInfo(t *testing.T) {
	empty := ShareInfo{
		ShareInfo: sharing.ShareInfo{
			Name:    "b",
			Objects: []sharing.SharedDataObject{},
		},
	}
	firstShare := ShareInfo{
		ShareInfo: sharing.ShareInfo{
			Name: "b",
			Objects: []sharing.SharedDataObject{
				{
					Name:           "main.b",
					DataObjectType: "TABLE",
					Comment:        "c",
				},
				{
					Name:           "main.a",
					DataObjectType: "TABLE",
					Comment:        "c",
				},
			},
		},
	}
	secondShare := ShareInfo{
		ShareInfo: sharing.ShareInfo{
			Name: "b",
			Objects: []sharing.SharedDataObject{
				{
					Name:           "main.c",
					DataObjectType: "TABLE",
					Comment:        "d",
				},
				{
					Name:           "main.a",
					DataObjectType: "TABLE",
					Comment:        "c",
				},
			},
		},
	}
	thirdShare := ShareInfo{
		ShareInfo: sharing.ShareInfo{
			Name: "b",
			Objects: []sharing.SharedDataObject{
				{
					Name:           "main.c",
					DataObjectType: "TABLE",
					Comment:        "d",
				},
				{
					Name:           "main.b",
					DataObjectType: "TABLE",
					Comment:        "d",
				},
			},
		},
	}
	fourthShare := ShareInfo{
		ShareInfo: sharing.ShareInfo{
			Name: "d",
			Objects: []sharing.SharedDataObject{
				{
					Name:           "main.b",
					DataObjectType: "TABLE",
					Comment:        "d",
				},
				{
					Name:           "main.a",
					DataObjectType: "TABLE",
					Comment:        "c",
				},
			},
		},
	}
	diffAdd := []sharing.SharedDataObjectUpdate{
		{
			Action: sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.a",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
	}
	diffRemove := []sharing.SharedDataObjectUpdate{
		{
			Action: sharing.SharedDataObjectUpdateActionRemove,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: sharing.SharedDataObjectUpdateActionRemove,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.a",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
	}
	diff12 := []sharing.SharedDataObjectUpdate{
		{
			Action: sharing.SharedDataObjectUpdateActionRemove,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.c",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
	}
	diff13 := []sharing.SharedDataObjectUpdate{
		{
			Action: sharing.SharedDataObjectUpdateActionRemove,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.a",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.c",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
		{
			Action: sharing.SharedDataObjectUpdateActionUpdate,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
	}
	diff14 := []sharing.SharedDataObjectUpdate{
		{
			Action: sharing.SharedDataObjectUpdateActionUpdate,
			DataObject: &sharing.SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
	}
	assert.Equal(t, firstShare.Diff(firstShare), []sharing.SharedDataObjectUpdate{}, "Should not have difference")
	assert.Equal(t, empty.Diff(firstShare), diffAdd, "Should have 2 ADDs")
	assert.Equal(t, firstShare.Diff(empty), diffRemove, "Should have 2 REMOVEs")
	assert.Equal(t, firstShare.Diff(secondShare), diff12, "Should have 1 ADD and 1 REMOVE")
	assert.Equal(t, firstShare.Diff(thirdShare), diff13, "Should have 1 ADD, 1 REMOVE and 1 UPDATE")
	assert.Equal(t, firstShare.Diff(fourthShare), diff14, "Should have 1 UPDATE")
}

func TestShareCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceShare())
}

func TestCreateShare(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/shares",
				ExpectedRequest: ShareInfo{
					ShareInfo: sharing.ShareInfo{
						Name:    "a",
						Comment: "b",
					},
				},
				Response: ShareInfo{
					ShareInfo: sharing.ShareInfo{
						Name:    "a",
						Comment: "b",
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/a",
				ExpectedRequest: sharing.UpdateShare{
					Owner:   "admin",
					Comment: "b",
					Updates: []sharing.SharedDataObjectUpdate{
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Name:           "main.a",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Name:           "main.b",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
					},
				},
				Response: ShareInfo{
					ShareInfo: sharing.ShareInfo{
						Name:    "a",
						Comment: "b",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/a?include_shared_data=true",
				Response: ShareInfo{
					ShareInfo: sharing.ShareInfo{
						Name:    "a",
						Comment: "b",
						Owner:   "admin",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "main.a",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
							{
								Name:           "main.b",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
					},
				},
			},
		},
		Resource: ResourceShare(),
		Create:   true,
		HCL: `
			name  = "a"
			comment = "b"
			owner = "admin"
			object {
				name = "main.a"
				comment = "c"
				data_object_type = "TABLE"
			}
			object {
				name = "main.b"
				comment = "c"
				data_object_type = "TABLE"
			}
		`,
	}.ApplyNoError(t)
}

func TestUpdateShare(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					ShareInfo: sharing.ShareInfo{
						Name: "abc",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "d",
								DataObjectType: "TABLE",
								Comment:        "d",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: sharing.UpdateShare{
					Owner: "admin",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: sharing.UpdateShare{
					Comment: "cba",
					Updates: []sharing.SharedDataObjectUpdate{
						{
							Action: "REMOVE",
							DataObject: &sharing.SharedDataObject{
								Comment:        "d",
								DataObjectType: "TABLE",
								Name:           "d",
							},
						},
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "a",
							},
						},
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "b",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					sharing.ShareInfo{
						Name:    "abc",
						Owner:   "admin",
						Comment: "cba",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "a",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
							{
								Name:           "b",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
						},
					},
				},
			},
		},
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name": "abc",
		},
		HCL: `
			name  = "abc"
			owner = "admin"
			comment = "cba"
			object {
				name = "a"
				comment = "c"
				data_object_type = "TABLE"
			}
			object {
				name = "b"
				comment = "c"
				data_object_type = "TABLE"
			}
		`,
		Resource: ResourceShare(),
	}.ApplyNoError(t)
}

func TestUpdateShareRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					sharing.ShareInfo{
						Name: "abc",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "d",
								DataObjectType: "TABLE",
								Comment:        "d",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: sharing.UpdateShare{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: sharing.UpdateShare{
					Comment: "updatedComment",
					Updates: []sharing.SharedDataObjectUpdate{
						{
							Action: "REMOVE",
							DataObject: &sharing.SharedDataObject{
								Comment:        "d",
								DataObjectType: "TABLE",
								Name:           "d",
							},
						},
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "a",
							},
						},
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "b",
							},
						},
					},
				},
				Response: apierr.APIError{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: sharing.UpdateShare{
					Owner: "admin",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: sharing.UpdateShare{
					Comment: "cba",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					sharing.ShareInfo{
						Name:    "abc",
						Comment: "cba",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "a",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
							{
								Name:           "b",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
						},
					},
				},
			},
		},
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":    "abc",
			"owner":   "admin",
			"comment": "cba",
		},
		HCL: `
			name  = "abc"
			comment = "updatedComment"
			owner = "updatedOwner"
			object {
				name = "a"
				comment = "c"
				data_object_type = "TABLE"
			}
			object {
				name = "b"
				comment = "c"
				data_object_type = "TABLE"
			}
		`,
		Resource: ResourceShare(),
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateShare_NoChanges(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					sharing.ShareInfo{
						Name: "abc",
						Objects: []sharing.SharedDataObject{
							{
								Name:            "d",
								DataObjectType:  "TABLE",
								Comment:         "c",
								SharedAs:        "",
								AddedAt:         0,
								AddedBy:         "",
								ForceSendFields: []string{"Name", "Comment", "DataObjectType"},
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					sharing.ShareInfo{
						Name: "abc",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "d",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
						},
					},
				},
			},
		},
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name": "abc",
		},
		HCL: `
			name = "abc"
				object {
					name = "d"
					comment = "c"
					data_object_type = "TABLE"
				}
		`,
		Resource: ResourceShare(),
	}.ApplyNoError(t)
}

func TestCreateShare_ThrowError(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/shares",
				ExpectedRequest: ShareInfo{
					sharing.ShareInfo{
						Name: "a",
					},
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceShare(),
		Create:   true,
		HCL: `
			name = "a"
			object {
				name = "main.a"
				comment = "c"
				data_object_type = "TABLE"
			}
			object {
				name = "main.b"
				comment = "c"
				data_object_type = "TABLE"
			}
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestCreateShareButPatchFails(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/shares",
				ExpectedRequest: ShareInfo{
					sharing.ShareInfo{
						Name: "a",
					},
				},
				Response: ShareInfo{
					sharing.ShareInfo{
						Name: "a",
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/a",
				ExpectedRequest: sharing.UpdateShare{
					Updates: []sharing.SharedDataObjectUpdate{
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Name:           "main.a",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Name:           "main.b",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
					},
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/shares/a?",
			},
		},
		Resource: ResourceShare(),
		Create:   true,
		HCL: `
			name = "a"
			object {
				name = "main.a"
				comment = "c"
				data_object_type = "TABLE"
			}
			object {
				name = "main.b"
				comment = "c"
				data_object_type = "TABLE"
			}
		`,
	}.Apply(t)

	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestUpdateShareComplexDiff(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					sharing.ShareInfo{
						Name: "abc",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "a",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "b",
								AddedAt:        0,
								AddedBy:        "",
							},
						},
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: sharing.UpdateShare{
					Updates: []sharing.SharedDataObjectUpdate{
						{
							Action: "ADD",
							DataObject: &sharing.SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "b",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					sharing.ShareInfo{
						Name: "abc",
						Objects: []sharing.SharedDataObject{
							{
								Name:           "a",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
							{
								Name:           "b",
								DataObjectType: "TABLE",
								Comment:        "c",
								SharedAs:       "",
								AddedAt:        0,
								AddedBy:        "",
							},
						},
					},
				},
			},
		},
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name": "abc",
		},
		HCL: `
			name = "abc"
			object {
				name = "a"
				comment = "c"
				data_object_type = "TABLE"
			}
			object {
				name = "b"
				comment = "c"
				data_object_type = "TABLE"
			}
		`,
		Resource: ResourceShare(),
	}.ApplyNoError(t)
}
