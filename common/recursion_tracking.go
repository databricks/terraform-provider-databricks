package common

import (
	"maps"
	"reflect"
	"strings"
)

type recursionTrackingContext struct {
	// Path is used for refernces from resourceProviderRegistry as prefix
	path             []string
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
	newPath := make([]string, len(rt.path))
	copy(newPath, rt.path)
	newTimesVisited := map[string]int{}
	maps.Copy(newTimesVisited, rt.timesVisited)
	return recursionTrackingContext{
		path:             newPath,
		timesVisited:     newTimesVisited,
		maxDepthForTypes: rt.maxDepthForTypes,
	}
}

func (rt recursionTrackingContext) visit(v reflect.Value) {
	rt.timesVisited[getNameForType(v.Type())] += 1
}

func (rt recursionTrackingContext) addToPath(fieldName string) recursionTrackingContext {
	newRt := rt.copy()
	// Special path element `"0"` is used to denote either arrays or sets of elements
	newRt.path = append(rt.path, fieldName, "0")
	return newRt
}

func getEmptyRecursionTrackingContext() recursionTrackingContext {
	return recursionTrackingContext{
		[]string{},
		map[string]int{},
		map[string]int{},
	}
}

func getRecursionTrackingContext(rp RecursiveResourceProvider) recursionTrackingContext {
	return recursionTrackingContext{
		[]string{},
		map[string]int{},
		rp.MaxDepthForTypes(),
	}
}

func getNameForType(t reflect.Type) string {
	return strings.TrimPrefix(t.String(), "*")
}
