package catalog

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

// These tests pin create responses whose full name the classic client can't put in a URL.
// Create keeps that full name as the ID, and no later read or delete can build its URL, so
// a replay journey's mandatory destroy could never remove the volume from state. The
// fixture server fails the test on a request it has no fixture for, so the tests also pin
// that the GET or PATCH after the POST is never sent. Outside these fixtures, Terraform
// prefixes each error with "cannot create volume: " and keeps the volume as tainted.

const controlCharacterVolumeError = `failed creating new request: parse ` +
	`"/api/2.1/unity-catalog/volumes/main.replay.raw\t1?": ` +
	`net/url: invalid control character in URL`

func TestVolumeCreateWithControlCharacterInReturnedFullNameFailsTheRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/volumes",
				ExpectedRequest: catalog.CreateVolumeRequestContent{
					CatalogName: "main",
					SchemaName:  "replay",
					Name:        "raw\t1",
					VolumeType:  catalog.VolumeTypeManaged,
				},
				Response: catalog.VolumeInfo{
					CatalogName: "main",
					SchemaName:  "replay",
					Name:        "raw\t1",
					FullName:    "main.replay.raw\t1",
					VolumeType:  catalog.VolumeTypeManaged,
					Owner:       "replay-admin@example.com",
				},
			},
		},
		Resource: ResourceVolume(),
		Create:   true,
		HCL: `
		name         = "raw\t1"
		catalog_name = "main"
		schema_name  = "replay"
		volume_type  = "MANAGED"
		`,
	}.Apply(t)
	assert.EqualError(t, err, controlCharacterVolumeError)
	assert.Equal(t, "main.replay.raw\t1", d.Id())
}

func TestVolumeDeleteWithControlCharacterInIDFailsBeforeTheDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceVolume(),
		Delete:   true,
		Removed:  true,
		ID:       "main.replay.raw\t1",
	}.ExpectError(t, controlCharacterVolumeError)
}

// TestVolumeDeleteWithPercentInNameFailsBeforeTheDelete pins the DELETE for the ID this
// test keeps.
func TestVolumeCreateWithOwnerAndUnbuildableReturnedFullNameFailsTheOwnerPatch(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/volumes",
				ExpectedRequest: catalog.CreateVolumeRequestContent{
					CatalogName: "main",
					SchemaName:  "replay",
					Name:        "50%_raw",
					VolumeType:  catalog.VolumeTypeManaged,
				},
				Response: catalog.VolumeInfo{
					CatalogName: "main",
					SchemaName:  "replay",
					Name:        "50%_raw",
					FullName:    "main.replay.50%_raw",
					VolumeType:  catalog.VolumeTypeManaged,
					Owner:       "replay-admin@example.com",
				},
			},
		},
		Resource: ResourceVolume(),
		Create:   true,
		HCL: `
		name         = "50%_raw"
		catalog_name = "main"
		schema_name  = "replay"
		volume_type  = "MANAGED"
		owner        = "data-engineers"
		`,
	}.Apply(t)
	assert.EqualError(t, err, `failed creating new request: parse `+
		`"/api/2.1/unity-catalog/volumes/main.replay.50%_raw": invalid URL escape "%_r"`)
	assert.Equal(t, "main.replay.50%_raw", d.Id())
}
