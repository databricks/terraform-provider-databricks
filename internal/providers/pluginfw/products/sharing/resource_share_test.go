package sharing

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/stretchr/testify/assert"
)

// TestDiff tests the diff function that compares two ShareInfo states
func TestDiff(t *testing.T) {
	empty := sharing.ShareInfo{
		Name:    "test-share",
		Objects: []sharing.SharedDataObject{},
	}

	firstShare := sharing.ShareInfo{
		Name: "test-share",
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
	}

	secondShare := sharing.ShareInfo{
		Name: "test-share",
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
	}

	thirdShare := sharing.ShareInfo{
		Name: "test-share",
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
	}

	fourthShare := sharing.ShareInfo{
		Name: "test-share",
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
	}

	// Test: No difference when comparing same share
	assert.Equal(t, []sharing.SharedDataObjectUpdate{}, diff(firstShare, firstShare), "Should not have difference")

	// Test: Adding objects to empty share
	diffAdd := diff(empty, firstShare)
	assert.Len(t, diffAdd, 2, "Should have 2 ADDs")
	for _, update := range diffAdd {
		assert.Equal(t, sharing.SharedDataObjectUpdateActionAdd, update.Action)
	}

	// Test: Removing all objects
	diffRemove := diff(firstShare, empty)
	assert.Len(t, diffRemove, 2, "Should have 2 REMOVEs")
	for _, update := range diffRemove {
		assert.Equal(t, sharing.SharedDataObjectUpdateActionRemove, update.Action)
	}

	// Test: One ADD and one REMOVE
	diff12 := diff(firstShare, secondShare)
	assert.Len(t, diff12, 2, "Should have 2 changes")
	var hasAdd, hasRemove bool
	for _, update := range diff12 {
		if update.Action == sharing.SharedDataObjectUpdateActionAdd {
			hasAdd = true
			assert.Equal(t, "main.c", update.DataObject.Name)
		}
		if update.Action == sharing.SharedDataObjectUpdateActionRemove {
			hasRemove = true
			assert.Equal(t, "main.b", update.DataObject.Name)
		}
	}
	assert.True(t, hasAdd, "Should have ADD action")
	assert.True(t, hasRemove, "Should have REMOVE action")

	// Test: One ADD, one REMOVE, one UPDATE
	diff13 := diff(firstShare, thirdShare)
	assert.Len(t, diff13, 3, "Should have 3 changes")
	var hasUpdate bool
	hasAdd, hasRemove = false, false
	for _, update := range diff13 {
		switch update.Action {
		case sharing.SharedDataObjectUpdateActionAdd:
			hasAdd = true
			assert.Equal(t, "main.c", update.DataObject.Name)
		case sharing.SharedDataObjectUpdateActionRemove:
			hasRemove = true
			assert.Equal(t, "main.a", update.DataObject.Name)
		case sharing.SharedDataObjectUpdateActionUpdate:
			hasUpdate = true
			assert.Equal(t, "main.b", update.DataObject.Name)
			assert.Equal(t, "d", update.DataObject.Comment)
		}
	}
	assert.True(t, hasAdd, "Should have ADD action")
	assert.True(t, hasRemove, "Should have REMOVE action")
	assert.True(t, hasUpdate, "Should have UPDATE action")

	// Test: Only UPDATE
	diff14 := diff(firstShare, fourthShare)
	assert.Len(t, diff14, 1, "Should have 1 UPDATE")
	assert.Equal(t, sharing.SharedDataObjectUpdateActionUpdate, diff14[0].Action)
	assert.Equal(t, "main.b", diff14[0].DataObject.Name)
	assert.Equal(t, "d", diff14[0].DataObject.Comment)
}

// TestEqual tests the equal function that compares SharedDataObjects
func TestEqual(t *testing.T) {
	obj1 := sharing.SharedDataObject{
		Name:           "main.table",
		DataObjectType: "TABLE",
		Comment:        "test comment",
		SharedAs:       "alias",
		AddedAt:        123456,
		AddedBy:        "user@example.com",
		Status:         "ACTIVE",
	}

	obj2 := sharing.SharedDataObject{
		Name:           "main.table",
		DataObjectType: "TABLE",
		Comment:        "test comment",
		SharedAs:       "",     // Empty SharedAs should be considered equal to obj1.SharedAs
		AddedAt:        999999, // Different computed fields should be ignored
		AddedBy:        "other@example.com",
		Status:         "INACTIVE",
	}

	obj3 := sharing.SharedDataObject{
		Name:           "main.table",
		DataObjectType: "TABLE",
		Comment:        "different comment", // Different comment
		SharedAs:       "alias",
		AddedAt:        123456,
		AddedBy:        "user@example.com",
		Status:         "ACTIVE",
	}

	// Test: Objects should be equal when only computed fields differ
	assert.True(t, equal(obj1, obj2), "Objects should be equal when only computed fields differ")

	// Test: Objects should not be equal when user-defined fields differ
	assert.False(t, equal(obj1, obj3), "Objects should not be equal when comment differs")
}

// TestMatchOrder tests the matchOrder function
func TestMatchOrder(t *testing.T) {
	reference := []sharing.SharedDataObject{
		{Name: "table1"},
		{Name: "table2"},
		{Name: "table3"},
	}

	target := []sharing.SharedDataObject{
		{Name: "table3"},
		{Name: "table1"},
		{Name: "table2"},
	}

	matchOrder(target, reference, func(obj sharing.SharedDataObject) string {
		return obj.Name
	})

	// Target should now have the same order as reference
	assert.Equal(t, "table1", target[0].Name)
	assert.Equal(t, "table2", target[1].Name)
	assert.Equal(t, "table3", target[2].Name)
}

// TestSuppressCDFEnabledDiff tests the suppressCDFEnabledDiff function
func TestSuppressCDFEnabledDiff(t *testing.T) {
	si := &sharing.ShareInfo{
		Name: "test-share",
		Objects: []sharing.SharedDataObject{
			{
				Name:                     "table1",
				CdfEnabled:               true,
				HistoryDataSharingStatus: "ENABLED",
			},
			{
				Name:                     "table2",
				CdfEnabled:               true,
				HistoryDataSharingStatus: "DISABLED",
			},
			{
				Name:       "table3",
				CdfEnabled: false,
			},
		},
	}

	suppressCDFEnabledDiff(si)

	// CdfEnabled should be false when HistoryDataSharingStatus is ENABLED
	assert.False(t, si.Objects[0].CdfEnabled, "CdfEnabled should be false when HistoryDataSharingStatus is ENABLED")
	// CdfEnabled should remain true when HistoryDataSharingStatus is DISABLED
	assert.True(t, si.Objects[1].CdfEnabled, "CdfEnabled should remain true when HistoryDataSharingStatus is DISABLED")
	// CdfEnabled should remain false when already false
	assert.False(t, si.Objects[2].CdfEnabled, "CdfEnabled should remain false")
}

// TestShareChanges tests the shareChanges function
func TestShareChanges(t *testing.T) {
	si := sharing.ShareInfo{
		Name:  "test-share",
		Owner: "test-owner",
		Objects: []sharing.SharedDataObject{
			{
				Name:           "table1",
				DataObjectType: "TABLE",
			},
			{
				Name:           "table2",
				DataObjectType: "TABLE",
			},
		},
	}

	// Test ADD action
	result := shareChanges(si, "ADD")
	assert.Equal(t, "test-share", result.Name)
	assert.Equal(t, "test-owner", result.Owner)
	assert.Len(t, result.Updates, 2)
	for _, update := range result.Updates {
		assert.Equal(t, sharing.SharedDataObjectUpdateActionAdd, update.Action)
	}

	// Test REMOVE action
	result = shareChanges(si, "REMOVE")
	assert.Len(t, result.Updates, 2)
	for _, update := range result.Updates {
		assert.Equal(t, sharing.SharedDataObjectUpdateActionRemove, update.Action)
	}
}
