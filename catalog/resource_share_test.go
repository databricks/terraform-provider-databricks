package catalog

import (
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestShareCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceShare())
}

func TestCreateShare(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/shares",
				ExpectedRequest: ShareInfo{
					Name: "a",
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
				},
				Response: RecipientInfo{
					Name: "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/shares/a",
				ExpectedRequest: ShareUpdates{
					Updates: []ShareDataChange{
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.b",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.a",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
					},
				},
				Response: RecipientInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/shares/a?include_shared_data=true",
				Response: ShareInfo{
					Name: "a",
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
				},
			},
		},
		Resource: ResourceShare(),
		Create:   true,
		HCL: `
		name = "a"
	    objects {
		  name = "main.a"
		  comment = "c"
		  data_object_type = "TABLE"
	    }
	    objects {
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
				Resource: "/api/2.0/unity-catalog/shares/abc?include_shared_data=true",
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
				Resource: "/api/2.0/unity-catalog/shares/abc",
				Response: ShareUpdates{
					Updates: []ShareDataChange{
						{
							Action: "REMOVE",
							DataObject: SharedDataObject{
								Comment:        "d",
								DataObjectType: "TABLE",
								Name:           "d",
							},
						},
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
								Comment:        "d",
								DataObjectType: "TABLE",
								Name:           "b",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/shares/abc?include_shared_data=true",
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
		HCL: `name = "abc"
		  objects {
			name = "a"
			comment = "c"
			data_object_type = "TABLE"
		  }
		  objects {
			name = "b"
			comment = "c"
			data_object_type = "TABLE"
		  }`,
		Resource: ResourceShare(),
	}.ApplyNoError(t)
}

func TestUpdateShare_NoChanges(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/shares/abc?include_shared_data=true",
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
				Resource: "/api/2.0/unity-catalog/shares/abc?include_shared_data=true",
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
		HCL: `name = "abc"
		  objects {
			name = "d"
			comment = "c"
			data_object_type = "TABLE"
		  }`,
		Resource: ResourceShare(),
	}.ApplyNoError(t)
}

func TestCreateShare_ThrowError(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/shares",
				ExpectedRequest: ShareInfo{
					Name: "a",
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
				},
				Response: RecipientInfo{
					Name: "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/shares/a",
				ExpectedRequest: ShareUpdates{
					Updates: []ShareDataChange{
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.b",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
						{
							Action: "ADD",
							DataObject: SharedDataObject{
								Name:           "main.a",
								DataObjectType: "TABLE",
								Comment:        "c",
							},
						},
					},
				},
				Response: common.APIErrorBody{
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
	    objects {
		  name = "main.a"
		  comment = "c"
		  data_object_type = "TABLE"
	    }
	    objects {
		  name = "main.b"
		  comment = "c"
		  data_object_type = "TABLE"
	    }
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}
