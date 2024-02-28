package common

import (
	"maps"
	"reflect"
	"strings"
)

type recursionTrackingContext struct {
	timesVisited     map[string]int
	maxDepthForTypes map[string]int
}

func (rt recursionTrackingContext) depthExceeded(typeField reflect.StructField) bool {
	typeName := rt.getNameForTypeField(typeField)
	if maxDepth, ok := rt.maxDepthForTypes[typeName]; ok {
		return rt.timesVisited[typeName]+1 > maxDepth
	}
	return false
}

func (rt recursionTrackingContext) getNameForTypeField(typeField reflect.StructField) string {
	return strings.TrimPrefix(typeField.Type.String(), "*")
}

func (rt recursionTrackingContext) getMaxDepthForTypeField(typeField reflect.StructField) int {
	typeName := rt.getNameForTypeField(typeField)
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
	rt.timesVisited[strings.TrimPrefix(v.Type().String(), "*")] += 1
}

func getEmptyRecursionTrackingContext() recursionTrackingContext {
	return recursionTrackingContext{
		map[string]int{},
		map[string]int{},
	}
}

func getRecursionTrackingContext(rp ResourceProvider) recursionTrackingContext {
	return recursionTrackingContext{
		map[string]int{},
		rp.MaxDepthForTypes(),
	}
}
