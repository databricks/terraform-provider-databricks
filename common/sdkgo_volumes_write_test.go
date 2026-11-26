package common

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/useragent"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sendBothWays sends one call through sdk-go and then the same call through the classic
// client, and returns the request each sent.
func sendBothWays(t *testing.T, response string,
	sdkGo func(context.Context, *SdkGoVolumes) error,
	classic func(context.Context, catalog.VolumesInterface) error) (sentRequest, sentRequest) {
	t.Helper()
	var sent []sentRequest
	record := func(w http.ResponseWriter, r *http.Request) {
		sent = append(sent, recordRequest(t, r))
		respond(http.StatusOK, response)(w, r)
	}
	server, _ := serveSequence(t, record, record)
	w := newVolumesTestWorkspaceClient(t, server, nil)
	ctx := useragent.InContext(context.Background(), "resource", "volume")
	vc, err := SdkGoVolumesClient(ctx, w)
	require.NoError(t, err)
	require.NoError(t, sdkGo(ctx, vc))
	require.NoError(t, classic(ctx, w.Volumes))
	require.Len(t, sent, 2)
	return sent[0], sent[1]
}

// assertSameRequest checks that sdk-go sent the classic request, and only added the
// sdk-go modules to the User-Agent.
func assertSameRequest(t *testing.T, sdkGo, classic sentRequest) {
	t.Helper()
	modules, ok := strings.CutPrefix(sdkGo.userAgent, classic.userAgent+" ")
	require.True(t, ok, "%q extends %q", sdkGo.userAgent, classic.userAgent)
	for _, module := range strings.Fields(modules) {
		assert.True(t, strings.HasPrefix(module, "sdk-go-"), module)
	}
	sdkGo.userAgent, classic.userAgent = "", ""
	assert.Equal(t, classic, sdkGo)
}

func TestSdkGoVolumesClientCreatesWithTheClassicRequest(t *testing.T) {
	catalogName, schemaName, name := "main", "replay", "a#b"
	comment, location := "<comment>", "s3://bucket/a"
	sdkGo, classic := sendBothWays(t, sdkGoTestVolume,
		func(ctx context.Context, vc *SdkGoVolumes) error {
			_, err := vc.Create(ctx, sdkvolumes.CreateVolumeRequest{
				CatalogName: &catalogName, SchemaName: &schemaName, Name: &name,
				VolumeType: sdkvolumes.VolumeType_External, StorageLocation: &location,
				Comment: &comment,
			})
			return err
		},
		func(ctx context.Context, volumes catalog.VolumesInterface) error {
			_, err := volumes.Create(ctx, catalog.CreateVolumeRequestContent{
				CatalogName: catalogName, SchemaName: schemaName, Name: name,
				VolumeType: catalog.VolumeTypeExternal, StorageLocation: location,
				Comment: comment,
			})
			return err
		})
	assertSameRequest(t, sdkGo, classic)
}

// The classic client formats the name into the URL of a PATCH or DELETE unescaped, so
// in a name "#" drops the rest, "?" starts a query and "%41" is decoded. sdk-go writes
// send the same request for every name, and the PATCH body has no full_name_arg.
func TestSdkGoVolumesClientWritesWithTheClassicRequestForEveryName(t *testing.T) {
	for _, name := range []string{"main.replay.volume", "main.replay.a#b",
		"main.replay.a?b", "main.replay.a%41", "main.replay.a b",
		"main.replay.café"} {
		t.Run(name, func(t *testing.T) {
			comment, newName, owner, empty := "<comment>", "renamed", "owner@example.com", ""
			for _, tc := range []struct {
				sdkGo   sdkvolumes.UpdateVolumeRequest
				classic catalog.UpdateVolumeRequestContent
			}{
				{sdkvolumes.UpdateVolumeRequest{FullNameArg: &name, Comment: &comment,
					NewName: &newName, Owner: &owner},
					catalog.UpdateVolumeRequestContent{Name: name, Comment: comment,
						NewName: newName, Owner: owner}},
				{sdkvolumes.UpdateVolumeRequest{FullNameArg: &name, Comment: &empty},
					catalog.UpdateVolumeRequestContent{Name: name,
						ForceSendFields: []string{"Comment"}}},
			} {
				sdkGo, classic := sendBothWays(t, sdkGoTestVolume,
					func(ctx context.Context, vc *SdkGoVolumes) error {
						_, err := vc.Update(ctx, tc.sdkGo)
						return err
					},
					func(ctx context.Context, volumes catalog.VolumesInterface) error {
						_, err := volumes.Update(ctx, tc.classic)
						return err
					})
				assertSameRequest(t, sdkGo, classic)
			}
			sdkGo, classic := sendBothWays(t, "{}",
				func(ctx context.Context, vc *SdkGoVolumes) error {
					return vc.DeleteByName(ctx, name)
				},
				func(ctx context.Context, volumes catalog.VolumesInterface) error {
					return volumes.DeleteByName(ctx, name)
				})
			assertSameRequest(t, sdkGo, classic)
		})
	}
}
