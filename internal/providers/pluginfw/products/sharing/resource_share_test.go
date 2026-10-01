package sharing

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShareCommentChanged(t *testing.T) {
	for _, tt := range []struct {
		name  string
		plan  types.String
		state types.String
		want  bool
	}{
		// omitted or deferred comment is never sent
		{"null plan, value in state", types.StringNull(), types.StringValue("x"), false},
		{"null plan, null state", types.StringNull(), types.StringNull(), false},
		{"unknown plan, value in state", types.StringUnknown(), types.StringValue("x"), false},
		// concrete, differing comment is a real change
		{"set from null", types.StringValue("x"), types.StringNull(), true},
		{"changed value", types.StringValue("y"), types.StringValue("x"), true},
		{"explicit clear", types.StringValue(""), types.StringValue("x"), true},
		// no-op when the known comment already matches state
		{"unchanged value", types.StringValue("x"), types.StringValue("x"), false},
		{"unchanged empty", types.StringValue(""), types.StringValue(""), false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shareCommentChanged(tt.plan, tt.state))
		})
	}
}

// resolvePlannedComment models the plan value Core produces for an Optional+Computed
// attribute: config wins when set, otherwise the refreshed state value carries forward.
func resolvePlannedComment(config, refreshed types.String) types.String {
	if !config.IsNull() {
		return config
	}
	return refreshed
}

// TestShareCommentConfigToState covers what the user writes in HCL (null == omitted)
// versus what ends up in state, given what the server already holds.
func TestShareCommentConfigToState(t *testing.T) {
	ctx := context.Background()

	for _, tt := range []struct {
		name         string
		config       types.String
		serverBefore types.String
		wantSent     bool
		wantState    types.String
	}{
		// explicit value: terraform manages it, state mirrors config
		{"writes value, server empty", types.StringValue("x"), types.StringNull(), true, types.StringValue("x")},
		{"writes value, server has different value", types.StringValue("x"), types.StringValue("other"), true, types.StringValue("x")},
		{"writes value, server already matches", types.StringValue("x"), types.StringValue("x"), false, types.StringValue("x")},
		// explicit "": clears the description
		{"writes empty, server has value", types.StringValue(""), types.StringValue("other"), true, types.StringValue("")},
		{"writes empty, server already empty", types.StringValue(""), types.StringValue(""), false, types.StringValue("")},
		{"writes empty, server null", types.StringValue(""), types.StringNull(), true, types.StringValue("")},
		// omitted: not managed, state adopts whatever the server holds
		{"omits comment, server null", types.StringNull(), types.StringNull(), false, types.StringNull()},
		{"omits comment, server has value", types.StringNull(), types.StringValue("set in UI"), false, types.StringValue("set in UI")},
		{"omits comment, server empty string", types.StringNull(), types.StringValue(""), false, types.StringValue("")},
	} {
		t.Run(tt.name, func(t *testing.T) {
			planned := resolvePlannedComment(tt.config, tt.serverBefore)

			sent := shareCommentChanged(planned, tt.serverBefore)
			assert.Equal(t, tt.wantSent, sent, "comment sent on the wire")

			serverAfter := tt.serverBefore
			if sent {
				serverAfter = planned
			}
			apiShare := sharing.ShareInfo{Name: "s"}
			if !serverAfter.IsNull() {
				// a present comment lands in ForceSendFields on unmarshal, so "" converts
				// to a known "" rather than null
				apiShare.Comment = serverAfter.ValueString()
				apiShare.ForceSendFields = append(apiShare.ForceSendFields, "Comment")
			}

			var state ShareInfoExtended
			assert.False(t, converters.GoSdkToTfSdkStruct(ctx, apiShare, &state).HasError())
			assert.True(t, state.Comment.Equal(tt.wantState),
				"state comment: got %v, want %v", state.Comment, tt.wantState)

			// planned must equal post-apply state or core rejects the apply
			assert.True(t, state.Comment.Equal(planned),
				"planned %v must equal post-apply state %v", planned, state.Comment)
		})
	}
}

func TestShareSyncEffectiveFields(t *testing.T) {
	shareName := "test-share-name"
	ctx := context.Background()
	shares := ShareResource{}

	tests := []struct {
		name       string
		planGoSDK  sharing.ShareInfo
		stateGoSDK sharing.ShareInfo
	}{
		{
			name: "plan with less objects",
			planGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
			stateGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-2",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-2"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
		},
		{
			name: "plan with more objects",
			planGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-2",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-2"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
			stateGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
		},
		{
			name: "empty plan",
			planGoSDK: sharing.ShareInfo{
				Name:    shareName,
				Objects: []sharing.SharedDataObject{},
			},
			stateGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var planTFSDK ShareInfoExtended
			diagnostics := converters.GoSdkToTfSdkStruct(ctx, tt.planGoSDK, &planTFSDK)
			assert.False(t, diagnostics.HasError())

			var stateTFSDK ShareInfoExtended
			diagnostics = converters.GoSdkToTfSdkStruct(ctx, tt.stateGoSDK, &stateTFSDK)
			assert.False(t, diagnostics.HasError())

			_, diagnostics = shares.syncEffectiveFields(ctx, planTFSDK, stateTFSDK, effectiveFieldsActionRead{})
			assert.False(t, diagnostics.HasError())
		})
	}
}

func TestReadShareStatePreservesUnsetCDF(t *testing.T) {
	ctx := context.Background()
	objectName := "catalog.schema.table"
	apiShare := &sharing.ShareInfo{
		Name: "share",
		Objects: []sharing.SharedDataObject{
			{
				Name:                     objectName,
				DataObjectType:           "TABLE",
				CdfEnabled:               true,
				HistoryDataSharingStatus: "ENABLED",
			},
		},
	}

	var existingState ShareInfoExtended
	require.False(t, converters.GoSdkToTfSdkStruct(ctx, apiShare, &existingState).HasError())
	existingObjects, ok := existingState.GetObjects(ctx)
	require.True(t, ok)
	existingObjects[0].CdfEnabled = types.BoolNull()
	existingObjects[0].EffectiveCdfEnabled = types.BoolValue(true)
	existingObjects[0].HistoryDataSharingStatus = types.StringNull()
	existingObjects[0].EffectiveHistoryDataSharingStatus = types.StringValue("ENABLED")
	existingState.SetObjects(ctx, existingObjects)

	var refreshedState ShareInfoExtended
	require.False(t, converters.GoSdkToTfSdkStruct(ctx, apiShare, &refreshedState).HasError())
	state, diagnostics := (&ShareResource{}).syncEffectiveFields(ctx, existingState, refreshedState, effectiveFieldsActionRead{})
	require.False(t, diagnostics.HasError())
	objects, ok := state.GetObjects(ctx)
	require.True(t, ok)
	require.Len(t, objects, 1)
	assert.True(t, objects[0].CdfEnabled.IsNull(), "cdf_enabled should remain unset")
	assert.True(t, objects[0].EffectiveCdfEnabled.ValueBool(), "effective_cdf_enabled should reflect the API")
}

func TestShareSyncEffectiveFieldsCreateOrUpdateSyncsObjectsOnceByName(t *testing.T) {
	ctx := context.Background()
	planGoSDK := sharing.ShareInfo{
		Name: "test-share-name",
		Objects: []sharing.SharedDataObject{
			{Name: "obj-1"},
			{Name: "obj-2", SharedAs: "configured-alias"},
		},
	}
	serverGoSDK := sharing.ShareInfo{
		Name: "test-share-name",
		Objects: []sharing.SharedDataObject{
			{Name: "obj-2", SharedAs: "server-alias-2"},
			{Name: "obj-1", SharedAs: "server-alias-1"},
		},
	}

	var plan ShareInfoExtended
	diagnostics := converters.GoSdkToTfSdkStruct(ctx, planGoSDK, &plan)
	assert.False(t, diagnostics.HasError())
	var serverState ShareInfoExtended
	diagnostics = converters.GoSdkToTfSdkStruct(ctx, serverGoSDK, &serverState)
	assert.False(t, diagnostics.HasError())

	result, diagnostics := (&ShareResource{}).syncEffectiveFields(ctx, plan, serverState, effectiveFieldsActionCreateOrUpdate{})
	assert.False(t, diagnostics.HasError())
	objects, ok := result.GetObjects(ctx)
	assert.True(t, ok)
	assert.Len(t, objects, 2)
	objectsByName := map[string]sharing_tf.SharedDataObject_SdkV2{}
	for _, object := range objects {
		objectsByName[object.Name.ValueString()] = object
	}

	assert.True(t, objectsByName["obj-1"].SharedAs.IsNull())
	assert.Equal(t, "server-alias-1", objectsByName["obj-1"].EffectiveSharedAs.ValueString())
	assert.Equal(t, "configured-alias", objectsByName["obj-2"].SharedAs.ValueString())
	assert.Equal(t, "server-alias-2", objectsByName["obj-2"].EffectiveSharedAs.ValueString())
}

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

// TestDiffUpdateObjectComment tests updating an object's comment field
func TestDiffUpdateObjectComment(t *testing.T) {
	before := sharing.ShareInfo{
		Name: "test-share",
		Objects: []sharing.SharedDataObject{
			{
				Name:           "catalog.schema.table",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table",
				Comment:        "original comment",
			},
		},
	}

	after := sharing.ShareInfo{
		Name: "test-share",
		Objects: []sharing.SharedDataObject{
			{
				Name:           "catalog.schema.table",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table",
				Comment:        "updated comment",
			},
		},
	}

	changes := diff(before, after)
	assert.Len(t, changes, 1, "Should generate UPDATE for comment change")
	assert.Equal(t, sharing.SharedDataObjectUpdateActionUpdate, changes[0].Action)
	assert.Equal(t, "updated comment", changes[0].DataObject.Comment)

	// SharedAs must be preserved in UPDATE operations
	assert.Equal(t, "schema.table", changes[0].DataObject.SharedAs,
		"SharedAs must be preserved in UPDATE operations")
}

// TestDiffAddAndUpdate tests a combination of ADD and UPDATE operations
func TestDiffAddAndUpdate(t *testing.T) {
	before := sharing.ShareInfo{
		Name: "test-share",
		Objects: []sharing.SharedDataObject{
			{
				Name:           "catalog.schema.table1",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table1",
				Comment:        "comment1",
			},
		},
	}

	after := sharing.ShareInfo{
		Name: "test-share",
		Objects: []sharing.SharedDataObject{
			{
				Name:           "catalog.schema.table1",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table1",
				Comment:        "updated comment1",
			},
			{
				Name:           "catalog.schema.table2",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table2",
				Comment:        "comment2",
			},
		},
	}

	changes := diff(before, after)
	assert.Len(t, changes, 2, "Should generate both UPDATE and ADD")

	var hasUpdate, hasAdd bool
	for _, change := range changes {
		if change.Action == sharing.SharedDataObjectUpdateActionUpdate {
			hasUpdate = true
			assert.Equal(t, "catalog.schema.table1", change.DataObject.Name)
			// SharedAs must be preserved in UPDATE operations
			assert.Equal(t, "schema.table1", change.DataObject.SharedAs,
				"SharedAs must be preserved in UPDATE operations")
		} else if change.Action == sharing.SharedDataObjectUpdateActionAdd {
			hasAdd = true
			assert.Equal(t, "catalog.schema.table2", change.DataObject.Name)
			assert.Equal(t, "schema.table2", change.DataObject.SharedAs)
		}
	}
	assert.True(t, hasUpdate, "Should have UPDATE action")
	assert.True(t, hasAdd, "Should have ADD action")
}

// TestDiffAddUpdateRemove tests all three operations together
func TestDiffAddUpdateRemove(t *testing.T) {
	before := sharing.ShareInfo{
		Name: "test-share",
		Objects: []sharing.SharedDataObject{
			{
				Name:           "catalog.schema.table1",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table1",
				Comment:        "comment1",
			},
			{
				Name:           "catalog.schema.table2",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table2",
				Comment:        "comment2",
			},
		},
	}

	after := sharing.ShareInfo{
		Name: "test-share",
		Objects: []sharing.SharedDataObject{
			{
				Name:           "catalog.schema.table1",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table1",
				Comment:        "updated comment1",
			},
			{
				Name:           "catalog.schema.table3",
				DataObjectType: "TABLE",
				SharedAs:       "schema.table3",
				Comment:        "comment3",
			},
		},
	}

	changes := diff(before, after)
	assert.Len(t, changes, 3, "Should generate REMOVE, UPDATE, and ADD")

	var hasUpdate, hasAdd, hasRemove bool
	for _, change := range changes {
		switch change.Action {
		case sharing.SharedDataObjectUpdateActionUpdate:
			hasUpdate = true
			assert.Equal(t, "catalog.schema.table1", change.DataObject.Name)
			// SharedAs must be preserved in UPDATE operations
			assert.Equal(t, "schema.table1", change.DataObject.SharedAs,
				"SharedAs must be preserved in UPDATE operations")
		case sharing.SharedDataObjectUpdateActionAdd:
			hasAdd = true
			assert.Equal(t, "catalog.schema.table3", change.DataObject.Name)
		case sharing.SharedDataObjectUpdateActionRemove:
			hasRemove = true
			assert.Equal(t, "catalog.schema.table2", change.DataObject.Name)
		}
	}
	assert.True(t, hasUpdate, "Should have UPDATE action")
	assert.True(t, hasAdd, "Should have ADD action")
	assert.True(t, hasRemove, "Should have REMOVE action")
}
