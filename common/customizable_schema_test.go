package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/stretchr/testify/assert"
)

var testCustomizableSchemaScm = StructToSchema(testStruct{}, nil)

func TestCustomizableSchemaSetOptional(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetOptional()
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].Optional, "optional should be overriden to true in field: non-optional")
}

func TestCustomizableSchemaSetRequired(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "float").SetRequired()
	assert.Truef(t, testCustomizableSchemaScm["float"].Required, "required should be overriden to true in field: float")
}

func TestCustomizableSchemaSetReadOnly(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "bool").SetReadOnly()
	assert.Truef(t, testCustomizableSchemaScm["bool"].Computed, "computed should be overriden to true in field: bool")
	assert.Falsef(t, testCustomizableSchemaScm["bool"].Optional, "optional should be overriden to false in field: bool")
	assert.Falsef(t, testCustomizableSchemaScm["bool"].Required, "required should be overriden to false in field: bool")
	assert.Truef(t, testCustomizableSchemaScm["bool"].MaxItems == 0, "maxItems should be overriden to 0 in field: bool")
}

func TestCustomizableSchemaSetComputed(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "string").SetComputed()
	assert.Truef(t, testCustomizableSchemaScm["string"].Computed, "computed should be overriden to true in field: string")
}

func TestCustomizableSchemaSetDefault(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetDefault("abc")
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].Default == "abc", "default should be overriden to abc in field: non_optional")
}

func TestCustomizableSchemaSetSuppressDiff(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetSuppressDiff()
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].DiffSuppressFunc != nil, "DiffSuppressfunc should be set in field: non_optional")
}

func TestCustomizableSchemaSetCustomSuppressDiffWithDefault(t *testing.T) {
	tc := []struct {
		name      string
		fieldName string
		dv        any
		old       string
		new       string
		result    bool
	}{
		{
			name:      "default string",
			fieldName: "string",
			dv:        "foobar",
			old:       "foobar",
			new:       "",
			result:    true,
		},
		{
			name:      "default int",
			fieldName: "integer",
			dv:        123,
			old:       "123",
			new:       "0",
			result:    true,
		},
		{
			name:      "default bool",
			fieldName: "bool",
			dv:        true,
			old:       "true",
			new:       "false",
			result:    true,
		},
		{
			name:      "default float",
			fieldName: "float",
			dv:        123.456,
			old:       "123.456",
			new:       "0",
			result:    true,
		},
		{
			name:      "non default string",
			fieldName: "string",
			dv:        "foobar",
			old:       "non-default-val",
			new:       "",
			result:    false,
		},
		{
			name:      "non default int",
			fieldName: "integer",
			dv:        123,
			old:       "non-default-val",
			new:       "0",
			result:    false,
		},
		{
			name:      "non default bool",
			fieldName: "bool",
			dv:        true,
			old:       "non-default-val",
			new:       "false",
			result:    false,
		},
		{
			name:      "non default float",
			fieldName: "float",
			dv:        123.456,
			old:       "non-default-val",
			new:       "0",
			result:    false,
		},
		{
			name:      "override in config",
			fieldName: "string",
			dv:        "foobar",
			old:       "foobar",
			new:       "new-val-in-config",
			result:    false,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			CustomizeSchemaPath(testCustomizableSchemaScm, tt.fieldName).SetSuppressDiffWithDefault(tt.dv)
			assert.Truef(t, testCustomizableSchemaScm[tt.fieldName].DiffSuppressFunc != nil, "DiffSuppressFunc should be set in field: %s", tt.fieldName)

			assert.Equal(t, tt.result, testCustomizableSchemaScm[tt.fieldName].DiffSuppressFunc("", tt.old, tt.new, nil))
		})
	}
}

func TestCustomizableSchemaSetCustomSuppressDiff(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetCustomSuppressDiff(diffSuppressor("test", CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").Schema))
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].DiffSuppressFunc != nil, "DiffSuppressfunc should be set in field: non_optional")
}

func TestCustomizableSchemaSetSensitive(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetSensitive()
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].Sensitive, "sensitive should be overriden to true in field: non_optional")
}

func TestCustomizableSchemaSetForceNew(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetForceNew()
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].ForceNew, "forcenew should be overriden to true in field: non_optional")
}

func TestCustomizableSchemaSetMaxItems(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "int_slice").SetMaxItems(5)
	assert.Truef(t, testCustomizableSchemaScm["int_slice"].MaxItems == 5, "maxItems should be overriden to 5 in field: int_slice")
}

func TestCustomizableSchemaSetMinItems(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "int_slice").SetMinItems(3)
	assert.Truef(t, testCustomizableSchemaScm["int_slice"].MinItems == 3, "maxItems should be overriden to 5 in field: int_slice")
}

func TestCustomizableSchemaSetConflictsWith(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetConflictsWith([]string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["non_optional"].ConflictsWith) == 1, "conflictsWith should be set in field: non_optional")
}

func TestCustomizableSchemaSetConflictsWith_PathInContext(t *testing.T) {
	fakeContextWithPath := schemaPathContext{
		path:       []string{"a", "0", "b"},
		schemaPath: []*schema.Schema{},
	}
	cs := CustomizeSchemaPath(testCustomizableSchemaScm, "float")
	cs.context = fakeContextWithPath
	cs.SetConflictsWith([]string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["float"].ConflictsWith) == 1, "conflictsWith should be set in field: float")
	assert.Truef(t, cs.Schema.ConflictsWith[0] == "a.0.b.abc", "conflictsWith should be set with the correct prefix")
}

func TestCustomizableSchemaSetConflictsWith_MultiItemList(t *testing.T) {
	fakeContextWithPath := schemaPathContext{
		path: []string{"a", "0", "b"},
		schemaPath: []*schema.Schema{
			{
				Type:     schema.TypeList,
				MaxItems: 10,
			},
		},
	}
	cs := CustomizeSchemaPath(testCustomizableSchemaScm, "bool")
	cs.context = fakeContextWithPath
	cs.SetConflictsWith([]string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["bool"].ConflictsWith) == 0, "conflictsWith should not be set when there's multi item list in the path")
}

func TestCustomizableSchemaSetExactlyOneOf(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetExactlyOneOf([]string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["non_optional"].ExactlyOneOf) == 1, "ExactlyOneOf should be set in field: non_optional")
}

func TestCustomizableSchemaAtLeastOneOf(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetAtLeastOneOf([]string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["non_optional"].AtLeastOneOf) == 1, "AtLeastOneOf should be set in field: non_optional")
}

func TestCustomizableSchemaSetRequiredWith(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetRequiredWith([]string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["non_optional"].RequiredWith) == 1, "RequiredWith should be set in field: non_optional")
}
func TestCustomizableSchemaSetDeprecated(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetDeprecated("test reason")
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].Deprecated == "test reason", "deprecated should be overriden in field: non_optional")
}

func TestCustomizableSchemaSetValidateFunc(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetValidateFunc(validation.StringInSlice([]string{"PHOTON", "STANDARD"}, false))
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].ValidateFunc != nil, "validateFunc should be set in field: non_optional")
}

func TestCustomizableSchemaSetValidateDiagFunc(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetValidateDiagFunc(validation.ToDiagFunc(validation.IntAtLeast(0)))
	assert.Truef(t, testCustomizableSchemaScm["non_optional"].ValidateDiagFunc != nil, "validateDiagFunc should be set in field: non_optional")
}

func TestCustomizableSchemaAddNewField(t *testing.T) {

	CustomizeSchemaPath(testCustomizableSchemaScm).AddNewField("test", &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	})

	assert.Truef(t, MustSchemaPath(testCustomizableSchemaScm, "test").Type == schema.TypeString, "field test should be added to the top level")

	CustomizeSchemaPath(testCustomizableSchemaScm, "ptr_item").AddNewField("test", &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	})

	assert.Truef(t, MustSchemaPath(testCustomizableSchemaScm, "ptr_item", "test").Type == schema.TypeString, "field ptr_item.test should be added")
}
