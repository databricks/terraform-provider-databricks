package common

import (
	"maps"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type recursionTrackingContext struct {
	timesVisited     map[string]int
	maxDepthForTypes map[string]int
}

func (rt recursionTrackingContext) depthExceeded(typeField reflect.StructField) bool {
	typeName := getNameForType(typeField.Type)
	if maxDepth, ok := rt.maxDepthForTypes[typeName]; ok {
		return rt.timesVisited[typeName]+1 > maxDepth
	}
	return false
}

func (rt recursionTrackingContext) getMaxDepthForTypeField(typeField reflect.StructField) int {
	typeName := getNameForType(typeField.Type)
	return rt.maxDepthForTypes[typeName]
}

func (rt recursionTrackingContext) copy() recursionTrackingContext {
	newTimesVisited := map[string]int{}
	maps.Copy(newTimesVisited, rt.timesVisited)
	return recursionTrackingContext{
		timesVisited:     newTimesVisited,
		maxDepthForTypes: rt.maxDepthForTypes,
	}
}

func (rt recursionTrackingContext) visit(v reflect.Value) {
	rt.timesVisited[getNameForType(v.Type())] += 1
}

type SchemaPathContext struct {
	// Path is used for refernces from resourceProviderRegistry as prefix
	path       []string
	schemaPath []*schema.Schema
}

func (spc SchemaPathContext) copy() SchemaPathContext {
	newPath := make([]string, len(spc.path))
	copy(newPath, spc.path)
	newSchemaPath := make([]*schema.Schema, len(spc.schemaPath))
	copy(newSchemaPath, spc.schemaPath)
	return SchemaPathContext{
		path:       newPath,
		schemaPath: newSchemaPath,
	}
}

func (spc SchemaPathContext) addToPath(fieldName string, schema *schema.Schema) SchemaPathContext {
	newSpc := spc.copy()
	// Special path element `"0"` is used to denote either arrays or sets of elements
	newSpc.path = append(spc.path, fieldName, "0")
	newSpc.schemaPath = append(spc.schemaPath, schema)
	return newSpc
}

func getEmptySchemaPathContext() SchemaPathContext {
	return SchemaPathContext{
		[]string{},
		[]*schema.Schema{},
	}
}

func getEmptyRecursionTrackingContext() recursionTrackingContext {
	return recursionTrackingContext{
		map[string]int{},
		map[string]int{},
	}
}

func getRecursionTrackingContext(rp RecursiveResourceProvider) recursionTrackingContext {
	return recursionTrackingContext{
		map[string]int{},
		rp.MaxDepthForTypes(),
	}
}

func getNameForType(t reflect.Type) string {
	return strings.TrimPrefix(t.String(), "*")
}
