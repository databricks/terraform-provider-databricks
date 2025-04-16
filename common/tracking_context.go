// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"maps"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type trackingContext struct {
	timesVisited     map[string]int
	maxDepthForTypes map[string]int
	pathCtx          schemaPathContext
}

type schemaPathContext struct {
	// Path is used for refernces from resourceProviderRegistry as prefix, follows the format of `[a,0,b]`
	path       []string
	schemaPath []*schema.Schema
}

func (tc trackingContext) withPath(fieldName string, schema *schema.Schema) trackingContext {
	newTc := tc.copy()
	// Special path element `"0"` is used to denote either arrays or sets of elements
	newTc.pathCtx.path = append(tc.pathCtx.path, fieldName, "0")
	newTc.pathCtx.schemaPath = append(tc.pathCtx.schemaPath, schema)
	return newTc
}

func (tc trackingContext) depthExceeded(typeField reflect.StructField) bool {
	typeName := getNameForType(typeField.Type)
	if maxDepth, ok := tc.maxDepthForTypes[typeName]; ok {
		return tc.timesVisited[typeName]+1 > maxDepth
	}
	return false
}

func (tc trackingContext) getMaxDepthForTypeField(typeField reflect.StructField) int {
	typeName := getNameForType(typeField.Type)
	return tc.maxDepthForTypes[typeName]
}

func (tc trackingContext) copy() trackingContext {
	newTimesVisited := map[string]int{}
	maps.Copy(newTimesVisited, tc.timesVisited)
	newPath := make([]string, len(tc.pathCtx.path))
	copy(newPath, tc.pathCtx.path)
	newSchemaPath := make([]*schema.Schema, len(tc.pathCtx.schemaPath))
	copy(newSchemaPath, tc.pathCtx.schemaPath)
	return trackingContext{
		timesVisited:     newTimesVisited,
		maxDepthForTypes: tc.maxDepthForTypes,
		pathCtx: schemaPathContext{
			path:       newPath,
			schemaPath: newSchemaPath,
		},
	}
}

func (tc trackingContext) withPathContext(scp schemaPathContext) trackingContext {
	newTc := tc.copy()
	newTc.pathCtx = scp
	return newTc
}

func (tc trackingContext) visit(v reflect.Value) trackingContext {
	newTc := tc.copy()
	newTc.timesVisited[getNameForType(v.Type())] += 1
	return newTc
}

func getEmptySchemaPathContext() schemaPathContext {
	return schemaPathContext{
		[]string{},
		[]*schema.Schema{},
	}
}

func getEmptyTrackingContext() trackingContext {
	return trackingContext{
		map[string]int{},
		map[string]int{},
		getEmptySchemaPathContext(),
	}
}

func getTrackingContext(rp RecursiveResourceProvider) trackingContext {
	return trackingContext{
		map[string]int{},
		rp.MaxDepthForTypes(),
		getEmptySchemaPathContext(),
	}
}

func getNameForType(t reflect.Type) string {
	return strings.TrimPrefix(t.String(), "*")
}
