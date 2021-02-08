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
	migrated, err := MigrateV0(context.Background(), map[string]interface{}{
		"overwrite": true,
		"source":    "../storage/testdata/tf-test-python.py",
		"foo":       "bar",
	}, nil)
	require.NoError(t, err)
	assert.Equal(t, "e4ba3a99cc1b65aff280ed8b016686b9", migrated["md5"])

	migrated, err = MigrateV0(context.Background(), map[string]interface{}{
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
