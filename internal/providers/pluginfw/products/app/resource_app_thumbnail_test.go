package app

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	mockapps "github.com/databricks/databricks-sdk-go/experimental/mocks/service/apps"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const pngSignature = "\x89PNG\r\n\x1a\n"

func writeThumbnail(t *testing.T, content string) string {
	p := filepath.Join(t.TempDir(), "thumbnail")
	require.NoError(t, os.WriteFile(p, []byte(content), 0o600))
	return p
}

func TestReadThumbnail(t *testing.T) {
	for _, tc := range []struct {
		name    string
		content string
		wantErr string
	}{
		{"png", pngSignature + "data", ""},
		{"jpeg", "\xff\xd8\xff\xe0data", ""},
		{"png at the size limit", pngSignature + strings.Repeat("x", maxThumbnailBytes-len(pngSignature)), ""},
		{"png over the size limit", pngSignature + strings.Repeat("x", maxThumbnailBytes-len(pngSignature)+1),
			"is 250001 bytes, larger than the 250000 bytes (250 KB)"},
		{"gif", "GIF89a data", "must be a PNG or JPEG image, but its content is image/gif"},
		{"empty", "", "must be a PNG or JPEG image"},
		{"truncated png signature", pngSignature[:4], "must be a PNG or JPEG image"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			content, err := readThumbnail(writeThumbnail(t, tc.content))
			if tc.wantErr != "" {
				assert.ErrorContains(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.content, string(content))
		})
	}
}

func TestReadThumbnail_MissingFile(t *testing.T) {
	_, err := readThumbnail(filepath.Join(t.TempDir(), "missing.png"))
	assert.ErrorIs(t, err, os.ErrNotExist)
}

func TestSyncThumbnail(t *testing.T) {
	content := pngSignature + "data"
	thumbnailPath := types.StringValue(writeThumbnail(t, content))
	md5 := common.CalculateMd5Hash([]byte(content))
	upload := apps.UpdateAppThumbnailRequest{
		Name:         "my-app",
		AppThumbnail: &apps.AppThumbnail{Thumbnail: base64.StdEncoding.EncodeToString([]byte(content))},
	}
	deleteRequest := apps.DeleteAppThumbnailRequest{Name: "my-app"}
	for _, tc := range []struct {
		name          string
		thumbnailPath types.String
		priorMd5      types.String
		expect        func(*mockapps.MockAppsInterface_Expecter)
		wantMd5       types.String
		wantErr       string
	}{
		{
			name:          "uploads a new thumbnail",
			thumbnailPath: thumbnailPath,
			priorMd5:      types.StringNull(),
			expect: func(e *mockapps.MockAppsInterface_Expecter) {
				e.UpdateAppThumbnail(mock.Anything, upload).Return(&apps.AppThumbnail{}, nil)
			},
			wantMd5: types.StringValue(md5),
		},
		{
			name:          "uploads a changed thumbnail",
			thumbnailPath: thumbnailPath,
			priorMd5:      types.StringValue("old-md5"),
			expect: func(e *mockapps.MockAppsInterface_Expecter) {
				e.UpdateAppThumbnail(mock.Anything, upload).Return(&apps.AppThumbnail{}, nil)
			},
			wantMd5: types.StringValue(md5),
		},
		{
			name:          "skips an unchanged thumbnail",
			thumbnailPath: thumbnailPath,
			priorMd5:      types.StringValue(md5),
			wantMd5:       types.StringValue(md5),
		},
		{
			name:          "keeps the prior hash when the upload fails",
			thumbnailPath: thumbnailPath,
			priorMd5:      types.StringValue("old-md5"),
			expect: func(e *mockapps.MockAppsInterface_Expecter) {
				e.UpdateAppThumbnail(mock.Anything, upload).Return(nil, fmt.Errorf("Unsupported image format"))
			},
			wantMd5: types.StringValue("old-md5"),
			wantErr: "failed to update app thumbnail: Unsupported image format",
		},
		{
			name:          "deletes a removed thumbnail",
			thumbnailPath: types.StringNull(),
			priorMd5:      types.StringValue(md5),
			expect: func(e *mockapps.MockAppsInterface_Expecter) {
				e.DeleteAppThumbnail(mock.Anything, deleteRequest).Return(nil)
			},
			wantMd5: types.StringNull(),
		},
		{
			name:          "treats a missing thumbnail as deleted",
			thumbnailPath: types.StringNull(),
			priorMd5:      types.StringValue(md5),
			expect: func(e *mockapps.MockAppsInterface_Expecter) {
				e.DeleteAppThumbnail(mock.Anything, deleteRequest).Return(apierr.ErrNotFound)
			},
			wantMd5: types.StringNull(),
		},
		{
			name:          "keeps the prior hash when the delete fails",
			thumbnailPath: types.StringNull(),
			priorMd5:      types.StringValue(md5),
			expect: func(e *mockapps.MockAppsInterface_Expecter) {
				e.DeleteAppThumbnail(mock.Anything, deleteRequest).Return(apierr.ErrPermissionDenied)
			},
			wantMd5: types.StringValue(md5),
			wantErr: "failed to delete app thumbnail",
		},
		{
			name:          "does nothing without a thumbnail",
			thumbnailPath: types.StringNull(),
			priorMd5:      types.StringNull(),
			wantMd5:       types.StringNull(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// The mock fails the test on any API call the case doesn't expect.
			a := mockapps.NewMockAppsInterface(t)
			if tc.expect != nil {
				tc.expect(a.EXPECT())
			}
			gotMd5, err := syncThumbnail(context.Background(), a, "my-app", tc.thumbnailPath, tc.priorMd5)
			if tc.wantErr != "" {
				assert.ErrorContains(t, err, tc.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.wantMd5, gotMd5)
		})
	}
}
