package common

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestStringIsUUID(t *testing.T) {
	assert.True(t, StringIsUUID("3f670caf-9a4b-4479-8143-1a0878da8f57"))
	assert.False(t, StringIsUUID("abc"))
}

func TestGetTerraformVersionFromContext(t *testing.T) {
	assert.Equal(t, "unknown", GetTerraformVersionFromContext(context.Background()))

	//
	p := &schema.Provider{}
	p.TerraformVersion = "exporter"
	ctx := context.WithValue(context.Background(), Provider, p)
	assert.Equal(t, "exporter", GetTerraformVersionFromContext(ctx))

	//
	assert.True(t, IsExporter(ctx))
}

func TestSuppressDiffWhitespaceChange(t *testing.T) {
	assert.True(t, SuppressDiffWhitespaceChange("k", "value", "  value  ", nil))
	assert.False(t, SuppressDiffWhitespaceChange("k", "value", "new_value", nil))
}

func TestMustInt64(t *testing.T) {
	assert.Equal(t, int64(123), MustInt64("123"))
}

func TestReadFileContent(t *testing.T) {
	tmpDir := fmt.Sprintf("/tmp/Dashboard-%f", rand.Float64())
	fileName := tmpDir + "/Dashboard.json"
	os.Mkdir(tmpDir, 0755)
	os.WriteFile(fileName, []byte("hello"), 0644)
	content, err := ReadFileContent(fileName)
	assert.Equal(t, []byte("hello"), content)
	assert.NoError(t, err)
}

func TestCalculateMd5Hash(t *testing.T) {
	hash := CalculateMd5Hash([]byte("hello"))
	assert.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte("hello"))), hash)
}

func TestReadSerializedJsonContent(t *testing.T) {
	_, md5Hash, err := ReadSerializedJsonContent("hello", "")
	assert.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte("hello"))), md5Hash)
	assert.NoError(t, err)

	tmpDir := fmt.Sprintf("/tmp/Dashboard-%f", rand.Float64())
	fileName := tmpDir + "/Dashboard.json"
	os.Mkdir(tmpDir, 0755)
	os.WriteFile(fileName, []byte("hello"), 0644)
	_, md5Hash, err = ReadSerializedJsonContent("", fileName)
	assert.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte("hello"))), md5Hash)
	assert.NoError(t, err)
}

func TestIgnoreNotFoundError(t *testing.T) {
	err := IgnoreNotFoundError(nil)
	assert.NoError(t, err)

	err = IgnoreNotFoundError(fmt.Errorf("error"))
	assert.EqualError(t, err, "error")

	err = IgnoreNotFoundError(&apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    "error",
	})
	assert.NoError(t, err)

	err = IgnoreNotFoundError(&apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    "cluster xyz does not exist",
	})
	assert.NoError(t, err)
}
