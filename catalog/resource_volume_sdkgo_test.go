package catalog

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

func ptr[T any](v T) *T {
	return &v
}

func TestVolumeInfoFromSdkGo(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   sdkvolumes.VolumeInfo
		want VolumeInfo
	}{
		{"managed volume", sdkvolumes.VolumeInfo{
			Name: ptr("managed"), CatalogName: ptr("main"), SchemaName: ptr("replay"),
			VolumeType: sdkvolumes.VolumeType_Managed, StorageLocation: ptr("s3://m/v1"),
			Owner: ptr("admins"), Comment: ptr("kept"), FullName: ptr("main.replay.managed"),
			VolumeId: ptr("00000000-0000-4000-8000-000000000001"),
		}, VolumeInfo{
			Name: "managed", CatalogName: "main", SchemaName: "replay",
			VolumeType: catalog.VolumeTypeManaged, StorageLocation: "s3://m/v1",
			Owner: "admins", Comment: "kept",
		}},
		{"external volume", sdkvolumes.VolumeInfo{
			Name: ptr("external"), CatalogName: ptr("main"), SchemaName: ptr("replay"),
			VolumeType: sdkvolumes.VolumeType_External, StorageLocation: ptr("s3://e/v2"),
			Owner: ptr("admins"),
		}, VolumeInfo{
			Name: "external", CatalogName: "main", SchemaName: "replay",
			VolumeType: catalog.VolumeTypeExternal, StorageLocation: "s3://e/v2",
			Owner: "admins",
		}},
		{"unknown volume type is kept", sdkvolumes.VolumeInfo{
			Name: ptr("foreign"), VolumeType: sdkvolumes.VolumeType("FOREIGN"),
		}, VolumeInfo{Name: "foreign", VolumeType: catalog.VolumeType("FOREIGN")}},
		{"absent comment", sdkvolumes.VolumeInfo{Name: ptr("v")}, VolumeInfo{Name: "v"}},
		{"empty comment", sdkvolumes.VolumeInfo{Name: ptr("v"), Comment: ptr("")},
			VolumeInfo{Name: "v"}},
		{"absent fields", sdkvolumes.VolumeInfo{}, VolumeInfo{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, volumeInfoFromSdkGo(&tc.in))
		})
	}
}

const (
	parityVolumeID = "main.replay.parity"
	parityVolume   = `"name": "parity", "catalog_name": "main", "schema_name": "replay",
		"full_name": "main.replay.parity", "owner": "admins"`
)

// parityPriorState is the state of a MANAGED volume created with comment "kept".
var parityPriorState = map[string]string{
	"id": parityVolumeID, "name": "parity", "catalog_name": "main",
	"schema_name": "replay", "volume_type": "MANAGED", "owner": "admins",
	"storage_location": "s3://m/v1", "comment": "kept",
	"volume_path": "/Volumes/main/replay/parity",
}

// noAmbientConfig keeps a test client from reading the environment or config files.
type noAmbientConfig struct{}

func (noAmbientConfig) Name() string                   { return "no ambient config" }
func (noAmbientConfig) Configure(*config.Config) error { return nil }

func noHostMetadata(context.Context, string) (*config.HostMetadata, error) {
	return nil, nil
}

// readVolumeOnPath runs the provider's refresh, or its importer when prior is nil,
// against a server that answers the GET with body, with reads through sdk-go or not.
func readVolumeOnPath(t *testing.T, sdkGo bool, prior map[string]string,
	body string) map[string]string {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		assert.Equal(t, "/api/2.1/unity-catalog/volumes/"+parityVolumeID, r.URL.Path)
		userAgent := r.Header.Get("User-Agent")
		assert.Equal(t, sdkGo, strings.Contains(userAgent, " sdk-go-uc-volumes/"), userAgent)
		fmt.Fprint(w, body)
	}))
	defer server.Close()
	inner, err := client.New(&config.Config{
		Host:                 server.URL,
		Token:                "token",
		Loaders:              []config.Loader{noAmbientConfig{}},
		HostMetadataResolver: noHostMetadata,
	})
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	c.SetCachedWorkspaceID(12345)
	r := ResourceVolume().ToResource()
	if sdkGo {
		common.EnableSdkGo(r)
	}
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	if prior == nil {
		d := r.Data(&terraform.InstanceState{ID: parityVolumeID})
		imported, err := r.Importer.StateContext(ctx, d, c)
		require.NoError(t, err)
		require.Len(t, imported, 1)
		return imported[0].State().Attributes
	}
	d := r.Data(&terraform.InstanceState{ID: parityVolumeID, Attributes: prior})
	require.Empty(t, r.ReadContext(ctx, d, c))
	return d.State().Attributes
}

// TestReadVolumeWithSdkGoSetsStateLikeExistingPath reads the same responses on both
// paths, so the converter and StructToData together must produce the existing state.
func TestReadVolumeWithSdkGoSetsStateLikeExistingPath(t *testing.T) {
	for _, tc := range []struct {
		name  string
		prior map[string]string
		body  string
		want  map[string]string
	}{
		{"refresh managed", parityPriorState, `{` + parityVolume + `,
			"volume_type": "MANAGED", "storage_location": "s3://m/v1", "comment": "kept"}`,
			map[string]string{"volume_type": "MANAGED", "comment": "kept"}},
		{"refresh external", parityPriorState, `{` + parityVolume + `,
			"volume_type": "EXTERNAL", "storage_location": "s3://e/v2"}`,
			map[string]string{"volume_type": "EXTERNAL", "storage_location": "s3://e/v2"}},
		{"refresh unknown volume type", parityPriorState, `{` + parityVolume + `,
			"volume_type": "FOREIGN", "storage_location": "s3://m/v1"}`,
			map[string]string{"volume_type": "FOREIGN"}},
		{"refresh keeps comment when absent", parityPriorState, `{` + parityVolume + `,
			"volume_type": "MANAGED", "storage_location": "s3://m/v1"}`,
			map[string]string{"comment": "kept"}},
		{"refresh keeps comment when empty", parityPriorState, `{` + parityVolume + `,
			"volume_type": "MANAGED", "storage_location": "s3://m/v1", "comment": ""}`,
			map[string]string{"comment": "kept"}},
		{"refresh reads changed comment", parityPriorState, `{` + parityVolume + `,
			"volume_type": "MANAGED", "storage_location": "s3://m/v1",
			"comment": "changed"}`,
			map[string]string{"comment": "changed"}},
		{"import with comment", nil, `{` + parityVolume + `,
			"volume_type": "EXTERNAL", "storage_location": "s3://e/v2", "comment": "kept"}`,
			map[string]string{"volume_type": "EXTERNAL", "comment": "kept",
				"volume_path": "/Volumes/main/replay/parity"}},
		{"import with absent comment", nil, `{` + parityVolume + `,
			"volume_type": "MANAGED", "storage_location": "s3://m/v1"}`,
			map[string]string{"comment": ""}},
		{"import with empty comment", nil, `{` + parityVolume + `,
			"volume_type": "MANAGED", "storage_location": "s3://m/v1", "comment": ""}`,
			map[string]string{"comment": ""}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			existing := readVolumeOnPath(t, false, tc.prior, tc.body)
			sdkGo := readVolumeOnPath(t, true, tc.prior, tc.body)
			assert.Equal(t, existing, sdkGo)
			for key, value := range tc.want {
				assert.Equal(t, value, sdkGo[key], key)
			}
		})
	}
}
