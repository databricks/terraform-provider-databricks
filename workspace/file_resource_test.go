package workspace

import (
	"context"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMigrateV0(t *testing.T) {
	migrated, err := MigrateV0(context.Background(), map[string]any{
		"overwrite": true,
		"source":    "../storage/testdata/tf-test-python.py",
		"foo":       "bar",
	}, nil)
	require.NoError(t, err)
	assert.Equal(t, "e4ba3a99cc1b65aff280ed8b016686b9", migrated["md5"])

	migrated, err = MigrateV0(context.Background(), map[string]any{
		"overwrite": true,
		"content":   "eA==",
		"foo":       "bar",
	}, nil)
	require.NoError(t, err)
	assert.Equal(t, "9dd4e461268c8034f5c8564e155c67a6", migrated["md5"])
}

func TestFileContentSchemaWrongSource(t *testing.T) {
	s := FileContentSchema(map[string]*schema.Schema{})
	d := s["source"].ValidateDiagFunc("__does_not_exist__", cty.GetAttrPath("x"))
	assert.NotNil(t, d)
	assert.True(t, d.HasError())
	assert.Equal(t, "File __does_not_exist__ does not exist", d[0].Summary)
}

func TestFileContentSchemaEmptyPath(t *testing.T) {
	s := FileContentSchema(map[string]*schema.Schema{})
	d := s["path"].ValidateDiagFunc("", cty.GetAttrPath("x"))
	assert.NotNil(t, d)
	assert.True(t, d.HasError())
	assert.Equal(t, "Path must not be empty", d[0].Summary)
}

func TestFileContentSchemaDbfs(t *testing.T) {
	s := FileContentSchema(map[string]*schema.Schema{})
	d := s["path"].ValidateDiagFunc("dbfs:/x", cty.GetAttrPath("x"))
	assert.NotNil(t, d)
	assert.True(t, d.HasError())
	assert.Equal(t, "Remove `dbfs:` prefix", d[0].Summary)
}

func TestFileContentSchemaClean(t *testing.T) {
	s := FileContentSchema(map[string]*schema.Schema{})
	d := s["path"].ValidateDiagFunc("/Foo/../Bar", cty.GetAttrPath("x"))
	assert.NotNil(t, d)
	assert.True(t, d.HasError())
	assert.Equal(t, "Clean path required", d[0].Summary)
}

func TestNormalizeWorkspacePath(t *testing.T) {
	testCases := []struct {
		name           string
		configuredPath string
		apiPath        string
		expected       string
	}{
		{
			name:           "API adds /Workspace prefix - should strip it",
			configuredPath: "/Users/user@example.com/notebook.py",
			apiPath:        "/Workspace/Users/user@example.com/notebook.py",
			expected:       "/Users/user@example.com/notebook.py",
		},
		{
			name:           "Config has /Workspace prefix but API doesn't - should add it",
			configuredPath: "/Workspace/Users/user@example.com/notebook.py",
			apiPath:        "/Users/user@example.com/notebook.py",
			expected:       "/Workspace/Users/user@example.com/notebook.py",
		},
		{
			name:           "Both have /Workspace prefix - no change",
			configuredPath: "/Workspace/Users/user@example.com/notebook.py",
			apiPath:        "/Workspace/Users/user@example.com/notebook.py",
			expected:       "/Workspace/Users/user@example.com/notebook.py",
		},
		{
			name:           "Neither has /Workspace prefix - no change",
			configuredPath: "/Users/user@example.com/notebook.py",
			apiPath:        "/Users/user@example.com/notebook.py",
			expected:       "/Users/user@example.com/notebook.py",
		},
		{
			name:           "Empty configured path - return API path as-is",
			configuredPath: "",
			apiPath:        "/Workspace/Users/user@example.com/notebook.py",
			expected:       "/Workspace/Users/user@example.com/notebook.py",
		},
		{
			name:           "Directory path without /Workspace in config, with /Workspace in API",
			configuredPath: "/Shared/test",
			apiPath:        "/Workspace/Shared/test",
			expected:       "/Shared/test",
		},
		{
			name:           "Directory path with /Workspace in config, without /Workspace in API",
			configuredPath: "/Workspace/Shared/test",
			apiPath:        "/Shared/test",
			expected:       "/Workspace/Shared/test",
		},
		{
			name:           "Service principal path - API adds /Workspace",
			configuredPath: "/Users/0b66cdac-04f8-408e-9290-13c058a2ebe1/file.py",
			apiPath:        "/Workspace/Users/0b66cdac-04f8-408e-9290-13c058a2ebe1/file.py",
			expected:       "/Users/0b66cdac-04f8-408e-9290-13c058a2ebe1/file.py",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NormalizeWorkspacePath(tc.configuredPath, tc.apiPath)
			assert.Equal(t, tc.expected, result)
		})
	}
}
