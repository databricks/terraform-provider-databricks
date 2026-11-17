package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

// These tests pin the PATCH and DELETE errors for a volume whose full name has a "%"
// that doesn't start a valid escape. The classic client puts the name in the URL
// unescaped, so building the request fails before anything is sent; the empty fixture
// list selects the real HTTP client, which fails on any request. Replay can't express
// these calls: such a volume can't be refreshed either, so a journey's mandatory
// destroy could never remove it from state. The replay journey
// name_with_percent_fails_to_build_the_url pins the GET and its full diagnostic.

const percentVolumeID = "main.replay.50%_raw"

func TestVolumeUpdateWithPercentInNameFailsBeforeThePatch(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceVolume(),
		Update:   true,
		ID:       percentVolumeID,
		InstanceState: map[string]string{
			"name":         "50%_raw",
			"catalog_name": "main",
			"schema_name":  "replay",
			"volume_type":  "MANAGED",
		},
		HCL: `
		name         = "50%_raw"
		catalog_name = "main"
		schema_name  = "replay"
		volume_type  = "MANAGED"
		comment      = "percent"
		`,
	}.ExpectError(t, `failed creating new request: parse `+
		`"/api/2.1/unity-catalog/volumes/main.replay.50%_raw": invalid URL escape "%_r"`)
}

func TestVolumeDeleteWithPercentInNameFailsBeforeTheDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceVolume(),
		Delete:   true,
		Removed:  true,
		ID:       percentVolumeID,
	}.ExpectError(t, `failed creating new request: parse `+
		`"/api/2.1/unity-catalog/volumes/main.replay.50%_raw?": invalid URL escape "%_r"`)
}
