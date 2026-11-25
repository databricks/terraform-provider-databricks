package catalog

import (
	"encoding/json"
	"maps"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests pin that the generated mapping in resource_volume_sdkgo_gen.go sends the
// bodies and leaves the state that the existing reflection path does.

const mappingVolumeID = "main.default.mapped_volume"

// volumeMappingConfigs are configurations whose request bodies must match.
var volumeMappingConfigs = map[string]map[string]any{
	"every attribute set": {
		"catalog_name":     "main",
		"schema_name":      "default",
		"name":             "mapped_volume",
		"volume_type":      "EXTERNAL",
		"storage_location": "s3://bucket/volumes/mapped",
		"comment":          "a comment",
		"owner":            "owner@example.com",
	},
	"optional attributes empty": {
		"catalog_name": "main",
		"schema_name":  "default",
		"name":         "mapped_volume",
		"volume_type":  "MANAGED",
	},
	"optional attributes set to empty strings": {
		"catalog_name":     "main",
		"schema_name":      "default",
		"name":             "mapped_volume",
		"volume_type":      "MANAGED",
		"storage_location": "",
		"comment":          "",
		"owner":            "",
	},
	"unicode and upper-case names": {
		"catalog_name": "Main_CATALOG",
		"schema_name":  "Schéma_Ü",
		"name":         "Volüme_名前",
		"volume_type":  "MANAGED",
		"comment":      "ünïcødé ✓",
		"owner":        "Group Ä",
	},
	"external volume": {
		"catalog_name":     "main",
		"schema_name":      "default",
		"name":             "mapped_volume",
		"volume_type":      "EXTERNAL",
		"storage_location": "abfss://container@account.dfs.core.windows.net/path",
	},
}

func TestVolumeCreateRequestFromDataSendsTheExistingBody(t *testing.T) {
	r := ResourceVolume()
	for name, config := range volumeMappingConfigs {
		t.Run(name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, r.Schema, config)
			var existing catalog.CreateVolumeRequestContent
			common.DataToStructPointer(d, r.Schema, &existing)
			assertSameJSON(t, existing, volumeCreateRequestFromData(d))
		})
	}
}

func TestVolumeUpdateRequestFromDataSendsTheExistingBody(t *testing.T) {
	r := ResourceVolume()
	for name, config := range volumeMappingConfigs {
		t.Run(name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, r.Schema, config)
			assertSameUpdateBody(t, r.Schema, d)
		})
	}
	t.Run("comment and owner cleared", func(t *testing.T) {
		state := &terraform.InstanceState{ID: mappingVolumeID}
		state.Attributes = map[string]string{
			"id":           mappingVolumeID,
			"name":         "mapped_volume",
			"catalog_name": "main",
			"schema_name":  "default",
			"volume_type":  "MANAGED",
			"comment":      "first",
			"owner":        "owner@example.com",
		}
		diff := &terraform.InstanceDiff{}
		diff.Attributes = map[string]*terraform.ResourceAttrDiff{
			"comment": {Old: "first", New: ""},
			"owner":   {Old: "owner@example.com", New: ""},
		}
		d, err := schema.InternalMap(r.Schema).Data(state, diff)
		require.NoError(t, err)
		assertSameUpdateBody(t, r.Schema, d)
	})
}

func assertSameUpdateBody(t *testing.T, s map[string]*schema.Schema,
	d *schema.ResourceData) {
	t.Helper()
	var existing catalog.UpdateVolumeRequestContent
	common.DataToStructPointer(d, s, &existing)
	assertSameJSON(t, existing, volumeUpdateRequestFromData(d))
}

func assertSameJSON(t *testing.T, existing, sdkGo any) {
	t.Helper()
	existingBody, err := json.Marshal(existing)
	require.NoError(t, err)
	sdkGoBody, err := json.Marshal(sdkGo)
	require.NoError(t, err)
	assert.JSONEq(t, string(existingBody), string(sdkGoBody))
}

// volumeMappingResponses are GetVolume responses whose decodes must leave the same
// state.
var volumeMappingResponses = map[string]string{
	"every field": `{
		"catalog_name": "main",
		"schema_name": "default",
		"name": "mapped_volume",
		"full_name": "main.default.mapped_volume",
		"volume_type": "EXTERNAL",
		"storage_location": "s3://bucket/volumes/mapped",
		"owner": "owner@example.com",
		"comment": "a comment",
		"volume_id": "01234567-89ab-cdef-0123-456789abcdef",
		"metastore_id": "fedcba98-7654-3210-fedc-ba9876543210",
		"created_at": 1700000000000,
		"created_by": "creator@example.com",
		"updated_at": 1700000001000,
		"updated_by": "updater@example.com",
		"access_point": "arn:aws:s3:us-west-2:123456789012:accesspoint/mapped",
		"browse_only": false
	}`,
	"optional fields missing": `{
		"catalog_name": "main",
		"schema_name": "default",
		"name": "mapped_volume",
		"full_name": "main.default.mapped_volume",
		"volume_type": "MANAGED"
	}`,
	"empty strings": `{
		"catalog_name": "",
		"schema_name": "",
		"name": "",
		"full_name": "",
		"volume_type": "",
		"storage_location": "",
		"owner": "",
		"comment": ""
	}`,
	"unicode and upper-case names": `{
		"catalog_name": "Main_CATALOG",
		"schema_name": "Schéma_Ü",
		"name": "Volüme_名前",
		"full_name": "Main_CATALOG.Schéma_Ü.Volüme_名前",
		"volume_type": "MANAGED",
		"owner": "Group Ä",
		"comment": "ünïcødé ✓"
	}`,
}

// volumeMappingPriorStates are the states that a read starts from.
var volumeMappingPriorStates = map[string]struct {
	attributes map[string]string
	// isNew is true for an import, which reads into a new resource.
	isNew bool
}{
	"import": {attributes: map[string]string{"id": mappingVolumeID}, isNew: true},
	"every attribute set": {attributes: map[string]string{
		"id":               mappingVolumeID,
		"catalog_name":     "old_catalog",
		"schema_name":      "old_schema",
		"name":             "old_name",
		"volume_type":      "MANAGED",
		"storage_location": "s3://bucket/old",
		"owner":            "old@example.com",
		"comment":          "old comment",
	}},
	"optional attributes unset": {attributes: map[string]string{
		"id":           mappingVolumeID,
		"catalog_name": "old_catalog",
		"schema_name":  "old_schema",
		"name":         "old_name",
		"volume_type":  "MANAGED",
	}},
}

func TestVolumeInfoToDataLeavesTheExistingState(t *testing.T) {
	r := ResourceVolume()
	for responseName, body := range volumeMappingResponses {
		var existing catalog.VolumeInfo
		require.NoError(t, json.Unmarshal([]byte(body), &existing), responseName)
		var sdkGo sdkvolumes.VolumeInfo
		require.NoError(t, json.Unmarshal([]byte(body), &sdkGo), responseName)
		for stateName, prior := range volumeMappingPriorStates {
			t.Run(responseName+" into "+stateName, func(t *testing.T) {
				existingData := priorVolumeData(r, prior.attributes, prior.isNew)
				require.NoError(t, common.StructToData(&existing, r.Schema, existingData))
				sdkGoData := priorVolumeData(r, prior.attributes, prior.isNew)
				require.NoError(t, volumeInfoToData(&sdkGo, sdkGoData))
				assert.Equal(t, existingData.State().Attributes,
					sdkGoData.State().Attributes)
			})
		}
	}
}

// TestVolumeInfoToDataSetsTheResponse guards the comparison above against a mapping and
// a reference that both leave the state unchanged.
func TestVolumeInfoToDataSetsTheResponse(t *testing.T) {
	r := ResourceVolume()
	var v sdkvolumes.VolumeInfo
	require.NoError(t, json.Unmarshal([]byte(volumeMappingResponses["every field"]), &v))
	d := priorVolumeData(r, volumeMappingPriorStates["import"].attributes, true)
	require.NoError(t, volumeInfoToData(&v, d))
	assert.Equal(t, map[string]any{
		"catalog_name":     "main",
		"schema_name":      "default",
		"name":             "mapped_volume",
		"volume_type":      "EXTERNAL",
		"storage_location": "s3://bucket/volumes/mapped",
		"owner":            "owner@example.com",
		"comment":          "a comment",
	}, map[string]any{
		"catalog_name":     d.Get("catalog_name"),
		"schema_name":      d.Get("schema_name"),
		"name":             d.Get("name"),
		"volume_type":      d.Get("volume_type"),
		"storage_location": d.Get("storage_location"),
		"owner":            d.Get("owner"),
		"comment":          d.Get("comment"),
	})
}

func priorVolumeData(r common.Resource, attributes map[string]string,
	isNew bool) *schema.ResourceData {
	d := r.ToResource().Data(&terraform.InstanceState{ID: mappingVolumeID,
		Attributes: maps.Clone(attributes)})
	if isNew {
		d.MarkNewResource()
	}
	return d
}
