package sdkgoutility

import (
	"reflect"
	"testing"

	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/stretchr/testify/assert"
)

func TestFieldName_RegisteredModelDatabricks(t *testing.T) {
	structType := reflect.TypeOf(modelregistry.RegisteredModelDatabricks{})
	want := map[string]string{
		"Name":                 "name",
		"CreationTimestamp":    "creation_timestamp",
		"LastUpdatedTimestamp": "last_updated_timestamp",
		"UserId":               "user_id",
		"Description":          "description",
		"LatestVersions":       "latest_versions",
		"Id":                   "id",
		"PermissionLevel":      "permission_level",
		"Tags":                 "tags",
	}
	assert.Equal(t, len(want), structType.NumField())
	for i := range structType.NumField() {
		field := structType.Field(i)
		assert.Equal(t, want[field.Name], FieldName(structType, field), field.Name)
	}
}
