package catalog

import (
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func volumeJSON(t *testing.T, v any) map[string]any {
	t.Helper()
	data, err := json.Marshal(v)
	require.NoError(t, err)
	var out map[string]any
	require.NoError(t, json.Unmarshal(data, &out))
	return out
}

func TestCreateVolumeRequestToSdkGoMatchesExistingBody(t *testing.T) {
	for name, r := range map[string]catalog.CreateVolumeRequestContent{
		"required_fields": {
			CatalogName: "main", SchemaName: "replay", Name: "managed",
			VolumeType: catalog.VolumeTypeManaged,
		},
		"every_field": {
			CatalogName: "main", SchemaName: "replay", Name: "external",
			VolumeType: catalog.VolumeTypeExternal, Comment: "comment",
			StorageLocation: "s3://bucket/external",
		},
		"forced_empty_comment": {
			CatalogName: "main", SchemaName: "replay", Name: "managed",
			VolumeType: catalog.VolumeTypeManaged, ForceSendFields: []string{"Comment"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, volumeJSON(t, r), volumeJSON(t, createVolumeRequestToSdkGo(r)))
		})
	}
}

func TestUpdateVolumeRequestToSdkGoAddsOnlyFullNameArg(t *testing.T) {
	for name, r := range map[string]catalog.UpdateVolumeRequestContent{
		"owner": {Name: "main.replay.volume", Owner: "data-engineers"},
		"owner_after_create": {
			Name: "main.replay.volume", Owner: "data-engineers", Comment: "comment",
		},
		"rename_with_comment": {
			Name: "main.replay.volume", NewName: "renamed", Comment: "comment",
		},
		"cleared_comment": {
			Name: "main.replay.volume", ForceSendFields: []string{"Comment"},
		},
		"empty_owner": {Name: "main.replay.volume"},
	} {
		t.Run(name, func(t *testing.T) {
			expected := volumeJSON(t, r)
			expected["full_name_arg"] = "main.replay.volume"
			assert.Equal(t, expected, volumeJSON(t, updateVolumeRequestToSdkGo(r)))
		})
	}
}

func TestUpdateVolumeRequestToSdkGoSendsEmptyComment(t *testing.T) {
	r := catalog.UpdateVolumeRequestContent{
		Name: "main.replay.volume", ForceSendFields: []string{"Comment"},
	}
	assert.Equal(t, map[string]any{"full_name_arg": "main.replay.volume", "comment": ""},
		volumeJSON(t, updateVolumeRequestToSdkGo(r)))
}
