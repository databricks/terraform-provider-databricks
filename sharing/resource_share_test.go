package sharing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDiffShareInfo(t *testing.T) {
	empty := ShareInfo{
		Name:    "b",
		Objects: []SharedDataObject{},
	}
	firstShare := ShareInfo{
		Name: "b",
		Objects: []SharedDataObject{
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
	}
	secondShare := ShareInfo{
		Name: "b",
		Objects: []SharedDataObject{
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
	}
	thirdShare := ShareInfo{
		Name: "b",
		Objects: []SharedDataObject{
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
	}
	fourthShare := ShareInfo{
		Name: "d",
		Objects: []SharedDataObject{
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
	}
	diffAdd := []ShareDataChange{
		{
			Action: ShareAdd,
			DataObject: SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: ShareAdd,
			DataObject: SharedDataObject{
				Name:           "main.a",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
	}
	diffRemove := []ShareDataChange{
		{
			Action: ShareRemove,
			DataObject: SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: ShareRemove,
			DataObject: SharedDataObject{
				Name:           "main.a",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
	}
	diff12 := []ShareDataChange{
		{
			Action: ShareRemove,
			DataObject: SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: ShareAdd,
			DataObject: SharedDataObject{
				Name:           "main.c",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
	}
	diff13 := []ShareDataChange{
		{
			Action: ShareRemove,
			DataObject: SharedDataObject{
				Name:           "main.a",
				DataObjectType: "TABLE",
				Comment:        "c",
			},
		},
		{
			Action: ShareAdd,
			DataObject: SharedDataObject{
				Name:           "main.c",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
		{
			Action: ShareUpdate,
			DataObject: SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
	}
	diff14 := []ShareDataChange{
		{
			Action: ShareUpdate,
			DataObject: SharedDataObject{
				Name:           "main.b",
				DataObjectType: "TABLE",
				Comment:        "d",
			},
		},
	}
	assert.Equal(t, firstShare.Diff(firstShare), []ShareDataChange{}, "Should not have difference")
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
					Name: "a",
				},
				Response: ShareInfo{
					Name: "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/a",
				ExpectedRequest: ShareUpdates{
					Owner: "admin",
					Updates: []ShareDataChange{
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.a",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.b",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
					},
				},
				Response: ShareInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/a?include_shared_data=true",
				Response: ShareInfo{
					Name:  "a",
					Owner: "admin",
					Objects: []SharedDataObject{
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
		Resource: ResourceShare(),
		Create:   true,
		HCL: `
			name  = "a"
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
					Name: "abc",
					Objects: []SharedDataObject{
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
				ExpectedRequest: ShareUpdates{
					Updates: []ShareDataChange{
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "a",
							},
						},
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "b",
							},
						},
						{
							Action: "REMOVE",
							DataObject: SharedDataObject{
								Comment:        "d",
								DataObjectType: "TABLE",
								Name:           "d",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					Name:  "abc",
					Owner: "admin",
					Objects: []SharedDataObject{
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
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name": "abc",
		},
		HCL: `
			name  = "abc"
			owner = "admin"
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
					Name: "abc",
					Objects: []SharedDataObject{
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
				ExpectedRequest: ShareUpdates{
					Updates: []ShareDataChange{
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "a",
							},
						},
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Comment:        "c",
								DataObjectType: "TABLE",
								Name:           "b",
							},
						},
						{
							Action: "REMOVE",
							DataObject: SharedDataObject{
								Comment:        "d",
								DataObjectType: "TABLE",
								Name:           "d",
							},
						},
					},
				},
				Response: apierr.APIErrorBody{
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
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					Name:  "abc",
					Owner: "admin",
					Objects: []SharedDataObject{
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
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":  "abc",
			"owner": "admin",
		},
		HCL: `
			name  = "abc"
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
					Name: "abc",
					Objects: []SharedDataObject{
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
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/abc?include_shared_data=true",
				Response: ShareInfo{
					Name: "abc",
					Objects: []SharedDataObject{
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
					Name: "a",
				},
				Response: apierr.APIErrorBody{
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
					Name: "a",
				},
				Response: ShareInfo{
					Name: "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/a",
				ExpectedRequest: ShareUpdates{
					Updates: []ShareDataChange{
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.a",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.b",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
					},
				},
				Response: apierr.APIErrorBody{
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
					Name: "abc",
					Objects: []SharedDataObject{
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
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/shares/abc",
				ExpectedRequest: ShareUpdates{
					Updates: []ShareDataChange{
						{
							Action: "ADD",
							DataObject: SharedDataObject{
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
					Name: "abc",
					Objects: []SharedDataObject{
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
