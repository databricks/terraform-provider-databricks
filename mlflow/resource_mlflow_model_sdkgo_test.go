package mlflow

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mlflowModelSdkGoEnv is the variable that opts resources in to the modular SDK.
const mlflowModelSdkGoEnv = "USE_SDK_GO_RESOURCES"

// mlflowModelOnBothPaths runs test once on the classic SDK path of
// databricks_mlflow_model and once on its modular SDK path. The classic run clears
// the variable, so that the environment of the test process can't switch it.
func mlflowModelOnBothPaths(t *testing.T, test func(t *testing.T)) {
	t.Run("classic", func(t *testing.T) {
		t.Setenv(mlflowModelSdkGoEnv, "")
		test(t)
	})
	t.Run("sdkgo", func(t *testing.T) {
		t.Setenv(mlflowModelSdkGoEnv, mlflowModelResourceType)
		test(t)
	})
}

// TestResourceMlflowModelSdkGoSchemaEqualsClassic compares the schema that the
// resource has on the modular SDK path with the classic one, in every property of
// every attribute at every depth, so that opting in can't change the schema.
func TestResourceMlflowModelSdkGoSchemaEqualsClassic(t *testing.T) {
	t.Setenv(mlflowModelSdkGoEnv, "")
	classic := ResourceMlflowModel().Schema
	t.Setenv(mlflowModelSdkGoEnv, mlflowModelResourceType)
	sdkGoSchemas := map[string]func() map[string]*schema.Schema{
		"schema_of_the_opted_in_resource": func() map[string]*schema.Schema { return ResourceMlflowModel().Schema },
		"schema_from_the_modular_request": mlflowModelSdkGoSchema,
	}
	for name, build := range sdkGoSchemas {
		t.Run(name, func(t *testing.T) {
			assert.Empty(t, mlflowModelSchemaDiff("", reflect.ValueOf(classic), reflect.ValueOf(build())))
		})
	}

	// The comparison has to notice a change of any property, also in a nested block.
	// Every field of schema.Schema is changed in turn on tags.key.
	nestedKey := func(s map[string]*schema.Schema) *schema.Schema {
		return s["tags"].Elem.(*schema.Resource).Schema["key"]
	}
	properties := reflect.TypeFor[schema.Schema]()
	for i := range properties.NumField() {
		property := properties.Field(i).Name
		t.Run("detects_a_changed_"+property, func(t *testing.T) {
			changed := mlflowModelSdkGoSchema()
			mlflowModelSchemaChange(t, reflect.ValueOf(nestedKey(changed)).Elem().Field(i))
			diff := mlflowModelSchemaDiff("", reflect.ValueOf(classic), reflect.ValueOf(changed))
			require.Len(t, diff, 1)
			assert.Contains(t, diff[0], ".tags.Elem.Schema.key."+property+":")
		})
	}
	t.Run("detects_an_attribute_on_one_side_only", func(t *testing.T) {
		changed := mlflowModelSdkGoSchema()
		delete(changed["tags"].Elem.(*schema.Resource).Schema, "value")
		changed["extra"] = &schema.Schema{Type: schema.TypeString, Optional: true}
		diff := mlflowModelSchemaDiff("", reflect.ValueOf(classic), reflect.ValueOf(changed))
		assert.ElementsMatch(t, []string{
			".tags.Elem.Schema.value: only in the classic schema",
			".extra: only in the modular schema",
		}, diff)
	})
}

// mlflowModelSchemaDiff returns one line for each place where classic and modular
// differ. It walks every field of the values by reflection, so it compares the
// properties that schema.Schema has today and those that a later version adds.
// Functions are equal when they are the same function or both nil.
func mlflowModelSchemaDiff(path string, classic, modular reflect.Value) []string {
	if classic.Type() != modular.Type() {
		return []string{path + ": " + classic.Type().String() + " in the classic schema, " +
			modular.Type().String() + " in the modular schema"}
	}
	var diff []string
	switch classic.Kind() {
	case reflect.Interface, reflect.Pointer:
		if classic.IsNil() != modular.IsNil() {
			return []string{path + ": nil in one schema only"}
		}
		if !classic.IsNil() {
			diff = mlflowModelSchemaDiff(path, classic.Elem(), modular.Elem())
		}
	case reflect.Struct:
		for i := range classic.NumField() {
			field := path + "." + classic.Type().Field(i).Name
			diff = append(diff, mlflowModelSchemaDiff(field, classic.Field(i), modular.Field(i))...)
		}
	case reflect.Map:
		for _, key := range classic.MapKeys() {
			entry := path + "." + key.String()
			if !modular.MapIndex(key).IsValid() {
				diff = append(diff, entry+": only in the classic schema")
				continue
			}
			diff = append(diff, mlflowModelSchemaDiff(entry, classic.MapIndex(key), modular.MapIndex(key))...)
		}
		for _, key := range modular.MapKeys() {
			if !classic.MapIndex(key).IsValid() {
				diff = append(diff, path+"."+key.String()+": only in the modular schema")
			}
		}
	case reflect.Slice:
		if classic.Len() != modular.Len() {
			return []string{path + ": the lengths differ"}
		}
		for i := range classic.Len() {
			diff = append(diff, mlflowModelSchemaDiff(path, classic.Index(i), modular.Index(i))...)
		}
	case reflect.Func:
		if classic.Pointer() != modular.Pointer() {
			diff = []string{path + ": the functions differ"}
		}
	default:
		if !classic.Equal(modular) {
			diff = []string{path + ": the values differ"}
		}
	}
	return diff
}

// mlflowModelSchemaChange gives property, a field of a schema.Schema, another value.
func mlflowModelSchemaChange(t *testing.T, property reflect.Value) {
	switch property.Kind() {
	case reflect.Bool:
		property.SetBool(!property.Bool())
	case reflect.Int:
		property.SetInt(property.Int() + 1)
	case reflect.String:
		property.SetString(property.String() + "changed")
	case reflect.Slice:
		property.Set(reflect.Append(property, reflect.Zero(property.Type().Elem())))
	case reflect.Interface:
		property.Set(reflect.ValueOf("changed"))
	case reflect.Func:
		property.Set(reflect.MakeFunc(property.Type(), func([]reflect.Value) []reflect.Value {
			panic("the comparison doesn't call the functions of a schema")
		}))
	default:
		t.Fatalf("schema.Schema has a property of kind %s, which this test can't change", property.Kind())
	}
}
