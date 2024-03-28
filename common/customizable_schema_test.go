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
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetConflictsWith([]string{}, []string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["non_optional"].ConflictsWith) == 1, "conflictsWith should be set in field: non_optional")
}

func TestCustomizableSchemaSetExactlyOneOf(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetExactlyOneOf([]string{}, []string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["non_optional"].ExactlyOneOf) == 1, "ExactlyOneOf should be set in field: non_optional")
}

func TestCustomizableSchemaAtLeastOneOf(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetAtLeastOneOf([]string{}, []string{"abc"})
	assert.Truef(t, len(testCustomizableSchemaScm["non_optional"].AtLeastOneOf) == 1, "AtLeastOneOf should be set in field: non_optional")
}

func TestCustomizableSchemaSetRequiredWith(t *testing.T) {
	CustomizeSchemaPath(testCustomizableSchemaScm, "non_optional").SetRequiredWith([]string{}, []string{"abc"})
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
